// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package SGXVerifier

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

// SGXVerifierPubkey is an auto generated low-level Go binding around an user-defined struct.
type SGXVerifierPubkey struct {
	X [32]byte
	Y [32]byte
}

// SGXVerifierMetaData contains all meta data concerning the SGXVerifier contract.
var SGXVerifierMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"attestationAddr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_chainId\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"attestValiditySeconds\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"attestedProvers\",\"outputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"x\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"y\",\"type\":\"bytes32\"}],\"internalType\":\"structSGXVerifier.Pubkey\",\"name\":\"pubkey\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"time\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"attestedReports\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"val\",\"type\":\"uint256\"}],\"name\":\"changeAttestValiditySeconds\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_newOwner\",\"type\":\"address\"}],\"name\":\"changeOwner\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"dcapAttestation\",\"outputs\":[{\"internalType\":\"contractIAttestation\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"pubkeyX\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"pubkeyY\",\"type\":\"bytes32\"}],\"name\":\"isProverRegistered\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"layer2ChainId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"report\",\"type\":\"bytes\"}],\"name\":\"register\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"threshold\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"pubkeyX\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"pubkeyY\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"verifyAttestation\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_mrenclave\",\"type\":\"bytes32\"}],\"name\":\"verifyMrEnclave\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_mrsigner\",\"type\":\"bytes32\"}],\"name\":\"verifyMrSigner\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// SGXVerifierABI is the input ABI used to generate the binding from.
// Deprecated: Use SGXVerifierMetaData.ABI instead.
var SGXVerifierABI = SGXVerifierMetaData.ABI

// SGXVerifier is an auto generated Go binding around an Ethereum contract.
type SGXVerifier struct {
	SGXVerifierCaller     // Read-only binding to the contract
	SGXVerifierTransactor // Write-only binding to the contract
	SGXVerifierFilterer   // Log filterer for contract events
}

