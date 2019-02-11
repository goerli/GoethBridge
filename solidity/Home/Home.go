// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package homecontract

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

// HomeABI is the input ABI used to generate the binding from.
const HomeABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"_recipient\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"},{\"name\":\"_fromChain\",\"type\":\"uint256\"},{\"name\":\"_txHash\",\"type\":\"bytes32\"}],\"name\":\"withdraw\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_addr\",\"type\":\"address\"}],\"name\":\"setBridge\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"bridge\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"fallback\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"ContractCreation\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"_addr\",\"type\":\"address\"}],\"name\":\"BridgeSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"_addr\",\"type\":\"address\"}],\"name\":\"BridgeFunded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"_recipient\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_value\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"_fromChain\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"_txHash\",\"type\":\"bytes32\"}],\"name\":\"Withdraw\",\"type\":\"event\"}]"

// Home is an auto generated Go binding around an Ethereum contract.
type Home struct {
	HomeCaller     // Read-only binding to the contract
	HomeTransactor // Write-only binding to the contract
	HomeFilterer   // Log filterer for contract events
}

// HomeCaller is an auto generated read-only Go binding around an Ethereum contract.
type HomeCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// HomeTransactor is an auto generated write-only Go binding around an Ethereum contract.
type HomeTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// HomeFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type HomeFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// HomeSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type HomeSession struct {
	Contract     *Home             // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// HomeCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type HomeCallerSession struct {
	Contract *HomeCaller   // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// HomeTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type HomeTransactorSession struct {
	Contract     *HomeTransactor   // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// HomeRaw is an auto generated low-level Go binding around an Ethereum contract.
type HomeRaw struct {
	Contract *Home // Generic contract binding to access the raw methods on
}

// HomeCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type HomeCallerRaw struct {
	Contract *HomeCaller // Generic read-only contract binding to access the raw methods on
}

// HomeTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type HomeTransactorRaw struct {
	Contract *HomeTransactor // Generic write-only contract binding to access the raw methods on
}

