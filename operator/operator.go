package operator

import (
	"bytes"
	"context"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"math/big"
	"os"
	"strings"
	"sync"
	"time"

	"github.com/Layr-Labs/eigensdk-go/nodeapi"

	"github.com/automata-network/multi-prover-avs/aggregator"
	"github.com/automata-network/multi-prover-avs/contracts/bindings"
	"github.com/automata-network/multi-prover-avs/contracts/bindings/TEELivenessVerifier"
	"github.com/automata-network/multi-prover-avs/utils"
	"github.com/automata-network/multi-prover-avs/xmetric"
	"github.com/automata-network/multi-prover-avs/xtask"
	"github.com/chzyer/logex"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
)

type Operator struct {
	cfg    *ConfigContext
	logger *utils.Logger

	operatorAddress common.Address
	metricName      string
	semVer          string

	aggregator     *aggregator.Client
	nodeApi        *nodeapi.NodeApi
	operatorMetric *xmetric.OperatorCollector

	operatorId    [32]byte
	metrics       *Metrics
	quorumNumbers []byte

	proverClient *xtask.ProverClient

	TEELivenessVerifier *TEELivenessVerifier.TEELivenessVerifier
}

func NewOperator(path string, semVer string) (*Operator, error) {
	cfg, err := ParseConfigContext(path, nil)
	if err != nil {
		return nil, logex.Trace(err)
	}

	proverClient, err := xtask.NewProverClient(cfg.Config.ProverURL)
	if err != nil {
		return nil, logex.Trace(err)
	}

	logger := utils.NewLogger(logex.NewLoggerEx(os.Stderr))

	TEELivenessVerifier, err := TEELivenessVerifier.NewTEELivenessVerifier(cfg.Config.TEELivenessVerifierAddress, cfg.AttestationClient)
	if err != nil {
		return nil, logex.Trace(err)
	}

	aggClient, err := aggregator.NewClient(cfg.Config.AggregatorURL)
	if err != nil {
		return nil, logex.Trace(err, "aggregatorURL:"+cfg.Config.AggregatorURL)
	}

	operatorAddress, err := cfg.QueryOperatorAddress()
	if err != nil {
		return nil, logex.Trace(err, "queryOperatorAddr")
	}

	if operatorAddress == utils.ZeroAddress {
		return nil, logex.NewErrorf("operator is not registered")
	}

	operatorMetric := xmetric.NewOperatorCollector("avs", cfg.EigenClients.PrometheusRegistry)
	metrics := NewMetrics(cfg.AvsName, cfg.EigenClients, logger, operatorAddress, cfg.Config.EigenMetricsIpPortAddress, xtask.GetQuorumNames())

	nodeApi := nodeapi.NewNodeApi(cfg.AvsName, semVer, cfg.Config.NodeApiIpPortAddress, logger)

	operator := &Operator{
		cfg:                 cfg,
		metricName:          fmt.Sprintf("%v_%v", operatorAddress, semVer),
		semVer:              semVer,
		proverClient:        proverClient,
		logger:              logger,
		aggregator:          aggClient,
		operatorAddress:     operatorAddress,
		metrics:             metrics,
		operatorMetric:      operatorMetric,
		nodeApi:             nodeApi,
		TEELivenessVerifier: TEELivenessVerifier,
	}

	return operator, nil
}

