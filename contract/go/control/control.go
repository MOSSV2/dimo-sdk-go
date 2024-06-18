// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package control

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

// ControlMetaData contains all meta data concerning the Control contract.
var ControlMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_b\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_a\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"_ep\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"_p\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"_d\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"_s\",\"type\":\"uint64\"}],\"name\":\"add\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"bank\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_a\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_money\",\"type\":\"uint256\"}],\"name\":\"checkNode\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_from\",\"type\":\"address\"},{\"internalType\":\"uint8\",\"name\":\"_typ\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"_m\",\"type\":\"uint256\"}],\"name\":\"lock\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_t\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_m\",\"type\":\"uint256\"}],\"name\":\"mint\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_from\",\"type\":\"address\"},{\"internalType\":\"uint8\",\"name\":\"_typ\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_m\",\"type\":\"uint256\"}],\"name\":\"punish\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_a\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_m\",\"type\":\"uint256\"}],\"name\":\"release\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_from\",\"type\":\"address\"},{\"internalType\":\"uint8\",\"name\":\"_typ\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"_m\",\"type\":\"uint256\"}],\"name\":\"unlock\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561000f575f80fd5b506040516123e63803806123e6833981810160405281019061003191906100d4565b805f806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550506100ff565b5f80fd5b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f6100a38261007a565b9050919050565b6100b381610099565b81146100bd575f80fd5b50565b5f815190506100ce816100aa565b92915050565b5f602082840312156100e9576100e8610076565b5b5f6100f6848285016100c0565b91505092915050565b6122da8061010c5f395ff3fe608060405234801561000f575f80fd5b5060043610610086575f3560e01c806340c10f191161005957806340c10f19146100fa578063738dddba1461011657806376cdb03b14610132578063a3f608691461015057610086565b80630357371d1461008a57806306f0b4f1146100a65780632165e20b146100c25780633acd28a9146100de575b5f80fd5b6100a4600480360381019061009f91906116ce565b61016c565b005b6100c060048036038101906100bb9190611742565b610370565b005b6100dc60048036038101906100d791906117a6565b6107a4565b005b6100f860048036038101906100f39190611833565b610bd5565b005b610114600480360381019061010f91906116ce565b610de2565b005b610130600480360381019061012b91906117a6565b610f70565b005b61013a6113a1565b60405161014791906118b9565b60405180910390f35b61016a600480360381019061016591906116ce565b6113c4565b005b5f8054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663693ec85e6040518163ffffffff1660e01b81526004016101c29061192c565b6020604051808303815f875af11580156101de573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610202919061195e565b73ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161461026f576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610266906119d3565b60405180910390fd5b5f8054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663693ec85e6040518163ffffffff1660e01b81526004016102c590611a3b565b6020604051808303815f875af11580156102e1573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610305919061195e565b73ffffffffffffffffffffffffffffffffffffffff16630357371d83836040518363ffffffff1660e01b815260040161033f929190611a68565b5f604051808303815f87803b158015610356575f80fd5b505af1158015610368573d5f803e3d5ffd5b505050505050565b60018360ff160361054e575f8054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663693ec85e6040518163ffffffff1660e01b81526004016103d190611ad9565b6020604051808303815f875af11580156103ed573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610411919061195e565b73ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16148061050a57505f8054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663693ec85e6040518163ffffffff1660e01b815260040161049b90611b41565b6020604051808303815f875af11580156104b7573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906104db919061195e565b73ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16145b610549576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161054090611ba9565b60405180910390fd5b61069d565b60038360ff1603610661575f8054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663693ec85e6040518163ffffffff1660e01b81526004016105af90611c11565b6020604051808303815f875af11580156105cb573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906105ef919061195e565b73ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161461065c576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161065390611c79565b60405180910390fd5b61069c565b6040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161069390611ce1565b60405180910390fd5b5b5f8054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663693ec85e6040518163ffffffff1660e01b81526004016106f390611d49565b6020604051808303815f875af115801561070f573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610733919061195e565b73ffffffffffffffffffffffffffffffffffffffff166306f0b4f1858585856040518563ffffffff1660e01b81526004016107719493929190611d76565b5f604051808303815f87803b158015610788575f80fd5b505af115801561079a573d5f803e3d5ffd5b5050505050505050565b60018260ff1603610982575f8054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663693ec85e6040518163ffffffff1660e01b815260040161080590611ad9565b6020604051808303815f875af1158015610821573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610845919061195e565b73ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16148061093e57505f8054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663693ec85e6040518163ffffffff1660e01b81526004016108cf90611b41565b6020604051808303815f875af11580156108eb573d5f803e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061090f919061195e565b73ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16145b61097d576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161097490611e03565b60405180910390fd5b610ad1565b60038260ff1603610a95575f8054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663693ec85e6040518163ffffffff1660e01b81526004016109e390611c11565b6020604051808303815f875af11580156109ff573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610a23919061195e565b73ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610a90576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610a8790611e6b565b60405180910390fd5b610ad0565b6040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610ac790611ce1565b60405180910390fd5b5b5f8054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663693ec85e6040518163ffffffff1660e01b8152600401610b2790611d49565b6020604051808303815f875af1158015610b43573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610b67919061195e565b73ffffffffffffffffffffffffffffffffffffffff16632165e20b8484846040518463ffffffff1660e01b8152600401610ba393929190611e89565b5f604051808303815f87803b158015610bba575f80fd5b505af1158015610bcc573d5f803e3d5ffd5b50505050505050565b5f8054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663693ec85e6040518163ffffffff1660e01b8152600401610c2b9061192c565b6020604051808303815f875af1158015610c47573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610c6b919061195e565b73ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610cd8576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610ccf90611f08565b60405180910390fd5b5f8054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663693ec85e6040518163ffffffff1660e01b8152600401610d2e90611a3b565b6020604051808303815f875af1158015610d4a573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610d6e919061195e565b73ffffffffffffffffffffffffffffffffffffffff16633acd28a986868686866040518663ffffffff1660e01b8152600401610dae959493929190611f35565b5f604051808303815f87803b158015610dc5575f80fd5b505af1158015610dd7573d5f803e3d5ffd5b505050505050505050565b5f8054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663693ec85e6040518163ffffffff1660e01b8152600401610e3890611a3b565b6020604051808303815f875af1158015610e54573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610e78919061195e565b73ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610ee5576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610edc90611fd0565b60405180910390fd5b5f8054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166340c10f1983836040518363ffffffff1660e01b8152600401610f3f929190611a68565b5f604051808303815f87803b158015610f56575f80fd5b505af1158015610f68573d5f803e3d5ffd5b505050505050565b60018260ff160361114e575f8054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663693ec85e6040518163ffffffff1660e01b8152600401610fd190611ad9565b6020604051808303815f875af1158015610fed573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190611011919061195e565b73ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16148061110a57505f8054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663693ec85e6040518163ffffffff1660e01b815260040161109b90611b41565b6020604051808303815f875af11580156110b7573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906110db919061195e565b73ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16145b611149576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161114090612038565b60405180910390fd5b61129d565b60038260ff1603611261575f8054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663693ec85e6040518163ffffffff1660e01b81526004016111af90611c11565b6020604051808303815f875af11580156111cb573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906111ef919061195e565b73ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161461125c576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401611253906120a0565b60405180910390fd5b61129c565b6040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161129390611ce1565b60405180910390fd5b5b5f8054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663693ec85e6040518163ffffffff1660e01b81526004016112f390611d49565b6020604051808303815f875af115801561130f573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190611333919061195e565b73ffffffffffffffffffffffffffffffffffffffff1663738dddba8484846040518463ffffffff1660e01b815260040161136f93929190611e89565b5f604051808303815f87803b158015611386575f80fd5b505af1158015611398573d5f803e3d5ffd5b50505050505050565b5f8054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b5f805f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663693ec85e6040518163ffffffff1660e01b815260040161141c90611d49565b6020604051808303815f875af1158015611438573d5f803e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061145c919061195e565b73ffffffffffffffffffffffffffffffffffffffff166361e728b18460016040518363ffffffff1660e01b8152600401611497929190612100565b602060405180830381865afa1580156114b2573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906114d6919061213b565b90505f8054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663693ec85e6040518163ffffffff1660e01b815260040161152e9061192c565b6020604051808303815f875af115801561154a573d5f803e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061156e919061195e565b73ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16036115f4578160076115ad9190612193565b8110156115ef576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016115e69061221e565b60405180910390fd5b611638565b81811015611637576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161162e90612286565b60405180910390fd5b5b505050565b5f80fd5b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f61166a82611641565b9050919050565b61167a81611660565b8114611684575f80fd5b50565b5f8135905061169581611671565b92915050565b5f819050919050565b6116ad8161169b565b81146116b7575f80fd5b50565b5f813590506116c8816116a4565b92915050565b5f80604083850312156116e4576116e361163d565b5b5f6116f185828601611687565b9250506020611702858286016116ba565b9150509250929050565b5f60ff82169050919050565b6117218161170c565b811461172b575f80fd5b50565b5f8135905061173c81611718565b92915050565b5f805f806080858703121561175a5761175961163d565b5b5f61176787828801611687565b94505060206117788782880161172e565b935050604061178987828801611687565b925050606061179a878288016116ba565b91505092959194509250565b5f805f606084860312156117bd576117bc61163d565b5b5f6117ca86828701611687565b93505060206117db8682870161172e565b92505060406117ec868287016116ba565b9150509250925092565b5f67ffffffffffffffff82169050919050565b611812816117f6565b811461181c575f80fd5b50565b5f8135905061182d81611809565b92915050565b5f805f805f60a0868803121561184c5761184b61163d565b5b5f61185988828901611687565b955050602061186a8882890161181f565b945050604061187b888289016116ba565b935050606061188c8882890161181f565b925050608061189d8882890161181f565b9150509295509295909350565b6118b381611660565b82525050565b5f6020820190506118cc5f8301846118aa565b92915050565b5f82825260208201905092915050565b7f70696563650000000000000000000000000000000000000000000000000000005f82015250565b5f6119166005836118d2565b9150611921826118e2565b602082019050919050565b5f6020820190508181035f8301526119438161190a565b9050919050565b5f8151905061195881611671565b92915050565b5f602082840312156119735761197261163d565b5b5f6119808482850161194a565b91505092915050565b7f696e762072656c656173650000000000000000000000000000000000000000005f82015250565b5f6119bd600b836118d2565b91506119c882611989565b602082019050919050565b5f6020820190508181035f8301526119ea816119b1565b9050919050565b7f72657761726400000000000000000000000000000000000000000000000000005f82015250565b5f611a256006836118d2565b9150611a30826119f1565b602082019050919050565b5f6020820190508181035f830152611a5281611a19565b9050919050565b611a628161169b565b82525050565b5f604082019050611a7b5f8301856118aa565b611a886020830184611a59565b9392505050565b7f6570726f6f6600000000000000000000000000000000000000000000000000005f82015250565b5f611ac36006836118d2565b9150611ace82611a8f565b602082019050919050565b5f6020820190508181035f830152611af081611ab7565b9050919050565b7f727370726f6f66000000000000000000000000000000000000000000000000005f82015250565b5f611b2b6007836118d2565b9150611b3682611af7565b602082019050919050565b5f6020820190508181035f830152611b5881611b1f565b9050919050565b7f696e762070756e697368310000000000000000000000000000000000000000005f82015250565b5f611b93600b836118d2565b9150611b9e82611b5f565b602082019050919050565b5f6020820190508181035f830152611bc081611b87565b9050919050565b7f73706163650000000000000000000000000000000000000000000000000000005f82015250565b5f611bfb6005836118d2565b9150611c0682611bc7565b602082019050919050565b5f6020820190508181035f830152611c2881611bef565b9050919050565b7f696e762070756e697368330000000000000000000000000000000000000000005f82015250565b5f611c63600b836118d2565b9150611c6e82611c2f565b602082019050919050565b5f6020820190508181035f830152611c9081611c57565b9050919050565b7f756e7375700000000000000000000000000000000000000000000000000000005f82015250565b5f611ccb6005836118d2565b9150611cd682611c97565b602082019050919050565b5f6020820190508181035f830152611cf881611cbf565b9050919050565b7f6e6f6465000000000000000000000000000000000000000000000000000000005f82015250565b5f611d336004836118d2565b9150611d3e82611cff565b602082019050919050565b5f6020820190508181035f830152611d6081611d27565b9050919050565b611d708161170c565b82525050565b5f608082019050611d895f8301876118aa565b611d966020830186611d67565b611da360408301856118aa565b611db06060830184611a59565b95945050505050565b7f696e7620756e6c6f636b310000000000000000000000000000000000000000005f82015250565b5f611ded600b836118d2565b9150611df882611db9565b602082019050919050565b5f6020820190508181035f830152611e1a81611de1565b9050919050565b7f696e7620756e6c6f636b330000000000000000000000000000000000000000005f82015250565b5f611e55600b836118d2565b9150611e6082611e21565b602082019050919050565b5f6020820190508181035f830152611e8281611e49565b9050919050565b5f606082019050611e9c5f8301866118aa565b611ea96020830185611d67565b611eb66040830184611a59565b949350505050565b7f696e7620616464000000000000000000000000000000000000000000000000005f82015250565b5f611ef26007836118d2565b9150611efd82611ebe565b602082019050919050565b5f6020820190508181035f830152611f1f81611ee6565b9050919050565b611f2f816117f6565b82525050565b5f60a082019050611f485f8301886118aa565b611f556020830187611f26565b611f626040830186611a59565b611f6f6060830185611f26565b611f7c6080830184611f26565b9695505050505050565b7f696e76206d696e740000000000000000000000000000000000000000000000005f82015250565b5f611fba6008836118d2565b9150611fc582611f86565b602082019050919050565b5f6020820190508181035f830152611fe781611fae565b9050919050565b7f696e76206c6f636b3100000000000000000000000000000000000000000000005f82015250565b5f6120226009836118d2565b915061202d82611fee565b602082019050919050565b5f6020820190508181035f83015261204f81612016565b9050919050565b7f696e76206c6f636b3300000000000000000000000000000000000000000000005f82015250565b5f61208a6009836118d2565b915061209582612056565b602082019050919050565b5f6020820190508181035f8301526120b78161207e565b9050919050565b5f819050919050565b5f819050919050565b5f6120ea6120e56120e0846120be565b6120c7565b61170c565b9050919050565b6120fa816120d0565b82525050565b5f6040820190506121135f8301856118aa565b61212060208301846120f1565b9392505050565b5f81519050612135816116a4565b92915050565b5f602082840312156121505761214f61163d565b5b5f61215d84828501612127565b91505092915050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b5f61219d8261169b565b91506121a88361169b565b92508282026121b68161169b565b915082820484148315176121cd576121cc612166565b5b5092915050565b7f696e73756620706c6564676537000000000000000000000000000000000000005f82015250565b5f612208600d836118d2565b9150612213826121d4565b602082019050919050565b5f6020820190508181035f830152612235816121fc565b9050919050565b7f696e73756620706c6564676500000000000000000000000000000000000000005f82015250565b5f612270600c836118d2565b915061227b8261223c565b602082019050919050565b5f6020820190508181035f83015261229d81612264565b905091905056fea26469706673582212200691bb95e2be21e51ff4425d7e24e09765211688e0661f3da6be96e59fc9207c64736f6c634300081a0033",
}

