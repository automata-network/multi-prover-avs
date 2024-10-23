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

// TEELivenessVerifierReportDataV2 is an auto generated low-level Go binding around an user-defined struct.
type TEELivenessVerifierReportDataV2 struct {
	Pubkey               TEELivenessVerifierPubkey
	ReferenceBlockNumber *big.Int
	ReferenceBlockHash   [32]byte
	ProverAddressHash    [32]byte
}

// TEELivenessVerifierMetaData contains all meta data concerning the TEELivenessVerifier contract.
var TEELivenessVerifierMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"attestValiditySeconds\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"attestedProverAddr\",\"inputs\":[{\"name\":\"proverKey\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"proverAddr\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"attestedProvers\",\"inputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"pubkey\",\"type\":\"tuple\",\"internalType\":\"structTEELivenessVerifier.Pubkey\",\"components\":[{\"name\":\"x\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"y\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"name\":\"time\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"attestedReports\",\"inputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"changeAttestValiditySeconds\",\"inputs\":[{\"name\":\"val\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"changeAttestationImpl\",\"inputs\":[{\"name\":\"_attestationAddr\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"changeMaxBlockNumberDiff\",\"inputs\":[{\"name\":\"_maxBlockNumberDiff\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"dcapAttestation\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIAttestation\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"initialize\",\"inputs\":[{\"name\":\"_initialOwner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_attestationAddr\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_maxBlockNumberDiff\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_attestValiditySeconds\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"maxBlockNumberDiff\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"reinitialize\",\"inputs\":[{\"name\":\"i\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"_initialOwner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_attestationAddr\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_maxBlockNumberDiff\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_attestValiditySeconds\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"renounceOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"submitLivenessProofV2\",\"inputs\":[{\"name\":\"_data\",\"type\":\"tuple\",\"internalType\":\"structTEELivenessVerifier.ReportDataV2\",\"components\":[{\"name\":\"pubkey\",\"type\":\"tuple\",\"internalType\":\"structTEELivenessVerifier.Pubkey\",\"components\":[{\"name\":\"x\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"y\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"name\":\"referenceBlockNumber\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"referenceBlockHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"proverAddressHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"name\":\"_report\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"verifyAttestationV2\",\"inputs\":[{\"name\":\"pubkeyX\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"pubkeyY\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"verifyLivenessProof\",\"inputs\":[{\"name\":\"pubkeyX\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"pubkeyY\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"verifyLivenessProofV2\",\"inputs\":[{\"name\":\"pubkeyX\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"pubkeyY\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"proverKey\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"verifyMrEnclave\",\"inputs\":[{\"name\":\"_mrenclave\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"verifyMrSigner\",\"inputs\":[{\"name\":\"_mrsigner\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"event\",\"name\":\"Initialized\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"uint8\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"previousOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false}]",
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

