package aggregator

import (
	"context"
	"errors"
	"fmt"
	"math/big"
	"sort"
	"sync"
	"time"

	"github.com/Layr-Labs/eigensdk-go/crypto/bls"
	"github.com/Layr-Labs/eigensdk-go/logging"
	"github.com/Layr-Labs/eigensdk-go/services/avsregistry"
	"github.com/Layr-Labs/eigensdk-go/types"
	"github.com/automata-network/multi-prover-avs/utils"
	"github.com/chzyer/logex"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
)

var (
	TaskAlreadyInitializedErrorFn = func(taskIndex types.TaskIndex) error {
		return fmt.Errorf("task %d already initialized", taskIndex)
	}
	TaskExpiredError = func(taskIndex types.TaskIndex, threadholdInfo map[types.QuorumNum]ThresholdInfo) error {
		return fmt.Errorf("task#%v expired: %+v", taskIndex, threadholdInfo)
	}
	TaskNotFoundErrorFn = func(taskIndex types.TaskIndex) error {
		return fmt.Errorf("task %d not initialized or already completed", taskIndex)
	}
	OperatorNotPartOfTaskQuorumErrorFn = func(operatorId types.OperatorId) error {
		return fmt.Errorf("operator %x not part of task's quorum", operatorId)
	}
	SignatureVerificationError = func(operatorId types.OperatorId, err error) error {
		return fmt.Errorf("operator %x Failed to verify signature: %w", operatorId, err)
	}
	OperatorG2KeyNotFound = func(operatorId types.OperatorId) error {
		return fmt.Errorf("operator %x g2 key not found", operatorId)
	}
	IncorrectSignatureError = errors.New("Signature verification failed. Incorrect Signature.")
)

// BlsAggregationServiceResponse is the response from the bls aggregation service
type BlsAggregationServiceResponse struct {
	Err                error                    // if Err is not nil, the other fields are not valid
	TaskIndex          types.TaskIndex          // unique identifier of the task
	TaskResponseDigest types.TaskResponseDigest // digest of the task response that was signed
	// The below 8 fields are the data needed to build the IBLSSignatureChecker.NonSignerStakesAndSignature struct
	// users of this service will need to build the struct themselves by converting the bls points
	// into the BN254.G1/G2Point structs that the IBLSSignatureChecker expects
	// given that those are different for each AVS service manager that individually inherits BLSSignatureChecker
	NonSignersPubkeysG1          []*bls.G1Point
	QuorumApksG1                 []*bls.G1Point
	SignersApkG2                 *bls.G2Point
	SignersAggSigG1              *bls.Signature
	NonSignerQuorumBitmapIndices []uint32
	QuorumApkIndices             []uint32
	TotalStakeIndices            []uint32
	NonSignerStakeIndices        [][]uint32
}

// BlsAggregatorService is a service that performs BLS signature aggregation for an AVS' tasks
// Assumptions:
//  1. BlsAggregatorService only verifies digest signatures, so avs code needs to verify that the digest
//     passed to ProcessNewSignature is indeed the digest of a valid taskResponse
//     (see the comment above checkSignature for more details)
//  2. BlsAggregatorService is VERY generic and makes very few assumptions about the tasks structure or
//     the time at which operators will send their signatures. It is mostly suitable for offchain computation
//     oracle (a la truebit) type of AVS, where tasks are sent onchain by users sporadically, and where
//     new tasks can start even before the previous ones have finished aggregation.
//     AVSs like eigenDA that have a much more controlled task submission schedule and where new tasks are
//     only submitted after the previous one's response has been aggregated and responded onchain, could have
//     a much simpler AggregationService without all the complicated parallel goroutines.
type BlsAggregatorService struct {
	// aggregatedResponsesC is the channel which all goroutines share to send their responses back to the
	// main thread after they are done aggregating (either they reached the threshold, or timeout expired)
	aggregatedResponsesC chan *BlsAggregationServiceResponse
	// signedTaskRespsCs are the channels to send the signed task responses to the goroutines processing them
	// each new task is assigned a new goroutine and a new channel
	signedTaskRespsCs map[types.TaskIndex]chan types.SignedTaskResponseDigest
	// we add chans to taskChans from the main thread (InitializeNewTask) when we create new tasks,
	// we read them in ProcessNewSignature from the main thread when we receive new signed tasks,
	// and remove them from its respective goroutine when the task is completed or reached timeout
	// we thus need a mutex to protect taskChans
	taskChansMutex     sync.RWMutex
	avsRegistryService avsregistry.AvsRegistryService
	logger             logging.Logger
}