func (o *Operator) Start(ctx context.Context) error {
	logex.Info("starting operator...")
	nodeApiErrChan := o.nodeApi.Start()

	if err := o.checkIsRegistered(); err != nil {
		return logex.Trace(err)
	}
	if err := o.RegisterAttestationReport(ctx); err != nil {
		return logex.Trace(err, utils.EcdsaAddress(o.cfg.AttestationEcdsaKey))
	}
	md, err := o.proverClient.Metadata(ctx)
	if err != nil {
		return logex.Trace(err, fmt.Sprintf("check prover metadata: %v", o.cfg.Config.ProverURL))
	}
	errChan := o.metrics.Serve(o.cfg.Config.EigenMetricsIpPortAddress, o.cfg.Config.ProverURL)
	go func() {
		select {
		case err := <-errChan:
			logex.Fatal(err)
		case err := <-nodeApiErrChan:
			logex.Fatal(err)
		}
	}()

	o.logger.Infof("Started Operator... operator info: operatorId=%v, operatorAddr=%v, operatorG1Pubkey=%v, operatorG2Pubkey=%v, proverVersion=%v, fetchTaskWithContext=%v",
		hex.EncodeToString(o.operatorId[:]),
		o.operatorAddress,
		o.cfg.BlsKey.GetPubKeyG1(),
		o.cfg.BlsKey.GetPubKeyG2(),
		md.Version, md.TaskWithContext,
	)

	go o.metricExporterLoop(ctx)
	go o.metadataExport(ctx)

	var wg sync.WaitGroup
	for _, quorum := range o.quorumNumbers {
		ty := xtask.NewTaskType(quorum)
		if !ty.IsValid() {
			logex.Errorf("unknown quorum: %v", quorum)
			continue
		}
		wg.Add(1)
		go func() {
			defer wg.Done()
			o.subscribeTask(ctx, ty, md.GetWithContext(ty))
		}()

	}
	wg.Wait()

	return nil
}

func (o *Operator) metadataExport(ctx context.Context) {
	ticker := time.NewTicker(60 * time.Second)
	defer ticker.Stop()
	export := func() {
		md, err := o.proverClient.Metadata(ctx)
		if err != nil {
			logex.Error(err)
			return
		}

		operatorAddr := o.operatorAddress
		version := o.semVer
		attestationAddr := utils.EcdsaAddress(o.cfg.AttestationEcdsaKey)
		proverUrlHash := utils.ProverAddrHash(o.cfg.Config.ProverURL)
		proverVersion := md.Version
		proverWithContext := md.WithContext
		o.operatorMetric.Metadata.WithLabelValues(
			o.cfg.AvsName,
			operatorAddr.String(),
			version,
			attestationAddr.String(),
			proverUrlHash.String(),
			proverVersion,
			fmt.Sprint(proverWithContext),
			fmt.Sprint(md.GetWithContext(xtask.ScrollTask)),
			fmt.Sprint(md.GetWithContext(xtask.LineaTask)),
		).Add(1)

	}

	export()
	for range ticker.C {
		export()
	}
}

func (o *Operator) subscribeTask(ctx context.Context, ty xtask.TaskType, withContext bool) {
	req := &aggregator.FetchTaskReq{
		PrevTaskID:  0,
		TaskType:    ty,
		MaxWaitSecs: 100,
		WithContext: withContext,
	}
	for {
		logex.Infof("fetch task[%v]: %#v", ty.Value(), req)
		resp, err := o.aggregator.FetchTask(ctx, req)
		if err != nil {
			time.Sleep(time.Second)
			continue
		}
		if resp.TaskID == 0 {
			time.Sleep(time.Second)
			continue
		}

		logex.Infof("accept new task: [%v] %v", resp.TaskType.Value(), resp.TaskID)
		req.PrevTaskID = resp.TaskID

		o.operatorMetric.FetchTask.WithLabelValues(o.cfg.AvsName, ty.Value(), fmt.Sprint(req.WithContext)).Add(1)
		o.operatorMetric.LatestTask.WithLabelValues(o.cfg.AvsName, ty.Value()).Set(float64(resp.TaskID))

		startProcessTask := time.Now()
		switch resp.TaskType {
		case xtask.ScrollTask:
			if err := o.processScrollTask(ctx, resp); err != nil {
				logex.Error(err)
			}
		case xtask.LineaTask:
			if err := o.processLineaTask(ctx, resp); err != nil {
				logex.Error(err)
			}
		}
		o.operatorMetric.ProcessTaskMs.WithLabelValues(o.cfg.AvsName, ty.Value()).Set(float64(time.Since(startProcessTask).Milliseconds()))
	}
}

