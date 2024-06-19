// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package everify

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

// IEVerifyCInfo is an auto generated low-level Go binding around an user-defined struct.
type IEVerifyCInfo struct {
	PreCnt  uint64
	Count   uint64
	Round   uint8
	RoundAt uint8
	CIndex  uint64
	QIndex  uint8
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

// EVerifyMetaData contains all meta data concerning the EVerify contract.
var EVerifyMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_b\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"KZG_G1\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"KZG_VK\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"bank\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_a\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"_ep\",\"type\":\"uint64\"},{\"internalType\":\"uint8\",\"name\":\"_qIndex\",\"type\":\"uint8\"}],\"name\":\"chalCom\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_a\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"_ep\",\"type\":\"uint64\"},{\"internalType\":\"uint8\",\"name\":\"_qIndex\",\"type\":\"uint8\"}],\"name\":\"chalOne\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_a\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"_ep\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"_sum\",\"type\":\"bytes\"}],\"name\":\"challenge\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_a\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"_seed\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"_count\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"_pcount\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"_index\",\"type\":\"uint64\"}],\"name\":\"choose\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"divide\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_a\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"_ep\",\"type\":\"uint64\"}],\"name\":\"getCInfo\",\"outputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"preCnt\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"count\",\"type\":\"uint64\"},{\"internalType\":\"uint8\",\"name\":\"round\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"roundAt\",\"type\":\"uint8\"},{\"internalType\":\"uint64\",\"name\":\"cIndex\",\"type\":\"uint64\"},{\"internalType\":\"uint8\",\"name\":\"qIndex\",\"type\":\"uint8\"}],\"internalType\":\"structIEVerify.CInfo\",\"name\":\"_ci\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_a\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"_ep\",\"type\":\"uint64\"},{\"internalType\":\"uint8\",\"name\":\"_i\",\"type\":\"uint8\"}],\"name\":\"getCommit\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"_cnt\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"_pcnt\",\"type\":\"uint64\"}],\"name\":\"getOrder\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_ptype\",\"type\":\"string\"}],\"name\":\"getVKRoot\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_a\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"_ep\",\"type\":\"uint64\"},{\"internalType\":\"bytes[]\",\"name\":\"_com\",\"type\":\"bytes[]\"},{\"internalType\":\"bytes\",\"name\":\"_proof\",\"type\":\"bytes\"}],\"name\":\"proveCom\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_a\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"_ep\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"_wroot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"_proof\",\"type\":\"bytes\"}],\"name\":\"proveKZG\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_a\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"_ep\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"_com\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"_proof\",\"type\":\"bytes\"}],\"name\":\"proveOne\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_ptype\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"_vkroot\",\"type\":\"uint256\"}],\"name\":\"setVeirfier\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x60806040526040518060800160405280606081526020016158b3606091396002908161002b9190610400565b5060405180606001604052806030815260200161591360309139600390816100539190610400565b50600860045f6101000a81548160ff021916908360ff16021790555034801561007a575f80fd5b50604051615943380380615943833981810160405281019061009c919061052d565b6100b86100ad6100fe60201b60201c565b61010560201b60201c565b8060015f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050610558565b5f33905090565b5f805f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050815f806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508173ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a35050565b5f81519050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b7f4e487b71000000000000000000000000000000000000000000000000000000005f52602260045260245ffd5b5f600282049050600182168061024157607f821691505b602082108103610254576102536101fd565b5b50919050565b5f819050815f5260205f209050919050565b5f6020601f8301049050919050565b5f82821b905092915050565b5f600883026102b67fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8261027b565b6102c0868361027b565b95508019841693508086168417925050509392505050565b5f819050919050565b5f819050919050565b5f6103046102ff6102fa846102d8565b6102e1565b6102d8565b9050919050565b5f819050919050565b61031d836102ea565b6103316103298261030b565b848454610287565b825550505050565b5f90565b610345610339565b610350818484610314565b505050565b5b81811015610373576103685f8261033d565b600181019050610356565b5050565b601f8211156103b8576103898161025a565b6103928461026c565b810160208510156103a1578190505b6103b56103ad8561026c565b830182610355565b50505b505050565b5f82821c905092915050565b5f6103d85f19846008026103bd565b1980831691505092915050565b5f6103f083836103c9565b9150826002028217905092915050565b610409826101c6565b67ffffffffffffffff811115610422576104216101d0565b5b61042c825461022a565b610437828285610377565b5f60209050601f831160018114610468575f8415610456578287015190505b61046085826103e5565b8655506104c7565b601f1984166104768661025a565b5f5b8281101561049d57848901518255600182019150602085019450602081019050610478565b868310156104ba57848901516104b6601f8916826103c9565b8355505b6001600288020188555050505b505050505050565b5f80fd5b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f6104fc826104d3565b9050919050565b61050c816104f2565b8114610516575f80fd5b50565b5f8151905061052781610503565b92915050565b5f60208284031215610542576105416104cf565b5b5f61054f84828501610519565b91505092915050565b61534e806105655f395ff3fe608060405234801561000f575f80fd5b506004361061011f575f3560e01c80639949edae116100ab578063b54753b81161006f578063b54753b814610319578063b8b55ddd14610349578063bd42f4c214610367578063ddbf603914610397578063f2fde38b146103c75761011f565b80639949edae1461023b5780639b00f9591461026b578063a22659111461029b578063b1ea66de146102cb578063b53b8d7a146102fb5761011f565b80634c077132116100f25780634c077132146101a95780635a9b427d146101d9578063715018a6146101f557806376cdb03b146101ff5780638da5cb5b1461021d5761011f565b8063141df27c146101235780632049af341461013f57806323178d221461015b5780633dcc3e1b14610179575b5f80fd5b61013d600480360381019061013891906137c1565b6103e3565b005b61015960048036038101906101549190613983565b610497565b005b610163610942565b6040516101709190613a7d565b60405180910390f35b610193600480360381019061018e9190613a9d565b6109ce565b6040516101a09190613af6565b60405180910390f35b6101c360048036038101906101be9190613b0f565b6109e1565b6040516101d09190613b95565b60405180910390f35b6101f360048036038101906101ee9190613bae565b6109fa565b005b6101fd610f42565b005b610207610f55565b6040516102149190613c29565b60405180910390f35b610225610f7a565b6040516102329190613c29565b60405180910390f35b61025560048036038101906102509190613d24565b610fa1565b6040516102629190613af6565b60405180910390f35b61028560048036038101906102809190613dc0565b611668565b6040516102929190613af6565b60405180910390f35b6102b560048036038101906102b09190613e5c565b612164565b6040516102c29190613eb2565b60405180910390f35b6102e560048036038101906102e09190613ecb565b61218b565b6040516102f29190613fa0565b60405180910390f35b6103036124dc565b6040516103109190613af6565b60405180910390f35b610333600480360381019061032e9190613fe3565b6124ee565b6040516103409190613af6565b60405180910390f35b6103516129d3565b60405161035e9190613a7d565b60405180910390f35b610381600480360381019061037c9190613fe3565b612a5f565b60405161038e9190613af6565b60405180910390f35b6103b160048036038101906103ac9190613fe3565b612f44565b6040516103be9190613a7d565b60405180910390f35b6103e160048036038101906103dc9190614033565b61304b565b005b6103eb6130cd565b5f6007836040516103fc91906140a2565b9081526020016040518091039020541461044b576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161044290614112565b60405180910390fd5b7f30644e72e131a029b85045b68181585d2833e84879b9709143e1f593f0000001810690508060078360405161048191906140a2565b9081526020016040518091039020819055505050565b5f60015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663693ec85e6040518163ffffffff1660e01b81526004016104f09061417a565b6020604051808303815f875af115801561050c573d5f803e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061053091906141ac565b73ffffffffffffffffffffffffffffffffffffffff166312a02c826001866105589190614204565b6040518263ffffffff1660e01b81526004016105749190613b95565b60408051808303815f875af115801561058f573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906105b39190614267565b9150505f81866040516020016105ca92919061430a565b604051602081830303815290604052805190602001205f1c90507f12ab655e9a2ca55660b44d1e5c37b00159aa76fed00000010a11800000000001810690505f846106148361314b565b60036040516020016106289392919061445e565b604051602081830303815290604052805190602001205f1c90507f30644e72e131a029b85045b68181585d2833e84879b9709143e1f593f0000001810690505f6007604051610676906144dc565b908152602001604051809103902054036106c5576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016106bc9061453a565b60405180910390fd5b5f600267ffffffffffffffff8111156106e1576106e061366a565b5b60405190808252806020026020018201604052801561070f5781602001602082028036833780820191505090505b5090506007604051610720906144dc565b908152602001604051809103902054815f8151811061074257610741614558565b5b602002602001018181525050818160018151811061076357610762614558565b5b6020026020010181815250505f60015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663693ec85e6040518163ffffffff1660e01b81526004016107c8906145a7565b6020604051808303815f875af11580156107e4573d5f803e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061080891906141ac565b90505f73ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1603610878576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161086f9061460f565b60405180910390fd5b5f8173ffffffffffffffffffffffffffffffffffffffff16637e4f7a8a88856040518363ffffffff1660e01b81526004016108b49291906146e4565b6020604051808303815f875af11580156108d0573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906108f4919061474e565b905080610936576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161092d906147c3565b60405180910390fd5b50505050505050505050565b6003805461094f9061439c565b80601f016020809104026020016040519081016040528092919081815260200182805461097b9061439c565b80156109c65780601f1061099d576101008083540402835291602001916109c6565b820191905f5260205f20905b8154815290600101906020018083116109a957829003601f168201915b505050505081565b5f6109d98383613221565b905092915050565b5f6109ef868686868661337c565b905095945050505050565b610a0261341b565b5f60015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663693ec85e6040518163ffffffff1660e01b8152600401610a5b9061482b565b6020604051808303815f875af1158015610a77573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610a9b91906141ac565b90505f8173ffffffffffffffffffffffffffffffffffffffff166350979d7e86866040518363ffffffff1660e01b8152600401610ad9929190614849565b6020604051808303815f875af1158015610af5573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610b199190614884565b90505f8273ffffffffffffffffffffffffffffffffffffffff166350979d7e87600188610b4691906148af565b6040518363ffffffff1660e01b8152600401610b63929190614849565b6020604051808303815f875af1158015610b7f573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610ba39190614884565b9050610baf8282613221565b60055f8873ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f8767ffffffffffffffff1667ffffffffffffffff1681526020019081526020015f205f0160106101000a81548160ff021916908360ff1602179055508060055f8873ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f8767ffffffffffffffff1667ffffffffffffffff1681526020019081526020015f205f015f6101000a81548167ffffffffffffffff021916908367ffffffffffffffff1602179055508160055f8873ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f8767ffffffffffffffff1667ffffffffffffffff1681526020019081526020015f205f0160086101000a81548167ffffffffffffffff021916908367ffffffffffffffff160217905550600160055f8873ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f8767ffffffffffffffff1667ffffffffffffffff1681526020019081526020015f205f0160116101000a81548160ff021916908360ff1602179055505f60055f8873ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f8767ffffffffffffffff1667ffffffffffffffff1681526020019081526020015f205f0160126101000a81548167ffffffffffffffff021916908367ffffffffffffffff1602179055505f60055f8873ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f8767ffffffffffffffff1667ffffffffffffffff1681526020019081526020015f205f01601a6101000a81548160ff021916908360ff1602179055508360065f8873ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f8767ffffffffffffffff1667ffffffffffffffff1681526020019081526020015f205f8060ff1681526020019081526020015f209081610f399190614a75565b50505050505050565b610f4a6130cd565b610f535f613522565b565b60015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b5f805f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b5f610faa61341b565b60045f9054906101000a900460ff1660ff16835114610ffe576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610ff590614b8e565b60405180910390fd5b5f60065f8773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f8667ffffffffffffffff1667ffffffffffffffff1681526020019081526020015f205f60055f8973ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f8867ffffffffffffffff1667ffffffffffffffff1681526020019081526020015f205f01601a9054906101000a900460ff1660ff1660ff1681526020019081526020015f2080546110ee9061439c565b80601f016020809104026020016040519081016040528092919081815260200182805461111a9061439c565b80156111655780601f1061113c57610100808354040283529160200191611165565b820191905f5260205f20905b81548152906001019060200180831161114857829003601f168201915b505050505090505f8151116111af576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016111a690614bf6565b60405180910390fd5b5f5b60045f9054906101000a900460ff1660ff168160ff16101561121b5781858260ff16815181106111e4576111e3614558565b5b60200260200101516040516020016111fd929190614c14565b604051602081830303815290604052915080806001019150506111b1565b505f81805190602001205f1c90507f30644e72e131a029b85045b68181585d2833e84879b9709143e1f593f0000001810690505f600760405161125d90614c81565b908152602001604051809103902054036112ac576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016112a39061453a565b60405180910390fd5b5f600267ffffffffffffffff8111156112c8576112c761366a565b5b6040519080825280602002602001820160405280156112f65781602001602082028036833780820191505090505b509050600760405161130790614c81565b908152602001604051809103902054815f8151811061132957611328614558565b5b602002602001018181525050818160018151811061134a57611349614558565b5b6020026020010181815250505f60015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663693ec85e6040518163ffffffff1660e01b81526004016113af90614cb7565b6020604051808303815f875af11580156113cb573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906113ef91906141ac565b90505f73ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff160361145f576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016114569061460f565b60405180910390fd5b5f8173ffffffffffffffffffffffffffffffffffffffff16637e4f7a8a88856040518363ffffffff1660e01b815260040161149b9291906146e4565b6020604051808303815f875af11580156114b7573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906114db919061474e565b90508061151d576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401611514906147c3565b60405180910390fd5b5f5b60045f9054906101000a900460ff1660ff168160ff1610156115e857888160ff168151811061155157611550614558565b5b602002602001015160065f8d73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f8c67ffffffffffffffff1667ffffffffffffffff1681526020019081526020015f205f8360ff1660ff1681526020019081526020015f2090816115da9190614a75565b50808060010191505061151f565b5060055f8b73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f8a67ffffffffffffffff1667ffffffffffffffff1681526020019081526020015f205f0160119054906101000a900460ff1695505050505050949350505050565b5f8060055f8773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f8667ffffffffffffffff1667ffffffffffffffff1681526020019081526020015f205f0160119054906101000a900460ff169050600160055f8873ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f8767ffffffffffffffff1667ffffffffffffffff1681526020019081526020015f205f0160109054906101000a900460ff166117589190614cd5565b60ff168160ff161461179f576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161179690614d53565b60405180910390fd5b5f60015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663693ec85e6040518163ffffffff1660e01b81526004016117f89061417a565b6020604051808303815f875af1158015611814573d5f803e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061183891906141ac565b73ffffffffffffffffffffffffffffffffffffffff166312a02c826001886118609190614204565b6040518263ffffffff1660e01b815260040161187c9190613b95565b60408051808303815f875af1158015611897573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906118bb9190614267565b9150505f611a2d888360055f8c73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f8b67ffffffffffffffff1667ffffffffffffffff1681526020019081526020015f205f0160089054906101000a900467ffffffffffffffff1660055f8d73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f8c67ffffffffffffffff1667ffffffffffffffff1681526020019081526020015f205f015f9054906101000a900467ffffffffffffffff1660055f8e73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f8d67ffffffffffffffff1667ffffffffffffffff1681526020019081526020015f205f0160129054906101000a900467ffffffffffffffff1661337c565b90505f60015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663693ec85e6040518163ffffffff1660e01b8152600401611a889061482b565b6020604051808303815f875af1158015611aa4573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190611ac891906141ac565b73ffffffffffffffffffffffffffffffffffffffff1663894d35f88a848a6040518463ffffffff1660e01b8152600401611b0493929190614d71565b602060405180830381865afa158015611b1f573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190611b439190614884565b90508767ffffffffffffffff168167ffffffffffffffff1611611bec5760028054611b6d9061439c565b80601f0160208091040260200160405190810160405280929190818152602001828054611b999061439c565b8015611be45780601f10611bbb57610100808354040283529160200191611be4565b820191905f5260205f20905b815481529060010190602001808311611bc757829003601f168201915b505050505096505b5f60065f8b73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f8a67ffffffffffffffff1667ffffffffffffffff1681526020019081526020015f205f60055f8d73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f8c67ffffffffffffffff1667ffffffffffffffff1681526020019081526020015f205f01601a9054906101000a900460ff1660ff1660ff1681526020019081526020015f208054611cdc9061439c565b80601f0160208091040260200160405190810160405280929190818152602001828054611d089061439c565b8015611d535780601f10611d2a57610100808354040283529160200191611d53565b820191905f5260205f20905b815481529060010190602001808311611d3657829003601f168201915b505050505090505f8460055f8d73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f8c67ffffffffffffffff1667ffffffffffffffff1681526020019081526020015f205f0160129054906101000a900467ffffffffffffffff16604051602001611de4929190614de1565b604051602081830303815290604052805190602001205f1c90507f12ab655e9a2ca55660b44d1e5c37b00159aa76fed00000010a11800000000001810690505f828a611e2f8461314b565b604051602001611e4193929190614e0c565b604051602081830303815290604052805190602001205f1c90507f30644e72e131a029b85045b68181585d2833e84879b9709143e1f593f0000001810690505f6007604051611e8f90614e86565b90815260200160405180910390205403611ede576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401611ed59061453a565b60405180910390fd5b5f600267ffffffffffffffff811115611efa57611ef961366a565b5b604051908082528060200260200182016040528015611f285781602001602082028036833780820191505090505b5090506007604051611f3990614e86565b908152602001604051809103902054815f81518110611f5b57611f5a614558565b5b6020026020010181815250508181600181518110611f7c57611f7b614558565b5b6020026020010181815250505f60015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663693ec85e6040518163ffffffff1660e01b8152600401611fe190614ebc565b6020604051808303815f875af1158015611ffd573d5f803e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061202191906141ac565b90505f73ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1603612091576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016120889061460f565b60405180910390fd5b5f8173ffffffffffffffffffffffffffffffffffffffff16637e4f7a8a8d856040518363ffffffff1660e01b81526004016120cd9291906146e4565b6020604051808303815f875af11580156120e9573d5f803e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061210d919061474e565b90508061214f576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401612146906147c3565b60405180910390fd5b899a5050505050505050505050949350505050565b5f60078260405161217591906140a2565b9081526020016040518091039020549050919050565b6121936135ea565b60055f8473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f8367ffffffffffffffff1667ffffffffffffffff1681526020019081526020015f205f015f9054906101000a900467ffffffffffffffff16815f019067ffffffffffffffff16908167ffffffffffffffff168152505060055f8473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f8367ffffffffffffffff1667ffffffffffffffff1681526020019081526020015f205f0160089054906101000a900467ffffffffffffffff16816020019067ffffffffffffffff16908167ffffffffffffffff168152505060055f8473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f8367ffffffffffffffff1667ffffffffffffffff1681526020019081526020015f205f0160109054906101000a900460ff16816040019060ff16908160ff168152505060055f8473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f8367ffffffffffffffff1667ffffffffffffffff1681526020019081526020015f205f0160119054906101000a900460ff16816060019060ff16908160ff168152505060055f8473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f8367ffffffffffffffff1667ffffffffffffffff1681526020019081526020015f205f01601a9054906101000a900460ff168160a0019060ff16908160ff168152505060055f8473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f8367ffffffffffffffff1667ffffffffffffffff1681526020019081526020015f205f0160129054906101000a900467ffffffffffffffff16816080019067ffffffffffffffff16908167ffffffffffffffff168152505092915050565b60045f9054906101000a900460ff1681565b5f6124f761341b565b60045f9054906101000a900460ff1660ff168260ff161061254d576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161254490614f24565b60405180910390fd5b5f60065f8673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f8567ffffffffffffffff1667ffffffffffffffff1681526020019081526020015f205f8460ff1660ff1681526020019081526020015f2080546125ce9061439c565b905011612610576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161260790614bf6565b60405180910390fd5b60055f8573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f8467ffffffffffffffff1667ffffffffffffffff1681526020019081526020015f205f0160109054906101000a900460ff1660ff1660055f8673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f8567ffffffffffffffff1667ffffffffffffffff1681526020019081526020015f205f0160119054906101000a900460ff1660ff1610612736576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161272d90614f8c565b60405180910390fd5b8160055f8673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f8567ffffffffffffffff1667ffffffffffffffff1681526020019081526020015f205f01601a6101000a81548160ff021916908360ff160217905550600160055f8673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f8567ffffffffffffffff1667ffffffffffffffff1681526020019081526020015f205f0160118282829054906101000a900460ff166128309190614cd5565b92506101000a81548160ff021916908360ff16021790555081600860055f8773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f8667ffffffffffffffff1667ffffffffffffffff1681526020019081526020015f205f01601a9054906101000a900460ff166128c59190614faa565b6128cf9190614cd5565b60ff1660055f8673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f8567ffffffffffffffff1667ffffffffffffffff1681526020019081526020015f205f0160126101000a81548167ffffffffffffffff021916908367ffffffffffffffff16021790555060055f8573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f8467ffffffffffffffff1667ffffffffffffffff1681526020019081526020015f205f0160119054906101000a900460ff1690509392505050565b600280546129e09061439c565b80601f0160208091040260200160405190810160405280929190818152602001828054612a0c9061439c565b8015612a575780601f10612a2e57610100808354040283529160200191612a57565b820191905f5260205f20905b815481529060010190602001808311612a3a57829003601f168201915b505050505081565b5f612a6861341b565b60055f8573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f8467ffffffffffffffff1667ffffffffffffffff1681526020019081526020015f205f0160109054906101000a900460ff1660ff1660055f8673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f8567ffffffffffffffff1667ffffffffffffffff1681526020019081526020015f205f0160119054906101000a900460ff1660ff1614612b8e576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401612b8590614d53565b60405180910390fd5b60045f9054906101000a900460ff1660ff168260ff1610612be4576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401612bdb90614f24565b60405180910390fd5b5f60065f8673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f8567ffffffffffffffff1667ffffffffffffffff1681526020019081526020015f205f8460ff1660ff1681526020019081526020015f208054612c659061439c565b905011612ca7576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401612c9e90614bf6565b60405180910390fd5b600160055f8673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f8567ffffffffffffffff1667ffffffffffffffff1681526020019081526020015f205f0160118282829054906101000a900460ff16612d269190614cd5565b92506101000a81548160ff021916908360ff16021790555081600860055f8773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f8667ffffffffffffffff1667ffffffffffffffff1681526020019081526020015f205f01601a9054906101000a900460ff16612dbb9190614faa565b612dc59190614cd5565b60ff1660055f8673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f8567ffffffffffffffff1667ffffffffffffffff1681526020019081526020015f205f0160126101000a81548167ffffffffffffffff021916908367ffffffffffffffff1602179055508160055f8673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f8567ffffffffffffffff1667ffffffffffffffff1681526020019081526020015f205f01601a6101000a81548160ff021916908360ff16021790555060055f8573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f8467ffffffffffffffff1667ffffffffffffffff1681526020019081526020015f205f0160119054906101000a900460ff1690509392505050565b606060065f8573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f8467ffffffffffffffff1667ffffffffffffffff1681526020019081526020015f205f8360ff1660ff1681526020019081526020015f208054612fc69061439c565b80601f0160208091040260200160405190810160405280929190818152602001828054612ff29061439c565b801561303d5780601f106130145761010080835404028352916020019161303d565b820191905f5260205f20905b81548152906001019060200180831161302057829003601f168201915b505050505090509392505050565b6130536130cd565b5f73ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16036130c1576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016130b890615056565b60405180910390fd5b6130ca81613522565b50565b6130d56135e3565b73ffffffffffffffffffffffffffffffffffffffff166130f3610f7a565b73ffffffffffffffffffffffffffffffffffffffff1614613149576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401613140906150be565b60405180910390fd5b565b60605f67ffffffffffffffff90505f60c0828516901b9050604082901b91505f6040838616901b9050808261318091906150dc565b9150604083901b92506040838616901c9050808261319e91906150dc565b9150604083901b925060c0838616901c905080826131bc91906150dc565b91505f603067ffffffffffffffff8111156131da576131d961366a565b5b6040519080825280601f01601f19166020018201604052801561320c5781602001600182028036833780820191505090505b50905082602082015280945050505050919050565b5f808367ffffffffffffffff160361323b575f9050613376565b5f600890505f600190505b8467ffffffffffffffff168267ffffffffffffffff1610156132955760045f9054906101000a900460ff1660ff168261327f919061510f565b915060018161328e9190614cd5565b9050613246565b60048160ff1611156132d05760045f9054906101000a900460ff1660ff16826132be9190615178565b91506001816132cd91906151a8565b90505b5b60068160ff1611156133105760045f9054906101000a900460ff1660ff16826132fa9190615178565b915060018161330991906151a8565b90506132d1565b5f848661331d91906148af565b90505b8067ffffffffffffffff168367ffffffffffffffff16101561336f5760045f9054906101000a900460ff1660ff1683613359919061510f565b92506001826133689190614cd5565b9150613320565b8193505050505b92915050565b5f8367ffffffffffffffff1682846133949190614204565b67ffffffffffffffff1610156133b75781836133b09190614204565b9050613412565b5f8367ffffffffffffffff16111561340e578267ffffffffffffffff168587846040516020016133e9939291906151dc565b604051602081830303815290604052805190602001205f1c61340b9190615218565b92505b8290505b95945050505050565b60015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663693ec85e6040518163ffffffff1660e01b815260040161347390615292565b6020604051808303815f875af115801561348f573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906134b391906141ac565b73ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614613520576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401613517906152fa565b60405180910390fd5b565b5f805f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050815f806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508173ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a35050565b5f33905090565b6040518060c001604052805f67ffffffffffffffff1681526020015f67ffffffffffffffff1681526020015f60ff1681526020015f60ff1681526020015f67ffffffffffffffff1681526020015f60ff1681525090565b5f604051905090565b5f80fd5b5f80fd5b5f80fd5b5f80fd5b5f601f19601f8301169050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b6136a08261365a565b810181811067ffffffffffffffff821117156136bf576136be61366a565b5b80604052505050565b5f6136d1613641565b90506136dd8282613697565b919050565b5f67ffffffffffffffff8211156136fc576136fb61366a565b5b6137058261365a565b9050602081019050919050565b828183375f83830152505050565b5f61373261372d846136e2565b6136c8565b90508281526020810184848401111561374e5761374d613656565b5b613759848285613712565b509392505050565b5f82601f83011261377557613774613652565b5b8135613785848260208601613720565b91505092915050565b5f819050919050565b6137a08161378e565b81146137aa575f80fd5b50565b5f813590506137bb81613797565b92915050565b5f80604083850312156137d7576137d661364a565b5b5f83013567ffffffffffffffff8111156137f4576137f361364e565b5b61380085828601613761565b9250506020613811858286016137ad565b9150509250929050565b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f6138448261381b565b9050919050565b6138548161383a565b811461385e575f80fd5b50565b5f8135905061386f8161384b565b92915050565b5f67ffffffffffffffff82169050919050565b61389181613875565b811461389b575f80fd5b50565b5f813590506138ac81613888565b92915050565b5f819050919050565b6138c4816138b2565b81146138ce575f80fd5b50565b5f813590506138df816138bb565b92915050565b5f67ffffffffffffffff8211156138ff576138fe61366a565b5b6139088261365a565b9050602081019050919050565b5f613927613922846138e5565b6136c8565b90508281526020810184848401111561394357613942613656565b5b61394e848285613712565b509392505050565b5f82601f83011261396a57613969613652565b5b813561397a848260208601613915565b91505092915050565b5f805f806080858703121561399b5761399a61364a565b5b5f6139a887828801613861565b94505060206139b98782880161389e565b93505060406139ca878288016138d1565b925050606085013567ffffffffffffffff8111156139eb576139ea61364e565b5b6139f787828801613956565b91505092959194509250565b5f81519050919050565b5f82825260208201905092915050565b5f5b83811015613a3a578082015181840152602081019050613a1f565b5f8484015250505050565b5f613a4f82613a03565b613a598185613a0d565b9350613a69818560208601613a1d565b613a728161365a565b840191505092915050565b5f6020820190508181035f830152613a958184613a45565b905092915050565b5f8060408385031215613ab357613ab261364a565b5b5f613ac08582860161389e565b9250506020613ad18582860161389e565b9150509250929050565b5f60ff82169050919050565b613af081613adb565b82525050565b5f602082019050613b095f830184613ae7565b92915050565b5f805f805f60a08688031215613b2857613b2761364a565b5b5f613b3588828901613861565b9550506020613b46888289016138d1565b9450506040613b578882890161389e565b9350506060613b688882890161389e565b9250506080613b798882890161389e565b9150509295509295909350565b613b8f81613875565b82525050565b5f602082019050613ba85f830184613b86565b92915050565b5f805f60608486031215613bc557613bc461364a565b5b5f613bd286828701613861565b9350506020613be38682870161389e565b925050604084013567ffffffffffffffff811115613c0457613c0361364e565b5b613c1086828701613956565b9150509250925092565b613c238161383a565b82525050565b5f602082019050613c3c5f830184613c1a565b92915050565b5f67ffffffffffffffff821115613c5c57613c5b61366a565b5b602082029050602081019050919050565b5f80fd5b5f613c83613c7e84613c42565b6136c8565b90508083825260208201905060208402830185811115613ca657613ca5613c6d565b5b835b81811015613ced57803567ffffffffffffffff811115613ccb57613cca613652565b5b808601613cd88982613956565b85526020850194505050602081019050613ca8565b5050509392505050565b5f82601f830112613d0b57613d0a613652565b5b8135613d1b848260208601613c71565b91505092915050565b5f805f8060808587031215613d3c57613d3b61364a565b5b5f613d4987828801613861565b9450506020613d5a8782880161389e565b935050604085013567ffffffffffffffff811115613d7b57613d7a61364e565b5b613d8787828801613cf7565b925050606085013567ffffffffffffffff811115613da857613da761364e565b5b613db487828801613956565b91505092959194509250565b5f805f8060808587031215613dd857613dd761364a565b5b5f613de587828801613861565b9450506020613df68782880161389e565b935050604085013567ffffffffffffffff811115613e1757613e1661364e565b5b613e2387828801613956565b925050606085013567ffffffffffffffff811115613e4457613e4361364e565b5b613e5087828801613956565b91505092959194509250565b5f60208284031215613e7157613e7061364a565b5b5f82013567ffffffffffffffff811115613e8e57613e8d61364e565b5b613e9a84828501613761565b91505092915050565b613eac8161378e565b82525050565b5f602082019050613ec55f830184613ea3565b92915050565b5f8060408385031215613ee157613ee061364a565b5b5f613eee85828601613861565b9250506020613eff8582860161389e565b9150509250929050565b613f1281613875565b82525050565b613f2181613adb565b82525050565b60c082015f820151613f3b5f850182613f09565b506020820151613f4e6020850182613f09565b506040820151613f616040850182613f18565b506060820151613f746060850182613f18565b506080820151613f876080850182613f09565b5060a0820151613f9a60a0850182613f18565b50505050565b5f60c082019050613fb35f830184613f27565b92915050565b613fc281613adb565b8114613fcc575f80fd5b50565b5f81359050613fdd81613fb9565b92915050565b5f805f60608486031215613ffa57613ff961364a565b5b5f61400786828701613861565b93505060206140188682870161389e565b925050604061402986828701613fcf565b9150509250925092565b5f602082840312156140485761404761364a565b5b5f61405584828501613861565b91505092915050565b5f81519050919050565b5f81905092915050565b5f61407c8261405e565b6140868185614068565b9350614096818560208601613a1d565b80840191505092915050565b5f6140ad8284614072565b915081905092915050565b5f82825260208201905092915050565b7f68617320766b726f6f74000000000000000000000000000000000000000000005f82015250565b5f6140fc600a836140b8565b9150614107826140c8565b602082019050919050565b5f6020820190508181035f830152614129816140f0565b9050919050565b7f65706f63680000000000000000000000000000000000000000000000000000005f82015250565b5f6141646005836140b8565b915061416f82614130565b602082019050919050565b5f6020820190508181035f83015261419181614158565b9050919050565b5f815190506141a68161384b565b92915050565b5f602082840312156141c1576141c061364a565b5b5f6141ce84828501614198565b91505092915050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b5f61420e82613875565b915061421983613875565b9250828201905067ffffffffffffffff811115614239576142386141d7565b5b92915050565b5f8151905061424d81613797565b92915050565b5f81519050614261816138bb565b92915050565b5f806040838503121561427d5761427c61364a565b5b5f61428a8582860161423f565b925050602061429b85828601614253565b9150509250929050565b5f819050919050565b6142bf6142ba826138b2565b6142a5565b82525050565b5f8160601b9050919050565b5f6142db826142c5565b9050919050565b5f6142ec826142d1565b9050919050565b6143046142ff8261383a565b6142e2565b82525050565b5f61431582856142ae565b60208201915061432582846142f3565b6014820191508190509392505050565b5f81905092915050565b5f61434982613a03565b6143538185614335565b9350614363818560208601613a1d565b80840191505092915050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52602260045260245ffd5b5f60028204905060018216806143b357607f821691505b6020821081036143c6576143c561436f565b5b50919050565b5f819050815f5260205f209050919050565b5f81546143ea8161439c565b6143f48186614335565b9450600182165f811461440e576001811461442357614455565b60ff1983168652811515820286019350614455565b61442c856143cc565b5f5b8381101561444d5781548189015260018201915060208101905061442e565b838801955050505b50505092915050565b5f61446982866142ae565b602082019150614479828561433f565b915061448582846143de565b9150819050949350505050565b7f6b7a6700000000000000000000000000000000000000000000000000000000005f82015250565b5f6144c6600383614068565b91506144d182614492565b600382019050919050565b5f6144e6826144ba565b9150819050919050565b7f6e6f20766b726f6f7400000000000000000000000000000000000000000000005f82015250565b5f6145246009836140b8565b915061452f826144f0565b602082019050919050565b5f6020820190508181035f83015261455181614518565b9050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52603260045260245ffd5b5f6145916003836140b8565b915061459c82614492565b602082019050919050565b5f6020820190508181035f8301526145be81614585565b9050919050565b7f6e6f2070760000000000000000000000000000000000000000000000000000005f82015250565b5f6145f96005836140b8565b9150614604826145c5565b602082019050919050565b5f6020820190508181035f830152614626816145ed565b9050919050565b5f81519050919050565b5f82825260208201905092915050565b5f819050602082019050919050565b61465f8161378e565b82525050565b5f6146708383614656565b60208301905092915050565b5f602082019050919050565b5f6146928261462d565b61469c8185614637565b93506146a783614647565b805f5b838110156146d75781516146be8882614665565b97506146c98361467c565b9250506001810190506146aa565b5085935050505092915050565b5f6040820190508181035f8301526146fc8185613a45565b905081810360208301526147108184614688565b90509392505050565b5f8115159050919050565b61472d81614719565b8114614737575f80fd5b50565b5f8151905061474881614724565b92915050565b5f602082840312156147635761476261364a565b5b5f6147708482850161473a565b91505092915050565b7f696e7620706600000000000000000000000000000000000000000000000000005f82015250565b5f6147ad6006836140b8565b91506147b882614779565b602082019050919050565b5f6020820190508181035f8301526147da816147a1565b9050919050565b7f70696563650000000000000000000000000000000000000000000000000000005f82015250565b5f6148156005836140b8565b9150614820826147e1565b602082019050919050565b5f6020820190508181035f83015261484281614809565b9050919050565b5f60408201905061485c5f830185613c1a565b6148696020830184613b86565b9392505050565b5f8151905061487e81613888565b92915050565b5f602082840312156148995761489861364a565b5b5f6148a684828501614870565b91505092915050565b5f6148b982613875565b91506148c483613875565b9250828203905067ffffffffffffffff8111156148e4576148e36141d7565b5b92915050565b5f6020601f8301049050919050565b5f82821b905092915050565b5f600883026149347fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff826148f9565b61493e86836148f9565b95508019841693508086168417925050509392505050565b5f819050919050565b5f61497961497461496f8461378e565b614956565b61378e565b9050919050565b5f819050919050565b6149928361495f565b6149a661499e82614980565b848454614905565b825550505050565b5f90565b6149ba6149ae565b6149c5818484614989565b505050565b5b818110156149e8576149dd5f826149b2565b6001810190506149cb565b5050565b601f821115614a2d576149fe816143cc565b614a07846148ea565b81016020851015614a16578190505b614a2a614a22856148ea565b8301826149ca565b50505b505050565b5f82821c905092915050565b5f614a4d5f1984600802614a32565b1980831691505092915050565b5f614a658383614a3e565b9150826002028217905092915050565b614a7e82613a03565b67ffffffffffffffff811115614a9757614a9661366a565b5b614aa1825461439c565b614aac8282856149ec565b5f60209050601f831160018114614add575f8415614acb578287015190505b614ad58582614a5a565b865550614b3c565b601f198416614aeb866143cc565b5f5b82811015614b1257848901518255600182019150602085019450602081019050614aed565b86831015614b2f5784890151614b2b601f891682614a3e565b8355505b6001600288020188555050505b505050505050565b7f696e7620636f6d000000000000000000000000000000000000000000000000005f82015250565b5f614b786007836140b8565b9150614b8382614b44565b602082019050919050565b5f6020820190508181035f830152614ba581614b6c565b9050919050565b7f6e6f2073756d00000000000000000000000000000000000000000000000000005f82015250565b5f614be06006836140b8565b9150614beb82614bac565b602082019050919050565b5f6020820190508181035f830152614c0d81614bd4565b9050919050565b5f614c1f828561433f565b9150614c2b828461433f565b91508190509392505050565b7f61646400000000000000000000000000000000000000000000000000000000005f82015250565b5f614c6b600383614068565b9150614c7682614c37565b600382019050919050565b5f614c8b82614c5f565b9150819050919050565b5f614ca16003836140b8565b9150614cac82614c37565b602082019050919050565b5f6020820190508181035f830152614cce81614c95565b9050919050565b5f614cdf82613adb565b9150614cea83613adb565b9250828201905060ff811115614d0357614d026141d7565b5b92915050565b7f696e73756620726f756e640000000000000000000000000000000000000000005f82015250565b5f614d3d600b836140b8565b9150614d4882614d09565b602082019050919050565b5f6020820190508181035f830152614d6a81614d31565b9050919050565b5f606082019050614d845f830186613c1a565b614d916020830185613b86565b8181036040830152614da38184613a45565b9050949350505050565b5f8160c01b9050919050565b5f614dc382614dad565b9050919050565b614ddb614dd682613875565b614db9565b82525050565b5f614dec82856142ae565b602082019150614dfc8284614dca565b6008820191508190509392505050565b5f614e17828661433f565b9150614e23828561433f565b9150614e2f828461433f565b9150819050949350505050565b7f6d756c00000000000000000000000000000000000000000000000000000000005f82015250565b5f614e70600383614068565b9150614e7b82614e3c565b600382019050919050565b5f614e9082614e64565b9150819050919050565b5f614ea66003836140b8565b9150614eb182614e3c565b602082019050919050565b5f6020820190508181035f830152614ed381614e9a565b9050919050565b7f696e7620710000000000000000000000000000000000000000000000000000005f82015250565b5f614f0e6005836140b8565b9150614f1982614eda565b602082019050919050565b5f6020820190508181035f830152614f3b81614f02565b9050919050565b7f65786365656420726f756e6400000000000000000000000000000000000000005f82015250565b5f614f76600c836140b8565b9150614f8182614f42565b602082019050919050565b5f6020820190508181035f830152614fa381614f6a565b9050919050565b5f614fb482613adb565b9150614fbf83613adb565b9250828202614fcd81613adb565b9150808214614fdf57614fde6141d7565b5b5092915050565b7f4f776e61626c653a206e6577206f776e657220697320746865207a65726f20615f8201527f6464726573730000000000000000000000000000000000000000000000000000602082015250565b5f6150406026836140b8565b915061504b82614fe6565b604082019050919050565b5f6020820190508181035f83015261506d81615034565b9050919050565b7f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e65725f82015250565b5f6150a86020836140b8565b91506150b382615074565b602082019050919050565b5f6020820190508181035f8301526150d58161509c565b9050919050565b5f6150e68261378e565b91506150f18361378e565b9250828201905080821115615109576151086141d7565b5b92915050565b5f61511982613875565b915061512483613875565b925082820261513281613875565b9150808214615144576151436141d7565b5b5092915050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601260045260245ffd5b5f61518282613875565b915061518d83613875565b92508261519d5761519c61514b565b5b828204905092915050565b5f6151b282613adb565b91506151bd83613adb565b9250828203905060ff8111156151d6576151d56141d7565b5b92915050565b5f6151e782866142ae565b6020820191506151f782856142f3565b6014820191506152078284614dca565b600882019150819050949350505050565b5f6152228261378e565b915061522d8361378e565b92508261523d5761523c61514b565b5b828206905092915050565b7f6570726f6f6600000000000000000000000000000000000000000000000000005f82015250565b5f61527c6006836140b8565b915061528782615248565b602082019050919050565b5f6020820190508181035f8301526152a981615270565b9050919050565b7f696e7620656300000000000000000000000000000000000000000000000000005f82015250565b5f6152e46006836140b8565b91506152ef826152b0565b602082019050919050565b5f6020820190508181035f830152615311816152d8565b905091905056fea26469706673582212201a40746d8bebbdd21acbc67eb6ebd2021f50c622e52eec9ad6da24c74863f61564736f6c634300081a0033eab9b16eb21be9efd5481512ffcd394e188282c8bd37cb5c85951e2caa9d41bbc8fc6225bf87ff54008848defe740a67fd82de55559c8ea6c2fe3d3634a9591a6d182ad44fb82305bd7fb348ca3e52d91f674f5d30afeec401914a69c5102eff3b8201b322c65a735690cf82c850b8624c29ec05400e06ba92a9aad12c37c1605812abbc9a1a11f500b3ab28b7751b52",
}

