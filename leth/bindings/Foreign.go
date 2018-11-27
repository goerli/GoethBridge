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

// ForeignABI is the input ABI used to generate the binding from.
const ForeignABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"_recipient\",\"type\":\"address\"},{\"name\":\"_toChain\",\"type\":\"uint256\"}],\"name\":\"deposit\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"ContractCreation\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"_recipient\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_value\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"_toChain\",\"type\":\"uint256\"}],\"name\":\"Deposit\",\"type\":\"event\"}]"

// Foreign is an auto generated Go binding around an Ethereum contract.
type Foreign struct {
	ForeignCaller     // Read-only binding to the contract
	ForeignTransactor // Write-only binding to the contract
	ForeignFilterer   // Log filterer for contract events
}

// ForeignCaller is an auto generated read-only Go binding around an Ethereum contract.
type ForeignCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ForeignTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ForeignTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ForeignFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ForeignFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ForeignSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ForeignSession struct {
	Contract     *Foreign          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ForeignCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ForeignCallerSession struct {
	Contract *ForeignCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// ForeignTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ForeignTransactorSession struct {
	Contract     *ForeignTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// ForeignRaw is an auto generated low-level Go binding around an Ethereum contract.
type ForeignRaw struct {
	Contract *Foreign // Generic contract binding to access the raw methods on
}

// ForeignCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ForeignCallerRaw struct {
	Contract *ForeignCaller // Generic read-only contract binding to access the raw methods on
}

// ForeignTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ForeignTransactorRaw struct {
	Contract *ForeignTransactor // Generic write-only contract binding to access the raw methods on
}

