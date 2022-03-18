// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package utils

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

// TestStructMetaData contains all meta data concerning the TestStruct contract.
var TestStructMetaData = &bind.MetaData{
	ABI: "[{\"constant\":false,\"inputs\":[{\"name\":\"name\",\"type\":\"string\"},{\"name\":\"haoma\",\"type\":\"string\"}],\"name\":\"buyBallot\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"name\",\"type\":\"string\"},{\"name\":\"password\",\"type\":\"string\"}],\"name\":\"register\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"accountList\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"accoutNum\",\"outputs\":[{\"name\":\"\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"name\",\"type\":\"string\"}],\"name\":\"getHaoMa\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"name\",\"type\":\"string\"}],\"name\":\"validName\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"name\",\"type\":\"string\"},{\"name\":\"password\",\"type\":\"string\"}],\"name\":\"validPassword\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// TestStructABI is the input ABI used to generate the binding from.
// Deprecated: Use TestStructMetaData.ABI instead.
var TestStructABI = TestStructMetaData.ABI

// TestStruct is an auto generated Go binding around an Ethereum contract.
type TestStruct struct {
	TestStructCaller     // Read-only binding to the contract
	TestStructTransactor // Write-only binding to the contract
	TestStructFilterer   // Log filterer for contract events
}

// TestStructCaller is an auto generated read-only Go binding around an Ethereum contract.
type TestStructCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TestStructTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TestStructTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TestStructFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TestStructFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TestStructSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TestStructSession struct {
	Contract     *TestStruct       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TestStructCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TestStructCallerSession struct {
	Contract *TestStructCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// TestStructTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TestStructTransactorSession struct {
	Contract     *TestStructTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// TestStructRaw is an auto generated low-level Go binding around an Ethereum contract.
type TestStructRaw struct {
	Contract *TestStruct // Generic contract binding to access the raw methods on
}

// TestStructCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TestStructCallerRaw struct {
	Contract *TestStructCaller // Generic read-only contract binding to access the raw methods on
}

// TestStructTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TestStructTransactorRaw struct {
	Contract *TestStructTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTestStruct creates a new instance of TestStruct, bound to a specific deployed contract.
func NewTestStruct(address common.Address, backend bind.ContractBackend) (*TestStruct, error) {
	contract, err := bindTestStruct(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &TestStruct{TestStructCaller: TestStructCaller{contract: contract}, TestStructTransactor: TestStructTransactor{contract: contract}, TestStructFilterer: TestStructFilterer{contract: contract}}, nil
}

// NewTestStructCaller creates a new read-only instance of TestStruct, bound to a specific deployed contract.
func NewTestStructCaller(address common.Address, caller bind.ContractCaller) (*TestStructCaller, error) {
	contract, err := bindTestStruct(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TestStructCaller{contract: contract}, nil
}

// NewTestStructTransactor creates a new write-only instance of TestStruct, bound to a specific deployed contract.
func NewTestStructTransactor(address common.Address, transactor bind.ContractTransactor) (*TestStructTransactor, error) {
	contract, err := bindTestStruct(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TestStructTransactor{contract: contract}, nil
}

// NewTestStructFilterer creates a new log filterer instance of TestStruct, bound to a specific deployed contract.
func NewTestStructFilterer(address common.Address, filterer bind.ContractFilterer) (*TestStructFilterer, error) {
	contract, err := bindTestStruct(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TestStructFilterer{contract: contract}, nil
}

// bindTestStruct binds a generic wrapper to an already deployed contract.
func bindTestStruct(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(TestStructABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TestStruct *TestStructRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TestStruct.Contract.TestStructCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TestStruct *TestStructRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TestStruct.Contract.TestStructTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TestStruct *TestStructRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TestStruct.Contract.TestStructTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TestStruct *TestStructCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TestStruct.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TestStruct *TestStructTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TestStruct.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TestStruct *TestStructTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TestStruct.Contract.contract.Transact(opts, method, params...)
}

// AccountList is a free data retrieval call binding the contract method 0x4ea98d16.
//
// Solidity: function accountList(uint256 ) view returns(string)
func (_TestStruct *TestStructCaller) AccountList(opts *bind.CallOpts, arg0 *big.Int) (string, error) {
	var out []interface{}
	err := _TestStruct.contract.Call(opts, &out, "accountList", arg0)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// AccountList is a free data retrieval call binding the contract method 0x4ea98d16.
//
// Solidity: function accountList(uint256 ) view returns(string)
func (_TestStruct *TestStructSession) AccountList(arg0 *big.Int) (string, error) {
	return _TestStruct.Contract.AccountList(&_TestStruct.CallOpts, arg0)
}

// AccountList is a free data retrieval call binding the contract method 0x4ea98d16.
//
// Solidity: function accountList(uint256 ) view returns(string)
func (_TestStruct *TestStructCallerSession) AccountList(arg0 *big.Int) (string, error) {
	return _TestStruct.Contract.AccountList(&_TestStruct.CallOpts, arg0)
}

// AccoutNum is a free data retrieval call binding the contract method 0x7e732842.
//
// Solidity: function accoutNum() view returns(uint8)
func (_TestStruct *TestStructCaller) AccoutNum(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _TestStruct.contract.Call(opts, &out, "accoutNum")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// AccoutNum is a free data retrieval call binding the contract method 0x7e732842.
//
// Solidity: function accoutNum() view returns(uint8)
func (_TestStruct *TestStructSession) AccoutNum() (uint8, error) {
	return _TestStruct.Contract.AccoutNum(&_TestStruct.CallOpts)
}

// AccoutNum is a free data retrieval call binding the contract method 0x7e732842.
//
// Solidity: function accoutNum() view returns(uint8)
func (_TestStruct *TestStructCallerSession) AccoutNum() (uint8, error) {
	return _TestStruct.Contract.AccoutNum(&_TestStruct.CallOpts)
}

// GetHaoMa is a free data retrieval call binding the contract method 0x8a23e876.
//
// Solidity: function getHaoMa(string name) view returns(string)
func (_TestStruct *TestStructCaller) GetHaoMa(opts *bind.CallOpts, name string) (string, error) {
	var out []interface{}
	err := _TestStruct.contract.Call(opts, &out, "getHaoMa", name)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// GetHaoMa is a free data retrieval call binding the contract method 0x8a23e876.
//
// Solidity: function getHaoMa(string name) view returns(string)
func (_TestStruct *TestStructSession) GetHaoMa(name string) (string, error) {
	return _TestStruct.Contract.GetHaoMa(&_TestStruct.CallOpts, name)
}

// GetHaoMa is a free data retrieval call binding the contract method 0x8a23e876.
//
// Solidity: function getHaoMa(string name) view returns(string)
func (_TestStruct *TestStructCallerSession) GetHaoMa(name string) (string, error) {
	return _TestStruct.Contract.GetHaoMa(&_TestStruct.CallOpts, name)
}

// ValidName is a free data retrieval call binding the contract method 0x7a921a1a.
//
// Solidity: function validName(string name) view returns(bool)
func (_TestStruct *TestStructCaller) ValidName(opts *bind.CallOpts, name string) (bool, error) {
	var out []interface{}
	err := _TestStruct.contract.Call(opts, &out, "validName", name)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// ValidName is a free data retrieval call binding the contract method 0x7a921a1a.
//
// Solidity: function validName(string name) view returns(bool)
func (_TestStruct *TestStructSession) ValidName(name string) (bool, error) {
	return _TestStruct.Contract.ValidName(&_TestStruct.CallOpts, name)
}

// ValidName is a free data retrieval call binding the contract method 0x7a921a1a.
//
// Solidity: function validName(string name) view returns(bool)
func (_TestStruct *TestStructCallerSession) ValidName(name string) (bool, error) {
	return _TestStruct.Contract.ValidName(&_TestStruct.CallOpts, name)
}

// ValidPassword is a free data retrieval call binding the contract method 0xcce7f668.
//
// Solidity: function validPassword(string name, string password) view returns(bool)
func (_TestStruct *TestStructCaller) ValidPassword(opts *bind.CallOpts, name string, password string) (bool, error) {
	var out []interface{}
	err := _TestStruct.contract.Call(opts, &out, "validPassword", name, password)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// ValidPassword is a free data retrieval call binding the contract method 0xcce7f668.
//
// Solidity: function validPassword(string name, string password) view returns(bool)
func (_TestStruct *TestStructSession) ValidPassword(name string, password string) (bool, error) {
	return _TestStruct.Contract.ValidPassword(&_TestStruct.CallOpts, name, password)
}

// ValidPassword is a free data retrieval call binding the contract method 0xcce7f668.
//
// Solidity: function validPassword(string name, string password) view returns(bool)
func (_TestStruct *TestStructCallerSession) ValidPassword(name string, password string) (bool, error) {
	return _TestStruct.Contract.ValidPassword(&_TestStruct.CallOpts, name, password)
}

// BuyBallot is a paid mutator transaction binding the contract method 0x86382ce0.
//
// Solidity: function buyBallot(string name, string haoma) returns(bool)
func (_TestStruct *TestStructTransactor) BuyBallot(opts *bind.TransactOpts, name string, haoma string) (*types.Transaction, error) {
	return _TestStruct.contract.Transact(opts, "buyBallot", name, haoma)
}

// BuyBallot is a paid mutator transaction binding the contract method 0x86382ce0.
//
// Solidity: function buyBallot(string name, string haoma) returns(bool)
func (_TestStruct *TestStructSession) BuyBallot(name string, haoma string) (*types.Transaction, error) {
	return _TestStruct.Contract.BuyBallot(&_TestStruct.TransactOpts, name, haoma)
}

// BuyBallot is a paid mutator transaction binding the contract method 0x86382ce0.
//
// Solidity: function buyBallot(string name, string haoma) returns(bool)
func (_TestStruct *TestStructTransactorSession) BuyBallot(name string, haoma string) (*types.Transaction, error) {
	return _TestStruct.Contract.BuyBallot(&_TestStruct.TransactOpts, name, haoma)
}

// Register is a paid mutator transaction binding the contract method 0x3ffbd47f.
//
// Solidity: function register(string name, string password) returns()
func (_TestStruct *TestStructTransactor) Register(opts *bind.TransactOpts, name string, password string) (*types.Transaction, error) {
	return _TestStruct.contract.Transact(opts, "register", name, password)
}

// Register is a paid mutator transaction binding the contract method 0x3ffbd47f.
//
// Solidity: function register(string name, string password) returns()
func (_TestStruct *TestStructSession) Register(name string, password string) (*types.Transaction, error) {
	return _TestStruct.Contract.Register(&_TestStruct.TransactOpts, name, password)
}

// Register is a paid mutator transaction binding the contract method 0x3ffbd47f.
//
// Solidity: function register(string name, string password) returns()
func (_TestStruct *TestStructTransactorSession) Register(name string, password string) (*types.Transaction, error) {
	return _TestStruct.Contract.Register(&_TestStruct.TransactOpts, name, password)
}
