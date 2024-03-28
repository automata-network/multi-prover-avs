// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package AutomataDcapV3Attestation

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

// EnclaveIdStructEnclaveId is an auto generated low-level Go binding around an user-defined struct.
type EnclaveIdStructEnclaveId struct {
	Miscselect     [4]byte
	MiscselectMask [4]byte
	Isvprodid      uint16
	Attributes     [16]byte
	AttributesMask [16]byte
	Mrsigner       [32]byte
	TcbLevels      []EnclaveIdStructTcbLevel
}

// EnclaveIdStructTcbLevel is an auto generated low-level Go binding around an user-defined struct.
type EnclaveIdStructTcbLevel struct {
	Tcb       EnclaveIdStructTcbObj
	TcbStatus uint8
}

// EnclaveIdStructTcbObj is an auto generated low-level Go binding around an user-defined struct.
type EnclaveIdStructTcbObj struct {
	Isvsvn uint16
}

// TCBInfoStructTCBInfo is an auto generated low-level Go binding around an user-defined struct.
type TCBInfoStructTCBInfo struct {
	Pceid     string
	Fmspc     string
	TcbLevels []TCBInfoStructTCBLevelObj
}

// TCBInfoStructTCBLevelObj is an auto generated low-level Go binding around an user-defined struct.
type TCBInfoStructTCBLevelObj struct {
	Pcesvn           *big.Int
	SgxTcbCompSvnArr []*big.Int
	Status           uint8
}

// V3StructCertificationData is an auto generated low-level Go binding around an user-defined struct.
type V3StructCertificationData struct {
	CertType             uint16
	CertDataSize         uint32
	DecodedCertDataArray [3][]byte
}

// V3StructECDSAQuoteV3AuthData is an auto generated low-level Go binding around an user-defined struct.
type V3StructECDSAQuoteV3AuthData struct {
	Ecdsa256BitSignature []byte
	EcdsaAttestationKey  []byte
	PckSignedQeReport    V3StructEnclaveReport
	QeReportSignature    []byte
	QeAuthData           V3StructQEAuthData
	Certification        V3StructCertificationData
}

// V3StructEnclaveReport is an auto generated low-level Go binding around an user-defined struct.
type V3StructEnclaveReport struct {
	CpuSvn     [16]byte
	MiscSelect [4]byte
	Reserved1  [28]byte
	Attributes [16]byte
	MrEnclave  [32]byte
	Reserved2  [32]byte
	MrSigner   [32]byte
	Reserved3  []byte
	IsvProdId  uint16
	IsvSvn     uint16
	Reserved4  []byte
	ReportData []byte
}

// V3StructHeader is an auto generated low-level Go binding around an user-defined struct.
type V3StructHeader struct {
	Version            [2]byte
	AttestationKeyType [2]byte
	TeeType            [4]byte
	QeSvn              [2]byte
	PceSvn             [2]byte
	QeVendorId         [16]byte
	UserData           [20]byte
}

// V3StructParsedV3Quote is an auto generated low-level Go binding around an user-defined struct.
type V3StructParsedV3Quote struct {
	Header             V3StructHeader
	LocalEnclaveReport V3StructEnclaveReport
	V3AuthData         V3StructECDSAQuoteV3AuthData
}

// V3StructQEAuthData is an auto generated low-level Go binding around an user-defined struct.
type V3StructQEAuthData struct {
	ParsedDataSize uint16
	Data           []byte
}

