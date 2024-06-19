// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package rsproof

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

// RSProofProofInfo is an auto generated low-level Go binding around an user-defined struct.
type RSProofProofInfo struct {
	Fake     bool
	Chaler   common.Address
	ChalTime *big.Int
}

// BytesLibMetaData contains all meta data concerning the BytesLib contract.
var BytesLibMetaData = &bind.MetaData{
	ABI: "[]",
	Bin: "0x6055604b600b8282823980515f1a607314603f577f4e487b71000000000000000000000000000000000000000000000000000000005f525f60045260245ffd5b305f52607381538281f3fe730000000000000000000000000000000000000000301460806040525f80fdfea26469706673582212200940c2268ae67147c242f52c2d31a571619dbffb60f722683b22f128cea00ad064736f6c634300081a0033",
}

// BytesLibABI is the input ABI used to generate the binding from.
// Deprecated: Use BytesLibMetaData.ABI instead.
var BytesLibABI = BytesLibMetaData.ABI

// BytesLibBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use BytesLibMetaData.Bin instead.
var BytesLibBin = BytesLibMetaData.Bin

// DeployBytesLib deploys a new Ethereum contract, binding an instance of BytesLib to it.
func DeployBytesLib(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *BytesLib, error) {
	parsed, err := BytesLibMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(BytesLibBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &BytesLib{BytesLibCaller: BytesLibCaller{contract: contract}, BytesLibTransactor: BytesLibTransactor{contract: contract}, BytesLibFilterer: BytesLibFilterer{contract: contract}}, nil
}

// BytesLib is an auto generated Go binding around an Ethereum contract.
type BytesLib struct {
	BytesLibCaller     // Read-only binding to the contract
	BytesLibTransactor // Write-only binding to the contract
	BytesLibFilterer   // Log filterer for contract events
}

// BytesLibCaller is an auto generated read-only Go binding around an Ethereum contract.
type BytesLibCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BytesLibTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BytesLibTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BytesLibFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BytesLibFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BytesLibSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BytesLibSession struct {
	Contract     *BytesLib         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BytesLibCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BytesLibCallerSession struct {
	Contract *BytesLibCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// BytesLibTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BytesLibTransactorSession struct {
	Contract     *BytesLibTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// BytesLibRaw is an auto generated low-level Go binding around an Ethereum contract.
type BytesLibRaw struct {
	Contract *BytesLib // Generic contract binding to access the raw methods on
}

// BytesLibCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BytesLibCallerRaw struct {
	Contract *BytesLibCaller // Generic read-only contract binding to access the raw methods on
}

// BytesLibTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BytesLibTransactorRaw struct {
	Contract *BytesLibTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBytesLib creates a new instance of BytesLib, bound to a specific deployed contract.
func NewBytesLib(address common.Address, backend bind.ContractBackend) (*BytesLib, error) {
	contract, err := bindBytesLib(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &BytesLib{BytesLibCaller: BytesLibCaller{contract: contract}, BytesLibTransactor: BytesLibTransactor{contract: contract}, BytesLibFilterer: BytesLibFilterer{contract: contract}}, nil
}

// NewBytesLibCaller creates a new read-only instance of BytesLib, bound to a specific deployed contract.
func NewBytesLibCaller(address common.Address, caller bind.ContractCaller) (*BytesLibCaller, error) {
	contract, err := bindBytesLib(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BytesLibCaller{contract: contract}, nil
}

// NewBytesLibTransactor creates a new write-only instance of BytesLib, bound to a specific deployed contract.
func NewBytesLibTransactor(address common.Address, transactor bind.ContractTransactor) (*BytesLibTransactor, error) {
	contract, err := bindBytesLib(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BytesLibTransactor{contract: contract}, nil
}

// NewBytesLibFilterer creates a new log filterer instance of BytesLib, bound to a specific deployed contract.
func NewBytesLibFilterer(address common.Address, filterer bind.ContractFilterer) (*BytesLibFilterer, error) {
	contract, err := bindBytesLib(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BytesLibFilterer{contract: contract}, nil
}

// bindBytesLib binds a generic wrapper to an already deployed contract.
func bindBytesLib(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := BytesLibMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BytesLib *BytesLibRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BytesLib.Contract.BytesLibCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BytesLib *BytesLibRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BytesLib.Contract.BytesLibTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BytesLib *BytesLibRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BytesLib.Contract.BytesLibTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BytesLib *BytesLibCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BytesLib.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BytesLib *BytesLibTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BytesLib.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BytesLib *BytesLibTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BytesLib.Contract.contract.Transact(opts, method, params...)
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
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_a\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"_ep\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"_p\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"_d\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"_s\",\"type\":\"uint64\"}],\"name\":\"add\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_a\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_money\",\"type\":\"uint256\"}],\"name\":\"checkNode\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_a\",\"type\":\"address\"},{\"internalType\":\"uint8\",\"name\":\"_typ\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"_m\",\"type\":\"uint256\"}],\"name\":\"lock\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_a\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_m\",\"type\":\"uint256\"}],\"name\":\"mint\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_a\",\"type\":\"address\"},{\"internalType\":\"uint8\",\"name\":\"_typ\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_m\",\"type\":\"uint256\"}],\"name\":\"punish\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_a\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_m\",\"type\":\"uint256\"}],\"name\":\"release\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_a\",\"type\":\"address\"},{\"internalType\":\"uint8\",\"name\":\"_typ\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"_m\",\"type\":\"uint256\"}],\"name\":\"unlock\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
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

// Add is a paid mutator transaction binding the contract method 0x3acd28a9.
//
// Solidity: function add(address _a, uint64 _ep, uint256 _p, uint64 _d, uint64 _s) returns()
func (_IControl *IControlTransactor) Add(opts *bind.TransactOpts, _a common.Address, _ep uint64, _p *big.Int, _d uint64, _s uint64) (*types.Transaction, error) {
	return _IControl.contract.Transact(opts, "add", _a, _ep, _p, _d, _s)
}

// Add is a paid mutator transaction binding the contract method 0x3acd28a9.
//
// Solidity: function add(address _a, uint64 _ep, uint256 _p, uint64 _d, uint64 _s) returns()
func (_IControl *IControlSession) Add(_a common.Address, _ep uint64, _p *big.Int, _d uint64, _s uint64) (*types.Transaction, error) {
	return _IControl.Contract.Add(&_IControl.TransactOpts, _a, _ep, _p, _d, _s)
}

// Add is a paid mutator transaction binding the contract method 0x3acd28a9.
//
// Solidity: function add(address _a, uint64 _ep, uint256 _p, uint64 _d, uint64 _s) returns()
func (_IControl *IControlTransactorSession) Add(_a common.Address, _ep uint64, _p *big.Int, _d uint64, _s uint64) (*types.Transaction, error) {
	return _IControl.Contract.Add(&_IControl.TransactOpts, _a, _ep, _p, _d, _s)
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
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_a\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"_pi\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"_e\",\"type\":\"uint64\"}],\"name\":\"AddPiece\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_a\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"_ri\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"_sri\",\"type\":\"uint64\"}],\"name\":\"AddReplica\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_a\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_m\",\"type\":\"uint256\"}],\"name\":\"Settle\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_a\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"_pi\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_m\",\"type\":\"uint256\"}],\"name\":\"Withdraw\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"_pi\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"_ri\",\"type\":\"uint64\"},{\"internalType\":\"uint8\",\"name\":\"_pri\",\"type\":\"uint8\"}],\"name\":\"checkPRI\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_a\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"_ri\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"_com\",\"type\":\"bytes\"}],\"name\":\"checkSReplica\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_a\",\"type\":\"address\"}],\"name\":\"checkStore\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_pn\",\"type\":\"bytes\"}],\"name\":\"getPIndex\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"_pi\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"_ri\",\"type\":\"uint64\"}],\"name\":\"getPRInfo\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"},{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_rn\",\"type\":\"bytes\"}],\"name\":\"getRIndex\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_a\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"_e\",\"type\":\"uint64\"}],\"name\":\"getStoreCount\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_a\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"_e\",\"type\":\"uint64\"}],\"name\":\"getStoreSalary\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
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

// CheckStore is a paid mutator transaction binding the contract method 0x62b98032.
//
// Solidity: function checkStore(address _a) returns()
func (_IPiece *IPieceTransactor) CheckStore(opts *bind.TransactOpts, _a common.Address) (*types.Transaction, error) {
	return _IPiece.contract.Transact(opts, "checkStore", _a)
}

// CheckStore is a paid mutator transaction binding the contract method 0x62b98032.
//
// Solidity: function checkStore(address _a) returns()
func (_IPiece *IPieceSession) CheckStore(_a common.Address) (*types.Transaction, error) {
	return _IPiece.Contract.CheckStore(&_IPiece.TransactOpts, _a)
}

// CheckStore is a paid mutator transaction binding the contract method 0x62b98032.
//
// Solidity: function checkStore(address _a) returns()
func (_IPiece *IPieceTransactorSession) CheckStore(_a common.Address) (*types.Transaction, error) {
	return _IPiece.Contract.CheckStore(&_IPiece.TransactOpts, _a)
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
	M   *big.Int
	Raw types.Log // Blockchain specific contextual infos
}

// FilterSettle is a free log retrieval operation binding the contract event 0xa3788eedc10ef3ec6982384227c562c6666cf2b6af4f6a583b6d5d0f2ba0d204.
//
// Solidity: event Settle(address indexed _a, uint256 _m)
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

// WatchSettle is a free log subscription operation binding the contract event 0xa3788eedc10ef3ec6982384227c562c6666cf2b6af4f6a583b6d5d0f2ba0d204.
//
// Solidity: event Settle(address indexed _a, uint256 _m)
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

// ParseSettle is a log parse operation binding the contract event 0xa3788eedc10ef3ec6982384227c562c6666cf2b6af4f6a583b6d5d0f2ba0d204.
//
// Solidity: event Settle(address indexed _a, uint256 _m)
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

// IPlonkMetaData contains all meta data concerning the IPlonk contract.
var IPlonkMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_proof\",\"type\":\"bytes\"},{\"internalType\":\"uint256[]\",\"name\":\"_public_inputs\",\"type\":\"uint256[]\"}],\"name\":\"Verify\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// IPlonkABI is the input ABI used to generate the binding from.
// Deprecated: Use IPlonkMetaData.ABI instead.
var IPlonkABI = IPlonkMetaData.ABI

// IPlonk is an auto generated Go binding around an Ethereum contract.
type IPlonk struct {
	IPlonkCaller     // Read-only binding to the contract
	IPlonkTransactor // Write-only binding to the contract
	IPlonkFilterer   // Log filterer for contract events
}

// IPlonkCaller is an auto generated read-only Go binding around an Ethereum contract.
type IPlonkCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IPlonkTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IPlonkTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IPlonkFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IPlonkFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IPlonkSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IPlonkSession struct {
	Contract     *IPlonk           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IPlonkCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IPlonkCallerSession struct {
	Contract *IPlonkCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// IPlonkTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IPlonkTransactorSession struct {
	Contract     *IPlonkTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IPlonkRaw is an auto generated low-level Go binding around an Ethereum contract.
type IPlonkRaw struct {
	Contract *IPlonk // Generic contract binding to access the raw methods on
}

// IPlonkCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IPlonkCallerRaw struct {
	Contract *IPlonkCaller // Generic read-only contract binding to access the raw methods on
}

// IPlonkTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IPlonkTransactorRaw struct {
	Contract *IPlonkTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIPlonk creates a new instance of IPlonk, bound to a specific deployed contract.
func NewIPlonk(address common.Address, backend bind.ContractBackend) (*IPlonk, error) {
	contract, err := bindIPlonk(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IPlonk{IPlonkCaller: IPlonkCaller{contract: contract}, IPlonkTransactor: IPlonkTransactor{contract: contract}, IPlonkFilterer: IPlonkFilterer{contract: contract}}, nil
}

// NewIPlonkCaller creates a new read-only instance of IPlonk, bound to a specific deployed contract.
func NewIPlonkCaller(address common.Address, caller bind.ContractCaller) (*IPlonkCaller, error) {
	contract, err := bindIPlonk(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IPlonkCaller{contract: contract}, nil
}

// NewIPlonkTransactor creates a new write-only instance of IPlonk, bound to a specific deployed contract.
func NewIPlonkTransactor(address common.Address, transactor bind.ContractTransactor) (*IPlonkTransactor, error) {
	contract, err := bindIPlonk(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IPlonkTransactor{contract: contract}, nil
}

// NewIPlonkFilterer creates a new log filterer instance of IPlonk, bound to a specific deployed contract.
func NewIPlonkFilterer(address common.Address, filterer bind.ContractFilterer) (*IPlonkFilterer, error) {
	contract, err := bindIPlonk(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IPlonkFilterer{contract: contract}, nil
}

// bindIPlonk binds a generic wrapper to an already deployed contract.
func bindIPlonk(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := IPlonkMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IPlonk *IPlonkRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IPlonk.Contract.IPlonkCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IPlonk *IPlonkRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IPlonk.Contract.IPlonkTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IPlonk *IPlonkRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IPlonk.Contract.IPlonkTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IPlonk *IPlonkCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IPlonk.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IPlonk *IPlonkTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IPlonk.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IPlonk *IPlonkTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IPlonk.Contract.contract.Transact(opts, method, params...)
}

// Verify is a paid mutator transaction binding the contract method 0x7e4f7a8a.
//
// Solidity: function Verify(bytes _proof, uint256[] _public_inputs) returns(bool)
func (_IPlonk *IPlonkTransactor) Verify(opts *bind.TransactOpts, _proof []byte, _public_inputs []*big.Int) (*types.Transaction, error) {
	return _IPlonk.contract.Transact(opts, "Verify", _proof, _public_inputs)
}

// Verify is a paid mutator transaction binding the contract method 0x7e4f7a8a.
//
// Solidity: function Verify(bytes _proof, uint256[] _public_inputs) returns(bool)
func (_IPlonk *IPlonkSession) Verify(_proof []byte, _public_inputs []*big.Int) (*types.Transaction, error) {
	return _IPlonk.Contract.Verify(&_IPlonk.TransactOpts, _proof, _public_inputs)
}

// Verify is a paid mutator transaction binding the contract method 0x7e4f7a8a.
//
// Solidity: function Verify(bytes _proof, uint256[] _public_inputs) returns(bool)
func (_IPlonk *IPlonkTransactorSession) Verify(_proof []byte, _public_inputs []*big.Int) (*types.Transaction, error) {
	return _IPlonk.Contract.Verify(&_IPlonk.TransactOpts, _proof, _public_inputs)
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

// RSProofMetaData contains all meta data concerning the RSProof contract.
var RSProofMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_b\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_s\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"_ri\",\"type\":\"uint64\"}],\"name\":\"Challenge\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_s\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"_ri\",\"type\":\"uint64\"}],\"name\":\"Fake\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_s\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"_ri\",\"type\":\"uint64\"}],\"name\":\"Prove\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"KZG_MAX\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"KZG_VK\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"bank\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"basePenalty\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_pn\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"_rn\",\"type\":\"bytes\"},{\"internalType\":\"uint8\",\"name\":\"_pri\",\"type\":\"uint8\"}],\"name\":\"challenge\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"_pi\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"_ri\",\"type\":\"uint64\"},{\"internalType\":\"uint8\",\"name\":\"_pri\",\"type\":\"uint8\"}],\"name\":\"check\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"current\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"_pi\",\"type\":\"uint64\"},{\"internalType\":\"uint8\",\"name\":\"_pri\",\"type\":\"uint8\"}],\"name\":\"getProof\",\"outputs\":[{\"components\":[{\"internalType\":\"bool\",\"name\":\"fake\",\"type\":\"bool\"},{\"internalType\":\"address\",\"name\":\"chaler\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"chalTime\",\"type\":\"uint256\"}],\"internalType\":\"structRSProof.ProofInfo\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"_rsn\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"_rsk\",\"type\":\"uint8\"}],\"name\":\"getVKRoot\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"minProveTime\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_pn\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"_rn\",\"type\":\"bytes\"},{\"internalType\":\"uint8\",\"name\":\"_pri\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"_proof\",\"type\":\"bytes\"}],\"name\":\"prove\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"_rsn\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"_rsk\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"_vkroot\",\"type\":\"uint256\"}],\"name\":\"setVeirfier\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x6080604052604051806060016040528060308152602001613b23603091396002908161002b91906103d7565b50630200000060035561012c60055567016345785d8a0000600655348015610051575f80fd5b50604051613b53380380613b5383398181016040528101906100739190610504565b61008f6100846100d560201b60201c565b6100dc60201b60201c565b8060015f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055505061052f565b5f33905090565b5f805f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050815f806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508173ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a35050565b5f81519050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b7f4e487b71000000000000000000000000000000000000000000000000000000005f52602260045260245ffd5b5f600282049050600182168061021857607f821691505b60208210810361022b5761022a6101d4565b5b50919050565b5f819050815f5260205f209050919050565b5f6020601f8301049050919050565b5f82821b905092915050565b5f6008830261028d7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82610252565b6102978683610252565b95508019841693508086168417925050509392505050565b5f819050919050565b5f819050919050565b5f6102db6102d66102d1846102af565b6102b8565b6102af565b9050919050565b5f819050919050565b6102f4836102c1565b610308610300826102e2565b84845461025e565b825550505050565b5f90565b61031c610310565b6103278184846102eb565b505050565b5b8181101561034a5761033f5f82610314565b60018101905061032d565b5050565b601f82111561038f5761036081610231565b61036984610243565b81016020851015610378578190505b61038c61038485610243565b83018261032c565b50505b505050565b5f82821c905092915050565b5f6103af5f1984600802610394565b1980831691505092915050565b5f6103c783836103a0565b9150826002028217905092915050565b6103e08261019d565b67ffffffffffffffff8111156103f9576103f86101a7565b5b6104038254610201565b61040e82828561034e565b5f60209050601f83116001811461043f575f841561042d578287015190505b61043785826103bc565b86555061049e565b601f19841661044d86610231565b5f5b828110156104745784890151825560018201915060208501945060208101905061044f565b86831015610491578489015161048d601f8916826103a0565b8355505b6001600288020188555050505b505050505050565b5f80fd5b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f6104d3826104aa565b9050919050565b6104e3816104c9565b81146104ed575f80fd5b50565b5f815190506104fe816104da565b92915050565b5f60208284031215610519576105186104a6565b5b5f610526848285016104f0565b91505092915050565b6135e78061053c5f395ff3fe608060405234801561000f575f80fd5b50600436106100f3575f3560e01c806376cdb03b11610095578063abe881d611610064578063abe881d614610239578063b29cc00914610257578063f2fde38b14610287578063f354cd5f146102a3576100f3565b806376cdb03b146101af5780638da5cb5b146101cd5780639fa6a6e3146101eb578063a617efec14610209576100f3565b80635712e98c116100d15780635712e98c1461014f5780635acbd1bf1461016d5780636c05dff814610189578063715018a6146101a5576100f3565b80630bd26cb5146100f757806323178d22146101155780633c530b0814610133575b5f80fd5b6100ff6102bf565b60405161010c91906122f3565b60405180910390f35b61011d6102c5565b60405161012a9190612396565b60405180910390f35b61014d60048036038101906101489190612529565b610351565b005b610157610be6565b60405161016491906122f3565b60405180910390f35b6101876004803603810190610182919061260b565b610bec565b005b6101a3600480360381019061019e919061265b565b610cbd565b005b6101ad611162565b005b6101b7611175565b6040516101c49190612722565b60405180910390f35b6101d561119a565b6040516101e29190612722565b60405180910390f35b6101f36111c1565b604051610200919061275d565b60405180910390f35b610223600480360381019061021e91906127a0565b6111da565b6040516102309190612856565b60405180910390f35b610241611326565b60405161024e91906122f3565b60405180910390f35b610271600480360381019061026c919061286f565b61132c565b60405161027e91906122f3565b60405180910390f35b6102a1600480360381019061029c91906128d7565b611362565b005b6102bd60048036038101906102b89190612902565b6113e4565b005b60055481565b600280546102d29061297f565b80601f01602080910402602001604051908101604052809291908181526020018280546102fe9061297f565b80156103495780601f1061032057610100808354040283529160200191610349565b820191905f5260205f20905b81548152906001019060200180831161032c57829003601f168201915b505050505081565b5f60015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663693ec85e6040518163ffffffff1660e01b81526004016103aa90612a09565b6020604051808303815f875af11580156103c6573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906103ea9190612a3b565b90505f8173ffffffffffffffffffffffffffffffffffffffff16631ce85e7c876040518263ffffffff1660e01b81526004016104269190612396565b602060405180830381865afa158015610441573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906104659190612a7a565b90505f8273ffffffffffffffffffffffffffffffffffffffff1663ba8b2618876040518263ffffffff1660e01b81526004016104a19190612396565b602060405180830381865afa1580156104bc573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906104e09190612a7a565b90505f8373ffffffffffffffffffffffffffffffffffffffff16632700f3a88484896040518463ffffffff1660e01b815260040161052093929190612ab4565b6020604051808303815f875af115801561053c573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906105609190612a3b565b90505f805f808773ffffffffffffffffffffffffffffffffffffffff1663b650b50488886040518363ffffffff1660e01b81526004016105a1929190612ae9565b608060405180830381865afa1580156105bc573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906105e09190612b57565b935093509350935060075f8867ffffffffffffffff1667ffffffffffffffff1681526020019081526020015f205f8b60ff1660ff1681526020019081526020015f205f015f9054906101000a900460ff1615610671576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161066890612c05565b60405180910390fd5b5f60075f8967ffffffffffffffff1667ffffffffffffffff1681526020019081526020015f205f8c60ff1660ff1681526020019081526020015f2060010154116106f0576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016106e790612c6d565b60405180910390fd5b60055460075f8967ffffffffffffffff1667ffffffffffffffff1681526020019081526020015f205f8c60ff1660ff1681526020019081526020015f206001015461073b9190612cb8565b431061077c576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161077390612d35565b60405180910390fd5b610785856117d2565b5f60075f8967ffffffffffffffff1667ffffffffffffffff1681526020019081526020015f205f8c60ff1660ff1681526020019081526020015f20600101819055505f815f1c90507f30644e72e131a029b85045b68181585d2833e84879b9709143e1f593f0000001810690505f8d8d6108018e60ff16611aa3565b6108148767ffffffffffffffff16611aa3565b61081f600354611aa3565b600260405160200161083696959493929190612e1f565b604051602081830303815290604052805190602001205f1c90507f30644e72e131a029b85045b68181585d2833e84879b9709143e1f593f0000001810690505f60085f8860ff1660ff1681526020019081526020015f205f8760ff1660ff1681526020019081526020015f2054036108e3576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016108da90612ec0565b60405180910390fd5b5f600367ffffffffffffffff8111156108ff576108fe6123cf565b5b60405190808252806020026020018201604052801561092d5781602001602082028036833780820191505090505b50905060085f8860ff1660ff1681526020019081526020015f205f8760ff1660ff1681526020019081526020015f2054815f815181106109705761096f612ede565b5b602002602001018181525050818160018151811061099157610990612ede565b5b60200260200101818152505082816002815181106109b2576109b1612ede565b5b6020026020010181815250505f60015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663693ec85e6040518163ffffffff1660e01b8152600401610a1790612f55565b6020604051808303815f875af1158015610a33573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610a579190612a3b565b90505f73ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1603610ac7576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610abe90612fbd565b60405180910390fd5b5f8173ffffffffffffffffffffffffffffffffffffffff16637e4f7a8a8f856040518363ffffffff1660e01b8152600401610b03929190613083565b6020604051808303815f875af1158015610b1f573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610b4391906130e2565b905080610b85576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610b7c90613157565b60405180910390fd5b8973ffffffffffffffffffffffffffffffffffffffff167f4387d8e339e908dfa927fdd9f9555518616803c51d321dce6e75d1a647bd16598c604051610bcb919061275d565b60405180910390a25050505050505050505050505050505050565b60065481565b610bf4611b79565b5f60085f8560ff1660ff1681526020019081526020015f205f8460ff1660ff1681526020019081526020015f205414610c62576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610c59906131bf565b60405180910390fd5b7f30644e72e131a029b85045b68181585d2833e84879b9709143e1f593f0000001810690508060085f8560ff1660ff1681526020019081526020015f205f8460ff1660ff1681526020019081526020015f2081905550505050565b5f60015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663693ec85e6040518163ffffffff1660e01b8152600401610d1690612a09565b6020604051808303815f875af1158015610d32573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610d569190612a3b565b90505f8173ffffffffffffffffffffffffffffffffffffffff16631ce85e7c866040518263ffffffff1660e01b8152600401610d929190612396565b602060405180830381865afa158015610dad573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610dd19190612a7a565b90505f8273ffffffffffffffffffffffffffffffffffffffff1663ba8b2618866040518263ffffffff1660e01b8152600401610e0d9190612396565b602060405180830381865afa158015610e28573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610e4c9190612a7a565b90505f8373ffffffffffffffffffffffffffffffffffffffff16632700f3a88484886040518463ffffffff1660e01b8152600401610e8c93929190612ab4565b6020604051808303815f875af1158015610ea8573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610ecc9190612a3b565b90505f73ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1603610f3c576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610f3390613227565b60405180910390fd5b60075f8467ffffffffffffffff1667ffffffffffffffff1681526020019081526020015f205f8660ff1660ff1681526020019081526020015f205f015f9054906101000a900460ff1615610fc5576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610fbc90612c05565b60405180910390fd5b5f60075f8567ffffffffffffffff1667ffffffffffffffff1681526020019081526020015f205f8760ff1660ff1681526020019081526020015f206001015414611044576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161103b9061328f565b60405180910390fd5b61104e8133611bf7565b4360075f8567ffffffffffffffff1667ffffffffffffffff1681526020019081526020015f205f8760ff1660ff1681526020019081526020015f20600101819055503360075f8567ffffffffffffffff1667ffffffffffffffff1681526020019081526020015f205f8760ff1660ff1681526020019081526020015f205f0160016101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508073ffffffffffffffffffffffffffffffffffffffff167f40871ff904cf6ca42f9dcbac1f7d50ab55381a1d0db3eedd5dd967b209f8d82383604051611151919061275d565b60405180910390a250505050505050565b61116a611b79565b6111735f611d8a565b565b60015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b5f805f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b60045f9054906101000a900467ffffffffffffffff1681565b6111e26122a5565b6111ea6122a5565b60075f8567ffffffffffffffff1667ffffffffffffffff1681526020019081526020015f205f8460ff1660ff1681526020019081526020015f205f015f9054906101000a900460ff16815f01901515908115158152505060075f8567ffffffffffffffff1667ffffffffffffffff1681526020019081526020015f205f8460ff1660ff1681526020019081526020015f205f0160019054906101000a900473ffffffffffffffffffffffffffffffffffffffff16816020019073ffffffffffffffffffffffffffffffffffffffff16908173ffffffffffffffffffffffffffffffffffffffff168152505060075f8567ffffffffffffffff1667ffffffffffffffff1681526020019081526020015f205f8460ff1660ff1681526020019081526020015f20600101548160400181815250508091505092915050565b60035481565b5f60085f8460ff1660ff1681526020019081526020015f205f8360ff1660ff1681526020019081526020015f2054905092915050565b61136a611b79565b5f73ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16036113d8576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016113cf9061331d565b60405180910390fd5b6113e181611d8a565b50565b60075f8467ffffffffffffffff1667ffffffffffffffff1681526020019081526020015f205f8260ff1660ff1681526020019081526020015f205f015f9054906101000a900460ff161561146d576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161146490612c05565b60405180910390fd5b5f60075f8567ffffffffffffffff1667ffffffffffffffff1681526020019081526020015f205f8360ff1660ff1681526020019081526020015f2060010154116114ec576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016114e390612c6d565b60405180910390fd5b5f60015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663693ec85e6040518163ffffffff1660e01b815260040161154590612a09565b6020604051808303815f875af1158015611561573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906115859190612a3b565b73ffffffffffffffffffffffffffffffffffffffff16632700f3a88585856040518463ffffffff1660e01b81526004016115c193929190612ab4565b6020604051808303815f875af11580156115dd573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906116019190612a3b565b90505f73ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1603611671576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161166890613227565b60405180910390fd5b60055460075f8667ffffffffffffffff1667ffffffffffffffff1681526020019081526020015f205f8460ff1660ff1681526020019081526020015f20600101546116bc9190612cb8565b4311156117cc576117298160075f8767ffffffffffffffff1667ffffffffffffffff1681526020019081526020015f205f8560ff1660ff1681526020019081526020015f205f0160019054906101000a900473ffffffffffffffffffffffffffffffffffffffff16611e4b565b600160075f8667ffffffffffffffff1667ffffffffffffffff1681526020019081526020015f205f8460ff1660ff1681526020019081526020015f205f015f6101000a81548160ff0219169083151502179055508073ffffffffffffffffffffffffffffffffffffffff167f0fcd3bb0ce91639d74185d625544f2b5714ecd14a6a6d9ccb65b2526e9ac12f5846040516117c3919061275d565b60405180910390a25b50505050565b5f60015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663693ec85e6040518163ffffffff1660e01b815260040161182b90613385565b6020604051808303815f875af1158015611847573d5f803e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061186b9190612a3b565b905060015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663693ec85e6040518163ffffffff1660e01b81526004016118c5906133ed565b6020604051808303815f875af11580156118e1573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906119059190612a3b565b73ffffffffffffffffffffffffffffffffffffffff16632165e20b8360016006546040518463ffffffff1660e01b81526004016119449392919061344d565b5f604051808303815f87803b15801561195b575f80fd5b505af115801561196d573d5f803e3d5ffd5b5050505060015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166376890c588360026006546119be91906134af565b6040518363ffffffff1660e01b81526004016119db9291906134df565b5f604051808303815f87803b1580156119f2575f80fd5b505af1158015611a04573d5f803e3d5ffd5b5050505060015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166376890c58826002600654611a5591906134af565b6040518363ffffffff1660e01b8152600401611a729291906134df565b5f604051808303815f87803b158015611a89575f80fd5b505af1158015611a9b573d5f803e3d5ffd5b505050505050565b60605f67ffffffffffffffff90505f60c0828516901b9050604082901b91505f6040838616901b90508082611ad89190612cb8565b9150604083901b92506040838616901c90508082611af69190612cb8565b9150604083901b925060c0838616901c90508082611b149190612cb8565b91505f603067ffffffffffffffff811115611b3257611b316123cf565b5b6040519080825280601f01601f191660200182016040528015611b645781602001600182028036833780820191505090505b50905082602082015280945050505050919050565b611b8161229e565b73ffffffffffffffffffffffffffffffffffffffff16611b9f61119a565b73ffffffffffffffffffffffffffffffffffffffff1614611bf5576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401611bec90613550565b60405180910390fd5b565b60015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663693ec85e6040518163ffffffff1660e01b8152600401611c4f906133ed565b6020604051808303815f875af1158015611c6b573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190611c8f9190612a3b565b73ffffffffffffffffffffffffffffffffffffffff1663738dddba8360016006546040518463ffffffff1660e01b8152600401611cce9392919061344d565b5f604051808303815f87803b158015611ce5575f80fd5b505af1158015611cf7573d5f803e3d5ffd5b5050505060015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663e8888915826006546040518363ffffffff1660e01b8152600401611d599291906134df565b5f604051808303815f87803b158015611d70575f80fd5b505af1158015611d82573d5f803e3d5ffd5b505050505050565b5f805f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050815f806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508173ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a35050565b5f60015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663693ec85e6040518163ffffffff1660e01b8152600401611ea490613385565b6020604051808303815f875af1158015611ec0573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190611ee49190612a3b565b905060015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663693ec85e6040518163ffffffff1660e01b8152600401611f3e906133ed565b6020604051808303815f875af1158015611f5a573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190611f7e9190612a3b565b73ffffffffffffffffffffffffffffffffffffffff16632165e20b8460016006546040518463ffffffff1660e01b8152600401611fbd9392919061344d565b5f604051808303815f87803b158015611fd4575f80fd5b505af1158015611fe6573d5f803e3d5ffd5b5050505060015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663693ec85e6040518163ffffffff1660e01b8152600401612042906133ed565b6020604051808303815f875af115801561205e573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906120829190612a3b565b73ffffffffffffffffffffffffffffffffffffffff166306f0b4f18460018560026006546120b091906134af565b6040518563ffffffff1660e01b81526004016120cf949392919061356e565b5f604051808303815f87803b1580156120e6575f80fd5b505af11580156120f8573d5f803e3d5ffd5b5050505060015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663693ec85e6040518163ffffffff1660e01b8152600401612154906133ed565b6020604051808303815f875af1158015612170573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906121949190612a3b565b73ffffffffffffffffffffffffffffffffffffffff166306f0b4f18460018460026006546121c291906134af565b6040518563ffffffff1660e01b81526004016121e1949392919061356e565b5f604051808303815f87803b1580156121f8575f80fd5b505af115801561220a573d5f803e3d5ffd5b5050505060015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166376890c58836006546040518363ffffffff1660e01b815260040161226c9291906134df565b5f604051808303815f87803b158015612283575f80fd5b505af1158015612295573d5f803e3d5ffd5b50505050505050565b5f33905090565b60405180606001604052805f151581526020015f73ffffffffffffffffffffffffffffffffffffffff1681526020015f81525090565b5f819050919050565b6122ed816122db565b82525050565b5f6020820190506123065f8301846122e4565b92915050565b5f81519050919050565b5f82825260208201905092915050565b5f5b83811015612343578082015181840152602081019050612328565b5f8484015250505050565b5f601f19601f8301169050919050565b5f6123688261230c565b6123728185612316565b9350612382818560208601612326565b61238b8161234e565b840191505092915050565b5f6020820190508181035f8301526123ae818461235e565b905092915050565b5f604051905090565b5f80fd5b5f80fd5b5f80fd5b5f80fd5b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b6124058261234e565b810181811067ffffffffffffffff82111715612424576124236123cf565b5b80604052505050565b5f6124366123b6565b905061244282826123fc565b919050565b5f67ffffffffffffffff821115612461576124606123cf565b5b61246a8261234e565b9050602081019050919050565b828183375f83830152505050565b5f61249761249284612447565b61242d565b9050828152602081018484840111156124b3576124b26123cb565b5b6124be848285612477565b509392505050565b5f82601f8301126124da576124d96123c7565b5b81356124ea848260208601612485565b91505092915050565b5f60ff82169050919050565b612508816124f3565b8114612512575f80fd5b50565b5f81359050612523816124ff565b92915050565b5f805f8060808587031215612541576125406123bf565b5b5f85013567ffffffffffffffff81111561255e5761255d6123c3565b5b61256a878288016124c6565b945050602085013567ffffffffffffffff81111561258b5761258a6123c3565b5b612597878288016124c6565b93505060406125a887828801612515565b925050606085013567ffffffffffffffff8111156125c9576125c86123c3565b5b6125d5878288016124c6565b91505092959194509250565b6125ea816122db565b81146125f4575f80fd5b50565b5f81359050612605816125e1565b92915050565b5f805f60608486031215612622576126216123bf565b5b5f61262f86828701612515565b935050602061264086828701612515565b9250506040612651868287016125f7565b9150509250925092565b5f805f60608486031215612672576126716123bf565b5b5f84013567ffffffffffffffff81111561268f5761268e6123c3565b5b61269b868287016124c6565b935050602084013567ffffffffffffffff8111156126bc576126bb6123c3565b5b6126c8868287016124c6565b92505060406126d986828701612515565b9150509250925092565b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f61270c826126e3565b9050919050565b61271c81612702565b82525050565b5f6020820190506127355f830184612713565b92915050565b5f67ffffffffffffffff82169050919050565b6127578161273b565b82525050565b5f6020820190506127705f83018461274e565b92915050565b61277f8161273b565b8114612789575f80fd5b50565b5f8135905061279a81612776565b92915050565b5f80604083850312156127b6576127b56123bf565b5b5f6127c38582860161278c565b92505060206127d485828601612515565b9150509250929050565b5f8115159050919050565b6127f2816127de565b82525050565b61280181612702565b82525050565b612810816122db565b82525050565b606082015f82015161282a5f8501826127e9565b50602082015161283d60208501826127f8565b5060408201516128506040850182612807565b50505050565b5f6060820190506128695f830184612816565b92915050565b5f8060408385031215612885576128846123bf565b5b5f61289285828601612515565b92505060206128a385828601612515565b9150509250929050565b6128b681612702565b81146128c0575f80fd5b50565b5f813590506128d1816128ad565b92915050565b5f602082840312156128ec576128eb6123bf565b5b5f6128f9848285016128c3565b91505092915050565b5f805f60608486031215612919576129186123bf565b5b5f6129268682870161278c565b93505060206129378682870161278c565b925050604061294886828701612515565b9150509250925092565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52602260045260245ffd5b5f600282049050600182168061299657607f821691505b6020821081036129a9576129a8612952565b5b50919050565b5f82825260208201905092915050565b7f70696563650000000000000000000000000000000000000000000000000000005f82015250565b5f6129f36005836129af565b91506129fe826129bf565b602082019050919050565b5f6020820190508181035f830152612a20816129e7565b9050919050565b5f81519050612a35816128ad565b92915050565b5f60208284031215612a5057612a4f6123bf565b5b5f612a5d84828501612a27565b91505092915050565b5f81519050612a7481612776565b92915050565b5f60208284031215612a8f57612a8e6123bf565b5b5f612a9c84828501612a66565b91505092915050565b612aae816124f3565b82525050565b5f606082019050612ac75f83018661274e565b612ad4602083018561274e565b612ae16040830184612aa5565b949350505050565b5f604082019050612afc5f83018561274e565b612b09602083018461274e565b9392505050565b5f81519050612b1e816124ff565b92915050565b5f819050919050565b612b3681612b24565b8114612b40575f80fd5b50565b5f81519050612b5181612b2d565b92915050565b5f805f8060808587031215612b6f57612b6e6123bf565b5b5f612b7c87828801612b10565b9450506020612b8d87828801612b10565b9350506040612b9e87828801612a66565b9250506060612baf87828801612b43565b91505092959194509250565b7f6661696c656400000000000000000000000000000000000000000000000000005f82015250565b5f612bef6006836129af565b9150612bfa82612bbb565b602082019050919050565b5f6020820190508181035f830152612c1c81612be3565b9050919050565b7f6e6f206368616c000000000000000000000000000000000000000000000000005f82015250565b5f612c576007836129af565b9150612c6282612c23565b602082019050919050565b5f6020820190508181035f830152612c8481612c4b565b9050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b5f612cc2826122db565b9150612ccd836122db565b9250828201905080821115612ce557612ce4612c8b565b5b92915050565b7f65786365656420707400000000000000000000000000000000000000000000005f82015250565b5f612d1f6009836129af565b9150612d2a82612ceb565b602082019050919050565b5f6020820190508181035f830152612d4c81612d13565b9050919050565b5f81905092915050565b5f612d678261230c565b612d718185612d53565b9350612d81818560208601612326565b80840191505092915050565b5f819050815f5260205f209050919050565b5f8154612dab8161297f565b612db58186612d53565b9450600182165f8114612dcf5760018114612de457612e16565b60ff1983168652811515820286019350612e16565b612ded85612d8d565b5f5b83811015612e0e57815481890152600182019150602081019050612def565b838801955050505b50505092915050565b5f612e2a8289612d5d565b9150612e368288612d5d565b9150612e428287612d5d565b9150612e4e8286612d5d565b9150612e5a8285612d5d565b9150612e668284612d9f565b9150819050979650505050505050565b7f6e6f20766b726f6f7400000000000000000000000000000000000000000000005f82015250565b5f612eaa6009836129af565b9150612eb582612e76565b602082019050919050565b5f6020820190508181035f830152612ed781612e9e565b9050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52603260045260245ffd5b7f72736f6e650000000000000000000000000000000000000000000000000000005f82015250565b5f612f3f6005836129af565b9150612f4a82612f0b565b602082019050919050565b5f6020820190508181035f830152612f6c81612f33565b9050919050565b7f6e6f2070760000000000000000000000000000000000000000000000000000005f82015250565b5f612fa76005836129af565b9150612fb282612f73565b602082019050919050565b5f6020820190508181035f830152612fd481612f9b565b9050919050565b5f81519050919050565b5f82825260208201905092915050565b5f819050602082019050919050565b5f61300f8383612807565b60208301905092915050565b5f602082019050919050565b5f61303182612fdb565b61303b8185612fe5565b935061304683612ff5565b805f5b8381101561307657815161305d8882613004565b97506130688361301b565b925050600181019050613049565b5085935050505092915050565b5f6040820190508181035f83015261309b818561235e565b905081810360208301526130af8184613027565b90509392505050565b6130c1816127de565b81146130cb575f80fd5b50565b5f815190506130dc816130b8565b92915050565b5f602082840312156130f7576130f66123bf565b5b5f613104848285016130ce565b91505092915050565b7f696e7620706600000000000000000000000000000000000000000000000000005f82015250565b5f6131416006836129af565b915061314c8261310d565b602082019050919050565b5f6020820190508181035f83015261316e81613135565b9050919050565b7f68617320766b726f6f74000000000000000000000000000000000000000000005f82015250565b5f6131a9600a836129af565b91506131b482613175565b602082019050919050565b5f6020820190508181035f8301526131d68161319d565b9050919050565b7f6e6f2073746f72650000000000000000000000000000000000000000000000005f82015250565b5f6132116008836129af565b915061321c826131dd565b602082019050919050565b5f6020820190508181035f83015261323e81613205565b9050919050565b7f696e206368616c000000000000000000000000000000000000000000000000005f82015250565b5f6132796007836129af565b915061328482613245565b602082019050919050565b5f6020820190508181035f8301526132a68161326d565b9050919050565b7f4f776e61626c653a206e6577206f776e657220697320746865207a65726f20615f8201527f6464726573730000000000000000000000000000000000000000000000000000602082015250565b5f6133076026836129af565b9150613312826132ad565b604082019050919050565b5f6020820190508181035f830152613334816132fb565b9050919050565b7f62617365000000000000000000000000000000000000000000000000000000005f82015250565b5f61336f6004836129af565b915061337a8261333b565b602082019050919050565b5f6020820190508181035f83015261339c81613363565b9050919050565b7f636f6e74726f6c000000000000000000000000000000000000000000000000005f82015250565b5f6133d76007836129af565b91506133e2826133a3565b602082019050919050565b5f6020820190508181035f830152613404816133cb565b9050919050565b5f819050919050565b5f819050919050565b5f61343761343261342d8461340b565b613414565b6124f3565b9050919050565b6134478161341d565b82525050565b5f6060820190506134605f830186612713565b61346d602083018561343e565b61347a60408301846122e4565b949350505050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601260045260245ffd5b5f6134b9826122db565b91506134c4836122db565b9250826134d4576134d3613482565b5b828204905092915050565b5f6040820190506134f25f830185612713565b6134ff60208301846122e4565b9392505050565b7f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e65725f82015250565b5f61353a6020836129af565b915061354582613506565b602082019050919050565b5f6020820190508181035f8301526135678161352e565b9050919050565b5f6080820190506135815f830187612713565b61358e602083018661343e565b61359b6040830185612713565b6135a860608301846122e4565b9594505050505056fea264697066735822122076b03931c6babfb82be855759fbd11385d32d7c8f68a05b0ffec79f9b909ea0864736f6c634300081a00333b8201b322c65a735690cf82c850b8624c29ec05400e06ba92a9aad12c37c1605812abbc9a1a11f500b3ab28b7751b52",
}

// RSProofABI is the input ABI used to generate the binding from.
// Deprecated: Use RSProofMetaData.ABI instead.
var RSProofABI = RSProofMetaData.ABI

// RSProofBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use RSProofMetaData.Bin instead.
var RSProofBin = RSProofMetaData.Bin

// DeployRSProof deploys a new Ethereum contract, binding an instance of RSProof to it.
func DeployRSProof(auth *bind.TransactOpts, backend bind.ContractBackend, _b common.Address) (common.Address, *types.Transaction, *RSProof, error) {
	parsed, err := RSProofMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(RSProofBin), backend, _b)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &RSProof{RSProofCaller: RSProofCaller{contract: contract}, RSProofTransactor: RSProofTransactor{contract: contract}, RSProofFilterer: RSProofFilterer{contract: contract}}, nil
}

// RSProof is an auto generated Go binding around an Ethereum contract.
type RSProof struct {
	RSProofCaller     // Read-only binding to the contract
	RSProofTransactor // Write-only binding to the contract
	RSProofFilterer   // Log filterer for contract events
}

// RSProofCaller is an auto generated read-only Go binding around an Ethereum contract.
type RSProofCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RSProofTransactor is an auto generated write-only Go binding around an Ethereum contract.
type RSProofTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RSProofFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type RSProofFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RSProofSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type RSProofSession struct {
	Contract     *RSProof          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// RSProofCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type RSProofCallerSession struct {
	Contract *RSProofCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// RSProofTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type RSProofTransactorSession struct {
	Contract     *RSProofTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// RSProofRaw is an auto generated low-level Go binding around an Ethereum contract.
type RSProofRaw struct {
	Contract *RSProof // Generic contract binding to access the raw methods on
}

// RSProofCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type RSProofCallerRaw struct {
	Contract *RSProofCaller // Generic read-only contract binding to access the raw methods on
}

// RSProofTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type RSProofTransactorRaw struct {
	Contract *RSProofTransactor // Generic write-only contract binding to access the raw methods on
}

// NewRSProof creates a new instance of RSProof, bound to a specific deployed contract.
func NewRSProof(address common.Address, backend bind.ContractBackend) (*RSProof, error) {
	contract, err := bindRSProof(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &RSProof{RSProofCaller: RSProofCaller{contract: contract}, RSProofTransactor: RSProofTransactor{contract: contract}, RSProofFilterer: RSProofFilterer{contract: contract}}, nil
}

// NewRSProofCaller creates a new read-only instance of RSProof, bound to a specific deployed contract.
func NewRSProofCaller(address common.Address, caller bind.ContractCaller) (*RSProofCaller, error) {
	contract, err := bindRSProof(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &RSProofCaller{contract: contract}, nil
}

// NewRSProofTransactor creates a new write-only instance of RSProof, bound to a specific deployed contract.
func NewRSProofTransactor(address common.Address, transactor bind.ContractTransactor) (*RSProofTransactor, error) {
	contract, err := bindRSProof(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &RSProofTransactor{contract: contract}, nil
}

// NewRSProofFilterer creates a new log filterer instance of RSProof, bound to a specific deployed contract.
func NewRSProofFilterer(address common.Address, filterer bind.ContractFilterer) (*RSProofFilterer, error) {
	contract, err := bindRSProof(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &RSProofFilterer{contract: contract}, nil
}

// bindRSProof binds a generic wrapper to an already deployed contract.
func bindRSProof(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := RSProofMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_RSProof *RSProofRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _RSProof.Contract.RSProofCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_RSProof *RSProofRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RSProof.Contract.RSProofTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_RSProof *RSProofRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RSProof.Contract.RSProofTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_RSProof *RSProofCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _RSProof.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_RSProof *RSProofTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RSProof.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_RSProof *RSProofTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RSProof.Contract.contract.Transact(opts, method, params...)
}

// KZGMAX is a free data retrieval call binding the contract method 0xabe881d6.
//
// Solidity: function KZG_MAX() view returns(uint256)
func (_RSProof *RSProofCaller) KZGMAX(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _RSProof.contract.Call(opts, &out, "KZG_MAX")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// KZGMAX is a free data retrieval call binding the contract method 0xabe881d6.
//
// Solidity: function KZG_MAX() view returns(uint256)
func (_RSProof *RSProofSession) KZGMAX() (*big.Int, error) {
	return _RSProof.Contract.KZGMAX(&_RSProof.CallOpts)
}

// KZGMAX is a free data retrieval call binding the contract method 0xabe881d6.
//
// Solidity: function KZG_MAX() view returns(uint256)
func (_RSProof *RSProofCallerSession) KZGMAX() (*big.Int, error) {
	return _RSProof.Contract.KZGMAX(&_RSProof.CallOpts)
}

// KZGVK is a free data retrieval call binding the contract method 0x23178d22.
//
// Solidity: function KZG_VK() view returns(bytes)
func (_RSProof *RSProofCaller) KZGVK(opts *bind.CallOpts) ([]byte, error) {
	var out []interface{}
	err := _RSProof.contract.Call(opts, &out, "KZG_VK")

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// KZGVK is a free data retrieval call binding the contract method 0x23178d22.
//
// Solidity: function KZG_VK() view returns(bytes)
func (_RSProof *RSProofSession) KZGVK() ([]byte, error) {
	return _RSProof.Contract.KZGVK(&_RSProof.CallOpts)
}

// KZGVK is a free data retrieval call binding the contract method 0x23178d22.
//
// Solidity: function KZG_VK() view returns(bytes)
func (_RSProof *RSProofCallerSession) KZGVK() ([]byte, error) {
	return _RSProof.Contract.KZGVK(&_RSProof.CallOpts)
}

// Bank is a free data retrieval call binding the contract method 0x76cdb03b.
//
// Solidity: function bank() view returns(address)
func (_RSProof *RSProofCaller) Bank(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _RSProof.contract.Call(opts, &out, "bank")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Bank is a free data retrieval call binding the contract method 0x76cdb03b.
//
// Solidity: function bank() view returns(address)
func (_RSProof *RSProofSession) Bank() (common.Address, error) {
	return _RSProof.Contract.Bank(&_RSProof.CallOpts)
}

// Bank is a free data retrieval call binding the contract method 0x76cdb03b.
//
// Solidity: function bank() view returns(address)
func (_RSProof *RSProofCallerSession) Bank() (common.Address, error) {
	return _RSProof.Contract.Bank(&_RSProof.CallOpts)
}

// BasePenalty is a free data retrieval call binding the contract method 0x5712e98c.
//
// Solidity: function basePenalty() view returns(uint256)
func (_RSProof *RSProofCaller) BasePenalty(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _RSProof.contract.Call(opts, &out, "basePenalty")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BasePenalty is a free data retrieval call binding the contract method 0x5712e98c.
//
// Solidity: function basePenalty() view returns(uint256)
func (_RSProof *RSProofSession) BasePenalty() (*big.Int, error) {
	return _RSProof.Contract.BasePenalty(&_RSProof.CallOpts)
}

// BasePenalty is a free data retrieval call binding the contract method 0x5712e98c.
//
// Solidity: function basePenalty() view returns(uint256)
func (_RSProof *RSProofCallerSession) BasePenalty() (*big.Int, error) {
	return _RSProof.Contract.BasePenalty(&_RSProof.CallOpts)
}

// Current is a free data retrieval call binding the contract method 0x9fa6a6e3.
//
// Solidity: function current() view returns(uint64)
func (_RSProof *RSProofCaller) Current(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _RSProof.contract.Call(opts, &out, "current")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// Current is a free data retrieval call binding the contract method 0x9fa6a6e3.
//
// Solidity: function current() view returns(uint64)
func (_RSProof *RSProofSession) Current() (uint64, error) {
	return _RSProof.Contract.Current(&_RSProof.CallOpts)
}

// Current is a free data retrieval call binding the contract method 0x9fa6a6e3.
//
// Solidity: function current() view returns(uint64)
func (_RSProof *RSProofCallerSession) Current() (uint64, error) {
	return _RSProof.Contract.Current(&_RSProof.CallOpts)
}

// GetProof is a free data retrieval call binding the contract method 0xa617efec.
//
// Solidity: function getProof(uint64 _pi, uint8 _pri) view returns((bool,address,uint256))
func (_RSProof *RSProofCaller) GetProof(opts *bind.CallOpts, _pi uint64, _pri uint8) (RSProofProofInfo, error) {
	var out []interface{}
	err := _RSProof.contract.Call(opts, &out, "getProof", _pi, _pri)

	if err != nil {
		return *new(RSProofProofInfo), err
	}

	out0 := *abi.ConvertType(out[0], new(RSProofProofInfo)).(*RSProofProofInfo)

	return out0, err

}

// GetProof is a free data retrieval call binding the contract method 0xa617efec.
//
// Solidity: function getProof(uint64 _pi, uint8 _pri) view returns((bool,address,uint256))
func (_RSProof *RSProofSession) GetProof(_pi uint64, _pri uint8) (RSProofProofInfo, error) {
	return _RSProof.Contract.GetProof(&_RSProof.CallOpts, _pi, _pri)
}

// GetProof is a free data retrieval call binding the contract method 0xa617efec.
//
// Solidity: function getProof(uint64 _pi, uint8 _pri) view returns((bool,address,uint256))
func (_RSProof *RSProofCallerSession) GetProof(_pi uint64, _pri uint8) (RSProofProofInfo, error) {
	return _RSProof.Contract.GetProof(&_RSProof.CallOpts, _pi, _pri)
}

// GetVKRoot is a free data retrieval call binding the contract method 0xb29cc009.
//
// Solidity: function getVKRoot(uint8 _rsn, uint8 _rsk) view returns(uint256)
func (_RSProof *RSProofCaller) GetVKRoot(opts *bind.CallOpts, _rsn uint8, _rsk uint8) (*big.Int, error) {
	var out []interface{}
	err := _RSProof.contract.Call(opts, &out, "getVKRoot", _rsn, _rsk)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetVKRoot is a free data retrieval call binding the contract method 0xb29cc009.
//
// Solidity: function getVKRoot(uint8 _rsn, uint8 _rsk) view returns(uint256)
func (_RSProof *RSProofSession) GetVKRoot(_rsn uint8, _rsk uint8) (*big.Int, error) {
	return _RSProof.Contract.GetVKRoot(&_RSProof.CallOpts, _rsn, _rsk)
}

// GetVKRoot is a free data retrieval call binding the contract method 0xb29cc009.
//
// Solidity: function getVKRoot(uint8 _rsn, uint8 _rsk) view returns(uint256)
func (_RSProof *RSProofCallerSession) GetVKRoot(_rsn uint8, _rsk uint8) (*big.Int, error) {
	return _RSProof.Contract.GetVKRoot(&_RSProof.CallOpts, _rsn, _rsk)
}

// MinProveTime is a free data retrieval call binding the contract method 0x0bd26cb5.
//
// Solidity: function minProveTime() view returns(uint256)
func (_RSProof *RSProofCaller) MinProveTime(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _RSProof.contract.Call(opts, &out, "minProveTime")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MinProveTime is a free data retrieval call binding the contract method 0x0bd26cb5.
//
// Solidity: function minProveTime() view returns(uint256)
func (_RSProof *RSProofSession) MinProveTime() (*big.Int, error) {
	return _RSProof.Contract.MinProveTime(&_RSProof.CallOpts)
}

// MinProveTime is a free data retrieval call binding the contract method 0x0bd26cb5.
//
// Solidity: function minProveTime() view returns(uint256)
func (_RSProof *RSProofCallerSession) MinProveTime() (*big.Int, error) {
	return _RSProof.Contract.MinProveTime(&_RSProof.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_RSProof *RSProofCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _RSProof.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_RSProof *RSProofSession) Owner() (common.Address, error) {
	return _RSProof.Contract.Owner(&_RSProof.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_RSProof *RSProofCallerSession) Owner() (common.Address, error) {
	return _RSProof.Contract.Owner(&_RSProof.CallOpts)
}

// Challenge is a paid mutator transaction binding the contract method 0x6c05dff8.
//
// Solidity: function challenge(bytes _pn, bytes _rn, uint8 _pri) returns()
func (_RSProof *RSProofTransactor) Challenge(opts *bind.TransactOpts, _pn []byte, _rn []byte, _pri uint8) (*types.Transaction, error) {
	return _RSProof.contract.Transact(opts, "challenge", _pn, _rn, _pri)
}

// Challenge is a paid mutator transaction binding the contract method 0x6c05dff8.
//
// Solidity: function challenge(bytes _pn, bytes _rn, uint8 _pri) returns()
func (_RSProof *RSProofSession) Challenge(_pn []byte, _rn []byte, _pri uint8) (*types.Transaction, error) {
	return _RSProof.Contract.Challenge(&_RSProof.TransactOpts, _pn, _rn, _pri)
}

// Challenge is a paid mutator transaction binding the contract method 0x6c05dff8.
//
// Solidity: function challenge(bytes _pn, bytes _rn, uint8 _pri) returns()
func (_RSProof *RSProofTransactorSession) Challenge(_pn []byte, _rn []byte, _pri uint8) (*types.Transaction, error) {
	return _RSProof.Contract.Challenge(&_RSProof.TransactOpts, _pn, _rn, _pri)
}

// Check is a paid mutator transaction binding the contract method 0xf354cd5f.
//
// Solidity: function check(uint64 _pi, uint64 _ri, uint8 _pri) returns()
func (_RSProof *RSProofTransactor) Check(opts *bind.TransactOpts, _pi uint64, _ri uint64, _pri uint8) (*types.Transaction, error) {
	return _RSProof.contract.Transact(opts, "check", _pi, _ri, _pri)
}

// Check is a paid mutator transaction binding the contract method 0xf354cd5f.
//
// Solidity: function check(uint64 _pi, uint64 _ri, uint8 _pri) returns()
func (_RSProof *RSProofSession) Check(_pi uint64, _ri uint64, _pri uint8) (*types.Transaction, error) {
	return _RSProof.Contract.Check(&_RSProof.TransactOpts, _pi, _ri, _pri)
}

// Check is a paid mutator transaction binding the contract method 0xf354cd5f.
//
// Solidity: function check(uint64 _pi, uint64 _ri, uint8 _pri) returns()
func (_RSProof *RSProofTransactorSession) Check(_pi uint64, _ri uint64, _pri uint8) (*types.Transaction, error) {
	return _RSProof.Contract.Check(&_RSProof.TransactOpts, _pi, _ri, _pri)
}

// Prove is a paid mutator transaction binding the contract method 0x3c530b08.
//
// Solidity: function prove(bytes _pn, bytes _rn, uint8 _pri, bytes _proof) returns()
func (_RSProof *RSProofTransactor) Prove(opts *bind.TransactOpts, _pn []byte, _rn []byte, _pri uint8, _proof []byte) (*types.Transaction, error) {
	return _RSProof.contract.Transact(opts, "prove", _pn, _rn, _pri, _proof)
}

// Prove is a paid mutator transaction binding the contract method 0x3c530b08.
//
// Solidity: function prove(bytes _pn, bytes _rn, uint8 _pri, bytes _proof) returns()
func (_RSProof *RSProofSession) Prove(_pn []byte, _rn []byte, _pri uint8, _proof []byte) (*types.Transaction, error) {
	return _RSProof.Contract.Prove(&_RSProof.TransactOpts, _pn, _rn, _pri, _proof)
}

// Prove is a paid mutator transaction binding the contract method 0x3c530b08.
//
// Solidity: function prove(bytes _pn, bytes _rn, uint8 _pri, bytes _proof) returns()
func (_RSProof *RSProofTransactorSession) Prove(_pn []byte, _rn []byte, _pri uint8, _proof []byte) (*types.Transaction, error) {
	return _RSProof.Contract.Prove(&_RSProof.TransactOpts, _pn, _rn, _pri, _proof)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_RSProof *RSProofTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RSProof.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_RSProof *RSProofSession) RenounceOwnership() (*types.Transaction, error) {
	return _RSProof.Contract.RenounceOwnership(&_RSProof.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_RSProof *RSProofTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _RSProof.Contract.RenounceOwnership(&_RSProof.TransactOpts)
}

// SetVeirfier is a paid mutator transaction binding the contract method 0x5acbd1bf.
//
// Solidity: function setVeirfier(uint8 _rsn, uint8 _rsk, uint256 _vkroot) returns()
func (_RSProof *RSProofTransactor) SetVeirfier(opts *bind.TransactOpts, _rsn uint8, _rsk uint8, _vkroot *big.Int) (*types.Transaction, error) {
	return _RSProof.contract.Transact(opts, "setVeirfier", _rsn, _rsk, _vkroot)
}

// SetVeirfier is a paid mutator transaction binding the contract method 0x5acbd1bf.
//
// Solidity: function setVeirfier(uint8 _rsn, uint8 _rsk, uint256 _vkroot) returns()
func (_RSProof *RSProofSession) SetVeirfier(_rsn uint8, _rsk uint8, _vkroot *big.Int) (*types.Transaction, error) {
	return _RSProof.Contract.SetVeirfier(&_RSProof.TransactOpts, _rsn, _rsk, _vkroot)
}

// SetVeirfier is a paid mutator transaction binding the contract method 0x5acbd1bf.
//
// Solidity: function setVeirfier(uint8 _rsn, uint8 _rsk, uint256 _vkroot) returns()
func (_RSProof *RSProofTransactorSession) SetVeirfier(_rsn uint8, _rsk uint8, _vkroot *big.Int) (*types.Transaction, error) {
	return _RSProof.Contract.SetVeirfier(&_RSProof.TransactOpts, _rsn, _rsk, _vkroot)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_RSProof *RSProofTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _RSProof.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_RSProof *RSProofSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _RSProof.Contract.TransferOwnership(&_RSProof.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_RSProof *RSProofTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _RSProof.Contract.TransferOwnership(&_RSProof.TransactOpts, newOwner)
}

// RSProofChallengeIterator is returned from FilterChallenge and is used to iterate over the raw logs and unpacked data for Challenge events raised by the RSProof contract.
type RSProofChallengeIterator struct {
	Event *RSProofChallenge // Event containing the contract specifics and raw log

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
func (it *RSProofChallengeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RSProofChallenge)
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
		it.Event = new(RSProofChallenge)
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
func (it *RSProofChallengeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RSProofChallengeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RSProofChallenge represents a Challenge event raised by the RSProof contract.
type RSProofChallenge struct {
	S   common.Address
	Ri  uint64
	Raw types.Log // Blockchain specific contextual infos
}

// FilterChallenge is a free log retrieval operation binding the contract event 0x40871ff904cf6ca42f9dcbac1f7d50ab55381a1d0db3eedd5dd967b209f8d823.
//
// Solidity: event Challenge(address indexed _s, uint64 _ri)
func (_RSProof *RSProofFilterer) FilterChallenge(opts *bind.FilterOpts, _s []common.Address) (*RSProofChallengeIterator, error) {

	var _sRule []interface{}
	for _, _sItem := range _s {
		_sRule = append(_sRule, _sItem)
	}

	logs, sub, err := _RSProof.contract.FilterLogs(opts, "Challenge", _sRule)
	if err != nil {
		return nil, err
	}
	return &RSProofChallengeIterator{contract: _RSProof.contract, event: "Challenge", logs: logs, sub: sub}, nil
}

// WatchChallenge is a free log subscription operation binding the contract event 0x40871ff904cf6ca42f9dcbac1f7d50ab55381a1d0db3eedd5dd967b209f8d823.
//
// Solidity: event Challenge(address indexed _s, uint64 _ri)
func (_RSProof *RSProofFilterer) WatchChallenge(opts *bind.WatchOpts, sink chan<- *RSProofChallenge, _s []common.Address) (event.Subscription, error) {

	var _sRule []interface{}
	for _, _sItem := range _s {
		_sRule = append(_sRule, _sItem)
	}

	logs, sub, err := _RSProof.contract.WatchLogs(opts, "Challenge", _sRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RSProofChallenge)
				if err := _RSProof.contract.UnpackLog(event, "Challenge", log); err != nil {
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

// ParseChallenge is a log parse operation binding the contract event 0x40871ff904cf6ca42f9dcbac1f7d50ab55381a1d0db3eedd5dd967b209f8d823.
//
// Solidity: event Challenge(address indexed _s, uint64 _ri)
func (_RSProof *RSProofFilterer) ParseChallenge(log types.Log) (*RSProofChallenge, error) {
	event := new(RSProofChallenge)
	if err := _RSProof.contract.UnpackLog(event, "Challenge", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RSProofFakeIterator is returned from FilterFake and is used to iterate over the raw logs and unpacked data for Fake events raised by the RSProof contract.
type RSProofFakeIterator struct {
	Event *RSProofFake // Event containing the contract specifics and raw log

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
func (it *RSProofFakeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RSProofFake)
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
		it.Event = new(RSProofFake)
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
func (it *RSProofFakeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RSProofFakeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RSProofFake represents a Fake event raised by the RSProof contract.
type RSProofFake struct {
	S   common.Address
	Ri  uint64
	Raw types.Log // Blockchain specific contextual infos
}

// FilterFake is a free log retrieval operation binding the contract event 0x0fcd3bb0ce91639d74185d625544f2b5714ecd14a6a6d9ccb65b2526e9ac12f5.
//
// Solidity: event Fake(address indexed _s, uint64 _ri)
func (_RSProof *RSProofFilterer) FilterFake(opts *bind.FilterOpts, _s []common.Address) (*RSProofFakeIterator, error) {

	var _sRule []interface{}
	for _, _sItem := range _s {
		_sRule = append(_sRule, _sItem)
	}

	logs, sub, err := _RSProof.contract.FilterLogs(opts, "Fake", _sRule)
	if err != nil {
		return nil, err
	}
	return &RSProofFakeIterator{contract: _RSProof.contract, event: "Fake", logs: logs, sub: sub}, nil
}

// WatchFake is a free log subscription operation binding the contract event 0x0fcd3bb0ce91639d74185d625544f2b5714ecd14a6a6d9ccb65b2526e9ac12f5.
//
// Solidity: event Fake(address indexed _s, uint64 _ri)
func (_RSProof *RSProofFilterer) WatchFake(opts *bind.WatchOpts, sink chan<- *RSProofFake, _s []common.Address) (event.Subscription, error) {

	var _sRule []interface{}
	for _, _sItem := range _s {
		_sRule = append(_sRule, _sItem)
	}

	logs, sub, err := _RSProof.contract.WatchLogs(opts, "Fake", _sRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RSProofFake)
				if err := _RSProof.contract.UnpackLog(event, "Fake", log); err != nil {
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

// ParseFake is a log parse operation binding the contract event 0x0fcd3bb0ce91639d74185d625544f2b5714ecd14a6a6d9ccb65b2526e9ac12f5.
//
// Solidity: event Fake(address indexed _s, uint64 _ri)
func (_RSProof *RSProofFilterer) ParseFake(log types.Log) (*RSProofFake, error) {
	event := new(RSProofFake)
	if err := _RSProof.contract.UnpackLog(event, "Fake", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RSProofOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the RSProof contract.
type RSProofOwnershipTransferredIterator struct {
	Event *RSProofOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *RSProofOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RSProofOwnershipTransferred)
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
		it.Event = new(RSProofOwnershipTransferred)
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
func (it *RSProofOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RSProofOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RSProofOwnershipTransferred represents a OwnershipTransferred event raised by the RSProof contract.
type RSProofOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_RSProof *RSProofFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*RSProofOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _RSProof.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &RSProofOwnershipTransferredIterator{contract: _RSProof.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_RSProof *RSProofFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *RSProofOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _RSProof.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RSProofOwnershipTransferred)
				if err := _RSProof.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_RSProof *RSProofFilterer) ParseOwnershipTransferred(log types.Log) (*RSProofOwnershipTransferred, error) {
	event := new(RSProofOwnershipTransferred)
	if err := _RSProof.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RSProofProveIterator is returned from FilterProve and is used to iterate over the raw logs and unpacked data for Prove events raised by the RSProof contract.
type RSProofProveIterator struct {
	Event *RSProofProve // Event containing the contract specifics and raw log

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
func (it *RSProofProveIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RSProofProve)
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
		it.Event = new(RSProofProve)
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
func (it *RSProofProveIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RSProofProveIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RSProofProve represents a Prove event raised by the RSProof contract.
type RSProofProve struct {
	S   common.Address
	Ri  uint64
	Raw types.Log // Blockchain specific contextual infos
}

// FilterProve is a free log retrieval operation binding the contract event 0x4387d8e339e908dfa927fdd9f9555518616803c51d321dce6e75d1a647bd1659.
//
// Solidity: event Prove(address indexed _s, uint64 _ri)
func (_RSProof *RSProofFilterer) FilterProve(opts *bind.FilterOpts, _s []common.Address) (*RSProofProveIterator, error) {

	var _sRule []interface{}
	for _, _sItem := range _s {
		_sRule = append(_sRule, _sItem)
	}

	logs, sub, err := _RSProof.contract.FilterLogs(opts, "Prove", _sRule)
	if err != nil {
		return nil, err
	}
	return &RSProofProveIterator{contract: _RSProof.contract, event: "Prove", logs: logs, sub: sub}, nil
}

// WatchProve is a free log subscription operation binding the contract event 0x4387d8e339e908dfa927fdd9f9555518616803c51d321dce6e75d1a647bd1659.
//
// Solidity: event Prove(address indexed _s, uint64 _ri)
func (_RSProof *RSProofFilterer) WatchProve(opts *bind.WatchOpts, sink chan<- *RSProofProve, _s []common.Address) (event.Subscription, error) {

	var _sRule []interface{}
	for _, _sItem := range _s {
		_sRule = append(_sRule, _sItem)
	}

	logs, sub, err := _RSProof.contract.WatchLogs(opts, "Prove", _sRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RSProofProve)
				if err := _RSProof.contract.UnpackLog(event, "Prove", log); err != nil {
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

// ParseProve is a log parse operation binding the contract event 0x4387d8e339e908dfa927fdd9f9555518616803c51d321dce6e75d1a647bd1659.
//
// Solidity: event Prove(address indexed _s, uint64 _ri)
func (_RSProof *RSProofFilterer) ParseProve(log types.Log) (*RSProofProve, error) {
	event := new(RSProofProve)
	if err := _RSProof.contract.UnpackLog(event, "Prove", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
