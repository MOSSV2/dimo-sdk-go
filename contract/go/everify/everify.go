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

// EVerifyVInfo is an auto generated low-level Go binding around an user-defined struct.
type EVerifyVInfo struct {
	Vkroot   *big.Int
	Verifier common.Address
}

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
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_b\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"KZG_G1\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"KZG_VK\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"bank\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_a\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"_ep\",\"type\":\"uint64\"},{\"internalType\":\"uint8\",\"name\":\"_qIndex\",\"type\":\"uint8\"}],\"name\":\"chalCom\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_a\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"_ep\",\"type\":\"uint64\"},{\"internalType\":\"uint8\",\"name\":\"_qIndex\",\"type\":\"uint8\"}],\"name\":\"chalOne\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_a\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"_ep\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"_sum\",\"type\":\"bytes\"}],\"name\":\"challenge\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_a\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"_seed\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"_count\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"_pcount\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"_index\",\"type\":\"uint64\"}],\"name\":\"choose\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"divide\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_a\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"_ep\",\"type\":\"uint64\"}],\"name\":\"getCInfo\",\"outputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"preCnt\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"count\",\"type\":\"uint64\"},{\"internalType\":\"uint8\",\"name\":\"round\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"roundAt\",\"type\":\"uint8\"},{\"internalType\":\"uint64\",\"name\":\"cIndex\",\"type\":\"uint64\"},{\"internalType\":\"uint8\",\"name\":\"qIndex\",\"type\":\"uint8\"}],\"internalType\":\"structIEVerify.CInfo\",\"name\":\"_ci\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"_cnt\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"_pcnt\",\"type\":\"uint64\"}],\"name\":\"getOrder\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"_ptype\",\"type\":\"uint8\"}],\"name\":\"getVerifier\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"vkroot\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"verifier\",\"type\":\"address\"}],\"internalType\":\"structEVerify.VInfo\",\"name\":\"_vi\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_a\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"_ep\",\"type\":\"uint64\"},{\"internalType\":\"bytes[]\",\"name\":\"_com\",\"type\":\"bytes[]\"},{\"internalType\":\"bytes\",\"name\":\"_proof\",\"type\":\"bytes\"}],\"name\":\"proveCom\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_a\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"_ep\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"_wroot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"_proof\",\"type\":\"bytes\"}],\"name\":\"proveKZG\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_a\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"_ep\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"_com\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"_proof\",\"type\":\"bytes\"}],\"name\":\"proveOne\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"_ptype\",\"type\":\"uint8\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"vkroot\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"verifier\",\"type\":\"address\"}],\"internalType\":\"structEVerify.VInfo\",\"name\":\"_vi\",\"type\":\"tuple\"}],\"name\":\"setVeirfier\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405260405180608001604052806060815260200161551b606091396002908161002b9190610400565b5060405180606001604052806030815260200161557b60309139600390816100539190610400565b50600860045f6101000a81548160ff021916908360ff16021790555034801561007a575f80fd5b506040516155ab3803806155ab833981810160405281019061009c919061052d565b6100b86100ad6100fe60201b60201c565b61010560201b60201c565b8060015f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050610558565b5f33905090565b5f805f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050815f806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508173ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a35050565b5f81519050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b7f4e487b71000000000000000000000000000000000000000000000000000000005f52602260045260245ffd5b5f600282049050600182168061024157607f821691505b602082108103610254576102536101fd565b5b50919050565b5f819050815f5260205f209050919050565b5f6020601f8301049050919050565b5f82821b905092915050565b5f600883026102b67fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8261027b565b6102c0868361027b565b95508019841693508086168417925050509392505050565b5f819050919050565b5f819050919050565b5f6103046102ff6102fa846102d8565b6102e1565b6102d8565b9050919050565b5f819050919050565b61031d836102ea565b6103316103298261030b565b848454610287565b825550505050565b5f90565b610345610339565b610350818484610314565b505050565b5b81811015610373576103685f8261033d565b600181019050610356565b5050565b601f8211156103b8576103898161025a565b6103928461026c565b810160208510156103a1578190505b6103b56103ad8561026c565b830182610355565b50505b505050565b5f82821c905092915050565b5f6103d85f19846008026103bd565b1980831691505092915050565b5f6103f083836103c9565b9150826002028217905092915050565b610409826101c6565b67ffffffffffffffff811115610422576104216101d0565b5b61042c825461022a565b610437828285610377565b5f60209050601f831160018114610468575f8415610456578287015190505b61046085826103e5565b8655506104c7565b601f1984166104768661025a565b5f5b8281101561049d57848901518255600182019150602085019450602081019050610478565b868310156104ba57848901516104b6601f8916826103c9565b8355505b6001600288020188555050505b505050505050565b5f80fd5b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f6104fc826104d3565b9050919050565b61050c816104f2565b8114610516575f80fd5b50565b5f8151905061052781610503565b92915050565b5f60208284031215610542576105416104cf565b5b5f61054f84828501610519565b91505092915050565b614fb6806105655f395ff3fe608060405234801561000f575f80fd5b5060043610610114575f3560e01c80639b00f959116100a0578063b54753b81161006f578063b54753b814610306578063b8b55ddd14610336578063bd42f4c214610354578063be12c9ea14610384578063f2fde38b146103a057610114565b80639b00f95914610258578063adaad46714610288578063b1ea66de146102b8578063b53b8d7a146102e857610114565b80635a9b427d116100e75780635a9b427d146101b2578063715018a6146101e257806376cdb03b146101ec5780638da5cb5b1461020a5780639949edae1461022857610114565b80632049af341461011857806323178d22146101345780633dcc3e1b146101525780634c07713214610182575b5f80fd5b610132600480360381019061012d919061374c565b6103bc565b005b61013c6107dd565b6040516101499190613846565b60405180910390f35b61016c60048036038101906101679190613866565b610869565b60405161017991906138bf565b60405180910390f35b61019c600480360381019061019791906138d8565b61087c565b6040516101a9919061395e565b60405180910390f35b6101cc60048036038101906101c79190613977565b610895565b6040516101d991906138bf565b60405180910390f35b6101ea610de4565b005b6101f4610df7565b60405161020191906139f2565b60405180910390f35b610212610e1c565b60405161021f91906139f2565b60405180910390f35b610242600480360381019061023d9190613aed565b610e43565b60405161024f91906138bf565b60405180910390f35b610272600480360381019061026d9190613b89565b611481565b60405161027f91906138bf565b60405180910390f35b6102a2600480360381019061029d9190613c4f565b611ef4565b6040516102af9190613cce565b60405180910390f35b6102d260048036038101906102cd9190613ce7565b611f94565b6040516102df9190613dbc565b60405180910390f35b6102f06122e5565b6040516102fd91906138bf565b60405180910390f35b610320600480360381019061031b9190613dd5565b6122f7565b60405161032d91906138bf565b60405180910390f35b61033e6127dc565b60405161034b9190613846565b60405180910390f35b61036e60048036038101906103699190613dd5565b612868565b60405161037b91906138bf565b60405180910390f35b61039e60048036038101906103999190613ea0565b612d4d565b005b6103ba60048036038101906103b59190613ede565b612f11565b005b5f60075f8060ff1681526020019081526020015f206040518060400160405290815f8201548152602001600182015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152505090505f73ffffffffffffffffffffffffffffffffffffffff16816020015173ffffffffffffffffffffffffffffffffffffffff16036104ae576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016104a590613f63565b60405180910390fd5b5f60015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663693ec85e6040518163ffffffff1660e01b815260040161050790613fcb565b6020604051808303815f875af1158015610523573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906105479190613ffd565b73ffffffffffffffffffffffffffffffffffffffff166312a02c8260018761056f9190614055565b6040518263ffffffff1660e01b815260040161058b919061395e565b60408051808303815f875af11580156105a6573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906105ca91906140b8565b9150505f81876040516020016105e192919061415b565b604051602081830303815290604052805190602001205f1c90507f12ab655e9a2ca55660b44d1e5c37b00159aa76fed00000010a11800000000001810690505f8561062b83612f93565b600360405160200161063f939291906142af565b604051602081830303815290604052805190602001205f1c90507f30644e72e131a029b85045b68181585d2833e84879b9709143e1f593f0000001810690505f600267ffffffffffffffff81111561069a57610699613628565b5b6040519080825280602002602001820160405280156106c85781602001602082028036833780820191505090505b509050845f0151815f815181106106e2576106e16142e3565b5b6020026020010181815250508181600181518110610703576107026142e3565b5b6020026020010181815250505f856020015173ffffffffffffffffffffffffffffffffffffffff16637e4f7a8a88846040518363ffffffff1660e01b815260040161074f9291906143b8565b6020604051808303815f875af115801561076b573d5f803e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061078f9190614422565b9050806107d1576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016107c890614497565b60405180910390fd5b50505050505050505050565b600380546107ea906141ed565b80601f0160208091040260200160405190810160405280929190818152602001828054610816906141ed565b80156108615780601f1061083857610100808354040283529160200191610861565b820191905f5260205f20905b81548152906001019060200180831161084457829003601f168201915b505050505081565b5f6108748383613069565b905092915050565b5f61088a86868686866131c4565b905095945050505050565b5f61089e613263565b5f60015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663693ec85e6040518163ffffffff1660e01b81526004016108f7906144ff565b6020604051808303815f875af1158015610913573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906109379190613ffd565b90505f8173ffffffffffffffffffffffffffffffffffffffff166350979d7e87876040518363ffffffff1660e01b815260040161097592919061451d565b6020604051808303815f875af1158015610991573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906109b59190614558565b90505f8273ffffffffffffffffffffffffffffffffffffffff166350979d7e886001896109e29190614583565b6040518363ffffffff1660e01b81526004016109ff92919061451d565b6020604051808303815f875af1158015610a1b573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610a3f9190614558565b9050610a4b8282613069565b60055f8973ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f8867ffffffffffffffff1667ffffffffffffffff1681526020019081526020015f205f0160106101000a81548160ff021916908360ff1602179055508060055f8973ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f8867ffffffffffffffff1667ffffffffffffffff1681526020019081526020015f205f015f6101000a81548167ffffffffffffffff021916908367ffffffffffffffff1602179055508160055f8973ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f8867ffffffffffffffff1667ffffffffffffffff1681526020019081526020015f205f0160086101000a81548167ffffffffffffffff021916908367ffffffffffffffff160217905550600160055f8973ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f8867ffffffffffffffff1667ffffffffffffffff1681526020019081526020015f205f0160116101000a81548160ff021916908360ff1602179055505f60055f8973ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f8867ffffffffffffffff1667ffffffffffffffff1681526020019081526020015f205f0160126101000a81548167ffffffffffffffff021916908367ffffffffffffffff1602179055505f60055f8973ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f8867ffffffffffffffff1667ffffffffffffffff1681526020019081526020015f205f01601a6101000a81548160ff021916908360ff1602179055508460065f8973ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f8867ffffffffffffffff1667ffffffffffffffff1681526020019081526020015f205f8060ff1681526020019081526020015f209081610dd59190614749565b50600193505050509392505050565b610dec61336a565b610df55f6133e8565b565b60015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b5f805f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b5f610e4c613263565b60045f9054906101000a900460ff1660ff16835114610ea0576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610e9790614862565b60405180910390fd5b5f60065f8773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f8667ffffffffffffffff1667ffffffffffffffff1681526020019081526020015f205f60055f8973ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f8867ffffffffffffffff1667ffffffffffffffff1681526020019081526020015f205f01601a9054906101000a900460ff1660ff1660ff1681526020019081526020015f208054610f90906141ed565b80601f0160208091040260200160405190810160405280929190818152602001828054610fbc906141ed565b80156110075780601f10610fde57610100808354040283529160200191611007565b820191905f5260205f20905b815481529060010190602001808311610fea57829003601f168201915b505050505090505f815111611051576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401611048906148ca565b60405180910390fd5b5f5b60045f9054906101000a900460ff1660ff168160ff1610156110bd5781858260ff1681518110611086576110856142e3565b5b602002602001015160405160200161109f9291906148e8565b60405160208183030381529060405291508080600101915050611053565b505f81805190602001205f1c90507f30644e72e131a029b85045b68181585d2833e84879b9709143e1f593f0000001810690505f60075f600260ff1681526020019081526020015f206040518060400160405290815f8201548152602001600182015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152505090505f73ffffffffffffffffffffffffffffffffffffffff16816020015173ffffffffffffffffffffffffffffffffffffffff16036111e3576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016111da90613f63565b60405180910390fd5b5f600267ffffffffffffffff8111156111ff576111fe613628565b5b60405190808252806020026020018201604052801561122d5781602001602082028036833780820191505090505b509050815f0151815f81518110611247576112466142e3565b5b6020026020010181815250508281600181518110611268576112676142e3565b5b6020026020010181815250505f826020015173ffffffffffffffffffffffffffffffffffffffff16637e4f7a8a88846040518363ffffffff1660e01b81526004016112b49291906143b8565b6020604051808303815f875af11580156112d0573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906112f49190614422565b905080611336576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161132d90614497565b60405180910390fd5b5f5b60045f9054906101000a900460ff1660ff168160ff16101561140157888160ff168151811061136a576113696142e3565b5b602002602001015160065f8d73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f8c67ffffffffffffffff1667ffffffffffffffff1681526020019081526020015f205f8360ff1660ff1681526020019081526020015f2090816113f39190614749565b508080600101915050611338565b5060055f8b73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f8a67ffffffffffffffff1667ffffffffffffffff1681526020019081526020015f205f0160119054906101000a900460ff1695505050505050949350505050565b5f8060055f8773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f8667ffffffffffffffff1667ffffffffffffffff1681526020019081526020015f205f0160119054906101000a900460ff169050600160055f8873ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f8767ffffffffffffffff1667ffffffffffffffff1681526020019081526020015f205f0160109054906101000a900460ff16611571919061490b565b60ff168160ff16146115b8576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016115af90614989565b60405180910390fd5b5f60015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663693ec85e6040518163ffffffff1660e01b815260040161161190613fcb565b6020604051808303815f875af115801561162d573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906116519190613ffd565b73ffffffffffffffffffffffffffffffffffffffff166312a02c826001886116799190614055565b6040518263ffffffff1660e01b8152600401611695919061395e565b60408051808303815f875af11580156116b0573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906116d491906140b8565b9150505f611846888360055f8c73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f8b67ffffffffffffffff1667ffffffffffffffff1681526020019081526020015f205f0160089054906101000a900467ffffffffffffffff1660055f8d73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f8c67ffffffffffffffff1667ffffffffffffffff1681526020019081526020015f205f015f9054906101000a900467ffffffffffffffff1660055f8e73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f8d67ffffffffffffffff1667ffffffffffffffff1681526020019081526020015f205f0160129054906101000a900467ffffffffffffffff166131c4565b90505f60015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663693ec85e6040518163ffffffff1660e01b81526004016118a1906144ff565b6020604051808303815f875af11580156118bd573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906118e19190613ffd565b73ffffffffffffffffffffffffffffffffffffffff1663894d35f88a848a6040518463ffffffff1660e01b815260040161191d939291906149a7565b602060405180830381865afa158015611938573d5f803e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061195c9190614558565b90508767ffffffffffffffff168167ffffffffffffffff1611611a055760028054611986906141ed565b80601f01602080910402602001604051908101604052809291908181526020018280546119b2906141ed565b80156119fd5780601f106119d4576101008083540402835291602001916119fd565b820191905f5260205f20905b8154815290600101906020018083116119e057829003601f168201915b505050505096505b5f60065f8b73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f8a67ffffffffffffffff1667ffffffffffffffff1681526020019081526020015f205f60055f8d73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f8c67ffffffffffffffff1667ffffffffffffffff1681526020019081526020015f205f01601a9054906101000a900460ff1660ff1660ff1681526020019081526020015f208054611af5906141ed565b80601f0160208091040260200160405190810160405280929190818152602001828054611b21906141ed565b8015611b6c5780601f10611b4357610100808354040283529160200191611b6c565b820191905f5260205f20905b815481529060010190602001808311611b4f57829003601f168201915b505050505090505f8460055f8d73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f8c67ffffffffffffffff1667ffffffffffffffff1681526020019081526020015f205f0160129054906101000a900467ffffffffffffffff16604051602001611bfd929190614a17565b604051602081830303815290604052805190602001205f1c90507f12ab655e9a2ca55660b44d1e5c37b00159aa76fed00000010a11800000000001810690505f828a611c4884612f93565b604051602001611c5a93929190614a42565b604051602081830303815290604052805190602001205f1c90507f30644e72e131a029b85045b68181585d2833e84879b9709143e1f593f0000001810690505f60075f600160ff1681526020019081526020015f206040518060400160405290815f8201548152602001600182015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152505090505f73ffffffffffffffffffffffffffffffffffffffff16816020015173ffffffffffffffffffffffffffffffffffffffff1603611d8c576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401611d8390613f63565b60405180910390fd5b5f600267ffffffffffffffff811115611da857611da7613628565b5b604051908082528060200260200182016040528015611dd65781602001602082028036833780820191505090505b509050815f0151815f81518110611df057611def6142e3565b5b6020026020010181815250508281600181518110611e1157611e106142e3565b5b6020026020010181815250505f826020015173ffffffffffffffffffffffffffffffffffffffff16637e4f7a8a8d846040518363ffffffff1660e01b8152600401611e5d9291906143b8565b6020604051808303815f875af1158015611e79573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190611e9d9190614422565b905080611edf576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401611ed690614497565b60405180910390fd5b899a5050505050505050505050949350505050565b611efc6134b0565b60075f8360ff1660ff1681526020019081526020015f205f0154815f01818152505060075f8360ff1660ff1681526020019081526020015f206001015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff16816020019073ffffffffffffffffffffffffffffffffffffffff16908173ffffffffffffffffffffffffffffffffffffffff1681525050919050565b611f9c6134de565b60055f8473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f8367ffffffffffffffff1667ffffffffffffffff1681526020019081526020015f205f015f9054906101000a900467ffffffffffffffff16815f019067ffffffffffffffff16908167ffffffffffffffff168152505060055f8473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f8367ffffffffffffffff1667ffffffffffffffff1681526020019081526020015f205f0160089054906101000a900467ffffffffffffffff16816020019067ffffffffffffffff16908167ffffffffffffffff168152505060055f8473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f8367ffffffffffffffff1667ffffffffffffffff1681526020019081526020015f205f0160109054906101000a900460ff16816040019060ff16908160ff168152505060055f8473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f8367ffffffffffffffff1667ffffffffffffffff1681526020019081526020015f205f0160119054906101000a900460ff16816060019060ff16908160ff168152505060055f8473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f8367ffffffffffffffff1667ffffffffffffffff1681526020019081526020015f205f01601a9054906101000a900460ff168160a0019060ff16908160ff168152505060055f8473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f8367ffffffffffffffff1667ffffffffffffffff1681526020019081526020015f205f0160129054906101000a900467ffffffffffffffff16816080019067ffffffffffffffff16908167ffffffffffffffff168152505092915050565b60045f9054906101000a900460ff1681565b5f612300613263565b60045f9054906101000a900460ff1660ff168260ff1610612356576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161234d90614abc565b60405180910390fd5b5f60065f8673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f8567ffffffffffffffff1667ffffffffffffffff1681526020019081526020015f205f8460ff1660ff1681526020019081526020015f2080546123d7906141ed565b905011612419576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401612410906148ca565b60405180910390fd5b60055f8573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f8467ffffffffffffffff1667ffffffffffffffff1681526020019081526020015f205f0160109054906101000a900460ff1660ff1660055f8673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f8567ffffffffffffffff1667ffffffffffffffff1681526020019081526020015f205f0160119054906101000a900460ff1660ff161061253f576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161253690614b24565b60405180910390fd5b8160055f8673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f8567ffffffffffffffff1667ffffffffffffffff1681526020019081526020015f205f01601a6101000a81548160ff021916908360ff160217905550600160055f8673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f8567ffffffffffffffff1667ffffffffffffffff1681526020019081526020015f205f0160118282829054906101000a900460ff16612639919061490b565b92506101000a81548160ff021916908360ff16021790555081600860055f8773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f8667ffffffffffffffff1667ffffffffffffffff1681526020019081526020015f205f01601a9054906101000a900460ff166126ce9190614b42565b6126d8919061490b565b60ff1660055f8673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f8567ffffffffffffffff1667ffffffffffffffff1681526020019081526020015f205f0160126101000a81548167ffffffffffffffff021916908367ffffffffffffffff16021790555060055f8573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f8467ffffffffffffffff1667ffffffffffffffff1681526020019081526020015f205f0160119054906101000a900460ff1690509392505050565b600280546127e9906141ed565b80601f0160208091040260200160405190810160405280929190818152602001828054612815906141ed565b80156128605780601f1061283757610100808354040283529160200191612860565b820191905f5260205f20905b81548152906001019060200180831161284357829003601f168201915b505050505081565b5f612871613263565b60055f8573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f8467ffffffffffffffff1667ffffffffffffffff1681526020019081526020015f205f0160109054906101000a900460ff1660ff1660055f8673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f8567ffffffffffffffff1667ffffffffffffffff1681526020019081526020015f205f0160119054906101000a900460ff1660ff1614612997576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161298e90614989565b60405180910390fd5b60045f9054906101000a900460ff1660ff168260ff16106129ed576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016129e490614abc565b60405180910390fd5b5f60065f8673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f8567ffffffffffffffff1667ffffffffffffffff1681526020019081526020015f205f8460ff1660ff1681526020019081526020015f208054612a6e906141ed565b905011612ab0576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401612aa7906148ca565b60405180910390fd5b600160055f8673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f8567ffffffffffffffff1667ffffffffffffffff1681526020019081526020015f205f0160118282829054906101000a900460ff16612b2f919061490b565b92506101000a81548160ff021916908360ff16021790555081600860055f8773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f8667ffffffffffffffff1667ffffffffffffffff1681526020019081526020015f205f01601a9054906101000a900460ff16612bc49190614b42565b612bce919061490b565b60ff1660055f8673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f8567ffffffffffffffff1667ffffffffffffffff1681526020019081526020015f205f0160126101000a81548167ffffffffffffffff021916908367ffffffffffffffff1602179055508160055f8673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f8567ffffffffffffffff1667ffffffffffffffff1681526020019081526020015f205f01601a6101000a81548160ff021916908360ff16021790555060055f8573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f8467ffffffffffffffff1667ffffffffffffffff1681526020019081526020015f205f0160119054906101000a900460ff1690509392505050565b612d5561336a565b5f73ffffffffffffffffffffffffffffffffffffffff16816020015173ffffffffffffffffffffffffffffffffffffffff1603612dc7576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401612dbe90614bc8565b60405180910390fd5b5f73ffffffffffffffffffffffffffffffffffffffff1660075f8460ff1660ff1681526020019081526020015f206001015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1614612e6e576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401612e6590614c30565b60405180910390fd5b5f815f015190507f30644e72e131a029b85045b68181585d2833e84879b9709143e1f593f00000018106905080825f0181815250508160075f8560ff1660ff1681526020019081526020015f205f820151815f01556020820151816001015f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550905050505050565b612f1961336a565b5f73ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1603612f87576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401612f7e90614cbe565b60405180910390fd5b612f90816133e8565b50565b60605f67ffffffffffffffff90505f60c0828516901b9050604082901b91505f6040838616901b90508082612fc89190614cdc565b9150604083901b92506040838616901c90508082612fe69190614cdc565b9150604083901b925060c0838616901c905080826130049190614cdc565b91505f603067ffffffffffffffff81111561302257613021613628565b5b6040519080825280601f01601f1916602001820160405280156130545781602001600182028036833780820191505090505b50905082602082015280945050505050919050565b5f808367ffffffffffffffff1603613083575f90506131be565b5f600890505f600190505b8467ffffffffffffffff168267ffffffffffffffff1610156130dd5760045f9054906101000a900460ff1660ff16826130c79190614d0f565b91506001816130d6919061490b565b905061308e565b60048160ff1611156131185760045f9054906101000a900460ff1660ff16826131069190614d78565b91506001816131159190614da8565b90505b5b60068160ff1611156131585760045f9054906101000a900460ff1660ff16826131429190614d78565b91506001816131519190614da8565b9050613119565b5f84866131659190614583565b90505b8067ffffffffffffffff168367ffffffffffffffff1610156131b75760045f9054906101000a900460ff1660ff16836131a19190614d0f565b92506001826131b0919061490b565b9150613168565b8193505050505b92915050565b5f8367ffffffffffffffff1682846131dc9190614055565b67ffffffffffffffff1610156131ff5781836131f89190614055565b905061325a565b5f8367ffffffffffffffff161115613256578267ffffffffffffffff1685878460405160200161323193929190614ddc565b604051602081830303815290604052805190602001205f1c6132539190614e18565b92505b8290505b95945050505050565b60015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663693ec85e6040518163ffffffff1660e01b81526004016132bb90614e92565b6020604051808303815f875af11580156132d7573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906132fb9190613ffd565b73ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614613368576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161335f90614efa565b60405180910390fd5b565b6133726134a9565b73ffffffffffffffffffffffffffffffffffffffff16613390610e1c565b73ffffffffffffffffffffffffffffffffffffffff16146133e6576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016133dd90614f62565b60405180910390fd5b565b5f805f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050815f806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508173ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a35050565b5f33905090565b60405180604001604052805f81526020015f73ffffffffffffffffffffffffffffffffffffffff1681525090565b6040518060c001604052805f67ffffffffffffffff1681526020015f67ffffffffffffffff1681526020015f60ff1681526020015f60ff1681526020015f67ffffffffffffffff1681526020015f60ff1681525090565b5f604051905090565b5f80fd5b5f80fd5b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f61356f82613546565b9050919050565b61357f81613565565b8114613589575f80fd5b50565b5f8135905061359a81613576565b92915050565b5f67ffffffffffffffff82169050919050565b6135bc816135a0565b81146135c6575f80fd5b50565b5f813590506135d7816135b3565b92915050565b5f819050919050565b6135ef816135dd565b81146135f9575f80fd5b50565b5f8135905061360a816135e6565b92915050565b5f80fd5b5f80fd5b5f601f19601f8301169050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b61365e82613618565b810181811067ffffffffffffffff8211171561367d5761367c613628565b5b80604052505050565b5f61368f613535565b905061369b8282613655565b919050565b5f67ffffffffffffffff8211156136ba576136b9613628565b5b6136c382613618565b9050602081019050919050565b828183375f83830152505050565b5f6136f06136eb846136a0565b613686565b90508281526020810184848401111561370c5761370b613614565b5b6137178482856136d0565b509392505050565b5f82601f83011261373357613732613610565b5b81356137438482602086016136de565b91505092915050565b5f805f80608085870312156137645761376361353e565b5b5f6137718782880161358c565b9450506020613782878288016135c9565b9350506040613793878288016135fc565b925050606085013567ffffffffffffffff8111156137b4576137b3613542565b5b6137c08782880161371f565b91505092959194509250565b5f81519050919050565b5f82825260208201905092915050565b5f5b838110156138035780820151818401526020810190506137e8565b5f8484015250505050565b5f613818826137cc565b61382281856137d6565b93506138328185602086016137e6565b61383b81613618565b840191505092915050565b5f6020820190508181035f83015261385e818461380e565b905092915050565b5f806040838503121561387c5761387b61353e565b5b5f613889858286016135c9565b925050602061389a858286016135c9565b9150509250929050565b5f60ff82169050919050565b6138b9816138a4565b82525050565b5f6020820190506138d25f8301846138b0565b92915050565b5f805f805f60a086880312156138f1576138f061353e565b5b5f6138fe8882890161358c565b955050602061390f888289016135fc565b9450506040613920888289016135c9565b9350506060613931888289016135c9565b9250506080613942888289016135c9565b9150509295509295909350565b613958816135a0565b82525050565b5f6020820190506139715f83018461394f565b92915050565b5f805f6060848603121561398e5761398d61353e565b5b5f61399b8682870161358c565b93505060206139ac868287016135c9565b925050604084013567ffffffffffffffff8111156139cd576139cc613542565b5b6139d98682870161371f565b9150509250925092565b6139ec81613565565b82525050565b5f602082019050613a055f8301846139e3565b92915050565b5f67ffffffffffffffff821115613a2557613a24613628565b5b602082029050602081019050919050565b5f80fd5b5f613a4c613a4784613a0b565b613686565b90508083825260208201905060208402830185811115613a6f57613a6e613a36565b5b835b81811015613ab657803567ffffffffffffffff811115613a9457613a93613610565b5b808601613aa1898261371f565b85526020850194505050602081019050613a71565b5050509392505050565b5f82601f830112613ad457613ad3613610565b5b8135613ae4848260208601613a3a565b91505092915050565b5f805f8060808587031215613b0557613b0461353e565b5b5f613b128782880161358c565b9450506020613b23878288016135c9565b935050604085013567ffffffffffffffff811115613b4457613b43613542565b5b613b5087828801613ac0565b925050606085013567ffffffffffffffff811115613b7157613b70613542565b5b613b7d8782880161371f565b91505092959194509250565b5f805f8060808587031215613ba157613ba061353e565b5b5f613bae8782880161358c565b9450506020613bbf878288016135c9565b935050604085013567ffffffffffffffff811115613be057613bdf613542565b5b613bec8782880161371f565b925050606085013567ffffffffffffffff811115613c0d57613c0c613542565b5b613c198782880161371f565b91505092959194509250565b613c2e816138a4565b8114613c38575f80fd5b50565b5f81359050613c4981613c25565b92915050565b5f60208284031215613c6457613c6361353e565b5b5f613c7184828501613c3b565b91505092915050565b5f819050919050565b613c8c81613c7a565b82525050565b613c9b81613565565b82525050565b604082015f820151613cb55f850182613c83565b506020820151613cc86020850182613c92565b50505050565b5f604082019050613ce15f830184613ca1565b92915050565b5f8060408385031215613cfd57613cfc61353e565b5b5f613d0a8582860161358c565b9250506020613d1b858286016135c9565b9150509250929050565b613d2e816135a0565b82525050565b613d3d816138a4565b82525050565b60c082015f820151613d575f850182613d25565b506020820151613d6a6020850182613d25565b506040820151613d7d6040850182613d34565b506060820151613d906060850182613d34565b506080820151613da36080850182613d25565b5060a0820151613db660a0850182613d34565b50505050565b5f60c082019050613dcf5f830184613d43565b92915050565b5f805f60608486031215613dec57613deb61353e565b5b5f613df98682870161358c565b9350506020613e0a868287016135c9565b9250506040613e1b86828701613c3b565b9150509250925092565b5f80fd5b613e3281613c7a565b8114613e3c575f80fd5b50565b5f81359050613e4d81613e29565b92915050565b5f60408284031215613e6857613e67613e25565b5b613e726040613686565b90505f613e8184828501613e3f565b5f830152506020613e948482850161358c565b60208301525092915050565b5f8060608385031215613eb657613eb561353e565b5b5f613ec385828601613c3b565b9250506020613ed485828601613e53565b9150509250929050565b5f60208284031215613ef357613ef261353e565b5b5f613f008482850161358c565b91505092915050565b5f82825260208201905092915050565b7f6e6f2076000000000000000000000000000000000000000000000000000000005f82015250565b5f613f4d600483613f09565b9150613f5882613f19565b602082019050919050565b5f6020820190508181035f830152613f7a81613f41565b9050919050565b7f65706f63680000000000000000000000000000000000000000000000000000005f82015250565b5f613fb5600583613f09565b9150613fc082613f81565b602082019050919050565b5f6020820190508181035f830152613fe281613fa9565b9050919050565b5f81519050613ff781613576565b92915050565b5f602082840312156140125761401161353e565b5b5f61401f84828501613fe9565b91505092915050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b5f61405f826135a0565b915061406a836135a0565b9250828201905067ffffffffffffffff81111561408a57614089614028565b5b92915050565b5f8151905061409e81613e29565b92915050565b5f815190506140b2816135e6565b92915050565b5f80604083850312156140ce576140cd61353e565b5b5f6140db85828601614090565b92505060206140ec858286016140a4565b9150509250929050565b5f819050919050565b61411061410b826135dd565b6140f6565b82525050565b5f8160601b9050919050565b5f61412c82614116565b9050919050565b5f61413d82614122565b9050919050565b61415561415082613565565b614133565b82525050565b5f61416682856140ff565b6020820191506141768284614144565b6014820191508190509392505050565b5f81905092915050565b5f61419a826137cc565b6141a48185614186565b93506141b48185602086016137e6565b80840191505092915050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52602260045260245ffd5b5f600282049050600182168061420457607f821691505b602082108103614217576142166141c0565b5b50919050565b5f819050815f5260205f209050919050565b5f815461423b816141ed565b6142458186614186565b9450600182165f811461425f5760018114614274576142a6565b60ff19831686528115158202860193506142a6565b61427d8561421d565b5f5b8381101561429e5781548189015260018201915060208101905061427f565b838801955050505b50505092915050565b5f6142ba82866140ff565b6020820191506142ca8285614190565b91506142d6828461422f565b9150819050949350505050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52603260045260245ffd5b5f81519050919050565b5f82825260208201905092915050565b5f819050602082019050919050565b5f6143448383613c83565b60208301905092915050565b5f602082019050919050565b5f61436682614310565b614370818561431a565b935061437b8361432a565b805f5b838110156143ab5781516143928882614339565b975061439d83614350565b92505060018101905061437e565b5085935050505092915050565b5f6040820190508181035f8301526143d0818561380e565b905081810360208301526143e4818461435c565b90509392505050565b5f8115159050919050565b614401816143ed565b811461440b575f80fd5b50565b5f8151905061441c816143f8565b92915050565b5f602082840312156144375761443661353e565b5b5f6144448482850161440e565b91505092915050565b7f696e7620706600000000000000000000000000000000000000000000000000005f82015250565b5f614481600683613f09565b915061448c8261444d565b602082019050919050565b5f6020820190508181035f8301526144ae81614475565b9050919050565b7f70696563650000000000000000000000000000000000000000000000000000005f82015250565b5f6144e9600583613f09565b91506144f4826144b5565b602082019050919050565b5f6020820190508181035f830152614516816144dd565b9050919050565b5f6040820190506145305f8301856139e3565b61453d602083018461394f565b9392505050565b5f81519050614552816135b3565b92915050565b5f6020828403121561456d5761456c61353e565b5b5f61457a84828501614544565b91505092915050565b5f61458d826135a0565b9150614598836135a0565b9250828203905067ffffffffffffffff8111156145b8576145b7614028565b5b92915050565b5f6020601f8301049050919050565b5f82821b905092915050565b5f600883026146087fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff826145cd565b61461286836145cd565b95508019841693508086168417925050509392505050565b5f819050919050565b5f61464d61464861464384613c7a565b61462a565b613c7a565b9050919050565b5f819050919050565b61466683614633565b61467a61467282614654565b8484546145d9565b825550505050565b5f90565b61468e614682565b61469981848461465d565b505050565b5b818110156146bc576146b15f82614686565b60018101905061469f565b5050565b601f821115614701576146d28161421d565b6146db846145be565b810160208510156146ea578190505b6146fe6146f6856145be565b83018261469e565b50505b505050565b5f82821c905092915050565b5f6147215f1984600802614706565b1980831691505092915050565b5f6147398383614712565b9150826002028217905092915050565b614752826137cc565b67ffffffffffffffff81111561476b5761476a613628565b5b61477582546141ed565b6147808282856146c0565b5f60209050601f8311600181146147b1575f841561479f578287015190505b6147a9858261472e565b865550614810565b601f1984166147bf8661421d565b5f5b828110156147e6578489015182556001820191506020850194506020810190506147c1565b8683101561480357848901516147ff601f891682614712565b8355505b6001600288020188555050505b505050505050565b7f696e7620636f6d000000000000000000000000000000000000000000000000005f82015250565b5f61484c600783613f09565b915061485782614818565b602082019050919050565b5f6020820190508181035f83015261487981614840565b9050919050565b7f6e6f2073756d00000000000000000000000000000000000000000000000000005f82015250565b5f6148b4600683613f09565b91506148bf82614880565b602082019050919050565b5f6020820190508181035f8301526148e1816148a8565b9050919050565b5f6148f38285614190565b91506148ff8284614190565b91508190509392505050565b5f614915826138a4565b9150614920836138a4565b9250828201905060ff81111561493957614938614028565b5b92915050565b7f696e73756620726f756e640000000000000000000000000000000000000000005f82015250565b5f614973600b83613f09565b915061497e8261493f565b602082019050919050565b5f6020820190508181035f8301526149a081614967565b9050919050565b5f6060820190506149ba5f8301866139e3565b6149c7602083018561394f565b81810360408301526149d9818461380e565b9050949350505050565b5f8160c01b9050919050565b5f6149f9826149e3565b9050919050565b614a11614a0c826135a0565b6149ef565b82525050565b5f614a2282856140ff565b602082019150614a328284614a00565b6008820191508190509392505050565b5f614a4d8286614190565b9150614a598285614190565b9150614a658284614190565b9150819050949350505050565b7f696e7620710000000000000000000000000000000000000000000000000000005f82015250565b5f614aa6600583613f09565b9150614ab182614a72565b602082019050919050565b5f6020820190508181035f830152614ad381614a9a565b9050919050565b7f65786365656420726f756e6400000000000000000000000000000000000000005f82015250565b5f614b0e600c83613f09565b9150614b1982614ada565b602082019050919050565b5f6020820190508181035f830152614b3b81614b02565b9050919050565b5f614b4c826138a4565b9150614b57836138a4565b9250828202614b65816138a4565b9150808214614b7757614b76614028565b5b5092915050565b7f696e7620766b00000000000000000000000000000000000000000000000000005f82015250565b5f614bb2600683613f09565b9150614bbd82614b7e565b602082019050919050565b5f6020820190508181035f830152614bdf81614ba6565b9050919050565b7f68617320766b00000000000000000000000000000000000000000000000000005f82015250565b5f614c1a600683613f09565b9150614c2582614be6565b602082019050919050565b5f6020820190508181035f830152614c4781614c0e565b9050919050565b7f4f776e61626c653a206e6577206f776e657220697320746865207a65726f20615f8201527f6464726573730000000000000000000000000000000000000000000000000000602082015250565b5f614ca8602683613f09565b9150614cb382614c4e565b604082019050919050565b5f6020820190508181035f830152614cd581614c9c565b9050919050565b5f614ce682613c7a565b9150614cf183613c7a565b9250828201905080821115614d0957614d08614028565b5b92915050565b5f614d19826135a0565b9150614d24836135a0565b9250828202614d32816135a0565b9150808214614d4457614d43614028565b5b5092915050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601260045260245ffd5b5f614d82826135a0565b9150614d8d836135a0565b925082614d9d57614d9c614d4b565b5b828204905092915050565b5f614db2826138a4565b9150614dbd836138a4565b9250828203905060ff811115614dd657614dd5614028565b5b92915050565b5f614de782866140ff565b602082019150614df78285614144565b601482019150614e078284614a00565b600882019150819050949350505050565b5f614e2282613c7a565b9150614e2d83613c7a565b925082614e3d57614e3c614d4b565b5b828206905092915050565b7f6570726f6f6600000000000000000000000000000000000000000000000000005f82015250565b5f614e7c600683613f09565b9150614e8782614e48565b602082019050919050565b5f6020820190508181035f830152614ea981614e70565b9050919050565b7f696e7620656300000000000000000000000000000000000000000000000000005f82015250565b5f614ee4600683613f09565b9150614eef82614eb0565b602082019050919050565b5f6020820190508181035f830152614f1181614ed8565b9050919050565b7f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e65725f82015250565b5f614f4c602083613f09565b9150614f5782614f18565b602082019050919050565b5f6020820190508181035f830152614f7981614f40565b905091905056fea264697066735822122064aac1463e2b0adcaa101076df301f0efd7ec8c38f8ac7bffdb067f3be1a657464736f6c634300081a0033eab9b16eb21be9efd5481512ffcd394e188282c8bd37cb5c85951e2caa9d41bbc8fc6225bf87ff54008848defe740a67fd82de55559c8ea6c2fe3d3634a9591a6d182ad44fb82305bd7fb348ca3e52d91f674f5d30afeec401914a69c5102eff3b8201b322c65a735690cf82c850b8624c29ec05400e06ba92a9aad12c37c1605812abbc9a1a11f500b3ab28b7751b52",
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