// AutomataDcapV3AttestationMetaData contains all meta data concerning the AutomataDcapV3Attestation contract.
var AutomataDcapV3AttestationMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sigVerifyLibAddr\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"pemCertLibAddr\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"internalType\":\"bytes[]\",\"name\":\"serialNumBatch\",\"type\":\"bytes[]\"}],\"name\":\"addRevokedCertSerialNum\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes4\",\"name\":\"miscselect\",\"type\":\"bytes4\"},{\"internalType\":\"bytes4\",\"name\":\"miscselectMask\",\"type\":\"bytes4\"},{\"internalType\":\"uint16\",\"name\":\"isvprodid\",\"type\":\"uint16\"},{\"internalType\":\"bytes16\",\"name\":\"attributes\",\"type\":\"bytes16\"},{\"internalType\":\"bytes16\",\"name\":\"attributesMask\",\"type\":\"bytes16\"},{\"internalType\":\"bytes32\",\"name\":\"mrsigner\",\"type\":\"bytes32\"},{\"components\":[{\"components\":[{\"internalType\":\"uint16\",\"name\":\"isvsvn\",\"type\":\"uint16\"}],\"internalType\":\"structEnclaveIdStruct.TcbObj\",\"name\":\"tcb\",\"type\":\"tuple\"},{\"internalType\":\"enumEnclaveIdStruct.EnclaveIdStatus\",\"name\":\"tcbStatus\",\"type\":\"uint8\"}],\"internalType\":\"structEnclaveIdStruct.TcbLevel[]\",\"name\":\"tcbLevels\",\"type\":\"tuple[]\"}],\"internalType\":\"structEnclaveIdStruct.EnclaveId\",\"name\":\"qeIdentityInput\",\"type\":\"tuple\"}],\"name\":\"configureQeIdentityJson\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"fmspc\",\"type\":\"string\"},{\"components\":[{\"internalType\":\"string\",\"name\":\"pceid\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"fmspc\",\"type\":\"string\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"pcesvn\",\"type\":\"uint256\"},{\"internalType\":\"uint256[]\",\"name\":\"sgxTcbCompSvnArr\",\"type\":\"uint256[]\"},{\"internalType\":\"enumTCBInfoStruct.TCBStatus\",\"name\":\"status\",\"type\":\"uint8\"}],\"internalType\":\"structTCBInfoStruct.TCBLevelObj[]\",\"name\":\"tcbLevels\",\"type\":\"tuple[]\"}],\"internalType\":\"structTCBInfoStruct.TCBInfo\",\"name\":\"tcbInfoInput\",\"type\":\"tuple\"}],\"name\":\"configureTcbInfoJson\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pemCertLib\",\"outputs\":[{\"internalType\":\"contractIPEMCertChainLib\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"qeIdentity\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"miscselect\",\"type\":\"bytes4\"},{\"internalType\":\"bytes4\",\"name\":\"miscselectMask\",\"type\":\"bytes4\"},{\"internalType\":\"uint16\",\"name\":\"isvprodid\",\"type\":\"uint16\"},{\"internalType\":\"bytes16\",\"name\":\"attributes\",\"type\":\"bytes16\"},{\"internalType\":\"bytes16\",\"name\":\"attributesMask\",\"type\":\"bytes16\"},{\"internalType\":\"bytes32\",\"name\":\"mrsigner\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"internalType\":\"bytes[]\",\"name\":\"serialNumBatch\",\"type\":\"bytes[]\"}],\"name\":\"removeRevokedCertSerialNum\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_mrEnclave\",\"type\":\"bytes32\"},{\"internalType\":\"bool\",\"name\":\"_trusted\",\"type\":\"bool\"}],\"name\":\"setMrEnclave\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_mrSigner\",\"type\":\"bytes32\"},{\"internalType\":\"bool\",\"name\":\"_trusted\",\"type\":\"bool\"}],\"name\":\"setMrSigner\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"sigVerifyLib\",\"outputs\":[{\"internalType\":\"contractISigVerifyLib\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"fmspc\",\"type\":\"string\"}],\"name\":\"tcbInfo\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"pceid\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"fmspc\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"toggleLocalReportCheck\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"verifyAttestation\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"components\":[{\"internalType\":\"bytes2\",\"name\":\"version\",\"type\":\"bytes2\"},{\"internalType\":\"bytes2\",\"name\":\"attestationKeyType\",\"type\":\"bytes2\"},{\"internalType\":\"bytes4\",\"name\":\"teeType\",\"type\":\"bytes4\"},{\"internalType\":\"bytes2\",\"name\":\"qeSvn\",\"type\":\"bytes2\"},{\"internalType\":\"bytes2\",\"name\":\"pceSvn\",\"type\":\"bytes2\"},{\"internalType\":\"bytes16\",\"name\":\"qeVendorId\",\"type\":\"bytes16\"},{\"internalType\":\"bytes20\",\"name\":\"userData\",\"type\":\"bytes20\"}],\"internalType\":\"structV3Struct.Header\",\"name\":\"header\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bytes16\",\"name\":\"cpuSvn\",\"type\":\"bytes16\"},{\"internalType\":\"bytes4\",\"name\":\"miscSelect\",\"type\":\"bytes4\"},{\"internalType\":\"bytes28\",\"name\":\"reserved1\",\"type\":\"bytes28\"},{\"internalType\":\"bytes16\",\"name\":\"attributes\",\"type\":\"bytes16\"},{\"internalType\":\"bytes32\",\"name\":\"mrEnclave\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"reserved2\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"mrSigner\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"reserved3\",\"type\":\"bytes\"},{\"internalType\":\"uint16\",\"name\":\"isvProdId\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"isvSvn\",\"type\":\"uint16\"},{\"internalType\":\"bytes\",\"name\":\"reserved4\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"reportData\",\"type\":\"bytes\"}],\"internalType\":\"structV3Struct.EnclaveReport\",\"name\":\"localEnclaveReport\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bytes\",\"name\":\"ecdsa256BitSignature\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"ecdsaAttestationKey\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"bytes16\",\"name\":\"cpuSvn\",\"type\":\"bytes16\"},{\"internalType\":\"bytes4\",\"name\":\"miscSelect\",\"type\":\"bytes4\"},{\"internalType\":\"bytes28\",\"name\":\"reserved1\",\"type\":\"bytes28\"},{\"internalType\":\"bytes16\",\"name\":\"attributes\",\"type\":\"bytes16\"},{\"internalType\":\"bytes32\",\"name\":\"mrEnclave\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"reserved2\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"mrSigner\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"reserved3\",\"type\":\"bytes\"},{\"internalType\":\"uint16\",\"name\":\"isvProdId\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"isvSvn\",\"type\":\"uint16\"},{\"internalType\":\"bytes\",\"name\":\"reserved4\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"reportData\",\"type\":\"bytes\"}],\"internalType\":\"structV3Struct.EnclaveReport\",\"name\":\"pckSignedQeReport\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"qeReportSignature\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"uint16\",\"name\":\"parsedDataSize\",\"type\":\"uint16\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"internalType\":\"structV3Struct.QEAuthData\",\"name\":\"qeAuthData\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint16\",\"name\":\"certType\",\"type\":\"uint16\"},{\"internalType\":\"uint32\",\"name\":\"certDataSize\",\"type\":\"uint32\"},{\"internalType\":\"bytes[3]\",\"name\":\"decodedCertDataArray\",\"type\":\"bytes[3]\"}],\"internalType\":\"structV3Struct.CertificationData\",\"name\":\"certification\",\"type\":\"tuple\"}],\"internalType\":\"structV3Struct.ECDSAQuoteV3AuthData\",\"name\":\"v3AuthData\",\"type\":\"tuple\"}],\"internalType\":\"structV3Struct.ParsedV3Quote\",\"name\":\"v3quote\",\"type\":\"tuple\"}],\"name\":\"verifyParsedQuote\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// AutomataDcapV3AttestationABI is the input ABI used to generate the binding from.
// Deprecated: Use AutomataDcapV3AttestationMetaData.ABI instead.
var AutomataDcapV3AttestationABI = AutomataDcapV3AttestationMetaData.ABI