func NewBlsAggregatorService(avsRegistryService avsregistry.AvsRegistryService, logger logging.Logger) *BlsAggregatorService {
	return &BlsAggregatorService{
		aggregatedResponsesC: make(chan *BlsAggregationServiceResponse, 4),
		signedTaskRespsCs:    make(map[types.TaskIndex]chan types.SignedTaskResponseDigest),
		taskChansMutex:       sync.RWMutex{},
		avsRegistryService:   avsRegistryService,
		logger:               logger,
	}
}

func (a *BlsAggregatorService) GetResponseChannel() <-chan *BlsAggregationServiceResponse {
	return a.aggregatedResponsesC
}

// InitializeNewTask creates a new task goroutine meant to process new signed task responses for that task
// (that are sent via ProcessNewSignature) and adds a channel to a.taskChans to send the signed task responses to it
// quorumNumbers and quorumThresholdPercentages set the requirements for this task to be considered complete, which happens
// when a particular TaskResponseDigest (received via the a.taskChans[taskIndex]) has been signed by signers whose stake
// in each of the listed quorums adds up to at least quorumThresholdPercentages[i] of the total stake in that quorum
func (a *BlsAggregatorService) InitializeNewTask(
	ctx context.Context,
	taskIndex types.TaskIndex,
	taskCreatedBlock uint32,
	quorumNumbers types.QuorumNums,
	minQuorumThresholdPercentages types.QuorumThresholdPercentages,
	minWait time.Duration,
	timeToExpiry time.Duration,
) error {
	a.logger.Info("AggregatorService initializing new task", "taskIndex", taskIndex, "taskCreatedBlock", taskCreatedBlock, "quorumNumbers", quorumNumbers, "minQuorumThresholdPercentages", minQuorumThresholdPercentages, "minWait", minWait, "timeToExpiry", timeToExpiry)

	a.taskChansMutex.Lock()
	signedTaskRespsC, taskExists := a.signedTaskRespsCs[taskIndex]
	if !taskExists {
		signedTaskRespsC = make(chan types.SignedTaskResponseDigest, 128)
		a.signedTaskRespsCs[taskIndex] = signedTaskRespsC
	}
	a.taskChansMutex.Unlock()
	if taskExists {
		return TaskAlreadyInitializedErrorFn(taskIndex)
	}

	operatorStates, err := NewOperatorStates(ctx, a.avsRegistryService, quorumNumbers, minQuorumThresholdPercentages, taskIndex, taskCreatedBlock)
	if err != nil {
		return logex.Trace(err)
	}

	go a.singleTaskAggregatorGoroutineFunc(operatorStates, minWait, timeToExpiry, signedTaskRespsC)
	return nil
}

func (a *BlsAggregatorService) ProcessNewSignature(
	ctx context.Context,
	taskIndex types.TaskIndex,
	taskResponseDigest types.TaskResponseDigest,
	blsSignature *bls.Signature,
	operatorId types.OperatorId,
) error {
	start := time.Now()
	defer func() {
		a.logger.Info("AggregatorService process new signature", "taskIndex", taskIndex, "costTime", time.Since(start))
	}()

	a.taskChansMutex.Lock()
	taskC, taskInitialized := a.signedTaskRespsCs[taskIndex]
	a.taskChansMutex.Unlock()
	if !taskInitialized {
		return TaskNotFoundErrorFn(taskIndex)
	}
	signatureVerificationErrorC := make(chan error)
	// send the task to the goroutine processing this task
	// and return the error (if any) returned by the signature verification routine
	select {
	// we need to send this as part of select because if the goroutine is processing another SignedTaskResponseDigest
	// and cannot receive this one, we want the context to be able to cancel the request
	case taskC <- types.SignedTaskResponseDigest{
		TaskResponseDigest:          taskResponseDigest,
		BlsSignature:                blsSignature,
		OperatorId:                  operatorId,
		SignatureVerificationErrorC: signatureVerificationErrorC,
	}:
		// note that we need to wait synchronously here for this response because we want to
		// send back an informative error message to the operator who sent his signature to the aggregator
		return <-signatureVerificationErrorC
	case <-ctx.Done():
		return ctx.Err()
	}
}