// EVerifyABI is the input ABI used to generate the binding from.
// Deprecated: Use EVerifyMetaData.ABI instead.
var EVerifyABI = EVerifyMetaData.ABI

// EVerifyBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use EVerifyMetaData.Bin instead.
var EVerifyBin = EVerifyMetaData.Bin

// DeployEVerify deploys a new Ethereum contract, binding an instance of EVerify to it.
func DeployEVerify(auth *bind.TransactOpts, backend bind.ContractBackend, _b common.Address) (common.Address, *types.Transaction, *EVerify, error) {
	parsed, err := EVerifyMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(EVerifyBin), backend, _b)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &EVerify{EVerifyCaller: EVerifyCaller{contract: contract}, EVerifyTransactor: EVerifyTransactor{contract: contract}, EVerifyFilterer: EVerifyFilterer{contract: contract}}, nil
}

// EVerify is an auto generated Go binding around an Ethereum contract.
type EVerify struct {
	EVerifyCaller     // Read-only binding to the contract
	EVerifyTransactor // Write-only binding to the contract
	EVerifyFilterer   // Log filterer for contract events
}

// EVerifyCaller is an auto generated read-only Go binding around an Ethereum contract.
type EVerifyCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EVerifyTransactor is an auto generated write-only Go binding around an Ethereum contract.
type EVerifyTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EVerifyFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type EVerifyFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EVerifySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type EVerifySession struct {
	Contract     *EVerify          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// EVerifyCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type EVerifyCallerSession struct {
	Contract *EVerifyCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// EVerifyTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type EVerifyTransactorSession struct {
	Contract     *EVerifyTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// EVerifyRaw is an auto generated low-level Go binding around an Ethereum contract.
type EVerifyRaw struct {
	Contract *EVerify // Generic contract binding to access the raw methods on
}

// EVerifyCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type EVerifyCallerRaw struct {
	Contract *EVerifyCaller // Generic read-only contract binding to access the raw methods on
}

// EVerifyTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type EVerifyTransactorRaw struct {
	Contract *EVerifyTransactor // Generic write-only contract binding to access the raw methods on
}

// NewEVerify creates a new instance of EVerify, bound to a specific deployed contract.
func NewEVerify(address common.Address, backend bind.ContractBackend) (*EVerify, error) {
	contract, err := bindEVerify(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &EVerify{EVerifyCaller: EVerifyCaller{contract: contract}, EVerifyTransactor: EVerifyTransactor{contract: contract}, EVerifyFilterer: EVerifyFilterer{contract: contract}}, nil
}

// NewEVerifyCaller creates a new read-only instance of EVerify, bound to a specific deployed contract.
func NewEVerifyCaller(address common.Address, caller bind.ContractCaller) (*EVerifyCaller, error) {
	contract, err := bindEVerify(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &EVerifyCaller{contract: contract}, nil
}

// NewEVerifyTransactor creates a new write-only instance of EVerify, bound to a specific deployed contract.
func NewEVerifyTransactor(address common.Address, transactor bind.ContractTransactor) (*EVerifyTransactor, error) {
	contract, err := bindEVerify(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &EVerifyTransactor{contract: contract}, nil
}

// NewEVerifyFilterer creates a new log filterer instance of EVerify, bound to a specific deployed contract.
func NewEVerifyFilterer(address common.Address, filterer bind.ContractFilterer) (*EVerifyFilterer, error) {
	contract, err := bindEVerify(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &EVerifyFilterer{contract: contract}, nil
}

// bindEVerify binds a generic wrapper to an already deployed contract.
func bindEVerify(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := EVerifyMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_EVerify *EVerifyRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _EVerify.Contract.EVerifyCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_EVerify *EVerifyRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EVerify.Contract.EVerifyTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_EVerify *EVerifyRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _EVerify.Contract.EVerifyTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_EVerify *EVerifyCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _EVerify.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_EVerify *EVerifyTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EVerify.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_EVerify *EVerifyTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _EVerify.Contract.contract.Transact(opts, method, params...)
}

// KZGG1 is a free data retrieval call binding the contract method 0xb8b55ddd.
//
// Solidity: function KZG_G1() view returns(bytes)
func (_EVerify *EVerifyCaller) KZGG1(opts *bind.CallOpts) ([]byte, error) {
	var out []interface{}
	err := _EVerify.contract.Call(opts, &out, "KZG_G1")

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// KZGG1 is a free data retrieval call binding the contract method 0xb8b55ddd.
//
// Solidity: function KZG_G1() view returns(bytes)
func (_EVerify *EVerifySession) KZGG1() ([]byte, error) {
	return _EVerify.Contract.KZGG1(&_EVerify.CallOpts)
}

// KZGG1 is a free data retrieval call binding the contract method 0xb8b55ddd.
//
// Solidity: function KZG_G1() view returns(bytes)
func (_EVerify *EVerifyCallerSession) KZGG1() ([]byte, error) {
	return _EVerify.Contract.KZGG1(&_EVerify.CallOpts)
}

// KZGVK is a free data retrieval call binding the contract method 0x23178d22.
//
// Solidity: function KZG_VK() view returns(bytes)
func (_EVerify *EVerifyCaller) KZGVK(opts *bind.CallOpts) ([]byte, error) {
	var out []interface{}
	err := _EVerify.contract.Call(opts, &out, "KZG_VK")

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// KZGVK is a free data retrieval call binding the contract method 0x23178d22.
//
// Solidity: function KZG_VK() view returns(bytes)
func (_EVerify *EVerifySession) KZGVK() ([]byte, error) {
	return _EVerify.Contract.KZGVK(&_EVerify.CallOpts)
}

// KZGVK is a free data retrieval call binding the contract method 0x23178d22.
//
// Solidity: function KZG_VK() view returns(bytes)
func (_EVerify *EVerifyCallerSession) KZGVK() ([]byte, error) {
	return _EVerify.Contract.KZGVK(&_EVerify.CallOpts)
}

// Bank is a free data retrieval call binding the contract method 0x76cdb03b.
//
// Solidity: function bank() view returns(address)
func (_EVerify *EVerifyCaller) Bank(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _EVerify.contract.Call(opts, &out, "bank")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Bank is a free data retrieval call binding the contract method 0x76cdb03b.
//
// Solidity: function bank() view returns(address)
func (_EVerify *EVerifySession) Bank() (common.Address, error) {
	return _EVerify.Contract.Bank(&_EVerify.CallOpts)
}

// Bank is a free data retrieval call binding the contract method 0x76cdb03b.
//
// Solidity: function bank() view returns(address)
func (_EVerify *EVerifyCallerSession) Bank() (common.Address, error) {
	return _EVerify.Contract.Bank(&_EVerify.CallOpts)
}

// Choose is a free data retrieval call binding the contract method 0x4c077132.
//
// Solidity: function choose(address _a, bytes32 _seed, uint64 _count, uint64 _pcount, uint64 _index) pure returns(uint64)
func (_EVerify *EVerifyCaller) Choose(opts *bind.CallOpts, _a common.Address, _seed [32]byte, _count uint64, _pcount uint64, _index uint64) (uint64, error) {
	var out []interface{}
	err := _EVerify.contract.Call(opts, &out, "choose", _a, _seed, _count, _pcount, _index)

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// Choose is a free data retrieval call binding the contract method 0x4c077132.
//
// Solidity: function choose(address _a, bytes32 _seed, uint64 _count, uint64 _pcount, uint64 _index) pure returns(uint64)
func (_EVerify *EVerifySession) Choose(_a common.Address, _seed [32]byte, _count uint64, _pcount uint64, _index uint64) (uint64, error) {
	return _EVerify.Contract.Choose(&_EVerify.CallOpts, _a, _seed, _count, _pcount, _index)
}

// Choose is a free data retrieval call binding the contract method 0x4c077132.
//
// Solidity: function choose(address _a, bytes32 _seed, uint64 _count, uint64 _pcount, uint64 _index) pure returns(uint64)
func (_EVerify *EVerifyCallerSession) Choose(_a common.Address, _seed [32]byte, _count uint64, _pcount uint64, _index uint64) (uint64, error) {
	return _EVerify.Contract.Choose(&_EVerify.CallOpts, _a, _seed, _count, _pcount, _index)
}

// Divide is a free data retrieval call binding the contract method 0xb53b8d7a.
//
// Solidity: function divide() view returns(uint8)
func (_EVerify *EVerifyCaller) Divide(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _EVerify.contract.Call(opts, &out, "divide")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Divide is a free data retrieval call binding the contract method 0xb53b8d7a.
//
// Solidity: function divide() view returns(uint8)
func (_EVerify *EVerifySession) Divide() (uint8, error) {
	return _EVerify.Contract.Divide(&_EVerify.CallOpts)
}

// Divide is a free data retrieval call binding the contract method 0xb53b8d7a.
//
// Solidity: function divide() view returns(uint8)
func (_EVerify *EVerifyCallerSession) Divide() (uint8, error) {
	return _EVerify.Contract.Divide(&_EVerify.CallOpts)
}

// GetCInfo is a free data retrieval call binding the contract method 0xb1ea66de.
//
// Solidity: function getCInfo(address _a, uint64 _ep) view returns((uint64,uint64,uint8,uint8,uint64,uint8) _ci)
func (_EVerify *EVerifyCaller) GetCInfo(opts *bind.CallOpts, _a common.Address, _ep uint64) (IEVerifyCInfo, error) {
	var out []interface{}
	err := _EVerify.contract.Call(opts, &out, "getCInfo", _a, _ep)

	if err != nil {
		return *new(IEVerifyCInfo), err
	}

	out0 := *abi.ConvertType(out[0], new(IEVerifyCInfo)).(*IEVerifyCInfo)

	return out0, err

}

// GetCInfo is a free data retrieval call binding the contract method 0xb1ea66de.
//
// Solidity: function getCInfo(address _a, uint64 _ep) view returns((uint64,uint64,uint8,uint8,uint64,uint8) _ci)
func (_EVerify *EVerifySession) GetCInfo(_a common.Address, _ep uint64) (IEVerifyCInfo, error) {
	return _EVerify.Contract.GetCInfo(&_EVerify.CallOpts, _a, _ep)
}

// GetCInfo is a free data retrieval call binding the contract method 0xb1ea66de.
//
// Solidity: function getCInfo(address _a, uint64 _ep) view returns((uint64,uint64,uint8,uint8,uint64,uint8) _ci)
func (_EVerify *EVerifyCallerSession) GetCInfo(_a common.Address, _ep uint64) (IEVerifyCInfo, error) {
	return _EVerify.Contract.GetCInfo(&_EVerify.CallOpts, _a, _ep)
}

// GetCommit is a free data retrieval call binding the contract method 0xddbf6039.
//
// Solidity: function getCommit(address _a, uint64 _ep, uint8 _i) view returns(bytes)
func (_EVerify *EVerifyCaller) GetCommit(opts *bind.CallOpts, _a common.Address, _ep uint64, _i uint8) ([]byte, error) {
	var out []interface{}
	err := _EVerify.contract.Call(opts, &out, "getCommit", _a, _ep, _i)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// GetCommit is a free data retrieval call binding the contract method 0xddbf6039.
//
// Solidity: function getCommit(address _a, uint64 _ep, uint8 _i) view returns(bytes)
func (_EVerify *EVerifySession) GetCommit(_a common.Address, _ep uint64, _i uint8) ([]byte, error) {
	return _EVerify.Contract.GetCommit(&_EVerify.CallOpts, _a, _ep, _i)
}

// GetCommit is a free data retrieval call binding the contract method 0xddbf6039.
//
// Solidity: function getCommit(address _a, uint64 _ep, uint8 _i) view returns(bytes)
func (_EVerify *EVerifyCallerSession) GetCommit(_a common.Address, _ep uint64, _i uint8) ([]byte, error) {
	return _EVerify.Contract.GetCommit(&_EVerify.CallOpts, _a, _ep, _i)
}

// GetOrder is a free data retrieval call binding the contract method 0x3dcc3e1b.
//
// Solidity: function getOrder(uint64 _cnt, uint64 _pcnt) view returns(uint8)
func (_EVerify *EVerifyCaller) GetOrder(opts *bind.CallOpts, _cnt uint64, _pcnt uint64) (uint8, error) {
	var out []interface{}
	err := _EVerify.contract.Call(opts, &out, "getOrder", _cnt, _pcnt)

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// GetOrder is a free data retrieval call binding the contract method 0x3dcc3e1b.
//
// Solidity: function getOrder(uint64 _cnt, uint64 _pcnt) view returns(uint8)
func (_EVerify *EVerifySession) GetOrder(_cnt uint64, _pcnt uint64) (uint8, error) {
	return _EVerify.Contract.GetOrder(&_EVerify.CallOpts, _cnt, _pcnt)
}

// GetOrder is a free data retrieval call binding the contract method 0x3dcc3e1b.
//
// Solidity: function getOrder(uint64 _cnt, uint64 _pcnt) view returns(uint8)
func (_EVerify *EVerifyCallerSession) GetOrder(_cnt uint64, _pcnt uint64) (uint8, error) {
	return _EVerify.Contract.GetOrder(&_EVerify.CallOpts, _cnt, _pcnt)
}

// GetVKRoot is a free data retrieval call binding the contract method 0xa2265911.
//
// Solidity: function getVKRoot(string _ptype) view returns(uint256)
func (_EVerify *EVerifyCaller) GetVKRoot(opts *bind.CallOpts, _ptype string) (*big.Int, error) {
	var out []interface{}
	err := _EVerify.contract.Call(opts, &out, "getVKRoot", _ptype)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetVKRoot is a free data retrieval call binding the contract method 0xa2265911.
//
// Solidity: function getVKRoot(string _ptype) view returns(uint256)
func (_EVerify *EVerifySession) GetVKRoot(_ptype string) (*big.Int, error) {
	return _EVerify.Contract.GetVKRoot(&_EVerify.CallOpts, _ptype)
}

// GetVKRoot is a free data retrieval call binding the contract method 0xa2265911.
//
// Solidity: function getVKRoot(string _ptype) view returns(uint256)
func (_EVerify *EVerifyCallerSession) GetVKRoot(_ptype string) (*big.Int, error) {
	return _EVerify.Contract.GetVKRoot(&_EVerify.CallOpts, _ptype)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_EVerify *EVerifyCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _EVerify.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_EVerify *EVerifySession) Owner() (common.Address, error) {
	return _EVerify.Contract.Owner(&_EVerify.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_EVerify *EVerifyCallerSession) Owner() (common.Address, error) {
	return _EVerify.Contract.Owner(&_EVerify.CallOpts)
}

// ChalCom is a paid mutator transaction binding the contract method 0xb54753b8.
//
// Solidity: function chalCom(address _a, uint64 _ep, uint8 _qIndex) returns(uint8)
func (_EVerify *EVerifyTransactor) ChalCom(opts *bind.TransactOpts, _a common.Address, _ep uint64, _qIndex uint8) (*types.Transaction, error) {
	return _EVerify.contract.Transact(opts, "chalCom", _a, _ep, _qIndex)
}

// ChalCom is a paid mutator transaction binding the contract method 0xb54753b8.
//
// Solidity: function chalCom(address _a, uint64 _ep, uint8 _qIndex) returns(uint8)
func (_EVerify *EVerifySession) ChalCom(_a common.Address, _ep uint64, _qIndex uint8) (*types.Transaction, error) {
	return _EVerify.Contract.ChalCom(&_EVerify.TransactOpts, _a, _ep, _qIndex)
}

// ChalCom is a paid mutator transaction binding the contract method 0xb54753b8.
//
// Solidity: function chalCom(address _a, uint64 _ep, uint8 _qIndex) returns(uint8)
func (_EVerify *EVerifyTransactorSession) ChalCom(_a common.Address, _ep uint64, _qIndex uint8) (*types.Transaction, error) {
	return _EVerify.Contract.ChalCom(&_EVerify.TransactOpts, _a, _ep, _qIndex)
}

// ChalOne is a paid mutator transaction binding the contract method 0xbd42f4c2.
//
// Solidity: function chalOne(address _a, uint64 _ep, uint8 _qIndex) returns(uint8)
func (_EVerify *EVerifyTransactor) ChalOne(opts *bind.TransactOpts, _a common.Address, _ep uint64, _qIndex uint8) (*types.Transaction, error) {
	return _EVerify.contract.Transact(opts, "chalOne", _a, _ep, _qIndex)
}

// ChalOne is a paid mutator transaction binding the contract method 0xbd42f4c2.
//
// Solidity: function chalOne(address _a, uint64 _ep, uint8 _qIndex) returns(uint8)
func (_EVerify *EVerifySession) ChalOne(_a common.Address, _ep uint64, _qIndex uint8) (*types.Transaction, error) {
	return _EVerify.Contract.ChalOne(&_EVerify.TransactOpts, _a, _ep, _qIndex)
}

// ChalOne is a paid mutator transaction binding the contract method 0xbd42f4c2.
//
// Solidity: function chalOne(address _a, uint64 _ep, uint8 _qIndex) returns(uint8)
func (_EVerify *EVerifyTransactorSession) ChalOne(_a common.Address, _ep uint64, _qIndex uint8) (*types.Transaction, error) {
	return _EVerify.Contract.ChalOne(&_EVerify.TransactOpts, _a, _ep, _qIndex)
}

// Challenge is a paid mutator transaction binding the contract method 0x5a9b427d.
//
// Solidity: function challenge(address _a, uint64 _ep, bytes _sum) returns()
func (_EVerify *EVerifyTransactor) Challenge(opts *bind.TransactOpts, _a common.Address, _ep uint64, _sum []byte) (*types.Transaction, error) {
	return _EVerify.contract.Transact(opts, "challenge", _a, _ep, _sum)
}

// Challenge is a paid mutator transaction binding the contract method 0x5a9b427d.
//
// Solidity: function challenge(address _a, uint64 _ep, bytes _sum) returns()
func (_EVerify *EVerifySession) Challenge(_a common.Address, _ep uint64, _sum []byte) (*types.Transaction, error) {
	return _EVerify.Contract.Challenge(&_EVerify.TransactOpts, _a, _ep, _sum)
}

// Challenge is a paid mutator transaction binding the contract method 0x5a9b427d.
//
// Solidity: function challenge(address _a, uint64 _ep, bytes _sum) returns()
func (_EVerify *EVerifyTransactorSession) Challenge(_a common.Address, _ep uint64, _sum []byte) (*types.Transaction, error) {
	return _EVerify.Contract.Challenge(&_EVerify.TransactOpts, _a, _ep, _sum)
}

// ProveCom is a paid mutator transaction binding the contract method 0x9949edae.
//
// Solidity: function proveCom(address _a, uint64 _ep, bytes[] _com, bytes _proof) returns(uint8)
func (_EVerify *EVerifyTransactor) ProveCom(opts *bind.TransactOpts, _a common.Address, _ep uint64, _com [][]byte, _proof []byte) (*types.Transaction, error) {
	return _EVerify.contract.Transact(opts, "proveCom", _a, _ep, _com, _proof)
}

// ProveCom is a paid mutator transaction binding the contract method 0x9949edae.
//
// Solidity: function proveCom(address _a, uint64 _ep, bytes[] _com, bytes _proof) returns(uint8)
func (_EVerify *EVerifySession) ProveCom(_a common.Address, _ep uint64, _com [][]byte, _proof []byte) (*types.Transaction, error) {
	return _EVerify.Contract.ProveCom(&_EVerify.TransactOpts, _a, _ep, _com, _proof)
}

// ProveCom is a paid mutator transaction binding the contract method 0x9949edae.
//
// Solidity: function proveCom(address _a, uint64 _ep, bytes[] _com, bytes _proof) returns(uint8)
func (_EVerify *EVerifyTransactorSession) ProveCom(_a common.Address, _ep uint64, _com [][]byte, _proof []byte) (*types.Transaction, error) {
	return _EVerify.Contract.ProveCom(&_EVerify.TransactOpts, _a, _ep, _com, _proof)
}

// ProveKZG is a paid mutator transaction binding the contract method 0x2049af34.
//
// Solidity: function proveKZG(address _a, uint64 _ep, bytes32 _wroot, bytes _proof) returns()
func (_EVerify *EVerifyTransactor) ProveKZG(opts *bind.TransactOpts, _a common.Address, _ep uint64, _wroot [32]byte, _proof []byte) (*types.Transaction, error) {
	return _EVerify.contract.Transact(opts, "proveKZG", _a, _ep, _wroot, _proof)
}

// ProveKZG is a paid mutator transaction binding the contract method 0x2049af34.
//
// Solidity: function proveKZG(address _a, uint64 _ep, bytes32 _wroot, bytes _proof) returns()
func (_EVerify *EVerifySession) ProveKZG(_a common.Address, _ep uint64, _wroot [32]byte, _proof []byte) (*types.Transaction, error) {
	return _EVerify.Contract.ProveKZG(&_EVerify.TransactOpts, _a, _ep, _wroot, _proof)
}

// ProveKZG is a paid mutator transaction binding the contract method 0x2049af34.
//
// Solidity: function proveKZG(address _a, uint64 _ep, bytes32 _wroot, bytes _proof) returns()
func (_EVerify *EVerifyTransactorSession) ProveKZG(_a common.Address, _ep uint64, _wroot [32]byte, _proof []byte) (*types.Transaction, error) {
	return _EVerify.Contract.ProveKZG(&_EVerify.TransactOpts, _a, _ep, _wroot, _proof)
}

// ProveOne is a paid mutator transaction binding the contract method 0x9b00f959.
//
// Solidity: function proveOne(address _a, uint64 _ep, bytes _com, bytes _proof) returns(uint8)
func (_EVerify *EVerifyTransactor) ProveOne(opts *bind.TransactOpts, _a common.Address, _ep uint64, _com []byte, _proof []byte) (*types.Transaction, error) {
	return _EVerify.contract.Transact(opts, "proveOne", _a, _ep, _com, _proof)
}

// ProveOne is a paid mutator transaction binding the contract method 0x9b00f959.
//
// Solidity: function proveOne(address _a, uint64 _ep, bytes _com, bytes _proof) returns(uint8)
func (_EVerify *EVerifySession) ProveOne(_a common.Address, _ep uint64, _com []byte, _proof []byte) (*types.Transaction, error) {
	return _EVerify.Contract.ProveOne(&_EVerify.TransactOpts, _a, _ep, _com, _proof)
}

// ProveOne is a paid mutator transaction binding the contract method 0x9b00f959.
//
// Solidity: function proveOne(address _a, uint64 _ep, bytes _com, bytes _proof) returns(uint8)
func (_EVerify *EVerifyTransactorSession) ProveOne(_a common.Address, _ep uint64, _com []byte, _proof []byte) (*types.Transaction, error) {
	return _EVerify.Contract.ProveOne(&_EVerify.TransactOpts, _a, _ep, _com, _proof)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_EVerify *EVerifyTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EVerify.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_EVerify *EVerifySession) RenounceOwnership() (*types.Transaction, error) {
	return _EVerify.Contract.RenounceOwnership(&_EVerify.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_EVerify *EVerifyTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _EVerify.Contract.RenounceOwnership(&_EVerify.TransactOpts)
}

// SetVeirfier is a paid mutator transaction binding the contract method 0x141df27c.
//
// Solidity: function setVeirfier(string _ptype, uint256 _vkroot) returns()
func (_EVerify *EVerifyTransactor) SetVeirfier(opts *bind.TransactOpts, _ptype string, _vkroot *big.Int) (*types.Transaction, error) {
	return _EVerify.contract.Transact(opts, "setVeirfier", _ptype, _vkroot)
}

// SetVeirfier is a paid mutator transaction binding the contract method 0x141df27c.
//
// Solidity: function setVeirfier(string _ptype, uint256 _vkroot) returns()
func (_EVerify *EVerifySession) SetVeirfier(_ptype string, _vkroot *big.Int) (*types.Transaction, error) {
	return _EVerify.Contract.SetVeirfier(&_EVerify.TransactOpts, _ptype, _vkroot)
}

// SetVeirfier is a paid mutator transaction binding the contract method 0x141df27c.
//
// Solidity: function setVeirfier(string _ptype, uint256 _vkroot) returns()
func (_EVerify *EVerifyTransactorSession) SetVeirfier(_ptype string, _vkroot *big.Int) (*types.Transaction, error) {
	return _EVerify.Contract.SetVeirfier(&_EVerify.TransactOpts, _ptype, _vkroot)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_EVerify *EVerifyTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _EVerify.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_EVerify *EVerifySession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _EVerify.Contract.TransferOwnership(&_EVerify.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_EVerify *EVerifyTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _EVerify.Contract.TransferOwnership(&_EVerify.TransactOpts, newOwner)
}

// EVerifyOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the EVerify contract.
type EVerifyOwnershipTransferredIterator struct {
	Event *EVerifyOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *EVerifyOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EVerifyOwnershipTransferred)
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
		it.Event = new(EVerifyOwnershipTransferred)
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
func (it *EVerifyOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EVerifyOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EVerifyOwnershipTransferred represents a OwnershipTransferred event raised by the EVerify contract.
type EVerifyOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_EVerify *EVerifyFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*EVerifyOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _EVerify.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &EVerifyOwnershipTransferredIterator{contract: _EVerify.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_EVerify *EVerifyFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *EVerifyOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _EVerify.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EVerifyOwnershipTransferred)
				if err := _EVerify.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_EVerify *EVerifyFilterer) ParseOwnershipTransferred(log types.Log) (*EVerifyOwnershipTransferred, error) {
	event := new(EVerifyOwnershipTransferred)
	if err := _EVerify.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
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

// IEVerifyMetaData contains all meta data concerning the IEVerify contract.
var IEVerifyMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_a\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"_ep\",\"type\":\"uint64\"},{\"internalType\":\"uint8\",\"name\":\"_qIndex\",\"type\":\"uint8\"}],\"name\":\"chalCom\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_a\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"_ep\",\"type\":\"uint64\"},{\"internalType\":\"uint8\",\"name\":\"_qIndex\",\"type\":\"uint8\"}],\"name\":\"chalOne\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_a\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"_ep\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"_com\",\"type\":\"bytes\"}],\"name\":\"challenge\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_a\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"_ep\",\"type\":\"uint64\"},{\"internalType\":\"bytes[]\",\"name\":\"_com\",\"type\":\"bytes[]\"},{\"internalType\":\"bytes\",\"name\":\"_proof\",\"type\":\"bytes\"}],\"name\":\"proveCom\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_a\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"_ep\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"_wroot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"_proof\",\"type\":\"bytes\"}],\"name\":\"proveKZG\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_a\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"_ep\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"_com\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"_proof\",\"type\":\"bytes\"}],\"name\":\"proveOne\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// IEVerifyABI is the input ABI used to generate the binding from.
// Deprecated: Use IEVerifyMetaData.ABI instead.
var IEVerifyABI = IEVerifyMetaData.ABI

// IEVerify is an auto generated Go binding around an Ethereum contract.
type IEVerify struct {
	IEVerifyCaller     // Read-only binding to the contract
	IEVerifyTransactor // Write-only binding to the contract
	IEVerifyFilterer   // Log filterer for contract events
}

// IEVerifyCaller is an auto generated read-only Go binding around an Ethereum contract.
type IEVerifyCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IEVerifyTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IEVerifyTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IEVerifyFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IEVerifyFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IEVerifySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IEVerifySession struct {
	Contract     *IEVerify         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IEVerifyCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IEVerifyCallerSession struct {
	Contract *IEVerifyCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// IEVerifyTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IEVerifyTransactorSession struct {
	Contract     *IEVerifyTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// IEVerifyRaw is an auto generated low-level Go binding around an Ethereum contract.
type IEVerifyRaw struct {
	Contract *IEVerify // Generic contract binding to access the raw methods on
}

// IEVerifyCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IEVerifyCallerRaw struct {
	Contract *IEVerifyCaller // Generic read-only contract binding to access the raw methods on
}

// IEVerifyTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IEVerifyTransactorRaw struct {
	Contract *IEVerifyTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIEVerify creates a new instance of IEVerify, bound to a specific deployed contract.
func NewIEVerify(address common.Address, backend bind.ContractBackend) (*IEVerify, error) {
	contract, err := bindIEVerify(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IEVerify{IEVerifyCaller: IEVerifyCaller{contract: contract}, IEVerifyTransactor: IEVerifyTransactor{contract: contract}, IEVerifyFilterer: IEVerifyFilterer{contract: contract}}, nil
}

// NewIEVerifyCaller creates a new read-only instance of IEVerify, bound to a specific deployed contract.
func NewIEVerifyCaller(address common.Address, caller bind.ContractCaller) (*IEVerifyCaller, error) {
	contract, err := bindIEVerify(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IEVerifyCaller{contract: contract}, nil
}

// NewIEVerifyTransactor creates a new write-only instance of IEVerify, bound to a specific deployed contract.
func NewIEVerifyTransactor(address common.Address, transactor bind.ContractTransactor) (*IEVerifyTransactor, error) {
	contract, err := bindIEVerify(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IEVerifyTransactor{contract: contract}, nil
}

// NewIEVerifyFilterer creates a new log filterer instance of IEVerify, bound to a specific deployed contract.
func NewIEVerifyFilterer(address common.Address, filterer bind.ContractFilterer) (*IEVerifyFilterer, error) {
	contract, err := bindIEVerify(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IEVerifyFilterer{contract: contract}, nil
}

// bindIEVerify binds a generic wrapper to an already deployed contract.
func bindIEVerify(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := IEVerifyMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IEVerify *IEVerifyRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IEVerify.Contract.IEVerifyCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IEVerify *IEVerifyRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IEVerify.Contract.IEVerifyTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IEVerify *IEVerifyRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IEVerify.Contract.IEVerifyTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IEVerify *IEVerifyCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IEVerify.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IEVerify *IEVerifyTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IEVerify.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IEVerify *IEVerifyTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IEVerify.Contract.contract.Transact(opts, method, params...)
}

// ChalCom is a paid mutator transaction binding the contract method 0xb54753b8.
//
// Solidity: function chalCom(address _a, uint64 _ep, uint8 _qIndex) returns(uint8)
func (_IEVerify *IEVerifyTransactor) ChalCom(opts *bind.TransactOpts, _a common.Address, _ep uint64, _qIndex uint8) (*types.Transaction, error) {
	return _IEVerify.contract.Transact(opts, "chalCom", _a, _ep, _qIndex)
}

// ChalCom is a paid mutator transaction binding the contract method 0xb54753b8.
//
// Solidity: function chalCom(address _a, uint64 _ep, uint8 _qIndex) returns(uint8)
func (_IEVerify *IEVerifySession) ChalCom(_a common.Address, _ep uint64, _qIndex uint8) (*types.Transaction, error) {
	return _IEVerify.Contract.ChalCom(&_IEVerify.TransactOpts, _a, _ep, _qIndex)
}

// ChalCom is a paid mutator transaction binding the contract method 0xb54753b8.
//
// Solidity: function chalCom(address _a, uint64 _ep, uint8 _qIndex) returns(uint8)
func (_IEVerify *IEVerifyTransactorSession) ChalCom(_a common.Address, _ep uint64, _qIndex uint8) (*types.Transaction, error) {
	return _IEVerify.Contract.ChalCom(&_IEVerify.TransactOpts, _a, _ep, _qIndex)
}

// ChalOne is a paid mutator transaction binding the contract method 0xbd42f4c2.
//
// Solidity: function chalOne(address _a, uint64 _ep, uint8 _qIndex) returns(uint8)
func (_IEVerify *IEVerifyTransactor) ChalOne(opts *bind.TransactOpts, _a common.Address, _ep uint64, _qIndex uint8) (*types.Transaction, error) {
	return _IEVerify.contract.Transact(opts, "chalOne", _a, _ep, _qIndex)
}

// ChalOne is a paid mutator transaction binding the contract method 0xbd42f4c2.
//
// Solidity: function chalOne(address _a, uint64 _ep, uint8 _qIndex) returns(uint8)
func (_IEVerify *IEVerifySession) ChalOne(_a common.Address, _ep uint64, _qIndex uint8) (*types.Transaction, error) {
	return _IEVerify.Contract.ChalOne(&_IEVerify.TransactOpts, _a, _ep, _qIndex)
}

// ChalOne is a paid mutator transaction binding the contract method 0xbd42f4c2.
//
// Solidity: function chalOne(address _a, uint64 _ep, uint8 _qIndex) returns(uint8)
func (_IEVerify *IEVerifyTransactorSession) ChalOne(_a common.Address, _ep uint64, _qIndex uint8) (*types.Transaction, error) {
	return _IEVerify.Contract.ChalOne(&_IEVerify.TransactOpts, _a, _ep, _qIndex)
}

// Challenge is a paid mutator transaction binding the contract method 0x5a9b427d.
//
// Solidity: function challenge(address _a, uint64 _ep, bytes _com) returns()
func (_IEVerify *IEVerifyTransactor) Challenge(opts *bind.TransactOpts, _a common.Address, _ep uint64, _com []byte) (*types.Transaction, error) {
	return _IEVerify.contract.Transact(opts, "challenge", _a, _ep, _com)
}

// Challenge is a paid mutator transaction binding the contract method 0x5a9b427d.
//
// Solidity: function challenge(address _a, uint64 _ep, bytes _com) returns()
func (_IEVerify *IEVerifySession) Challenge(_a common.Address, _ep uint64, _com []byte) (*types.Transaction, error) {
	return _IEVerify.Contract.Challenge(&_IEVerify.TransactOpts, _a, _ep, _com)
}

// Challenge is a paid mutator transaction binding the contract method 0x5a9b427d.
//
// Solidity: function challenge(address _a, uint64 _ep, bytes _com) returns()
func (_IEVerify *IEVerifyTransactorSession) Challenge(_a common.Address, _ep uint64, _com []byte) (*types.Transaction, error) {
	return _IEVerify.Contract.Challenge(&_IEVerify.TransactOpts, _a, _ep, _com)
}

// ProveCom is a paid mutator transaction binding the contract method 0x9949edae.
//
// Solidity: function proveCom(address _a, uint64 _ep, bytes[] _com, bytes _proof) returns(uint8)
func (_IEVerify *IEVerifyTransactor) ProveCom(opts *bind.TransactOpts, _a common.Address, _ep uint64, _com [][]byte, _proof []byte) (*types.Transaction, error) {
	return _IEVerify.contract.Transact(opts, "proveCom", _a, _ep, _com, _proof)
}

// ProveCom is a paid mutator transaction binding the contract method 0x9949edae.
//
// Solidity: function proveCom(address _a, uint64 _ep, bytes[] _com, bytes _proof) returns(uint8)
func (_IEVerify *IEVerifySession) ProveCom(_a common.Address, _ep uint64, _com [][]byte, _proof []byte) (*types.Transaction, error) {
	return _IEVerify.Contract.ProveCom(&_IEVerify.TransactOpts, _a, _ep, _com, _proof)
}

// ProveCom is a paid mutator transaction binding the contract method 0x9949edae.
//
// Solidity: function proveCom(address _a, uint64 _ep, bytes[] _com, bytes _proof) returns(uint8)
func (_IEVerify *IEVerifyTransactorSession) ProveCom(_a common.Address, _ep uint64, _com [][]byte, _proof []byte) (*types.Transaction, error) {
	return _IEVerify.Contract.ProveCom(&_IEVerify.TransactOpts, _a, _ep, _com, _proof)
}

// ProveKZG is a paid mutator transaction binding the contract method 0x2049af34.
//
// Solidity: function proveKZG(address _a, uint64 _ep, bytes32 _wroot, bytes _proof) returns()
func (_IEVerify *IEVerifyTransactor) ProveKZG(opts *bind.TransactOpts, _a common.Address, _ep uint64, _wroot [32]byte, _proof []byte) (*types.Transaction, error) {
	return _IEVerify.contract.Transact(opts, "proveKZG", _a, _ep, _wroot, _proof)
}

// ProveKZG is a paid mutator transaction binding the contract method 0x2049af34.
//
// Solidity: function proveKZG(address _a, uint64 _ep, bytes32 _wroot, bytes _proof) returns()
func (_IEVerify *IEVerifySession) ProveKZG(_a common.Address, _ep uint64, _wroot [32]byte, _proof []byte) (*types.Transaction, error) {
	return _IEVerify.Contract.ProveKZG(&_IEVerify.TransactOpts, _a, _ep, _wroot, _proof)
}

// ProveKZG is a paid mutator transaction binding the contract method 0x2049af34.
//
// Solidity: function proveKZG(address _a, uint64 _ep, bytes32 _wroot, bytes _proof) returns()
func (_IEVerify *IEVerifyTransactorSession) ProveKZG(_a common.Address, _ep uint64, _wroot [32]byte, _proof []byte) (*types.Transaction, error) {
	return _IEVerify.Contract.ProveKZG(&_IEVerify.TransactOpts, _a, _ep, _wroot, _proof)
}

// ProveOne is a paid mutator transaction binding the contract method 0x9b00f959.
//
// Solidity: function proveOne(address _a, uint64 _ep, bytes _com, bytes _proof) returns(uint8)
func (_IEVerify *IEVerifyTransactor) ProveOne(opts *bind.TransactOpts, _a common.Address, _ep uint64, _com []byte, _proof []byte) (*types.Transaction, error) {
	return _IEVerify.contract.Transact(opts, "proveOne", _a, _ep, _com, _proof)
}

// ProveOne is a paid mutator transaction binding the contract method 0x9b00f959.
//
// Solidity: function proveOne(address _a, uint64 _ep, bytes _com, bytes _proof) returns(uint8)
func (_IEVerify *IEVerifySession) ProveOne(_a common.Address, _ep uint64, _com []byte, _proof []byte) (*types.Transaction, error) {
	return _IEVerify.Contract.ProveOne(&_IEVerify.TransactOpts, _a, _ep, _com, _proof)
}

// ProveOne is a paid mutator transaction binding the contract method 0x9b00f959.
//
// Solidity: function proveOne(address _a, uint64 _ep, bytes _com, bytes _proof) returns(uint8)
func (_IEVerify *IEVerifyTransactorSession) ProveOne(_a common.Address, _ep uint64, _com []byte, _proof []byte) (*types.Transaction, error) {
	return _IEVerify.Contract.ProveOne(&_IEVerify.TransactOpts, _a, _ep, _com, _proof)
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
