// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package add

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

// PlonkVerifierMetaData contains all meta data concerning the PlonkVerifier contract.
var PlonkVerifierMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"proof\",\"type\":\"bytes\"},{\"internalType\":\"uint256[]\",\"name\":\"public_inputs\",\"type\":\"uint256[]\"}],\"name\":\"Verify\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x6080604052348015600e575f80fd5b50612aa38061001c5f395ff3fe608060405234801561000f575f80fd5b5060043610610029575f3560e01c80637e4f7a8a1461002d575b5f80fd5b610047600480360381019061004291906129bc565b61005d565b6040516100549190612a54565b60405180910390f35b5f604051610240810161006f846104a1565b61007985856104b5565b61008286610507565b61008b87610524565b5f61009786868a6106d2565b90506100a281610a18565b90506100ae8189610a7f565b90506100ba8189610b14565b60608301517f30644e72e131a029b85045b68181585d2833e84879b9709143e1f593f000000160017f30644e72e131a029b85045b68181585d2833e84879b9709143e1f593f00000010361011385630200000085612895565b08806101c086015261012684888a610b82565b61013185898d610ee3565b7f30644e72e131a029b85045b68181585d2833e84879b9709143e1f593f00000018282089150816101a0880152610166611298565b61016f8c61242c565b6101788c61236c565b6101818c611f67565b61018a8c611a65565b6101938c6117bd565b61019c8c6113ec565b61020087015197506128ef565b6040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601d60248201527f77726f6e67206e756d626572206f66207075626c696320696e707574730000006044820152606481fd5b6040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601260248201527f6572726f72206563206f7065726174696f6e00000000000000000000000000006044820152606481fd5b6040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601860248201527f696e707574732061726520626967676572207468616e207200000000000000006044820152606481fd5b6040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601060248201527f77726f6e672070726f6f662073697a65000000000000000000000000000000006044820152606481fd5b6040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601660248201527f6f70656e696e677320626967676572207468616e2072000000000000000000006044820152606481fd5b6040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152600d60248201527f6572726f722070616972696e67000000000000000000000000000000000000006044820152606481fd5b6040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152600c60248201527f6572726f722076657269667900000000000000000000000000000000000000006044820152606481fd5b6040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601460248201527f6572726f722072616e646f6d2067656e206b7a670000000000000000000000006044820152606481fd5b600281146104b2576104b16101a9565b5b50565b5f5b81811015610502577f30644e72e131a029b85045b68181585d2833e84879b9709143e1f593f0000000833511156104f1576104f0610267565b5b6020830192506001810190506104b7565b505050565b6060600102610300018082146105205761051f6102c6565b5b5050565b61018081017f30644e72e131a029b85045b68181585d2833e84879b9709143e1f593f00000008135111561055b5761055a610325565b5b6101a0820190507f30644e72e131a029b85045b68181585d2833e84879b9709143e1f593f00000008135111561059457610593610325565b5b6101c0820190507f30644e72e131a029b85045b68181585d2833e84879b9709143e1f593f0000000813511156105cd576105cc610325565b5b6101e0820190507f30644e72e131a029b85045b68181585d2833e84879b9709143e1f593f00000008135111561060657610605610325565b5b610200820190507f30644e72e131a029b85045b68181585d2833e84879b9709143e1f593f00000008135111561063f5761063e610325565b5b610260820190507f30644e72e131a029b85045b68181585d2833e84879b9709143e1f593f00000008135111561067857610677610325565b5b610300820190505f5b60018110156106cd577f30644e72e131a029b85045b68181585d2833e84879b9709143e1f593f0000000823511156106bc576106bb610325565b5b602082019150600181019050610681565b505050565b5f60405161024081016467616d6d6181527f2202ae9749401ff80f200f5503fcf7cb703f260ab5cdf53c6ac1575b16a6f2af60208201527f184bc2816f03dd34395e8e9c655823a2107c09e11b52f85e6f3036e5c4b0eab860408201527f19327f815af09aedb8dd29d4cb469ca68efa66908892800e710685e787c4d57b60608201527e3743e437eebf7b8bba489b1107723f167f190265794db3a8579a213941dd5160808201527f0dcb39983eefc79002d3831d1d0c4f364c0f3baa4f36296912255e4f3c934b1260a08201527f1fc809d49edf6dad23bbe2432fa9194e15ebc1a9884fc9712e95eb9c0cf2978f60c08201527f1cfa73b5fe54f37e6ea2bbc833c5ead8c2955977d301082f291a983197744ad360e08201527f25d821f6f59b0b09619c2fdc351b414d2d84d813691dad5f9263861f4890824f6101008201527f2306e3a4b6cfef7968ab23948bc6fc61ace6c4eaa6cfa5db86d6d70751d5b2686101208201527f27197bb086e59e480c24d178591c3a740df918b513b1e83c98adbf067dc79bae6101408201527f0d1111111d9b137a3b39f21e8cd799952eba0f7dd8333611e0035fd1349f34756101608201527f0c221cc28b9a07f9bd9631813507ac41ce2a2ec9d34e2e65baf8dc9232b3f9036101808201527f2ce433a5c59044073404dd2a195e5ff8e78ea910b902ddfccbf332deea64ce986101a08201527f16823c42ca4da6ffa758eb5fd1fe638b42bcc141e29070855c6fea185be35d446101c08201527f186d0b17a5f25bac608065334595513b8c66d7f134b4ab659eec839d991271f96101e08201527f171f4834df6d206fafd32cb2cc9f3fad7cd7a051e174c30259a9f6b281b5cd686102008201527f20de84e793641606a44bc3e44308ec88dfb2cd8e4238f9777ad6b9cea5e35f426102208201527f20aedafc165518cdebf1ac75e4c2be5a5d0675fc824d05e5a89b84c1ab726c8961024082015261026081016020860280888337808201915060c0808784378083019250816102c50160406001028101905060208582601b880160025afa806109de576109dd6103e3565b5b855197507f30644e72e131a029b85045b68181585d2833e84879b9709143e1f593f000000188066040880152505050505050509392505050565b5f60405161024060405101636265746181528360208201526020816024601c840160025afa80610a4b57610a4a6103e3565b5b815193507f30644e72e131a029b85045b68181585d2833e84879b9709143e1f593f000000184066020840152505050919050565b5f60405161024060405101606564616c70686182526020820186815260208101905061032086016001604002808284378083019250808401935060406102208901843760208585601b880160025afa80610adc57610adb6103e3565b5b855197507f30644e72e131a029b85045b68181585d2833e84879b9709143e1f593f000000188065f8801525050505050505092915050565b60405161024060405101637a657461815283602082015260c0808401604083013760208160e4601c840160025afa80610b5057610b4f6103e3565b5b81517f30644e72e131a029b85045b68181585d2833e84879b9709143e1f593f000000181066060850152505050505050565b5f60405160608101516101c082015186610b9e81888486610c1b565b5f805b88811015610c0e577f30644e72e131a029b85045b68181585d2833e84879b9709143e1f593f0000001883584510991507f30644e72e131a029b85045b68181585d2833e84879b9709143e1f593f00000018288089650602083019250602088019750600181019050610ba1565b5050505050509392505050565b7f30644e72e131a029b85045b68181585d2833e84879b9709143e1f593f00000017f30644e5aaf0a66b91f8030da595e7d1c6787b9b45fc54c546729acf1ff05360983096001855f5b86811015610d0b577f30644e72e131a029b85045b68181585d2833e84879b9709143e1f593f0000001837f30644e72e131a029b85045b68181585d2833e84879b9709143e1f593f000000103860882527f30644e72e131a029b85045b68181585d2833e84879b9709143e1f593f00000017f2a734ebb326341efa19b0361d9130cd47b26b7488dc6d26eeccd4f3eb878331a84099250602082019150600181019050610c64565b50610d17818789610dd3565b869050600191505f5b86811015610dc9577f30644e72e131a029b85045b68181585d2833e84879b9709143e1f593f0000001837f30644e72e131a029b85045b68181585d2833e84879b9709143e1f593f0000001868551090982526020820191507f30644e72e131a029b85045b68181585d2833e84879b9709143e1f593f00000017f2a734ebb326341efa19b0361d9130cd47b26b7488dc6d26eeccd4f3eb878331a84099250600181019050610d20565b5050505050505050565b600183525f805b83811015610e285781850151828401517f30644e72e131a029b85045b68181585d2833e84879b9709143e1f593f0000001818309905060208401935080848801525050600181019050610dda565b5060208103820191508084019350610e686020850160027f30644e72e131a029b85045b68181585d2833e84879b9709143e1f593f0000001038651612895565b5f5b84811015610edb5760208603955083517f30644e72e131a029b85045b68181585d2833e84879b9709143e1f593f0000001875184098086527f30644e72e131a029b85045b68181585d2833e84879b9709143e1f593f000000182850993506020860395505050600181019050610e6a565b505050505050565b5f60405160608101516101c082015161032085015f80610f098a60208501358535611099565b9150610f1c8a62d33de98b018688610f5a565b90507f30644e72e131a029b85045b68181585d2833e84879b9709143e1f593f000000180828409880896506040830192505050505050509392505050565b5f610f8685857f2a734ebb326341efa19b0361d9130cd47b26b7488dc6d26eeccd4f3eb878331a612895565b7f30644e72e131a029b85045b68181585d2833e84879b9709143e1f593f0000001817f30644e72e131a029b85045b68181585d2833e84879b9709143e1f593f000000103840894507f30644e72e131a029b85045b68181585d2833e84879b9709143e1f593f00000017f30644e5aaf0a66b91f8030da595e7d1c6787b9b45fc54c546729acf1ff053609820990506110428660027f30644e72e131a029b85045b68181585d2833e84879b9709143e1f593f00000010387612895565b94507f30644e72e131a029b85045b68181585d2833e84879b9709143e1f593f000000185820990507f30644e72e131a029b85045b68181585d2833e84879b9709143e1f593f0000001848209915050949350505050565b5f8084525f60208501528160408501528260608501525f6080850153603060818501535f60828501536042608385015360536084850153604260858501536032608685015360326087850153602d608885015360506089850153606c608a850153606f608b850153606e608c850153606b608d850153600b608e850153602084608f8660025afa8061112e5761112d6103e3565b5b8451600160208701536042602187015360536022870153604260238701536032602487015360326025870153602d602687015360506027870153606c6028870153606f6029870153606e602a870153606b602b870153600b602c870153602086602d8860025afa9150816111a5576111a46103e3565b5b808651186020870152600260408701536042604187015360536042870153604260438701536032604487015360326045870153602d604687015360506047870153606c6048870153606f6049870153606e604a870153606b604b870153600b604c87015360208601602081602d8360025afa925082611227576112266103e3565b5b7f30644e72e131a029b85045b68181585d2833e84879b9709143e1f593f00000017001000000000000000000000000000000008851099350602087015160801c7f30644e72e131a029b85045b68181585d2833e84879b9709143e1f593f00000018186089450505050509392505050565b604051610240604051016101c08201517f30644e72e131a029b85045b68181585d2833e84879b9709143e1f593f000000160017f30644e72e131a029b85045b68181585d2833e84879b9709143e1f593f0000001036060850151086113218360027f30644e72e131a029b85045b68181585d2833e84879b9709143e1f593f00000010383612895565b90507f30644e72e131a029b85045b68181585d2833e84879b9709143e1f593f00000017f30644e5aaf0a66b91f8030da595e7d1c6787b9b45fc54c546729acf1ff053609820990507f30644e72e131a029b85045b68181585d2833e84879b9709143e1f593f000000182820991505f8401517f30644e72e131a029b85045b68181585d2833e84879b9709143e1f593f000000181840992507f30644e72e131a029b85045b68181585d2833e84879b9709143e1f593f000000181840992508260808601525050505050565b60405161024081016101608201518152610180820151602082015261028083013560408201526102a08301356060820152610220830135608082015261024083013560a08201526102c083013560c08201526102e083013560e082015260608201516101008201526101e08201516101208201526020816101408360025afa8061147957611478610442565b5b7f30644e72e131a029b85045b68181585d2833e84879b9709143e1f593f000000182510690508160408101925061028085013581526102a085013560208201526114c983836102c08801846127ed565b61016084016114de84846102208901846127ed565b61014085016114f284610260890183612843565b8460408101955060018152600260208201528151604082015260408160608360075afa80611523576115226103e3565b5b6020820180517f30644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd47038152611559888487886126b9565b876040890198506115748960608c01516102808e0184612768565b7f30644e72e131a029b85045b68181585d2833e84879b9709143e1f593f00000017f2a734ebb326341efa19b0361d9130cd47b26b7488dc6d26eeccd4f3eb878331a60608c0151097f30644e72e131a029b85045b68181585d2833e84879b9709143e1f593f0000001818a0998506115f28a8a6102c08f01856127ed565b6115fe8a83898a6126b9565b6020880180517f30644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd4703815287518b52602088015160208c01527f198e9393920d483a7260bfb731fb5d25f1aa493335a9e71297e485b7aef312c260408c01527f1800deef121f1e76426a00665e5c4479674322d4f75edadd46debd5cd992f6ed60608c01527f090689d0585ff075ec9e99ad690c3395bc4b313370b38ef355acdadcd122975b60808c01527f12c85ea5db8c6deb4aab71808dcb408fe3d1e7690c43d37b4ce6cc0166fa7daa60a08c0152885160c08c0152602089015160e08c01527ef30dd7827a6227853d4e57718c1d58d9bbd944cf7b6413c4e3b830c69e20126101008c01527f1ff585af5763615da952b63dd7752d725a9f3936111f099466996af171a90ef26101208c01527f047bc7e18e26e6453665d28c64153bebc3becc11898a6cbd297cb63d63bbf87e6101408c01527f0108ad3114af3e6761a8099fb001c3f2ff7f4b31bad965a08c627c510ea30f2d6101608c01526117838b611792565b50505050505050505050505050565b60405160205f6101808460085afa806117ae576117ad610384565b5b5f518061020084015250505050565b6040516102406040510160208101604082016101e084015180610160860160e087015161016088015261010087015161018088015261012087015161014088015261181086835f8b016101608b016127ed565b611823826101808a016101408a01612843565b7f30644e72e131a029b85045b68181585d2833e84879b9709143e1f593f00000018383099150611858868360408b01846127ed565b61186b826101a08a016101408a01612843565b7f30644e72e131a029b85045b68181585d2833e84879b9709143e1f593f000000183830991506118a0868360808b01846127ed565b6118b3826101c08a016101408a01612843565b7f30644e72e131a029b85045b68181585d2833e84879b9709143e1f593f000000183830991507f2202ae9749401ff80f200f5503fcf7cb703f260ab5cdf53c6ac1575b16a6f2af86527f184bc2816f03dd34395e8e9c655823a2107c09e11b52f85e6f3036e5c4b0eab8855261192b8483888461279b565b61193e826101e08a016101408a01612843565b7f30644e72e131a029b85045b68181585d2833e84879b9709143e1f593f000000183830991507f19327f815af09aedb8dd29d4cb469ca68efa66908892800e710685e787c4d57b86527e3743e437eebf7b8bba489b1107723f167f190265794db3a8579a213941dd5185526119b58483888461279b565b6119c8826102008a016101408a01612843565b61030088017f30644e72e131a029b85045b68181585d2833e84879b9709143e1f593f000000184840992507f20de84e793641606a44bc3e44308ec88dfb2cd8e4238f9777ad6b9cea5e35f4287527f20aedafc165518cdebf1ac75e4c2be5a5d0675fc824d05e5a89b84c1ab726c898652611a458584898561279b565b611a5483826101408b01612843565b602081019050505050505050505050565b604051610240604051016467616d6d6181526060820151602082015260e08201516040820152610100820151606082015260c05f840160808301377f2202ae9749401ff80f200f5503fcf7cb703f260ab5cdf53c6ac1575b16a6f2af6101408201527f184bc2816f03dd34395e8e9c655823a2107c09e11b52f85e6f3036e5c4b0eab86101608201527f19327f815af09aedb8dd29d4cb469ca68efa66908892800e710685e787c4d57b6101808201527e3743e437eebf7b8bba489b1107723f167f190265794db3a8579a213941dd516101a08201526101c07f20de84e793641606a44bc3e44308ec88dfb2cd8e4238f9777ad6b9cea5e35f42818301527f20aedafc165518cdebf1ac75e4c2be5a5d0675fc824d05e5a89b84c1ab726c896020820183015260408101905061012083015181830152610180840135602082018301526101a0840135604082018301526101c0840135606082018301526101e08401356080820183015261020084013560a0820183015260c081018201610300850160206001028183376020600102820191506102608601358252601b600360010260140160208102600501905060206101e088018284890160025afa80611c3057611c2f6103e3565b5b7f30644e72e131a029b85045b68181585d2833e84879b9709143e1f593f00000016101e0890151066101e0890152505050505050505050565b604051610240604051017f1cfa73b5fe54f37e6ea2bbc833c5ead8c2955977d301082f291a983197744ad381527f25d821f6f59b0b09619c2fdc351b414d2d84d813691dad5f9263861f4890824f6020820152611cd3604082016101808501358360e08601612735565b7f2306e3a4b6cfef7968ab23948bc6fc61ace6c4eaa6cfa5db86d6d70751d5b26881527f27197bb086e59e480c24d178591c3a740df918b513b1e83c98adbf067dc79bae6020820152611d33604082016101a08501358360e0860161279b565b7f30644e72e131a029b85045b68181585d2833e84879b9709143e1f593f00000016101a0840135610180850135097f0d1111111d9b137a3b39f21e8cd799952eba0f7dd8333611e0035fd1349f347582527f0c221cc28b9a07f9bd9631813507ac41ce2a2ec9d34e2e65baf8dc9232b3f9036020830152611dbc60408301828460e0870161279b565b7f2ce433a5c59044073404dd2a195e5ff8e78ea910b902ddfccbf332deea64ce9882527f16823c42ca4da6ffa758eb5fd1fe638b42bcc141e29070855c6fea185be35d446020830152611e1c604083016101c08601358460e0870161279b565b7f186d0b17a5f25bac608065334595513b8c66d7f134b4ab659eec839d991271f982527f171f4834df6d206fafd32cb2cc9f3fad7cd7a051e174c30259a9f6b281b5cd686020830152611e7a604083018360e0860160e087016126b9565b610300840161032085015f5b6001811015611ec7578135855260208201356020860152611eb06040860184358760e08a0161279b565b602083019250604082019150600181019050611e86565b507f0dcb39983eefc79002d3831d1d0c4f364c0f3baa4f36296912255e4f3c934b1284527f1fc809d49edf6dad23bbe2432fa9194e15ebc1a9884fc9712e95eb9c0cf2978f6020850152611f2360408501888660e0890161279b565b61022086013584526102408601356020850152611f4860408501898660e0890161279b565b611f5d8460a0870160e0880160e089016126b9565b5050505050505050565b6040516020810151604082015160608301515f8401517f30644e72e131a029b85045b68181585d2833e84879b9709143e1f593f000000184610260880135097f30644e72e131a029b85045b68181585d2833e84879b9709143e1f593f00000016101e088013586097f30644e72e131a029b85045b68181585d2833e84879b9709143e1f593f0000001610180890135820890507f30644e72e131a029b85045b68181585d2833e84879b9709143e1f593f000000185820890507f30644e72e131a029b85045b68181585d2833e84879b9709143e1f593f000000161020089013587097f30644e72e131a029b85045b68181585d2833e84879b9709143e1f593f00000016101a08a0135820890507f30644e72e131a029b85045b68181585d2833e84879b9709143e1f593f000000186820890507f30644e72e131a029b85045b68181585d2833e84879b9709143e1f593f00000018284097f30644e72e131a029b85045b68181585d2833e84879b9709143e1f593f000000182820990507f30644e72e131a029b85045b68181585d2833e84879b9709143e1f593f000000185820990507f30644e72e131a029b85045b68181585d2833e84879b9709143e1f593f0000001600580097f30644e72e131a029b85045b68181585d2833e84879b9709143e1f593f0000001878a097f30644e72e131a029b85045b68181585d2833e84879b9709143e1f593f00000016101808d0135820895507f30644e72e131a029b85045b68181585d2833e84879b9709143e1f593f000000189870895507f30644e72e131a029b85045b68181585d2833e84879b9709143e1f593f00000016005820994507f30644e72e131a029b85045b68181585d2833e84879b9709143e1f593f00000016101a08d0135860894507f30644e72e131a029b85045b68181585d2833e84879b9709143e1f593f000000189860894507f30644e72e131a029b85045b68181585d2833e84879b9709143e1f593f000000182820993507f30644e72e131a029b85045b68181585d2833e84879b9709143e1f593f00000016101c08d0135850893507f30644e72e131a029b85045b68181585d2833e84879b9709143e1f593f000000189850893507f30644e72e131a029b85045b68181585d2833e84879b9709143e1f593f00000018587097f30644e72e131a029b85045b68181585d2833e84879b9709143e1f593f00000018582099050807f30644e72e131a029b85045b68181585d2833e84879b9709143e1f593f00000010390507f30644e72e131a029b85045b68181585d2833e84879b9709143e1f593f000000188820990507f30644e72e131a029b85045b68181585d2833e84879b9709143e1f593f000000160808d01518208905061235d81858f611c69565b50505050505050505050505050565b60405160026302000000016102406040510161238d81836060860151612895565b6123a08282610140880160a08801612768565b6123b682610100870160a0870160a088016126f7565b6123c8828260a0870160a08801612735565b6123dd8260c0870160a0870160a088016126f7565b6123f4826101c086015160a0870160a08801612735565b60c0840151807f30644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd470390508060c0860152505050505050565b6040515f7f30644e72e131a029b85045b68181585d2833e84879b9709143e1f593f000000160208301516101e08501350990507f30644e72e131a029b85045b68181585d2833e84879b9709143e1f593f00000016040830151820890507f30644e72e131a029b85045b68181585d2833e84879b9709143e1f593f0000001610180840135820890505f7f30644e72e131a029b85045b68181585d2833e84879b9709143e1f593f000000160208401516102008601350990507f30644e72e131a029b85045b68181585d2833e84879b9709143e1f593f00000016040840151820890507f30644e72e131a029b85045b68181585d2833e84879b9709143e1f593f00000016101a0850135820890505f7f30644e72e131a029b85045b68181585d2833e84879b9709143e1f593f000000160408501516101c08701350890507f30644e72e131a029b85045b68181585d2833e84879b9709143e1f593f000000182840992507f30644e72e131a029b85045b68181585d2833e84879b9709143e1f593f000000181840992507f30644e72e131a029b85045b68181585d2833e84879b9709143e1f593f00000015f850151840992507f30644e72e131a029b85045b68181585d2833e84879b9709143e1f593f0000001610260860135840992507f30644e72e131a029b85045b68181585d2833e84879b9709143e1f593f00000016101a08501518408925060808401519150817f30644e72e131a029b85045b68181585d2833e84879b9709143e1f593f00000010391507f30644e72e131a029b85045b68181585d2833e84879b9709143e1f593f00000018284089250827f30644e72e131a029b85045b68181585d2833e84879b9709143e1f593f0000001039250826101208501525050505050565b8151845260208201516020850152825160408501526020830151606085015260408160808660065afa806126f0576126ef610208565b5b5050505050565b8151845260208201516020850152823560408501526020830135606085015260408160808660065afa8061272e5761272d610208565b5b5050505050565b815184526020820151602085015282604085015260408160608660075afa8061276157612760610208565b5b5050505050565b813584526020820135602085015282604085015260408160608660075afa8061279457612793610208565b5b5050505050565b815184526020820151602085015282604085015260408460608660075afa815160408601526020820151606086015260408260808760065afa81169050806127e6576127e5610208565b5b5050505050565b604051823585526020830135602086015283604086015260408560608760075afa825160408701526020830151606087015260408360808860065afa811690508061283b5761283a610208565b5b505050505050565b7f30644e72e131a029b85045b68181585d2833e84879b9709143e1f593f0000001838335097f30644e72e131a029b85045b68181585d2833e84879b9709143e1f593f000000181835108825250505050565b5f60208452602080850152602060408501528160608501528260808501527f30644e72e131a029b85045b68181585d2833e84879b9709143e1f593f000000160a085015260208460c08660055afa84519150509392505050565b50505050505050949350505050565b5f80fd5b5f80fd5b5f80fd5b5f80fd5b5f80fd5b5f8083601f84011261292757612926612906565b5b8235905067ffffffffffffffff8111156129445761294361290a565b5b6020830191508360018202830111156129605761295f61290e565b5b9250929050565b5f8083601f84011261297c5761297b612906565b5b8235905067ffffffffffffffff8111156129995761299861290a565b5b6020830191508360208202830111156129b5576129b461290e565b5b9250929050565b5f805f80604085870312156129d4576129d36128fe565b5b5f85013567ffffffffffffffff8111156129f1576129f0612902565b5b6129fd87828801612912565b9450945050602085013567ffffffffffffffff811115612a2057612a1f612902565b5b612a2c87828801612967565b925092505092959194509250565b5f8115159050919050565b612a4e81612a3a565b82525050565b5f602082019050612a675f830184612a45565b9291505056fea2646970667358221220e81e6df87cdce883062e711af50d0a694b79175d6fa953b970c7ab078054310964736f6c634300081a0033",
}