// AutomataDcapV3Attestation is an auto generated Go binding around an Ethereum contract.
type AutomataDcapV3Attestation struct {
	AutomataDcapV3AttestationCaller     // Read-only binding to the contract
	AutomataDcapV3AttestationTransactor // Write-only binding to the contract
	AutomataDcapV3AttestationFilterer   // Log filterer for contract events
}

// AutomataDcapV3AttestationCaller is an auto generated read-only Go binding around an Ethereum contract.
type AutomataDcapV3AttestationCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AutomataDcapV3AttestationTransactor is an auto generated write-only Go binding around an Ethereum contract.
type AutomataDcapV3AttestationTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AutomataDcapV3AttestationFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AutomataDcapV3AttestationFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AutomataDcapV3AttestationSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AutomataDcapV3AttestationSession struct {
	Contract     *AutomataDcapV3Attestation // Generic contract binding to set the session for
	CallOpts     bind.CallOpts              // Call options to use throughout this session
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// AutomataDcapV3AttestationCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AutomataDcapV3AttestationCallerSession struct {
	Contract *AutomataDcapV3AttestationCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                    // Call options to use throughout this session
}

// AutomataDcapV3AttestationTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AutomataDcapV3AttestationTransactorSession struct {
	Contract     *AutomataDcapV3AttestationTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                    // Transaction auth options to use throughout this session
}

// AutomataDcapV3AttestationRaw is an auto generated low-level Go binding around an Ethereum contract.
type AutomataDcapV3AttestationRaw struct {
	Contract *AutomataDcapV3Attestation // Generic contract binding to access the raw methods on
}

// AutomataDcapV3AttestationCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AutomataDcapV3AttestationCallerRaw struct {
	Contract *AutomataDcapV3AttestationCaller // Generic read-only contract binding to access the raw methods on
}

// AutomataDcapV3AttestationTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AutomataDcapV3AttestationTransactorRaw struct {
	Contract *AutomataDcapV3AttestationTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAutomataDcapV3Attestation creates a new instance of AutomataDcapV3Attestation, bound to a specific deployed contract.
