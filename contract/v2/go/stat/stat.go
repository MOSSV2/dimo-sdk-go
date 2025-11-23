// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package stat

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

// StatMetaData contains all meta data concerning the Stat contract.
var StatMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"UPGRADE_INTERFACE_VERSION\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"VERSION\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"checkPiece\",\"inputs\":[{\"name\":\"_pn\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"epoch\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIEpoch\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"eproof\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIEProof\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"initialize\",\"inputs\":[{\"name\":\"_epoch\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_piece\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_rsproof\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_eproof\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"initialOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"piece\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIPiece\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"proxiableUUID\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"renounceOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"rsproof\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIRSProof\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"upgradeToAndCall\",\"inputs\":[{\"name\":\"newImplementation\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"event\",\"name\":\"Initialized\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"previousOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Upgraded\",\"inputs\":[{\"name\":\"implementation\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"AddressEmptyCode\",\"inputs\":[{\"name\":\"target\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ERC1967InvalidImplementation\",\"inputs\":[{\"name\":\"implementation\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ERC1967NonPayable\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"FailedCall\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidInitialization\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotInitializing\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OwnableInvalidOwner\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"OwnableUnauthorizedAccount\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"UUPSUnauthorizedCallContext\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"UUPSUnsupportedProxiableUUID\",\"inputs\":[{\"name\":\"slot\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]}]",
	Bin: "0x60a06040523073ffffffffffffffffffffffffffffffffffffffff1660809073ffffffffffffffffffffffffffffffffffffffff1681525034801562000043575f80fd5b50620000546200005a60201b60201c565b620001e1565b5f6200006b6200015e60201b60201c565b9050805f0160089054906101000a900460ff1615620000b6576040517ff92ee8a900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b67ffffffffffffffff8016815f015f9054906101000a900467ffffffffffffffff1667ffffffffffffffff16146200015b5767ffffffffffffffff815f015f6101000a81548167ffffffffffffffff021916908367ffffffffffffffff1602179055507fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d267ffffffffffffffff604051620001529190620001c6565b60405180910390a15b50565b5f80620001706200017960201b60201c565b90508091505090565b5f7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a005f1b905090565b5f67ffffffffffffffff82169050919050565b620001c081620001a2565b82525050565b5f602082019050620001db5f830184620001b5565b92915050565b608051611d90620002085f395f8181610c3001528181610c850152610e3f0152611d905ff3fe6080604052600436106100c1575f3560e01c80638da5cb5b1161007e578063ad3cb1cc11610058578063ad3cb1cc1461021b578063c2e8e71714610245578063f2fde38b14610281578063ffa1ad74146102a9576100c1565b80638da5cb5b1461019d578063900cf0cf146101c757806395ad21dc146101f1576100c1565b80631459457a146100c55780634f1ef286146100ed57806352d1902d14610109578063715018a61461013357806379ca7e0f1461014957806381cc0c7a14610173575b5f80fd5b3480156100d0575f80fd5b506100eb60048036038101906100e691906114c7565b6102d3565b005b6101076004803603810190610102919061167a565b610557565b005b348015610114575f80fd5b5061011d610576565b60405161012a91906116ec565b60405180910390f35b34801561013e575f80fd5b506101476105a7565b005b348015610154575f80fd5b5061015d6105ba565b60405161016a9190611760565b60405180910390f35b34801561017e575f80fd5b506101876105df565b6040516101949190611799565b60405180910390f35b3480156101a8575f80fd5b506101b1610604565b6040516101be91906117c1565b60405180910390f35b3480156101d2575f80fd5b506101db610639565b6040516101e891906117fa565b60405180910390f35b3480156101fc575f80fd5b5061020561065c565b6040516102129190611833565b60405180910390f35b348015610226575f80fd5b5061022f610681565b60405161023c91906118c6565b60405180910390f35b348015610250575f80fd5b5061026b600480360381019061026691906118e6565b6106ba565b6040516102789190611947565b60405180910390f35b34801561028c575f80fd5b506102a760048036038101906102a29190611960565b610b4a565b005b3480156102b4575f80fd5b506102bd610bce565b6040516102ca91906118c6565b60405180910390f35b5f6102dc610c07565b90505f815f0160089054906101000a900460ff161590505f825f015f9054906101000a900467ffffffffffffffff1690505f808267ffffffffffffffff161480156103245750825b90505f60018367ffffffffffffffff1614801561035757505f3073ffffffffffffffffffffffffffffffffffffffff163b145b905081158015610365575080155b1561039c576040517ff92ee8a900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6001855f015f6101000a81548167ffffffffffffffff021916908367ffffffffffffffff16021790555083156103e9576001855f0160086101000a81548160ff0219169083151502179055505b6103f286610c1a565b895f806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508860015f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508760025f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508660035f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550831561054b575f855f0160086101000a81548160ff0219169083151502179055507fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2600160405161054291906119d7565b60405180910390a15b50505050505050505050565b61055f610c2e565b61056882610d14565b6105728282610d1f565b5050565b5f61057f610e3d565b7f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc5f1b905090565b6105af610ec4565b6105b85f610f4b565b565b60025f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b60035f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b5f8061060e61101c565b9050805f015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1691505090565b5f8054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b60015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b6040518060400160405280600581526020017f352e302e3000000000000000000000000000000000000000000000000000000081525081565b5f805f8054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16639fa6a6e36040518163ffffffff1660e01b8152600401602060405180830381865afa158015610724573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906107489190611a1a565b90505f60015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16631ce85e7c856040518263ffffffff1660e01b81526004016107a59190611a97565b602060405180830381865afa1580156107c0573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906107e49190611a1a565b90505f8167ffffffffffffffff1603610801575f92505050610b45565b5f805f60015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16636aadfac7856040518263ffffffff1660e01b815260040161085e9190611ac6565b606060405180830381865afa158015610879573d5f803e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061089d9190611b15565b9250925092508467ffffffffffffffff168167ffffffffffffffff16116108cb575f95505050505050610b45565b6002856108d89190611b92565b94505f805f90505b8460ff168160ff161015610b1d575f8060015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663ee91365b89856040518363ffffffff1660e01b815260040161094c929190611bdc565b6040805180830381865afa158015610966573d5f803e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061098a9190611c17565b915091505f8267ffffffffffffffff16036109a6575050610b10565b5f60025f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166383f036278a866040518363ffffffff1660e01b8152600401610a03929190611bdc565b602060405180830381865afa158015610a1e573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610a429190611c7f565b90508015610a5257505050610b10565b60035f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663df6387fa838c6040518363ffffffff1660e01b8152600401610aae929190611caa565b602060405180830381865afa158015610ac9573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610aed9190611c7f565b90508015610afd57505050610b10565b600185610b0a9190611cd1565b94505050505b80806001019150506108e0565b508260ff168160ff161015610b3a575f9650505050505050610b45565b600196505050505050505b919050565b610b52610ec4565b5f73ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1603610bc2575f6040517f1e4fbdf7000000000000000000000000000000000000000000000000000000008152600401610bb991906117c1565b60405180910390fd5b610bcb81610f4b565b50565b6040518060400160405280600581526020017f322e302e3000000000000000000000000000000000000000000000000000000081525081565b5f80610c11611043565b90508091505090565b610c2261106c565b610c2b816110ac565b50565b7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff163073ffffffffffffffffffffffffffffffffffffffff161480610cdb57507f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff16610cc2611130565b73ffffffffffffffffffffffffffffffffffffffff1614155b15610d12576040517fe07c8dba00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b565b610d1c610ec4565b50565b8173ffffffffffffffffffffffffffffffffffffffff166352d1902d6040518163ffffffff1660e01b8152600401602060405180830381865afa925050508015610d8757506040513d601f19601f82011682018060405250810190610d849190611d2f565b60015b610dc857816040517f4c9c8ce3000000000000000000000000000000000000000000000000000000008152600401610dbf91906117c1565b60405180910390fd5b7f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc5f1b8114610e2e57806040517faa1d49a4000000000000000000000000000000000000000000000000000000008152600401610e2591906116ec565b60405180910390fd5b610e388383611183565b505050565b7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff163073ffffffffffffffffffffffffffffffffffffffff1614610ec2576040517fe07c8dba00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b565b610ecc6111f5565b73ffffffffffffffffffffffffffffffffffffffff16610eea610604565b73ffffffffffffffffffffffffffffffffffffffff1614610f4957610f0d6111f5565b6040517f118cdaa7000000000000000000000000000000000000000000000000000000008152600401610f4091906117c1565b60405180910390fd5b565b5f610f5461101c565b90505f815f015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905082825f015f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508273ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a3505050565b5f7f9016d09d72d40fdae2fd8ceac6b6234c7706214fd39c1cd1e609a0528c199300905090565b5f7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a005f1b905090565b6110746111fc565b6110aa576040517fd7e6bcf800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b565b6110b461106c565b5f73ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1603611124575f6040517f1e4fbdf700000000000000000000000000000000000000000000000000000000815260040161111b91906117c1565b60405180910390fd5b61112d81610f4b565b50565b5f61115c7f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc5f1b61121a565b5f015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b61118c82611223565b8173ffffffffffffffffffffffffffffffffffffffff167fbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b60405160405180910390a25f815111156111e8576111e282826112ec565b506111f1565b6111f06113dd565b5b5050565b5f33905090565b5f611205610c07565b5f0160089054906101000a900460ff16905090565b5f819050919050565b5f8173ffffffffffffffffffffffffffffffffffffffff163b0361127e57806040517f4c9c8ce300000000000000000000000000000000000000000000000000000000815260040161127591906117c1565b60405180910390fd5b806112aa7f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc5f1b61121a565b5f015f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050565b60605f6112f98484611419565b905080801561132f57505f61130c61142d565b118061132e57505f8473ffffffffffffffffffffffffffffffffffffffff163b115b5b156113445761133c611434565b9150506113d7565b801561138757836040517f9996b31500000000000000000000000000000000000000000000000000000000815260040161137e91906117c1565b60405180910390fd5b5f61139061142d565b11156113a35761139e611451565b6113d5565b6040517fd6bda27500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b505b92915050565b5f341115611417576040517fb398979f00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b565b5f805f835160208501865af4905092915050565b5f3d905090565b606060405190503d81523d5f602083013e3d602001810160405290565b6040513d5f823e3d81fd5b5f604051905090565b5f80fd5b5f80fd5b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f6114968261146d565b9050919050565b6114a68161148c565b81146114b0575f80fd5b50565b5f813590506114c18161149d565b92915050565b5f805f805f60a086880312156114e0576114df611465565b5b5f6114ed888289016114b3565b95505060206114fe888289016114b3565b945050604061150f888289016114b3565b9350506060611520888289016114b3565b9250506080611531888289016114b3565b9150509295509295909350565b5f80fd5b5f80fd5b5f601f19601f8301169050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b61158c82611546565b810181811067ffffffffffffffff821117156115ab576115aa611556565b5b80604052505050565b5f6115bd61145c565b90506115c98282611583565b919050565b5f67ffffffffffffffff8211156115e8576115e7611556565b5b6115f182611546565b9050602081019050919050565b828183375f83830152505050565b5f61161e611619846115ce565b6115b4565b90508281526020810184848401111561163a57611639611542565b5b6116458482856115fe565b509392505050565b5f82601f8301126116615761166061153e565b5b813561167184826020860161160c565b91505092915050565b5f80604083850312156116905761168f611465565b5b5f61169d858286016114b3565b925050602083013567ffffffffffffffff8111156116be576116bd611469565b5b6116ca8582860161164d565b9150509250929050565b5f819050919050565b6116e6816116d4565b82525050565b5f6020820190506116ff5f8301846116dd565b92915050565b5f819050919050565b5f61172861172361171e8461146d565b611705565b61146d565b9050919050565b5f6117398261170e565b9050919050565b5f61174a8261172f565b9050919050565b61175a81611740565b82525050565b5f6020820190506117735f830184611751565b92915050565b5f6117838261172f565b9050919050565b61179381611779565b82525050565b5f6020820190506117ac5f83018461178a565b92915050565b6117bb8161148c565b82525050565b5f6020820190506117d45f8301846117b2565b92915050565b5f6117e48261172f565b9050919050565b6117f4816117da565b82525050565b5f60208201905061180d5f8301846117eb565b92915050565b5f61181d8261172f565b9050919050565b61182d81611813565b82525050565b5f6020820190506118465f830184611824565b92915050565b5f81519050919050565b5f82825260208201905092915050565b5f5b83811015611883578082015181840152602081019050611868565b5f8484015250505050565b5f6118988261184c565b6118a28185611856565b93506118b2818560208601611866565b6118bb81611546565b840191505092915050565b5f6020820190508181035f8301526118de818461188e565b905092915050565b5f602082840312156118fb576118fa611465565b5b5f82013567ffffffffffffffff81111561191857611917611469565b5b6119248482850161164d565b91505092915050565b5f8115159050919050565b6119418161192d565b82525050565b5f60208201905061195a5f830184611938565b92915050565b5f6020828403121561197557611974611465565b5b5f611982848285016114b3565b91505092915050565b5f819050919050565b5f67ffffffffffffffff82169050919050565b5f6119c16119bc6119b78461198b565b611705565b611994565b9050919050565b6119d1816119a7565b82525050565b5f6020820190506119ea5f8301846119c8565b92915050565b6119f981611994565b8114611a03575f80fd5b50565b5f81519050611a14816119f0565b92915050565b5f60208284031215611a2f57611a2e611465565b5b5f611a3c84828501611a06565b91505092915050565b5f81519050919050565b5f82825260208201905092915050565b5f611a6982611a45565b611a738185611a4f565b9350611a83818560208601611866565b611a8c81611546565b840191505092915050565b5f6020820190508181035f830152611aaf8184611a5f565b905092915050565b611ac081611994565b82525050565b5f602082019050611ad95f830184611ab7565b92915050565b5f60ff82169050919050565b611af481611adf565b8114611afe575f80fd5b50565b5f81519050611b0f81611aeb565b92915050565b5f805f60608486031215611b2c57611b2b611465565b5b5f611b3986828701611b01565b9350506020611b4a86828701611b01565b9250506040611b5b86828701611a06565b9150509250925092565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b5f611b9c82611994565b9150611ba783611994565b9250828203905067ffffffffffffffff811115611bc757611bc6611b65565b5b92915050565b611bd681611adf565b82525050565b5f604082019050611bef5f830185611ab7565b611bfc6020830184611bcd565b9392505050565b5f81519050611c118161149d565b92915050565b5f8060408385031215611c2d57611c2c611465565b5b5f611c3a85828601611a06565b9250506020611c4b85828601611c03565b9150509250929050565b611c5e8161192d565b8114611c68575f80fd5b50565b5f81519050611c7981611c55565b92915050565b5f60208284031215611c9457611c93611465565b5b5f611ca184828501611c6b565b91505092915050565b5f604082019050611cbd5f8301856117b2565b611cca6020830184611ab7565b9392505050565b5f611cdb82611adf565b9150611ce683611adf565b9250828201905060ff811115611cff57611cfe611b65565b5b92915050565b611d0e816116d4565b8114611d18575f80fd5b50565b5f81519050611d2981611d05565b92915050565b5f60208284031215611d4457611d43611465565b5b5f611d5184828501611d1b565b9150509291505056fea264697066735822122009eb075919d2f4cedb082cba1a2a4dcb41a659953bfc0190c85d2774d016436c64736f6c63430008160033",
}