func (a *BlsAggregatorService) singleTaskAggregatorGoroutineFunc(
	operatorStates *OperatorStates,
	minWait time.Duration,
	timeToExpiry time.Duration,
	signedTaskRespsC <-chan types.SignedTaskResponseDigest,
) {
	defer a.closeTaskGoroutine(operatorStates.taskIndex)

	fastSend := false
	if minWait == time.Duration(0) {
		minWait = timeToExpiry
		fastSend = true
	}

	taskExpiredTimer := time.NewTimer(timeToExpiry)
	defer taskExpiredTimer.Stop()
	fireTimer := time.NewTimer(minWait)
	defer fireTimer.Stop()

	task := newBlsAggTask(operatorStates)

	for {
		select {
		case signer := <-signedTaskRespsC:
			a.logger.Info("Task goroutine received new signed task response digest", "taskIndex", operatorStates.taskIndex, "signedTaskResponseDigest", signer)
			if !task.AddNewSigner(signer) {
				a.logger.Info("Signature Verification Failed", "taskIndex", operatorStates.taskIndex, "signedTaskResponseDigest", signer)
				continue
			}

			a.logger.Infof("shares update for taskIndex=%v: digest=%x, %+v", operatorStates.taskIndex, task.GetHighestDigest(), task.getThreshold(signer.TaskResponseDigest))

			if fastSend {
				result := task.Result(a.avsRegistryService)
				if result != nil {
					a.aggregatedResponsesC <- result
					return
				}
			}
		case <-fireTimer.C:
			result := task.Result(a.avsRegistryService)
			if result == nil {
				// We could not make it in time, in this case, so we should send the result ASAP
				fastSend = true
				continue
			}
			a.aggregatedResponsesC <- result
			return
		case <-taskExpiredTimer.C:
			a.aggregatedResponsesC <- task.Err(fmt.Errorf("task expired"))
			return
		}
	}

}

// closeTaskGoroutine is run when the goroutine processing taskIndex's task responses ends (for whatever reason)
// it deletes the response channel for taskIndex from a.taskChans
// so that the main thread knows that this task goroutine is no longer running
// and doesn't try to send new signatures to it
func (a *BlsAggregatorService) closeTaskGoroutine(taskIndex types.TaskIndex) {
	a.logger.Infof("close task: %v", taskIndex)
	a.taskChansMutex.Lock()
	delete(a.signedTaskRespsCs, taskIndex)
	a.taskChansMutex.Unlock()
}

// verifySignature verifies that a signature is valid against the operator pubkey stored in the
// operatorsAvsStateDict for that particular task
// TODO(samlaf): right now we are only checking that the *digest* is signed correctly!!
// we could be sent a signature of any kind of garbage and we would happily aggregate it
// this forces the avs code to verify that the digest is indeed the digest of a valid taskResponse
// we could take taskResponse as an interface{} and have avs code pass us a taskResponseHashFunction
// that we could use to hash and verify the taskResponse itself
func (o *OperatorStates) verifySignature(
	signedTaskResponseDigest types.SignedTaskResponseDigest,
) error {
	operatorId := signedTaskResponseDigest.OperatorId
	avsState, ok := o.operatorsAvsStateDict[operatorId]
	if !ok {
		return logex.Trace(OperatorNotPartOfTaskQuorumErrorFn(operatorId))
	}

	// 0. verify that the msg actually came from the correct operator
	operatorG2Pubkey := avsState.Pubkeys.G2Pubkey
	if operatorG2Pubkey == nil {
		return logex.Trace(OperatorG2KeyNotFound(operatorId))
	}
	logex.Debug("Verifying signed task response digest signature",
		"operatorG2Pubkey", operatorG2Pubkey,
		"taskResponseDigest", signedTaskResponseDigest.TaskResponseDigest,
		"blsSignature", signedTaskResponseDigest.BlsSignature,
	)
	signatureVerified, err := signedTaskResponseDigest.BlsSignature.Verify(operatorG2Pubkey, signedTaskResponseDigest.TaskResponseDigest)
	if err != nil {
		return logex.Trace(err, operatorId)
	}
	if !signatureVerified {
		return logex.Trace(IncorrectSignatureError, operatorId)
	}
	return nil
}

