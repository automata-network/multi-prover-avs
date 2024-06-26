// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package MultiProverServiceManager

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
	_ = abi.ConvertType
)

// BN254G1Point is an auto generated low-level Go binding around an user-defined struct.
type BN254G1Point struct {
	X *big.Int
	Y *big.Int
}

// BN254G2Point is an auto generated low-level Go binding around an user-defined struct.
type BN254G2Point struct {
	X [2]*big.Int
	Y [2]*big.Int
}

// IBLSSignatureCheckerNonSignerStakesAndSignature is an auto generated low-level Go binding around an user-defined struct.
type IBLSSignatureCheckerNonSignerStakesAndSignature struct {
	NonSignerQuorumBitmapIndices []uint32
	NonSignerPubkeys             []BN254G1Point
	QuorumApks                   []BN254G1Point
	ApkG2                        BN254G2Point
	Sigma                        BN254G1Point
	QuorumApkIndices             []uint32
	TotalStakeIndices            []uint32
	NonSignerStakeIndices        [][]uint32
}

// IBLSSignatureCheckerQuorumStakeTotals is an auto generated low-level Go binding around an user-defined struct.
type IBLSSignatureCheckerQuorumStakeTotals struct {
	SignedStakeForQuorum []*big.Int
	TotalStakeForQuorum  []*big.Int
}

// IMultiProverServiceManagerCommittee is an auto generated low-level Go binding around an user-defined struct.
type IMultiProverServiceManagerCommittee struct {
	Id               *big.Int
	Description      string
	Metadata         []byte
	TeeQuorumNumbers []byte
}

// IMultiProverServiceManagerReducedStateHeader is an auto generated low-level Go binding around an user-defined struct.
type IMultiProverServiceManagerReducedStateHeader struct {
	CommitteeId          *big.Int
	Metadata             []byte
	State                []byte
	ReferenceBlockNumber uint32
}

// IMultiProverServiceManagerStateHeader is an auto generated low-level Go binding around an user-defined struct.
type IMultiProverServiceManagerStateHeader struct {
	CommitteeId                *big.Int
	Metadata                   []byte
	State                      []byte
	QuorumNumbers              []byte
	QuorumThresholdPercentages []byte
	ReferenceBlockNumber       uint32
}

// IMultiProverServiceManagerTEEQuorum is an auto generated low-level Go binding around an user-defined struct.
type IMultiProverServiceManagerTEEQuorum struct {
	TeeType      uint8
	QuorumNumber uint8
}

// ISignatureUtilsSignatureWithSaltAndExpiry is an auto generated low-level Go binding around an user-defined struct.
type ISignatureUtilsSignatureWithSaltAndExpiry struct {
	Signature []byte
	Salt      [32]byte
	Expiry    *big.Int
}

// MultiProverServiceManagerMetaData contains all meta data concerning the MultiProverServiceManager contract.
var MultiProverServiceManagerMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"__avsDirectory\",\"type\":\"address\",\"internalType\":\"contractIAVSDirectory\"},{\"name\":\"__registryCoordinator\",\"type\":\"address\",\"internalType\":\"contractIRegistryCoordinator\"},{\"name\":\"__stakeRegistry\",\"type\":\"address\",\"internalType\":\"contractIStakeRegistry\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"PAUSED_OPERTOR_REGISTRATION\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"PAUSED_SUBMIT_STATE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"THRESHOLD_DENOMINATOR\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"_hashReducedStateHeader\",\"inputs\":[{\"name\":\"reducedStateHeader\",\"type\":\"tuple\",\"internalType\":\"structIMultiProverServiceManager.ReducedStateHeader\",\"components\":[{\"name\":\"committeeId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"metadata\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"state\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"referenceBlockNumber\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"addCommittee\",\"inputs\":[{\"name\":\"committee\",\"type\":\"tuple\",\"internalType\":\"structIMultiProverServiceManager.Committee\",\"components\":[{\"name\":\"id\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"description\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"metadata\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"teeQuorumNumbers\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"addTEEQuorum\",\"inputs\":[{\"name\":\"teeQuorum\",\"type\":\"tuple\",\"internalType\":\"structIMultiProverServiceManager.TEEQuorum\",\"components\":[{\"name\":\"teeType\",\"type\":\"uint8\",\"internalType\":\"enumIMultiProverServiceManager.TEE\"},{\"name\":\"quorumNumber\",\"type\":\"uint8\",\"internalType\":\"uint8\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"avsDirectory\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"blacklistOperator\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"blsApkRegistry\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIBLSApkRegistry\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"checkSignatures\",\"inputs\":[{\"name\":\"msgHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"quorumNumbers\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"referenceBlockNumber\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"params\",\"type\":\"tuple\",\"internalType\":\"structIBLSSignatureChecker.NonSignerStakesAndSignature\",\"components\":[{\"name\":\"nonSignerQuorumBitmapIndices\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"},{\"name\":\"nonSignerPubkeys\",\"type\":\"tuple[]\",\"internalType\":\"structBN254.G1Point[]\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"quorumApks\",\"type\":\"tuple[]\",\"internalType\":\"structBN254.G1Point[]\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"apkG2\",\"type\":\"tuple\",\"internalType\":\"structBN254.G2Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256[2]\",\"internalType\":\"uint256[2]\"},{\"name\":\"Y\",\"type\":\"uint256[2]\",\"internalType\":\"uint256[2]\"}]},{\"name\":\"sigma\",\"type\":\"tuple\",\"internalType\":\"structBN254.G1Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"quorumApkIndices\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"},{\"name\":\"totalStakeIndices\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"},{\"name\":\"nonSignerStakeIndices\",\"type\":\"uint32[][]\",\"internalType\":\"uint32[][]\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structIBLSSignatureChecker.QuorumStakeTotals\",\"components\":[{\"name\":\"signedStakeForQuorum\",\"type\":\"uint96[]\",\"internalType\":\"uint96[]\"},{\"name\":\"totalStakeForQuorum\",\"type\":\"uint96[]\",\"internalType\":\"uint96[]\"}]},{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"committeeManager\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"committees\",\"inputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"id\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"description\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"metadata\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"teeQuorumNumbers\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"confirmState\",\"inputs\":[{\"name\":\"stateHeader\",\"type\":\"tuple\",\"internalType\":\"structIMultiProverServiceManager.StateHeader\",\"components\":[{\"name\":\"committeeId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"metadata\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"state\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"quorumNumbers\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"quorumThresholdPercentages\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"referenceBlockNumber\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"nonSignerStakesAndSignature\",\"type\":\"tuple\",\"internalType\":\"structIBLSSignatureChecker.NonSignerStakesAndSignature\",\"components\":[{\"name\":\"nonSignerQuorumBitmapIndices\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"},{\"name\":\"nonSignerPubkeys\",\"type\":\"tuple[]\",\"internalType\":\"structBN254.G1Point[]\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"quorumApks\",\"type\":\"tuple[]\",\"internalType\":\"structBN254.G1Point[]\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"apkG2\",\"type\":\"tuple\",\"internalType\":\"structBN254.G2Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256[2]\",\"internalType\":\"uint256[2]\"},{\"name\":\"Y\",\"type\":\"uint256[2]\",\"internalType\":\"uint256[2]\"}]},{\"name\":\"sigma\",\"type\":\"tuple\",\"internalType\":\"structBN254.G1Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"quorumApkIndices\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"},{\"name\":\"totalStakeIndices\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"},{\"name\":\"nonSignerStakeIndices\",\"type\":\"uint32[][]\",\"internalType\":\"uint32[][]\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"delegation\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIDelegationManager\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"deregisterOperatorFromAVS\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"disablePoA\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"enablePoA\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getOperatorRestakedStrategies\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRestakeableStrategies\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"initialize\",\"inputs\":[{\"name\":\"_pauserRegistry\",\"type\":\"address\",\"internalType\":\"contractIPauserRegistry\"},{\"name\":\"_initialPausedStatus\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_initialOwner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_stateConfirmer\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_poaManager\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_committeeManager\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_poaEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"isOperatorWhitelisted\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isPoAEnabled\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"pause\",\"inputs\":[{\"name\":\"newPausedStatus\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"pauseAll\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"paused\",\"inputs\":[{\"name\":\"index\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"paused\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"pauserRegistry\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIPauserRegistry\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"poaEnabled\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"poaManager\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"quorumIdToCommitteeId\",\"inputs\":[{\"name\":\"\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"registerOperatorToAVS\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"operatorSignature\",\"type\":\"tuple\",\"internalType\":\"structISignatureUtils.SignatureWithSaltAndExpiry\",\"components\":[{\"name\":\"signature\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"salt\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"expiry\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"registryCoordinator\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIRegistryCoordinator\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"removeCommittee\",\"inputs\":[{\"name\":\"committeeId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"removeTEEQuorum\",\"inputs\":[{\"name\":\"quorumNumber\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"renounceOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setCommitteeManager\",\"inputs\":[{\"name\":\"_committeeManager\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setPauserRegistry\",\"inputs\":[{\"name\":\"newPauserRegistry\",\"type\":\"address\",\"internalType\":\"contractIPauserRegistry\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setPoAManager\",\"inputs\":[{\"name\":\"_poaManager\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setStaleStakesForbidden\",\"inputs\":[{\"name\":\"value\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setStateConfirmer\",\"inputs\":[{\"name\":\"_stateConfirmer\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"stakeRegistry\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIStakeRegistry\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"staleStakesForbidden\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"stateConfirmer\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"taskId\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"taskIdToMetadataHash\",\"inputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"teeQuorums\",\"inputs\":[{\"name\":\"\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"outputs\":[{\"name\":\"teeType\",\"type\":\"uint8\",\"internalType\":\"enumIMultiProverServiceManager.TEE\"},{\"name\":\"quorumNumber\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"trySignatureAndApkVerification\",\"inputs\":[{\"name\":\"msgHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"apk\",\"type\":\"tuple\",\"internalType\":\"structBN254.G1Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"apkG2\",\"type\":\"tuple\",\"internalType\":\"structBN254.G2Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256[2]\",\"internalType\":\"uint256[2]\"},{\"name\":\"Y\",\"type\":\"uint256[2]\",\"internalType\":\"uint256[2]\"}]},{\"name\":\"sigma\",\"type\":\"tuple\",\"internalType\":\"structBN254.G1Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"outputs\":[{\"name\":\"pairingSuccessful\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"siganatureIsValid\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"unpause\",\"inputs\":[{\"name\":\"newPausedStatus\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"updateAVSMetadataURI\",\"inputs\":[{\"name\":\"_metadataURI\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"updateCommittee\",\"inputs\":[{\"name\":\"committee\",\"type\":\"tuple\",\"internalType\":\"structIMultiProverServiceManager.Committee\",\"components\":[{\"name\":\"id\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"description\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"metadata\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"teeQuorumNumbers\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"whitelistOperator\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"CommitteeManagerUpdated\",\"inputs\":[{\"name\":\"previousCommitteeManager\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"currentCommitteeManager\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Initialized\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"uint8\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"previousOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Paused\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newPausedStatus\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"PauserRegistrySet\",\"inputs\":[{\"name\":\"pauserRegistry\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"contractIPauserRegistry\"},{\"name\":\"newPauserRegistry\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"contractIPauserRegistry\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"PoAManagerUpdated\",\"inputs\":[{\"name\":\"previousPoaManager\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"currentPoaManager\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"StaleStakesForbiddenUpdate\",\"inputs\":[{\"name\":\"value\",\"type\":\"bool\",\"indexed\":false,\"internalType\":\"bool\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"StateConfirmed\",\"inputs\":[{\"name\":\"committeeId\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"metadata\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"},{\"name\":\"state\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"StateConfirmerUpdated\",\"inputs\":[{\"name\":\"previousConfirmer\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"currentConfirmer\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Unpaused\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newPausedStatus\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"CommitteeExist\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"CommitteeNotExist\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InsufficientThreshold\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidQuorum\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidQuorumParam\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidSender\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NoPermission\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotWhitelisted\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"QuorumNotInitialized\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"TEEQuorumExist\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"TEEQuorumNotExist\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"TEEQuorumUsed\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ZeroAddr\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ZeroId\",\"inputs\":[]}]",
}

// MultiProverServiceManagerABI is the input ABI used to generate the binding from.
// Deprecated: Use MultiProverServiceManagerMetaData.ABI instead.
var MultiProverServiceManagerABI = MultiProverServiceManagerMetaData.ABI

// MultiProverServiceManager is an auto generated Go binding around an Ethereum contract.
type MultiProverServiceManager struct {
	MultiProverServiceManagerCaller     // Read-only binding to the contract
	MultiProverServiceManagerTransactor // Write-only binding to the contract
	MultiProverServiceManagerFilterer   // Log filterer for contract events
}

// MultiProverServiceManagerCaller is an auto generated read-only Go binding around an Ethereum contract.
type MultiProverServiceManagerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MultiProverServiceManagerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MultiProverServiceManagerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MultiProverServiceManagerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MultiProverServiceManagerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MultiProverServiceManagerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MultiProverServiceManagerSession struct {
	Contract     *MultiProverServiceManager // Generic contract binding to set the session for
	CallOpts     bind.CallOpts              // Call options to use throughout this session
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// MultiProverServiceManagerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MultiProverServiceManagerCallerSession struct {
	Contract *MultiProverServiceManagerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                    // Call options to use throughout this session
}

// MultiProverServiceManagerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MultiProverServiceManagerTransactorSession struct {
	Contract     *MultiProverServiceManagerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                    // Transaction auth options to use throughout this session
}

// MultiProverServiceManagerRaw is an auto generated low-level Go binding around an Ethereum contract.
type MultiProverServiceManagerRaw struct {
	Contract *MultiProverServiceManager // Generic contract binding to access the raw methods on
}

// MultiProverServiceManagerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MultiProverServiceManagerCallerRaw struct {
	Contract *MultiProverServiceManagerCaller // Generic read-only contract binding to access the raw methods on
}

// MultiProverServiceManagerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MultiProverServiceManagerTransactorRaw struct {
	Contract *MultiProverServiceManagerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMultiProverServiceManager creates a new instance of MultiProverServiceManager, bound to a specific deployed contract.
