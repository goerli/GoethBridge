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

// TimelockABI is the input ABI used to generate the binding from.
const TimelockABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"_recipient\",\"type\":\"address\"},{\"name\":\"_toChain\",\"type\":\"uint256\"}],\"name\":\"deposit\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"release\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"ContractCreation\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"_recipient\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_value\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"_toChain\",\"type\":\"uint256\"}],\"name\":\"Deposit\",\"type\":\"event\"}]"

// Timelock is an auto generated Go binding around an Ethereum contract.
type Timelock struct {
	TimelockCaller     // Read-only binding to the contract
	TimelockTransactor // Write-only binding to the contract
	TimelockFilterer   // Log filterer for contract events
}

// TimelockCaller is an auto generated read-only Go binding around an Ethereum contract.
type TimelockCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TimelockTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TimelockTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TimelockFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TimelockFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TimelockSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TimelockSession struct {
	Contract     *Timelock         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TimelockCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TimelockCallerSession struct {
	Contract *TimelockCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// TimelockTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TimelockTransactorSession struct {
	Contract     *TimelockTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// TimelockRaw is an auto generated low-level Go binding around an Ethereum contract.
type TimelockRaw struct {
	Contract *Timelock // Generic contract binding to access the raw methods on
}

// TimelockCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TimelockCallerRaw struct {
	Contract *TimelockCaller // Generic read-only contract binding to access the raw methods on
}

// TimelockTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TimelockTransactorRaw struct {
	Contract *TimelockTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTimelock creates a new instance of Timelock, bound to a specific deployed contract.
func NewTimelock(address common.Address, backend bind.ContractBackend) (*Timelock, error) {
	contract, err := bindTimelock(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Timelock{TimelockCaller: TimelockCaller{contract: contract}, TimelockTransactor: TimelockTransactor{contract: contract}, TimelockFilterer: TimelockFilterer{contract: contract}}, nil
}

// NewTimelockCaller creates a new read-only instance of Timelock, bound to a specific deployed contract.
func NewTimelockCaller(address common.Address, caller bind.ContractCaller) (*TimelockCaller, error) {
	contract, err := bindTimelock(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TimelockCaller{contract: contract}, nil
}

// NewTimelockTransactor creates a new write-only instance of Timelock, bound to a specific deployed contract.
func NewTimelockTransactor(address common.Address, transactor bind.ContractTransactor) (*TimelockTransactor, error) {
	contract, err := bindTimelock(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TimelockTransactor{contract: contract}, nil
}

// NewTimelockFilterer creates a new log filterer instance of Timelock, bound to a specific deployed contract.
func NewTimelockFilterer(address common.Address, filterer bind.ContractFilterer) (*TimelockFilterer, error) {
	contract, err := bindTimelock(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TimelockFilterer{contract: contract}, nil
}

// bindTimelock binds a generic wrapper to an already deployed contract.
func bindTimelock(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(TimelockABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Timelock *TimelockRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Timelock.Contract.TimelockCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Timelock *TimelockRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Timelock.Contract.TimelockTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Timelock *TimelockRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Timelock.Contract.TimelockTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Timelock *TimelockCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Timelock.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Timelock *TimelockTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Timelock.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Timelock *TimelockTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Timelock.Contract.contract.Transact(opts, method, params...)
}

// Deposit is a paid mutator transaction binding the contract method 0x47e7ef24.
//
// Solidity: function deposit(_recipient address, _toChain uint256) returns()
func (_Timelock *TimelockTransactor) Deposit(opts *bind.TransactOpts, _recipient common.Address, _toChain *big.Int) (*types.Transaction, error) {
	return _Timelock.contract.Transact(opts, "deposit", _recipient, _toChain)
}

// Deposit is a paid mutator transaction binding the contract method 0x47e7ef24.
//
// Solidity: function deposit(_recipient address, _toChain uint256) returns()
func (_Timelock *TimelockSession) Deposit(_recipient common.Address, _toChain *big.Int) (*types.Transaction, error) {
	return _Timelock.Contract.Deposit(&_Timelock.TransactOpts, _recipient, _toChain)
}

// Deposit is a paid mutator transaction binding the contract method 0x47e7ef24.
//
// Solidity: function deposit(_recipient address, _toChain uint256) returns()
func (_Timelock *TimelockTransactorSession) Deposit(_recipient common.Address, _toChain *big.Int) (*types.Transaction, error) {
	return _Timelock.Contract.Deposit(&_Timelock.TransactOpts, _recipient, _toChain)
}

// Release is a paid mutator transaction binding the contract method 0x86d1a69f.
//
// Solidity: function release() returns()
func (_Timelock *TimelockTransactor) Release(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Timelock.contract.Transact(opts, "release")
}

// Release is a paid mutator transaction binding the contract method 0x86d1a69f.
//
// Solidity: function release() returns()
func (_Timelock *TimelockSession) Release() (*types.Transaction, error) {
	return _Timelock.Contract.Release(&_Timelock.TransactOpts)
}

// Release is a paid mutator transaction binding the contract method 0x86d1a69f.
//
// Solidity: function release() returns()
func (_Timelock *TimelockTransactorSession) Release() (*types.Transaction, error) {
	return _Timelock.Contract.Release(&_Timelock.TransactOpts)
}

// TimelockContractCreationIterator is returned from FilterContractCreation and is used to iterate over the raw logs and unpacked data for ContractCreation events raised by the Timelock contract.
type TimelockContractCreationIterator struct {
	Event *TimelockContractCreation // Event containing the contract specifics and raw log

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
func (it *TimelockContractCreationIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TimelockContractCreation)
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
		it.Event = new(TimelockContractCreation)
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
func (it *TimelockContractCreationIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TimelockContractCreationIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TimelockContractCreation represents a ContractCreation event raised by the Timelock contract.
type TimelockContractCreation struct {
	Owner common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterContractCreation is a free log retrieval operation binding the contract event 0x4db17dd5e4732fb6da34a148104a592783ca119a1e7bb8829eba6cbadef0b511.
//
// Solidity: e ContractCreation(_owner address)
func (_Timelock *TimelockFilterer) FilterContractCreation(opts *bind.FilterOpts) (*TimelockContractCreationIterator, error) {

	logs, sub, err := _Timelock.contract.FilterLogs(opts, "ContractCreation")
	if err != nil {
		return nil, err
	}
	return &TimelockContractCreationIterator{contract: _Timelock.contract, event: "ContractCreation", logs: logs, sub: sub}, nil
}

// WatchContractCreation is a free log subscription operation binding the contract event 0x4db17dd5e4732fb6da34a148104a592783ca119a1e7bb8829eba6cbadef0b511.
//
// Solidity: e ContractCreation(_owner address)
func (_Timelock *TimelockFilterer) WatchContractCreation(opts *bind.WatchOpts, sink chan<- *TimelockContractCreation) (event.Subscription, error) {

	logs, sub, err := _Timelock.contract.WatchLogs(opts, "ContractCreation")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TimelockContractCreation)
				if err := _Timelock.contract.UnpackLog(event, "ContractCreation", log); err != nil {
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

// TimelockDepositIterator is returned from FilterDeposit and is used to iterate over the raw logs and unpacked data for Deposit events raised by the Timelock contract.
type TimelockDepositIterator struct {
	Event *TimelockDeposit // Event containing the contract specifics and raw log

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
func (it *TimelockDepositIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TimelockDeposit)
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
		it.Event = new(TimelockDeposit)
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
func (it *TimelockDepositIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TimelockDepositIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TimelockDeposit represents a Deposit event raised by the Timelock contract.
type TimelockDeposit struct {
	Recipient common.Address
	Value     *big.Int
	ToChain   *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterDeposit is a free log retrieval operation binding the contract event 0x90890809c654f11d6e72a28fa60149770a0d11ec6c92319d6ceb2bb0a4ea1a15.
//
// Solidity: e Deposit(_recipient address, _value uint256, _toChain uint256)
func (_Timelock *TimelockFilterer) FilterDeposit(opts *bind.FilterOpts) (*TimelockDepositIterator, error) {

	logs, sub, err := _Timelock.contract.FilterLogs(opts, "Deposit")
	if err != nil {
		return nil, err
	}
	return &TimelockDepositIterator{contract: _Timelock.contract, event: "Deposit", logs: logs, sub: sub}, nil
}

// WatchDeposit is a free log subscription operation binding the contract event 0x90890809c654f11d6e72a28fa60149770a0d11ec6c92319d6ceb2bb0a4ea1a15.
//
// Solidity: e Deposit(_recipient address, _value uint256, _toChain uint256)
func (_Timelock *TimelockFilterer) WatchDeposit(opts *bind.WatchOpts, sink chan<- *TimelockDeposit) (event.Subscription, error) {

	logs, sub, err := _Timelock.contract.WatchLogs(opts, "Deposit")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TimelockDeposit)
				if err := _Timelock.contract.UnpackLog(event, "Deposit", log); err != nil {
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
