// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package piece

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

// IPiecePieceInfo is an auto generated low-level Go binding around an user-defined struct.
type IPiecePieceInfo struct {
	Owner   common.Address
	Balance *big.Int
	Price   *big.Int
	Size    uint64
	Start   uint64
	Expire  uint64
	Rsn     uint8
	Rsk     uint8
}

// IPieceReplicaInfo is an auto generated low-level Go binding around an user-defined struct.
type IPieceReplicaInfo struct {
	Piece    uint64
	StoredOn common.Address
	Root     [32]byte
}

// IPieceStoreInfo is an auto generated low-level Go binding around an user-defined struct.
type IPieceStoreInfo struct {
	Epoch  uint64
	Count  uint64
	Salary *big.Int
}

// IPieceStoreStat is an auto generated low-level Go binding around an user-defined struct.
type IPieceStoreStat struct {
	Count  uint64
	Salary *big.Int
}

// ContextMetaData contains all meta data concerning the Context contract.
var ContextMetaData = &bind.MetaData{
	ABI: "[]",
}

// ContextABI is the input ABI used to generate the binding from.
// Deprecated: Use ContextMetaData.ABI instead.
var ContextABI = ContextMetaData.ABI

// Context is an auto generated Go binding around an Ethereum contract.
type Context struct {
	ContextCaller     // Read-only binding to the contract
	ContextTransactor // Write-only binding to the contract
	ContextFilterer   // Log filterer for contract events
}

// ContextCaller is an auto generated read-only Go binding around an Ethereum contract.
type ContextCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContextTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ContextTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContextFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ContextFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContextSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ContextSession struct {
	Contract     *Context          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ContextCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ContextCallerSession struct {
	Contract *ContextCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// ContextTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ContextTransactorSession struct {
	Contract     *ContextTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// ContextRaw is an auto generated low-level Go binding around an Ethereum contract.
type ContextRaw struct {
	Contract *Context // Generic contract binding to access the raw methods on
}

// ContextCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ContextCallerRaw struct {
	Contract *ContextCaller // Generic read-only contract binding to access the raw methods on
}

// ContextTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ContextTransactorRaw struct {
	Contract *ContextTransactor // Generic write-only contract binding to access the raw methods on
}

