// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package eproof

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

// IEProofInitParams is an auto generated low-level Go binding around an user-defined struct.
type IEProofInitParams struct {
	Epoch        common.Address
	Node         common.Address
	Piece        common.Address
	Token        common.Address
	Everify      common.Address
	Base         common.Address
	MinProveTime *big.Int
}

// IEProofProofInfo is an auto generated low-level Go binding around an user-defined struct.
type IEProofProofInfo struct {
	Fake      bool
	Submit    bool
	Chaler    common.Address
	ChalTime  *big.Int
	ProveTime *big.Int
}

// EProofMetaData contains all meta data concerning the EProof contract.
var EProofMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"UPGRADE_INTERFACE_VERSION\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"VERSION\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"base\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"basePenalty\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"chalCom\",\"inputs\":[{\"name\":\"_a\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_ep\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"_qIndex\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"chalKZG\",\"inputs\":[{\"name\":\"_a\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_ep\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"chalOne\",\"inputs\":[{\"name\":\"_a\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_ep\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"_qIndex\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"challenge\",\"inputs\":[{\"name\":\"_a\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_ep\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"_sum\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"check\",\"inputs\":[{\"name\":\"_a\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_ep\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"epoch\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIEpoch\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"everify\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIEVerify\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getEProof\",\"inputs\":[{\"name\":\"_a\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_e\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structIEProof.ProofInfo\",\"components\":[{\"name\":\"fake\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"submit\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"chaler\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"chalTime\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"proveTime\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getStat\",\"inputs\":[{\"name\":\"_a\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_e\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"initialize\",\"inputs\":[{\"name\":\"params\",\"type\":\"tuple\",\"internalType\":\"structIEProof.InitParams\",\"components\":[{\"name\":\"epoch\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"node\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"piece\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"everify\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"base\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"minProveTime\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"initialOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"minProveTime\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"node\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractINode\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"piece\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIPiece\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"proveCom\",\"inputs\":[{\"name\":\"_ep\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"_com\",\"type\":\"bytes[]\",\"internalType\":\"bytes[]\"},{\"name\":\"_proof\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"proveKZG\",\"inputs\":[{\"name\":\"_ep\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"_wroot\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"_proof\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"proveOne\",\"inputs\":[{\"name\":\"_ep\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"_com\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"_proof\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"proxiableUUID\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"renounceOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setMinProveTime\",\"inputs\":[{\"name\":\"_t\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"submit\",\"inputs\":[{\"name\":\"_ep\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"_com\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"_proof\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"token\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIERC20\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"upgradeToAndCall\",\"inputs\":[{\"name\":\"newImplementation\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"event\",\"name\":\"Challenge\",\"inputs\":[{\"name\":\"_s\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"_e\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"},{\"name\":\"_round\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"uint8\"},{\"name\":\"_qi\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"uint8\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Fake\",\"inputs\":[{\"name\":\"_s\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"_e\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Initialized\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"previousOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Prove\",\"inputs\":[{\"name\":\"_s\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"_e\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"},{\"name\":\"_round\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"uint8\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Submit\",\"inputs\":[{\"name\":\"_s\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"_e\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Upgraded\",\"inputs\":[{\"name\":\"implementation\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"AddressEmptyCode\",\"inputs\":[{\"name\":\"target\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ERC1967InvalidImplementation\",\"inputs\":[{\"name\":\"implementation\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ERC1967NonPayable\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"FailedCall\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidInitialization\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotInitializing\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OwnableInvalidOwner\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"OwnableUnauthorizedAccount\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"SafeERC20FailedOperation\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"UUPSUnauthorizedCallContext\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"UUPSUnsupportedProxiableUUID\",\"inputs\":[{\"name\":\"slot\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]}]",
	Bin: "0x60a06040523073ffffffffffffffffffffffffffffffffffffffff1660809073ffffffffffffffffffffffffffffffffffffffff16815250670de0b6b3a76400006007553480156200004f575f80fd5b50620000606200006660201b60201c565b620001ed565b5f620000776200016a60201b60201c565b9050805f0160089054906101000a900460ff1615620000c2576040517ff92ee8a900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b67ffffffffffffffff8016815f015f9054906101000a900467ffffffffffffffff1667ffffffffffffffff1614620001675767ffffffffffffffff815f015f6101000a81548167ffffffffffffffff021916908367ffffffffffffffff1602179055507fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d267ffffffffffffffff6040516200015e9190620001d2565b60405180910390a15b50565b5f806200017c6200018560201b60201c565b90508091505090565b5f7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a005f1b905090565b5f67ffffffffffffffff82169050919050565b620001cc81620001ae565b82525050565b5f602082019050620001e75f830184620001c1565b92915050565b608051615c15620002145f395f81816131f50152818161324a01526134040152615c155ff3fe6080604052600436106101b6575f3560e01c806395ad21dc116100eb578063dd48931511610089578063e70b775e11610063578063e70b775e14610598578063f2fde38b146105c0578063fc0c546a146105e8578063ffa1ad7414610612576101b6565b8063dd4893151461050c578063df6387fa14610534578063e58e8e0514610570576101b6565b8063b7c1e98e116100c5578063b7c1e98e1461046a578063bd42f4c214610492578063ca489a87146104ba578063d70754ec146104e2576101b6565b806395ad21dc146103ee578063ad3cb1cc14610418578063b54753b814610442576101b6565b80635a9b427d11610158578063715018a611610132578063715018a61461035c5780637ecaca98146103725780638da5cb5b1461039a578063900cf0cf146103c4576101b6565b80635a9b427d146102e25780635eb24b671461030a57806362b26a7214610332576101b6565b80634f1ef286116101945780634f1ef286146102485780635001f3b51461026457806352d1902d1461028e5780635712e98c146102b8576101b6565b80630bd26cb5146101ba578063266d12cb146101e45780634ca2a89014610220575b5f80fd5b3480156101c5575f80fd5b506101ce61063c565b6040516101db91906142dd565b60405180910390f35b3480156101ef575f80fd5b5061020a6004803603810190610205919061439e565b610642565b604051610217919061447a565b60405180910390f35b34801561022b575f80fd5b50610246600480360381019061024191906145cf565b6108ec565b005b610262600480360381019061025d9190614657565b610b7c565b005b34801561026f575f80fd5b50610278610b9b565b60405161028591906146c0565b60405180910390f35b348015610299575f80fd5b506102a2610bc0565b6040516102af91906146f1565b60405180910390f35b3480156102c3575f80fd5b506102cc610bf1565b6040516102d991906142dd565b60405180910390f35b3480156102ed575f80fd5b506103086004803603810190610303919061470a565b610bf7565b005b348015610315575f80fd5b50610330600480360381019061032b919061439e565b610dac565b005b34801561033d575f80fd5b506103466111a1565b60405161035391906147d1565b60405180910390f35b348015610367575f80fd5b506103706111c6565b005b34801561037d575f80fd5b506103986004803603810190610393919061439e565b6111d9565b005b3480156103a5575f80fd5b506103ae6112be565b6040516103bb91906146c0565b60405180910390f35b3480156103cf575f80fd5b506103d86112f3565b6040516103e5919061480a565b60405180910390f35b3480156103f9575f80fd5b50610402611318565b60405161040f9190614843565b60405180910390f35b348015610423575f80fd5b5061042c61133b565b60405161043991906148d6565b60405180910390f35b34801561044d575f80fd5b506104686004803603810190610463919061492c565b611374565b005b348015610475575f80fd5b50610490600480360381019061048b9190614a5b565b6117c6565b005b34801561049d575f80fd5b506104b860048036038101906104b3919061492c565b611ae9565b005b3480156104c5575f80fd5b506104e060048036038101906104db9190614ac4565b611f3b565b005b3480156104ed575f80fd5b506104f661219e565b6040516105039190614b50565b60405180910390f35b348015610517575f80fd5b50610532600480360381019061052d9190614b69565b6121c3565b005b34801561053f575f80fd5b5061055a6004803603810190610555919061439e565b6121d5565b6040516105679190614ba3565b60405180910390f35b34801561057b575f80fd5b5061059660048036038101906105919190614c9e565b61224d565b005b3480156105a3575f80fd5b506105be60048036038101906105b991906145cf565b6125af565b005b3480156105cb575f80fd5b506105e660048036038101906105e19190614d26565b612cba565b005b3480156105f3575f80fd5b506105fc612d3e565b6040516106099190614d71565b60405180910390f35b34801561061d575f80fd5b50610626612d63565b60405161063391906148d6565b60405180910390f35b60065481565b61064a614281565b610652614281565b60085f8573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f8467ffffffffffffffff1667ffffffffffffffff1681526020019081526020015f205f015f9054906101000a900460ff16815f01901515908115158152505060085f8573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f8467ffffffffffffffff1667ffffffffffffffff1681526020019081526020015f205f0160019054906101000a900460ff1681602001901515908115158152505060085f8573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f8467ffffffffffffffff1667ffffffffffffffff1681526020019081526020015f205f0160029054906101000a900473ffffffffffffffffffffffffffffffffffffffff16816040019073ffffffffffffffffffffffffffffffffffffffff16908173ffffffffffffffffffffffffffffffffffffffff168152505060085f8573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f8467ffffffffffffffff1667ffffffffffffffff1681526020019081526020015f206001015481606001818152505060085f8573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f8467ffffffffffffffff1667ffffffffffffffff1681526020019081526020015f20600201548160800181815250508091505092915050565b6108f63384612d9c565b60085f3373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f8467ffffffffffffffff1667ffffffffffffffff1681526020019081526020015f206002015460085f3373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f8567ffffffffffffffff1667ffffffffffffffff1681526020019081526020015f2060010154118015610a3a575060065460085f3373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f8567ffffffffffffffff1667ffffffffffffffff1681526020019081526020015f2060010154610a379190614db7565b43105b610a79576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610a7090614e34565b60405180910390fd5b5f60045f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16639b00f959338686866040518563ffffffff1660e01b8152600401610ada9493929190614eb3565b6020604051808303815f875af1158015610af6573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610b1a9190614f18565b9050610b263385612fbf565b3373ffffffffffffffffffffffffffffffffffffffff167f074353199408473b546546e1626a6b8ac62ff72784909ac8d0708a73125484af8583604051610b6e929190614f52565b60405180910390a250505050565b610b846131f3565b610b8d826132d9565b610b9782826132e4565b5050565b60055f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b5f610bc9613402565b7f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc5f1b905090565b60075481565b610c018383613489565b60095f8473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f8367ffffffffffffffff1667ffffffffffffffff1681526020019081526020015f2081604051610c6f9190614fb3565b90815260200160405180910390205f9054906101000a900460ff16610cc9576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610cc090615013565b60405180910390fd5b60045f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16635a9b427d8484846040518463ffffffff1660e01b8152600401610d2793929190615031565b5f604051808303815f87803b158015610d3e575f80fd5b505af1158015610d50573d5f803e3d5ffd5b505050508273ffffffffffffffffffffffffffffffffffffffff167f1cf322418045f261f2817d06ac5cf24890e870ce5863c6434a8893f03d6eae4c8360015f604051610d9f939291906150df565b60405180910390a2505050565b610db68282612d9c565b5f60085f8473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f8367ffffffffffffffff1667ffffffffffffffff1681526020019081526020015f206001015411610e5b576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610e529061515e565b60405180910390fd5b600160085f8473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f8367ffffffffffffffff1667ffffffffffffffff1681526020019081526020015f2060020154610ecb9190614db7565b60085f8473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f8367ffffffffffffffff1667ffffffffffffffff1681526020019081526020015f20600101540361103f5760065460085f8473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f8367ffffffffffffffff1667ffffffffffffffff1681526020019081526020015f2060010154610fa59190614db7565b43111561103e57611039828260085f8673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f8567ffffffffffffffff1667ffffffffffffffff1681526020019081526020015f205f0160029054906101000a900473ffffffffffffffffffffffffffffffffffffffff16613829565b61119d565b5b60085f8373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f8267ffffffffffffffff1667ffffffffffffffff1681526020019081526020015f2060020154600160085f8573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f8467ffffffffffffffff1667ffffffffffffffff1681526020019081526020015f20600101546111139190614db7565b0361119c5760065460085f8473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f8367ffffffffffffffff1667ffffffffffffffff1681526020019081526020015f20600201546111899190614db7565b43111561119b5761119a8282612fbf565b5b5b5b5050565b60045f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b6111ce613a68565b6111d75f613aef565b565b6111e38282613489565b60085f8373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f8267ffffffffffffffff1667ffffffffffffffff1681526020019081526020015f205f0160019054906101000a900460ff1661126757611262828233613829565b6112ba565b8173ffffffffffffffffffffffffffffffffffffffff167f1cf322418045f261f2817d06ac5cf24890e870ce5863c6434a8893f03d6eae4c825f806040516112b19392919061517c565b60405180910390a25b5050565b5f806112c8613bc0565b9050805f015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1691505090565b60015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b5f8054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b6040518060400160405280600581526020017f352e302e3000000000000000000000000000000000000000000000000000000081525081565b61137e8383612d9c565b60085f8473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f8367ffffffffffffffff1667ffffffffffffffff1681526020019081526020015f205f0160029054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161461146e576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401611465906151fb565b60405180910390fd5b60085f8473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f8367ffffffffffffffff1667ffffffffffffffff1681526020019081526020015f206001015460085f8573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f8467ffffffffffffffff1667ffffffffffffffff1681526020019081526020015f20600201541180156115b2575060065460085f8573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f8467ffffffffffffffff1667ffffffffffffffff1681526020019081526020015f20600101546115af9190614db7565b43105b6115f1576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016115e890614e34565b60405180910390fd5b4360085f8573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f8467ffffffffffffffff1667ffffffffffffffff1681526020019081526020015f20600101819055506001436116669190615219565b60085f8573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f8467ffffffffffffffff1667ffffffffffffffff1681526020019081526020015f20600201819055505f60045f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663b54753b88585856040518463ffffffff1660e01b815260040161172c9392919061524c565b6020604051808303815f875af1158015611748573d5f803e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061176c9190614f18565b90508373ffffffffffffffffffffffffffffffffffffffff167f1cf322418045f261f2817d06ac5cf24890e870ce5863c6434a8893f03d6eae4c8483856040516117b893929190615281565b60405180910390a250505050565b5f6117cf613be7565b90505f815f0160089054906101000a900460ff161590505f825f015f9054906101000a900467ffffffffffffffff1690505f808267ffffffffffffffff161480156118175750825b90505f60018367ffffffffffffffff1614801561184a57505f3073ffffffffffffffffffffffffffffffffffffffff163b145b905081158015611858575080155b1561188f576040517ff92ee8a900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6001855f015f6101000a81548167ffffffffffffffff021916908367ffffffffffffffff16021790555083156118dc576001855f0160086101000a81548160ff0219169083151502179055505b6118e586613bfa565b865f015160015f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555086604001515f806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550866060015160025f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550866020015160035f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550866080015160045f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508660a0015160055f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508660c001516006819055508315611ae0575f855f0160086101000a81548160ff0219169083151502179055507fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d26001604051611ad791906152e6565b60405180910390a15b50505050505050565b611af38383612d9c565b60085f8473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f8367ffffffffffffffff1667ffffffffffffffff1681526020019081526020015f205f0160029054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614611be3576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401611bda90615349565b60405180910390fd5b60085f8473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f8367ffffffffffffffff1667ffffffffffffffff1681526020019081526020015f206001015460085f8573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f8467ffffffffffffffff1667ffffffffffffffff1681526020019081526020015f2060020154118015611d27575060065460085f8573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f8467ffffffffffffffff1667ffffffffffffffff1681526020019081526020015f2060020154611d249190614db7565b43105b611d66576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401611d5d90614e34565b60405180910390fd5b4360085f8573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f8467ffffffffffffffff1667ffffffffffffffff1681526020019081526020015f2060010181905550600143611ddb9190615219565b60085f8573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f8467ffffffffffffffff1667ffffffffffffffff1681526020019081526020015f20600201819055505f60045f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663bd42f4c28585856040518463ffffffff1660e01b8152600401611ea19392919061524c565b6020604051808303815f875af1158015611ebd573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190611ee19190614f18565b90508373ffffffffffffffffffffffffffffffffffffffff167f1cf322418045f261f2817d06ac5cf24890e870ce5863c6434a8893f03d6eae4c848385604051611f2d93929190615281565b60405180910390a250505050565b611f453384612d9c565b600a5f3373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f8467ffffffffffffffff1667ffffffffffffffff1681526020019081526020015f205f8381526020019081526020015f205f9054906101000a900460ff16612000576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401611ff7906153b1565b60405180910390fd5b60065460085f3373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f8567ffffffffffffffff1667ffffffffffffffff1681526020019081526020015f20600101546120719190614db7565b43106120b2576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016120a990614e34565b60405180910390fd5b60045f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16632049af34338585856040518563ffffffff1660e01b815260040161211294939291906153cf565b5f604051808303815f87803b158015612129575f80fd5b505af115801561213b573d5f803e3d5ffd5b505050506121493384612fbf565b3373ffffffffffffffffffffffffffffffffffffffff167f074353199408473b546546e1626a6b8ac62ff72784909ac8d0708a73125484af845f604051612191929190615419565b60405180910390a2505050565b60035f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b6121cb613a68565b8060068190555050565b5f60085f8473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f8367ffffffffffffffff1667ffffffffffffffff1681526020019081526020015f205f015f9054906101000a900460ff16905092915050565b6122573384612d9c565b60085f3373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f8467ffffffffffffffff1667ffffffffffffffff1681526020019081526020015f206002015460085f3373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f8567ffffffffffffffff1667ffffffffffffffff1681526020019081526020015f206001015411801561239b575060065460085f3373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f8567ffffffffffffffff1667ffffffffffffffff1681526020019081526020015f20600101546123989190614db7565b43105b6123da576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016123d190614e34565b60405180910390fd5b6001436123e79190615219565b60085f3373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f8567ffffffffffffffff1667ffffffffffffffff1681526020019081526020015f20600101819055504360085f3373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f8567ffffffffffffffff1667ffffffffffffffff1681526020019081526020015f20600201819055505f60045f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16639949edae338686866040518563ffffffff1660e01b81526004016125179493929190615543565b6020604051808303815f875af1158015612533573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906125579190614f18565b90503373ffffffffffffffffffffffffffffffffffffffff167f074353199408473b546546e1626a6b8ac62ff72784909ac8d0708a73125484af85836040516125a1929190614f52565b60405180910390a250505050565b5f8251116125f2576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016125e9906155de565b60405180910390fd5b60015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663919840ad6040518163ffffffff1660e01b81526004015f604051808303815f87803b158015612658575f80fd5b505af115801561266a573d5f803e3d5ffd5b505050505f60015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16639fa6a6e36040518163ffffffff1660e01b8152600401602060405180830381865afa1580156126d9573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906126fd9190615610565b90508067ffffffffffffffff16600185612717919061563b565b67ffffffffffffffff1614612761576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401612758906156c0565b60405180910390fd5b60085f3373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f8567ffffffffffffffff1667ffffffffffffffff1681526020019081526020015f205f0160019054906101000a900460ff1615612811576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161280890615728565b60405180910390fd5b5f8054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166362b98032336040518263ffffffff1660e01b815260040161286991906146c0565b5f604051808303815f87803b158015612880575f80fd5b505af1158015612892573d5f803e3d5ffd5b505050505f805f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166318198ad733876040518363ffffffff1660e01b81526004016128f2929190615746565b6020604051808303815f875af115801561290e573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906129329190615781565b90505f8111612976576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161296d906157f6565b60405180910390fd5b5f8060035f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166361e728b13360016040518363ffffffff1660e01b81526004016129d5929190615814565b6040805180830381865afa1580156129ef573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190612a139190615865565b915091508115612a655782811015612a60576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401612a57906158ed565b60405180910390fd5b612aa0565b6040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401612a9790615955565b60405180910390fd5b600160095f3373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f8967ffffffffffffffff1667ffffffffffffffff1681526020019081526020015f2087604051612b109190614fb3565b90815260200160405180910390205f6101000a81548160ff0219169083151502179055505f8686604051602001612b48929190615973565b6040516020818303038152906040528051906020012090506001600a5f3373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f8a67ffffffffffffffff1667ffffffffffffffff1681526020019081526020015f205f8381526020019081526020015f205f6101000a81548160ff021916908315150217905550600160085f3373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f8a67ffffffffffffffff1667ffffffffffffffff1681526020019081526020015f205f0160016101000a81548160ff0219169083151502179055503373ffffffffffffffffffffffffffffffffffffffff167f98e23f0b72b943d75df730f1b4c0fbb9baffa041d18e1819f2d851a5c2c1c44489604051612ca89190615996565b60405180910390a25050505050505050565b612cc2613a68565b5f73ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1603612d32575f6040517f1e4fbdf7000000000000000000000000000000000000000000000000000000008152600401612d2991906146c0565b60405180910390fd5b612d3b81613aef565b50565b60025f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b6040518060400160405280600581526020017f322e302e3000000000000000000000000000000000000000000000000000000081525081565b60015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663919840ad6040518163ffffffff1660e01b81526004015f604051808303815f87803b158015612e02575f80fd5b505af1158015612e14573d5f803e3d5ffd5b505050505f60015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16639fa6a6e36040518163ffffffff1660e01b8152600401602060405180830381865afa158015612e83573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190612ea79190615610565b90508067ffffffffffffffff16600183612ec1919061563b565b67ffffffffffffffff1610612f0b576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401612f02906159f9565b60405180910390fd5b60085f8473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f8367ffffffffffffffff1667ffffffffffffffff1681526020019081526020015f205f015f9054906101000a900460ff1615612fba576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401612fb190615a61565b60405180910390fd5b505050565b60035f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16637eee288d836007546040518363ffffffff1660e01b815260040161301d929190615a7f565b5f604051808303815f87803b158015613034575f80fd5b505af1158015613046573d5f803e3d5ffd5b505050506130a482600260075461305d9190615ad3565b60025f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16613c0e9092919063ffffffff16565b61311f60055f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1660026007546130d89190615ad3565b60025f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16613c0e9092919063ffffffff16565b5f60085f8473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f8367ffffffffffffffff1667ffffffffffffffff1681526020019081526020015f20600201819055505f60085f8473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f8367ffffffffffffffff1667ffffffffffffffff1681526020019081526020015f20600101819055505050565b7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff163073ffffffffffffffffffffffffffffffffffffffff1614806132a057507f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff16613287613c61565b73ffffffffffffffffffffffffffffffffffffffff1614155b156132d7576040517fe07c8dba00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b565b6132e1613a68565b50565b8173ffffffffffffffffffffffffffffffffffffffff166352d1902d6040518163ffffffff1660e01b8152600401602060405180830381865afa92505050801561334c57506040513d601f19601f820116820180604052508101906133499190615b17565b60015b61338d57816040517f4c9c8ce300000000000000000000000000000000000000000000000000000000815260040161338491906146c0565b60405180910390fd5b7f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc5f1b81146133f357806040517faa1d49a40000000000000000000000000000000000000000000000000000000081526004016133ea91906146f1565b60405180910390fd5b6133fd8383613cb4565b505050565b7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff163073ffffffffffffffffffffffffffffffffffffffff1614613487576040517fe07c8dba00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b565b6134938282612d9c565b5f8054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166362b98032836040518263ffffffff1660e01b81526004016134eb91906146c0565b5f604051808303815f87803b158015613502575f80fd5b505af1158015613514573d5f803e3d5ffd5b505050505f805f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166318198ad784846040518363ffffffff1660e01b8152600401613574929190615746565b6020604051808303815f875af1158015613590573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906135b49190615781565b90505f81116135f8576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016135ef906157f6565b60405180910390fd5b5f60085f8573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f8467ffffffffffffffff1667ffffffffffffffff1681526020019081526020015f20600101541461369d576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161369490615b8c565b60405180910390fd5b6136a78333613d26565b4360085f8573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f8467ffffffffffffffff1667ffffffffffffffff1681526020019081526020015f206001018190555060014361371c9190615219565b60085f8573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f8467ffffffffffffffff1667ffffffffffffffff1681526020019081526020015f20600201819055503360085f8573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f8467ffffffffffffffff1667ffffffffffffffff1681526020019081526020015f205f0160026101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550505050565b60035f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16637eee288d846007546040518363ffffffff1660e01b8152600401613887929190615a7f565b5f604051808303815f87803b15801561389e575f80fd5b505af11580156138b0573d5f803e3d5ffd5b5050505060035f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16639748dcdc848360026007546139029190615ad3565b6040518463ffffffff1660e01b815260040161392093929190615baa565b5f604051808303815f87803b158015613937575f80fd5b505af1158015613949573d5f803e3d5ffd5b5050505061399b8160075460025f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16613c0e9092919063ffffffff16565b600160085f8573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f8467ffffffffffffffff1667ffffffffffffffff1681526020019081526020015f205f015f6101000a81548160ff0219169083151502179055508273ffffffffffffffffffffffffffffffffffffffff167f0fcd3bb0ce91639d74185d625544f2b5714ecd14a6a6d9ccb65b2526e9ac12f583604051613a5b9190615996565b60405180910390a2505050565b613a70613e05565b73ffffffffffffffffffffffffffffffffffffffff16613a8e6112be565b73ffffffffffffffffffffffffffffffffffffffff1614613aed57613ab1613e05565b6040517f118cdaa7000000000000000000000000000000000000000000000000000000008152600401613ae491906146c0565b60405180910390fd5b565b5f613af8613bc0565b90505f815f015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905082825f015f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508273ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a3505050565b5f7f9016d09d72d40fdae2fd8ceac6b6234c7706214fd39c1cd1e609a0528c199300905090565b5f80613bf1613e0c565b90508091505090565b613c02613e35565b613c0b81613e75565b50565b613c1b8383836001613ef9565b613c5c57826040517f5274afe7000000000000000000000000000000000000000000000000000000008152600401613c5391906146c0565b60405180910390fd5b505050565b5f613c8d7f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc5f1b613f5b565b5f015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b613cbd82613f64565b8173ffffffffffffffffffffffffffffffffffffffff167fbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b60405160405180910390a25f81511115613d1957613d13828261402d565b50613d22565b613d2161411e565b5b5050565b613d76813060075460025f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1661415a909392919063ffffffff16565b60035f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663282d3fdf836007546040518363ffffffff1660e01b8152600401613dd4929190615a7f565b5f604051808303815f87803b158015613deb575f80fd5b505af1158015613dfd573d5f803e3d5ffd5b505050505050565b5f33905090565b5f7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a005f1b905090565b613e3d6141af565b613e73576040517fd7e6bcf800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b565b613e7d613e35565b5f73ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1603613eed575f6040517f1e4fbdf7000000000000000000000000000000000000000000000000000000008152600401613ee491906146c0565b60405180910390fd5b613ef681613aef565b50565b5f8063a9059cbb60e01b9050604051815f525f1960601c86166004528460245260205f60445f808b5af1925060015f51148316613f4d578383151615613f41573d5f823e3d81fd5b5f873b113d1516831692505b806040525050949350505050565b5f819050919050565b5f8173ffffffffffffffffffffffffffffffffffffffff163b03613fbf57806040517f4c9c8ce3000000000000000000000000000000000000000000000000000000008152600401613fb691906146c0565b60405180910390fd5b80613feb7f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc5f1b613f5b565b5f015f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050565b60605f61403a84846141cd565b905080801561407057505f61404d6141e1565b118061406f57505f8473ffffffffffffffffffffffffffffffffffffffff163b115b5b156140855761407d6141e8565b915050614118565b80156140c857836040517f9996b3150000000000000000000000000000000000000000000000000000000081526004016140bf91906146c0565b60405180910390fd5b5f6140d16141e1565b11156140e4576140df614205565b614116565b6040517fd6bda27500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b505b92915050565b5f341115614158576040517fb398979f00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b565b614168848484846001614210565b6141a957836040517f5274afe70000000000000000000000000000000000000000000000000000000081526004016141a091906146c0565b60405180910390fd5b50505050565b5f6141b8613be7565b5f0160089054906101000a900460ff16905090565b5f805f835160208501865af4905092915050565b5f3d905090565b606060405190503d81523d5f602083013e3d602001810160405290565b6040513d5f823e3d81fd5b5f806323b872dd60e01b9050604051815f525f1960601c87166004525f1960601c86166024528460445260205f60645f808c5af1925060015f5114831661426e578383151615614262573d5f823e3d81fd5b5f883b113d1516831692505b806040525f606052505095945050505050565b6040518060a001604052805f151581526020015f151581526020015f73ffffffffffffffffffffffffffffffffffffffff1681526020015f81526020015f81525090565b5f819050919050565b6142d7816142c5565b82525050565b5f6020820190506142f05f8301846142ce565b92915050565b5f604051905090565b5f80fd5b5f80fd5b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f61433082614307565b9050919050565b61434081614326565b811461434a575f80fd5b50565b5f8135905061435b81614337565b92915050565b5f67ffffffffffffffff82169050919050565b61437d81614361565b8114614387575f80fd5b50565b5f8135905061439881614374565b92915050565b5f80604083850312156143b4576143b36142ff565b5b5f6143c18582860161434d565b92505060206143d28582860161438a565b9150509250929050565b5f8115159050919050565b6143f0816143dc565b82525050565b6143ff81614326565b82525050565b61440e816142c5565b82525050565b60a082015f8201516144285f8501826143e7565b50602082015161443b60208501826143e7565b50604082015161444e60408501826143f6565b5060608201516144616060850182614405565b5060808201516144746080850182614405565b50505050565b5f60a08201905061448d5f830184614414565b92915050565b5f80fd5b5f80fd5b5f601f19601f8301169050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b6144e18261449b565b810181811067ffffffffffffffff82111715614500576144ff6144ab565b5b80604052505050565b5f6145126142f6565b905061451e82826144d8565b919050565b5f67ffffffffffffffff82111561453d5761453c6144ab565b5b6145468261449b565b9050602081019050919050565b828183375f83830152505050565b5f61457361456e84614523565b614509565b90508281526020810184848401111561458f5761458e614497565b5b61459a848285614553565b509392505050565b5f82601f8301126145b6576145b5614493565b5b81356145c6848260208601614561565b91505092915050565b5f805f606084860312156145e6576145e56142ff565b5b5f6145f38682870161438a565b935050602084013567ffffffffffffffff81111561461457614613614303565b5b614620868287016145a2565b925050604084013567ffffffffffffffff81111561464157614640614303565b5b61464d868287016145a2565b9150509250925092565b5f806040838503121561466d5761466c6142ff565b5b5f61467a8582860161434d565b925050602083013567ffffffffffffffff81111561469b5761469a614303565b5b6146a7858286016145a2565b9150509250929050565b6146ba81614326565b82525050565b5f6020820190506146d35f8301846146b1565b92915050565b5f819050919050565b6146eb816146d9565b82525050565b5f6020820190506147045f8301846146e2565b92915050565b5f805f60608486031215614721576147206142ff565b5b5f61472e8682870161434d565b935050602061473f8682870161438a565b925050604084013567ffffffffffffffff8111156147605761475f614303565b5b61476c868287016145a2565b9150509250925092565b5f819050919050565b5f61479961479461478f84614307565b614776565b614307565b9050919050565b5f6147aa8261477f565b9050919050565b5f6147bb826147a0565b9050919050565b6147cb816147b1565b82525050565b5f6020820190506147e45f8301846147c2565b92915050565b5f6147f4826147a0565b9050919050565b614804816147ea565b82525050565b5f60208201905061481d5f8301846147fb565b92915050565b5f61482d826147a0565b9050919050565b61483d81614823565b82525050565b5f6020820190506148565f830184614834565b92915050565b5f81519050919050565b5f82825260208201905092915050565b5f5b83811015614893578082015181840152602081019050614878565b5f8484015250505050565b5f6148a88261485c565b6148b28185614866565b93506148c2818560208601614876565b6148cb8161449b565b840191505092915050565b5f6020820190508181035f8301526148ee818461489e565b905092915050565b5f60ff82169050919050565b61490b816148f6565b8114614915575f80fd5b50565b5f8135905061492681614902565b92915050565b5f805f60608486031215614943576149426142ff565b5b5f6149508682870161434d565b93505060206149618682870161438a565b925050604061497286828701614918565b9150509250925092565b5f80fd5b614989816142c5565b8114614993575f80fd5b50565b5f813590506149a481614980565b92915050565b5f60e082840312156149bf576149be61497c565b5b6149c960e0614509565b90505f6149d88482850161434d565b5f8301525060206149eb8482850161434d565b60208301525060406149ff8482850161434d565b6040830152506060614a138482850161434d565b6060830152506080614a278482850161434d565b60808301525060a0614a3b8482850161434d565b60a08301525060c0614a4f84828501614996565b60c08301525092915050565b5f806101008385031215614a7257614a716142ff565b5b5f614a7f858286016149aa565b92505060e0614a908582860161434d565b9150509250929050565b614aa3816146d9565b8114614aad575f80fd5b50565b5f81359050614abe81614a9a565b92915050565b5f805f60608486031215614adb57614ada6142ff565b5b5f614ae88682870161438a565b9350506020614af986828701614ab0565b925050604084013567ffffffffffffffff811115614b1a57614b19614303565b5b614b26868287016145a2565b9150509250925092565b5f614b3a826147a0565b9050919050565b614b4a81614b30565b82525050565b5f602082019050614b635f830184614b41565b92915050565b5f60208284031215614b7e57614b7d6142ff565b5b5f614b8b84828501614996565b91505092915050565b614b9d816143dc565b82525050565b5f602082019050614bb65f830184614b94565b92915050565b5f67ffffffffffffffff821115614bd657614bd56144ab565b5b602082029050602081019050919050565b5f80fd5b5f614bfd614bf884614bbc565b614509565b90508083825260208201905060208402830185811115614c2057614c1f614be7565b5b835b81811015614c6757803567ffffffffffffffff811115614c4557614c44614493565b5b808601614c5289826145a2565b85526020850194505050602081019050614c22565b5050509392505050565b5f82601f830112614c8557614c84614493565b5b8135614c95848260208601614beb565b91505092915050565b5f805f60608486031215614cb557614cb46142ff565b5b5f614cc28682870161438a565b935050602084013567ffffffffffffffff811115614ce357614ce2614303565b5b614cef86828701614c71565b925050604084013567ffffffffffffffff811115614d1057614d0f614303565b5b614d1c868287016145a2565b9150509250925092565b5f60208284031215614d3b57614d3a6142ff565b5b5f614d488482850161434d565b91505092915050565b5f614d5b826147a0565b9050919050565b614d6b81614d51565b82525050565b5f602082019050614d845f830184614d62565b92915050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b5f614dc1826142c5565b9150614dcc836142c5565b9250828201905080821115614de457614de3614d8a565b5b92915050565b7f65786365656420707400000000000000000000000000000000000000000000005f82015250565b5f614e1e600983614866565b9150614e2982614dea565b602082019050919050565b5f6020820190508181035f830152614e4b81614e12565b9050919050565b614e5b81614361565b82525050565b5f81519050919050565b5f82825260208201905092915050565b5f614e8582614e61565b614e8f8185614e6b565b9350614e9f818560208601614876565b614ea88161449b565b840191505092915050565b5f608082019050614ec65f8301876146b1565b614ed36020830186614e52565b8181036040830152614ee58185614e7b565b90508181036060830152614ef98184614e7b565b905095945050505050565b5f81519050614f1281614902565b92915050565b5f60208284031215614f2d57614f2c6142ff565b5b5f614f3a84828501614f04565b91505092915050565b614f4c816148f6565b82525050565b5f604082019050614f655f830185614e52565b614f726020830184614f43565b9392505050565b5f81905092915050565b5f614f8d82614e61565b614f978185614f79565b9350614fa7818560208601614876565b80840191505092915050565b5f614fbe8284614f83565b915081905092915050565b7f6e6f2073756d00000000000000000000000000000000000000000000000000005f82015250565b5f614ffd600683614866565b915061500882614fc9565b602082019050919050565b5f6020820190508181035f83015261502a81614ff1565b9050919050565b5f6060820190506150445f8301866146b1565b6150516020830185614e52565b81810360408301526150638184614e7b565b9050949350505050565b5f819050919050565b5f61509061508b6150868461506d565b614776565b6148f6565b9050919050565b6150a081615076565b82525050565b5f819050919050565b5f6150c96150c46150bf846150a6565b614776565b6148f6565b9050919050565b6150d9816150af565b82525050565b5f6060820190506150f25f830186614e52565b6150ff6020830185615097565b61510c60408301846150d0565b949350505050565b7f6e6f206368616c000000000000000000000000000000000000000000000000005f82015250565b5f615148600783614866565b915061515382615114565b602082019050919050565b5f6020820190508181035f8301526151758161513c565b9050919050565b5f60608201905061518f5f830186614e52565b61519c60208301856150d0565b6151a960408301846150d0565b949350505050565b7f696e76206368616c436f6d0000000000000000000000000000000000000000005f82015250565b5f6151e5600b83614866565b91506151f0826151b1565b602082019050919050565b5f6020820190508181035f830152615212816151d9565b9050919050565b5f615223826142c5565b915061522e836142c5565b925082820390508181111561524657615245614d8a565b5b92915050565b5f60608201905061525f5f8301866146b1565b61526c6020830185614e52565b6152796040830184614f43565b949350505050565b5f6060820190506152945f830186614e52565b6152a16020830185614f43565b6152ae6040830184614f43565b949350505050565b5f6152d06152cb6152c68461506d565b614776565b614361565b9050919050565b6152e0816152b6565b82525050565b5f6020820190506152f95f8301846152d7565b92915050565b7f696e76206368616c4f6e650000000000000000000000000000000000000000005f82015250565b5f615333600b83614866565b915061533e826152ff565b602082019050919050565b5f6020820190508181035f83015261536081615327565b9050919050565b7f6e6f2077697400000000000000000000000000000000000000000000000000005f82015250565b5f61539b600683614866565b91506153a682615367565b602082019050919050565b5f6020820190508181035f8301526153c88161538f565b9050919050565b5f6080820190506153e25f8301876146b1565b6153ef6020830186614e52565b6153fc60408301856146e2565b818103606083015261540e8184614e7b565b905095945050505050565b5f60408201905061542c5f830185614e52565b61543960208301846150d0565b9392505050565b5f81519050919050565b5f82825260208201905092915050565b5f819050602082019050919050565b5f82825260208201905092915050565b5f61548382614e61565b61548d8185615469565b935061549d818560208601614876565b6154a68161449b565b840191505092915050565b5f6154bc8383615479565b905092915050565b5f602082019050919050565b5f6154da82615440565b6154e4818561544a565b9350836020820285016154f68561545a565b805f5b85811015615531578484038952815161551285826154b1565b945061551d836154c4565b925060208a019950506001810190506154f9565b50829750879550505050505092915050565b5f6080820190506155565f8301876146b1565b6155636020830186614e52565b818103604083015261557581856154d0565b905081810360608301526155898184614e7b565b905095945050505050565b7f696e7620636f6d000000000000000000000000000000000000000000000000005f82015250565b5f6155c8600783614866565b91506155d382615594565b602082019050919050565b5f6020820190508181035f8301526155f5816155bc565b9050919050565b5f8151905061560a81614374565b92915050565b5f60208284031215615625576156246142ff565b5b5f615632848285016155fc565b91505092915050565b5f61564582614361565b915061565083614361565b9250828201905067ffffffffffffffff8111156156705761566f614d8a565b5b92915050565b7f696e7620657000000000000000000000000000000000000000000000000000005f82015250565b5f6156aa600683614866565b91506156b582615676565b602082019050919050565b5f6020820190508181035f8301526156d78161569e565b9050919050565b7f70726f6f662065786973740000000000000000000000000000000000000000005f82015250565b5f615712600b83614866565b915061571d826156de565b602082019050919050565b5f6020820190508181035f83015261573f81615706565b9050919050565b5f6040820190506157595f8301856146b1565b6157666020830184614e52565b9392505050565b5f8151905061577b81614980565b92915050565b5f60208284031215615796576157956142ff565b5b5f6157a38482850161576d565b91505092915050565b7f6e6f2061637469766500000000000000000000000000000000000000000000005f82015250565b5f6157e0600983614866565b91506157eb826157ac565b602082019050919050565b5f6020820190508181035f83015261580d816157d4565b9050919050565b5f6040820190506158275f8301856146b1565b6158346020830184615097565b9392505050565b615844816143dc565b811461584e575f80fd5b50565b5f8151905061585f8161583b565b92915050565b5f806040838503121561587b5761587a6142ff565b5b5f61588885828601615851565b92505060206158998582860161576d565b9150509250929050565b7f696e73756620706c6564676500000000000000000000000000000000000000005f82015250565b5f6158d7600c83614866565b91506158e2826158a3565b602082019050919050565b5f6020820190508181035f830152615904816158cb565b9050919050565b7f696e76206e6f64650000000000000000000000000000000000000000000000005f82015250565b5f61593f600883614866565b915061594a8261590b565b602082019050919050565b5f6020820190508181035f83015261596c81615933565b9050919050565b5f61597e8285614f83565b915061598a8284614f83565b91508190509392505050565b5f6020820190506159a95f830184614e52565b92915050565b7f6561726c790000000000000000000000000000000000000000000000000000005f82015250565b5f6159e3600583614866565b91506159ee826159af565b602082019050919050565b5f6020820190508181035f830152615a10816159d7565b9050919050565b7f6661696c656400000000000000000000000000000000000000000000000000005f82015250565b5f615a4b600683614866565b9150615a5682615a17565b602082019050919050565b5f6020820190508181035f830152615a7881615a3f565b9050919050565b5f604082019050615a925f8301856146b1565b615a9f60208301846142ce565b9392505050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601260045260245ffd5b5f615add826142c5565b9150615ae8836142c5565b925082615af857615af7615aa6565b5b828204905092915050565b5f81519050615b1181614a9a565b92915050565b5f60208284031215615b2c57615b2b6142ff565b5b5f615b3984828501615b03565b91505092915050565b7f696e206368616c000000000000000000000000000000000000000000000000005f82015250565b5f615b76600783614866565b9150615b8182615b42565b602082019050919050565b5f6020820190508181035f830152615ba381615b6a565b9050919050565b5f606082019050615bbd5f8301866146b1565b615bca60208301856146b1565b615bd760408301846142ce565b94935050505056fea26469706673582212205c1c193ca377b6d752867ab4d8e91dc47f2e16a04f773a8d82ee04a41dbb8c7f64736f6c63430008160033",
}

