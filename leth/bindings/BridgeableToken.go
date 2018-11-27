// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package bindings

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

// BridgeableTokenABI is the input ABI used to generate the binding from.
const BridgeableTokenABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_spender\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_recipient\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"},{\"name\":\"_toChain\",\"type\":\"uint256\"}],\"name\":\"deposit\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_from\",\"type\":\"address\"},{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"name\":\"\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_spender\",\"type\":\"address\"},{\"name\":\"_subtractedValue\",\"type\":\"uint256\"}],\"name\":\"decreaseApproval\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_addr\",\"type\":\"address\"}],\"name\":\"setBridge\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_recipient\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"},{\"name\":\"_fromChain\",\"type\":\"uint256\"}],\"name\":\"withdraw\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_spender\",\"type\":\"address\"},{\"name\":\"_addedValue\",\"type\":\"uint256\"}],\"name\":\"increaseApproval\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\"},{\"name\":\"_spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"_recipient\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_value\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"_toChain\",\"type\":\"uint256\"}],\"name\":\"Deposit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"_recipient\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_value\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"_fromChain\",\"type\":\"uint256\"}],\"name\":\"Withdraw\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"}]"

// BridgeableToken is an auto generated Go binding around an Ethereum contract.
type BridgeableToken struct {
	BridgeableTokenCaller     // Read-only binding to the contract
	BridgeableTokenTransactor // Write-only binding to the contract
	BridgeableTokenFilterer   // Log filterer for contract events
}