// NewContext creates a new instance of Context, bound to a specific deployed contract.
func NewContext(address common.Address, backend bind.ContractBackend) (*Context, error) {
	contract, err := bindContext(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Context{ContextCaller: ContextCaller{contract: contract}, ContextTransactor: ContextTransactor{contract: contract}, ContextFilterer: ContextFilterer{contract: contract}}, nil
}

// NewContextCaller creates a new read-only instance of Context, bound to a specific deployed contract.
func NewContextCaller(address common.Address, caller bind.ContractCaller) (*ContextCaller, error) {
	contract, err := bindContext(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ContextCaller{contract: contract}, nil
}

// NewContextTransactor creates a new write-only instance of Context, bound to a specific deployed contract.
func NewContextTransactor(address common.Address, transactor bind.ContractTransactor) (*ContextTransactor, error) {
	contract, err := bindContext(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ContextTransactor{contract: contract}, nil
}

// NewContextFilterer creates a new log filterer instance of Context, bound to a specific deployed contract.
func NewContextFilterer(address common.Address, filterer bind.ContractFilterer) (*ContextFilterer, error) {
	contract, err := bindContext(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ContextFilterer{contract: contract}, nil
}

// bindContext binds a generic wrapper to an already deployed contract.
func bindContext(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ContextMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Context *ContextRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Context.Contract.ContextCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Context *ContextRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Context.Contract.ContextTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Context *ContextRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Context.Contract.ContextTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Context *ContextCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Context.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Context *ContextTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Context.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Context *ContextTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Context.Contract.contract.Transact(opts, method, params...)
}

// IBankMetaData contains all meta data concerning the IBank contract.
var IBankMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_m\",\"type\":\"uint256\"}],\"name\":\"Mint\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"_s\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"}],\"name\":\"Set\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_c\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_m\",\"type\":\"uint256\"}],\"name\":\"TransferIn\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_c\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_m\",\"type\":\"uint256\"}],\"name\":\"TransferOut\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_s\",\"type\":\"string\"}],\"name\":\"get\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_a\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_m\",\"type\":\"uint256\"}],\"name\":\"mint\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_s\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"_a\",\"type\":\"address\"}],\"name\":\"set\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_a\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_m\",\"type\":\"uint256\"}],\"name\":\"transferIn\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_a\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_m\",\"type\":\"uint256\"}],\"name\":\"transferOut\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"}]",
}

// IBankABI is the input ABI used to generate the binding from.
// Deprecated: Use IBankMetaData.ABI instead.
var IBankABI = IBankMetaData.ABI

// IBank is an auto generated Go binding around an Ethereum contract.
type IBank struct {
	IBankCaller     // Read-only binding to the contract
	IBankTransactor // Write-only binding to the contract
	IBankFilterer   // Log filterer for contract events
}

// IBankCaller is an auto generated read-only Go binding around an Ethereum contract.
type IBankCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IBankTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IBankTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IBankFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IBankFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IBankSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IBankSession struct {
	Contract     *IBank            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IBankCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IBankCallerSession struct {
	Contract *IBankCaller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// IBankTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IBankTransactorSession struct {
	Contract     *IBankTransactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IBankRaw is an auto generated low-level Go binding around an Ethereum contract.
type IBankRaw struct {
	Contract *IBank // Generic contract binding to access the raw methods on
}

// IBankCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IBankCallerRaw struct {
	Contract *IBankCaller // Generic read-only contract binding to access the raw methods on
}

// IBankTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IBankTransactorRaw struct {
	Contract *IBankTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIBank creates a new instance of IBank, bound to a specific deployed contract.
func NewIBank(address common.Address, backend bind.ContractBackend) (*IBank, error) {
	contract, err := bindIBank(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IBank{IBankCaller: IBankCaller{contract: contract}, IBankTransactor: IBankTransactor{contract: contract}, IBankFilterer: IBankFilterer{contract: contract}}, nil
}

// NewIBankCaller creates a new read-only instance of IBank, bound to a specific deployed contract.
func NewIBankCaller(address common.Address, caller bind.ContractCaller) (*IBankCaller, error) {
	contract, err := bindIBank(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IBankCaller{contract: contract}, nil
}

// NewIBankTransactor creates a new write-only instance of IBank, bound to a specific deployed contract.
func NewIBankTransactor(address common.Address, transactor bind.ContractTransactor) (*IBankTransactor, error) {
	contract, err := bindIBank(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IBankTransactor{contract: contract}, nil
}

// NewIBankFilterer creates a new log filterer instance of IBank, bound to a specific deployed contract.
func NewIBankFilterer(address common.Address, filterer bind.ContractFilterer) (*IBankFilterer, error) {
	contract, err := bindIBank(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IBankFilterer{contract: contract}, nil
}

// bindIBank binds a generic wrapper to an already deployed contract.
func bindIBank(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := IBankMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IBank *IBankRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IBank.Contract.IBankCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IBank *IBankRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IBank.Contract.IBankTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IBank *IBankRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IBank.Contract.IBankTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IBank *IBankCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IBank.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IBank *IBankTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IBank.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IBank *IBankTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IBank.Contract.contract.Transact(opts, method, params...)
}

// Get is a paid mutator transaction binding the contract method 0x693ec85e.
//
// Solidity: function get(string _s) returns(address)
func (_IBank *IBankTransactor) Get(opts *bind.TransactOpts, _s string) (*types.Transaction, error) {
	return _IBank.contract.Transact(opts, "get", _s)
}

// Get is a paid mutator transaction binding the contract method 0x693ec85e.
//
// Solidity: function get(string _s) returns(address)
func (_IBank *IBankSession) Get(_s string) (*types.Transaction, error) {
	return _IBank.Contract.Get(&_IBank.TransactOpts, _s)
}

// Get is a paid mutator transaction binding the contract method 0x693ec85e.
//
// Solidity: function get(string _s) returns(address)
func (_IBank *IBankTransactorSession) Get(_s string) (*types.Transaction, error) {
	return _IBank.Contract.Get(&_IBank.TransactOpts, _s)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address _a, uint256 _m) payable returns()
func (_IBank *IBankTransactor) Mint(opts *bind.TransactOpts, _a common.Address, _m *big.Int) (*types.Transaction, error) {
	return _IBank.contract.Transact(opts, "mint", _a, _m)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address _a, uint256 _m) payable returns()
func (_IBank *IBankSession) Mint(_a common.Address, _m *big.Int) (*types.Transaction, error) {
	return _IBank.Contract.Mint(&_IBank.TransactOpts, _a, _m)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address _a, uint256 _m) payable returns()
func (_IBank *IBankTransactorSession) Mint(_a common.Address, _m *big.Int) (*types.Transaction, error) {
	return _IBank.Contract.Mint(&_IBank.TransactOpts, _a, _m)
}

// Set is a paid mutator transaction binding the contract method 0xa815ff15.
//
// Solidity: function set(string _s, address _a) returns()
func (_IBank *IBankTransactor) Set(opts *bind.TransactOpts, _s string, _a common.Address) (*types.Transaction, error) {
	return _IBank.contract.Transact(opts, "set", _s, _a)
}

// Set is a paid mutator transaction binding the contract method 0xa815ff15.
//
// Solidity: function set(string _s, address _a) returns()
func (_IBank *IBankSession) Set(_s string, _a common.Address) (*types.Transaction, error) {
	return _IBank.Contract.Set(&_IBank.TransactOpts, _s, _a)
}

// Set is a paid mutator transaction binding the contract method 0xa815ff15.
//
// Solidity: function set(string _s, address _a) returns()
func (_IBank *IBankTransactorSession) Set(_s string, _a common.Address) (*types.Transaction, error) {
	return _IBank.Contract.Set(&_IBank.TransactOpts, _s, _a)
}

// TransferIn is a paid mutator transaction binding the contract method 0xe8888915.
//
// Solidity: function transferIn(address _a, uint256 _m) payable returns()
func (_IBank *IBankTransactor) TransferIn(opts *bind.TransactOpts, _a common.Address, _m *big.Int) (*types.Transaction, error) {
	return _IBank.contract.Transact(opts, "transferIn", _a, _m)
}

// TransferIn is a paid mutator transaction binding the contract method 0xe8888915.
//
// Solidity: function transferIn(address _a, uint256 _m) payable returns()
func (_IBank *IBankSession) TransferIn(_a common.Address, _m *big.Int) (*types.Transaction, error) {
	return _IBank.Contract.TransferIn(&_IBank.TransactOpts, _a, _m)
}

// TransferIn is a paid mutator transaction binding the contract method 0xe8888915.
//
// Solidity: function transferIn(address _a, uint256 _m) payable returns()
func (_IBank *IBankTransactorSession) TransferIn(_a common.Address, _m *big.Int) (*types.Transaction, error) {
	return _IBank.Contract.TransferIn(&_IBank.TransactOpts, _a, _m)
}

// TransferOut is a paid mutator transaction binding the contract method 0x76890c58.
//
// Solidity: function transferOut(address _a, uint256 _m) payable returns()
func (_IBank *IBankTransactor) TransferOut(opts *bind.TransactOpts, _a common.Address, _m *big.Int) (*types.Transaction, error) {
	return _IBank.contract.Transact(opts, "transferOut", _a, _m)
}

// TransferOut is a paid mutator transaction binding the contract method 0x76890c58.
//
// Solidity: function transferOut(address _a, uint256 _m) payable returns()
func (_IBank *IBankSession) TransferOut(_a common.Address, _m *big.Int) (*types.Transaction, error) {
	return _IBank.Contract.TransferOut(&_IBank.TransactOpts, _a, _m)
}

// TransferOut is a paid mutator transaction binding the contract method 0x76890c58.
//
// Solidity: function transferOut(address _a, uint256 _m) payable returns()
func (_IBank *IBankTransactorSession) TransferOut(_a common.Address, _m *big.Int) (*types.Transaction, error) {
	return _IBank.Contract.TransferOut(&_IBank.TransactOpts, _a, _m)
}

// IBankMintIterator is returned from FilterMint and is used to iterate over the raw logs and unpacked data for Mint events raised by the IBank contract.
type IBankMintIterator struct {
	Event *IBankMint // Event containing the contract specifics and raw log

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
func (it *IBankMintIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IBankMint)
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
		it.Event = new(IBankMint)
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
func (it *IBankMintIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IBankMintIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IBankMint represents a Mint event raised by the IBank contract.
type IBankMint struct {
	To  common.Address
	M   *big.Int
	Raw types.Log // Blockchain specific contextual infos
}

// FilterMint is a free log retrieval operation binding the contract event 0x0f6798a560793a54c3bcfe86a93cde1e73087d944c0ea20544137d4121396885.
//
// Solidity: event Mint(address indexed _to, uint256 _m)
func (_IBank *IBankFilterer) FilterMint(opts *bind.FilterOpts, _to []common.Address) (*IBankMintIterator, error) {

	var _toRule []interface{}
	for _, _toItem := range _to {
		_toRule = append(_toRule, _toItem)
	}

	logs, sub, err := _IBank.contract.FilterLogs(opts, "Mint", _toRule)
	if err != nil {
		return nil, err
	}
	return &IBankMintIterator{contract: _IBank.contract, event: "Mint", logs: logs, sub: sub}, nil
}

// WatchMint is a free log subscription operation binding the contract event 0x0f6798a560793a54c3bcfe86a93cde1e73087d944c0ea20544137d4121396885.
//
// Solidity: event Mint(address indexed _to, uint256 _m)
func (_IBank *IBankFilterer) WatchMint(opts *bind.WatchOpts, sink chan<- *IBankMint, _to []common.Address) (event.Subscription, error) {

	var _toRule []interface{}
	for _, _toItem := range _to {
		_toRule = append(_toRule, _toItem)
	}

	logs, sub, err := _IBank.contract.WatchLogs(opts, "Mint", _toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IBankMint)
				if err := _IBank.contract.UnpackLog(event, "Mint", log); err != nil {
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

// ParseMint is a log parse operation binding the contract event 0x0f6798a560793a54c3bcfe86a93cde1e73087d944c0ea20544137d4121396885.
//
// Solidity: event Mint(address indexed _to, uint256 _m)
func (_IBank *IBankFilterer) ParseMint(log types.Log) (*IBankMint, error) {
	event := new(IBankMint)
	if err := _IBank.contract.UnpackLog(event, "Mint", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IBankSetIterator is returned from FilterSet and is used to iterate over the raw logs and unpacked data for Set events raised by the IBank contract.
type IBankSetIterator struct {
	Event *IBankSet // Event containing the contract specifics and raw log

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
func (it *IBankSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IBankSet)
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
		it.Event = new(IBankSet)
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
func (it *IBankSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IBankSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IBankSet represents a Set event raised by the IBank contract.
type IBankSet struct {
	S   string
	To  common.Address
	Raw types.Log // Blockchain specific contextual infos
}

// FilterSet is a free log retrieval operation binding the contract event 0x496595ced95720268cf8bc60bae3f35024ff2a130f73ac4e20f5c1eaca35db99.
//
// Solidity: event Set(string _s, address _to)
func (_IBank *IBankFilterer) FilterSet(opts *bind.FilterOpts) (*IBankSetIterator, error) {

	logs, sub, err := _IBank.contract.FilterLogs(opts, "Set")
	if err != nil {
		return nil, err
	}
	return &IBankSetIterator{contract: _IBank.contract, event: "Set", logs: logs, sub: sub}, nil
}

// WatchSet is a free log subscription operation binding the contract event 0x496595ced95720268cf8bc60bae3f35024ff2a130f73ac4e20f5c1eaca35db99.
//
// Solidity: event Set(string _s, address _to)
func (_IBank *IBankFilterer) WatchSet(opts *bind.WatchOpts, sink chan<- *IBankSet) (event.Subscription, error) {

	logs, sub, err := _IBank.contract.WatchLogs(opts, "Set")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IBankSet)
				if err := _IBank.contract.UnpackLog(event, "Set", log); err != nil {
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

// ParseSet is a log parse operation binding the contract event 0x496595ced95720268cf8bc60bae3f35024ff2a130f73ac4e20f5c1eaca35db99.
//
// Solidity: event Set(string _s, address _to)
func (_IBank *IBankFilterer) ParseSet(log types.Log) (*IBankSet, error) {
	event := new(IBankSet)
	if err := _IBank.contract.UnpackLog(event, "Set", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IBankTransferInIterator is returned from FilterTransferIn and is used to iterate over the raw logs and unpacked data for TransferIn events raised by the IBank contract.
type IBankTransferInIterator struct {
	Event *IBankTransferIn // Event containing the contract specifics and raw log

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
func (it *IBankTransferInIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IBankTransferIn)
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
		it.Event = new(IBankTransferIn)
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
func (it *IBankTransferInIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IBankTransferInIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IBankTransferIn represents a TransferIn event raised by the IBank contract.
type IBankTransferIn struct {
	C    common.Address
	From common.Address
	M    *big.Int
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterTransferIn is a free log retrieval operation binding the contract event 0x8ab008cb8444c6d8f2fa3dc0b159b5a86c62b00092219a6c69757805cbf7bde7.
//
// Solidity: event TransferIn(address indexed _c, address indexed _from, uint256 _m)
func (_IBank *IBankFilterer) FilterTransferIn(opts *bind.FilterOpts, _c []common.Address, _from []common.Address) (*IBankTransferInIterator, error) {

	var _cRule []interface{}
	for _, _cItem := range _c {
		_cRule = append(_cRule, _cItem)
	}
	var _fromRule []interface{}
	for _, _fromItem := range _from {
		_fromRule = append(_fromRule, _fromItem)
	}

	logs, sub, err := _IBank.contract.FilterLogs(opts, "TransferIn", _cRule, _fromRule)
	if err != nil {
		return nil, err
	}
	return &IBankTransferInIterator{contract: _IBank.contract, event: "TransferIn", logs: logs, sub: sub}, nil
}

// WatchTransferIn is a free log subscription operation binding the contract event 0x8ab008cb8444c6d8f2fa3dc0b159b5a86c62b00092219a6c69757805cbf7bde7.
//
// Solidity: event TransferIn(address indexed _c, address indexed _from, uint256 _m)
func (_IBank *IBankFilterer) WatchTransferIn(opts *bind.WatchOpts, sink chan<- *IBankTransferIn, _c []common.Address, _from []common.Address) (event.Subscription, error) {

	var _cRule []interface{}
	for _, _cItem := range _c {
		_cRule = append(_cRule, _cItem)
	}
	var _fromRule []interface{}
	for _, _fromItem := range _from {
		_fromRule = append(_fromRule, _fromItem)
	}

	logs, sub, err := _IBank.contract.WatchLogs(opts, "TransferIn", _cRule, _fromRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IBankTransferIn)
				if err := _IBank.contract.UnpackLog(event, "TransferIn", log); err != nil {
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

// ParseTransferIn is a log parse operation binding the contract event 0x8ab008cb8444c6d8f2fa3dc0b159b5a86c62b00092219a6c69757805cbf7bde7.
//
// Solidity: event TransferIn(address indexed _c, address indexed _from, uint256 _m)
func (_IBank *IBankFilterer) ParseTransferIn(log types.Log) (*IBankTransferIn, error) {
	event := new(IBankTransferIn)
	if err := _IBank.contract.UnpackLog(event, "TransferIn", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IBankTransferOutIterator is returned from FilterTransferOut and is used to iterate over the raw logs and unpacked data for TransferOut events raised by the IBank contract.
type IBankTransferOutIterator struct {
	Event *IBankTransferOut // Event containing the contract specifics and raw log

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
func (it *IBankTransferOutIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IBankTransferOut)
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
		it.Event = new(IBankTransferOut)
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
func (it *IBankTransferOutIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IBankTransferOutIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IBankTransferOut represents a TransferOut event raised by the IBank contract.
type IBankTransferOut struct {
	C   common.Address
	To  common.Address
	M   *big.Int
	Raw types.Log // Blockchain specific contextual infos
}

// FilterTransferOut is a free log retrieval operation binding the contract event 0x5d2c285d7e86ec84b7a404ba6d7627ca0a71709acce9c1cc05fada583e3ded81.
//
// Solidity: event TransferOut(address indexed _c, address indexed _to, uint256 _m)
func (_IBank *IBankFilterer) FilterTransferOut(opts *bind.FilterOpts, _c []common.Address, _to []common.Address) (*IBankTransferOutIterator, error) {

	var _cRule []interface{}
	for _, _cItem := range _c {
		_cRule = append(_cRule, _cItem)
	}
	var _toRule []interface{}
	for _, _toItem := range _to {
		_toRule = append(_toRule, _toItem)
	}

	logs, sub, err := _IBank.contract.FilterLogs(opts, "TransferOut", _cRule, _toRule)
	if err != nil {
		return nil, err
	}
	return &IBankTransferOutIterator{contract: _IBank.contract, event: "TransferOut", logs: logs, sub: sub}, nil
}

// WatchTransferOut is a free log subscription operation binding the contract event 0x5d2c285d7e86ec84b7a404ba6d7627ca0a71709acce9c1cc05fada583e3ded81.
//
// Solidity: event TransferOut(address indexed _c, address indexed _to, uint256 _m)
func (_IBank *IBankFilterer) WatchTransferOut(opts *bind.WatchOpts, sink chan<- *IBankTransferOut, _c []common.Address, _to []common.Address) (event.Subscription, error) {

	var _cRule []interface{}
	for _, _cItem := range _c {
		_cRule = append(_cRule, _cItem)
	}
	var _toRule []interface{}
	for _, _toItem := range _to {
		_toRule = append(_toRule, _toItem)
	}

	logs, sub, err := _IBank.contract.WatchLogs(opts, "TransferOut", _cRule, _toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IBankTransferOut)
				if err := _IBank.contract.UnpackLog(event, "TransferOut", log); err != nil {
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

// ParseTransferOut is a log parse operation binding the contract event 0x5d2c285d7e86ec84b7a404ba6d7627ca0a71709acce9c1cc05fada583e3ded81.
//
// Solidity: event TransferOut(address indexed _c, address indexed _to, uint256 _m)
func (_IBank *IBankFilterer) ParseTransferOut(log types.Log) (*IBankTransferOut, error) {
	event := new(IBankTransferOut)
	if err := _IBank.contract.UnpackLog(event, "TransferOut", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IControlMetaData contains all meta data concerning the IControl contract.
var IControlMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_a\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"_ep\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"_m\",\"type\":\"uint256\"}],\"name\":\"add\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_a\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_money\",\"type\":\"uint256\"}],\"name\":\"checkNode\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_a\",\"type\":\"address\"},{\"internalType\":\"uint8\",\"name\":\"_typ\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"_m\",\"type\":\"uint256\"}],\"name\":\"lock\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_a\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_m\",\"type\":\"uint256\"}],\"name\":\"mint\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_a\",\"type\":\"address\"},{\"internalType\":\"uint8\",\"name\":\"_typ\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_m\",\"type\":\"uint256\"}],\"name\":\"punish\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_a\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_m\",\"type\":\"uint256\"}],\"name\":\"release\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_a\",\"type\":\"address\"},{\"internalType\":\"uint8\",\"name\":\"_typ\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"_m\",\"type\":\"uint256\"}],\"name\":\"unlock\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// IControlABI is the input ABI used to generate the binding from.
// Deprecated: Use IControlMetaData.ABI instead.
var IControlABI = IControlMetaData.ABI

// IControl is an auto generated Go binding around an Ethereum contract.
type IControl struct {
	IControlCaller     // Read-only binding to the contract
	IControlTransactor // Write-only binding to the contract
	IControlFilterer   // Log filterer for contract events
}

// IControlCaller is an auto generated read-only Go binding around an Ethereum contract.
type IControlCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IControlTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IControlTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IControlFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IControlFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IControlSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IControlSession struct {
	Contract     *IControl         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IControlCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IControlCallerSession struct {
	Contract *IControlCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// IControlTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IControlTransactorSession struct {
	Contract     *IControlTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// IControlRaw is an auto generated low-level Go binding around an Ethereum contract.
type IControlRaw struct {
	Contract *IControl // Generic contract binding to access the raw methods on
}

// IControlCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IControlCallerRaw struct {
	Contract *IControlCaller // Generic read-only contract binding to access the raw methods on
}

// IControlTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IControlTransactorRaw struct {
	Contract *IControlTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIControl creates a new instance of IControl, bound to a specific deployed contract.
func NewIControl(address common.Address, backend bind.ContractBackend) (*IControl, error) {
	contract, err := bindIControl(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IControl{IControlCaller: IControlCaller{contract: contract}, IControlTransactor: IControlTransactor{contract: contract}, IControlFilterer: IControlFilterer{contract: contract}}, nil
}

// NewIControlCaller creates a new read-only instance of IControl, bound to a specific deployed contract.
func NewIControlCaller(address common.Address, caller bind.ContractCaller) (*IControlCaller, error) {
	contract, err := bindIControl(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IControlCaller{contract: contract}, nil
}

// NewIControlTransactor creates a new write-only instance of IControl, bound to a specific deployed contract.
func NewIControlTransactor(address common.Address, transactor bind.ContractTransactor) (*IControlTransactor, error) {
	contract, err := bindIControl(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IControlTransactor{contract: contract}, nil
}

// NewIControlFilterer creates a new log filterer instance of IControl, bound to a specific deployed contract.
func NewIControlFilterer(address common.Address, filterer bind.ContractFilterer) (*IControlFilterer, error) {
	contract, err := bindIControl(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IControlFilterer{contract: contract}, nil
}

// bindIControl binds a generic wrapper to an already deployed contract.
func bindIControl(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := IControlMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IControl *IControlRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IControl.Contract.IControlCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IControl *IControlRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IControl.Contract.IControlTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IControl *IControlRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IControl.Contract.IControlTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IControl *IControlCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IControl.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IControl *IControlTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IControl.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IControl *IControlTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IControl.Contract.contract.Transact(opts, method, params...)
}

// Add is a paid mutator transaction binding the contract method 0xdae448b9.
//
// Solidity: function add(address _a, uint64 _ep, uint256 _m) returns()
func (_IControl *IControlTransactor) Add(opts *bind.TransactOpts, _a common.Address, _ep uint64, _m *big.Int) (*types.Transaction, error) {
	return _IControl.contract.Transact(opts, "add", _a, _ep, _m)
}

// Add is a paid mutator transaction binding the contract method 0xdae448b9.
//
// Solidity: function add(address _a, uint64 _ep, uint256 _m) returns()
func (_IControl *IControlSession) Add(_a common.Address, _ep uint64, _m *big.Int) (*types.Transaction, error) {
	return _IControl.Contract.Add(&_IControl.TransactOpts, _a, _ep, _m)
}

// Add is a paid mutator transaction binding the contract method 0xdae448b9.
//
// Solidity: function add(address _a, uint64 _ep, uint256 _m) returns()
func (_IControl *IControlTransactorSession) Add(_a common.Address, _ep uint64, _m *big.Int) (*types.Transaction, error) {
	return _IControl.Contract.Add(&_IControl.TransactOpts, _a, _ep, _m)
}

// CheckNode is a paid mutator transaction binding the contract method 0xa3f60869.
//
// Solidity: function checkNode(address _a, uint256 _money) returns()
func (_IControl *IControlTransactor) CheckNode(opts *bind.TransactOpts, _a common.Address, _money *big.Int) (*types.Transaction, error) {
	return _IControl.contract.Transact(opts, "checkNode", _a, _money)
}

// CheckNode is a paid mutator transaction binding the contract method 0xa3f60869.
//
// Solidity: function checkNode(address _a, uint256 _money) returns()
func (_IControl *IControlSession) CheckNode(_a common.Address, _money *big.Int) (*types.Transaction, error) {
	return _IControl.Contract.CheckNode(&_IControl.TransactOpts, _a, _money)
}

// CheckNode is a paid mutator transaction binding the contract method 0xa3f60869.
//
// Solidity: function checkNode(address _a, uint256 _money) returns()
func (_IControl *IControlTransactorSession) CheckNode(_a common.Address, _money *big.Int) (*types.Transaction, error) {
	return _IControl.Contract.CheckNode(&_IControl.TransactOpts, _a, _money)
}

// Lock is a paid mutator transaction binding the contract method 0x738dddba.
//
// Solidity: function lock(address _a, uint8 _typ, uint256 _m) returns()
func (_IControl *IControlTransactor) Lock(opts *bind.TransactOpts, _a common.Address, _typ uint8, _m *big.Int) (*types.Transaction, error) {
	return _IControl.contract.Transact(opts, "lock", _a, _typ, _m)
}

// Lock is a paid mutator transaction binding the contract method 0x738dddba.
//
// Solidity: function lock(address _a, uint8 _typ, uint256 _m) returns()
func (_IControl *IControlSession) Lock(_a common.Address, _typ uint8, _m *big.Int) (*types.Transaction, error) {
	return _IControl.Contract.Lock(&_IControl.TransactOpts, _a, _typ, _m)
}

// Lock is a paid mutator transaction binding the contract method 0x738dddba.
//
// Solidity: function lock(address _a, uint8 _typ, uint256 _m) returns()
func (_IControl *IControlTransactorSession) Lock(_a common.Address, _typ uint8, _m *big.Int) (*types.Transaction, error) {
	return _IControl.Contract.Lock(&_IControl.TransactOpts, _a, _typ, _m)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address _a, uint256 _m) returns()
func (_IControl *IControlTransactor) Mint(opts *bind.TransactOpts, _a common.Address, _m *big.Int) (*types.Transaction, error) {
	return _IControl.contract.Transact(opts, "mint", _a, _m)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address _a, uint256 _m) returns()
func (_IControl *IControlSession) Mint(_a common.Address, _m *big.Int) (*types.Transaction, error) {
	return _IControl.Contract.Mint(&_IControl.TransactOpts, _a, _m)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address _a, uint256 _m) returns()
func (_IControl *IControlTransactorSession) Mint(_a common.Address, _m *big.Int) (*types.Transaction, error) {
	return _IControl.Contract.Mint(&_IControl.TransactOpts, _a, _m)
}

// Punish is a paid mutator transaction binding the contract method 0x06f0b4f1.
//
// Solidity: function punish(address _a, uint8 _typ, address _to, uint256 _m) returns()
func (_IControl *IControlTransactor) Punish(opts *bind.TransactOpts, _a common.Address, _typ uint8, _to common.Address, _m *big.Int) (*types.Transaction, error) {
	return _IControl.contract.Transact(opts, "punish", _a, _typ, _to, _m)
}

// Punish is a paid mutator transaction binding the contract method 0x06f0b4f1.
//
// Solidity: function punish(address _a, uint8 _typ, address _to, uint256 _m) returns()
func (_IControl *IControlSession) Punish(_a common.Address, _typ uint8, _to common.Address, _m *big.Int) (*types.Transaction, error) {
	return _IControl.Contract.Punish(&_IControl.TransactOpts, _a, _typ, _to, _m)
}

// Punish is a paid mutator transaction binding the contract method 0x06f0b4f1.
//
// Solidity: function punish(address _a, uint8 _typ, address _to, uint256 _m) returns()
func (_IControl *IControlTransactorSession) Punish(_a common.Address, _typ uint8, _to common.Address, _m *big.Int) (*types.Transaction, error) {
	return _IControl.Contract.Punish(&_IControl.TransactOpts, _a, _typ, _to, _m)
}

// Release is a paid mutator transaction binding the contract method 0x0357371d.
//
// Solidity: function release(address _a, uint256 _m) returns()
func (_IControl *IControlTransactor) Release(opts *bind.TransactOpts, _a common.Address, _m *big.Int) (*types.Transaction, error) {
	return _IControl.contract.Transact(opts, "release", _a, _m)
}

// Release is a paid mutator transaction binding the contract method 0x0357371d.
//
// Solidity: function release(address _a, uint256 _m) returns()
func (_IControl *IControlSession) Release(_a common.Address, _m *big.Int) (*types.Transaction, error) {
	return _IControl.Contract.Release(&_IControl.TransactOpts, _a, _m)
}

// Release is a paid mutator transaction binding the contract method 0x0357371d.
//
// Solidity: function release(address _a, uint256 _m) returns()
func (_IControl *IControlTransactorSession) Release(_a common.Address, _m *big.Int) (*types.Transaction, error) {
	return _IControl.Contract.Release(&_IControl.TransactOpts, _a, _m)
}

// Unlock is a paid mutator transaction binding the contract method 0x2165e20b.
//
// Solidity: function unlock(address _a, uint8 _typ, uint256 _m) returns()
func (_IControl *IControlTransactor) Unlock(opts *bind.TransactOpts, _a common.Address, _typ uint8, _m *big.Int) (*types.Transaction, error) {
	return _IControl.contract.Transact(opts, "unlock", _a, _typ, _m)
}

// Unlock is a paid mutator transaction binding the contract method 0x2165e20b.
//
// Solidity: function unlock(address _a, uint8 _typ, uint256 _m) returns()
func (_IControl *IControlSession) Unlock(_a common.Address, _typ uint8, _m *big.Int) (*types.Transaction, error) {
	return _IControl.Contract.Unlock(&_IControl.TransactOpts, _a, _typ, _m)
}

// Unlock is a paid mutator transaction binding the contract method 0x2165e20b.
//
// Solidity: function unlock(address _a, uint8 _typ, uint256 _m) returns()
func (_IControl *IControlTransactorSession) Unlock(_a common.Address, _typ uint8, _m *big.Int) (*types.Transaction, error) {
	return _IControl.Contract.Unlock(&_IControl.TransactOpts, _a, _typ, _m)
}

// IEpochMetaData contains all meta data concerning the IEpoch contract.
var IEpochMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"_e\",\"type\":\"uint64\"}],\"name\":\"SetEpoch\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"check\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"current\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"_epoch\",\"type\":\"uint64\"}],\"name\":\"getEpoch\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// IEpochABI is the input ABI used to generate the binding from.
// Deprecated: Use IEpochMetaData.ABI instead.
var IEpochABI = IEpochMetaData.ABI

// IEpoch is an auto generated Go binding around an Ethereum contract.
type IEpoch struct {
	IEpochCaller     // Read-only binding to the contract
	IEpochTransactor // Write-only binding to the contract
	IEpochFilterer   // Log filterer for contract events
}

// IEpochCaller is an auto generated read-only Go binding around an Ethereum contract.
type IEpochCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IEpochTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IEpochTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IEpochFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IEpochFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IEpochSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IEpochSession struct {
	Contract     *IEpoch           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IEpochCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IEpochCallerSession struct {
	Contract *IEpochCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// IEpochTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IEpochTransactorSession struct {
	Contract     *IEpochTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IEpochRaw is an auto generated low-level Go binding around an Ethereum contract.
type IEpochRaw struct {
	Contract *IEpoch // Generic contract binding to access the raw methods on
}

// IEpochCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IEpochCallerRaw struct {
	Contract *IEpochCaller // Generic read-only contract binding to access the raw methods on
}

// IEpochTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IEpochTransactorRaw struct {
	Contract *IEpochTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIEpoch creates a new instance of IEpoch, bound to a specific deployed contract.
func NewIEpoch(address common.Address, backend bind.ContractBackend) (*IEpoch, error) {
	contract, err := bindIEpoch(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IEpoch{IEpochCaller: IEpochCaller{contract: contract}, IEpochTransactor: IEpochTransactor{contract: contract}, IEpochFilterer: IEpochFilterer{contract: contract}}, nil
}

// NewIEpochCaller creates a new read-only instance of IEpoch, bound to a specific deployed contract.
func NewIEpochCaller(address common.Address, caller bind.ContractCaller) (*IEpochCaller, error) {
	contract, err := bindIEpoch(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IEpochCaller{contract: contract}, nil
}

// NewIEpochTransactor creates a new write-only instance of IEpoch, bound to a specific deployed contract.
func NewIEpochTransactor(address common.Address, transactor bind.ContractTransactor) (*IEpochTransactor, error) {
	contract, err := bindIEpoch(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IEpochTransactor{contract: contract}, nil
}

// NewIEpochFilterer creates a new log filterer instance of IEpoch, bound to a specific deployed contract.
func NewIEpochFilterer(address common.Address, filterer bind.ContractFilterer) (*IEpochFilterer, error) {
	contract, err := bindIEpoch(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IEpochFilterer{contract: contract}, nil
}

// bindIEpoch binds a generic wrapper to an already deployed contract.
func bindIEpoch(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := IEpochMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IEpoch *IEpochRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IEpoch.Contract.IEpochCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IEpoch *IEpochRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IEpoch.Contract.IEpochTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IEpoch *IEpochRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IEpoch.Contract.IEpochTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IEpoch *IEpochCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IEpoch.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IEpoch *IEpochTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IEpoch.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IEpoch *IEpochTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IEpoch.Contract.contract.Transact(opts, method, params...)
}

// Check is a paid mutator transaction binding the contract method 0x919840ad.
//
// Solidity: function check() returns()
func (_IEpoch *IEpochTransactor) Check(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IEpoch.contract.Transact(opts, "check")
}

// Check is a paid mutator transaction binding the contract method 0x919840ad.
//
// Solidity: function check() returns()
func (_IEpoch *IEpochSession) Check() (*types.Transaction, error) {
	return _IEpoch.Contract.Check(&_IEpoch.TransactOpts)
}

// Check is a paid mutator transaction binding the contract method 0x919840ad.
//
// Solidity: function check() returns()
func (_IEpoch *IEpochTransactorSession) Check() (*types.Transaction, error) {
	return _IEpoch.Contract.Check(&_IEpoch.TransactOpts)
}

// Current is a paid mutator transaction binding the contract method 0x9fa6a6e3.
//
// Solidity: function current() returns(uint64)
func (_IEpoch *IEpochTransactor) Current(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IEpoch.contract.Transact(opts, "current")
}

// Current is a paid mutator transaction binding the contract method 0x9fa6a6e3.
//
// Solidity: function current() returns(uint64)
func (_IEpoch *IEpochSession) Current() (*types.Transaction, error) {
	return _IEpoch.Contract.Current(&_IEpoch.TransactOpts)
}

// Current is a paid mutator transaction binding the contract method 0x9fa6a6e3.
//
// Solidity: function current() returns(uint64)
func (_IEpoch *IEpochTransactorSession) Current() (*types.Transaction, error) {
	return _IEpoch.Contract.Current(&_IEpoch.TransactOpts)
}

// GetEpoch is a paid mutator transaction binding the contract method 0x12a02c82.
//
// Solidity: function getEpoch(uint64 _epoch) returns(uint256, bytes32)
func (_IEpoch *IEpochTransactor) GetEpoch(opts *bind.TransactOpts, _epoch uint64) (*types.Transaction, error) {
	return _IEpoch.contract.Transact(opts, "getEpoch", _epoch)
}

// GetEpoch is a paid mutator transaction binding the contract method 0x12a02c82.
//
// Solidity: function getEpoch(uint64 _epoch) returns(uint256, bytes32)
func (_IEpoch *IEpochSession) GetEpoch(_epoch uint64) (*types.Transaction, error) {
	return _IEpoch.Contract.GetEpoch(&_IEpoch.TransactOpts, _epoch)
}

// GetEpoch is a paid mutator transaction binding the contract method 0x12a02c82.
//
// Solidity: function getEpoch(uint64 _epoch) returns(uint256, bytes32)
func (_IEpoch *IEpochTransactorSession) GetEpoch(_epoch uint64) (*types.Transaction, error) {
	return _IEpoch.Contract.GetEpoch(&_IEpoch.TransactOpts, _epoch)
}

// IEpochSetEpochIterator is returned from FilterSetEpoch and is used to iterate over the raw logs and unpacked data for SetEpoch events raised by the IEpoch contract.
type IEpochSetEpochIterator struct {
	Event *IEpochSetEpoch // Event containing the contract specifics and raw log

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
func (it *IEpochSetEpochIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IEpochSetEpoch)
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
		it.Event = new(IEpochSetEpoch)
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
func (it *IEpochSetEpochIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IEpochSetEpochIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IEpochSetEpoch represents a SetEpoch event raised by the IEpoch contract.
type IEpochSetEpoch struct {
	E   uint64
	Raw types.Log // Blockchain specific contextual infos
}

// FilterSetEpoch is a free log retrieval operation binding the contract event 0x9d4ccb161ea809df14334516cc3070025c80baddb8e8364afaca6a6fe31bfd75.
//
// Solidity: event SetEpoch(uint64 _e)
func (_IEpoch *IEpochFilterer) FilterSetEpoch(opts *bind.FilterOpts) (*IEpochSetEpochIterator, error) {

	logs, sub, err := _IEpoch.contract.FilterLogs(opts, "SetEpoch")
	if err != nil {
		return nil, err
	}
	return &IEpochSetEpochIterator{contract: _IEpoch.contract, event: "SetEpoch", logs: logs, sub: sub}, nil
}

// WatchSetEpoch is a free log subscription operation binding the contract event 0x9d4ccb161ea809df14334516cc3070025c80baddb8e8364afaca6a6fe31bfd75.
//
// Solidity: event SetEpoch(uint64 _e)
func (_IEpoch *IEpochFilterer) WatchSetEpoch(opts *bind.WatchOpts, sink chan<- *IEpochSetEpoch) (event.Subscription, error) {

	logs, sub, err := _IEpoch.contract.WatchLogs(opts, "SetEpoch")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IEpochSetEpoch)
				if err := _IEpoch.contract.UnpackLog(event, "SetEpoch", log); err != nil {
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

// ParseSetEpoch is a log parse operation binding the contract event 0x9d4ccb161ea809df14334516cc3070025c80baddb8e8364afaca6a6fe31bfd75.
//
// Solidity: event SetEpoch(uint64 _e)
func (_IEpoch *IEpochFilterer) ParseSetEpoch(log types.Log) (*IEpochSetEpoch, error) {
	event := new(IEpochSetEpoch)
	if err := _IEpoch.contract.UnpackLog(event, "SetEpoch", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IPieceMetaData contains all meta data concerning the IPiece contract.
var IPieceMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_a\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"_pi\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"_e\",\"type\":\"uint64\"}],\"name\":\"AddPiece\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_a\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"_ri\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"_sri\",\"type\":\"uint64\"}],\"name\":\"AddReplica\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_a\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"_e\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_m\",\"type\":\"uint256\"}],\"name\":\"Settle\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_a\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"_pi\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_m\",\"type\":\"uint256\"}],\"name\":\"Withdraw\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"_pi\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"_ri\",\"type\":\"uint64\"},{\"internalType\":\"uint8\",\"name\":\"_pri\",\"type\":\"uint8\"}],\"name\":\"checkPRI\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_a\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"_ri\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"_com\",\"type\":\"bytes\"}],\"name\":\"checkSReplica\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_pn\",\"type\":\"bytes\"}],\"name\":\"getPIndex\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"_pi\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"_ri\",\"type\":\"uint64\"}],\"name\":\"getPRInfo\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"},{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_rn\",\"type\":\"bytes\"}],\"name\":\"getRIndex\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_a\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"_e\",\"type\":\"uint64\"}],\"name\":\"getStoreCount\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_a\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"_e\",\"type\":\"uint64\"}],\"name\":\"getStoreSalary\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// IPieceABI is the input ABI used to generate the binding from.
// Deprecated: Use IPieceMetaData.ABI instead.
var IPieceABI = IPieceMetaData.ABI

// IPiece is an auto generated Go binding around an Ethereum contract.
type IPiece struct {
	IPieceCaller     // Read-only binding to the contract
	IPieceTransactor // Write-only binding to the contract
	IPieceFilterer   // Log filterer for contract events
}

// IPieceCaller is an auto generated read-only Go binding around an Ethereum contract.
type IPieceCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IPieceTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IPieceTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IPieceFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IPieceFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IPieceSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IPieceSession struct {
	Contract     *IPiece           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IPieceCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IPieceCallerSession struct {
	Contract *IPieceCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// IPieceTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IPieceTransactorSession struct {
	Contract     *IPieceTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IPieceRaw is an auto generated low-level Go binding around an Ethereum contract.
type IPieceRaw struct {
	Contract *IPiece // Generic contract binding to access the raw methods on
}

// IPieceCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IPieceCallerRaw struct {
	Contract *IPieceCaller // Generic read-only contract binding to access the raw methods on
}

// IPieceTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IPieceTransactorRaw struct {
	Contract *IPieceTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIPiece creates a new instance of IPiece, bound to a specific deployed contract.
func NewIPiece(address common.Address, backend bind.ContractBackend) (*IPiece, error) {
	contract, err := bindIPiece(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IPiece{IPieceCaller: IPieceCaller{contract: contract}, IPieceTransactor: IPieceTransactor{contract: contract}, IPieceFilterer: IPieceFilterer{contract: contract}}, nil
}

// NewIPieceCaller creates a new read-only instance of IPiece, bound to a specific deployed contract.
func NewIPieceCaller(address common.Address, caller bind.ContractCaller) (*IPieceCaller, error) {
	contract, err := bindIPiece(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IPieceCaller{contract: contract}, nil
}

// NewIPieceTransactor creates a new write-only instance of IPiece, bound to a specific deployed contract.
func NewIPieceTransactor(address common.Address, transactor bind.ContractTransactor) (*IPieceTransactor, error) {
	contract, err := bindIPiece(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IPieceTransactor{contract: contract}, nil
}

// NewIPieceFilterer creates a new log filterer instance of IPiece, bound to a specific deployed contract.
func NewIPieceFilterer(address common.Address, filterer bind.ContractFilterer) (*IPieceFilterer, error) {
	contract, err := bindIPiece(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IPieceFilterer{contract: contract}, nil
}

// bindIPiece binds a generic wrapper to an already deployed contract.
func bindIPiece(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := IPieceMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IPiece *IPieceRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IPiece.Contract.IPieceCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IPiece *IPieceRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IPiece.Contract.IPieceTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IPiece *IPieceRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IPiece.Contract.IPieceTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IPiece *IPieceCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IPiece.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IPiece *IPieceTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IPiece.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IPiece *IPieceTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IPiece.Contract.contract.Transact(opts, method, params...)
}

// CheckSReplica is a free data retrieval call binding the contract method 0x894d35f8.
//
// Solidity: function checkSReplica(address _a, uint64 _ri, bytes _com) view returns(uint64)
func (_IPiece *IPieceCaller) CheckSReplica(opts *bind.CallOpts, _a common.Address, _ri uint64, _com []byte) (uint64, error) {
	var out []interface{}
	err := _IPiece.contract.Call(opts, &out, "checkSReplica", _a, _ri, _com)

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// CheckSReplica is a free data retrieval call binding the contract method 0x894d35f8.
//
// Solidity: function checkSReplica(address _a, uint64 _ri, bytes _com) view returns(uint64)
func (_IPiece *IPieceSession) CheckSReplica(_a common.Address, _ri uint64, _com []byte) (uint64, error) {
	return _IPiece.Contract.CheckSReplica(&_IPiece.CallOpts, _a, _ri, _com)
}

// CheckSReplica is a free data retrieval call binding the contract method 0x894d35f8.
//
// Solidity: function checkSReplica(address _a, uint64 _ri, bytes _com) view returns(uint64)
func (_IPiece *IPieceCallerSession) CheckSReplica(_a common.Address, _ri uint64, _com []byte) (uint64, error) {
	return _IPiece.Contract.CheckSReplica(&_IPiece.CallOpts, _a, _ri, _com)
}

// GetPIndex is a free data retrieval call binding the contract method 0x1ce85e7c.
//
// Solidity: function getPIndex(bytes _pn) view returns(uint64)
func (_IPiece *IPieceCaller) GetPIndex(opts *bind.CallOpts, _pn []byte) (uint64, error) {
	var out []interface{}
	err := _IPiece.contract.Call(opts, &out, "getPIndex", _pn)

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// GetPIndex is a free data retrieval call binding the contract method 0x1ce85e7c.
//
// Solidity: function getPIndex(bytes _pn) view returns(uint64)
func (_IPiece *IPieceSession) GetPIndex(_pn []byte) (uint64, error) {
	return _IPiece.Contract.GetPIndex(&_IPiece.CallOpts, _pn)
}

// GetPIndex is a free data retrieval call binding the contract method 0x1ce85e7c.
//
// Solidity: function getPIndex(bytes _pn) view returns(uint64)
func (_IPiece *IPieceCallerSession) GetPIndex(_pn []byte) (uint64, error) {
	return _IPiece.Contract.GetPIndex(&_IPiece.CallOpts, _pn)
}

// GetPRInfo is a free data retrieval call binding the contract method 0xb650b504.
//
// Solidity: function getPRInfo(uint64 _pi, uint64 _ri) view returns(uint8, uint8, uint64, bytes32)
func (_IPiece *IPieceCaller) GetPRInfo(opts *bind.CallOpts, _pi uint64, _ri uint64) (uint8, uint8, uint64, [32]byte, error) {
	var out []interface{}
	err := _IPiece.contract.Call(opts, &out, "getPRInfo", _pi, _ri)

	if err != nil {
		return *new(uint8), *new(uint8), *new(uint64), *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)
	out1 := *abi.ConvertType(out[1], new(uint8)).(*uint8)
	out2 := *abi.ConvertType(out[2], new(uint64)).(*uint64)
	out3 := *abi.ConvertType(out[3], new([32]byte)).(*[32]byte)

	return out0, out1, out2, out3, err

}

// GetPRInfo is a free data retrieval call binding the contract method 0xb650b504.
//
// Solidity: function getPRInfo(uint64 _pi, uint64 _ri) view returns(uint8, uint8, uint64, bytes32)
func (_IPiece *IPieceSession) GetPRInfo(_pi uint64, _ri uint64) (uint8, uint8, uint64, [32]byte, error) {
	return _IPiece.Contract.GetPRInfo(&_IPiece.CallOpts, _pi, _ri)
}

// GetPRInfo is a free data retrieval call binding the contract method 0xb650b504.
//
// Solidity: function getPRInfo(uint64 _pi, uint64 _ri) view returns(uint8, uint8, uint64, bytes32)
func (_IPiece *IPieceCallerSession) GetPRInfo(_pi uint64, _ri uint64) (uint8, uint8, uint64, [32]byte, error) {
	return _IPiece.Contract.GetPRInfo(&_IPiece.CallOpts, _pi, _ri)
}

// GetRIndex is a free data retrieval call binding the contract method 0xba8b2618.
//
// Solidity: function getRIndex(bytes _rn) view returns(uint64)
func (_IPiece *IPieceCaller) GetRIndex(opts *bind.CallOpts, _rn []byte) (uint64, error) {
	var out []interface{}
	err := _IPiece.contract.Call(opts, &out, "getRIndex", _rn)

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// GetRIndex is a free data retrieval call binding the contract method 0xba8b2618.
//
// Solidity: function getRIndex(bytes _rn) view returns(uint64)
func (_IPiece *IPieceSession) GetRIndex(_rn []byte) (uint64, error) {
	return _IPiece.Contract.GetRIndex(&_IPiece.CallOpts, _rn)
}

// GetRIndex is a free data retrieval call binding the contract method 0xba8b2618.
//
// Solidity: function getRIndex(bytes _rn) view returns(uint64)
func (_IPiece *IPieceCallerSession) GetRIndex(_rn []byte) (uint64, error) {
	return _IPiece.Contract.GetRIndex(&_IPiece.CallOpts, _rn)
}

// CheckPRI is a paid mutator transaction binding the contract method 0x2700f3a8.
//
// Solidity: function checkPRI(uint64 _pi, uint64 _ri, uint8 _pri) returns(address)
func (_IPiece *IPieceTransactor) CheckPRI(opts *bind.TransactOpts, _pi uint64, _ri uint64, _pri uint8) (*types.Transaction, error) {
	return _IPiece.contract.Transact(opts, "checkPRI", _pi, _ri, _pri)
}

// CheckPRI is a paid mutator transaction binding the contract method 0x2700f3a8.
//
// Solidity: function checkPRI(uint64 _pi, uint64 _ri, uint8 _pri) returns(address)
func (_IPiece *IPieceSession) CheckPRI(_pi uint64, _ri uint64, _pri uint8) (*types.Transaction, error) {
	return _IPiece.Contract.CheckPRI(&_IPiece.TransactOpts, _pi, _ri, _pri)
}

// CheckPRI is a paid mutator transaction binding the contract method 0x2700f3a8.
//
// Solidity: function checkPRI(uint64 _pi, uint64 _ri, uint8 _pri) returns(address)
func (_IPiece *IPieceTransactorSession) CheckPRI(_pi uint64, _ri uint64, _pri uint8) (*types.Transaction, error) {
	return _IPiece.Contract.CheckPRI(&_IPiece.TransactOpts, _pi, _ri, _pri)
}

// GetStoreCount is a paid mutator transaction binding the contract method 0x50979d7e.
//
// Solidity: function getStoreCount(address _a, uint64 _e) returns(uint64)
func (_IPiece *IPieceTransactor) GetStoreCount(opts *bind.TransactOpts, _a common.Address, _e uint64) (*types.Transaction, error) {
	return _IPiece.contract.Transact(opts, "getStoreCount", _a, _e)
}

// GetStoreCount is a paid mutator transaction binding the contract method 0x50979d7e.
//
// Solidity: function getStoreCount(address _a, uint64 _e) returns(uint64)
func (_IPiece *IPieceSession) GetStoreCount(_a common.Address, _e uint64) (*types.Transaction, error) {
	return _IPiece.Contract.GetStoreCount(&_IPiece.TransactOpts, _a, _e)
}

// GetStoreCount is a paid mutator transaction binding the contract method 0x50979d7e.
//
// Solidity: function getStoreCount(address _a, uint64 _e) returns(uint64)
func (_IPiece *IPieceTransactorSession) GetStoreCount(_a common.Address, _e uint64) (*types.Transaction, error) {
	return _IPiece.Contract.GetStoreCount(&_IPiece.TransactOpts, _a, _e)
}

// GetStoreSalary is a paid mutator transaction binding the contract method 0x18198ad7.
//
// Solidity: function getStoreSalary(address _a, uint64 _e) returns(uint256)
func (_IPiece *IPieceTransactor) GetStoreSalary(opts *bind.TransactOpts, _a common.Address, _e uint64) (*types.Transaction, error) {
	return _IPiece.contract.Transact(opts, "getStoreSalary", _a, _e)
}

// GetStoreSalary is a paid mutator transaction binding the contract method 0x18198ad7.
//
// Solidity: function getStoreSalary(address _a, uint64 _e) returns(uint256)
func (_IPiece *IPieceSession) GetStoreSalary(_a common.Address, _e uint64) (*types.Transaction, error) {
	return _IPiece.Contract.GetStoreSalary(&_IPiece.TransactOpts, _a, _e)
}

// GetStoreSalary is a paid mutator transaction binding the contract method 0x18198ad7.
//
// Solidity: function getStoreSalary(address _a, uint64 _e) returns(uint256)
func (_IPiece *IPieceTransactorSession) GetStoreSalary(_a common.Address, _e uint64) (*types.Transaction, error) {
	return _IPiece.Contract.GetStoreSalary(&_IPiece.TransactOpts, _a, _e)
}

// IPieceAddPieceIterator is returned from FilterAddPiece and is used to iterate over the raw logs and unpacked data for AddPiece events raised by the IPiece contract.
type IPieceAddPieceIterator struct {
	Event *IPieceAddPiece // Event containing the contract specifics and raw log

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
func (it *IPieceAddPieceIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IPieceAddPiece)
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
		it.Event = new(IPieceAddPiece)
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
func (it *IPieceAddPieceIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IPieceAddPieceIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IPieceAddPiece represents a AddPiece event raised by the IPiece contract.
type IPieceAddPiece struct {
	A   common.Address
	Pi  uint64
	E   uint64
	Raw types.Log // Blockchain specific contextual infos
}

// FilterAddPiece is a free log retrieval operation binding the contract event 0x09ff0fb5f488978fdffbf3223e63e914f2a9b0c228844b8e2caedad8a85d220c.
//
// Solidity: event AddPiece(address indexed _a, uint64 _pi, uint64 _e)
func (_IPiece *IPieceFilterer) FilterAddPiece(opts *bind.FilterOpts, _a []common.Address) (*IPieceAddPieceIterator, error) {

	var _aRule []interface{}
	for _, _aItem := range _a {
		_aRule = append(_aRule, _aItem)
	}

	logs, sub, err := _IPiece.contract.FilterLogs(opts, "AddPiece", _aRule)
	if err != nil {
		return nil, err
	}
	return &IPieceAddPieceIterator{contract: _IPiece.contract, event: "AddPiece", logs: logs, sub: sub}, nil
}

// WatchAddPiece is a free log subscription operation binding the contract event 0x09ff0fb5f488978fdffbf3223e63e914f2a9b0c228844b8e2caedad8a85d220c.
//
// Solidity: event AddPiece(address indexed _a, uint64 _pi, uint64 _e)
func (_IPiece *IPieceFilterer) WatchAddPiece(opts *bind.WatchOpts, sink chan<- *IPieceAddPiece, _a []common.Address) (event.Subscription, error) {

	var _aRule []interface{}
	for _, _aItem := range _a {
		_aRule = append(_aRule, _aItem)
	}

	logs, sub, err := _IPiece.contract.WatchLogs(opts, "AddPiece", _aRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IPieceAddPiece)
				if err := _IPiece.contract.UnpackLog(event, "AddPiece", log); err != nil {
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

// ParseAddPiece is a log parse operation binding the contract event 0x09ff0fb5f488978fdffbf3223e63e914f2a9b0c228844b8e2caedad8a85d220c.
//
// Solidity: event AddPiece(address indexed _a, uint64 _pi, uint64 _e)
func (_IPiece *IPieceFilterer) ParseAddPiece(log types.Log) (*IPieceAddPiece, error) {
	event := new(IPieceAddPiece)
	if err := _IPiece.contract.UnpackLog(event, "AddPiece", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IPieceAddReplicaIterator is returned from FilterAddReplica and is used to iterate over the raw logs and unpacked data for AddReplica events raised by the IPiece contract.
type IPieceAddReplicaIterator struct {
	Event *IPieceAddReplica // Event containing the contract specifics and raw log

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
func (it *IPieceAddReplicaIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IPieceAddReplica)
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
		it.Event = new(IPieceAddReplica)
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
func (it *IPieceAddReplicaIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IPieceAddReplicaIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IPieceAddReplica represents a AddReplica event raised by the IPiece contract.
type IPieceAddReplica struct {
	A   common.Address
	Ri  uint64
	Sri uint64
	Raw types.Log // Blockchain specific contextual infos
}

// FilterAddReplica is a free log retrieval operation binding the contract event 0x4b079e585085a3c94b2865036e598391fad4359764b36da39e21897be56349ef.
//
// Solidity: event AddReplica(address indexed _a, uint64 _ri, uint64 _sri)
func (_IPiece *IPieceFilterer) FilterAddReplica(opts *bind.FilterOpts, _a []common.Address) (*IPieceAddReplicaIterator, error) {

	var _aRule []interface{}
	for _, _aItem := range _a {
		_aRule = append(_aRule, _aItem)
	}

	logs, sub, err := _IPiece.contract.FilterLogs(opts, "AddReplica", _aRule)
	if err != nil {
		return nil, err
	}
	return &IPieceAddReplicaIterator{contract: _IPiece.contract, event: "AddReplica", logs: logs, sub: sub}, nil
}

// WatchAddReplica is a free log subscription operation binding the contract event 0x4b079e585085a3c94b2865036e598391fad4359764b36da39e21897be56349ef.
//
// Solidity: event AddReplica(address indexed _a, uint64 _ri, uint64 _sri)
func (_IPiece *IPieceFilterer) WatchAddReplica(opts *bind.WatchOpts, sink chan<- *IPieceAddReplica, _a []common.Address) (event.Subscription, error) {

	var _aRule []interface{}
	for _, _aItem := range _a {
		_aRule = append(_aRule, _aItem)
	}

	logs, sub, err := _IPiece.contract.WatchLogs(opts, "AddReplica", _aRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IPieceAddReplica)
				if err := _IPiece.contract.UnpackLog(event, "AddReplica", log); err != nil {
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

// ParseAddReplica is a log parse operation binding the contract event 0x4b079e585085a3c94b2865036e598391fad4359764b36da39e21897be56349ef.
//
// Solidity: event AddReplica(address indexed _a, uint64 _ri, uint64 _sri)
func (_IPiece *IPieceFilterer) ParseAddReplica(log types.Log) (*IPieceAddReplica, error) {
	event := new(IPieceAddReplica)
	if err := _IPiece.contract.UnpackLog(event, "AddReplica", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IPieceSettleIterator is returned from FilterSettle and is used to iterate over the raw logs and unpacked data for Settle events raised by the IPiece contract.
type IPieceSettleIterator struct {
	Event *IPieceSettle // Event containing the contract specifics and raw log

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
func (it *IPieceSettleIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IPieceSettle)
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
		it.Event = new(IPieceSettle)
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
func (it *IPieceSettleIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IPieceSettleIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IPieceSettle represents a Settle event raised by the IPiece contract.
type IPieceSettle struct {
	A   common.Address
	E   uint64
	M   *big.Int
	Raw types.Log // Blockchain specific contextual infos
}

// FilterSettle is a free log retrieval operation binding the contract event 0x22b9b452d488700bdfeab2aa62972ac1fce3160583db89ab4c799708bf5a2c6c.
//
// Solidity: event Settle(address indexed _a, uint64 _e, uint256 _m)
func (_IPiece *IPieceFilterer) FilterSettle(opts *bind.FilterOpts, _a []common.Address) (*IPieceSettleIterator, error) {

	var _aRule []interface{}
	for _, _aItem := range _a {
		_aRule = append(_aRule, _aItem)
	}

	logs, sub, err := _IPiece.contract.FilterLogs(opts, "Settle", _aRule)
	if err != nil {
		return nil, err
	}
	return &IPieceSettleIterator{contract: _IPiece.contract, event: "Settle", logs: logs, sub: sub}, nil
}

// WatchSettle is a free log subscription operation binding the contract event 0x22b9b452d488700bdfeab2aa62972ac1fce3160583db89ab4c799708bf5a2c6c.
//
// Solidity: event Settle(address indexed _a, uint64 _e, uint256 _m)
func (_IPiece *IPieceFilterer) WatchSettle(opts *bind.WatchOpts, sink chan<- *IPieceSettle, _a []common.Address) (event.Subscription, error) {

	var _aRule []interface{}
	for _, _aItem := range _a {
		_aRule = append(_aRule, _aItem)
	}

	logs, sub, err := _IPiece.contract.WatchLogs(opts, "Settle", _aRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IPieceSettle)
				if err := _IPiece.contract.UnpackLog(event, "Settle", log); err != nil {
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

// ParseSettle is a log parse operation binding the contract event 0x22b9b452d488700bdfeab2aa62972ac1fce3160583db89ab4c799708bf5a2c6c.
//
// Solidity: event Settle(address indexed _a, uint64 _e, uint256 _m)
func (_IPiece *IPieceFilterer) ParseSettle(log types.Log) (*IPieceSettle, error) {
	event := new(IPieceSettle)
	if err := _IPiece.contract.UnpackLog(event, "Settle", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IPieceWithdrawIterator is returned from FilterWithdraw and is used to iterate over the raw logs and unpacked data for Withdraw events raised by the IPiece contract.
type IPieceWithdrawIterator struct {
	Event *IPieceWithdraw // Event containing the contract specifics and raw log

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
func (it *IPieceWithdrawIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IPieceWithdraw)
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
		it.Event = new(IPieceWithdraw)
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
func (it *IPieceWithdrawIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IPieceWithdrawIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IPieceWithdraw represents a Withdraw event raised by the IPiece contract.
type IPieceWithdraw struct {
	A   common.Address
	Pi  uint64
	M   *big.Int
	Raw types.Log // Blockchain specific contextual infos
}

// FilterWithdraw is a free log retrieval operation binding the contract event 0xb283270b87db7ad5d1fbb15af2039324aa28bebf00c89e37579882f7cb261d19.
//
// Solidity: event Withdraw(address indexed _a, uint64 _pi, uint256 _m)
func (_IPiece *IPieceFilterer) FilterWithdraw(opts *bind.FilterOpts, _a []common.Address) (*IPieceWithdrawIterator, error) {

	var _aRule []interface{}
	for _, _aItem := range _a {
		_aRule = append(_aRule, _aItem)
	}

	logs, sub, err := _IPiece.contract.FilterLogs(opts, "Withdraw", _aRule)
	if err != nil {
		return nil, err
	}
	return &IPieceWithdrawIterator{contract: _IPiece.contract, event: "Withdraw", logs: logs, sub: sub}, nil
}

// WatchWithdraw is a free log subscription operation binding the contract event 0xb283270b87db7ad5d1fbb15af2039324aa28bebf00c89e37579882f7cb261d19.
//
// Solidity: event Withdraw(address indexed _a, uint64 _pi, uint256 _m)
func (_IPiece *IPieceFilterer) WatchWithdraw(opts *bind.WatchOpts, sink chan<- *IPieceWithdraw, _a []common.Address) (event.Subscription, error) {

	var _aRule []interface{}
	for _, _aItem := range _a {
		_aRule = append(_aRule, _aItem)
	}

	logs, sub, err := _IPiece.contract.WatchLogs(opts, "Withdraw", _aRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IPieceWithdraw)
				if err := _IPiece.contract.UnpackLog(event, "Withdraw", log); err != nil {
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

// ParseWithdraw is a log parse operation binding the contract event 0xb283270b87db7ad5d1fbb15af2039324aa28bebf00c89e37579882f7cb261d19.
//
// Solidity: event Withdraw(address indexed _a, uint64 _pi, uint256 _m)
func (_IPiece *IPieceFilterer) ParseWithdraw(log types.Log) (*IPieceWithdraw, error) {
	event := new(IPieceWithdraw)
	if err := _IPiece.contract.UnpackLog(event, "Withdraw", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OwnableMetaData contains all meta data concerning the Ownable contract.
var OwnableMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// OwnableABI is the input ABI used to generate the binding from.
// Deprecated: Use OwnableMetaData.ABI instead.
var OwnableABI = OwnableMetaData.ABI

// Ownable is an auto generated Go binding around an Ethereum contract.
type Ownable struct {
	OwnableCaller     // Read-only binding to the contract
	OwnableTransactor // Write-only binding to the contract
	OwnableFilterer   // Log filterer for contract events
}

// OwnableCaller is an auto generated read-only Go binding around an Ethereum contract.
type OwnableCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OwnableTransactor is an auto generated write-only Go binding around an Ethereum contract.
type OwnableTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OwnableFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type OwnableFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OwnableSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type OwnableSession struct {
	Contract     *Ownable          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// OwnableCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type OwnableCallerSession struct {
	Contract *OwnableCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// OwnableTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type OwnableTransactorSession struct {
	Contract     *OwnableTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// OwnableRaw is an auto generated low-level Go binding around an Ethereum contract.
type OwnableRaw struct {
	Contract *Ownable // Generic contract binding to access the raw methods on
}

// OwnableCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type OwnableCallerRaw struct {
	Contract *OwnableCaller // Generic read-only contract binding to access the raw methods on
}

// OwnableTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type OwnableTransactorRaw struct {
	Contract *OwnableTransactor // Generic write-only contract binding to access the raw methods on
}

// NewOwnable creates a new instance of Ownable, bound to a specific deployed contract.
func NewOwnable(address common.Address, backend bind.ContractBackend) (*Ownable, error) {
	contract, err := bindOwnable(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Ownable{OwnableCaller: OwnableCaller{contract: contract}, OwnableTransactor: OwnableTransactor{contract: contract}, OwnableFilterer: OwnableFilterer{contract: contract}}, nil
}

// NewOwnableCaller creates a new read-only instance of Ownable, bound to a specific deployed contract.
func NewOwnableCaller(address common.Address, caller bind.ContractCaller) (*OwnableCaller, error) {
	contract, err := bindOwnable(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &OwnableCaller{contract: contract}, nil
}

// NewOwnableTransactor creates a new write-only instance of Ownable, bound to a specific deployed contract.
func NewOwnableTransactor(address common.Address, transactor bind.ContractTransactor) (*OwnableTransactor, error) {
	contract, err := bindOwnable(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &OwnableTransactor{contract: contract}, nil
}

// NewOwnableFilterer creates a new log filterer instance of Ownable, bound to a specific deployed contract.
func NewOwnableFilterer(address common.Address, filterer bind.ContractFilterer) (*OwnableFilterer, error) {
	contract, err := bindOwnable(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &OwnableFilterer{contract: contract}, nil
}

// bindOwnable binds a generic wrapper to an already deployed contract.
func bindOwnable(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := OwnableMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Ownable *OwnableRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Ownable.Contract.OwnableCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Ownable *OwnableRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Ownable.Contract.OwnableTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Ownable *OwnableRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Ownable.Contract.OwnableTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Ownable *OwnableCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Ownable.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Ownable *OwnableTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Ownable.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Ownable *OwnableTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Ownable.Contract.contract.Transact(opts, method, params...)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Ownable *OwnableCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Ownable.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Ownable *OwnableSession) Owner() (common.Address, error) {
	return _Ownable.Contract.Owner(&_Ownable.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Ownable *OwnableCallerSession) Owner() (common.Address, error) {
	return _Ownable.Contract.Owner(&_Ownable.CallOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Ownable *OwnableTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Ownable.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Ownable *OwnableSession) RenounceOwnership() (*types.Transaction, error) {
	return _Ownable.Contract.RenounceOwnership(&_Ownable.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Ownable *OwnableTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Ownable.Contract.RenounceOwnership(&_Ownable.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Ownable *OwnableTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Ownable.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Ownable *OwnableSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Ownable.Contract.TransferOwnership(&_Ownable.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Ownable *OwnableTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Ownable.Contract.TransferOwnership(&_Ownable.TransactOpts, newOwner)
}

// OwnableOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Ownable contract.
type OwnableOwnershipTransferredIterator struct {
	Event *OwnableOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *OwnableOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OwnableOwnershipTransferred)
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
		it.Event = new(OwnableOwnershipTransferred)
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
func (it *OwnableOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OwnableOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OwnableOwnershipTransferred represents a OwnershipTransferred event raised by the Ownable contract.
type OwnableOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Ownable *OwnableFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*OwnableOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Ownable.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &OwnableOwnershipTransferredIterator{contract: _Ownable.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Ownable *OwnableFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *OwnableOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Ownable.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OwnableOwnershipTransferred)
				if err := _Ownable.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_Ownable *OwnableFilterer) ParseOwnershipTransferred(log types.Log) (*OwnableOwnershipTransferred, error) {
	event := new(OwnableOwnershipTransferred)
	if err := _Ownable.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PieceMetaData contains all meta data concerning the Piece contract.
var PieceMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_b\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_a\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"_pi\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"_e\",\"type\":\"uint64\"}],\"name\":\"AddPiece\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_a\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"_ri\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"_sri\",\"type\":\"uint64\"}],\"name\":\"AddReplica\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_a\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"_e\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_m\",\"type\":\"uint256\"}],\"name\":\"Settle\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_a\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"_pi\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_m\",\"type\":\"uint256\"}],\"name\":\"Withdraw\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_pn\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"_price\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"_size\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"_expire\",\"type\":\"uint64\"},{\"internalType\":\"uint8\",\"name\":\"rsn\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"rsk\",\"type\":\"uint8\"}],\"name\":\"addPiece\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_rn\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"_pi\",\"type\":\"uint64\"},{\"internalType\":\"uint8\",\"name\":\"_pri\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"_proof\",\"type\":\"bytes\"}],\"name\":\"addReplica\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"bank\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"_pi\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"_ri\",\"type\":\"uint64\"},{\"internalType\":\"uint8\",\"name\":\"_pri\",\"type\":\"uint8\"}],\"name\":\"checkPRI\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_a\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"_ri\",\"type\":\"uint64\"}],\"name\":\"checkSRAt\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_a\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"_ri\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"_com\",\"type\":\"bytes\"}],\"name\":\"checkSReplica\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"current\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"delay\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_pn\",\"type\":\"bytes\"}],\"name\":\"getPIndex\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"_pi\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"_ri\",\"type\":\"uint64\"}],\"name\":\"getPRInfo\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"},{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"_pi\",\"type\":\"uint64\"}],\"name\":\"getPiece\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"size\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"start\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"expire\",\"type\":\"uint64\"},{\"internalType\":\"uint8\",\"name\":\"rsn\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"rsk\",\"type\":\"uint8\"}],\"internalType\":\"structIPiece.PieceInfo\",\"name\":\"_pinfo\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_rn\",\"type\":\"bytes\"}],\"name\":\"getRIndex\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"_r\",\"type\":\"uint64\"}],\"name\":\"getReplica\",\"outputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"piece\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"storedOn\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"root\",\"type\":\"bytes32\"}],\"internalType\":\"structIPiece.ReplicaInfo\",\"name\":\"_ri\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_a\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"_e\",\"type\":\"uint64\"}],\"name\":\"getSStat\",\"outputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"count\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"salary\",\"type\":\"uint256\"}],\"internalType\":\"structIPiece.StoreStat\",\"name\":\"_si\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_a\",\"type\":\"address\"}],\"name\":\"getStore\",\"outputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"epoch\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"count\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"salary\",\"type\":\"uint256\"}],\"internalType\":\"structIPiece.StoreInfo\",\"name\":\"_si\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_a\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"_e\",\"type\":\"uint64\"}],\"name\":\"getStoreCount\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_a\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"_e\",\"type\":\"uint64\"}],\"name\":\"getStoreSalary\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"maxSize\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"maxStore\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"minPrice\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"minStore\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"_ep\",\"type\":\"uint64\"}],\"name\":\"settle\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"_pi\",\"type\":\"uint64\"}],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x6080604052600760025f6101000a81548167ffffffffffffffff021916908367ffffffffffffffff1602179055506064600260086101000a81548167ffffffffffffffff021916908367ffffffffffffffff1602179055506103e8600260106101000a81548167ffffffffffffffff021916908367ffffffffffffffff1602179055506302000000600260186101000a81548167ffffffffffffffff021916908367ffffffffffffffff160217905550655af3107a40006003553480156100c4575f80fd5b5060405161509138038061509183398181016040528101906100e69190610530565b6101026100f761035260201b60201c565b61035960201b60201c565b8060015f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555061014a61041a565b600481908060018154018082558091505060019003905f5260205f2090600402015f909190919091505f820151815f015f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060208201518160010155604082015181600201556060820151816003015f6101000a81548167ffffffffffffffff021916908367ffffffffffffffff16021790555060808201518160030160086101000a81548167ffffffffffffffff021916908367ffffffffffffffff16021790555060a08201518160030160106101000a81548167ffffffffffffffff021916908367ffffffffffffffff16021790555060c08201518160030160186101000a81548160ff021916908360ff16021790555060e08201518160030160196101000a81548160ff021916908360ff16021790555050506102a3610491565b600881908060018154018082558091505060019003905f5260205f2090600202015f909190919091505f820151815f015f6101000a81548167ffffffffffffffff021916908367ffffffffffffffff1602179055506020820151815f0160086101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060408201518160010155505050505061055b565b5f33905090565b5f805f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050815f806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508173ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a35050565b6040518061010001604052805f73ffffffffffffffffffffffffffffffffffffffff1681526020015f81526020015f81526020015f67ffffffffffffffff1681526020015f67ffffffffffffffff1681526020015f67ffffffffffffffff1681526020015f60ff1681526020015f60ff1681525090565b60405180606001604052805f67ffffffffffffffff1681526020015f73ffffffffffffffffffffffffffffffffffffffff1681526020015f80191681525090565b5f80fd5b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f6104ff826104d6565b9050919050565b61050f816104f5565b8114610519575f80fd5b50565b5f8151905061052a81610506565b92915050565b5f60208284031215610545576105446104d2565b5b5f6105528482850161051c565b91505092915050565b614b29806105685f395ff3fe608060405234801561000f575f80fd5b506004361061018c575f3560e01c806376cdb03b116100dc5780639fa6a6e311610095578063c3aa38041161006f578063c3aa3804146104d1578063e45be8eb14610501578063f2fde38b1461051f578063fca5be7a1461053b5761018c565b80639fa6a6e314610450578063b650b5041461046e578063ba8b2618146104a15761018c565b806376cdb03b146103665780637a87f0c41461038457806388795125146103b4578063894d35f8146103d25780638da5cb5b146104025780639108544c146104205761018c565b80634b9ba4461161014957806367439a2d1161012357806367439a2d146102f25780636a42b8f814610322578063715018a614610340578063750f0acc1461034a5761018c565b80634b9ba4461461028a5780634d9c3d4b146102a657806350979d7e146102c25761018c565b806318198ad714610190578063194cbaf5146101c05780631ce85e7c146101dc5780632565b1591461020c5780632700f3a81461022a5780634b4ffccd1461025a575b5f80fd5b6101aa60048036038101906101a59190613796565b610559565b6040516101b791906137ec565b60405180910390f35b6101da60048036038101906101d591906139d4565b6105c6565b005b6101f660048036038101906101f19190613a74565b611081565b6040516102039190613aca565b60405180910390f35b6102146110bb565b6040516102219190613aca565b60405180910390f35b610244600480360381019061023f9190613ae3565b6110d5565b6040516102519190613b42565b60405180910390f35b610274600480360381019061026f9190613b5b565b6111cf565b6040516102819190613be4565b60405180910390f35b6102a4600480360381019061029f9190613bfd565b6112f2565b005b6102c060048036038101906102bb9190613c52565b61177a565b005b6102dc60048036038101906102d79190613796565b611e78565b6040516102e99190613aca565b60405180910390f35b61030c60048036038101906103079190613796565b611ef7565b6040516103199190613aca565b60405180910390f35b61032a611f74565b6040516103379190613aca565b60405180910390f35b610348611f8d565b005b610364600480360381019061035f9190613bfd565b611fa0565b005b61036e612296565b60405161037b9190613b42565b60405180910390f35b61039e60048036038101906103999190613bfd565b6122bb565b6040516103ab9190613db5565b60405180910390f35b6103bc61253b565b6040516103c99190613aca565b60405180910390f35b6103ec60048036038101906103e79190613dcf565b612555565b6040516103f99190613aca565b60405180910390f35b61040a6126e0565b6040516104179190613b42565b60405180910390f35b61043a60048036038101906104359190613bfd565b612707565b6040516104479190613e93565b60405180910390f35b610458612809565b6040516104659190613aca565b60405180910390f35b61048860048036038101906104839190613eac565b612823565b6040516104989493929190613f08565b60405180910390f35b6104bb60048036038101906104b69190613a74565b61298e565b6040516104c89190613aca565b60405180910390f35b6104eb60048036038101906104e69190613796565b6129c8565b6040516104f89190613f78565b60405180910390f35b610509612acf565b60405161051691906137ec565b60405180910390f35b61053960048036038101906105349190613b5b565b612ad5565b005b610543612b57565b6040516105509190613aca565b60405180910390f35b5f600c5f8473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f8367ffffffffffffffff1667ffffffffffffffff1681526020019081526020015f2060010154905092915050565b6105ce612b71565b600160149054906101000a900467ffffffffffffffff1667ffffffffffffffff1660025f9054906101000a900467ffffffffffffffff1660048667ffffffffffffffff168154811061062357610622613f91565b5b905f5260205f20906004020160030160089054906101000a900467ffffffffffffffff166106519190613feb565b67ffffffffffffffff161161069b576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161069290614080565b60405180910390fd5b8260ff1660048567ffffffffffffffff16815481106106bd576106bc613f91565b5b905f5260205f20906004020160030160189054906101000a900460ff1660ff16101561071e576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610715906140e8565b60405180910390fd5b5f60098660405161072f9190614158565b90815260200160405180910390205f9054906101000a900467ffffffffffffffff1667ffffffffffffffff161461079b576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610792906141b8565b60405180910390fd5b5f60065f8667ffffffffffffffff1667ffffffffffffffff1681526020019081526020015f205f8560ff1660ff1681526020019081526020015f205f9054906101000a900467ffffffffffffffff1667ffffffffffffffff1614610834576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161082b90614220565b60405180910390fd5b60075f8567ffffffffffffffff1667ffffffffffffffff1681526020019081526020015f205f3373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f9054906101000a900460ff16156108e1576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016108d890614288565b60405180910390fd5b5f60048567ffffffffffffffff1681548110610900576108ff613f91565b5b905f5260205f20906004020160030160199054906101000a900460ff1660ff16620f800061092e91906142a6565b600160048767ffffffffffffffff168154811061094e5761094d613f91565b5b905f5260205f2090600402016003015f9054906101000a900467ffffffffffffffff1661097b91906142e2565b610985919061434a565b60016109919190613feb565b90505f8167ffffffffffffffff16600160149054906101000a900467ffffffffffffffff1660048867ffffffffffffffff16815481106109d4576109d3613f91565b5b905f5260205f20906004020160030160109054906101000a900467ffffffffffffffff16610a0291906142e2565b67ffffffffffffffff1660048867ffffffffffffffff1681548110610a2a57610a29613f91565b5b905f5260205f20906004020160020154610a44919061437a565b610a4e919061437a565b90508060048767ffffffffffffffff1681548110610a6f57610a6e613f91565b5b905f5260205f2090600402016001015f828254610a8c91906143bb565b9250508190555060015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663693ec85e6040518163ffffffff1660e01b8152600401610aeb90614438565b6020604051808303815f875af1158015610b07573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610b2b919061446a565b73ffffffffffffffffffffffffffffffffffffffff1663dae448b933600160149054906101000a900467ffffffffffffffff16846040518463ffffffff1660e01b8152600401610b7d93929190614495565b5f604051808303815f87803b158015610b94575f80fd5b505af1158015610ba6573d5f803e3d5ffd5b50505050610bb26135e2565b86815f019067ffffffffffffffff16908167ffffffffffffffff168152505033816020019073ffffffffffffffffffffffffffffffffffffffff16908173ffffffffffffffffffffffffffffffffffffffff16815250508484604051610c199291906144ee565b60405180910390208160400181815250505f6008805490509050600882908060018154018082558091505060019003905f5260205f2090600202015f909190919091505f820151815f015f6101000a81548167ffffffffffffffff021916908367ffffffffffffffff1602179055506020820151815f0160086101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506040820151816001015550508060098a604051610ceb9190614158565b90815260200160405180910390205f6101000a81548167ffffffffffffffff021916908367ffffffffffffffff1602179055508060065f8a67ffffffffffffffff1667ffffffffffffffff1681526020019081526020015f205f8960ff1660ff1681526020019081526020015f205f6101000a81548167ffffffffffffffff021916908367ffffffffffffffff160217905550600160075f8a67ffffffffffffffff1667ffffffffffffffff1681526020019081526020015f205f3373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f6101000a81548160ff0219169083151502179055508367ffffffffffffffff1660048967ffffffffffffffff1681548110610e1f57610e1e613f91565b5b905f5260205f20906004020160020154610e39919061437a565b9250610e88338260048b67ffffffffffffffff1681548110610e5e57610e5d613f91565b5b905f5260205f20906004020160030160109054906101000a900467ffffffffffffffff1686612d38565b60015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663693ec85e6040518163ffffffff1660e01b8152600401610ee090614438565b6020604051808303815f875af1158015610efc573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610f20919061446a565b73ffffffffffffffffffffffffffffffffffffffff1663a3f6086933600a5f3373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f20600101546040518363ffffffff1660e01b8152600401610f9a929190614506565b5f604051808303815f87803b158015610fb1575f80fd5b505af1158015610fc3573d5f803e3d5ffd5b505050503373ffffffffffffffffffffffffffffffffffffffff167f4b079e585085a3c94b2865036e598391fad4359764b36da39e21897be56349ef826001600a5f3373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f0160089054906101000a900467ffffffffffffffff1661106091906142e2565b60405161106e92919061452d565b60405180910390a2505050505050505050565b5f6005826040516110929190614158565b90815260200160405180910390205f9054906101000a900467ffffffffffffffff169050919050565b600260189054906101000a900467ffffffffffffffff1681565b5f8267ffffffffffffffff1660065f8667ffffffffffffffff1667ffffffffffffffff1681526020019081526020015f205f8460ff1660ff1681526020019081526020015f205f9054906101000a900467ffffffffffffffff1667ffffffffffffffff1614611179576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401611170906140e8565b60405180910390fd5b60088367ffffffffffffffff168154811061119757611196613f91565b5b905f5260205f2090600202015f0160089054906101000a900473ffffffffffffffffffffffffffffffffffffffff1690509392505050565b6111d7613623565b6040518060600160405280600a5f8573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f015f9054906101000a900467ffffffffffffffff1667ffffffffffffffff168152602001600a5f8573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f0160089054906101000a900467ffffffffffffffff1667ffffffffffffffff168152602001600a5f8573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f20600101548152509050919050565b6112fa612b71565b61130333612f78565b60015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663693ec85e6040518163ffffffff1660e01b815260040161135b90614438565b6020604051808303815f875af1158015611377573d5f803e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061139b919061446a565b73ffffffffffffffffffffffffffffffffffffffff1663a3f6086933600a5f3373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f20600101546040518363ffffffff1660e01b8152600401611415929190614506565b5f604051808303815f87803b15801561142c575f80fd5b505af115801561143e573d5f803e3d5ffd5b50505050600160149054906101000a900467ffffffffffffffff1667ffffffffffffffff1660025f9054906101000a900467ffffffffffffffff16826114849190613feb565b67ffffffffffffffff1611156114cf576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016114c69061459e565b60405180910390fd5b5f600c5f3373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f8367ffffffffffffffff1667ffffffffffffffff1681526020019081526020015f2060010154905060015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663693ec85e6040518163ffffffff1660e01b815260040161158e90614438565b6020604051808303815f875af11580156115aa573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906115ce919061446a565b73ffffffffffffffffffffffffffffffffffffffff16630357371d33836040518363ffffffff1660e01b8152600401611608929190614506565b5f604051808303815f87803b15801561161f575f80fd5b505af1158015611631573d5f803e3d5ffd5b5050505060015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663e888891533836040518363ffffffff1660e01b8152600401611691929190614506565b5f604051808303815f87803b1580156116a8575f80fd5b505af11580156116ba573d5f803e3d5ffd5b505050503373ffffffffffffffffffffffffffffffffffffffff167f22b9b452d488700bdfeab2aa62972ac1fce3160583db89ab4c799708bf5a2c6c83836040516117069291906145bc565b60405180910390a25f600c5f3373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f8467ffffffffffffffff1667ffffffffffffffff1681526020019081526020015f20600101819055505050565b611782612b71565b5f8467ffffffffffffffff16116117ce576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016117c59061462d565b60405180910390fd5b5f8160ff1611611813576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161180a9061462d565b60405180910390fd5b8160ff168160ff161061185b576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161185290614695565b60405180910390fd5b600260089054906101000a900467ffffffffffffffff16600160149054906101000a900467ffffffffffffffff166118939190613feb565b67ffffffffffffffff168367ffffffffffffffff1610156118e9576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016118e0906146fd565b60405180910390fd5b600260109054906101000a900467ffffffffffffffff16600160149054906101000a900467ffffffffffffffff166119219190613feb565b67ffffffffffffffff168367ffffffffffffffff161115611977576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161196e90614765565b60405180910390fd5b6003548510156119bc576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016119b3906147cd565b60405180910390fd5b5f8160ff16601f6119cd91906142a6565b6001866119da91906142e2565b6119e4919061434a565b60016119f09190613feb565b9050600260189054906101000a900467ffffffffffffffff1667ffffffffffffffff168260ff1682611a2291906142a6565b67ffffffffffffffff161115611a6d576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401611a6490614835565b60405180910390fd5b618000600186611a7d91906142e2565b611a87919061434a565b6001611a939190613feb565b90505f8167ffffffffffffffff168460ff16600160149054906101000a900467ffffffffffffffff1687611ac791906142e2565b67ffffffffffffffff1689611adc919061437a565b611ae6919061437a565b611af0919061437a565b905060015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663e888891533836040518363ffffffff1660e01b8152600401611b4e929190614506565b5f604051808303815f87803b158015611b65575f80fd5b505af1158015611b77573d5f803e3d5ffd5b505050505f6004805490509050611b8c613655565b33815f019073ffffffffffffffffffffffffffffffffffffffff16908173ffffffffffffffffffffffffffffffffffffffff1681525050828160200181815250508881604001818152505087816060019067ffffffffffffffff16908167ffffffffffffffff1681525050600160149054906101000a900467ffffffffffffffff16816080019067ffffffffffffffff16908167ffffffffffffffff1681525050868160a0019067ffffffffffffffff16908167ffffffffffffffff1681525050858160c0019060ff16908160ff1681525050848160e0019060ff16908160ff1681525050600481908060018154018082558091505060019003905f5260205f2090600402015f909190919091505f820151815f015f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060208201518160010155604082015181600201556060820151816003015f6101000a81548167ffffffffffffffff021916908367ffffffffffffffff16021790555060808201518160030160086101000a81548167ffffffffffffffff021916908367ffffffffffffffff16021790555060a08201518160030160106101000a81548167ffffffffffffffff021916908367ffffffffffffffff16021790555060c08201518160030160186101000a81548160ff021916908360ff16021790555060e08201518160030160196101000a81548160ff021916908360ff16021790555050508160058b604051611dd39190614158565b90815260200160405180910390205f6101000a81548167ffffffffffffffff021916908367ffffffffffffffff1602179055503373ffffffffffffffffffffffffffffffffffffffff167f09ff0fb5f488978fdffbf3223e63e914f2a9b0c228844b8e2caedad8a85d220c83600160149054906101000a900467ffffffffffffffff16604051611e6492919061452d565b60405180910390a250505050505050505050565b5f600c5f8473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f8367ffffffffffffffff1667ffffffffffffffff1681526020019081526020015f205f015f9054906101000a900467ffffffffffffffff16905092915050565b5f600b5f8473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f8367ffffffffffffffff1667ffffffffffffffff1681526020019081526020015f205f9054906101000a900467ffffffffffffffff16905092915050565b60025f9054906101000a900467ffffffffffffffff1681565b611f9561349c565b611f9e5f61351a565b565b611fa8612b71565b3373ffffffffffffffffffffffffffffffffffffffff1660048267ffffffffffffffff1681548110611fdd57611fdc613f91565b5b905f5260205f2090600402015f015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1614612061576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016120589061489d565b60405180910390fd5b600160149054906101000a900467ffffffffffffffff1667ffffffffffffffff1660025f9054906101000a900467ffffffffffffffff1660048367ffffffffffffffff16815481106120b6576120b5613f91565b5b905f5260205f20906004020160030160089054906101000a900467ffffffffffffffff166120e49190613feb565b67ffffffffffffffff161061212e576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016121259061459e565b60405180910390fd5b60015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166376890c583360048467ffffffffffffffff168154811061218a57612189613f91565b5b905f5260205f209060040201600101546040518363ffffffff1660e01b81526004016121b7929190614506565b5f604051808303815f87803b1580156121ce575f80fd5b505af11580156121e0573d5f803e3d5ffd5b505050503373ffffffffffffffffffffffffffffffffffffffff167fb283270b87db7ad5d1fbb15af2039324aa28bebf00c89e37579882f7cb261d198260048467ffffffffffffffff168154811061223b5761223a613f91565b5b905f5260205f209060040201600101546040516122599291906145bc565b60405180910390a25f60048267ffffffffffffffff16815481106122805761227f613f91565b5b905f5260205f2090600402016001018190555050565b60015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b6122c3613655565b60048267ffffffffffffffff16815481106122e1576122e0613f91565b5b905f5260205f2090600402015f015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff16815f019073ffffffffffffffffffffffffffffffffffffffff16908173ffffffffffffffffffffffffffffffffffffffff168152505060048267ffffffffffffffff168154811061236357612362613f91565b5b905f5260205f2090600402016002015481604001818152505060048267ffffffffffffffff168154811061239a57612399613f91565b5b905f5260205f2090600402016003015f9054906101000a900467ffffffffffffffff16816060019067ffffffffffffffff16908167ffffffffffffffff168152505060048267ffffffffffffffff16815481106123fa576123f9613f91565b5b905f5260205f20906004020160030160089054906101000a900467ffffffffffffffff16816080019067ffffffffffffffff16908167ffffffffffffffff168152505060048267ffffffffffffffff168154811061245b5761245a613f91565b5b905f5260205f20906004020160030160109054906101000a900467ffffffffffffffff168160a0019067ffffffffffffffff16908167ffffffffffffffff168152505060048267ffffffffffffffff16815481106124bc576124bb613f91565b5b905f5260205f20906004020160030160189054906101000a900460ff168160c0019060ff16908160ff168152505060048267ffffffffffffffff168154811061250857612507613f91565b5b905f5260205f20906004020160030160199054906101000a900460ff168160e0019060ff16908160ff1681525050919050565b600260089054906101000a900467ffffffffffffffff1681565b5f600b5f8573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f8467ffffffffffffffff1667ffffffffffffffff1681526020019081526020015f205f9054906101000a900467ffffffffffffffff1692508267ffffffffffffffff166009836040516125e79190614158565b90815260200160405180910390205f9054906101000a900467ffffffffffffffff1667ffffffffffffffff1614612653576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161264a90614905565b60405180910390fd5b60088367ffffffffffffffff168154811061267157612670613f91565b5b905f5260205f2090600202015f015f9054906101000a900467ffffffffffffffff16925060048367ffffffffffffffff16815481106126b3576126b2613f91565b5b905f5260205f20906004020160030160109054906101000a900467ffffffffffffffff1690509392505050565b5f805f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b61270f6135e2565b604051806060016040528060088467ffffffffffffffff168154811061273857612737613f91565b5b905f5260205f2090600202015f015f9054906101000a900467ffffffffffffffff1667ffffffffffffffff16815260200160088467ffffffffffffffff168154811061278757612786613f91565b5b905f5260205f2090600202015f0160089054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200160088467ffffffffffffffff16815481106127ef576127ee613f91565b5b905f5260205f209060020201600101548152509050919050565b600160149054906101000a900467ffffffffffffffff1681565b5f805f805f60048767ffffffffffffffff168154811061284657612845613f91565b5b905f5260205f20906004020160030160199054906101000a900460ff1660ff16601f61287291906142a6565b600160048967ffffffffffffffff168154811061289257612891613f91565b5b905f5260205f2090600402016003015f9054906101000a900467ffffffffffffffff166128bf91906142e2565b6128c9919061434a565b60016128d59190613feb565b905060048767ffffffffffffffff16815481106128f5576128f4613f91565b5b905f5260205f20906004020160030160189054906101000a900460ff1660048867ffffffffffffffff16815481106129305761292f613f91565b5b905f5260205f20906004020160030160199054906101000a900460ff168260088967ffffffffffffffff168154811061296c5761296b613f91565b5b905f5260205f2090600202016001015494509450945094505092959194509250565b5f60098260405161299f9190614158565b90815260200160405180910390205f9054906101000a900467ffffffffffffffff169050919050565b6129d06136cc565b6040518060400160405280600c5f8673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f8567ffffffffffffffff1667ffffffffffffffff1681526020019081526020015f205f015f9054906101000a900467ffffffffffffffff1667ffffffffffffffff168152602001600c5f8673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f8567ffffffffffffffff1667ffffffffffffffff1681526020019081526020015f2060010154815250905092915050565b60035481565b612add61349c565b5f73ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1603612b4b576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401612b4290614993565b60405180910390fd5b612b548161351a565b50565b600260109054906101000a900467ffffffffffffffff1681565b5f60015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663693ec85e6040518163ffffffff1660e01b8152600401612bca906149fb565b6020604051808303815f875af1158015612be6573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190612c0a919061446a565b90508073ffffffffffffffffffffffffffffffffffffffff1663919840ad6040518163ffffffff1660e01b81526004015f604051808303815f87803b158015612c51575f80fd5b505af1158015612c63573d5f803e3d5ffd5b505050505f8173ffffffffffffffffffffffffffffffffffffffff16639fa6a6e36040518163ffffffff1660e01b81526004016020604051808303815f875af1158015612cb2573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190612cd69190614a2d565b9050600160149054906101000a900467ffffffffffffffff1667ffffffffffffffff168167ffffffffffffffff161115612d345780600160146101000a81548167ffffffffffffffff021916908367ffffffffffffffff1602179055505b5050565b612d4184612f78565b82600b5f8673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f600a5f8873ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f0160089054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681526020019081526020015f205f6101000a81548167ffffffffffffffff021916908367ffffffffffffffff1602179055506001600a5f8673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f0160088282829054906101000a900467ffffffffffffffff16612e7d9190613feb565b92506101000a81548167ffffffffffffffff021916908367ffffffffffffffff16021790555080600a5f8673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f206001015f828254612ef29190614a58565b9250508190555080600c5f8673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f8467ffffffffffffffff1667ffffffffffffffff1681526020019081526020015f206001015f828254612f6b9190614a58565b9250508190555050505050565b5f600a5f8373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f015f9054906101000a900467ffffffffffffffff1667ffffffffffffffff160361305757600160149054906101000a900467ffffffffffffffff16600a5f8373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f015f6101000a81548167ffffffffffffffff021916908367ffffffffffffffff1602179055505b600160149054906101000a900467ffffffffffffffff1667ffffffffffffffff16600a5f8373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f015f9054906101000a900467ffffffffffffffff1667ffffffffffffffff161015613499575f6001600a5f8473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f015f9054906101000a900467ffffffffffffffff1661313b9190613feb565b90505b600160149054906101000a900467ffffffffffffffff1667ffffffffffffffff168167ffffffffffffffff161161341c57600c5f8373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f8267ffffffffffffffff1667ffffffffffffffff1681526020019081526020015f2060010154600a5f8473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f206001015461321e91906143bb565b600c5f8473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f8367ffffffffffffffff1667ffffffffffffffff1681526020019081526020015f2060010181905550600a5f8373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f0160089054906101000a900467ffffffffffffffff16600c5f8473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f8367ffffffffffffffff1667ffffffffffffffff1681526020019081526020015f205f015f6101000a81548167ffffffffffffffff021916908367ffffffffffffffff160217905550600c5f8373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f8267ffffffffffffffff1667ffffffffffffffff1681526020019081526020015f2060010154600a5f8473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f20600101819055506001816134159190613feb565b905061313e565b50600160149054906101000a900467ffffffffffffffff16600a5f8373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f015f6101000a81548167ffffffffffffffff021916908367ffffffffffffffff1602179055505b50565b6134a46135db565b73ffffffffffffffffffffffffffffffffffffffff166134c26126e0565b73ffffffffffffffffffffffffffffffffffffffff1614613518576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161350f90614ad5565b60405180910390fd5b565b5f805f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050815f806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508173ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a35050565b5f33905090565b60405180606001604052805f67ffffffffffffffff1681526020015f73ffffffffffffffffffffffffffffffffffffffff1681526020015f80191681525090565b60405180606001604052805f67ffffffffffffffff1681526020015f67ffffffffffffffff1681526020015f81525090565b6040518061010001604052805f73ffffffffffffffffffffffffffffffffffffffff1681526020015f81526020015f81526020015f67ffffffffffffffff1681526020015f67ffffffffffffffff1681526020015f67ffffffffffffffff1681526020015f60ff1681526020015f60ff1681525090565b60405180604001604052805f67ffffffffffffffff1681526020015f81525090565b5f604051905090565b5f80fd5b5f80fd5b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f613728826136ff565b9050919050565b6137388161371e565b8114613742575f80fd5b50565b5f813590506137538161372f565b92915050565b5f67ffffffffffffffff82169050919050565b61377581613759565b811461377f575f80fd5b50565b5f813590506137908161376c565b92915050565b5f80604083850312156137ac576137ab6136f7565b5b5f6137b985828601613745565b92505060206137ca85828601613782565b9150509250929050565b5f819050919050565b6137e6816137d4565b82525050565b5f6020820190506137ff5f8301846137dd565b92915050565b5f80fd5b5f80fd5b5f601f19601f8301169050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b6138538261380d565b810181811067ffffffffffffffff821117156138725761387161381d565b5b80604052505050565b5f6138846136ee565b9050613890828261384a565b919050565b5f67ffffffffffffffff8211156138af576138ae61381d565b5b6138b88261380d565b9050602081019050919050565b828183375f83830152505050565b5f6138e56138e084613895565b61387b565b90508281526020810184848401111561390157613900613809565b5b61390c8482856138c5565b509392505050565b5f82601f83011261392857613927613805565b5b81356139388482602086016138d3565b91505092915050565b5f60ff82169050919050565b61395681613941565b8114613960575f80fd5b50565b5f813590506139718161394d565b92915050565b5f80fd5b5f80fd5b5f8083601f84011261399457613993613805565b5b8235905067ffffffffffffffff8111156139b1576139b0613977565b5b6020830191508360018202830111156139cd576139cc61397b565b5b9250929050565b5f805f805f608086880312156139ed576139ec6136f7565b5b5f86013567ffffffffffffffff811115613a0a57613a096136fb565b5b613a1688828901613914565b9550506020613a2788828901613782565b9450506040613a3888828901613963565b935050606086013567ffffffffffffffff811115613a5957613a586136fb565b5b613a658882890161397f565b92509250509295509295909350565b5f60208284031215613a8957613a886136f7565b5b5f82013567ffffffffffffffff811115613aa657613aa56136fb565b5b613ab284828501613914565b91505092915050565b613ac481613759565b82525050565b5f602082019050613add5f830184613abb565b92915050565b5f805f60608486031215613afa57613af96136f7565b5b5f613b0786828701613782565b9350506020613b1886828701613782565b9250506040613b2986828701613963565b9150509250925092565b613b3c8161371e565b82525050565b5f602082019050613b555f830184613b33565b92915050565b5f60208284031215613b7057613b6f6136f7565b5b5f613b7d84828501613745565b91505092915050565b613b8f81613759565b82525050565b613b9e816137d4565b82525050565b606082015f820151613bb85f850182613b86565b506020820151613bcb6020850182613b86565b506040820151613bde6040850182613b95565b50505050565b5f606082019050613bf75f830184613ba4565b92915050565b5f60208284031215613c1257613c116136f7565b5b5f613c1f84828501613782565b91505092915050565b613c31816137d4565b8114613c3b575f80fd5b50565b5f81359050613c4c81613c28565b92915050565b5f805f805f8060c08789031215613c6c57613c6b6136f7565b5b5f87013567ffffffffffffffff811115613c8957613c886136fb565b5b613c9589828a01613914565b9650506020613ca689828a01613c3e565b9550506040613cb789828a01613782565b9450506060613cc889828a01613782565b9350506080613cd989828a01613963565b92505060a0613cea89828a01613963565b9150509295509295509295565b613d008161371e565b82525050565b613d0f81613941565b82525050565b61010082015f820151613d2a5f850182613cf7565b506020820151613d3d6020850182613b95565b506040820151613d506040850182613b95565b506060820151613d636060850182613b86565b506080820151613d766080850182613b86565b5060a0820151613d8960a0850182613b86565b5060c0820151613d9c60c0850182613d06565b5060e0820151613daf60e0850182613d06565b50505050565b5f61010082019050613dc95f830184613d15565b92915050565b5f805f60608486031215613de657613de56136f7565b5b5f613df386828701613745565b9350506020613e0486828701613782565b925050604084013567ffffffffffffffff811115613e2557613e246136fb565b5b613e3186828701613914565b9150509250925092565b5f819050919050565b613e4d81613e3b565b82525050565b606082015f820151613e675f850182613b86565b506020820151613e7a6020850182613cf7565b506040820151613e8d6040850182613e44565b50505050565b5f606082019050613ea65f830184613e53565b92915050565b5f8060408385031215613ec257613ec16136f7565b5b5f613ecf85828601613782565b9250506020613ee085828601613782565b9150509250929050565b613ef381613941565b82525050565b613f0281613e3b565b82525050565b5f608082019050613f1b5f830187613eea565b613f286020830186613eea565b613f356040830185613abb565b613f426060830184613ef9565b95945050505050565b604082015f820151613f5f5f850182613b86565b506020820151613f726020850182613b95565b50505050565b5f604082019050613f8b5f830184613f4b565b92915050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52603260045260245ffd5b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b5f613ff582613759565b915061400083613759565b9250828201905067ffffffffffffffff8111156140205761401f613fbe565b5b92915050565b5f82825260208201905092915050565b7f746f6f206c6174650000000000000000000000000000000000000000000000005f82015250565b5f61406a600883614026565b915061407582614036565b602082019050919050565b5f6020820190508181035f8301526140978161405e565b9050919050565b7f696e7620707269000000000000000000000000000000000000000000000000005f82015250565b5f6140d2600783614026565b91506140dd8261409e565b602082019050919050565b5f6020820190508181035f8301526140ff816140c6565b9050919050565b5f81519050919050565b5f81905092915050565b8281835e5f83830152505050565b5f61413282614106565b61413c8185614110565b935061414c81856020860161411a565b80840191505092915050565b5f6141638284614128565b915081905092915050565b7f65786973742072000000000000000000000000000000000000000000000000005f82015250565b5f6141a2600783614026565b91506141ad8261416e565b602082019050919050565b5f6020820190508181035f8301526141cf81614196565b9050919050565b7f65786973742070726900000000000000000000000000000000000000000000005f82015250565b5f61420a600983614026565b9150614215826141d6565b602082019050919050565b5f6020820190508181035f830152614237816141fe565b9050919050565b7f64757020616464720000000000000000000000000000000000000000000000005f82015250565b5f614272600883614026565b915061427d8261423e565b602082019050919050565b5f6020820190508181035f83015261429f81614266565b9050919050565b5f6142b082613759565b91506142bb83613759565b92508282026142c981613759565b91508082146142db576142da613fbe565b5b5092915050565b5f6142ec82613759565b91506142f783613759565b9250828203905067ffffffffffffffff81111561431757614316613fbe565b5b92915050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601260045260245ffd5b5f61435482613759565b915061435f83613759565b92508261436f5761436e61431d565b5b828204905092915050565b5f614384826137d4565b915061438f836137d4565b925082820261439d816137d4565b915082820484148315176143b4576143b3613fbe565b5b5092915050565b5f6143c5826137d4565b91506143d0836137d4565b92508282039050818111156143e8576143e7613fbe565b5b92915050565b7f636f6e74726f6c000000000000000000000000000000000000000000000000005f82015250565b5f614422600783614026565b915061442d826143ee565b602082019050919050565b5f6020820190508181035f83015261444f81614416565b9050919050565b5f815190506144648161372f565b92915050565b5f6020828403121561447f5761447e6136f7565b5b5f61448c84828501614456565b91505092915050565b5f6060820190506144a85f830186613b33565b6144b56020830185613abb565b6144c260408301846137dd565b949350505050565b5f6144d58385614110565b93506144e28385846138c5565b82840190509392505050565b5f6144fa8284866144ca565b91508190509392505050565b5f6040820190506145195f830185613b33565b61452660208301846137dd565b9392505050565b5f6040820190506145405f830185613abb565b61454d6020830184613abb565b9392505050565b7f6561726c790000000000000000000000000000000000000000000000000000005f82015250565b5f614588600583614026565b915061459382614554565b602082019050919050565b5f6020820190508181035f8301526145b58161457c565b9050919050565b5f6040820190506145cf5f830185613abb565b6145dc60208301846137dd565b9392505050565b7f696e762072736b000000000000000000000000000000000000000000000000005f82015250565b5f614617600783614026565b9150614622826145e3565b602082019050919050565b5f6020820190508181035f8301526146448161460b565b9050919050565b7f696e7620727300000000000000000000000000000000000000000000000000005f82015250565b5f61467f600683614026565b915061468a8261464b565b602082019050919050565b5f6020820190508181035f8301526146ac81614673565b9050919050565b7f73686f72742064757200000000000000000000000000000000000000000000005f82015250565b5f6146e7600983614026565b91506146f2826146b3565b602082019050919050565b5f6020820190508181035f830152614714816146db565b9050919050565b7f65786365656420647572000000000000000000000000000000000000000000005f82015250565b5f61474f600a83614026565b915061475a8261471b565b602082019050919050565b5f6020820190508181035f83015261477c81614743565b9050919050565b7f6d696e20707269636500000000000000000000000000000000000000000000005f82015250565b5f6147b7600983614026565b91506147c282614783565b602082019050919050565b5f6020820190508181035f8301526147e4816147ab565b9050919050565b7f6578636565642073697a650000000000000000000000000000000000000000005f82015250565b5f61481f600b83614026565b915061482a826147eb565b602082019050919050565b5f6020820190508181035f83015261484c81614813565b9050919050565b7f696e7620776300000000000000000000000000000000000000000000000000005f82015250565b5f614887600683614026565b915061489282614853565b602082019050919050565b5f6020820190508181035f8301526148b48161487b565b9050919050565b7f696e7620737200000000000000000000000000000000000000000000000000005f82015250565b5f6148ef600683614026565b91506148fa826148bb565b602082019050919050565b5f6020820190508181035f83015261491c816148e3565b9050919050565b7f4f776e61626c653a206e6577206f776e657220697320746865207a65726f20615f8201527f6464726573730000000000000000000000000000000000000000000000000000602082015250565b5f61497d602683614026565b915061498882614923565b604082019050919050565b5f6020820190508181035f8301526149aa81614971565b9050919050565b7f65706f63680000000000000000000000000000000000000000000000000000005f82015250565b5f6149e5600583614026565b91506149f0826149b1565b602082019050919050565b5f6020820190508181035f830152614a12816149d9565b9050919050565b5f81519050614a278161376c565b92915050565b5f60208284031215614a4257614a416136f7565b5b5f614a4f84828501614a19565b91505092915050565b5f614a62826137d4565b9150614a6d836137d4565b9250828201905080821115614a8557614a84613fbe565b5b92915050565b7f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e65725f82015250565b5f614abf602083614026565b9150614aca82614a8b565b602082019050919050565b5f6020820190508181035f830152614aec81614ab3565b905091905056fea2646970667358221220ae500e2de28575922d922bc6cb299ab86aef5a41f7742f63e8f12a8d20dbab3d64736f6c634300081a0033",
}

// PieceABI is the input ABI used to generate the binding from.
// Deprecated: Use PieceMetaData.ABI instead.
var PieceABI = PieceMetaData.ABI

// PieceBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use PieceMetaData.Bin instead.
var PieceBin = PieceMetaData.Bin

// DeployPiece deploys a new Ethereum contract, binding an instance of Piece to it.
func DeployPiece(auth *bind.TransactOpts, backend bind.ContractBackend, _b common.Address) (common.Address, *types.Transaction, *Piece, error) {
	parsed, err := PieceMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(PieceBin), backend, _b)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Piece{PieceCaller: PieceCaller{contract: contract}, PieceTransactor: PieceTransactor{contract: contract}, PieceFilterer: PieceFilterer{contract: contract}}, nil
}

// Piece is an auto generated Go binding around an Ethereum contract.
type Piece struct {
	PieceCaller     // Read-only binding to the contract
	PieceTransactor // Write-only binding to the contract
	PieceFilterer   // Log filterer for contract events
}

// PieceCaller is an auto generated read-only Go binding around an Ethereum contract.
type PieceCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PieceTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PieceTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PieceFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PieceFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PieceSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PieceSession struct {
	Contract     *Piece            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PieceCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PieceCallerSession struct {
	Contract *PieceCaller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// PieceTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PieceTransactorSession struct {
	Contract     *PieceTransactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PieceRaw is an auto generated low-level Go binding around an Ethereum contract.
type PieceRaw struct {
	Contract *Piece // Generic contract binding to access the raw methods on
}

// PieceCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PieceCallerRaw struct {
	Contract *PieceCaller // Generic read-only contract binding to access the raw methods on
}

// PieceTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PieceTransactorRaw struct {
	Contract *PieceTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPiece creates a new instance of Piece, bound to a specific deployed contract.
func NewPiece(address common.Address, backend bind.ContractBackend) (*Piece, error) {
	contract, err := bindPiece(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Piece{PieceCaller: PieceCaller{contract: contract}, PieceTransactor: PieceTransactor{contract: contract}, PieceFilterer: PieceFilterer{contract: contract}}, nil
}

// NewPieceCaller creates a new read-only instance of Piece, bound to a specific deployed contract.
func NewPieceCaller(address common.Address, caller bind.ContractCaller) (*PieceCaller, error) {
	contract, err := bindPiece(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PieceCaller{contract: contract}, nil
}

// NewPieceTransactor creates a new write-only instance of Piece, bound to a specific deployed contract.
func NewPieceTransactor(address common.Address, transactor bind.ContractTransactor) (*PieceTransactor, error) {
	contract, err := bindPiece(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PieceTransactor{contract: contract}, nil
}

// NewPieceFilterer creates a new log filterer instance of Piece, bound to a specific deployed contract.
func NewPieceFilterer(address common.Address, filterer bind.ContractFilterer) (*PieceFilterer, error) {
	contract, err := bindPiece(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PieceFilterer{contract: contract}, nil
}

// bindPiece binds a generic wrapper to an already deployed contract.
func bindPiece(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := PieceMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Piece *PieceRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Piece.Contract.PieceCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Piece *PieceRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Piece.Contract.PieceTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Piece *PieceRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Piece.Contract.PieceTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Piece *PieceCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Piece.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Piece *PieceTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Piece.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Piece *PieceTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Piece.Contract.contract.Transact(opts, method, params...)
}

// Bank is a free data retrieval call binding the contract method 0x76cdb03b.
//
// Solidity: function bank() view returns(address)
func (_Piece *PieceCaller) Bank(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Piece.contract.Call(opts, &out, "bank")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Bank is a free data retrieval call binding the contract method 0x76cdb03b.
//
// Solidity: function bank() view returns(address)
func (_Piece *PieceSession) Bank() (common.Address, error) {
	return _Piece.Contract.Bank(&_Piece.CallOpts)
}

// Bank is a free data retrieval call binding the contract method 0x76cdb03b.
//
// Solidity: function bank() view returns(address)
func (_Piece *PieceCallerSession) Bank() (common.Address, error) {
	return _Piece.Contract.Bank(&_Piece.CallOpts)
}

// CheckPRI is a free data retrieval call binding the contract method 0x2700f3a8.
//
// Solidity: function checkPRI(uint64 _pi, uint64 _ri, uint8 _pri) view returns(address)
func (_Piece *PieceCaller) CheckPRI(opts *bind.CallOpts, _pi uint64, _ri uint64, _pri uint8) (common.Address, error) {
	var out []interface{}
	err := _Piece.contract.Call(opts, &out, "checkPRI", _pi, _ri, _pri)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// CheckPRI is a free data retrieval call binding the contract method 0x2700f3a8.
//
// Solidity: function checkPRI(uint64 _pi, uint64 _ri, uint8 _pri) view returns(address)
func (_Piece *PieceSession) CheckPRI(_pi uint64, _ri uint64, _pri uint8) (common.Address, error) {
	return _Piece.Contract.CheckPRI(&_Piece.CallOpts, _pi, _ri, _pri)
}

// CheckPRI is a free data retrieval call binding the contract method 0x2700f3a8.
//
// Solidity: function checkPRI(uint64 _pi, uint64 _ri, uint8 _pri) view returns(address)
func (_Piece *PieceCallerSession) CheckPRI(_pi uint64, _ri uint64, _pri uint8) (common.Address, error) {
	return _Piece.Contract.CheckPRI(&_Piece.CallOpts, _pi, _ri, _pri)
}

// CheckSRAt is a free data retrieval call binding the contract method 0x67439a2d.
//
// Solidity: function checkSRAt(address _a, uint64 _ri) view returns(uint64)
func (_Piece *PieceCaller) CheckSRAt(opts *bind.CallOpts, _a common.Address, _ri uint64) (uint64, error) {
	var out []interface{}
	err := _Piece.contract.Call(opts, &out, "checkSRAt", _a, _ri)

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// CheckSRAt is a free data retrieval call binding the contract method 0x67439a2d.
//
// Solidity: function checkSRAt(address _a, uint64 _ri) view returns(uint64)
func (_Piece *PieceSession) CheckSRAt(_a common.Address, _ri uint64) (uint64, error) {
	return _Piece.Contract.CheckSRAt(&_Piece.CallOpts, _a, _ri)
}

// CheckSRAt is a free data retrieval call binding the contract method 0x67439a2d.
//
// Solidity: function checkSRAt(address _a, uint64 _ri) view returns(uint64)
func (_Piece *PieceCallerSession) CheckSRAt(_a common.Address, _ri uint64) (uint64, error) {
	return _Piece.Contract.CheckSRAt(&_Piece.CallOpts, _a, _ri)
}

// CheckSReplica is a free data retrieval call binding the contract method 0x894d35f8.
//
// Solidity: function checkSReplica(address _a, uint64 _ri, bytes _com) view returns(uint64)
func (_Piece *PieceCaller) CheckSReplica(opts *bind.CallOpts, _a common.Address, _ri uint64, _com []byte) (uint64, error) {
	var out []interface{}
	err := _Piece.contract.Call(opts, &out, "checkSReplica", _a, _ri, _com)

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// CheckSReplica is a free data retrieval call binding the contract method 0x894d35f8.
//
// Solidity: function checkSReplica(address _a, uint64 _ri, bytes _com) view returns(uint64)
func (_Piece *PieceSession) CheckSReplica(_a common.Address, _ri uint64, _com []byte) (uint64, error) {
	return _Piece.Contract.CheckSReplica(&_Piece.CallOpts, _a, _ri, _com)
}

// CheckSReplica is a free data retrieval call binding the contract method 0x894d35f8.
//
// Solidity: function checkSReplica(address _a, uint64 _ri, bytes _com) view returns(uint64)
func (_Piece *PieceCallerSession) CheckSReplica(_a common.Address, _ri uint64, _com []byte) (uint64, error) {
	return _Piece.Contract.CheckSReplica(&_Piece.CallOpts, _a, _ri, _com)
}

// Current is a free data retrieval call binding the contract method 0x9fa6a6e3.
//
// Solidity: function current() view returns(uint64)
func (_Piece *PieceCaller) Current(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _Piece.contract.Call(opts, &out, "current")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// Current is a free data retrieval call binding the contract method 0x9fa6a6e3.
//
// Solidity: function current() view returns(uint64)
func (_Piece *PieceSession) Current() (uint64, error) {
	return _Piece.Contract.Current(&_Piece.CallOpts)
}

// Current is a free data retrieval call binding the contract method 0x9fa6a6e3.
//
// Solidity: function current() view returns(uint64)
func (_Piece *PieceCallerSession) Current() (uint64, error) {
	return _Piece.Contract.Current(&_Piece.CallOpts)
}

// Delay is a free data retrieval call binding the contract method 0x6a42b8f8.
//
// Solidity: function delay() view returns(uint64)
func (_Piece *PieceCaller) Delay(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _Piece.contract.Call(opts, &out, "delay")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// Delay is a free data retrieval call binding the contract method 0x6a42b8f8.
//
// Solidity: function delay() view returns(uint64)
func (_Piece *PieceSession) Delay() (uint64, error) {
	return _Piece.Contract.Delay(&_Piece.CallOpts)
}

// Delay is a free data retrieval call binding the contract method 0x6a42b8f8.
//
// Solidity: function delay() view returns(uint64)
func (_Piece *PieceCallerSession) Delay() (uint64, error) {
	return _Piece.Contract.Delay(&_Piece.CallOpts)
}

// GetPIndex is a free data retrieval call binding the contract method 0x1ce85e7c.
//
// Solidity: function getPIndex(bytes _pn) view returns(uint64)
func (_Piece *PieceCaller) GetPIndex(opts *bind.CallOpts, _pn []byte) (uint64, error) {
	var out []interface{}
	err := _Piece.contract.Call(opts, &out, "getPIndex", _pn)

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// GetPIndex is a free data retrieval call binding the contract method 0x1ce85e7c.
//
// Solidity: function getPIndex(bytes _pn) view returns(uint64)
func (_Piece *PieceSession) GetPIndex(_pn []byte) (uint64, error) {
	return _Piece.Contract.GetPIndex(&_Piece.CallOpts, _pn)
}

// GetPIndex is a free data retrieval call binding the contract method 0x1ce85e7c.
//
// Solidity: function getPIndex(bytes _pn) view returns(uint64)
func (_Piece *PieceCallerSession) GetPIndex(_pn []byte) (uint64, error) {
	return _Piece.Contract.GetPIndex(&_Piece.CallOpts, _pn)
}

// GetPRInfo is a free data retrieval call binding the contract method 0xb650b504.
//
// Solidity: function getPRInfo(uint64 _pi, uint64 _ri) view returns(uint8, uint8, uint64, bytes32)
func (_Piece *PieceCaller) GetPRInfo(opts *bind.CallOpts, _pi uint64, _ri uint64) (uint8, uint8, uint64, [32]byte, error) {
	var out []interface{}
	err := _Piece.contract.Call(opts, &out, "getPRInfo", _pi, _ri)

	if err != nil {
		return *new(uint8), *new(uint8), *new(uint64), *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)
	out1 := *abi.ConvertType(out[1], new(uint8)).(*uint8)
	out2 := *abi.ConvertType(out[2], new(uint64)).(*uint64)
	out3 := *abi.ConvertType(out[3], new([32]byte)).(*[32]byte)

	return out0, out1, out2, out3, err

}

// GetPRInfo is a free data retrieval call binding the contract method 0xb650b504.
//
// Solidity: function getPRInfo(uint64 _pi, uint64 _ri) view returns(uint8, uint8, uint64, bytes32)
func (_Piece *PieceSession) GetPRInfo(_pi uint64, _ri uint64) (uint8, uint8, uint64, [32]byte, error) {
	return _Piece.Contract.GetPRInfo(&_Piece.CallOpts, _pi, _ri)
}

// GetPRInfo is a free data retrieval call binding the contract method 0xb650b504.
//
// Solidity: function getPRInfo(uint64 _pi, uint64 _ri) view returns(uint8, uint8, uint64, bytes32)
func (_Piece *PieceCallerSession) GetPRInfo(_pi uint64, _ri uint64) (uint8, uint8, uint64, [32]byte, error) {
	return _Piece.Contract.GetPRInfo(&_Piece.CallOpts, _pi, _ri)
}

// GetPiece is a free data retrieval call binding the contract method 0x7a87f0c4.
//
// Solidity: function getPiece(uint64 _pi) view returns((address,uint256,uint256,uint64,uint64,uint64,uint8,uint8) _pinfo)
func (_Piece *PieceCaller) GetPiece(opts *bind.CallOpts, _pi uint64) (IPiecePieceInfo, error) {
	var out []interface{}
	err := _Piece.contract.Call(opts, &out, "getPiece", _pi)

	if err != nil {
		return *new(IPiecePieceInfo), err
	}

	out0 := *abi.ConvertType(out[0], new(IPiecePieceInfo)).(*IPiecePieceInfo)

	return out0, err

}

// GetPiece is a free data retrieval call binding the contract method 0x7a87f0c4.
//
// Solidity: function getPiece(uint64 _pi) view returns((address,uint256,uint256,uint64,uint64,uint64,uint8,uint8) _pinfo)
func (_Piece *PieceSession) GetPiece(_pi uint64) (IPiecePieceInfo, error) {
	return _Piece.Contract.GetPiece(&_Piece.CallOpts, _pi)
}

// GetPiece is a free data retrieval call binding the contract method 0x7a87f0c4.
//
// Solidity: function getPiece(uint64 _pi) view returns((address,uint256,uint256,uint64,uint64,uint64,uint8,uint8) _pinfo)
func (_Piece *PieceCallerSession) GetPiece(_pi uint64) (IPiecePieceInfo, error) {
	return _Piece.Contract.GetPiece(&_Piece.CallOpts, _pi)
}

// GetRIndex is a free data retrieval call binding the contract method 0xba8b2618.
//
// Solidity: function getRIndex(bytes _rn) view returns(uint64)
func (_Piece *PieceCaller) GetRIndex(opts *bind.CallOpts, _rn []byte) (uint64, error) {
	var out []interface{}
	err := _Piece.contract.Call(opts, &out, "getRIndex", _rn)

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// GetRIndex is a free data retrieval call binding the contract method 0xba8b2618.
//
// Solidity: function getRIndex(bytes _rn) view returns(uint64)
func (_Piece *PieceSession) GetRIndex(_rn []byte) (uint64, error) {
	return _Piece.Contract.GetRIndex(&_Piece.CallOpts, _rn)
}

// GetRIndex is a free data retrieval call binding the contract method 0xba8b2618.
//
// Solidity: function getRIndex(bytes _rn) view returns(uint64)
func (_Piece *PieceCallerSession) GetRIndex(_rn []byte) (uint64, error) {
	return _Piece.Contract.GetRIndex(&_Piece.CallOpts, _rn)
}

// GetReplica is a free data retrieval call binding the contract method 0x9108544c.
//
// Solidity: function getReplica(uint64 _r) view returns((uint64,address,bytes32) _ri)
func (_Piece *PieceCaller) GetReplica(opts *bind.CallOpts, _r uint64) (IPieceReplicaInfo, error) {
	var out []interface{}
	err := _Piece.contract.Call(opts, &out, "getReplica", _r)

	if err != nil {
		return *new(IPieceReplicaInfo), err
	}

	out0 := *abi.ConvertType(out[0], new(IPieceReplicaInfo)).(*IPieceReplicaInfo)

	return out0, err

}

// GetReplica is a free data retrieval call binding the contract method 0x9108544c.
//
// Solidity: function getReplica(uint64 _r) view returns((uint64,address,bytes32) _ri)
func (_Piece *PieceSession) GetReplica(_r uint64) (IPieceReplicaInfo, error) {
	return _Piece.Contract.GetReplica(&_Piece.CallOpts, _r)
}

// GetReplica is a free data retrieval call binding the contract method 0x9108544c.
//
// Solidity: function getReplica(uint64 _r) view returns((uint64,address,bytes32) _ri)
func (_Piece *PieceCallerSession) GetReplica(_r uint64) (IPieceReplicaInfo, error) {
	return _Piece.Contract.GetReplica(&_Piece.CallOpts, _r)
}

// GetSStat is a free data retrieval call binding the contract method 0xc3aa3804.
//
// Solidity: function getSStat(address _a, uint64 _e) view returns((uint64,uint256) _si)
func (_Piece *PieceCaller) GetSStat(opts *bind.CallOpts, _a common.Address, _e uint64) (IPieceStoreStat, error) {
	var out []interface{}
	err := _Piece.contract.Call(opts, &out, "getSStat", _a, _e)

	if err != nil {
		return *new(IPieceStoreStat), err
	}

	out0 := *abi.ConvertType(out[0], new(IPieceStoreStat)).(*IPieceStoreStat)

	return out0, err

}

// GetSStat is a free data retrieval call binding the contract method 0xc3aa3804.
//
// Solidity: function getSStat(address _a, uint64 _e) view returns((uint64,uint256) _si)
func (_Piece *PieceSession) GetSStat(_a common.Address, _e uint64) (IPieceStoreStat, error) {
	return _Piece.Contract.GetSStat(&_Piece.CallOpts, _a, _e)
}

// GetSStat is a free data retrieval call binding the contract method 0xc3aa3804.
//
// Solidity: function getSStat(address _a, uint64 _e) view returns((uint64,uint256) _si)
func (_Piece *PieceCallerSession) GetSStat(_a common.Address, _e uint64) (IPieceStoreStat, error) {
	return _Piece.Contract.GetSStat(&_Piece.CallOpts, _a, _e)
}

// GetStore is a free data retrieval call binding the contract method 0x4b4ffccd.
//
// Solidity: function getStore(address _a) view returns((uint64,uint64,uint256) _si)
func (_Piece *PieceCaller) GetStore(opts *bind.CallOpts, _a common.Address) (IPieceStoreInfo, error) {
	var out []interface{}
	err := _Piece.contract.Call(opts, &out, "getStore", _a)

	if err != nil {
		return *new(IPieceStoreInfo), err
	}

	out0 := *abi.ConvertType(out[0], new(IPieceStoreInfo)).(*IPieceStoreInfo)

	return out0, err

}

// GetStore is a free data retrieval call binding the contract method 0x4b4ffccd.
//
// Solidity: function getStore(address _a) view returns((uint64,uint64,uint256) _si)
func (_Piece *PieceSession) GetStore(_a common.Address) (IPieceStoreInfo, error) {
	return _Piece.Contract.GetStore(&_Piece.CallOpts, _a)
}

// GetStore is a free data retrieval call binding the contract method 0x4b4ffccd.
//
// Solidity: function getStore(address _a) view returns((uint64,uint64,uint256) _si)
func (_Piece *PieceCallerSession) GetStore(_a common.Address) (IPieceStoreInfo, error) {
	return _Piece.Contract.GetStore(&_Piece.CallOpts, _a)
}

// GetStoreCount is a free data retrieval call binding the contract method 0x50979d7e.
//
// Solidity: function getStoreCount(address _a, uint64 _e) view returns(uint64)
func (_Piece *PieceCaller) GetStoreCount(opts *bind.CallOpts, _a common.Address, _e uint64) (uint64, error) {
	var out []interface{}
	err := _Piece.contract.Call(opts, &out, "getStoreCount", _a, _e)

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// GetStoreCount is a free data retrieval call binding the contract method 0x50979d7e.
//
// Solidity: function getStoreCount(address _a, uint64 _e) view returns(uint64)
func (_Piece *PieceSession) GetStoreCount(_a common.Address, _e uint64) (uint64, error) {
	return _Piece.Contract.GetStoreCount(&_Piece.CallOpts, _a, _e)
}

// GetStoreCount is a free data retrieval call binding the contract method 0x50979d7e.
//
// Solidity: function getStoreCount(address _a, uint64 _e) view returns(uint64)
func (_Piece *PieceCallerSession) GetStoreCount(_a common.Address, _e uint64) (uint64, error) {
	return _Piece.Contract.GetStoreCount(&_Piece.CallOpts, _a, _e)
}

// GetStoreSalary is a free data retrieval call binding the contract method 0x18198ad7.
//
// Solidity: function getStoreSalary(address _a, uint64 _e) view returns(uint256)
func (_Piece *PieceCaller) GetStoreSalary(opts *bind.CallOpts, _a common.Address, _e uint64) (*big.Int, error) {
	var out []interface{}
	err := _Piece.contract.Call(opts, &out, "getStoreSalary", _a, _e)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetStoreSalary is a free data retrieval call binding the contract method 0x18198ad7.
//
// Solidity: function getStoreSalary(address _a, uint64 _e) view returns(uint256)
func (_Piece *PieceSession) GetStoreSalary(_a common.Address, _e uint64) (*big.Int, error) {
	return _Piece.Contract.GetStoreSalary(&_Piece.CallOpts, _a, _e)
}

// GetStoreSalary is a free data retrieval call binding the contract method 0x18198ad7.
//
// Solidity: function getStoreSalary(address _a, uint64 _e) view returns(uint256)
func (_Piece *PieceCallerSession) GetStoreSalary(_a common.Address, _e uint64) (*big.Int, error) {
	return _Piece.Contract.GetStoreSalary(&_Piece.CallOpts, _a, _e)
}

// MaxSize is a free data retrieval call binding the contract method 0x2565b159.
//
// Solidity: function maxSize() view returns(uint64)
func (_Piece *PieceCaller) MaxSize(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _Piece.contract.Call(opts, &out, "maxSize")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// MaxSize is a free data retrieval call binding the contract method 0x2565b159.
//
// Solidity: function maxSize() view returns(uint64)
func (_Piece *PieceSession) MaxSize() (uint64, error) {
	return _Piece.Contract.MaxSize(&_Piece.CallOpts)
}

// MaxSize is a free data retrieval call binding the contract method 0x2565b159.
//
// Solidity: function maxSize() view returns(uint64)
func (_Piece *PieceCallerSession) MaxSize() (uint64, error) {
	return _Piece.Contract.MaxSize(&_Piece.CallOpts)
}

// MaxStore is a free data retrieval call binding the contract method 0xfca5be7a.
//
// Solidity: function maxStore() view returns(uint64)
func (_Piece *PieceCaller) MaxStore(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _Piece.contract.Call(opts, &out, "maxStore")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// MaxStore is a free data retrieval call binding the contract method 0xfca5be7a.
//
// Solidity: function maxStore() view returns(uint64)
func (_Piece *PieceSession) MaxStore() (uint64, error) {
	return _Piece.Contract.MaxStore(&_Piece.CallOpts)
}

// MaxStore is a free data retrieval call binding the contract method 0xfca5be7a.
//
// Solidity: function maxStore() view returns(uint64)
func (_Piece *PieceCallerSession) MaxStore() (uint64, error) {
	return _Piece.Contract.MaxStore(&_Piece.CallOpts)
}

// MinPrice is a free data retrieval call binding the contract method 0xe45be8eb.
//
// Solidity: function minPrice() view returns(uint256)
func (_Piece *PieceCaller) MinPrice(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Piece.contract.Call(opts, &out, "minPrice")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MinPrice is a free data retrieval call binding the contract method 0xe45be8eb.
//
// Solidity: function minPrice() view returns(uint256)
func (_Piece *PieceSession) MinPrice() (*big.Int, error) {
	return _Piece.Contract.MinPrice(&_Piece.CallOpts)
}

// MinPrice is a free data retrieval call binding the contract method 0xe45be8eb.
//
// Solidity: function minPrice() view returns(uint256)
func (_Piece *PieceCallerSession) MinPrice() (*big.Int, error) {
	return _Piece.Contract.MinPrice(&_Piece.CallOpts)
}

// MinStore is a free data retrieval call binding the contract method 0x88795125.
//
// Solidity: function minStore() view returns(uint64)
func (_Piece *PieceCaller) MinStore(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _Piece.contract.Call(opts, &out, "minStore")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// MinStore is a free data retrieval call binding the contract method 0x88795125.
//
// Solidity: function minStore() view returns(uint64)
func (_Piece *PieceSession) MinStore() (uint64, error) {
	return _Piece.Contract.MinStore(&_Piece.CallOpts)
}

// MinStore is a free data retrieval call binding the contract method 0x88795125.
//
// Solidity: function minStore() view returns(uint64)
func (_Piece *PieceCallerSession) MinStore() (uint64, error) {
	return _Piece.Contract.MinStore(&_Piece.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Piece *PieceCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Piece.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Piece *PieceSession) Owner() (common.Address, error) {
	return _Piece.Contract.Owner(&_Piece.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Piece *PieceCallerSession) Owner() (common.Address, error) {
	return _Piece.Contract.Owner(&_Piece.CallOpts)
}

// AddPiece is a paid mutator transaction binding the contract method 0x4d9c3d4b.
//
// Solidity: function addPiece(bytes _pn, uint256 _price, uint64 _size, uint64 _expire, uint8 rsn, uint8 rsk) returns()
func (_Piece *PieceTransactor) AddPiece(opts *bind.TransactOpts, _pn []byte, _price *big.Int, _size uint64, _expire uint64, rsn uint8, rsk uint8) (*types.Transaction, error) {
	return _Piece.contract.Transact(opts, "addPiece", _pn, _price, _size, _expire, rsn, rsk)
}

// AddPiece is a paid mutator transaction binding the contract method 0x4d9c3d4b.
//
// Solidity: function addPiece(bytes _pn, uint256 _price, uint64 _size, uint64 _expire, uint8 rsn, uint8 rsk) returns()
func (_Piece *PieceSession) AddPiece(_pn []byte, _price *big.Int, _size uint64, _expire uint64, rsn uint8, rsk uint8) (*types.Transaction, error) {
	return _Piece.Contract.AddPiece(&_Piece.TransactOpts, _pn, _price, _size, _expire, rsn, rsk)
}

// AddPiece is a paid mutator transaction binding the contract method 0x4d9c3d4b.
//
// Solidity: function addPiece(bytes _pn, uint256 _price, uint64 _size, uint64 _expire, uint8 rsn, uint8 rsk) returns()
func (_Piece *PieceTransactorSession) AddPiece(_pn []byte, _price *big.Int, _size uint64, _expire uint64, rsn uint8, rsk uint8) (*types.Transaction, error) {
	return _Piece.Contract.AddPiece(&_Piece.TransactOpts, _pn, _price, _size, _expire, rsn, rsk)
}

// AddReplica is a paid mutator transaction binding the contract method 0x194cbaf5.
//
// Solidity: function addReplica(bytes _rn, uint64 _pi, uint8 _pri, bytes _proof) returns()
func (_Piece *PieceTransactor) AddReplica(opts *bind.TransactOpts, _rn []byte, _pi uint64, _pri uint8, _proof []byte) (*types.Transaction, error) {
	return _Piece.contract.Transact(opts, "addReplica", _rn, _pi, _pri, _proof)
}

// AddReplica is a paid mutator transaction binding the contract method 0x194cbaf5.
//
// Solidity: function addReplica(bytes _rn, uint64 _pi, uint8 _pri, bytes _proof) returns()
func (_Piece *PieceSession) AddReplica(_rn []byte, _pi uint64, _pri uint8, _proof []byte) (*types.Transaction, error) {
	return _Piece.Contract.AddReplica(&_Piece.TransactOpts, _rn, _pi, _pri, _proof)
}

// AddReplica is a paid mutator transaction binding the contract method 0x194cbaf5.
//
// Solidity: function addReplica(bytes _rn, uint64 _pi, uint8 _pri, bytes _proof) returns()
func (_Piece *PieceTransactorSession) AddReplica(_rn []byte, _pi uint64, _pri uint8, _proof []byte) (*types.Transaction, error) {
	return _Piece.Contract.AddReplica(&_Piece.TransactOpts, _rn, _pi, _pri, _proof)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Piece *PieceTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Piece.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Piece *PieceSession) RenounceOwnership() (*types.Transaction, error) {
	return _Piece.Contract.RenounceOwnership(&_Piece.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Piece *PieceTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Piece.Contract.RenounceOwnership(&_Piece.TransactOpts)
}

// Settle is a paid mutator transaction binding the contract method 0x4b9ba446.
//
// Solidity: function settle(uint64 _ep) returns()
func (_Piece *PieceTransactor) Settle(opts *bind.TransactOpts, _ep uint64) (*types.Transaction, error) {
	return _Piece.contract.Transact(opts, "settle", _ep)
}

// Settle is a paid mutator transaction binding the contract method 0x4b9ba446.
//
// Solidity: function settle(uint64 _ep) returns()
func (_Piece *PieceSession) Settle(_ep uint64) (*types.Transaction, error) {
	return _Piece.Contract.Settle(&_Piece.TransactOpts, _ep)
}

// Settle is a paid mutator transaction binding the contract method 0x4b9ba446.
//
// Solidity: function settle(uint64 _ep) returns()
func (_Piece *PieceTransactorSession) Settle(_ep uint64) (*types.Transaction, error) {
	return _Piece.Contract.Settle(&_Piece.TransactOpts, _ep)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Piece *PieceTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Piece.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Piece *PieceSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Piece.Contract.TransferOwnership(&_Piece.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Piece *PieceTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Piece.Contract.TransferOwnership(&_Piece.TransactOpts, newOwner)
}

// Withdraw is a paid mutator transaction binding the contract method 0x750f0acc.
//
// Solidity: function withdraw(uint64 _pi) returns()
func (_Piece *PieceTransactor) Withdraw(opts *bind.TransactOpts, _pi uint64) (*types.Transaction, error) {
	return _Piece.contract.Transact(opts, "withdraw", _pi)
}

// Withdraw is a paid mutator transaction binding the contract method 0x750f0acc.
//
// Solidity: function withdraw(uint64 _pi) returns()
func (_Piece *PieceSession) Withdraw(_pi uint64) (*types.Transaction, error) {
	return _Piece.Contract.Withdraw(&_Piece.TransactOpts, _pi)
}

// Withdraw is a paid mutator transaction binding the contract method 0x750f0acc.
//
// Solidity: function withdraw(uint64 _pi) returns()
func (_Piece *PieceTransactorSession) Withdraw(_pi uint64) (*types.Transaction, error) {
	return _Piece.Contract.Withdraw(&_Piece.TransactOpts, _pi)
}

// PieceAddPieceIterator is returned from FilterAddPiece and is used to iterate over the raw logs and unpacked data for AddPiece events raised by the Piece contract.
type PieceAddPieceIterator struct {
	Event *PieceAddPiece // Event containing the contract specifics and raw log

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
func (it *PieceAddPieceIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PieceAddPiece)
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
		it.Event = new(PieceAddPiece)
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
func (it *PieceAddPieceIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PieceAddPieceIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PieceAddPiece represents a AddPiece event raised by the Piece contract.
type PieceAddPiece struct {
	A   common.Address
	Pi  uint64
	E   uint64
	Raw types.Log // Blockchain specific contextual infos
}

// FilterAddPiece is a free log retrieval operation binding the contract event 0x09ff0fb5f488978fdffbf3223e63e914f2a9b0c228844b8e2caedad8a85d220c.
//
// Solidity: event AddPiece(address indexed _a, uint64 _pi, uint64 _e)
func (_Piece *PieceFilterer) FilterAddPiece(opts *bind.FilterOpts, _a []common.Address) (*PieceAddPieceIterator, error) {

	var _aRule []interface{}
	for _, _aItem := range _a {
		_aRule = append(_aRule, _aItem)
	}

	logs, sub, err := _Piece.contract.FilterLogs(opts, "AddPiece", _aRule)
	if err != nil {
		return nil, err
	}
	return &PieceAddPieceIterator{contract: _Piece.contract, event: "AddPiece", logs: logs, sub: sub}, nil
}

// WatchAddPiece is a free log subscription operation binding the contract event 0x09ff0fb5f488978fdffbf3223e63e914f2a9b0c228844b8e2caedad8a85d220c.
//
// Solidity: event AddPiece(address indexed _a, uint64 _pi, uint64 _e)
func (_Piece *PieceFilterer) WatchAddPiece(opts *bind.WatchOpts, sink chan<- *PieceAddPiece, _a []common.Address) (event.Subscription, error) {

	var _aRule []interface{}
	for _, _aItem := range _a {
		_aRule = append(_aRule, _aItem)
	}

	logs, sub, err := _Piece.contract.WatchLogs(opts, "AddPiece", _aRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PieceAddPiece)
				if err := _Piece.contract.UnpackLog(event, "AddPiece", log); err != nil {
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

// ParseAddPiece is a log parse operation binding the contract event 0x09ff0fb5f488978fdffbf3223e63e914f2a9b0c228844b8e2caedad8a85d220c.
//
// Solidity: event AddPiece(address indexed _a, uint64 _pi, uint64 _e)
func (_Piece *PieceFilterer) ParseAddPiece(log types.Log) (*PieceAddPiece, error) {
	event := new(PieceAddPiece)
	if err := _Piece.contract.UnpackLog(event, "AddPiece", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PieceAddReplicaIterator is returned from FilterAddReplica and is used to iterate over the raw logs and unpacked data for AddReplica events raised by the Piece contract.
type PieceAddReplicaIterator struct {
	Event *PieceAddReplica // Event containing the contract specifics and raw log

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
func (it *PieceAddReplicaIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PieceAddReplica)
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
		it.Event = new(PieceAddReplica)
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
func (it *PieceAddReplicaIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PieceAddReplicaIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PieceAddReplica represents a AddReplica event raised by the Piece contract.
type PieceAddReplica struct {
	A   common.Address
	Ri  uint64
	Sri uint64
	Raw types.Log // Blockchain specific contextual infos
}

// FilterAddReplica is a free log retrieval operation binding the contract event 0x4b079e585085a3c94b2865036e598391fad4359764b36da39e21897be56349ef.
//
// Solidity: event AddReplica(address indexed _a, uint64 _ri, uint64 _sri)
func (_Piece *PieceFilterer) FilterAddReplica(opts *bind.FilterOpts, _a []common.Address) (*PieceAddReplicaIterator, error) {

	var _aRule []interface{}
	for _, _aItem := range _a {
		_aRule = append(_aRule, _aItem)
	}

	logs, sub, err := _Piece.contract.FilterLogs(opts, "AddReplica", _aRule)
	if err != nil {
		return nil, err
	}
	return &PieceAddReplicaIterator{contract: _Piece.contract, event: "AddReplica", logs: logs, sub: sub}, nil
}

// WatchAddReplica is a free log subscription operation binding the contract event 0x4b079e585085a3c94b2865036e598391fad4359764b36da39e21897be56349ef.
//
// Solidity: event AddReplica(address indexed _a, uint64 _ri, uint64 _sri)
func (_Piece *PieceFilterer) WatchAddReplica(opts *bind.WatchOpts, sink chan<- *PieceAddReplica, _a []common.Address) (event.Subscription, error) {

	var _aRule []interface{}
	for _, _aItem := range _a {
		_aRule = append(_aRule, _aItem)
	}

	logs, sub, err := _Piece.contract.WatchLogs(opts, "AddReplica", _aRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PieceAddReplica)
				if err := _Piece.contract.UnpackLog(event, "AddReplica", log); err != nil {
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

// ParseAddReplica is a log parse operation binding the contract event 0x4b079e585085a3c94b2865036e598391fad4359764b36da39e21897be56349ef.
//
// Solidity: event AddReplica(address indexed _a, uint64 _ri, uint64 _sri)
func (_Piece *PieceFilterer) ParseAddReplica(log types.Log) (*PieceAddReplica, error) {
	event := new(PieceAddReplica)
	if err := _Piece.contract.UnpackLog(event, "AddReplica", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PieceOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Piece contract.
type PieceOwnershipTransferredIterator struct {
	Event *PieceOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *PieceOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PieceOwnershipTransferred)
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
		it.Event = new(PieceOwnershipTransferred)
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
func (it *PieceOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PieceOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PieceOwnershipTransferred represents a OwnershipTransferred event raised by the Piece contract.
type PieceOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Piece *PieceFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*PieceOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Piece.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &PieceOwnershipTransferredIterator{contract: _Piece.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Piece *PieceFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *PieceOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Piece.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PieceOwnershipTransferred)
				if err := _Piece.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_Piece *PieceFilterer) ParseOwnershipTransferred(log types.Log) (*PieceOwnershipTransferred, error) {
	event := new(PieceOwnershipTransferred)
	if err := _Piece.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PieceSettleIterator is returned from FilterSettle and is used to iterate over the raw logs and unpacked data for Settle events raised by the Piece contract.
type PieceSettleIterator struct {
	Event *PieceSettle // Event containing the contract specifics and raw log

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
func (it *PieceSettleIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PieceSettle)
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
		it.Event = new(PieceSettle)
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
func (it *PieceSettleIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PieceSettleIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PieceSettle represents a Settle event raised by the Piece contract.
type PieceSettle struct {
	A   common.Address
	E   uint64
	M   *big.Int
	Raw types.Log // Blockchain specific contextual infos
}

// FilterSettle is a free log retrieval operation binding the contract event 0x22b9b452d488700bdfeab2aa62972ac1fce3160583db89ab4c799708bf5a2c6c.
//
// Solidity: event Settle(address indexed _a, uint64 _e, uint256 _m)
func (_Piece *PieceFilterer) FilterSettle(opts *bind.FilterOpts, _a []common.Address) (*PieceSettleIterator, error) {

	var _aRule []interface{}
	for _, _aItem := range _a {
		_aRule = append(_aRule, _aItem)
	}

	logs, sub, err := _Piece.contract.FilterLogs(opts, "Settle", _aRule)
	if err != nil {
		return nil, err
	}
	return &PieceSettleIterator{contract: _Piece.contract, event: "Settle", logs: logs, sub: sub}, nil
}

// WatchSettle is a free log subscription operation binding the contract event 0x22b9b452d488700bdfeab2aa62972ac1fce3160583db89ab4c799708bf5a2c6c.
//
// Solidity: event Settle(address indexed _a, uint64 _e, uint256 _m)
func (_Piece *PieceFilterer) WatchSettle(opts *bind.WatchOpts, sink chan<- *PieceSettle, _a []common.Address) (event.Subscription, error) {

	var _aRule []interface{}
	for _, _aItem := range _a {
		_aRule = append(_aRule, _aItem)
	}

	logs, sub, err := _Piece.contract.WatchLogs(opts, "Settle", _aRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PieceSettle)
				if err := _Piece.contract.UnpackLog(event, "Settle", log); err != nil {
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

// ParseSettle is a log parse operation binding the contract event 0x22b9b452d488700bdfeab2aa62972ac1fce3160583db89ab4c799708bf5a2c6c.
//
// Solidity: event Settle(address indexed _a, uint64 _e, uint256 _m)
func (_Piece *PieceFilterer) ParseSettle(log types.Log) (*PieceSettle, error) {
	event := new(PieceSettle)
	if err := _Piece.contract.UnpackLog(event, "Settle", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PieceWithdrawIterator is returned from FilterWithdraw and is used to iterate over the raw logs and unpacked data for Withdraw events raised by the Piece contract.
type PieceWithdrawIterator struct {
	Event *PieceWithdraw // Event containing the contract specifics and raw log

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
func (it *PieceWithdrawIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PieceWithdraw)
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
		it.Event = new(PieceWithdraw)
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
func (it *PieceWithdrawIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PieceWithdrawIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PieceWithdraw represents a Withdraw event raised by the Piece contract.
type PieceWithdraw struct {
	A   common.Address
	Pi  uint64
	M   *big.Int
	Raw types.Log // Blockchain specific contextual infos
}

// FilterWithdraw is a free log retrieval operation binding the contract event 0xb283270b87db7ad5d1fbb15af2039324aa28bebf00c89e37579882f7cb261d19.
//
// Solidity: event Withdraw(address indexed _a, uint64 _pi, uint256 _m)
func (_Piece *PieceFilterer) FilterWithdraw(opts *bind.FilterOpts, _a []common.Address) (*PieceWithdrawIterator, error) {

	var _aRule []interface{}
	for _, _aItem := range _a {
		_aRule = append(_aRule, _aItem)
	}

	logs, sub, err := _Piece.contract.FilterLogs(opts, "Withdraw", _aRule)
	if err != nil {
		return nil, err
	}
	return &PieceWithdrawIterator{contract: _Piece.contract, event: "Withdraw", logs: logs, sub: sub}, nil
}

// WatchWithdraw is a free log subscription operation binding the contract event 0xb283270b87db7ad5d1fbb15af2039324aa28bebf00c89e37579882f7cb261d19.
//
// Solidity: event Withdraw(address indexed _a, uint64 _pi, uint256 _m)
func (_Piece *PieceFilterer) WatchWithdraw(opts *bind.WatchOpts, sink chan<- *PieceWithdraw, _a []common.Address) (event.Subscription, error) {

	var _aRule []interface{}
	for _, _aItem := range _a {
		_aRule = append(_aRule, _aItem)
	}

	logs, sub, err := _Piece.contract.WatchLogs(opts, "Withdraw", _aRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PieceWithdraw)
				if err := _Piece.contract.UnpackLog(event, "Withdraw", log); err != nil {
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

// ParseWithdraw is a log parse operation binding the contract event 0xb283270b87db7ad5d1fbb15af2039324aa28bebf00c89e37579882f7cb261d19.
//
// Solidity: event Withdraw(address indexed _a, uint64 _pi, uint256 _m)
func (_Piece *PieceFilterer) ParseWithdraw(log types.Log) (*PieceWithdraw, error) {
	event := new(PieceWithdraw)
	if err := _Piece.contract.UnpackLog(event, "Withdraw", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
