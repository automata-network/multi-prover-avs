// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package TEELivenessVerifier

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

// TEELivenessVerifierPubkey is an auto generated low-level Go binding around an user-defined struct.
type TEELivenessVerifierPubkey struct {
	X [32]byte
	Y [32]byte
}

// TEELivenessVerifierMetaData contains all meta data concerning the TEELivenessVerifier contract.
var TEELivenessVerifierMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"_attestationAddr\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_simulation\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"attestValiditySeconds\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"attestedProvers\",\"inputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"pubkey\",\"type\":\"tuple\",\"internalType\":\"structTEELivenessVerifier.Pubkey\",\"components\":[{\"name\":\"x\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"y\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"name\":\"time\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"attestedReports\",\"inputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"changeAttestValiditySeconds\",\"inputs\":[{\"name\":\"val\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"changeOwner\",\"inputs\":[{\"name\":\"_newOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"dcapAttestation\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIAttestation\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"simulation\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"submitLivenessProof\",\"inputs\":[{\"name\":\"_report\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"verifyAttestation\",\"inputs\":[{\"name\":\"pubkeyX\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"pubkeyY\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"verifyLivenessProof\",\"inputs\":[{\"name\":\"pubkeyX\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"pubkeyY\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"verifyMrEnclave\",\"inputs\":[{\"name\":\"_mrenclave\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"verifyMrSigner\",\"inputs\":[{\"name\":\"_mrsigner\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"}]",
}

// TEELivenessVerifierABI is the input ABI used to generate the binding from.
// Deprecated: Use TEELivenessVerifierMetaData.ABI instead.
var TEELivenessVerifierABI = TEELivenessVerifierMetaData.ABI

// TEELivenessVerifier is an auto generated Go binding around an Ethereum contract.
type TEELivenessVerifier struct {
	TEELivenessVerifierCaller     // Read-only binding to the contract
	TEELivenessVerifierTransactor // Write-only binding to the contract
	TEELivenessVerifierFilterer   // Log filterer for contract events
}

// TEELivenessVerifierCaller is an auto generated read-only Go binding around an Ethereum contract.
type TEELivenessVerifierCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TEELivenessVerifierTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TEELivenessVerifierTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TEELivenessVerifierFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TEELivenessVerifierFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TEELivenessVerifierSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TEELivenessVerifierSession struct {
	Contract     *TEELivenessVerifier // Generic contract binding to set the session for
	CallOpts     bind.CallOpts        // Call options to use throughout this session
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// TEELivenessVerifierCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TEELivenessVerifierCallerSession struct {
	Contract *TEELivenessVerifierCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts              // Call options to use throughout this session
}

// TEELivenessVerifierTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TEELivenessVerifierTransactorSession struct {
	Contract     *TEELivenessVerifierTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts              // Transaction auth options to use throughout this session
}

// TEELivenessVerifierRaw is an auto generated low-level Go binding around an Ethereum contract.
type TEELivenessVerifierRaw struct {
	Contract *TEELivenessVerifier // Generic contract binding to access the raw methods on
}

// TEELivenessVerifierCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TEELivenessVerifierCallerRaw struct {
	Contract *TEELivenessVerifierCaller // Generic read-only contract binding to access the raw methods on
}

// TEELivenessVerifierTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TEELivenessVerifierTransactorRaw struct {
	Contract *TEELivenessVerifierTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTEELivenessVerifier creates a new instance of TEELivenessVerifier, bound to a specific deployed contract.
