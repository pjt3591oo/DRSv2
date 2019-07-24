// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package abi

import (
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
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = abi.U256
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// BufferABI is the input ABI used to generate the binding from.
const BufferABI = "[]"

// BufferBin is the compiled bytecode used for deploying new contracts.
const BufferBin = `0x604c6025600b82828239805160001a6073141515601857fe5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea165627a7a72305820b06f2d629060e4904bababa4f7df0b91f31a7e69e1d738323de555e24786f9750029`

// DeployBuffer deploys a new Ethereum contract, binding an instance of Buffer to it.
func DeployBuffer(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Buffer, error) {
	parsed, err := abi.JSON(strings.NewReader(BufferABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(BufferBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Buffer{BufferCaller: BufferCaller{contract: contract}, BufferTransactor: BufferTransactor{contract: contract}, BufferFilterer: BufferFilterer{contract: contract}}, nil
}

// Buffer is an auto generated Go binding around an Ethereum contract.
type Buffer struct {
	BufferCaller     // Read-only binding to the contract
	BufferTransactor // Write-only binding to the contract
	BufferFilterer   // Log filterer for contract events
}

// BufferCaller is an auto generated read-only Go binding around an Ethereum contract.
type BufferCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BufferTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BufferTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BufferFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BufferFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BufferSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BufferSession struct {
	Contract     *Buffer           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BufferCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BufferCallerSession struct {
	Contract *BufferCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// BufferTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BufferTransactorSession struct {
	Contract     *BufferTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BufferRaw is an auto generated low-level Go binding around an Ethereum contract.
type BufferRaw struct {
	Contract *Buffer // Generic contract binding to access the raw methods on
}

// BufferCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BufferCallerRaw struct {
	Contract *BufferCaller // Generic read-only contract binding to access the raw methods on
}

// BufferTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BufferTransactorRaw struct {
	Contract *BufferTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBuffer creates a new instance of Buffer, bound to a specific deployed contract.
func NewBuffer(address common.Address, backend bind.ContractBackend) (*Buffer, error) {
	contract, err := bindBuffer(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Buffer{BufferCaller: BufferCaller{contract: contract}, BufferTransactor: BufferTransactor{contract: contract}, BufferFilterer: BufferFilterer{contract: contract}}, nil
}

// NewBufferCaller creates a new read-only instance of Buffer, bound to a specific deployed contract.
func NewBufferCaller(address common.Address, caller bind.ContractCaller) (*BufferCaller, error) {
	contract, err := bindBuffer(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BufferCaller{contract: contract}, nil
}

// NewBufferTransactor creates a new write-only instance of Buffer, bound to a specific deployed contract.
func NewBufferTransactor(address common.Address, transactor bind.ContractTransactor) (*BufferTransactor, error) {
	contract, err := bindBuffer(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BufferTransactor{contract: contract}, nil
}

// NewBufferFilterer creates a new log filterer instance of Buffer, bound to a specific deployed contract.
func NewBufferFilterer(address common.Address, filterer bind.ContractFilterer) (*BufferFilterer, error) {
	contract, err := bindBuffer(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BufferFilterer{contract: contract}, nil
}

// bindBuffer binds a generic wrapper to an already deployed contract.
func bindBuffer(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(BufferABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Buffer *BufferRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Buffer.Contract.BufferCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Buffer *BufferRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Buffer.Contract.BufferTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Buffer *BufferRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Buffer.Contract.BufferTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Buffer *BufferCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Buffer.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Buffer *BufferTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Buffer.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Buffer *BufferTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Buffer.Contract.contract.Transact(opts, method, params...)
}

// CBORABI is the input ABI used to generate the binding from.
const CBORABI = "[]"

// CBORBin is the compiled bytecode used for deploying new contracts.
const CBORBin = `0x604c6025600b82828239805160001a6073141515601857fe5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea165627a7a723058208b4674eef169b7a2210f3206842ffc2cff57c580b249913d4265161cad9349290029`

// DeployCBOR deploys a new Ethereum contract, binding an instance of CBOR to it.
func DeployCBOR(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *CBOR, error) {
	parsed, err := abi.JSON(strings.NewReader(CBORABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(CBORBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &CBOR{CBORCaller: CBORCaller{contract: contract}, CBORTransactor: CBORTransactor{contract: contract}, CBORFilterer: CBORFilterer{contract: contract}}, nil
}

// CBOR is an auto generated Go binding around an Ethereum contract.
type CBOR struct {
	CBORCaller     // Read-only binding to the contract
	CBORTransactor // Write-only binding to the contract
	CBORFilterer   // Log filterer for contract events
}

// CBORCaller is an auto generated read-only Go binding around an Ethereum contract.
type CBORCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CBORTransactor is an auto generated write-only Go binding around an Ethereum contract.
type CBORTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CBORFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type CBORFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CBORSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type CBORSession struct {
	Contract     *CBOR             // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// CBORCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type CBORCallerSession struct {
	Contract *CBORCaller   // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// CBORTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type CBORTransactorSession struct {
	Contract     *CBORTransactor   // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// CBORRaw is an auto generated low-level Go binding around an Ethereum contract.
type CBORRaw struct {
	Contract *CBOR // Generic contract binding to access the raw methods on
}

// CBORCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type CBORCallerRaw struct {
	Contract *CBORCaller // Generic read-only contract binding to access the raw methods on
}

// CBORTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type CBORTransactorRaw struct {
	Contract *CBORTransactor // Generic write-only contract binding to access the raw methods on
}

// NewCBOR creates a new instance of CBOR, bound to a specific deployed contract.
func NewCBOR(address common.Address, backend bind.ContractBackend) (*CBOR, error) {
	contract, err := bindCBOR(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &CBOR{CBORCaller: CBORCaller{contract: contract}, CBORTransactor: CBORTransactor{contract: contract}, CBORFilterer: CBORFilterer{contract: contract}}, nil
}

// NewCBORCaller creates a new read-only instance of CBOR, bound to a specific deployed contract.
func NewCBORCaller(address common.Address, caller bind.ContractCaller) (*CBORCaller, error) {
	contract, err := bindCBOR(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &CBORCaller{contract: contract}, nil
}

// NewCBORTransactor creates a new write-only instance of CBOR, bound to a specific deployed contract.
func NewCBORTransactor(address common.Address, transactor bind.ContractTransactor) (*CBORTransactor, error) {
	contract, err := bindCBOR(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &CBORTransactor{contract: contract}, nil
}

// NewCBORFilterer creates a new log filterer instance of CBOR, bound to a specific deployed contract.
func NewCBORFilterer(address common.Address, filterer bind.ContractFilterer) (*CBORFilterer, error) {
	contract, err := bindCBOR(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &CBORFilterer{contract: contract}, nil
}

// bindCBOR binds a generic wrapper to an already deployed contract.
func bindCBOR(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(CBORABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CBOR *CBORRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _CBOR.Contract.CBORCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CBOR *CBORRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CBOR.Contract.CBORTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CBOR *CBORRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CBOR.Contract.CBORTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CBOR *CBORCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _CBOR.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CBOR *CBORTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CBOR.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CBOR *CBORTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CBOR.Contract.contract.Transact(opts, method, params...)
}

// DigitalReserveSystemABI is the input ABI used to generate the binding from.
const DigitalReserveSystemABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"currency\",\"type\":\"bytes32\"},{\"name\":\"newPrice\",\"type\":\"uint256\"}],\"name\":\"setPrice\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"creditIssuanceFee\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getCollectedFee\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"peggedValue\",\"type\":\"uint256\"},{\"name\":\"peggedCurrency\",\"type\":\"bytes32\"},{\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"checkRequiredCollateralWithFee\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"}],\"name\":\"trustedPartners\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"currency\",\"type\":\"bytes32\"}],\"name\":\"getPrice\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"renounceWhitelistAdmin\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"assetName\",\"type\":\"bytes32\"},{\"name\":\"issuer\",\"type\":\"address\"}],\"name\":\"collateralOf\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"peggedValue\",\"type\":\"uint256\"},{\"name\":\"peggedCurrency\",\"type\":\"bytes32\"},{\"name\":\"assetName\",\"type\":\"bytes32\"},{\"name\":\"amount\",\"type\":\"uint256\"},{\"name\":\"stellarAddr\",\"type\":\"string\"},{\"name\":\"hashSecret\",\"type\":\"bytes32\"}],\"name\":\"mint\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"prices\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"priceFeeders\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"assetName\",\"type\":\"bytes32\"},{\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"redeem\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"setTrustedPartner\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"account\",\"type\":\"address\"}],\"name\":\"addWhitelistAdmin\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"issuer\",\"type\":\"address\"},{\"name\":\"assetName\",\"type\":\"bytes32\"},{\"name\":\"stellarIssuerAddress\",\"type\":\"string\"}],\"name\":\"__callbackConfirmIssueCredit\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"currency\",\"type\":\"bytes32\"},{\"name\":\"priceFeederAddr\",\"type\":\"address\"}],\"name\":\"setPriceFeeder\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newFee\",\"type\":\"uint256\"}],\"name\":\"setCreditIssuanceFee\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"account\",\"type\":\"address\"}],\"name\":\"isWhitelistAdmin\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"withdrawFee\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"digitalCredits\",\"outputs\":[{\"name\":\"collateral\",\"type\":\"uint256\"},{\"name\":\"assetAmount\",\"type\":\"uint256\"},{\"name\":\"peggedValue\",\"type\":\"uint256\"},{\"name\":\"peggedCurrency\",\"type\":\"bytes32\"},{\"name\":\"stellarIssuerAddress\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"collectedFee\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"veloTokenAddr\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"assetName\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"collateral\",\"type\":\"uint256\"},{\"indexed\":true,\"name\":\"issuer\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"stellarAddr\",\"type\":\"string\"},{\"indexed\":false,\"name\":\"hashSecret\",\"type\":\"bytes32\"}],\"name\":\"Mint\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"account\",\"type\":\"address\"}],\"name\":\"WhitelistAdminAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"account\",\"type\":\"address\"}],\"name\":\"WhitelistAdminRemoved\",\"type\":\"event\"}]"

// DigitalReserveSystemBin is the compiled bytecode used for deploying new contracts.
const DigitalReserveSystemBin = `0x60806040523480156200001157600080fd5b50604051602080620018ed833981018060405260208110156200003357600080fd5b505162000047336200006d602090811b901c565b600180546001600160a01b0319166001600160a01b0392909216919091179055620001e8565b62000088816000620000bf60201b6200148e1790919060201c565b6040516001600160a01b038216907f22380c05984257a1cb900161c713dd71d39e74820f1aea43bd3f1bdd2096129990600090a250565b620000d182826200016360201b60201c565b156200013e57604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601f60248201527f526f6c65733a206163636f756e7420616c72656164792068617320726f6c6500604482015290519081900360640190fd5b6001600160a01b0316600090815260209190915260409020805460ff19166001179055565b60006001600160a01b0382161515620001c8576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401808060200182810382526022815260200180620018cb6022913960400191505060405180910390fd5b506001600160a01b03166000908152602091909152604090205460ff1690565b6116d380620001f86000396000f3fe60806040526004361061012a5760003560e01c806364b456ed116100ab578063a7aa2b791161006f578063a7aa2b7914610539578063aedb08d414610572578063bb5f747b1461059c578063be357616146105cf578063c574ccb4146105f9578063e811f50a1461064e5761012a565b806364b456ed14610395578063673da154146103db578063686984061461040b5780637362d9c81461043e578063747f43b6146104715761012a565b806331d98b3f116100f257806331d98b3f146102335780634c5a628c1461025d5780634e1bfbe7146102725780635762517e146102ab57806360846bc61461036b5761012a565b806310d8d74d1461012f5780631f6f207214610161578063271ec3051461018857806328893f5f1461019d57806329fc9ec8146101ec575b600080fd5b34801561013b57600080fd5b5061015f6004803603604081101561015257600080fd5b5080359060200135610663565b005b34801561016d57600080fd5b506101766106cd565b60408051918252519081900360200190f35b34801561019457600080fd5b506101766106d3565b3480156101a957600080fd5b506101d3600480360360608110156101c057600080fd5b50803590602081013590604001356106d9565b6040805192835260208301919091528051918290030190f35b3480156101f857600080fd5b5061021f6004803603602081101561020f57600080fd5b50356001600160a01b03166106f8565b604080519115158252519081900360200190f35b34801561023f57600080fd5b506101766004803603602081101561025657600080fd5b503561070d565b34801561026957600080fd5b5061015f61071f565b34801561027e57600080fd5b506101766004803603604081101561029557600080fd5b50803590602001356001600160a01b031661072a565b61015f600480360360c08110156102c157600080fd5b81359160208101359160408201359160608101359181019060a0810160808201356401000000008111156102f457600080fd5b82018360208201111561030657600080fd5b8035906020019184600183028401116401000000008311171561032857600080fd5b91908080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152509295505091359250610779915050565b34801561037757600080fd5b506101766004803603602081101561038e57600080fd5b5035610ce8565b3480156103a157600080fd5b506103bf600480360360208110156103b857600080fd5b5035610cfa565b604080516001600160a01b039092168252519081900360200190f35b3480156103e757600080fd5b5061015f600480360360408110156103fe57600080fd5b5080359060200135610d15565b34801561041757600080fd5b5061015f6004803603602081101561042e57600080fd5b50356001600160a01b0316610d19565b34801561044a57600080fd5b5061015f6004803603602081101561046157600080fd5b50356001600160a01b0316610d86565b34801561047d57600080fd5b5061015f6004803603606081101561049457600080fd5b6001600160a01b03823516916020810135918101906060810160408201356401000000008111156104c457600080fd5b8201836020820111156104d657600080fd5b803590602001918460018302840111640100000000831117156104f857600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250929550610ddb945050505050565b34801561054557600080fd5b5061015f6004803603604081101561055c57600080fd5b50803590602001356001600160a01b0316610ed9565b34801561057e57600080fd5b5061015f6004803603602081101561059557600080fd5b5035610f50565b3480156105a857600080fd5b5061021f600480360360208110156105bf57600080fd5b50356001600160a01b0316610f9e565b3480156105db57600080fd5b5061015f600480360360208110156105f257600080fd5b5035610fb0565b34801561060557600080fd5b506106236004803603602081101561061c57600080fd5b50356110ef565b6040805195865260208601949094528484019290925260608401526080830152519081900360a00190f35b34801561065a57600080fd5b5061017661111e565b6000828152600660205260409020546001600160a01b031633146106bb57604051600160e51b62461bcd0281526004018080602001828103825260268152602001806115136026913960400191505060405180910390fd5b60009182526007602052604090912055565b60025481565b60035490565b6000806000806106ea878787611124565b909890975095505050505050565b60046020526000908152604090205460ff1681565b60009081526007602052604090205490565b61072833611193565b565b604080516001600160a01b03831660601b602080830191909152603480830186905283518084039091018152605490920183528151918101919091206000908152600590915220545b92915050565b3360009081526004602052604090205460ff1615156001146107cf57604051600160e51b62461bcd0281526004018080602001828103825260318152602001806115a46031913960400191505060405180910390fd5b6000858152600660205260409020546001600160a01b0316151561082757604051600160e51b62461bcd02815260040180806020018281038252602f815260200180611679602f913960400191505060405180910390fd5b604080513360601b602080830191909152603480830188905283518084039091018152605490920183528151918101919091206000818152600590925291902060020154156108cf5786156108b057604051600160e51b62461bcd0281526004018080602001828103825260418152602001806115636041913960600191505060405180910390fd5b6000818152600560205260409020600281015460039091015490975095505b6000806108dd898988611124565b909250905081151561092357604051600160e51b62461bcd02815260040180806020018281038252602a815260200180611539602a913960400191505060405180910390fd5b60015460408051600160e01b6370a08231028152336004820152905184926001600160a01b0316916370a08231916024808301926020929190829003018186803b15801561097057600080fd5b505afa158015610984573d6000803e3d6000fd5b505050506040513d602081101561099a57600080fd5b5051116109f15760408051600160e51b62461bcd02815260206004820152600f60248201527f6e6f7420656e6f7567682056454c4f0000000000000000000000000000000000604482015290519081900360640190fd5b610a01828263ffffffff6111db16565b60015460408051600160e11b636eb1769f02815233600482015230602482015290516001600160a01b039092169163dd62ed3e91604480820192602092909190829003018186803b158015610a5557600080fd5b505afa158015610a69573d6000803e3d6000fd5b505050506040513d6020811015610a7f57600080fd5b505111610ad65760408051600160e51b62461bcd02815260206004820152601460248201527f6e6f7420656e6f75676820616c6c6f77616e6365000000000000000000000000604482015290519081900360640190fd5b6001546001600160a01b03166323b872dd3330610af9868663ffffffff6111db16565b6040805163ffffffff861660e01b81526001600160a01b0394851660048201529290931660248301526044820152905160648083019260209291908290030181600087803b158015610b4a57600080fd5b505af1158015610b5e573d6000803e3d6000fd5b505050506040513d6020811015610b7457600080fd5b50506040805160a0810182526000858152600560205291909120548190610ba1908563ffffffff6111db16565b8152600085815260056020908152604090912060010154910190610bcb908963ffffffff6111db16565b815260208082018c905260408083018c905260008781526005808452828220600481018054606097880152928a90529084528551815592850151600184015590840151600283015591830151600391820155608090920151905554610c36908263ffffffff6111db16565b6003556040518551869190819060208401908083835b60208310610c6b5780518252601f199092019160209182019101610c4c565b51815160001960209485036101000a01908116901991909116179052604080519490920184900384208d85529084018c9052838201889052606084018a905290519094503393507faa2a61bb785a72b9e47fb58125440d23bbad4b6c8dda3e17ab80714bf056e4e3928190036080019150a3505050505050505050565b60076020526000908152604090205481565b6006602052600090815260409020546001600160a01b031681565b5050565b610d2233610f9e565b1515610d6257604051600160e51b62461bcd0281526004018080602001828103825260408152602001806116396040913960400191505060405180910390fd5b6001600160a01b03166000908152600460205260409020805460ff19166001179055565b610d8f33610f9e565b1515610dcf57604051600160e51b62461bcd0281526004018080602001828103825260408152602001806116396040913960400191505060405180910390fd5b610dd88161123f565b50565b610de433610f9e565b1515610e2457604051600160e51b62461bcd0281526004018080602001828103825260408152602001806116396040913960400191505060405180910390fd5b604080513360601b602080830191909152603480830186905283518084039091018152605483019093528251928101929092208351909284926074019182918401908083835b60208310610e895780518252601f199092019160209182019101610e6a565b51815160209384036101000a60001901801990921691161790526040805192909401828103601f190183528452815191810191909120600096875260059091529190942060040155505050505050565b610ee233610f9e565b1515610f2257604051600160e51b62461bcd0281526004018080602001828103825260408152602001806116396040913960400191505060405180910390fd5b60009182526006602052604090912080546001600160a01b0319166001600160a01b03909216919091179055565b610f5933610f9e565b1515610f9957604051600160e51b62461bcd0281526004018080602001828103825260408152602001806116396040913960400191505060405180910390fd5b600255565b6000610773818363ffffffff61128716565b610fb933610f9e565b1515610ff957604051600160e51b62461bcd0281526004018080602001828103825260408152602001806116396040913960400191505060405180910390fd5b6003548111156110535760408051600160e51b62461bcd02815260206004820152601e60248201527f616d6f756e74206d757374203c3d20746f20636f6c6c65637465644665650000604482015290519081900360640190fd5b60015460408051600160e01b63a9059cbb0281523360048201526024810184905290516001600160a01b039092169163a9059cbb916044808201926020929091908290030181600087803b1580156110aa57600080fd5b505af11580156110be573d6000803e3d6000fd5b505050506040513d60208110156110d457600080fd5b50506003546110e9908263ffffffff6112f316565b60035550565b600560205260009081526040902080546001820154600283015460038401546004909401549293919290919085565b60035481565b6000828152600760205260408120548190819061115a90859061114e90899063ffffffff61135316565b9063ffffffff6113c216565b90506000611185612710611179600254856113c290919063ffffffff16565b9063ffffffff61135316565b919791965090945050505050565b6111a460008263ffffffff61142216565b6040516001600160a01b038216907f0a8eb35e5ca14b3d6f28e4abf2f128dbab231a58b56e89beb5d636115001e16590600090a250565b6000828201838110156112385760408051600160e51b62461bcd02815260206004820152601b60248201527f536166654d6174683a206164646974696f6e206f766572666c6f770000000000604482015290519081900360640190fd5b9392505050565b61125060008263ffffffff61148e16565b6040516001600160a01b038216907f22380c05984257a1cb900161c713dd71d39e74820f1aea43bd3f1bdd2096129990600090a250565b60006001600160a01b03821615156112d357604051600160e51b62461bcd0281526004018080602001828103825260228152602001806116176022913960400191505060405180910390fd5b506001600160a01b03166000908152602091909152604090205460ff1690565b60008282111561134d5760408051600160e51b62461bcd02815260206004820152601e60248201527f536166654d6174683a207375627472616374696f6e206f766572666c6f770000604482015290519081900360640190fd5b50900390565b60008115156113ac5760408051600160e51b62461bcd02815260206004820152601a60248201527f536166654d6174683a206469766973696f6e206279207a65726f000000000000604482015290519081900360640190fd5b600082848115156113b957fe5b04949350505050565b60008215156113d357506000610773565b8282028284828115156113e257fe5b041461123857604051600160e51b62461bcd0281526004018080602001828103825260218152602001806115f66021913960400191505060405180910390fd5b61142c8282611287565b151561146c57604051600160e51b62461bcd0281526004018080602001828103825260218152602001806115d56021913960400191505060405180910390fd5b6001600160a01b0316600090815260209190915260409020805460ff19169055565b6114988282611287565b156114ed5760408051600160e51b62461bcd02815260206004820152601f60248201527f526f6c65733a206163636f756e7420616c72656164792068617320726f6c6500604482015290519081900360640190fd5b6001600160a01b0316600090815260209190915260409020805460ff1916600117905556fe6f6e6c79207072696365206665656465722063616e20757064617465207468652070726963656576657279206469676974616c2063726564697473206d757374206861766520636f6c6c61746572616c70656767656456616c75652063616e206f6e6c792062652073657420617420746865203173742074696d65206469676974616c20637265646974206d696e7465646f6e6c79207472757374656420706172746e65722063616e206d696e7420746865206469676974616c2063726564697473526f6c65733a206163636f756e7420646f6573206e6f74206861766520726f6c65536166654d6174683a206d756c7469706c69636174696f6e206f766572666c6f77526f6c65733a206163636f756e7420697320746865207a65726f206164647265737357686974656c69737441646d696e526f6c653a2063616c6c657220646f6573206e6f742068617665207468652057686974656c69737441646d696e20726f6c6570726963652066656564657220666f7220676976656e2063757272656e6379206861736e2774206265656e20736574a165627a7a723058205b59849bb979eb5a58b283e12743176bb76df2a73855e551bb0cbe65504ddcc80029526f6c65733a206163636f756e7420697320746865207a65726f2061646472657373`

// DeployDigitalReserveSystem deploys a new Ethereum contract, binding an instance of DigitalReserveSystem to it.
func DeployDigitalReserveSystem(auth *bind.TransactOpts, backend bind.ContractBackend, veloTokenAddr common.Address) (common.Address, *types.Transaction, *DigitalReserveSystem, error) {
	parsed, err := abi.JSON(strings.NewReader(DigitalReserveSystemABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(DigitalReserveSystemBin), backend, veloTokenAddr)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &DigitalReserveSystem{DigitalReserveSystemCaller: DigitalReserveSystemCaller{contract: contract}, DigitalReserveSystemTransactor: DigitalReserveSystemTransactor{contract: contract}, DigitalReserveSystemFilterer: DigitalReserveSystemFilterer{contract: contract}}, nil
}

// DigitalReserveSystem is an auto generated Go binding around an Ethereum contract.
type DigitalReserveSystem struct {
	DigitalReserveSystemCaller     // Read-only binding to the contract
	DigitalReserveSystemTransactor // Write-only binding to the contract
	DigitalReserveSystemFilterer   // Log filterer for contract events
}

// DigitalReserveSystemCaller is an auto generated read-only Go binding around an Ethereum contract.
type DigitalReserveSystemCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DigitalReserveSystemTransactor is an auto generated write-only Go binding around an Ethereum contract.
type DigitalReserveSystemTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DigitalReserveSystemFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type DigitalReserveSystemFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DigitalReserveSystemSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type DigitalReserveSystemSession struct {
	Contract     *DigitalReserveSystem // Generic contract binding to set the session for
	CallOpts     bind.CallOpts         // Call options to use throughout this session
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// DigitalReserveSystemCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type DigitalReserveSystemCallerSession struct {
	Contract *DigitalReserveSystemCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts               // Call options to use throughout this session
}

// DigitalReserveSystemTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type DigitalReserveSystemTransactorSession struct {
	Contract     *DigitalReserveSystemTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts               // Transaction auth options to use throughout this session
}

// DigitalReserveSystemRaw is an auto generated low-level Go binding around an Ethereum contract.
type DigitalReserveSystemRaw struct {
	Contract *DigitalReserveSystem // Generic contract binding to access the raw methods on
}

// DigitalReserveSystemCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type DigitalReserveSystemCallerRaw struct {
	Contract *DigitalReserveSystemCaller // Generic read-only contract binding to access the raw methods on
}

// DigitalReserveSystemTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type DigitalReserveSystemTransactorRaw struct {
	Contract *DigitalReserveSystemTransactor // Generic write-only contract binding to access the raw methods on
}

// NewDigitalReserveSystem creates a new instance of DigitalReserveSystem, bound to a specific deployed contract.
func NewDigitalReserveSystem(address common.Address, backend bind.ContractBackend) (*DigitalReserveSystem, error) {
	contract, err := bindDigitalReserveSystem(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &DigitalReserveSystem{DigitalReserveSystemCaller: DigitalReserveSystemCaller{contract: contract}, DigitalReserveSystemTransactor: DigitalReserveSystemTransactor{contract: contract}, DigitalReserveSystemFilterer: DigitalReserveSystemFilterer{contract: contract}}, nil
}

// NewDigitalReserveSystemCaller creates a new read-only instance of DigitalReserveSystem, bound to a specific deployed contract.
func NewDigitalReserveSystemCaller(address common.Address, caller bind.ContractCaller) (*DigitalReserveSystemCaller, error) {
	contract, err := bindDigitalReserveSystem(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &DigitalReserveSystemCaller{contract: contract}, nil
}

// NewDigitalReserveSystemTransactor creates a new write-only instance of DigitalReserveSystem, bound to a specific deployed contract.
func NewDigitalReserveSystemTransactor(address common.Address, transactor bind.ContractTransactor) (*DigitalReserveSystemTransactor, error) {
	contract, err := bindDigitalReserveSystem(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &DigitalReserveSystemTransactor{contract: contract}, nil
}

// NewDigitalReserveSystemFilterer creates a new log filterer instance of DigitalReserveSystem, bound to a specific deployed contract.
func NewDigitalReserveSystemFilterer(address common.Address, filterer bind.ContractFilterer) (*DigitalReserveSystemFilterer, error) {
	contract, err := bindDigitalReserveSystem(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &DigitalReserveSystemFilterer{contract: contract}, nil
}

// bindDigitalReserveSystem binds a generic wrapper to an already deployed contract.
func bindDigitalReserveSystem(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(DigitalReserveSystemABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DigitalReserveSystem *DigitalReserveSystemRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _DigitalReserveSystem.Contract.DigitalReserveSystemCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DigitalReserveSystem *DigitalReserveSystemRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DigitalReserveSystem.Contract.DigitalReserveSystemTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DigitalReserveSystem *DigitalReserveSystemRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DigitalReserveSystem.Contract.DigitalReserveSystemTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DigitalReserveSystem *DigitalReserveSystemCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _DigitalReserveSystem.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DigitalReserveSystem *DigitalReserveSystemTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DigitalReserveSystem.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DigitalReserveSystem *DigitalReserveSystemTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DigitalReserveSystem.Contract.contract.Transact(opts, method, params...)
}

// CheckRequiredCollateralWithFee is a free data retrieval call binding the contract method 0x28893f5f.
//
// Solidity: function checkRequiredCollateralWithFee(uint256 peggedValue, bytes32 peggedCurrency, uint256 amount) constant returns(uint256, uint256)
func (_DigitalReserveSystem *DigitalReserveSystemCaller) CheckRequiredCollateralWithFee(opts *bind.CallOpts, peggedValue *big.Int, peggedCurrency [32]byte, amount *big.Int) (*big.Int, *big.Int, error) {
	var (
		ret0 = new(*big.Int)
		ret1 = new(*big.Int)
	)
	out := &[]interface{}{
		ret0,
		ret1,
	}
	err := _DigitalReserveSystem.contract.Call(opts, out, "checkRequiredCollateralWithFee", peggedValue, peggedCurrency, amount)
	return *ret0, *ret1, err
}

// CheckRequiredCollateralWithFee is a free data retrieval call binding the contract method 0x28893f5f.
//
// Solidity: function checkRequiredCollateralWithFee(uint256 peggedValue, bytes32 peggedCurrency, uint256 amount) constant returns(uint256, uint256)
func (_DigitalReserveSystem *DigitalReserveSystemSession) CheckRequiredCollateralWithFee(peggedValue *big.Int, peggedCurrency [32]byte, amount *big.Int) (*big.Int, *big.Int, error) {
	return _DigitalReserveSystem.Contract.CheckRequiredCollateralWithFee(&_DigitalReserveSystem.CallOpts, peggedValue, peggedCurrency, amount)
}

// CheckRequiredCollateralWithFee is a free data retrieval call binding the contract method 0x28893f5f.
//
// Solidity: function checkRequiredCollateralWithFee(uint256 peggedValue, bytes32 peggedCurrency, uint256 amount) constant returns(uint256, uint256)
func (_DigitalReserveSystem *DigitalReserveSystemCallerSession) CheckRequiredCollateralWithFee(peggedValue *big.Int, peggedCurrency [32]byte, amount *big.Int) (*big.Int, *big.Int, error) {
	return _DigitalReserveSystem.Contract.CheckRequiredCollateralWithFee(&_DigitalReserveSystem.CallOpts, peggedValue, peggedCurrency, amount)
}

// CollateralOf is a free data retrieval call binding the contract method 0x4e1bfbe7.
//
// Solidity: function collateralOf(bytes32 assetName, address issuer) constant returns(uint256)
func (_DigitalReserveSystem *DigitalReserveSystemCaller) CollateralOf(opts *bind.CallOpts, assetName [32]byte, issuer common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _DigitalReserveSystem.contract.Call(opts, out, "collateralOf", assetName, issuer)
	return *ret0, err
}

// CollateralOf is a free data retrieval call binding the contract method 0x4e1bfbe7.
//
// Solidity: function collateralOf(bytes32 assetName, address issuer) constant returns(uint256)
func (_DigitalReserveSystem *DigitalReserveSystemSession) CollateralOf(assetName [32]byte, issuer common.Address) (*big.Int, error) {
	return _DigitalReserveSystem.Contract.CollateralOf(&_DigitalReserveSystem.CallOpts, assetName, issuer)
}

// CollateralOf is a free data retrieval call binding the contract method 0x4e1bfbe7.
//
// Solidity: function collateralOf(bytes32 assetName, address issuer) constant returns(uint256)
func (_DigitalReserveSystem *DigitalReserveSystemCallerSession) CollateralOf(assetName [32]byte, issuer common.Address) (*big.Int, error) {
	return _DigitalReserveSystem.Contract.CollateralOf(&_DigitalReserveSystem.CallOpts, assetName, issuer)
}

// CollectedFee is a free data retrieval call binding the contract method 0xe811f50a.
//
// Solidity: function collectedFee() constant returns(uint256)
func (_DigitalReserveSystem *DigitalReserveSystemCaller) CollectedFee(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _DigitalReserveSystem.contract.Call(opts, out, "collectedFee")
	return *ret0, err
}

// CollectedFee is a free data retrieval call binding the contract method 0xe811f50a.
//
// Solidity: function collectedFee() constant returns(uint256)
func (_DigitalReserveSystem *DigitalReserveSystemSession) CollectedFee() (*big.Int, error) {
	return _DigitalReserveSystem.Contract.CollectedFee(&_DigitalReserveSystem.CallOpts)
}

// CollectedFee is a free data retrieval call binding the contract method 0xe811f50a.
//
// Solidity: function collectedFee() constant returns(uint256)
func (_DigitalReserveSystem *DigitalReserveSystemCallerSession) CollectedFee() (*big.Int, error) {
	return _DigitalReserveSystem.Contract.CollectedFee(&_DigitalReserveSystem.CallOpts)
}

// CreditIssuanceFee is a free data retrieval call binding the contract method 0x1f6f2072.
//
// Solidity: function creditIssuanceFee() constant returns(uint256)
func (_DigitalReserveSystem *DigitalReserveSystemCaller) CreditIssuanceFee(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _DigitalReserveSystem.contract.Call(opts, out, "creditIssuanceFee")
	return *ret0, err
}

// CreditIssuanceFee is a free data retrieval call binding the contract method 0x1f6f2072.
//
// Solidity: function creditIssuanceFee() constant returns(uint256)
func (_DigitalReserveSystem *DigitalReserveSystemSession) CreditIssuanceFee() (*big.Int, error) {
	return _DigitalReserveSystem.Contract.CreditIssuanceFee(&_DigitalReserveSystem.CallOpts)
}

// CreditIssuanceFee is a free data retrieval call binding the contract method 0x1f6f2072.
//
// Solidity: function creditIssuanceFee() constant returns(uint256)
func (_DigitalReserveSystem *DigitalReserveSystemCallerSession) CreditIssuanceFee() (*big.Int, error) {
	return _DigitalReserveSystem.Contract.CreditIssuanceFee(&_DigitalReserveSystem.CallOpts)
}

// DigitalCredits is a free data retrieval call binding the contract method 0xc574ccb4.
//
// Solidity: function digitalCredits(bytes32 ) constant returns(uint256 collateral, uint256 assetAmount, uint256 peggedValue, bytes32 peggedCurrency, bytes32 stellarIssuerAddress)
func (_DigitalReserveSystem *DigitalReserveSystemCaller) DigitalCredits(opts *bind.CallOpts, arg0 [32]byte) (struct {
	Collateral           *big.Int
	AssetAmount          *big.Int
	PeggedValue          *big.Int
	PeggedCurrency       [32]byte
	StellarIssuerAddress [32]byte
}, error) {
	ret := new(struct {
		Collateral           *big.Int
		AssetAmount          *big.Int
		PeggedValue          *big.Int
		PeggedCurrency       [32]byte
		StellarIssuerAddress [32]byte
	})
	out := ret
	err := _DigitalReserveSystem.contract.Call(opts, out, "digitalCredits", arg0)
	return *ret, err
}

// DigitalCredits is a free data retrieval call binding the contract method 0xc574ccb4.
//
// Solidity: function digitalCredits(bytes32 ) constant returns(uint256 collateral, uint256 assetAmount, uint256 peggedValue, bytes32 peggedCurrency, bytes32 stellarIssuerAddress)
func (_DigitalReserveSystem *DigitalReserveSystemSession) DigitalCredits(arg0 [32]byte) (struct {
	Collateral           *big.Int
	AssetAmount          *big.Int
	PeggedValue          *big.Int
	PeggedCurrency       [32]byte
	StellarIssuerAddress [32]byte
}, error) {
	return _DigitalReserveSystem.Contract.DigitalCredits(&_DigitalReserveSystem.CallOpts, arg0)
}

// DigitalCredits is a free data retrieval call binding the contract method 0xc574ccb4.
//
// Solidity: function digitalCredits(bytes32 ) constant returns(uint256 collateral, uint256 assetAmount, uint256 peggedValue, bytes32 peggedCurrency, bytes32 stellarIssuerAddress)
func (_DigitalReserveSystem *DigitalReserveSystemCallerSession) DigitalCredits(arg0 [32]byte) (struct {
	Collateral           *big.Int
	AssetAmount          *big.Int
	PeggedValue          *big.Int
	PeggedCurrency       [32]byte
	StellarIssuerAddress [32]byte
}, error) {
	return _DigitalReserveSystem.Contract.DigitalCredits(&_DigitalReserveSystem.CallOpts, arg0)
}

// GetCollectedFee is a free data retrieval call binding the contract method 0x271ec305.
//
// Solidity: function getCollectedFee() constant returns(uint256)
func (_DigitalReserveSystem *DigitalReserveSystemCaller) GetCollectedFee(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _DigitalReserveSystem.contract.Call(opts, out, "getCollectedFee")
	return *ret0, err
}

// GetCollectedFee is a free data retrieval call binding the contract method 0x271ec305.
//
// Solidity: function getCollectedFee() constant returns(uint256)
func (_DigitalReserveSystem *DigitalReserveSystemSession) GetCollectedFee() (*big.Int, error) {
	return _DigitalReserveSystem.Contract.GetCollectedFee(&_DigitalReserveSystem.CallOpts)
}

// GetCollectedFee is a free data retrieval call binding the contract method 0x271ec305.
//
// Solidity: function getCollectedFee() constant returns(uint256)
func (_DigitalReserveSystem *DigitalReserveSystemCallerSession) GetCollectedFee() (*big.Int, error) {
	return _DigitalReserveSystem.Contract.GetCollectedFee(&_DigitalReserveSystem.CallOpts)
}

// GetPrice is a free data retrieval call binding the contract method 0x31d98b3f.
//
// Solidity: function getPrice(bytes32 currency) constant returns(uint256)
func (_DigitalReserveSystem *DigitalReserveSystemCaller) GetPrice(opts *bind.CallOpts, currency [32]byte) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _DigitalReserveSystem.contract.Call(opts, out, "getPrice", currency)
	return *ret0, err
}

// GetPrice is a free data retrieval call binding the contract method 0x31d98b3f.
//
// Solidity: function getPrice(bytes32 currency) constant returns(uint256)
func (_DigitalReserveSystem *DigitalReserveSystemSession) GetPrice(currency [32]byte) (*big.Int, error) {
	return _DigitalReserveSystem.Contract.GetPrice(&_DigitalReserveSystem.CallOpts, currency)
}

// GetPrice is a free data retrieval call binding the contract method 0x31d98b3f.
//
// Solidity: function getPrice(bytes32 currency) constant returns(uint256)
func (_DigitalReserveSystem *DigitalReserveSystemCallerSession) GetPrice(currency [32]byte) (*big.Int, error) {
	return _DigitalReserveSystem.Contract.GetPrice(&_DigitalReserveSystem.CallOpts, currency)
}

// IsWhitelistAdmin is a free data retrieval call binding the contract method 0xbb5f747b.
//
// Solidity: function isWhitelistAdmin(address account) constant returns(bool)
func (_DigitalReserveSystem *DigitalReserveSystemCaller) IsWhitelistAdmin(opts *bind.CallOpts, account common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _DigitalReserveSystem.contract.Call(opts, out, "isWhitelistAdmin", account)
	return *ret0, err
}

// IsWhitelistAdmin is a free data retrieval call binding the contract method 0xbb5f747b.
//
// Solidity: function isWhitelistAdmin(address account) constant returns(bool)
func (_DigitalReserveSystem *DigitalReserveSystemSession) IsWhitelistAdmin(account common.Address) (bool, error) {
	return _DigitalReserveSystem.Contract.IsWhitelistAdmin(&_DigitalReserveSystem.CallOpts, account)
}

// IsWhitelistAdmin is a free data retrieval call binding the contract method 0xbb5f747b.
//
// Solidity: function isWhitelistAdmin(address account) constant returns(bool)
func (_DigitalReserveSystem *DigitalReserveSystemCallerSession) IsWhitelistAdmin(account common.Address) (bool, error) {
	return _DigitalReserveSystem.Contract.IsWhitelistAdmin(&_DigitalReserveSystem.CallOpts, account)
}

// PriceFeeders is a free data retrieval call binding the contract method 0x64b456ed.
//
// Solidity: function priceFeeders(bytes32 ) constant returns(address)
func (_DigitalReserveSystem *DigitalReserveSystemCaller) PriceFeeders(opts *bind.CallOpts, arg0 [32]byte) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _DigitalReserveSystem.contract.Call(opts, out, "priceFeeders", arg0)
	return *ret0, err
}

// PriceFeeders is a free data retrieval call binding the contract method 0x64b456ed.
//
// Solidity: function priceFeeders(bytes32 ) constant returns(address)
func (_DigitalReserveSystem *DigitalReserveSystemSession) PriceFeeders(arg0 [32]byte) (common.Address, error) {
	return _DigitalReserveSystem.Contract.PriceFeeders(&_DigitalReserveSystem.CallOpts, arg0)
}

// PriceFeeders is a free data retrieval call binding the contract method 0x64b456ed.
//
// Solidity: function priceFeeders(bytes32 ) constant returns(address)
func (_DigitalReserveSystem *DigitalReserveSystemCallerSession) PriceFeeders(arg0 [32]byte) (common.Address, error) {
	return _DigitalReserveSystem.Contract.PriceFeeders(&_DigitalReserveSystem.CallOpts, arg0)
}

// Prices is a free data retrieval call binding the contract method 0x60846bc6.
//
// Solidity: function prices(bytes32 ) constant returns(uint256)
func (_DigitalReserveSystem *DigitalReserveSystemCaller) Prices(opts *bind.CallOpts, arg0 [32]byte) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _DigitalReserveSystem.contract.Call(opts, out, "prices", arg0)
	return *ret0, err
}

// Prices is a free data retrieval call binding the contract method 0x60846bc6.
//
// Solidity: function prices(bytes32 ) constant returns(uint256)
func (_DigitalReserveSystem *DigitalReserveSystemSession) Prices(arg0 [32]byte) (*big.Int, error) {
	return _DigitalReserveSystem.Contract.Prices(&_DigitalReserveSystem.CallOpts, arg0)
}

// Prices is a free data retrieval call binding the contract method 0x60846bc6.
//
// Solidity: function prices(bytes32 ) constant returns(uint256)
func (_DigitalReserveSystem *DigitalReserveSystemCallerSession) Prices(arg0 [32]byte) (*big.Int, error) {
	return _DigitalReserveSystem.Contract.Prices(&_DigitalReserveSystem.CallOpts, arg0)
}

// TrustedPartners is a free data retrieval call binding the contract method 0x29fc9ec8.
//
// Solidity: function trustedPartners(address ) constant returns(bool)
func (_DigitalReserveSystem *DigitalReserveSystemCaller) TrustedPartners(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _DigitalReserveSystem.contract.Call(opts, out, "trustedPartners", arg0)
	return *ret0, err
}

// TrustedPartners is a free data retrieval call binding the contract method 0x29fc9ec8.
//
// Solidity: function trustedPartners(address ) constant returns(bool)
func (_DigitalReserveSystem *DigitalReserveSystemSession) TrustedPartners(arg0 common.Address) (bool, error) {
	return _DigitalReserveSystem.Contract.TrustedPartners(&_DigitalReserveSystem.CallOpts, arg0)
}

// TrustedPartners is a free data retrieval call binding the contract method 0x29fc9ec8.
//
// Solidity: function trustedPartners(address ) constant returns(bool)
func (_DigitalReserveSystem *DigitalReserveSystemCallerSession) TrustedPartners(arg0 common.Address) (bool, error) {
	return _DigitalReserveSystem.Contract.TrustedPartners(&_DigitalReserveSystem.CallOpts, arg0)
}

// CallbackConfirmIssueCredit is a paid mutator transaction binding the contract method 0x747f43b6.
//
// Solidity: function __callbackConfirmIssueCredit(address issuer, bytes32 assetName, string stellarIssuerAddress) returns()
func (_DigitalReserveSystem *DigitalReserveSystemTransactor) CallbackConfirmIssueCredit(opts *bind.TransactOpts, issuer common.Address, assetName [32]byte, stellarIssuerAddress string) (*types.Transaction, error) {
	return _DigitalReserveSystem.contract.Transact(opts, "__callbackConfirmIssueCredit", issuer, assetName, stellarIssuerAddress)
}

// CallbackConfirmIssueCredit is a paid mutator transaction binding the contract method 0x747f43b6.
//
// Solidity: function __callbackConfirmIssueCredit(address issuer, bytes32 assetName, string stellarIssuerAddress) returns()
func (_DigitalReserveSystem *DigitalReserveSystemSession) CallbackConfirmIssueCredit(issuer common.Address, assetName [32]byte, stellarIssuerAddress string) (*types.Transaction, error) {
	return _DigitalReserveSystem.Contract.CallbackConfirmIssueCredit(&_DigitalReserveSystem.TransactOpts, issuer, assetName, stellarIssuerAddress)
}

// CallbackConfirmIssueCredit is a paid mutator transaction binding the contract method 0x747f43b6.
//
// Solidity: function __callbackConfirmIssueCredit(address issuer, bytes32 assetName, string stellarIssuerAddress) returns()
func (_DigitalReserveSystem *DigitalReserveSystemTransactorSession) CallbackConfirmIssueCredit(issuer common.Address, assetName [32]byte, stellarIssuerAddress string) (*types.Transaction, error) {
	return _DigitalReserveSystem.Contract.CallbackConfirmIssueCredit(&_DigitalReserveSystem.TransactOpts, issuer, assetName, stellarIssuerAddress)
}

// AddWhitelistAdmin is a paid mutator transaction binding the contract method 0x7362d9c8.
//
// Solidity: function addWhitelistAdmin(address account) returns()
func (_DigitalReserveSystem *DigitalReserveSystemTransactor) AddWhitelistAdmin(opts *bind.TransactOpts, account common.Address) (*types.Transaction, error) {
	return _DigitalReserveSystem.contract.Transact(opts, "addWhitelistAdmin", account)
}

// AddWhitelistAdmin is a paid mutator transaction binding the contract method 0x7362d9c8.
//
// Solidity: function addWhitelistAdmin(address account) returns()
func (_DigitalReserveSystem *DigitalReserveSystemSession) AddWhitelistAdmin(account common.Address) (*types.Transaction, error) {
	return _DigitalReserveSystem.Contract.AddWhitelistAdmin(&_DigitalReserveSystem.TransactOpts, account)
}

// AddWhitelistAdmin is a paid mutator transaction binding the contract method 0x7362d9c8.
//
// Solidity: function addWhitelistAdmin(address account) returns()
func (_DigitalReserveSystem *DigitalReserveSystemTransactorSession) AddWhitelistAdmin(account common.Address) (*types.Transaction, error) {
	return _DigitalReserveSystem.Contract.AddWhitelistAdmin(&_DigitalReserveSystem.TransactOpts, account)
}

// Mint is a paid mutator transaction binding the contract method 0x5762517e.
//
// Solidity: function mint(uint256 peggedValue, bytes32 peggedCurrency, bytes32 assetName, uint256 amount, string stellarAddr, bytes32 hashSecret) returns()
func (_DigitalReserveSystem *DigitalReserveSystemTransactor) Mint(opts *bind.TransactOpts, peggedValue *big.Int, peggedCurrency [32]byte, assetName [32]byte, amount *big.Int, stellarAddr string, hashSecret [32]byte) (*types.Transaction, error) {
	return _DigitalReserveSystem.contract.Transact(opts, "mint", peggedValue, peggedCurrency, assetName, amount, stellarAddr, hashSecret)
}

// Mint is a paid mutator transaction binding the contract method 0x5762517e.
//
// Solidity: function mint(uint256 peggedValue, bytes32 peggedCurrency, bytes32 assetName, uint256 amount, string stellarAddr, bytes32 hashSecret) returns()
func (_DigitalReserveSystem *DigitalReserveSystemSession) Mint(peggedValue *big.Int, peggedCurrency [32]byte, assetName [32]byte, amount *big.Int, stellarAddr string, hashSecret [32]byte) (*types.Transaction, error) {
	return _DigitalReserveSystem.Contract.Mint(&_DigitalReserveSystem.TransactOpts, peggedValue, peggedCurrency, assetName, amount, stellarAddr, hashSecret)
}

// Mint is a paid mutator transaction binding the contract method 0x5762517e.
//
// Solidity: function mint(uint256 peggedValue, bytes32 peggedCurrency, bytes32 assetName, uint256 amount, string stellarAddr, bytes32 hashSecret) returns()
func (_DigitalReserveSystem *DigitalReserveSystemTransactorSession) Mint(peggedValue *big.Int, peggedCurrency [32]byte, assetName [32]byte, amount *big.Int, stellarAddr string, hashSecret [32]byte) (*types.Transaction, error) {
	return _DigitalReserveSystem.Contract.Mint(&_DigitalReserveSystem.TransactOpts, peggedValue, peggedCurrency, assetName, amount, stellarAddr, hashSecret)
}

// Redeem is a paid mutator transaction binding the contract method 0x673da154.
//
// Solidity: function redeem(bytes32 assetName, uint256 amount) returns()
func (_DigitalReserveSystem *DigitalReserveSystemTransactor) Redeem(opts *bind.TransactOpts, assetName [32]byte, amount *big.Int) (*types.Transaction, error) {
	return _DigitalReserveSystem.contract.Transact(opts, "redeem", assetName, amount)
}

// Redeem is a paid mutator transaction binding the contract method 0x673da154.
//
// Solidity: function redeem(bytes32 assetName, uint256 amount) returns()
func (_DigitalReserveSystem *DigitalReserveSystemSession) Redeem(assetName [32]byte, amount *big.Int) (*types.Transaction, error) {
	return _DigitalReserveSystem.Contract.Redeem(&_DigitalReserveSystem.TransactOpts, assetName, amount)
}

// Redeem is a paid mutator transaction binding the contract method 0x673da154.
//
// Solidity: function redeem(bytes32 assetName, uint256 amount) returns()
func (_DigitalReserveSystem *DigitalReserveSystemTransactorSession) Redeem(assetName [32]byte, amount *big.Int) (*types.Transaction, error) {
	return _DigitalReserveSystem.Contract.Redeem(&_DigitalReserveSystem.TransactOpts, assetName, amount)
}

// RenounceWhitelistAdmin is a paid mutator transaction binding the contract method 0x4c5a628c.
//
// Solidity: function renounceWhitelistAdmin() returns()
func (_DigitalReserveSystem *DigitalReserveSystemTransactor) RenounceWhitelistAdmin(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DigitalReserveSystem.contract.Transact(opts, "renounceWhitelistAdmin")
}

// RenounceWhitelistAdmin is a paid mutator transaction binding the contract method 0x4c5a628c.
//
// Solidity: function renounceWhitelistAdmin() returns()
func (_DigitalReserveSystem *DigitalReserveSystemSession) RenounceWhitelistAdmin() (*types.Transaction, error) {
	return _DigitalReserveSystem.Contract.RenounceWhitelistAdmin(&_DigitalReserveSystem.TransactOpts)
}

// RenounceWhitelistAdmin is a paid mutator transaction binding the contract method 0x4c5a628c.
//
// Solidity: function renounceWhitelistAdmin() returns()
func (_DigitalReserveSystem *DigitalReserveSystemTransactorSession) RenounceWhitelistAdmin() (*types.Transaction, error) {
	return _DigitalReserveSystem.Contract.RenounceWhitelistAdmin(&_DigitalReserveSystem.TransactOpts)
}

// SetCreditIssuanceFee is a paid mutator transaction binding the contract method 0xaedb08d4.
//
// Solidity: function setCreditIssuanceFee(uint256 newFee) returns()
func (_DigitalReserveSystem *DigitalReserveSystemTransactor) SetCreditIssuanceFee(opts *bind.TransactOpts, newFee *big.Int) (*types.Transaction, error) {
	return _DigitalReserveSystem.contract.Transact(opts, "setCreditIssuanceFee", newFee)
}

// SetCreditIssuanceFee is a paid mutator transaction binding the contract method 0xaedb08d4.
//
// Solidity: function setCreditIssuanceFee(uint256 newFee) returns()
func (_DigitalReserveSystem *DigitalReserveSystemSession) SetCreditIssuanceFee(newFee *big.Int) (*types.Transaction, error) {
	return _DigitalReserveSystem.Contract.SetCreditIssuanceFee(&_DigitalReserveSystem.TransactOpts, newFee)
}

// SetCreditIssuanceFee is a paid mutator transaction binding the contract method 0xaedb08d4.
//
// Solidity: function setCreditIssuanceFee(uint256 newFee) returns()
func (_DigitalReserveSystem *DigitalReserveSystemTransactorSession) SetCreditIssuanceFee(newFee *big.Int) (*types.Transaction, error) {
	return _DigitalReserveSystem.Contract.SetCreditIssuanceFee(&_DigitalReserveSystem.TransactOpts, newFee)
}

// SetPrice is a paid mutator transaction binding the contract method 0x10d8d74d.
//
// Solidity: function setPrice(bytes32 currency, uint256 newPrice) returns()
func (_DigitalReserveSystem *DigitalReserveSystemTransactor) SetPrice(opts *bind.TransactOpts, currency [32]byte, newPrice *big.Int) (*types.Transaction, error) {
	return _DigitalReserveSystem.contract.Transact(opts, "setPrice", currency, newPrice)
}

// SetPrice is a paid mutator transaction binding the contract method 0x10d8d74d.
//
// Solidity: function setPrice(bytes32 currency, uint256 newPrice) returns()
func (_DigitalReserveSystem *DigitalReserveSystemSession) SetPrice(currency [32]byte, newPrice *big.Int) (*types.Transaction, error) {
	return _DigitalReserveSystem.Contract.SetPrice(&_DigitalReserveSystem.TransactOpts, currency, newPrice)
}

// SetPrice is a paid mutator transaction binding the contract method 0x10d8d74d.
//
// Solidity: function setPrice(bytes32 currency, uint256 newPrice) returns()
func (_DigitalReserveSystem *DigitalReserveSystemTransactorSession) SetPrice(currency [32]byte, newPrice *big.Int) (*types.Transaction, error) {
	return _DigitalReserveSystem.Contract.SetPrice(&_DigitalReserveSystem.TransactOpts, currency, newPrice)
}

// SetPriceFeeder is a paid mutator transaction binding the contract method 0xa7aa2b79.
//
// Solidity: function setPriceFeeder(bytes32 currency, address priceFeederAddr) returns()
func (_DigitalReserveSystem *DigitalReserveSystemTransactor) SetPriceFeeder(opts *bind.TransactOpts, currency [32]byte, priceFeederAddr common.Address) (*types.Transaction, error) {
	return _DigitalReserveSystem.contract.Transact(opts, "setPriceFeeder", currency, priceFeederAddr)
}

// SetPriceFeeder is a paid mutator transaction binding the contract method 0xa7aa2b79.
//
// Solidity: function setPriceFeeder(bytes32 currency, address priceFeederAddr) returns()
func (_DigitalReserveSystem *DigitalReserveSystemSession) SetPriceFeeder(currency [32]byte, priceFeederAddr common.Address) (*types.Transaction, error) {
	return _DigitalReserveSystem.Contract.SetPriceFeeder(&_DigitalReserveSystem.TransactOpts, currency, priceFeederAddr)
}

// SetPriceFeeder is a paid mutator transaction binding the contract method 0xa7aa2b79.
//
// Solidity: function setPriceFeeder(bytes32 currency, address priceFeederAddr) returns()
func (_DigitalReserveSystem *DigitalReserveSystemTransactorSession) SetPriceFeeder(currency [32]byte, priceFeederAddr common.Address) (*types.Transaction, error) {
	return _DigitalReserveSystem.Contract.SetPriceFeeder(&_DigitalReserveSystem.TransactOpts, currency, priceFeederAddr)
}

// SetTrustedPartner is a paid mutator transaction binding the contract method 0x68698406.
//
// Solidity: function setTrustedPartner(address addr) returns()
func (_DigitalReserveSystem *DigitalReserveSystemTransactor) SetTrustedPartner(opts *bind.TransactOpts, addr common.Address) (*types.Transaction, error) {
	return _DigitalReserveSystem.contract.Transact(opts, "setTrustedPartner", addr)
}

// SetTrustedPartner is a paid mutator transaction binding the contract method 0x68698406.
//
// Solidity: function setTrustedPartner(address addr) returns()
func (_DigitalReserveSystem *DigitalReserveSystemSession) SetTrustedPartner(addr common.Address) (*types.Transaction, error) {
	return _DigitalReserveSystem.Contract.SetTrustedPartner(&_DigitalReserveSystem.TransactOpts, addr)
}

// SetTrustedPartner is a paid mutator transaction binding the contract method 0x68698406.
//
// Solidity: function setTrustedPartner(address addr) returns()
func (_DigitalReserveSystem *DigitalReserveSystemTransactorSession) SetTrustedPartner(addr common.Address) (*types.Transaction, error) {
	return _DigitalReserveSystem.Contract.SetTrustedPartner(&_DigitalReserveSystem.TransactOpts, addr)
}

// WithdrawFee is a paid mutator transaction binding the contract method 0xbe357616.
//
// Solidity: function withdrawFee(uint256 amount) returns()
func (_DigitalReserveSystem *DigitalReserveSystemTransactor) WithdrawFee(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _DigitalReserveSystem.contract.Transact(opts, "withdrawFee", amount)
}

// WithdrawFee is a paid mutator transaction binding the contract method 0xbe357616.
//
// Solidity: function withdrawFee(uint256 amount) returns()
func (_DigitalReserveSystem *DigitalReserveSystemSession) WithdrawFee(amount *big.Int) (*types.Transaction, error) {
	return _DigitalReserveSystem.Contract.WithdrawFee(&_DigitalReserveSystem.TransactOpts, amount)
}

// WithdrawFee is a paid mutator transaction binding the contract method 0xbe357616.
//
// Solidity: function withdrawFee(uint256 amount) returns()
func (_DigitalReserveSystem *DigitalReserveSystemTransactorSession) WithdrawFee(amount *big.Int) (*types.Transaction, error) {
	return _DigitalReserveSystem.Contract.WithdrawFee(&_DigitalReserveSystem.TransactOpts, amount)
}

// DigitalReserveSystemMintIterator is returned from FilterMint and is used to iterate over the raw logs and unpacked data for Mint events raised by the DigitalReserveSystem contract.
type DigitalReserveSystemMintIterator struct {
	Event *DigitalReserveSystemMint // Event containing the contract specifics and raw log

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
func (it *DigitalReserveSystemMintIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DigitalReserveSystemMint)
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
		it.Event = new(DigitalReserveSystemMint)
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
func (it *DigitalReserveSystemMintIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DigitalReserveSystemMintIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DigitalReserveSystemMint represents a Mint event raised by the DigitalReserveSystem contract.
type DigitalReserveSystemMint struct {
	AssetName   [32]byte
	Amount      *big.Int
	Collateral  *big.Int
	Issuer      common.Address
	StellarAddr common.Hash
	HashSecret  [32]byte
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterMint is a free log retrieval operation binding the contract event 0xaa2a61bb785a72b9e47fb58125440d23bbad4b6c8dda3e17ab80714bf056e4e3.
//
// Solidity: event Mint(bytes32 assetName, uint256 amount, uint256 collateral, address indexed issuer, string indexed stellarAddr, bytes32 hashSecret)
func (_DigitalReserveSystem *DigitalReserveSystemFilterer) FilterMint(opts *bind.FilterOpts, issuer []common.Address, stellarAddr []string) (*DigitalReserveSystemMintIterator, error) {

	var issuerRule []interface{}
	for _, issuerItem := range issuer {
		issuerRule = append(issuerRule, issuerItem)
	}
	var stellarAddrRule []interface{}
	for _, stellarAddrItem := range stellarAddr {
		stellarAddrRule = append(stellarAddrRule, stellarAddrItem)
	}

	logs, sub, err := _DigitalReserveSystem.contract.FilterLogs(opts, "Mint", issuerRule, stellarAddrRule)
	if err != nil {
		return nil, err
	}
	return &DigitalReserveSystemMintIterator{contract: _DigitalReserveSystem.contract, event: "Mint", logs: logs, sub: sub}, nil
}

// WatchMint is a free log subscription operation binding the contract event 0xaa2a61bb785a72b9e47fb58125440d23bbad4b6c8dda3e17ab80714bf056e4e3.
//
// Solidity: event Mint(bytes32 assetName, uint256 amount, uint256 collateral, address indexed issuer, string indexed stellarAddr, bytes32 hashSecret)
func (_DigitalReserveSystem *DigitalReserveSystemFilterer) WatchMint(opts *bind.WatchOpts, sink chan<- *DigitalReserveSystemMint, issuer []common.Address, stellarAddr []string) (event.Subscription, error) {

	var issuerRule []interface{}
	for _, issuerItem := range issuer {
		issuerRule = append(issuerRule, issuerItem)
	}
	var stellarAddrRule []interface{}
	for _, stellarAddrItem := range stellarAddr {
		stellarAddrRule = append(stellarAddrRule, stellarAddrItem)
	}

	logs, sub, err := _DigitalReserveSystem.contract.WatchLogs(opts, "Mint", issuerRule, stellarAddrRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DigitalReserveSystemMint)
				if err := _DigitalReserveSystem.contract.UnpackLog(event, "Mint", log); err != nil {
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

// DigitalReserveSystemWhitelistAdminAddedIterator is returned from FilterWhitelistAdminAdded and is used to iterate over the raw logs and unpacked data for WhitelistAdminAdded events raised by the DigitalReserveSystem contract.
type DigitalReserveSystemWhitelistAdminAddedIterator struct {
	Event *DigitalReserveSystemWhitelistAdminAdded // Event containing the contract specifics and raw log

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
func (it *DigitalReserveSystemWhitelistAdminAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DigitalReserveSystemWhitelistAdminAdded)
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
		it.Event = new(DigitalReserveSystemWhitelistAdminAdded)
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
func (it *DigitalReserveSystemWhitelistAdminAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DigitalReserveSystemWhitelistAdminAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DigitalReserveSystemWhitelistAdminAdded represents a WhitelistAdminAdded event raised by the DigitalReserveSystem contract.
type DigitalReserveSystemWhitelistAdminAdded struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterWhitelistAdminAdded is a free log retrieval operation binding the contract event 0x22380c05984257a1cb900161c713dd71d39e74820f1aea43bd3f1bdd20961299.
//
// Solidity: event WhitelistAdminAdded(address indexed account)
func (_DigitalReserveSystem *DigitalReserveSystemFilterer) FilterWhitelistAdminAdded(opts *bind.FilterOpts, account []common.Address) (*DigitalReserveSystemWhitelistAdminAddedIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _DigitalReserveSystem.contract.FilterLogs(opts, "WhitelistAdminAdded", accountRule)
	if err != nil {
		return nil, err
	}
	return &DigitalReserveSystemWhitelistAdminAddedIterator{contract: _DigitalReserveSystem.contract, event: "WhitelistAdminAdded", logs: logs, sub: sub}, nil
}

// WatchWhitelistAdminAdded is a free log subscription operation binding the contract event 0x22380c05984257a1cb900161c713dd71d39e74820f1aea43bd3f1bdd20961299.
//
// Solidity: event WhitelistAdminAdded(address indexed account)
func (_DigitalReserveSystem *DigitalReserveSystemFilterer) WatchWhitelistAdminAdded(opts *bind.WatchOpts, sink chan<- *DigitalReserveSystemWhitelistAdminAdded, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _DigitalReserveSystem.contract.WatchLogs(opts, "WhitelistAdminAdded", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DigitalReserveSystemWhitelistAdminAdded)
				if err := _DigitalReserveSystem.contract.UnpackLog(event, "WhitelistAdminAdded", log); err != nil {
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

// DigitalReserveSystemWhitelistAdminRemovedIterator is returned from FilterWhitelistAdminRemoved and is used to iterate over the raw logs and unpacked data for WhitelistAdminRemoved events raised by the DigitalReserveSystem contract.
type DigitalReserveSystemWhitelistAdminRemovedIterator struct {
	Event *DigitalReserveSystemWhitelistAdminRemoved // Event containing the contract specifics and raw log

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
func (it *DigitalReserveSystemWhitelistAdminRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DigitalReserveSystemWhitelistAdminRemoved)
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
		it.Event = new(DigitalReserveSystemWhitelistAdminRemoved)
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
func (it *DigitalReserveSystemWhitelistAdminRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DigitalReserveSystemWhitelistAdminRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DigitalReserveSystemWhitelistAdminRemoved represents a WhitelistAdminRemoved event raised by the DigitalReserveSystem contract.
type DigitalReserveSystemWhitelistAdminRemoved struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterWhitelistAdminRemoved is a free log retrieval operation binding the contract event 0x0a8eb35e5ca14b3d6f28e4abf2f128dbab231a58b56e89beb5d636115001e165.
//
// Solidity: event WhitelistAdminRemoved(address indexed account)
func (_DigitalReserveSystem *DigitalReserveSystemFilterer) FilterWhitelistAdminRemoved(opts *bind.FilterOpts, account []common.Address) (*DigitalReserveSystemWhitelistAdminRemovedIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _DigitalReserveSystem.contract.FilterLogs(opts, "WhitelistAdminRemoved", accountRule)
	if err != nil {
		return nil, err
	}
	return &DigitalReserveSystemWhitelistAdminRemovedIterator{contract: _DigitalReserveSystem.contract, event: "WhitelistAdminRemoved", logs: logs, sub: sub}, nil
}

// WatchWhitelistAdminRemoved is a free log subscription operation binding the contract event 0x0a8eb35e5ca14b3d6f28e4abf2f128dbab231a58b56e89beb5d636115001e165.
//
// Solidity: event WhitelistAdminRemoved(address indexed account)
func (_DigitalReserveSystem *DigitalReserveSystemFilterer) WatchWhitelistAdminRemoved(opts *bind.WatchOpts, sink chan<- *DigitalReserveSystemWhitelistAdminRemoved, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _DigitalReserveSystem.contract.WatchLogs(opts, "WhitelistAdminRemoved", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DigitalReserveSystemWhitelistAdminRemoved)
				if err := _DigitalReserveSystem.contract.UnpackLog(event, "WhitelistAdminRemoved", log); err != nil {
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

// ERC20ABI is the input ABI used to generate the binding from.
const ERC20ABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"spender\",\"type\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"sender\",\"type\":\"address\"},{\"name\":\"recipient\",\"type\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"spender\",\"type\":\"address\"},{\"name\":\"addedValue\",\"type\":\"uint256\"}],\"name\":\"increaseAllowance\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"spender\",\"type\":\"address\"},{\"name\":\"subtractedValue\",\"type\":\"uint256\"}],\"name\":\"decreaseAllowance\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"recipient\",\"type\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"owner\",\"type\":\"address\"},{\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"}]"

// ERC20Bin is the compiled bytecode used for deploying new contracts.
const ERC20Bin = `0x608060405234801561001057600080fd5b506106f3806100206000396000f3fe608060405234801561001057600080fd5b50600436106100885760003560e01c806370a082311161005b57806370a0823114610149578063a457c2d71461016f578063a9059cbb1461019b578063dd62ed3e146101c757610088565b8063095ea7b31461008d57806318160ddd146100cd57806323b872dd146100e7578063395093511461011d575b600080fd5b6100b9600480360360408110156100a357600080fd5b506001600160a01b0381351690602001356101f5565b604080519115158252519081900360200190f35b6100d561020b565b60408051918252519081900360200190f35b6100b9600480360360608110156100fd57600080fd5b506001600160a01b03813581169160208101359091169060400135610211565b6100b96004803603604081101561013357600080fd5b506001600160a01b038135169060200135610268565b6100d56004803603602081101561015f57600080fd5b50356001600160a01b03166102a4565b6100b96004803603604081101561018557600080fd5b506001600160a01b0381351690602001356102bf565b6100b9600480360360408110156101b157600080fd5b506001600160a01b0381351690602001356102fb565b6100d5600480360360408110156101dd57600080fd5b506001600160a01b0381358116916020013516610308565b6000610202338484610333565b50600192915050565b60025490565b600061021e848484610429565b6001600160a01b03841660009081526001602090815260408083203380855292529091205461025e918691610259908663ffffffff61057516565b610333565b5060019392505050565b3360008181526001602090815260408083206001600160a01b03871684529091528120549091610202918590610259908663ffffffff6105d516565b6001600160a01b031660009081526020819052604090205490565b3360008181526001602090815260408083206001600160a01b03871684529091528120549091610202918590610259908663ffffffff61057516565b6000610202338484610429565b6001600160a01b03918216600090815260016020908152604080832093909416825291909152205490565b6001600160a01b038316151561037d57604051600160e51b62461bcd0281526004018080602001828103825260248152602001806106a46024913960400191505060405180910390fd5b6001600160a01b03821615156103c757604051600160e51b62461bcd02815260040180806020018281038252602281526020018061065d6022913960400191505060405180910390fd5b6001600160a01b03808416600081815260016020908152604080832094871680845294825291829020859055815185815291517f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b9259281900390910190a3505050565b6001600160a01b038316151561047357604051600160e51b62461bcd02815260040180806020018281038252602581526020018061067f6025913960400191505060405180910390fd5b6001600160a01b03821615156104bd57604051600160e51b62461bcd02815260040180806020018281038252602381526020018061063a6023913960400191505060405180910390fd5b6001600160a01b0383166000908152602081905260409020546104e6908263ffffffff61057516565b6001600160a01b03808516600090815260208190526040808220939093559084168152205461051b908263ffffffff6105d516565b6001600160a01b038084166000818152602081815260409182902094909455805185815290519193928716927fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef92918290030190a3505050565b6000828211156105cf5760408051600160e51b62461bcd02815260206004820152601e60248201527f536166654d6174683a207375627472616374696f6e206f766572666c6f770000604482015290519081900360640190fd5b50900390565b6000828201838110156106325760408051600160e51b62461bcd02815260206004820152601b60248201527f536166654d6174683a206164646974696f6e206f766572666c6f770000000000604482015290519081900360640190fd5b939250505056fe45524332303a207472616e7366657220746f20746865207a65726f206164647265737345524332303a20617070726f766520746f20746865207a65726f206164647265737345524332303a207472616e736665722066726f6d20746865207a65726f206164647265737345524332303a20617070726f76652066726f6d20746865207a65726f2061646472657373a165627a7a72305820dc5dba64afb4457ee29bfd043cba5f758ead09bf4b1daed5d77ab92a7a3d385a0029`

// DeployERC20 deploys a new Ethereum contract, binding an instance of ERC20 to it.
func DeployERC20(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *ERC20, error) {
	parsed, err := abi.JSON(strings.NewReader(ERC20ABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ERC20Bin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ERC20{ERC20Caller: ERC20Caller{contract: contract}, ERC20Transactor: ERC20Transactor{contract: contract}, ERC20Filterer: ERC20Filterer{contract: contract}}, nil
}

// ERC20 is an auto generated Go binding around an Ethereum contract.
type ERC20 struct {
	ERC20Caller     // Read-only binding to the contract
	ERC20Transactor // Write-only binding to the contract
	ERC20Filterer   // Log filterer for contract events
}

// ERC20Caller is an auto generated read-only Go binding around an Ethereum contract.
type ERC20Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC20Transactor is an auto generated write-only Go binding around an Ethereum contract.
type ERC20Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC20Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ERC20Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC20Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ERC20Session struct {
	Contract     *ERC20            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ERC20CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ERC20CallerSession struct {
	Contract *ERC20Caller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// ERC20TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ERC20TransactorSession struct {
	Contract     *ERC20Transactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ERC20Raw is an auto generated low-level Go binding around an Ethereum contract.
type ERC20Raw struct {
	Contract *ERC20 // Generic contract binding to access the raw methods on
}

// ERC20CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ERC20CallerRaw struct {
	Contract *ERC20Caller // Generic read-only contract binding to access the raw methods on
}

// ERC20TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ERC20TransactorRaw struct {
	Contract *ERC20Transactor // Generic write-only contract binding to access the raw methods on
}

// NewERC20 creates a new instance of ERC20, bound to a specific deployed contract.
func NewERC20(address common.Address, backend bind.ContractBackend) (*ERC20, error) {
	contract, err := bindERC20(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ERC20{ERC20Caller: ERC20Caller{contract: contract}, ERC20Transactor: ERC20Transactor{contract: contract}, ERC20Filterer: ERC20Filterer{contract: contract}}, nil
}

// NewERC20Caller creates a new read-only instance of ERC20, bound to a specific deployed contract.
func NewERC20Caller(address common.Address, caller bind.ContractCaller) (*ERC20Caller, error) {
	contract, err := bindERC20(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ERC20Caller{contract: contract}, nil
}

// NewERC20Transactor creates a new write-only instance of ERC20, bound to a specific deployed contract.
func NewERC20Transactor(address common.Address, transactor bind.ContractTransactor) (*ERC20Transactor, error) {
	contract, err := bindERC20(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ERC20Transactor{contract: contract}, nil
}

// NewERC20Filterer creates a new log filterer instance of ERC20, bound to a specific deployed contract.
func NewERC20Filterer(address common.Address, filterer bind.ContractFilterer) (*ERC20Filterer, error) {
	contract, err := bindERC20(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ERC20Filterer{contract: contract}, nil
}

// bindERC20 binds a generic wrapper to an already deployed contract.
func bindERC20(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ERC20ABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC20 *ERC20Raw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ERC20.Contract.ERC20Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC20 *ERC20Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC20.Contract.ERC20Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERC20 *ERC20Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ERC20.Contract.ERC20Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC20 *ERC20CallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ERC20.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC20 *ERC20TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC20.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERC20 *ERC20TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ERC20.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) constant returns(uint256)
func (_ERC20 *ERC20Caller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ERC20.contract.Call(opts, out, "allowance", owner, spender)
	return *ret0, err
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) constant returns(uint256)
func (_ERC20 *ERC20Session) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _ERC20.Contract.Allowance(&_ERC20.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) constant returns(uint256)
func (_ERC20 *ERC20CallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _ERC20.Contract.Allowance(&_ERC20.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) constant returns(uint256)
func (_ERC20 *ERC20Caller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ERC20.contract.Call(opts, out, "balanceOf", account)
	return *ret0, err
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) constant returns(uint256)
func (_ERC20 *ERC20Session) BalanceOf(account common.Address) (*big.Int, error) {
	return _ERC20.Contract.BalanceOf(&_ERC20.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) constant returns(uint256)
func (_ERC20 *ERC20CallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _ERC20.Contract.BalanceOf(&_ERC20.CallOpts, account)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_ERC20 *ERC20Caller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ERC20.contract.Call(opts, out, "totalSupply")
	return *ret0, err
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_ERC20 *ERC20Session) TotalSupply() (*big.Int, error) {
	return _ERC20.Contract.TotalSupply(&_ERC20.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_ERC20 *ERC20CallerSession) TotalSupply() (*big.Int, error) {
	return _ERC20.Contract.TotalSupply(&_ERC20.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_ERC20 *ERC20Transactor) Approve(opts *bind.TransactOpts, spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _ERC20.contract.Transact(opts, "approve", spender, value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_ERC20 *ERC20Session) Approve(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _ERC20.Contract.Approve(&_ERC20.TransactOpts, spender, value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_ERC20 *ERC20TransactorSession) Approve(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _ERC20.Contract.Approve(&_ERC20.TransactOpts, spender, value)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_ERC20 *ERC20Transactor) DecreaseAllowance(opts *bind.TransactOpts, spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _ERC20.contract.Transact(opts, "decreaseAllowance", spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_ERC20 *ERC20Session) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _ERC20.Contract.DecreaseAllowance(&_ERC20.TransactOpts, spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_ERC20 *ERC20TransactorSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _ERC20.Contract.DecreaseAllowance(&_ERC20.TransactOpts, spender, subtractedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_ERC20 *ERC20Transactor) IncreaseAllowance(opts *bind.TransactOpts, spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _ERC20.contract.Transact(opts, "increaseAllowance", spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_ERC20 *ERC20Session) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _ERC20.Contract.IncreaseAllowance(&_ERC20.TransactOpts, spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_ERC20 *ERC20TransactorSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _ERC20.Contract.IncreaseAllowance(&_ERC20.TransactOpts, spender, addedValue)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_ERC20 *ERC20Transactor) Transfer(opts *bind.TransactOpts, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ERC20.contract.Transact(opts, "transfer", recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_ERC20 *ERC20Session) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ERC20.Contract.Transfer(&_ERC20.TransactOpts, recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_ERC20 *ERC20TransactorSession) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ERC20.Contract.Transfer(&_ERC20.TransactOpts, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_ERC20 *ERC20Transactor) TransferFrom(opts *bind.TransactOpts, sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ERC20.contract.Transact(opts, "transferFrom", sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_ERC20 *ERC20Session) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ERC20.Contract.TransferFrom(&_ERC20.TransactOpts, sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_ERC20 *ERC20TransactorSession) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ERC20.Contract.TransferFrom(&_ERC20.TransactOpts, sender, recipient, amount)
}

// ERC20ApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the ERC20 contract.
type ERC20ApprovalIterator struct {
	Event *ERC20Approval // Event containing the contract specifics and raw log

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
func (it *ERC20ApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC20Approval)
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
		it.Event = new(ERC20Approval)
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
func (it *ERC20ApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC20ApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC20Approval represents a Approval event raised by the ERC20 contract.
type ERC20Approval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_ERC20 *ERC20Filterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*ERC20ApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _ERC20.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &ERC20ApprovalIterator{contract: _ERC20.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_ERC20 *ERC20Filterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *ERC20Approval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _ERC20.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC20Approval)
				if err := _ERC20.contract.UnpackLog(event, "Approval", log); err != nil {
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

// ERC20TransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the ERC20 contract.
type ERC20TransferIterator struct {
	Event *ERC20Transfer // Event containing the contract specifics and raw log

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
func (it *ERC20TransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC20Transfer)
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
		it.Event = new(ERC20Transfer)
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
func (it *ERC20TransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC20TransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC20Transfer represents a Transfer event raised by the ERC20 contract.
type ERC20Transfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_ERC20 *ERC20Filterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*ERC20TransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ERC20.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &ERC20TransferIterator{contract: _ERC20.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_ERC20 *ERC20Filterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *ERC20Transfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ERC20.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC20Transfer)
				if err := _ERC20.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// ERC20DetailedABI is the input ABI used to generate the binding from.
const ERC20DetailedABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"spender\",\"type\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"sender\",\"type\":\"address\"},{\"name\":\"recipient\",\"type\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"name\":\"\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"recipient\",\"type\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"owner\",\"type\":\"address\"},{\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"name\",\"type\":\"string\"},{\"name\":\"symbol\",\"type\":\"string\"},{\"name\":\"decimals\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"}]"

// ERC20DetailedBin is the compiled bytecode used for deploying new contracts.
const ERC20DetailedBin = `0x`

// DeployERC20Detailed deploys a new Ethereum contract, binding an instance of ERC20Detailed to it.
func DeployERC20Detailed(auth *bind.TransactOpts, backend bind.ContractBackend, name string, symbol string, decimals uint8) (common.Address, *types.Transaction, *ERC20Detailed, error) {
	parsed, err := abi.JSON(strings.NewReader(ERC20DetailedABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ERC20DetailedBin), backend, name, symbol, decimals)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ERC20Detailed{ERC20DetailedCaller: ERC20DetailedCaller{contract: contract}, ERC20DetailedTransactor: ERC20DetailedTransactor{contract: contract}, ERC20DetailedFilterer: ERC20DetailedFilterer{contract: contract}}, nil
}

// ERC20Detailed is an auto generated Go binding around an Ethereum contract.
type ERC20Detailed struct {
	ERC20DetailedCaller     // Read-only binding to the contract
	ERC20DetailedTransactor // Write-only binding to the contract
	ERC20DetailedFilterer   // Log filterer for contract events
}

// ERC20DetailedCaller is an auto generated read-only Go binding around an Ethereum contract.
type ERC20DetailedCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC20DetailedTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ERC20DetailedTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC20DetailedFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ERC20DetailedFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC20DetailedSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ERC20DetailedSession struct {
	Contract     *ERC20Detailed    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ERC20DetailedCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ERC20DetailedCallerSession struct {
	Contract *ERC20DetailedCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// ERC20DetailedTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ERC20DetailedTransactorSession struct {
	Contract     *ERC20DetailedTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// ERC20DetailedRaw is an auto generated low-level Go binding around an Ethereum contract.
type ERC20DetailedRaw struct {
	Contract *ERC20Detailed // Generic contract binding to access the raw methods on
}

// ERC20DetailedCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ERC20DetailedCallerRaw struct {
	Contract *ERC20DetailedCaller // Generic read-only contract binding to access the raw methods on
}

// ERC20DetailedTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ERC20DetailedTransactorRaw struct {
	Contract *ERC20DetailedTransactor // Generic write-only contract binding to access the raw methods on
}

// NewERC20Detailed creates a new instance of ERC20Detailed, bound to a specific deployed contract.
func NewERC20Detailed(address common.Address, backend bind.ContractBackend) (*ERC20Detailed, error) {
	contract, err := bindERC20Detailed(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ERC20Detailed{ERC20DetailedCaller: ERC20DetailedCaller{contract: contract}, ERC20DetailedTransactor: ERC20DetailedTransactor{contract: contract}, ERC20DetailedFilterer: ERC20DetailedFilterer{contract: contract}}, nil
}

// NewERC20DetailedCaller creates a new read-only instance of ERC20Detailed, bound to a specific deployed contract.
func NewERC20DetailedCaller(address common.Address, caller bind.ContractCaller) (*ERC20DetailedCaller, error) {
	contract, err := bindERC20Detailed(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ERC20DetailedCaller{contract: contract}, nil
}

// NewERC20DetailedTransactor creates a new write-only instance of ERC20Detailed, bound to a specific deployed contract.
func NewERC20DetailedTransactor(address common.Address, transactor bind.ContractTransactor) (*ERC20DetailedTransactor, error) {
	contract, err := bindERC20Detailed(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ERC20DetailedTransactor{contract: contract}, nil
}

// NewERC20DetailedFilterer creates a new log filterer instance of ERC20Detailed, bound to a specific deployed contract.
func NewERC20DetailedFilterer(address common.Address, filterer bind.ContractFilterer) (*ERC20DetailedFilterer, error) {
	contract, err := bindERC20Detailed(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ERC20DetailedFilterer{contract: contract}, nil
}

// bindERC20Detailed binds a generic wrapper to an already deployed contract.
func bindERC20Detailed(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ERC20DetailedABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC20Detailed *ERC20DetailedRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ERC20Detailed.Contract.ERC20DetailedCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC20Detailed *ERC20DetailedRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC20Detailed.Contract.ERC20DetailedTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERC20Detailed *ERC20DetailedRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ERC20Detailed.Contract.ERC20DetailedTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC20Detailed *ERC20DetailedCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ERC20Detailed.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC20Detailed *ERC20DetailedTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC20Detailed.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERC20Detailed *ERC20DetailedTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ERC20Detailed.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) constant returns(uint256)
func (_ERC20Detailed *ERC20DetailedCaller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ERC20Detailed.contract.Call(opts, out, "allowance", owner, spender)
	return *ret0, err
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) constant returns(uint256)
func (_ERC20Detailed *ERC20DetailedSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _ERC20Detailed.Contract.Allowance(&_ERC20Detailed.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) constant returns(uint256)
func (_ERC20Detailed *ERC20DetailedCallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _ERC20Detailed.Contract.Allowance(&_ERC20Detailed.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) constant returns(uint256)
func (_ERC20Detailed *ERC20DetailedCaller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ERC20Detailed.contract.Call(opts, out, "balanceOf", account)
	return *ret0, err
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) constant returns(uint256)
func (_ERC20Detailed *ERC20DetailedSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _ERC20Detailed.Contract.BalanceOf(&_ERC20Detailed.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) constant returns(uint256)
func (_ERC20Detailed *ERC20DetailedCallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _ERC20Detailed.Contract.BalanceOf(&_ERC20Detailed.CallOpts, account)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() constant returns(uint8)
func (_ERC20Detailed *ERC20DetailedCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var (
		ret0 = new(uint8)
	)
	out := ret0
	err := _ERC20Detailed.contract.Call(opts, out, "decimals")
	return *ret0, err
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() constant returns(uint8)
func (_ERC20Detailed *ERC20DetailedSession) Decimals() (uint8, error) {
	return _ERC20Detailed.Contract.Decimals(&_ERC20Detailed.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() constant returns(uint8)
func (_ERC20Detailed *ERC20DetailedCallerSession) Decimals() (uint8, error) {
	return _ERC20Detailed.Contract.Decimals(&_ERC20Detailed.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_ERC20Detailed *ERC20DetailedCaller) Name(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _ERC20Detailed.contract.Call(opts, out, "name")
	return *ret0, err
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_ERC20Detailed *ERC20DetailedSession) Name() (string, error) {
	return _ERC20Detailed.Contract.Name(&_ERC20Detailed.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_ERC20Detailed *ERC20DetailedCallerSession) Name() (string, error) {
	return _ERC20Detailed.Contract.Name(&_ERC20Detailed.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() constant returns(string)
func (_ERC20Detailed *ERC20DetailedCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _ERC20Detailed.contract.Call(opts, out, "symbol")
	return *ret0, err
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() constant returns(string)
func (_ERC20Detailed *ERC20DetailedSession) Symbol() (string, error) {
	return _ERC20Detailed.Contract.Symbol(&_ERC20Detailed.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() constant returns(string)
func (_ERC20Detailed *ERC20DetailedCallerSession) Symbol() (string, error) {
	return _ERC20Detailed.Contract.Symbol(&_ERC20Detailed.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_ERC20Detailed *ERC20DetailedCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ERC20Detailed.contract.Call(opts, out, "totalSupply")
	return *ret0, err
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_ERC20Detailed *ERC20DetailedSession) TotalSupply() (*big.Int, error) {
	return _ERC20Detailed.Contract.TotalSupply(&_ERC20Detailed.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_ERC20Detailed *ERC20DetailedCallerSession) TotalSupply() (*big.Int, error) {
	return _ERC20Detailed.Contract.TotalSupply(&_ERC20Detailed.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_ERC20Detailed *ERC20DetailedTransactor) Approve(opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ERC20Detailed.contract.Transact(opts, "approve", spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_ERC20Detailed *ERC20DetailedSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ERC20Detailed.Contract.Approve(&_ERC20Detailed.TransactOpts, spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_ERC20Detailed *ERC20DetailedTransactorSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ERC20Detailed.Contract.Approve(&_ERC20Detailed.TransactOpts, spender, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_ERC20Detailed *ERC20DetailedTransactor) Transfer(opts *bind.TransactOpts, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ERC20Detailed.contract.Transact(opts, "transfer", recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_ERC20Detailed *ERC20DetailedSession) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ERC20Detailed.Contract.Transfer(&_ERC20Detailed.TransactOpts, recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_ERC20Detailed *ERC20DetailedTransactorSession) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ERC20Detailed.Contract.Transfer(&_ERC20Detailed.TransactOpts, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_ERC20Detailed *ERC20DetailedTransactor) TransferFrom(opts *bind.TransactOpts, sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ERC20Detailed.contract.Transact(opts, "transferFrom", sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_ERC20Detailed *ERC20DetailedSession) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ERC20Detailed.Contract.TransferFrom(&_ERC20Detailed.TransactOpts, sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_ERC20Detailed *ERC20DetailedTransactorSession) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ERC20Detailed.Contract.TransferFrom(&_ERC20Detailed.TransactOpts, sender, recipient, amount)
}

// ERC20DetailedApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the ERC20Detailed contract.
type ERC20DetailedApprovalIterator struct {
	Event *ERC20DetailedApproval // Event containing the contract specifics and raw log

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
func (it *ERC20DetailedApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC20DetailedApproval)
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
		it.Event = new(ERC20DetailedApproval)
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
func (it *ERC20DetailedApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC20DetailedApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC20DetailedApproval represents a Approval event raised by the ERC20Detailed contract.
type ERC20DetailedApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_ERC20Detailed *ERC20DetailedFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*ERC20DetailedApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _ERC20Detailed.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &ERC20DetailedApprovalIterator{contract: _ERC20Detailed.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_ERC20Detailed *ERC20DetailedFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *ERC20DetailedApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _ERC20Detailed.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC20DetailedApproval)
				if err := _ERC20Detailed.contract.UnpackLog(event, "Approval", log); err != nil {
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

// ERC20DetailedTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the ERC20Detailed contract.
type ERC20DetailedTransferIterator struct {
	Event *ERC20DetailedTransfer // Event containing the contract specifics and raw log

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
func (it *ERC20DetailedTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC20DetailedTransfer)
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
		it.Event = new(ERC20DetailedTransfer)
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
func (it *ERC20DetailedTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC20DetailedTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC20DetailedTransfer represents a Transfer event raised by the ERC20Detailed contract.
type ERC20DetailedTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_ERC20Detailed *ERC20DetailedFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*ERC20DetailedTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ERC20Detailed.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &ERC20DetailedTransferIterator{contract: _ERC20Detailed.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_ERC20Detailed *ERC20DetailedFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *ERC20DetailedTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ERC20Detailed.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC20DetailedTransfer)
				if err := _ERC20Detailed.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// IERC20ABI is the input ABI used to generate the binding from.
const IERC20ABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"spender\",\"type\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"sender\",\"type\":\"address\"},{\"name\":\"recipient\",\"type\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"recipient\",\"type\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"owner\",\"type\":\"address\"},{\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"}]"

// IERC20Bin is the compiled bytecode used for deploying new contracts.
const IERC20Bin = `0x`

// DeployIERC20 deploys a new Ethereum contract, binding an instance of IERC20 to it.
func DeployIERC20(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *IERC20, error) {
	parsed, err := abi.JSON(strings.NewReader(IERC20ABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(IERC20Bin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &IERC20{IERC20Caller: IERC20Caller{contract: contract}, IERC20Transactor: IERC20Transactor{contract: contract}, IERC20Filterer: IERC20Filterer{contract: contract}}, nil
}

// IERC20 is an auto generated Go binding around an Ethereum contract.
type IERC20 struct {
	IERC20Caller     // Read-only binding to the contract
	IERC20Transactor // Write-only binding to the contract
	IERC20Filterer   // Log filterer for contract events
}

// IERC20Caller is an auto generated read-only Go binding around an Ethereum contract.
type IERC20Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IERC20Transactor is an auto generated write-only Go binding around an Ethereum contract.
type IERC20Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IERC20Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IERC20Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IERC20Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IERC20Session struct {
	Contract     *IERC20           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IERC20CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IERC20CallerSession struct {
	Contract *IERC20Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// IERC20TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IERC20TransactorSession struct {
	Contract     *IERC20Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IERC20Raw is an auto generated low-level Go binding around an Ethereum contract.
type IERC20Raw struct {
	Contract *IERC20 // Generic contract binding to access the raw methods on
}

// IERC20CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IERC20CallerRaw struct {
	Contract *IERC20Caller // Generic read-only contract binding to access the raw methods on
}

// IERC20TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IERC20TransactorRaw struct {
	Contract *IERC20Transactor // Generic write-only contract binding to access the raw methods on
}

// NewIERC20 creates a new instance of IERC20, bound to a specific deployed contract.
func NewIERC20(address common.Address, backend bind.ContractBackend) (*IERC20, error) {
	contract, err := bindIERC20(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IERC20{IERC20Caller: IERC20Caller{contract: contract}, IERC20Transactor: IERC20Transactor{contract: contract}, IERC20Filterer: IERC20Filterer{contract: contract}}, nil
}

// NewIERC20Caller creates a new read-only instance of IERC20, bound to a specific deployed contract.
func NewIERC20Caller(address common.Address, caller bind.ContractCaller) (*IERC20Caller, error) {
	contract, err := bindIERC20(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IERC20Caller{contract: contract}, nil
}

// NewIERC20Transactor creates a new write-only instance of IERC20, bound to a specific deployed contract.
func NewIERC20Transactor(address common.Address, transactor bind.ContractTransactor) (*IERC20Transactor, error) {
	contract, err := bindIERC20(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IERC20Transactor{contract: contract}, nil
}

// NewIERC20Filterer creates a new log filterer instance of IERC20, bound to a specific deployed contract.
func NewIERC20Filterer(address common.Address, filterer bind.ContractFilterer) (*IERC20Filterer, error) {
	contract, err := bindIERC20(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IERC20Filterer{contract: contract}, nil
}

// bindIERC20 binds a generic wrapper to an already deployed contract.
func bindIERC20(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IERC20ABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IERC20 *IERC20Raw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _IERC20.Contract.IERC20Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IERC20 *IERC20Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IERC20.Contract.IERC20Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IERC20 *IERC20Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IERC20.Contract.IERC20Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IERC20 *IERC20CallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _IERC20.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IERC20 *IERC20TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IERC20.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IERC20 *IERC20TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IERC20.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) constant returns(uint256)
func (_IERC20 *IERC20Caller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _IERC20.contract.Call(opts, out, "allowance", owner, spender)
	return *ret0, err
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) constant returns(uint256)
func (_IERC20 *IERC20Session) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _IERC20.Contract.Allowance(&_IERC20.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) constant returns(uint256)
func (_IERC20 *IERC20CallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _IERC20.Contract.Allowance(&_IERC20.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) constant returns(uint256)
func (_IERC20 *IERC20Caller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _IERC20.contract.Call(opts, out, "balanceOf", account)
	return *ret0, err
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) constant returns(uint256)
func (_IERC20 *IERC20Session) BalanceOf(account common.Address) (*big.Int, error) {
	return _IERC20.Contract.BalanceOf(&_IERC20.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) constant returns(uint256)
func (_IERC20 *IERC20CallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _IERC20.Contract.BalanceOf(&_IERC20.CallOpts, account)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_IERC20 *IERC20Caller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _IERC20.contract.Call(opts, out, "totalSupply")
	return *ret0, err
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_IERC20 *IERC20Session) TotalSupply() (*big.Int, error) {
	return _IERC20.Contract.TotalSupply(&_IERC20.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_IERC20 *IERC20CallerSession) TotalSupply() (*big.Int, error) {
	return _IERC20.Contract.TotalSupply(&_IERC20.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_IERC20 *IERC20Transactor) Approve(opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20.contract.Transact(opts, "approve", spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_IERC20 *IERC20Session) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20.Contract.Approve(&_IERC20.TransactOpts, spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_IERC20 *IERC20TransactorSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20.Contract.Approve(&_IERC20.TransactOpts, spender, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_IERC20 *IERC20Transactor) Transfer(opts *bind.TransactOpts, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20.contract.Transact(opts, "transfer", recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_IERC20 *IERC20Session) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20.Contract.Transfer(&_IERC20.TransactOpts, recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_IERC20 *IERC20TransactorSession) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20.Contract.Transfer(&_IERC20.TransactOpts, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_IERC20 *IERC20Transactor) TransferFrom(opts *bind.TransactOpts, sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20.contract.Transact(opts, "transferFrom", sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_IERC20 *IERC20Session) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20.Contract.TransferFrom(&_IERC20.TransactOpts, sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_IERC20 *IERC20TransactorSession) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20.Contract.TransferFrom(&_IERC20.TransactOpts, sender, recipient, amount)
}

// IERC20ApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the IERC20 contract.
type IERC20ApprovalIterator struct {
	Event *IERC20Approval // Event containing the contract specifics and raw log

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
func (it *IERC20ApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IERC20Approval)
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
		it.Event = new(IERC20Approval)
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
func (it *IERC20ApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IERC20ApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IERC20Approval represents a Approval event raised by the IERC20 contract.
type IERC20Approval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_IERC20 *IERC20Filterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*IERC20ApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _IERC20.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &IERC20ApprovalIterator{contract: _IERC20.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_IERC20 *IERC20Filterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *IERC20Approval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _IERC20.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IERC20Approval)
				if err := _IERC20.contract.UnpackLog(event, "Approval", log); err != nil {
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

// IERC20TransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the IERC20 contract.
type IERC20TransferIterator struct {
	Event *IERC20Transfer // Event containing the contract specifics and raw log

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
func (it *IERC20TransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IERC20Transfer)
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
		it.Event = new(IERC20Transfer)
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
func (it *IERC20TransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IERC20TransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IERC20Transfer represents a Transfer event raised by the IERC20 contract.
type IERC20Transfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_IERC20 *IERC20Filterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*IERC20TransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _IERC20.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &IERC20TransferIterator{contract: _IERC20.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_IERC20 *IERC20Filterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *IERC20Transfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _IERC20.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IERC20Transfer)
				if err := _IERC20.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// OraclizeAddrResolverIABI is the input ABI used to generate the binding from.
const OraclizeAddrResolverIABI = "[{\"constant\":false,\"inputs\":[],\"name\":\"getAddress\",\"outputs\":[{\"name\":\"_address\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// OraclizeAddrResolverIBin is the compiled bytecode used for deploying new contracts.
const OraclizeAddrResolverIBin = `0x`

// DeployOraclizeAddrResolverI deploys a new Ethereum contract, binding an instance of OraclizeAddrResolverI to it.
func DeployOraclizeAddrResolverI(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *OraclizeAddrResolverI, error) {
	parsed, err := abi.JSON(strings.NewReader(OraclizeAddrResolverIABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(OraclizeAddrResolverIBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &OraclizeAddrResolverI{OraclizeAddrResolverICaller: OraclizeAddrResolverICaller{contract: contract}, OraclizeAddrResolverITransactor: OraclizeAddrResolverITransactor{contract: contract}, OraclizeAddrResolverIFilterer: OraclizeAddrResolverIFilterer{contract: contract}}, nil
}

// OraclizeAddrResolverI is an auto generated Go binding around an Ethereum contract.
type OraclizeAddrResolverI struct {
	OraclizeAddrResolverICaller     // Read-only binding to the contract
	OraclizeAddrResolverITransactor // Write-only binding to the contract
	OraclizeAddrResolverIFilterer   // Log filterer for contract events
}

// OraclizeAddrResolverICaller is an auto generated read-only Go binding around an Ethereum contract.
type OraclizeAddrResolverICaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OraclizeAddrResolverITransactor is an auto generated write-only Go binding around an Ethereum contract.
type OraclizeAddrResolverITransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OraclizeAddrResolverIFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type OraclizeAddrResolverIFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OraclizeAddrResolverISession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type OraclizeAddrResolverISession struct {
	Contract     *OraclizeAddrResolverI // Generic contract binding to set the session for
	CallOpts     bind.CallOpts          // Call options to use throughout this session
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// OraclizeAddrResolverICallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type OraclizeAddrResolverICallerSession struct {
	Contract *OraclizeAddrResolverICaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                // Call options to use throughout this session
}

// OraclizeAddrResolverITransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type OraclizeAddrResolverITransactorSession struct {
	Contract     *OraclizeAddrResolverITransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                // Transaction auth options to use throughout this session
}

// OraclizeAddrResolverIRaw is an auto generated low-level Go binding around an Ethereum contract.
type OraclizeAddrResolverIRaw struct {
	Contract *OraclizeAddrResolverI // Generic contract binding to access the raw methods on
}

// OraclizeAddrResolverICallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type OraclizeAddrResolverICallerRaw struct {
	Contract *OraclizeAddrResolverICaller // Generic read-only contract binding to access the raw methods on
}

// OraclizeAddrResolverITransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type OraclizeAddrResolverITransactorRaw struct {
	Contract *OraclizeAddrResolverITransactor // Generic write-only contract binding to access the raw methods on
}

// NewOraclizeAddrResolverI creates a new instance of OraclizeAddrResolverI, bound to a specific deployed contract.
func NewOraclizeAddrResolverI(address common.Address, backend bind.ContractBackend) (*OraclizeAddrResolverI, error) {
	contract, err := bindOraclizeAddrResolverI(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &OraclizeAddrResolverI{OraclizeAddrResolverICaller: OraclizeAddrResolverICaller{contract: contract}, OraclizeAddrResolverITransactor: OraclizeAddrResolverITransactor{contract: contract}, OraclizeAddrResolverIFilterer: OraclizeAddrResolverIFilterer{contract: contract}}, nil
}

// NewOraclizeAddrResolverICaller creates a new read-only instance of OraclizeAddrResolverI, bound to a specific deployed contract.
func NewOraclizeAddrResolverICaller(address common.Address, caller bind.ContractCaller) (*OraclizeAddrResolverICaller, error) {
	contract, err := bindOraclizeAddrResolverI(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &OraclizeAddrResolverICaller{contract: contract}, nil
}

// NewOraclizeAddrResolverITransactor creates a new write-only instance of OraclizeAddrResolverI, bound to a specific deployed contract.
func NewOraclizeAddrResolverITransactor(address common.Address, transactor bind.ContractTransactor) (*OraclizeAddrResolverITransactor, error) {
	contract, err := bindOraclizeAddrResolverI(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &OraclizeAddrResolverITransactor{contract: contract}, nil
}

// NewOraclizeAddrResolverIFilterer creates a new log filterer instance of OraclizeAddrResolverI, bound to a specific deployed contract.
func NewOraclizeAddrResolverIFilterer(address common.Address, filterer bind.ContractFilterer) (*OraclizeAddrResolverIFilterer, error) {
	contract, err := bindOraclizeAddrResolverI(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &OraclizeAddrResolverIFilterer{contract: contract}, nil
}

// bindOraclizeAddrResolverI binds a generic wrapper to an already deployed contract.
func bindOraclizeAddrResolverI(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(OraclizeAddrResolverIABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_OraclizeAddrResolverI *OraclizeAddrResolverIRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _OraclizeAddrResolverI.Contract.OraclizeAddrResolverICaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_OraclizeAddrResolverI *OraclizeAddrResolverIRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OraclizeAddrResolverI.Contract.OraclizeAddrResolverITransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_OraclizeAddrResolverI *OraclizeAddrResolverIRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _OraclizeAddrResolverI.Contract.OraclizeAddrResolverITransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_OraclizeAddrResolverI *OraclizeAddrResolverICallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _OraclizeAddrResolverI.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_OraclizeAddrResolverI *OraclizeAddrResolverITransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OraclizeAddrResolverI.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_OraclizeAddrResolverI *OraclizeAddrResolverITransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _OraclizeAddrResolverI.Contract.contract.Transact(opts, method, params...)
}

// GetAddress is a paid mutator transaction binding the contract method 0x38cc4831.
//
// Solidity: function getAddress() returns(address _address)
func (_OraclizeAddrResolverI *OraclizeAddrResolverITransactor) GetAddress(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OraclizeAddrResolverI.contract.Transact(opts, "getAddress")
}

// GetAddress is a paid mutator transaction binding the contract method 0x38cc4831.
//
// Solidity: function getAddress() returns(address _address)
func (_OraclizeAddrResolverI *OraclizeAddrResolverISession) GetAddress() (*types.Transaction, error) {
	return _OraclizeAddrResolverI.Contract.GetAddress(&_OraclizeAddrResolverI.TransactOpts)
}

// GetAddress is a paid mutator transaction binding the contract method 0x38cc4831.
//
// Solidity: function getAddress() returns(address _address)
func (_OraclizeAddrResolverI *OraclizeAddrResolverITransactorSession) GetAddress() (*types.Transaction, error) {
	return _OraclizeAddrResolverI.Contract.GetAddress(&_OraclizeAddrResolverI.TransactOpts)
}

// OraclizeIABI is the input ABI used to generate the binding from.
const OraclizeIABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"_datasource\",\"type\":\"string\"},{\"name\":\"_gasLimit\",\"type\":\"uint256\"}],\"name\":\"getPrice\",\"outputs\":[{\"name\":\"_dsprice\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_datasource\",\"type\":\"string\"}],\"name\":\"getPrice\",\"outputs\":[{\"name\":\"_dsprice\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_proofType\",\"type\":\"bytes1\"}],\"name\":\"setProofType\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_timestamp\",\"type\":\"uint256\"},{\"name\":\"_datasource\",\"type\":\"string\"},{\"name\":\"_arg1\",\"type\":\"string\"},{\"name\":\"_arg2\",\"type\":\"string\"}],\"name\":\"query2\",\"outputs\":[{\"name\":\"_id\",\"type\":\"bytes32\"}],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_timestamp\",\"type\":\"uint256\"},{\"name\":\"_datasource\",\"type\":\"string\"},{\"name\":\"_argN\",\"type\":\"bytes\"}],\"name\":\"queryN\",\"outputs\":[{\"name\":\"_id\",\"type\":\"bytes32\"}],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_timestamp\",\"type\":\"uint256\"},{\"name\":\"_datasource\",\"type\":\"string\"},{\"name\":\"_arg1\",\"type\":\"string\"},{\"name\":\"_arg2\",\"type\":\"string\"},{\"name\":\"_gasLimit\",\"type\":\"uint256\"}],\"name\":\"query2_withGasLimit\",\"outputs\":[{\"name\":\"_id\",\"type\":\"bytes32\"}],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"randomDS_getSessionPubKeyHash\",\"outputs\":[{\"name\":\"_sessionKeyHash\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_timestamp\",\"type\":\"uint256\"},{\"name\":\"_datasource\",\"type\":\"string\"},{\"name\":\"_arg\",\"type\":\"string\"}],\"name\":\"query\",\"outputs\":[{\"name\":\"_id\",\"type\":\"bytes32\"}],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"cbAddress\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_timestamp\",\"type\":\"uint256\"},{\"name\":\"_datasource\",\"type\":\"string\"},{\"name\":\"_arg\",\"type\":\"string\"},{\"name\":\"_gasLimit\",\"type\":\"uint256\"}],\"name\":\"query_withGasLimit\",\"outputs\":[{\"name\":\"_id\",\"type\":\"bytes32\"}],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_timestamp\",\"type\":\"uint256\"},{\"name\":\"_datasource\",\"type\":\"string\"},{\"name\":\"_argN\",\"type\":\"bytes\"},{\"name\":\"_gasLimit\",\"type\":\"uint256\"}],\"name\":\"queryN_withGasLimit\",\"outputs\":[{\"name\":\"_id\",\"type\":\"bytes32\"}],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_gasPrice\",\"type\":\"uint256\"}],\"name\":\"setCustomGasPrice\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// OraclizeIBin is the compiled bytecode used for deploying new contracts.
const OraclizeIBin = `0x`

// DeployOraclizeI deploys a new Ethereum contract, binding an instance of OraclizeI to it.
func DeployOraclizeI(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *OraclizeI, error) {
	parsed, err := abi.JSON(strings.NewReader(OraclizeIABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(OraclizeIBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &OraclizeI{OraclizeICaller: OraclizeICaller{contract: contract}, OraclizeITransactor: OraclizeITransactor{contract: contract}, OraclizeIFilterer: OraclizeIFilterer{contract: contract}}, nil
}

// OraclizeI is an auto generated Go binding around an Ethereum contract.
type OraclizeI struct {
	OraclizeICaller     // Read-only binding to the contract
	OraclizeITransactor // Write-only binding to the contract
	OraclizeIFilterer   // Log filterer for contract events
}

// OraclizeICaller is an auto generated read-only Go binding around an Ethereum contract.
type OraclizeICaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OraclizeITransactor is an auto generated write-only Go binding around an Ethereum contract.
type OraclizeITransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OraclizeIFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type OraclizeIFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OraclizeISession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type OraclizeISession struct {
	Contract     *OraclizeI        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// OraclizeICallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type OraclizeICallerSession struct {
	Contract *OraclizeICaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// OraclizeITransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type OraclizeITransactorSession struct {
	Contract     *OraclizeITransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// OraclizeIRaw is an auto generated low-level Go binding around an Ethereum contract.
type OraclizeIRaw struct {
	Contract *OraclizeI // Generic contract binding to access the raw methods on
}

// OraclizeICallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type OraclizeICallerRaw struct {
	Contract *OraclizeICaller // Generic read-only contract binding to access the raw methods on
}

// OraclizeITransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type OraclizeITransactorRaw struct {
	Contract *OraclizeITransactor // Generic write-only contract binding to access the raw methods on
}

// NewOraclizeI creates a new instance of OraclizeI, bound to a specific deployed contract.
func NewOraclizeI(address common.Address, backend bind.ContractBackend) (*OraclizeI, error) {
	contract, err := bindOraclizeI(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &OraclizeI{OraclizeICaller: OraclizeICaller{contract: contract}, OraclizeITransactor: OraclizeITransactor{contract: contract}, OraclizeIFilterer: OraclizeIFilterer{contract: contract}}, nil
}

// NewOraclizeICaller creates a new read-only instance of OraclizeI, bound to a specific deployed contract.
func NewOraclizeICaller(address common.Address, caller bind.ContractCaller) (*OraclizeICaller, error) {
	contract, err := bindOraclizeI(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &OraclizeICaller{contract: contract}, nil
}

// NewOraclizeITransactor creates a new write-only instance of OraclizeI, bound to a specific deployed contract.
func NewOraclizeITransactor(address common.Address, transactor bind.ContractTransactor) (*OraclizeITransactor, error) {
	contract, err := bindOraclizeI(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &OraclizeITransactor{contract: contract}, nil
}

// NewOraclizeIFilterer creates a new log filterer instance of OraclizeI, bound to a specific deployed contract.
func NewOraclizeIFilterer(address common.Address, filterer bind.ContractFilterer) (*OraclizeIFilterer, error) {
	contract, err := bindOraclizeI(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &OraclizeIFilterer{contract: contract}, nil
}

// bindOraclizeI binds a generic wrapper to an already deployed contract.
func bindOraclizeI(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(OraclizeIABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_OraclizeI *OraclizeIRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _OraclizeI.Contract.OraclizeICaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_OraclizeI *OraclizeIRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OraclizeI.Contract.OraclizeITransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_OraclizeI *OraclizeIRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _OraclizeI.Contract.OraclizeITransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_OraclizeI *OraclizeICallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _OraclizeI.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_OraclizeI *OraclizeITransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OraclizeI.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_OraclizeI *OraclizeITransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _OraclizeI.Contract.contract.Transact(opts, method, params...)
}

// CbAddress is a free data retrieval call binding the contract method 0xc281d19e.
//
// Solidity: function cbAddress() constant returns(address)
func (_OraclizeI *OraclizeICaller) CbAddress(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _OraclizeI.contract.Call(opts, out, "cbAddress")
	return *ret0, err
}

// CbAddress is a free data retrieval call binding the contract method 0xc281d19e.
//
// Solidity: function cbAddress() constant returns(address)
func (_OraclizeI *OraclizeISession) CbAddress() (common.Address, error) {
	return _OraclizeI.Contract.CbAddress(&_OraclizeI.CallOpts)
}

// CbAddress is a free data retrieval call binding the contract method 0xc281d19e.
//
// Solidity: function cbAddress() constant returns(address)
func (_OraclizeI *OraclizeICallerSession) CbAddress() (common.Address, error) {
	return _OraclizeI.Contract.CbAddress(&_OraclizeI.CallOpts)
}

// RandomDSGetSessionPubKeyHash is a free data retrieval call binding the contract method 0xabaa5f3e.
//
// Solidity: function randomDS_getSessionPubKeyHash() constant returns(bytes32 _sessionKeyHash)
func (_OraclizeI *OraclizeICaller) RandomDSGetSessionPubKeyHash(opts *bind.CallOpts) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _OraclizeI.contract.Call(opts, out, "randomDS_getSessionPubKeyHash")
	return *ret0, err
}

// RandomDSGetSessionPubKeyHash is a free data retrieval call binding the contract method 0xabaa5f3e.
//
// Solidity: function randomDS_getSessionPubKeyHash() constant returns(bytes32 _sessionKeyHash)
func (_OraclizeI *OraclizeISession) RandomDSGetSessionPubKeyHash() ([32]byte, error) {
	return _OraclizeI.Contract.RandomDSGetSessionPubKeyHash(&_OraclizeI.CallOpts)
}

// RandomDSGetSessionPubKeyHash is a free data retrieval call binding the contract method 0xabaa5f3e.
//
// Solidity: function randomDS_getSessionPubKeyHash() constant returns(bytes32 _sessionKeyHash)
func (_OraclizeI *OraclizeICallerSession) RandomDSGetSessionPubKeyHash() ([32]byte, error) {
	return _OraclizeI.Contract.RandomDSGetSessionPubKeyHash(&_OraclizeI.CallOpts)
}

// GetPrice is a paid mutator transaction binding the contract method 0x524f3889.
//
// Solidity: function getPrice(string _datasource) returns(uint256 _dsprice)
func (_OraclizeI *OraclizeITransactor) GetPrice(opts *bind.TransactOpts, _datasource string) (*types.Transaction, error) {
	return _OraclizeI.contract.Transact(opts, "getPrice", _datasource)
}

// GetPrice is a paid mutator transaction binding the contract method 0x524f3889.
//
// Solidity: function getPrice(string _datasource) returns(uint256 _dsprice)
func (_OraclizeI *OraclizeISession) GetPrice(_datasource string) (*types.Transaction, error) {
	return _OraclizeI.Contract.GetPrice(&_OraclizeI.TransactOpts, _datasource)
}

// GetPrice is a paid mutator transaction binding the contract method 0x524f3889.
//
// Solidity: function getPrice(string _datasource) returns(uint256 _dsprice)
func (_OraclizeI *OraclizeITransactorSession) GetPrice(_datasource string) (*types.Transaction, error) {
	return _OraclizeI.Contract.GetPrice(&_OraclizeI.TransactOpts, _datasource)
}

// Query is a paid mutator transaction binding the contract method 0xadf59f99.
//
// Solidity: function query(uint256 _timestamp, string _datasource, string _arg) returns(bytes32 _id)
func (_OraclizeI *OraclizeITransactor) Query(opts *bind.TransactOpts, _timestamp *big.Int, _datasource string, _arg string) (*types.Transaction, error) {
	return _OraclizeI.contract.Transact(opts, "query", _timestamp, _datasource, _arg)
}

// Query is a paid mutator transaction binding the contract method 0xadf59f99.
//
// Solidity: function query(uint256 _timestamp, string _datasource, string _arg) returns(bytes32 _id)
func (_OraclizeI *OraclizeISession) Query(_timestamp *big.Int, _datasource string, _arg string) (*types.Transaction, error) {
	return _OraclizeI.Contract.Query(&_OraclizeI.TransactOpts, _timestamp, _datasource, _arg)
}

// Query is a paid mutator transaction binding the contract method 0xadf59f99.
//
// Solidity: function query(uint256 _timestamp, string _datasource, string _arg) returns(bytes32 _id)
func (_OraclizeI *OraclizeITransactorSession) Query(_timestamp *big.Int, _datasource string, _arg string) (*types.Transaction, error) {
	return _OraclizeI.Contract.Query(&_OraclizeI.TransactOpts, _timestamp, _datasource, _arg)
}

// Query2 is a paid mutator transaction binding the contract method 0x77228659.
//
// Solidity: function query2(uint256 _timestamp, string _datasource, string _arg1, string _arg2) returns(bytes32 _id)
func (_OraclizeI *OraclizeITransactor) Query2(opts *bind.TransactOpts, _timestamp *big.Int, _datasource string, _arg1 string, _arg2 string) (*types.Transaction, error) {
	return _OraclizeI.contract.Transact(opts, "query2", _timestamp, _datasource, _arg1, _arg2)
}

// Query2 is a paid mutator transaction binding the contract method 0x77228659.
//
// Solidity: function query2(uint256 _timestamp, string _datasource, string _arg1, string _arg2) returns(bytes32 _id)
func (_OraclizeI *OraclizeISession) Query2(_timestamp *big.Int, _datasource string, _arg1 string, _arg2 string) (*types.Transaction, error) {
	return _OraclizeI.Contract.Query2(&_OraclizeI.TransactOpts, _timestamp, _datasource, _arg1, _arg2)
}

// Query2 is a paid mutator transaction binding the contract method 0x77228659.
//
// Solidity: function query2(uint256 _timestamp, string _datasource, string _arg1, string _arg2) returns(bytes32 _id)
func (_OraclizeI *OraclizeITransactorSession) Query2(_timestamp *big.Int, _datasource string, _arg1 string, _arg2 string) (*types.Transaction, error) {
	return _OraclizeI.Contract.Query2(&_OraclizeI.TransactOpts, _timestamp, _datasource, _arg1, _arg2)
}

// Query2WithGasLimit is a paid mutator transaction binding the contract method 0x85dee34c.
//
// Solidity: function query2_withGasLimit(uint256 _timestamp, string _datasource, string _arg1, string _arg2, uint256 _gasLimit) returns(bytes32 _id)
func (_OraclizeI *OraclizeITransactor) Query2WithGasLimit(opts *bind.TransactOpts, _timestamp *big.Int, _datasource string, _arg1 string, _arg2 string, _gasLimit *big.Int) (*types.Transaction, error) {
	return _OraclizeI.contract.Transact(opts, "query2_withGasLimit", _timestamp, _datasource, _arg1, _arg2, _gasLimit)
}

// Query2WithGasLimit is a paid mutator transaction binding the contract method 0x85dee34c.
//
// Solidity: function query2_withGasLimit(uint256 _timestamp, string _datasource, string _arg1, string _arg2, uint256 _gasLimit) returns(bytes32 _id)
func (_OraclizeI *OraclizeISession) Query2WithGasLimit(_timestamp *big.Int, _datasource string, _arg1 string, _arg2 string, _gasLimit *big.Int) (*types.Transaction, error) {
	return _OraclizeI.Contract.Query2WithGasLimit(&_OraclizeI.TransactOpts, _timestamp, _datasource, _arg1, _arg2, _gasLimit)
}

// Query2WithGasLimit is a paid mutator transaction binding the contract method 0x85dee34c.
//
// Solidity: function query2_withGasLimit(uint256 _timestamp, string _datasource, string _arg1, string _arg2, uint256 _gasLimit) returns(bytes32 _id)
func (_OraclizeI *OraclizeITransactorSession) Query2WithGasLimit(_timestamp *big.Int, _datasource string, _arg1 string, _arg2 string, _gasLimit *big.Int) (*types.Transaction, error) {
	return _OraclizeI.Contract.Query2WithGasLimit(&_OraclizeI.TransactOpts, _timestamp, _datasource, _arg1, _arg2, _gasLimit)
}

// QueryN is a paid mutator transaction binding the contract method 0x83eed3d5.
//
// Solidity: function queryN(uint256 _timestamp, string _datasource, bytes _argN) returns(bytes32 _id)
func (_OraclizeI *OraclizeITransactor) QueryN(opts *bind.TransactOpts, _timestamp *big.Int, _datasource string, _argN []byte) (*types.Transaction, error) {
	return _OraclizeI.contract.Transact(opts, "queryN", _timestamp, _datasource, _argN)
}

// QueryN is a paid mutator transaction binding the contract method 0x83eed3d5.
//
// Solidity: function queryN(uint256 _timestamp, string _datasource, bytes _argN) returns(bytes32 _id)
func (_OraclizeI *OraclizeISession) QueryN(_timestamp *big.Int, _datasource string, _argN []byte) (*types.Transaction, error) {
	return _OraclizeI.Contract.QueryN(&_OraclizeI.TransactOpts, _timestamp, _datasource, _argN)
}

// QueryN is a paid mutator transaction binding the contract method 0x83eed3d5.
//
// Solidity: function queryN(uint256 _timestamp, string _datasource, bytes _argN) returns(bytes32 _id)
func (_OraclizeI *OraclizeITransactorSession) QueryN(_timestamp *big.Int, _datasource string, _argN []byte) (*types.Transaction, error) {
	return _OraclizeI.Contract.QueryN(&_OraclizeI.TransactOpts, _timestamp, _datasource, _argN)
}

// QueryNWithGasLimit is a paid mutator transaction binding the contract method 0xc55c1cb6.
//
// Solidity: function queryN_withGasLimit(uint256 _timestamp, string _datasource, bytes _argN, uint256 _gasLimit) returns(bytes32 _id)
func (_OraclizeI *OraclizeITransactor) QueryNWithGasLimit(opts *bind.TransactOpts, _timestamp *big.Int, _datasource string, _argN []byte, _gasLimit *big.Int) (*types.Transaction, error) {
	return _OraclizeI.contract.Transact(opts, "queryN_withGasLimit", _timestamp, _datasource, _argN, _gasLimit)
}

// QueryNWithGasLimit is a paid mutator transaction binding the contract method 0xc55c1cb6.
//
// Solidity: function queryN_withGasLimit(uint256 _timestamp, string _datasource, bytes _argN, uint256 _gasLimit) returns(bytes32 _id)
func (_OraclizeI *OraclizeISession) QueryNWithGasLimit(_timestamp *big.Int, _datasource string, _argN []byte, _gasLimit *big.Int) (*types.Transaction, error) {
	return _OraclizeI.Contract.QueryNWithGasLimit(&_OraclizeI.TransactOpts, _timestamp, _datasource, _argN, _gasLimit)
}

// QueryNWithGasLimit is a paid mutator transaction binding the contract method 0xc55c1cb6.
//
// Solidity: function queryN_withGasLimit(uint256 _timestamp, string _datasource, bytes _argN, uint256 _gasLimit) returns(bytes32 _id)
func (_OraclizeI *OraclizeITransactorSession) QueryNWithGasLimit(_timestamp *big.Int, _datasource string, _argN []byte, _gasLimit *big.Int) (*types.Transaction, error) {
	return _OraclizeI.Contract.QueryNWithGasLimit(&_OraclizeI.TransactOpts, _timestamp, _datasource, _argN, _gasLimit)
}

// QueryWithGasLimit is a paid mutator transaction binding the contract method 0xc51be90f.
//
// Solidity: function query_withGasLimit(uint256 _timestamp, string _datasource, string _arg, uint256 _gasLimit) returns(bytes32 _id)
func (_OraclizeI *OraclizeITransactor) QueryWithGasLimit(opts *bind.TransactOpts, _timestamp *big.Int, _datasource string, _arg string, _gasLimit *big.Int) (*types.Transaction, error) {
	return _OraclizeI.contract.Transact(opts, "query_withGasLimit", _timestamp, _datasource, _arg, _gasLimit)
}

// QueryWithGasLimit is a paid mutator transaction binding the contract method 0xc51be90f.
//
// Solidity: function query_withGasLimit(uint256 _timestamp, string _datasource, string _arg, uint256 _gasLimit) returns(bytes32 _id)
func (_OraclizeI *OraclizeISession) QueryWithGasLimit(_timestamp *big.Int, _datasource string, _arg string, _gasLimit *big.Int) (*types.Transaction, error) {
	return _OraclizeI.Contract.QueryWithGasLimit(&_OraclizeI.TransactOpts, _timestamp, _datasource, _arg, _gasLimit)
}

// QueryWithGasLimit is a paid mutator transaction binding the contract method 0xc51be90f.
//
// Solidity: function query_withGasLimit(uint256 _timestamp, string _datasource, string _arg, uint256 _gasLimit) returns(bytes32 _id)
func (_OraclizeI *OraclizeITransactorSession) QueryWithGasLimit(_timestamp *big.Int, _datasource string, _arg string, _gasLimit *big.Int) (*types.Transaction, error) {
	return _OraclizeI.Contract.QueryWithGasLimit(&_OraclizeI.TransactOpts, _timestamp, _datasource, _arg, _gasLimit)
}

// SetCustomGasPrice is a paid mutator transaction binding the contract method 0xca6ad1e4.
//
// Solidity: function setCustomGasPrice(uint256 _gasPrice) returns()
func (_OraclizeI *OraclizeITransactor) SetCustomGasPrice(opts *bind.TransactOpts, _gasPrice *big.Int) (*types.Transaction, error) {
	return _OraclizeI.contract.Transact(opts, "setCustomGasPrice", _gasPrice)
}

// SetCustomGasPrice is a paid mutator transaction binding the contract method 0xca6ad1e4.
//
// Solidity: function setCustomGasPrice(uint256 _gasPrice) returns()
func (_OraclizeI *OraclizeISession) SetCustomGasPrice(_gasPrice *big.Int) (*types.Transaction, error) {
	return _OraclizeI.Contract.SetCustomGasPrice(&_OraclizeI.TransactOpts, _gasPrice)
}

// SetCustomGasPrice is a paid mutator transaction binding the contract method 0xca6ad1e4.
//
// Solidity: function setCustomGasPrice(uint256 _gasPrice) returns()
func (_OraclizeI *OraclizeITransactorSession) SetCustomGasPrice(_gasPrice *big.Int) (*types.Transaction, error) {
	return _OraclizeI.Contract.SetCustomGasPrice(&_OraclizeI.TransactOpts, _gasPrice)
}

// SetProofType is a paid mutator transaction binding the contract method 0x688dcfd7.
//
// Solidity: function setProofType(bytes1 _proofType) returns()
func (_OraclizeI *OraclizeITransactor) SetProofType(opts *bind.TransactOpts, _proofType [1]byte) (*types.Transaction, error) {
	return _OraclizeI.contract.Transact(opts, "setProofType", _proofType)
}

// SetProofType is a paid mutator transaction binding the contract method 0x688dcfd7.
//
// Solidity: function setProofType(bytes1 _proofType) returns()
func (_OraclizeI *OraclizeISession) SetProofType(_proofType [1]byte) (*types.Transaction, error) {
	return _OraclizeI.Contract.SetProofType(&_OraclizeI.TransactOpts, _proofType)
}

// SetProofType is a paid mutator transaction binding the contract method 0x688dcfd7.
//
// Solidity: function setProofType(bytes1 _proofType) returns()
func (_OraclizeI *OraclizeITransactorSession) SetProofType(_proofType [1]byte) (*types.Transaction, error) {
	return _OraclizeI.Contract.SetProofType(&_OraclizeI.TransactOpts, _proofType)
}

// PriceFeederABI is the input ABI used to generate the binding from.
const PriceFeederABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"id\",\"type\":\"bytes32\"},{\"name\":\"result\",\"type\":\"string\"}],\"name\":\"__callback\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_myid\",\"type\":\"bytes32\"},{\"name\":\"_result\",\"type\":\"string\"},{\"name\":\"_proof\",\"type\":\"bytes\"}],\"name\":\"__callback\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"delay\",\"type\":\"uint256\"}],\"name\":\"updatePrice\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"digitalReserveSystemAddr\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"_digitalReserveSystemAddr\",\"type\":\"address\"}],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"nextStep\",\"type\":\"string\"}],\"name\":\"LogConstructorInitiated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"id\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"price\",\"type\":\"string\"}],\"name\":\"LogPriceUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"description\",\"type\":\"string\"}],\"name\":\"LogNewOraclizeQuery\",\"type\":\"event\"}]"

// PriceFeederBin is the compiled bytecode used for deploying new contracts.
const PriceFeederBin = `0x608060405260405160208061111d8339810180604052602081101561002357600080fd5b50516001805473845abb4b697dc2049db1fa0fd71ade29f268010e6001600160a01b031991821617909155600580549091166001600160a01b03831617905560408051602080825260259082018190527f11a3fca63f87bd67d7f9f72b744acc8be2193705e7a734ac3a773d35d259e87b928291908201906110f8823960400191505060405180910390a15061103a806100be6000396000f3fe60806040526004361061003f5760003560e01c806327dc297e1461004457806338bbfa50146101005780638d6cc56d14610241578063ec8820d81461025e575b600080fd5b34801561005057600080fd5b506100fe6004803603604081101561006757600080fd5b8135919081019060408101602082013564010000000081111561008957600080fd5b82018360208201111561009b57600080fd5b803590602001918460018302840111640100000000831117156100bd57600080fd5b91908080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525092955061028f945050505050565b005b34801561010c57600080fd5b506100fe6004803603606081101561012357600080fd5b8135919081019060408101602082013564010000000081111561014557600080fd5b82018360208201111561015757600080fd5b8035906020019184600183028401116401000000008311171561017957600080fd5b91908080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525092959493602081019350359150506401000000008111156101cc57600080fd5b8201836020820111156101de57600080fd5b8035906020019184600183028401116401000000008311171561020057600080fd5b91908080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152509295506103cb945050505050565b6100fe6004803603602081101561025757600080fd5b50356103fc565b34801561026a57600080fd5b5061027361048c565b604080516001600160a01b039092168252519081900360200190f35b61029761049b565b6001600160a01b031633146102ab57600080fd5b6005546001600160a01b03166310d8d74d6102c7836007610685565b6040518263ffffffff1660e01b81526004018080600160ea1b621554d102815250602001828152602001915050600060405180830381600087803b15801561030e57600080fd5b505af1158015610322573d6000803e3d6000fd5b505050507f81bb956107492f3a02fe41016dcffb7e279b1f695887cd1c6ca7c44e4825067f82826040518083815260200180602001828103825283818151815260200191508051906020019080838360005b8381101561038c578181015183820152602001610374565b50505050905090810190601f1680156103b95780820380516001836020036101000a031916815260200191505b50935050505060405180910390a15050565b5050600080805260036020527f3617319a054d772f909f7c479a2cebe5066e836a939412e32403c99029b92eff5550565b7f621c2856e3b87f81235f8ac8a22bbb40a0142961960710d00b2b6c380902b57e604051808060200182810382526035815260200180610fda6035913960400191505060405180910390a161048881604051806040016040528060038152602001600160ea1b62155493028152506040518060a0016040528060648152602001610f7660649139610784565b5050565b6005546001600160a01b031681565b6001546000906001600160a01b031615806104c857506001546104c6906001600160a01b0316610b50565b155b156104d9576104d76000610b54565b505b600160009054906101000a90046001600160a01b03166001600160a01b03166338cc48316040518163ffffffff1660e01b8152600401602060405180830381600087803b15801561052957600080fd5b505af115801561053d573d6000803e3d6000fd5b505050506040513d602081101561055357600080fd5b50516000546001600160a01b0390811691161461060657600160009054906101000a90046001600160a01b03166001600160a01b03166338cc48316040518163ffffffff1660e01b8152600401602060405180830381600087803b1580156105ba57600080fd5b505af11580156105ce573d6000803e3d6000fd5b505050506040513d60208110156105e457600080fd5b5051600080546001600160a01b0319166001600160a01b039092169190911790555b6000809054906101000a90046001600160a01b03166001600160a01b031663c281d19e6040518163ffffffff1660e01b815260040160206040518083038186803b15801561065357600080fd5b505afa158015610667573d6000803e3d6000fd5b505050506040513d602081101561067d57600080fd5b505190505b90565b6000828180805b835181101561076b57603084828151811015156106a557fe5b90602001015160f81c60f81b60f81c60ff16101580156106e55750603984828151811015156106d057fe5b90602001015160f81c60f81b60f81c60ff1611155b15610736578115610704578515156106fc5761076b565b600019909501945b600a830292506030848281518110151561071a57fe5b90602001015160f81c60f81b60f81c60ff160383019250610763565b838181518110151561074457fe5b90602001015160f81c60f81b60f81c60ff16602e141561076357600191505b60010161068c565b50841561077b5784600a0a820291505b50949350505050565b6001546000906001600160a01b031615806107b157506001546107af906001600160a01b0316610b50565b155b156107c2576107c06000610b54565b505b600160009054906101000a90046001600160a01b03166001600160a01b03166338cc48316040518163ffffffff1660e01b8152600401602060405180830381600087803b15801561081257600080fd5b505af1158015610826573d6000803e3d6000fd5b505050506040513d602081101561083c57600080fd5b50516000546001600160a01b039081169116146108ef57600160009054906101000a90046001600160a01b03166001600160a01b03166338cc48316040518163ffffffff1660e01b8152600401602060405180830381600087803b1580156108a357600080fd5b505af11580156108b7573d6000803e3d6000fd5b505050506040513d60208110156108cd57600080fd5b5051600080546001600160a01b0319166001600160a01b039092169190911790555b60008054604051600160e01b63524f38890281526020600482018181528751602484015287516001600160a01b039094169363524f38899389938392604490920191908501908083838b5b8381101561095257818101518382015260200161093a565b50505050905090810190601f16801561097f5780820380516001836020036101000a031916815260200191505b5092505050602060405180830381600087803b15801561099e57600080fd5b505af11580156109b2573d6000803e3d6000fd5b505050506040513d60208110156109c857600080fd5b50519050670de0b6b3a764000062030d403a02018111156109ed575060009050610b49565b6000809054906101000a90046001600160a01b03166001600160a01b031663adf59f99828787876040518563ffffffff1660e01b8152600401808481526020018060200180602001838103835285818151815260200191508051906020019080838360005b83811015610a6a578181015183820152602001610a52565b50505050905090810190601f168015610a975780820380516001836020036101000a031916815260200191505b50838103825284518152845160209182019186019080838360005b83811015610aca578181015183820152602001610ab2565b50505050905090810190601f168015610af75780820380516001836020036101000a031916815260200191505b50955050505050506020604051808303818588803b158015610b1857600080fd5b505af1158015610b2c573d6000803e3d6000fd5b50505050506040513d6020811015610b4357600080fd5b50519150505b9392505050565b3b90565b6000610b5e610b64565b92915050565b600080610b84731d3b2638a7cc9f2cb3d298a3da7a90b67e5506ed610b50565b1115610be657600180546001600160a01b031916731d3b2638a7cc9f2cb3d298a3da7a90b67e5506ed17905560408051808201909152600b8152600160aa1b6a195d1a17db585a5b9b995d026020820152610bde90610ecf565b506001610682565b6000610c0573c03a2615d5efaf5f49f60b7bb6583eaec212fdf1610b50565b1115610c6e57600180546001600160a01b03191673c03a2615d5efaf5f49f60b7bb6583eaec212fdf117905560408051808201909152600c81527f6574685f726f707374656e3300000000000000000000000000000000000000006020820152610bde90610ecf565b6000610c8d73b7a07bcf2ba2f2703b24c0691b5278999c59ac7e610b50565b1115610ce557600180546001600160a01b03191673b7a07bcf2ba2f2703b24c0691b5278999c59ac7e1790556040805180820190915260098152600160b91b6832ba342fb5b7bb30b7026020820152610bde90610ecf565b6000610d0473146500cfd35b22e4a392fe0adc06de1a1368ed48610b50565b1115610d5e57600180546001600160a01b03191673146500cfd35b22e4a392fe0adc06de1a1368ed4817905560408051808201909152600b8152600160a81b6a6574685f72696e6b656279026020820152610bde90610ecf565b6000610d7d73a2998efd205fb9d4b4963afb70778d6354ad3a41610b50565b1115610dd657600180546001600160a01b03191673a2998efd205fb9d4b4963afb70778d6354ad3a4117905560408051808201909152600a8152600160b01b696574685f676f65726c69026020820152610bde90610ecf565b6000610df5736f485c8bf6fc43ea212e93bbf8ce046c7f1cb475610b50565b1115610e275750600180546001600160a01b031916736f485c8bf6fc43ea212e93bbf8ce046c7f1cb475178155610682565b6000610e467320e12a1f859b3feae5fb2a0a32c18f5a65555bbf610b50565b1115610e785750600180546001600160a01b0319167320e12a1f859b3feae5fb2a0a32c18f5a65555bbf178155610682565b6000610e977351efaf4c8b3c9afbd5ab9f4bbc82784ab6ef8faa610b50565b1115610ec95750600180546001600160a01b0319167351efaf4c8b3c9afbd5ab9f4bbc82784ab6ef8faa178155610682565b50600090565b8051610488906002906020840190828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f10610f1e57805160ff1916838001178555610f4b565b82800160010185558215610f4b579182015b82811115610f4b578251825591602001919060010190610f30565b50610f57929150610f5b565b5090565b61068291905b80821115610f575760008155600101610f6156fe6a736f6e2868747470733a2f2f6465762d696e7465726e616c2d73797374656d2e6c696768746e6574617069732e696f2f76656c6f2d66782d72617465733f636f6e766572743d555344292e646174612e2231222e71756f74652e5553442e70726963656f7261636c697a65207175657279207761732073656e742c207374616e64696e6720627920666f722074686520616e737765722e2ea165627a7a72305820752e280f24d34bf61cfef3a947441e227be84fd3a6d3fb10e32ba6dc0550bc1400296050726963654665656465726020636f6e74726163742077617320696e697469617465642e`

// DeployPriceFeeder deploys a new Ethereum contract, binding an instance of PriceFeeder to it.
func DeployPriceFeeder(auth *bind.TransactOpts, backend bind.ContractBackend, _digitalReserveSystemAddr common.Address) (common.Address, *types.Transaction, *PriceFeeder, error) {
	parsed, err := abi.JSON(strings.NewReader(PriceFeederABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(PriceFeederBin), backend, _digitalReserveSystemAddr)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &PriceFeeder{PriceFeederCaller: PriceFeederCaller{contract: contract}, PriceFeederTransactor: PriceFeederTransactor{contract: contract}, PriceFeederFilterer: PriceFeederFilterer{contract: contract}}, nil
}

// PriceFeeder is an auto generated Go binding around an Ethereum contract.
type PriceFeeder struct {
	PriceFeederCaller     // Read-only binding to the contract
	PriceFeederTransactor // Write-only binding to the contract
	PriceFeederFilterer   // Log filterer for contract events
}

// PriceFeederCaller is an auto generated read-only Go binding around an Ethereum contract.
type PriceFeederCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PriceFeederTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PriceFeederTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PriceFeederFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PriceFeederFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PriceFeederSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PriceFeederSession struct {
	Contract     *PriceFeeder      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PriceFeederCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PriceFeederCallerSession struct {
	Contract *PriceFeederCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// PriceFeederTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PriceFeederTransactorSession struct {
	Contract     *PriceFeederTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// PriceFeederRaw is an auto generated low-level Go binding around an Ethereum contract.
type PriceFeederRaw struct {
	Contract *PriceFeeder // Generic contract binding to access the raw methods on
}

// PriceFeederCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PriceFeederCallerRaw struct {
	Contract *PriceFeederCaller // Generic read-only contract binding to access the raw methods on
}

// PriceFeederTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PriceFeederTransactorRaw struct {
	Contract *PriceFeederTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPriceFeeder creates a new instance of PriceFeeder, bound to a specific deployed contract.
func NewPriceFeeder(address common.Address, backend bind.ContractBackend) (*PriceFeeder, error) {
	contract, err := bindPriceFeeder(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &PriceFeeder{PriceFeederCaller: PriceFeederCaller{contract: contract}, PriceFeederTransactor: PriceFeederTransactor{contract: contract}, PriceFeederFilterer: PriceFeederFilterer{contract: contract}}, nil
}

// NewPriceFeederCaller creates a new read-only instance of PriceFeeder, bound to a specific deployed contract.
func NewPriceFeederCaller(address common.Address, caller bind.ContractCaller) (*PriceFeederCaller, error) {
	contract, err := bindPriceFeeder(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PriceFeederCaller{contract: contract}, nil
}

// NewPriceFeederTransactor creates a new write-only instance of PriceFeeder, bound to a specific deployed contract.
func NewPriceFeederTransactor(address common.Address, transactor bind.ContractTransactor) (*PriceFeederTransactor, error) {
	contract, err := bindPriceFeeder(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PriceFeederTransactor{contract: contract}, nil
}

// NewPriceFeederFilterer creates a new log filterer instance of PriceFeeder, bound to a specific deployed contract.
func NewPriceFeederFilterer(address common.Address, filterer bind.ContractFilterer) (*PriceFeederFilterer, error) {
	contract, err := bindPriceFeeder(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PriceFeederFilterer{contract: contract}, nil
}

// bindPriceFeeder binds a generic wrapper to an already deployed contract.
func bindPriceFeeder(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(PriceFeederABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PriceFeeder *PriceFeederRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _PriceFeeder.Contract.PriceFeederCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PriceFeeder *PriceFeederRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PriceFeeder.Contract.PriceFeederTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PriceFeeder *PriceFeederRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PriceFeeder.Contract.PriceFeederTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PriceFeeder *PriceFeederCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _PriceFeeder.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PriceFeeder *PriceFeederTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PriceFeeder.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PriceFeeder *PriceFeederTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PriceFeeder.Contract.contract.Transact(opts, method, params...)
}

// DigitalReserveSystemAddr is a free data retrieval call binding the contract method 0xec8820d8.
//
// Solidity: function digitalReserveSystemAddr() constant returns(address)
func (_PriceFeeder *PriceFeederCaller) DigitalReserveSystemAddr(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _PriceFeeder.contract.Call(opts, out, "digitalReserveSystemAddr")
	return *ret0, err
}

// DigitalReserveSystemAddr is a free data retrieval call binding the contract method 0xec8820d8.
//
// Solidity: function digitalReserveSystemAddr() constant returns(address)
func (_PriceFeeder *PriceFeederSession) DigitalReserveSystemAddr() (common.Address, error) {
	return _PriceFeeder.Contract.DigitalReserveSystemAddr(&_PriceFeeder.CallOpts)
}

// DigitalReserveSystemAddr is a free data retrieval call binding the contract method 0xec8820d8.
//
// Solidity: function digitalReserveSystemAddr() constant returns(address)
func (_PriceFeeder *PriceFeederCallerSession) DigitalReserveSystemAddr() (common.Address, error) {
	return _PriceFeeder.Contract.DigitalReserveSystemAddr(&_PriceFeeder.CallOpts)
}

// Callback is a paid mutator transaction binding the contract method 0x38bbfa50.
//
// Solidity: function __callback(bytes32 _myid, string _result, bytes _proof) returns()
func (_PriceFeeder *PriceFeederTransactor) Callback(opts *bind.TransactOpts, _myid [32]byte, _result string, _proof []byte) (*types.Transaction, error) {
	return _PriceFeeder.contract.Transact(opts, "__callback", _myid, _result, _proof)
}

// Callback is a paid mutator transaction binding the contract method 0x38bbfa50.
//
// Solidity: function __callback(bytes32 _myid, string _result, bytes _proof) returns()
func (_PriceFeeder *PriceFeederSession) Callback(_myid [32]byte, _result string, _proof []byte) (*types.Transaction, error) {
	return _PriceFeeder.Contract.Callback(&_PriceFeeder.TransactOpts, _myid, _result, _proof)
}

// Callback is a paid mutator transaction binding the contract method 0x38bbfa50.
//
// Solidity: function __callback(bytes32 _myid, string _result, bytes _proof) returns()
func (_PriceFeeder *PriceFeederTransactorSession) Callback(_myid [32]byte, _result string, _proof []byte) (*types.Transaction, error) {
	return _PriceFeeder.Contract.Callback(&_PriceFeeder.TransactOpts, _myid, _result, _proof)
}

// UpdatePrice is a paid mutator transaction binding the contract method 0x8d6cc56d.
//
// Solidity: function updatePrice(uint256 delay) returns()
func (_PriceFeeder *PriceFeederTransactor) UpdatePrice(opts *bind.TransactOpts, delay *big.Int) (*types.Transaction, error) {
	return _PriceFeeder.contract.Transact(opts, "updatePrice", delay)
}

// UpdatePrice is a paid mutator transaction binding the contract method 0x8d6cc56d.
//
// Solidity: function updatePrice(uint256 delay) returns()
func (_PriceFeeder *PriceFeederSession) UpdatePrice(delay *big.Int) (*types.Transaction, error) {
	return _PriceFeeder.Contract.UpdatePrice(&_PriceFeeder.TransactOpts, delay)
}

// UpdatePrice is a paid mutator transaction binding the contract method 0x8d6cc56d.
//
// Solidity: function updatePrice(uint256 delay) returns()
func (_PriceFeeder *PriceFeederTransactorSession) UpdatePrice(delay *big.Int) (*types.Transaction, error) {
	return _PriceFeeder.Contract.UpdatePrice(&_PriceFeeder.TransactOpts, delay)
}

// PriceFeederLogConstructorInitiatedIterator is returned from FilterLogConstructorInitiated and is used to iterate over the raw logs and unpacked data for LogConstructorInitiated events raised by the PriceFeeder contract.
type PriceFeederLogConstructorInitiatedIterator struct {
	Event *PriceFeederLogConstructorInitiated // Event containing the contract specifics and raw log

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
func (it *PriceFeederLogConstructorInitiatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PriceFeederLogConstructorInitiated)
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
		it.Event = new(PriceFeederLogConstructorInitiated)
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
func (it *PriceFeederLogConstructorInitiatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PriceFeederLogConstructorInitiatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PriceFeederLogConstructorInitiated represents a LogConstructorInitiated event raised by the PriceFeeder contract.
type PriceFeederLogConstructorInitiated struct {
	NextStep string
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterLogConstructorInitiated is a free log retrieval operation binding the contract event 0x11a3fca63f87bd67d7f9f72b744acc8be2193705e7a734ac3a773d35d259e87b.
//
// Solidity: event LogConstructorInitiated(string nextStep)
func (_PriceFeeder *PriceFeederFilterer) FilterLogConstructorInitiated(opts *bind.FilterOpts) (*PriceFeederLogConstructorInitiatedIterator, error) {

	logs, sub, err := _PriceFeeder.contract.FilterLogs(opts, "LogConstructorInitiated")
	if err != nil {
		return nil, err
	}
	return &PriceFeederLogConstructorInitiatedIterator{contract: _PriceFeeder.contract, event: "LogConstructorInitiated", logs: logs, sub: sub}, nil
}

// WatchLogConstructorInitiated is a free log subscription operation binding the contract event 0x11a3fca63f87bd67d7f9f72b744acc8be2193705e7a734ac3a773d35d259e87b.
//
// Solidity: event LogConstructorInitiated(string nextStep)
func (_PriceFeeder *PriceFeederFilterer) WatchLogConstructorInitiated(opts *bind.WatchOpts, sink chan<- *PriceFeederLogConstructorInitiated) (event.Subscription, error) {

	logs, sub, err := _PriceFeeder.contract.WatchLogs(opts, "LogConstructorInitiated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PriceFeederLogConstructorInitiated)
				if err := _PriceFeeder.contract.UnpackLog(event, "LogConstructorInitiated", log); err != nil {
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

// PriceFeederLogNewOraclizeQueryIterator is returned from FilterLogNewOraclizeQuery and is used to iterate over the raw logs and unpacked data for LogNewOraclizeQuery events raised by the PriceFeeder contract.
type PriceFeederLogNewOraclizeQueryIterator struct {
	Event *PriceFeederLogNewOraclizeQuery // Event containing the contract specifics and raw log

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
func (it *PriceFeederLogNewOraclizeQueryIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PriceFeederLogNewOraclizeQuery)
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
		it.Event = new(PriceFeederLogNewOraclizeQuery)
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
func (it *PriceFeederLogNewOraclizeQueryIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PriceFeederLogNewOraclizeQueryIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PriceFeederLogNewOraclizeQuery represents a LogNewOraclizeQuery event raised by the PriceFeeder contract.
type PriceFeederLogNewOraclizeQuery struct {
	Description string
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterLogNewOraclizeQuery is a free log retrieval operation binding the contract event 0x621c2856e3b87f81235f8ac8a22bbb40a0142961960710d00b2b6c380902b57e.
//
// Solidity: event LogNewOraclizeQuery(string description)
func (_PriceFeeder *PriceFeederFilterer) FilterLogNewOraclizeQuery(opts *bind.FilterOpts) (*PriceFeederLogNewOraclizeQueryIterator, error) {

	logs, sub, err := _PriceFeeder.contract.FilterLogs(opts, "LogNewOraclizeQuery")
	if err != nil {
		return nil, err
	}
	return &PriceFeederLogNewOraclizeQueryIterator{contract: _PriceFeeder.contract, event: "LogNewOraclizeQuery", logs: logs, sub: sub}, nil
}

// WatchLogNewOraclizeQuery is a free log subscription operation binding the contract event 0x621c2856e3b87f81235f8ac8a22bbb40a0142961960710d00b2b6c380902b57e.
//
// Solidity: event LogNewOraclizeQuery(string description)
func (_PriceFeeder *PriceFeederFilterer) WatchLogNewOraclizeQuery(opts *bind.WatchOpts, sink chan<- *PriceFeederLogNewOraclizeQuery) (event.Subscription, error) {

	logs, sub, err := _PriceFeeder.contract.WatchLogs(opts, "LogNewOraclizeQuery")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PriceFeederLogNewOraclizeQuery)
				if err := _PriceFeeder.contract.UnpackLog(event, "LogNewOraclizeQuery", log); err != nil {
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

// PriceFeederLogPriceUpdatedIterator is returned from FilterLogPriceUpdated and is used to iterate over the raw logs and unpacked data for LogPriceUpdated events raised by the PriceFeeder contract.
type PriceFeederLogPriceUpdatedIterator struct {
	Event *PriceFeederLogPriceUpdated // Event containing the contract specifics and raw log

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
func (it *PriceFeederLogPriceUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PriceFeederLogPriceUpdated)
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
		it.Event = new(PriceFeederLogPriceUpdated)
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
func (it *PriceFeederLogPriceUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PriceFeederLogPriceUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PriceFeederLogPriceUpdated represents a LogPriceUpdated event raised by the PriceFeeder contract.
type PriceFeederLogPriceUpdated struct {
	Id    [32]byte
	Price string
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterLogPriceUpdated is a free log retrieval operation binding the contract event 0x81bb956107492f3a02fe41016dcffb7e279b1f695887cd1c6ca7c44e4825067f.
//
// Solidity: event LogPriceUpdated(bytes32 id, string price)
func (_PriceFeeder *PriceFeederFilterer) FilterLogPriceUpdated(opts *bind.FilterOpts) (*PriceFeederLogPriceUpdatedIterator, error) {

	logs, sub, err := _PriceFeeder.contract.FilterLogs(opts, "LogPriceUpdated")
	if err != nil {
		return nil, err
	}
	return &PriceFeederLogPriceUpdatedIterator{contract: _PriceFeeder.contract, event: "LogPriceUpdated", logs: logs, sub: sub}, nil
}

// WatchLogPriceUpdated is a free log subscription operation binding the contract event 0x81bb956107492f3a02fe41016dcffb7e279b1f695887cd1c6ca7c44e4825067f.
//
// Solidity: event LogPriceUpdated(bytes32 id, string price)
func (_PriceFeeder *PriceFeederFilterer) WatchLogPriceUpdated(opts *bind.WatchOpts, sink chan<- *PriceFeederLogPriceUpdated) (event.Subscription, error) {

	logs, sub, err := _PriceFeeder.contract.WatchLogs(opts, "LogPriceUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PriceFeederLogPriceUpdated)
				if err := _PriceFeeder.contract.UnpackLog(event, "LogPriceUpdated", log); err != nil {
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

// RolesABI is the input ABI used to generate the binding from.
const RolesABI = "[]"

// RolesBin is the compiled bytecode used for deploying new contracts.
const RolesBin = `0x604c6025600b82828239805160001a6073141515601857fe5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea165627a7a723058201c97c31d7ec32620cb48c1da9153fcee0d5cc770477a7a31b8e9c125b55de3cc0029`

// DeployRoles deploys a new Ethereum contract, binding an instance of Roles to it.
func DeployRoles(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Roles, error) {
	parsed, err := abi.JSON(strings.NewReader(RolesABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(RolesBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Roles{RolesCaller: RolesCaller{contract: contract}, RolesTransactor: RolesTransactor{contract: contract}, RolesFilterer: RolesFilterer{contract: contract}}, nil
}

// Roles is an auto generated Go binding around an Ethereum contract.
type Roles struct {
	RolesCaller     // Read-only binding to the contract
	RolesTransactor // Write-only binding to the contract
	RolesFilterer   // Log filterer for contract events
}

// RolesCaller is an auto generated read-only Go binding around an Ethereum contract.
type RolesCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RolesTransactor is an auto generated write-only Go binding around an Ethereum contract.
type RolesTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RolesFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type RolesFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RolesSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type RolesSession struct {
	Contract     *Roles            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// RolesCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type RolesCallerSession struct {
	Contract *RolesCaller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// RolesTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type RolesTransactorSession struct {
	Contract     *RolesTransactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// RolesRaw is an auto generated low-level Go binding around an Ethereum contract.
type RolesRaw struct {
	Contract *Roles // Generic contract binding to access the raw methods on
}

// RolesCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type RolesCallerRaw struct {
	Contract *RolesCaller // Generic read-only contract binding to access the raw methods on
}

// RolesTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type RolesTransactorRaw struct {
	Contract *RolesTransactor // Generic write-only contract binding to access the raw methods on
}

// NewRoles creates a new instance of Roles, bound to a specific deployed contract.
func NewRoles(address common.Address, backend bind.ContractBackend) (*Roles, error) {
	contract, err := bindRoles(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Roles{RolesCaller: RolesCaller{contract: contract}, RolesTransactor: RolesTransactor{contract: contract}, RolesFilterer: RolesFilterer{contract: contract}}, nil
}

// NewRolesCaller creates a new read-only instance of Roles, bound to a specific deployed contract.
func NewRolesCaller(address common.Address, caller bind.ContractCaller) (*RolesCaller, error) {
	contract, err := bindRoles(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &RolesCaller{contract: contract}, nil
}

// NewRolesTransactor creates a new write-only instance of Roles, bound to a specific deployed contract.
func NewRolesTransactor(address common.Address, transactor bind.ContractTransactor) (*RolesTransactor, error) {
	contract, err := bindRoles(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &RolesTransactor{contract: contract}, nil
}

// NewRolesFilterer creates a new log filterer instance of Roles, bound to a specific deployed contract.
func NewRolesFilterer(address common.Address, filterer bind.ContractFilterer) (*RolesFilterer, error) {
	contract, err := bindRoles(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &RolesFilterer{contract: contract}, nil
}

// bindRoles binds a generic wrapper to an already deployed contract.
func bindRoles(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(RolesABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Roles *RolesRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Roles.Contract.RolesCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Roles *RolesRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Roles.Contract.RolesTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Roles *RolesRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Roles.Contract.RolesTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Roles *RolesCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Roles.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Roles *RolesTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Roles.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Roles *RolesTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Roles.Contract.contract.Transact(opts, method, params...)
}

// SafeMathABI is the input ABI used to generate the binding from.
const SafeMathABI = "[]"

// SafeMathBin is the compiled bytecode used for deploying new contracts.
const SafeMathBin = `0x604c6025600b82828239805160001a6073141515601857fe5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea165627a7a72305820063ea101e6c0f9188778d13299307409b207d0ef28e28e97ada05c1abf880a7b0029`

// DeploySafeMath deploys a new Ethereum contract, binding an instance of SafeMath to it.
func DeploySafeMath(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *SafeMath, error) {
	parsed, err := abi.JSON(strings.NewReader(SafeMathABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(SafeMathBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &SafeMath{SafeMathCaller: SafeMathCaller{contract: contract}, SafeMathTransactor: SafeMathTransactor{contract: contract}, SafeMathFilterer: SafeMathFilterer{contract: contract}}, nil
}

// SafeMath is an auto generated Go binding around an Ethereum contract.
type SafeMath struct {
	SafeMathCaller     // Read-only binding to the contract
	SafeMathTransactor // Write-only binding to the contract
	SafeMathFilterer   // Log filterer for contract events
}

// SafeMathCaller is an auto generated read-only Go binding around an Ethereum contract.
type SafeMathCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SafeMathTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SafeMathTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SafeMathFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SafeMathFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SafeMathSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SafeMathSession struct {
	Contract     *SafeMath         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SafeMathCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SafeMathCallerSession struct {
	Contract *SafeMathCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// SafeMathTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SafeMathTransactorSession struct {
	Contract     *SafeMathTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// SafeMathRaw is an auto generated low-level Go binding around an Ethereum contract.
type SafeMathRaw struct {
	Contract *SafeMath // Generic contract binding to access the raw methods on
}

// SafeMathCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SafeMathCallerRaw struct {
	Contract *SafeMathCaller // Generic read-only contract binding to access the raw methods on
}

// SafeMathTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SafeMathTransactorRaw struct {
	Contract *SafeMathTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSafeMath creates a new instance of SafeMath, bound to a specific deployed contract.
func NewSafeMath(address common.Address, backend bind.ContractBackend) (*SafeMath, error) {
	contract, err := bindSafeMath(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SafeMath{SafeMathCaller: SafeMathCaller{contract: contract}, SafeMathTransactor: SafeMathTransactor{contract: contract}, SafeMathFilterer: SafeMathFilterer{contract: contract}}, nil
}

// NewSafeMathCaller creates a new read-only instance of SafeMath, bound to a specific deployed contract.
func NewSafeMathCaller(address common.Address, caller bind.ContractCaller) (*SafeMathCaller, error) {
	contract, err := bindSafeMath(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SafeMathCaller{contract: contract}, nil
}

// NewSafeMathTransactor creates a new write-only instance of SafeMath, bound to a specific deployed contract.
func NewSafeMathTransactor(address common.Address, transactor bind.ContractTransactor) (*SafeMathTransactor, error) {
	contract, err := bindSafeMath(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SafeMathTransactor{contract: contract}, nil
}

// NewSafeMathFilterer creates a new log filterer instance of SafeMath, bound to a specific deployed contract.
func NewSafeMathFilterer(address common.Address, filterer bind.ContractFilterer) (*SafeMathFilterer, error) {
	contract, err := bindSafeMath(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SafeMathFilterer{contract: contract}, nil
}

// bindSafeMath binds a generic wrapper to an already deployed contract.
func bindSafeMath(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(SafeMathABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SafeMath *SafeMathRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _SafeMath.Contract.SafeMathCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SafeMath *SafeMathRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SafeMath.Contract.SafeMathTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SafeMath *SafeMathRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SafeMath.Contract.SafeMathTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SafeMath *SafeMathCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _SafeMath.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SafeMath *SafeMathTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SafeMath.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SafeMath *SafeMathTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SafeMath.Contract.contract.Transact(opts, method, params...)
}

// VeloTokenABI is the input ABI used to generate the binding from.
const VeloTokenABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"spender\",\"type\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"issuer\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"sender\",\"type\":\"address\"},{\"name\":\"recipient\",\"type\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"name\":\"\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"spender\",\"type\":\"address\"},{\"name\":\"addedValue\",\"type\":\"uint256\"}],\"name\":\"increaseAllowance\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"account\",\"type\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"mint\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"renounceWhitelistAdmin\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"account\",\"type\":\"address\"}],\"name\":\"addWhitelistAdmin\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"account\",\"type\":\"address\"},{\"name\":\"tokens\",\"type\":\"uint256\"}],\"name\":\"burn\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"spender\",\"type\":\"address\"},{\"name\":\"subtractedValue\",\"type\":\"uint256\"}],\"name\":\"decreaseAllowance\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"recipient\",\"type\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"account\",\"type\":\"address\"}],\"name\":\"isWhitelistAdmin\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"owner\",\"type\":\"address\"},{\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"issuer\",\"type\":\"string\"},{\"name\":\"name\",\"type\":\"string\"},{\"name\":\"code\",\"type\":\"string\"},{\"name\":\"decimals\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"account\",\"type\":\"address\"}],\"name\":\"WhitelistAdminAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"account\",\"type\":\"address\"}],\"name\":\"WhitelistAdminRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"}]"

// VeloTokenBin is the compiled bytecode used for deploying new contracts.
const VeloTokenBin = `0x60806040523480156200001157600080fd5b50604051620013a7380380620013a7833981018060405260808110156200003757600080fd5b8101908080516401000000008111156200005057600080fd5b820160208101848111156200006457600080fd5b81516401000000008111828201871017156200007f57600080fd5b505092919060200180516401000000008111156200009c57600080fd5b82016020810184811115620000b057600080fd5b8151640100000000811182820187101715620000cb57600080fd5b50509291906020018051640100000000811115620000e857600080fd5b82016020810184811115620000fc57600080fd5b81516401000000008111828201871017156200011757600080fd5b5050602091820151855191945092508491849184916200013d9160039186019062000318565b5081516200015390600490602085019062000318565b506005805460ff191660ff92909216919091179055506200017d9050336200019d602090811b901c565b83516200019290600790602087019062000318565b5050505050620003bd565b620001b8816006620001ef60201b62000dd61790919060201c565b6040516001600160a01b038216907f22380c05984257a1cb900161c713dd71d39e74820f1aea43bd3f1bdd2096129990600090a250565b6200020182826200029360201b60201c565b156200026e57604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601f60248201527f526f6c65733a206163636f756e7420616c72656164792068617320726f6c6500604482015290519081900360640190fd5b6001600160a01b0316600090815260209190915260409020805460ff19166001179055565b60006001600160a01b0382161515620002f8576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401808060200182810382526022815260200180620013856022913960400191505060405180910390fd5b506001600160a01b03166000908152602091909152604090205460ff1690565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f106200035b57805160ff19168380011785556200038b565b828001600101855582156200038b579182015b828111156200038b5782518255916020019190600101906200036e565b50620003999291506200039d565b5090565b620003ba91905b80821115620003995760008155600101620003a4565b90565b610fb880620003cd6000396000f3fe608060405234801561001057600080fd5b506004361061010b5760003560e01c80634c5a628c116100a25780639dc29fac116100715780639dc29fac146102f9578063a457c2d714610325578063a9059cbb14610351578063bb5f747b1461037d578063dd62ed3e146103a35761010b565b80634c5a628c1461029b57806370a08231146102a55780637362d9c8146102cb57806395d89b41146102f15761010b565b806323b872dd116100de57806323b872dd146101ef578063313ce56714610225578063395093511461024357806340c10f191461026f5761010b565b806306fdde0314610110578063095ea7b31461018d57806318160ddd146101cd5780631d143848146101e7575b600080fd5b6101186103d1565b6040805160208082528351818301528351919283929083019185019080838360005b8381101561015257818101518382015260200161013a565b50505050905090810190601f16801561017f5780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b6101b9600480360360408110156101a357600080fd5b506001600160a01b038135169060200135610467565b604080519115158252519081900360200190f35b6101d561047d565b60408051918252519081900360200190f35b610118610483565b6101b96004803603606081101561020557600080fd5b506001600160a01b038135811691602081013590911690604001356104e4565b61022d61053b565b6040805160ff9092168252519081900360200190f35b6101b96004803603604081101561025957600080fd5b506001600160a01b038135169060200135610544565b6101b96004803603604081101561028557600080fd5b506001600160a01b038135169060200135610580565b6102a36105d5565b005b6101d5600480360360208110156102bb57600080fd5b50356001600160a01b03166105e0565b6102a3600480360360208110156102e157600080fd5b50356001600160a01b03166105fb565b610118610650565b6102a36004803603604081101561030f57600080fd5b506001600160a01b0381351690602001356106b1565b6101b96004803603604081101561033b57600080fd5b506001600160a01b038135169060200135610708565b6101b96004803603604081101561036757600080fd5b506001600160a01b038135169060200135610744565b6101b96004803603602081101561039357600080fd5b50356001600160a01b0316610751565b6101d5600480360360408110156103b957600080fd5b506001600160a01b038135811691602001351661076a565b60038054604080516020601f600260001961010060018816150201909516949094049384018190048102820181019092528281526060939092909183018282801561045d5780601f106104325761010080835404028352916020019161045d565b820191906000526020600020905b81548152906001019060200180831161044057829003601f168201915b5050505050905090565b6000610474338484610795565b50600192915050565b60025490565b60078054604080516020601f600260001961010060018816150201909516949094049384018190048102820181019092528281526060939092909183018282801561045d5780601f106104325761010080835404028352916020019161045d565b60006104f184848461088b565b6001600160a01b03841660009081526001602090815260408083203380855292529091205461053191869161052c908663ffffffff6109d716565b610795565b5060019392505050565b60055460ff1690565b3360008181526001602090815260408083206001600160a01b0387168452909152812054909161047491859061052c908663ffffffff610a3716565b600061058b33610751565b15156105cb57604051600160e51b62461bcd028152600401808060200182810382526040815260200180610ee36040913960400191505060405180910390fd5b6104748383610a9b565b6105de33610b90565b565b6001600160a01b031660009081526020819052604090205490565b61060433610751565b151561064457604051600160e51b62461bcd028152600401808060200182810382526040815260200180610ee36040913960400191505060405180910390fd5b61064d81610bd8565b50565b60048054604080516020601f600260001961010060018816150201909516949094049384018190048102820181019092528281526060939092909183018282801561045d5780601f106104325761010080835404028352916020019161045d565b6106ba33610751565b15156106fa57604051600160e51b62461bcd028152600401808060200182810382526040815260200180610ee36040913960400191505060405180910390fd5b6107048282610c20565b5050565b3360008181526001602090815260408083206001600160a01b0387168452909152812054909161047491859061052c908663ffffffff6109d716565b600061047433848461088b565b600061076460068363ffffffff610cfe16565b92915050565b6001600160a01b03918216600090815260016020908152604080832093909416825291909152205490565b6001600160a01b03831615156107df57604051600160e51b62461bcd028152600401808060200182810382526024815260200180610f696024913960400191505060405180910390fd5b6001600160a01b038216151561082957604051600160e51b62461bcd028152600401808060200182810382526022815260200180610e7e6022913960400191505060405180910390fd5b6001600160a01b03808416600081815260016020908152604080832094871680845294825291829020859055815185815291517f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b9259281900390910190a3505050565b6001600160a01b03831615156108d557604051600160e51b62461bcd028152600401808060200182810382526025815260200180610f446025913960400191505060405180910390fd5b6001600160a01b038216151561091f57604051600160e51b62461bcd028152600401808060200182810382526023815260200180610e5b6023913960400191505060405180910390fd5b6001600160a01b038316600090815260208190526040902054610948908263ffffffff6109d716565b6001600160a01b03808516600090815260208190526040808220939093559084168152205461097d908263ffffffff610a3716565b6001600160a01b038084166000818152602081815260409182902094909455805185815290519193928716927fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef92918290030190a3505050565b600082821115610a315760408051600160e51b62461bcd02815260206004820152601e60248201527f536166654d6174683a207375627472616374696f6e206f766572666c6f770000604482015290519081900360640190fd5b50900390565b600082820183811015610a945760408051600160e51b62461bcd02815260206004820152601b60248201527f536166654d6174683a206164646974696f6e206f766572666c6f770000000000604482015290519081900360640190fd5b9392505050565b6001600160a01b0382161515610afb5760408051600160e51b62461bcd02815260206004820152601f60248201527f45524332303a206d696e7420746f20746865207a65726f206164647265737300604482015290519081900360640190fd5b600254610b0e908263ffffffff610a3716565b6002556001600160a01b038216600090815260208190526040902054610b3a908263ffffffff610a3716565b6001600160a01b0383166000818152602081815260408083209490945583518581529351929391927fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef9281900390910190a35050565b610ba160068263ffffffff610d6a16565b6040516001600160a01b038216907f0a8eb35e5ca14b3d6f28e4abf2f128dbab231a58b56e89beb5d636115001e16590600090a250565b610be960068263ffffffff610dd616565b6040516001600160a01b038216907f22380c05984257a1cb900161c713dd71d39e74820f1aea43bd3f1bdd2096129990600090a250565b6001600160a01b0382161515610c6a57604051600160e51b62461bcd028152600401808060200182810382526021815260200180610f236021913960400191505060405180910390fd5b600254610c7d908263ffffffff6109d716565b6002556001600160a01b038216600090815260208190526040902054610ca9908263ffffffff6109d716565b6001600160a01b038316600081815260208181526040808320949094558351858152935191937fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef929081900390910190a35050565b60006001600160a01b0382161515610d4a57604051600160e51b62461bcd028152600401808060200182810382526022815260200180610ec16022913960400191505060405180910390fd5b506001600160a01b03166000908152602091909152604090205460ff1690565b610d748282610cfe565b1515610db457604051600160e51b62461bcd028152600401808060200182810382526021815260200180610ea06021913960400191505060405180910390fd5b6001600160a01b0316600090815260209190915260409020805460ff19169055565b610de08282610cfe565b15610e355760408051600160e51b62461bcd02815260206004820152601f60248201527f526f6c65733a206163636f756e7420616c72656164792068617320726f6c6500604482015290519081900360640190fd5b6001600160a01b0316600090815260209190915260409020805460ff1916600117905556fe45524332303a207472616e7366657220746f20746865207a65726f206164647265737345524332303a20617070726f766520746f20746865207a65726f2061646472657373526f6c65733a206163636f756e7420646f6573206e6f74206861766520726f6c65526f6c65733a206163636f756e7420697320746865207a65726f206164647265737357686974656c69737441646d696e526f6c653a2063616c6c657220646f6573206e6f742068617665207468652057686974656c69737441646d696e20726f6c6545524332303a206275726e2066726f6d20746865207a65726f206164647265737345524332303a207472616e736665722066726f6d20746865207a65726f206164647265737345524332303a20617070726f76652066726f6d20746865207a65726f2061646472657373a165627a7a72305820bad9617bb73db856345bc85523f0f32c47777af0347ae15dc20c7cc0d1d4ebdc0029526f6c65733a206163636f756e7420697320746865207a65726f2061646472657373`

// DeployVeloToken deploys a new Ethereum contract, binding an instance of VeloToken to it.
func DeployVeloToken(auth *bind.TransactOpts, backend bind.ContractBackend, issuer string, name string, code string, decimals uint8) (common.Address, *types.Transaction, *VeloToken, error) {
	parsed, err := abi.JSON(strings.NewReader(VeloTokenABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(VeloTokenBin), backend, issuer, name, code, decimals)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &VeloToken{VeloTokenCaller: VeloTokenCaller{contract: contract}, VeloTokenTransactor: VeloTokenTransactor{contract: contract}, VeloTokenFilterer: VeloTokenFilterer{contract: contract}}, nil
}

// VeloToken is an auto generated Go binding around an Ethereum contract.
type VeloToken struct {
	VeloTokenCaller     // Read-only binding to the contract
	VeloTokenTransactor // Write-only binding to the contract
	VeloTokenFilterer   // Log filterer for contract events
}

// VeloTokenCaller is an auto generated read-only Go binding around an Ethereum contract.
type VeloTokenCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VeloTokenTransactor is an auto generated write-only Go binding around an Ethereum contract.
type VeloTokenTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VeloTokenFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type VeloTokenFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VeloTokenSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type VeloTokenSession struct {
	Contract     *VeloToken        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// VeloTokenCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type VeloTokenCallerSession struct {
	Contract *VeloTokenCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// VeloTokenTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type VeloTokenTransactorSession struct {
	Contract     *VeloTokenTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// VeloTokenRaw is an auto generated low-level Go binding around an Ethereum contract.
type VeloTokenRaw struct {
	Contract *VeloToken // Generic contract binding to access the raw methods on
}

// VeloTokenCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type VeloTokenCallerRaw struct {
	Contract *VeloTokenCaller // Generic read-only contract binding to access the raw methods on
}

// VeloTokenTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type VeloTokenTransactorRaw struct {
	Contract *VeloTokenTransactor // Generic write-only contract binding to access the raw methods on
}

// NewVeloToken creates a new instance of VeloToken, bound to a specific deployed contract.
func NewVeloToken(address common.Address, backend bind.ContractBackend) (*VeloToken, error) {
	contract, err := bindVeloToken(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &VeloToken{VeloTokenCaller: VeloTokenCaller{contract: contract}, VeloTokenTransactor: VeloTokenTransactor{contract: contract}, VeloTokenFilterer: VeloTokenFilterer{contract: contract}}, nil
}

// NewVeloTokenCaller creates a new read-only instance of VeloToken, bound to a specific deployed contract.
func NewVeloTokenCaller(address common.Address, caller bind.ContractCaller) (*VeloTokenCaller, error) {
	contract, err := bindVeloToken(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &VeloTokenCaller{contract: contract}, nil
}

// NewVeloTokenTransactor creates a new write-only instance of VeloToken, bound to a specific deployed contract.
func NewVeloTokenTransactor(address common.Address, transactor bind.ContractTransactor) (*VeloTokenTransactor, error) {
	contract, err := bindVeloToken(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &VeloTokenTransactor{contract: contract}, nil
}

// NewVeloTokenFilterer creates a new log filterer instance of VeloToken, bound to a specific deployed contract.
func NewVeloTokenFilterer(address common.Address, filterer bind.ContractFilterer) (*VeloTokenFilterer, error) {
	contract, err := bindVeloToken(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &VeloTokenFilterer{contract: contract}, nil
}

// bindVeloToken binds a generic wrapper to an already deployed contract.
func bindVeloToken(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(VeloTokenABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_VeloToken *VeloTokenRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _VeloToken.Contract.VeloTokenCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_VeloToken *VeloTokenRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VeloToken.Contract.VeloTokenTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_VeloToken *VeloTokenRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _VeloToken.Contract.VeloTokenTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_VeloToken *VeloTokenCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _VeloToken.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_VeloToken *VeloTokenTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VeloToken.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_VeloToken *VeloTokenTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _VeloToken.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) constant returns(uint256)
func (_VeloToken *VeloTokenCaller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _VeloToken.contract.Call(opts, out, "allowance", owner, spender)
	return *ret0, err
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) constant returns(uint256)
func (_VeloToken *VeloTokenSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _VeloToken.Contract.Allowance(&_VeloToken.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) constant returns(uint256)
func (_VeloToken *VeloTokenCallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _VeloToken.Contract.Allowance(&_VeloToken.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) constant returns(uint256)
func (_VeloToken *VeloTokenCaller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _VeloToken.contract.Call(opts, out, "balanceOf", account)
	return *ret0, err
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) constant returns(uint256)
func (_VeloToken *VeloTokenSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _VeloToken.Contract.BalanceOf(&_VeloToken.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) constant returns(uint256)
func (_VeloToken *VeloTokenCallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _VeloToken.Contract.BalanceOf(&_VeloToken.CallOpts, account)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() constant returns(uint8)
func (_VeloToken *VeloTokenCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var (
		ret0 = new(uint8)
	)
	out := ret0
	err := _VeloToken.contract.Call(opts, out, "decimals")
	return *ret0, err
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() constant returns(uint8)
func (_VeloToken *VeloTokenSession) Decimals() (uint8, error) {
	return _VeloToken.Contract.Decimals(&_VeloToken.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() constant returns(uint8)
func (_VeloToken *VeloTokenCallerSession) Decimals() (uint8, error) {
	return _VeloToken.Contract.Decimals(&_VeloToken.CallOpts)
}

// IsWhitelistAdmin is a free data retrieval call binding the contract method 0xbb5f747b.
//
// Solidity: function isWhitelistAdmin(address account) constant returns(bool)
func (_VeloToken *VeloTokenCaller) IsWhitelistAdmin(opts *bind.CallOpts, account common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _VeloToken.contract.Call(opts, out, "isWhitelistAdmin", account)
	return *ret0, err
}

// IsWhitelistAdmin is a free data retrieval call binding the contract method 0xbb5f747b.
//
// Solidity: function isWhitelistAdmin(address account) constant returns(bool)
func (_VeloToken *VeloTokenSession) IsWhitelistAdmin(account common.Address) (bool, error) {
	return _VeloToken.Contract.IsWhitelistAdmin(&_VeloToken.CallOpts, account)
}

// IsWhitelistAdmin is a free data retrieval call binding the contract method 0xbb5f747b.
//
// Solidity: function isWhitelistAdmin(address account) constant returns(bool)
func (_VeloToken *VeloTokenCallerSession) IsWhitelistAdmin(account common.Address) (bool, error) {
	return _VeloToken.Contract.IsWhitelistAdmin(&_VeloToken.CallOpts, account)
}

// Issuer is a free data retrieval call binding the contract method 0x1d143848.
//
// Solidity: function issuer() constant returns(string)
func (_VeloToken *VeloTokenCaller) Issuer(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _VeloToken.contract.Call(opts, out, "issuer")
	return *ret0, err
}

// Issuer is a free data retrieval call binding the contract method 0x1d143848.
//
// Solidity: function issuer() constant returns(string)
func (_VeloToken *VeloTokenSession) Issuer() (string, error) {
	return _VeloToken.Contract.Issuer(&_VeloToken.CallOpts)
}

// Issuer is a free data retrieval call binding the contract method 0x1d143848.
//
// Solidity: function issuer() constant returns(string)
func (_VeloToken *VeloTokenCallerSession) Issuer() (string, error) {
	return _VeloToken.Contract.Issuer(&_VeloToken.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_VeloToken *VeloTokenCaller) Name(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _VeloToken.contract.Call(opts, out, "name")
	return *ret0, err
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_VeloToken *VeloTokenSession) Name() (string, error) {
	return _VeloToken.Contract.Name(&_VeloToken.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_VeloToken *VeloTokenCallerSession) Name() (string, error) {
	return _VeloToken.Contract.Name(&_VeloToken.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() constant returns(string)
func (_VeloToken *VeloTokenCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _VeloToken.contract.Call(opts, out, "symbol")
	return *ret0, err
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() constant returns(string)
func (_VeloToken *VeloTokenSession) Symbol() (string, error) {
	return _VeloToken.Contract.Symbol(&_VeloToken.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() constant returns(string)
func (_VeloToken *VeloTokenCallerSession) Symbol() (string, error) {
	return _VeloToken.Contract.Symbol(&_VeloToken.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_VeloToken *VeloTokenCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _VeloToken.contract.Call(opts, out, "totalSupply")
	return *ret0, err
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_VeloToken *VeloTokenSession) TotalSupply() (*big.Int, error) {
	return _VeloToken.Contract.TotalSupply(&_VeloToken.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_VeloToken *VeloTokenCallerSession) TotalSupply() (*big.Int, error) {
	return _VeloToken.Contract.TotalSupply(&_VeloToken.CallOpts)
}

// AddWhitelistAdmin is a paid mutator transaction binding the contract method 0x7362d9c8.
//
// Solidity: function addWhitelistAdmin(address account) returns()
func (_VeloToken *VeloTokenTransactor) AddWhitelistAdmin(opts *bind.TransactOpts, account common.Address) (*types.Transaction, error) {
	return _VeloToken.contract.Transact(opts, "addWhitelistAdmin", account)
}

// AddWhitelistAdmin is a paid mutator transaction binding the contract method 0x7362d9c8.
//
// Solidity: function addWhitelistAdmin(address account) returns()
func (_VeloToken *VeloTokenSession) AddWhitelistAdmin(account common.Address) (*types.Transaction, error) {
	return _VeloToken.Contract.AddWhitelistAdmin(&_VeloToken.TransactOpts, account)
}

// AddWhitelistAdmin is a paid mutator transaction binding the contract method 0x7362d9c8.
//
// Solidity: function addWhitelistAdmin(address account) returns()
func (_VeloToken *VeloTokenTransactorSession) AddWhitelistAdmin(account common.Address) (*types.Transaction, error) {
	return _VeloToken.Contract.AddWhitelistAdmin(&_VeloToken.TransactOpts, account)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_VeloToken *VeloTokenTransactor) Approve(opts *bind.TransactOpts, spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _VeloToken.contract.Transact(opts, "approve", spender, value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_VeloToken *VeloTokenSession) Approve(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _VeloToken.Contract.Approve(&_VeloToken.TransactOpts, spender, value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_VeloToken *VeloTokenTransactorSession) Approve(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _VeloToken.Contract.Approve(&_VeloToken.TransactOpts, spender, value)
}

// Burn is a paid mutator transaction binding the contract method 0x9dc29fac.
//
// Solidity: function burn(address account, uint256 tokens) returns()
func (_VeloToken *VeloTokenTransactor) Burn(opts *bind.TransactOpts, account common.Address, tokens *big.Int) (*types.Transaction, error) {
	return _VeloToken.contract.Transact(opts, "burn", account, tokens)
}

// Burn is a paid mutator transaction binding the contract method 0x9dc29fac.
//
// Solidity: function burn(address account, uint256 tokens) returns()
func (_VeloToken *VeloTokenSession) Burn(account common.Address, tokens *big.Int) (*types.Transaction, error) {
	return _VeloToken.Contract.Burn(&_VeloToken.TransactOpts, account, tokens)
}

// Burn is a paid mutator transaction binding the contract method 0x9dc29fac.
//
// Solidity: function burn(address account, uint256 tokens) returns()
func (_VeloToken *VeloTokenTransactorSession) Burn(account common.Address, tokens *big.Int) (*types.Transaction, error) {
	return _VeloToken.Contract.Burn(&_VeloToken.TransactOpts, account, tokens)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_VeloToken *VeloTokenTransactor) DecreaseAllowance(opts *bind.TransactOpts, spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _VeloToken.contract.Transact(opts, "decreaseAllowance", spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_VeloToken *VeloTokenSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _VeloToken.Contract.DecreaseAllowance(&_VeloToken.TransactOpts, spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_VeloToken *VeloTokenTransactorSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _VeloToken.Contract.DecreaseAllowance(&_VeloToken.TransactOpts, spender, subtractedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_VeloToken *VeloTokenTransactor) IncreaseAllowance(opts *bind.TransactOpts, spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _VeloToken.contract.Transact(opts, "increaseAllowance", spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_VeloToken *VeloTokenSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _VeloToken.Contract.IncreaseAllowance(&_VeloToken.TransactOpts, spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_VeloToken *VeloTokenTransactorSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _VeloToken.Contract.IncreaseAllowance(&_VeloToken.TransactOpts, spender, addedValue)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address account, uint256 amount) returns(bool)
func (_VeloToken *VeloTokenTransactor) Mint(opts *bind.TransactOpts, account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _VeloToken.contract.Transact(opts, "mint", account, amount)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address account, uint256 amount) returns(bool)
func (_VeloToken *VeloTokenSession) Mint(account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _VeloToken.Contract.Mint(&_VeloToken.TransactOpts, account, amount)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address account, uint256 amount) returns(bool)
func (_VeloToken *VeloTokenTransactorSession) Mint(account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _VeloToken.Contract.Mint(&_VeloToken.TransactOpts, account, amount)
}

// RenounceWhitelistAdmin is a paid mutator transaction binding the contract method 0x4c5a628c.
//
// Solidity: function renounceWhitelistAdmin() returns()
func (_VeloToken *VeloTokenTransactor) RenounceWhitelistAdmin(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VeloToken.contract.Transact(opts, "renounceWhitelistAdmin")
}

// RenounceWhitelistAdmin is a paid mutator transaction binding the contract method 0x4c5a628c.
//
// Solidity: function renounceWhitelistAdmin() returns()
func (_VeloToken *VeloTokenSession) RenounceWhitelistAdmin() (*types.Transaction, error) {
	return _VeloToken.Contract.RenounceWhitelistAdmin(&_VeloToken.TransactOpts)
}

// RenounceWhitelistAdmin is a paid mutator transaction binding the contract method 0x4c5a628c.
//
// Solidity: function renounceWhitelistAdmin() returns()
func (_VeloToken *VeloTokenTransactorSession) RenounceWhitelistAdmin() (*types.Transaction, error) {
	return _VeloToken.Contract.RenounceWhitelistAdmin(&_VeloToken.TransactOpts)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_VeloToken *VeloTokenTransactor) Transfer(opts *bind.TransactOpts, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _VeloToken.contract.Transact(opts, "transfer", recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_VeloToken *VeloTokenSession) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _VeloToken.Contract.Transfer(&_VeloToken.TransactOpts, recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_VeloToken *VeloTokenTransactorSession) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _VeloToken.Contract.Transfer(&_VeloToken.TransactOpts, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_VeloToken *VeloTokenTransactor) TransferFrom(opts *bind.TransactOpts, sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _VeloToken.contract.Transact(opts, "transferFrom", sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_VeloToken *VeloTokenSession) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _VeloToken.Contract.TransferFrom(&_VeloToken.TransactOpts, sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_VeloToken *VeloTokenTransactorSession) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _VeloToken.Contract.TransferFrom(&_VeloToken.TransactOpts, sender, recipient, amount)
}

// VeloTokenApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the VeloToken contract.
type VeloTokenApprovalIterator struct {
	Event *VeloTokenApproval // Event containing the contract specifics and raw log

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
func (it *VeloTokenApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VeloTokenApproval)
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
		it.Event = new(VeloTokenApproval)
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
func (it *VeloTokenApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VeloTokenApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VeloTokenApproval represents a Approval event raised by the VeloToken contract.
type VeloTokenApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_VeloToken *VeloTokenFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*VeloTokenApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _VeloToken.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &VeloTokenApprovalIterator{contract: _VeloToken.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_VeloToken *VeloTokenFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *VeloTokenApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _VeloToken.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VeloTokenApproval)
				if err := _VeloToken.contract.UnpackLog(event, "Approval", log); err != nil {
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

// VeloTokenTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the VeloToken contract.
type VeloTokenTransferIterator struct {
	Event *VeloTokenTransfer // Event containing the contract specifics and raw log

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
func (it *VeloTokenTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VeloTokenTransfer)
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
		it.Event = new(VeloTokenTransfer)
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
func (it *VeloTokenTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VeloTokenTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VeloTokenTransfer represents a Transfer event raised by the VeloToken contract.
type VeloTokenTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_VeloToken *VeloTokenFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*VeloTokenTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _VeloToken.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &VeloTokenTransferIterator{contract: _VeloToken.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_VeloToken *VeloTokenFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *VeloTokenTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _VeloToken.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VeloTokenTransfer)
				if err := _VeloToken.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// VeloTokenWhitelistAdminAddedIterator is returned from FilterWhitelistAdminAdded and is used to iterate over the raw logs and unpacked data for WhitelistAdminAdded events raised by the VeloToken contract.
type VeloTokenWhitelistAdminAddedIterator struct {
	Event *VeloTokenWhitelistAdminAdded // Event containing the contract specifics and raw log

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
func (it *VeloTokenWhitelistAdminAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VeloTokenWhitelistAdminAdded)
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
		it.Event = new(VeloTokenWhitelistAdminAdded)
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
func (it *VeloTokenWhitelistAdminAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VeloTokenWhitelistAdminAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VeloTokenWhitelistAdminAdded represents a WhitelistAdminAdded event raised by the VeloToken contract.
type VeloTokenWhitelistAdminAdded struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterWhitelistAdminAdded is a free log retrieval operation binding the contract event 0x22380c05984257a1cb900161c713dd71d39e74820f1aea43bd3f1bdd20961299.
//
// Solidity: event WhitelistAdminAdded(address indexed account)
func (_VeloToken *VeloTokenFilterer) FilterWhitelistAdminAdded(opts *bind.FilterOpts, account []common.Address) (*VeloTokenWhitelistAdminAddedIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _VeloToken.contract.FilterLogs(opts, "WhitelistAdminAdded", accountRule)
	if err != nil {
		return nil, err
	}
	return &VeloTokenWhitelistAdminAddedIterator{contract: _VeloToken.contract, event: "WhitelistAdminAdded", logs: logs, sub: sub}, nil
}

// WatchWhitelistAdminAdded is a free log subscription operation binding the contract event 0x22380c05984257a1cb900161c713dd71d39e74820f1aea43bd3f1bdd20961299.
//
// Solidity: event WhitelistAdminAdded(address indexed account)
func (_VeloToken *VeloTokenFilterer) WatchWhitelistAdminAdded(opts *bind.WatchOpts, sink chan<- *VeloTokenWhitelistAdminAdded, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _VeloToken.contract.WatchLogs(opts, "WhitelistAdminAdded", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VeloTokenWhitelistAdminAdded)
				if err := _VeloToken.contract.UnpackLog(event, "WhitelistAdminAdded", log); err != nil {
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

// VeloTokenWhitelistAdminRemovedIterator is returned from FilterWhitelistAdminRemoved and is used to iterate over the raw logs and unpacked data for WhitelistAdminRemoved events raised by the VeloToken contract.
type VeloTokenWhitelistAdminRemovedIterator struct {
	Event *VeloTokenWhitelistAdminRemoved // Event containing the contract specifics and raw log

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
func (it *VeloTokenWhitelistAdminRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VeloTokenWhitelistAdminRemoved)
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
		it.Event = new(VeloTokenWhitelistAdminRemoved)
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
func (it *VeloTokenWhitelistAdminRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VeloTokenWhitelistAdminRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VeloTokenWhitelistAdminRemoved represents a WhitelistAdminRemoved event raised by the VeloToken contract.
type VeloTokenWhitelistAdminRemoved struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterWhitelistAdminRemoved is a free log retrieval operation binding the contract event 0x0a8eb35e5ca14b3d6f28e4abf2f128dbab231a58b56e89beb5d636115001e165.
//
// Solidity: event WhitelistAdminRemoved(address indexed account)
func (_VeloToken *VeloTokenFilterer) FilterWhitelistAdminRemoved(opts *bind.FilterOpts, account []common.Address) (*VeloTokenWhitelistAdminRemovedIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _VeloToken.contract.FilterLogs(opts, "WhitelistAdminRemoved", accountRule)
	if err != nil {
		return nil, err
	}
	return &VeloTokenWhitelistAdminRemovedIterator{contract: _VeloToken.contract, event: "WhitelistAdminRemoved", logs: logs, sub: sub}, nil
}

// WatchWhitelistAdminRemoved is a free log subscription operation binding the contract event 0x0a8eb35e5ca14b3d6f28e4abf2f128dbab231a58b56e89beb5d636115001e165.
//
// Solidity: event WhitelistAdminRemoved(address indexed account)
func (_VeloToken *VeloTokenFilterer) WatchWhitelistAdminRemoved(opts *bind.WatchOpts, sink chan<- *VeloTokenWhitelistAdminRemoved, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _VeloToken.contract.WatchLogs(opts, "WhitelistAdminRemoved", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VeloTokenWhitelistAdminRemoved)
				if err := _VeloToken.contract.UnpackLog(event, "WhitelistAdminRemoved", log); err != nil {
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

// WhitelistAdminRoleABI is the input ABI used to generate the binding from.
const WhitelistAdminRoleABI = "[{\"constant\":false,\"inputs\":[],\"name\":\"renounceWhitelistAdmin\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"account\",\"type\":\"address\"}],\"name\":\"addWhitelistAdmin\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"account\",\"type\":\"address\"}],\"name\":\"isWhitelistAdmin\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"account\",\"type\":\"address\"}],\"name\":\"WhitelistAdminAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"account\",\"type\":\"address\"}],\"name\":\"WhitelistAdminRemoved\",\"type\":\"event\"}]"

// WhitelistAdminRoleBin is the compiled bytecode used for deploying new contracts.
const WhitelistAdminRoleBin = `0x`

// DeployWhitelistAdminRole deploys a new Ethereum contract, binding an instance of WhitelistAdminRole to it.
func DeployWhitelistAdminRole(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *WhitelistAdminRole, error) {
	parsed, err := abi.JSON(strings.NewReader(WhitelistAdminRoleABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(WhitelistAdminRoleBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &WhitelistAdminRole{WhitelistAdminRoleCaller: WhitelistAdminRoleCaller{contract: contract}, WhitelistAdminRoleTransactor: WhitelistAdminRoleTransactor{contract: contract}, WhitelistAdminRoleFilterer: WhitelistAdminRoleFilterer{contract: contract}}, nil
}

// WhitelistAdminRole is an auto generated Go binding around an Ethereum contract.
type WhitelistAdminRole struct {
	WhitelistAdminRoleCaller     // Read-only binding to the contract
	WhitelistAdminRoleTransactor // Write-only binding to the contract
	WhitelistAdminRoleFilterer   // Log filterer for contract events
}

// WhitelistAdminRoleCaller is an auto generated read-only Go binding around an Ethereum contract.
type WhitelistAdminRoleCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WhitelistAdminRoleTransactor is an auto generated write-only Go binding around an Ethereum contract.
type WhitelistAdminRoleTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WhitelistAdminRoleFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type WhitelistAdminRoleFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WhitelistAdminRoleSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type WhitelistAdminRoleSession struct {
	Contract     *WhitelistAdminRole // Generic contract binding to set the session for
	CallOpts     bind.CallOpts       // Call options to use throughout this session
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// WhitelistAdminRoleCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type WhitelistAdminRoleCallerSession struct {
	Contract *WhitelistAdminRoleCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts             // Call options to use throughout this session
}

// WhitelistAdminRoleTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type WhitelistAdminRoleTransactorSession struct {
	Contract     *WhitelistAdminRoleTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts             // Transaction auth options to use throughout this session
}

// WhitelistAdminRoleRaw is an auto generated low-level Go binding around an Ethereum contract.
type WhitelistAdminRoleRaw struct {
	Contract *WhitelistAdminRole // Generic contract binding to access the raw methods on
}

// WhitelistAdminRoleCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type WhitelistAdminRoleCallerRaw struct {
	Contract *WhitelistAdminRoleCaller // Generic read-only contract binding to access the raw methods on
}

// WhitelistAdminRoleTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type WhitelistAdminRoleTransactorRaw struct {
	Contract *WhitelistAdminRoleTransactor // Generic write-only contract binding to access the raw methods on
}

// NewWhitelistAdminRole creates a new instance of WhitelistAdminRole, bound to a specific deployed contract.
func NewWhitelistAdminRole(address common.Address, backend bind.ContractBackend) (*WhitelistAdminRole, error) {
	contract, err := bindWhitelistAdminRole(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &WhitelistAdminRole{WhitelistAdminRoleCaller: WhitelistAdminRoleCaller{contract: contract}, WhitelistAdminRoleTransactor: WhitelistAdminRoleTransactor{contract: contract}, WhitelistAdminRoleFilterer: WhitelistAdminRoleFilterer{contract: contract}}, nil
}

// NewWhitelistAdminRoleCaller creates a new read-only instance of WhitelistAdminRole, bound to a specific deployed contract.
func NewWhitelistAdminRoleCaller(address common.Address, caller bind.ContractCaller) (*WhitelistAdminRoleCaller, error) {
	contract, err := bindWhitelistAdminRole(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &WhitelistAdminRoleCaller{contract: contract}, nil
}

// NewWhitelistAdminRoleTransactor creates a new write-only instance of WhitelistAdminRole, bound to a specific deployed contract.
func NewWhitelistAdminRoleTransactor(address common.Address, transactor bind.ContractTransactor) (*WhitelistAdminRoleTransactor, error) {
	contract, err := bindWhitelistAdminRole(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &WhitelistAdminRoleTransactor{contract: contract}, nil
}

// NewWhitelistAdminRoleFilterer creates a new log filterer instance of WhitelistAdminRole, bound to a specific deployed contract.
func NewWhitelistAdminRoleFilterer(address common.Address, filterer bind.ContractFilterer) (*WhitelistAdminRoleFilterer, error) {
	contract, err := bindWhitelistAdminRole(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &WhitelistAdminRoleFilterer{contract: contract}, nil
}

// bindWhitelistAdminRole binds a generic wrapper to an already deployed contract.
func bindWhitelistAdminRole(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(WhitelistAdminRoleABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_WhitelistAdminRole *WhitelistAdminRoleRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _WhitelistAdminRole.Contract.WhitelistAdminRoleCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_WhitelistAdminRole *WhitelistAdminRoleRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _WhitelistAdminRole.Contract.WhitelistAdminRoleTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_WhitelistAdminRole *WhitelistAdminRoleRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _WhitelistAdminRole.Contract.WhitelistAdminRoleTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_WhitelistAdminRole *WhitelistAdminRoleCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _WhitelistAdminRole.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_WhitelistAdminRole *WhitelistAdminRoleTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _WhitelistAdminRole.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_WhitelistAdminRole *WhitelistAdminRoleTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _WhitelistAdminRole.Contract.contract.Transact(opts, method, params...)
}

// IsWhitelistAdmin is a free data retrieval call binding the contract method 0xbb5f747b.
//
// Solidity: function isWhitelistAdmin(address account) constant returns(bool)
func (_WhitelistAdminRole *WhitelistAdminRoleCaller) IsWhitelistAdmin(opts *bind.CallOpts, account common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _WhitelistAdminRole.contract.Call(opts, out, "isWhitelistAdmin", account)
	return *ret0, err
}

// IsWhitelistAdmin is a free data retrieval call binding the contract method 0xbb5f747b.
//
// Solidity: function isWhitelistAdmin(address account) constant returns(bool)
func (_WhitelistAdminRole *WhitelistAdminRoleSession) IsWhitelistAdmin(account common.Address) (bool, error) {
	return _WhitelistAdminRole.Contract.IsWhitelistAdmin(&_WhitelistAdminRole.CallOpts, account)
}

// IsWhitelistAdmin is a free data retrieval call binding the contract method 0xbb5f747b.
//
// Solidity: function isWhitelistAdmin(address account) constant returns(bool)
func (_WhitelistAdminRole *WhitelistAdminRoleCallerSession) IsWhitelistAdmin(account common.Address) (bool, error) {
	return _WhitelistAdminRole.Contract.IsWhitelistAdmin(&_WhitelistAdminRole.CallOpts, account)
}

// AddWhitelistAdmin is a paid mutator transaction binding the contract method 0x7362d9c8.
//
// Solidity: function addWhitelistAdmin(address account) returns()
func (_WhitelistAdminRole *WhitelistAdminRoleTransactor) AddWhitelistAdmin(opts *bind.TransactOpts, account common.Address) (*types.Transaction, error) {
	return _WhitelistAdminRole.contract.Transact(opts, "addWhitelistAdmin", account)
}

// AddWhitelistAdmin is a paid mutator transaction binding the contract method 0x7362d9c8.
//
// Solidity: function addWhitelistAdmin(address account) returns()
func (_WhitelistAdminRole *WhitelistAdminRoleSession) AddWhitelistAdmin(account common.Address) (*types.Transaction, error) {
	return _WhitelistAdminRole.Contract.AddWhitelistAdmin(&_WhitelistAdminRole.TransactOpts, account)
}

// AddWhitelistAdmin is a paid mutator transaction binding the contract method 0x7362d9c8.
//
// Solidity: function addWhitelistAdmin(address account) returns()
func (_WhitelistAdminRole *WhitelistAdminRoleTransactorSession) AddWhitelistAdmin(account common.Address) (*types.Transaction, error) {
	return _WhitelistAdminRole.Contract.AddWhitelistAdmin(&_WhitelistAdminRole.TransactOpts, account)
}

// RenounceWhitelistAdmin is a paid mutator transaction binding the contract method 0x4c5a628c.
//
// Solidity: function renounceWhitelistAdmin() returns()
func (_WhitelistAdminRole *WhitelistAdminRoleTransactor) RenounceWhitelistAdmin(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _WhitelistAdminRole.contract.Transact(opts, "renounceWhitelistAdmin")
}

// RenounceWhitelistAdmin is a paid mutator transaction binding the contract method 0x4c5a628c.
//
// Solidity: function renounceWhitelistAdmin() returns()
func (_WhitelistAdminRole *WhitelistAdminRoleSession) RenounceWhitelistAdmin() (*types.Transaction, error) {
	return _WhitelistAdminRole.Contract.RenounceWhitelistAdmin(&_WhitelistAdminRole.TransactOpts)
}

// RenounceWhitelistAdmin is a paid mutator transaction binding the contract method 0x4c5a628c.
//
// Solidity: function renounceWhitelistAdmin() returns()
func (_WhitelistAdminRole *WhitelistAdminRoleTransactorSession) RenounceWhitelistAdmin() (*types.Transaction, error) {
	return _WhitelistAdminRole.Contract.RenounceWhitelistAdmin(&_WhitelistAdminRole.TransactOpts)
}

// WhitelistAdminRoleWhitelistAdminAddedIterator is returned from FilterWhitelistAdminAdded and is used to iterate over the raw logs and unpacked data for WhitelistAdminAdded events raised by the WhitelistAdminRole contract.
type WhitelistAdminRoleWhitelistAdminAddedIterator struct {
	Event *WhitelistAdminRoleWhitelistAdminAdded // Event containing the contract specifics and raw log

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
func (it *WhitelistAdminRoleWhitelistAdminAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WhitelistAdminRoleWhitelistAdminAdded)
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
		it.Event = new(WhitelistAdminRoleWhitelistAdminAdded)
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
func (it *WhitelistAdminRoleWhitelistAdminAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WhitelistAdminRoleWhitelistAdminAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WhitelistAdminRoleWhitelistAdminAdded represents a WhitelistAdminAdded event raised by the WhitelistAdminRole contract.
type WhitelistAdminRoleWhitelistAdminAdded struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterWhitelistAdminAdded is a free log retrieval operation binding the contract event 0x22380c05984257a1cb900161c713dd71d39e74820f1aea43bd3f1bdd20961299.
//
// Solidity: event WhitelistAdminAdded(address indexed account)
func (_WhitelistAdminRole *WhitelistAdminRoleFilterer) FilterWhitelistAdminAdded(opts *bind.FilterOpts, account []common.Address) (*WhitelistAdminRoleWhitelistAdminAddedIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _WhitelistAdminRole.contract.FilterLogs(opts, "WhitelistAdminAdded", accountRule)
	if err != nil {
		return nil, err
	}
	return &WhitelistAdminRoleWhitelistAdminAddedIterator{contract: _WhitelistAdminRole.contract, event: "WhitelistAdminAdded", logs: logs, sub: sub}, nil
}

// WatchWhitelistAdminAdded is a free log subscription operation binding the contract event 0x22380c05984257a1cb900161c713dd71d39e74820f1aea43bd3f1bdd20961299.
//
// Solidity: event WhitelistAdminAdded(address indexed account)
func (_WhitelistAdminRole *WhitelistAdminRoleFilterer) WatchWhitelistAdminAdded(opts *bind.WatchOpts, sink chan<- *WhitelistAdminRoleWhitelistAdminAdded, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _WhitelistAdminRole.contract.WatchLogs(opts, "WhitelistAdminAdded", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WhitelistAdminRoleWhitelistAdminAdded)
				if err := _WhitelistAdminRole.contract.UnpackLog(event, "WhitelistAdminAdded", log); err != nil {
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

// WhitelistAdminRoleWhitelistAdminRemovedIterator is returned from FilterWhitelistAdminRemoved and is used to iterate over the raw logs and unpacked data for WhitelistAdminRemoved events raised by the WhitelistAdminRole contract.
type WhitelistAdminRoleWhitelistAdminRemovedIterator struct {
	Event *WhitelistAdminRoleWhitelistAdminRemoved // Event containing the contract specifics and raw log

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
func (it *WhitelistAdminRoleWhitelistAdminRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WhitelistAdminRoleWhitelistAdminRemoved)
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
		it.Event = new(WhitelistAdminRoleWhitelistAdminRemoved)
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
func (it *WhitelistAdminRoleWhitelistAdminRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WhitelistAdminRoleWhitelistAdminRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WhitelistAdminRoleWhitelistAdminRemoved represents a WhitelistAdminRemoved event raised by the WhitelistAdminRole contract.
type WhitelistAdminRoleWhitelistAdminRemoved struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterWhitelistAdminRemoved is a free log retrieval operation binding the contract event 0x0a8eb35e5ca14b3d6f28e4abf2f128dbab231a58b56e89beb5d636115001e165.
//
// Solidity: event WhitelistAdminRemoved(address indexed account)
func (_WhitelistAdminRole *WhitelistAdminRoleFilterer) FilterWhitelistAdminRemoved(opts *bind.FilterOpts, account []common.Address) (*WhitelistAdminRoleWhitelistAdminRemovedIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _WhitelistAdminRole.contract.FilterLogs(opts, "WhitelistAdminRemoved", accountRule)
	if err != nil {
		return nil, err
	}
	return &WhitelistAdminRoleWhitelistAdminRemovedIterator{contract: _WhitelistAdminRole.contract, event: "WhitelistAdminRemoved", logs: logs, sub: sub}, nil
}

// WatchWhitelistAdminRemoved is a free log subscription operation binding the contract event 0x0a8eb35e5ca14b3d6f28e4abf2f128dbab231a58b56e89beb5d636115001e165.
//
// Solidity: event WhitelistAdminRemoved(address indexed account)
func (_WhitelistAdminRole *WhitelistAdminRoleFilterer) WatchWhitelistAdminRemoved(opts *bind.WatchOpts, sink chan<- *WhitelistAdminRoleWhitelistAdminRemoved, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _WhitelistAdminRole.contract.WatchLogs(opts, "WhitelistAdminRemoved", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WhitelistAdminRoleWhitelistAdminRemoved)
				if err := _WhitelistAdminRole.contract.UnpackLog(event, "WhitelistAdminRemoved", log); err != nil {
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

// SolcCheckerABI is the input ABI used to generate the binding from.
const SolcCheckerABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"x\",\"type\":\"bytes\"}],\"name\":\"f\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// SolcCheckerBin is the compiled bytecode used for deploying new contracts.
const SolcCheckerBin = `0x`

// DeploySolcChecker deploys a new Ethereum contract, binding an instance of SolcChecker to it.
func DeploySolcChecker(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *SolcChecker, error) {
	parsed, err := abi.JSON(strings.NewReader(SolcCheckerABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(SolcCheckerBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &SolcChecker{SolcCheckerCaller: SolcCheckerCaller{contract: contract}, SolcCheckerTransactor: SolcCheckerTransactor{contract: contract}, SolcCheckerFilterer: SolcCheckerFilterer{contract: contract}}, nil
}

// SolcChecker is an auto generated Go binding around an Ethereum contract.
type SolcChecker struct {
	SolcCheckerCaller     // Read-only binding to the contract
	SolcCheckerTransactor // Write-only binding to the contract
	SolcCheckerFilterer   // Log filterer for contract events
}

// SolcCheckerCaller is an auto generated read-only Go binding around an Ethereum contract.
type SolcCheckerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SolcCheckerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SolcCheckerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SolcCheckerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SolcCheckerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SolcCheckerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SolcCheckerSession struct {
	Contract     *SolcChecker      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SolcCheckerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SolcCheckerCallerSession struct {
	Contract *SolcCheckerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// SolcCheckerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SolcCheckerTransactorSession struct {
	Contract     *SolcCheckerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// SolcCheckerRaw is an auto generated low-level Go binding around an Ethereum contract.
type SolcCheckerRaw struct {
	Contract *SolcChecker // Generic contract binding to access the raw methods on
}

// SolcCheckerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SolcCheckerCallerRaw struct {
	Contract *SolcCheckerCaller // Generic read-only contract binding to access the raw methods on
}

// SolcCheckerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SolcCheckerTransactorRaw struct {
	Contract *SolcCheckerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSolcChecker creates a new instance of SolcChecker, bound to a specific deployed contract.
func NewSolcChecker(address common.Address, backend bind.ContractBackend) (*SolcChecker, error) {
	contract, err := bindSolcChecker(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SolcChecker{SolcCheckerCaller: SolcCheckerCaller{contract: contract}, SolcCheckerTransactor: SolcCheckerTransactor{contract: contract}, SolcCheckerFilterer: SolcCheckerFilterer{contract: contract}}, nil
}

// NewSolcCheckerCaller creates a new read-only instance of SolcChecker, bound to a specific deployed contract.
func NewSolcCheckerCaller(address common.Address, caller bind.ContractCaller) (*SolcCheckerCaller, error) {
	contract, err := bindSolcChecker(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SolcCheckerCaller{contract: contract}, nil
}

// NewSolcCheckerTransactor creates a new write-only instance of SolcChecker, bound to a specific deployed contract.
func NewSolcCheckerTransactor(address common.Address, transactor bind.ContractTransactor) (*SolcCheckerTransactor, error) {
	contract, err := bindSolcChecker(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SolcCheckerTransactor{contract: contract}, nil
}

// NewSolcCheckerFilterer creates a new log filterer instance of SolcChecker, bound to a specific deployed contract.
func NewSolcCheckerFilterer(address common.Address, filterer bind.ContractFilterer) (*SolcCheckerFilterer, error) {
	contract, err := bindSolcChecker(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SolcCheckerFilterer{contract: contract}, nil
}

// bindSolcChecker binds a generic wrapper to an already deployed contract.
func bindSolcChecker(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(SolcCheckerABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SolcChecker *SolcCheckerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _SolcChecker.Contract.SolcCheckerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SolcChecker *SolcCheckerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SolcChecker.Contract.SolcCheckerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SolcChecker *SolcCheckerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SolcChecker.Contract.SolcCheckerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SolcChecker *SolcCheckerCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _SolcChecker.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SolcChecker *SolcCheckerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SolcChecker.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SolcChecker *SolcCheckerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SolcChecker.Contract.contract.Transact(opts, method, params...)
}

// F is a paid mutator transaction binding the contract method 0xd45754f8.
//
// Solidity: function f(bytes x) returns()
func (_SolcChecker *SolcCheckerTransactor) F(opts *bind.TransactOpts, x []byte) (*types.Transaction, error) {
	return _SolcChecker.contract.Transact(opts, "f", x)
}

// F is a paid mutator transaction binding the contract method 0xd45754f8.
//
// Solidity: function f(bytes x) returns()
func (_SolcChecker *SolcCheckerSession) F(x []byte) (*types.Transaction, error) {
	return _SolcChecker.Contract.F(&_SolcChecker.TransactOpts, x)
}

// F is a paid mutator transaction binding the contract method 0xd45754f8.
//
// Solidity: function f(bytes x) returns()
func (_SolcChecker *SolcCheckerTransactorSession) F(x []byte) (*types.Transaction, error) {
	return _SolcChecker.Contract.F(&_SolcChecker.TransactOpts, x)
}

// UsingOraclizeABI is the input ABI used to generate the binding from.
const UsingOraclizeABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"_myid\",\"type\":\"bytes32\"},{\"name\":\"_result\",\"type\":\"string\"}],\"name\":\"__callback\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_myid\",\"type\":\"bytes32\"},{\"name\":\"_result\",\"type\":\"string\"},{\"name\":\"_proof\",\"type\":\"bytes\"}],\"name\":\"__callback\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// UsingOraclizeBin is the compiled bytecode used for deploying new contracts.
const UsingOraclizeBin = `0x608060405234801561001057600080fd5b5061029b806100206000396000f3fe608060405234801561001057600080fd5b50600436106100365760003560e01c806327dc297e1461003b57806338bbfa50146100ea575b600080fd5b6100e86004803603604081101561005157600080fd5b8135919081019060408101602082013564010000000081111561007357600080fd5b82018360208201111561008557600080fd5b803590602001918460018302840111640100000000831117156100a757600080fd5b91908080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525092955061021e945050505050565b005b6100e86004803603606081101561010057600080fd5b8135919081019060408101602082013564010000000081111561012257600080fd5b82018360208201111561013457600080fd5b8035906020019184600183028401116401000000008311171561015657600080fd5b91908080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525092959493602081019350359150506401000000008111156101a957600080fd5b8201836020820111156101bb57600080fd5b803590602001918460018302840111640100000000831117156101dd57600080fd5b91908080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525092955061023e945050505050565b60408051600081526020810190915261023a908390839061023e565b5050565b5050600080805260036020527f3617319a054d772f909f7c479a2cebe5066e836a939412e32403c99029b92eff555056fea165627a7a723058201d21159cff3ef727cf00e7075106502f4cbd2935445dfdf9dbc7ddcc9af8c8ae0029`

// DeployUsingOraclize deploys a new Ethereum contract, binding an instance of UsingOraclize to it.
func DeployUsingOraclize(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *UsingOraclize, error) {
	parsed, err := abi.JSON(strings.NewReader(UsingOraclizeABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(UsingOraclizeBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &UsingOraclize{UsingOraclizeCaller: UsingOraclizeCaller{contract: contract}, UsingOraclizeTransactor: UsingOraclizeTransactor{contract: contract}, UsingOraclizeFilterer: UsingOraclizeFilterer{contract: contract}}, nil
}

// UsingOraclize is an auto generated Go binding around an Ethereum contract.
type UsingOraclize struct {
	UsingOraclizeCaller     // Read-only binding to the contract
	UsingOraclizeTransactor // Write-only binding to the contract
	UsingOraclizeFilterer   // Log filterer for contract events
}

// UsingOraclizeCaller is an auto generated read-only Go binding around an Ethereum contract.
type UsingOraclizeCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UsingOraclizeTransactor is an auto generated write-only Go binding around an Ethereum contract.
type UsingOraclizeTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UsingOraclizeFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type UsingOraclizeFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UsingOraclizeSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type UsingOraclizeSession struct {
	Contract     *UsingOraclize    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// UsingOraclizeCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type UsingOraclizeCallerSession struct {
	Contract *UsingOraclizeCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// UsingOraclizeTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type UsingOraclizeTransactorSession struct {
	Contract     *UsingOraclizeTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// UsingOraclizeRaw is an auto generated low-level Go binding around an Ethereum contract.
type UsingOraclizeRaw struct {
	Contract *UsingOraclize // Generic contract binding to access the raw methods on
}

// UsingOraclizeCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type UsingOraclizeCallerRaw struct {
	Contract *UsingOraclizeCaller // Generic read-only contract binding to access the raw methods on
}

// UsingOraclizeTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type UsingOraclizeTransactorRaw struct {
	Contract *UsingOraclizeTransactor // Generic write-only contract binding to access the raw methods on
}

// NewUsingOraclize creates a new instance of UsingOraclize, bound to a specific deployed contract.
func NewUsingOraclize(address common.Address, backend bind.ContractBackend) (*UsingOraclize, error) {
	contract, err := bindUsingOraclize(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &UsingOraclize{UsingOraclizeCaller: UsingOraclizeCaller{contract: contract}, UsingOraclizeTransactor: UsingOraclizeTransactor{contract: contract}, UsingOraclizeFilterer: UsingOraclizeFilterer{contract: contract}}, nil
}

// NewUsingOraclizeCaller creates a new read-only instance of UsingOraclize, bound to a specific deployed contract.
func NewUsingOraclizeCaller(address common.Address, caller bind.ContractCaller) (*UsingOraclizeCaller, error) {
	contract, err := bindUsingOraclize(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &UsingOraclizeCaller{contract: contract}, nil
}

// NewUsingOraclizeTransactor creates a new write-only instance of UsingOraclize, bound to a specific deployed contract.
func NewUsingOraclizeTransactor(address common.Address, transactor bind.ContractTransactor) (*UsingOraclizeTransactor, error) {
	contract, err := bindUsingOraclize(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &UsingOraclizeTransactor{contract: contract}, nil
}

// NewUsingOraclizeFilterer creates a new log filterer instance of UsingOraclize, bound to a specific deployed contract.
func NewUsingOraclizeFilterer(address common.Address, filterer bind.ContractFilterer) (*UsingOraclizeFilterer, error) {
	contract, err := bindUsingOraclize(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &UsingOraclizeFilterer{contract: contract}, nil
}

// bindUsingOraclize binds a generic wrapper to an already deployed contract.
func bindUsingOraclize(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(UsingOraclizeABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_UsingOraclize *UsingOraclizeRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _UsingOraclize.Contract.UsingOraclizeCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_UsingOraclize *UsingOraclizeRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _UsingOraclize.Contract.UsingOraclizeTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_UsingOraclize *UsingOraclizeRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _UsingOraclize.Contract.UsingOraclizeTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_UsingOraclize *UsingOraclizeCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _UsingOraclize.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_UsingOraclize *UsingOraclizeTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _UsingOraclize.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_UsingOraclize *UsingOraclizeTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _UsingOraclize.Contract.contract.Transact(opts, method, params...)
}

// Callback is a paid mutator transaction binding the contract method 0x38bbfa50.
//
// Solidity: function __callback(bytes32 _myid, string _result, bytes _proof) returns()
func (_UsingOraclize *UsingOraclizeTransactor) Callback(opts *bind.TransactOpts, _myid [32]byte, _result string, _proof []byte) (*types.Transaction, error) {
	return _UsingOraclize.contract.Transact(opts, "__callback", _myid, _result, _proof)
}

// Callback is a paid mutator transaction binding the contract method 0x38bbfa50.
//
// Solidity: function __callback(bytes32 _myid, string _result, bytes _proof) returns()
func (_UsingOraclize *UsingOraclizeSession) Callback(_myid [32]byte, _result string, _proof []byte) (*types.Transaction, error) {
	return _UsingOraclize.Contract.Callback(&_UsingOraclize.TransactOpts, _myid, _result, _proof)
}

// Callback is a paid mutator transaction binding the contract method 0x38bbfa50.
//
// Solidity: function __callback(bytes32 _myid, string _result, bytes _proof) returns()
func (_UsingOraclize *UsingOraclizeTransactorSession) Callback(_myid [32]byte, _result string, _proof []byte) (*types.Transaction, error) {
	return _UsingOraclize.Contract.Callback(&_UsingOraclize.TransactOpts, _myid, _result, _proof)
}