func (o *Operator) metricExporterLoop(ctx context.Context) {
	export := func() {
		metrics, err := o.metrics.Gather()
		if err != nil {
			logex.Error(err)
		}
		newMetric := make([]*xmetric.MetricFamily, 0, len(metrics))
		for _, item := range metrics {
			if strings.HasPrefix(*item.Name, "avs") || strings.HasPrefix(*item.Name, "eigen_") {
				newMetric = append(newMetric, item)
			}
		}
		err = o.aggregator.SubmitMetrics(ctx, &aggregator.SubmitMetricsReq{
			Name:    o.metricName,
			Metrics: newMetric,
		})
		if err != nil {
			logex.Error(err)
		}
	}

	export()
	for range time.Tick(600 * time.Second) {
		export()
	}
}

func (o *Operator) processLineaTask(ctx context.Context, resp *aggregator.FetchTaskResp) (err error) {
	ty := xtask.LineaTask
	var ext xtask.LineaTaskExt
	if err := json.Unmarshal(resp.Ext, &ext); err != nil {
		return logex.Trace(err)
	}
	var taskCtx *xtask.ScrollContext
	if len(resp.Context) == 0 {
		logex.Info("[linea] generating task context for:", ext.StartBlock.ToInt(), ext.EndBlock.ToInt())
		var skip bool
		taskCtx, skip, err = o.proverClient.GenerateLineaContext(ctx, ext.StartBlock.ToInt().Int64(), ext.EndBlock.ToInt().Int64(), xtask.LineaTask)
		if err != nil {
			return logex.Trace(err)
		}
		if skip {
			return nil
		}
	} else {
		if err := json.Unmarshal(resp.Context, &taskCtx); err != nil {
			return logex.Trace(err)
		}
	}

	genPoeStart := time.Now()
	poe, skip, err := o.proverClient.ProveLinea(ctx, o.operatorAddress, &ext, taskCtx)
	if err != nil {
		return logex.Trace(err)
	}
	o.operatorMetric.GenPoeMs.WithLabelValues(o.cfg.AvsName, ty.Value()).Set(float64(time.Since(genPoeStart).Milliseconds()))

	logex.Pretty(poe)
	if skip {
		return nil
	}

	md := &aggregator.Metadata{
		BatchId:    poe.BatchId,
		StartBlock: poe.StartBlock,
		EndBlock:   poe.EndBlock,
	}
	mdBytes, err := json.Marshal(md)
	if err != nil {
		return logex.Trace(err)
	}

	stateHeader := &aggregator.StateHeader{
		Identifier:                 (*hexutil.Big)(big.NewInt(int64(ty))),
		Metadata:                   mdBytes,
		State:                      poe.Poe.Pack(),
		QuorumNumbers:              []byte{ty.GetQuorum()},
		QuorumThresholdPercentages: []byte{0},
		ReferenceBlockNumber:       uint32(ext.ReferenceBlockNumber),
	}

	digest, err := stateHeader.Digest()
	if err != nil {
		return logex.Trace(err)
	}
	sig := o.cfg.BlsKey.SignMessage(digest)

	o.operatorMetric.SubmitTask.WithLabelValues(o.cfg.AvsName, ty.Value()).Add(1)
	submitTaskTime := time.Now()
	// submit to aggregator
	if err := o.aggregator.SubmitTask(ctx, &aggregator.TaskRequest{
		Task:            stateHeader,
		Signature:       sig,
		OperatorId:      o.operatorId,
		ProverSignature: poe.PoeSignature,
	}); err != nil {
		return logex.Trace(err)
	}
	o.operatorMetric.SubmitTaskMs.WithLabelValues(o.cfg.AvsName, ty.Value()).Set(float64(time.Since(submitTaskTime).Milliseconds()))
	// logex.Info(poe)
	return nil
}