// EProofABI is the input ABI used to generate the binding from.
// Deprecated: Use EProofMetaData.ABI instead.
var EProofABI = EProofMetaData.ABI

// EProofBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use EProofMetaData.Bin instead.
var EProofBin = EProofMetaData.Bin

// DeployEProof deploys a new Ethereum contract, binding an instance of EProof to it.
func DeployEProof(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *EProof, error) {
	parsed, err := EProofMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(EProofBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &EProof{EProofCaller: EProofCaller{contract: contract}, EProofTransactor: EProofTransactor{contract: contract}, EProofFilterer: EProofFilterer{contract: contract}}, nil
}

// EProof is an auto generated Go binding around an Ethereum contract.
type EProof struct {
	EProofCaller     // Read-only binding to the contract
	EProofTransactor // Write-only binding to the contract
	EProofFilterer   // Log filterer for contract events
}

// EProofCaller is an auto generated read-only Go binding around an Ethereum contract.
type EProofCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EProofTransactor is an auto generated write-only Go binding around an Ethereum contract.
type EProofTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EProofFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type EProofFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EProofSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type EProofSession struct {
	Contract     *EProof           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// EProofCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type EProofCallerSession struct {
	Contract *EProofCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// EProofTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type EProofTransactorSession struct {
	Contract     *EProofTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// EProofRaw is an auto generated low-level Go binding around an Ethereum contract.
type EProofRaw struct {
	Contract *EProof // Generic contract binding to access the raw methods on
}

// EProofCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type EProofCallerRaw struct {
	Contract *EProofCaller // Generic read-only contract binding to access the raw methods on
}

// EProofTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type EProofTransactorRaw struct {
	Contract *EProofTransactor // Generic write-only contract binding to access the raw methods on
}

