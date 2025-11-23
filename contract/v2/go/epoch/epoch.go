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

// EpochMetaData contains all meta data concerning the Epoch contract.
var EpochMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"UPGRADE_INTERFACE_VERSION\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"VERSION\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"check\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"current\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getEpoch\",\"inputs\":[{\"name\":\"_epoch\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"initialize\",\"inputs\":[{\"name\":\"_slots\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"initialOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"proxiableUUID\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"renounceOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setSlots\",\"inputs\":[{\"name\":\"_slots\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"slots\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"upgradeToAndCall\",\"inputs\":[{\"name\":\"newImplementation\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"event\",\"name\":\"Initialized\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"previousOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"SetEpoch\",\"inputs\":[{\"name\":\"_e\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"UpdateSlots\",\"inputs\":[{\"name\":\"_s\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Upgraded\",\"inputs\":[{\"name\":\"implementation\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"AddressEmptyCode\",\"inputs\":[{\"name\":\"target\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ERC1967InvalidImplementation\",\"inputs\":[{\"name\":\"implementation\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ERC1967NonPayable\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"FailedCall\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidInitialization\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotInitializing\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OwnableInvalidOwner\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"OwnableUnauthorizedAccount\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"UUPSUnauthorizedCallContext\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"UUPSUnsupportedProxiableUUID\",\"inputs\":[{\"name\":\"slot\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]}]",
	Bin: "0x60a06040523073ffffffffffffffffffffffffffffffffffffffff1660809073ffffffffffffffffffffffffffffffffffffffff1681525034801562000043575f80fd5b50620000546200005a60201b60201c565b620001e1565b5f6200006b6200015e60201b60201c565b9050805f0160089054906101000a900460ff1615620000b6576040517ff92ee8a900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b67ffffffffffffffff8016815f015f9054906101000a900467ffffffffffffffff1667ffffffffffffffff16146200015b5767ffffffffffffffff815f015f6101000a81548167ffffffffffffffff021916908367ffffffffffffffff1602179055507fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d267ffffffffffffffff604051620001529190620001c6565b60405180910390a15b50565b5f80620001706200017960201b60201c565b90508091505090565b5f7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a005f1b905090565b5f67ffffffffffffffff82169050919050565b620001c081620001a2565b82525050565b5f602082019050620001db5f830184620001b5565b92915050565b608051611b3a620002085f395f8181610b4401528181610b990152610d530152611b3a5ff3fe6080604052600436106100c1575f3560e01c8063919840ad1161007e578063cbccf09111610058578063cbccf0911461021c578063d7eecc7e14610244578063f2fde38b1461026c578063ffa1ad7414610294576100c1565b8063919840ad146101b25780639fa6a6e3146101c8578063ad3cb1cc146101f2576100c1565b806312a02c82146100c557806348547d69146101025780634f1ef2861461012c57806352d1902d14610148578063715018a6146101725780638da5cb5b14610188575b5f80fd5b3480156100d0575f80fd5b506100eb60048036038101906100e69190611400565b6102be565b6040516100f992919061145b565b60405180910390f35b34801561010d575f80fd5b50610116610324565b6040516101239190611491565b60405180910390f35b61014660048036038101906101419190611640565b61033d565b005b348015610153575f80fd5b5061015c61035c565b604051610169919061169a565b60405180910390f35b34801561017d575f80fd5b5061018661038d565b005b348015610193575f80fd5b5061019c6103a0565b6040516101a991906116c2565b60405180910390f35b3480156101bd575f80fd5b506101c66103d5565b005b3480156101d3575f80fd5b506101dc6105cb565b6040516101e99190611491565b60405180910390f35b3480156101fd575f80fd5b506102066105e2565b6040516102139190611755565b60405180910390f35b348015610227575f80fd5b50610242600480360381019061023d9190611400565b61061b565b005b34801561024f575f80fd5b5061026a60048036038101906102659190611775565b6106d1565b005b348015610277575f80fd5b50610292600480360381019061028d91906117b3565b610a85565b005b34801561029f575f80fd5b506102a8610b09565b6040516102b59190611755565b60405180910390f35b5f8060028367ffffffffffffffff16815481106102de576102dd6117de565b5b905f5260205f2090600202015f015460028467ffffffffffffffff168154811061030b5761030a6117de565b5b905f5260205f2090600202016001015491509150915091565b5f60089054906101000a900467ffffffffffffffff1681565b610345610b42565b61034e82610c28565b6103588282610c33565b5050565b5f610365610d51565b7f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc5f1b905090565b610395610dd8565b61039e5f610e5f565b565b5f806103aa610f30565b9050805f015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1691505090565b5a60015f8282546103e69190611838565b925050819055505f60089054906101000a900467ffffffffffffffff1667ffffffffffffffff1660025f8054906101000a900467ffffffffffffffff1667ffffffffffffffff168154811061043e5761043d6117de565b5b905f5260205f2090600202015f015443610458919061186b565b106105c957610465611397565b43815f01818152505060025f8054906101000a900467ffffffffffffffff1667ffffffffffffffff168154811061049f5761049e6117de565b5b905f5260205f2090600202016001015433426001436104be919061186b565b406001546040516020016104d6959493929190611923565b60405160208183030381529060405280519060200120816020018181525050600281908060018154018082558091505060019003905f5260205f2090600202015f909190919091505f820151815f01556020820151816001015550505f8081819054906101000a900467ffffffffffffffff168092919061055690611981565b91906101000a81548167ffffffffffffffff021916908367ffffffffffffffff160217905550507f9d4ccb161ea809df14334516cc3070025c80baddb8e8364afaca6a6fe31bfd755f8054906101000a900467ffffffffffffffff166040516105bf9190611491565b60405180910390a1505b565b5f8054906101000a900467ffffffffffffffff1681565b6040518060400160405280600581526020017f352e302e3000000000000000000000000000000000000000000000000000000081525081565b610623610dd8565b5f8167ffffffffffffffff161161066f576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610666906119fa565b60405180910390fd5b805f60086101000a81548167ffffffffffffffff021916908367ffffffffffffffff1602179055507fba1da09884224d94b905b0a3a828942e734c70ad3abf9cb4a66bda28fa7e6486816040516106c69190611491565b60405180910390a150565b5f6106da610f57565b90505f815f0160089054906101000a900460ff161590505f825f015f9054906101000a900467ffffffffffffffff1690505f808267ffffffffffffffff161480156107225750825b90505f60018367ffffffffffffffff1614801561075557505f3073ffffffffffffffffffffffffffffffffffffffff163b145b905081158015610763575080155b1561079a576040517ff92ee8a900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6001855f015f6101000a81548167ffffffffffffffff021916908367ffffffffffffffff16021790555083156107e7576001855f0160086101000a81548160ff0219169083151502179055505b6107f086610f6a565b865f60086101000a81548167ffffffffffffffff021916908367ffffffffffffffff160217905550610820611397565b5f815f0181815250503342600143610838919061186b565b4060405160200161084b93929190611a18565b60405160208183030381529060405280519060200120816020018181525050600281908060018154018082558091505060019003905f5260205f2090600202015f909190919091505f820151815f01556020820151816001015550505a60015f8282546108b89190611838565b9250508190555043815f01818152505060025f8054906101000a900467ffffffffffffffff1667ffffffffffffffff16815481106108f9576108f86117de565b5b905f5260205f209060020201600101543342600143610918919061186b565b40600154604051602001610930959493929190611923565b60405160208183030381529060405280519060200120816020018181525050600281908060018154018082558091505060019003905f5260205f2090600202015f909190919091505f820151815f01556020820151816001015550505f8081819054906101000a900467ffffffffffffffff16809291906109b090611981565b91906101000a81548167ffffffffffffffff021916908367ffffffffffffffff160217905550507f9d4ccb161ea809df14334516cc3070025c80baddb8e8364afaca6a6fe31bfd755f8054906101000a900467ffffffffffffffff16604051610a199190611491565b60405180910390a1508315610a7c575f855f0160086101000a81548160ff0219169083151502179055507fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d26001604051610a739190611a96565b60405180910390a15b50505050505050565b610a8d610dd8565b5f73ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1603610afd575f6040517f1e4fbdf7000000000000000000000000000000000000000000000000000000008152600401610af491906116c2565b60405180910390fd5b610b0681610e5f565b50565b6040518060400160405280600581526020017f322e302e3000000000000000000000000000000000000000000000000000000081525081565b7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff163073ffffffffffffffffffffffffffffffffffffffff161480610bef57507f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff16610bd6610f7e565b73ffffffffffffffffffffffffffffffffffffffff1614155b15610c26576040517fe07c8dba00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b565b610c30610dd8565b50565b8173ffffffffffffffffffffffffffffffffffffffff166352d1902d6040518163ffffffff1660e01b8152600401602060405180830381865afa925050508015610c9b57506040513d601f19601f82011682018060405250810190610c989190611ad9565b60015b610cdc57816040517f4c9c8ce3000000000000000000000000000000000000000000000000000000008152600401610cd391906116c2565b60405180910390fd5b7f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc5f1b8114610d4257806040517faa1d49a4000000000000000000000000000000000000000000000000000000008152600401610d39919061169a565b60405180910390fd5b610d4c8383610fd1565b505050565b7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff163073ffffffffffffffffffffffffffffffffffffffff1614610dd6576040517fe07c8dba00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b565b610de0611043565b73ffffffffffffffffffffffffffffffffffffffff16610dfe6103a0565b73ffffffffffffffffffffffffffffffffffffffff1614610e5d57610e21611043565b6040517f118cdaa7000000000000000000000000000000000000000000000000000000008152600401610e5491906116c2565b60405180910390fd5b565b5f610e68610f30565b90505f815f015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905082825f015f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508273ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a3505050565b5f7f9016d09d72d40fdae2fd8ceac6b6234c7706214fd39c1cd1e609a0528c199300905090565b5f80610f6161104a565b90508091505090565b610f72611073565b610f7b816110b3565b50565b5f610faa7f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc5f1b611137565b5f015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b610fda82611140565b8173ffffffffffffffffffffffffffffffffffffffff167fbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b60405160405180910390a25f81511115611036576110308282611209565b5061103f565b61103e6112fa565b5b5050565b5f33905090565b5f7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a005f1b905090565b61107b611336565b6110b1576040517fd7e6bcf800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b565b6110bb611073565b5f73ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff160361112b575f6040517f1e4fbdf700000000000000000000000000000000000000000000000000000000815260040161112291906116c2565b60405180910390fd5b61113481610e5f565b50565b5f819050919050565b5f8173ffffffffffffffffffffffffffffffffffffffff163b0361119b57806040517f4c9c8ce300000000000000000000000000000000000000000000000000000000815260040161119291906116c2565b60405180910390fd5b806111c77f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc5f1b611137565b5f015f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050565b60605f6112168484611354565b905080801561124c57505f611229611368565b118061124b57505f8473ffffffffffffffffffffffffffffffffffffffff163b115b5b156112615761125961136f565b9150506112f4565b80156112a457836040517f9996b31500000000000000000000000000000000000000000000000000000000815260040161129b91906116c2565b60405180910390fd5b5f6112ad611368565b11156112c0576112bb61138c565b6112f2565b6040517fd6bda27500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b505b92915050565b5f341115611334576040517fb398979f00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b565b5f61133f610f57565b5f0160089054906101000a900460ff16905090565b5f805f835160208501865af4905092915050565b5f3d905090565b606060405190503d81523d5f602083013e3d602001810160405290565b6040513d5f823e3d81fd5b60405180604001604052805f81526020015f80191681525090565b5f604051905090565b5f80fd5b5f80fd5b5f67ffffffffffffffff82169050919050565b6113df816113c3565b81146113e9575f80fd5b50565b5f813590506113fa816113d6565b92915050565b5f60208284031215611415576114146113bb565b5b5f611422848285016113ec565b91505092915050565b5f819050919050565b61143d8161142b565b82525050565b5f819050919050565b61145581611443565b82525050565b5f60408201905061146e5f830185611434565b61147b602083018461144c565b9392505050565b61148b816113c3565b82525050565b5f6020820190506114a45f830184611482565b92915050565b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f6114d3826114aa565b9050919050565b6114e3816114c9565b81146114ed575f80fd5b50565b5f813590506114fe816114da565b92915050565b5f80fd5b5f80fd5b5f601f19601f8301169050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b6115528261150c565b810181811067ffffffffffffffff821117156115715761157061151c565b5b80604052505050565b5f6115836113b2565b905061158f8282611549565b919050565b5f67ffffffffffffffff8211156115ae576115ad61151c565b5b6115b78261150c565b9050602081019050919050565b828183375f83830152505050565b5f6115e46115df84611594565b61157a565b905082815260208101848484011115611600576115ff611508565b5b61160b8482856115c4565b509392505050565b5f82601f83011261162757611626611504565b5b81356116378482602086016115d2565b91505092915050565b5f8060408385031215611656576116556113bb565b5b5f611663858286016114f0565b925050602083013567ffffffffffffffff811115611684576116836113bf565b5b61169085828601611613565b9150509250929050565b5f6020820190506116ad5f83018461144c565b92915050565b6116bc816114c9565b82525050565b5f6020820190506116d55f8301846116b3565b92915050565b5f81519050919050565b5f82825260208201905092915050565b5f5b838110156117125780820151818401526020810190506116f7565b5f8484015250505050565b5f611727826116db565b61173181856116e5565b93506117418185602086016116f5565b61174a8161150c565b840191505092915050565b5f6020820190508181035f83015261176d818461171d565b905092915050565b5f806040838503121561178b5761178a6113bb565b5b5f611798858286016113ec565b92505060206117a9858286016114f0565b9150509250929050565b5f602082840312156117c8576117c76113bb565b5b5f6117d5848285016114f0565b91505092915050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52603260045260245ffd5b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b5f6118428261142b565b915061184d8361142b565b92508282019050808211156118655761186461180b565b5b92915050565b5f6118758261142b565b91506118808361142b565b92508282039050818111156118985761189761180b565b5b92915050565b5f819050919050565b6118b86118b382611443565b61189e565b82525050565b5f8160601b9050919050565b5f6118d4826118be565b9050919050565b5f6118e5826118ca565b9050919050565b6118fd6118f8826114c9565b6118db565b82525050565b5f819050919050565b61191d6119188261142b565b611903565b82525050565b5f61192e82886118a7565b60208201915061193e82876118ec565b60148201915061194e828661190c565b60208201915061195e82856118a7565b60208201915061196e828461190c565b6020820191508190509695505050505050565b5f61198b826113c3565b915067ffffffffffffffff82036119a5576119a461180b565b5b600182019050919050565b7f7a65726f20736c6f7473000000000000000000000000000000000000000000005f82015250565b5f6119e4600a836116e5565b91506119ef826119b0565b602082019050919050565b5f6020820190508181035f830152611a11816119d8565b9050919050565b5f611a2382866118ec565b601482019150611a33828561190c565b602082019150611a4382846118a7565b602082019150819050949350505050565b5f819050919050565b5f819050919050565b5f611a80611a7b611a7684611a54565b611a5d565b6113c3565b9050919050565b611a9081611a66565b82525050565b5f602082019050611aa95f830184611a87565b92915050565b611ab881611443565b8114611ac2575f80fd5b50565b5f81519050611ad381611aaf565b92915050565b5f60208284031215611aee57611aed6113bb565b5b5f611afb84828501611ac5565b9150509291505056fea26469706673582212202265db12a13427cf55b397210e51fb58c3943917f5d9646bc28c1c0317878f2364736f6c63430008160033",
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

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_Epoch *EpochCaller) UPGRADEINTERFACEVERSION(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Epoch.contract.Call(opts, &out, "UPGRADE_INTERFACE_VERSION")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_Epoch *EpochSession) UPGRADEINTERFACEVERSION() (string, error) {
	return _Epoch.Contract.UPGRADEINTERFACEVERSION(&_Epoch.CallOpts)
}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_Epoch *EpochCallerSession) UPGRADEINTERFACEVERSION() (string, error) {
	return _Epoch.Contract.UPGRADEINTERFACEVERSION(&_Epoch.CallOpts)
}