func (o *Operator) processScrollTask(ctx context.Context, resp *aggregator.FetchTaskResp) (err error) {
	ty := xtask.ScrollTask
	var ext xtask.ScrollTaskExt
	if err := json.Unmarshal(resp.Ext, &ext); err != nil {
		return logex.Trace(err)
	}
	var taskCtx *xtask.ScrollContext
	if len(resp.Context) == 0 {
		logex.Info("[scroll] generating task context for:", ext.StartBlock.ToInt(), ext.EndBlock.ToInt())
		var skip bool
		taskCtx, skip, err = o.proverClient.GenerateScrollContext(ctx, ext.StartBlock.ToInt().Int64(), ext.EndBlock.ToInt().Int64(), ty)
		if err != nil {
			return logex.Trace(err)
		}
		if skip {
			return nil
		}
	} else {
		if err := json.Unmarshal(resp.Context, &taskCtx); err != nil {
			return logex.Trace(err)
		}
	}

	genPoeStart := time.Now()
	poe, skip, err := o.proverClient.GetPoeByPob(ctx, o.operatorAddress, ext.BatchData, taskCtx)
	if err != nil {
		return logex.Trace(err)
	}
	o.operatorMetric.GenPoeMs.WithLabelValues(o.cfg.AvsName, ty.Value()).Set(float64(time.Since(genPoeStart).Milliseconds()))

	logex.Pretty(poe)
	if skip {
		return nil
	}

	md := &aggregator.Metadata{
		BatchId:    poe.BatchId,
		StartBlock: poe.StartBlock,
		EndBlock:   poe.EndBlock,
	}
	mdBytes, err := json.Marshal(md)
	if err != nil {
		return logex.Trace(err)
	}

	stateHeader := &aggregator.StateHeader{
		Identifier:                 (*hexutil.Big)(big.NewInt(int64(ty))),
		Metadata:                   mdBytes,
		State:                      poe.Poe.Pack(),
		QuorumNumbers:              []byte{ty.GetQuorum()},
		QuorumThresholdPercentages: []byte{0},
		ReferenceBlockNumber:       uint32(ext.ReferenceBlockNumber),
	}

	digest, err := stateHeader.Digest()
	if err != nil {
		return logex.Trace(err)
	}
	sig := o.cfg.BlsKey.SignMessage(digest)

	o.operatorMetric.SubmitTask.WithLabelValues(o.cfg.AvsName, ty.Value()).Add(1)
	submitTaskTime := time.Now()
	// submit to aggregator
	if err := o.aggregator.SubmitTask(ctx, &aggregator.TaskRequest{
		Task:            stateHeader,
		Signature:       sig,
		OperatorId:      o.operatorId,
		ProverSignature: poe.PoeSignature,
	}); err != nil {
		return logex.Trace(err)
	}
	o.operatorMetric.SubmitTaskMs.WithLabelValues(o.cfg.AvsName, ty.Value()).Set(float64(time.Since(submitTaskTime).Milliseconds()))
	// logex.Info(poe)
	return nil
}

func (o *Operator) checkIsRegistered() error {
	operatorIsRegistered, err := o.cfg.EigenClients.AvsRegistryChainReader.IsOperatorRegistered(nil, o.operatorAddress)
	if err != nil {
		return logex.Trace(err)
	}
	if !operatorIsRegistered {
		return logex.NewErrorf("operator[%v] is not registered", o.operatorAddress)
	}
	o.operatorId, err = o.cfg.EigenClients.AvsRegistryChainReader.GetOperatorId(nil, o.operatorAddress)
	if err != nil {
		return logex.Trace(err)
	}
	quorumNumbers, _, err := o.cfg.EigenClients.AvsRegistryChainReader.GetOperatorsStakeInQuorumsOfOperatorAtCurrentBlock(&bind.CallOpts{}, o.operatorId)
	if err != nil {
		return logex.Trace(err)
	}
	o.quorumNumbers = make([]byte, len(quorumNumbers))
	for i, qn := range quorumNumbers {
		o.quorumNumbers[i] = byte(qn)
	}
	return nil
}