// StatABI is the input ABI used to generate the binding from.
// Deprecated: Use StatMetaData.ABI instead.
var StatABI = StatMetaData.ABI

// StatBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use StatMetaData.Bin instead.
var StatBin = StatMetaData.Bin

// DeployStat deploys a new Ethereum contract, binding an instance of Stat to it.
func DeployStat(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Stat, error) {
	parsed, err := StatMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(StatBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Stat{StatCaller: StatCaller{contract: contract}, StatTransactor: StatTransactor{contract: contract}, StatFilterer: StatFilterer{contract: contract}}, nil
}

// Stat is an auto generated Go binding around an Ethereum contract.
type Stat struct {
	StatCaller     // Read-only binding to the contract
	StatTransactor // Write-only binding to the contract
	StatFilterer   // Log filterer for contract events
}

// StatCaller is an auto generated read-only Go binding around an Ethereum contract.
type StatCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StatTransactor is an auto generated write-only Go binding around an Ethereum contract.
type StatTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StatFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type StatFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StatSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type StatSession struct {
	Contract     *Stat             // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// StatCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type StatCallerSession struct {
	Contract *StatCaller   // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// StatTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type StatTransactorSession struct {
	Contract     *StatTransactor   // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// StatRaw is an auto generated low-level Go binding around an Ethereum contract.
type StatRaw struct {
	Contract *Stat // Generic contract binding to access the raw methods on
}

// StatCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type StatCallerRaw struct {
	Contract *StatCaller // Generic read-only contract binding to access the raw methods on
}

// StatTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type StatTransactorRaw struct {
	Contract *StatTransactor // Generic write-only contract binding to access the raw methods on
}