// BridgeableTokenCaller is an auto generated read-only Go binding around an Ethereum contract.
type BridgeableTokenCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BridgeableTokenTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BridgeableTokenTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BridgeableTokenFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BridgeableTokenFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BridgeableTokenSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BridgeableTokenSession struct {
	Contract     *BridgeableToken  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BridgeableTokenCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BridgeableTokenCallerSession struct {
	Contract *BridgeableTokenCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// BridgeableTokenTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BridgeableTokenTransactorSession struct {
	Contract     *BridgeableTokenTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// BridgeableTokenRaw is an auto generated low-level Go binding around an Ethereum contract.
type BridgeableTokenRaw struct {
	Contract *BridgeableToken // Generic contract binding to access the raw methods on
}

// BridgeableTokenCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BridgeableTokenCallerRaw struct {
	Contract *BridgeableTokenCaller // Generic read-only contract binding to access the raw methods on
}

// BridgeableTokenTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BridgeableTokenTransactorRaw struct {
	Contract *BridgeableTokenTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBridgeableToken creates a new instance of BridgeableToken, bound to a specific deployed contract.
func NewBridgeableToken(address common.Address, backend bind.ContractBackend) (*BridgeableToken, error) {
	contract, err := bindBridgeableToken(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &BridgeableToken{BridgeableTokenCaller: BridgeableTokenCaller{contract: contract}, BridgeableTokenTransactor: BridgeableTokenTransactor{contract: contract}, BridgeableTokenFilterer: BridgeableTokenFilterer{contract: contract}}, nil
}

// NewBridgeableTokenCaller creates a new read-only instance of BridgeableToken, bound to a specific deployed contract.
func NewBridgeableTokenCaller(address common.Address, caller bind.ContractCaller) (*BridgeableTokenCaller, error) {
	contract, err := bindBridgeableToken(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BridgeableTokenCaller{contract: contract}, nil
}

// NewBridgeableTokenTransactor creates a new write-only instance of BridgeableToken, bound to a specific deployed contract.
func NewBridgeableTokenTransactor(address common.Address, transactor bind.ContractTransactor) (*BridgeableTokenTransactor, error) {
	contract, err := bindBridgeableToken(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BridgeableTokenTransactor{contract: contract}, nil
}

// NewBridgeableTokenFilterer creates a new log filterer instance of BridgeableToken, bound to a specific deployed contract.
func NewBridgeableTokenFilterer(address common.Address, filterer bind.ContractFilterer) (*BridgeableTokenFilterer, error) {
	contract, err := bindBridgeableToken(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BridgeableTokenFilterer{contract: contract}, nil
}

// bindBridgeableToken binds a generic wrapper to an already deployed contract.
func bindBridgeableToken(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(BridgeableTokenABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BridgeableToken *BridgeableTokenRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _BridgeableToken.Contract.BridgeableTokenCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BridgeableToken *BridgeableTokenRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BridgeableToken.Contract.BridgeableTokenTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BridgeableToken *BridgeableTokenRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BridgeableToken.Contract.BridgeableTokenTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BridgeableToken *BridgeableTokenCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _BridgeableToken.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BridgeableToken *BridgeableTokenTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BridgeableToken.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BridgeableToken *BridgeableTokenTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BridgeableToken.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(_owner address, _spender address) constant returns(uint256)
func (_BridgeableToken *BridgeableTokenCaller) Allowance(opts *bind.CallOpts, _owner common.Address, _spender common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _BridgeableToken.contract.Call(opts, out, "allowance", _owner, _spender)
	return *ret0, err
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(_owner address, _spender address) constant returns(uint256)
func (_BridgeableToken *BridgeableTokenSession) Allowance(_owner common.Address, _spender common.Address) (*big.Int, error) {
	return _BridgeableToken.Contract.Allowance(&_BridgeableToken.CallOpts, _owner, _spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(_owner address, _spender address) constant returns(uint256)
func (_BridgeableToken *BridgeableTokenCallerSession) Allowance(_owner common.Address, _spender common.Address) (*big.Int, error) {
	return _BridgeableToken.Contract.Allowance(&_BridgeableToken.CallOpts, _owner, _spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(_owner address) constant returns(uint256)
func (_BridgeableToken *BridgeableTokenCaller) BalanceOf(opts *bind.CallOpts, _owner common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _BridgeableToken.contract.Call(opts, out, "balanceOf", _owner)
	return *ret0, err
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(_owner address) constant returns(uint256)
func (_BridgeableToken *BridgeableTokenSession) BalanceOf(_owner common.Address) (*big.Int, error) {
	return _BridgeableToken.Contract.BalanceOf(&_BridgeableToken.CallOpts, _owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(_owner address) constant returns(uint256)
func (_BridgeableToken *BridgeableTokenCallerSession) BalanceOf(_owner common.Address) (*big.Int, error) {
	return _BridgeableToken.Contract.BalanceOf(&_BridgeableToken.CallOpts, _owner)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() constant returns(uint8)
func (_BridgeableToken *BridgeableTokenCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var (
		ret0 = new(uint8)
	)
	out := ret0
	err := _BridgeableToken.contract.Call(opts, out, "decimals")
	return *ret0, err
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() constant returns(uint8)
func (_BridgeableToken *BridgeableTokenSession) Decimals() (uint8, error) {
	return _BridgeableToken.Contract.Decimals(&_BridgeableToken.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() constant returns(uint8)
func (_BridgeableToken *BridgeableTokenCallerSession) Decimals() (uint8, error) {
	return _BridgeableToken.Contract.Decimals(&_BridgeableToken.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_BridgeableToken *BridgeableTokenCaller) Name(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _BridgeableToken.contract.Call(opts, out, "name")
	return *ret0, err
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_BridgeableToken *BridgeableTokenSession) Name() (string, error) {
	return _BridgeableToken.Contract.Name(&_BridgeableToken.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_BridgeableToken *BridgeableTokenCallerSession) Name() (string, error) {
	return _BridgeableToken.Contract.Name(&_BridgeableToken.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() constant returns(string)
func (_BridgeableToken *BridgeableTokenCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _BridgeableToken.contract.Call(opts, out, "symbol")
	return *ret0, err
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() constant returns(string)
func (_BridgeableToken *BridgeableTokenSession) Symbol() (string, error) {
	return _BridgeableToken.Contract.Symbol(&_BridgeableToken.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() constant returns(string)
func (_BridgeableToken *BridgeableTokenCallerSession) Symbol() (string, error) {
	return _BridgeableToken.Contract.Symbol(&_BridgeableToken.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_BridgeableToken *BridgeableTokenCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _BridgeableToken.contract.Call(opts, out, "totalSupply")
	return *ret0, err
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_BridgeableToken *BridgeableTokenSession) TotalSupply() (*big.Int, error) {
	return _BridgeableToken.Contract.TotalSupply(&_BridgeableToken.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_BridgeableToken *BridgeableTokenCallerSession) TotalSupply() (*big.Int, error) {
	return _BridgeableToken.Contract.TotalSupply(&_BridgeableToken.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(_spender address, _value uint256) returns(bool)
func (_BridgeableToken *BridgeableTokenTransactor) Approve(opts *bind.TransactOpts, _spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _BridgeableToken.contract.Transact(opts, "approve", _spender, _value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(_spender address, _value uint256) returns(bool)
func (_BridgeableToken *BridgeableTokenSession) Approve(_spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _BridgeableToken.Contract.Approve(&_BridgeableToken.TransactOpts, _spender, _value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(_spender address, _value uint256) returns(bool)
func (_BridgeableToken *BridgeableTokenTransactorSession) Approve(_spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _BridgeableToken.Contract.Approve(&_BridgeableToken.TransactOpts, _spender, _value)
}

// DecreaseApproval is a paid mutator transaction binding the contract method 0x66188463.
//
// Solidity: function decreaseApproval(_spender address, _subtractedValue uint256) returns(bool)
func (_BridgeableToken *BridgeableTokenTransactor) DecreaseApproval(opts *bind.TransactOpts, _spender common.Address, _subtractedValue *big.Int) (*types.Transaction, error) {
	return _BridgeableToken.contract.Transact(opts, "decreaseApproval", _spender, _subtractedValue)
}

// DecreaseApproval is a paid mutator transaction binding the contract method 0x66188463.
//
// Solidity: function decreaseApproval(_spender address, _subtractedValue uint256) returns(bool)
func (_BridgeableToken *BridgeableTokenSession) DecreaseApproval(_spender common.Address, _subtractedValue *big.Int) (*types.Transaction, error) {
	return _BridgeableToken.Contract.DecreaseApproval(&_BridgeableToken.TransactOpts, _spender, _subtractedValue)
}

// DecreaseApproval is a paid mutator transaction binding the contract method 0x66188463.
//
// Solidity: function decreaseApproval(_spender address, _subtractedValue uint256) returns(bool)
func (_BridgeableToken *BridgeableTokenTransactorSession) DecreaseApproval(_spender common.Address, _subtractedValue *big.Int) (*types.Transaction, error) {
	return _BridgeableToken.Contract.DecreaseApproval(&_BridgeableToken.TransactOpts, _spender, _subtractedValue)
}

// Deposit is a paid mutator transaction binding the contract method 0x0efe6a8b.
//
// Solidity: function deposit(_recipient address, _value uint256, _toChain uint256) returns()
func (_BridgeableToken *BridgeableTokenTransactor) Deposit(opts *bind.TransactOpts, _recipient common.Address, _value *big.Int, _toChain *big.Int) (*types.Transaction, error) {
	return _BridgeableToken.contract.Transact(opts, "deposit", _recipient, _value, _toChain)
}

// Deposit is a paid mutator transaction binding the contract method 0x0efe6a8b.
//
// Solidity: function deposit(_recipient address, _value uint256, _toChain uint256) returns()
func (_BridgeableToken *BridgeableTokenSession) Deposit(_recipient common.Address, _value *big.Int, _toChain *big.Int) (*types.Transaction, error) {
	return _BridgeableToken.Contract.Deposit(&_BridgeableToken.TransactOpts, _recipient, _value, _toChain)
}

// Deposit is a paid mutator transaction binding the contract method 0x0efe6a8b.
//
// Solidity: function deposit(_recipient address, _value uint256, _toChain uint256) returns()
func (_BridgeableToken *BridgeableTokenTransactorSession) Deposit(_recipient common.Address, _value *big.Int, _toChain *big.Int) (*types.Transaction, error) {
	return _BridgeableToken.Contract.Deposit(&_BridgeableToken.TransactOpts, _recipient, _value, _toChain)
}

// IncreaseApproval is a paid mutator transaction binding the contract method 0xd73dd623.
//
// Solidity: function increaseApproval(_spender address, _addedValue uint256) returns(bool)
func (_BridgeableToken *BridgeableTokenTransactor) IncreaseApproval(opts *bind.TransactOpts, _spender common.Address, _addedValue *big.Int) (*types.Transaction, error) {
	return _BridgeableToken.contract.Transact(opts, "increaseApproval", _spender, _addedValue)
}

// IncreaseApproval is a paid mutator transaction binding the contract method 0xd73dd623.
//
// Solidity: function increaseApproval(_spender address, _addedValue uint256) returns(bool)
func (_BridgeableToken *BridgeableTokenSession) IncreaseApproval(_spender common.Address, _addedValue *big.Int) (*types.Transaction, error) {
	return _BridgeableToken.Contract.IncreaseApproval(&_BridgeableToken.TransactOpts, _spender, _addedValue)
}

// IncreaseApproval is a paid mutator transaction binding the contract method 0xd73dd623.
//
// Solidity: function increaseApproval(_spender address, _addedValue uint256) returns(bool)
func (_BridgeableToken *BridgeableTokenTransactorSession) IncreaseApproval(_spender common.Address, _addedValue *big.Int) (*types.Transaction, error) {
	return _BridgeableToken.Contract.IncreaseApproval(&_BridgeableToken.TransactOpts, _spender, _addedValue)
}

// SetBridge is a paid mutator transaction binding the contract method 0x8dd14802.
//
// Solidity: function setBridge(_addr address) returns()
func (_BridgeableToken *BridgeableTokenTransactor) SetBridge(opts *bind.TransactOpts, _addr common.Address) (*types.Transaction, error) {
	return _BridgeableToken.contract.Transact(opts, "setBridge", _addr)
}

// SetBridge is a paid mutator transaction binding the contract method 0x8dd14802.
//
// Solidity: function setBridge(_addr address) returns()
func (_BridgeableToken *BridgeableTokenSession) SetBridge(_addr common.Address) (*types.Transaction, error) {
	return _BridgeableToken.Contract.SetBridge(&_BridgeableToken.TransactOpts, _addr)
}

// SetBridge is a paid mutator transaction binding the contract method 0x8dd14802.
//
// Solidity: function setBridge(_addr address) returns()
func (_BridgeableToken *BridgeableTokenTransactorSession) SetBridge(_addr common.Address) (*types.Transaction, error) {
	return _BridgeableToken.Contract.SetBridge(&_BridgeableToken.TransactOpts, _addr)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(_to address, _value uint256) returns(bool)
func (_BridgeableToken *BridgeableTokenTransactor) Transfer(opts *bind.TransactOpts, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _BridgeableToken.contract.Transact(opts, "transfer", _to, _value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(_to address, _value uint256) returns(bool)
func (_BridgeableToken *BridgeableTokenSession) Transfer(_to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _BridgeableToken.Contract.Transfer(&_BridgeableToken.TransactOpts, _to, _value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(_to address, _value uint256) returns(bool)
func (_BridgeableToken *BridgeableTokenTransactorSession) Transfer(_to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _BridgeableToken.Contract.Transfer(&_BridgeableToken.TransactOpts, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(_from address, _to address, _value uint256) returns(bool)
func (_BridgeableToken *BridgeableTokenTransactor) TransferFrom(opts *bind.TransactOpts, _from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _BridgeableToken.contract.Transact(opts, "transferFrom", _from, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(_from address, _to address, _value uint256) returns(bool)
func (_BridgeableToken *BridgeableTokenSession) TransferFrom(_from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _BridgeableToken.Contract.TransferFrom(&_BridgeableToken.TransactOpts, _from, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(_from address, _to address, _value uint256) returns(bool)
func (_BridgeableToken *BridgeableTokenTransactorSession) TransferFrom(_from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _BridgeableToken.Contract.TransferFrom(&_BridgeableToken.TransactOpts, _from, _to, _value)
}

// Withdraw is a paid mutator transaction binding the contract method 0xb5c5f672.
//
// Solidity: function withdraw(_recipient address, _value uint256, _fromChain uint256) returns()
func (_BridgeableToken *BridgeableTokenTransactor) Withdraw(opts *bind.TransactOpts, _recipient common.Address, _value *big.Int, _fromChain *big.Int) (*types.Transaction, error) {
	return _BridgeableToken.contract.Transact(opts, "withdraw", _recipient, _value, _fromChain)
}

// Withdraw is a paid mutator transaction binding the contract method 0xb5c5f672.
//
// Solidity: function withdraw(_recipient address, _value uint256, _fromChain uint256) returns()
func (_BridgeableToken *BridgeableTokenSession) Withdraw(_recipient common.Address, _value *big.Int, _fromChain *big.Int) (*types.Transaction, error) {
	return _BridgeableToken.Contract.Withdraw(&_BridgeableToken.TransactOpts, _recipient, _value, _fromChain)
}

// Withdraw is a paid mutator transaction binding the contract method 0xb5c5f672.
//
// Solidity: function withdraw(_recipient address, _value uint256, _fromChain uint256) returns()
func (_BridgeableToken *BridgeableTokenTransactorSession) Withdraw(_recipient common.Address, _value *big.Int, _fromChain *big.Int) (*types.Transaction, error) {
	return _BridgeableToken.Contract.Withdraw(&_BridgeableToken.TransactOpts, _recipient, _value, _fromChain)
}

// BridgeableTokenApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the BridgeableToken contract.
type BridgeableTokenApprovalIterator struct {
	Event *BridgeableTokenApproval // Event containing the contract specifics and raw log

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
func (it *BridgeableTokenApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeableTokenApproval)
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
		it.Event = new(BridgeableTokenApproval)
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
func (it *BridgeableTokenApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeableTokenApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeableTokenApproval represents a Approval event raised by the BridgeableToken contract.
type BridgeableTokenApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: e Approval(owner indexed address, spender indexed address, value uint256)
func (_BridgeableToken *BridgeableTokenFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*BridgeableTokenApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _BridgeableToken.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &BridgeableTokenApprovalIterator{contract: _BridgeableToken.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: e Approval(owner indexed address, spender indexed address, value uint256)
func (_BridgeableToken *BridgeableTokenFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *BridgeableTokenApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _BridgeableToken.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeableTokenApproval)
				if err := _BridgeableToken.contract.UnpackLog(event, "Approval", log); err != nil {
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

// BridgeableTokenDepositIterator is returned from FilterDeposit and is used to iterate over the raw logs and unpacked data for Deposit events raised by the BridgeableToken contract.
type BridgeableTokenDepositIterator struct {
	Event *BridgeableTokenDeposit // Event containing the contract specifics and raw log

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
func (it *BridgeableTokenDepositIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeableTokenDeposit)
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
		it.Event = new(BridgeableTokenDeposit)
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
func (it *BridgeableTokenDepositIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeableTokenDepositIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeableTokenDeposit represents a Deposit event raised by the BridgeableToken contract.
type BridgeableTokenDeposit struct {
	Recipient common.Address
	Value     *big.Int
	ToChain   *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterDeposit is a free log retrieval operation binding the contract event 0x90890809c654f11d6e72a28fa60149770a0d11ec6c92319d6ceb2bb0a4ea1a15.
//
// Solidity: e Deposit(_recipient address, _value uint256, _toChain uint256)
func (_BridgeableToken *BridgeableTokenFilterer) FilterDeposit(opts *bind.FilterOpts) (*BridgeableTokenDepositIterator, error) {

	logs, sub, err := _BridgeableToken.contract.FilterLogs(opts, "Deposit")
	if err != nil {
		return nil, err
	}
	return &BridgeableTokenDepositIterator{contract: _BridgeableToken.contract, event: "Deposit", logs: logs, sub: sub}, nil
}

// WatchDeposit is a free log subscription operation binding the contract event 0x90890809c654f11d6e72a28fa60149770a0d11ec6c92319d6ceb2bb0a4ea1a15.
//
// Solidity: e Deposit(_recipient address, _value uint256, _toChain uint256)
func (_BridgeableToken *BridgeableTokenFilterer) WatchDeposit(opts *bind.WatchOpts, sink chan<- *BridgeableTokenDeposit) (event.Subscription, error) {

	logs, sub, err := _BridgeableToken.contract.WatchLogs(opts, "Deposit")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeableTokenDeposit)
				if err := _BridgeableToken.contract.UnpackLog(event, "Deposit", log); err != nil {
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

// BridgeableTokenTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the BridgeableToken contract.
type BridgeableTokenTransferIterator struct {
	Event *BridgeableTokenTransfer // Event containing the contract specifics and raw log

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
func (it *BridgeableTokenTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeableTokenTransfer)
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
		it.Event = new(BridgeableTokenTransfer)
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
func (it *BridgeableTokenTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeableTokenTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeableTokenTransfer represents a Transfer event raised by the BridgeableToken contract.
type BridgeableTokenTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: e Transfer(from indexed address, to indexed address, value uint256)
func (_BridgeableToken *BridgeableTokenFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*BridgeableTokenTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _BridgeableToken.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &BridgeableTokenTransferIterator{contract: _BridgeableToken.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: e Transfer(from indexed address, to indexed address, value uint256)
func (_BridgeableToken *BridgeableTokenFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *BridgeableTokenTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _BridgeableToken.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeableTokenTransfer)
				if err := _BridgeableToken.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// BridgeableTokenWithdrawIterator is returned from FilterWithdraw and is used to iterate over the raw logs and unpacked data for Withdraw events raised by the BridgeableToken contract.
type BridgeableTokenWithdrawIterator struct {
	Event *BridgeableTokenWithdraw // Event containing the contract specifics and raw log

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
func (it *BridgeableTokenWithdrawIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeableTokenWithdraw)
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
		it.Event = new(BridgeableTokenWithdraw)
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
func (it *BridgeableTokenWithdrawIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeableTokenWithdrawIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeableTokenWithdraw represents a Withdraw event raised by the BridgeableToken contract.
type BridgeableTokenWithdraw struct {
	Recipient common.Address
	Value     *big.Int
	FromChain *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterWithdraw is a free log retrieval operation binding the contract event 0xf279e6a1f5e320cca91135676d9cb6e44ca8a08c0b88342bcdb1144f6511b568.
//
// Solidity: e Withdraw(_recipient address, _value uint256, _fromChain uint256)
func (_BridgeableToken *BridgeableTokenFilterer) FilterWithdraw(opts *bind.FilterOpts) (*BridgeableTokenWithdrawIterator, error) {

	logs, sub, err := _BridgeableToken.contract.FilterLogs(opts, "Withdraw")
	if err != nil {
		return nil, err
	}
	return &BridgeableTokenWithdrawIterator{contract: _BridgeableToken.contract, event: "Withdraw", logs: logs, sub: sub}, nil
}

// WatchWithdraw is a free log subscription operation binding the contract event 0xf279e6a1f5e320cca91135676d9cb6e44ca8a08c0b88342bcdb1144f6511b568.
//
// Solidity: e Withdraw(_recipient address, _value uint256, _fromChain uint256)
func (_BridgeableToken *BridgeableTokenFilterer) WatchWithdraw(opts *bind.WatchOpts, sink chan<- *BridgeableTokenWithdraw) (event.Subscription, error) {

	logs, sub, err := _BridgeableToken.contract.WatchLogs(opts, "Withdraw")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeableTokenWithdraw)
				if err := _BridgeableToken.contract.UnpackLog(event, "Withdraw", log); err != nil {
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