// GetVerifier is a free data retrieval call binding the contract method 0xadaad467.
//
// Solidity: function getVerifier(uint8 _ptype) view returns((uint256,address) _vi)
func (_EVerify *EVerifyCaller) GetVerifier(opts *bind.CallOpts, _ptype uint8) (EVerifyVInfo, error) {
	var out []interface{}
	err := _EVerify.contract.Call(opts, &out, "getVerifier", _ptype)

	if err != nil {
		return *new(EVerifyVInfo), err
	}

	out0 := *abi.ConvertType(out[0], new(EVerifyVInfo)).(*EVerifyVInfo)

	return out0, err

}

// GetVerifier is a free data retrieval call binding the contract method 0xadaad467.
//
// Solidity: function getVerifier(uint8 _ptype) view returns((uint256,address) _vi)
func (_EVerify *EVerifySession) GetVerifier(_ptype uint8) (EVerifyVInfo, error) {
	return _EVerify.Contract.GetVerifier(&_EVerify.CallOpts, _ptype)
}

// GetVerifier is a free data retrieval call binding the contract method 0xadaad467.
//
// Solidity: function getVerifier(uint8 _ptype) view returns((uint256,address) _vi)
func (_EVerify *EVerifyCallerSession) GetVerifier(_ptype uint8) (EVerifyVInfo, error) {
	return _EVerify.Contract.GetVerifier(&_EVerify.CallOpts, _ptype)
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
// Solidity: function challenge(address _a, uint64 _ep, bytes _sum) returns(uint8)
func (_EVerify *EVerifyTransactor) Challenge(opts *bind.TransactOpts, _a common.Address, _ep uint64, _sum []byte) (*types.Transaction, error) {
	return _EVerify.contract.Transact(opts, "challenge", _a, _ep, _sum)
}

// Challenge is a paid mutator transaction binding the contract method 0x5a9b427d.
//
// Solidity: function challenge(address _a, uint64 _ep, bytes _sum) returns(uint8)
func (_EVerify *EVerifySession) Challenge(_a common.Address, _ep uint64, _sum []byte) (*types.Transaction, error) {
	return _EVerify.Contract.Challenge(&_EVerify.TransactOpts, _a, _ep, _sum)
}

// Challenge is a paid mutator transaction binding the contract method 0x5a9b427d.
//
// Solidity: function challenge(address _a, uint64 _ep, bytes _sum) returns(uint8)
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

// SetVeirfier is a paid mutator transaction binding the contract method 0xbe12c9ea.
//
// Solidity: function setVeirfier(uint8 _ptype, (uint256,address) _vi) returns()
func (_EVerify *EVerifyTransactor) SetVeirfier(opts *bind.TransactOpts, _ptype uint8, _vi EVerifyVInfo) (*types.Transaction, error) {
	return _EVerify.contract.Transact(opts, "setVeirfier", _ptype, _vi)
}

// SetVeirfier is a paid mutator transaction binding the contract method 0xbe12c9ea.
//
// Solidity: function setVeirfier(uint8 _ptype, (uint256,address) _vi) returns()
func (_EVerify *EVerifySession) SetVeirfier(_ptype uint8, _vi EVerifyVInfo) (*types.Transaction, error) {
	return _EVerify.Contract.SetVeirfier(&_EVerify.TransactOpts, _ptype, _vi)
}

// SetVeirfier is a paid mutator transaction binding the contract method 0xbe12c9ea.
//
// Solidity: function setVeirfier(uint8 _ptype, (uint256,address) _vi) returns()
func (_EVerify *EVerifyTransactorSession) SetVeirfier(_ptype uint8, _vi EVerifyVInfo) (*types.Transaction, error) {
	return _EVerify.Contract.SetVeirfier(&_EVerify.TransactOpts, _ptype, _vi)
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
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_a\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"_ep\",\"type\":\"uint64\"},{\"internalType\":\"uint8\",\"name\":\"_qIndex\",\"type\":\"uint8\"}],\"name\":\"chalCom\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_a\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"_ep\",\"type\":\"uint64\"},{\"internalType\":\"uint8\",\"name\":\"_qIndex\",\"type\":\"uint8\"}],\"name\":\"chalOne\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_a\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"_ep\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"_com\",\"type\":\"bytes\"}],\"name\":\"challenge\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_a\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"_ep\",\"type\":\"uint64\"},{\"internalType\":\"bytes[]\",\"name\":\"_com\",\"type\":\"bytes[]\"},{\"internalType\":\"bytes\",\"name\":\"_proof\",\"type\":\"bytes\"}],\"name\":\"proveCom\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_a\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"_ep\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"_wroot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"_proof\",\"type\":\"bytes\"}],\"name\":\"proveKZG\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_a\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"_ep\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"_com\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"_proof\",\"type\":\"bytes\"}],\"name\":\"proveOne\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
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
// Solidity: function challenge(address _a, uint64 _ep, bytes _com) returns(uint8)
func (_IEVerify *IEVerifyTransactor) Challenge(opts *bind.TransactOpts, _a common.Address, _ep uint64, _com []byte) (*types.Transaction, error) {
	return _IEVerify.contract.Transact(opts, "challenge", _a, _ep, _com)
}

// Challenge is a paid mutator transaction binding the contract method 0x5a9b427d.
//
// Solidity: function challenge(address _a, uint64 _ep, bytes _com) returns(uint8)
func (_IEVerify *IEVerifySession) Challenge(_a common.Address, _ep uint64, _com []byte) (*types.Transaction, error) {
	return _IEVerify.Contract.Challenge(&_IEVerify.TransactOpts, _a, _ep, _com)
}

// Challenge is a paid mutator transaction binding the contract method 0x5a9b427d.
//
// Solidity: function challenge(address _a, uint64 _ep, bytes _com) returns(uint8)
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

// IEproofMetaData contains all meta data concerning the IEproof contract.
var IEproofMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_s\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"_e\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"_round\",\"type\":\"uint8\"}],\"name\":\"Challenge\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_s\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"_e\",\"type\":\"uint64\"}],\"name\":\"Fake\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_s\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"_e\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"_round\",\"type\":\"uint8\"}],\"name\":\"Prove\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_s\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"_e\",\"type\":\"uint64\"}],\"name\":\"Submit\",\"type\":\"event\"}]",
}

// IEproofABI is the input ABI used to generate the binding from.
// Deprecated: Use IEproofMetaData.ABI instead.
var IEproofABI = IEproofMetaData.ABI

// IEproof is an auto generated Go binding around an Ethereum contract.
type IEproof struct {
	IEproofCaller     // Read-only binding to the contract
	IEproofTransactor // Write-only binding to the contract
	IEproofFilterer   // Log filterer for contract events
}

// IEproofCaller is an auto generated read-only Go binding around an Ethereum contract.
type IEproofCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IEproofTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IEproofTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IEproofFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IEproofFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IEproofSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IEproofSession struct {
	Contract     *IEproof          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IEproofCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IEproofCallerSession struct {
	Contract *IEproofCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// IEproofTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IEproofTransactorSession struct {
	Contract     *IEproofTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// IEproofRaw is an auto generated low-level Go binding around an Ethereum contract.
type IEproofRaw struct {
	Contract *IEproof // Generic contract binding to access the raw methods on
}

// IEproofCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IEproofCallerRaw struct {
	Contract *IEproofCaller // Generic read-only contract binding to access the raw methods on
}

// IEproofTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IEproofTransactorRaw struct {
	Contract *IEproofTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIEproof creates a new instance of IEproof, bound to a specific deployed contract.
func NewIEproof(address common.Address, backend bind.ContractBackend) (*IEproof, error) {
	contract, err := bindIEproof(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IEproof{IEproofCaller: IEproofCaller{contract: contract}, IEproofTransactor: IEproofTransactor{contract: contract}, IEproofFilterer: IEproofFilterer{contract: contract}}, nil
}

// NewIEproofCaller creates a new read-only instance of IEproof, bound to a specific deployed contract.
func NewIEproofCaller(address common.Address, caller bind.ContractCaller) (*IEproofCaller, error) {
	contract, err := bindIEproof(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IEproofCaller{contract: contract}, nil
}

// NewIEproofTransactor creates a new write-only instance of IEproof, bound to a specific deployed contract.
func NewIEproofTransactor(address common.Address, transactor bind.ContractTransactor) (*IEproofTransactor, error) {
	contract, err := bindIEproof(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IEproofTransactor{contract: contract}, nil
}

// NewIEproofFilterer creates a new log filterer instance of IEproof, bound to a specific deployed contract.
func NewIEproofFilterer(address common.Address, filterer bind.ContractFilterer) (*IEproofFilterer, error) {
	contract, err := bindIEproof(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IEproofFilterer{contract: contract}, nil
}

// bindIEproof binds a generic wrapper to an already deployed contract.
func bindIEproof(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := IEproofMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IEproof *IEproofRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IEproof.Contract.IEproofCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IEproof *IEproofRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IEproof.Contract.IEproofTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IEproof *IEproofRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IEproof.Contract.IEproofTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IEproof *IEproofCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IEproof.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IEproof *IEproofTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IEproof.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IEproof *IEproofTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IEproof.Contract.contract.Transact(opts, method, params...)
}

// IEproofChallengeIterator is returned from FilterChallenge and is used to iterate over the raw logs and unpacked data for Challenge events raised by the IEproof contract.
type IEproofChallengeIterator struct {
	Event *IEproofChallenge // Event containing the contract specifics and raw log

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
func (it *IEproofChallengeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IEproofChallenge)
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
		it.Event = new(IEproofChallenge)
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
func (it *IEproofChallengeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IEproofChallengeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IEproofChallenge represents a Challenge event raised by the IEproof contract.
type IEproofChallenge struct {
	S     common.Address
	E     uint64
	Round uint8
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterChallenge is a free log retrieval operation binding the contract event 0x6fb32eb2afe1710cb5810cfbfc9177d3c89ee79998070eec250d3255d4c26ddb.
//
// Solidity: event Challenge(address indexed _s, uint64 _e, uint8 _round)
func (_IEproof *IEproofFilterer) FilterChallenge(opts *bind.FilterOpts, _s []common.Address) (*IEproofChallengeIterator, error) {

	var _sRule []interface{}
	for _, _sItem := range _s {
		_sRule = append(_sRule, _sItem)
	}

	logs, sub, err := _IEproof.contract.FilterLogs(opts, "Challenge", _sRule)
	if err != nil {
		return nil, err
	}
	return &IEproofChallengeIterator{contract: _IEproof.contract, event: "Challenge", logs: logs, sub: sub}, nil
}

// WatchChallenge is a free log subscription operation binding the contract event 0x6fb32eb2afe1710cb5810cfbfc9177d3c89ee79998070eec250d3255d4c26ddb.
//
// Solidity: event Challenge(address indexed _s, uint64 _e, uint8 _round)
func (_IEproof *IEproofFilterer) WatchChallenge(opts *bind.WatchOpts, sink chan<- *IEproofChallenge, _s []common.Address) (event.Subscription, error) {

	var _sRule []interface{}
	for _, _sItem := range _s {
		_sRule = append(_sRule, _sItem)
	}

	logs, sub, err := _IEproof.contract.WatchLogs(opts, "Challenge", _sRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IEproofChallenge)
				if err := _IEproof.contract.UnpackLog(event, "Challenge", log); err != nil {
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

// ParseChallenge is a log parse operation binding the contract event 0x6fb32eb2afe1710cb5810cfbfc9177d3c89ee79998070eec250d3255d4c26ddb.
//
// Solidity: event Challenge(address indexed _s, uint64 _e, uint8 _round)
func (_IEproof *IEproofFilterer) ParseChallenge(log types.Log) (*IEproofChallenge, error) {
	event := new(IEproofChallenge)
	if err := _IEproof.contract.UnpackLog(event, "Challenge", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IEproofFakeIterator is returned from FilterFake and is used to iterate over the raw logs and unpacked data for Fake events raised by the IEproof contract.
type IEproofFakeIterator struct {
	Event *IEproofFake // Event containing the contract specifics and raw log

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
func (it *IEproofFakeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IEproofFake)
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
		it.Event = new(IEproofFake)
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
func (it *IEproofFakeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IEproofFakeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IEproofFake represents a Fake event raised by the IEproof contract.
type IEproofFake struct {
	S   common.Address
	E   uint64
	Raw types.Log // Blockchain specific contextual infos
}

// FilterFake is a free log retrieval operation binding the contract event 0x0fcd3bb0ce91639d74185d625544f2b5714ecd14a6a6d9ccb65b2526e9ac12f5.
//
// Solidity: event Fake(address indexed _s, uint64 _e)
func (_IEproof *IEproofFilterer) FilterFake(opts *bind.FilterOpts, _s []common.Address) (*IEproofFakeIterator, error) {

	var _sRule []interface{}
	for _, _sItem := range _s {
		_sRule = append(_sRule, _sItem)
	}

	logs, sub, err := _IEproof.contract.FilterLogs(opts, "Fake", _sRule)
	if err != nil {
		return nil, err
	}
	return &IEproofFakeIterator{contract: _IEproof.contract, event: "Fake", logs: logs, sub: sub}, nil
}

// WatchFake is a free log subscription operation binding the contract event 0x0fcd3bb0ce91639d74185d625544f2b5714ecd14a6a6d9ccb65b2526e9ac12f5.
//
// Solidity: event Fake(address indexed _s, uint64 _e)
func (_IEproof *IEproofFilterer) WatchFake(opts *bind.WatchOpts, sink chan<- *IEproofFake, _s []common.Address) (event.Subscription, error) {

	var _sRule []interface{}
	for _, _sItem := range _s {
		_sRule = append(_sRule, _sItem)
	}

	logs, sub, err := _IEproof.contract.WatchLogs(opts, "Fake", _sRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IEproofFake)
				if err := _IEproof.contract.UnpackLog(event, "Fake", log); err != nil {
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
// Solidity: event Fake(address indexed _s, uint64 _e)
func (_IEproof *IEproofFilterer) ParseFake(log types.Log) (*IEproofFake, error) {
	event := new(IEproofFake)
	if err := _IEproof.contract.UnpackLog(event, "Fake", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IEproofProveIterator is returned from FilterProve and is used to iterate over the raw logs and unpacked data for Prove events raised by the IEproof contract.
type IEproofProveIterator struct {
	Event *IEproofProve // Event containing the contract specifics and raw log

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
func (it *IEproofProveIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IEproofProve)
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
		it.Event = new(IEproofProve)
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
func (it *IEproofProveIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IEproofProveIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IEproofProve represents a Prove event raised by the IEproof contract.
type IEproofProve struct {
	S     common.Address
	E     uint64
	Round uint8
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterProve is a free log retrieval operation binding the contract event 0x074353199408473b546546e1626a6b8ac62ff72784909ac8d0708a73125484af.
//
// Solidity: event Prove(address indexed _s, uint64 _e, uint8 _round)
func (_IEproof *IEproofFilterer) FilterProve(opts *bind.FilterOpts, _s []common.Address) (*IEproofProveIterator, error) {

	var _sRule []interface{}
	for _, _sItem := range _s {
		_sRule = append(_sRule, _sItem)
	}

	logs, sub, err := _IEproof.contract.FilterLogs(opts, "Prove", _sRule)
	if err != nil {
		return nil, err
	}
	return &IEproofProveIterator{contract: _IEproof.contract, event: "Prove", logs: logs, sub: sub}, nil
}

// WatchProve is a free log subscription operation binding the contract event 0x074353199408473b546546e1626a6b8ac62ff72784909ac8d0708a73125484af.
//
// Solidity: event Prove(address indexed _s, uint64 _e, uint8 _round)
func (_IEproof *IEproofFilterer) WatchProve(opts *bind.WatchOpts, sink chan<- *IEproofProve, _s []common.Address) (event.Subscription, error) {

	var _sRule []interface{}
	for _, _sItem := range _s {
		_sRule = append(_sRule, _sItem)
	}

	logs, sub, err := _IEproof.contract.WatchLogs(opts, "Prove", _sRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IEproofProve)
				if err := _IEproof.contract.UnpackLog(event, "Prove", log); err != nil {
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

// ParseProve is a log parse operation binding the contract event 0x074353199408473b546546e1626a6b8ac62ff72784909ac8d0708a73125484af.
//
// Solidity: event Prove(address indexed _s, uint64 _e, uint8 _round)
func (_IEproof *IEproofFilterer) ParseProve(log types.Log) (*IEproofProve, error) {
	event := new(IEproofProve)
	if err := _IEproof.contract.UnpackLog(event, "Prove", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IEproofSubmitIterator is returned from FilterSubmit and is used to iterate over the raw logs and unpacked data for Submit events raised by the IEproof contract.
type IEproofSubmitIterator struct {
	Event *IEproofSubmit // Event containing the contract specifics and raw log

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
func (it *IEproofSubmitIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IEproofSubmit)
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
		it.Event = new(IEproofSubmit)
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
func (it *IEproofSubmitIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IEproofSubmitIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IEproofSubmit represents a Submit event raised by the IEproof contract.
type IEproofSubmit struct {
	S   common.Address
	E   uint64
	Raw types.Log // Blockchain specific contextual infos
}

// FilterSubmit is a free log retrieval operation binding the contract event 0x98e23f0b72b943d75df730f1b4c0fbb9baffa041d18e1819f2d851a5c2c1c444.
//
// Solidity: event Submit(address indexed _s, uint64 _e)
func (_IEproof *IEproofFilterer) FilterSubmit(opts *bind.FilterOpts, _s []common.Address) (*IEproofSubmitIterator, error) {

	var _sRule []interface{}
	for _, _sItem := range _s {
		_sRule = append(_sRule, _sItem)
	}

	logs, sub, err := _IEproof.contract.FilterLogs(opts, "Submit", _sRule)
	if err != nil {
		return nil, err
	}
	return &IEproofSubmitIterator{contract: _IEproof.contract, event: "Submit", logs: logs, sub: sub}, nil
}

// WatchSubmit is a free log subscription operation binding the contract event 0x98e23f0b72b943d75df730f1b4c0fbb9baffa041d18e1819f2d851a5c2c1c444.
//
// Solidity: event Submit(address indexed _s, uint64 _e)
func (_IEproof *IEproofFilterer) WatchSubmit(opts *bind.WatchOpts, sink chan<- *IEproofSubmit, _s []common.Address) (event.Subscription, error) {

	var _sRule []interface{}
	for _, _sItem := range _s {
		_sRule = append(_sRule, _sItem)
	}

	logs, sub, err := _IEproof.contract.WatchLogs(opts, "Submit", _sRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IEproofSubmit)
				if err := _IEproof.contract.UnpackLog(event, "Submit", log); err != nil {
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

// ParseSubmit is a log parse operation binding the contract event 0x98e23f0b72b943d75df730f1b4c0fbb9baffa041d18e1819f2d851a5c2c1c444.
//
// Solidity: event Submit(address indexed _s, uint64 _e)
func (_IEproof *IEproofFilterer) ParseSubmit(log types.Log) (*IEproofSubmit, error) {
	event := new(IEproofSubmit)
	if err := _IEproof.contract.UnpackLog(event, "Submit", log); err != nil {
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