func NewMultiProverServiceManager(address common.Address, backend bind.ContractBackend) (*MultiProverServiceManager, error) {
	contract, err := bindMultiProverServiceManager(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &MultiProverServiceManager{MultiProverServiceManagerCaller: MultiProverServiceManagerCaller{contract: contract}, MultiProverServiceManagerTransactor: MultiProverServiceManagerTransactor{contract: contract}, MultiProverServiceManagerFilterer: MultiProverServiceManagerFilterer{contract: contract}}, nil
}

// NewMultiProverServiceManagerCaller creates a new read-only instance of MultiProverServiceManager, bound to a specific deployed contract.
func NewMultiProverServiceManagerCaller(address common.Address, caller bind.ContractCaller) (*MultiProverServiceManagerCaller, error) {
	contract, err := bindMultiProverServiceManager(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MultiProverServiceManagerCaller{contract: contract}, nil
}

// NewMultiProverServiceManagerTransactor creates a new write-only instance of MultiProverServiceManager, bound to a specific deployed contract.
func NewMultiProverServiceManagerTransactor(address common.Address, transactor bind.ContractTransactor) (*MultiProverServiceManagerTransactor, error) {
	contract, err := bindMultiProverServiceManager(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MultiProverServiceManagerTransactor{contract: contract}, nil
}

// NewMultiProverServiceManagerFilterer creates a new log filterer instance of MultiProverServiceManager, bound to a specific deployed contract.
func NewMultiProverServiceManagerFilterer(address common.Address, filterer bind.ContractFilterer) (*MultiProverServiceManagerFilterer, error) {
	contract, err := bindMultiProverServiceManager(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MultiProverServiceManagerFilterer{contract: contract}, nil
}

// bindMultiProverServiceManager binds a generic wrapper to an already deployed contract.
func bindMultiProverServiceManager(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := MultiProverServiceManagerMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MultiProverServiceManager *MultiProverServiceManagerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MultiProverServiceManager.Contract.MultiProverServiceManagerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MultiProverServiceManager *MultiProverServiceManagerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MultiProverServiceManager.Contract.MultiProverServiceManagerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MultiProverServiceManager *MultiProverServiceManagerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MultiProverServiceManager.Contract.MultiProverServiceManagerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MultiProverServiceManager *MultiProverServiceManagerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MultiProverServiceManager.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MultiProverServiceManager *MultiProverServiceManagerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MultiProverServiceManager.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MultiProverServiceManager *MultiProverServiceManagerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MultiProverServiceManager.Contract.contract.Transact(opts, method, params...)
}

// PAUSEDOPERTORREGISTRATION is a free data retrieval call binding the contract method 0xe9f7903f.
//
// Solidity: function PAUSED_OPERTOR_REGISTRATION() view returns(uint8)
func (_MultiProverServiceManager *MultiProverServiceManagerCaller) PAUSEDOPERTORREGISTRATION(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _MultiProverServiceManager.contract.Call(opts, &out, "PAUSED_OPERTOR_REGISTRATION")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// PAUSEDOPERTORREGISTRATION is a free data retrieval call binding the contract method 0xe9f7903f.
//
// Solidity: function PAUSED_OPERTOR_REGISTRATION() view returns(uint8)
func (_MultiProverServiceManager *MultiProverServiceManagerSession) PAUSEDOPERTORREGISTRATION() (uint8, error) {
	return _MultiProverServiceManager.Contract.PAUSEDOPERTORREGISTRATION(&_MultiProverServiceManager.CallOpts)
}

// PAUSEDOPERTORREGISTRATION is a free data retrieval call binding the contract method 0xe9f7903f.
//
// Solidity: function PAUSED_OPERTOR_REGISTRATION() view returns(uint8)
func (_MultiProverServiceManager *MultiProverServiceManagerCallerSession) PAUSEDOPERTORREGISTRATION() (uint8, error) {
	return _MultiProverServiceManager.Contract.PAUSEDOPERTORREGISTRATION(&_MultiProverServiceManager.CallOpts)
}

// PAUSEDSUBMITSTATE is a free data retrieval call binding the contract method 0x8a99056c.
//
// Solidity: function PAUSED_SUBMIT_STATE() view returns(uint8)
func (_MultiProverServiceManager *MultiProverServiceManagerCaller) PAUSEDSUBMITSTATE(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _MultiProverServiceManager.contract.Call(opts, &out, "PAUSED_SUBMIT_STATE")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// PAUSEDSUBMITSTATE is a free data retrieval call binding the contract method 0x8a99056c.
//
// Solidity: function PAUSED_SUBMIT_STATE() view returns(uint8)
func (_MultiProverServiceManager *MultiProverServiceManagerSession) PAUSEDSUBMITSTATE() (uint8, error) {
	return _MultiProverServiceManager.Contract.PAUSEDSUBMITSTATE(&_MultiProverServiceManager.CallOpts)
}

// PAUSEDSUBMITSTATE is a free data retrieval call binding the contract method 0x8a99056c.
//
// Solidity: function PAUSED_SUBMIT_STATE() view returns(uint8)
func (_MultiProverServiceManager *MultiProverServiceManagerCallerSession) PAUSEDSUBMITSTATE() (uint8, error) {
	return _MultiProverServiceManager.Contract.PAUSEDSUBMITSTATE(&_MultiProverServiceManager.CallOpts)
}

// THRESHOLDDENOMINATOR is a free data retrieval call binding the contract method 0xef024458.
//
// Solidity: function THRESHOLD_DENOMINATOR() view returns(uint256)
func (_MultiProverServiceManager *MultiProverServiceManagerCaller) THRESHOLDDENOMINATOR(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _MultiProverServiceManager.contract.Call(opts, &out, "THRESHOLD_DENOMINATOR")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// THRESHOLDDENOMINATOR is a free data retrieval call binding the contract method 0xef024458.
//
// Solidity: function THRESHOLD_DENOMINATOR() view returns(uint256)
func (_MultiProverServiceManager *MultiProverServiceManagerSession) THRESHOLDDENOMINATOR() (*big.Int, error) {
	return _MultiProverServiceManager.Contract.THRESHOLDDENOMINATOR(&_MultiProverServiceManager.CallOpts)
}

// THRESHOLDDENOMINATOR is a free data retrieval call binding the contract method 0xef024458.
//
// Solidity: function THRESHOLD_DENOMINATOR() view returns(uint256)
func (_MultiProverServiceManager *MultiProverServiceManagerCallerSession) THRESHOLDDENOMINATOR() (*big.Int, error) {
	return _MultiProverServiceManager.Contract.THRESHOLDDENOMINATOR(&_MultiProverServiceManager.CallOpts)
}

// HashReducedStateHeader is a free data retrieval call binding the contract method 0xe6d52594.
//
// Solidity: function _hashReducedStateHeader((uint256,bytes,bytes,uint32) reducedStateHeader) pure returns(bytes32)
func (_MultiProverServiceManager *MultiProverServiceManagerCaller) HashReducedStateHeader(opts *bind.CallOpts, reducedStateHeader IMultiProverServiceManagerReducedStateHeader) ([32]byte, error) {
	var out []interface{}
	err := _MultiProverServiceManager.contract.Call(opts, &out, "_hashReducedStateHeader", reducedStateHeader)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// HashReducedStateHeader is a free data retrieval call binding the contract method 0xe6d52594.
//
// Solidity: function _hashReducedStateHeader((uint256,bytes,bytes,uint32) reducedStateHeader) pure returns(bytes32)
func (_MultiProverServiceManager *MultiProverServiceManagerSession) HashReducedStateHeader(reducedStateHeader IMultiProverServiceManagerReducedStateHeader) ([32]byte, error) {
	return _MultiProverServiceManager.Contract.HashReducedStateHeader(&_MultiProverServiceManager.CallOpts, reducedStateHeader)
}

// HashReducedStateHeader is a free data retrieval call binding the contract method 0xe6d52594.
//
// Solidity: function _hashReducedStateHeader((uint256,bytes,bytes,uint32) reducedStateHeader) pure returns(bytes32)
func (_MultiProverServiceManager *MultiProverServiceManagerCallerSession) HashReducedStateHeader(reducedStateHeader IMultiProverServiceManagerReducedStateHeader) ([32]byte, error) {
	return _MultiProverServiceManager.Contract.HashReducedStateHeader(&_MultiProverServiceManager.CallOpts, reducedStateHeader)
}

// AvsDirectory is a free data retrieval call binding the contract method 0x6b3aa72e.
//
// Solidity: function avsDirectory() view returns(address)
func (_MultiProverServiceManager *MultiProverServiceManagerCaller) AvsDirectory(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MultiProverServiceManager.contract.Call(opts, &out, "avsDirectory")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AvsDirectory is a free data retrieval call binding the contract method 0x6b3aa72e.
//
// Solidity: function avsDirectory() view returns(address)
func (_MultiProverServiceManager *MultiProverServiceManagerSession) AvsDirectory() (common.Address, error) {
	return _MultiProverServiceManager.Contract.AvsDirectory(&_MultiProverServiceManager.CallOpts)
}

// AvsDirectory is a free data retrieval call binding the contract method 0x6b3aa72e.
//
// Solidity: function avsDirectory() view returns(address)
func (_MultiProverServiceManager *MultiProverServiceManagerCallerSession) AvsDirectory() (common.Address, error) {
	return _MultiProverServiceManager.Contract.AvsDirectory(&_MultiProverServiceManager.CallOpts)
}

// BlsApkRegistry is a free data retrieval call binding the contract method 0x5df45946.
//
// Solidity: function blsApkRegistry() view returns(address)
func (_MultiProverServiceManager *MultiProverServiceManagerCaller) BlsApkRegistry(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MultiProverServiceManager.contract.Call(opts, &out, "blsApkRegistry")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// BlsApkRegistry is a free data retrieval call binding the contract method 0x5df45946.
//
// Solidity: function blsApkRegistry() view returns(address)
func (_MultiProverServiceManager *MultiProverServiceManagerSession) BlsApkRegistry() (common.Address, error) {
	return _MultiProverServiceManager.Contract.BlsApkRegistry(&_MultiProverServiceManager.CallOpts)
}

// BlsApkRegistry is a free data retrieval call binding the contract method 0x5df45946.
//
// Solidity: function blsApkRegistry() view returns(address)
func (_MultiProverServiceManager *MultiProverServiceManagerCallerSession) BlsApkRegistry() (common.Address, error) {
	return _MultiProverServiceManager.Contract.BlsApkRegistry(&_MultiProverServiceManager.CallOpts)
}

// CheckSignatures is a free data retrieval call binding the contract method 0x6efb4636.
//
// Solidity: function checkSignatures(bytes32 msgHash, bytes quorumNumbers, uint32 referenceBlockNumber, (uint32[],(uint256,uint256)[],(uint256,uint256)[],(uint256[2],uint256[2]),(uint256,uint256),uint32[],uint32[],uint32[][]) params) view returns((uint96[],uint96[]), bytes32)
func (_MultiProverServiceManager *MultiProverServiceManagerCaller) CheckSignatures(opts *bind.CallOpts, msgHash [32]byte, quorumNumbers []byte, referenceBlockNumber uint32, params IBLSSignatureCheckerNonSignerStakesAndSignature) (IBLSSignatureCheckerQuorumStakeTotals, [32]byte, error) {
	var out []interface{}
	err := _MultiProverServiceManager.contract.Call(opts, &out, "checkSignatures", msgHash, quorumNumbers, referenceBlockNumber, params)

	if err != nil {
		return *new(IBLSSignatureCheckerQuorumStakeTotals), *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new(IBLSSignatureCheckerQuorumStakeTotals)).(*IBLSSignatureCheckerQuorumStakeTotals)
	out1 := *abi.ConvertType(out[1], new([32]byte)).(*[32]byte)

	return out0, out1, err

}

// CheckSignatures is a free data retrieval call binding the contract method 0x6efb4636.
//
// Solidity: function checkSignatures(bytes32 msgHash, bytes quorumNumbers, uint32 referenceBlockNumber, (uint32[],(uint256,uint256)[],(uint256,uint256)[],(uint256[2],uint256[2]),(uint256,uint256),uint32[],uint32[],uint32[][]) params) view returns((uint96[],uint96[]), bytes32)
func (_MultiProverServiceManager *MultiProverServiceManagerSession) CheckSignatures(msgHash [32]byte, quorumNumbers []byte, referenceBlockNumber uint32, params IBLSSignatureCheckerNonSignerStakesAndSignature) (IBLSSignatureCheckerQuorumStakeTotals, [32]byte, error) {
	return _MultiProverServiceManager.Contract.CheckSignatures(&_MultiProverServiceManager.CallOpts, msgHash, quorumNumbers, referenceBlockNumber, params)
}

// CheckSignatures is a free data retrieval call binding the contract method 0x6efb4636.
//
// Solidity: function checkSignatures(bytes32 msgHash, bytes quorumNumbers, uint32 referenceBlockNumber, (uint32[],(uint256,uint256)[],(uint256,uint256)[],(uint256[2],uint256[2]),(uint256,uint256),uint32[],uint32[],uint32[][]) params) view returns((uint96[],uint96[]), bytes32)
func (_MultiProverServiceManager *MultiProverServiceManagerCallerSession) CheckSignatures(msgHash [32]byte, quorumNumbers []byte, referenceBlockNumber uint32, params IBLSSignatureCheckerNonSignerStakesAndSignature) (IBLSSignatureCheckerQuorumStakeTotals, [32]byte, error) {
	return _MultiProverServiceManager.Contract.CheckSignatures(&_MultiProverServiceManager.CallOpts, msgHash, quorumNumbers, referenceBlockNumber, params)
}

// CommitteeManager is a free data retrieval call binding the contract method 0x7e8611e2.
//
// Solidity: function committeeManager() view returns(address)
func (_MultiProverServiceManager *MultiProverServiceManagerCaller) CommitteeManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MultiProverServiceManager.contract.Call(opts, &out, "committeeManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// CommitteeManager is a free data retrieval call binding the contract method 0x7e8611e2.
//
// Solidity: function committeeManager() view returns(address)
func (_MultiProverServiceManager *MultiProverServiceManagerSession) CommitteeManager() (common.Address, error) {
	return _MultiProverServiceManager.Contract.CommitteeManager(&_MultiProverServiceManager.CallOpts)
}

// CommitteeManager is a free data retrieval call binding the contract method 0x7e8611e2.
//
// Solidity: function committeeManager() view returns(address)
func (_MultiProverServiceManager *MultiProverServiceManagerCallerSession) CommitteeManager() (common.Address, error) {
	return _MultiProverServiceManager.Contract.CommitteeManager(&_MultiProverServiceManager.CallOpts)
}

// Committees is a free data retrieval call binding the contract method 0xf5e820fd.
//
// Solidity: function committees(uint256 ) view returns(uint256 id, string description, bytes metadata, bytes teeQuorumNumbers)
func (_MultiProverServiceManager *MultiProverServiceManagerCaller) Committees(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Id               *big.Int
	Description      string
	Metadata         []byte
	TeeQuorumNumbers []byte
}, error) {
	var out []interface{}
	err := _MultiProverServiceManager.contract.Call(opts, &out, "committees", arg0)

	outstruct := new(struct {
		Id               *big.Int
		Description      string
		Metadata         []byte
		TeeQuorumNumbers []byte
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Id = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Description = *abi.ConvertType(out[1], new(string)).(*string)
	outstruct.Metadata = *abi.ConvertType(out[2], new([]byte)).(*[]byte)
	outstruct.TeeQuorumNumbers = *abi.ConvertType(out[3], new([]byte)).(*[]byte)

	return *outstruct, err

}

// Committees is a free data retrieval call binding the contract method 0xf5e820fd.
//
// Solidity: function committees(uint256 ) view returns(uint256 id, string description, bytes metadata, bytes teeQuorumNumbers)
func (_MultiProverServiceManager *MultiProverServiceManagerSession) Committees(arg0 *big.Int) (struct {
	Id               *big.Int
	Description      string
	Metadata         []byte
	TeeQuorumNumbers []byte
}, error) {
	return _MultiProverServiceManager.Contract.Committees(&_MultiProverServiceManager.CallOpts, arg0)
}

// Committees is a free data retrieval call binding the contract method 0xf5e820fd.
//
// Solidity: function committees(uint256 ) view returns(uint256 id, string description, bytes metadata, bytes teeQuorumNumbers)
func (_MultiProverServiceManager *MultiProverServiceManagerCallerSession) Committees(arg0 *big.Int) (struct {
	Id               *big.Int
	Description      string
	Metadata         []byte
	TeeQuorumNumbers []byte
}, error) {
	return _MultiProverServiceManager.Contract.Committees(&_MultiProverServiceManager.CallOpts, arg0)
}

// Delegation is a free data retrieval call binding the contract method 0xdf5cf723.
//
// Solidity: function delegation() view returns(address)
func (_MultiProverServiceManager *MultiProverServiceManagerCaller) Delegation(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MultiProverServiceManager.contract.Call(opts, &out, "delegation")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Delegation is a free data retrieval call binding the contract method 0xdf5cf723.
//
// Solidity: function delegation() view returns(address)
func (_MultiProverServiceManager *MultiProverServiceManagerSession) Delegation() (common.Address, error) {
	return _MultiProverServiceManager.Contract.Delegation(&_MultiProverServiceManager.CallOpts)
}

// Delegation is a free data retrieval call binding the contract method 0xdf5cf723.
//
// Solidity: function delegation() view returns(address)
func (_MultiProverServiceManager *MultiProverServiceManagerCallerSession) Delegation() (common.Address, error) {
	return _MultiProverServiceManager.Contract.Delegation(&_MultiProverServiceManager.CallOpts)
}

// GetOperatorRestakedStrategies is a free data retrieval call binding the contract method 0x33cfb7b7.
//
// Solidity: function getOperatorRestakedStrategies(address operator) view returns(address[])
func (_MultiProverServiceManager *MultiProverServiceManagerCaller) GetOperatorRestakedStrategies(opts *bind.CallOpts, operator common.Address) ([]common.Address, error) {
	var out []interface{}
	err := _MultiProverServiceManager.contract.Call(opts, &out, "getOperatorRestakedStrategies", operator)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetOperatorRestakedStrategies is a free data retrieval call binding the contract method 0x33cfb7b7.
//
// Solidity: function getOperatorRestakedStrategies(address operator) view returns(address[])
func (_MultiProverServiceManager *MultiProverServiceManagerSession) GetOperatorRestakedStrategies(operator common.Address) ([]common.Address, error) {
	return _MultiProverServiceManager.Contract.GetOperatorRestakedStrategies(&_MultiProverServiceManager.CallOpts, operator)
}

// GetOperatorRestakedStrategies is a free data retrieval call binding the contract method 0x33cfb7b7.
//
// Solidity: function getOperatorRestakedStrategies(address operator) view returns(address[])
func (_MultiProverServiceManager *MultiProverServiceManagerCallerSession) GetOperatorRestakedStrategies(operator common.Address) ([]common.Address, error) {
	return _MultiProverServiceManager.Contract.GetOperatorRestakedStrategies(&_MultiProverServiceManager.CallOpts, operator)
}

// GetRestakeableStrategies is a free data retrieval call binding the contract method 0xe481af9d.
//
// Solidity: function getRestakeableStrategies() view returns(address[])
func (_MultiProverServiceManager *MultiProverServiceManagerCaller) GetRestakeableStrategies(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _MultiProverServiceManager.contract.Call(opts, &out, "getRestakeableStrategies")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetRestakeableStrategies is a free data retrieval call binding the contract method 0xe481af9d.
//
// Solidity: function getRestakeableStrategies() view returns(address[])
func (_MultiProverServiceManager *MultiProverServiceManagerSession) GetRestakeableStrategies() ([]common.Address, error) {
	return _MultiProverServiceManager.Contract.GetRestakeableStrategies(&_MultiProverServiceManager.CallOpts)
}

// GetRestakeableStrategies is a free data retrieval call binding the contract method 0xe481af9d.
//
// Solidity: function getRestakeableStrategies() view returns(address[])
func (_MultiProverServiceManager *MultiProverServiceManagerCallerSession) GetRestakeableStrategies() ([]common.Address, error) {
	return _MultiProverServiceManager.Contract.GetRestakeableStrategies(&_MultiProverServiceManager.CallOpts)
}

// IsOperatorWhitelisted is a free data retrieval call binding the contract method 0x2e8da829.
//
// Solidity: function isOperatorWhitelisted(address operator) view returns(bool)
func (_MultiProverServiceManager *MultiProverServiceManagerCaller) IsOperatorWhitelisted(opts *bind.CallOpts, operator common.Address) (bool, error) {
	var out []interface{}
	err := _MultiProverServiceManager.contract.Call(opts, &out, "isOperatorWhitelisted", operator)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsOperatorWhitelisted is a free data retrieval call binding the contract method 0x2e8da829.
//
// Solidity: function isOperatorWhitelisted(address operator) view returns(bool)
func (_MultiProverServiceManager *MultiProverServiceManagerSession) IsOperatorWhitelisted(operator common.Address) (bool, error) {
	return _MultiProverServiceManager.Contract.IsOperatorWhitelisted(&_MultiProverServiceManager.CallOpts, operator)
}

// IsOperatorWhitelisted is a free data retrieval call binding the contract method 0x2e8da829.
//
// Solidity: function isOperatorWhitelisted(address operator) view returns(bool)
func (_MultiProverServiceManager *MultiProverServiceManagerCallerSession) IsOperatorWhitelisted(operator common.Address) (bool, error) {
	return _MultiProverServiceManager.Contract.IsOperatorWhitelisted(&_MultiProverServiceManager.CallOpts, operator)
}

// IsPoAEnabled is a free data retrieval call binding the contract method 0x64330892.
//
// Solidity: function isPoAEnabled() view returns(bool)
func (_MultiProverServiceManager *MultiProverServiceManagerCaller) IsPoAEnabled(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _MultiProverServiceManager.contract.Call(opts, &out, "isPoAEnabled")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsPoAEnabled is a free data retrieval call binding the contract method 0x64330892.
//
// Solidity: function isPoAEnabled() view returns(bool)
func (_MultiProverServiceManager *MultiProverServiceManagerSession) IsPoAEnabled() (bool, error) {
	return _MultiProverServiceManager.Contract.IsPoAEnabled(&_MultiProverServiceManager.CallOpts)
}

// IsPoAEnabled is a free data retrieval call binding the contract method 0x64330892.
//
// Solidity: function isPoAEnabled() view returns(bool)
func (_MultiProverServiceManager *MultiProverServiceManagerCallerSession) IsPoAEnabled() (bool, error) {
	return _MultiProverServiceManager.Contract.IsPoAEnabled(&_MultiProverServiceManager.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_MultiProverServiceManager *MultiProverServiceManagerCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MultiProverServiceManager.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_MultiProverServiceManager *MultiProverServiceManagerSession) Owner() (common.Address, error) {
	return _MultiProverServiceManager.Contract.Owner(&_MultiProverServiceManager.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_MultiProverServiceManager *MultiProverServiceManagerCallerSession) Owner() (common.Address, error) {
	return _MultiProverServiceManager.Contract.Owner(&_MultiProverServiceManager.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5ac86ab7.
//
// Solidity: function paused(uint8 index) view returns(bool)
func (_MultiProverServiceManager *MultiProverServiceManagerCaller) Paused(opts *bind.CallOpts, index uint8) (bool, error) {
	var out []interface{}
	err := _MultiProverServiceManager.contract.Call(opts, &out, "paused", index)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5ac86ab7.
//
// Solidity: function paused(uint8 index) view returns(bool)
func (_MultiProverServiceManager *MultiProverServiceManagerSession) Paused(index uint8) (bool, error) {
	return _MultiProverServiceManager.Contract.Paused(&_MultiProverServiceManager.CallOpts, index)
}

// Paused is a free data retrieval call binding the contract method 0x5ac86ab7.
//
// Solidity: function paused(uint8 index) view returns(bool)
func (_MultiProverServiceManager *MultiProverServiceManagerCallerSession) Paused(index uint8) (bool, error) {
	return _MultiProverServiceManager.Contract.Paused(&_MultiProverServiceManager.CallOpts, index)
}

// Paused0 is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(uint256)
func (_MultiProverServiceManager *MultiProverServiceManagerCaller) Paused0(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _MultiProverServiceManager.contract.Call(opts, &out, "paused0")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Paused0 is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(uint256)
func (_MultiProverServiceManager *MultiProverServiceManagerSession) Paused0() (*big.Int, error) {
	return _MultiProverServiceManager.Contract.Paused0(&_MultiProverServiceManager.CallOpts)
}

// Paused0 is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(uint256)
func (_MultiProverServiceManager *MultiProverServiceManagerCallerSession) Paused0() (*big.Int, error) {
	return _MultiProverServiceManager.Contract.Paused0(&_MultiProverServiceManager.CallOpts)
}

// PauserRegistry is a free data retrieval call binding the contract method 0x886f1195.
//
// Solidity: function pauserRegistry() view returns(address)
func (_MultiProverServiceManager *MultiProverServiceManagerCaller) PauserRegistry(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MultiProverServiceManager.contract.Call(opts, &out, "pauserRegistry")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PauserRegistry is a free data retrieval call binding the contract method 0x886f1195.
//
// Solidity: function pauserRegistry() view returns(address)
func (_MultiProverServiceManager *MultiProverServiceManagerSession) PauserRegistry() (common.Address, error) {
	return _MultiProverServiceManager.Contract.PauserRegistry(&_MultiProverServiceManager.CallOpts)
}

// PauserRegistry is a free data retrieval call binding the contract method 0x886f1195.
//
// Solidity: function pauserRegistry() view returns(address)
func (_MultiProverServiceManager *MultiProverServiceManagerCallerSession) PauserRegistry() (common.Address, error) {
	return _MultiProverServiceManager.Contract.PauserRegistry(&_MultiProverServiceManager.CallOpts)
}

// PoaEnabled is a free data retrieval call binding the contract method 0xf7c089f1.
//
// Solidity: function poaEnabled() view returns(bool)
func (_MultiProverServiceManager *MultiProverServiceManagerCaller) PoaEnabled(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _MultiProverServiceManager.contract.Call(opts, &out, "poaEnabled")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// PoaEnabled is a free data retrieval call binding the contract method 0xf7c089f1.
//
// Solidity: function poaEnabled() view returns(bool)
func (_MultiProverServiceManager *MultiProverServiceManagerSession) PoaEnabled() (bool, error) {
	return _MultiProverServiceManager.Contract.PoaEnabled(&_MultiProverServiceManager.CallOpts)
}

// PoaEnabled is a free data retrieval call binding the contract method 0xf7c089f1.
//
// Solidity: function poaEnabled() view returns(bool)
func (_MultiProverServiceManager *MultiProverServiceManagerCallerSession) PoaEnabled() (bool, error) {
	return _MultiProverServiceManager.Contract.PoaEnabled(&_MultiProverServiceManager.CallOpts)
}

// PoaManager is a free data retrieval call binding the contract method 0x4fbc0f9a.
//
// Solidity: function poaManager() view returns(address)
func (_MultiProverServiceManager *MultiProverServiceManagerCaller) PoaManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MultiProverServiceManager.contract.Call(opts, &out, "poaManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PoaManager is a free data retrieval call binding the contract method 0x4fbc0f9a.
//
// Solidity: function poaManager() view returns(address)
func (_MultiProverServiceManager *MultiProverServiceManagerSession) PoaManager() (common.Address, error) {
	return _MultiProverServiceManager.Contract.PoaManager(&_MultiProverServiceManager.CallOpts)
}

// PoaManager is a free data retrieval call binding the contract method 0x4fbc0f9a.
//
// Solidity: function poaManager() view returns(address)
func (_MultiProverServiceManager *MultiProverServiceManagerCallerSession) PoaManager() (common.Address, error) {
	return _MultiProverServiceManager.Contract.PoaManager(&_MultiProverServiceManager.CallOpts)
}

// QuorumIdToCommitteeId is a free data retrieval call binding the contract method 0xcb6b33b5.
//
// Solidity: function quorumIdToCommitteeId(uint8 ) view returns(uint256)
func (_MultiProverServiceManager *MultiProverServiceManagerCaller) QuorumIdToCommitteeId(opts *bind.CallOpts, arg0 uint8) (*big.Int, error) {
	var out []interface{}
	err := _MultiProverServiceManager.contract.Call(opts, &out, "quorumIdToCommitteeId", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// QuorumIdToCommitteeId is a free data retrieval call binding the contract method 0xcb6b33b5.
//
// Solidity: function quorumIdToCommitteeId(uint8 ) view returns(uint256)
func (_MultiProverServiceManager *MultiProverServiceManagerSession) QuorumIdToCommitteeId(arg0 uint8) (*big.Int, error) {
	return _MultiProverServiceManager.Contract.QuorumIdToCommitteeId(&_MultiProverServiceManager.CallOpts, arg0)
}

// QuorumIdToCommitteeId is a free data retrieval call binding the contract method 0xcb6b33b5.
//
// Solidity: function quorumIdToCommitteeId(uint8 ) view returns(uint256)
func (_MultiProverServiceManager *MultiProverServiceManagerCallerSession) QuorumIdToCommitteeId(arg0 uint8) (*big.Int, error) {
	return _MultiProverServiceManager.Contract.QuorumIdToCommitteeId(&_MultiProverServiceManager.CallOpts, arg0)
}

// RegistryCoordinator is a free data retrieval call binding the contract method 0x6d14a987.
//
// Solidity: function registryCoordinator() view returns(address)
func (_MultiProverServiceManager *MultiProverServiceManagerCaller) RegistryCoordinator(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MultiProverServiceManager.contract.Call(opts, &out, "registryCoordinator")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// RegistryCoordinator is a free data retrieval call binding the contract method 0x6d14a987.
//
// Solidity: function registryCoordinator() view returns(address)
func (_MultiProverServiceManager *MultiProverServiceManagerSession) RegistryCoordinator() (common.Address, error) {
	return _MultiProverServiceManager.Contract.RegistryCoordinator(&_MultiProverServiceManager.CallOpts)
}

// RegistryCoordinator is a free data retrieval call binding the contract method 0x6d14a987.
//
// Solidity: function registryCoordinator() view returns(address)
func (_MultiProverServiceManager *MultiProverServiceManagerCallerSession) RegistryCoordinator() (common.Address, error) {
	return _MultiProverServiceManager.Contract.RegistryCoordinator(&_MultiProverServiceManager.CallOpts)
}

// StakeRegistry is a free data retrieval call binding the contract method 0x68304835.
//
// Solidity: function stakeRegistry() view returns(address)
func (_MultiProverServiceManager *MultiProverServiceManagerCaller) StakeRegistry(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MultiProverServiceManager.contract.Call(opts, &out, "stakeRegistry")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// StakeRegistry is a free data retrieval call binding the contract method 0x68304835.
//
// Solidity: function stakeRegistry() view returns(address)
func (_MultiProverServiceManager *MultiProverServiceManagerSession) StakeRegistry() (common.Address, error) {
	return _MultiProverServiceManager.Contract.StakeRegistry(&_MultiProverServiceManager.CallOpts)
}

// StakeRegistry is a free data retrieval call binding the contract method 0x68304835.
//
// Solidity: function stakeRegistry() view returns(address)
func (_MultiProverServiceManager *MultiProverServiceManagerCallerSession) StakeRegistry() (common.Address, error) {
	return _MultiProverServiceManager.Contract.StakeRegistry(&_MultiProverServiceManager.CallOpts)
}

// StaleStakesForbidden is a free data retrieval call binding the contract method 0xb98d0908.
//
// Solidity: function staleStakesForbidden() view returns(bool)
func (_MultiProverServiceManager *MultiProverServiceManagerCaller) StaleStakesForbidden(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _MultiProverServiceManager.contract.Call(opts, &out, "staleStakesForbidden")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// StaleStakesForbidden is a free data retrieval call binding the contract method 0xb98d0908.
//
// Solidity: function staleStakesForbidden() view returns(bool)
func (_MultiProverServiceManager *MultiProverServiceManagerSession) StaleStakesForbidden() (bool, error) {
	return _MultiProverServiceManager.Contract.StaleStakesForbidden(&_MultiProverServiceManager.CallOpts)
}

// StaleStakesForbidden is a free data retrieval call binding the contract method 0xb98d0908.
//
// Solidity: function staleStakesForbidden() view returns(bool)
func (_MultiProverServiceManager *MultiProverServiceManagerCallerSession) StaleStakesForbidden() (bool, error) {
	return _MultiProverServiceManager.Contract.StaleStakesForbidden(&_MultiProverServiceManager.CallOpts)
}

// StateConfirmer is a free data retrieval call binding the contract method 0x60268489.
//
// Solidity: function stateConfirmer() view returns(address)
func (_MultiProverServiceManager *MultiProverServiceManagerCaller) StateConfirmer(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MultiProverServiceManager.contract.Call(opts, &out, "stateConfirmer")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// StateConfirmer is a free data retrieval call binding the contract method 0x60268489.
//
// Solidity: function stateConfirmer() view returns(address)
func (_MultiProverServiceManager *MultiProverServiceManagerSession) StateConfirmer() (common.Address, error) {
	return _MultiProverServiceManager.Contract.StateConfirmer(&_MultiProverServiceManager.CallOpts)
}

// StateConfirmer is a free data retrieval call binding the contract method 0x60268489.
//
// Solidity: function stateConfirmer() view returns(address)
func (_MultiProverServiceManager *MultiProverServiceManagerCallerSession) StateConfirmer() (common.Address, error) {
	return _MultiProverServiceManager.Contract.StateConfirmer(&_MultiProverServiceManager.CallOpts)
}

// TaskId is a free data retrieval call binding the contract method 0x3322b23d.
//
// Solidity: function taskId() view returns(uint32)
func (_MultiProverServiceManager *MultiProverServiceManagerCaller) TaskId(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _MultiProverServiceManager.contract.Call(opts, &out, "taskId")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// TaskId is a free data retrieval call binding the contract method 0x3322b23d.
//
// Solidity: function taskId() view returns(uint32)
func (_MultiProverServiceManager *MultiProverServiceManagerSession) TaskId() (uint32, error) {
	return _MultiProverServiceManager.Contract.TaskId(&_MultiProverServiceManager.CallOpts)
}

// TaskId is a free data retrieval call binding the contract method 0x3322b23d.
//
// Solidity: function taskId() view returns(uint32)
func (_MultiProverServiceManager *MultiProverServiceManagerCallerSession) TaskId() (uint32, error) {
	return _MultiProverServiceManager.Contract.TaskId(&_MultiProverServiceManager.CallOpts)
}

// TaskIdToMetadataHash is a free data retrieval call binding the contract method 0xf2a5973a.
//
// Solidity: function taskIdToMetadataHash(uint32 ) view returns(bytes32)
func (_MultiProverServiceManager *MultiProverServiceManagerCaller) TaskIdToMetadataHash(opts *bind.CallOpts, arg0 uint32) ([32]byte, error) {
	var out []interface{}
	err := _MultiProverServiceManager.contract.Call(opts, &out, "taskIdToMetadataHash", arg0)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// TaskIdToMetadataHash is a free data retrieval call binding the contract method 0xf2a5973a.
//
// Solidity: function taskIdToMetadataHash(uint32 ) view returns(bytes32)
func (_MultiProverServiceManager *MultiProverServiceManagerSession) TaskIdToMetadataHash(arg0 uint32) ([32]byte, error) {
	return _MultiProverServiceManager.Contract.TaskIdToMetadataHash(&_MultiProverServiceManager.CallOpts, arg0)
}

// TaskIdToMetadataHash is a free data retrieval call binding the contract method 0xf2a5973a.
//
// Solidity: function taskIdToMetadataHash(uint32 ) view returns(bytes32)
func (_MultiProverServiceManager *MultiProverServiceManagerCallerSession) TaskIdToMetadataHash(arg0 uint32) ([32]byte, error) {
	return _MultiProverServiceManager.Contract.TaskIdToMetadataHash(&_MultiProverServiceManager.CallOpts, arg0)
}

// TeeQuorums is a free data retrieval call binding the contract method 0x4b8606e2.
//
// Solidity: function teeQuorums(uint8 ) view returns(uint8 teeType, uint8 quorumNumber)
func (_MultiProverServiceManager *MultiProverServiceManagerCaller) TeeQuorums(opts *bind.CallOpts, arg0 uint8) (struct {
	TeeType      uint8
	QuorumNumber uint8
}, error) {
	var out []interface{}
	err := _MultiProverServiceManager.contract.Call(opts, &out, "teeQuorums", arg0)

	outstruct := new(struct {
		TeeType      uint8
		QuorumNumber uint8
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.TeeType = *abi.ConvertType(out[0], new(uint8)).(*uint8)
	outstruct.QuorumNumber = *abi.ConvertType(out[1], new(uint8)).(*uint8)

	return *outstruct, err

}

// TeeQuorums is a free data retrieval call binding the contract method 0x4b8606e2.
//
// Solidity: function teeQuorums(uint8 ) view returns(uint8 teeType, uint8 quorumNumber)
func (_MultiProverServiceManager *MultiProverServiceManagerSession) TeeQuorums(arg0 uint8) (struct {
	TeeType      uint8
	QuorumNumber uint8
}, error) {
	return _MultiProverServiceManager.Contract.TeeQuorums(&_MultiProverServiceManager.CallOpts, arg0)
}

// TeeQuorums is a free data retrieval call binding the contract method 0x4b8606e2.
//
// Solidity: function teeQuorums(uint8 ) view returns(uint8 teeType, uint8 quorumNumber)
func (_MultiProverServiceManager *MultiProverServiceManagerCallerSession) TeeQuorums(arg0 uint8) (struct {
	TeeType      uint8
	QuorumNumber uint8
}, error) {
	return _MultiProverServiceManager.Contract.TeeQuorums(&_MultiProverServiceManager.CallOpts, arg0)
}

// TrySignatureAndApkVerification is a free data retrieval call binding the contract method 0x171f1d5b.
//
// Solidity: function trySignatureAndApkVerification(bytes32 msgHash, (uint256,uint256) apk, (uint256[2],uint256[2]) apkG2, (uint256,uint256) sigma) view returns(bool pairingSuccessful, bool siganatureIsValid)
func (_MultiProverServiceManager *MultiProverServiceManagerCaller) TrySignatureAndApkVerification(opts *bind.CallOpts, msgHash [32]byte, apk BN254G1Point, apkG2 BN254G2Point, sigma BN254G1Point) (struct {
	PairingSuccessful bool
	SiganatureIsValid bool
}, error) {
	var out []interface{}
	err := _MultiProverServiceManager.contract.Call(opts, &out, "trySignatureAndApkVerification", msgHash, apk, apkG2, sigma)

	outstruct := new(struct {
		PairingSuccessful bool
		SiganatureIsValid bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.PairingSuccessful = *abi.ConvertType(out[0], new(bool)).(*bool)
	outstruct.SiganatureIsValid = *abi.ConvertType(out[1], new(bool)).(*bool)

	return *outstruct, err

}

// TrySignatureAndApkVerification is a free data retrieval call binding the contract method 0x171f1d5b.
//
// Solidity: function trySignatureAndApkVerification(bytes32 msgHash, (uint256,uint256) apk, (uint256[2],uint256[2]) apkG2, (uint256,uint256) sigma) view returns(bool pairingSuccessful, bool siganatureIsValid)
func (_MultiProverServiceManager *MultiProverServiceManagerSession) TrySignatureAndApkVerification(msgHash [32]byte, apk BN254G1Point, apkG2 BN254G2Point, sigma BN254G1Point) (struct {
	PairingSuccessful bool
	SiganatureIsValid bool
}, error) {
	return _MultiProverServiceManager.Contract.TrySignatureAndApkVerification(&_MultiProverServiceManager.CallOpts, msgHash, apk, apkG2, sigma)
}

// TrySignatureAndApkVerification is a free data retrieval call binding the contract method 0x171f1d5b.
//
// Solidity: function trySignatureAndApkVerification(bytes32 msgHash, (uint256,uint256) apk, (uint256[2],uint256[2]) apkG2, (uint256,uint256) sigma) view returns(bool pairingSuccessful, bool siganatureIsValid)
func (_MultiProverServiceManager *MultiProverServiceManagerCallerSession) TrySignatureAndApkVerification(msgHash [32]byte, apk BN254G1Point, apkG2 BN254G2Point, sigma BN254G1Point) (struct {
	PairingSuccessful bool
	SiganatureIsValid bool
}, error) {
	return _MultiProverServiceManager.Contract.TrySignatureAndApkVerification(&_MultiProverServiceManager.CallOpts, msgHash, apk, apkG2, sigma)
}

// AddCommittee is a paid mutator transaction binding the contract method 0x5fd6c531.
//
// Solidity: function addCommittee((uint256,string,bytes,bytes) committee) returns()
func (_MultiProverServiceManager *MultiProverServiceManagerTransactor) AddCommittee(opts *bind.TransactOpts, committee IMultiProverServiceManagerCommittee) (*types.Transaction, error) {
	return _MultiProverServiceManager.contract.Transact(opts, "addCommittee", committee)
}

// AddCommittee is a paid mutator transaction binding the contract method 0x5fd6c531.
//
// Solidity: function addCommittee((uint256,string,bytes,bytes) committee) returns()
func (_MultiProverServiceManager *MultiProverServiceManagerSession) AddCommittee(committee IMultiProverServiceManagerCommittee) (*types.Transaction, error) {
	return _MultiProverServiceManager.Contract.AddCommittee(&_MultiProverServiceManager.TransactOpts, committee)
}

// AddCommittee is a paid mutator transaction binding the contract method 0x5fd6c531.
//
// Solidity: function addCommittee((uint256,string,bytes,bytes) committee) returns()
func (_MultiProverServiceManager *MultiProverServiceManagerTransactorSession) AddCommittee(committee IMultiProverServiceManagerCommittee) (*types.Transaction, error) {
	return _MultiProverServiceManager.Contract.AddCommittee(&_MultiProverServiceManager.TransactOpts, committee)
}

// AddTEEQuorum is a paid mutator transaction binding the contract method 0xbce475dc.
//
// Solidity: function addTEEQuorum((uint8,uint8) teeQuorum) returns()
func (_MultiProverServiceManager *MultiProverServiceManagerTransactor) AddTEEQuorum(opts *bind.TransactOpts, teeQuorum IMultiProverServiceManagerTEEQuorum) (*types.Transaction, error) {
	return _MultiProverServiceManager.contract.Transact(opts, "addTEEQuorum", teeQuorum)
}

// AddTEEQuorum is a paid mutator transaction binding the contract method 0xbce475dc.
//
// Solidity: function addTEEQuorum((uint8,uint8) teeQuorum) returns()
func (_MultiProverServiceManager *MultiProverServiceManagerSession) AddTEEQuorum(teeQuorum IMultiProverServiceManagerTEEQuorum) (*types.Transaction, error) {
	return _MultiProverServiceManager.Contract.AddTEEQuorum(&_MultiProverServiceManager.TransactOpts, teeQuorum)
}

// AddTEEQuorum is a paid mutator transaction binding the contract method 0xbce475dc.
//
// Solidity: function addTEEQuorum((uint8,uint8) teeQuorum) returns()
func (_MultiProverServiceManager *MultiProverServiceManagerTransactorSession) AddTEEQuorum(teeQuorum IMultiProverServiceManagerTEEQuorum) (*types.Transaction, error) {
	return _MultiProverServiceManager.Contract.AddTEEQuorum(&_MultiProverServiceManager.TransactOpts, teeQuorum)
}

// BlacklistOperator is a paid mutator transaction binding the contract method 0xb4492ecc.
//
// Solidity: function blacklistOperator(address operator) returns()
func (_MultiProverServiceManager *MultiProverServiceManagerTransactor) BlacklistOperator(opts *bind.TransactOpts, operator common.Address) (*types.Transaction, error) {
	return _MultiProverServiceManager.contract.Transact(opts, "blacklistOperator", operator)
}

// BlacklistOperator is a paid mutator transaction binding the contract method 0xb4492ecc.
//
// Solidity: function blacklistOperator(address operator) returns()
func (_MultiProverServiceManager *MultiProverServiceManagerSession) BlacklistOperator(operator common.Address) (*types.Transaction, error) {
	return _MultiProverServiceManager.Contract.BlacklistOperator(&_MultiProverServiceManager.TransactOpts, operator)
}

// BlacklistOperator is a paid mutator transaction binding the contract method 0xb4492ecc.
//
// Solidity: function blacklistOperator(address operator) returns()
func (_MultiProverServiceManager *MultiProverServiceManagerTransactorSession) BlacklistOperator(operator common.Address) (*types.Transaction, error) {
	return _MultiProverServiceManager.Contract.BlacklistOperator(&_MultiProverServiceManager.TransactOpts, operator)
}

// ConfirmState is a paid mutator transaction binding the contract method 0x19718835.
//
// Solidity: function confirmState((uint256,bytes,bytes,bytes,bytes,uint32) stateHeader, (uint32[],(uint256,uint256)[],(uint256,uint256)[],(uint256[2],uint256[2]),(uint256,uint256),uint32[],uint32[],uint32[][]) nonSignerStakesAndSignature) returns()
func (_MultiProverServiceManager *MultiProverServiceManagerTransactor) ConfirmState(opts *bind.TransactOpts, stateHeader IMultiProverServiceManagerStateHeader, nonSignerStakesAndSignature IBLSSignatureCheckerNonSignerStakesAndSignature) (*types.Transaction, error) {
	return _MultiProverServiceManager.contract.Transact(opts, "confirmState", stateHeader, nonSignerStakesAndSignature)
}

// ConfirmState is a paid mutator transaction binding the contract method 0x19718835.
//
// Solidity: function confirmState((uint256,bytes,bytes,bytes,bytes,uint32) stateHeader, (uint32[],(uint256,uint256)[],(uint256,uint256)[],(uint256[2],uint256[2]),(uint256,uint256),uint32[],uint32[],uint32[][]) nonSignerStakesAndSignature) returns()
func (_MultiProverServiceManager *MultiProverServiceManagerSession) ConfirmState(stateHeader IMultiProverServiceManagerStateHeader, nonSignerStakesAndSignature IBLSSignatureCheckerNonSignerStakesAndSignature) (*types.Transaction, error) {
	return _MultiProverServiceManager.Contract.ConfirmState(&_MultiProverServiceManager.TransactOpts, stateHeader, nonSignerStakesAndSignature)
}

// ConfirmState is a paid mutator transaction binding the contract method 0x19718835.
//
// Solidity: function confirmState((uint256,bytes,bytes,bytes,bytes,uint32) stateHeader, (uint32[],(uint256,uint256)[],(uint256,uint256)[],(uint256[2],uint256[2]),(uint256,uint256),uint32[],uint32[],uint32[][]) nonSignerStakesAndSignature) returns()
func (_MultiProverServiceManager *MultiProverServiceManagerTransactorSession) ConfirmState(stateHeader IMultiProverServiceManagerStateHeader, nonSignerStakesAndSignature IBLSSignatureCheckerNonSignerStakesAndSignature) (*types.Transaction, error) {
	return _MultiProverServiceManager.Contract.ConfirmState(&_MultiProverServiceManager.TransactOpts, stateHeader, nonSignerStakesAndSignature)
}

// DeregisterOperatorFromAVS is a paid mutator transaction binding the contract method 0xa364f4da.
//
// Solidity: function deregisterOperatorFromAVS(address operator) returns()
func (_MultiProverServiceManager *MultiProverServiceManagerTransactor) DeregisterOperatorFromAVS(opts *bind.TransactOpts, operator common.Address) (*types.Transaction, error) {
	return _MultiProverServiceManager.contract.Transact(opts, "deregisterOperatorFromAVS", operator)
}

// DeregisterOperatorFromAVS is a paid mutator transaction binding the contract method 0xa364f4da.
//
// Solidity: function deregisterOperatorFromAVS(address operator) returns()
func (_MultiProverServiceManager *MultiProverServiceManagerSession) DeregisterOperatorFromAVS(operator common.Address) (*types.Transaction, error) {
	return _MultiProverServiceManager.Contract.DeregisterOperatorFromAVS(&_MultiProverServiceManager.TransactOpts, operator)
}

// DeregisterOperatorFromAVS is a paid mutator transaction binding the contract method 0xa364f4da.
//
// Solidity: function deregisterOperatorFromAVS(address operator) returns()
func (_MultiProverServiceManager *MultiProverServiceManagerTransactorSession) DeregisterOperatorFromAVS(operator common.Address) (*types.Transaction, error) {
	return _MultiProverServiceManager.Contract.DeregisterOperatorFromAVS(&_MultiProverServiceManager.TransactOpts, operator)
}

// DisablePoA is a paid mutator transaction binding the contract method 0x05db3698.
//
// Solidity: function disablePoA() returns()
func (_MultiProverServiceManager *MultiProverServiceManagerTransactor) DisablePoA(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MultiProverServiceManager.contract.Transact(opts, "disablePoA")
}

// DisablePoA is a paid mutator transaction binding the contract method 0x05db3698.
//
// Solidity: function disablePoA() returns()
func (_MultiProverServiceManager *MultiProverServiceManagerSession) DisablePoA() (*types.Transaction, error) {
	return _MultiProverServiceManager.Contract.DisablePoA(&_MultiProverServiceManager.TransactOpts)
}

// DisablePoA is a paid mutator transaction binding the contract method 0x05db3698.
//
// Solidity: function disablePoA() returns()
func (_MultiProverServiceManager *MultiProverServiceManagerTransactorSession) DisablePoA() (*types.Transaction, error) {
	return _MultiProverServiceManager.Contract.DisablePoA(&_MultiProverServiceManager.TransactOpts)
}

// EnablePoA is a paid mutator transaction binding the contract method 0x12f8b905.
//
// Solidity: function enablePoA() returns()
func (_MultiProverServiceManager *MultiProverServiceManagerTransactor) EnablePoA(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MultiProverServiceManager.contract.Transact(opts, "enablePoA")
}

// EnablePoA is a paid mutator transaction binding the contract method 0x12f8b905.
//
// Solidity: function enablePoA() returns()
func (_MultiProverServiceManager *MultiProverServiceManagerSession) EnablePoA() (*types.Transaction, error) {
	return _MultiProverServiceManager.Contract.EnablePoA(&_MultiProverServiceManager.TransactOpts)
}

// EnablePoA is a paid mutator transaction binding the contract method 0x12f8b905.
//
// Solidity: function enablePoA() returns()
func (_MultiProverServiceManager *MultiProverServiceManagerTransactorSession) EnablePoA() (*types.Transaction, error) {
	return _MultiProverServiceManager.Contract.EnablePoA(&_MultiProverServiceManager.TransactOpts)
}

// Initialize is a paid mutator transaction binding the contract method 0x7ac525a9.
//
// Solidity: function initialize(address _pauserRegistry, uint256 _initialPausedStatus, address _initialOwner, address _stateConfirmer, address _poaManager, address _committeeManager, bool _poaEnabled) returns()
func (_MultiProverServiceManager *MultiProverServiceManagerTransactor) Initialize(opts *bind.TransactOpts, _pauserRegistry common.Address, _initialPausedStatus *big.Int, _initialOwner common.Address, _stateConfirmer common.Address, _poaManager common.Address, _committeeManager common.Address, _poaEnabled bool) (*types.Transaction, error) {
	return _MultiProverServiceManager.contract.Transact(opts, "initialize", _pauserRegistry, _initialPausedStatus, _initialOwner, _stateConfirmer, _poaManager, _committeeManager, _poaEnabled)
}

// Initialize is a paid mutator transaction binding the contract method 0x7ac525a9.
//
// Solidity: function initialize(address _pauserRegistry, uint256 _initialPausedStatus, address _initialOwner, address _stateConfirmer, address _poaManager, address _committeeManager, bool _poaEnabled) returns()
func (_MultiProverServiceManager *MultiProverServiceManagerSession) Initialize(_pauserRegistry common.Address, _initialPausedStatus *big.Int, _initialOwner common.Address, _stateConfirmer common.Address, _poaManager common.Address, _committeeManager common.Address, _poaEnabled bool) (*types.Transaction, error) {
	return _MultiProverServiceManager.Contract.Initialize(&_MultiProverServiceManager.TransactOpts, _pauserRegistry, _initialPausedStatus, _initialOwner, _stateConfirmer, _poaManager, _committeeManager, _poaEnabled)
}

// Initialize is a paid mutator transaction binding the contract method 0x7ac525a9.
//
// Solidity: function initialize(address _pauserRegistry, uint256 _initialPausedStatus, address _initialOwner, address _stateConfirmer, address _poaManager, address _committeeManager, bool _poaEnabled) returns()
func (_MultiProverServiceManager *MultiProverServiceManagerTransactorSession) Initialize(_pauserRegistry common.Address, _initialPausedStatus *big.Int, _initialOwner common.Address, _stateConfirmer common.Address, _poaManager common.Address, _committeeManager common.Address, _poaEnabled bool) (*types.Transaction, error) {
	return _MultiProverServiceManager.Contract.Initialize(&_MultiProverServiceManager.TransactOpts, _pauserRegistry, _initialPausedStatus, _initialOwner, _stateConfirmer, _poaManager, _committeeManager, _poaEnabled)
}

// Pause is a paid mutator transaction binding the contract method 0x136439dd.
//
// Solidity: function pause(uint256 newPausedStatus) returns()
func (_MultiProverServiceManager *MultiProverServiceManagerTransactor) Pause(opts *bind.TransactOpts, newPausedStatus *big.Int) (*types.Transaction, error) {
	return _MultiProverServiceManager.contract.Transact(opts, "pause", newPausedStatus)
}

// Pause is a paid mutator transaction binding the contract method 0x136439dd.
//
// Solidity: function pause(uint256 newPausedStatus) returns()
func (_MultiProverServiceManager *MultiProverServiceManagerSession) Pause(newPausedStatus *big.Int) (*types.Transaction, error) {
	return _MultiProverServiceManager.Contract.Pause(&_MultiProverServiceManager.TransactOpts, newPausedStatus)
}

// Pause is a paid mutator transaction binding the contract method 0x136439dd.
//
// Solidity: function pause(uint256 newPausedStatus) returns()
func (_MultiProverServiceManager *MultiProverServiceManagerTransactorSession) Pause(newPausedStatus *big.Int) (*types.Transaction, error) {
	return _MultiProverServiceManager.Contract.Pause(&_MultiProverServiceManager.TransactOpts, newPausedStatus)
}

// PauseAll is a paid mutator transaction binding the contract method 0x595c6a67.
//
// Solidity: function pauseAll() returns()
func (_MultiProverServiceManager *MultiProverServiceManagerTransactor) PauseAll(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MultiProverServiceManager.contract.Transact(opts, "pauseAll")
}

// PauseAll is a paid mutator transaction binding the contract method 0x595c6a67.
//
// Solidity: function pauseAll() returns()
func (_MultiProverServiceManager *MultiProverServiceManagerSession) PauseAll() (*types.Transaction, error) {
	return _MultiProverServiceManager.Contract.PauseAll(&_MultiProverServiceManager.TransactOpts)
}

// PauseAll is a paid mutator transaction binding the contract method 0x595c6a67.
//
// Solidity: function pauseAll() returns()
func (_MultiProverServiceManager *MultiProverServiceManagerTransactorSession) PauseAll() (*types.Transaction, error) {
	return _MultiProverServiceManager.Contract.PauseAll(&_MultiProverServiceManager.TransactOpts)
}

// RegisterOperatorToAVS is a paid mutator transaction binding the contract method 0x9926ee7d.
//
// Solidity: function registerOperatorToAVS(address operator, (bytes,bytes32,uint256) operatorSignature) returns()
func (_MultiProverServiceManager *MultiProverServiceManagerTransactor) RegisterOperatorToAVS(opts *bind.TransactOpts, operator common.Address, operatorSignature ISignatureUtilsSignatureWithSaltAndExpiry) (*types.Transaction, error) {
	return _MultiProverServiceManager.contract.Transact(opts, "registerOperatorToAVS", operator, operatorSignature)
}

// RegisterOperatorToAVS is a paid mutator transaction binding the contract method 0x9926ee7d.
//
// Solidity: function registerOperatorToAVS(address operator, (bytes,bytes32,uint256) operatorSignature) returns()
func (_MultiProverServiceManager *MultiProverServiceManagerSession) RegisterOperatorToAVS(operator common.Address, operatorSignature ISignatureUtilsSignatureWithSaltAndExpiry) (*types.Transaction, error) {
	return _MultiProverServiceManager.Contract.RegisterOperatorToAVS(&_MultiProverServiceManager.TransactOpts, operator, operatorSignature)
}

// RegisterOperatorToAVS is a paid mutator transaction binding the contract method 0x9926ee7d.
//
// Solidity: function registerOperatorToAVS(address operator, (bytes,bytes32,uint256) operatorSignature) returns()
func (_MultiProverServiceManager *MultiProverServiceManagerTransactorSession) RegisterOperatorToAVS(operator common.Address, operatorSignature ISignatureUtilsSignatureWithSaltAndExpiry) (*types.Transaction, error) {
	return _MultiProverServiceManager.Contract.RegisterOperatorToAVS(&_MultiProverServiceManager.TransactOpts, operator, operatorSignature)
}

// RemoveCommittee is a paid mutator transaction binding the contract method 0x38628003.
//
// Solidity: function removeCommittee(uint256 committeeId) returns()
func (_MultiProverServiceManager *MultiProverServiceManagerTransactor) RemoveCommittee(opts *bind.TransactOpts, committeeId *big.Int) (*types.Transaction, error) {
	return _MultiProverServiceManager.contract.Transact(opts, "removeCommittee", committeeId)
}

// RemoveCommittee is a paid mutator transaction binding the contract method 0x38628003.
//
// Solidity: function removeCommittee(uint256 committeeId) returns()
func (_MultiProverServiceManager *MultiProverServiceManagerSession) RemoveCommittee(committeeId *big.Int) (*types.Transaction, error) {
	return _MultiProverServiceManager.Contract.RemoveCommittee(&_MultiProverServiceManager.TransactOpts, committeeId)
}

// RemoveCommittee is a paid mutator transaction binding the contract method 0x38628003.
//
// Solidity: function removeCommittee(uint256 committeeId) returns()
func (_MultiProverServiceManager *MultiProverServiceManagerTransactorSession) RemoveCommittee(committeeId *big.Int) (*types.Transaction, error) {
	return _MultiProverServiceManager.Contract.RemoveCommittee(&_MultiProverServiceManager.TransactOpts, committeeId)
}

// RemoveTEEQuorum is a paid mutator transaction binding the contract method 0x7135aad4.
//
// Solidity: function removeTEEQuorum(uint8 quorumNumber) returns()
func (_MultiProverServiceManager *MultiProverServiceManagerTransactor) RemoveTEEQuorum(opts *bind.TransactOpts, quorumNumber uint8) (*types.Transaction, error) {
	return _MultiProverServiceManager.contract.Transact(opts, "removeTEEQuorum", quorumNumber)
}

// RemoveTEEQuorum is a paid mutator transaction binding the contract method 0x7135aad4.
//
// Solidity: function removeTEEQuorum(uint8 quorumNumber) returns()
func (_MultiProverServiceManager *MultiProverServiceManagerSession) RemoveTEEQuorum(quorumNumber uint8) (*types.Transaction, error) {
	return _MultiProverServiceManager.Contract.RemoveTEEQuorum(&_MultiProverServiceManager.TransactOpts, quorumNumber)
}

// RemoveTEEQuorum is a paid mutator transaction binding the contract method 0x7135aad4.
//
// Solidity: function removeTEEQuorum(uint8 quorumNumber) returns()
func (_MultiProverServiceManager *MultiProverServiceManagerTransactorSession) RemoveTEEQuorum(quorumNumber uint8) (*types.Transaction, error) {
	return _MultiProverServiceManager.Contract.RemoveTEEQuorum(&_MultiProverServiceManager.TransactOpts, quorumNumber)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_MultiProverServiceManager *MultiProverServiceManagerTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MultiProverServiceManager.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_MultiProverServiceManager *MultiProverServiceManagerSession) RenounceOwnership() (*types.Transaction, error) {
	return _MultiProverServiceManager.Contract.RenounceOwnership(&_MultiProverServiceManager.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_MultiProverServiceManager *MultiProverServiceManagerTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _MultiProverServiceManager.Contract.RenounceOwnership(&_MultiProverServiceManager.TransactOpts)
}

// SetCommitteeManager is a paid mutator transaction binding the contract method 0x1f8d48d3.
//
// Solidity: function setCommitteeManager(address _committeeManager) returns()
func (_MultiProverServiceManager *MultiProverServiceManagerTransactor) SetCommitteeManager(opts *bind.TransactOpts, _committeeManager common.Address) (*types.Transaction, error) {
	return _MultiProverServiceManager.contract.Transact(opts, "setCommitteeManager", _committeeManager)
}

// SetCommitteeManager is a paid mutator transaction binding the contract method 0x1f8d48d3.
//
// Solidity: function setCommitteeManager(address _committeeManager) returns()
func (_MultiProverServiceManager *MultiProverServiceManagerSession) SetCommitteeManager(_committeeManager common.Address) (*types.Transaction, error) {
	return _MultiProverServiceManager.Contract.SetCommitteeManager(&_MultiProverServiceManager.TransactOpts, _committeeManager)
}

// SetCommitteeManager is a paid mutator transaction binding the contract method 0x1f8d48d3.
//
// Solidity: function setCommitteeManager(address _committeeManager) returns()
func (_MultiProverServiceManager *MultiProverServiceManagerTransactorSession) SetCommitteeManager(_committeeManager common.Address) (*types.Transaction, error) {
	return _MultiProverServiceManager.Contract.SetCommitteeManager(&_MultiProverServiceManager.TransactOpts, _committeeManager)
}

// SetPauserRegistry is a paid mutator transaction binding the contract method 0x10d67a2f.
//
// Solidity: function setPauserRegistry(address newPauserRegistry) returns()
func (_MultiProverServiceManager *MultiProverServiceManagerTransactor) SetPauserRegistry(opts *bind.TransactOpts, newPauserRegistry common.Address) (*types.Transaction, error) {
	return _MultiProverServiceManager.contract.Transact(opts, "setPauserRegistry", newPauserRegistry)
}

// SetPauserRegistry is a paid mutator transaction binding the contract method 0x10d67a2f.
//
// Solidity: function setPauserRegistry(address newPauserRegistry) returns()
func (_MultiProverServiceManager *MultiProverServiceManagerSession) SetPauserRegistry(newPauserRegistry common.Address) (*types.Transaction, error) {
	return _MultiProverServiceManager.Contract.SetPauserRegistry(&_MultiProverServiceManager.TransactOpts, newPauserRegistry)
}

// SetPauserRegistry is a paid mutator transaction binding the contract method 0x10d67a2f.
//
// Solidity: function setPauserRegistry(address newPauserRegistry) returns()
func (_MultiProverServiceManager *MultiProverServiceManagerTransactorSession) SetPauserRegistry(newPauserRegistry common.Address) (*types.Transaction, error) {
	return _MultiProverServiceManager.Contract.SetPauserRegistry(&_MultiProverServiceManager.TransactOpts, newPauserRegistry)
}

// SetPoAManager is a paid mutator transaction binding the contract method 0x9d71ad88.
//
// Solidity: function setPoAManager(address _poaManager) returns()
func (_MultiProverServiceManager *MultiProverServiceManagerTransactor) SetPoAManager(opts *bind.TransactOpts, _poaManager common.Address) (*types.Transaction, error) {
	return _MultiProverServiceManager.contract.Transact(opts, "setPoAManager", _poaManager)
}

// SetPoAManager is a paid mutator transaction binding the contract method 0x9d71ad88.
//
// Solidity: function setPoAManager(address _poaManager) returns()
func (_MultiProverServiceManager *MultiProverServiceManagerSession) SetPoAManager(_poaManager common.Address) (*types.Transaction, error) {
	return _MultiProverServiceManager.Contract.SetPoAManager(&_MultiProverServiceManager.TransactOpts, _poaManager)
}

// SetPoAManager is a paid mutator transaction binding the contract method 0x9d71ad88.
//
// Solidity: function setPoAManager(address _poaManager) returns()
func (_MultiProverServiceManager *MultiProverServiceManagerTransactorSession) SetPoAManager(_poaManager common.Address) (*types.Transaction, error) {
	return _MultiProverServiceManager.Contract.SetPoAManager(&_MultiProverServiceManager.TransactOpts, _poaManager)
}

// SetStaleStakesForbidden is a paid mutator transaction binding the contract method 0x416c7e5e.
//
// Solidity: function setStaleStakesForbidden(bool value) returns()
func (_MultiProverServiceManager *MultiProverServiceManagerTransactor) SetStaleStakesForbidden(opts *bind.TransactOpts, value bool) (*types.Transaction, error) {
	return _MultiProverServiceManager.contract.Transact(opts, "setStaleStakesForbidden", value)
}

// SetStaleStakesForbidden is a paid mutator transaction binding the contract method 0x416c7e5e.
//
// Solidity: function setStaleStakesForbidden(bool value) returns()
func (_MultiProverServiceManager *MultiProverServiceManagerSession) SetStaleStakesForbidden(value bool) (*types.Transaction, error) {
	return _MultiProverServiceManager.Contract.SetStaleStakesForbidden(&_MultiProverServiceManager.TransactOpts, value)
}

// SetStaleStakesForbidden is a paid mutator transaction binding the contract method 0x416c7e5e.
//
// Solidity: function setStaleStakesForbidden(bool value) returns()
func (_MultiProverServiceManager *MultiProverServiceManagerTransactorSession) SetStaleStakesForbidden(value bool) (*types.Transaction, error) {
	return _MultiProverServiceManager.Contract.SetStaleStakesForbidden(&_MultiProverServiceManager.TransactOpts, value)
}

// SetStateConfirmer is a paid mutator transaction binding the contract method 0x93df6af6.
//
// Solidity: function setStateConfirmer(address _stateConfirmer) returns()
func (_MultiProverServiceManager *MultiProverServiceManagerTransactor) SetStateConfirmer(opts *bind.TransactOpts, _stateConfirmer common.Address) (*types.Transaction, error) {
	return _MultiProverServiceManager.contract.Transact(opts, "setStateConfirmer", _stateConfirmer)
}

// SetStateConfirmer is a paid mutator transaction binding the contract method 0x93df6af6.
//
// Solidity: function setStateConfirmer(address _stateConfirmer) returns()
func (_MultiProverServiceManager *MultiProverServiceManagerSession) SetStateConfirmer(_stateConfirmer common.Address) (*types.Transaction, error) {
	return _MultiProverServiceManager.Contract.SetStateConfirmer(&_MultiProverServiceManager.TransactOpts, _stateConfirmer)
}

// SetStateConfirmer is a paid mutator transaction binding the contract method 0x93df6af6.
//
// Solidity: function setStateConfirmer(address _stateConfirmer) returns()
func (_MultiProverServiceManager *MultiProverServiceManagerTransactorSession) SetStateConfirmer(_stateConfirmer common.Address) (*types.Transaction, error) {
	return _MultiProverServiceManager.Contract.SetStateConfirmer(&_MultiProverServiceManager.TransactOpts, _stateConfirmer)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_MultiProverServiceManager *MultiProverServiceManagerTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _MultiProverServiceManager.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_MultiProverServiceManager *MultiProverServiceManagerSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _MultiProverServiceManager.Contract.TransferOwnership(&_MultiProverServiceManager.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_MultiProverServiceManager *MultiProverServiceManagerTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _MultiProverServiceManager.Contract.TransferOwnership(&_MultiProverServiceManager.TransactOpts, newOwner)
}

// Unpause is a paid mutator transaction binding the contract method 0xfabc1cbc.
//
// Solidity: function unpause(uint256 newPausedStatus) returns()
func (_MultiProverServiceManager *MultiProverServiceManagerTransactor) Unpause(opts *bind.TransactOpts, newPausedStatus *big.Int) (*types.Transaction, error) {
	return _MultiProverServiceManager.contract.Transact(opts, "unpause", newPausedStatus)
}

// Unpause is a paid mutator transaction binding the contract method 0xfabc1cbc.
//
// Solidity: function unpause(uint256 newPausedStatus) returns()
func (_MultiProverServiceManager *MultiProverServiceManagerSession) Unpause(newPausedStatus *big.Int) (*types.Transaction, error) {
	return _MultiProverServiceManager.Contract.Unpause(&_MultiProverServiceManager.TransactOpts, newPausedStatus)
}

// Unpause is a paid mutator transaction binding the contract method 0xfabc1cbc.
//
// Solidity: function unpause(uint256 newPausedStatus) returns()
func (_MultiProverServiceManager *MultiProverServiceManagerTransactorSession) Unpause(newPausedStatus *big.Int) (*types.Transaction, error) {
	return _MultiProverServiceManager.Contract.Unpause(&_MultiProverServiceManager.TransactOpts, newPausedStatus)
}

// UpdateAVSMetadataURI is a paid mutator transaction binding the contract method 0xa98fb355.
//
// Solidity: function updateAVSMetadataURI(string _metadataURI) returns()
func (_MultiProverServiceManager *MultiProverServiceManagerTransactor) UpdateAVSMetadataURI(opts *bind.TransactOpts, _metadataURI string) (*types.Transaction, error) {
	return _MultiProverServiceManager.contract.Transact(opts, "updateAVSMetadataURI", _metadataURI)
}

// UpdateAVSMetadataURI is a paid mutator transaction binding the contract method 0xa98fb355.
//
// Solidity: function updateAVSMetadataURI(string _metadataURI) returns()
func (_MultiProverServiceManager *MultiProverServiceManagerSession) UpdateAVSMetadataURI(_metadataURI string) (*types.Transaction, error) {
	return _MultiProverServiceManager.Contract.UpdateAVSMetadataURI(&_MultiProverServiceManager.TransactOpts, _metadataURI)
}

// UpdateAVSMetadataURI is a paid mutator transaction binding the contract method 0xa98fb355.
//
// Solidity: function updateAVSMetadataURI(string _metadataURI) returns()
func (_MultiProverServiceManager *MultiProverServiceManagerTransactorSession) UpdateAVSMetadataURI(_metadataURI string) (*types.Transaction, error) {
	return _MultiProverServiceManager.Contract.UpdateAVSMetadataURI(&_MultiProverServiceManager.TransactOpts, _metadataURI)
}

// UpdateCommittee is a paid mutator transaction binding the contract method 0xb36487b9.
//
// Solidity: function updateCommittee((uint256,string,bytes,bytes) committee) returns()
func (_MultiProverServiceManager *MultiProverServiceManagerTransactor) UpdateCommittee(opts *bind.TransactOpts, committee IMultiProverServiceManagerCommittee) (*types.Transaction, error) {
	return _MultiProverServiceManager.contract.Transact(opts, "updateCommittee", committee)
}

// UpdateCommittee is a paid mutator transaction binding the contract method 0xb36487b9.
//
// Solidity: function updateCommittee((uint256,string,bytes,bytes) committee) returns()
func (_MultiProverServiceManager *MultiProverServiceManagerSession) UpdateCommittee(committee IMultiProverServiceManagerCommittee) (*types.Transaction, error) {
	return _MultiProverServiceManager.Contract.UpdateCommittee(&_MultiProverServiceManager.TransactOpts, committee)
}

// UpdateCommittee is a paid mutator transaction binding the contract method 0xb36487b9.
//
// Solidity: function updateCommittee((uint256,string,bytes,bytes) committee) returns()
func (_MultiProverServiceManager *MultiProverServiceManagerTransactorSession) UpdateCommittee(committee IMultiProverServiceManagerCommittee) (*types.Transaction, error) {
	return _MultiProverServiceManager.Contract.UpdateCommittee(&_MultiProverServiceManager.TransactOpts, committee)
}

// WhitelistOperator is a paid mutator transaction binding the contract method 0x2a2b0bcd.
//
// Solidity: function whitelistOperator(address operator) returns()
func (_MultiProverServiceManager *MultiProverServiceManagerTransactor) WhitelistOperator(opts *bind.TransactOpts, operator common.Address) (*types.Transaction, error) {
	return _MultiProverServiceManager.contract.Transact(opts, "whitelistOperator", operator)
}

// WhitelistOperator is a paid mutator transaction binding the contract method 0x2a2b0bcd.
//
// Solidity: function whitelistOperator(address operator) returns()
func (_MultiProverServiceManager *MultiProverServiceManagerSession) WhitelistOperator(operator common.Address) (*types.Transaction, error) {
	return _MultiProverServiceManager.Contract.WhitelistOperator(&_MultiProverServiceManager.TransactOpts, operator)
}

// WhitelistOperator is a paid mutator transaction binding the contract method 0x2a2b0bcd.
//
// Solidity: function whitelistOperator(address operator) returns()
func (_MultiProverServiceManager *MultiProverServiceManagerTransactorSession) WhitelistOperator(operator common.Address) (*types.Transaction, error) {
	return _MultiProverServiceManager.Contract.WhitelistOperator(&_MultiProverServiceManager.TransactOpts, operator)
}

// MultiProverServiceManagerCommitteeManagerUpdatedIterator is returned from FilterCommitteeManagerUpdated and is used to iterate over the raw logs and unpacked data for CommitteeManagerUpdated events raised by the MultiProverServiceManager contract.
type MultiProverServiceManagerCommitteeManagerUpdatedIterator struct {
	Event *MultiProverServiceManagerCommitteeManagerUpdated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *MultiProverServiceManagerCommitteeManagerUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MultiProverServiceManagerCommitteeManagerUpdated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(MultiProverServiceManagerCommitteeManagerUpdated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *MultiProverServiceManagerCommitteeManagerUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MultiProverServiceManagerCommitteeManagerUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MultiProverServiceManagerCommitteeManagerUpdated represents a CommitteeManagerUpdated event raised by the MultiProverServiceManager contract.
type MultiProverServiceManagerCommitteeManagerUpdated struct {
	PreviousCommitteeManager common.Address
	CurrentCommitteeManager  common.Address
	Raw                      types.Log // Blockchain specific contextual infos
}

// FilterCommitteeManagerUpdated is a free log retrieval operation binding the contract event 0xb0bb5f603e798b07a529de2cb13eee1f02b35628fe487e1b3cd99adf48e9a0b0.
//
// Solidity: event CommitteeManagerUpdated(address previousCommitteeManager, address currentCommitteeManager)
func (_MultiProverServiceManager *MultiProverServiceManagerFilterer) FilterCommitteeManagerUpdated(opts *bind.FilterOpts) (*MultiProverServiceManagerCommitteeManagerUpdatedIterator, error) {

	logs, sub, err := _MultiProverServiceManager.contract.FilterLogs(opts, "CommitteeManagerUpdated")
	if err != nil {
		return nil, err
	}
	return &MultiProverServiceManagerCommitteeManagerUpdatedIterator{contract: _MultiProverServiceManager.contract, event: "CommitteeManagerUpdated", logs: logs, sub: sub}, nil
}

// WatchCommitteeManagerUpdated is a free log subscription operation binding the contract event 0xb0bb5f603e798b07a529de2cb13eee1f02b35628fe487e1b3cd99adf48e9a0b0.
//
// Solidity: event CommitteeManagerUpdated(address previousCommitteeManager, address currentCommitteeManager)
func (_MultiProverServiceManager *MultiProverServiceManagerFilterer) WatchCommitteeManagerUpdated(opts *bind.WatchOpts, sink chan<- *MultiProverServiceManagerCommitteeManagerUpdated) (event.Subscription, error) {

	logs, sub, err := _MultiProverServiceManager.contract.WatchLogs(opts, "CommitteeManagerUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MultiProverServiceManagerCommitteeManagerUpdated)
				if err := _MultiProverServiceManager.contract.UnpackLog(event, "CommitteeManagerUpdated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseCommitteeManagerUpdated is a log parse operation binding the contract event 0xb0bb5f603e798b07a529de2cb13eee1f02b35628fe487e1b3cd99adf48e9a0b0.
//
// Solidity: event CommitteeManagerUpdated(address previousCommitteeManager, address currentCommitteeManager)
func (_MultiProverServiceManager *MultiProverServiceManagerFilterer) ParseCommitteeManagerUpdated(log types.Log) (*MultiProverServiceManagerCommitteeManagerUpdated, error) {
	event := new(MultiProverServiceManagerCommitteeManagerUpdated)
	if err := _MultiProverServiceManager.contract.UnpackLog(event, "CommitteeManagerUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MultiProverServiceManagerInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the MultiProverServiceManager contract.
type MultiProverServiceManagerInitializedIterator struct {
	Event *MultiProverServiceManagerInitialized // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *MultiProverServiceManagerInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MultiProverServiceManagerInitialized)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(MultiProverServiceManagerInitialized)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *MultiProverServiceManagerInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MultiProverServiceManagerInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MultiProverServiceManagerInitialized represents a Initialized event raised by the MultiProverServiceManager contract.
type MultiProverServiceManagerInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_MultiProverServiceManager *MultiProverServiceManagerFilterer) FilterInitialized(opts *bind.FilterOpts) (*MultiProverServiceManagerInitializedIterator, error) {

	logs, sub, err := _MultiProverServiceManager.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &MultiProverServiceManagerInitializedIterator{contract: _MultiProverServiceManager.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_MultiProverServiceManager *MultiProverServiceManagerFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *MultiProverServiceManagerInitialized) (event.Subscription, error) {

	logs, sub, err := _MultiProverServiceManager.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MultiProverServiceManagerInitialized)
				if err := _MultiProverServiceManager.contract.UnpackLog(event, "Initialized", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseInitialized is a log parse operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_MultiProverServiceManager *MultiProverServiceManagerFilterer) ParseInitialized(log types.Log) (*MultiProverServiceManagerInitialized, error) {
	event := new(MultiProverServiceManagerInitialized)
	if err := _MultiProverServiceManager.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MultiProverServiceManagerOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the MultiProverServiceManager contract.
type MultiProverServiceManagerOwnershipTransferredIterator struct {
	Event *MultiProverServiceManagerOwnershipTransferred // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *MultiProverServiceManagerOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MultiProverServiceManagerOwnershipTransferred)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(MultiProverServiceManagerOwnershipTransferred)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *MultiProverServiceManagerOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MultiProverServiceManagerOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MultiProverServiceManagerOwnershipTransferred represents a OwnershipTransferred event raised by the MultiProverServiceManager contract.
type MultiProverServiceManagerOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_MultiProverServiceManager *MultiProverServiceManagerFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*MultiProverServiceManagerOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _MultiProverServiceManager.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &MultiProverServiceManagerOwnershipTransferredIterator{contract: _MultiProverServiceManager.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_MultiProverServiceManager *MultiProverServiceManagerFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *MultiProverServiceManagerOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _MultiProverServiceManager.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MultiProverServiceManagerOwnershipTransferred)
				if err := _MultiProverServiceManager.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_MultiProverServiceManager *MultiProverServiceManagerFilterer) ParseOwnershipTransferred(log types.Log) (*MultiProverServiceManagerOwnershipTransferred, error) {
	event := new(MultiProverServiceManagerOwnershipTransferred)
	if err := _MultiProverServiceManager.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MultiProverServiceManagerPausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the MultiProverServiceManager contract.
type MultiProverServiceManagerPausedIterator struct {
	Event *MultiProverServiceManagerPaused // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *MultiProverServiceManagerPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MultiProverServiceManagerPaused)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(MultiProverServiceManagerPaused)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *MultiProverServiceManagerPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MultiProverServiceManagerPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MultiProverServiceManagerPaused represents a Paused event raised by the MultiProverServiceManager contract.
type MultiProverServiceManagerPaused struct {
	Account         common.Address
	NewPausedStatus *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0xab40a374bc51de372200a8bc981af8c9ecdc08dfdaef0bb6e09f88f3c616ef3d.
//
// Solidity: event Paused(address indexed account, uint256 newPausedStatus)
func (_MultiProverServiceManager *MultiProverServiceManagerFilterer) FilterPaused(opts *bind.FilterOpts, account []common.Address) (*MultiProverServiceManagerPausedIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _MultiProverServiceManager.contract.FilterLogs(opts, "Paused", accountRule)
	if err != nil {
		return nil, err
	}
	return &MultiProverServiceManagerPausedIterator{contract: _MultiProverServiceManager.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0xab40a374bc51de372200a8bc981af8c9ecdc08dfdaef0bb6e09f88f3c616ef3d.
//
// Solidity: event Paused(address indexed account, uint256 newPausedStatus)
func (_MultiProverServiceManager *MultiProverServiceManagerFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *MultiProverServiceManagerPaused, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _MultiProverServiceManager.contract.WatchLogs(opts, "Paused", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MultiProverServiceManagerPaused)
				if err := _MultiProverServiceManager.contract.UnpackLog(event, "Paused", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParsePaused is a log parse operation binding the contract event 0xab40a374bc51de372200a8bc981af8c9ecdc08dfdaef0bb6e09f88f3c616ef3d.
//
// Solidity: event Paused(address indexed account, uint256 newPausedStatus)
func (_MultiProverServiceManager *MultiProverServiceManagerFilterer) ParsePaused(log types.Log) (*MultiProverServiceManagerPaused, error) {
	event := new(MultiProverServiceManagerPaused)
	if err := _MultiProverServiceManager.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MultiProverServiceManagerPauserRegistrySetIterator is returned from FilterPauserRegistrySet and is used to iterate over the raw logs and unpacked data for PauserRegistrySet events raised by the MultiProverServiceManager contract.
type MultiProverServiceManagerPauserRegistrySetIterator struct {
	Event *MultiProverServiceManagerPauserRegistrySet // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *MultiProverServiceManagerPauserRegistrySetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MultiProverServiceManagerPauserRegistrySet)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(MultiProverServiceManagerPauserRegistrySet)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *MultiProverServiceManagerPauserRegistrySetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MultiProverServiceManagerPauserRegistrySetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MultiProverServiceManagerPauserRegistrySet represents a PauserRegistrySet event raised by the MultiProverServiceManager contract.
type MultiProverServiceManagerPauserRegistrySet struct {
	PauserRegistry    common.Address
	NewPauserRegistry common.Address
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterPauserRegistrySet is a free log retrieval operation binding the contract event 0x6e9fcd539896fca60e8b0f01dd580233e48a6b0f7df013b89ba7f565869acdb6.
//
// Solidity: event PauserRegistrySet(address pauserRegistry, address newPauserRegistry)
func (_MultiProverServiceManager *MultiProverServiceManagerFilterer) FilterPauserRegistrySet(opts *bind.FilterOpts) (*MultiProverServiceManagerPauserRegistrySetIterator, error) {

	logs, sub, err := _MultiProverServiceManager.contract.FilterLogs(opts, "PauserRegistrySet")
	if err != nil {
		return nil, err
	}
	return &MultiProverServiceManagerPauserRegistrySetIterator{contract: _MultiProverServiceManager.contract, event: "PauserRegistrySet", logs: logs, sub: sub}, nil
}

// WatchPauserRegistrySet is a free log subscription operation binding the contract event 0x6e9fcd539896fca60e8b0f01dd580233e48a6b0f7df013b89ba7f565869acdb6.
//
// Solidity: event PauserRegistrySet(address pauserRegistry, address newPauserRegistry)
func (_MultiProverServiceManager *MultiProverServiceManagerFilterer) WatchPauserRegistrySet(opts *bind.WatchOpts, sink chan<- *MultiProverServiceManagerPauserRegistrySet) (event.Subscription, error) {

	logs, sub, err := _MultiProverServiceManager.contract.WatchLogs(opts, "PauserRegistrySet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MultiProverServiceManagerPauserRegistrySet)
				if err := _MultiProverServiceManager.contract.UnpackLog(event, "PauserRegistrySet", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParsePauserRegistrySet is a log parse operation binding the contract event 0x6e9fcd539896fca60e8b0f01dd580233e48a6b0f7df013b89ba7f565869acdb6.
//
// Solidity: event PauserRegistrySet(address pauserRegistry, address newPauserRegistry)
func (_MultiProverServiceManager *MultiProverServiceManagerFilterer) ParsePauserRegistrySet(log types.Log) (*MultiProverServiceManagerPauserRegistrySet, error) {
	event := new(MultiProverServiceManagerPauserRegistrySet)
	if err := _MultiProverServiceManager.contract.UnpackLog(event, "PauserRegistrySet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MultiProverServiceManagerPoAManagerUpdatedIterator is returned from FilterPoAManagerUpdated and is used to iterate over the raw logs and unpacked data for PoAManagerUpdated events raised by the MultiProverServiceManager contract.
type MultiProverServiceManagerPoAManagerUpdatedIterator struct {
	Event *MultiProverServiceManagerPoAManagerUpdated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *MultiProverServiceManagerPoAManagerUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MultiProverServiceManagerPoAManagerUpdated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(MultiProverServiceManagerPoAManagerUpdated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *MultiProverServiceManagerPoAManagerUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MultiProverServiceManagerPoAManagerUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MultiProverServiceManagerPoAManagerUpdated represents a PoAManagerUpdated event raised by the MultiProverServiceManager contract.
type MultiProverServiceManagerPoAManagerUpdated struct {
	PreviousPoaManager common.Address
	CurrentPoaManager  common.Address
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterPoAManagerUpdated is a free log retrieval operation binding the contract event 0x24a5170be7563b6bf7dd80d648dd7f8971f3836fc77d85c9ce0923ee7365e31c.
//
// Solidity: event PoAManagerUpdated(address previousPoaManager, address currentPoaManager)
func (_MultiProverServiceManager *MultiProverServiceManagerFilterer) FilterPoAManagerUpdated(opts *bind.FilterOpts) (*MultiProverServiceManagerPoAManagerUpdatedIterator, error) {

	logs, sub, err := _MultiProverServiceManager.contract.FilterLogs(opts, "PoAManagerUpdated")
	if err != nil {
		return nil, err
	}
	return &MultiProverServiceManagerPoAManagerUpdatedIterator{contract: _MultiProverServiceManager.contract, event: "PoAManagerUpdated", logs: logs, sub: sub}, nil
}

// WatchPoAManagerUpdated is a free log subscription operation binding the contract event 0x24a5170be7563b6bf7dd80d648dd7f8971f3836fc77d85c9ce0923ee7365e31c.
//
// Solidity: event PoAManagerUpdated(address previousPoaManager, address currentPoaManager)
func (_MultiProverServiceManager *MultiProverServiceManagerFilterer) WatchPoAManagerUpdated(opts *bind.WatchOpts, sink chan<- *MultiProverServiceManagerPoAManagerUpdated) (event.Subscription, error) {

	logs, sub, err := _MultiProverServiceManager.contract.WatchLogs(opts, "PoAManagerUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MultiProverServiceManagerPoAManagerUpdated)
				if err := _MultiProverServiceManager.contract.UnpackLog(event, "PoAManagerUpdated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParsePoAManagerUpdated is a log parse operation binding the contract event 0x24a5170be7563b6bf7dd80d648dd7f8971f3836fc77d85c9ce0923ee7365e31c.
//
// Solidity: event PoAManagerUpdated(address previousPoaManager, address currentPoaManager)
func (_MultiProverServiceManager *MultiProverServiceManagerFilterer) ParsePoAManagerUpdated(log types.Log) (*MultiProverServiceManagerPoAManagerUpdated, error) {
	event := new(MultiProverServiceManagerPoAManagerUpdated)
	if err := _MultiProverServiceManager.contract.UnpackLog(event, "PoAManagerUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MultiProverServiceManagerStaleStakesForbiddenUpdateIterator is returned from FilterStaleStakesForbiddenUpdate and is used to iterate over the raw logs and unpacked data for StaleStakesForbiddenUpdate events raised by the MultiProverServiceManager contract.
type MultiProverServiceManagerStaleStakesForbiddenUpdateIterator struct {
	Event *MultiProverServiceManagerStaleStakesForbiddenUpdate // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *MultiProverServiceManagerStaleStakesForbiddenUpdateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MultiProverServiceManagerStaleStakesForbiddenUpdate)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(MultiProverServiceManagerStaleStakesForbiddenUpdate)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *MultiProverServiceManagerStaleStakesForbiddenUpdateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MultiProverServiceManagerStaleStakesForbiddenUpdateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MultiProverServiceManagerStaleStakesForbiddenUpdate represents a StaleStakesForbiddenUpdate event raised by the MultiProverServiceManager contract.
type MultiProverServiceManagerStaleStakesForbiddenUpdate struct {
	Value bool
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterStaleStakesForbiddenUpdate is a free log retrieval operation binding the contract event 0x40e4ed880a29e0f6ddce307457fb75cddf4feef7d3ecb0301bfdf4976a0e2dfc.
//
// Solidity: event StaleStakesForbiddenUpdate(bool value)
func (_MultiProverServiceManager *MultiProverServiceManagerFilterer) FilterStaleStakesForbiddenUpdate(opts *bind.FilterOpts) (*MultiProverServiceManagerStaleStakesForbiddenUpdateIterator, error) {

	logs, sub, err := _MultiProverServiceManager.contract.FilterLogs(opts, "StaleStakesForbiddenUpdate")
	if err != nil {
		return nil, err
	}
	return &MultiProverServiceManagerStaleStakesForbiddenUpdateIterator{contract: _MultiProverServiceManager.contract, event: "StaleStakesForbiddenUpdate", logs: logs, sub: sub}, nil
}

// WatchStaleStakesForbiddenUpdate is a free log subscription operation binding the contract event 0x40e4ed880a29e0f6ddce307457fb75cddf4feef7d3ecb0301bfdf4976a0e2dfc.
//
// Solidity: event StaleStakesForbiddenUpdate(bool value)
func (_MultiProverServiceManager *MultiProverServiceManagerFilterer) WatchStaleStakesForbiddenUpdate(opts *bind.WatchOpts, sink chan<- *MultiProverServiceManagerStaleStakesForbiddenUpdate) (event.Subscription, error) {

	logs, sub, err := _MultiProverServiceManager.contract.WatchLogs(opts, "StaleStakesForbiddenUpdate")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MultiProverServiceManagerStaleStakesForbiddenUpdate)
				if err := _MultiProverServiceManager.contract.UnpackLog(event, "StaleStakesForbiddenUpdate", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseStaleStakesForbiddenUpdate is a log parse operation binding the contract event 0x40e4ed880a29e0f6ddce307457fb75cddf4feef7d3ecb0301bfdf4976a0e2dfc.
//
// Solidity: event StaleStakesForbiddenUpdate(bool value)
func (_MultiProverServiceManager *MultiProverServiceManagerFilterer) ParseStaleStakesForbiddenUpdate(log types.Log) (*MultiProverServiceManagerStaleStakesForbiddenUpdate, error) {
	event := new(MultiProverServiceManagerStaleStakesForbiddenUpdate)
	if err := _MultiProverServiceManager.contract.UnpackLog(event, "StaleStakesForbiddenUpdate", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MultiProverServiceManagerStateConfirmedIterator is returned from FilterStateConfirmed and is used to iterate over the raw logs and unpacked data for StateConfirmed events raised by the MultiProverServiceManager contract.
type MultiProverServiceManagerStateConfirmedIterator struct {
	Event *MultiProverServiceManagerStateConfirmed // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *MultiProverServiceManagerStateConfirmedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MultiProverServiceManagerStateConfirmed)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(MultiProverServiceManagerStateConfirmed)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *MultiProverServiceManagerStateConfirmedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MultiProverServiceManagerStateConfirmedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MultiProverServiceManagerStateConfirmed represents a StateConfirmed event raised by the MultiProverServiceManager contract.
type MultiProverServiceManagerStateConfirmed struct {
	CommitteeId *big.Int
	Metadata    []byte
	State       []byte
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterStateConfirmed is a free log retrieval operation binding the contract event 0xfa10e7f61e3e060beb2a9dab524d6d58b04c90b1ef9ca10367825cf50870e65d.
//
// Solidity: event StateConfirmed(uint256 indexed committeeId, bytes metadata, bytes state)
func (_MultiProverServiceManager *MultiProverServiceManagerFilterer) FilterStateConfirmed(opts *bind.FilterOpts, committeeId []*big.Int) (*MultiProverServiceManagerStateConfirmedIterator, error) {

	var committeeIdRule []interface{}
	for _, committeeIdItem := range committeeId {
		committeeIdRule = append(committeeIdRule, committeeIdItem)
	}

	logs, sub, err := _MultiProverServiceManager.contract.FilterLogs(opts, "StateConfirmed", committeeIdRule)
	if err != nil {
		return nil, err
	}
	return &MultiProverServiceManagerStateConfirmedIterator{contract: _MultiProverServiceManager.contract, event: "StateConfirmed", logs: logs, sub: sub}, nil
}

// WatchStateConfirmed is a free log subscription operation binding the contract event 0xfa10e7f61e3e060beb2a9dab524d6d58b04c90b1ef9ca10367825cf50870e65d.
//
// Solidity: event StateConfirmed(uint256 indexed committeeId, bytes metadata, bytes state)
func (_MultiProverServiceManager *MultiProverServiceManagerFilterer) WatchStateConfirmed(opts *bind.WatchOpts, sink chan<- *MultiProverServiceManagerStateConfirmed, committeeId []*big.Int) (event.Subscription, error) {

	var committeeIdRule []interface{}
	for _, committeeIdItem := range committeeId {
		committeeIdRule = append(committeeIdRule, committeeIdItem)
	}

	logs, sub, err := _MultiProverServiceManager.contract.WatchLogs(opts, "StateConfirmed", committeeIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MultiProverServiceManagerStateConfirmed)
				if err := _MultiProverServiceManager.contract.UnpackLog(event, "StateConfirmed", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseStateConfirmed is a log parse operation binding the contract event 0xfa10e7f61e3e060beb2a9dab524d6d58b04c90b1ef9ca10367825cf50870e65d.
//
// Solidity: event StateConfirmed(uint256 indexed committeeId, bytes metadata, bytes state)
func (_MultiProverServiceManager *MultiProverServiceManagerFilterer) ParseStateConfirmed(log types.Log) (*MultiProverServiceManagerStateConfirmed, error) {
	event := new(MultiProverServiceManagerStateConfirmed)
	if err := _MultiProverServiceManager.contract.UnpackLog(event, "StateConfirmed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MultiProverServiceManagerStateConfirmerUpdatedIterator is returned from FilterStateConfirmerUpdated and is used to iterate over the raw logs and unpacked data for StateConfirmerUpdated events raised by the MultiProverServiceManager contract.
type MultiProverServiceManagerStateConfirmerUpdatedIterator struct {
	Event *MultiProverServiceManagerStateConfirmerUpdated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *MultiProverServiceManagerStateConfirmerUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MultiProverServiceManagerStateConfirmerUpdated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(MultiProverServiceManagerStateConfirmerUpdated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *MultiProverServiceManagerStateConfirmerUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MultiProverServiceManagerStateConfirmerUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MultiProverServiceManagerStateConfirmerUpdated represents a StateConfirmerUpdated event raised by the MultiProverServiceManager contract.
type MultiProverServiceManagerStateConfirmerUpdated struct {
	PreviousConfirmer common.Address
	CurrentConfirmer  common.Address
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterStateConfirmerUpdated is a free log retrieval operation binding the contract event 0xec6e98ff4cc81b828fb5fdd40168f0959b59c3cd95c004ddc2aeb39d44352e98.
//
// Solidity: event StateConfirmerUpdated(address previousConfirmer, address currentConfirmer)
func (_MultiProverServiceManager *MultiProverServiceManagerFilterer) FilterStateConfirmerUpdated(opts *bind.FilterOpts) (*MultiProverServiceManagerStateConfirmerUpdatedIterator, error) {

	logs, sub, err := _MultiProverServiceManager.contract.FilterLogs(opts, "StateConfirmerUpdated")
	if err != nil {
		return nil, err
	}
	return &MultiProverServiceManagerStateConfirmerUpdatedIterator{contract: _MultiProverServiceManager.contract, event: "StateConfirmerUpdated", logs: logs, sub: sub}, nil
}

// WatchStateConfirmerUpdated is a free log subscription operation binding the contract event 0xec6e98ff4cc81b828fb5fdd40168f0959b59c3cd95c004ddc2aeb39d44352e98.
//
// Solidity: event StateConfirmerUpdated(address previousConfirmer, address currentConfirmer)
func (_MultiProverServiceManager *MultiProverServiceManagerFilterer) WatchStateConfirmerUpdated(opts *bind.WatchOpts, sink chan<- *MultiProverServiceManagerStateConfirmerUpdated) (event.Subscription, error) {

	logs, sub, err := _MultiProverServiceManager.contract.WatchLogs(opts, "StateConfirmerUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MultiProverServiceManagerStateConfirmerUpdated)
				if err := _MultiProverServiceManager.contract.UnpackLog(event, "StateConfirmerUpdated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseStateConfirmerUpdated is a log parse operation binding the contract event 0xec6e98ff4cc81b828fb5fdd40168f0959b59c3cd95c004ddc2aeb39d44352e98.
//
// Solidity: event StateConfirmerUpdated(address previousConfirmer, address currentConfirmer)
func (_MultiProverServiceManager *MultiProverServiceManagerFilterer) ParseStateConfirmerUpdated(log types.Log) (*MultiProverServiceManagerStateConfirmerUpdated, error) {
	event := new(MultiProverServiceManagerStateConfirmerUpdated)
	if err := _MultiProverServiceManager.contract.UnpackLog(event, "StateConfirmerUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MultiProverServiceManagerUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the MultiProverServiceManager contract.
type MultiProverServiceManagerUnpausedIterator struct {
	Event *MultiProverServiceManagerUnpaused // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *MultiProverServiceManagerUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MultiProverServiceManagerUnpaused)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(MultiProverServiceManagerUnpaused)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *MultiProverServiceManagerUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MultiProverServiceManagerUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MultiProverServiceManagerUnpaused represents a Unpaused event raised by the MultiProverServiceManager contract.
type MultiProverServiceManagerUnpaused struct {
	Account         common.Address
	NewPausedStatus *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x3582d1828e26bf56bd801502bc021ac0bc8afb57c826e4986b45593c8fad389c.
//
// Solidity: event Unpaused(address indexed account, uint256 newPausedStatus)
func (_MultiProverServiceManager *MultiProverServiceManagerFilterer) FilterUnpaused(opts *bind.FilterOpts, account []common.Address) (*MultiProverServiceManagerUnpausedIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _MultiProverServiceManager.contract.FilterLogs(opts, "Unpaused", accountRule)
	if err != nil {
		return nil, err
	}
	return &MultiProverServiceManagerUnpausedIterator{contract: _MultiProverServiceManager.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x3582d1828e26bf56bd801502bc021ac0bc8afb57c826e4986b45593c8fad389c.
//
// Solidity: event Unpaused(address indexed account, uint256 newPausedStatus)
func (_MultiProverServiceManager *MultiProverServiceManagerFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *MultiProverServiceManagerUnpaused, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _MultiProverServiceManager.contract.WatchLogs(opts, "Unpaused", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MultiProverServiceManagerUnpaused)
				if err := _MultiProverServiceManager.contract.UnpackLog(event, "Unpaused", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseUnpaused is a log parse operation binding the contract event 0x3582d1828e26bf56bd801502bc021ac0bc8afb57c826e4986b45593c8fad389c.
//
// Solidity: event Unpaused(address indexed account, uint256 newPausedStatus)
func (_MultiProverServiceManager *MultiProverServiceManagerFilterer) ParseUnpaused(log types.Log) (*MultiProverServiceManagerUnpaused, error) {
	event := new(MultiProverServiceManagerUnpaused)
	if err := _MultiProverServiceManager.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