// NewForeign creates a new instance of Foreign, bound to a specific deployed contract.
func NewForeign(address common.Address, backend bind.ContractBackend) (*Foreign, error) {
	contract, err := bindForeign(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Foreign{ForeignCaller: ForeignCaller{contract: contract}, ForeignTransactor: ForeignTransactor{contract: contract}, ForeignFilterer: ForeignFilterer{contract: contract}}, nil
}

// NewForeignCaller creates a new read-only instance of Foreign, bound to a specific deployed contract.
func NewForeignCaller(address common.Address, caller bind.ContractCaller) (*ForeignCaller, error) {
	contract, err := bindForeign(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ForeignCaller{contract: contract}, nil
}

// NewForeignTransactor creates a new write-only instance of Foreign, bound to a specific deployed contract.
func NewForeignTransactor(address common.Address, transactor bind.ContractTransactor) (*ForeignTransactor, error) {
	contract, err := bindForeign(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ForeignTransactor{contract: contract}, nil
}

// NewForeignFilterer creates a new log filterer instance of Foreign, bound to a specific deployed contract.
func NewForeignFilterer(address common.Address, filterer bind.ContractFilterer) (*ForeignFilterer, error) {
	contract, err := bindForeign(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ForeignFilterer{contract: contract}, nil
}

// bindForeign binds a generic wrapper to an already deployed contract.
func bindForeign(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ForeignABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Foreign *ForeignRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Foreign.Contract.ForeignCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Foreign *ForeignRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Foreign.Contract.ForeignTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Foreign *ForeignRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Foreign.Contract.ForeignTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Foreign *ForeignCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Foreign.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Foreign *ForeignTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Foreign.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Foreign *ForeignTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Foreign.Contract.contract.Transact(opts, method, params...)
}

// Deposit is a paid mutator transaction binding the contract method 0x47e7ef24.
//
// Solidity: function deposit(_recipient address, _toChain uint256) returns()
func (_Foreign *ForeignTransactor) Deposit(opts *bind.TransactOpts, _recipient common.Address, _toChain *big.Int) (*types.Transaction, error) {
	return _Foreign.contract.Transact(opts, "deposit", _recipient, _toChain)
}

// Deposit is a paid mutator transaction binding the contract method 0x47e7ef24.
//
// Solidity: function deposit(_recipient address, _toChain uint256) returns()
func (_Foreign *ForeignSession) Deposit(_recipient common.Address, _toChain *big.Int) (*types.Transaction, error) {
	return _Foreign.Contract.Deposit(&_Foreign.TransactOpts, _recipient, _toChain)
}

// Deposit is a paid mutator transaction binding the contract method 0x47e7ef24.
//
// Solidity: function deposit(_recipient address, _toChain uint256) returns()
func (_Foreign *ForeignTransactorSession) Deposit(_recipient common.Address, _toChain *big.Int) (*types.Transaction, error) {
	return _Foreign.Contract.Deposit(&_Foreign.TransactOpts, _recipient, _toChain)
}

// ForeignContractCreationIterator is returned from FilterContractCreation and is used to iterate over the raw logs and unpacked data for ContractCreation events raised by the Foreign contract.
type ForeignContractCreationIterator struct {
	Event *ForeignContractCreation // Event containing the contract specifics and raw log

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
func (it *ForeignContractCreationIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ForeignContractCreation)
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
		it.Event = new(ForeignContractCreation)
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
func (it *ForeignContractCreationIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ForeignContractCreationIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ForeignContractCreation represents a ContractCreation event raised by the Foreign contract.
type ForeignContractCreation struct {
	Owner common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterContractCreation is a free log retrieval operation binding the contract event 0x4db17dd5e4732fb6da34a148104a592783ca119a1e7bb8829eba6cbadef0b511.
//
// Solidity: e ContractCreation(_owner address)
func (_Foreign *ForeignFilterer) FilterContractCreation(opts *bind.FilterOpts) (*ForeignContractCreationIterator, error) {

	logs, sub, err := _Foreign.contract.FilterLogs(opts, "ContractCreation")
	if err != nil {
		return nil, err
	}
	return &ForeignContractCreationIterator{contract: _Foreign.contract, event: "ContractCreation", logs: logs, sub: sub}, nil
}

// WatchContractCreation is a free log subscription operation binding the contract event 0x4db17dd5e4732fb6da34a148104a592783ca119a1e7bb8829eba6cbadef0b511.
//
// Solidity: e ContractCreation(_owner address)
func (_Foreign *ForeignFilterer) WatchContractCreation(opts *bind.WatchOpts, sink chan<- *ForeignContractCreation) (event.Subscription, error) {

	logs, sub, err := _Foreign.contract.WatchLogs(opts, "ContractCreation")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ForeignContractCreation)
				if err := _Foreign.contract.UnpackLog(event, "ContractCreation", log); err != nil {
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

// ForeignDepositIterator is returned from FilterDeposit and is used to iterate over the raw logs and unpacked data for Deposit events raised by the Foreign contract.
type ForeignDepositIterator struct {
	Event *ForeignDeposit // Event containing the contract specifics and raw log

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
func (it *ForeignDepositIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ForeignDeposit)
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
		it.Event = new(ForeignDeposit)
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
func (it *ForeignDepositIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ForeignDepositIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ForeignDeposit represents a Deposit event raised by the Foreign contract.
type ForeignDeposit struct {
	Recipient common.Address
	Value     *big.Int
	ToChain   *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterDeposit is a free log retrieval operation binding the contract event 0x90890809c654f11d6e72a28fa60149770a0d11ec6c92319d6ceb2bb0a4ea1a15.
//
// Solidity: e Deposit(_recipient address, _value uint256, _toChain uint256)
func (_Foreign *ForeignFilterer) FilterDeposit(opts *bind.FilterOpts) (*ForeignDepositIterator, error) {

	logs, sub, err := _Foreign.contract.FilterLogs(opts, "Deposit")
	if err != nil {
		return nil, err
	}
	return &ForeignDepositIterator{contract: _Foreign.contract, event: "Deposit", logs: logs, sub: sub}, nil
}

// WatchDeposit is a free log subscription operation binding the contract event 0x90890809c654f11d6e72a28fa60149770a0d11ec6c92319d6ceb2bb0a4ea1a15.
//
// Solidity: e Deposit(_recipient address, _value uint256, _toChain uint256)
func (_Foreign *ForeignFilterer) WatchDeposit(opts *bind.WatchOpts, sink chan<- *ForeignDeposit) (event.Subscription, error) {

	logs, sub, err := _Foreign.contract.WatchLogs(opts, "Deposit")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ForeignDeposit)
				if err := _Foreign.contract.UnpackLog(event, "Deposit", log); err != nil {
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