// ControlABI is the input ABI used to generate the binding from.
// Deprecated: Use ControlMetaData.ABI instead.
var ControlABI = ControlMetaData.ABI

// ControlBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ControlMetaData.Bin instead.
var ControlBin = ControlMetaData.Bin

// DeployControl deploys a new Ethereum contract, binding an instance of Control to it.
func DeployControl(auth *bind.TransactOpts, backend bind.ContractBackend, _b common.Address) (common.Address, *types.Transaction, *Control, error) {
	parsed, err := ControlMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ControlBin), backend, _b)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Control{ControlCaller: ControlCaller{contract: contract}, ControlTransactor: ControlTransactor{contract: contract}, ControlFilterer: ControlFilterer{contract: contract}}, nil
}

// Control is an auto generated Go binding around an Ethereum contract.
type Control struct {
	ControlCaller     // Read-only binding to the contract
	ControlTransactor // Write-only binding to the contract
	ControlFilterer   // Log filterer for contract events
}

// ControlCaller is an auto generated read-only Go binding around an Ethereum contract.
type ControlCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ControlTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ControlTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ControlFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ControlFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ControlSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ControlSession struct {
	Contract     *Control          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ControlCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ControlCallerSession struct {
	Contract *ControlCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// ControlTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ControlTransactorSession struct {
	Contract     *ControlTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// ControlRaw is an auto generated low-level Go binding around an Ethereum contract.
type ControlRaw struct {
	Contract *Control // Generic contract binding to access the raw methods on
}

// ControlCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ControlCallerRaw struct {
	Contract *ControlCaller // Generic read-only contract binding to access the raw methods on
}

// ControlTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ControlTransactorRaw struct {
	Contract *ControlTransactor // Generic write-only contract binding to access the raw methods on
}