type ThresholdInfo struct {
	Signed    *big.Int
	Total     *big.Int
	Threshold types.QuorumThresholdPercentage
	Percent   *big.Float
}

// aggregatedOperators is meant to be used as a value in a map
// map[taskResponseDigest]aggregatedOperators
type aggregatedOperators struct {
	// aggregate g2 pubkey of all operatos who signed on this taskResponseDigest
	signersApkG2 *bls.G2Point
	// aggregate signature of all operators who signed on this taskResponseDigest
	signersAggSigG1 *bls.Signature
	// aggregate stake of all operators who signed on this header for each quorum
	signersTotalStakePerQuorum map[types.QuorumNum]*big.Int
	// set of OperatorId of operators who signed on this header
	signersOperatorIdsSet map[types.OperatorId]bool
}

func newAggregatedOperators() *aggregatedOperators {
	return &aggregatedOperators{
		// we've already verified that the operator is part of the task's quorum, so we don't need checks here
		signersApkG2:               bls.NewZeroG2Point(),
		signersAggSigG1:            bls.NewZeroSignature(),
		signersOperatorIdsSet:      make(map[types.OperatorId]bool),
		signersTotalStakePerQuorum: make(map[types.QuorumNum]*big.Int),
	}
}

func (a *aggregatedOperators) Add(sig types.SignedTaskResponseDigest, states types.OperatorAvsState) {
	a.signersAggSigG1.Add(sig.BlsSignature)
	a.signersApkG2.Add(states.Pubkeys.G2Pubkey)
	a.signersOperatorIdsSet[sig.OperatorId] = true
	for quorumNum, stake := range states.StakePerQuorum {
		oldStake, ok := a.signersTotalStakePerQuorum[quorumNum]
		if !ok {
			oldStake = big.NewInt(0)
			a.signersTotalStakePerQuorum[quorumNum] = oldStake
		}
		oldStake.Add(oldStake, stake)
	}
}

type blsAggTask struct {
	states                  *OperatorStates
	aggregatedOperatorsDict map[types.TaskResponseDigest]*aggregatedOperators
}

func newBlsAggTask(operatorStates *OperatorStates) *blsAggTask {
	return &blsAggTask{
		states:                  operatorStates,
		aggregatedOperatorsDict: make(map[types.TaskResponseDigest]*aggregatedOperators),
	}
}

func (t *blsAggTask) getThreshold(digest types.TaskResponseDigest) map[types.QuorumNum]ThresholdInfo {
	return t.states.formatThreshold(t.getDigestOperators(digest))
}

func (t *blsAggTask) getDigestOperators(digest types.TaskResponseDigest) *aggregatedOperators {
	digestAggregatedOperators, ok := t.aggregatedOperatorsDict[digest]
	if !ok {
		digestAggregatedOperators = newAggregatedOperators()
		t.aggregatedOperatorsDict[digest] = digestAggregatedOperators
	}
	return digestAggregatedOperators
}

func (t *blsAggTask) GetHighestDigest() types.TaskResponseDigest {
	var highestDigest *types.TaskResponseDigest
	highestPct := big.NewFloat(0)
	for digest, oprs := range t.aggregatedOperatorsDict {
		thresholds := t.states.formatThreshold(oprs)
		minPct := big.NewFloat(1)
		for _, n := range thresholds {
			if n.Percent.Cmp(minPct) < 0 {
				minPct = n.Percent
			}
		}
		if minPct.Cmp(highestPct) > 0 {
			highestPct = minPct
			highestDigest = &digest
		}
	}
	return *highestDigest
}