// PlonkVerifierABI is the input ABI used to generate the binding from.
// Deprecated: Use PlonkVerifierMetaData.ABI instead.
var PlonkVerifierABI = PlonkVerifierMetaData.ABI

// PlonkVerifierBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use PlonkVerifierMetaData.Bin instead.
var PlonkVerifierBin = PlonkVerifierMetaData.Bin

// DeployPlonkVerifier deploys a new Ethereum contract, binding an instance of PlonkVerifier to it.
func DeployPlonkVerifier(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *PlonkVerifier, error) {
	parsed, err := PlonkVerifierMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(PlonkVerifierBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &PlonkVerifier{PlonkVerifierCaller: PlonkVerifierCaller{contract: contract}, PlonkVerifierTransactor: PlonkVerifierTransactor{contract: contract}, PlonkVerifierFilterer: PlonkVerifierFilterer{contract: contract}}, nil
}

// PlonkVerifier is an auto generated Go binding around an Ethereum contract.
type PlonkVerifier struct {
	PlonkVerifierCaller     // Read-only binding to the contract
	PlonkVerifierTransactor // Write-only binding to the contract
	PlonkVerifierFilterer   // Log filterer for contract events
}

// PlonkVerifierCaller is an auto generated read-only Go binding around an Ethereum contract.
type PlonkVerifierCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PlonkVerifierTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PlonkVerifierTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PlonkVerifierFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PlonkVerifierFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PlonkVerifierSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PlonkVerifierSession struct {
	Contract     *PlonkVerifier    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PlonkVerifierCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PlonkVerifierCallerSession struct {
	Contract *PlonkVerifierCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// PlonkVerifierTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PlonkVerifierTransactorSession struct {
	Contract     *PlonkVerifierTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// PlonkVerifierRaw is an auto generated low-level Go binding around an Ethereum contract.
type PlonkVerifierRaw struct {
	Contract *PlonkVerifier // Generic contract binding to access the raw methods on
}

// PlonkVerifierCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PlonkVerifierCallerRaw struct {
	Contract *PlonkVerifierCaller // Generic read-only contract binding to access the raw methods on
}

// PlonkVerifierTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PlonkVerifierTransactorRaw struct {
	Contract *PlonkVerifierTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPlonkVerifier creates a new instance of PlonkVerifier, bound to a specific deployed contract.
func NewPlonkVerifier(address common.Address, backend bind.ContractBackend) (*PlonkVerifier, error) {
	contract, err := bindPlonkVerifier(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &PlonkVerifier{PlonkVerifierCaller: PlonkVerifierCaller{contract: contract}, PlonkVerifierTransactor: PlonkVerifierTransactor{contract: contract}, PlonkVerifierFilterer: PlonkVerifierFilterer{contract: contract}}, nil
}

// NewPlonkVerifierCaller creates a new read-only instance of PlonkVerifier, bound to a specific deployed contract.
func NewPlonkVerifierCaller(address common.Address, caller bind.ContractCaller) (*PlonkVerifierCaller, error) {
	contract, err := bindPlonkVerifier(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PlonkVerifierCaller{contract: contract}, nil
}

// NewPlonkVerifierTransactor creates a new write-only instance of PlonkVerifier, bound to a specific deployed contract.
func NewPlonkVerifierTransactor(address common.Address, transactor bind.ContractTransactor) (*PlonkVerifierTransactor, error) {
	contract, err := bindPlonkVerifier(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PlonkVerifierTransactor{contract: contract}, nil
}

// NewPlonkVerifierFilterer creates a new log filterer instance of PlonkVerifier, bound to a specific deployed contract.
func NewPlonkVerifierFilterer(address common.Address, filterer bind.ContractFilterer) (*PlonkVerifierFilterer, error) {
	contract, err := bindPlonkVerifier(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PlonkVerifierFilterer{contract: contract}, nil
}

// bindPlonkVerifier binds a generic wrapper to an already deployed contract.
func bindPlonkVerifier(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := PlonkVerifierMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PlonkVerifier *PlonkVerifierRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PlonkVerifier.Contract.PlonkVerifierCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PlonkVerifier *PlonkVerifierRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PlonkVerifier.Contract.PlonkVerifierTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PlonkVerifier *PlonkVerifierRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PlonkVerifier.Contract.PlonkVerifierTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PlonkVerifier *PlonkVerifierCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PlonkVerifier.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PlonkVerifier *PlonkVerifierTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PlonkVerifier.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PlonkVerifier *PlonkVerifierTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PlonkVerifier.Contract.contract.Transact(opts, method, params...)
}

// Verify is a free data retrieval call binding the contract method 0x7e4f7a8a.
//
// Solidity: function Verify(bytes proof, uint256[] public_inputs) view returns(bool success)
func (_PlonkVerifier *PlonkVerifierCaller) Verify(opts *bind.CallOpts, proof []byte, public_inputs []*big.Int) (bool, error) {
	var out []interface{}
	err := _PlonkVerifier.contract.Call(opts, &out, "Verify", proof, public_inputs)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Verify is a free data retrieval call binding the contract method 0x7e4f7a8a.
//
// Solidity: function Verify(bytes proof, uint256[] public_inputs) view returns(bool success)
func (_PlonkVerifier *PlonkVerifierSession) Verify(proof []byte, public_inputs []*big.Int) (bool, error) {
	return _PlonkVerifier.Contract.Verify(&_PlonkVerifier.CallOpts, proof, public_inputs)
}

// Verify is a free data retrieval call binding the contract method 0x7e4f7a8a.
//
// Solidity: function Verify(bytes proof, uint256[] public_inputs) view returns(bool success)
func (_PlonkVerifier *PlonkVerifierCallerSession) Verify(proof []byte, public_inputs []*big.Int) (bool, error) {
	return _PlonkVerifier.Contract.Verify(&_PlonkVerifier.CallOpts, proof, public_inputs)
}
