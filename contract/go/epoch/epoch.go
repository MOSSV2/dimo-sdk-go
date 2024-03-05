// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package epoch

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

// IEpochEpochInfo is an auto generated low-level Go binding around an user-defined struct.
type IEpochEpochInfo struct {
	Start *big.Int
	Seed  [32]byte
}

// EpochMetaData contains all meta data concerning the Epoch contract.
var EpochMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"_e\",\"type\":\"uint64\"}],\"name\":\"SetEpoch\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"check\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"curEpoch\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"_epoch\",\"type\":\"uint64\"}],\"name\":\"getEpoch\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"start\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"seed\",\"type\":\"bytes32\"}],\"internalType\":\"structIEpoch.EpochInfo\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"_epoch\",\"type\":\"uint64\"}],\"name\":\"getEpochSeed\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"slotsInEpoch\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x60806040526121c06000806101000a81548167ffffffffffffffff021916908367ffffffffffffffff16021790555034801561003a57600080fd5b506100436100d5565b60008160000181815250504360014261005c919061012b565b4060405160200161006e9291906101ab565b60405160208183030381529060405280519060200120816020018181525050600181908060018154018082558091505060019003906000526020600020906002020160009091909190915060008201518160000155602082015181600101555050506101d7565b604051806040016040528060008152602001600080191681525090565b6000819050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b6000610136826100f2565b9150610141836100f2565b9250828203905081811115610159576101586100fc565b5b92915050565b6000819050919050565b61017a610175826100f2565b61015f565b82525050565b6000819050919050565b6000819050919050565b6101a56101a082610180565b61018a565b82525050565b60006101b78285610169565b6020820191506101c78284610194565b6020820191508190509392505050565b610755806101e66000396000f3fe608060405234801561001057600080fd5b50600436106100575760003560e01c806312a02c821461005c5780632249084c1461008c578063919840ad146100bc578063bfb3cbb9146100c6578063d8887315146100e4575b600080fd5b61007660048036038101906100719190610456565b610102565b60405161008391906104e4565b60405180910390f35b6100a660048036038101906100a19190610456565b610190565b6040516100b3919061050e565b60405180910390f35b6100c46101c9565b005b6100ce6101d3565b6040516100db9190610538565b60405180910390f35b6100ec6101eb565b6040516100f99190610538565b60405180910390f35b61010a6103f4565b6101126103f4565b604051806040016040528060018567ffffffffffffffff168154811061013b5761013a610553565b5b906000526020600020906002020160000154815260200160018567ffffffffffffffff16815481106101705761016f610553565b5b906000526020600020906002020160010154815250905080915050919050565b600060018267ffffffffffffffff16815481106101b0576101af610553565b5b9060005260206000209060020201600101549050919050565b6101d1610205565b565b60008054906101000a900467ffffffffffffffff1681565b600060089054906101000a900467ffffffffffffffff1681565b60008054906101000a900467ffffffffffffffff1667ffffffffffffffff166001600060089054906101000a900467ffffffffffffffff1667ffffffffffffffff168154811061025857610257610553565b5b9060005260206000209060020201600001544361027591906105b1565b106103f2576102826103f4565b438160000181815250506001600060089054906101000a900467ffffffffffffffff1667ffffffffffffffff16815481106102c0576102bf610553565b5b90600052602060002090600202016001015433436001426102e191906105b1565b406040516020016102f594939291906106a1565b604051602081830303815290604052805190602001208160200181815250506001819080600181540180825580915050600190039060005260206000209060020201600090919091909150600082015181600001556020820151816001015550506000600881819054906101000a900467ffffffffffffffff168092919061037c906106ef565b91906101000a81548167ffffffffffffffff021916908367ffffffffffffffff160217905550507f9d4ccb161ea809df14334516cc3070025c80baddb8e8364afaca6a6fe31bfd75600060089054906101000a900467ffffffffffffffff166040516103e89190610538565b60405180910390a1505b565b604051806040016040528060008152602001600080191681525090565b600080fd5b600067ffffffffffffffff82169050919050565b61043381610416565b811461043e57600080fd5b50565b6000813590506104508161042a565b92915050565b60006020828403121561046c5761046b610411565b5b600061047a84828501610441565b91505092915050565b6000819050919050565b61049681610483565b82525050565b6000819050919050565b6104af8161049c565b82525050565b6040820160008201516104cb600085018261048d565b5060208201516104de60208501826104a6565b50505050565b60006040820190506104f960008301846104b5565b92915050565b6105088161049c565b82525050565b600060208201905061052360008301846104ff565b92915050565b61053281610416565b82525050565b600060208201905061054d6000830184610529565b92915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b60006105bc82610483565b91506105c783610483565b92508282039050818111156105df576105de610582565b5b92915050565b6000819050919050565b6106006105fb8261049c565b6105e5565b82525050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b600061063182610606565b9050919050565b60008160601b9050919050565b600061065082610638565b9050919050565b600061066282610645565b9050919050565b61067a61067582610626565b610657565b82525050565b6000819050919050565b61069b61069682610483565b610680565b82525050565b60006106ad82876105ef565b6020820191506106bd8286610669565b6014820191506106cd828561068a565b6020820191506106dd82846105ef565b60208201915081905095945050505050565b60006106fa82610416565b915067ffffffffffffffff820361071457610713610582565b5b60018201905091905056fea26469706673582212205e51382b6e97b1c2849137284441f70fe15d29c754aa23d7b9ed62da7f36621c64736f6c63430008180033",
}