// NewStat creates a new instance of Stat, bound to a specific deployed contract.
func NewStat(address common.Address, backend bind.ContractBackend) (*Stat, error) {
	contract, err := bindStat(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Stat{StatCaller: StatCaller{contract: contract}, StatTransactor: StatTransactor{contract: contract}, StatFilterer: StatFilterer{contract: contract}}, nil
}

// NewStatCaller creates a new read-only instance of Stat, bound to a specific deployed contract.
func NewStatCaller(address common.Address, caller bind.ContractCaller) (*StatCaller, error) {
	contract, err := bindStat(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &StatCaller{contract: contract}, nil
}

// NewStatTransactor creates a new write-only instance of Stat, bound to a specific deployed contract.
func NewStatTransactor(address common.Address, transactor bind.ContractTransactor) (*StatTransactor, error) {
	contract, err := bindStat(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &StatTransactor{contract: contract}, nil
}

// NewStatFilterer creates a new log filterer instance of Stat, bound to a specific deployed contract.
func NewStatFilterer(address common.Address, filterer bind.ContractFilterer) (*StatFilterer, error) {
	contract, err := bindStat(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &StatFilterer{contract: contract}, nil
}

// bindStat binds a generic wrapper to an already deployed contract.
func bindStat(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := StatMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Stat *StatRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Stat.Contract.StatCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Stat *StatRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Stat.Contract.StatTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Stat *StatRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Stat.Contract.StatTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Stat *StatCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Stat.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Stat *StatTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Stat.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Stat *StatTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Stat.Contract.contract.Transact(opts, method, params...)
}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_Stat *StatCaller) UPGRADEINTERFACEVERSION(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Stat.contract.Call(opts, &out, "UPGRADE_INTERFACE_VERSION")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_Stat *StatSession) UPGRADEINTERFACEVERSION() (string, error) {
	return _Stat.Contract.UPGRADEINTERFACEVERSION(&_Stat.CallOpts)
}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_Stat *StatCallerSession) UPGRADEINTERFACEVERSION() (string, error) {
	return _Stat.Contract.UPGRADEINTERFACEVERSION(&_Stat.CallOpts)
}

// VERSION is a free data retrieval call binding the contract method 0xffa1ad74.
//
// Solidity: function VERSION() view returns(string)
func (_Stat *StatCaller) VERSION(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Stat.contract.Call(opts, &out, "VERSION")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// VERSION is a free data retrieval call binding the contract method 0xffa1ad74.
//
// Solidity: function VERSION() view returns(string)
func (_Stat *StatSession) VERSION() (string, error) {
	return _Stat.Contract.VERSION(&_Stat.CallOpts)
}

// VERSION is a free data retrieval call binding the contract method 0xffa1ad74.
//
// Solidity: function VERSION() view returns(string)
func (_Stat *StatCallerSession) VERSION() (string, error) {
	return _Stat.Contract.VERSION(&_Stat.CallOpts)
}

// CheckPiece is a free data retrieval call binding the contract method 0xc2e8e717.
//
// Solidity: function checkPiece(bytes _pn) view returns(bool)
func (_Stat *StatCaller) CheckPiece(opts *bind.CallOpts, _pn []byte) (bool, error) {
	var out []interface{}
	err := _Stat.contract.Call(opts, &out, "checkPiece", _pn)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// CheckPiece is a free data retrieval call binding the contract method 0xc2e8e717.
//
// Solidity: function checkPiece(bytes _pn) view returns(bool)
func (_Stat *StatSession) CheckPiece(_pn []byte) (bool, error) {
	return _Stat.Contract.CheckPiece(&_Stat.CallOpts, _pn)
}

// CheckPiece is a free data retrieval call binding the contract method 0xc2e8e717.
//
// Solidity: function checkPiece(bytes _pn) view returns(bool)
func (_Stat *StatCallerSession) CheckPiece(_pn []byte) (bool, error) {
	return _Stat.Contract.CheckPiece(&_Stat.CallOpts, _pn)
}

// Epoch is a free data retrieval call binding the contract method 0x900cf0cf.
//
// Solidity: function epoch() view returns(address)
func (_Stat *StatCaller) Epoch(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Stat.contract.Call(opts, &out, "epoch")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Epoch is a free data retrieval call binding the contract method 0x900cf0cf.
//
// Solidity: function epoch() view returns(address)
func (_Stat *StatSession) Epoch() (common.Address, error) {
	return _Stat.Contract.Epoch(&_Stat.CallOpts)
}

// Epoch is a free data retrieval call binding the contract method 0x900cf0cf.
//
// Solidity: function epoch() view returns(address)
func (_Stat *StatCallerSession) Epoch() (common.Address, error) {
	return _Stat.Contract.Epoch(&_Stat.CallOpts)
}

// Eproof is a free data retrieval call binding the contract method 0x81cc0c7a.
//
// Solidity: function eproof() view returns(address)
func (_Stat *StatCaller) Eproof(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Stat.contract.Call(opts, &out, "eproof")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Eproof is a free data retrieval call binding the contract method 0x81cc0c7a.
//
// Solidity: function eproof() view returns(address)
func (_Stat *StatSession) Eproof() (common.Address, error) {
	return _Stat.Contract.Eproof(&_Stat.CallOpts)
}

// Eproof is a free data retrieval call binding the contract method 0x81cc0c7a.
//
// Solidity: function eproof() view returns(address)
func (_Stat *StatCallerSession) Eproof() (common.Address, error) {
	return _Stat.Contract.Eproof(&_Stat.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Stat *StatCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Stat.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Stat *StatSession) Owner() (common.Address, error) {
	return _Stat.Contract.Owner(&_Stat.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Stat *StatCallerSession) Owner() (common.Address, error) {
	return _Stat.Contract.Owner(&_Stat.CallOpts)
}

// Piece is a free data retrieval call binding the contract method 0x95ad21dc.
//
// Solidity: function piece() view returns(address)
func (_Stat *StatCaller) Piece(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Stat.contract.Call(opts, &out, "piece")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Piece is a free data retrieval call binding the contract method 0x95ad21dc.
//
// Solidity: function piece() view returns(address)
func (_Stat *StatSession) Piece() (common.Address, error) {
	return _Stat.Contract.Piece(&_Stat.CallOpts)
}

// Piece is a free data retrieval call binding the contract method 0x95ad21dc.
//
// Solidity: function piece() view returns(address)
func (_Stat *StatCallerSession) Piece() (common.Address, error) {
	return _Stat.Contract.Piece(&_Stat.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_Stat *StatCaller) ProxiableUUID(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Stat.contract.Call(opts, &out, "proxiableUUID")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_Stat *StatSession) ProxiableUUID() ([32]byte, error) {
	return _Stat.Contract.ProxiableUUID(&_Stat.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_Stat *StatCallerSession) ProxiableUUID() ([32]byte, error) {
	return _Stat.Contract.ProxiableUUID(&_Stat.CallOpts)
}

// Rsproof is a free data retrieval call binding the contract method 0x79ca7e0f.
//
// Solidity: function rsproof() view returns(address)
func (_Stat *StatCaller) Rsproof(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Stat.contract.Call(opts, &out, "rsproof")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Rsproof is a free data retrieval call binding the contract method 0x79ca7e0f.
//
// Solidity: function rsproof() view returns(address)
func (_Stat *StatSession) Rsproof() (common.Address, error) {
	return _Stat.Contract.Rsproof(&_Stat.CallOpts)
}

// Rsproof is a free data retrieval call binding the contract method 0x79ca7e0f.
//
// Solidity: function rsproof() view returns(address)
func (_Stat *StatCallerSession) Rsproof() (common.Address, error) {
	return _Stat.Contract.Rsproof(&_Stat.CallOpts)
}

// Initialize is a paid mutator transaction binding the contract method 0x1459457a.
//
// Solidity: function initialize(address _epoch, address _piece, address _rsproof, address _eproof, address initialOwner) returns()
func (_Stat *StatTransactor) Initialize(opts *bind.TransactOpts, _epoch common.Address, _piece common.Address, _rsproof common.Address, _eproof common.Address, initialOwner common.Address) (*types.Transaction, error) {
	return _Stat.contract.Transact(opts, "initialize", _epoch, _piece, _rsproof, _eproof, initialOwner)
}

// Initialize is a paid mutator transaction binding the contract method 0x1459457a.
//
// Solidity: function initialize(address _epoch, address _piece, address _rsproof, address _eproof, address initialOwner) returns()
func (_Stat *StatSession) Initialize(_epoch common.Address, _piece common.Address, _rsproof common.Address, _eproof common.Address, initialOwner common.Address) (*types.Transaction, error) {
	return _Stat.Contract.Initialize(&_Stat.TransactOpts, _epoch, _piece, _rsproof, _eproof, initialOwner)
}

// Initialize is a paid mutator transaction binding the contract method 0x1459457a.
//
// Solidity: function initialize(address _epoch, address _piece, address _rsproof, address _eproof, address initialOwner) returns()
func (_Stat *StatTransactorSession) Initialize(_epoch common.Address, _piece common.Address, _rsproof common.Address, _eproof common.Address, initialOwner common.Address) (*types.Transaction, error) {
	return _Stat.Contract.Initialize(&_Stat.TransactOpts, _epoch, _piece, _rsproof, _eproof, initialOwner)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Stat *StatTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Stat.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Stat *StatSession) RenounceOwnership() (*types.Transaction, error) {
	return _Stat.Contract.RenounceOwnership(&_Stat.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Stat *StatTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Stat.Contract.RenounceOwnership(&_Stat.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Stat *StatTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Stat.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Stat *StatSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Stat.Contract.TransferOwnership(&_Stat.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Stat *StatTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Stat.Contract.TransferOwnership(&_Stat.TransactOpts, newOwner)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_Stat *StatTransactor) UpgradeToAndCall(opts *bind.TransactOpts, newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _Stat.contract.Transact(opts, "upgradeToAndCall", newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_Stat *StatSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _Stat.Contract.UpgradeToAndCall(&_Stat.TransactOpts, newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_Stat *StatTransactorSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _Stat.Contract.UpgradeToAndCall(&_Stat.TransactOpts, newImplementation, data)
}

// StatInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the Stat contract.
type StatInitializedIterator struct {
	Event *StatInitialized // Event containing the contract specifics and raw log

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
func (it *StatInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StatInitialized)
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
		it.Event = new(StatInitialized)
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
func (it *StatInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StatInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StatInitialized represents a Initialized event raised by the Stat contract.
type StatInitialized struct {
	Version uint64
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_Stat *StatFilterer) FilterInitialized(opts *bind.FilterOpts) (*StatInitializedIterator, error) {

	logs, sub, err := _Stat.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &StatInitializedIterator{contract: _Stat.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_Stat *StatFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *StatInitialized) (event.Subscription, error) {

	logs, sub, err := _Stat.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StatInitialized)
				if err := _Stat.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_Stat *StatFilterer) ParseInitialized(log types.Log) (*StatInitialized, error) {
	event := new(StatInitialized)
	if err := _Stat.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StatOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Stat contract.
type StatOwnershipTransferredIterator struct {
	Event *StatOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *StatOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StatOwnershipTransferred)
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
		it.Event = new(StatOwnershipTransferred)
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
func (it *StatOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StatOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StatOwnershipTransferred represents a OwnershipTransferred event raised by the Stat contract.
type StatOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Stat *StatFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*StatOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Stat.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &StatOwnershipTransferredIterator{contract: _Stat.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Stat *StatFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *StatOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Stat.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StatOwnershipTransferred)
				if err := _Stat.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_Stat *StatFilterer) ParseOwnershipTransferred(log types.Log) (*StatOwnershipTransferred, error) {
	event := new(StatOwnershipTransferred)
	if err := _Stat.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StatUpgradedIterator is returned from FilterUpgraded and is used to iterate over the raw logs and unpacked data for Upgraded events raised by the Stat contract.
type StatUpgradedIterator struct {
	Event *StatUpgraded // Event containing the contract specifics and raw log

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
func (it *StatUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StatUpgraded)
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
		it.Event = new(StatUpgraded)
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
func (it *StatUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StatUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StatUpgraded represents a Upgraded event raised by the Stat contract.
type StatUpgraded struct {
	Implementation common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterUpgraded is a free log retrieval operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_Stat *StatFilterer) FilterUpgraded(opts *bind.FilterOpts, implementation []common.Address) (*StatUpgradedIterator, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _Stat.contract.FilterLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return &StatUpgradedIterator{contract: _Stat.contract, event: "Upgraded", logs: logs, sub: sub}, nil
}

// WatchUpgraded is a free log subscription operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_Stat *StatFilterer) WatchUpgraded(opts *bind.WatchOpts, sink chan<- *StatUpgraded, implementation []common.Address) (event.Subscription, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _Stat.contract.WatchLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StatUpgraded)
				if err := _Stat.contract.UnpackLog(event, "Upgraded", log); err != nil {
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
func (_Stat *StatFilterer) ParseUpgraded(log types.Log) (*StatUpgraded, error) {
	event := new(StatUpgraded)
	if err := _Stat.contract.UnpackLog(event, "Upgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