// AttestedProverAddr is a free data retrieval call binding the contract method 0x80c19970.
//
// Solidity: function attestedProverAddr(bytes32 proverKey) view returns(address proverAddr)
func (_TEELivenessVerifier *TEELivenessVerifierCaller) AttestedProverAddr(opts *bind.CallOpts, proverKey [32]byte) (common.Address, error) {
	var out []interface{}
	err := _TEELivenessVerifier.contract.Call(opts, &out, "attestedProverAddr", proverKey)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AttestedProverAddr is a free data retrieval call binding the contract method 0x80c19970.
//
// Solidity: function attestedProverAddr(bytes32 proverKey) view returns(address proverAddr)
func (_TEELivenessVerifier *TEELivenessVerifierSession) AttestedProverAddr(proverKey [32]byte) (common.Address, error) {
	return _TEELivenessVerifier.Contract.AttestedProverAddr(&_TEELivenessVerifier.CallOpts, proverKey)
}

// AttestedProverAddr is a free data retrieval call binding the contract method 0x80c19970.
//
// Solidity: function attestedProverAddr(bytes32 proverKey) view returns(address proverAddr)
func (_TEELivenessVerifier *TEELivenessVerifierCallerSession) AttestedProverAddr(proverKey [32]byte) (common.Address, error) {
	return _TEELivenessVerifier.Contract.AttestedProverAddr(&_TEELivenessVerifier.CallOpts, proverKey)
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

// MaxBlockNumberDiff is a free data retrieval call binding the contract method 0xfeb1cb80.
//
// Solidity: function maxBlockNumberDiff() view returns(uint256)
func (_TEELivenessVerifier *TEELivenessVerifierCaller) MaxBlockNumberDiff(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _TEELivenessVerifier.contract.Call(opts, &out, "maxBlockNumberDiff")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxBlockNumberDiff is a free data retrieval call binding the contract method 0xfeb1cb80.
//
// Solidity: function maxBlockNumberDiff() view returns(uint256)
func (_TEELivenessVerifier *TEELivenessVerifierSession) MaxBlockNumberDiff() (*big.Int, error) {
	return _TEELivenessVerifier.Contract.MaxBlockNumberDiff(&_TEELivenessVerifier.CallOpts)
}

// MaxBlockNumberDiff is a free data retrieval call binding the contract method 0xfeb1cb80.
//
// Solidity: function maxBlockNumberDiff() view returns(uint256)
func (_TEELivenessVerifier *TEELivenessVerifierCallerSession) MaxBlockNumberDiff() (*big.Int, error) {
	return _TEELivenessVerifier.Contract.MaxBlockNumberDiff(&_TEELivenessVerifier.CallOpts)
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

// VerifyAttestationV2 is a free data retrieval call binding the contract method 0xce3234b8.
//
// Solidity: function verifyAttestationV2(bytes32 pubkeyX, bytes32 pubkeyY, bytes data) view returns(bool)
func (_TEELivenessVerifier *TEELivenessVerifierCaller) VerifyAttestationV2(opts *bind.CallOpts, pubkeyX [32]byte, pubkeyY [32]byte, data []byte) (bool, error) {
	var out []interface{}
	err := _TEELivenessVerifier.contract.Call(opts, &out, "verifyAttestationV2", pubkeyX, pubkeyY, data)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// VerifyAttestationV2 is a free data retrieval call binding the contract method 0xce3234b8.
//
// Solidity: function verifyAttestationV2(bytes32 pubkeyX, bytes32 pubkeyY, bytes data) view returns(bool)
func (_TEELivenessVerifier *TEELivenessVerifierSession) VerifyAttestationV2(pubkeyX [32]byte, pubkeyY [32]byte, data []byte) (bool, error) {
	return _TEELivenessVerifier.Contract.VerifyAttestationV2(&_TEELivenessVerifier.CallOpts, pubkeyX, pubkeyY, data)
}

// VerifyAttestationV2 is a free data retrieval call binding the contract method 0xce3234b8.
//
// Solidity: function verifyAttestationV2(bytes32 pubkeyX, bytes32 pubkeyY, bytes data) view returns(bool)
func (_TEELivenessVerifier *TEELivenessVerifierCallerSession) VerifyAttestationV2(pubkeyX [32]byte, pubkeyY [32]byte, data []byte) (bool, error) {
	return _TEELivenessVerifier.Contract.VerifyAttestationV2(&_TEELivenessVerifier.CallOpts, pubkeyX, pubkeyY, data)
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

// VerifyLivenessProofV2 is a free data retrieval call binding the contract method 0xa14065dc.
//
// Solidity: function verifyLivenessProofV2(bytes32 pubkeyX, bytes32 pubkeyY, address proverKey) view returns(bool)
func (_TEELivenessVerifier *TEELivenessVerifierCaller) VerifyLivenessProofV2(opts *bind.CallOpts, pubkeyX [32]byte, pubkeyY [32]byte, proverKey common.Address) (bool, error) {
	var out []interface{}
	err := _TEELivenessVerifier.contract.Call(opts, &out, "verifyLivenessProofV2", pubkeyX, pubkeyY, proverKey)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// VerifyLivenessProofV2 is a free data retrieval call binding the contract method 0xa14065dc.
//
// Solidity: function verifyLivenessProofV2(bytes32 pubkeyX, bytes32 pubkeyY, address proverKey) view returns(bool)
func (_TEELivenessVerifier *TEELivenessVerifierSession) VerifyLivenessProofV2(pubkeyX [32]byte, pubkeyY [32]byte, proverKey common.Address) (bool, error) {
	return _TEELivenessVerifier.Contract.VerifyLivenessProofV2(&_TEELivenessVerifier.CallOpts, pubkeyX, pubkeyY, proverKey)
}

// VerifyLivenessProofV2 is a free data retrieval call binding the contract method 0xa14065dc.
//
// Solidity: function verifyLivenessProofV2(bytes32 pubkeyX, bytes32 pubkeyY, address proverKey) view returns(bool)
func (_TEELivenessVerifier *TEELivenessVerifierCallerSession) VerifyLivenessProofV2(pubkeyX [32]byte, pubkeyY [32]byte, proverKey common.Address) (bool, error) {
	return _TEELivenessVerifier.Contract.VerifyLivenessProofV2(&_TEELivenessVerifier.CallOpts, pubkeyX, pubkeyY, proverKey)
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

// ChangeAttestationImpl is a paid mutator transaction binding the contract method 0xdd6e7951.
//
// Solidity: function changeAttestationImpl(address _attestationAddr) returns()
func (_TEELivenessVerifier *TEELivenessVerifierTransactor) ChangeAttestationImpl(opts *bind.TransactOpts, _attestationAddr common.Address) (*types.Transaction, error) {
	return _TEELivenessVerifier.contract.Transact(opts, "changeAttestationImpl", _attestationAddr)
}

// ChangeAttestationImpl is a paid mutator transaction binding the contract method 0xdd6e7951.
//
// Solidity: function changeAttestationImpl(address _attestationAddr) returns()
func (_TEELivenessVerifier *TEELivenessVerifierSession) ChangeAttestationImpl(_attestationAddr common.Address) (*types.Transaction, error) {
	return _TEELivenessVerifier.Contract.ChangeAttestationImpl(&_TEELivenessVerifier.TransactOpts, _attestationAddr)
}

// ChangeAttestationImpl is a paid mutator transaction binding the contract method 0xdd6e7951.
//
// Solidity: function changeAttestationImpl(address _attestationAddr) returns()
func (_TEELivenessVerifier *TEELivenessVerifierTransactorSession) ChangeAttestationImpl(_attestationAddr common.Address) (*types.Transaction, error) {
	return _TEELivenessVerifier.Contract.ChangeAttestationImpl(&_TEELivenessVerifier.TransactOpts, _attestationAddr)
}

// ChangeMaxBlockNumberDiff is a paid mutator transaction binding the contract method 0x6cb5ec95.
//
// Solidity: function changeMaxBlockNumberDiff(uint256 _maxBlockNumberDiff) returns()
func (_TEELivenessVerifier *TEELivenessVerifierTransactor) ChangeMaxBlockNumberDiff(opts *bind.TransactOpts, _maxBlockNumberDiff *big.Int) (*types.Transaction, error) {
	return _TEELivenessVerifier.contract.Transact(opts, "changeMaxBlockNumberDiff", _maxBlockNumberDiff)
}

// ChangeMaxBlockNumberDiff is a paid mutator transaction binding the contract method 0x6cb5ec95.
//
// Solidity: function changeMaxBlockNumberDiff(uint256 _maxBlockNumberDiff) returns()
func (_TEELivenessVerifier *TEELivenessVerifierSession) ChangeMaxBlockNumberDiff(_maxBlockNumberDiff *big.Int) (*types.Transaction, error) {
	return _TEELivenessVerifier.Contract.ChangeMaxBlockNumberDiff(&_TEELivenessVerifier.TransactOpts, _maxBlockNumberDiff)
}

// ChangeMaxBlockNumberDiff is a paid mutator transaction binding the contract method 0x6cb5ec95.
//
// Solidity: function changeMaxBlockNumberDiff(uint256 _maxBlockNumberDiff) returns()
func (_TEELivenessVerifier *TEELivenessVerifierTransactorSession) ChangeMaxBlockNumberDiff(_maxBlockNumberDiff *big.Int) (*types.Transaction, error) {
	return _TEELivenessVerifier.Contract.ChangeMaxBlockNumberDiff(&_TEELivenessVerifier.TransactOpts, _maxBlockNumberDiff)
}

// Initialize is a paid mutator transaction binding the contract method 0xeb990c59.
//
// Solidity: function initialize(address _initialOwner, address _attestationAddr, uint256 _maxBlockNumberDiff, uint256 _attestValiditySeconds) returns()
func (_TEELivenessVerifier *TEELivenessVerifierTransactor) Initialize(opts *bind.TransactOpts, _initialOwner common.Address, _attestationAddr common.Address, _maxBlockNumberDiff *big.Int, _attestValiditySeconds *big.Int) (*types.Transaction, error) {
	return _TEELivenessVerifier.contract.Transact(opts, "initialize", _initialOwner, _attestationAddr, _maxBlockNumberDiff, _attestValiditySeconds)
}

// Initialize is a paid mutator transaction binding the contract method 0xeb990c59.
//
// Solidity: function initialize(address _initialOwner, address _attestationAddr, uint256 _maxBlockNumberDiff, uint256 _attestValiditySeconds) returns()
func (_TEELivenessVerifier *TEELivenessVerifierSession) Initialize(_initialOwner common.Address, _attestationAddr common.Address, _maxBlockNumberDiff *big.Int, _attestValiditySeconds *big.Int) (*types.Transaction, error) {
	return _TEELivenessVerifier.Contract.Initialize(&_TEELivenessVerifier.TransactOpts, _initialOwner, _attestationAddr, _maxBlockNumberDiff, _attestValiditySeconds)
}

// Initialize is a paid mutator transaction binding the contract method 0xeb990c59.
//
// Solidity: function initialize(address _initialOwner, address _attestationAddr, uint256 _maxBlockNumberDiff, uint256 _attestValiditySeconds) returns()
func (_TEELivenessVerifier *TEELivenessVerifierTransactorSession) Initialize(_initialOwner common.Address, _attestationAddr common.Address, _maxBlockNumberDiff *big.Int, _attestValiditySeconds *big.Int) (*types.Transaction, error) {
	return _TEELivenessVerifier.Contract.Initialize(&_TEELivenessVerifier.TransactOpts, _initialOwner, _attestationAddr, _maxBlockNumberDiff, _attestValiditySeconds)
}

// Reinitialize is a paid mutator transaction binding the contract method 0x8087b269.
//
// Solidity: function reinitialize(uint8 i, address _initialOwner, address _attestationAddr, uint256 _maxBlockNumberDiff, uint256 _attestValiditySeconds) returns()
func (_TEELivenessVerifier *TEELivenessVerifierTransactor) Reinitialize(opts *bind.TransactOpts, i uint8, _initialOwner common.Address, _attestationAddr common.Address, _maxBlockNumberDiff *big.Int, _attestValiditySeconds *big.Int) (*types.Transaction, error) {
	return _TEELivenessVerifier.contract.Transact(opts, "reinitialize", i, _initialOwner, _attestationAddr, _maxBlockNumberDiff, _attestValiditySeconds)
}

// Reinitialize is a paid mutator transaction binding the contract method 0x8087b269.
//
// Solidity: function reinitialize(uint8 i, address _initialOwner, address _attestationAddr, uint256 _maxBlockNumberDiff, uint256 _attestValiditySeconds) returns()
func (_TEELivenessVerifier *TEELivenessVerifierSession) Reinitialize(i uint8, _initialOwner common.Address, _attestationAddr common.Address, _maxBlockNumberDiff *big.Int, _attestValiditySeconds *big.Int) (*types.Transaction, error) {
	return _TEELivenessVerifier.Contract.Reinitialize(&_TEELivenessVerifier.TransactOpts, i, _initialOwner, _attestationAddr, _maxBlockNumberDiff, _attestValiditySeconds)
}

// Reinitialize is a paid mutator transaction binding the contract method 0x8087b269.
//
// Solidity: function reinitialize(uint8 i, address _initialOwner, address _attestationAddr, uint256 _maxBlockNumberDiff, uint256 _attestValiditySeconds) returns()
func (_TEELivenessVerifier *TEELivenessVerifierTransactorSession) Reinitialize(i uint8, _initialOwner common.Address, _attestationAddr common.Address, _maxBlockNumberDiff *big.Int, _attestValiditySeconds *big.Int) (*types.Transaction, error) {
	return _TEELivenessVerifier.Contract.Reinitialize(&_TEELivenessVerifier.TransactOpts, i, _initialOwner, _attestationAddr, _maxBlockNumberDiff, _attestValiditySeconds)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_TEELivenessVerifier *TEELivenessVerifierTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TEELivenessVerifier.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_TEELivenessVerifier *TEELivenessVerifierSession) RenounceOwnership() (*types.Transaction, error) {
	return _TEELivenessVerifier.Contract.RenounceOwnership(&_TEELivenessVerifier.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_TEELivenessVerifier *TEELivenessVerifierTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _TEELivenessVerifier.Contract.RenounceOwnership(&_TEELivenessVerifier.TransactOpts)
}

// SubmitLivenessProofV2 is a paid mutator transaction binding the contract method 0xfb691bbe.
//
// Solidity: function submitLivenessProofV2(((bytes32,bytes32),uint256,bytes32,bytes32) _data, bytes _report) returns()
func (_TEELivenessVerifier *TEELivenessVerifierTransactor) SubmitLivenessProofV2(opts *bind.TransactOpts, _data TEELivenessVerifierReportDataV2, _report []byte) (*types.Transaction, error) {
	return _TEELivenessVerifier.contract.Transact(opts, "submitLivenessProofV2", _data, _report)
}

// SubmitLivenessProofV2 is a paid mutator transaction binding the contract method 0xfb691bbe.
//
// Solidity: function submitLivenessProofV2(((bytes32,bytes32),uint256,bytes32,bytes32) _data, bytes _report) returns()
func (_TEELivenessVerifier *TEELivenessVerifierSession) SubmitLivenessProofV2(_data TEELivenessVerifierReportDataV2, _report []byte) (*types.Transaction, error) {
	return _TEELivenessVerifier.Contract.SubmitLivenessProofV2(&_TEELivenessVerifier.TransactOpts, _data, _report)
}

// SubmitLivenessProofV2 is a paid mutator transaction binding the contract method 0xfb691bbe.
//
// Solidity: function submitLivenessProofV2(((bytes32,bytes32),uint256,bytes32,bytes32) _data, bytes _report) returns()
func (_TEELivenessVerifier *TEELivenessVerifierTransactorSession) SubmitLivenessProofV2(_data TEELivenessVerifierReportDataV2, _report []byte) (*types.Transaction, error) {
	return _TEELivenessVerifier.Contract.SubmitLivenessProofV2(&_TEELivenessVerifier.TransactOpts, _data, _report)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_TEELivenessVerifier *TEELivenessVerifierTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _TEELivenessVerifier.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_TEELivenessVerifier *TEELivenessVerifierSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _TEELivenessVerifier.Contract.TransferOwnership(&_TEELivenessVerifier.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_TEELivenessVerifier *TEELivenessVerifierTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _TEELivenessVerifier.Contract.TransferOwnership(&_TEELivenessVerifier.TransactOpts, newOwner)
}

// TEELivenessVerifierInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the TEELivenessVerifier contract.
type TEELivenessVerifierInitializedIterator struct {
	Event *TEELivenessVerifierInitialized // Event containing the contract specifics and raw log

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
func (it *TEELivenessVerifierInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TEELivenessVerifierInitialized)
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
		it.Event = new(TEELivenessVerifierInitialized)
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
func (it *TEELivenessVerifierInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TEELivenessVerifierInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TEELivenessVerifierInitialized represents a Initialized event raised by the TEELivenessVerifier contract.
type TEELivenessVerifierInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_TEELivenessVerifier *TEELivenessVerifierFilterer) FilterInitialized(opts *bind.FilterOpts) (*TEELivenessVerifierInitializedIterator, error) {

	logs, sub, err := _TEELivenessVerifier.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &TEELivenessVerifierInitializedIterator{contract: _TEELivenessVerifier.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_TEELivenessVerifier *TEELivenessVerifierFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *TEELivenessVerifierInitialized) (event.Subscription, error) {

	logs, sub, err := _TEELivenessVerifier.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TEELivenessVerifierInitialized)
				if err := _TEELivenessVerifier.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_TEELivenessVerifier *TEELivenessVerifierFilterer) ParseInitialized(log types.Log) (*TEELivenessVerifierInitialized, error) {
	event := new(TEELivenessVerifierInitialized)
	if err := _TEELivenessVerifier.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TEELivenessVerifierOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the TEELivenessVerifier contract.
type TEELivenessVerifierOwnershipTransferredIterator struct {
	Event *TEELivenessVerifierOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *TEELivenessVerifierOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TEELivenessVerifierOwnershipTransferred)
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
		it.Event = new(TEELivenessVerifierOwnershipTransferred)
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
func (it *TEELivenessVerifierOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TEELivenessVerifierOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TEELivenessVerifierOwnershipTransferred represents a OwnershipTransferred event raised by the TEELivenessVerifier contract.
type TEELivenessVerifierOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_TEELivenessVerifier *TEELivenessVerifierFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*TEELivenessVerifierOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _TEELivenessVerifier.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &TEELivenessVerifierOwnershipTransferredIterator{contract: _TEELivenessVerifier.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_TEELivenessVerifier *TEELivenessVerifierFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *TEELivenessVerifierOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _TEELivenessVerifier.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TEELivenessVerifierOwnershipTransferred)
				if err := _TEELivenessVerifier.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_TEELivenessVerifier *TEELivenessVerifierFilterer) ParseOwnershipTransferred(log types.Log) (*TEELivenessVerifierOwnershipTransferred, error) {
	event := new(TEELivenessVerifierOwnershipTransferred)
	if err := _TEELivenessVerifier.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