func (t *blsAggTask) Err(err error) *BlsAggregationServiceResponse {
	digest := t.GetHighestDigest()
	oprs := t.aggregatedOperatorsDict[digest]
	threshold := t.states.formatThreshold(oprs)

	return &BlsAggregationServiceResponse{
		Err:                logex.Trace(err, fmt.Sprintf("threshold:%v", threshold), fmt.Sprintf("taskIndex:%v", t.states.taskIndex)),
		TaskIndex:          t.states.taskIndex,
		TaskResponseDigest: digest,
	}
}

func (t *blsAggTask) Result(avsRegistryService avsregistry.AvsRegistryService) *BlsAggregationServiceResponse {
	digest := t.GetHighestDigest()
	oprs := t.aggregatedOperatorsDict[digest]

	thresholds := t.states.formatThreshold(oprs)
	if !t.states.checkIfStakeThresholdsMet(thresholds) {
		return nil
	}

	nonSignersOperatorIds := t.states.NonSignerOperatorIds(oprs)
	quorumApksG1 := t.states.QuorumApksG1()

	nonSignersG1Pubkeys := make([]*bls.G1Point, 0, len(nonSignersOperatorIds))
	for _, operatorId := range nonSignersOperatorIds {
		operator := t.states.operatorsAvsStateDict[operatorId]
		nonSignersG1Pubkeys = append(nonSignersG1Pubkeys, operator.Pubkeys.G1Pubkey)
	}

	response := &BlsAggregationServiceResponse{
		Err:                 nil,
		TaskIndex:           t.states.taskIndex,
		TaskResponseDigest:  digest,
		NonSignersPubkeysG1: nonSignersG1Pubkeys,
		QuorumApksG1:        quorumApksG1,
		SignersApkG2:        oprs.signersApkG2,
		SignersAggSigG1:     oprs.signersAggSigG1,
	}

	indices, err := avsRegistryService.GetCheckSignaturesIndices(&bind.CallOpts{}, t.states.taskCreatedBlock, t.states.quorumNumbers, nonSignersOperatorIds)
	if err != nil {
		response.Err = types.WrapError(errors.New("failed to get check signatures indices"), err)
		return response
	}
	response.NonSignerQuorumBitmapIndices = indices.NonSignerQuorumBitmapIndices
	response.QuorumApkIndices = indices.QuorumApkIndices
	response.TotalStakeIndices = indices.TotalStakeIndices
	response.NonSignerStakeIndices = indices.NonSignerStakeIndices
	return response
}

func (t *blsAggTask) AddNewSigner(signedTaskResponseDigest types.SignedTaskResponseDigest) bool {
	operatorId := signedTaskResponseDigest.OperatorId
	if err := t.states.verifySignature(signedTaskResponseDigest); err != nil {
		signedTaskResponseDigest.SignatureVerificationErrorC <- logex.Trace(err, t.states.taskIndex)
		return false
	}
	signedTaskResponseDigest.SignatureVerificationErrorC <- nil

	avsState := t.states.operatorsAvsStateDict[operatorId]

	oprs := t.getDigestOperators(signedTaskResponseDigest.TaskResponseDigest)
	oprs.Add(signedTaskResponseDigest, avsState)

	return true
}

type OperatorStates struct {
	taskIndex                     types.TaskIndex
	quorumNumbers                 types.QuorumNums
	taskCreatedBlock              uint32
	minQuorumThresholdPercentages []types.QuorumThresholdPercentage
	operatorsAvsStateDict         map[types.OperatorId]types.OperatorAvsState
	quorumsAvsStakeDict           map[types.QuorumNum]types.QuorumAvsState
}