// EpochABI is the input ABI used to generate the binding from.
// Deprecated: Use EpochMetaData.ABI instead.
var EpochABI = EpochMetaData.ABI

// EpochBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use EpochMetaData.Bin instead.
var EpochBin = EpochMetaData.Bin

// DeployEpoch deploys a new Ethereum contract, binding an instance of Epoch to it.
func DeployEpoch(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Epoch, error) {
	parsed, err := EpochMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(EpochBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Epoch{EpochCaller: EpochCaller{contract: contract}, EpochTransactor: EpochTransactor{contract: contract}, EpochFilterer: EpochFilterer{contract: contract}}, nil
}

// Epoch is an auto generated Go binding around an Ethereum contract.
type Epoch struct {
	EpochCaller     // Read-only binding to the contract
	EpochTransactor // Write-only binding to the contract
	EpochFilterer   // Log filterer for contract events
}

// EpochCaller is an auto generated read-only Go binding around an Ethereum contract.
type EpochCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EpochTransactor is an auto generated write-only Go binding around an Ethereum contract.
type EpochTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EpochFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type EpochFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EpochSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type EpochSession struct {
	Contract     *Epoch            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// EpochCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type EpochCallerSession struct {
	Contract *EpochCaller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// EpochTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type EpochTransactorSession struct {
	Contract     *EpochTransactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// EpochRaw is an auto generated low-level Go binding around an Ethereum contract.
type EpochRaw struct {
	Contract *Epoch // Generic contract binding to access the raw methods on
}

// EpochCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type EpochCallerRaw struct {
	Contract *EpochCaller // Generic read-only contract binding to access the raw methods on
}

// EpochTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type EpochTransactorRaw struct {
	Contract *EpochTransactor // Generic write-only contract binding to access the raw methods on
}

