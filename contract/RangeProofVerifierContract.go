// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contract

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
)

// RangeProofVerifierContractMetaData contains all meta data concerning the RangeProofVerifierContract contract.
var RangeProofVerifierContractMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint256[2]\",\"name\":\"a\",\"type\":\"uint256[2]\"},{\"internalType\":\"uint256[2][2]\",\"name\":\"b\",\"type\":\"uint256[2][2]\"},{\"internalType\":\"uint256[2]\",\"name\":\"c\",\"type\":\"uint256[2]\"},{\"internalType\":\"uint256[1]\",\"name\":\"input\",\"type\":\"uint256[1]\"}],\"name\":\"verifyProof\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"r\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// RangeProofVerifierContractABI is the input ABI used to generate the binding from.
// Deprecated: Use RangeProofVerifierContractMetaData.ABI instead.
var RangeProofVerifierContractABI = RangeProofVerifierContractMetaData.ABI

// RangeProofVerifierContract is an auto generated Go binding around an Ethereum contract.
type RangeProofVerifierContract struct {
	RangeProofVerifierContractCaller     // Read-only binding to the contract
	RangeProofVerifierContractTransactor // Write-only binding to the contract
	RangeProofVerifierContractFilterer   // Log filterer for contract events
}

// RangeProofVerifierContractCaller is an auto generated read-only Go binding around an Ethereum contract.
type RangeProofVerifierContractCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RangeProofVerifierContractTransactor is an auto generated write-only Go binding around an Ethereum contract.
type RangeProofVerifierContractTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RangeProofVerifierContractFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type RangeProofVerifierContractFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RangeProofVerifierContractSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type RangeProofVerifierContractSession struct {
	Contract     *RangeProofVerifierContract // Generic contract binding to set the session for
	CallOpts     bind.CallOpts               // Call options to use throughout this session
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// RangeProofVerifierContractCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type RangeProofVerifierContractCallerSession struct {
	Contract *RangeProofVerifierContractCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                     // Call options to use throughout this session
}

// RangeProofVerifierContractTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type RangeProofVerifierContractTransactorSession struct {
	Contract     *RangeProofVerifierContractTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                     // Transaction auth options to use throughout this session
}

// RangeProofVerifierContractRaw is an auto generated low-level Go binding around an Ethereum contract.
type RangeProofVerifierContractRaw struct {
	Contract *RangeProofVerifierContract // Generic contract binding to access the raw methods on
}

// RangeProofVerifierContractCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type RangeProofVerifierContractCallerRaw struct {
	Contract *RangeProofVerifierContractCaller // Generic read-only contract binding to access the raw methods on
}

// RangeProofVerifierContractTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type RangeProofVerifierContractTransactorRaw struct {
	Contract *RangeProofVerifierContractTransactor // Generic write-only contract binding to access the raw methods on
}

// NewRangeProofVerifierContract creates a new instance of RangeProofVerifierContract, bound to a specific deployed contract.
func NewRangeProofVerifierContract(address common.Address, backend bind.ContractBackend) (*RangeProofVerifierContract, error) {
	contract, err := bindRangeProofVerifierContract(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &RangeProofVerifierContract{RangeProofVerifierContractCaller: RangeProofVerifierContractCaller{contract: contract}, RangeProofVerifierContractTransactor: RangeProofVerifierContractTransactor{contract: contract}, RangeProofVerifierContractFilterer: RangeProofVerifierContractFilterer{contract: contract}}, nil
}

// NewRangeProofVerifierContractCaller creates a new read-only instance of RangeProofVerifierContract, bound to a specific deployed contract.
func NewRangeProofVerifierContractCaller(address common.Address, caller bind.ContractCaller) (*RangeProofVerifierContractCaller, error) {
	contract, err := bindRangeProofVerifierContract(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &RangeProofVerifierContractCaller{contract: contract}, nil
}

// NewRangeProofVerifierContractTransactor creates a new write-only instance of RangeProofVerifierContract, bound to a specific deployed contract.
func NewRangeProofVerifierContractTransactor(address common.Address, transactor bind.ContractTransactor) (*RangeProofVerifierContractTransactor, error) {
	contract, err := bindRangeProofVerifierContract(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &RangeProofVerifierContractTransactor{contract: contract}, nil
}

// NewRangeProofVerifierContractFilterer creates a new log filterer instance of RangeProofVerifierContract, bound to a specific deployed contract.
func NewRangeProofVerifierContractFilterer(address common.Address, filterer bind.ContractFilterer) (*RangeProofVerifierContractFilterer, error) {
	contract, err := bindRangeProofVerifierContract(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &RangeProofVerifierContractFilterer{contract: contract}, nil
}

// bindRangeProofVerifierContract binds a generic wrapper to an already deployed contract.
func bindRangeProofVerifierContract(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(RangeProofVerifierContractABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_RangeProofVerifierContract *RangeProofVerifierContractRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _RangeProofVerifierContract.Contract.RangeProofVerifierContractCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_RangeProofVerifierContract *RangeProofVerifierContractRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RangeProofVerifierContract.Contract.RangeProofVerifierContractTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_RangeProofVerifierContract *RangeProofVerifierContractRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RangeProofVerifierContract.Contract.RangeProofVerifierContractTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_RangeProofVerifierContract *RangeProofVerifierContractCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _RangeProofVerifierContract.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_RangeProofVerifierContract *RangeProofVerifierContractTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RangeProofVerifierContract.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_RangeProofVerifierContract *RangeProofVerifierContractTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RangeProofVerifierContract.Contract.contract.Transact(opts, method, params...)
}

// VerifyProof is a free data retrieval call binding the contract method 0x43753b4d.
//
// Solidity: function verifyProof(uint256[2] a, uint256[2][2] b, uint256[2] c, uint256[1] input) view returns(bool r)
func (_RangeProofVerifierContract *RangeProofVerifierContractCaller) VerifyProof(opts *bind.CallOpts, a [2]*big.Int, b [2][2]*big.Int, c [2]*big.Int, input [1]*big.Int) (bool, error) {
	var out []interface{}
	err := _RangeProofVerifierContract.contract.Call(opts, &out, "verifyProof", a, b, c, input)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// VerifyProof is a free data retrieval call binding the contract method 0x43753b4d.
//
// Solidity: function verifyProof(uint256[2] a, uint256[2][2] b, uint256[2] c, uint256[1] input) view returns(bool r)
func (_RangeProofVerifierContract *RangeProofVerifierContractSession) VerifyProof(a [2]*big.Int, b [2][2]*big.Int, c [2]*big.Int, input [1]*big.Int) (bool, error) {
	return _RangeProofVerifierContract.Contract.VerifyProof(&_RangeProofVerifierContract.CallOpts, a, b, c, input)
}

// VerifyProof is a free data retrieval call binding the contract method 0x43753b4d.
//
// Solidity: function verifyProof(uint256[2] a, uint256[2][2] b, uint256[2] c, uint256[1] input) view returns(bool r)
func (_RangeProofVerifierContract *RangeProofVerifierContractCallerSession) VerifyProof(a [2]*big.Int, b [2][2]*big.Int, c [2]*big.Int, input [1]*big.Int) (bool, error) {
	return _RangeProofVerifierContract.Contract.VerifyProof(&_RangeProofVerifierContract.CallOpts, a, b, c, input)
}