var ABI = func() abi.ABI {
	ty := `[{"inputs":[{"internalType":"uint8","name":"_version","type":"uint8"},{"internalType":"bytes","name":"_parentBatchHeader","type":"bytes"},{"internalType":"bytes[]","name":"_chunks","type":"bytes[]"},{"internalType":"bytes","name":"_skippedL1MessageBitmap","type":"bytes"}],"name":"commitBatch","outputs":[],"stateMutability":"nonpayable","type":"function"}]`
	result, err := abi.JSON(bytes.NewReader([]byte(ty)))
	if err != nil {
		panic(err)
	}
	return result
}()

func (o *Operator) proverGetPoe(ctx context.Context, txHash common.Hash, topics []common.Hash) (*xtask.PoeResponse, bool, error) {
	logex.Infof("fetching poe for batch %v", topics[2])
	poe, skip, err := o.proverClient.GetPoe(ctx, txHash)
	if err != nil {
		return nil, skip, logex.Trace(err)
	}
	return poe, skip, nil
}

func (o *Operator) proverGetAttestationReport(ctx context.Context, pubkey []byte) ([]byte, error) {
	quote, err := o.proverClient.GenerateAttestaionReport(ctx, pubkey)
	if err != nil {
		return nil, logex.Trace(err)
	}
	return quote, nil
}

func (o *Operator) selectReferenceBlock(ctx context.Context, reportData *TEELivenessVerifier.TEELivenessVerifierReportDataV2) error {
	// corner case:
	//  1. block numbers are not sequential
	//  2. the types.Header.Hash() may not compatible with the chain
	headBlock, err := o.cfg.AttestationClient.HeaderByNumber(ctx, nil)
	if err != nil {
		return logex.Trace(err)
	}
	reportData.ReferenceBlockHash = headBlock.ParentHash
	referenceBlock, err := o.cfg.AttestationClient.HeaderByHash(ctx, headBlock.ParentHash)
	if err != nil {
		return logex.Trace(err)
	}
	reportData.ReferenceBlockNumber = referenceBlock.Number
	return nil
}

func (o *Operator) registerAttestationReport(ctx context.Context, pubkeyBytes []byte) error {
	var reportData TEELivenessVerifier.TEELivenessVerifierReportDataV2
	copy(reportData.Pubkey.X[:], pubkeyBytes[:32])
	copy(reportData.Pubkey.Y[:], pubkeyBytes[32:])
	reportData.ProverAddressHash = utils.ProverAddrHash(o.cfg.Config.ProverURL)
	if err := o.selectReferenceBlock(ctx, &reportData); err != nil {
		return logex.Trace(err)
	}

	dataHash, err := bindings.ReportDataHash(&reportData)
	if err != nil {
		return logex.Trace(err)
	}
	reportDataBytes := make([]byte, 64)
	copy(reportDataBytes[32:], dataHash[:])

	startGenReport := time.Now()
	report, err := o.proverGetAttestationReport(ctx, reportDataBytes)
	if err != nil {
		return logex.Trace(err)
	}
	o.operatorMetric.GenReportMs.WithLabelValues(o.cfg.AvsName).Set(float64(time.Since(startGenReport).Milliseconds()))

	chainId, err := o.cfg.AttestationClient.ChainID(ctx)
	if err != nil {
		return logex.Trace(err)
	}
	opt, err := bind.NewKeyedTransactorWithChainID(o.cfg.AttestationEcdsaKey, chainId)
	if err != nil {
		return logex.Trace(err)
	}

	tx, err := o.TEELivenessVerifier.SubmitLivenessProofV2(opt, reportData, report)
	if err != nil {
		balance, checkErr := o.cfg.AttestationClient.BalanceAt(ctx, opt.From, nil)
		if checkErr != nil {
			return logex.Trace(err, "check balance")
		}
		if rdb, err := bindings.ReportDataBytes(&reportData); err == nil {
			logex.Infof("attestation report userdata: 0x%x", rdb)
		}
		logex.Infof("attestation report data: 0x%x", report)
		return logex.Trace(
			bindings.MultiProverError(err),
			fmt.Sprintf("balance:%.2f", utils.WeiToF64(balance, 18)),
			fmt.Sprintf("verifierAddr:%v", o.cfg.Config.TEELivenessVerifierAddress),
		)
	}

	logex.Infof("submitted liveness proof: %v", tx.Hash())
	receipt, err := utils.WaitTx(ctx, o.cfg.AttestationClient, tx, nil)
	if err != nil {
		return logex.Trace(err)
	}
	gasCost := utils.GetTxCost(receipt)

	balance, err := o.cfg.AttestationClient.BalanceAt(ctx, opt.From, nil)
	if err != nil {
		return logex.Trace(err, "check balance")
	}
	o.operatorMetric.AttestationAccBalance.WithLabelValues(o.cfg.AvsName).Set(utils.WeiToF64(balance, 18))

	o.operatorMetric.LastAttestationCost.WithLabelValues(o.cfg.AvsName).Set(gasCost)
	logex.Infof("registered in TEELivenessVerifier: %v", tx.Hash())
	return nil
}

