// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package BLSApkRegistry

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

// IBLSApkRegistryApkUpdate is an auto generated low-level Go binding around an user-defined struct.
type IBLSApkRegistryApkUpdate struct {
	ApkHash               [24]byte
	UpdateBlockNumber     uint32
	NextUpdateBlockNumber uint32
}

// IBLSApkRegistryPubkeyRegistrationParams is an auto generated low-level Go binding around an user-defined struct.
type IBLSApkRegistryPubkeyRegistrationParams struct {
	PubkeyRegistrationSignature BN254G1Point
	PubkeyG1                    BN254G1Point
	PubkeyG2                    BN254G2Point
}

// BLSApkRegistryMetaData contains all meta data concerning the BLSApkRegistry contract.
var BLSApkRegistryMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"_registryCoordinator\",\"type\":\"address\",\"internalType\":\"contractIRegistryCoordinator\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"apkHistory\",\"inputs\":[{\"name\":\"\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"apkHash\",\"type\":\"bytes24\",\"internalType\":\"bytes24\"},{\"name\":\"updateBlockNumber\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"nextUpdateBlockNumber\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"currentApk\",\"inputs\":[{\"name\":\"\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"outputs\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"deregisterOperator\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"quorumNumbers\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getApk\",\"inputs\":[{\"name\":\"quorumNumber\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structBN254.G1Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getApkHashAtBlockNumberAndIndex\",\"inputs\":[{\"name\":\"quorumNumber\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"blockNumber\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"index\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes24\",\"internalType\":\"bytes24\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getApkHistoryLength\",\"inputs\":[{\"name\":\"quorumNumber\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getApkIndicesAtBlockNumber\",\"inputs\":[{\"name\":\"quorumNumbers\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"blockNumber\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getApkUpdateAtIndex\",\"inputs\":[{\"name\":\"quorumNumber\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"index\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structIBLSApkRegistry.ApkUpdate\",\"components\":[{\"name\":\"apkHash\",\"type\":\"bytes24\",\"internalType\":\"bytes24\"},{\"name\":\"updateBlockNumber\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"nextUpdateBlockNumber\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getOperatorFromPubkeyHash\",\"inputs\":[{\"name\":\"pubkeyHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getOperatorId\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRegisteredPubkey\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structBN254.G1Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"initializeQuorum\",\"inputs\":[{\"name\":\"quorumNumber\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"operatorToPubkey\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"operatorToPubkeyHash\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"pubkeyHashToOperator\",\"inputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"registerBLSPublicKey\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"params\",\"type\":\"tuple\",\"internalType\":\"structIBLSApkRegistry.PubkeyRegistrationParams\",\"components\":[{\"name\":\"pubkeyRegistrationSignature\",\"type\":\"tuple\",\"internalType\":\"structBN254.G1Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"pubkeyG1\",\"type\":\"tuple\",\"internalType\":\"structBN254.G1Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"pubkeyG2\",\"type\":\"tuple\",\"internalType\":\"structBN254.G2Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256[2]\",\"internalType\":\"uint256[2]\"},{\"name\":\"Y\",\"type\":\"uint256[2]\",\"internalType\":\"uint256[2]\"}]}]},{\"name\":\"pubkeyRegistrationMessageHash\",\"type\":\"tuple\",\"internalType\":\"structBN254.G1Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"outputs\":[{\"name\":\"operatorId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"registerOperator\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"quorumNumbers\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"registryCoordinator\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"event\",\"name\":\"Initialized\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"uint8\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"NewPubkeyRegistration\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"pubkeyG1\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structBN254.G1Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"pubkeyG2\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structBN254.G2Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256[2]\",\"internalType\":\"uint256[2]\"},{\"name\":\"Y\",\"type\":\"uint256[2]\",\"internalType\":\"uint256[2]\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OperatorAddedToQuorums\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"operatorId\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"},{\"name\":\"quorumNumbers\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OperatorRemovedFromQuorums\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"operatorId\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"},{\"name\":\"quorumNumbers\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false}]",
}

// BLSApkRegistryABI is the input ABI used to generate the binding from.
// Deprecated: Use BLSApkRegistryMetaData.ABI instead.
var BLSApkRegistryABI = BLSApkRegistryMetaData.ABI

// BLSApkRegistry is an auto generated Go binding around an Ethereum contract.
type BLSApkRegistry struct {
	BLSApkRegistryCaller     // Read-only binding to the contract
	BLSApkRegistryTransactor // Write-only binding to the contract
	BLSApkRegistryFilterer   // Log filterer for contract events
}