func NewTEELivenessVerifier(address common.Address, backend bind.ContractBackend) (*TEELivenessVerifier, error) {
	contract, err := bindTEELivenessVerifier(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &TEELivenessVerifier{TEELivenessVerifierCaller: TEELivenessVerifierCaller{contract: contract}, TEELivenessVerifierTransactor: TEELivenessVerifierTransactor{contract: contract}, TEELivenessVerifierFilterer: TEELivenessVerifierFilterer{contract: contract}}, nil
}

// NewTEELivenessVerifierCaller creates a new read-only instance of TEELivenessVerifier, bound to a specific deployed contract.
func NewTEELivenessVerifierCaller(address common.Address, caller bind.ContractCaller) (*TEELivenessVerifierCaller, error) {
	contract, err := bindTEELivenessVerifier(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TEELivenessVerifierCaller{contract: contract}, nil
}

// NewTEELivenessVerifierTransactor creates a new write-only instance of TEELivenessVerifier, bound to a specific deployed contract.
func NewTEELivenessVerifierTransactor(address common.Address, transactor bind.ContractTransactor) (*TEELivenessVerifierTransactor, error) {
	contract, err := bindTEELivenessVerifier(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TEELivenessVerifierTransactor{contract: contract}, nil
}

// NewTEELivenessVerifierFilterer creates a new log filterer instance of TEELivenessVerifier, bound to a specific deployed contract.
func NewTEELivenessVerifierFilterer(address common.Address, filterer bind.ContractFilterer) (*TEELivenessVerifierFilterer, error) {
	contract, err := bindTEELivenessVerifier(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TEELivenessVerifierFilterer{contract: contract}, nil
}

// bindTEELivenessVerifier binds a generic wrapper to an already deployed contract.
func bindTEELivenessVerifier(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := TEELivenessVerifierMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TEELivenessVerifier *TEELivenessVerifierRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TEELivenessVerifier.Contract.TEELivenessVerifierCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TEELivenessVerifier *TEELivenessVerifierRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TEELivenessVerifier.Contract.TEELivenessVerifierTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TEELivenessVerifier *TEELivenessVerifierRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TEELivenessVerifier.Contract.TEELivenessVerifierTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TEELivenessVerifier *TEELivenessVerifierCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TEELivenessVerifier.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TEELivenessVerifier *TEELivenessVerifierTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TEELivenessVerifier.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TEELivenessVerifier *TEELivenessVerifierTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TEELivenessVerifier.Contract.contract.Transact(opts, method, params...)
}

// AttestValiditySeconds is a free data retrieval call binding the contract method 0x95aa79ca.
//
// Solidity: function attestValiditySeconds() view returns(uint256)
func (_TEELivenessVerifier *TEELivenessVerifierCaller) AttestValiditySeconds(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _TEELivenessVerifier.contract.Call(opts, &out, "attestValiditySeconds")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AttestValiditySeconds is a free data retrieval call binding the contract method 0x95aa79ca.
//
// Solidity: function attestValiditySeconds() view returns(uint256)
func (_TEELivenessVerifier *TEELivenessVerifierSession) AttestValiditySeconds() (*big.Int, error) {
	return _TEELivenessVerifier.Contract.AttestValiditySeconds(&_TEELivenessVerifier.CallOpts)
}

// AttestValiditySeconds is a free data retrieval call binding the contract method 0x95aa79ca.
//
// Solidity: function attestValiditySeconds() view returns(uint256)
func (_TEELivenessVerifier *TEELivenessVerifierCallerSession) AttestValiditySeconds() (*big.Int, error) {
	return _TEELivenessVerifier.Contract.AttestValiditySeconds(&_TEELivenessVerifier.CallOpts)
}

// AttestedProvers is a free data retrieval call binding the contract method 0x3b7a3d82.
//
// Solidity: function attestedProvers(bytes32 ) view returns((bytes32,bytes32) pubkey, uint256 time)
func (_TEELivenessVerifier *TEELivenessVerifierCaller) AttestedProvers(opts *bind.CallOpts, arg0 [32]byte) (struct {
	Pubkey TEELivenessVerifierPubkey
	Time   *big.Int
}, error) {
	var out []interface{}
	err := _TEELivenessVerifier.contract.Call(opts, &out, "attestedProvers", arg0)

	outstruct := new(struct {
		Pubkey TEELivenessVerifierPubkey
		Time   *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Pubkey = *abi.ConvertType(out[0], new(TEELivenessVerifierPubkey)).(*TEELivenessVerifierPubkey)
	outstruct.Time = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// AttestedProvers is a free data retrieval call binding the contract method 0x3b7a3d82.
//
// Solidity: function attestedProvers(bytes32 ) view returns((bytes32,bytes32) pubkey, uint256 time)
func (_TEELivenessVerifier *TEELivenessVerifierSession) AttestedProvers(arg0 [32]byte) (struct {
	Pubkey TEELivenessVerifierPubkey
	Time   *big.Int
}, error) {
	return _TEELivenessVerifier.Contract.AttestedProvers(&_TEELivenessVerifier.CallOpts, arg0)
}

// AttestedProvers is a free data retrieval call binding the contract method 0x3b7a3d82.
//
// Solidity: function attestedProvers(bytes32 ) view returns((bytes32,bytes32) pubkey, uint256 time)
func (_TEELivenessVerifier *TEELivenessVerifierCallerSession) AttestedProvers(arg0 [32]byte) (struct {
	Pubkey TEELivenessVerifierPubkey
	Time   *big.Int
}, error) {
	return _TEELivenessVerifier.Contract.AttestedProvers(&_TEELivenessVerifier.CallOpts, arg0)
}

// AttestedReports is a free data retrieval call binding the contract method 0x74e1a553.
//
// Solidity: function attestedReports(bytes32 ) view returns(bool)
func (_TEELivenessVerifier *TEELivenessVerifierCaller) AttestedReports(opts *bind.CallOpts, arg0 [32]byte) (bool, error) {
	var out []interface{}
	err := _TEELivenessVerifier.contract.Call(opts, &out, "attestedReports", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// AttestedReports is a free data retrieval call binding the contract method 0x74e1a553.
//
// Solidity: function attestedReports(bytes32 ) view returns(bool)
func (_TEELivenessVerifier *TEELivenessVerifierSession) AttestedReports(arg0 [32]byte) (bool, error) {
	return _TEELivenessVerifier.Contract.AttestedReports(&_TEELivenessVerifier.CallOpts, arg0)
}

// AttestedReports is a free data retrieval call binding the contract method 0x74e1a553.
//
// Solidity: function attestedReports(bytes32 ) view returns(bool)
func (_TEELivenessVerifier *TEELivenessVerifierCallerSession) AttestedReports(arg0 [32]byte) (bool, error) {
	return _TEELivenessVerifier.Contract.AttestedReports(&_TEELivenessVerifier.CallOpts, arg0)
}

// DcapAttestation is a free data retrieval call binding the contract method 0x48e9b136.
//
// Solidity: function dcapAttestation() view returns(address)
func (_TEELivenessVerifier *TEELivenessVerifierCaller) DcapAttestation(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _TEELivenessVerifier.contract.Call(opts, &out, "dcapAttestation")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// DcapAttestation is a free data retrieval call binding the contract method 0x48e9b136.
//
// Solidity: function dcapAttestation() view returns(address)
func (_TEELivenessVerifier *TEELivenessVerifierSession) DcapAttestation() (common.Address, error) {
	return _TEELivenessVerifier.Contract.DcapAttestation(&_TEELivenessVerifier.CallOpts)
}

// DcapAttestation is a free data retrieval call binding the contract method 0x48e9b136.
//
// Solidity: function dcapAttestation() view returns(address)
func (_TEELivenessVerifier *TEELivenessVerifierCallerSession) DcapAttestation() (common.Address, error) {
	return _TEELivenessVerifier.Contract.DcapAttestation(&_TEELivenessVerifier.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_TEELivenessVerifier *TEELivenessVerifierCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _TEELivenessVerifier.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_TEELivenessVerifier *TEELivenessVerifierSession) Owner() (common.Address, error) {
	return _TEELivenessVerifier.Contract.Owner(&_TEELivenessVerifier.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_TEELivenessVerifier *TEELivenessVerifierCallerSession) Owner() (common.Address, error) {
	return _TEELivenessVerifier.Contract.Owner(&_TEELivenessVerifier.CallOpts)
}

// Simulation is a free data retrieval call binding the contract method 0x95f76bdd.
//
// Solidity: function simulation() view returns(bool)
func (_TEELivenessVerifier *TEELivenessVerifierCaller) Simulation(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _TEELivenessVerifier.contract.Call(opts, &out, "simulation")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Simulation is a free data retrieval call binding the contract method 0x95f76bdd.
//
// Solidity: function simulation() view returns(bool)
func (_TEELivenessVerifier *TEELivenessVerifierSession) Simulation() (bool, error) {
	return _TEELivenessVerifier.Contract.Simulation(&_TEELivenessVerifier.CallOpts)
}

// Simulation is a free data retrieval call binding the contract method 0x95f76bdd.
//
// Solidity: function simulation() view returns(bool)
func (_TEELivenessVerifier *TEELivenessVerifierCallerSession) Simulation() (bool, error) {
	return _TEELivenessVerifier.Contract.Simulation(&_TEELivenessVerifier.CallOpts)
}

// VerifyAttestation is a free data retrieval call binding the contract method 0xa65542ca.
//
// Solidity: function verifyAttestation(bytes32 pubkeyX, bytes32 pubkeyY, bytes data) view returns(bool)
func (_TEELivenessVerifier *TEELivenessVerifierCaller) VerifyAttestation(opts *bind.CallOpts, pubkeyX [32]byte, pubkeyY [32]byte, data []byte) (bool, error) {
	var out []interface{}
	err := _TEELivenessVerifier.contract.Call(opts, &out, "verifyAttestation", pubkeyX, pubkeyY, data)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// VerifyAttestation is a free data retrieval call binding the contract method 0xa65542ca.
//
// Solidity: function verifyAttestation(bytes32 pubkeyX, bytes32 pubkeyY, bytes data) view returns(bool)
func (_TEELivenessVerifier *TEELivenessVerifierSession) VerifyAttestation(pubkeyX [32]byte, pubkeyY [32]byte, data []byte) (bool, error) {
	return _TEELivenessVerifier.Contract.VerifyAttestation(&_TEELivenessVerifier.CallOpts, pubkeyX, pubkeyY, data)
}

// VerifyAttestation is a free data retrieval call binding the contract method 0xa65542ca.
//
// Solidity: function verifyAttestation(bytes32 pubkeyX, bytes32 pubkeyY, bytes data) view returns(bool)
func (_TEELivenessVerifier *TEELivenessVerifierCallerSession) VerifyAttestation(pubkeyX [32]byte, pubkeyY [32]byte, data []byte) (bool, error) {
	return _TEELivenessVerifier.Contract.VerifyAttestation(&_TEELivenessVerifier.CallOpts, pubkeyX, pubkeyY, data)
}

// VerifyLivenessProof is a free data retrieval call binding the contract method 0xb72c58b1.
//
// Solidity: function verifyLivenessProof(bytes32 pubkeyX, bytes32 pubkeyY) view returns(bool)
func (_TEELivenessVerifier *TEELivenessVerifierCaller) VerifyLivenessProof(opts *bind.CallOpts, pubkeyX [32]byte, pubkeyY [32]byte) (bool, error) {
	var out []interface{}
	err := _TEELivenessVerifier.contract.Call(opts, &out, "verifyLivenessProof", pubkeyX, pubkeyY)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// VerifyLivenessProof is a free data retrieval call binding the contract method 0xb72c58b1.
//
// Solidity: function verifyLivenessProof(bytes32 pubkeyX, bytes32 pubkeyY) view returns(bool)
func (_TEELivenessVerifier *TEELivenessVerifierSession) VerifyLivenessProof(pubkeyX [32]byte, pubkeyY [32]byte) (bool, error) {
	return _TEELivenessVerifier.Contract.VerifyLivenessProof(&_TEELivenessVerifier.CallOpts, pubkeyX, pubkeyY)
}

// VerifyLivenessProof is a free data retrieval call binding the contract method 0xb72c58b1.
//
// Solidity: function verifyLivenessProof(bytes32 pubkeyX, bytes32 pubkeyY) view returns(bool)
func (_TEELivenessVerifier *TEELivenessVerifierCallerSession) VerifyLivenessProof(pubkeyX [32]byte, pubkeyY [32]byte) (bool, error) {
	return _TEELivenessVerifier.Contract.VerifyLivenessProof(&_TEELivenessVerifier.CallOpts, pubkeyX, pubkeyY)
}

// VerifyMrEnclave is a free data retrieval call binding the contract method 0x18a98d6f.
//
// Solidity: function verifyMrEnclave(bytes32 _mrenclave) view returns(bool)
func (_TEELivenessVerifier *TEELivenessVerifierCaller) VerifyMrEnclave(opts *bind.CallOpts, _mrenclave [32]byte) (bool, error) {
	var out []interface{}
	err := _TEELivenessVerifier.contract.Call(opts, &out, "verifyMrEnclave", _mrenclave)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// VerifyMrEnclave is a free data retrieval call binding the contract method 0x18a98d6f.
//
// Solidity: function verifyMrEnclave(bytes32 _mrenclave) view returns(bool)
func (_TEELivenessVerifier *TEELivenessVerifierSession) VerifyMrEnclave(_mrenclave [32]byte) (bool, error) {
	return _TEELivenessVerifier.Contract.VerifyMrEnclave(&_TEELivenessVerifier.CallOpts, _mrenclave)
}

// VerifyMrEnclave is a free data retrieval call binding the contract method 0x18a98d6f.
//
// Solidity: function verifyMrEnclave(bytes32 _mrenclave) view returns(bool)
func (_TEELivenessVerifier *TEELivenessVerifierCallerSession) VerifyMrEnclave(_mrenclave [32]byte) (bool, error) {
	return _TEELivenessVerifier.Contract.VerifyMrEnclave(&_TEELivenessVerifier.CallOpts, _mrenclave)
}

// VerifyMrSigner is a free data retrieval call binding the contract method 0xc64b860f.
//
// Solidity: function verifyMrSigner(bytes32 _mrsigner) view returns(bool)
func (_TEELivenessVerifier *TEELivenessVerifierCaller) VerifyMrSigner(opts *bind.CallOpts, _mrsigner [32]byte) (bool, error) {
	var out []interface{}
	err := _TEELivenessVerifier.contract.Call(opts, &out, "verifyMrSigner", _mrsigner)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// VerifyMrSigner is a free data retrieval call binding the contract method 0xc64b860f.
//
// Solidity: function verifyMrSigner(bytes32 _mrsigner) view returns(bool)
func (_TEELivenessVerifier *TEELivenessVerifierSession) VerifyMrSigner(_mrsigner [32]byte) (bool, error) {
	return _TEELivenessVerifier.Contract.VerifyMrSigner(&_TEELivenessVerifier.CallOpts, _mrsigner)
}

// VerifyMrSigner is a free data retrieval call binding the contract method 0xc64b860f.
//
// Solidity: function verifyMrSigner(bytes32 _mrsigner) view returns(bool)
func (_TEELivenessVerifier *TEELivenessVerifierCallerSession) VerifyMrSigner(_mrsigner [32]byte) (bool, error) {
	return _TEELivenessVerifier.Contract.VerifyMrSigner(&_TEELivenessVerifier.CallOpts, _mrsigner)
}

// ChangeAttestValiditySeconds is a paid mutator transaction binding the contract method 0x315aa7d9.
//
// Solidity: function changeAttestValiditySeconds(uint256 val) returns()
func (_TEELivenessVerifier *TEELivenessVerifierTransactor) ChangeAttestValiditySeconds(opts *bind.TransactOpts, val *big.Int) (*types.Transaction, error) {
	return _TEELivenessVerifier.contract.Transact(opts, "changeAttestValiditySeconds", val)
}

// ChangeAttestValiditySeconds is a paid mutator transaction binding the contract method 0x315aa7d9.
//
// Solidity: function changeAttestValiditySeconds(uint256 val) returns()
func (_TEELivenessVerifier *TEELivenessVerifierSession) ChangeAttestValiditySeconds(val *big.Int) (*types.Transaction, error) {
	return _TEELivenessVerifier.Contract.ChangeAttestValiditySeconds(&_TEELivenessVerifier.TransactOpts, val)
}

// ChangeAttestValiditySeconds is a paid mutator transaction binding the contract method 0x315aa7d9.
//
// Solidity: function changeAttestValiditySeconds(uint256 val) returns()
func (_TEELivenessVerifier *TEELivenessVerifierTransactorSession) ChangeAttestValiditySeconds(val *big.Int) (*types.Transaction, error) {
	return _TEELivenessVerifier.Contract.ChangeAttestValiditySeconds(&_TEELivenessVerifier.TransactOpts, val)
}

// ChangeOwner is a paid mutator transaction binding the contract method 0xa6f9dae1.
//
// Solidity: function changeOwner(address _newOwner) returns()
func (_TEELivenessVerifier *TEELivenessVerifierTransactor) ChangeOwner(opts *bind.TransactOpts, _newOwner common.Address) (*types.Transaction, error) {
	return _TEELivenessVerifier.contract.Transact(opts, "changeOwner", _newOwner)
}

// ChangeOwner is a paid mutator transaction binding the contract method 0xa6f9dae1.
//
// Solidity: function changeOwner(address _newOwner) returns()
func (_TEELivenessVerifier *TEELivenessVerifierSession) ChangeOwner(_newOwner common.Address) (*types.Transaction, error) {
	return _TEELivenessVerifier.Contract.ChangeOwner(&_TEELivenessVerifier.TransactOpts, _newOwner)
}

// ChangeOwner is a paid mutator transaction binding the contract method 0xa6f9dae1.
//
// Solidity: function changeOwner(address _newOwner) returns()
func (_TEELivenessVerifier *TEELivenessVerifierTransactorSession) ChangeOwner(_newOwner common.Address) (*types.Transaction, error) {
	return _TEELivenessVerifier.Contract.ChangeOwner(&_TEELivenessVerifier.TransactOpts, _newOwner)
}

// SubmitLivenessProof is a paid mutator transaction binding the contract method 0x6cad7294.
//
// Solidity: function submitLivenessProof(bytes _report) returns()
func (_TEELivenessVerifier *TEELivenessVerifierTransactor) SubmitLivenessProof(opts *bind.TransactOpts, _report []byte) (*types.Transaction, error) {
	return _TEELivenessVerifier.contract.Transact(opts, "submitLivenessProof", _report)
}

// SubmitLivenessProof is a paid mutator transaction binding the contract method 0x6cad7294.
//
// Solidity: function submitLivenessProof(bytes _report) returns()
func (_TEELivenessVerifier *TEELivenessVerifierSession) SubmitLivenessProof(_report []byte) (*types.Transaction, error) {
	return _TEELivenessVerifier.Contract.SubmitLivenessProof(&_TEELivenessVerifier.TransactOpts, _report)
}

// SubmitLivenessProof is a paid mutator transaction binding the contract method 0x6cad7294.
//
// Solidity: function submitLivenessProof(bytes _report) returns()
func (_TEELivenessVerifier *TEELivenessVerifierTransactorSession) SubmitLivenessProof(_report []byte) (*types.Transaction, error) {
	return _TEELivenessVerifier.Contract.SubmitLivenessProof(&_TEELivenessVerifier.TransactOpts, _report)
}