// NewEpoch creates a new instance of Epoch, bound to a specific deployed contract.
func NewEpoch(address common.Address, backend bind.ContractBackend) (*Epoch, error) {
	contract, err := bindEpoch(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Epoch{EpochCaller: EpochCaller{contract: contract}, EpochTransactor: EpochTransactor{contract: contract}, EpochFilterer: EpochFilterer{contract: contract}}, nil
}

// NewEpochCaller creates a new read-only instance of Epoch, bound to a specific deployed contract.
func NewEpochCaller(address common.Address, caller bind.ContractCaller) (*EpochCaller, error) {
	contract, err := bindEpoch(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &EpochCaller{contract: contract}, nil
}

// NewEpochTransactor creates a new write-only instance of Epoch, bound to a specific deployed contract.
func NewEpochTransactor(address common.Address, transactor bind.ContractTransactor) (*EpochTransactor, error) {
	contract, err := bindEpoch(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &EpochTransactor{contract: contract}, nil
}

// NewEpochFilterer creates a new log filterer instance of Epoch, bound to a specific deployed contract.
func NewEpochFilterer(address common.Address, filterer bind.ContractFilterer) (*EpochFilterer, error) {
	contract, err := bindEpoch(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &EpochFilterer{contract: contract}, nil
}

// bindEpoch binds a generic wrapper to an already deployed contract.
func bindEpoch(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := EpochMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Epoch *EpochRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Epoch.Contract.EpochCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Epoch *EpochRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Epoch.Contract.EpochTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Epoch *EpochRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Epoch.Contract.EpochTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Epoch *EpochCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Epoch.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Epoch *EpochTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Epoch.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Epoch *EpochTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Epoch.Contract.contract.Transact(opts, method, params...)
}

// CurEpoch is a free data retrieval call binding the contract method 0xd8887315.
//
// Solidity: function curEpoch() view returns(uint64)
func (_Epoch *EpochCaller) CurEpoch(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _Epoch.contract.Call(opts, &out, "curEpoch")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// CurEpoch is a free data retrieval call binding the contract method 0xd8887315.
//
// Solidity: function curEpoch() view returns(uint64)
func (_Epoch *EpochSession) CurEpoch() (uint64, error) {
	return _Epoch.Contract.CurEpoch(&_Epoch.CallOpts)
}

// CurEpoch is a free data retrieval call binding the contract method 0xd8887315.
//
// Solidity: function curEpoch() view returns(uint64)
func (_Epoch *EpochCallerSession) CurEpoch() (uint64, error) {
	return _Epoch.Contract.CurEpoch(&_Epoch.CallOpts)
}

// GetEpoch is a free data retrieval call binding the contract method 0x12a02c82.
//
// Solidity: function getEpoch(uint64 _epoch) view returns((uint256,bytes32))
func (_Epoch *EpochCaller) GetEpoch(opts *bind.CallOpts, _epoch uint64) (IEpochEpochInfo, error) {
	var out []interface{}
	err := _Epoch.contract.Call(opts, &out, "getEpoch", _epoch)

	if err != nil {
		return *new(IEpochEpochInfo), err
	}

	out0 := *abi.ConvertType(out[0], new(IEpochEpochInfo)).(*IEpochEpochInfo)

	return out0, err

}

// GetEpoch is a free data retrieval call binding the contract method 0x12a02c82.
//
// Solidity: function getEpoch(uint64 _epoch) view returns((uint256,bytes32))
func (_Epoch *EpochSession) GetEpoch(_epoch uint64) (IEpochEpochInfo, error) {
	return _Epoch.Contract.GetEpoch(&_Epoch.CallOpts, _epoch)
}

// GetEpoch is a free data retrieval call binding the contract method 0x12a02c82.
//
// Solidity: function getEpoch(uint64 _epoch) view returns((uint256,bytes32))
func (_Epoch *EpochCallerSession) GetEpoch(_epoch uint64) (IEpochEpochInfo, error) {
	return _Epoch.Contract.GetEpoch(&_Epoch.CallOpts, _epoch)
}

// GetEpochSeed is a free data retrieval call binding the contract method 0x2249084c.
//
// Solidity: function getEpochSeed(uint64 _epoch) view returns(bytes32)
func (_Epoch *EpochCaller) GetEpochSeed(opts *bind.CallOpts, _epoch uint64) ([32]byte, error) {
	var out []interface{}
	err := _Epoch.contract.Call(opts, &out, "getEpochSeed", _epoch)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetEpochSeed is a free data retrieval call binding the contract method 0x2249084c.
//
// Solidity: function getEpochSeed(uint64 _epoch) view returns(bytes32)
func (_Epoch *EpochSession) GetEpochSeed(_epoch uint64) ([32]byte, error) {
	return _Epoch.Contract.GetEpochSeed(&_Epoch.CallOpts, _epoch)
}

// GetEpochSeed is a free data retrieval call binding the contract method 0x2249084c.
//
// Solidity: function getEpochSeed(uint64 _epoch) view returns(bytes32)
func (_Epoch *EpochCallerSession) GetEpochSeed(_epoch uint64) ([32]byte, error) {
	return _Epoch.Contract.GetEpochSeed(&_Epoch.CallOpts, _epoch)
}

// SlotsInEpoch is a free data retrieval call binding the contract method 0xbfb3cbb9.
//
// Solidity: function slotsInEpoch() view returns(uint64)
func (_Epoch *EpochCaller) SlotsInEpoch(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _Epoch.contract.Call(opts, &out, "slotsInEpoch")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// SlotsInEpoch is a free data retrieval call binding the contract method 0xbfb3cbb9.
//
// Solidity: function slotsInEpoch() view returns(uint64)
func (_Epoch *EpochSession) SlotsInEpoch() (uint64, error) {
	return _Epoch.Contract.SlotsInEpoch(&_Epoch.CallOpts)
}

// SlotsInEpoch is a free data retrieval call binding the contract method 0xbfb3cbb9.
//
// Solidity: function slotsInEpoch() view returns(uint64)
func (_Epoch *EpochCallerSession) SlotsInEpoch() (uint64, error) {
	return _Epoch.Contract.SlotsInEpoch(&_Epoch.CallOpts)
}

// Check is a paid mutator transaction binding the contract method 0x919840ad.
//
// Solidity: function check() returns()
func (_Epoch *EpochTransactor) Check(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Epoch.contract.Transact(opts, "check")
}

// Check is a paid mutator transaction binding the contract method 0x919840ad.
//
// Solidity: function check() returns()
func (_Epoch *EpochSession) Check() (*types.Transaction, error) {
	return _Epoch.Contract.Check(&_Epoch.TransactOpts)
}

// Check is a paid mutator transaction binding the contract method 0x919840ad.
//
// Solidity: function check() returns()
func (_Epoch *EpochTransactorSession) Check() (*types.Transaction, error) {
	return _Epoch.Contract.Check(&_Epoch.TransactOpts)
}

// EpochSetEpochIterator is returned from FilterSetEpoch and is used to iterate over the raw logs and unpacked data for SetEpoch events raised by the Epoch contract.
type EpochSetEpochIterator struct {
	Event *EpochSetEpoch // Event containing the contract specifics and raw log

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
func (it *EpochSetEpochIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EpochSetEpoch)
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
		it.Event = new(EpochSetEpoch)
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
func (it *EpochSetEpochIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EpochSetEpochIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EpochSetEpoch represents a SetEpoch event raised by the Epoch contract.
type EpochSetEpoch struct {
	E   uint64
	Raw types.Log // Blockchain specific contextual infos
}

// FilterSetEpoch is a free log retrieval operation binding the contract event 0x9d4ccb161ea809df14334516cc3070025c80baddb8e8364afaca6a6fe31bfd75.
//
// Solidity: event SetEpoch(uint64 _e)
func (_Epoch *EpochFilterer) FilterSetEpoch(opts *bind.FilterOpts) (*EpochSetEpochIterator, error) {

	logs, sub, err := _Epoch.contract.FilterLogs(opts, "SetEpoch")
	if err != nil {
		return nil, err
	}
	return &EpochSetEpochIterator{contract: _Epoch.contract, event: "SetEpoch", logs: logs, sub: sub}, nil
}

// WatchSetEpoch is a free log subscription operation binding the contract event 0x9d4ccb161ea809df14334516cc3070025c80baddb8e8364afaca6a6fe31bfd75.
//
// Solidity: event SetEpoch(uint64 _e)
func (_Epoch *EpochFilterer) WatchSetEpoch(opts *bind.WatchOpts, sink chan<- *EpochSetEpoch) (event.Subscription, error) {

	logs, sub, err := _Epoch.contract.WatchLogs(opts, "SetEpoch")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EpochSetEpoch)
				if err := _Epoch.contract.UnpackLog(event, "SetEpoch", log); err != nil {
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
func (_Epoch *EpochFilterer) ParseSetEpoch(log types.Log) (*EpochSetEpoch, error) {
	event := new(EpochSetEpoch)
	if err := _Epoch.contract.UnpackLog(event, "SetEpoch", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