func (o *Operator) RegisterAttestationReport(ctx context.Context) error {
	attestationAddr := utils.EcdsaAddress(o.cfg.AttestationEcdsaKey)
	logex.Infof("checking tee liveness... attestationLayerEcdsaAddress=%v", attestationAddr)
	pubkeyBytes := o.cfg.BlsKey.PubKey.Serialize()
	if len(pubkeyBytes) != 64 {
		return logex.NewErrorf("invalid pubkey")
	}

	var x, y [32]byte
	copy(x[:], pubkeyBytes[:32])
	copy(y[:], pubkeyBytes[32:64])
	isRegistered, err := o.TEELivenessVerifier.VerifyLivenessProof(nil, x, y)
	if err != nil {
		return logex.Trace(err)
	}

	balance, err := o.cfg.AttestationClient.BalanceAt(ctx, attestationAddr, nil)
	if err != nil {
		return logex.Trace(err, "check balance")
	}
	o.operatorMetric.AttestationAccBalance.WithLabelValues(o.cfg.AvsName).Set(utils.WeiToF64(balance, 18))

	if isRegistered {
		logex.Info("Operater has registered on TEE Liveness Verifier")
	} else {
		if err := o.registerAttestationReport(ctx, pubkeyBytes); err != nil {
			return logex.Trace(err)
		}
	}

	checkNext := func(ctx context.Context) error {
		validSecs, err := o.TEELivenessVerifier.AttestValiditySeconds(nil)
		if err != nil {
			return logex.Trace(err)
		}
		key := crypto.Keccak256Hash(pubkeyBytes)
		prover, err := o.TEELivenessVerifier.AttestedProvers(nil, key)
		if err != nil {
			return logex.Trace(err)
		}
		o.operatorMetric.LivenessTs.WithLabelValues(o.cfg.AvsName).Set(float64(prover.Time.Int64()))
		deadline := prover.Time.Int64() + validSecs.Int64()
		now := time.Now().Unix()
		o.operatorMetric.NextAttestationTs.WithLabelValues(o.cfg.AvsName).Set(float64(deadline))
		logex.Info("next attestation will be at", time.Unix(deadline, 0), ",validSecs=", validSecs.Int64())
		if deadline > now+300 {
			time.Sleep(time.Duration(deadline-now-300) * time.Second)
		}
		return o.registerAttestationReport(ctx, pubkeyBytes)
	}
	go func() {
		ctx := context.Background()
		for {
			if err := checkNext(ctx); err != nil {
				logex.Error(err)
				time.Sleep(10 * time.Second)
			}
		}
	}()

	return nil
}
