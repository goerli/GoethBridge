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

// BridgeABI is the input ABI used to generate the binding from.
const BridgeABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"_addr\",\"type\":\"address\"}],\"name\":\"isAuthority\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_addr\",\"type\":\"address\"}],\"name\":\"addAuthority\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_recipient\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"},{\"name\":\"_fromChain\",\"type\":\"uint256\"},{\"name\":\"_txHash\",\"type\":\"bytes32\"}],\"name\":\"withdraw\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_recipient\",\"type\":\"address\"},{\"name\":\"_toChain\",\"type\":\"uint256\"}],\"name\":\"deposit\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"increaseThreshold\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_recipient\",\"type\":\"address\"},{\"name\":\"_toChain\",\"type\":\"uint256\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"withdrawTo\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_threshold\",\"type\":\"uint256\"}],\"name\":\"setThreshold\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"fundBridge\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_addr\",\"type\":\"address\"}],\"name\":\"removeAuthority\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"decreaseThreshold\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"bridge\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"fallback\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"ContractCreation\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"_addr\",\"type\":\"address\"}],\"name\":\"BridgeSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"_addr\",\"type\":\"address\"}],\"name\":\"BridgeFunded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"_addr\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"Paid\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"_addr\",\"type\":\"address\"}],\"name\":\"AuthorityAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"_addr\",\"type\":\"address\"}],\"name\":\"AuthorityRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"_threshold\",\"type\":\"uint256\"}],\"name\":\"ThresholdUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"_recipient\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_value\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"_toChain\",\"type\":\"uint256\"}],\"name\":\"Deposit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"_recipient\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_value\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"_fromChain\",\"type\":\"uint256\"}],\"name\":\"Withdraw\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"_txHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"_authority\",\"type\":\"address\"}],\"name\":\"SignedForWithdraw\",\"type\":\"event\"}]"

// Bridge is an auto generated Go binding around an Ethereum contract.
type Bridge struct {
	BridgeCaller     // Read-only binding to the contract
	BridgeTransactor // Write-only binding to the contract
	BridgeFilterer   // Log filterer for contract events
}

// BridgeCaller is an auto generated read-only Go binding around an Ethereum contract.
type BridgeCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BridgeTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BridgeTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BridgeFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BridgeFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BridgeSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BridgeSession struct {
	Contract     *Bridge           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BridgeCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BridgeCallerSession struct {
	Contract *BridgeCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// BridgeTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BridgeTransactorSession struct {
	Contract     *BridgeTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BridgeRaw is an auto generated low-level Go binding around an Ethereum contract.
type BridgeRaw struct {
	Contract *Bridge // Generic contract binding to access the raw methods on
}

// BridgeCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BridgeCallerRaw struct {
	Contract *BridgeCaller // Generic read-only contract binding to access the raw methods on
}

// BridgeTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BridgeTransactorRaw struct {
	Contract *BridgeTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBridge creates a new instance of Bridge, bound to a specific deployed contract.
func NewBridge(address common.Address, backend bind.ContractBackend) (*Bridge, error) {
	contract, err := bindBridge(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Bridge{BridgeCaller: BridgeCaller{contract: contract}, BridgeTransactor: BridgeTransactor{contract: contract}, BridgeFilterer: BridgeFilterer{contract: contract}}, nil
}

// NewBridgeCaller creates a new read-only instance of Bridge, bound to a specific deployed contract.
func NewBridgeCaller(address common.Address, caller bind.ContractCaller) (*BridgeCaller, error) {
	contract, err := bindBridge(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BridgeCaller{contract: contract}, nil
}

// NewBridgeTransactor creates a new write-only instance of Bridge, bound to a specific deployed contract.
func NewBridgeTransactor(address common.Address, transactor bind.ContractTransactor) (*BridgeTransactor, error) {
	contract, err := bindBridge(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BridgeTransactor{contract: contract}, nil
}

// NewBridgeFilterer creates a new log filterer instance of Bridge, bound to a specific deployed contract.
func NewBridgeFilterer(address common.Address, filterer bind.ContractFilterer) (*BridgeFilterer, error) {
	contract, err := bindBridge(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BridgeFilterer{contract: contract}, nil
}

// bindBridge binds a generic wrapper to an already deployed contract.
func bindBridge(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(BridgeABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Bridge *BridgeRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Bridge.Contract.BridgeCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Bridge *BridgeRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Bridge.Contract.BridgeTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Bridge *BridgeRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Bridge.Contract.BridgeTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Bridge *BridgeCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Bridge.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Bridge *BridgeTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Bridge.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Bridge *BridgeTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Bridge.Contract.contract.Transact(opts, method, params...)
}

// Bridge is a free data retrieval call binding the contract method 0xe78cea92.
//
// Solidity: function bridge() constant returns(address)
func (_Bridge *BridgeCaller) Bridge(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Bridge.contract.Call(opts, out, "bridge")
	return *ret0, err
}

// Bridge is a free data retrieval call binding the contract method 0xe78cea92.
//
// Solidity: function bridge() constant returns(address)
func (_Bridge *BridgeSession) Bridge() (common.Address, error) {
	return _Bridge.Contract.Bridge(&_Bridge.CallOpts)
}

// Bridge is a free data retrieval call binding the contract method 0xe78cea92.
//
// Solidity: function bridge() constant returns(address)
func (_Bridge *BridgeCallerSession) Bridge() (common.Address, error) {
	return _Bridge.Contract.Bridge(&_Bridge.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Bridge *BridgeCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Bridge.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Bridge *BridgeSession) Owner() (common.Address, error) {
	return _Bridge.Contract.Owner(&_Bridge.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Bridge *BridgeCallerSession) Owner() (common.Address, error) {
	return _Bridge.Contract.Owner(&_Bridge.CallOpts)
}

// AddAuthority is a paid mutator transaction binding the contract method 0x26defa73.
//
// Solidity: function addAuthority(_addr address) returns()
func (_Bridge *BridgeTransactor) AddAuthority(opts *bind.TransactOpts, _addr common.Address) (*types.Transaction, error) {
	return _Bridge.contract.Transact(opts, "addAuthority", _addr)
}

// AddAuthority is a paid mutator transaction binding the contract method 0x26defa73.
//
// Solidity: function addAuthority(_addr address) returns()
func (_Bridge *BridgeSession) AddAuthority(_addr common.Address) (*types.Transaction, error) {
	return _Bridge.Contract.AddAuthority(&_Bridge.TransactOpts, _addr)
}

// AddAuthority is a paid mutator transaction binding the contract method 0x26defa73.
//
// Solidity: function addAuthority(_addr address) returns()
func (_Bridge *BridgeTransactorSession) AddAuthority(_addr common.Address) (*types.Transaction, error) {
	return _Bridge.Contract.AddAuthority(&_Bridge.TransactOpts, _addr)
}

// DecreaseThreshold is a paid mutator transaction binding the contract method 0xe7508cc6.
//
// Solidity: function decreaseThreshold() returns()
func (_Bridge *BridgeTransactor) DecreaseThreshold(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Bridge.contract.Transact(opts, "decreaseThreshold")
}

// DecreaseThreshold is a paid mutator transaction binding the contract method 0xe7508cc6.
//
// Solidity: function decreaseThreshold() returns()
func (_Bridge *BridgeSession) DecreaseThreshold() (*types.Transaction, error) {
	return _Bridge.Contract.DecreaseThreshold(&_Bridge.TransactOpts)
}

// DecreaseThreshold is a paid mutator transaction binding the contract method 0xe7508cc6.
//
// Solidity: function decreaseThreshold() returns()
func (_Bridge *BridgeTransactorSession) DecreaseThreshold() (*types.Transaction, error) {
	return _Bridge.Contract.DecreaseThreshold(&_Bridge.TransactOpts)
}

// Deposit is a paid mutator transaction binding the contract method 0x47e7ef24.
//
// Solidity: function deposit(_recipient address, _toChain uint256) returns()
func (_Bridge *BridgeTransactor) Deposit(opts *bind.TransactOpts, _recipient common.Address, _toChain *big.Int) (*types.Transaction, error) {
	return _Bridge.contract.Transact(opts, "deposit", _recipient, _toChain)
}

// Deposit is a paid mutator transaction binding the contract method 0x47e7ef24.
//
// Solidity: function deposit(_recipient address, _toChain uint256) returns()
func (_Bridge *BridgeSession) Deposit(_recipient common.Address, _toChain *big.Int) (*types.Transaction, error) {
	return _Bridge.Contract.Deposit(&_Bridge.TransactOpts, _recipient, _toChain)
}

// Deposit is a paid mutator transaction binding the contract method 0x47e7ef24.
//
// Solidity: function deposit(_recipient address, _toChain uint256) returns()
func (_Bridge *BridgeTransactorSession) Deposit(_recipient common.Address, _toChain *big.Int) (*types.Transaction, error) {
	return _Bridge.Contract.Deposit(&_Bridge.TransactOpts, _recipient, _toChain)
}

// FundBridge is a paid mutator transaction binding the contract method 0xc9c0909f.
//
// Solidity: function fundBridge() returns()
func (_Bridge *BridgeTransactor) FundBridge(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Bridge.contract.Transact(opts, "fundBridge")
}

// FundBridge is a paid mutator transaction binding the contract method 0xc9c0909f.
//
// Solidity: function fundBridge() returns()
func (_Bridge *BridgeSession) FundBridge() (*types.Transaction, error) {
	return _Bridge.Contract.FundBridge(&_Bridge.TransactOpts)
}

// FundBridge is a paid mutator transaction binding the contract method 0xc9c0909f.
//
// Solidity: function fundBridge() returns()
func (_Bridge *BridgeTransactorSession) FundBridge() (*types.Transaction, error) {
	return _Bridge.Contract.FundBridge(&_Bridge.TransactOpts)
}

// IncreaseThreshold is a paid mutator transaction binding the contract method 0x4f548aae.
//
// Solidity: function increaseThreshold() returns()
func (_Bridge *BridgeTransactor) IncreaseThreshold(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Bridge.contract.Transact(opts, "increaseThreshold")
}

// IncreaseThreshold is a paid mutator transaction binding the contract method 0x4f548aae.
//
// Solidity: function increaseThreshold() returns()
func (_Bridge *BridgeSession) IncreaseThreshold() (*types.Transaction, error) {
	return _Bridge.Contract.IncreaseThreshold(&_Bridge.TransactOpts)
}

// IncreaseThreshold is a paid mutator transaction binding the contract method 0x4f548aae.
//
// Solidity: function increaseThreshold() returns()
func (_Bridge *BridgeTransactorSession) IncreaseThreshold() (*types.Transaction, error) {
	return _Bridge.Contract.IncreaseThreshold(&_Bridge.TransactOpts)
}

// IsAuthority is a paid mutator transaction binding the contract method 0x2330f247.
//
// Solidity: function isAuthority(_addr address) returns(bool)
func (_Bridge *BridgeTransactor) IsAuthority(opts *bind.TransactOpts, _addr common.Address) (*types.Transaction, error) {
	return _Bridge.contract.Transact(opts, "isAuthority", _addr)
}

// IsAuthority is a paid mutator transaction binding the contract method 0x2330f247.
//
// Solidity: function isAuthority(_addr address) returns(bool)
func (_Bridge *BridgeSession) IsAuthority(_addr common.Address) (*types.Transaction, error) {
	return _Bridge.Contract.IsAuthority(&_Bridge.TransactOpts, _addr)
}

// IsAuthority is a paid mutator transaction binding the contract method 0x2330f247.
//
// Solidity: function isAuthority(_addr address) returns(bool)
func (_Bridge *BridgeTransactorSession) IsAuthority(_addr common.Address) (*types.Transaction, error) {
	return _Bridge.Contract.IsAuthority(&_Bridge.TransactOpts, _addr)
}

// RemoveAuthority is a paid mutator transaction binding the contract method 0xd544e010.
//
// Solidity: function removeAuthority(_addr address) returns()
func (_Bridge *BridgeTransactor) RemoveAuthority(opts *bind.TransactOpts, _addr common.Address) (*types.Transaction, error) {
	return _Bridge.contract.Transact(opts, "removeAuthority", _addr)
}

// RemoveAuthority is a paid mutator transaction binding the contract method 0xd544e010.
//
// Solidity: function removeAuthority(_addr address) returns()
func (_Bridge *BridgeSession) RemoveAuthority(_addr common.Address) (*types.Transaction, error) {
	return _Bridge.Contract.RemoveAuthority(&_Bridge.TransactOpts, _addr)
}

// RemoveAuthority is a paid mutator transaction binding the contract method 0xd544e010.
//
// Solidity: function removeAuthority(_addr address) returns()
func (_Bridge *BridgeTransactorSession) RemoveAuthority(_addr common.Address) (*types.Transaction, error) {
	return _Bridge.Contract.RemoveAuthority(&_Bridge.TransactOpts, _addr)
}

// SetThreshold is a paid mutator transaction binding the contract method 0x960bfe04.
//
// Solidity: function setThreshold(_threshold uint256) returns()
func (_Bridge *BridgeTransactor) SetThreshold(opts *bind.TransactOpts, _threshold *big.Int) (*types.Transaction, error) {
	return _Bridge.contract.Transact(opts, "setThreshold", _threshold)
}

// SetThreshold is a paid mutator transaction binding the contract method 0x960bfe04.
//
// Solidity: function setThreshold(_threshold uint256) returns()
func (_Bridge *BridgeSession) SetThreshold(_threshold *big.Int) (*types.Transaction, error) {
	return _Bridge.Contract.SetThreshold(&_Bridge.TransactOpts, _threshold)
}

// SetThreshold is a paid mutator transaction binding the contract method 0x960bfe04.
//
// Solidity: function setThreshold(_threshold uint256) returns()
func (_Bridge *BridgeTransactorSession) SetThreshold(_threshold *big.Int) (*types.Transaction, error) {
	return _Bridge.Contract.SetThreshold(&_Bridge.TransactOpts, _threshold)
}

// Withdraw is a paid mutator transaction binding the contract method 0x4250a6f3.
//
// Solidity: function withdraw(_recipient address, _value uint256, _fromChain uint256, _txHash bytes32) returns()
func (_Bridge *BridgeTransactor) Withdraw(opts *bind.TransactOpts, _recipient common.Address, _value *big.Int, _fromChain *big.Int, _txHash [32]byte) (*types.Transaction, error) {
	return _Bridge.contract.Transact(opts, "withdraw", _recipient, _value, _fromChain, _txHash)
}

// Withdraw is a paid mutator transaction binding the contract method 0x4250a6f3.
//
// Solidity: function withdraw(_recipient address, _value uint256, _fromChain uint256, _txHash bytes32) returns()
func (_Bridge *BridgeSession) Withdraw(_recipient common.Address, _value *big.Int, _fromChain *big.Int, _txHash [32]byte) (*types.Transaction, error) {
	return _Bridge.Contract.Withdraw(&_Bridge.TransactOpts, _recipient, _value, _fromChain, _txHash)
}

// Withdraw is a paid mutator transaction binding the contract method 0x4250a6f3.
//
// Solidity: function withdraw(_recipient address, _value uint256, _fromChain uint256, _txHash bytes32) returns()
func (_Bridge *BridgeTransactorSession) Withdraw(_recipient common.Address, _value *big.Int, _fromChain *big.Int, _txHash [32]byte) (*types.Transaction, error) {
	return _Bridge.Contract.Withdraw(&_Bridge.TransactOpts, _recipient, _value, _fromChain, _txHash)
}

// WithdrawTo is a paid mutator transaction binding the contract method 0x5fcbc20e.
//
// Solidity: function withdrawTo(_recipient address, _toChain uint256, _value uint256) returns()
func (_Bridge *BridgeTransactor) WithdrawTo(opts *bind.TransactOpts, _recipient common.Address, _toChain *big.Int, _value *big.Int) (*types.Transaction, error) {
	return _Bridge.contract.Transact(opts, "withdrawTo", _recipient, _toChain, _value)
}

// WithdrawTo is a paid mutator transaction binding the contract method 0x5fcbc20e.
//
// Solidity: function withdrawTo(_recipient address, _toChain uint256, _value uint256) returns()
func (_Bridge *BridgeSession) WithdrawTo(_recipient common.Address, _toChain *big.Int, _value *big.Int) (*types.Transaction, error) {
	return _Bridge.Contract.WithdrawTo(&_Bridge.TransactOpts, _recipient, _toChain, _value)
}

// WithdrawTo is a paid mutator transaction binding the contract method 0x5fcbc20e.
//
// Solidity: function withdrawTo(_recipient address, _toChain uint256, _value uint256) returns()
func (_Bridge *BridgeTransactorSession) WithdrawTo(_recipient common.Address, _toChain *big.Int, _value *big.Int) (*types.Transaction, error) {
	return _Bridge.Contract.WithdrawTo(&_Bridge.TransactOpts, _recipient, _toChain, _value)
}

// BridgeAuthorityAddedIterator is returned from FilterAuthorityAdded and is used to iterate over the raw logs and unpacked data for AuthorityAdded events raised by the Bridge contract.
type BridgeAuthorityAddedIterator struct {
	Event *BridgeAuthorityAdded // Event containing the contract specifics and raw log

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
func (it *BridgeAuthorityAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeAuthorityAdded)
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
		it.Event = new(BridgeAuthorityAdded)
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
func (it *BridgeAuthorityAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeAuthorityAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeAuthorityAdded represents a AuthorityAdded event raised by the Bridge contract.
type BridgeAuthorityAdded struct {
	Addr common.Address
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterAuthorityAdded is a free log retrieval operation binding the contract event 0x550a8ae64ec9d6640b6f168a26d3e6364b90defe8110c92135aa775b279e54ea.
//
// Solidity: e AuthorityAdded(_addr address)
func (_Bridge *BridgeFilterer) FilterAuthorityAdded(opts *bind.FilterOpts) (*BridgeAuthorityAddedIterator, error) {

	logs, sub, err := _Bridge.contract.FilterLogs(opts, "AuthorityAdded")
	if err != nil {
		return nil, err
	}
	return &BridgeAuthorityAddedIterator{contract: _Bridge.contract, event: "AuthorityAdded", logs: logs, sub: sub}, nil
}

// WatchAuthorityAdded is a free log subscription operation binding the contract event 0x550a8ae64ec9d6640b6f168a26d3e6364b90defe8110c92135aa775b279e54ea.
//
// Solidity: e AuthorityAdded(_addr address)
func (_Bridge *BridgeFilterer) WatchAuthorityAdded(opts *bind.WatchOpts, sink chan<- *BridgeAuthorityAdded) (event.Subscription, error) {

	logs, sub, err := _Bridge.contract.WatchLogs(opts, "AuthorityAdded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeAuthorityAdded)
				if err := _Bridge.contract.UnpackLog(event, "AuthorityAdded", log); err != nil {
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

// BridgeAuthorityRemovedIterator is returned from FilterAuthorityRemoved and is used to iterate over the raw logs and unpacked data for AuthorityRemoved events raised by the Bridge contract.
type BridgeAuthorityRemovedIterator struct {
	Event *BridgeAuthorityRemoved // Event containing the contract specifics and raw log

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
func (it *BridgeAuthorityRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeAuthorityRemoved)
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
		it.Event = new(BridgeAuthorityRemoved)
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
func (it *BridgeAuthorityRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeAuthorityRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeAuthorityRemoved represents a AuthorityRemoved event raised by the Bridge contract.
type BridgeAuthorityRemoved struct {
	Addr common.Address
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterAuthorityRemoved is a free log retrieval operation binding the contract event 0x272215cde179041f7a3e8da6f8aabc7c8fc1336ccd73aba698cb825a80d3be48.
//
// Solidity: e AuthorityRemoved(_addr address)
func (_Bridge *BridgeFilterer) FilterAuthorityRemoved(opts *bind.FilterOpts) (*BridgeAuthorityRemovedIterator, error) {

	logs, sub, err := _Bridge.contract.FilterLogs(opts, "AuthorityRemoved")
	if err != nil {
		return nil, err
	}
	return &BridgeAuthorityRemovedIterator{contract: _Bridge.contract, event: "AuthorityRemoved", logs: logs, sub: sub}, nil
}

// WatchAuthorityRemoved is a free log subscription operation binding the contract event 0x272215cde179041f7a3e8da6f8aabc7c8fc1336ccd73aba698cb825a80d3be48.
//
// Solidity: e AuthorityRemoved(_addr address)
func (_Bridge *BridgeFilterer) WatchAuthorityRemoved(opts *bind.WatchOpts, sink chan<- *BridgeAuthorityRemoved) (event.Subscription, error) {

	logs, sub, err := _Bridge.contract.WatchLogs(opts, "AuthorityRemoved")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeAuthorityRemoved)
				if err := _Bridge.contract.UnpackLog(event, "AuthorityRemoved", log); err != nil {
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

// BridgeBridgeFundedIterator is returned from FilterBridgeFunded and is used to iterate over the raw logs and unpacked data for BridgeFunded events raised by the Bridge contract.
type BridgeBridgeFundedIterator struct {
	Event *BridgeBridgeFunded // Event containing the contract specifics and raw log

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
func (it *BridgeBridgeFundedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeBridgeFunded)
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
		it.Event = new(BridgeBridgeFunded)
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
func (it *BridgeBridgeFundedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeBridgeFundedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeBridgeFunded represents a BridgeFunded event raised by the Bridge contract.
type BridgeBridgeFunded struct {
	Addr common.Address
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterBridgeFunded is a free log retrieval operation binding the contract event 0xc2520f24142cb24b12b04df358be485159ec7ec1a3c3ad25fa65e1a226e4eec3.
//
// Solidity: e BridgeFunded(_addr address)
func (_Bridge *BridgeFilterer) FilterBridgeFunded(opts *bind.FilterOpts) (*BridgeBridgeFundedIterator, error) {

	logs, sub, err := _Bridge.contract.FilterLogs(opts, "BridgeFunded")
	if err != nil {
		return nil, err
	}
	return &BridgeBridgeFundedIterator{contract: _Bridge.contract, event: "BridgeFunded", logs: logs, sub: sub}, nil
}

// WatchBridgeFunded is a free log subscription operation binding the contract event 0xc2520f24142cb24b12b04df358be485159ec7ec1a3c3ad25fa65e1a226e4eec3.
//
// Solidity: e BridgeFunded(_addr address)
func (_Bridge *BridgeFilterer) WatchBridgeFunded(opts *bind.WatchOpts, sink chan<- *BridgeBridgeFunded) (event.Subscription, error) {

	logs, sub, err := _Bridge.contract.WatchLogs(opts, "BridgeFunded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeBridgeFunded)
				if err := _Bridge.contract.UnpackLog(event, "BridgeFunded", log); err != nil {
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

// BridgeBridgeSetIterator is returned from FilterBridgeSet and is used to iterate over the raw logs and unpacked data for BridgeSet events raised by the Bridge contract.
type BridgeBridgeSetIterator struct {
	Event *BridgeBridgeSet // Event containing the contract specifics and raw log

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
func (it *BridgeBridgeSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeBridgeSet)
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
		it.Event = new(BridgeBridgeSet)
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
func (it *BridgeBridgeSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeBridgeSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeBridgeSet represents a BridgeSet event raised by the Bridge contract.
type BridgeBridgeSet struct {
	Addr common.Address
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterBridgeSet is a free log retrieval operation binding the contract event 0xa49730bff544fd0b716395c592e39c6fd2d2481a19b9229b5b240483db95a495.
//
// Solidity: e BridgeSet(_addr address)
func (_Bridge *BridgeFilterer) FilterBridgeSet(opts *bind.FilterOpts) (*BridgeBridgeSetIterator, error) {

	logs, sub, err := _Bridge.contract.FilterLogs(opts, "BridgeSet")
	if err != nil {
		return nil, err
	}
	return &BridgeBridgeSetIterator{contract: _Bridge.contract, event: "BridgeSet", logs: logs, sub: sub}, nil
}

// WatchBridgeSet is a free log subscription operation binding the contract event 0xa49730bff544fd0b716395c592e39c6fd2d2481a19b9229b5b240483db95a495.
//
// Solidity: e BridgeSet(_addr address)
func (_Bridge *BridgeFilterer) WatchBridgeSet(opts *bind.WatchOpts, sink chan<- *BridgeBridgeSet) (event.Subscription, error) {

	logs, sub, err := _Bridge.contract.WatchLogs(opts, "BridgeSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeBridgeSet)
				if err := _Bridge.contract.UnpackLog(event, "BridgeSet", log); err != nil {
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

// BridgeContractCreationIterator is returned from FilterContractCreation and is used to iterate over the raw logs and unpacked data for ContractCreation events raised by the Bridge contract.
type BridgeContractCreationIterator struct {
	Event *BridgeContractCreation // Event containing the contract specifics and raw log

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
func (it *BridgeContractCreationIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeContractCreation)
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
		it.Event = new(BridgeContractCreation)
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
func (it *BridgeContractCreationIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeContractCreationIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeContractCreation represents a ContractCreation event raised by the Bridge contract.
type BridgeContractCreation struct {
	Owner common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterContractCreation is a free log retrieval operation binding the contract event 0x4db17dd5e4732fb6da34a148104a592783ca119a1e7bb8829eba6cbadef0b511.
//
// Solidity: e ContractCreation(_owner address)
func (_Bridge *BridgeFilterer) FilterContractCreation(opts *bind.FilterOpts) (*BridgeContractCreationIterator, error) {

	logs, sub, err := _Bridge.contract.FilterLogs(opts, "ContractCreation")
	if err != nil {
		return nil, err
	}
	return &BridgeContractCreationIterator{contract: _Bridge.contract, event: "ContractCreation", logs: logs, sub: sub}, nil
}

// WatchContractCreation is a free log subscription operation binding the contract event 0x4db17dd5e4732fb6da34a148104a592783ca119a1e7bb8829eba6cbadef0b511.
//
// Solidity: e ContractCreation(_owner address)
func (_Bridge *BridgeFilterer) WatchContractCreation(opts *bind.WatchOpts, sink chan<- *BridgeContractCreation) (event.Subscription, error) {

	logs, sub, err := _Bridge.contract.WatchLogs(opts, "ContractCreation")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeContractCreation)
				if err := _Bridge.contract.UnpackLog(event, "ContractCreation", log); err != nil {
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

// BridgeDepositIterator is returned from FilterDeposit and is used to iterate over the raw logs and unpacked data for Deposit events raised by the Bridge contract.
type BridgeDepositIterator struct {
	Event *BridgeDeposit // Event containing the contract specifics and raw log

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
func (it *BridgeDepositIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeDeposit)
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
		it.Event = new(BridgeDeposit)
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
func (it *BridgeDepositIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeDepositIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeDeposit represents a Deposit event raised by the Bridge contract.
type BridgeDeposit struct {
	Recipient common.Address
	Value     *big.Int
	ToChain   *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterDeposit is a free log retrieval operation binding the contract event 0x90890809c654f11d6e72a28fa60149770a0d11ec6c92319d6ceb2bb0a4ea1a15.
//
// Solidity: e Deposit(_recipient address, _value uint256, _toChain uint256)
func (_Bridge *BridgeFilterer) FilterDeposit(opts *bind.FilterOpts) (*BridgeDepositIterator, error) {

	logs, sub, err := _Bridge.contract.FilterLogs(opts, "Deposit")
	if err != nil {
		return nil, err
	}
	return &BridgeDepositIterator{contract: _Bridge.contract, event: "Deposit", logs: logs, sub: sub}, nil
}

// WatchDeposit is a free log subscription operation binding the contract event 0x90890809c654f11d6e72a28fa60149770a0d11ec6c92319d6ceb2bb0a4ea1a15.
//
// Solidity: e Deposit(_recipient address, _value uint256, _toChain uint256)
func (_Bridge *BridgeFilterer) WatchDeposit(opts *bind.WatchOpts, sink chan<- *BridgeDeposit) (event.Subscription, error) {

	logs, sub, err := _Bridge.contract.WatchLogs(opts, "Deposit")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeDeposit)
				if err := _Bridge.contract.UnpackLog(event, "Deposit", log); err != nil {
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

// BridgePaidIterator is returned from FilterPaid and is used to iterate over the raw logs and unpacked data for Paid events raised by the Bridge contract.
type BridgePaidIterator struct {
	Event *BridgePaid // Event containing the contract specifics and raw log

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
func (it *BridgePaidIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgePaid)
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
		it.Event = new(BridgePaid)
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
func (it *BridgePaidIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgePaidIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgePaid represents a Paid event raised by the Bridge contract.
type BridgePaid struct {
	Addr  common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterPaid is a free log retrieval operation binding the contract event 0x737c69225d647e5994eab1a6c301bf6d9232beb2759ae1e27a8966b4732bc489.
//
// Solidity: e Paid(_addr address, _value uint256)
func (_Bridge *BridgeFilterer) FilterPaid(opts *bind.FilterOpts) (*BridgePaidIterator, error) {

	logs, sub, err := _Bridge.contract.FilterLogs(opts, "Paid")
	if err != nil {
		return nil, err
	}
	return &BridgePaidIterator{contract: _Bridge.contract, event: "Paid", logs: logs, sub: sub}, nil
}

// WatchPaid is a free log subscription operation binding the contract event 0x737c69225d647e5994eab1a6c301bf6d9232beb2759ae1e27a8966b4732bc489.
//
// Solidity: e Paid(_addr address, _value uint256)
func (_Bridge *BridgeFilterer) WatchPaid(opts *bind.WatchOpts, sink chan<- *BridgePaid) (event.Subscription, error) {

	logs, sub, err := _Bridge.contract.WatchLogs(opts, "Paid")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgePaid)
				if err := _Bridge.contract.UnpackLog(event, "Paid", log); err != nil {
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

// BridgeSignedForWithdrawIterator is returned from FilterSignedForWithdraw and is used to iterate over the raw logs and unpacked data for SignedForWithdraw events raised by the Bridge contract.
type BridgeSignedForWithdrawIterator struct {
	Event *BridgeSignedForWithdraw // Event containing the contract specifics and raw log

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
func (it *BridgeSignedForWithdrawIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeSignedForWithdraw)
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
		it.Event = new(BridgeSignedForWithdraw)
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
func (it *BridgeSignedForWithdrawIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeSignedForWithdrawIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeSignedForWithdraw represents a SignedForWithdraw event raised by the Bridge contract.
type BridgeSignedForWithdraw struct {
	TxHash    [32]byte
	Authority common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterSignedForWithdraw is a free log retrieval operation binding the contract event 0x9bada88c355a68c8a9dd9bd8f64aa19f89cb105e7fb179d7857ea7fb34878ad3.
//
// Solidity: e SignedForWithdraw(_txHash bytes32, _authority address)
func (_Bridge *BridgeFilterer) FilterSignedForWithdraw(opts *bind.FilterOpts) (*BridgeSignedForWithdrawIterator, error) {

	logs, sub, err := _Bridge.contract.FilterLogs(opts, "SignedForWithdraw")
	if err != nil {
		return nil, err
	}
	return &BridgeSignedForWithdrawIterator{contract: _Bridge.contract, event: "SignedForWithdraw", logs: logs, sub: sub}, nil
}

// WatchSignedForWithdraw is a free log subscription operation binding the contract event 0x9bada88c355a68c8a9dd9bd8f64aa19f89cb105e7fb179d7857ea7fb34878ad3.
//
// Solidity: e SignedForWithdraw(_txHash bytes32, _authority address)
func (_Bridge *BridgeFilterer) WatchSignedForWithdraw(opts *bind.WatchOpts, sink chan<- *BridgeSignedForWithdraw) (event.Subscription, error) {

	logs, sub, err := _Bridge.contract.WatchLogs(opts, "SignedForWithdraw")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeSignedForWithdraw)
				if err := _Bridge.contract.UnpackLog(event, "SignedForWithdraw", log); err != nil {
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

// BridgeThresholdUpdatedIterator is returned from FilterThresholdUpdated and is used to iterate over the raw logs and unpacked data for ThresholdUpdated events raised by the Bridge contract.
type BridgeThresholdUpdatedIterator struct {
	Event *BridgeThresholdUpdated // Event containing the contract specifics and raw log

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
func (it *BridgeThresholdUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeThresholdUpdated)
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
		it.Event = new(BridgeThresholdUpdated)
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
func (it *BridgeThresholdUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeThresholdUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeThresholdUpdated represents a ThresholdUpdated event raised by the Bridge contract.
type BridgeThresholdUpdated struct {
	Threshold *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterThresholdUpdated is a free log retrieval operation binding the contract event 0xadfa8ecb21b6962ebcd0adbd9ab985b7b4c5b5eb3b0dead683171565c7bfe171.
//
// Solidity: e ThresholdUpdated(_threshold uint256)
func (_Bridge *BridgeFilterer) FilterThresholdUpdated(opts *bind.FilterOpts) (*BridgeThresholdUpdatedIterator, error) {

	logs, sub, err := _Bridge.contract.FilterLogs(opts, "ThresholdUpdated")
	if err != nil {
		return nil, err
	}
	return &BridgeThresholdUpdatedIterator{contract: _Bridge.contract, event: "ThresholdUpdated", logs: logs, sub: sub}, nil
}

// WatchThresholdUpdated is a free log subscription operation binding the contract event 0xadfa8ecb21b6962ebcd0adbd9ab985b7b4c5b5eb3b0dead683171565c7bfe171.
//
// Solidity: e ThresholdUpdated(_threshold uint256)
func (_Bridge *BridgeFilterer) WatchThresholdUpdated(opts *bind.WatchOpts, sink chan<- *BridgeThresholdUpdated) (event.Subscription, error) {

	logs, sub, err := _Bridge.contract.WatchLogs(opts, "ThresholdUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeThresholdUpdated)
				if err := _Bridge.contract.UnpackLog(event, "ThresholdUpdated", log); err != nil {
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

// BridgeWithdrawIterator is returned from FilterWithdraw and is used to iterate over the raw logs and unpacked data for Withdraw events raised by the Bridge contract.
type BridgeWithdrawIterator struct {
	Event *BridgeWithdraw // Event containing the contract specifics and raw log

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
func (it *BridgeWithdrawIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeWithdraw)
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
		it.Event = new(BridgeWithdraw)
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
func (it *BridgeWithdrawIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeWithdrawIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeWithdraw represents a Withdraw event raised by the Bridge contract.
type BridgeWithdraw struct {
	Recipient common.Address
	Value     *big.Int
	FromChain *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterWithdraw is a free log retrieval operation binding the contract event 0xf279e6a1f5e320cca91135676d9cb6e44ca8a08c0b88342bcdb1144f6511b568.
//
// Solidity: e Withdraw(_recipient address, _value uint256, _fromChain uint256)
func (_Bridge *BridgeFilterer) FilterWithdraw(opts *bind.FilterOpts) (*BridgeWithdrawIterator, error) {

	logs, sub, err := _Bridge.contract.FilterLogs(opts, "Withdraw")
	if err != nil {
		return nil, err
	}
	return &BridgeWithdrawIterator{contract: _Bridge.contract, event: "Withdraw", logs: logs, sub: sub}, nil
}

// WatchWithdraw is a free log subscription operation binding the contract event 0xf279e6a1f5e320cca91135676d9cb6e44ca8a08c0b88342bcdb1144f6511b568.
//
// Solidity: e Withdraw(_recipient address, _value uint256, _fromChain uint256)
func (_Bridge *BridgeFilterer) WatchWithdraw(opts *bind.WatchOpts, sink chan<- *BridgeWithdraw) (event.Subscription, error) {

	logs, sub, err := _Bridge.contract.WatchLogs(opts, "Withdraw")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeWithdraw)
				if err := _Bridge.contract.UnpackLog(event, "Withdraw", log); err != nil {
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