func NewAutomataDcapV3Attestation(address common.Address, backend bind.ContractBackend) (*AutomataDcapV3Attestation, error) {
	contract, err := bindAutomataDcapV3Attestation(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &AutomataDcapV3Attestation{AutomataDcapV3AttestationCaller: AutomataDcapV3AttestationCaller{contract: contract}, AutomataDcapV3AttestationTransactor: AutomataDcapV3AttestationTransactor{contract: contract}, AutomataDcapV3AttestationFilterer: AutomataDcapV3AttestationFilterer{contract: contract}}, nil
}

// NewAutomataDcapV3AttestationCaller creates a new read-only instance of AutomataDcapV3Attestation, bound to a specific deployed contract.
func NewAutomataDcapV3AttestationCaller(address common.Address, caller bind.ContractCaller) (*AutomataDcapV3AttestationCaller, error) {
	contract, err := bindAutomataDcapV3Attestation(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AutomataDcapV3AttestationCaller{contract: contract}, nil
}

// NewAutomataDcapV3AttestationTransactor creates a new write-only instance of AutomataDcapV3Attestation, bound to a specific deployed contract.
func NewAutomataDcapV3AttestationTransactor(address common.Address, transactor bind.ContractTransactor) (*AutomataDcapV3AttestationTransactor, error) {
	contract, err := bindAutomataDcapV3Attestation(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AutomataDcapV3AttestationTransactor{contract: contract}, nil
}

// NewAutomataDcapV3AttestationFilterer creates a new log filterer instance of AutomataDcapV3Attestation, bound to a specific deployed contract.
func NewAutomataDcapV3AttestationFilterer(address common.Address, filterer bind.ContractFilterer) (*AutomataDcapV3AttestationFilterer, error) {
	contract, err := bindAutomataDcapV3Attestation(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AutomataDcapV3AttestationFilterer{contract: contract}, nil
}

// bindAutomataDcapV3Attestation binds a generic wrapper to an already deployed contract.
func bindAutomataDcapV3Attestation(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := AutomataDcapV3AttestationMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AutomataDcapV3Attestation *AutomataDcapV3AttestationRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AutomataDcapV3Attestation.Contract.AutomataDcapV3AttestationCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AutomataDcapV3Attestation *AutomataDcapV3AttestationRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AutomataDcapV3Attestation.Contract.AutomataDcapV3AttestationTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AutomataDcapV3Attestation *AutomataDcapV3AttestationRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AutomataDcapV3Attestation.Contract.AutomataDcapV3AttestationTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AutomataDcapV3Attestation *AutomataDcapV3AttestationCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AutomataDcapV3Attestation.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AutomataDcapV3Attestation *AutomataDcapV3AttestationTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AutomataDcapV3Attestation.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AutomataDcapV3Attestation *AutomataDcapV3AttestationTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AutomataDcapV3Attestation.Contract.contract.Transact(opts, method, params...)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_AutomataDcapV3Attestation *AutomataDcapV3AttestationCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AutomataDcapV3Attestation.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_AutomataDcapV3Attestation *AutomataDcapV3AttestationSession) Owner() (common.Address, error) {
	return _AutomataDcapV3Attestation.Contract.Owner(&_AutomataDcapV3Attestation.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_AutomataDcapV3Attestation *AutomataDcapV3AttestationCallerSession) Owner() (common.Address, error) {
	return _AutomataDcapV3Attestation.Contract.Owner(&_AutomataDcapV3Attestation.CallOpts)
}

// PemCertLib is a free data retrieval call binding the contract method 0x01d711f4.
//
// Solidity: function pemCertLib() view returns(address)
func (_AutomataDcapV3Attestation *AutomataDcapV3AttestationCaller) PemCertLib(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AutomataDcapV3Attestation.contract.Call(opts, &out, "pemCertLib")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PemCertLib is a free data retrieval call binding the contract method 0x01d711f4.
//
// Solidity: function pemCertLib() view returns(address)
func (_AutomataDcapV3Attestation *AutomataDcapV3AttestationSession) PemCertLib() (common.Address, error) {
	return _AutomataDcapV3Attestation.Contract.PemCertLib(&_AutomataDcapV3Attestation.CallOpts)
}

// PemCertLib is a free data retrieval call binding the contract method 0x01d711f4.
//
// Solidity: function pemCertLib() view returns(address)
func (_AutomataDcapV3Attestation *AutomataDcapV3AttestationCallerSession) PemCertLib() (common.Address, error) {
	return _AutomataDcapV3Attestation.Contract.PemCertLib(&_AutomataDcapV3Attestation.CallOpts)
}

// QeIdentity is a free data retrieval call binding the contract method 0xb684252f.
//
// Solidity: function qeIdentity() view returns(bytes4 miscselect, bytes4 miscselectMask, uint16 isvprodid, bytes16 attributes, bytes16 attributesMask, bytes32 mrsigner)
func (_AutomataDcapV3Attestation *AutomataDcapV3AttestationCaller) QeIdentity(opts *bind.CallOpts) (struct {
	Miscselect     [4]byte
	MiscselectMask [4]byte
	Isvprodid      uint16
	Attributes     [16]byte
	AttributesMask [16]byte
	Mrsigner       [32]byte
}, error) {
	var out []interface{}
	err := _AutomataDcapV3Attestation.contract.Call(opts, &out, "qeIdentity")

	outstruct := new(struct {
		Miscselect     [4]byte
		MiscselectMask [4]byte
		Isvprodid      uint16
		Attributes     [16]byte
		AttributesMask [16]byte
		Mrsigner       [32]byte
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Miscselect = *abi.ConvertType(out[0], new([4]byte)).(*[4]byte)
	outstruct.MiscselectMask = *abi.ConvertType(out[1], new([4]byte)).(*[4]byte)
	outstruct.Isvprodid = *abi.ConvertType(out[2], new(uint16)).(*uint16)
	outstruct.Attributes = *abi.ConvertType(out[3], new([16]byte)).(*[16]byte)
	outstruct.AttributesMask = *abi.ConvertType(out[4], new([16]byte)).(*[16]byte)
	outstruct.Mrsigner = *abi.ConvertType(out[5], new([32]byte)).(*[32]byte)

	return *outstruct, err

}

// QeIdentity is a free data retrieval call binding the contract method 0xb684252f.
//
// Solidity: function qeIdentity() view returns(bytes4 miscselect, bytes4 miscselectMask, uint16 isvprodid, bytes16 attributes, bytes16 attributesMask, bytes32 mrsigner)
func (_AutomataDcapV3Attestation *AutomataDcapV3AttestationSession) QeIdentity() (struct {
	Miscselect     [4]byte
	MiscselectMask [4]byte
	Isvprodid      uint16
	Attributes     [16]byte
	AttributesMask [16]byte
	Mrsigner       [32]byte
}, error) {
	return _AutomataDcapV3Attestation.Contract.QeIdentity(&_AutomataDcapV3Attestation.CallOpts)
}

// QeIdentity is a free data retrieval call binding the contract method 0xb684252f.
//
// Solidity: function qeIdentity() view returns(bytes4 miscselect, bytes4 miscselectMask, uint16 isvprodid, bytes16 attributes, bytes16 attributesMask, bytes32 mrsigner)
func (_AutomataDcapV3Attestation *AutomataDcapV3AttestationCallerSession) QeIdentity() (struct {
	Miscselect     [4]byte
	MiscselectMask [4]byte
	Isvprodid      uint16
	Attributes     [16]byte
	AttributesMask [16]byte
	Mrsigner       [32]byte
}, error) {
	return _AutomataDcapV3Attestation.Contract.QeIdentity(&_AutomataDcapV3Attestation.CallOpts)
}

// SigVerifyLib is a free data retrieval call binding the contract method 0x0d23d71b.
//
// Solidity: function sigVerifyLib() view returns(address)
func (_AutomataDcapV3Attestation *AutomataDcapV3AttestationCaller) SigVerifyLib(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AutomataDcapV3Attestation.contract.Call(opts, &out, "sigVerifyLib")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// SigVerifyLib is a free data retrieval call binding the contract method 0x0d23d71b.
//
// Solidity: function sigVerifyLib() view returns(address)
func (_AutomataDcapV3Attestation *AutomataDcapV3AttestationSession) SigVerifyLib() (common.Address, error) {
	return _AutomataDcapV3Attestation.Contract.SigVerifyLib(&_AutomataDcapV3Attestation.CallOpts)
}

// SigVerifyLib is a free data retrieval call binding the contract method 0x0d23d71b.
//
// Solidity: function sigVerifyLib() view returns(address)
func (_AutomataDcapV3Attestation *AutomataDcapV3AttestationCallerSession) SigVerifyLib() (common.Address, error) {
	return _AutomataDcapV3Attestation.Contract.SigVerifyLib(&_AutomataDcapV3Attestation.CallOpts)
}

// TcbInfo is a free data retrieval call binding the contract method 0x4c0977a9.
//
// Solidity: function tcbInfo(string fmspc) view returns(string pceid, string fmspc)
func (_AutomataDcapV3Attestation *AutomataDcapV3AttestationCaller) TcbInfo(opts *bind.CallOpts, fmspc string) (struct {
	Pceid string
	Fmspc string
}, error) {
	var out []interface{}
	err := _AutomataDcapV3Attestation.contract.Call(opts, &out, "tcbInfo", fmspc)

	outstruct := new(struct {
		Pceid string
		Fmspc string
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Pceid = *abi.ConvertType(out[0], new(string)).(*string)
	outstruct.Fmspc = *abi.ConvertType(out[1], new(string)).(*string)

	return *outstruct, err

}

// TcbInfo is a free data retrieval call binding the contract method 0x4c0977a9.
//
// Solidity: function tcbInfo(string fmspc) view returns(string pceid, string fmspc)
func (_AutomataDcapV3Attestation *AutomataDcapV3AttestationSession) TcbInfo(fmspc string) (struct {
	Pceid string
	Fmspc string
}, error) {
	return _AutomataDcapV3Attestation.Contract.TcbInfo(&_AutomataDcapV3Attestation.CallOpts, fmspc)
}

// TcbInfo is a free data retrieval call binding the contract method 0x4c0977a9.
//
// Solidity: function tcbInfo(string fmspc) view returns(string pceid, string fmspc)
func (_AutomataDcapV3Attestation *AutomataDcapV3AttestationCallerSession) TcbInfo(fmspc string) (struct {
	Pceid string
	Fmspc string
}, error) {
	return _AutomataDcapV3Attestation.Contract.TcbInfo(&_AutomataDcapV3Attestation.CallOpts, fmspc)
}

// VerifyAttestation is a free data retrieval call binding the contract method 0x769d87e7.
//
// Solidity: function verifyAttestation(bytes data) view returns(bool, bytes)
func (_AutomataDcapV3Attestation *AutomataDcapV3AttestationCaller) VerifyAttestation(opts *bind.CallOpts, data []byte) (bool, []byte, error) {
	var out []interface{}
	err := _AutomataDcapV3Attestation.contract.Call(opts, &out, "verifyAttestation", data)

	if err != nil {
		return *new(bool), *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)
	out1 := *abi.ConvertType(out[1], new([]byte)).(*[]byte)

	return out0, out1, err

}

// VerifyAttestation is a free data retrieval call binding the contract method 0x769d87e7.
//
// Solidity: function verifyAttestation(bytes data) view returns(bool, bytes)
func (_AutomataDcapV3Attestation *AutomataDcapV3AttestationSession) VerifyAttestation(data []byte) (bool, []byte, error) {
	return _AutomataDcapV3Attestation.Contract.VerifyAttestation(&_AutomataDcapV3Attestation.CallOpts, data)
}

// VerifyAttestation is a free data retrieval call binding the contract method 0x769d87e7.
//
// Solidity: function verifyAttestation(bytes data) view returns(bool, bytes)
func (_AutomataDcapV3Attestation *AutomataDcapV3AttestationCallerSession) VerifyAttestation(data []byte) (bool, []byte, error) {
	return _AutomataDcapV3Attestation.Contract.VerifyAttestation(&_AutomataDcapV3Attestation.CallOpts, data)
}

// VerifyParsedQuote is a free data retrieval call binding the contract method 0x089a168f.
//
// Solidity: function verifyParsedQuote(((bytes2,bytes2,bytes4,bytes2,bytes2,bytes16,bytes20),(bytes16,bytes4,bytes28,bytes16,bytes32,bytes32,bytes32,bytes,uint16,uint16,bytes,bytes),(bytes,bytes,(bytes16,bytes4,bytes28,bytes16,bytes32,bytes32,bytes32,bytes,uint16,uint16,bytes,bytes),bytes,(uint16,bytes),(uint16,uint32,bytes[3]))) v3quote) view returns(bool, bytes)
func (_AutomataDcapV3Attestation *AutomataDcapV3AttestationCaller) VerifyParsedQuote(opts *bind.CallOpts, v3quote V3StructParsedV3Quote) (bool, []byte, error) {
	var out []interface{}
	err := _AutomataDcapV3Attestation.contract.Call(opts, &out, "verifyParsedQuote", v3quote)

	if err != nil {
		return *new(bool), *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)
	out1 := *abi.ConvertType(out[1], new([]byte)).(*[]byte)

	return out0, out1, err

}

// VerifyParsedQuote is a free data retrieval call binding the contract method 0x089a168f.
//
// Solidity: function verifyParsedQuote(((bytes2,bytes2,bytes4,bytes2,bytes2,bytes16,bytes20),(bytes16,bytes4,bytes28,bytes16,bytes32,bytes32,bytes32,bytes,uint16,uint16,bytes,bytes),(bytes,bytes,(bytes16,bytes4,bytes28,bytes16,bytes32,bytes32,bytes32,bytes,uint16,uint16,bytes,bytes),bytes,(uint16,bytes),(uint16,uint32,bytes[3]))) v3quote) view returns(bool, bytes)
func (_AutomataDcapV3Attestation *AutomataDcapV3AttestationSession) VerifyParsedQuote(v3quote V3StructParsedV3Quote) (bool, []byte, error) {
	return _AutomataDcapV3Attestation.Contract.VerifyParsedQuote(&_AutomataDcapV3Attestation.CallOpts, v3quote)
}

// VerifyParsedQuote is a free data retrieval call binding the contract method 0x089a168f.
//
// Solidity: function verifyParsedQuote(((bytes2,bytes2,bytes4,bytes2,bytes2,bytes16,bytes20),(bytes16,bytes4,bytes28,bytes16,bytes32,bytes32,bytes32,bytes,uint16,uint16,bytes,bytes),(bytes,bytes,(bytes16,bytes4,bytes28,bytes16,bytes32,bytes32,bytes32,bytes,uint16,uint16,bytes,bytes),bytes,(uint16,bytes),(uint16,uint32,bytes[3]))) v3quote) view returns(bool, bytes)
func (_AutomataDcapV3Attestation *AutomataDcapV3AttestationCallerSession) VerifyParsedQuote(v3quote V3StructParsedV3Quote) (bool, []byte, error) {
	return _AutomataDcapV3Attestation.Contract.VerifyParsedQuote(&_AutomataDcapV3Attestation.CallOpts, v3quote)
}

// AddRevokedCertSerialNum is a paid mutator transaction binding the contract method 0x610de480.
//
// Solidity: function addRevokedCertSerialNum(uint256 index, bytes[] serialNumBatch) returns()
func (_AutomataDcapV3Attestation *AutomataDcapV3AttestationTransactor) AddRevokedCertSerialNum(opts *bind.TransactOpts, index *big.Int, serialNumBatch [][]byte) (*types.Transaction, error) {
	return _AutomataDcapV3Attestation.contract.Transact(opts, "addRevokedCertSerialNum", index, serialNumBatch)
}

// AddRevokedCertSerialNum is a paid mutator transaction binding the contract method 0x610de480.
//
// Solidity: function addRevokedCertSerialNum(uint256 index, bytes[] serialNumBatch) returns()
func (_AutomataDcapV3Attestation *AutomataDcapV3AttestationSession) AddRevokedCertSerialNum(index *big.Int, serialNumBatch [][]byte) (*types.Transaction, error) {
	return _AutomataDcapV3Attestation.Contract.AddRevokedCertSerialNum(&_AutomataDcapV3Attestation.TransactOpts, index, serialNumBatch)
}

// AddRevokedCertSerialNum is a paid mutator transaction binding the contract method 0x610de480.
//
// Solidity: function addRevokedCertSerialNum(uint256 index, bytes[] serialNumBatch) returns()
func (_AutomataDcapV3Attestation *AutomataDcapV3AttestationTransactorSession) AddRevokedCertSerialNum(index *big.Int, serialNumBatch [][]byte) (*types.Transaction, error) {
	return _AutomataDcapV3Attestation.Contract.AddRevokedCertSerialNum(&_AutomataDcapV3Attestation.TransactOpts, index, serialNumBatch)
}

// ConfigureQeIdentityJson is a paid mutator transaction binding the contract method 0x123ac29e.
//
// Solidity: function configureQeIdentityJson((bytes4,bytes4,uint16,bytes16,bytes16,bytes32,((uint16),uint8)[]) qeIdentityInput) returns()
func (_AutomataDcapV3Attestation *AutomataDcapV3AttestationTransactor) ConfigureQeIdentityJson(opts *bind.TransactOpts, qeIdentityInput EnclaveIdStructEnclaveId) (*types.Transaction, error) {
	return _AutomataDcapV3Attestation.contract.Transact(opts, "configureQeIdentityJson", qeIdentityInput)
}

// ConfigureQeIdentityJson is a paid mutator transaction binding the contract method 0x123ac29e.
//
// Solidity: function configureQeIdentityJson((bytes4,bytes4,uint16,bytes16,bytes16,bytes32,((uint16),uint8)[]) qeIdentityInput) returns()
func (_AutomataDcapV3Attestation *AutomataDcapV3AttestationSession) ConfigureQeIdentityJson(qeIdentityInput EnclaveIdStructEnclaveId) (*types.Transaction, error) {
	return _AutomataDcapV3Attestation.Contract.ConfigureQeIdentityJson(&_AutomataDcapV3Attestation.TransactOpts, qeIdentityInput)
}

// ConfigureQeIdentityJson is a paid mutator transaction binding the contract method 0x123ac29e.
//
// Solidity: function configureQeIdentityJson((bytes4,bytes4,uint16,bytes16,bytes16,bytes32,((uint16),uint8)[]) qeIdentityInput) returns()
func (_AutomataDcapV3Attestation *AutomataDcapV3AttestationTransactorSession) ConfigureQeIdentityJson(qeIdentityInput EnclaveIdStructEnclaveId) (*types.Transaction, error) {
	return _AutomataDcapV3Attestation.Contract.ConfigureQeIdentityJson(&_AutomataDcapV3Attestation.TransactOpts, qeIdentityInput)
}

// ConfigureTcbInfoJson is a paid mutator transaction binding the contract method 0x0581f14e.
//
// Solidity: function configureTcbInfoJson(string fmspc, (string,string,(uint256,uint256[],uint8)[]) tcbInfoInput) returns()
func (_AutomataDcapV3Attestation *AutomataDcapV3AttestationTransactor) ConfigureTcbInfoJson(opts *bind.TransactOpts, fmspc string, tcbInfoInput TCBInfoStructTCBInfo) (*types.Transaction, error) {
	return _AutomataDcapV3Attestation.contract.Transact(opts, "configureTcbInfoJson", fmspc, tcbInfoInput)
}

// ConfigureTcbInfoJson is a paid mutator transaction binding the contract method 0x0581f14e.
//
// Solidity: function configureTcbInfoJson(string fmspc, (string,string,(uint256,uint256[],uint8)[]) tcbInfoInput) returns()
func (_AutomataDcapV3Attestation *AutomataDcapV3AttestationSession) ConfigureTcbInfoJson(fmspc string, tcbInfoInput TCBInfoStructTCBInfo) (*types.Transaction, error) {
	return _AutomataDcapV3Attestation.Contract.ConfigureTcbInfoJson(&_AutomataDcapV3Attestation.TransactOpts, fmspc, tcbInfoInput)
}

// ConfigureTcbInfoJson is a paid mutator transaction binding the contract method 0x0581f14e.
//
// Solidity: function configureTcbInfoJson(string fmspc, (string,string,(uint256,uint256[],uint8)[]) tcbInfoInput) returns()
func (_AutomataDcapV3Attestation *AutomataDcapV3AttestationTransactorSession) ConfigureTcbInfoJson(fmspc string, tcbInfoInput TCBInfoStructTCBInfo) (*types.Transaction, error) {
	return _AutomataDcapV3Attestation.Contract.ConfigureTcbInfoJson(&_AutomataDcapV3Attestation.TransactOpts, fmspc, tcbInfoInput)
}

// RemoveRevokedCertSerialNum is a paid mutator transaction binding the contract method 0x1f3be096.
//
// Solidity: function removeRevokedCertSerialNum(uint256 index, bytes[] serialNumBatch) returns()
func (_AutomataDcapV3Attestation *AutomataDcapV3AttestationTransactor) RemoveRevokedCertSerialNum(opts *bind.TransactOpts, index *big.Int, serialNumBatch [][]byte) (*types.Transaction, error) {
	return _AutomataDcapV3Attestation.contract.Transact(opts, "removeRevokedCertSerialNum", index, serialNumBatch)
}

// RemoveRevokedCertSerialNum is a paid mutator transaction binding the contract method 0x1f3be096.
//
// Solidity: function removeRevokedCertSerialNum(uint256 index, bytes[] serialNumBatch) returns()
func (_AutomataDcapV3Attestation *AutomataDcapV3AttestationSession) RemoveRevokedCertSerialNum(index *big.Int, serialNumBatch [][]byte) (*types.Transaction, error) {
	return _AutomataDcapV3Attestation.Contract.RemoveRevokedCertSerialNum(&_AutomataDcapV3Attestation.TransactOpts, index, serialNumBatch)
}

// RemoveRevokedCertSerialNum is a paid mutator transaction binding the contract method 0x1f3be096.
//
// Solidity: function removeRevokedCertSerialNum(uint256 index, bytes[] serialNumBatch) returns()
func (_AutomataDcapV3Attestation *AutomataDcapV3AttestationTransactorSession) RemoveRevokedCertSerialNum(index *big.Int, serialNumBatch [][]byte) (*types.Transaction, error) {
	return _AutomataDcapV3Attestation.Contract.RemoveRevokedCertSerialNum(&_AutomataDcapV3Attestation.TransactOpts, index, serialNumBatch)
}

// SetMrEnclave is a paid mutator transaction binding the contract method 0x3a343014.
//
// Solidity: function setMrEnclave(bytes32 _mrEnclave, bool _trusted) returns()
func (_AutomataDcapV3Attestation *AutomataDcapV3AttestationTransactor) SetMrEnclave(opts *bind.TransactOpts, _mrEnclave [32]byte, _trusted bool) (*types.Transaction, error) {
	return _AutomataDcapV3Attestation.contract.Transact(opts, "setMrEnclave", _mrEnclave, _trusted)
}

// SetMrEnclave is a paid mutator transaction binding the contract method 0x3a343014.
//
// Solidity: function setMrEnclave(bytes32 _mrEnclave, bool _trusted) returns()
func (_AutomataDcapV3Attestation *AutomataDcapV3AttestationSession) SetMrEnclave(_mrEnclave [32]byte, _trusted bool) (*types.Transaction, error) {
	return _AutomataDcapV3Attestation.Contract.SetMrEnclave(&_AutomataDcapV3Attestation.TransactOpts, _mrEnclave, _trusted)
}

// SetMrEnclave is a paid mutator transaction binding the contract method 0x3a343014.
//
// Solidity: function setMrEnclave(bytes32 _mrEnclave, bool _trusted) returns()
func (_AutomataDcapV3Attestation *AutomataDcapV3AttestationTransactorSession) SetMrEnclave(_mrEnclave [32]byte, _trusted bool) (*types.Transaction, error) {
	return _AutomataDcapV3Attestation.Contract.SetMrEnclave(&_AutomataDcapV3Attestation.TransactOpts, _mrEnclave, _trusted)
}

// SetMrSigner is a paid mutator transaction binding the contract method 0xe2e28294.
//
// Solidity: function setMrSigner(bytes32 _mrSigner, bool _trusted) returns()
func (_AutomataDcapV3Attestation *AutomataDcapV3AttestationTransactor) SetMrSigner(opts *bind.TransactOpts, _mrSigner [32]byte, _trusted bool) (*types.Transaction, error) {
	return _AutomataDcapV3Attestation.contract.Transact(opts, "setMrSigner", _mrSigner, _trusted)
}

// SetMrSigner is a paid mutator transaction binding the contract method 0xe2e28294.
//
// Solidity: function setMrSigner(bytes32 _mrSigner, bool _trusted) returns()
func (_AutomataDcapV3Attestation *AutomataDcapV3AttestationSession) SetMrSigner(_mrSigner [32]byte, _trusted bool) (*types.Transaction, error) {
	return _AutomataDcapV3Attestation.Contract.SetMrSigner(&_AutomataDcapV3Attestation.TransactOpts, _mrSigner, _trusted)
}

// SetMrSigner is a paid mutator transaction binding the contract method 0xe2e28294.
//
// Solidity: function setMrSigner(bytes32 _mrSigner, bool _trusted) returns()
func (_AutomataDcapV3Attestation *AutomataDcapV3AttestationTransactorSession) SetMrSigner(_mrSigner [32]byte, _trusted bool) (*types.Transaction, error) {
	return _AutomataDcapV3Attestation.Contract.SetMrSigner(&_AutomataDcapV3Attestation.TransactOpts, _mrSigner, _trusted)
}

// ToggleLocalReportCheck is a paid mutator transaction binding the contract method 0x83801580.
//
// Solidity: function toggleLocalReportCheck() returns()
func (_AutomataDcapV3Attestation *AutomataDcapV3AttestationTransactor) ToggleLocalReportCheck(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AutomataDcapV3Attestation.contract.Transact(opts, "toggleLocalReportCheck")
}

// ToggleLocalReportCheck is a paid mutator transaction binding the contract method 0x83801580.
//
// Solidity: function toggleLocalReportCheck() returns()
func (_AutomataDcapV3Attestation *AutomataDcapV3AttestationSession) ToggleLocalReportCheck() (*types.Transaction, error) {
	return _AutomataDcapV3Attestation.Contract.ToggleLocalReportCheck(&_AutomataDcapV3Attestation.TransactOpts)
}

// ToggleLocalReportCheck is a paid mutator transaction binding the contract method 0x83801580.
//
// Solidity: function toggleLocalReportCheck() returns()
func (_AutomataDcapV3Attestation *AutomataDcapV3AttestationTransactorSession) ToggleLocalReportCheck() (*types.Transaction, error) {
	return _AutomataDcapV3Attestation.Contract.ToggleLocalReportCheck(&_AutomataDcapV3Attestation.TransactOpts)
}