// NewHome creates a new instance of Home, bound to a specific deployed contract.
func NewHome(address common.Address, backend bind.ContractBackend) (*Home, error) {
	contract, err := bindHome(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Home{HomeCaller: HomeCaller{contract: contract}, HomeTransactor: HomeTransactor{contract: contract}, HomeFilterer: HomeFilterer{contract: contract}}, nil
}

// NewHomeCaller creates a new read-only instance of Home, bound to a specific deployed contract.
func NewHomeCaller(address common.Address, caller bind.ContractCaller) (*HomeCaller, error) {
	contract, err := bindHome(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &HomeCaller{contract: contract}, nil
}

// NewHomeTransactor creates a new write-only instance of Home, bound to a specific deployed contract.
func NewHomeTransactor(address common.Address, transactor bind.ContractTransactor) (*HomeTransactor, error) {
	contract, err := bindHome(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &HomeTransactor{contract: contract}, nil
}

// NewHomeFilterer creates a new log filterer instance of Home, bound to a specific deployed contract.
func NewHomeFilterer(address common.Address, filterer bind.ContractFilterer) (*HomeFilterer, error) {
	contract, err := bindHome(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &HomeFilterer{contract: contract}, nil
}

// bindHome binds a generic wrapper to an already deployed contract.
func bindHome(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(HomeABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Home *HomeRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Home.Contract.HomeCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Home *HomeRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Home.Contract.HomeTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Home *HomeRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Home.Contract.HomeTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Home *HomeCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Home.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Home *HomeTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Home.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Home *HomeTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Home.Contract.contract.Transact(opts, method, params...)
}

// Bridge is a free data retrieval call binding the contract method 0xe78cea92.
//
// Solidity: function bridge() constant returns(address)
func (_Home *HomeCaller) Bridge(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Home.contract.Call(opts, out, "bridge")
	return *ret0, err
}

// Bridge is a free data retrieval call binding the contract method 0xe78cea92.
//
// Solidity: function bridge() constant returns(address)
func (_Home *HomeSession) Bridge() (common.Address, error) {
	return _Home.Contract.Bridge(&_Home.CallOpts)
}

// Bridge is a free data retrieval call binding the contract method 0xe78cea92.
//
// Solidity: function bridge() constant returns(address)
func (_Home *HomeCallerSession) Bridge() (common.Address, error) {
	return _Home.Contract.Bridge(&_Home.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Home *HomeCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Home.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Home *HomeSession) Owner() (common.Address, error) {
	return _Home.Contract.Owner(&_Home.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Home *HomeCallerSession) Owner() (common.Address, error) {
	return _Home.Contract.Owner(&_Home.CallOpts)
}

// SetBridge is a paid mutator transaction binding the contract method 0x8dd14802.
//
// Solidity: function setBridge(_addr address) returns()
func (_Home *HomeTransactor) SetBridge(opts *bind.TransactOpts, _addr common.Address) (*types.Transaction, error) {
	return _Home.contract.Transact(opts, "setBridge", _addr)
}

// SetBridge is a paid mutator transaction binding the contract method 0x8dd14802.
//
// Solidity: function setBridge(_addr address) returns()
func (_Home *HomeSession) SetBridge(_addr common.Address) (*types.Transaction, error) {
	return _Home.Contract.SetBridge(&_Home.TransactOpts, _addr)
}

// SetBridge is a paid mutator transaction binding the contract method 0x8dd14802.
//
// Solidity: function setBridge(_addr address) returns()
func (_Home *HomeTransactorSession) SetBridge(_addr common.Address) (*types.Transaction, error) {
	return _Home.Contract.SetBridge(&_Home.TransactOpts, _addr)
}

// Withdraw is a paid mutator transaction binding the contract method 0x4250a6f3.
//
// Solidity: function withdraw(_recipient address, _value uint256, _fromChain uint256, _txHash bytes32) returns()
func (_Home *HomeTransactor) Withdraw(opts *bind.TransactOpts, _recipient common.Address, _value *big.Int, _fromChain *big.Int, _txHash [32]byte) (*types.Transaction, error) {
	return _Home.contract.Transact(opts, "withdraw", _recipient, _value, _fromChain, _txHash)
}

// Withdraw is a paid mutator transaction binding the contract method 0x4250a6f3.
//
// Solidity: function withdraw(_recipient address, _value uint256, _fromChain uint256, _txHash bytes32) returns()
func (_Home *HomeSession) Withdraw(_recipient common.Address, _value *big.Int, _fromChain *big.Int, _txHash [32]byte) (*types.Transaction, error) {
	return _Home.Contract.Withdraw(&_Home.TransactOpts, _recipient, _value, _fromChain, _txHash)
}

// Withdraw is a paid mutator transaction binding the contract method 0x4250a6f3.
//
// Solidity: function withdraw(_recipient address, _value uint256, _fromChain uint256, _txHash bytes32) returns()
func (_Home *HomeTransactorSession) Withdraw(_recipient common.Address, _value *big.Int, _fromChain *big.Int, _txHash [32]byte) (*types.Transaction, error) {
	return _Home.Contract.Withdraw(&_Home.TransactOpts, _recipient, _value, _fromChain, _txHash)
}

// HomeBridgeFundedIterator is returned from FilterBridgeFunded and is used to iterate over the raw logs and unpacked data for BridgeFunded events raised by the Home contract.
type HomeBridgeFundedIterator struct {
	Event *HomeBridgeFunded // Event containing the contract specifics and raw log

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
func (it *HomeBridgeFundedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HomeBridgeFunded)
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
		it.Event = new(HomeBridgeFunded)
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
func (it *HomeBridgeFundedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *HomeBridgeFundedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// HomeBridgeFunded represents a BridgeFunded event raised by the Home contract.
type HomeBridgeFunded struct {
	Addr common.Address
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterBridgeFunded is a free log retrieval operation binding the contract event 0xc2520f24142cb24b12b04df358be485159ec7ec1a3c3ad25fa65e1a226e4eec3.
//
// Solidity: e BridgeFunded(_addr address)
func (_Home *HomeFilterer) FilterBridgeFunded(opts *bind.FilterOpts) (*HomeBridgeFundedIterator, error) {

	logs, sub, err := _Home.contract.FilterLogs(opts, "BridgeFunded")
	if err != nil {
		return nil, err
	}
	return &HomeBridgeFundedIterator{contract: _Home.contract, event: "BridgeFunded", logs: logs, sub: sub}, nil
}

// WatchBridgeFunded is a free log subscription operation binding the contract event 0xc2520f24142cb24b12b04df358be485159ec7ec1a3c3ad25fa65e1a226e4eec3.
//
// Solidity: e BridgeFunded(_addr address)
func (_Home *HomeFilterer) WatchBridgeFunded(opts *bind.WatchOpts, sink chan<- *HomeBridgeFunded) (event.Subscription, error) {

	logs, sub, err := _Home.contract.WatchLogs(opts, "BridgeFunded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(HomeBridgeFunded)
				if err := _Home.contract.UnpackLog(event, "BridgeFunded", log); err != nil {
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

// HomeBridgeSetIterator is returned from FilterBridgeSet and is used to iterate over the raw logs and unpacked data for BridgeSet events raised by the Home contract.
type HomeBridgeSetIterator struct {
	Event *HomeBridgeSet // Event containing the contract specifics and raw log

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
func (it *HomeBridgeSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HomeBridgeSet)
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
		it.Event = new(HomeBridgeSet)
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
func (it *HomeBridgeSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *HomeBridgeSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// HomeBridgeSet represents a BridgeSet event raised by the Home contract.
type HomeBridgeSet struct {
	Addr common.Address
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterBridgeSet is a free log retrieval operation binding the contract event 0xa49730bff544fd0b716395c592e39c6fd2d2481a19b9229b5b240483db95a495.
//
// Solidity: e BridgeSet(_addr address)
func (_Home *HomeFilterer) FilterBridgeSet(opts *bind.FilterOpts) (*HomeBridgeSetIterator, error) {

	logs, sub, err := _Home.contract.FilterLogs(opts, "BridgeSet")
	if err != nil {
		return nil, err
	}
	return &HomeBridgeSetIterator{contract: _Home.contract, event: "BridgeSet", logs: logs, sub: sub}, nil
}

// WatchBridgeSet is a free log subscription operation binding the contract event 0xa49730bff544fd0b716395c592e39c6fd2d2481a19b9229b5b240483db95a495.
//
// Solidity: e BridgeSet(_addr address)
func (_Home *HomeFilterer) WatchBridgeSet(opts *bind.WatchOpts, sink chan<- *HomeBridgeSet) (event.Subscription, error) {

	logs, sub, err := _Home.contract.WatchLogs(opts, "BridgeSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(HomeBridgeSet)
				if err := _Home.contract.UnpackLog(event, "BridgeSet", log); err != nil {
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

// HomeContractCreationIterator is returned from FilterContractCreation and is used to iterate over the raw logs and unpacked data for ContractCreation events raised by the Home contract.
type HomeContractCreationIterator struct {
	Event *HomeContractCreation // Event containing the contract specifics and raw log

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
func (it *HomeContractCreationIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HomeContractCreation)
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
		it.Event = new(HomeContractCreation)
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
func (it *HomeContractCreationIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *HomeContractCreationIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// HomeContractCreation represents a ContractCreation event raised by the Home contract.
type HomeContractCreation struct {
	Owner common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterContractCreation is a free log retrieval operation binding the contract event 0x4db17dd5e4732fb6da34a148104a592783ca119a1e7bb8829eba6cbadef0b511.
//
// Solidity: e ContractCreation(_owner address)
func (_Home *HomeFilterer) FilterContractCreation(opts *bind.FilterOpts) (*HomeContractCreationIterator, error) {

	logs, sub, err := _Home.contract.FilterLogs(opts, "ContractCreation")
	if err != nil {
		return nil, err
	}
	return &HomeContractCreationIterator{contract: _Home.contract, event: "ContractCreation", logs: logs, sub: sub}, nil
}

// WatchContractCreation is a free log subscription operation binding the contract event 0x4db17dd5e4732fb6da34a148104a592783ca119a1e7bb8829eba6cbadef0b511.
//
// Solidity: e ContractCreation(_owner address)
func (_Home *HomeFilterer) WatchContractCreation(opts *bind.WatchOpts, sink chan<- *HomeContractCreation) (event.Subscription, error) {

	logs, sub, err := _Home.contract.WatchLogs(opts, "ContractCreation")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(HomeContractCreation)
				if err := _Home.contract.UnpackLog(event, "ContractCreation", log); err != nil {
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

// HomeWithdrawIterator is returned from FilterWithdraw and is used to iterate over the raw logs and unpacked data for Withdraw events raised by the Home contract.
type HomeWithdrawIterator struct {
	Event *HomeWithdraw // Event containing the contract specifics and raw log

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
func (it *HomeWithdrawIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HomeWithdraw)
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
		it.Event = new(HomeWithdraw)
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
func (it *HomeWithdrawIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *HomeWithdrawIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// HomeWithdraw represents a Withdraw event raised by the Home contract.
type HomeWithdraw struct {
	Recipient common.Address
	Value     *big.Int
	FromChain *big.Int
	TxHash    [32]byte
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterWithdraw is a free log retrieval operation binding the contract event 0x80bd7e915c065f64ef8e41fa6334d4a2b752591cb4d462c204b8e2449eb0a1f7.
//
// Solidity: e Withdraw(_recipient address, _value uint256, _fromChain uint256, _txHash bytes32)
func (_Home *HomeFilterer) FilterWithdraw(opts *bind.FilterOpts) (*HomeWithdrawIterator, error) {

	logs, sub, err := _Home.contract.FilterLogs(opts, "Withdraw")
	if err != nil {
		return nil, err
	}
	return &HomeWithdrawIterator{contract: _Home.contract, event: "Withdraw", logs: logs, sub: sub}, nil
}

// WatchWithdraw is a free log subscription operation binding the contract event 0x80bd7e915c065f64ef8e41fa6334d4a2b752591cb4d462c204b8e2449eb0a1f7.
//
// Solidity: e Withdraw(_recipient address, _value uint256, _fromChain uint256, _txHash bytes32)
func (_Home *HomeFilterer) WatchWithdraw(opts *bind.WatchOpts, sink chan<- *HomeWithdraw) (event.Subscription, error) {

	logs, sub, err := _Home.contract.WatchLogs(opts, "Withdraw")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(HomeWithdraw)
				if err := _Home.contract.UnpackLog(event, "Withdraw", log); err != nil {
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