// NewControl creates a new instance of Control, bound to a specific deployed contract.
func NewControl(address common.Address, backend bind.ContractBackend) (*Control, error) {
	contract, err := bindControl(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Control{ControlCaller: ControlCaller{contract: contract}, ControlTransactor: ControlTransactor{contract: contract}, ControlFilterer: ControlFilterer{contract: contract}}, nil
}

// NewControlCaller creates a new read-only instance of Control, bound to a specific deployed contract.
func NewControlCaller(address common.Address, caller bind.ContractCaller) (*ControlCaller, error) {
	contract, err := bindControl(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ControlCaller{contract: contract}, nil
}

// NewControlTransactor creates a new write-only instance of Control, bound to a specific deployed contract.
func NewControlTransactor(address common.Address, transactor bind.ContractTransactor) (*ControlTransactor, error) {
	contract, err := bindControl(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ControlTransactor{contract: contract}, nil
}

// NewControlFilterer creates a new log filterer instance of Control, bound to a specific deployed contract.
func NewControlFilterer(address common.Address, filterer bind.ContractFilterer) (*ControlFilterer, error) {
	contract, err := bindControl(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ControlFilterer{contract: contract}, nil
}

// bindControl binds a generic wrapper to an already deployed contract.
func bindControl(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ControlMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Control *ControlRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Control.Contract.ControlCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Control *ControlRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Control.Contract.ControlTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Control *ControlRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Control.Contract.ControlTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Control *ControlCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Control.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Control *ControlTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Control.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Control *ControlTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Control.Contract.contract.Transact(opts, method, params...)
}

// Bank is a free data retrieval call binding the contract method 0x76cdb03b.
//
// Solidity: function bank() view returns(address)
func (_Control *ControlCaller) Bank(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Control.contract.Call(opts, &out, "bank")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Bank is a free data retrieval call binding the contract method 0x76cdb03b.
//
// Solidity: function bank() view returns(address)
func (_Control *ControlSession) Bank() (common.Address, error) {
	return _Control.Contract.Bank(&_Control.CallOpts)
}

// Bank is a free data retrieval call binding the contract method 0x76cdb03b.
//
// Solidity: function bank() view returns(address)
func (_Control *ControlCallerSession) Bank() (common.Address, error) {
	return _Control.Contract.Bank(&_Control.CallOpts)
}

// Add is a paid mutator transaction binding the contract method 0x3acd28a9.
//
// Solidity: function add(address _a, uint64 _ep, uint256 _p, uint64 _d, uint64 _s) returns()
func (_Control *ControlTransactor) Add(opts *bind.TransactOpts, _a common.Address, _ep uint64, _p *big.Int, _d uint64, _s uint64) (*types.Transaction, error) {
	return _Control.contract.Transact(opts, "add", _a, _ep, _p, _d, _s)
}

// Add is a paid mutator transaction binding the contract method 0x3acd28a9.
//
// Solidity: function add(address _a, uint64 _ep, uint256 _p, uint64 _d, uint64 _s) returns()
func (_Control *ControlSession) Add(_a common.Address, _ep uint64, _p *big.Int, _d uint64, _s uint64) (*types.Transaction, error) {
	return _Control.Contract.Add(&_Control.TransactOpts, _a, _ep, _p, _d, _s)
}

// Add is a paid mutator transaction binding the contract method 0x3acd28a9.
//
// Solidity: function add(address _a, uint64 _ep, uint256 _p, uint64 _d, uint64 _s) returns()
func (_Control *ControlTransactorSession) Add(_a common.Address, _ep uint64, _p *big.Int, _d uint64, _s uint64) (*types.Transaction, error) {
	return _Control.Contract.Add(&_Control.TransactOpts, _a, _ep, _p, _d, _s)
}

// CheckNode is a paid mutator transaction binding the contract method 0xa3f60869.
//
// Solidity: function checkNode(address _a, uint256 _money) returns()
func (_Control *ControlTransactor) CheckNode(opts *bind.TransactOpts, _a common.Address, _money *big.Int) (*types.Transaction, error) {
	return _Control.contract.Transact(opts, "checkNode", _a, _money)
}

// CheckNode is a paid mutator transaction binding the contract method 0xa3f60869.
//
// Solidity: function checkNode(address _a, uint256 _money) returns()
func (_Control *ControlSession) CheckNode(_a common.Address, _money *big.Int) (*types.Transaction, error) {
	return _Control.Contract.CheckNode(&_Control.TransactOpts, _a, _money)
}

// CheckNode is a paid mutator transaction binding the contract method 0xa3f60869.
//
// Solidity: function checkNode(address _a, uint256 _money) returns()
func (_Control *ControlTransactorSession) CheckNode(_a common.Address, _money *big.Int) (*types.Transaction, error) {
	return _Control.Contract.CheckNode(&_Control.TransactOpts, _a, _money)
}

// Lock is a paid mutator transaction binding the contract method 0x738dddba.
//
// Solidity: function lock(address _from, uint8 _typ, uint256 _m) returns()
func (_Control *ControlTransactor) Lock(opts *bind.TransactOpts, _from common.Address, _typ uint8, _m *big.Int) (*types.Transaction, error) {
	return _Control.contract.Transact(opts, "lock", _from, _typ, _m)
}

// Lock is a paid mutator transaction binding the contract method 0x738dddba.
//
// Solidity: function lock(address _from, uint8 _typ, uint256 _m) returns()
func (_Control *ControlSession) Lock(_from common.Address, _typ uint8, _m *big.Int) (*types.Transaction, error) {
	return _Control.Contract.Lock(&_Control.TransactOpts, _from, _typ, _m)
}

// Lock is a paid mutator transaction binding the contract method 0x738dddba.
//
// Solidity: function lock(address _from, uint8 _typ, uint256 _m) returns()
func (_Control *ControlTransactorSession) Lock(_from common.Address, _typ uint8, _m *big.Int) (*types.Transaction, error) {
	return _Control.Contract.Lock(&_Control.TransactOpts, _from, _typ, _m)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address _t, uint256 _m) returns()
func (_Control *ControlTransactor) Mint(opts *bind.TransactOpts, _t common.Address, _m *big.Int) (*types.Transaction, error) {
	return _Control.contract.Transact(opts, "mint", _t, _m)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address _t, uint256 _m) returns()
func (_Control *ControlSession) Mint(_t common.Address, _m *big.Int) (*types.Transaction, error) {
	return _Control.Contract.Mint(&_Control.TransactOpts, _t, _m)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address _t, uint256 _m) returns()
func (_Control *ControlTransactorSession) Mint(_t common.Address, _m *big.Int) (*types.Transaction, error) {
	return _Control.Contract.Mint(&_Control.TransactOpts, _t, _m)
}

// Punish is a paid mutator transaction binding the contract method 0x06f0b4f1.
//
// Solidity: function punish(address _from, uint8 _typ, address _to, uint256 _m) returns()
func (_Control *ControlTransactor) Punish(opts *bind.TransactOpts, _from common.Address, _typ uint8, _to common.Address, _m *big.Int) (*types.Transaction, error) {
	return _Control.contract.Transact(opts, "punish", _from, _typ, _to, _m)
}

// Punish is a paid mutator transaction binding the contract method 0x06f0b4f1.
//
// Solidity: function punish(address _from, uint8 _typ, address _to, uint256 _m) returns()
func (_Control *ControlSession) Punish(_from common.Address, _typ uint8, _to common.Address, _m *big.Int) (*types.Transaction, error) {
	return _Control.Contract.Punish(&_Control.TransactOpts, _from, _typ, _to, _m)
}

// Punish is a paid mutator transaction binding the contract method 0x06f0b4f1.
//
// Solidity: function punish(address _from, uint8 _typ, address _to, uint256 _m) returns()
func (_Control *ControlTransactorSession) Punish(_from common.Address, _typ uint8, _to common.Address, _m *big.Int) (*types.Transaction, error) {
	return _Control.Contract.Punish(&_Control.TransactOpts, _from, _typ, _to, _m)
}

// Release is a paid mutator transaction binding the contract method 0x0357371d.
//
// Solidity: function release(address _a, uint256 _m) returns()
func (_Control *ControlTransactor) Release(opts *bind.TransactOpts, _a common.Address, _m *big.Int) (*types.Transaction, error) {
	return _Control.contract.Transact(opts, "release", _a, _m)
}

// Release is a paid mutator transaction binding the contract method 0x0357371d.
//
// Solidity: function release(address _a, uint256 _m) returns()
func (_Control *ControlSession) Release(_a common.Address, _m *big.Int) (*types.Transaction, error) {
	return _Control.Contract.Release(&_Control.TransactOpts, _a, _m)
}

// Release is a paid mutator transaction binding the contract method 0x0357371d.
//
// Solidity: function release(address _a, uint256 _m) returns()
func (_Control *ControlTransactorSession) Release(_a common.Address, _m *big.Int) (*types.Transaction, error) {
	return _Control.Contract.Release(&_Control.TransactOpts, _a, _m)
}

// Unlock is a paid mutator transaction binding the contract method 0x2165e20b.
//
// Solidity: function unlock(address _from, uint8 _typ, uint256 _m) returns()
func (_Control *ControlTransactor) Unlock(opts *bind.TransactOpts, _from common.Address, _typ uint8, _m *big.Int) (*types.Transaction, error) {
	return _Control.contract.Transact(opts, "unlock", _from, _typ, _m)
}

// Unlock is a paid mutator transaction binding the contract method 0x2165e20b.
//
// Solidity: function unlock(address _from, uint8 _typ, uint256 _m) returns()
func (_Control *ControlSession) Unlock(_from common.Address, _typ uint8, _m *big.Int) (*types.Transaction, error) {
	return _Control.Contract.Unlock(&_Control.TransactOpts, _from, _typ, _m)
}

// Unlock is a paid mutator transaction binding the contract method 0x2165e20b.
//
// Solidity: function unlock(address _from, uint8 _typ, uint256 _m) returns()
func (_Control *ControlTransactorSession) Unlock(_from common.Address, _typ uint8, _m *big.Int) (*types.Transaction, error) {
	return _Control.Contract.Unlock(&_Control.TransactOpts, _from, _typ, _m)
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

// INodeMetaData contains all meta data concerning the INode contract.
var INodeMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_a\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"_type\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_money\",\"type\":\"uint256\"}],\"name\":\"Pledge\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_a\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"_typ\",\"type\":\"uint8\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_money\",\"type\":\"uint256\"}],\"name\":\"Punish\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"_type\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_money\",\"type\":\"uint256\"}],\"name\":\"Set\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_a\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"_type\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_money\",\"type\":\"uint256\"}],\"name\":\"Withdraw\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_a\",\"type\":\"address\"},{\"internalType\":\"uint8\",\"name\":\"_type\",\"type\":\"uint8\"}],\"name\":\"check\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_a\",\"type\":\"address\"},{\"internalType\":\"uint8\",\"name\":\"_type\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"_m\",\"type\":\"uint256\"}],\"name\":\"lock\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_a\",\"type\":\"address\"},{\"internalType\":\"uint8\",\"name\":\"_type\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_m\",\"type\":\"uint256\"}],\"name\":\"punish\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_a\",\"type\":\"address\"},{\"internalType\":\"uint8\",\"name\":\"_type\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"_m\",\"type\":\"uint256\"}],\"name\":\"unlock\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// INodeABI is the input ABI used to generate the binding from.
// Deprecated: Use INodeMetaData.ABI instead.
var INodeABI = INodeMetaData.ABI

// INode is an auto generated Go binding around an Ethereum contract.
type INode struct {
	INodeCaller     // Read-only binding to the contract
	INodeTransactor // Write-only binding to the contract
	INodeFilterer   // Log filterer for contract events
}

// INodeCaller is an auto generated read-only Go binding around an Ethereum contract.
type INodeCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// INodeTransactor is an auto generated write-only Go binding around an Ethereum contract.
type INodeTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// INodeFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type INodeFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// INodeSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type INodeSession struct {
	Contract     *INode            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// INodeCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type INodeCallerSession struct {
	Contract *INodeCaller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// INodeTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type INodeTransactorSession struct {
	Contract     *INodeTransactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// INodeRaw is an auto generated low-level Go binding around an Ethereum contract.
type INodeRaw struct {
	Contract *INode // Generic contract binding to access the raw methods on
}

// INodeCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type INodeCallerRaw struct {
	Contract *INodeCaller // Generic read-only contract binding to access the raw methods on
}

// INodeTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type INodeTransactorRaw struct {
	Contract *INodeTransactor // Generic write-only contract binding to access the raw methods on
}

// NewINode creates a new instance of INode, bound to a specific deployed contract.
func NewINode(address common.Address, backend bind.ContractBackend) (*INode, error) {
	contract, err := bindINode(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &INode{INodeCaller: INodeCaller{contract: contract}, INodeTransactor: INodeTransactor{contract: contract}, INodeFilterer: INodeFilterer{contract: contract}}, nil
}

// NewINodeCaller creates a new read-only instance of INode, bound to a specific deployed contract.
func NewINodeCaller(address common.Address, caller bind.ContractCaller) (*INodeCaller, error) {
	contract, err := bindINode(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &INodeCaller{contract: contract}, nil
}

// NewINodeTransactor creates a new write-only instance of INode, bound to a specific deployed contract.
func NewINodeTransactor(address common.Address, transactor bind.ContractTransactor) (*INodeTransactor, error) {
	contract, err := bindINode(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &INodeTransactor{contract: contract}, nil
}

// NewINodeFilterer creates a new log filterer instance of INode, bound to a specific deployed contract.
func NewINodeFilterer(address common.Address, filterer bind.ContractFilterer) (*INodeFilterer, error) {
	contract, err := bindINode(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &INodeFilterer{contract: contract}, nil
}

// bindINode binds a generic wrapper to an already deployed contract.
func bindINode(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := INodeMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_INode *INodeRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _INode.Contract.INodeCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_INode *INodeRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _INode.Contract.INodeTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_INode *INodeRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _INode.Contract.INodeTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_INode *INodeCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _INode.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_INode *INodeTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _INode.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_INode *INodeTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _INode.Contract.contract.Transact(opts, method, params...)
}

// Check is a free data retrieval call binding the contract method 0x61e728b1.
//
// Solidity: function check(address _a, uint8 _type) view returns(uint256)
func (_INode *INodeCaller) Check(opts *bind.CallOpts, _a common.Address, _type uint8) (*big.Int, error) {
	var out []interface{}
	err := _INode.contract.Call(opts, &out, "check", _a, _type)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Check is a free data retrieval call binding the contract method 0x61e728b1.
//
// Solidity: function check(address _a, uint8 _type) view returns(uint256)
func (_INode *INodeSession) Check(_a common.Address, _type uint8) (*big.Int, error) {
	return _INode.Contract.Check(&_INode.CallOpts, _a, _type)
}

// Check is a free data retrieval call binding the contract method 0x61e728b1.
//
// Solidity: function check(address _a, uint8 _type) view returns(uint256)
func (_INode *INodeCallerSession) Check(_a common.Address, _type uint8) (*big.Int, error) {
	return _INode.Contract.Check(&_INode.CallOpts, _a, _type)
}

// Lock is a paid mutator transaction binding the contract method 0x738dddba.
//
// Solidity: function lock(address _a, uint8 _type, uint256 _m) returns()
func (_INode *INodeTransactor) Lock(opts *bind.TransactOpts, _a common.Address, _type uint8, _m *big.Int) (*types.Transaction, error) {
	return _INode.contract.Transact(opts, "lock", _a, _type, _m)
}

// Lock is a paid mutator transaction binding the contract method 0x738dddba.
//
// Solidity: function lock(address _a, uint8 _type, uint256 _m) returns()
func (_INode *INodeSession) Lock(_a common.Address, _type uint8, _m *big.Int) (*types.Transaction, error) {
	return _INode.Contract.Lock(&_INode.TransactOpts, _a, _type, _m)
}

// Lock is a paid mutator transaction binding the contract method 0x738dddba.
//
// Solidity: function lock(address _a, uint8 _type, uint256 _m) returns()
func (_INode *INodeTransactorSession) Lock(_a common.Address, _type uint8, _m *big.Int) (*types.Transaction, error) {
	return _INode.Contract.Lock(&_INode.TransactOpts, _a, _type, _m)
}

// Punish is a paid mutator transaction binding the contract method 0x06f0b4f1.
//
// Solidity: function punish(address _a, uint8 _type, address _to, uint256 _m) returns()
func (_INode *INodeTransactor) Punish(opts *bind.TransactOpts, _a common.Address, _type uint8, _to common.Address, _m *big.Int) (*types.Transaction, error) {
	return _INode.contract.Transact(opts, "punish", _a, _type, _to, _m)
}

// Punish is a paid mutator transaction binding the contract method 0x06f0b4f1.
//
// Solidity: function punish(address _a, uint8 _type, address _to, uint256 _m) returns()
func (_INode *INodeSession) Punish(_a common.Address, _type uint8, _to common.Address, _m *big.Int) (*types.Transaction, error) {
	return _INode.Contract.Punish(&_INode.TransactOpts, _a, _type, _to, _m)
}

// Punish is a paid mutator transaction binding the contract method 0x06f0b4f1.
//
// Solidity: function punish(address _a, uint8 _type, address _to, uint256 _m) returns()
func (_INode *INodeTransactorSession) Punish(_a common.Address, _type uint8, _to common.Address, _m *big.Int) (*types.Transaction, error) {
	return _INode.Contract.Punish(&_INode.TransactOpts, _a, _type, _to, _m)
}

// Unlock is a paid mutator transaction binding the contract method 0x2165e20b.
//
// Solidity: function unlock(address _a, uint8 _type, uint256 _m) returns()
func (_INode *INodeTransactor) Unlock(opts *bind.TransactOpts, _a common.Address, _type uint8, _m *big.Int) (*types.Transaction, error) {
	return _INode.contract.Transact(opts, "unlock", _a, _type, _m)
}

// Unlock is a paid mutator transaction binding the contract method 0x2165e20b.
//
// Solidity: function unlock(address _a, uint8 _type, uint256 _m) returns()
func (_INode *INodeSession) Unlock(_a common.Address, _type uint8, _m *big.Int) (*types.Transaction, error) {
	return _INode.Contract.Unlock(&_INode.TransactOpts, _a, _type, _m)
}

// Unlock is a paid mutator transaction binding the contract method 0x2165e20b.
//
// Solidity: function unlock(address _a, uint8 _type, uint256 _m) returns()
func (_INode *INodeTransactorSession) Unlock(_a common.Address, _type uint8, _m *big.Int) (*types.Transaction, error) {
	return _INode.Contract.Unlock(&_INode.TransactOpts, _a, _type, _m)
}

// INodePledgeIterator is returned from FilterPledge and is used to iterate over the raw logs and unpacked data for Pledge events raised by the INode contract.
type INodePledgeIterator struct {
	Event *INodePledge // Event containing the contract specifics and raw log

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
func (it *INodePledgeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(INodePledge)
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
		it.Event = new(INodePledge)
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
func (it *INodePledgeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *INodePledgeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// INodePledge represents a Pledge event raised by the INode contract.
type INodePledge struct {
	A     common.Address
	Type  uint8
	Money *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterPledge is a free log retrieval operation binding the contract event 0xedec38b4b62433445c926815411650aa7d417efa9da307f6aa99e69e6f4493ee.
//
// Solidity: event Pledge(address indexed _a, uint8 _type, uint256 _money)
func (_INode *INodeFilterer) FilterPledge(opts *bind.FilterOpts, _a []common.Address) (*INodePledgeIterator, error) {

	var _aRule []interface{}
	for _, _aItem := range _a {
		_aRule = append(_aRule, _aItem)
	}

	logs, sub, err := _INode.contract.FilterLogs(opts, "Pledge", _aRule)
	if err != nil {
		return nil, err
	}
	return &INodePledgeIterator{contract: _INode.contract, event: "Pledge", logs: logs, sub: sub}, nil
}

// WatchPledge is a free log subscription operation binding the contract event 0xedec38b4b62433445c926815411650aa7d417efa9da307f6aa99e69e6f4493ee.
//
// Solidity: event Pledge(address indexed _a, uint8 _type, uint256 _money)
func (_INode *INodeFilterer) WatchPledge(opts *bind.WatchOpts, sink chan<- *INodePledge, _a []common.Address) (event.Subscription, error) {

	var _aRule []interface{}
	for _, _aItem := range _a {
		_aRule = append(_aRule, _aItem)
	}

	logs, sub, err := _INode.contract.WatchLogs(opts, "Pledge", _aRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(INodePledge)
				if err := _INode.contract.UnpackLog(event, "Pledge", log); err != nil {
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

// ParsePledge is a log parse operation binding the contract event 0xedec38b4b62433445c926815411650aa7d417efa9da307f6aa99e69e6f4493ee.
//
// Solidity: event Pledge(address indexed _a, uint8 _type, uint256 _money)
func (_INode *INodeFilterer) ParsePledge(log types.Log) (*INodePledge, error) {
	event := new(INodePledge)
	if err := _INode.contract.UnpackLog(event, "Pledge", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// INodePunishIterator is returned from FilterPunish and is used to iterate over the raw logs and unpacked data for Punish events raised by the INode contract.
type INodePunishIterator struct {
	Event *INodePunish // Event containing the contract specifics and raw log

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
func (it *INodePunishIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(INodePunish)
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
		it.Event = new(INodePunish)
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
func (it *INodePunishIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *INodePunishIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// INodePunish represents a Punish event raised by the INode contract.
type INodePunish struct {
	A     common.Address
	Typ   uint8
	To    common.Address
	Money *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterPunish is a free log retrieval operation binding the contract event 0xc6ef96923e613455515c6723eff1723445b22427fe442e8bf742e9d29b4b3c32.
//
// Solidity: event Punish(address indexed _a, uint8 _typ, address indexed _to, uint256 _money)
func (_INode *INodeFilterer) FilterPunish(opts *bind.FilterOpts, _a []common.Address, _to []common.Address) (*INodePunishIterator, error) {

	var _aRule []interface{}
	for _, _aItem := range _a {
		_aRule = append(_aRule, _aItem)
	}

	var _toRule []interface{}
	for _, _toItem := range _to {
		_toRule = append(_toRule, _toItem)
	}

	logs, sub, err := _INode.contract.FilterLogs(opts, "Punish", _aRule, _toRule)
	if err != nil {
		return nil, err
	}
	return &INodePunishIterator{contract: _INode.contract, event: "Punish", logs: logs, sub: sub}, nil
}

// WatchPunish is a free log subscription operation binding the contract event 0xc6ef96923e613455515c6723eff1723445b22427fe442e8bf742e9d29b4b3c32.
//
// Solidity: event Punish(address indexed _a, uint8 _typ, address indexed _to, uint256 _money)
func (_INode *INodeFilterer) WatchPunish(opts *bind.WatchOpts, sink chan<- *INodePunish, _a []common.Address, _to []common.Address) (event.Subscription, error) {

	var _aRule []interface{}
	for _, _aItem := range _a {
		_aRule = append(_aRule, _aItem)
	}

	var _toRule []interface{}
	for _, _toItem := range _to {
		_toRule = append(_toRule, _toItem)
	}

	logs, sub, err := _INode.contract.WatchLogs(opts, "Punish", _aRule, _toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(INodePunish)
				if err := _INode.contract.UnpackLog(event, "Punish", log); err != nil {
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

// ParsePunish is a log parse operation binding the contract event 0xc6ef96923e613455515c6723eff1723445b22427fe442e8bf742e9d29b4b3c32.
//
// Solidity: event Punish(address indexed _a, uint8 _typ, address indexed _to, uint256 _money)
func (_INode *INodeFilterer) ParsePunish(log types.Log) (*INodePunish, error) {
	event := new(INodePunish)
	if err := _INode.contract.UnpackLog(event, "Punish", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// INodeSetIterator is returned from FilterSet and is used to iterate over the raw logs and unpacked data for Set events raised by the INode contract.
type INodeSetIterator struct {
	Event *INodeSet // Event containing the contract specifics and raw log

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
func (it *INodeSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(INodeSet)
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
		it.Event = new(INodeSet)
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
func (it *INodeSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *INodeSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// INodeSet represents a Set event raised by the INode contract.
type INodeSet struct {
	Type  uint8
	Money *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterSet is a free log retrieval operation binding the contract event 0xc4b70ab905e9fd7aab427fb9e73cae1480cfdc41c22053b20745349a7ef67881.
//
// Solidity: event Set(uint8 _type, uint256 _money)
func (_INode *INodeFilterer) FilterSet(opts *bind.FilterOpts) (*INodeSetIterator, error) {

	logs, sub, err := _INode.contract.FilterLogs(opts, "Set")
	if err != nil {
		return nil, err
	}
	return &INodeSetIterator{contract: _INode.contract, event: "Set", logs: logs, sub: sub}, nil
}

// WatchSet is a free log subscription operation binding the contract event 0xc4b70ab905e9fd7aab427fb9e73cae1480cfdc41c22053b20745349a7ef67881.
//
// Solidity: event Set(uint8 _type, uint256 _money)
func (_INode *INodeFilterer) WatchSet(opts *bind.WatchOpts, sink chan<- *INodeSet) (event.Subscription, error) {

	logs, sub, err := _INode.contract.WatchLogs(opts, "Set")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(INodeSet)
				if err := _INode.contract.UnpackLog(event, "Set", log); err != nil {
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

// ParseSet is a log parse operation binding the contract event 0xc4b70ab905e9fd7aab427fb9e73cae1480cfdc41c22053b20745349a7ef67881.
//
// Solidity: event Set(uint8 _type, uint256 _money)
func (_INode *INodeFilterer) ParseSet(log types.Log) (*INodeSet, error) {
	event := new(INodeSet)
	if err := _INode.contract.UnpackLog(event, "Set", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// INodeWithdrawIterator is returned from FilterWithdraw and is used to iterate over the raw logs and unpacked data for Withdraw events raised by the INode contract.
type INodeWithdrawIterator struct {
	Event *INodeWithdraw // Event containing the contract specifics and raw log

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
func (it *INodeWithdrawIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(INodeWithdraw)
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
		it.Event = new(INodeWithdraw)
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
func (it *INodeWithdrawIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *INodeWithdrawIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// INodeWithdraw represents a Withdraw event raised by the INode contract.
type INodeWithdraw struct {
	A     common.Address
	Type  uint8
	Money *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterWithdraw is a free log retrieval operation binding the contract event 0xe3054ec6f352b09a86839beeeceb27ed4684f9164eb61f4a5d6d159ee8789d74.
//
// Solidity: event Withdraw(address indexed _a, uint8 _type, uint256 _money)
func (_INode *INodeFilterer) FilterWithdraw(opts *bind.FilterOpts, _a []common.Address) (*INodeWithdrawIterator, error) {

	var _aRule []interface{}
	for _, _aItem := range _a {
		_aRule = append(_aRule, _aItem)
	}

	logs, sub, err := _INode.contract.FilterLogs(opts, "Withdraw", _aRule)
	if err != nil {
		return nil, err
	}
	return &INodeWithdrawIterator{contract: _INode.contract, event: "Withdraw", logs: logs, sub: sub}, nil
}

// WatchWithdraw is a free log subscription operation binding the contract event 0xe3054ec6f352b09a86839beeeceb27ed4684f9164eb61f4a5d6d159ee8789d74.
//
// Solidity: event Withdraw(address indexed _a, uint8 _type, uint256 _money)
func (_INode *INodeFilterer) WatchWithdraw(opts *bind.WatchOpts, sink chan<- *INodeWithdraw, _a []common.Address) (event.Subscription, error) {

	var _aRule []interface{}
	for _, _aItem := range _a {
		_aRule = append(_aRule, _aItem)
	}

	logs, sub, err := _INode.contract.WatchLogs(opts, "Withdraw", _aRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(INodeWithdraw)
				if err := _INode.contract.UnpackLog(event, "Withdraw", log); err != nil {
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

// ParseWithdraw is a log parse operation binding the contract event 0xe3054ec6f352b09a86839beeeceb27ed4684f9164eb61f4a5d6d159ee8789d74.
//
// Solidity: event Withdraw(address indexed _a, uint8 _type, uint256 _money)
func (_INode *INodeFilterer) ParseWithdraw(log types.Log) (*INodeWithdraw, error) {
	event := new(INodeWithdraw)
	if err := _INode.contract.UnpackLog(event, "Withdraw", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IRewardMetaData contains all meta data concerning the IReward contract.
var IRewardMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_a\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"_ep\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"_p\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"_d\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"_s\",\"type\":\"uint64\"}],\"name\":\"add\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_a\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_m\",\"type\":\"uint256\"}],\"name\":\"release\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// IRewardABI is the input ABI used to generate the binding from.
// Deprecated: Use IRewardMetaData.ABI instead.
var IRewardABI = IRewardMetaData.ABI

// IReward is an auto generated Go binding around an Ethereum contract.
type IReward struct {
	IRewardCaller     // Read-only binding to the contract
	IRewardTransactor // Write-only binding to the contract
	IRewardFilterer   // Log filterer for contract events
}

// IRewardCaller is an auto generated read-only Go binding around an Ethereum contract.
type IRewardCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IRewardTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IRewardTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IRewardFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IRewardFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IRewardSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IRewardSession struct {
	Contract     *IReward          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IRewardCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IRewardCallerSession struct {
	Contract *IRewardCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// IRewardTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IRewardTransactorSession struct {
	Contract     *IRewardTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// IRewardRaw is an auto generated low-level Go binding around an Ethereum contract.
type IRewardRaw struct {
	Contract *IReward // Generic contract binding to access the raw methods on
}

// IRewardCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IRewardCallerRaw struct {
	Contract *IRewardCaller // Generic read-only contract binding to access the raw methods on
}

// IRewardTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IRewardTransactorRaw struct {
	Contract *IRewardTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIReward creates a new instance of IReward, bound to a specific deployed contract.
func NewIReward(address common.Address, backend bind.ContractBackend) (*IReward, error) {
	contract, err := bindIReward(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IReward{IRewardCaller: IRewardCaller{contract: contract}, IRewardTransactor: IRewardTransactor{contract: contract}, IRewardFilterer: IRewardFilterer{contract: contract}}, nil
}

// NewIRewardCaller creates a new read-only instance of IReward, bound to a specific deployed contract.
func NewIRewardCaller(address common.Address, caller bind.ContractCaller) (*IRewardCaller, error) {
	contract, err := bindIReward(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IRewardCaller{contract: contract}, nil
}

// NewIRewardTransactor creates a new write-only instance of IReward, bound to a specific deployed contract.
func NewIRewardTransactor(address common.Address, transactor bind.ContractTransactor) (*IRewardTransactor, error) {
	contract, err := bindIReward(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IRewardTransactor{contract: contract}, nil
}

// NewIRewardFilterer creates a new log filterer instance of IReward, bound to a specific deployed contract.
func NewIRewardFilterer(address common.Address, filterer bind.ContractFilterer) (*IRewardFilterer, error) {
	contract, err := bindIReward(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IRewardFilterer{contract: contract}, nil
}

// bindIReward binds a generic wrapper to an already deployed contract.
func bindIReward(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := IRewardMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IReward *IRewardRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IReward.Contract.IRewardCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IReward *IRewardRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IReward.Contract.IRewardTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IReward *IRewardRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IReward.Contract.IRewardTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IReward *IRewardCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IReward.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IReward *IRewardTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IReward.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IReward *IRewardTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IReward.Contract.contract.Transact(opts, method, params...)
}

// Add is a paid mutator transaction binding the contract method 0x3acd28a9.
//
// Solidity: function add(address _a, uint64 _ep, uint256 _p, uint64 _d, uint64 _s) returns()
func (_IReward *IRewardTransactor) Add(opts *bind.TransactOpts, _a common.Address, _ep uint64, _p *big.Int, _d uint64, _s uint64) (*types.Transaction, error) {
	return _IReward.contract.Transact(opts, "add", _a, _ep, _p, _d, _s)
}

// Add is a paid mutator transaction binding the contract method 0x3acd28a9.
//
// Solidity: function add(address _a, uint64 _ep, uint256 _p, uint64 _d, uint64 _s) returns()
func (_IReward *IRewardSession) Add(_a common.Address, _ep uint64, _p *big.Int, _d uint64, _s uint64) (*types.Transaction, error) {
	return _IReward.Contract.Add(&_IReward.TransactOpts, _a, _ep, _p, _d, _s)
}

// Add is a paid mutator transaction binding the contract method 0x3acd28a9.
//
// Solidity: function add(address _a, uint64 _ep, uint256 _p, uint64 _d, uint64 _s) returns()
func (_IReward *IRewardTransactorSession) Add(_a common.Address, _ep uint64, _p *big.Int, _d uint64, _s uint64) (*types.Transaction, error) {
	return _IReward.Contract.Add(&_IReward.TransactOpts, _a, _ep, _p, _d, _s)
}

// Release is a paid mutator transaction binding the contract method 0x0357371d.
//
// Solidity: function release(address _a, uint256 _m) returns()
func (_IReward *IRewardTransactor) Release(opts *bind.TransactOpts, _a common.Address, _m *big.Int) (*types.Transaction, error) {
	return _IReward.contract.Transact(opts, "release", _a, _m)
}

// Release is a paid mutator transaction binding the contract method 0x0357371d.
//
// Solidity: function release(address _a, uint256 _m) returns()
func (_IReward *IRewardSession) Release(_a common.Address, _m *big.Int) (*types.Transaction, error) {
	return _IReward.Contract.Release(&_IReward.TransactOpts, _a, _m)
}

// Release is a paid mutator transaction binding the contract method 0x0357371d.
//
// Solidity: function release(address _a, uint256 _m) returns()
func (_IReward *IRewardTransactorSession) Release(_a common.Address, _m *big.Int) (*types.Transaction, error) {
	return _IReward.Contract.Release(&_IReward.TransactOpts, _a, _m)
}
