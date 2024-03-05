// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package node

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

// INodePledgeInfo is an auto generated low-level Go binding around an user-defined struct.
type INodePledgeInfo struct {
	Time  *big.Int
	Value *big.Int
}

// NodeMetaData contains all meta data concerning the Node contract.
var NodeMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_b\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_a\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"_type\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_money\",\"type\":\"uint256\"}],\"name\":\"Pledge\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_a\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"_typ\",\"type\":\"uint8\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_money\",\"type\":\"uint256\"}],\"name\":\"Punish\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"_type\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_money\",\"type\":\"uint256\"}],\"name\":\"SetPledge\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_a\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"_type\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_money\",\"type\":\"uint256\"}],\"name\":\"Withdraw\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"bank\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_a\",\"type\":\"address\"},{\"internalType\":\"uint8\",\"name\":\"_type\",\"type\":\"uint8\"}],\"name\":\"check\",\"outputs\":[],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"curEpoch\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_detail\",\"type\":\"bytes\"}],\"name\":\"declare\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"_type\",\"type\":\"uint8\"}],\"name\":\"getMinPledge\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_a\",\"type\":\"address\"},{\"internalType\":\"uint8\",\"name\":\"_type\",\"type\":\"uint8\"}],\"name\":\"getPledge\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"time\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"internalType\":\"structINode.PledgeInfo\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_a\",\"type\":\"address\"},{\"internalType\":\"uint8\",\"name\":\"_type\",\"type\":\"uint8\"}],\"name\":\"getPledgeValue\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"}],\"name\":\"migrate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"_typ\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"_m\",\"type\":\"uint256\"}],\"name\":\"pledge\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_from\",\"type\":\"address\"},{\"internalType\":\"uint8\",\"name\":\"_typ\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_m\",\"type\":\"uint256\"}],\"name\":\"punish\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"_type\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"_money\",\"type\":\"uint256\"}],\"name\":\"setPledge\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"_typ\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"_m\",\"type\":\"uint256\"}],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405260646002553480156200001657600080fd5b5060405162001fa738038062001fa783398181016040528101906200003c9190620001da565b6200005c62000050620000a460201b60201c565b620000ac60201b60201c565b80600160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550506200020c565b600033905090565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050816000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508173ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a35050565b600080fd5b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000620001a28262000175565b9050919050565b620001b48162000195565b8114620001c057600080fd5b50565b600081519050620001d481620001a9565b92915050565b600060208284031215620001f357620001f262000170565b5b60006200020384828501620001c3565b91505092915050565b611d8b806200021c6000396000f3fe608060405234801561001057600080fd5b50600436106100f55760003560e01c806390d8cf0711610097578063d888731511610066578063d888731514610248578063e3fdacb614610266578063e6990adf14610296578063f2fde38b146102b2576100f5565b806390d8cf07146101c45780639708cab3146101e057806398c365d5146101fc578063ce5494bb1461022c576100f5565b806361e728b1116100d357806361e728b114610162578063715018a61461017e57806376cdb03b146101885780638da5cb5b146101a6576100f5565b806306f0b4f1146100fa5780633f489914146101165780634339ceca14610132575b600080fd5b610114600480360381019061010f91906111f3565b6102ce565b005b610130600480360381019061012b919061125a565b610543565b005b61014c6004803603810190610147919061129a565b610766565b60405161015991906112e9565b60405180910390f35b61017c6004803603810190610177919061129a565b6107ca565b005b610186610883565b005b610190610897565b60405161019d9190611313565b60405180910390f35b6101ae6108bd565b6040516101bb9190611313565b60405180910390f35b6101de60048036038101906101d99190611474565b6108e6565b005b6101fa60048036038101906101f5919061125a565b610939565b005b610216600480360381019061021191906114bd565b610b74565b60405161022391906112e9565b60405180910390f35b610246600480360381019061024191906114ea565b610b97565b005b610250610c2f565b60405161025d919061153a565b60405180910390f35b610280600480360381019061027b919061129a565b610c49565b60405161028d9190611593565b60405180910390f35b6102b060048036038101906102ab919061125a565b610d2c565b005b6102cc60048036038101906102c791906114ea565b610d8f565b005b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663693ec85e6040518163ffffffff1660e01b81526004016103279061160b565b6020604051808303816000875af1158015610346573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061036a9190611640565b73ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16146103d7576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016103ce906116b9565b60405180910390fd5b80600560008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008560ff1660ff16815260200190815260200160002060010160008282546104409190611708565b92505081905550600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166376890c5883836040518363ffffffff1660e01b81526004016104a492919061173c565b600060405180830381600087803b1580156104be57600080fd5b505af11580156104d2573d6000803e3d6000fd5b505050508173ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff167fc6ef96923e613455515c6723eff1723445b22427fe442e8bf742e9d29b4b3c328584604051610535929190611774565b60405180910390a350505050565b61054b610e12565b600160149054906101000a900467ffffffffffffffff1667ffffffffffffffff16600254600560003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008560ff1660ff168152602001908152602001600020600001546105d3919061179d565b10610613576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161060a9061181d565b60405180910390fd5b80600560003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008460ff1660ff168152602001908152602001600020600101600082825461067c9190611708565b92505081905550600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166376890c5833836040518363ffffffff1660e01b81526004016106e092919061173c565b600060405180830381600087803b1580156106fa57600080fd5b505af115801561070e573d6000803e3d6000fd5b505050503373ffffffffffffffffffffffffffffffffffffffff167fe3054ec6f352b09a86839beeeceb27ed4684f9164eb61f4a5d6d159ee8789d74838360405161075a929190611774565b60405180910390a25050565b6000600560008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008360ff1660ff16815260200190815260200160002060010154905092915050565b600360008260ff1660ff16815260200190815260200160002054600560008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008360ff1660ff16815260200190815260200160002060010154101561087f576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161087690611889565b60405180910390fd5b5050565b61088b610fae565b610895600061102c565b565b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b80600460003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060000190816109359190611ac0565b5050565b610941610e12565b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663e888891533836040518363ffffffff1660e01b815260040161099e92919061173c565b600060405180830381600087803b1580156109b857600080fd5b505af11580156109cc573d6000803e3d6000fd5b505050506000600560003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008460ff1660ff1681526020019081526020016000206000015403610ab057600160149054906101000a900467ffffffffffffffff1667ffffffffffffffff16600560003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008460ff1660ff168152602001908152602001600020600001819055505b80600560003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008460ff1660ff1681526020019081526020016000206001016000828254610b19919061179d565b925050819055503373ffffffffffffffffffffffffffffffffffffffff167fedec38b4b62433445c926815411650aa7d417efa9da307f6aa99e69e6f4493ee8383604051610b68929190611774565b60405180910390a25050565b6000600360008360ff1660ff168152602001908152602001600020549050919050565b610b9f610fae565b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663ce5494bb826040518263ffffffff1660e01b8152600401610bfa9190611313565b600060405180830381600087803b158015610c1457600080fd5b505af1158015610c28573d6000803e3d6000fd5b5050505050565b600160149054906101000a900467ffffffffffffffff1681565b610c516110f8565b610c596110f8565b6040518060400160405280600560008773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008660ff1660ff168152602001908152602001600020600001548152602001600560008773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008660ff1660ff1681526020019081526020016000206001015481525090508091505092915050565b610d34610fae565b80600360008460ff1660ff168152602001908152602001600020819055507fd0373c2693b53f4b1f900ea4ed085eb4b9f149c2a5c54ecc9b5a1b190b08fb208282604051610d83929190611774565b60405180910390a15050565b610d97610fae565b600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1603610e06576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610dfd90611c04565b60405180910390fd5b610e0f8161102c565b50565b6000600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663693ec85e6040518163ffffffff1660e01b8152600401610e6d90611c70565b6020604051808303816000875af1158015610e8c573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610eb09190611640565b90508073ffffffffffffffffffffffffffffffffffffffff1663919840ad6040518163ffffffff1660e01b8152600401600060405180830381600087803b158015610efa57600080fd5b505af1158015610f0e573d6000803e3d6000fd5b505050508073ffffffffffffffffffffffffffffffffffffffff1663d88873156040518163ffffffff1660e01b81526004016020604051808303816000875af1158015610f5f573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610f839190611cbc565b600160146101000a81548167ffffffffffffffff021916908367ffffffffffffffff16021790555050565b610fb66110f0565b73ffffffffffffffffffffffffffffffffffffffff16610fd46108bd565b73ffffffffffffffffffffffffffffffffffffffff161461102a576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161102190611d35565b60405180910390fd5b565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050816000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508173ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a35050565b600033905090565b604051806040016040528060008152602001600081525090565b6000604051905090565b600080fd5b600080fd5b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b600061115182611126565b9050919050565b61116181611146565b811461116c57600080fd5b50565b60008135905061117e81611158565b92915050565b600060ff82169050919050565b61119a81611184565b81146111a557600080fd5b50565b6000813590506111b781611191565b92915050565b6000819050919050565b6111d0816111bd565b81146111db57600080fd5b50565b6000813590506111ed816111c7565b92915050565b6000806000806080858703121561120d5761120c61111c565b5b600061121b8782880161116f565b945050602061122c878288016111a8565b935050604061123d8782880161116f565b925050606061124e878288016111de565b91505092959194509250565b600080604083850312156112715761127061111c565b5b600061127f858286016111a8565b9250506020611290858286016111de565b9150509250929050565b600080604083850312156112b1576112b061111c565b5b60006112bf8582860161116f565b92505060206112d0858286016111a8565b9150509250929050565b6112e3816111bd565b82525050565b60006020820190506112fe60008301846112da565b92915050565b61130d81611146565b82525050565b60006020820190506113286000830184611304565b92915050565b600080fd5b600080fd5b6000601f19601f8301169050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b61138182611338565b810181811067ffffffffffffffff821117156113a05761139f611349565b5b80604052505050565b60006113b3611112565b90506113bf8282611378565b919050565b600067ffffffffffffffff8211156113df576113de611349565b5b6113e882611338565b9050602081019050919050565b82818337600083830152505050565b6000611417611412846113c4565b6113a9565b90508281526020810184848401111561143357611432611333565b5b61143e8482856113f5565b509392505050565b600082601f83011261145b5761145a61132e565b5b813561146b848260208601611404565b91505092915050565b60006020828403121561148a5761148961111c565b5b600082013567ffffffffffffffff8111156114a8576114a7611121565b5b6114b484828501611446565b91505092915050565b6000602082840312156114d3576114d261111c565b5b60006114e1848285016111a8565b91505092915050565b600060208284031215611500576114ff61111c565b5b600061150e8482850161116f565b91505092915050565b600067ffffffffffffffff82169050919050565b61153481611517565b82525050565b600060208201905061154f600083018461152b565b92915050565b61155e816111bd565b82525050565b60408201600082015161157a6000850182611555565b50602082015161158d6020850182611555565b50505050565b60006040820190506115a86000830184611564565b92915050565b600082825260208201905092915050565b7f70726f6f66000000000000000000000000000000000000000000000000000000600082015250565b60006115f56005836115ae565b9150611600826115bf565b602082019050919050565b60006020820190508181036000830152611624816115e8565b9050919050565b60008151905061163a81611158565b92915050565b6000602082840312156116565761165561111c565b5b60006116648482850161162b565b91505092915050565b7f696e76616c69642063616c6c6572000000000000000000000000000000000000600082015250565b60006116a3600e836115ae565b91506116ae8261166d565b602082019050919050565b600060208201905081810360008301526116d281611696565b9050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b6000611713826111bd565b915061171e836111bd565b9250828203905081811115611736576117356116d9565b5b92915050565b60006040820190506117516000830185611304565b61175e60208301846112da565b9392505050565b61176e81611184565b82525050565b60006040820190506117896000830185611765565b61179660208301846112da565b9392505050565b60006117a8826111bd565b91506117b3836111bd565b92508282019050808211156117cb576117ca6116d9565b5b92915050565b7f6c6f636b65640000000000000000000000000000000000000000000000000000600082015250565b60006118076006836115ae565b9150611812826117d1565b602082019050919050565b60006020820190508181036000830152611836816117fa565b9050919050565b7f696e73756620706c656467650000000000000000000000000000000000000000600082015250565b6000611873600c836115ae565b915061187e8261183d565b602082019050919050565b600060208201905081810360008301526118a281611866565b9050919050565b600081519050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b600060028204905060018216806118fb57607f821691505b60208210810361190e5761190d6118b4565b5b50919050565b60008190508160005260206000209050919050565b60006020601f8301049050919050565b600082821b905092915050565b6000600883026119767fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82611939565b6119808683611939565b95508019841693508086168417925050509392505050565b6000819050919050565b60006119bd6119b86119b3846111bd565b611998565b6111bd565b9050919050565b6000819050919050565b6119d7836119a2565b6119eb6119e3826119c4565b848454611946565b825550505050565b600090565b611a006119f3565b611a0b8184846119ce565b505050565b5b81811015611a2f57611a246000826119f8565b600181019050611a11565b5050565b601f821115611a7457611a4581611914565b611a4e84611929565b81016020851015611a5d578190505b611a71611a6985611929565b830182611a10565b50505b505050565b600082821c905092915050565b6000611a9760001984600802611a79565b1980831691505092915050565b6000611ab08383611a86565b9150826002028217905092915050565b611ac9826118a9565b67ffffffffffffffff811115611ae257611ae1611349565b5b611aec82546118e3565b611af7828285611a33565b600060209050601f831160018114611b2a5760008415611b18578287015190505b611b228582611aa4565b865550611b8a565b601f198416611b3886611914565b60005b82811015611b6057848901518255600182019150602085019450602081019050611b3b565b86831015611b7d5784890151611b79601f891682611a86565b8355505b6001600288020188555050505b505050505050565b7f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160008201527f6464726573730000000000000000000000000000000000000000000000000000602082015250565b6000611bee6026836115ae565b9150611bf982611b92565b604082019050919050565b60006020820190508181036000830152611c1d81611be1565b9050919050565b7f65706f6368000000000000000000000000000000000000000000000000000000600082015250565b6000611c5a6005836115ae565b9150611c6582611c24565b602082019050919050565b60006020820190508181036000830152611c8981611c4d565b9050919050565b611c9981611517565b8114611ca457600080fd5b50565b600081519050611cb681611c90565b92915050565b600060208284031215611cd257611cd161111c565b5b6000611ce084828501611ca7565b91505092915050565b7f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572600082015250565b6000611d1f6020836115ae565b9150611d2a82611ce9565b602082019050919050565b60006020820190508181036000830152611d4e81611d12565b905091905056fea26469706673582212207262a9cf072f2f7ec785510128863d8d80ce6eb43809ca956b4a3cc0dfa1217864736f6c63430008180033",
}