// BLSApkRegistryCaller is an auto generated read-only Go binding around an Ethereum contract.
type BLSApkRegistryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BLSApkRegistryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BLSApkRegistryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BLSApkRegistryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BLSApkRegistryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BLSApkRegistrySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BLSApkRegistrySession struct {
	Contract     *BLSApkRegistry   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BLSApkRegistryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BLSApkRegistryCallerSession struct {
	Contract *BLSApkRegistryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// BLSApkRegistryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BLSApkRegistryTransactorSession struct {
	Contract     *BLSApkRegistryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// BLSApkRegistryRaw is an auto generated low-level Go binding around an Ethereum contract.
type BLSApkRegistryRaw struct {
	Contract *BLSApkRegistry // Generic contract binding to access the raw methods on
}

// BLSApkRegistryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BLSApkRegistryCallerRaw struct {
	Contract *BLSApkRegistryCaller // Generic read-only contract binding to access the raw methods on
}

// BLSApkRegistryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BLSApkRegistryTransactorRaw struct {
	Contract *BLSApkRegistryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBLSApkRegistry creates a new instance of BLSApkRegistry, bound to a specific deployed contract.
func NewBLSApkRegistry(address common.Address, backend bind.ContractBackend) (*BLSApkRegistry, error) {
	contract, err := bindBLSApkRegistry(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &BLSApkRegistry{BLSApkRegistryCaller: BLSApkRegistryCaller{contract: contract}, BLSApkRegistryTransactor: BLSApkRegistryTransactor{contract: contract}, BLSApkRegistryFilterer: BLSApkRegistryFilterer{contract: contract}}, nil
}

// NewBLSApkRegistryCaller creates a new read-only instance of BLSApkRegistry, bound to a specific deployed contract.
func NewBLSApkRegistryCaller(address common.Address, caller bind.ContractCaller) (*BLSApkRegistryCaller, error) {
	contract, err := bindBLSApkRegistry(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BLSApkRegistryCaller{contract: contract}, nil
}

// NewBLSApkRegistryTransactor creates a new write-only instance of BLSApkRegistry, bound to a specific deployed contract.
func NewBLSApkRegistryTransactor(address common.Address, transactor bind.ContractTransactor) (*BLSApkRegistryTransactor, error) {
	contract, err := bindBLSApkRegistry(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BLSApkRegistryTransactor{contract: contract}, nil
}

// NewBLSApkRegistryFilterer creates a new log filterer instance of BLSApkRegistry, bound to a specific deployed contract.
func NewBLSApkRegistryFilterer(address common.Address, filterer bind.ContractFilterer) (*BLSApkRegistryFilterer, error) {
	contract, err := bindBLSApkRegistry(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BLSApkRegistryFilterer{contract: contract}, nil
}

// bindBLSApkRegistry binds a generic wrapper to an already deployed contract.
func bindBLSApkRegistry(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := BLSApkRegistryMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BLSApkRegistry *BLSApkRegistryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BLSApkRegistry.Contract.BLSApkRegistryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BLSApkRegistry *BLSApkRegistryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BLSApkRegistry.Contract.BLSApkRegistryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BLSApkRegistry *BLSApkRegistryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BLSApkRegistry.Contract.BLSApkRegistryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BLSApkRegistry *BLSApkRegistryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BLSApkRegistry.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BLSApkRegistry *BLSApkRegistryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BLSApkRegistry.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BLSApkRegistry *BLSApkRegistryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BLSApkRegistry.Contract.contract.Transact(opts, method, params...)
}

// ApkHistory is a free data retrieval call binding the contract method 0x7916cea6.
//
// Solidity: function apkHistory(uint8 , uint256 ) view returns(bytes24 apkHash, uint32 updateBlockNumber, uint32 nextUpdateBlockNumber)
func (_BLSApkRegistry *BLSApkRegistryCaller) ApkHistory(opts *bind.CallOpts, arg0 uint8, arg1 *big.Int) (struct {
	ApkHash               [24]byte
	UpdateBlockNumber     uint32
	NextUpdateBlockNumber uint32
}, error) {
	var out []interface{}
	err := _BLSApkRegistry.contract.Call(opts, &out, "apkHistory", arg0, arg1)

	outstruct := new(struct {
		ApkHash               [24]byte
		UpdateBlockNumber     uint32
		NextUpdateBlockNumber uint32
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.ApkHash = *abi.ConvertType(out[0], new([24]byte)).(*[24]byte)
	outstruct.UpdateBlockNumber = *abi.ConvertType(out[1], new(uint32)).(*uint32)
	outstruct.NextUpdateBlockNumber = *abi.ConvertType(out[2], new(uint32)).(*uint32)

	return *outstruct, err

}

// ApkHistory is a free data retrieval call binding the contract method 0x7916cea6.
//
// Solidity: function apkHistory(uint8 , uint256 ) view returns(bytes24 apkHash, uint32 updateBlockNumber, uint32 nextUpdateBlockNumber)
func (_BLSApkRegistry *BLSApkRegistrySession) ApkHistory(arg0 uint8, arg1 *big.Int) (struct {
	ApkHash               [24]byte
	UpdateBlockNumber     uint32
	NextUpdateBlockNumber uint32
}, error) {
	return _BLSApkRegistry.Contract.ApkHistory(&_BLSApkRegistry.CallOpts, arg0, arg1)
}

// ApkHistory is a free data retrieval call binding the contract method 0x7916cea6.
//
// Solidity: function apkHistory(uint8 , uint256 ) view returns(bytes24 apkHash, uint32 updateBlockNumber, uint32 nextUpdateBlockNumber)
func (_BLSApkRegistry *BLSApkRegistryCallerSession) ApkHistory(arg0 uint8, arg1 *big.Int) (struct {
	ApkHash               [24]byte
	UpdateBlockNumber     uint32
	NextUpdateBlockNumber uint32
}, error) {
	return _BLSApkRegistry.Contract.ApkHistory(&_BLSApkRegistry.CallOpts, arg0, arg1)
}

// CurrentApk is a free data retrieval call binding the contract method 0xa3db80e2.
//
// Solidity: function currentApk(uint8 ) view returns(uint256 X, uint256 Y)
func (_BLSApkRegistry *BLSApkRegistryCaller) CurrentApk(opts *bind.CallOpts, arg0 uint8) (struct {
	X *big.Int
	Y *big.Int
}, error) {
	var out []interface{}
	err := _BLSApkRegistry.contract.Call(opts, &out, "currentApk", arg0)

	outstruct := new(struct {
		X *big.Int
		Y *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.X = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Y = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// CurrentApk is a free data retrieval call binding the contract method 0xa3db80e2.
//
// Solidity: function currentApk(uint8 ) view returns(uint256 X, uint256 Y)
func (_BLSApkRegistry *BLSApkRegistrySession) CurrentApk(arg0 uint8) (struct {
	X *big.Int
	Y *big.Int
}, error) {
	return _BLSApkRegistry.Contract.CurrentApk(&_BLSApkRegistry.CallOpts, arg0)
}

// CurrentApk is a free data retrieval call binding the contract method 0xa3db80e2.
//
// Solidity: function currentApk(uint8 ) view returns(uint256 X, uint256 Y)
func (_BLSApkRegistry *BLSApkRegistryCallerSession) CurrentApk(arg0 uint8) (struct {
	X *big.Int
	Y *big.Int
}, error) {
	return _BLSApkRegistry.Contract.CurrentApk(&_BLSApkRegistry.CallOpts, arg0)
}

// GetApk is a free data retrieval call binding the contract method 0x5f61a884.
//
// Solidity: function getApk(uint8 quorumNumber) view returns((uint256,uint256))
func (_BLSApkRegistry *BLSApkRegistryCaller) GetApk(opts *bind.CallOpts, quorumNumber uint8) (BN254G1Point, error) {
	var out []interface{}
	err := _BLSApkRegistry.contract.Call(opts, &out, "getApk", quorumNumber)

	if err != nil {
		return *new(BN254G1Point), err
	}

	out0 := *abi.ConvertType(out[0], new(BN254G1Point)).(*BN254G1Point)

	return out0, err

}

// GetApk is a free data retrieval call binding the contract method 0x5f61a884.
//
// Solidity: function getApk(uint8 quorumNumber) view returns((uint256,uint256))
func (_BLSApkRegistry *BLSApkRegistrySession) GetApk(quorumNumber uint8) (BN254G1Point, error) {
	return _BLSApkRegistry.Contract.GetApk(&_BLSApkRegistry.CallOpts, quorumNumber)
}

// GetApk is a free data retrieval call binding the contract method 0x5f61a884.
//
// Solidity: function getApk(uint8 quorumNumber) view returns((uint256,uint256))
func (_BLSApkRegistry *BLSApkRegistryCallerSession) GetApk(quorumNumber uint8) (BN254G1Point, error) {
	return _BLSApkRegistry.Contract.GetApk(&_BLSApkRegistry.CallOpts, quorumNumber)
}

// GetApkHashAtBlockNumberAndIndex is a free data retrieval call binding the contract method 0x68bccaac.
//
// Solidity: function getApkHashAtBlockNumberAndIndex(uint8 quorumNumber, uint32 blockNumber, uint256 index) view returns(bytes24)
func (_BLSApkRegistry *BLSApkRegistryCaller) GetApkHashAtBlockNumberAndIndex(opts *bind.CallOpts, quorumNumber uint8, blockNumber uint32, index *big.Int) ([24]byte, error) {
	var out []interface{}
	err := _BLSApkRegistry.contract.Call(opts, &out, "getApkHashAtBlockNumberAndIndex", quorumNumber, blockNumber, index)

	if err != nil {
		return *new([24]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([24]byte)).(*[24]byte)

	return out0, err

}

// GetApkHashAtBlockNumberAndIndex is a free data retrieval call binding the contract method 0x68bccaac.
//
// Solidity: function getApkHashAtBlockNumberAndIndex(uint8 quorumNumber, uint32 blockNumber, uint256 index) view returns(bytes24)
func (_BLSApkRegistry *BLSApkRegistrySession) GetApkHashAtBlockNumberAndIndex(quorumNumber uint8, blockNumber uint32, index *big.Int) ([24]byte, error) {
	return _BLSApkRegistry.Contract.GetApkHashAtBlockNumberAndIndex(&_BLSApkRegistry.CallOpts, quorumNumber, blockNumber, index)
}

// GetApkHashAtBlockNumberAndIndex is a free data retrieval call binding the contract method 0x68bccaac.
//
// Solidity: function getApkHashAtBlockNumberAndIndex(uint8 quorumNumber, uint32 blockNumber, uint256 index) view returns(bytes24)
func (_BLSApkRegistry *BLSApkRegistryCallerSession) GetApkHashAtBlockNumberAndIndex(quorumNumber uint8, blockNumber uint32, index *big.Int) ([24]byte, error) {
	return _BLSApkRegistry.Contract.GetApkHashAtBlockNumberAndIndex(&_BLSApkRegistry.CallOpts, quorumNumber, blockNumber, index)
}

// GetApkHistoryLength is a free data retrieval call binding the contract method 0x377ed99d.
//
// Solidity: function getApkHistoryLength(uint8 quorumNumber) view returns(uint32)
func (_BLSApkRegistry *BLSApkRegistryCaller) GetApkHistoryLength(opts *bind.CallOpts, quorumNumber uint8) (uint32, error) {
	var out []interface{}
	err := _BLSApkRegistry.contract.Call(opts, &out, "getApkHistoryLength", quorumNumber)

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// GetApkHistoryLength is a free data retrieval call binding the contract method 0x377ed99d.
//
// Solidity: function getApkHistoryLength(uint8 quorumNumber) view returns(uint32)
func (_BLSApkRegistry *BLSApkRegistrySession) GetApkHistoryLength(quorumNumber uint8) (uint32, error) {
	return _BLSApkRegistry.Contract.GetApkHistoryLength(&_BLSApkRegistry.CallOpts, quorumNumber)
}

// GetApkHistoryLength is a free data retrieval call binding the contract method 0x377ed99d.
//
// Solidity: function getApkHistoryLength(uint8 quorumNumber) view returns(uint32)
func (_BLSApkRegistry *BLSApkRegistryCallerSession) GetApkHistoryLength(quorumNumber uint8) (uint32, error) {
	return _BLSApkRegistry.Contract.GetApkHistoryLength(&_BLSApkRegistry.CallOpts, quorumNumber)
}

// GetApkIndicesAtBlockNumber is a free data retrieval call binding the contract method 0xd5254a8c.
//
// Solidity: function getApkIndicesAtBlockNumber(bytes quorumNumbers, uint256 blockNumber) view returns(uint32[])
func (_BLSApkRegistry *BLSApkRegistryCaller) GetApkIndicesAtBlockNumber(opts *bind.CallOpts, quorumNumbers []byte, blockNumber *big.Int) ([]uint32, error) {
	var out []interface{}
	err := _BLSApkRegistry.contract.Call(opts, &out, "getApkIndicesAtBlockNumber", quorumNumbers, blockNumber)

	if err != nil {
		return *new([]uint32), err
	}

	out0 := *abi.ConvertType(out[0], new([]uint32)).(*[]uint32)

	return out0, err

}

// GetApkIndicesAtBlockNumber is a free data retrieval call binding the contract method 0xd5254a8c.
//
// Solidity: function getApkIndicesAtBlockNumber(bytes quorumNumbers, uint256 blockNumber) view returns(uint32[])
func (_BLSApkRegistry *BLSApkRegistrySession) GetApkIndicesAtBlockNumber(quorumNumbers []byte, blockNumber *big.Int) ([]uint32, error) {
	return _BLSApkRegistry.Contract.GetApkIndicesAtBlockNumber(&_BLSApkRegistry.CallOpts, quorumNumbers, blockNumber)
}

// GetApkIndicesAtBlockNumber is a free data retrieval call binding the contract method 0xd5254a8c.
//
// Solidity: function getApkIndicesAtBlockNumber(bytes quorumNumbers, uint256 blockNumber) view returns(uint32[])
func (_BLSApkRegistry *BLSApkRegistryCallerSession) GetApkIndicesAtBlockNumber(quorumNumbers []byte, blockNumber *big.Int) ([]uint32, error) {
	return _BLSApkRegistry.Contract.GetApkIndicesAtBlockNumber(&_BLSApkRegistry.CallOpts, quorumNumbers, blockNumber)
}

// GetApkUpdateAtIndex is a free data retrieval call binding the contract method 0x605747d5.
//
// Solidity: function getApkUpdateAtIndex(uint8 quorumNumber, uint256 index) view returns((bytes24,uint32,uint32))
func (_BLSApkRegistry *BLSApkRegistryCaller) GetApkUpdateAtIndex(opts *bind.CallOpts, quorumNumber uint8, index *big.Int) (IBLSApkRegistryApkUpdate, error) {
	var out []interface{}
	err := _BLSApkRegistry.contract.Call(opts, &out, "getApkUpdateAtIndex", quorumNumber, index)

	if err != nil {
		return *new(IBLSApkRegistryApkUpdate), err
	}

	out0 := *abi.ConvertType(out[0], new(IBLSApkRegistryApkUpdate)).(*IBLSApkRegistryApkUpdate)

	return out0, err

}

// GetApkUpdateAtIndex is a free data retrieval call binding the contract method 0x605747d5.
//
// Solidity: function getApkUpdateAtIndex(uint8 quorumNumber, uint256 index) view returns((bytes24,uint32,uint32))
func (_BLSApkRegistry *BLSApkRegistrySession) GetApkUpdateAtIndex(quorumNumber uint8, index *big.Int) (IBLSApkRegistryApkUpdate, error) {
	return _BLSApkRegistry.Contract.GetApkUpdateAtIndex(&_BLSApkRegistry.CallOpts, quorumNumber, index)
}

// GetApkUpdateAtIndex is a free data retrieval call binding the contract method 0x605747d5.
//
// Solidity: function getApkUpdateAtIndex(uint8 quorumNumber, uint256 index) view returns((bytes24,uint32,uint32))
func (_BLSApkRegistry *BLSApkRegistryCallerSession) GetApkUpdateAtIndex(quorumNumber uint8, index *big.Int) (IBLSApkRegistryApkUpdate, error) {
	return _BLSApkRegistry.Contract.GetApkUpdateAtIndex(&_BLSApkRegistry.CallOpts, quorumNumber, index)
}

// GetOperatorFromPubkeyHash is a free data retrieval call binding the contract method 0x47b314e8.
//
// Solidity: function getOperatorFromPubkeyHash(bytes32 pubkeyHash) view returns(address)
func (_BLSApkRegistry *BLSApkRegistryCaller) GetOperatorFromPubkeyHash(opts *bind.CallOpts, pubkeyHash [32]byte) (common.Address, error) {
	var out []interface{}
	err := _BLSApkRegistry.contract.Call(opts, &out, "getOperatorFromPubkeyHash", pubkeyHash)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetOperatorFromPubkeyHash is a free data retrieval call binding the contract method 0x47b314e8.
//
// Solidity: function getOperatorFromPubkeyHash(bytes32 pubkeyHash) view returns(address)
func (_BLSApkRegistry *BLSApkRegistrySession) GetOperatorFromPubkeyHash(pubkeyHash [32]byte) (common.Address, error) {
	return _BLSApkRegistry.Contract.GetOperatorFromPubkeyHash(&_BLSApkRegistry.CallOpts, pubkeyHash)
}

// GetOperatorFromPubkeyHash is a free data retrieval call binding the contract method 0x47b314e8.
//
// Solidity: function getOperatorFromPubkeyHash(bytes32 pubkeyHash) view returns(address)
func (_BLSApkRegistry *BLSApkRegistryCallerSession) GetOperatorFromPubkeyHash(pubkeyHash [32]byte) (common.Address, error) {
	return _BLSApkRegistry.Contract.GetOperatorFromPubkeyHash(&_BLSApkRegistry.CallOpts, pubkeyHash)
}

// GetOperatorId is a free data retrieval call binding the contract method 0x13542a4e.
//
// Solidity: function getOperatorId(address operator) view returns(bytes32)
func (_BLSApkRegistry *BLSApkRegistryCaller) GetOperatorId(opts *bind.CallOpts, operator common.Address) ([32]byte, error) {
	var out []interface{}
	err := _BLSApkRegistry.contract.Call(opts, &out, "getOperatorId", operator)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetOperatorId is a free data retrieval call binding the contract method 0x13542a4e.
//
// Solidity: function getOperatorId(address operator) view returns(bytes32)
func (_BLSApkRegistry *BLSApkRegistrySession) GetOperatorId(operator common.Address) ([32]byte, error) {
	return _BLSApkRegistry.Contract.GetOperatorId(&_BLSApkRegistry.CallOpts, operator)
}

// GetOperatorId is a free data retrieval call binding the contract method 0x13542a4e.
//
// Solidity: function getOperatorId(address operator) view returns(bytes32)
func (_BLSApkRegistry *BLSApkRegistryCallerSession) GetOperatorId(operator common.Address) ([32]byte, error) {
	return _BLSApkRegistry.Contract.GetOperatorId(&_BLSApkRegistry.CallOpts, operator)
}

// GetRegisteredPubkey is a free data retrieval call binding the contract method 0x7ff81a87.
//
// Solidity: function getRegisteredPubkey(address operator) view returns((uint256,uint256), bytes32)
func (_BLSApkRegistry *BLSApkRegistryCaller) GetRegisteredPubkey(opts *bind.CallOpts, operator common.Address) (BN254G1Point, [32]byte, error) {
	var out []interface{}
	err := _BLSApkRegistry.contract.Call(opts, &out, "getRegisteredPubkey", operator)

	if err != nil {
		return *new(BN254G1Point), *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new(BN254G1Point)).(*BN254G1Point)
	out1 := *abi.ConvertType(out[1], new([32]byte)).(*[32]byte)

	return out0, out1, err

}

// GetRegisteredPubkey is a free data retrieval call binding the contract method 0x7ff81a87.
//
// Solidity: function getRegisteredPubkey(address operator) view returns((uint256,uint256), bytes32)
func (_BLSApkRegistry *BLSApkRegistrySession) GetRegisteredPubkey(operator common.Address) (BN254G1Point, [32]byte, error) {
	return _BLSApkRegistry.Contract.GetRegisteredPubkey(&_BLSApkRegistry.CallOpts, operator)
}

// GetRegisteredPubkey is a free data retrieval call binding the contract method 0x7ff81a87.
//
// Solidity: function getRegisteredPubkey(address operator) view returns((uint256,uint256), bytes32)
func (_BLSApkRegistry *BLSApkRegistryCallerSession) GetRegisteredPubkey(operator common.Address) (BN254G1Point, [32]byte, error) {
	return _BLSApkRegistry.Contract.GetRegisteredPubkey(&_BLSApkRegistry.CallOpts, operator)
}

// OperatorToPubkey is a free data retrieval call binding the contract method 0x00a1f4cb.
//
// Solidity: function operatorToPubkey(address ) view returns(uint256 X, uint256 Y)
func (_BLSApkRegistry *BLSApkRegistryCaller) OperatorToPubkey(opts *bind.CallOpts, arg0 common.Address) (struct {
	X *big.Int
	Y *big.Int
}, error) {
	var out []interface{}
	err := _BLSApkRegistry.contract.Call(opts, &out, "operatorToPubkey", arg0)

	outstruct := new(struct {
		X *big.Int
		Y *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.X = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Y = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// OperatorToPubkey is a free data retrieval call binding the contract method 0x00a1f4cb.
//
// Solidity: function operatorToPubkey(address ) view returns(uint256 X, uint256 Y)
func (_BLSApkRegistry *BLSApkRegistrySession) OperatorToPubkey(arg0 common.Address) (struct {
	X *big.Int
	Y *big.Int
}, error) {
	return _BLSApkRegistry.Contract.OperatorToPubkey(&_BLSApkRegistry.CallOpts, arg0)
}

// OperatorToPubkey is a free data retrieval call binding the contract method 0x00a1f4cb.
//
// Solidity: function operatorToPubkey(address ) view returns(uint256 X, uint256 Y)
func (_BLSApkRegistry *BLSApkRegistryCallerSession) OperatorToPubkey(arg0 common.Address) (struct {
	X *big.Int
	Y *big.Int
}, error) {
	return _BLSApkRegistry.Contract.OperatorToPubkey(&_BLSApkRegistry.CallOpts, arg0)
}

// OperatorToPubkeyHash is a free data retrieval call binding the contract method 0xde29fac0.
//
// Solidity: function operatorToPubkeyHash(address ) view returns(bytes32)
func (_BLSApkRegistry *BLSApkRegistryCaller) OperatorToPubkeyHash(opts *bind.CallOpts, arg0 common.Address) ([32]byte, error) {
	var out []interface{}
	err := _BLSApkRegistry.contract.Call(opts, &out, "operatorToPubkeyHash", arg0)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// OperatorToPubkeyHash is a free data retrieval call binding the contract method 0xde29fac0.
//
// Solidity: function operatorToPubkeyHash(address ) view returns(bytes32)
func (_BLSApkRegistry *BLSApkRegistrySession) OperatorToPubkeyHash(arg0 common.Address) ([32]byte, error) {
	return _BLSApkRegistry.Contract.OperatorToPubkeyHash(&_BLSApkRegistry.CallOpts, arg0)
}

// OperatorToPubkeyHash is a free data retrieval call binding the contract method 0xde29fac0.
//
// Solidity: function operatorToPubkeyHash(address ) view returns(bytes32)
func (_BLSApkRegistry *BLSApkRegistryCallerSession) OperatorToPubkeyHash(arg0 common.Address) ([32]byte, error) {
	return _BLSApkRegistry.Contract.OperatorToPubkeyHash(&_BLSApkRegistry.CallOpts, arg0)
}

// PubkeyHashToOperator is a free data retrieval call binding the contract method 0xe8bb9ae6.
//
// Solidity: function pubkeyHashToOperator(bytes32 ) view returns(address)
func (_BLSApkRegistry *BLSApkRegistryCaller) PubkeyHashToOperator(opts *bind.CallOpts, arg0 [32]byte) (common.Address, error) {
	var out []interface{}
	err := _BLSApkRegistry.contract.Call(opts, &out, "pubkeyHashToOperator", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PubkeyHashToOperator is a free data retrieval call binding the contract method 0xe8bb9ae6.
//
// Solidity: function pubkeyHashToOperator(bytes32 ) view returns(address)
func (_BLSApkRegistry *BLSApkRegistrySession) PubkeyHashToOperator(arg0 [32]byte) (common.Address, error) {
	return _BLSApkRegistry.Contract.PubkeyHashToOperator(&_BLSApkRegistry.CallOpts, arg0)
}

// PubkeyHashToOperator is a free data retrieval call binding the contract method 0xe8bb9ae6.
//
// Solidity: function pubkeyHashToOperator(bytes32 ) view returns(address)
func (_BLSApkRegistry *BLSApkRegistryCallerSession) PubkeyHashToOperator(arg0 [32]byte) (common.Address, error) {
	return _BLSApkRegistry.Contract.PubkeyHashToOperator(&_BLSApkRegistry.CallOpts, arg0)
}

// RegistryCoordinator is a free data retrieval call binding the contract method 0x6d14a987.
//
// Solidity: function registryCoordinator() view returns(address)
func (_BLSApkRegistry *BLSApkRegistryCaller) RegistryCoordinator(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BLSApkRegistry.contract.Call(opts, &out, "registryCoordinator")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// RegistryCoordinator is a free data retrieval call binding the contract method 0x6d14a987.
//
// Solidity: function registryCoordinator() view returns(address)
func (_BLSApkRegistry *BLSApkRegistrySession) RegistryCoordinator() (common.Address, error) {
	return _BLSApkRegistry.Contract.RegistryCoordinator(&_BLSApkRegistry.CallOpts)
}

// RegistryCoordinator is a free data retrieval call binding the contract method 0x6d14a987.
//
// Solidity: function registryCoordinator() view returns(address)
func (_BLSApkRegistry *BLSApkRegistryCallerSession) RegistryCoordinator() (common.Address, error) {
	return _BLSApkRegistry.Contract.RegistryCoordinator(&_BLSApkRegistry.CallOpts)
}

// DeregisterOperator is a paid mutator transaction binding the contract method 0xf4e24fe5.
//
// Solidity: function deregisterOperator(address operator, bytes quorumNumbers) returns()
func (_BLSApkRegistry *BLSApkRegistryTransactor) DeregisterOperator(opts *bind.TransactOpts, operator common.Address, quorumNumbers []byte) (*types.Transaction, error) {
	return _BLSApkRegistry.contract.Transact(opts, "deregisterOperator", operator, quorumNumbers)
}

// DeregisterOperator is a paid mutator transaction binding the contract method 0xf4e24fe5.
//
// Solidity: function deregisterOperator(address operator, bytes quorumNumbers) returns()
func (_BLSApkRegistry *BLSApkRegistrySession) DeregisterOperator(operator common.Address, quorumNumbers []byte) (*types.Transaction, error) {
	return _BLSApkRegistry.Contract.DeregisterOperator(&_BLSApkRegistry.TransactOpts, operator, quorumNumbers)
}

// DeregisterOperator is a paid mutator transaction binding the contract method 0xf4e24fe5.
//
// Solidity: function deregisterOperator(address operator, bytes quorumNumbers) returns()
func (_BLSApkRegistry *BLSApkRegistryTransactorSession) DeregisterOperator(operator common.Address, quorumNumbers []byte) (*types.Transaction, error) {
	return _BLSApkRegistry.Contract.DeregisterOperator(&_BLSApkRegistry.TransactOpts, operator, quorumNumbers)
}

// InitializeQuorum is a paid mutator transaction binding the contract method 0x26d941f2.
//
// Solidity: function initializeQuorum(uint8 quorumNumber) returns()
func (_BLSApkRegistry *BLSApkRegistryTransactor) InitializeQuorum(opts *bind.TransactOpts, quorumNumber uint8) (*types.Transaction, error) {
	return _BLSApkRegistry.contract.Transact(opts, "initializeQuorum", quorumNumber)
}

// InitializeQuorum is a paid mutator transaction binding the contract method 0x26d941f2.
//
// Solidity: function initializeQuorum(uint8 quorumNumber) returns()
func (_BLSApkRegistry *BLSApkRegistrySession) InitializeQuorum(quorumNumber uint8) (*types.Transaction, error) {
	return _BLSApkRegistry.Contract.InitializeQuorum(&_BLSApkRegistry.TransactOpts, quorumNumber)
}

// InitializeQuorum is a paid mutator transaction binding the contract method 0x26d941f2.
//
// Solidity: function initializeQuorum(uint8 quorumNumber) returns()
func (_BLSApkRegistry *BLSApkRegistryTransactorSession) InitializeQuorum(quorumNumber uint8) (*types.Transaction, error) {
	return _BLSApkRegistry.Contract.InitializeQuorum(&_BLSApkRegistry.TransactOpts, quorumNumber)
}

// RegisterBLSPublicKey is a paid mutator transaction binding the contract method 0xbf79ce58.
//
// Solidity: function registerBLSPublicKey(address operator, ((uint256,uint256),(uint256,uint256),(uint256[2],uint256[2])) params, (uint256,uint256) pubkeyRegistrationMessageHash) returns(bytes32 operatorId)
func (_BLSApkRegistry *BLSApkRegistryTransactor) RegisterBLSPublicKey(opts *bind.TransactOpts, operator common.Address, params IBLSApkRegistryPubkeyRegistrationParams, pubkeyRegistrationMessageHash BN254G1Point) (*types.Transaction, error) {
	return _BLSApkRegistry.contract.Transact(opts, "registerBLSPublicKey", operator, params, pubkeyRegistrationMessageHash)
}

// RegisterBLSPublicKey is a paid mutator transaction binding the contract method 0xbf79ce58.
//
// Solidity: function registerBLSPublicKey(address operator, ((uint256,uint256),(uint256,uint256),(uint256[2],uint256[2])) params, (uint256,uint256) pubkeyRegistrationMessageHash) returns(bytes32 operatorId)
func (_BLSApkRegistry *BLSApkRegistrySession) RegisterBLSPublicKey(operator common.Address, params IBLSApkRegistryPubkeyRegistrationParams, pubkeyRegistrationMessageHash BN254G1Point) (*types.Transaction, error) {
	return _BLSApkRegistry.Contract.RegisterBLSPublicKey(&_BLSApkRegistry.TransactOpts, operator, params, pubkeyRegistrationMessageHash)
}

// RegisterBLSPublicKey is a paid mutator transaction binding the contract method 0xbf79ce58.
//
// Solidity: function registerBLSPublicKey(address operator, ((uint256,uint256),(uint256,uint256),(uint256[2],uint256[2])) params, (uint256,uint256) pubkeyRegistrationMessageHash) returns(bytes32 operatorId)
func (_BLSApkRegistry *BLSApkRegistryTransactorSession) RegisterBLSPublicKey(operator common.Address, params IBLSApkRegistryPubkeyRegistrationParams, pubkeyRegistrationMessageHash BN254G1Point) (*types.Transaction, error) {
	return _BLSApkRegistry.Contract.RegisterBLSPublicKey(&_BLSApkRegistry.TransactOpts, operator, params, pubkeyRegistrationMessageHash)
}

// RegisterOperator is a paid mutator transaction binding the contract method 0x3fb27952.
//
// Solidity: function registerOperator(address operator, bytes quorumNumbers) returns()
func (_BLSApkRegistry *BLSApkRegistryTransactor) RegisterOperator(opts *bind.TransactOpts, operator common.Address, quorumNumbers []byte) (*types.Transaction, error) {
	return _BLSApkRegistry.contract.Transact(opts, "registerOperator", operator, quorumNumbers)
}

// RegisterOperator is a paid mutator transaction binding the contract method 0x3fb27952.
//
// Solidity: function registerOperator(address operator, bytes quorumNumbers) returns()
func (_BLSApkRegistry *BLSApkRegistrySession) RegisterOperator(operator common.Address, quorumNumbers []byte) (*types.Transaction, error) {
	return _BLSApkRegistry.Contract.RegisterOperator(&_BLSApkRegistry.TransactOpts, operator, quorumNumbers)
}

// RegisterOperator is a paid mutator transaction binding the contract method 0x3fb27952.
//
// Solidity: function registerOperator(address operator, bytes quorumNumbers) returns()
func (_BLSApkRegistry *BLSApkRegistryTransactorSession) RegisterOperator(operator common.Address, quorumNumbers []byte) (*types.Transaction, error) {
	return _BLSApkRegistry.Contract.RegisterOperator(&_BLSApkRegistry.TransactOpts, operator, quorumNumbers)
}

// BLSApkRegistryInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the BLSApkRegistry contract.
type BLSApkRegistryInitializedIterator struct {
	Event *BLSApkRegistryInitialized // Event containing the contract specifics and raw log

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
func (it *BLSApkRegistryInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BLSApkRegistryInitialized)
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
		it.Event = new(BLSApkRegistryInitialized)
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
func (it *BLSApkRegistryInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BLSApkRegistryInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BLSApkRegistryInitialized represents a Initialized event raised by the BLSApkRegistry contract.
type BLSApkRegistryInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_BLSApkRegistry *BLSApkRegistryFilterer) FilterInitialized(opts *bind.FilterOpts) (*BLSApkRegistryInitializedIterator, error) {

	logs, sub, err := _BLSApkRegistry.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &BLSApkRegistryInitializedIterator{contract: _BLSApkRegistry.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_BLSApkRegistry *BLSApkRegistryFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *BLSApkRegistryInitialized) (event.Subscription, error) {

	logs, sub, err := _BLSApkRegistry.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BLSApkRegistryInitialized)
				if err := _BLSApkRegistry.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_BLSApkRegistry *BLSApkRegistryFilterer) ParseInitialized(log types.Log) (*BLSApkRegistryInitialized, error) {
	event := new(BLSApkRegistryInitialized)
	if err := _BLSApkRegistry.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BLSApkRegistryNewPubkeyRegistrationIterator is returned from FilterNewPubkeyRegistration and is used to iterate over the raw logs and unpacked data for NewPubkeyRegistration events raised by the BLSApkRegistry contract.
type BLSApkRegistryNewPubkeyRegistrationIterator struct {
	Event *BLSApkRegistryNewPubkeyRegistration // Event containing the contract specifics and raw log

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
func (it *BLSApkRegistryNewPubkeyRegistrationIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BLSApkRegistryNewPubkeyRegistration)
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
		it.Event = new(BLSApkRegistryNewPubkeyRegistration)
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
func (it *BLSApkRegistryNewPubkeyRegistrationIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BLSApkRegistryNewPubkeyRegistrationIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BLSApkRegistryNewPubkeyRegistration represents a NewPubkeyRegistration event raised by the BLSApkRegistry contract.
type BLSApkRegistryNewPubkeyRegistration struct {
	Operator common.Address
	PubkeyG1 BN254G1Point
	PubkeyG2 BN254G2Point
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterNewPubkeyRegistration is a free log retrieval operation binding the contract event 0xe3fb6613af2e8930cf85d47fcf6db10192224a64c6cbe8023e0eee1ba3828041.
//
// Solidity: event NewPubkeyRegistration(address indexed operator, (uint256,uint256) pubkeyG1, (uint256[2],uint256[2]) pubkeyG2)
func (_BLSApkRegistry *BLSApkRegistryFilterer) FilterNewPubkeyRegistration(opts *bind.FilterOpts, operator []common.Address) (*BLSApkRegistryNewPubkeyRegistrationIterator, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _BLSApkRegistry.contract.FilterLogs(opts, "NewPubkeyRegistration", operatorRule)
	if err != nil {
		return nil, err
	}
	return &BLSApkRegistryNewPubkeyRegistrationIterator{contract: _BLSApkRegistry.contract, event: "NewPubkeyRegistration", logs: logs, sub: sub}, nil
}

// WatchNewPubkeyRegistration is a free log subscription operation binding the contract event 0xe3fb6613af2e8930cf85d47fcf6db10192224a64c6cbe8023e0eee1ba3828041.
//
// Solidity: event NewPubkeyRegistration(address indexed operator, (uint256,uint256) pubkeyG1, (uint256[2],uint256[2]) pubkeyG2)
func (_BLSApkRegistry *BLSApkRegistryFilterer) WatchNewPubkeyRegistration(opts *bind.WatchOpts, sink chan<- *BLSApkRegistryNewPubkeyRegistration, operator []common.Address) (event.Subscription, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _BLSApkRegistry.contract.WatchLogs(opts, "NewPubkeyRegistration", operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BLSApkRegistryNewPubkeyRegistration)
				if err := _BLSApkRegistry.contract.UnpackLog(event, "NewPubkeyRegistration", log); err != nil {
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

// ParseNewPubkeyRegistration is a log parse operation binding the contract event 0xe3fb6613af2e8930cf85d47fcf6db10192224a64c6cbe8023e0eee1ba3828041.
//
// Solidity: event NewPubkeyRegistration(address indexed operator, (uint256,uint256) pubkeyG1, (uint256[2],uint256[2]) pubkeyG2)
func (_BLSApkRegistry *BLSApkRegistryFilterer) ParseNewPubkeyRegistration(log types.Log) (*BLSApkRegistryNewPubkeyRegistration, error) {
	event := new(BLSApkRegistryNewPubkeyRegistration)
	if err := _BLSApkRegistry.contract.UnpackLog(event, "NewPubkeyRegistration", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BLSApkRegistryOperatorAddedToQuorumsIterator is returned from FilterOperatorAddedToQuorums and is used to iterate over the raw logs and unpacked data for OperatorAddedToQuorums events raised by the BLSApkRegistry contract.
type BLSApkRegistryOperatorAddedToQuorumsIterator struct {
	Event *BLSApkRegistryOperatorAddedToQuorums // Event containing the contract specifics and raw log

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
func (it *BLSApkRegistryOperatorAddedToQuorumsIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BLSApkRegistryOperatorAddedToQuorums)
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
		it.Event = new(BLSApkRegistryOperatorAddedToQuorums)
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
func (it *BLSApkRegistryOperatorAddedToQuorumsIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BLSApkRegistryOperatorAddedToQuorumsIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BLSApkRegistryOperatorAddedToQuorums represents a OperatorAddedToQuorums event raised by the BLSApkRegistry contract.
type BLSApkRegistryOperatorAddedToQuorums struct {
	Operator      common.Address
	OperatorId    [32]byte
	QuorumNumbers []byte
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOperatorAddedToQuorums is a free log retrieval operation binding the contract event 0x73a2b7fb844724b971802ae9b15db094d4b7192df9d7350e14eb466b9b22eb4e.
//
// Solidity: event OperatorAddedToQuorums(address operator, bytes32 operatorId, bytes quorumNumbers)
func (_BLSApkRegistry *BLSApkRegistryFilterer) FilterOperatorAddedToQuorums(opts *bind.FilterOpts) (*BLSApkRegistryOperatorAddedToQuorumsIterator, error) {

	logs, sub, err := _BLSApkRegistry.contract.FilterLogs(opts, "OperatorAddedToQuorums")
	if err != nil {
		return nil, err
	}
	return &BLSApkRegistryOperatorAddedToQuorumsIterator{contract: _BLSApkRegistry.contract, event: "OperatorAddedToQuorums", logs: logs, sub: sub}, nil
}

// WatchOperatorAddedToQuorums is a free log subscription operation binding the contract event 0x73a2b7fb844724b971802ae9b15db094d4b7192df9d7350e14eb466b9b22eb4e.
//
// Solidity: event OperatorAddedToQuorums(address operator, bytes32 operatorId, bytes quorumNumbers)
func (_BLSApkRegistry *BLSApkRegistryFilterer) WatchOperatorAddedToQuorums(opts *bind.WatchOpts, sink chan<- *BLSApkRegistryOperatorAddedToQuorums) (event.Subscription, error) {

	logs, sub, err := _BLSApkRegistry.contract.WatchLogs(opts, "OperatorAddedToQuorums")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BLSApkRegistryOperatorAddedToQuorums)
				if err := _BLSApkRegistry.contract.UnpackLog(event, "OperatorAddedToQuorums", log); err != nil {
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

// ParseOperatorAddedToQuorums is a log parse operation binding the contract event 0x73a2b7fb844724b971802ae9b15db094d4b7192df9d7350e14eb466b9b22eb4e.
//
// Solidity: event OperatorAddedToQuorums(address operator, bytes32 operatorId, bytes quorumNumbers)
func (_BLSApkRegistry *BLSApkRegistryFilterer) ParseOperatorAddedToQuorums(log types.Log) (*BLSApkRegistryOperatorAddedToQuorums, error) {
	event := new(BLSApkRegistryOperatorAddedToQuorums)
	if err := _BLSApkRegistry.contract.UnpackLog(event, "OperatorAddedToQuorums", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BLSApkRegistryOperatorRemovedFromQuorumsIterator is returned from FilterOperatorRemovedFromQuorums and is used to iterate over the raw logs and unpacked data for OperatorRemovedFromQuorums events raised by the BLSApkRegistry contract.
type BLSApkRegistryOperatorRemovedFromQuorumsIterator struct {
	Event *BLSApkRegistryOperatorRemovedFromQuorums // Event containing the contract specifics and raw log

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
func (it *BLSApkRegistryOperatorRemovedFromQuorumsIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BLSApkRegistryOperatorRemovedFromQuorums)
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
		it.Event = new(BLSApkRegistryOperatorRemovedFromQuorums)
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
func (it *BLSApkRegistryOperatorRemovedFromQuorumsIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BLSApkRegistryOperatorRemovedFromQuorumsIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BLSApkRegistryOperatorRemovedFromQuorums represents a OperatorRemovedFromQuorums event raised by the BLSApkRegistry contract.
type BLSApkRegistryOperatorRemovedFromQuorums struct {
	Operator      common.Address
	OperatorId    [32]byte
	QuorumNumbers []byte
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOperatorRemovedFromQuorums is a free log retrieval operation binding the contract event 0xf843ecd53a563675e62107be1494fdde4a3d49aeedaf8d88c616d85346e3500e.
//
// Solidity: event OperatorRemovedFromQuorums(address operator, bytes32 operatorId, bytes quorumNumbers)
func (_BLSApkRegistry *BLSApkRegistryFilterer) FilterOperatorRemovedFromQuorums(opts *bind.FilterOpts) (*BLSApkRegistryOperatorRemovedFromQuorumsIterator, error) {

	logs, sub, err := _BLSApkRegistry.contract.FilterLogs(opts, "OperatorRemovedFromQuorums")
	if err != nil {
		return nil, err
	}
	return &BLSApkRegistryOperatorRemovedFromQuorumsIterator{contract: _BLSApkRegistry.contract, event: "OperatorRemovedFromQuorums", logs: logs, sub: sub}, nil
}

// WatchOperatorRemovedFromQuorums is a free log subscription operation binding the contract event 0xf843ecd53a563675e62107be1494fdde4a3d49aeedaf8d88c616d85346e3500e.
//
// Solidity: event OperatorRemovedFromQuorums(address operator, bytes32 operatorId, bytes quorumNumbers)
func (_BLSApkRegistry *BLSApkRegistryFilterer) WatchOperatorRemovedFromQuorums(opts *bind.WatchOpts, sink chan<- *BLSApkRegistryOperatorRemovedFromQuorums) (event.Subscription, error) {

	logs, sub, err := _BLSApkRegistry.contract.WatchLogs(opts, "OperatorRemovedFromQuorums")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BLSApkRegistryOperatorRemovedFromQuorums)
				if err := _BLSApkRegistry.contract.UnpackLog(event, "OperatorRemovedFromQuorums", log); err != nil {
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

// ParseOperatorRemovedFromQuorums is a log parse operation binding the contract event 0xf843ecd53a563675e62107be1494fdde4a3d49aeedaf8d88c616d85346e3500e.
//
// Solidity: event OperatorRemovedFromQuorums(address operator, bytes32 operatorId, bytes quorumNumbers)
func (_BLSApkRegistry *BLSApkRegistryFilterer) ParseOperatorRemovedFromQuorums(log types.Log) (*BLSApkRegistryOperatorRemovedFromQuorums, error) {
	event := new(BLSApkRegistryOperatorRemovedFromQuorums)
	if err := _BLSApkRegistry.contract.UnpackLog(event, "OperatorRemovedFromQuorums", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