// NewEProof creates a new instance of EProof, bound to a specific deployed contract.
func NewEProof(address common.Address, backend bind.ContractBackend) (*EProof, error) {
	contract, err := bindEProof(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &EProof{EProofCaller: EProofCaller{contract: contract}, EProofTransactor: EProofTransactor{contract: contract}, EProofFilterer: EProofFilterer{contract: contract}}, nil
}

// NewEProofCaller creates a new read-only instance of EProof, bound to a specific deployed contract.
func NewEProofCaller(address common.Address, caller bind.ContractCaller) (*EProofCaller, error) {
	contract, err := bindEProof(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &EProofCaller{contract: contract}, nil
}

// NewEProofTransactor creates a new write-only instance of EProof, bound to a specific deployed contract.
func NewEProofTransactor(address common.Address, transactor bind.ContractTransactor) (*EProofTransactor, error) {
	contract, err := bindEProof(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &EProofTransactor{contract: contract}, nil
}

// NewEProofFilterer creates a new log filterer instance of EProof, bound to a specific deployed contract.
func NewEProofFilterer(address common.Address, filterer bind.ContractFilterer) (*EProofFilterer, error) {
	contract, err := bindEProof(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &EProofFilterer{contract: contract}, nil
}

// bindEProof binds a generic wrapper to an already deployed contract.
func bindEProof(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := EProofMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_EProof *EProofRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _EProof.Contract.EProofCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_EProof *EProofRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EProof.Contract.EProofTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_EProof *EProofRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _EProof.Contract.EProofTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_EProof *EProofCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _EProof.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_EProof *EProofTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EProof.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_EProof *EProofTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _EProof.Contract.contract.Transact(opts, method, params...)
}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_EProof *EProofCaller) UPGRADEINTERFACEVERSION(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _EProof.contract.Call(opts, &out, "UPGRADE_INTERFACE_VERSION")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_EProof *EProofSession) UPGRADEINTERFACEVERSION() (string, error) {
	return _EProof.Contract.UPGRADEINTERFACEVERSION(&_EProof.CallOpts)
}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_EProof *EProofCallerSession) UPGRADEINTERFACEVERSION() (string, error) {
	return _EProof.Contract.UPGRADEINTERFACEVERSION(&_EProof.CallOpts)
}

// VERSION is a free data retrieval call binding the contract method 0xffa1ad74.
//
// Solidity: function VERSION() view returns(string)
func (_EProof *EProofCaller) VERSION(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _EProof.contract.Call(opts, &out, "VERSION")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// VERSION is a free data retrieval call binding the contract method 0xffa1ad74.
//
// Solidity: function VERSION() view returns(string)
func (_EProof *EProofSession) VERSION() (string, error) {
	return _EProof.Contract.VERSION(&_EProof.CallOpts)
}

// VERSION is a free data retrieval call binding the contract method 0xffa1ad74.
//
// Solidity: function VERSION() view returns(string)
func (_EProof *EProofCallerSession) VERSION() (string, error) {
	return _EProof.Contract.VERSION(&_EProof.CallOpts)
}

// Base is a free data retrieval call binding the contract method 0x5001f3b5.
//
// Solidity: function base() view returns(address)
func (_EProof *EProofCaller) Base(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _EProof.contract.Call(opts, &out, "base")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Base is a free data retrieval call binding the contract method 0x5001f3b5.
//
// Solidity: function base() view returns(address)
func (_EProof *EProofSession) Base() (common.Address, error) {
	return _EProof.Contract.Base(&_EProof.CallOpts)
}

// Base is a free data retrieval call binding the contract method 0x5001f3b5.
//
// Solidity: function base() view returns(address)
func (_EProof *EProofCallerSession) Base() (common.Address, error) {
	return _EProof.Contract.Base(&_EProof.CallOpts)
}

// BasePenalty is a free data retrieval call binding the contract method 0x5712e98c.
//
// Solidity: function basePenalty() view returns(uint256)
func (_EProof *EProofCaller) BasePenalty(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _EProof.contract.Call(opts, &out, "basePenalty")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BasePenalty is a free data retrieval call binding the contract method 0x5712e98c.
//
// Solidity: function basePenalty() view returns(uint256)
func (_EProof *EProofSession) BasePenalty() (*big.Int, error) {
	return _EProof.Contract.BasePenalty(&_EProof.CallOpts)
}

// BasePenalty is a free data retrieval call binding the contract method 0x5712e98c.
//
// Solidity: function basePenalty() view returns(uint256)
func (_EProof *EProofCallerSession) BasePenalty() (*big.Int, error) {
	return _EProof.Contract.BasePenalty(&_EProof.CallOpts)
}

// Epoch is a free data retrieval call binding the contract method 0x900cf0cf.
//
// Solidity: function epoch() view returns(address)
func (_EProof *EProofCaller) Epoch(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _EProof.contract.Call(opts, &out, "epoch")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Epoch is a free data retrieval call binding the contract method 0x900cf0cf.
//
// Solidity: function epoch() view returns(address)
func (_EProof *EProofSession) Epoch() (common.Address, error) {
	return _EProof.Contract.Epoch(&_EProof.CallOpts)
}

// Epoch is a free data retrieval call binding the contract method 0x900cf0cf.
//
// Solidity: function epoch() view returns(address)
func (_EProof *EProofCallerSession) Epoch() (common.Address, error) {
	return _EProof.Contract.Epoch(&_EProof.CallOpts)
}

// Everify is a free data retrieval call binding the contract method 0x62b26a72.
//
// Solidity: function everify() view returns(address)
func (_EProof *EProofCaller) Everify(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _EProof.contract.Call(opts, &out, "everify")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Everify is a free data retrieval call binding the contract method 0x62b26a72.
//
// Solidity: function everify() view returns(address)
func (_EProof *EProofSession) Everify() (common.Address, error) {
	return _EProof.Contract.Everify(&_EProof.CallOpts)
}

// Everify is a free data retrieval call binding the contract method 0x62b26a72.
//
// Solidity: function everify() view returns(address)
func (_EProof *EProofCallerSession) Everify() (common.Address, error) {
	return _EProof.Contract.Everify(&_EProof.CallOpts)
}

// GetEProof is a free data retrieval call binding the contract method 0x266d12cb.
//
// Solidity: function getEProof(address _a, uint64 _e) view returns((bool,bool,address,uint256,uint256))
func (_EProof *EProofCaller) GetEProof(opts *bind.CallOpts, _a common.Address, _e uint64) (IEProofProofInfo, error) {
	var out []interface{}
	err := _EProof.contract.Call(opts, &out, "getEProof", _a, _e)

	if err != nil {
		return *new(IEProofProofInfo), err
	}

	out0 := *abi.ConvertType(out[0], new(IEProofProofInfo)).(*IEProofProofInfo)

	return out0, err

}

// GetEProof is a free data retrieval call binding the contract method 0x266d12cb.
//
// Solidity: function getEProof(address _a, uint64 _e) view returns((bool,bool,address,uint256,uint256))
func (_EProof *EProofSession) GetEProof(_a common.Address, _e uint64) (IEProofProofInfo, error) {
	return _EProof.Contract.GetEProof(&_EProof.CallOpts, _a, _e)
}

// GetEProof is a free data retrieval call binding the contract method 0x266d12cb.
//
// Solidity: function getEProof(address _a, uint64 _e) view returns((bool,bool,address,uint256,uint256))
func (_EProof *EProofCallerSession) GetEProof(_a common.Address, _e uint64) (IEProofProofInfo, error) {
	return _EProof.Contract.GetEProof(&_EProof.CallOpts, _a, _e)
}

// GetStat is a free data retrieval call binding the contract method 0xdf6387fa.
//
// Solidity: function getStat(address _a, uint64 _e) view returns(bool)
func (_EProof *EProofCaller) GetStat(opts *bind.CallOpts, _a common.Address, _e uint64) (bool, error) {
	var out []interface{}
	err := _EProof.contract.Call(opts, &out, "getStat", _a, _e)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// GetStat is a free data retrieval call binding the contract method 0xdf6387fa.
//
// Solidity: function getStat(address _a, uint64 _e) view returns(bool)
func (_EProof *EProofSession) GetStat(_a common.Address, _e uint64) (bool, error) {
	return _EProof.Contract.GetStat(&_EProof.CallOpts, _a, _e)
}

// GetStat is a free data retrieval call binding the contract method 0xdf6387fa.
//
// Solidity: function getStat(address _a, uint64 _e) view returns(bool)
func (_EProof *EProofCallerSession) GetStat(_a common.Address, _e uint64) (bool, error) {
	return _EProof.Contract.GetStat(&_EProof.CallOpts, _a, _e)
}

// MinProveTime is a free data retrieval call binding the contract method 0x0bd26cb5.
//
// Solidity: function minProveTime() view returns(uint256)
func (_EProof *EProofCaller) MinProveTime(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _EProof.contract.Call(opts, &out, "minProveTime")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MinProveTime is a free data retrieval call binding the contract method 0x0bd26cb5.
//
// Solidity: function minProveTime() view returns(uint256)
func (_EProof *EProofSession) MinProveTime() (*big.Int, error) {
	return _EProof.Contract.MinProveTime(&_EProof.CallOpts)
}

// MinProveTime is a free data retrieval call binding the contract method 0x0bd26cb5.
//
// Solidity: function minProveTime() view returns(uint256)
func (_EProof *EProofCallerSession) MinProveTime() (*big.Int, error) {
	return _EProof.Contract.MinProveTime(&_EProof.CallOpts)
}

// Node is a free data retrieval call binding the contract method 0xd70754ec.
//
// Solidity: function node() view returns(address)
func (_EProof *EProofCaller) Node(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _EProof.contract.Call(opts, &out, "node")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Node is a free data retrieval call binding the contract method 0xd70754ec.
//
// Solidity: function node() view returns(address)
func (_EProof *EProofSession) Node() (common.Address, error) {
	return _EProof.Contract.Node(&_EProof.CallOpts)
}

// Node is a free data retrieval call binding the contract method 0xd70754ec.
//
// Solidity: function node() view returns(address)
func (_EProof *EProofCallerSession) Node() (common.Address, error) {
	return _EProof.Contract.Node(&_EProof.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_EProof *EProofCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _EProof.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_EProof *EProofSession) Owner() (common.Address, error) {
	return _EProof.Contract.Owner(&_EProof.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_EProof *EProofCallerSession) Owner() (common.Address, error) {
	return _EProof.Contract.Owner(&_EProof.CallOpts)
}

// Piece is a free data retrieval call binding the contract method 0x95ad21dc.
//
// Solidity: function piece() view returns(address)
func (_EProof *EProofCaller) Piece(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _EProof.contract.Call(opts, &out, "piece")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Piece is a free data retrieval call binding the contract method 0x95ad21dc.
//
// Solidity: function piece() view returns(address)
func (_EProof *EProofSession) Piece() (common.Address, error) {
	return _EProof.Contract.Piece(&_EProof.CallOpts)
}

// Piece is a free data retrieval call binding the contract method 0x95ad21dc.
//
// Solidity: function piece() view returns(address)
func (_EProof *EProofCallerSession) Piece() (common.Address, error) {
	return _EProof.Contract.Piece(&_EProof.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_EProof *EProofCaller) ProxiableUUID(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _EProof.contract.Call(opts, &out, "proxiableUUID")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_EProof *EProofSession) ProxiableUUID() ([32]byte, error) {
	return _EProof.Contract.ProxiableUUID(&_EProof.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_EProof *EProofCallerSession) ProxiableUUID() ([32]byte, error) {
	return _EProof.Contract.ProxiableUUID(&_EProof.CallOpts)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_EProof *EProofCaller) Token(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _EProof.contract.Call(opts, &out, "token")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_EProof *EProofSession) Token() (common.Address, error) {
	return _EProof.Contract.Token(&_EProof.CallOpts)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_EProof *EProofCallerSession) Token() (common.Address, error) {
	return _EProof.Contract.Token(&_EProof.CallOpts)
}

// ChalCom is a paid mutator transaction binding the contract method 0xb54753b8.
//
// Solidity: function chalCom(address _a, uint64 _ep, uint8 _qIndex) returns()
func (_EProof *EProofTransactor) ChalCom(opts *bind.TransactOpts, _a common.Address, _ep uint64, _qIndex uint8) (*types.Transaction, error) {
	return _EProof.contract.Transact(opts, "chalCom", _a, _ep, _qIndex)
}

// ChalCom is a paid mutator transaction binding the contract method 0xb54753b8.
//
// Solidity: function chalCom(address _a, uint64 _ep, uint8 _qIndex) returns()
func (_EProof *EProofSession) ChalCom(_a common.Address, _ep uint64, _qIndex uint8) (*types.Transaction, error) {
	return _EProof.Contract.ChalCom(&_EProof.TransactOpts, _a, _ep, _qIndex)
}

// ChalCom is a paid mutator transaction binding the contract method 0xb54753b8.
//
// Solidity: function chalCom(address _a, uint64 _ep, uint8 _qIndex) returns()
func (_EProof *EProofTransactorSession) ChalCom(_a common.Address, _ep uint64, _qIndex uint8) (*types.Transaction, error) {
	return _EProof.Contract.ChalCom(&_EProof.TransactOpts, _a, _ep, _qIndex)
}

// ChalKZG is a paid mutator transaction binding the contract method 0x7ecaca98.
//
// Solidity: function chalKZG(address _a, uint64 _ep) returns()
func (_EProof *EProofTransactor) ChalKZG(opts *bind.TransactOpts, _a common.Address, _ep uint64) (*types.Transaction, error) {
	return _EProof.contract.Transact(opts, "chalKZG", _a, _ep)
}

// ChalKZG is a paid mutator transaction binding the contract method 0x7ecaca98.
//
// Solidity: function chalKZG(address _a, uint64 _ep) returns()
func (_EProof *EProofSession) ChalKZG(_a common.Address, _ep uint64) (*types.Transaction, error) {
	return _EProof.Contract.ChalKZG(&_EProof.TransactOpts, _a, _ep)
}

// ChalKZG is a paid mutator transaction binding the contract method 0x7ecaca98.
//
// Solidity: function chalKZG(address _a, uint64 _ep) returns()
func (_EProof *EProofTransactorSession) ChalKZG(_a common.Address, _ep uint64) (*types.Transaction, error) {
	return _EProof.Contract.ChalKZG(&_EProof.TransactOpts, _a, _ep)
}

// ChalOne is a paid mutator transaction binding the contract method 0xbd42f4c2.
//
// Solidity: function chalOne(address _a, uint64 _ep, uint8 _qIndex) returns()
func (_EProof *EProofTransactor) ChalOne(opts *bind.TransactOpts, _a common.Address, _ep uint64, _qIndex uint8) (*types.Transaction, error) {
	return _EProof.contract.Transact(opts, "chalOne", _a, _ep, _qIndex)
}

// ChalOne is a paid mutator transaction binding the contract method 0xbd42f4c2.
//
// Solidity: function chalOne(address _a, uint64 _ep, uint8 _qIndex) returns()
func (_EProof *EProofSession) ChalOne(_a common.Address, _ep uint64, _qIndex uint8) (*types.Transaction, error) {
	return _EProof.Contract.ChalOne(&_EProof.TransactOpts, _a, _ep, _qIndex)
}

// ChalOne is a paid mutator transaction binding the contract method 0xbd42f4c2.
//
// Solidity: function chalOne(address _a, uint64 _ep, uint8 _qIndex) returns()
func (_EProof *EProofTransactorSession) ChalOne(_a common.Address, _ep uint64, _qIndex uint8) (*types.Transaction, error) {
	return _EProof.Contract.ChalOne(&_EProof.TransactOpts, _a, _ep, _qIndex)
}

// Challenge is a paid mutator transaction binding the contract method 0x5a9b427d.
//
// Solidity: function challenge(address _a, uint64 _ep, bytes _sum) returns()
func (_EProof *EProofTransactor) Challenge(opts *bind.TransactOpts, _a common.Address, _ep uint64, _sum []byte) (*types.Transaction, error) {
	return _EProof.contract.Transact(opts, "challenge", _a, _ep, _sum)
}

// Challenge is a paid mutator transaction binding the contract method 0x5a9b427d.
//
// Solidity: function challenge(address _a, uint64 _ep, bytes _sum) returns()
func (_EProof *EProofSession) Challenge(_a common.Address, _ep uint64, _sum []byte) (*types.Transaction, error) {
	return _EProof.Contract.Challenge(&_EProof.TransactOpts, _a, _ep, _sum)
}

// Challenge is a paid mutator transaction binding the contract method 0x5a9b427d.
//
// Solidity: function challenge(address _a, uint64 _ep, bytes _sum) returns()
func (_EProof *EProofTransactorSession) Challenge(_a common.Address, _ep uint64, _sum []byte) (*types.Transaction, error) {
	return _EProof.Contract.Challenge(&_EProof.TransactOpts, _a, _ep, _sum)
}

// Check is a paid mutator transaction binding the contract method 0x5eb24b67.
//
// Solidity: function check(address _a, uint64 _ep) returns()
func (_EProof *EProofTransactor) Check(opts *bind.TransactOpts, _a common.Address, _ep uint64) (*types.Transaction, error) {
	return _EProof.contract.Transact(opts, "check", _a, _ep)
}

// Check is a paid mutator transaction binding the contract method 0x5eb24b67.
//
// Solidity: function check(address _a, uint64 _ep) returns()
func (_EProof *EProofSession) Check(_a common.Address, _ep uint64) (*types.Transaction, error) {
	return _EProof.Contract.Check(&_EProof.TransactOpts, _a, _ep)
}

// Check is a paid mutator transaction binding the contract method 0x5eb24b67.
//
// Solidity: function check(address _a, uint64 _ep) returns()
func (_EProof *EProofTransactorSession) Check(_a common.Address, _ep uint64) (*types.Transaction, error) {
	return _EProof.Contract.Check(&_EProof.TransactOpts, _a, _ep)
}

// Initialize is a paid mutator transaction binding the contract method 0xb7c1e98e.
//
// Solidity: function initialize((address,address,address,address,address,address,uint256) params, address initialOwner) returns()
func (_EProof *EProofTransactor) Initialize(opts *bind.TransactOpts, params IEProofInitParams, initialOwner common.Address) (*types.Transaction, error) {
	return _EProof.contract.Transact(opts, "initialize", params, initialOwner)
}

// Initialize is a paid mutator transaction binding the contract method 0xb7c1e98e.
//
// Solidity: function initialize((address,address,address,address,address,address,uint256) params, address initialOwner) returns()
func (_EProof *EProofSession) Initialize(params IEProofInitParams, initialOwner common.Address) (*types.Transaction, error) {
	return _EProof.Contract.Initialize(&_EProof.TransactOpts, params, initialOwner)
}

// Initialize is a paid mutator transaction binding the contract method 0xb7c1e98e.
//
// Solidity: function initialize((address,address,address,address,address,address,uint256) params, address initialOwner) returns()
func (_EProof *EProofTransactorSession) Initialize(params IEProofInitParams, initialOwner common.Address) (*types.Transaction, error) {
	return _EProof.Contract.Initialize(&_EProof.TransactOpts, params, initialOwner)
}

// ProveCom is a paid mutator transaction binding the contract method 0xe58e8e05.
//
// Solidity: function proveCom(uint64 _ep, bytes[] _com, bytes _proof) returns()
func (_EProof *EProofTransactor) ProveCom(opts *bind.TransactOpts, _ep uint64, _com [][]byte, _proof []byte) (*types.Transaction, error) {
	return _EProof.contract.Transact(opts, "proveCom", _ep, _com, _proof)
}

// ProveCom is a paid mutator transaction binding the contract method 0xe58e8e05.
//
// Solidity: function proveCom(uint64 _ep, bytes[] _com, bytes _proof) returns()
func (_EProof *EProofSession) ProveCom(_ep uint64, _com [][]byte, _proof []byte) (*types.Transaction, error) {
	return _EProof.Contract.ProveCom(&_EProof.TransactOpts, _ep, _com, _proof)
}

// ProveCom is a paid mutator transaction binding the contract method 0xe58e8e05.
//
// Solidity: function proveCom(uint64 _ep, bytes[] _com, bytes _proof) returns()
func (_EProof *EProofTransactorSession) ProveCom(_ep uint64, _com [][]byte, _proof []byte) (*types.Transaction, error) {
	return _EProof.Contract.ProveCom(&_EProof.TransactOpts, _ep, _com, _proof)
}

// ProveKZG is a paid mutator transaction binding the contract method 0xca489a87.
//
// Solidity: function proveKZG(uint64 _ep, bytes32 _wroot, bytes _proof) returns()
func (_EProof *EProofTransactor) ProveKZG(opts *bind.TransactOpts, _ep uint64, _wroot [32]byte, _proof []byte) (*types.Transaction, error) {
	return _EProof.contract.Transact(opts, "proveKZG", _ep, _wroot, _proof)
}

// ProveKZG is a paid mutator transaction binding the contract method 0xca489a87.
//
// Solidity: function proveKZG(uint64 _ep, bytes32 _wroot, bytes _proof) returns()
func (_EProof *EProofSession) ProveKZG(_ep uint64, _wroot [32]byte, _proof []byte) (*types.Transaction, error) {
	return _EProof.Contract.ProveKZG(&_EProof.TransactOpts, _ep, _wroot, _proof)
}

// ProveKZG is a paid mutator transaction binding the contract method 0xca489a87.
//
// Solidity: function proveKZG(uint64 _ep, bytes32 _wroot, bytes _proof) returns()
func (_EProof *EProofTransactorSession) ProveKZG(_ep uint64, _wroot [32]byte, _proof []byte) (*types.Transaction, error) {
	return _EProof.Contract.ProveKZG(&_EProof.TransactOpts, _ep, _wroot, _proof)
}

// ProveOne is a paid mutator transaction binding the contract method 0x4ca2a890.
//
// Solidity: function proveOne(uint64 _ep, bytes _com, bytes _proof) returns()
func (_EProof *EProofTransactor) ProveOne(opts *bind.TransactOpts, _ep uint64, _com []byte, _proof []byte) (*types.Transaction, error) {
	return _EProof.contract.Transact(opts, "proveOne", _ep, _com, _proof)
}

// ProveOne is a paid mutator transaction binding the contract method 0x4ca2a890.
//
// Solidity: function proveOne(uint64 _ep, bytes _com, bytes _proof) returns()
func (_EProof *EProofSession) ProveOne(_ep uint64, _com []byte, _proof []byte) (*types.Transaction, error) {
	return _EProof.Contract.ProveOne(&_EProof.TransactOpts, _ep, _com, _proof)
}

// ProveOne is a paid mutator transaction binding the contract method 0x4ca2a890.
//
// Solidity: function proveOne(uint64 _ep, bytes _com, bytes _proof) returns()
func (_EProof *EProofTransactorSession) ProveOne(_ep uint64, _com []byte, _proof []byte) (*types.Transaction, error) {
	return _EProof.Contract.ProveOne(&_EProof.TransactOpts, _ep, _com, _proof)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_EProof *EProofTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EProof.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_EProof *EProofSession) RenounceOwnership() (*types.Transaction, error) {
	return _EProof.Contract.RenounceOwnership(&_EProof.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_EProof *EProofTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _EProof.Contract.RenounceOwnership(&_EProof.TransactOpts)
}

// SetMinProveTime is a paid mutator transaction binding the contract method 0xdd489315.
//
// Solidity: function setMinProveTime(uint256 _t) returns()
func (_EProof *EProofTransactor) SetMinProveTime(opts *bind.TransactOpts, _t *big.Int) (*types.Transaction, error) {
	return _EProof.contract.Transact(opts, "setMinProveTime", _t)
}

// SetMinProveTime is a paid mutator transaction binding the contract method 0xdd489315.
//
// Solidity: function setMinProveTime(uint256 _t) returns()
func (_EProof *EProofSession) SetMinProveTime(_t *big.Int) (*types.Transaction, error) {
	return _EProof.Contract.SetMinProveTime(&_EProof.TransactOpts, _t)
}

// SetMinProveTime is a paid mutator transaction binding the contract method 0xdd489315.
//
// Solidity: function setMinProveTime(uint256 _t) returns()
func (_EProof *EProofTransactorSession) SetMinProveTime(_t *big.Int) (*types.Transaction, error) {
	return _EProof.Contract.SetMinProveTime(&_EProof.TransactOpts, _t)
}

// Submit is a paid mutator transaction binding the contract method 0xe70b775e.
//
// Solidity: function submit(uint64 _ep, bytes _com, bytes _proof) returns()
func (_EProof *EProofTransactor) Submit(opts *bind.TransactOpts, _ep uint64, _com []byte, _proof []byte) (*types.Transaction, error) {
	return _EProof.contract.Transact(opts, "submit", _ep, _com, _proof)
}

// Submit is a paid mutator transaction binding the contract method 0xe70b775e.
//
// Solidity: function submit(uint64 _ep, bytes _com, bytes _proof) returns()
func (_EProof *EProofSession) Submit(_ep uint64, _com []byte, _proof []byte) (*types.Transaction, error) {
	return _EProof.Contract.Submit(&_EProof.TransactOpts, _ep, _com, _proof)
}

// Submit is a paid mutator transaction binding the contract method 0xe70b775e.
//
// Solidity: function submit(uint64 _ep, bytes _com, bytes _proof) returns()
func (_EProof *EProofTransactorSession) Submit(_ep uint64, _com []byte, _proof []byte) (*types.Transaction, error) {
	return _EProof.Contract.Submit(&_EProof.TransactOpts, _ep, _com, _proof)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_EProof *EProofTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _EProof.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_EProof *EProofSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _EProof.Contract.TransferOwnership(&_EProof.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_EProof *EProofTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _EProof.Contract.TransferOwnership(&_EProof.TransactOpts, newOwner)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_EProof *EProofTransactor) UpgradeToAndCall(opts *bind.TransactOpts, newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _EProof.contract.Transact(opts, "upgradeToAndCall", newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_EProof *EProofSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _EProof.Contract.UpgradeToAndCall(&_EProof.TransactOpts, newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_EProof *EProofTransactorSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _EProof.Contract.UpgradeToAndCall(&_EProof.TransactOpts, newImplementation, data)
}

// EProofChallengeIterator is returned from FilterChallenge and is used to iterate over the raw logs and unpacked data for Challenge events raised by the EProof contract.
type EProofChallengeIterator struct {
	Event *EProofChallenge // Event containing the contract specifics and raw log

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
func (it *EProofChallengeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EProofChallenge)
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
		it.Event = new(EProofChallenge)
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
func (it *EProofChallengeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EProofChallengeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EProofChallenge represents a Challenge event raised by the EProof contract.
type EProofChallenge struct {
	S     common.Address
	E     uint64
	Round uint8
	Qi    uint8
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterChallenge is a free log retrieval operation binding the contract event 0x1cf322418045f261f2817d06ac5cf24890e870ce5863c6434a8893f03d6eae4c.
//
// Solidity: event Challenge(address indexed _s, uint64 _e, uint8 _round, uint8 _qi)
func (_EProof *EProofFilterer) FilterChallenge(opts *bind.FilterOpts, _s []common.Address) (*EProofChallengeIterator, error) {

	var _sRule []interface{}
	for _, _sItem := range _s {
		_sRule = append(_sRule, _sItem)
	}

	logs, sub, err := _EProof.contract.FilterLogs(opts, "Challenge", _sRule)
	if err != nil {
		return nil, err
	}
	return &EProofChallengeIterator{contract: _EProof.contract, event: "Challenge", logs: logs, sub: sub}, nil
}

// WatchChallenge is a free log subscription operation binding the contract event 0x1cf322418045f261f2817d06ac5cf24890e870ce5863c6434a8893f03d6eae4c.
//
// Solidity: event Challenge(address indexed _s, uint64 _e, uint8 _round, uint8 _qi)
func (_EProof *EProofFilterer) WatchChallenge(opts *bind.WatchOpts, sink chan<- *EProofChallenge, _s []common.Address) (event.Subscription, error) {

	var _sRule []interface{}
	for _, _sItem := range _s {
		_sRule = append(_sRule, _sItem)
	}

	logs, sub, err := _EProof.contract.WatchLogs(opts, "Challenge", _sRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EProofChallenge)
				if err := _EProof.contract.UnpackLog(event, "Challenge", log); err != nil {
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

// ParseChallenge is a log parse operation binding the contract event 0x1cf322418045f261f2817d06ac5cf24890e870ce5863c6434a8893f03d6eae4c.
//
// Solidity: event Challenge(address indexed _s, uint64 _e, uint8 _round, uint8 _qi)
func (_EProof *EProofFilterer) ParseChallenge(log types.Log) (*EProofChallenge, error) {
	event := new(EProofChallenge)
	if err := _EProof.contract.UnpackLog(event, "Challenge", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EProofFakeIterator is returned from FilterFake and is used to iterate over the raw logs and unpacked data for Fake events raised by the EProof contract.
type EProofFakeIterator struct {
	Event *EProofFake // Event containing the contract specifics and raw log

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
func (it *EProofFakeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EProofFake)
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
		it.Event = new(EProofFake)
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
func (it *EProofFakeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EProofFakeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EProofFake represents a Fake event raised by the EProof contract.
type EProofFake struct {
	S   common.Address
	E   uint64
	Raw types.Log // Blockchain specific contextual infos
}

// FilterFake is a free log retrieval operation binding the contract event 0x0fcd3bb0ce91639d74185d625544f2b5714ecd14a6a6d9ccb65b2526e9ac12f5.
//
// Solidity: event Fake(address indexed _s, uint64 _e)
func (_EProof *EProofFilterer) FilterFake(opts *bind.FilterOpts, _s []common.Address) (*EProofFakeIterator, error) {

	var _sRule []interface{}
	for _, _sItem := range _s {
		_sRule = append(_sRule, _sItem)
	}

	logs, sub, err := _EProof.contract.FilterLogs(opts, "Fake", _sRule)
	if err != nil {
		return nil, err
	}
	return &EProofFakeIterator{contract: _EProof.contract, event: "Fake", logs: logs, sub: sub}, nil
}

// WatchFake is a free log subscription operation binding the contract event 0x0fcd3bb0ce91639d74185d625544f2b5714ecd14a6a6d9ccb65b2526e9ac12f5.
//
// Solidity: event Fake(address indexed _s, uint64 _e)
func (_EProof *EProofFilterer) WatchFake(opts *bind.WatchOpts, sink chan<- *EProofFake, _s []common.Address) (event.Subscription, error) {

	var _sRule []interface{}
	for _, _sItem := range _s {
		_sRule = append(_sRule, _sItem)
	}

	logs, sub, err := _EProof.contract.WatchLogs(opts, "Fake", _sRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EProofFake)
				if err := _EProof.contract.UnpackLog(event, "Fake", log); err != nil {
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
func (_EProof *EProofFilterer) ParseFake(log types.Log) (*EProofFake, error) {
	event := new(EProofFake)
	if err := _EProof.contract.UnpackLog(event, "Fake", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EProofInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the EProof contract.
type EProofInitializedIterator struct {
	Event *EProofInitialized // Event containing the contract specifics and raw log

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
func (it *EProofInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EProofInitialized)
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
		it.Event = new(EProofInitialized)
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
func (it *EProofInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EProofInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EProofInitialized represents a Initialized event raised by the EProof contract.
type EProofInitialized struct {
	Version uint64
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_EProof *EProofFilterer) FilterInitialized(opts *bind.FilterOpts) (*EProofInitializedIterator, error) {

	logs, sub, err := _EProof.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &EProofInitializedIterator{contract: _EProof.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_EProof *EProofFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *EProofInitialized) (event.Subscription, error) {

	logs, sub, err := _EProof.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EProofInitialized)
				if err := _EProof.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_EProof *EProofFilterer) ParseInitialized(log types.Log) (*EProofInitialized, error) {
	event := new(EProofInitialized)
	if err := _EProof.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EProofOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the EProof contract.
type EProofOwnershipTransferredIterator struct {
	Event *EProofOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *EProofOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EProofOwnershipTransferred)
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
		it.Event = new(EProofOwnershipTransferred)
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
func (it *EProofOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EProofOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EProofOwnershipTransferred represents a OwnershipTransferred event raised by the EProof contract.
type EProofOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_EProof *EProofFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*EProofOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _EProof.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &EProofOwnershipTransferredIterator{contract: _EProof.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_EProof *EProofFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *EProofOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _EProof.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EProofOwnershipTransferred)
				if err := _EProof.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_EProof *EProofFilterer) ParseOwnershipTransferred(log types.Log) (*EProofOwnershipTransferred, error) {
	event := new(EProofOwnershipTransferred)
	if err := _EProof.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EProofProveIterator is returned from FilterProve and is used to iterate over the raw logs and unpacked data for Prove events raised by the EProof contract.
type EProofProveIterator struct {
	Event *EProofProve // Event containing the contract specifics and raw log

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
func (it *EProofProveIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EProofProve)
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
		it.Event = new(EProofProve)
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
func (it *EProofProveIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EProofProveIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EProofProve represents a Prove event raised by the EProof contract.
type EProofProve struct {
	S     common.Address
	E     uint64
	Round uint8
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterProve is a free log retrieval operation binding the contract event 0x074353199408473b546546e1626a6b8ac62ff72784909ac8d0708a73125484af.
//
// Solidity: event Prove(address indexed _s, uint64 _e, uint8 _round)
func (_EProof *EProofFilterer) FilterProve(opts *bind.FilterOpts, _s []common.Address) (*EProofProveIterator, error) {

	var _sRule []interface{}
	for _, _sItem := range _s {
		_sRule = append(_sRule, _sItem)
	}

	logs, sub, err := _EProof.contract.FilterLogs(opts, "Prove", _sRule)
	if err != nil {
		return nil, err
	}
	return &EProofProveIterator{contract: _EProof.contract, event: "Prove", logs: logs, sub: sub}, nil
}

// WatchProve is a free log subscription operation binding the contract event 0x074353199408473b546546e1626a6b8ac62ff72784909ac8d0708a73125484af.
//
// Solidity: event Prove(address indexed _s, uint64 _e, uint8 _round)
func (_EProof *EProofFilterer) WatchProve(opts *bind.WatchOpts, sink chan<- *EProofProve, _s []common.Address) (event.Subscription, error) {

	var _sRule []interface{}
	for _, _sItem := range _s {
		_sRule = append(_sRule, _sItem)
	}

	logs, sub, err := _EProof.contract.WatchLogs(opts, "Prove", _sRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EProofProve)
				if err := _EProof.contract.UnpackLog(event, "Prove", log); err != nil {
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
func (_EProof *EProofFilterer) ParseProve(log types.Log) (*EProofProve, error) {
	event := new(EProofProve)
	if err := _EProof.contract.UnpackLog(event, "Prove", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EProofSubmitIterator is returned from FilterSubmit and is used to iterate over the raw logs and unpacked data for Submit events raised by the EProof contract.
type EProofSubmitIterator struct {
	Event *EProofSubmit // Event containing the contract specifics and raw log

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
func (it *EProofSubmitIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EProofSubmit)
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
		it.Event = new(EProofSubmit)
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
func (it *EProofSubmitIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EProofSubmitIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EProofSubmit represents a Submit event raised by the EProof contract.
type EProofSubmit struct {
	S   common.Address
	E   uint64
	Raw types.Log // Blockchain specific contextual infos
}

// FilterSubmit is a free log retrieval operation binding the contract event 0x98e23f0b72b943d75df730f1b4c0fbb9baffa041d18e1819f2d851a5c2c1c444.
//
// Solidity: event Submit(address indexed _s, uint64 _e)
func (_EProof *EProofFilterer) FilterSubmit(opts *bind.FilterOpts, _s []common.Address) (*EProofSubmitIterator, error) {

	var _sRule []interface{}
	for _, _sItem := range _s {
		_sRule = append(_sRule, _sItem)
	}

	logs, sub, err := _EProof.contract.FilterLogs(opts, "Submit", _sRule)
	if err != nil {
		return nil, err
	}
	return &EProofSubmitIterator{contract: _EProof.contract, event: "Submit", logs: logs, sub: sub}, nil
}

// WatchSubmit is a free log subscription operation binding the contract event 0x98e23f0b72b943d75df730f1b4c0fbb9baffa041d18e1819f2d851a5c2c1c444.
//
// Solidity: event Submit(address indexed _s, uint64 _e)
func (_EProof *EProofFilterer) WatchSubmit(opts *bind.WatchOpts, sink chan<- *EProofSubmit, _s []common.Address) (event.Subscription, error) {

	var _sRule []interface{}
	for _, _sItem := range _s {
		_sRule = append(_sRule, _sItem)
	}

	logs, sub, err := _EProof.contract.WatchLogs(opts, "Submit", _sRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EProofSubmit)
				if err := _EProof.contract.UnpackLog(event, "Submit", log); err != nil {
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
func (_EProof *EProofFilterer) ParseSubmit(log types.Log) (*EProofSubmit, error) {
	event := new(EProofSubmit)
	if err := _EProof.contract.UnpackLog(event, "Submit", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EProofUpgradedIterator is returned from FilterUpgraded and is used to iterate over the raw logs and unpacked data for Upgraded events raised by the EProof contract.
type EProofUpgradedIterator struct {
	Event *EProofUpgraded // Event containing the contract specifics and raw log

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
func (it *EProofUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EProofUpgraded)
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
		it.Event = new(EProofUpgraded)
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
func (it *EProofUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EProofUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EProofUpgraded represents a Upgraded event raised by the EProof contract.
type EProofUpgraded struct {
	Implementation common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterUpgraded is a free log retrieval operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_EProof *EProofFilterer) FilterUpgraded(opts *bind.FilterOpts, implementation []common.Address) (*EProofUpgradedIterator, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _EProof.contract.FilterLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return &EProofUpgradedIterator{contract: _EProof.contract, event: "Upgraded", logs: logs, sub: sub}, nil
}

// WatchUpgraded is a free log subscription operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_EProof *EProofFilterer) WatchUpgraded(opts *bind.WatchOpts, sink chan<- *EProofUpgraded, implementation []common.Address) (event.Subscription, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _EProof.contract.WatchLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EProofUpgraded)
				if err := _EProof.contract.UnpackLog(event, "Upgraded", log); err != nil {
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
func (_EProof *EProofFilterer) ParseUpgraded(log types.Log) (*EProofUpgraded, error) {
	event := new(EProofUpgraded)
	if err := _EProof.contract.UnpackLog(event, "Upgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