// SGXVerifierCaller is an auto generated read-only Go binding around an Ethereum contract.
type SGXVerifierCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SGXVerifierTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SGXVerifierTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SGXVerifierFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SGXVerifierFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SGXVerifierSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SGXVerifierSession struct {
	Contract     *SGXVerifier      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SGXVerifierCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SGXVerifierCallerSession struct {
	Contract *SGXVerifierCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// SGXVerifierTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SGXVerifierTransactorSession struct {
	Contract     *SGXVerifierTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// SGXVerifierRaw is an auto generated low-level Go binding around an Ethereum contract.
type SGXVerifierRaw struct {
	Contract *SGXVerifier // Generic contract binding to access the raw methods on
}

// SGXVerifierCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SGXVerifierCallerRaw struct {
	Contract *SGXVerifierCaller // Generic read-only contract binding to access the raw methods on
}

// SGXVerifierTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SGXVerifierTransactorRaw struct {
	Contract *SGXVerifierTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSGXVerifier creates a new instance of SGXVerifier, bound to a specific deployed contract.
func NewSGXVerifier(address common.Address, backend bind.ContractBackend) (*SGXVerifier, error) {
	contract, err := bindSGXVerifier(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SGXVerifier{SGXVerifierCaller: SGXVerifierCaller{contract: contract}, SGXVerifierTransactor: SGXVerifierTransactor{contract: contract}, SGXVerifierFilterer: SGXVerifierFilterer{contract: contract}}, nil
}

// NewSGXVerifierCaller creates a new read-only instance of SGXVerifier, bound to a specific deployed contract.
func NewSGXVerifierCaller(address common.Address, caller bind.ContractCaller) (*SGXVerifierCaller, error) {
	contract, err := bindSGXVerifier(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SGXVerifierCaller{contract: contract}, nil
}

// NewSGXVerifierTransactor creates a new write-only instance of SGXVerifier, bound to a specific deployed contract.
func NewSGXVerifierTransactor(address common.Address, transactor bind.ContractTransactor) (*SGXVerifierTransactor, error) {
	contract, err := bindSGXVerifier(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SGXVerifierTransactor{contract: contract}, nil
}

// NewSGXVerifierFilterer creates a new log filterer instance of SGXVerifier, bound to a specific deployed contract.
func NewSGXVerifierFilterer(address common.Address, filterer bind.ContractFilterer) (*SGXVerifierFilterer, error) {
	contract, err := bindSGXVerifier(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SGXVerifierFilterer{contract: contract}, nil
}

// bindSGXVerifier binds a generic wrapper to an already deployed contract.
func bindSGXVerifier(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := SGXVerifierMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SGXVerifier *SGXVerifierRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SGXVerifier.Contract.SGXVerifierCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SGXVerifier *SGXVerifierRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SGXVerifier.Contract.SGXVerifierTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SGXVerifier *SGXVerifierRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SGXVerifier.Contract.SGXVerifierTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SGXVerifier *SGXVerifierCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SGXVerifier.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SGXVerifier *SGXVerifierTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SGXVerifier.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SGXVerifier *SGXVerifierTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SGXVerifier.Contract.contract.Transact(opts, method, params...)
}

// AttestValiditySeconds is a free data retrieval call binding the contract method 0x95aa79ca.
//
// Solidity: function attestValiditySeconds() view returns(uint256)
func (_SGXVerifier *SGXVerifierCaller) AttestValiditySeconds(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _SGXVerifier.contract.Call(opts, &out, "attestValiditySeconds")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AttestValiditySeconds is a free data retrieval call binding the contract method 0x95aa79ca.
//
// Solidity: function attestValiditySeconds() view returns(uint256)
func (_SGXVerifier *SGXVerifierSession) AttestValiditySeconds() (*big.Int, error) {
	return _SGXVerifier.Contract.AttestValiditySeconds(&_SGXVerifier.CallOpts)
}

// AttestValiditySeconds is a free data retrieval call binding the contract method 0x95aa79ca.
//
// Solidity: function attestValiditySeconds() view returns(uint256)
func (_SGXVerifier *SGXVerifierCallerSession) AttestValiditySeconds() (*big.Int, error) {
	return _SGXVerifier.Contract.AttestValiditySeconds(&_SGXVerifier.CallOpts)
}

// AttestedProvers is a free data retrieval call binding the contract method 0x3b7a3d82.
//
// Solidity: function attestedProvers(bytes32 ) view returns((bytes32,bytes32) pubkey, uint256 time)
func (_SGXVerifier *SGXVerifierCaller) AttestedProvers(opts *bind.CallOpts, arg0 [32]byte) (struct {
	Pubkey SGXVerifierPubkey
	Time   *big.Int
}, error) {
	var out []interface{}
	err := _SGXVerifier.contract.Call(opts, &out, "attestedProvers", arg0)

	outstruct := new(struct {
		Pubkey SGXVerifierPubkey
		Time   *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Pubkey = *abi.ConvertType(out[0], new(SGXVerifierPubkey)).(*SGXVerifierPubkey)
	outstruct.Time = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// AttestedProvers is a free data retrieval call binding the contract method 0x3b7a3d82.
//
// Solidity: function attestedProvers(bytes32 ) view returns((bytes32,bytes32) pubkey, uint256 time)
func (_SGXVerifier *SGXVerifierSession) AttestedProvers(arg0 [32]byte) (struct {
	Pubkey SGXVerifierPubkey
	Time   *big.Int
}, error) {
	return _SGXVerifier.Contract.AttestedProvers(&_SGXVerifier.CallOpts, arg0)
}

// AttestedProvers is a free data retrieval call binding the contract method 0x3b7a3d82.
//
// Solidity: function attestedProvers(bytes32 ) view returns((bytes32,bytes32) pubkey, uint256 time)
func (_SGXVerifier *SGXVerifierCallerSession) AttestedProvers(arg0 [32]byte) (struct {
	Pubkey SGXVerifierPubkey
	Time   *big.Int
}, error) {
	return _SGXVerifier.Contract.AttestedProvers(&_SGXVerifier.CallOpts, arg0)
}

// AttestedReports is a free data retrieval call binding the contract method 0x74e1a553.
//
// Solidity: function attestedReports(bytes32 ) view returns(bool)
func (_SGXVerifier *SGXVerifierCaller) AttestedReports(opts *bind.CallOpts, arg0 [32]byte) (bool, error) {
	var out []interface{}
	err := _SGXVerifier.contract.Call(opts, &out, "attestedReports", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// AttestedReports is a free data retrieval call binding the contract method 0x74e1a553.
//
// Solidity: function attestedReports(bytes32 ) view returns(bool)
func (_SGXVerifier *SGXVerifierSession) AttestedReports(arg0 [32]byte) (bool, error) {
	return _SGXVerifier.Contract.AttestedReports(&_SGXVerifier.CallOpts, arg0)
}

// AttestedReports is a free data retrieval call binding the contract method 0x74e1a553.
//
// Solidity: function attestedReports(bytes32 ) view returns(bool)
func (_SGXVerifier *SGXVerifierCallerSession) AttestedReports(arg0 [32]byte) (bool, error) {
	return _SGXVerifier.Contract.AttestedReports(&_SGXVerifier.CallOpts, arg0)
}

// DcapAttestation is a free data retrieval call binding the contract method 0x48e9b136.
//
// Solidity: function dcapAttestation() view returns(address)
func (_SGXVerifier *SGXVerifierCaller) DcapAttestation(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _SGXVerifier.contract.Call(opts, &out, "dcapAttestation")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// DcapAttestation is a free data retrieval call binding the contract method 0x48e9b136.
//
// Solidity: function dcapAttestation() view returns(address)
func (_SGXVerifier *SGXVerifierSession) DcapAttestation() (common.Address, error) {
	return _SGXVerifier.Contract.DcapAttestation(&_SGXVerifier.CallOpts)
}

// DcapAttestation is a free data retrieval call binding the contract method 0x48e9b136.
//
// Solidity: function dcapAttestation() view returns(address)
func (_SGXVerifier *SGXVerifierCallerSession) DcapAttestation() (common.Address, error) {
	return _SGXVerifier.Contract.DcapAttestation(&_SGXVerifier.CallOpts)
}

// IsProverRegistered is a free data retrieval call binding the contract method 0x9484bf05.
//
// Solidity: function isProverRegistered(bytes32 pubkeyX, bytes32 pubkeyY) view returns(bool)
func (_SGXVerifier *SGXVerifierCaller) IsProverRegistered(opts *bind.CallOpts, pubkeyX [32]byte, pubkeyY [32]byte) (bool, error) {
	var out []interface{}
	err := _SGXVerifier.contract.Call(opts, &out, "isProverRegistered", pubkeyX, pubkeyY)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsProverRegistered is a free data retrieval call binding the contract method 0x9484bf05.
//
// Solidity: function isProverRegistered(bytes32 pubkeyX, bytes32 pubkeyY) view returns(bool)
func (_SGXVerifier *SGXVerifierSession) IsProverRegistered(pubkeyX [32]byte, pubkeyY [32]byte) (bool, error) {
	return _SGXVerifier.Contract.IsProverRegistered(&_SGXVerifier.CallOpts, pubkeyX, pubkeyY)
}

// IsProverRegistered is a free data retrieval call binding the contract method 0x9484bf05.
//
// Solidity: function isProverRegistered(bytes32 pubkeyX, bytes32 pubkeyY) view returns(bool)
func (_SGXVerifier *SGXVerifierCallerSession) IsProverRegistered(pubkeyX [32]byte, pubkeyY [32]byte) (bool, error) {
	return _SGXVerifier.Contract.IsProverRegistered(&_SGXVerifier.CallOpts, pubkeyX, pubkeyY)
}

// Layer2ChainId is a free data retrieval call binding the contract method 0x03c7f4af.
//
// Solidity: function layer2ChainId() view returns(uint256)
func (_SGXVerifier *SGXVerifierCaller) Layer2ChainId(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _SGXVerifier.contract.Call(opts, &out, "layer2ChainId")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Layer2ChainId is a free data retrieval call binding the contract method 0x03c7f4af.
//
// Solidity: function layer2ChainId() view returns(uint256)
func (_SGXVerifier *SGXVerifierSession) Layer2ChainId() (*big.Int, error) {
	return _SGXVerifier.Contract.Layer2ChainId(&_SGXVerifier.CallOpts)
}

// Layer2ChainId is a free data retrieval call binding the contract method 0x03c7f4af.
//
// Solidity: function layer2ChainId() view returns(uint256)
func (_SGXVerifier *SGXVerifierCallerSession) Layer2ChainId() (*big.Int, error) {
	return _SGXVerifier.Contract.Layer2ChainId(&_SGXVerifier.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_SGXVerifier *SGXVerifierCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _SGXVerifier.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_SGXVerifier *SGXVerifierSession) Owner() (common.Address, error) {
	return _SGXVerifier.Contract.Owner(&_SGXVerifier.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_SGXVerifier *SGXVerifierCallerSession) Owner() (common.Address, error) {
	return _SGXVerifier.Contract.Owner(&_SGXVerifier.CallOpts)
}

// Threshold is a free data retrieval call binding the contract method 0x42cde4e8.
//
// Solidity: function threshold() view returns(uint256)
func (_SGXVerifier *SGXVerifierCaller) Threshold(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _SGXVerifier.contract.Call(opts, &out, "threshold")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Threshold is a free data retrieval call binding the contract method 0x42cde4e8.
//
// Solidity: function threshold() view returns(uint256)
func (_SGXVerifier *SGXVerifierSession) Threshold() (*big.Int, error) {
	return _SGXVerifier.Contract.Threshold(&_SGXVerifier.CallOpts)
}

// Threshold is a free data retrieval call binding the contract method 0x42cde4e8.
//
// Solidity: function threshold() view returns(uint256)
func (_SGXVerifier *SGXVerifierCallerSession) Threshold() (*big.Int, error) {
	return _SGXVerifier.Contract.Threshold(&_SGXVerifier.CallOpts)
}

// VerifyAttestation is a free data retrieval call binding the contract method 0xa65542ca.
//
// Solidity: function verifyAttestation(bytes32 pubkeyX, bytes32 pubkeyY, bytes data) view returns(bool)
func (_SGXVerifier *SGXVerifierCaller) VerifyAttestation(opts *bind.CallOpts, pubkeyX [32]byte, pubkeyY [32]byte, data []byte) (bool, error) {
	var out []interface{}
	err := _SGXVerifier.contract.Call(opts, &out, "verifyAttestation", pubkeyX, pubkeyY, data)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// VerifyAttestation is a free data retrieval call binding the contract method 0xa65542ca.
//
// Solidity: function verifyAttestation(bytes32 pubkeyX, bytes32 pubkeyY, bytes data) view returns(bool)
func (_SGXVerifier *SGXVerifierSession) VerifyAttestation(pubkeyX [32]byte, pubkeyY [32]byte, data []byte) (bool, error) {
	return _SGXVerifier.Contract.VerifyAttestation(&_SGXVerifier.CallOpts, pubkeyX, pubkeyY, data)
}

// VerifyAttestation is a free data retrieval call binding the contract method 0xa65542ca.
//
// Solidity: function verifyAttestation(bytes32 pubkeyX, bytes32 pubkeyY, bytes data) view returns(bool)
func (_SGXVerifier *SGXVerifierCallerSession) VerifyAttestation(pubkeyX [32]byte, pubkeyY [32]byte, data []byte) (bool, error) {
	return _SGXVerifier.Contract.VerifyAttestation(&_SGXVerifier.CallOpts, pubkeyX, pubkeyY, data)
}

// VerifyMrEnclave is a free data retrieval call binding the contract method 0x18a98d6f.
//
// Solidity: function verifyMrEnclave(bytes32 _mrenclave) view returns(bool)
func (_SGXVerifier *SGXVerifierCaller) VerifyMrEnclave(opts *bind.CallOpts, _mrenclave [32]byte) (bool, error) {
	var out []interface{}
	err := _SGXVerifier.contract.Call(opts, &out, "verifyMrEnclave", _mrenclave)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// VerifyMrEnclave is a free data retrieval call binding the contract method 0x18a98d6f.
//
// Solidity: function verifyMrEnclave(bytes32 _mrenclave) view returns(bool)
func (_SGXVerifier *SGXVerifierSession) VerifyMrEnclave(_mrenclave [32]byte) (bool, error) {
	return _SGXVerifier.Contract.VerifyMrEnclave(&_SGXVerifier.CallOpts, _mrenclave)
}

// VerifyMrEnclave is a free data retrieval call binding the contract method 0x18a98d6f.
//
// Solidity: function verifyMrEnclave(bytes32 _mrenclave) view returns(bool)
func (_SGXVerifier *SGXVerifierCallerSession) VerifyMrEnclave(_mrenclave [32]byte) (bool, error) {
	return _SGXVerifier.Contract.VerifyMrEnclave(&_SGXVerifier.CallOpts, _mrenclave)
}

// VerifyMrSigner is a free data retrieval call binding the contract method 0xc64b860f.
//
// Solidity: function verifyMrSigner(bytes32 _mrsigner) view returns(bool)
func (_SGXVerifier *SGXVerifierCaller) VerifyMrSigner(opts *bind.CallOpts, _mrsigner [32]byte) (bool, error) {
	var out []interface{}
	err := _SGXVerifier.contract.Call(opts, &out, "verifyMrSigner", _mrsigner)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// VerifyMrSigner is a free data retrieval call binding the contract method 0xc64b860f.
//
// Solidity: function verifyMrSigner(bytes32 _mrsigner) view returns(bool)
func (_SGXVerifier *SGXVerifierSession) VerifyMrSigner(_mrsigner [32]byte) (bool, error) {
	return _SGXVerifier.Contract.VerifyMrSigner(&_SGXVerifier.CallOpts, _mrsigner)
}

// VerifyMrSigner is a free data retrieval call binding the contract method 0xc64b860f.
//
// Solidity: function verifyMrSigner(bytes32 _mrsigner) view returns(bool)
func (_SGXVerifier *SGXVerifierCallerSession) VerifyMrSigner(_mrsigner [32]byte) (bool, error) {
	return _SGXVerifier.Contract.VerifyMrSigner(&_SGXVerifier.CallOpts, _mrsigner)
}

// ChangeAttestValiditySeconds is a paid mutator transaction binding the contract method 0x315aa7d9.
//
// Solidity: function changeAttestValiditySeconds(uint256 val) returns()
func (_SGXVerifier *SGXVerifierTransactor) ChangeAttestValiditySeconds(opts *bind.TransactOpts, val *big.Int) (*types.Transaction, error) {
	return _SGXVerifier.contract.Transact(opts, "changeAttestValiditySeconds", val)
}

// ChangeAttestValiditySeconds is a paid mutator transaction binding the contract method 0x315aa7d9.
//
// Solidity: function changeAttestValiditySeconds(uint256 val) returns()
func (_SGXVerifier *SGXVerifierSession) ChangeAttestValiditySeconds(val *big.Int) (*types.Transaction, error) {
	return _SGXVerifier.Contract.ChangeAttestValiditySeconds(&_SGXVerifier.TransactOpts, val)
}

// ChangeAttestValiditySeconds is a paid mutator transaction binding the contract method 0x315aa7d9.
//
// Solidity: function changeAttestValiditySeconds(uint256 val) returns()
func (_SGXVerifier *SGXVerifierTransactorSession) ChangeAttestValiditySeconds(val *big.Int) (*types.Transaction, error) {
	return _SGXVerifier.Contract.ChangeAttestValiditySeconds(&_SGXVerifier.TransactOpts, val)
}

// ChangeOwner is a paid mutator transaction binding the contract method 0xa6f9dae1.
//
// Solidity: function changeOwner(address _newOwner) returns()
func (_SGXVerifier *SGXVerifierTransactor) ChangeOwner(opts *bind.TransactOpts, _newOwner common.Address) (*types.Transaction, error) {
	return _SGXVerifier.contract.Transact(opts, "changeOwner", _newOwner)
}

// ChangeOwner is a paid mutator transaction binding the contract method 0xa6f9dae1.
//
// Solidity: function changeOwner(address _newOwner) returns()
func (_SGXVerifier *SGXVerifierSession) ChangeOwner(_newOwner common.Address) (*types.Transaction, error) {
	return _SGXVerifier.Contract.ChangeOwner(&_SGXVerifier.TransactOpts, _newOwner)
}

// ChangeOwner is a paid mutator transaction binding the contract method 0xa6f9dae1.
//
// Solidity: function changeOwner(address _newOwner) returns()
func (_SGXVerifier *SGXVerifierTransactorSession) ChangeOwner(_newOwner common.Address) (*types.Transaction, error) {
	return _SGXVerifier.Contract.ChangeOwner(&_SGXVerifier.TransactOpts, _newOwner)
}

// Register is a paid mutator transaction binding the contract method 0x82fbdc9c.
//
// Solidity: function register(bytes report) returns()
func (_SGXVerifier *SGXVerifierTransactor) Register(opts *bind.TransactOpts, report []byte) (*types.Transaction, error) {
	return _SGXVerifier.contract.Transact(opts, "register", report)
}

// Register is a paid mutator transaction binding the contract method 0x82fbdc9c.
//
// Solidity: function register(bytes report) returns()
func (_SGXVerifier *SGXVerifierSession) Register(report []byte) (*types.Transaction, error) {
	return _SGXVerifier.Contract.Register(&_SGXVerifier.TransactOpts, report)
}

// Register is a paid mutator transaction binding the contract method 0x82fbdc9c.
//
// Solidity: function register(bytes report) returns()
func (_SGXVerifier *SGXVerifierTransactorSession) Register(report []byte) (*types.Transaction, error) {
	return _SGXVerifier.Contract.Register(&_SGXVerifier.TransactOpts, report)
}