// VERSION is a free data retrieval call binding the contract method 0xffa1ad74.
//
// Solidity: function VERSION() view returns(string)
func (_Epoch *EpochCaller) VERSION(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Epoch.contract.Call(opts, &out, "VERSION")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// VERSION is a free data retrieval call binding the contract method 0xffa1ad74.
//
// Solidity: function VERSION() view returns(string)
func (_Epoch *EpochSession) VERSION() (string, error) {
	return _Epoch.Contract.VERSION(&_Epoch.CallOpts)
}

// VERSION is a free data retrieval call binding the contract method 0xffa1ad74.
//
// Solidity: function VERSION() view returns(string)
func (_Epoch *EpochCallerSession) VERSION() (string, error) {
	return _Epoch.Contract.VERSION(&_Epoch.CallOpts)
}

// Current is a free data retrieval call binding the contract method 0x9fa6a6e3.
//
// Solidity: function current() view returns(uint64)
func (_Epoch *EpochCaller) Current(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _Epoch.contract.Call(opts, &out, "current")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// Current is a free data retrieval call binding the contract method 0x9fa6a6e3.
//
// Solidity: function current() view returns(uint64)
func (_Epoch *EpochSession) Current() (uint64, error) {
	return _Epoch.Contract.Current(&_Epoch.CallOpts)
}

// Current is a free data retrieval call binding the contract method 0x9fa6a6e3.
//
// Solidity: function current() view returns(uint64)
func (_Epoch *EpochCallerSession) Current() (uint64, error) {
	return _Epoch.Contract.Current(&_Epoch.CallOpts)
}

// GetEpoch is a free data retrieval call binding the contract method 0x12a02c82.
//
// Solidity: function getEpoch(uint64 _epoch) view returns(uint256, bytes32)
func (_Epoch *EpochCaller) GetEpoch(opts *bind.CallOpts, _epoch uint64) (*big.Int, [32]byte, error) {
	var out []interface{}
	err := _Epoch.contract.Call(opts, &out, "getEpoch", _epoch)

	if err != nil {
		return *new(*big.Int), *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new([32]byte)).(*[32]byte)

	return out0, out1, err

}

// GetEpoch is a free data retrieval call binding the contract method 0x12a02c82.
//
// Solidity: function getEpoch(uint64 _epoch) view returns(uint256, bytes32)
func (_Epoch *EpochSession) GetEpoch(_epoch uint64) (*big.Int, [32]byte, error) {
	return _Epoch.Contract.GetEpoch(&_Epoch.CallOpts, _epoch)
}

// GetEpoch is a free data retrieval call binding the contract method 0x12a02c82.
//
// Solidity: function getEpoch(uint64 _epoch) view returns(uint256, bytes32)
func (_Epoch *EpochCallerSession) GetEpoch(_epoch uint64) (*big.Int, [32]byte, error) {
	return _Epoch.Contract.GetEpoch(&_Epoch.CallOpts, _epoch)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Epoch *EpochCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Epoch.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Epoch *EpochSession) Owner() (common.Address, error) {
	return _Epoch.Contract.Owner(&_Epoch.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Epoch *EpochCallerSession) Owner() (common.Address, error) {
	return _Epoch.Contract.Owner(&_Epoch.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_Epoch *EpochCaller) ProxiableUUID(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Epoch.contract.Call(opts, &out, "proxiableUUID")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_Epoch *EpochSession) ProxiableUUID() ([32]byte, error) {
	return _Epoch.Contract.ProxiableUUID(&_Epoch.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_Epoch *EpochCallerSession) ProxiableUUID() ([32]byte, error) {
	return _Epoch.Contract.ProxiableUUID(&_Epoch.CallOpts)
}

// Slots is a free data retrieval call binding the contract method 0x48547d69.
//
// Solidity: function slots() view returns(uint64)
func (_Epoch *EpochCaller) Slots(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _Epoch.contract.Call(opts, &out, "slots")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// Slots is a free data retrieval call binding the contract method 0x48547d69.
//
// Solidity: function slots() view returns(uint64)
func (_Epoch *EpochSession) Slots() (uint64, error) {
	return _Epoch.Contract.Slots(&_Epoch.CallOpts)
}

// Slots is a free data retrieval call binding the contract method 0x48547d69.
//
// Solidity: function slots() view returns(uint64)
func (_Epoch *EpochCallerSession) Slots() (uint64, error) {
	return _Epoch.Contract.Slots(&_Epoch.CallOpts)
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

// Initialize is a paid mutator transaction binding the contract method 0xd7eecc7e.
//
// Solidity: function initialize(uint64 _slots, address initialOwner) returns()
func (_Epoch *EpochTransactor) Initialize(opts *bind.TransactOpts, _slots uint64, initialOwner common.Address) (*types.Transaction, error) {
	return _Epoch.contract.Transact(opts, "initialize", _slots, initialOwner)
}

// Initialize is a paid mutator transaction binding the contract method 0xd7eecc7e.
//
// Solidity: function initialize(uint64 _slots, address initialOwner) returns()
func (_Epoch *EpochSession) Initialize(_slots uint64, initialOwner common.Address) (*types.Transaction, error) {
	return _Epoch.Contract.Initialize(&_Epoch.TransactOpts, _slots, initialOwner)
}

// Initialize is a paid mutator transaction binding the contract method 0xd7eecc7e.
//
// Solidity: function initialize(uint64 _slots, address initialOwner) returns()
func (_Epoch *EpochTransactorSession) Initialize(_slots uint64, initialOwner common.Address) (*types.Transaction, error) {
	return _Epoch.Contract.Initialize(&_Epoch.TransactOpts, _slots, initialOwner)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Epoch *EpochTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Epoch.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Epoch *EpochSession) RenounceOwnership() (*types.Transaction, error) {
	return _Epoch.Contract.RenounceOwnership(&_Epoch.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Epoch *EpochTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Epoch.Contract.RenounceOwnership(&_Epoch.TransactOpts)
}

// SetSlots is a paid mutator transaction binding the contract method 0xcbccf091.
//
// Solidity: function setSlots(uint64 _slots) returns()
func (_Epoch *EpochTransactor) SetSlots(opts *bind.TransactOpts, _slots uint64) (*types.Transaction, error) {
	return _Epoch.contract.Transact(opts, "setSlots", _slots)
}

// SetSlots is a paid mutator transaction binding the contract method 0xcbccf091.
//
// Solidity: function setSlots(uint64 _slots) returns()
func (_Epoch *EpochSession) SetSlots(_slots uint64) (*types.Transaction, error) {
	return _Epoch.Contract.SetSlots(&_Epoch.TransactOpts, _slots)
}

// SetSlots is a paid mutator transaction binding the contract method 0xcbccf091.
//
// Solidity: function setSlots(uint64 _slots) returns()
func (_Epoch *EpochTransactorSession) SetSlots(_slots uint64) (*types.Transaction, error) {
	return _Epoch.Contract.SetSlots(&_Epoch.TransactOpts, _slots)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Epoch *EpochTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Epoch.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Epoch *EpochSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Epoch.Contract.TransferOwnership(&_Epoch.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Epoch *EpochTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Epoch.Contract.TransferOwnership(&_Epoch.TransactOpts, newOwner)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_Epoch *EpochTransactor) UpgradeToAndCall(opts *bind.TransactOpts, newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _Epoch.contract.Transact(opts, "upgradeToAndCall", newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_Epoch *EpochSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _Epoch.Contract.UpgradeToAndCall(&_Epoch.TransactOpts, newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_Epoch *EpochTransactorSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _Epoch.Contract.UpgradeToAndCall(&_Epoch.TransactOpts, newImplementation, data)
}

// EpochInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the Epoch contract.
type EpochInitializedIterator struct {
	Event *EpochInitialized // Event containing the contract specifics and raw log

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
func (it *EpochInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EpochInitialized)
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
		it.Event = new(EpochInitialized)
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
func (it *EpochInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EpochInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EpochInitialized represents a Initialized event raised by the Epoch contract.
type EpochInitialized struct {
	Version uint64
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_Epoch *EpochFilterer) FilterInitialized(opts *bind.FilterOpts) (*EpochInitializedIterator, error) {

	logs, sub, err := _Epoch.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &EpochInitializedIterator{contract: _Epoch.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_Epoch *EpochFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *EpochInitialized) (event.Subscription, error) {

	logs, sub, err := _Epoch.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EpochInitialized)
				if err := _Epoch.contract.UnpackLog(event, "Initialized", log); err != nil {
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

// ParseInitialized is a log parse operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_Epoch *EpochFilterer) ParseInitialized(log types.Log) (*EpochInitialized, error) {
	event := new(EpochInitialized)
	if err := _Epoch.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EpochOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Epoch contract.
type EpochOwnershipTransferredIterator struct {
	Event *EpochOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *EpochOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EpochOwnershipTransferred)
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
		it.Event = new(EpochOwnershipTransferred)
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
func (it *EpochOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EpochOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EpochOwnershipTransferred represents a OwnershipTransferred event raised by the Epoch contract.
type EpochOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Epoch *EpochFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*EpochOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Epoch.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &EpochOwnershipTransferredIterator{contract: _Epoch.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Epoch *EpochFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *EpochOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Epoch.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EpochOwnershipTransferred)
				if err := _Epoch.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_Epoch *EpochFilterer) ParseOwnershipTransferred(log types.Log) (*EpochOwnershipTransferred, error) {
	event := new(EpochOwnershipTransferred)
	if err := _Epoch.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
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

// EpochUpdateSlotsIterator is returned from FilterUpdateSlots and is used to iterate over the raw logs and unpacked data for UpdateSlots events raised by the Epoch contract.
type EpochUpdateSlotsIterator struct {
	Event *EpochUpdateSlots // Event containing the contract specifics and raw log

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
func (it *EpochUpdateSlotsIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EpochUpdateSlots)
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
		it.Event = new(EpochUpdateSlots)
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
func (it *EpochUpdateSlotsIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EpochUpdateSlotsIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EpochUpdateSlots represents a UpdateSlots event raised by the Epoch contract.
type EpochUpdateSlots struct {
	S   uint64
	Raw types.Log // Blockchain specific contextual infos
}

// FilterUpdateSlots is a free log retrieval operation binding the contract event 0xba1da09884224d94b905b0a3a828942e734c70ad3abf9cb4a66bda28fa7e6486.
//
// Solidity: event UpdateSlots(uint64 _s)
func (_Epoch *EpochFilterer) FilterUpdateSlots(opts *bind.FilterOpts) (*EpochUpdateSlotsIterator, error) {

	logs, sub, err := _Epoch.contract.FilterLogs(opts, "UpdateSlots")
	if err != nil {
		return nil, err
	}
	return &EpochUpdateSlotsIterator{contract: _Epoch.contract, event: "UpdateSlots", logs: logs, sub: sub}, nil
}

// WatchUpdateSlots is a free log subscription operation binding the contract event 0xba1da09884224d94b905b0a3a828942e734c70ad3abf9cb4a66bda28fa7e6486.
//
// Solidity: event UpdateSlots(uint64 _s)
func (_Epoch *EpochFilterer) WatchUpdateSlots(opts *bind.WatchOpts, sink chan<- *EpochUpdateSlots) (event.Subscription, error) {

	logs, sub, err := _Epoch.contract.WatchLogs(opts, "UpdateSlots")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EpochUpdateSlots)
				if err := _Epoch.contract.UnpackLog(event, "UpdateSlots", log); err != nil {
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

// ParseUpdateSlots is a log parse operation binding the contract event 0xba1da09884224d94b905b0a3a828942e734c70ad3abf9cb4a66bda28fa7e6486.
//
// Solidity: event UpdateSlots(uint64 _s)
func (_Epoch *EpochFilterer) ParseUpdateSlots(log types.Log) (*EpochUpdateSlots, error) {
	event := new(EpochUpdateSlots)
	if err := _Epoch.contract.UnpackLog(event, "UpdateSlots", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EpochUpgradedIterator is returned from FilterUpgraded and is used to iterate over the raw logs and unpacked data for Upgraded events raised by the Epoch contract.
type EpochUpgradedIterator struct {
	Event *EpochUpgraded // Event containing the contract specifics and raw log

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
func (it *EpochUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EpochUpgraded)
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
		it.Event = new(EpochUpgraded)
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
func (it *EpochUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EpochUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EpochUpgraded represents a Upgraded event raised by the Epoch contract.
type EpochUpgraded struct {
	Implementation common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterUpgraded is a free log retrieval operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_Epoch *EpochFilterer) FilterUpgraded(opts *bind.FilterOpts, implementation []common.Address) (*EpochUpgradedIterator, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _Epoch.contract.FilterLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return &EpochUpgradedIterator{contract: _Epoch.contract, event: "Upgraded", logs: logs, sub: sub}, nil
}

// WatchUpgraded is a free log subscription operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_Epoch *EpochFilterer) WatchUpgraded(opts *bind.WatchOpts, sink chan<- *EpochUpgraded, implementation []common.Address) (event.Subscription, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _Epoch.contract.WatchLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EpochUpgraded)
				if err := _Epoch.contract.UnpackLog(event, "Upgraded", log); err != nil {
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

// ParseUpgraded is a log parse operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_Epoch *EpochFilterer) ParseUpgraded(log types.Log) (*EpochUpgraded, error) {
	event := new(EpochUpgraded)
	if err := _Epoch.contract.UnpackLog(event, "Upgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