// NodeABI is the input ABI used to generate the binding from.
// Deprecated: Use NodeMetaData.ABI instead.
var NodeABI = NodeMetaData.ABI

// NodeBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use NodeMetaData.Bin instead.
var NodeBin = NodeMetaData.Bin

// DeployNode deploys a new Ethereum contract, binding an instance of Node to it.
func DeployNode(auth *bind.TransactOpts, backend bind.ContractBackend, _b common.Address) (common.Address, *types.Transaction, *Node, error) {
	parsed, err := NodeMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(NodeBin), backend, _b)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Node{NodeCaller: NodeCaller{contract: contract}, NodeTransactor: NodeTransactor{contract: contract}, NodeFilterer: NodeFilterer{contract: contract}}, nil
}

// Node is an auto generated Go binding around an Ethereum contract.
type Node struct {
	NodeCaller     // Read-only binding to the contract
	NodeTransactor // Write-only binding to the contract
	NodeFilterer   // Log filterer for contract events
}

// NodeCaller is an auto generated read-only Go binding around an Ethereum contract.
type NodeCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NodeTransactor is an auto generated write-only Go binding around an Ethereum contract.
type NodeTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NodeFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type NodeFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NodeSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type NodeSession struct {
	Contract     *Node             // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// NodeCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type NodeCallerSession struct {
	Contract *NodeCaller   // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// NodeTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type NodeTransactorSession struct {
	Contract     *NodeTransactor   // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// NodeRaw is an auto generated low-level Go binding around an Ethereum contract.
type NodeRaw struct {
	Contract *Node // Generic contract binding to access the raw methods on
}

// NodeCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type NodeCallerRaw struct {
	Contract *NodeCaller // Generic read-only contract binding to access the raw methods on
}

// NodeTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type NodeTransactorRaw struct {
	Contract *NodeTransactor // Generic write-only contract binding to access the raw methods on
}

// NewNode creates a new instance of Node, bound to a specific deployed contract.
func NewNode(address common.Address, backend bind.ContractBackend) (*Node, error) {
	contract, err := bindNode(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Node{NodeCaller: NodeCaller{contract: contract}, NodeTransactor: NodeTransactor{contract: contract}, NodeFilterer: NodeFilterer{contract: contract}}, nil
}

// NewNodeCaller creates a new read-only instance of Node, bound to a specific deployed contract.
func NewNodeCaller(address common.Address, caller bind.ContractCaller) (*NodeCaller, error) {
	contract, err := bindNode(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &NodeCaller{contract: contract}, nil
}

// NewNodeTransactor creates a new write-only instance of Node, bound to a specific deployed contract.
func NewNodeTransactor(address common.Address, transactor bind.ContractTransactor) (*NodeTransactor, error) {
	contract, err := bindNode(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &NodeTransactor{contract: contract}, nil
}

// NewNodeFilterer creates a new log filterer instance of Node, bound to a specific deployed contract.
func NewNodeFilterer(address common.Address, filterer bind.ContractFilterer) (*NodeFilterer, error) {
	contract, err := bindNode(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &NodeFilterer{contract: contract}, nil
}

// bindNode binds a generic wrapper to an already deployed contract.
func bindNode(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := NodeMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Node *NodeRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Node.Contract.NodeCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Node *NodeRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Node.Contract.NodeTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Node *NodeRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Node.Contract.NodeTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Node *NodeCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Node.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Node *NodeTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Node.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Node *NodeTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Node.Contract.contract.Transact(opts, method, params...)
}

// Bank is a free data retrieval call binding the contract method 0x76cdb03b.
//
// Solidity: function bank() view returns(address)
func (_Node *NodeCaller) Bank(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Node.contract.Call(opts, &out, "bank")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Bank is a free data retrieval call binding the contract method 0x76cdb03b.
//
// Solidity: function bank() view returns(address)
func (_Node *NodeSession) Bank() (common.Address, error) {
	return _Node.Contract.Bank(&_Node.CallOpts)
}

// Bank is a free data retrieval call binding the contract method 0x76cdb03b.
//
// Solidity: function bank() view returns(address)
func (_Node *NodeCallerSession) Bank() (common.Address, error) {
	return _Node.Contract.Bank(&_Node.CallOpts)
}

// Check is a free data retrieval call binding the contract method 0x61e728b1.
//
// Solidity: function check(address _a, uint8 _type) view returns()
func (_Node *NodeCaller) Check(opts *bind.CallOpts, _a common.Address, _type uint8) error {
	var out []interface{}
	err := _Node.contract.Call(opts, &out, "check", _a, _type)

	if err != nil {
		return err
	}

	return err

}

// Check is a free data retrieval call binding the contract method 0x61e728b1.
//
// Solidity: function check(address _a, uint8 _type) view returns()
func (_Node *NodeSession) Check(_a common.Address, _type uint8) error {
	return _Node.Contract.Check(&_Node.CallOpts, _a, _type)
}

// Check is a free data retrieval call binding the contract method 0x61e728b1.
//
// Solidity: function check(address _a, uint8 _type) view returns()
func (_Node *NodeCallerSession) Check(_a common.Address, _type uint8) error {
	return _Node.Contract.Check(&_Node.CallOpts, _a, _type)
}

// CurEpoch is a free data retrieval call binding the contract method 0xd8887315.
//
// Solidity: function curEpoch() view returns(uint64)
func (_Node *NodeCaller) CurEpoch(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _Node.contract.Call(opts, &out, "curEpoch")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// CurEpoch is a free data retrieval call binding the contract method 0xd8887315.
//
// Solidity: function curEpoch() view returns(uint64)
func (_Node *NodeSession) CurEpoch() (uint64, error) {
	return _Node.Contract.CurEpoch(&_Node.CallOpts)
}

// CurEpoch is a free data retrieval call binding the contract method 0xd8887315.
//
// Solidity: function curEpoch() view returns(uint64)
func (_Node *NodeCallerSession) CurEpoch() (uint64, error) {
	return _Node.Contract.CurEpoch(&_Node.CallOpts)
}

// GetMinPledge is a free data retrieval call binding the contract method 0x98c365d5.
//
// Solidity: function getMinPledge(uint8 _type) view returns(uint256)
func (_Node *NodeCaller) GetMinPledge(opts *bind.CallOpts, _type uint8) (*big.Int, error) {
	var out []interface{}
	err := _Node.contract.Call(opts, &out, "getMinPledge", _type)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetMinPledge is a free data retrieval call binding the contract method 0x98c365d5.
//
// Solidity: function getMinPledge(uint8 _type) view returns(uint256)
func (_Node *NodeSession) GetMinPledge(_type uint8) (*big.Int, error) {
	return _Node.Contract.GetMinPledge(&_Node.CallOpts, _type)
}

// GetMinPledge is a free data retrieval call binding the contract method 0x98c365d5.
//
// Solidity: function getMinPledge(uint8 _type) view returns(uint256)
func (_Node *NodeCallerSession) GetMinPledge(_type uint8) (*big.Int, error) {
	return _Node.Contract.GetMinPledge(&_Node.CallOpts, _type)
}

// GetPledge is a free data retrieval call binding the contract method 0xe3fdacb6.
//
// Solidity: function getPledge(address _a, uint8 _type) view returns((uint256,uint256))
func (_Node *NodeCaller) GetPledge(opts *bind.CallOpts, _a common.Address, _type uint8) (INodePledgeInfo, error) {
	var out []interface{}
	err := _Node.contract.Call(opts, &out, "getPledge", _a, _type)

	if err != nil {
		return *new(INodePledgeInfo), err
	}

	out0 := *abi.ConvertType(out[0], new(INodePledgeInfo)).(*INodePledgeInfo)

	return out0, err

}

// GetPledge is a free data retrieval call binding the contract method 0xe3fdacb6.
//
// Solidity: function getPledge(address _a, uint8 _type) view returns((uint256,uint256))
func (_Node *NodeSession) GetPledge(_a common.Address, _type uint8) (INodePledgeInfo, error) {
	return _Node.Contract.GetPledge(&_Node.CallOpts, _a, _type)
}

// GetPledge is a free data retrieval call binding the contract method 0xe3fdacb6.
//
// Solidity: function getPledge(address _a, uint8 _type) view returns((uint256,uint256))
func (_Node *NodeCallerSession) GetPledge(_a common.Address, _type uint8) (INodePledgeInfo, error) {
	return _Node.Contract.GetPledge(&_Node.CallOpts, _a, _type)
}

// GetPledgeValue is a free data retrieval call binding the contract method 0x4339ceca.
//
// Solidity: function getPledgeValue(address _a, uint8 _type) view returns(uint256)
func (_Node *NodeCaller) GetPledgeValue(opts *bind.CallOpts, _a common.Address, _type uint8) (*big.Int, error) {
	var out []interface{}
	err := _Node.contract.Call(opts, &out, "getPledgeValue", _a, _type)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetPledgeValue is a free data retrieval call binding the contract method 0x4339ceca.
//
// Solidity: function getPledgeValue(address _a, uint8 _type) view returns(uint256)
func (_Node *NodeSession) GetPledgeValue(_a common.Address, _type uint8) (*big.Int, error) {
	return _Node.Contract.GetPledgeValue(&_Node.CallOpts, _a, _type)
}

// GetPledgeValue is a free data retrieval call binding the contract method 0x4339ceca.
//
// Solidity: function getPledgeValue(address _a, uint8 _type) view returns(uint256)
func (_Node *NodeCallerSession) GetPledgeValue(_a common.Address, _type uint8) (*big.Int, error) {
	return _Node.Contract.GetPledgeValue(&_Node.CallOpts, _a, _type)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Node *NodeCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Node.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Node *NodeSession) Owner() (common.Address, error) {
	return _Node.Contract.Owner(&_Node.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Node *NodeCallerSession) Owner() (common.Address, error) {
	return _Node.Contract.Owner(&_Node.CallOpts)
}

// Declare is a paid mutator transaction binding the contract method 0x90d8cf07.
//
// Solidity: function declare(bytes _detail) returns()
func (_Node *NodeTransactor) Declare(opts *bind.TransactOpts, _detail []byte) (*types.Transaction, error) {
	return _Node.contract.Transact(opts, "declare", _detail)
}

// Declare is a paid mutator transaction binding the contract method 0x90d8cf07.
//
// Solidity: function declare(bytes _detail) returns()
func (_Node *NodeSession) Declare(_detail []byte) (*types.Transaction, error) {
	return _Node.Contract.Declare(&_Node.TransactOpts, _detail)
}

// Declare is a paid mutator transaction binding the contract method 0x90d8cf07.
//
// Solidity: function declare(bytes _detail) returns()
func (_Node *NodeTransactorSession) Declare(_detail []byte) (*types.Transaction, error) {
	return _Node.Contract.Declare(&_Node.TransactOpts, _detail)
}

// Migrate is a paid mutator transaction binding the contract method 0xce5494bb.
//
// Solidity: function migrate(address _to) returns()
func (_Node *NodeTransactor) Migrate(opts *bind.TransactOpts, _to common.Address) (*types.Transaction, error) {
	return _Node.contract.Transact(opts, "migrate", _to)
}

// Migrate is a paid mutator transaction binding the contract method 0xce5494bb.
//
// Solidity: function migrate(address _to) returns()
func (_Node *NodeSession) Migrate(_to common.Address) (*types.Transaction, error) {
	return _Node.Contract.Migrate(&_Node.TransactOpts, _to)
}

// Migrate is a paid mutator transaction binding the contract method 0xce5494bb.
//
// Solidity: function migrate(address _to) returns()
func (_Node *NodeTransactorSession) Migrate(_to common.Address) (*types.Transaction, error) {
	return _Node.Contract.Migrate(&_Node.TransactOpts, _to)
}

// Pledge is a paid mutator transaction binding the contract method 0x9708cab3.
//
// Solidity: function pledge(uint8 _typ, uint256 _m) returns()
func (_Node *NodeTransactor) Pledge(opts *bind.TransactOpts, _typ uint8, _m *big.Int) (*types.Transaction, error) {
	return _Node.contract.Transact(opts, "pledge", _typ, _m)
}

// Pledge is a paid mutator transaction binding the contract method 0x9708cab3.
//
// Solidity: function pledge(uint8 _typ, uint256 _m) returns()
func (_Node *NodeSession) Pledge(_typ uint8, _m *big.Int) (*types.Transaction, error) {
	return _Node.Contract.Pledge(&_Node.TransactOpts, _typ, _m)
}

// Pledge is a paid mutator transaction binding the contract method 0x9708cab3.
//
// Solidity: function pledge(uint8 _typ, uint256 _m) returns()
func (_Node *NodeTransactorSession) Pledge(_typ uint8, _m *big.Int) (*types.Transaction, error) {
	return _Node.Contract.Pledge(&_Node.TransactOpts, _typ, _m)
}

// Punish is a paid mutator transaction binding the contract method 0x06f0b4f1.
//
// Solidity: function punish(address _from, uint8 _typ, address _to, uint256 _m) returns()
func (_Node *NodeTransactor) Punish(opts *bind.TransactOpts, _from common.Address, _typ uint8, _to common.Address, _m *big.Int) (*types.Transaction, error) {
	return _Node.contract.Transact(opts, "punish", _from, _typ, _to, _m)
}

// Punish is a paid mutator transaction binding the contract method 0x06f0b4f1.
//
// Solidity: function punish(address _from, uint8 _typ, address _to, uint256 _m) returns()
func (_Node *NodeSession) Punish(_from common.Address, _typ uint8, _to common.Address, _m *big.Int) (*types.Transaction, error) {
	return _Node.Contract.Punish(&_Node.TransactOpts, _from, _typ, _to, _m)
}

// Punish is a paid mutator transaction binding the contract method 0x06f0b4f1.
//
// Solidity: function punish(address _from, uint8 _typ, address _to, uint256 _m) returns()
func (_Node *NodeTransactorSession) Punish(_from common.Address, _typ uint8, _to common.Address, _m *big.Int) (*types.Transaction, error) {
	return _Node.Contract.Punish(&_Node.TransactOpts, _from, _typ, _to, _m)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Node *NodeTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Node.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Node *NodeSession) RenounceOwnership() (*types.Transaction, error) {
	return _Node.Contract.RenounceOwnership(&_Node.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Node *NodeTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Node.Contract.RenounceOwnership(&_Node.TransactOpts)
}

// SetPledge is a paid mutator transaction binding the contract method 0xe6990adf.
//
// Solidity: function setPledge(uint8 _type, uint256 _money) returns()
func (_Node *NodeTransactor) SetPledge(opts *bind.TransactOpts, _type uint8, _money *big.Int) (*types.Transaction, error) {
	return _Node.contract.Transact(opts, "setPledge", _type, _money)
}

// SetPledge is a paid mutator transaction binding the contract method 0xe6990adf.
//
// Solidity: function setPledge(uint8 _type, uint256 _money) returns()
func (_Node *NodeSession) SetPledge(_type uint8, _money *big.Int) (*types.Transaction, error) {
	return _Node.Contract.SetPledge(&_Node.TransactOpts, _type, _money)
}

// SetPledge is a paid mutator transaction binding the contract method 0xe6990adf.
//
// Solidity: function setPledge(uint8 _type, uint256 _money) returns()
func (_Node *NodeTransactorSession) SetPledge(_type uint8, _money *big.Int) (*types.Transaction, error) {
	return _Node.Contract.SetPledge(&_Node.TransactOpts, _type, _money)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Node *NodeTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Node.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Node *NodeSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Node.Contract.TransferOwnership(&_Node.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Node *NodeTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Node.Contract.TransferOwnership(&_Node.TransactOpts, newOwner)
}

// Withdraw is a paid mutator transaction binding the contract method 0x3f489914.
//
// Solidity: function withdraw(uint8 _typ, uint256 _m) returns()
func (_Node *NodeTransactor) Withdraw(opts *bind.TransactOpts, _typ uint8, _m *big.Int) (*types.Transaction, error) {
	return _Node.contract.Transact(opts, "withdraw", _typ, _m)
}

// Withdraw is a paid mutator transaction binding the contract method 0x3f489914.
//
// Solidity: function withdraw(uint8 _typ, uint256 _m) returns()
func (_Node *NodeSession) Withdraw(_typ uint8, _m *big.Int) (*types.Transaction, error) {
	return _Node.Contract.Withdraw(&_Node.TransactOpts, _typ, _m)
}

// Withdraw is a paid mutator transaction binding the contract method 0x3f489914.
//
// Solidity: function withdraw(uint8 _typ, uint256 _m) returns()
func (_Node *NodeTransactorSession) Withdraw(_typ uint8, _m *big.Int) (*types.Transaction, error) {
	return _Node.Contract.Withdraw(&_Node.TransactOpts, _typ, _m)
}

// NodeOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Node contract.
type NodeOwnershipTransferredIterator struct {
	Event *NodeOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *NodeOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NodeOwnershipTransferred)
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
		it.Event = new(NodeOwnershipTransferred)
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
func (it *NodeOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NodeOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NodeOwnershipTransferred represents a OwnershipTransferred event raised by the Node contract.
type NodeOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Node *NodeFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*NodeOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Node.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &NodeOwnershipTransferredIterator{contract: _Node.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Node *NodeFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *NodeOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Node.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NodeOwnershipTransferred)
				if err := _Node.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_Node *NodeFilterer) ParseOwnershipTransferred(log types.Log) (*NodeOwnershipTransferred, error) {
	event := new(NodeOwnershipTransferred)
	if err := _Node.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NodePledgeIterator is returned from FilterPledge and is used to iterate over the raw logs and unpacked data for Pledge events raised by the Node contract.
type NodePledgeIterator struct {
	Event *NodePledge // Event containing the contract specifics and raw log

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
func (it *NodePledgeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NodePledge)
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
		it.Event = new(NodePledge)
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
func (it *NodePledgeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NodePledgeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NodePledge represents a Pledge event raised by the Node contract.
type NodePledge struct {
	A     common.Address
	Type  uint8
	Money *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterPledge is a free log retrieval operation binding the contract event 0xedec38b4b62433445c926815411650aa7d417efa9da307f6aa99e69e6f4493ee.
//
// Solidity: event Pledge(address indexed _a, uint8 _type, uint256 _money)
func (_Node *NodeFilterer) FilterPledge(opts *bind.FilterOpts, _a []common.Address) (*NodePledgeIterator, error) {

	var _aRule []interface{}
	for _, _aItem := range _a {
		_aRule = append(_aRule, _aItem)
	}

	logs, sub, err := _Node.contract.FilterLogs(opts, "Pledge", _aRule)
	if err != nil {
		return nil, err
	}
	return &NodePledgeIterator{contract: _Node.contract, event: "Pledge", logs: logs, sub: sub}, nil
}

// WatchPledge is a free log subscription operation binding the contract event 0xedec38b4b62433445c926815411650aa7d417efa9da307f6aa99e69e6f4493ee.
//
// Solidity: event Pledge(address indexed _a, uint8 _type, uint256 _money)
func (_Node *NodeFilterer) WatchPledge(opts *bind.WatchOpts, sink chan<- *NodePledge, _a []common.Address) (event.Subscription, error) {

	var _aRule []interface{}
	for _, _aItem := range _a {
		_aRule = append(_aRule, _aItem)
	}

	logs, sub, err := _Node.contract.WatchLogs(opts, "Pledge", _aRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NodePledge)
				if err := _Node.contract.UnpackLog(event, "Pledge", log); err != nil {
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
func (_Node *NodeFilterer) ParsePledge(log types.Log) (*NodePledge, error) {
	event := new(NodePledge)
	if err := _Node.contract.UnpackLog(event, "Pledge", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NodePunishIterator is returned from FilterPunish and is used to iterate over the raw logs and unpacked data for Punish events raised by the Node contract.
type NodePunishIterator struct {
	Event *NodePunish // Event containing the contract specifics and raw log

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
func (it *NodePunishIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NodePunish)
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
		it.Event = new(NodePunish)
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
func (it *NodePunishIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NodePunishIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NodePunish represents a Punish event raised by the Node contract.
type NodePunish struct {
	A     common.Address
	Typ   uint8
	To    common.Address
	Money *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterPunish is a free log retrieval operation binding the contract event 0xc6ef96923e613455515c6723eff1723445b22427fe442e8bf742e9d29b4b3c32.
//
// Solidity: event Punish(address indexed _a, uint8 _typ, address indexed _to, uint256 _money)
func (_Node *NodeFilterer) FilterPunish(opts *bind.FilterOpts, _a []common.Address, _to []common.Address) (*NodePunishIterator, error) {

	var _aRule []interface{}
	for _, _aItem := range _a {
		_aRule = append(_aRule, _aItem)
	}

	var _toRule []interface{}
	for _, _toItem := range _to {
		_toRule = append(_toRule, _toItem)
	}

	logs, sub, err := _Node.contract.FilterLogs(opts, "Punish", _aRule, _toRule)
	if err != nil {
		return nil, err
	}
	return &NodePunishIterator{contract: _Node.contract, event: "Punish", logs: logs, sub: sub}, nil
}

// WatchPunish is a free log subscription operation binding the contract event 0xc6ef96923e613455515c6723eff1723445b22427fe442e8bf742e9d29b4b3c32.
//
// Solidity: event Punish(address indexed _a, uint8 _typ, address indexed _to, uint256 _money)
func (_Node *NodeFilterer) WatchPunish(opts *bind.WatchOpts, sink chan<- *NodePunish, _a []common.Address, _to []common.Address) (event.Subscription, error) {

	var _aRule []interface{}
	for _, _aItem := range _a {
		_aRule = append(_aRule, _aItem)
	}

	var _toRule []interface{}
	for _, _toItem := range _to {
		_toRule = append(_toRule, _toItem)
	}

	logs, sub, err := _Node.contract.WatchLogs(opts, "Punish", _aRule, _toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NodePunish)
				if err := _Node.contract.UnpackLog(event, "Punish", log); err != nil {
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
func (_Node *NodeFilterer) ParsePunish(log types.Log) (*NodePunish, error) {
	event := new(NodePunish)
	if err := _Node.contract.UnpackLog(event, "Punish", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NodeSetPledgeIterator is returned from FilterSetPledge and is used to iterate over the raw logs and unpacked data for SetPledge events raised by the Node contract.
type NodeSetPledgeIterator struct {
	Event *NodeSetPledge // Event containing the contract specifics and raw log

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
func (it *NodeSetPledgeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NodeSetPledge)
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
		it.Event = new(NodeSetPledge)
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
func (it *NodeSetPledgeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NodeSetPledgeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NodeSetPledge represents a SetPledge event raised by the Node contract.
type NodeSetPledge struct {
	Type  uint8
	Money *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterSetPledge is a free log retrieval operation binding the contract event 0xd0373c2693b53f4b1f900ea4ed085eb4b9f149c2a5c54ecc9b5a1b190b08fb20.
//
// Solidity: event SetPledge(uint8 _type, uint256 _money)
func (_Node *NodeFilterer) FilterSetPledge(opts *bind.FilterOpts) (*NodeSetPledgeIterator, error) {

	logs, sub, err := _Node.contract.FilterLogs(opts, "SetPledge")
	if err != nil {
		return nil, err
	}
	return &NodeSetPledgeIterator{contract: _Node.contract, event: "SetPledge", logs: logs, sub: sub}, nil
}

// WatchSetPledge is a free log subscription operation binding the contract event 0xd0373c2693b53f4b1f900ea4ed085eb4b9f149c2a5c54ecc9b5a1b190b08fb20.
//
// Solidity: event SetPledge(uint8 _type, uint256 _money)
func (_Node *NodeFilterer) WatchSetPledge(opts *bind.WatchOpts, sink chan<- *NodeSetPledge) (event.Subscription, error) {

	logs, sub, err := _Node.contract.WatchLogs(opts, "SetPledge")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NodeSetPledge)
				if err := _Node.contract.UnpackLog(event, "SetPledge", log); err != nil {
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

// ParseSetPledge is a log parse operation binding the contract event 0xd0373c2693b53f4b1f900ea4ed085eb4b9f149c2a5c54ecc9b5a1b190b08fb20.
//
// Solidity: event SetPledge(uint8 _type, uint256 _money)
func (_Node *NodeFilterer) ParseSetPledge(log types.Log) (*NodeSetPledge, error) {
	event := new(NodeSetPledge)
	if err := _Node.contract.UnpackLog(event, "SetPledge", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NodeWithdrawIterator is returned from FilterWithdraw and is used to iterate over the raw logs and unpacked data for Withdraw events raised by the Node contract.
type NodeWithdrawIterator struct {
	Event *NodeWithdraw // Event containing the contract specifics and raw log

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
func (it *NodeWithdrawIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NodeWithdraw)
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
		it.Event = new(NodeWithdraw)
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
func (it *NodeWithdrawIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NodeWithdrawIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NodeWithdraw represents a Withdraw event raised by the Node contract.
type NodeWithdraw struct {
	A     common.Address
	Type  uint8
	Money *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterWithdraw is a free log retrieval operation binding the contract event 0xe3054ec6f352b09a86839beeeceb27ed4684f9164eb61f4a5d6d159ee8789d74.
//
// Solidity: event Withdraw(address indexed _a, uint8 _type, uint256 _money)
func (_Node *NodeFilterer) FilterWithdraw(opts *bind.FilterOpts, _a []common.Address) (*NodeWithdrawIterator, error) {

	var _aRule []interface{}
	for _, _aItem := range _a {
		_aRule = append(_aRule, _aItem)
	}

	logs, sub, err := _Node.contract.FilterLogs(opts, "Withdraw", _aRule)
	if err != nil {
		return nil, err
	}
	return &NodeWithdrawIterator{contract: _Node.contract, event: "Withdraw", logs: logs, sub: sub}, nil
}

// WatchWithdraw is a free log subscription operation binding the contract event 0xe3054ec6f352b09a86839beeeceb27ed4684f9164eb61f4a5d6d159ee8789d74.
//
// Solidity: event Withdraw(address indexed _a, uint8 _type, uint256 _money)
func (_Node *NodeFilterer) WatchWithdraw(opts *bind.WatchOpts, sink chan<- *NodeWithdraw, _a []common.Address) (event.Subscription, error) {

	var _aRule []interface{}
	for _, _aItem := range _a {
		_aRule = append(_aRule, _aItem)
	}

	logs, sub, err := _Node.contract.WatchLogs(opts, "Withdraw", _aRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NodeWithdraw)
				if err := _Node.contract.UnpackLog(event, "Withdraw", log); err != nil {
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
func (_Node *NodeFilterer) ParseWithdraw(log types.Log) (*NodeWithdraw, error) {
	event := new(NodeWithdraw)
	if err := _Node.contract.UnpackLog(event, "Withdraw", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