func NewOperatorStates(
	ctx context.Context,
	avsRegistryService avsregistry.AvsRegistryService,
	quorumNumbers types.QuorumNums,
	minQuorumThresholdPercentages []types.QuorumThresholdPercentage,
	taskIndex uint32,
	taskCreatedBlock uint32,
) (*OperatorStates, error) {
	operatorsAvsStateDict, err := utils.Retry(5, time.Second, func() (map[types.OperatorId]types.OperatorAvsState, error) {
		return avsRegistryService.GetOperatorsAvsStateAtBlock(ctx, quorumNumbers, taskCreatedBlock)
	}, "taskIndex", taskIndex, "taskCreatedBlock", taskCreatedBlock)
	if err != nil {
		return nil, logex.Trace(err,
			fmt.Sprintf("taskIndex:%v", taskIndex),
			fmt.Sprintf("taskCreatedBlock:%v", taskCreatedBlock),
		)
	}
	quorumsAvsStakeDict, err := utils.Retry(5, time.Second, func() (map[types.QuorumNum]types.QuorumAvsState, error) {
		return avsRegistryService.GetQuorumsAvsStateAtBlock(ctx, quorumNumbers, taskCreatedBlock)
	}, "taskIndex", taskIndex, "taskCreatedBlock", taskCreatedBlock)
	if err != nil {
		return nil, logex.Trace(err,
			fmt.Sprintf("taskIndex:%v", taskIndex),
			fmt.Sprintf("taskCreatedBlock:%v", taskCreatedBlock),
		)
	}

	return &OperatorStates{
		taskIndex,
		quorumNumbers,
		taskCreatedBlock,
		minQuorumThresholdPercentages,
		operatorsAvsStateDict,
		quorumsAvsStakeDict,
	}, nil
}

func (o *OperatorStates) QuorumApksG1() []*bls.G1Point {
	quorumApksG1 := make([]*bls.G1Point, 0, len(o.quorumNumbers))
	for _, quorumNumber := range o.quorumNumbers {
		quorumApksG1 = append(quorumApksG1, o.quorumsAvsStakeDict[quorumNumber].AggPubkeyG1)
	}
	return quorumApksG1
}

func (o *OperatorStates) NonSignerOperatorIds(oprs *aggregatedOperators) []types.OperatorId {
	nonSignersOperatorIds := make([]types.OperatorId, 0, len(o.operatorsAvsStateDict))
	for operatorId := range o.operatorsAvsStateDict {
		if _, ok := oprs.signersOperatorIdsSet[operatorId]; !ok {
			nonSignersOperatorIds = append(nonSignersOperatorIds, operatorId)
		}
	}

	sort.SliceStable(nonSignersOperatorIds, func(i, j int) bool {
		iOprInt := new(big.Int).SetBytes(nonSignersOperatorIds[i][:])
		jOprInt := new(big.Int).SetBytes(nonSignersOperatorIds[j][:])
		return iOprInt.Cmp(jOprInt) == -1
	})
	return nonSignersOperatorIds
}

func (o *OperatorStates) formatThreshold(oprs *aggregatedOperators) map[types.QuorumNum]ThresholdInfo {
	out := make(map[types.QuorumNum]ThresholdInfo)
	for idx, quorumNum := range o.quorumNumbers {
		info := ThresholdInfo{
			Signed:    oprs.signersTotalStakePerQuorum[quorumNum],
			Total:     o.quorumsAvsStakeDict[quorumNum].TotalStake,
			Threshold: o.minQuorumThresholdPercentages[idx],
		}
		signed := new(big.Float).SetInt(info.Signed)
		total := new(big.Float).SetInt(info.Total)
		info.Percent = new(big.Float).Quo(signed, total)
		out[quorumNum] = info
	}
	return out
}

func (o *OperatorStates) checkIfStakeThresholdsMet(info map[types.QuorumNum]ThresholdInfo) bool {
	for _, info := range info {
		signedStake := new(big.Int).Mul(info.Signed, big.NewInt(100))
		thresholdStake := new(big.Int).Mul(info.Total, big.NewInt(int64(info.Threshold)))
		if signedStake.Cmp(thresholdStake) < 0 {
			return false
		}
	}
	return true
}
