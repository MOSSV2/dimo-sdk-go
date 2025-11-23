// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package piece

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

// IPieceInitParams is an auto generated low-level Go binding around an user-defined struct.
type IPieceInitParams struct {
	Delay       uint64
	MinStore    uint64
	MaxStore    uint64
	MaxSize     uint64
	MinPrice    *big.Int
	StreamPrice *big.Int
}

// IPiecePieceInfo is an auto generated low-level Go binding around an user-defined struct.
type IPiecePieceInfo struct {
	Rsn      uint8
	Rsk      uint8
	Size     uint64
	Start    uint64
	Expire   uint64
	Owner    common.Address
	Streamer common.Address
	Price    *big.Int
	Balance  *big.Int
}

// IPieceReplicaInfo is an auto generated low-level Go binding around an user-defined struct.
type IPieceReplicaInfo struct {
	Piece    uint64
	StoredOn common.Address
	Root     [32]byte
}

// IPieceStoreInfo is an auto generated low-level Go binding around an user-defined struct.
type IPieceStoreInfo struct {
	Epoch   uint64
	Count   uint64
	Salary  *big.Int
	Revenue *big.Int
}

// IPieceStoreStat is an auto generated low-level Go binding around an user-defined struct.
type IPieceStoreStat struct {
	Count  uint64
	Salary *big.Int
}

// PieceMetaData contains all meta data concerning the Piece contract.
var PieceMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"UPGRADE_INTERFACE_VERSION\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"VERSION\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"addPiece\",\"inputs\":[{\"name\":\"_pn\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"_price\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_size\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"_expire\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"rsn\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"rsk\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"_s\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"addReplica\",\"inputs\":[{\"name\":\"_rn\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"_pi\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"_pri\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"_proof\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"checkSReplica\",\"inputs\":[{\"name\":\"_a\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_ri\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"_com\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"checkStore\",\"inputs\":[{\"name\":\"_a\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"delay\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"epoch\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIEpoch\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getPIndex\",\"inputs\":[{\"name\":\"_pn\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getPRI\",\"inputs\":[{\"name\":\"_pi\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"_pri\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getPRInfo\",\"inputs\":[{\"name\":\"_pi\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"_ri\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getPiece\",\"inputs\":[{\"name\":\"_pi\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"_pinfo\",\"type\":\"tuple\",\"internalType\":\"structIPiece.PieceInfo\",\"components\":[{\"name\":\"rsn\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"rsk\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"size\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"start\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"expire\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"streamer\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"price\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"balance\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRIndex\",\"inputs\":[{\"name\":\"_rn\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRS\",\"inputs\":[{\"name\":\"_pi\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getReplica\",\"inputs\":[{\"name\":\"_r\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"_ri\",\"type\":\"tuple\",\"internalType\":\"structIPiece.ReplicaInfo\",\"components\":[{\"name\":\"piece\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"storedOn\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"root\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRevenue\",\"inputs\":[{\"name\":\"_a\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getSRAt\",\"inputs\":[{\"name\":\"_a\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_ri\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getSStat\",\"inputs\":[{\"name\":\"_a\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_e\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"_si\",\"type\":\"tuple\",\"internalType\":\"structIPiece.StoreStat\",\"components\":[{\"name\":\"count\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"salary\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getStore\",\"inputs\":[{\"name\":\"_a\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"_si\",\"type\":\"tuple\",\"internalType\":\"structIPiece.StoreInfo\",\"components\":[{\"name\":\"epoch\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"count\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"salary\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"revenue\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getStoreCount\",\"inputs\":[{\"name\":\"_a\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_e\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getStoreSalary\",\"inputs\":[{\"name\":\"_a\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_e\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"initialize\",\"inputs\":[{\"name\":\"_token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_epoch\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_node\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"params\",\"type\":\"tuple\",\"internalType\":\"structIPiece.InitParams\",\"components\":[{\"name\":\"delay\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"minStore\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"maxStore\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"maxSize\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"minPrice\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"streamPrice\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"initialOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"maxSize\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"maxStore\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"minPrice\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"minStore\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"node\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractINode\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"proxiableUUID\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"renounceOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"retake\",\"inputs\":[{\"name\":\"_pi\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"settle\",\"inputs\":[{\"name\":\"_m\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"streamPrice\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"token\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIERC20\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"upgradeToAndCall\",\"inputs\":[{\"name\":\"newImplementation\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"withdraw\",\"inputs\":[{\"name\":\"_m\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"AddPiece\",\"inputs\":[{\"name\":\"_a\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"_pi\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"},{\"name\":\"_e\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"AddReplica\",\"inputs\":[{\"name\":\"_a\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"_ri\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"},{\"name\":\"_sri\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Initialized\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"previousOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Retake\",\"inputs\":[{\"name\":\"_a\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"_pi\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"},{\"name\":\"_m\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Settle\",\"inputs\":[{\"name\":\"_a\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"_m\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Upgraded\",\"inputs\":[{\"name\":\"implementation\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Withdraw\",\"inputs\":[{\"name\":\"_a\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"_m\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"AddressEmptyCode\",\"inputs\":[{\"name\":\"target\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ERC1967InvalidImplementation\",\"inputs\":[{\"name\":\"implementation\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ERC1967NonPayable\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"FailedCall\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidInitialization\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotInitializing\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OwnableInvalidOwner\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"OwnableUnauthorizedAccount\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"SafeERC20FailedOperation\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"UUPSUnauthorizedCallContext\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"UUPSUnsupportedProxiableUUID\",\"inputs\":[{\"name\":\"slot\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]}]",
	Bin: "0x60a06040523073ffffffffffffffffffffffffffffffffffffffff1660809073ffffffffffffffffffffffffffffffffffffffff1681525034801562000043575f80fd5b50620000546200005a60201b60201c565b620001e1565b5f6200006b6200015e60201b60201c565b9050805f0160089054906101000a900460ff1615620000b6576040517ff92ee8a900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b67ffffffffffffffff8016815f015f9054906101000a900467ffffffffffffffff1667ffffffffffffffff16146200015b5767ffffffffffffffff815f015f6101000a81548167ffffffffffffffff021916908367ffffffffffffffff1602179055507fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d267ffffffffffffffff604051620001529190620001c6565b60405180910390a15b50565b5f80620001706200017960201b60201c565b90508091505090565b5f7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a005f1b905090565b5f67ffffffffffffffff82169050919050565b620001c081620001a2565b82525050565b5f602082019050620001db5f830184620001b5565b92915050565b6080516165dc620002085f395f8181613cd601528181613d2b0152613ee501526165dc5ff3fe608060405260043610610219575f3560e01c80638879512511610122578063ba8b2618116100aa578063ee91365b1161006e578063ee91365b1461081d578063f2fde38b1461085a578063fc0c546a14610882578063fca5be7a146108ac578063ffa1ad74146108d657610219565b8063ba8b261814610715578063bcbfd56b14610751578063c3aa38041461078d578063d70754ec146107c9578063e45be8eb146107f357610219565b80638ef0e895116100f15780638ef0e8951461061d578063900cf0cf146106455780639108544c1461066f578063ad3cb1cc146106ab578063b650b504146106d557610219565b80638879512514610565578063894d35f81461058f5780638da5cb5b146105cb5780638df82800146105f557610219565b806352d1902d116101a5578063715018a611610174578063715018a61461048557806375011aba1461049b5780637a87f0c4146104c55780637cf1ee911461050157806385b670951461052957610219565b806352d1902d146103cb57806362b98032146103f55780636a42b8f81461041d5780636aadfac71461044757610219565b80632565b159116101ec5780632565b159146102e55780632e1a7d4d1461030f5780634b4ffccd146103375780634f1ef2861461037357806350979d7e1461038f57610219565b8063123f84591461021d57806318198ad714610245578063194cbaf5146102815780631ce85e7c146102a9575b5f80fd5b348015610228575f80fd5b50610243600480360381019061023e9190614d15565b610900565b005b348015610250575f80fd5b5061026b60048036038101906102669190614d9a565b610c80565b6040516102789190614df0565b60405180910390f35b34801561028c575f80fd5b506102a760048036038101906102a29190614fd8565b610ced565b005b3480156102b4575f80fd5b506102cf60048036038101906102ca9190615078565b61191b565b6040516102dc91906150ce565b60405180910390f35b3480156102f0575f80fd5b506102f9611955565b60405161030691906150ce565b60405180910390f35b34801561031a575f80fd5b5061033560048036038101906103309190615111565b61196f565b005b348015610342575f80fd5b5061035d6004803603810190610358919061513c565b611a5d565b60405161036a91906151d8565b60405180910390f35b61038d600480360381019061038891906151f1565b611bc6565b005b34801561039a575f80fd5b506103b560048036038101906103b09190614d9a565b611be5565b6040516103c291906150ce565b60405180910390f35b3480156103d6575f80fd5b506103df611c64565b6040516103ec9190615263565b60405180910390f35b348015610400575f80fd5b5061041b6004803603810190610416919061513c565b611c95565b005b348015610428575f80fd5b50610431611db0565b60405161043e91906150ce565b60405180910390f35b348015610452575f80fd5b5061046d60048036038101906104689190614d15565b611dca565b60405161047c9392919061528b565b60405180910390f35b348015610490575f80fd5b50610499611e8e565b005b3480156104a6575f80fd5b506104af611ea1565b6040516104bc9190614df0565b60405180910390f35b3480156104d0575f80fd5b506104eb60048036038101906104e69190614d15565b611ea7565b6040516104f89190615393565b60405180910390f35b34801561050c575f80fd5b506105276004803603810190610522919061544e565b612169565b005b348015610534575f80fd5b5061054f600480360381019061054a919061513c565b612475565b60405161055c9190614df0565b60405180910390f35b348015610570575f80fd5b506105796124bb565b60405161058691906150ce565b60405180910390f35b34801561059a575f80fd5b506105b560048036038101906105b091906154c7565b6124d4565b6040516105c291906150ce565b60405180910390f35b3480156105d6575f80fd5b506105df61265e565b6040516105ec9190615542565b60405180910390f35b348015610600575f80fd5b5061061b60048036038101906106169190615111565b612693565b005b348015610628575f80fd5b50610643600480360381019061063e919061555b565b612a14565b005b348015610650575f80fd5b50610659613273565b604051610666919061566f565b60405180910390f35b34801561067a575f80fd5b5061069560048036038101906106909190614d15565b613298565b6040516106a291906156d7565b60405180910390f35b3480156106b6575f80fd5b506106bf61339a565b6040516106cc919061576a565b60405180910390f35b3480156106e0575f80fd5b506106fb60048036038101906106f6919061578a565b6133d3565b60405161070c9594939291906157c8565b60405180910390f35b348015610720575f80fd5b5061073b60048036038101906107369190615078565b61358b565b60405161074891906150ce565b60405180910390f35b34801561075c575f80fd5b5061077760048036038101906107729190614d9a565b6135c5565b60405161078491906150ce565b60405180910390f35b348015610798575f80fd5b506107b360048036038101906107ae9190614d9a565b613642565b6040516107c09190615846565b60405180910390f35b3480156107d4575f80fd5b506107dd613749565b6040516107ea919061587f565b60405180910390f35b3480156107fe575f80fd5b5061080761376e565b6040516108149190614df0565b60405180910390f35b348015610828575f80fd5b50610843600480360381019061083e9190615898565b613774565b6040516108519291906158d6565b60405180910390f35b348015610865575f80fd5b50610880600480360381019061087b919061513c565b613821565b005b34801561088d575f80fd5b506108966138a5565b6040516108a3919061591d565b60405180910390f35b3480156108b7575f80fd5b506108c06138c8565b6040516108cd91906150ce565b60405180910390f35b3480156108e1575f80fd5b506108ea6138e2565b6040516108f7919061576a565b60405180910390f35b60015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663919840ad6040518163ffffffff1660e01b81526004015f604051808303815f87803b158015610966575f80fd5b505af1158015610978573d5f803e3d5ffd5b505050505f60015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16639fa6a6e36040518163ffffffff1660e01b8152600401602060405180830381865afa1580156109e7573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610a0b919061594a565b90503373ffffffffffffffffffffffffffffffffffffffff1660068367ffffffffffffffff1681548110610a4257610a41615975565b5b905f5260205f2090600502016001015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1614610ac7576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610abe906159ec565b60405180910390fd5b8067ffffffffffffffff16600260149054906101000a900467ffffffffffffffff1660068467ffffffffffffffff1681548110610b0757610b06615975565b5b905f5260205f2090600502015f01600a9054906101000a900467ffffffffffffffff16610b349190615a37565b67ffffffffffffffff1610610b7e576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610b7590615abc565b60405180910390fd5b5f60068367ffffffffffffffff1681548110610b9d57610b9c615975565b5b905f5260205f2090600502016004015490505f60068467ffffffffffffffff1681548110610bce57610bcd615975565b5b905f5260205f20906005020160040181905550610c2b33825f8054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1661391b9092919063ffffffff16565b3373ffffffffffffffffffffffffffffffffffffffff167f44647ea87405778ba90a8ca28d2b772b1350bcd5950e5111bbc8c506a5c9632a8483604051610c73929190615ada565b60405180910390a2505050565b5f600e5f8473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f8367ffffffffffffffff1667ffffffffffffffff1681526020019081526020015f2060010154905092915050565b60015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663919840ad6040518163ffffffff1660e01b81526004015f604051808303815f87803b158015610d53575f80fd5b505af1158015610d65573d5f803e3d5ffd5b505050505f60015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16639fa6a6e36040518163ffffffff1660e01b8152600401602060405180830381865afa158015610dd4573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610df8919061594a565b90505f60068667ffffffffffffffff1681548110610e1957610e18615975565b5b905f5260205f2090600502016004015411610e69576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610e6090615b4b565b60405180910390fd5b8067ffffffffffffffff16600260149054906101000a900467ffffffffffffffff1660068767ffffffffffffffff1681548110610ea957610ea8615975565b5b905f5260205f2090600502015f01600a9054906101000a900467ffffffffffffffff16610ed69190615a37565b67ffffffffffffffff1611610f20576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610f1790615bb3565b60405180910390fd5b8360ff1660068667ffffffffffffffff1681548110610f4257610f41615975565b5b905f5260205f2090600502015f015f9054906101000a900460ff1660ff161015610fa1576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610f9890615c1b565b60405180910390fd5b5f600b87604051610fb29190615c7d565b90815260200160405180910390205f9054906101000a900467ffffffffffffffff1667ffffffffffffffff161461101e576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161101590615cdd565b60405180910390fd5b5f60085f8767ffffffffffffffff1667ffffffffffffffff1681526020019081526020015f205f8660ff1660ff1681526020019081526020015f205f9054906101000a900467ffffffffffffffff1667ffffffffffffffff16146110b7576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016110ae90615d45565b60405180910390fd5b60095f8667ffffffffffffffff1667ffffffffffffffff1681526020019081526020015f205f3373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f9054906101000a900460ff1615611164576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161115b90615dad565b60405180910390fd5b5f60068667ffffffffffffffff168154811061118357611182615975565b5b905f5260205f2090600502015f0160019054906101000a900460ff1660ff16601f6111ae9190615dcb565b600160068867ffffffffffffffff16815481106111ce576111cd615975565b5b905f5260205f2090600502015f0160029054906101000a900467ffffffffffffffff166111fb9190615e07565b6112059190615e6f565b60016112119190615a37565b90506180006001826112239190615e07565b61122d9190615e6f565b60016112399190615a37565b90505f8167ffffffffffffffff168360068967ffffffffffffffff168154811061126657611265615975565b5b905f5260205f2090600502015f0160129054906101000a900467ffffffffffffffff166112939190615e07565b67ffffffffffffffff1660068967ffffffffffffffff16815481106112bb576112ba615975565b5b905f5260205f209060050201600301546112d59190615e9f565b6112df9190615e9f565b90508060068867ffffffffffffffff1681548110611300576112ff615975565b5b905f5260205f2090600502016004015f82825461131d9190615ee0565b9250508190555060055460068867ffffffffffffffff168154811061134557611344615975565b5b905f5260205f2090600502016004015f8282546113629190615ee0565b92505081905550600554600f5f60068a67ffffffffffffffff168154811061138d5761138c615975565b5b905f5260205f2090600502016002015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f8282546114039190615f13565b92505081905550611412614b99565b87815f019067ffffffffffffffff16908167ffffffffffffffff168152505033816020019073ffffffffffffffffffffffffffffffffffffffff16908173ffffffffffffffffffffffffffffffffffffffff16815250508585604051611479929190615f6a565b60405180910390208160400181815250505f600a805490509050600a82908060018154018082558091505060019003905f5260205f2090600202015f909190919091505f820151815f015f6101000a81548167ffffffffffffffff021916908367ffffffffffffffff1602179055506020820151815f0160086101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060408201518160010155505080600b8b60405161154b9190615c7d565b90815260200160405180910390205f6101000a81548167ffffffffffffffff021916908367ffffffffffffffff1602179055508060085f8b67ffffffffffffffff1667ffffffffffffffff1681526020019081526020015f205f8a60ff1660ff1681526020019081526020015f205f6101000a81548167ffffffffffffffff021916908367ffffffffffffffff160217905550600160095f8b67ffffffffffffffff1667ffffffffffffffff1681526020019081526020015f205f3373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f6101000a81548160ff0219169083151502179055508367ffffffffffffffff1660068a67ffffffffffffffff168154811061167f5761167e615975565b5b905f5260205f209060050201600301546116999190615e9f565b92506116e8338260068c67ffffffffffffffff16815481106116be576116bd615975565b5b905f5260205f2090600502015f0160129054906101000a900467ffffffffffffffff16888761396e565b5f8060025f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166361e728b13360016040518363ffffffff1660e01b8152600401611747929190615fbb565b6040805180830381865afa158015611761573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190611785919061602b565b91509150811561182357600c5f3373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f206001015460076117dc9190615e9f565b81101561181e576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401611815906160b3565b60405180910390fd5b61185e565b6040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016118559061611b565b60405180910390fd5b3373ffffffffffffffffffffffffffffffffffffffff167f4b079e585085a3c94b2865036e598391fad4359764b36da39e21897be56349ef846001600c5f3373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f0160089054906101000a900467ffffffffffffffff166118f79190615e07565b604051611905929190616139565b60405180910390a2505050505050505050505050565b5f60078260405161192c9190615c7d565b90815260200160405180910390205f9054906101000a900467ffffffffffffffff169050919050565b600360109054906101000a900467ffffffffffffffff1681565b80600f5f3373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f8282546119bb9190615ee0565b92505081905550611a0c33825f8054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1661391b9092919063ffffffff16565b3373ffffffffffffffffffffffffffffffffffffffff167f884edad9ce6fa2440d8a54cc123490eb96d2768479d49ff9c7366125a942436482604051611a529190614df0565b60405180910390a250565b611a65614bda565b6040518060800160405280600c5f8573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f015f9054906101000a900467ffffffffffffffff1667ffffffffffffffff168152602001600c5f8573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f0160089054906101000a900467ffffffffffffffff1667ffffffffffffffff168152602001600c5f8573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f20600101548152602001600c5f8573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f20600201548152509050919050565b611bce613cd4565b611bd782613dba565b611be18282613dc5565b5050565b5f600e5f8473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f8367ffffffffffffffff1667ffffffffffffffff1681526020019081526020015f205f015f9054906101000a900467ffffffffffffffff16905092915050565b5f611c6d613ee3565b7f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc5f1b905090565b60015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663919840ad6040518163ffffffff1660e01b81526004015f604051808303815f87803b158015611cfb575f80fd5b505af1158015611d0d573d5f803e3d5ffd5b505050505f60015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16639fa6a6e36040518163ffffffff1660e01b8152600401602060405180830381865afa158015611d7c573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190611da0919061594a565b9050611dac8282613f6a565b5050565b600260149054906101000a900467ffffffffffffffff1681565b5f805f60068467ffffffffffffffff1681548110611deb57611dea615975565b5b905f5260205f2090600502015f015f9054906101000a900460ff1660068567ffffffffffffffff1681548110611e2457611e23615975565b5b905f5260205f2090600502015f0160019054906101000a900460ff1660068667ffffffffffffffff1681548110611e5e57611e5d615975565b5b905f5260205f2090600502015f0160129054906101000a900467ffffffffffffffff169250925092509193909250565b611e966144b2565b611e9f5f614539565b565b60055481565b611eaf614c12565b60405180610120016040528060068467ffffffffffffffff1681548110611ed957611ed8615975565b5b905f5260205f2090600502015f015f9054906101000a900460ff1660ff16815260200160068467ffffffffffffffff1681548110611f1a57611f19615975565b5b905f5260205f2090600502015f0160019054906101000a900460ff1660ff16815260200160068467ffffffffffffffff1681548110611f5c57611f5b615975565b5b905f5260205f2090600502015f0160029054906101000a900467ffffffffffffffff1667ffffffffffffffff16815260200160068467ffffffffffffffff1681548110611fac57611fab615975565b5b905f5260205f2090600502015f01600a9054906101000a900467ffffffffffffffff1667ffffffffffffffff16815260200160068467ffffffffffffffff1681548110611ffc57611ffb615975565b5b905f5260205f2090600502015f0160129054906101000a900467ffffffffffffffff1667ffffffffffffffff16815260200160068467ffffffffffffffff168154811061204c5761204b615975565b5b905f5260205f2090600502016001015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200160068467ffffffffffffffff16815481106120b4576120b3615975565b5b905f5260205f2090600502016002015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200160068467ffffffffffffffff168154811061211c5761211b615975565b5b905f5260205f20906005020160030154815260200160068467ffffffffffffffff168154811061214f5761214e615975565b5b905f5260205f209060050201600401548152509050919050565b5f61217261460a565b90505f815f0160089054906101000a900460ff161590505f825f015f9054906101000a900467ffffffffffffffff1690505f808267ffffffffffffffff161480156121ba5750825b90505f60018367ffffffffffffffff161480156121ed57505f3073ffffffffffffffffffffffffffffffffffffffff163b145b9050811580156121fb575080155b15612232576040517ff92ee8a900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6001855f015f6101000a81548167ffffffffffffffff021916908367ffffffffffffffff160217905550831561227f576001855f0160086101000a81548160ff0219169083151502179055505b6122888661461d565b895f806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508860015f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508760025f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550865f0151600260146101000a81548167ffffffffffffffff021916908367ffffffffffffffff160217905550866020015160035f6101000a81548167ffffffffffffffff021916908367ffffffffffffffff1602179055508660400151600360086101000a81548167ffffffffffffffff021916908367ffffffffffffffff1602179055508660600151600360106101000a81548167ffffffffffffffff021916908367ffffffffffffffff16021790555086608001516004819055508660a001516005819055508315612469575f855f0160086101000a81548160ff0219169083151502179055507fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d260016040516124609190616190565b60405180910390a15b50505050505050505050565b5f600f5f8373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f20549050919050565b60035f9054906101000a900467ffffffffffffffff1681565b5f600d5f8573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f8467ffffffffffffffff1667ffffffffffffffff1681526020019081526020015f205f9054906101000a900467ffffffffffffffff1692508267ffffffffffffffff16600b836040516125669190615c7d565b90815260200160405180910390205f9054906101000a900467ffffffffffffffff1667ffffffffffffffff16146125d2576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016125c9906161f3565b60405180910390fd5b600a8367ffffffffffffffff16815481106125f0576125ef615975565b5b905f5260205f2090600202015f015f9054906101000a900467ffffffffffffffff16925060068367ffffffffffffffff168154811061263257612631615975565b5b905f5260205f2090600502015f0160129054906101000a900467ffffffffffffffff1690509392505050565b5f80612668614631565b9050805f015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1691505090565b60015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663919840ad6040518163ffffffff1660e01b81526004015f604051808303815f87803b1580156126f9575f80fd5b505af115801561270b573d5f803e3d5ffd5b505050505f60015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16639fa6a6e36040518163ffffffff1660e01b8152600401602060405180830381865afa15801561277a573d5f803e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061279e919061594a565b90506127aa3382613f6a565b5f8060025f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166361e728b13360016040518363ffffffff1660e01b8152600401612809929190615fbb565b6040805180830381865afa158015612823573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190612847919061602b565b9150915081156128e557600c5f3373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f2060010154600761289e9190615e9f565b8110156128e0576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016128d7906160b3565b60405180910390fd5b612920565b6040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016129179061611b565b60405180910390fd5b83600c5f3373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f206002015f82825461296f9190615ee0565b925050819055506129c033855f8054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1661391b9092919063ffffffff16565b3373ffffffffffffffffffffffffffffffffffffffff167fa3788eedc10ef3ec6982384227c562c6666cf2b6af4f6a583b6d5d0f2ba0d20485604051612a069190614df0565b60405180910390a250505050565b60015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663919840ad6040518163ffffffff1660e01b81526004015f604051808303815f87803b158015612a7a575f80fd5b505af1158015612a8c573d5f803e3d5ffd5b505050505f60015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16639fa6a6e36040518163ffffffff1660e01b8152600401602060405180830381865afa158015612afb573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190612b1f919061594a565b90505f8667ffffffffffffffff1611612b6d576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401612b649061625b565b60405180910390fd5b5f8360ff1611612bb2576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401612ba9906162c3565b60405180910390fd5b8360ff168360ff1610612bfa576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401612bf19061632b565b60405180910390fd5b60035f9054906101000a900467ffffffffffffffff1681612c1b9190615a37565b67ffffffffffffffff168567ffffffffffffffff161015612c71576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401612c6890616393565b60405180910390fd5b600360089054906101000a900467ffffffffffffffff1681612c939190615a37565b67ffffffffffffffff168567ffffffffffffffff161115612ce9576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401612ce0906163fb565b60405180910390fd5b600454871015612d2e576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401612d2590616463565b60405180910390fd5b5f600789604051612d3f9190615c7d565b90815260200160405180910390205f9054906101000a900467ffffffffffffffff1667ffffffffffffffff1614612dab576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401612da2906164cb565b60405180910390fd5b5f8360ff16601f612dbc9190615dcb565b600188612dc99190615e07565b612dd39190615e6f565b6001612ddf9190615a37565b9050600360109054906101000a900467ffffffffffffffff1667ffffffffffffffff168460ff1682612e119190615dcb565b67ffffffffffffffff161115612e5c576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401612e5390616533565b60405180910390fd5b618000600182612e6c9190615e07565b612e769190615e6f565b6001612e829190615a37565b90505f8560ff166005548367ffffffffffffffff16858a612ea39190615e07565b67ffffffffffffffff168c612eb89190615e9f565b612ec29190615e9f565b612ecc9190615f13565b612ed69190615e9f565b9050612f243330835f8054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16614658909392919063ffffffff16565b5f6006805490509050612f35614c12565b338160a0019073ffffffffffffffffffffffffffffffffffffffff16908173ffffffffffffffffffffffffffffffffffffffff168152505082816101000181815250508a8160e001818152505089816040019067ffffffffffffffff16908167ffffffffffffffff168152505084816060019067ffffffffffffffff16908167ffffffffffffffff168152505088816080019067ffffffffffffffff16908167ffffffffffffffff168152505087815f019060ff16908160ff168152505086816020019060ff16908160ff1681525050858160c0019073ffffffffffffffffffffffffffffffffffffffff16908173ffffffffffffffffffffffffffffffffffffffff1681525050600681908060018154018082558091505060019003905f5260205f2090600502015f909190919091505f820151815f015f6101000a81548160ff021916908360ff1602179055506020820151815f0160016101000a81548160ff021916908360ff1602179055506040820151815f0160026101000a81548167ffffffffffffffff021916908367ffffffffffffffff1602179055506060820151815f01600a6101000a81548167ffffffffffffffff021916908367ffffffffffffffff1602179055506080820151815f0160126101000a81548167ffffffffffffffff021916908367ffffffffffffffff16021790555060a0820151816001015f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060c0820151816002015f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060e08201518160030155610100820151816004015550508160078d6040516131e29190615c7d565b90815260200160405180910390205f6101000a81548167ffffffffffffffff021916908367ffffffffffffffff1602179055503373ffffffffffffffffffffffffffffffffffffffff167f09ff0fb5f488978fdffbf3223e63e914f2a9b0c228844b8e2caedad8a85d220c838760405161325d929190616139565b60405180910390a2505050505050505050505050565b60015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b6132a0614b99565b6040518060600160405280600a8467ffffffffffffffff16815481106132c9576132c8615975565b5b905f5260205f2090600202015f015f9054906101000a900467ffffffffffffffff1667ffffffffffffffff168152602001600a8467ffffffffffffffff168154811061331857613317615975565b5b905f5260205f2090600202015f0160089054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001600a8467ffffffffffffffff16815481106133805761337f615975565b5b905f5260205f209060020201600101548152509050919050565b6040518060400160405280600581526020017f352e302e3000000000000000000000000000000000000000000000000000000081525081565b5f805f805f8060068867ffffffffffffffff16815481106133f7576133f6615975565b5b905f5260205f2090600502015f0160019054906101000a900460ff1660ff16601f6134229190615dcb565b600160068a67ffffffffffffffff168154811061344257613441615975565b5b905f5260205f2090600502015f0160029054906101000a900467ffffffffffffffff1661346f9190615e07565b6134799190615e6f565b60016134859190615a37565b905060068867ffffffffffffffff16815481106134a5576134a4615975565b5b905f5260205f2090600502015f015f9054906101000a900460ff1660068967ffffffffffffffff16815481106134de576134dd615975565b5b905f5260205f2090600502015f0160019054906101000a900460ff168260068b67ffffffffffffffff168154811061351957613518615975565b5b905f5260205f2090600502016002015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff16600a8b67ffffffffffffffff168154811061356657613565615975565b5b905f5260205f2090600202016001015495509550955095509550509295509295909350565b5f600b8260405161359c9190615c7d565b90815260200160405180910390205f9054906101000a900467ffffffffffffffff169050919050565b5f600d5f8473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f8367ffffffffffffffff1667ffffffffffffffff1681526020019081526020015f205f9054906101000a900467ffffffffffffffff16905092915050565b61364a614ca5565b6040518060400160405280600e5f8673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f8567ffffffffffffffff1667ffffffffffffffff1681526020019081526020015f205f015f9054906101000a900467ffffffffffffffff1667ffffffffffffffff168152602001600e5f8673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f8567ffffffffffffffff1667ffffffffffffffff1681526020019081526020015f2060010154815250905092915050565b60025f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b60045481565b5f805f60085f8667ffffffffffffffff1667ffffffffffffffff1681526020019081526020015f205f8560ff1660ff1681526020019081526020015f205f9054906101000a900467ffffffffffffffff16905080600a8267ffffffffffffffff16815481106137e6576137e5615975565b5b905f5260205f2090600202015f0160089054906101000a900473ffffffffffffffffffffffffffffffffffffffff1692509250509250929050565b6138296144b2565b5f73ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1603613899575f6040517f1e4fbdf70000000000000000000000000000000000000000000000000000000081526004016138909190615542565b60405180910390fd5b6138a281614539565b50565b5f8054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b600360089054906101000a900467ffffffffffffffff1681565b6040518060400160405280600581526020017f322e302e3000000000000000000000000000000000000000000000000000000081525081565b61392883838360016146ad565b61396957826040517f5274afe70000000000000000000000000000000000000000000000000000000081526004016139609190615542565b60405180910390fd5b505050565b6139788583613f6a565b83600d5f8773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f600c5f8973ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f0160089054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681526020019081526020015f205f6101000a81548167ffffffffffffffff021916908367ffffffffffffffff1602179055506001600c5f8773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f0160088282829054906101000a900467ffffffffffffffff16613ab49190615a37565b92506101000a81548167ffffffffffffffff021916908367ffffffffffffffff16021790555080600c5f8773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f206001015f828254613b299190615f13565b925050819055506001600e5f8773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f8467ffffffffffffffff1667ffffffffffffffff1681526020019081526020015f205f015f8282829054906101000a900467ffffffffffffffff16613bb59190615a37565b92506101000a81548167ffffffffffffffff021916908367ffffffffffffffff16021790555080600e5f8773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f8467ffffffffffffffff1667ffffffffffffffff1681526020019081526020015f206001015f828254613c4d9190615f13565b9250508190555080600e5f8773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f8567ffffffffffffffff1667ffffffffffffffff1681526020019081526020015f206001015f828254613cc69190615f13565b925050819055505050505050565b7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff163073ffffffffffffffffffffffffffffffffffffffff161480613d8157507f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff16613d6861470f565b73ffffffffffffffffffffffffffffffffffffffff1614155b15613db8576040517fe07c8dba00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b565b613dc26144b2565b50565b8173ffffffffffffffffffffffffffffffffffffffff166352d1902d6040518163ffffffff1660e01b8152600401602060405180830381865afa925050508015613e2d57506040513d601f19601f82011682018060405250810190613e2a919061657b565b60015b613e6e57816040517f4c9c8ce3000000000000000000000000000000000000000000000000000000008152600401613e659190615542565b60405180910390fd5b7f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc5f1b8114613ed457806040517faa1d49a4000000000000000000000000000000000000000000000000000000008152600401613ecb9190615263565b60405180910390fd5b613ede8383614762565b505050565b7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff163073ffffffffffffffffffffffffffffffffffffffff1614613f68576040517fe07c8dba00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b565b5f600c5f8473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f015f9054906101000a900467ffffffffffffffff1667ffffffffffffffff16036140335780600c5f8473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f015f6101000a81548167ffffffffffffffff021916908367ffffffffffffffff1602179055505b8067ffffffffffffffff16600c5f8473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f015f9054906101000a900467ffffffffffffffff1667ffffffffffffffff1610156144ae575f6001600c5f8573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f015f9054906101000a900467ffffffffffffffff166141019190615a37565b90505b8167ffffffffffffffff168167ffffffffffffffff161161444757600e5f8473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f8267ffffffffffffffff1667ffffffffffffffff1681526020019081526020015f2060010154600c5f8573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f206001015f8282546141d19190615ee0565b92505081905550600c5f8473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f2060010154600e5f8573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f8367ffffffffffffffff1667ffffffffffffffff1681526020019081526020015f2060010181905550600c5f8473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f0160089054906101000a900467ffffffffffffffff16600e5f8573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f8367ffffffffffffffff1667ffffffffffffffff1681526020019081526020015f205f015f6101000a81548167ffffffffffffffff021916908367ffffffffffffffff16021790555060028167ffffffffffffffff161061443357600e5f8473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f6002836143b89190615e07565b67ffffffffffffffff1667ffffffffffffffff1681526020019081526020015f2060010154600c5f8573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f206002015f82825461442b9190615f13565b925050819055505b6001816144409190615a37565b9050614104565b5080600c5f8473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f015f6101000a81548167ffffffffffffffff021916908367ffffffffffffffff1602179055505b5050565b6144ba6147d4565b73ffffffffffffffffffffffffffffffffffffffff166144d861265e565b73ffffffffffffffffffffffffffffffffffffffff1614614537576144fb6147d4565b6040517f118cdaa700000000000000000000000000000000000000000000000000000000815260040161452e9190615542565b60405180910390fd5b565b5f614542614631565b90505f815f015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905082825f015f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508273ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a3505050565b5f806146146147db565b90508091505090565b614625614804565b61462e81614844565b50565b5f7f9016d09d72d40fdae2fd8ceac6b6234c7706214fd39c1cd1e609a0528c199300905090565b6146668484848460016148c8565b6146a757836040517f5274afe700000000000000000000000000000000000000000000000000000000815260040161469e9190615542565b60405180910390fd5b50505050565b5f8063a9059cbb60e01b9050604051815f525f1960601c86166004528460245260205f60445f808b5af1925060015f511483166147015783831516156146f5573d5f823e3d81fd5b5f873b113d1516831692505b806040525050949350505050565b5f61473b7f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc5f1b614939565b5f015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b61476b82614942565b8173ffffffffffffffffffffffffffffffffffffffff167fbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b60405160405180910390a25f815111156147c7576147c18282614a0b565b506147d0565b6147cf614afc565b5b5050565b5f33905090565b5f7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a005f1b905090565b61480c614b38565b614842576040517fd7e6bcf800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b565b61484c614804565b5f73ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16036148bc575f6040517f1e4fbdf70000000000000000000000000000000000000000000000000000000081526004016148b39190615542565b60405180910390fd5b6148c581614539565b50565b5f806323b872dd60e01b9050604051815f525f1960601c87166004525f1960601c86166024528460445260205f60645f808c5af1925060015f5114831661492657838315161561491a573d5f823e3d81fd5b5f883b113d1516831692505b806040525f606052505095945050505050565b5f819050919050565b5f8173ffffffffffffffffffffffffffffffffffffffff163b0361499d57806040517f4c9c8ce30000000000000000000000000000000000000000000000000000000081526004016149949190615542565b60405180910390fd5b806149c97f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc5f1b614939565b5f015f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050565b60605f614a188484614b56565b9050808015614a4e57505f614a2b614b6a565b1180614a4d57505f8473ffffffffffffffffffffffffffffffffffffffff163b115b5b15614a6357614a5b614b71565b915050614af6565b8015614aa657836040517f9996b315000000000000000000000000000000000000000000000000000000008152600401614a9d9190615542565b60405180910390fd5b5f614aaf614b6a565b1115614ac257614abd614b8e565b614af4565b6040517fd6bda27500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b505b92915050565b5f341115614b36576040517fb398979f00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b565b5f614b4161460a565b5f0160089054906101000a900460ff16905090565b5f805f835160208501865af4905092915050565b5f3d905090565b606060405190503d81523d5f602083013e3d602001810160405290565b6040513d5f823e3d81fd5b60405180606001604052805f67ffffffffffffffff1681526020015f73ffffffffffffffffffffffffffffffffffffffff1681526020015f80191681525090565b60405180608001604052805f67ffffffffffffffff1681526020015f67ffffffffffffffff1681526020015f81526020015f81525090565b6040518061012001604052805f60ff1681526020015f60ff1681526020015f67ffffffffffffffff1681526020015f67ffffffffffffffff1681526020015f67ffffffffffffffff1681526020015f73ffffffffffffffffffffffffffffffffffffffff1681526020015f73ffffffffffffffffffffffffffffffffffffffff1681526020015f81526020015f81525090565b60405180604001604052805f67ffffffffffffffff1681526020015f81525090565b5f604051905090565b5f80fd5b5f80fd5b5f67ffffffffffffffff82169050919050565b614cf481614cd8565b8114614cfe575f80fd5b50565b5f81359050614d0f81614ceb565b92915050565b5f60208284031215614d2a57614d29614cd0565b5b5f614d3784828501614d01565b91505092915050565b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f614d6982614d40565b9050919050565b614d7981614d5f565b8114614d83575f80fd5b50565b5f81359050614d9481614d70565b92915050565b5f8060408385031215614db057614daf614cd0565b5b5f614dbd85828601614d86565b9250506020614dce85828601614d01565b9150509250929050565b5f819050919050565b614dea81614dd8565b82525050565b5f602082019050614e035f830184614de1565b92915050565b5f80fd5b5f80fd5b5f601f19601f8301169050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b614e5782614e11565b810181811067ffffffffffffffff82111715614e7657614e75614e21565b5b80604052505050565b5f614e88614cc7565b9050614e948282614e4e565b919050565b5f67ffffffffffffffff821115614eb357614eb2614e21565b5b614ebc82614e11565b9050602081019050919050565b828183375f83830152505050565b5f614ee9614ee484614e99565b614e7f565b905082815260208101848484011115614f0557614f04614e0d565b5b614f10848285614ec9565b509392505050565b5f82601f830112614f2c57614f2b614e09565b5b8135614f3c848260208601614ed7565b91505092915050565b5f60ff82169050919050565b614f5a81614f45565b8114614f64575f80fd5b50565b5f81359050614f7581614f51565b92915050565b5f80fd5b5f80fd5b5f8083601f840112614f9857614f97614e09565b5b8235905067ffffffffffffffff811115614fb557614fb4614f7b565b5b602083019150836001820283011115614fd157614fd0614f7f565b5b9250929050565b5f805f805f60808688031215614ff157614ff0614cd0565b5b5f86013567ffffffffffffffff81111561500e5761500d614cd4565b5b61501a88828901614f18565b955050602061502b88828901614d01565b945050604061503c88828901614f67565b935050606086013567ffffffffffffffff81111561505d5761505c614cd4565b5b61506988828901614f83565b92509250509295509295909350565b5f6020828403121561508d5761508c614cd0565b5b5f82013567ffffffffffffffff8111156150aa576150a9614cd4565b5b6150b684828501614f18565b91505092915050565b6150c881614cd8565b82525050565b5f6020820190506150e15f8301846150bf565b92915050565b6150f081614dd8565b81146150fa575f80fd5b50565b5f8135905061510b816150e7565b92915050565b5f6020828403121561512657615125614cd0565b5b5f615133848285016150fd565b91505092915050565b5f6020828403121561515157615150614cd0565b5b5f61515e84828501614d86565b91505092915050565b61517081614cd8565b82525050565b61517f81614dd8565b82525050565b608082015f8201516151995f850182615167565b5060208201516151ac6020850182615167565b5060408201516151bf6040850182615176565b5060608201516151d26060850182615176565b50505050565b5f6080820190506151eb5f830184615185565b92915050565b5f806040838503121561520757615206614cd0565b5b5f61521485828601614d86565b925050602083013567ffffffffffffffff81111561523557615234614cd4565b5b61524185828601614f18565b9150509250929050565b5f819050919050565b61525d8161524b565b82525050565b5f6020820190506152765f830184615254565b92915050565b61528581614f45565b82525050565b5f60608201905061529e5f83018661527c565b6152ab602083018561527c565b6152b860408301846150bf565b949350505050565b6152c981614f45565b82525050565b6152d881614d5f565b82525050565b61012082015f8201516152f35f8501826152c0565b50602082015161530660208501826152c0565b5060408201516153196040850182615167565b50606082015161532c6060850182615167565b50608082015161533f6080850182615167565b5060a082015161535260a08501826152cf565b5060c082015161536560c08501826152cf565b5060e082015161537860e0850182615176565b5061010082015161538d610100850182615176565b50505050565b5f610120820190506153a75f8301846152de565b92915050565b5f80fd5b5f60c082840312156153c6576153c56153ad565b5b6153d060c0614e7f565b90505f6153df84828501614d01565b5f8301525060206153f284828501614d01565b602083015250604061540684828501614d01565b604083015250606061541a84828501614d01565b606083015250608061542e848285016150fd565b60808301525060a0615442848285016150fd565b60a08301525092915050565b5f805f805f610140868803121561546857615467614cd0565b5b5f61547588828901614d86565b955050602061548688828901614d86565b945050604061549788828901614d86565b93505060606154a8888289016153b1565b9250506101206154ba88828901614d86565b9150509295509295909350565b5f805f606084860312156154de576154dd614cd0565b5b5f6154eb86828701614d86565b93505060206154fc86828701614d01565b925050604084013567ffffffffffffffff81111561551d5761551c614cd4565b5b61552986828701614f18565b9150509250925092565b61553c81614d5f565b82525050565b5f6020820190506155555f830184615533565b92915050565b5f805f805f805f60e0888a03121561557657615575614cd0565b5b5f88013567ffffffffffffffff81111561559357615592614cd4565b5b61559f8a828b01614f18565b97505060206155b08a828b016150fd565b96505060406155c18a828b01614d01565b95505060606155d28a828b01614d01565b94505060806155e38a828b01614f67565b93505060a06155f48a828b01614f67565b92505060c06156058a828b01614d86565b91505092959891949750929550565b5f819050919050565b5f61563761563261562d84614d40565b615614565b614d40565b9050919050565b5f6156488261561d565b9050919050565b5f6156598261563e565b9050919050565b6156698161564f565b82525050565b5f6020820190506156825f830184615660565b92915050565b6156918161524b565b82525050565b606082015f8201516156ab5f850182615167565b5060208201516156be60208501826152cf565b5060408201516156d16040850182615688565b50505050565b5f6060820190506156ea5f830184615697565b92915050565b5f81519050919050565b5f82825260208201905092915050565b5f5b8381101561572757808201518184015260208101905061570c565b5f8484015250505050565b5f61573c826156f0565b61574681856156fa565b935061575681856020860161570a565b61575f81614e11565b840191505092915050565b5f6020820190508181035f8301526157828184615732565b905092915050565b5f80604083850312156157a05761579f614cd0565b5b5f6157ad85828601614d01565b92505060206157be85828601614d01565b9150509250929050565b5f60a0820190506157db5f83018861527c565b6157e8602083018761527c565b6157f560408301866150bf565b6158026060830185615533565b61580f6080830184615254565b9695505050505050565b604082015f82015161582d5f850182615167565b5060208201516158406020850182615176565b50505050565b5f6040820190506158595f830184615819565b92915050565b5f6158698261563e565b9050919050565b6158798161585f565b82525050565b5f6020820190506158925f830184615870565b92915050565b5f80604083850312156158ae576158ad614cd0565b5b5f6158bb85828601614d01565b92505060206158cc85828601614f67565b9150509250929050565b5f6040820190506158e95f8301856150bf565b6158f66020830184615533565b9392505050565b5f6159078261563e565b9050919050565b615917816158fd565b82525050565b5f6020820190506159305f83018461590e565b92915050565b5f8151905061594481614ceb565b92915050565b5f6020828403121561595f5761595e614cd0565b5b5f61596c84828501615936565b91505092915050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52603260045260245ffd5b7f696e7620776300000000000000000000000000000000000000000000000000005f82015250565b5f6159d66006836156fa565b91506159e1826159a2565b602082019050919050565b5f6020820190508181035f830152615a03816159ca565b9050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b5f615a4182614cd8565b9150615a4c83614cd8565b9250828201905067ffffffffffffffff811115615a6c57615a6b615a0a565b5b92915050565b7f6561726c790000000000000000000000000000000000000000000000000000005f82015250565b5f615aa66005836156fa565b9150615ab182615a72565b602082019050919050565b5f6020820190508181035f830152615ad381615a9a565b9050919050565b5f604082019050615aed5f8301856150bf565b615afa6020830184614de1565b9392505050565b7f6e6f2070696563650000000000000000000000000000000000000000000000005f82015250565b5f615b356008836156fa565b9150615b4082615b01565b602082019050919050565b5f6020820190508181035f830152615b6281615b29565b9050919050565b7f746f6f206c6174650000000000000000000000000000000000000000000000005f82015250565b5f615b9d6008836156fa565b9150615ba882615b69565b602082019050919050565b5f6020820190508181035f830152615bca81615b91565b9050919050565b7f696e7620707269000000000000000000000000000000000000000000000000005f82015250565b5f615c056007836156fa565b9150615c1082615bd1565b602082019050919050565b5f6020820190508181035f830152615c3281615bf9565b9050919050565b5f81519050919050565b5f81905092915050565b5f615c5782615c39565b615c618185615c43565b9350615c7181856020860161570a565b80840191505092915050565b5f615c888284615c4d565b915081905092915050565b7f65786973742072000000000000000000000000000000000000000000000000005f82015250565b5f615cc76007836156fa565b9150615cd282615c93565b602082019050919050565b5f6020820190508181035f830152615cf481615cbb565b9050919050565b7f65786973742070726900000000000000000000000000000000000000000000005f82015250565b5f615d2f6009836156fa565b9150615d3a82615cfb565b602082019050919050565b5f6020820190508181035f830152615d5c81615d23565b9050919050565b7f64757020616464720000000000000000000000000000000000000000000000005f82015250565b5f615d976008836156fa565b9150615da282615d63565b602082019050919050565b5f6020820190508181035f830152615dc481615d8b565b9050919050565b5f615dd582614cd8565b9150615de083614cd8565b9250828202615dee81614cd8565b9150808214615e0057615dff615a0a565b5b5092915050565b5f615e1182614cd8565b9150615e1c83614cd8565b9250828203905067ffffffffffffffff811115615e3c57615e3b615a0a565b5b92915050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601260045260245ffd5b5f615e7982614cd8565b9150615e8483614cd8565b925082615e9457615e93615e42565b5b828204905092915050565b5f615ea982614dd8565b9150615eb483614dd8565b9250828202615ec281614dd8565b91508282048414831517615ed957615ed8615a0a565b5b5092915050565b5f615eea82614dd8565b9150615ef583614dd8565b9250828203905081811115615f0d57615f0c615a0a565b5b92915050565b5f615f1d82614dd8565b9150615f2883614dd8565b9250828201905080821115615f4057615f3f615a0a565b5b92915050565b5f615f518385615c43565b9350615f5e838584614ec9565b82840190509392505050565b5f615f76828486615f46565b91508190509392505050565b5f819050919050565b5f615fa5615fa0615f9b84615f82565b615614565b614f45565b9050919050565b615fb581615f8b565b82525050565b5f604082019050615fce5f830185615533565b615fdb6020830184615fac565b9392505050565b5f8115159050919050565b615ff681615fe2565b8114616000575f80fd5b50565b5f8151905061601181615fed565b92915050565b5f81519050616025816150e7565b92915050565b5f806040838503121561604157616040614cd0565b5b5f61604e85828601616003565b925050602061605f85828601616017565b9150509250929050565b7f696e73756620706c6564676537000000000000000000000000000000000000005f82015250565b5f61609d600d836156fa565b91506160a882616069565b602082019050919050565b5f6020820190508181035f8301526160ca81616091565b9050919050565b7f696e76206e6f64650000000000000000000000000000000000000000000000005f82015250565b5f6161056008836156fa565b9150616110826160d1565b602082019050919050565b5f6020820190508181035f830152616132816160f9565b9050919050565b5f60408201905061614c5f8301856150bf565b61615960208301846150bf565b9392505050565b5f61617a61617561617084615f82565b615614565b614cd8565b9050919050565b61618a81616160565b82525050565b5f6020820190506161a35f830184616181565b92915050565b7f696e7620737200000000000000000000000000000000000000000000000000005f82015250565b5f6161dd6006836156fa565b91506161e8826161a9565b602082019050919050565b5f6020820190508181035f83015261620a816161d1565b9050919050565b7f696e762073697a650000000000000000000000000000000000000000000000005f82015250565b5f6162456008836156fa565b915061625082616211565b602082019050919050565b5f6020820190508181035f83015261627281616239565b9050919050565b7f696e762072736b000000000000000000000000000000000000000000000000005f82015250565b5f6162ad6007836156fa565b91506162b882616279565b602082019050919050565b5f6020820190508181035f8301526162da816162a1565b9050919050565b7f696e7620727300000000000000000000000000000000000000000000000000005f82015250565b5f6163156006836156fa565b9150616320826162e1565b602082019050919050565b5f6020820190508181035f83015261634281616309565b9050919050565b7f73686f72742064757200000000000000000000000000000000000000000000005f82015250565b5f61637d6009836156fa565b915061638882616349565b602082019050919050565b5f6020820190508181035f8301526163aa81616371565b9050919050565b7f65786365656420647572000000000000000000000000000000000000000000005f82015250565b5f6163e5600a836156fa565b91506163f0826163b1565b602082019050919050565b5f6020820190508181035f830152616412816163d9565b9050919050565b7f6d696e20707269636500000000000000000000000000000000000000000000005f82015250565b5f61644d6009836156fa565b915061645882616419565b602082019050919050565b5f6020820190508181035f83015261647a81616441565b9050919050565b7f70696563652065786973740000000000000000000000000000000000000000005f82015250565b5f6164b5600b836156fa565b91506164c082616481565b602082019050919050565b5f6020820190508181035f8301526164e2816164a9565b9050919050565b7f6578636565642073697a650000000000000000000000000000000000000000005f82015250565b5f61651d600b836156fa565b9150616528826164e9565b602082019050919050565b5f6020820190508181035f83015261654a81616511565b9050919050565b61655a8161524b565b8114616564575f80fd5b50565b5f8151905061657581616551565b92915050565b5f602082840312156165905761658f614cd0565b5b5f61659d84828501616567565b9150509291505056fea26469706673582212207c6c1e2c413f06fd678d2f8971059037398e319f752765985800ae1f08ba138a64736f6c63430008160033",
}

// PieceABI is the input ABI used to generate the binding from.
// Deprecated: Use PieceMetaData.ABI instead.
var PieceABI = PieceMetaData.ABI

// PieceBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use PieceMetaData.Bin instead.
var PieceBin = PieceMetaData.Bin

// DeployPiece deploys a new Ethereum contract, binding an instance of Piece to it.
func DeployPiece(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Piece, error) {
	parsed, err := PieceMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(PieceBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Piece{PieceCaller: PieceCaller{contract: contract}, PieceTransactor: PieceTransactor{contract: contract}, PieceFilterer: PieceFilterer{contract: contract}}, nil
}

// Piece is an auto generated Go binding around an Ethereum contract.
type Piece struct {
	PieceCaller     // Read-only binding to the contract
	PieceTransactor // Write-only binding to the contract
	PieceFilterer   // Log filterer for contract events
}

// PieceCaller is an auto generated read-only Go binding around an Ethereum contract.
type PieceCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PieceTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PieceTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PieceFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PieceFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PieceSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PieceSession struct {
	Contract     *Piece            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PieceCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PieceCallerSession struct {
	Contract *PieceCaller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// PieceTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PieceTransactorSession struct {
	Contract     *PieceTransactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PieceRaw is an auto generated low-level Go binding around an Ethereum contract.
type PieceRaw struct {
	Contract *Piece // Generic contract binding to access the raw methods on
}

// PieceCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PieceCallerRaw struct {
	Contract *PieceCaller // Generic read-only contract binding to access the raw methods on
}

// PieceTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PieceTransactorRaw struct {
	Contract *PieceTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPiece creates a new instance of Piece, bound to a specific deployed contract.
func NewPiece(address common.Address, backend bind.ContractBackend) (*Piece, error) {
	contract, err := bindPiece(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Piece{PieceCaller: PieceCaller{contract: contract}, PieceTransactor: PieceTransactor{contract: contract}, PieceFilterer: PieceFilterer{contract: contract}}, nil
}

// NewPieceCaller creates a new read-only instance of Piece, bound to a specific deployed contract.
func NewPieceCaller(address common.Address, caller bind.ContractCaller) (*PieceCaller, error) {
	contract, err := bindPiece(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PieceCaller{contract: contract}, nil
}

// NewPieceTransactor creates a new write-only instance of Piece, bound to a specific deployed contract.
func NewPieceTransactor(address common.Address, transactor bind.ContractTransactor) (*PieceTransactor, error) {
	contract, err := bindPiece(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PieceTransactor{contract: contract}, nil
}

// NewPieceFilterer creates a new log filterer instance of Piece, bound to a specific deployed contract.
func NewPieceFilterer(address common.Address, filterer bind.ContractFilterer) (*PieceFilterer, error) {
	contract, err := bindPiece(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PieceFilterer{contract: contract}, nil
}

// bindPiece binds a generic wrapper to an already deployed contract.
func bindPiece(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := PieceMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Piece *PieceRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Piece.Contract.PieceCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Piece *PieceRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Piece.Contract.PieceTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Piece *PieceRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Piece.Contract.PieceTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Piece *PieceCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Piece.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Piece *PieceTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Piece.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Piece *PieceTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Piece.Contract.contract.Transact(opts, method, params...)
}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_Piece *PieceCaller) UPGRADEINTERFACEVERSION(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Piece.contract.Call(opts, &out, "UPGRADE_INTERFACE_VERSION")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_Piece *PieceSession) UPGRADEINTERFACEVERSION() (string, error) {
	return _Piece.Contract.UPGRADEINTERFACEVERSION(&_Piece.CallOpts)
}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_Piece *PieceCallerSession) UPGRADEINTERFACEVERSION() (string, error) {
	return _Piece.Contract.UPGRADEINTERFACEVERSION(&_Piece.CallOpts)
}

// VERSION is a free data retrieval call binding the contract method 0xffa1ad74.
//
// Solidity: function VERSION() view returns(string)
func (_Piece *PieceCaller) VERSION(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Piece.contract.Call(opts, &out, "VERSION")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// VERSION is a free data retrieval call binding the contract method 0xffa1ad74.
//
// Solidity: function VERSION() view returns(string)
func (_Piece *PieceSession) VERSION() (string, error) {
	return _Piece.Contract.VERSION(&_Piece.CallOpts)
}

// VERSION is a free data retrieval call binding the contract method 0xffa1ad74.
//
// Solidity: function VERSION() view returns(string)
func (_Piece *PieceCallerSession) VERSION() (string, error) {
	return _Piece.Contract.VERSION(&_Piece.CallOpts)
}

// CheckSReplica is a free data retrieval call binding the contract method 0x894d35f8.
//
// Solidity: function checkSReplica(address _a, uint64 _ri, bytes _com) view returns(uint64)
func (_Piece *PieceCaller) CheckSReplica(opts *bind.CallOpts, _a common.Address, _ri uint64, _com []byte) (uint64, error) {
	var out []interface{}
	err := _Piece.contract.Call(opts, &out, "checkSReplica", _a, _ri, _com)

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// CheckSReplica is a free data retrieval call binding the contract method 0x894d35f8.
//
// Solidity: function checkSReplica(address _a, uint64 _ri, bytes _com) view returns(uint64)
func (_Piece *PieceSession) CheckSReplica(_a common.Address, _ri uint64, _com []byte) (uint64, error) {
	return _Piece.Contract.CheckSReplica(&_Piece.CallOpts, _a, _ri, _com)
}

// CheckSReplica is a free data retrieval call binding the contract method 0x894d35f8.
//
// Solidity: function checkSReplica(address _a, uint64 _ri, bytes _com) view returns(uint64)
func (_Piece *PieceCallerSession) CheckSReplica(_a common.Address, _ri uint64, _com []byte) (uint64, error) {
	return _Piece.Contract.CheckSReplica(&_Piece.CallOpts, _a, _ri, _com)
}

// Delay is a free data retrieval call binding the contract method 0x6a42b8f8.
//
// Solidity: function delay() view returns(uint64)
func (_Piece *PieceCaller) Delay(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _Piece.contract.Call(opts, &out, "delay")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// Delay is a free data retrieval call binding the contract method 0x6a42b8f8.
//
// Solidity: function delay() view returns(uint64)
func (_Piece *PieceSession) Delay() (uint64, error) {
	return _Piece.Contract.Delay(&_Piece.CallOpts)
}

// Delay is a free data retrieval call binding the contract method 0x6a42b8f8.
//
// Solidity: function delay() view returns(uint64)
func (_Piece *PieceCallerSession) Delay() (uint64, error) {
	return _Piece.Contract.Delay(&_Piece.CallOpts)
}

// Epoch is a free data retrieval call binding the contract method 0x900cf0cf.
//
// Solidity: function epoch() view returns(address)
func (_Piece *PieceCaller) Epoch(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Piece.contract.Call(opts, &out, "epoch")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Epoch is a free data retrieval call binding the contract method 0x900cf0cf.
//
// Solidity: function epoch() view returns(address)
func (_Piece *PieceSession) Epoch() (common.Address, error) {
	return _Piece.Contract.Epoch(&_Piece.CallOpts)
}

// Epoch is a free data retrieval call binding the contract method 0x900cf0cf.
//
// Solidity: function epoch() view returns(address)
func (_Piece *PieceCallerSession) Epoch() (common.Address, error) {
	return _Piece.Contract.Epoch(&_Piece.CallOpts)
}

// GetPIndex is a free data retrieval call binding the contract method 0x1ce85e7c.
//
// Solidity: function getPIndex(bytes _pn) view returns(uint64)
func (_Piece *PieceCaller) GetPIndex(opts *bind.CallOpts, _pn []byte) (uint64, error) {
	var out []interface{}
	err := _Piece.contract.Call(opts, &out, "getPIndex", _pn)

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// GetPIndex is a free data retrieval call binding the contract method 0x1ce85e7c.
//
// Solidity: function getPIndex(bytes _pn) view returns(uint64)
func (_Piece *PieceSession) GetPIndex(_pn []byte) (uint64, error) {
	return _Piece.Contract.GetPIndex(&_Piece.CallOpts, _pn)
}

// GetPIndex is a free data retrieval call binding the contract method 0x1ce85e7c.
//
// Solidity: function getPIndex(bytes _pn) view returns(uint64)
func (_Piece *PieceCallerSession) GetPIndex(_pn []byte) (uint64, error) {
	return _Piece.Contract.GetPIndex(&_Piece.CallOpts, _pn)
}

// GetPRI is a free data retrieval call binding the contract method 0xee91365b.
//
// Solidity: function getPRI(uint64 _pi, uint8 _pri) view returns(uint64, address)
func (_Piece *PieceCaller) GetPRI(opts *bind.CallOpts, _pi uint64, _pri uint8) (uint64, common.Address, error) {
	var out []interface{}
	err := _Piece.contract.Call(opts, &out, "getPRI", _pi, _pri)

	if err != nil {
		return *new(uint64), *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)
	out1 := *abi.ConvertType(out[1], new(common.Address)).(*common.Address)

	return out0, out1, err

}

// GetPRI is a free data retrieval call binding the contract method 0xee91365b.
//
// Solidity: function getPRI(uint64 _pi, uint8 _pri) view returns(uint64, address)
func (_Piece *PieceSession) GetPRI(_pi uint64, _pri uint8) (uint64, common.Address, error) {
	return _Piece.Contract.GetPRI(&_Piece.CallOpts, _pi, _pri)
}

// GetPRI is a free data retrieval call binding the contract method 0xee91365b.
//
// Solidity: function getPRI(uint64 _pi, uint8 _pri) view returns(uint64, address)
func (_Piece *PieceCallerSession) GetPRI(_pi uint64, _pri uint8) (uint64, common.Address, error) {
	return _Piece.Contract.GetPRI(&_Piece.CallOpts, _pi, _pri)
}

// GetPRInfo is a free data retrieval call binding the contract method 0xb650b504.
//
// Solidity: function getPRInfo(uint64 _pi, uint64 _ri) view returns(uint8, uint8, uint64, address, bytes32)
func (_Piece *PieceCaller) GetPRInfo(opts *bind.CallOpts, _pi uint64, _ri uint64) (uint8, uint8, uint64, common.Address, [32]byte, error) {
	var out []interface{}
	err := _Piece.contract.Call(opts, &out, "getPRInfo", _pi, _ri)

	if err != nil {
		return *new(uint8), *new(uint8), *new(uint64), *new(common.Address), *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)
	out1 := *abi.ConvertType(out[1], new(uint8)).(*uint8)
	out2 := *abi.ConvertType(out[2], new(uint64)).(*uint64)
	out3 := *abi.ConvertType(out[3], new(common.Address)).(*common.Address)
	out4 := *abi.ConvertType(out[4], new([32]byte)).(*[32]byte)

	return out0, out1, out2, out3, out4, err

}

// GetPRInfo is a free data retrieval call binding the contract method 0xb650b504.
//
// Solidity: function getPRInfo(uint64 _pi, uint64 _ri) view returns(uint8, uint8, uint64, address, bytes32)
func (_Piece *PieceSession) GetPRInfo(_pi uint64, _ri uint64) (uint8, uint8, uint64, common.Address, [32]byte, error) {
	return _Piece.Contract.GetPRInfo(&_Piece.CallOpts, _pi, _ri)
}

// GetPRInfo is a free data retrieval call binding the contract method 0xb650b504.
//
// Solidity: function getPRInfo(uint64 _pi, uint64 _ri) view returns(uint8, uint8, uint64, address, bytes32)
func (_Piece *PieceCallerSession) GetPRInfo(_pi uint64, _ri uint64) (uint8, uint8, uint64, common.Address, [32]byte, error) {
	return _Piece.Contract.GetPRInfo(&_Piece.CallOpts, _pi, _ri)
}

// GetPiece is a free data retrieval call binding the contract method 0x7a87f0c4.
//
// Solidity: function getPiece(uint64 _pi) view returns((uint8,uint8,uint64,uint64,uint64,address,address,uint256,uint256) _pinfo)
func (_Piece *PieceCaller) GetPiece(opts *bind.CallOpts, _pi uint64) (IPiecePieceInfo, error) {
	var out []interface{}
	err := _Piece.contract.Call(opts, &out, "getPiece", _pi)

	if err != nil {
		return *new(IPiecePieceInfo), err
	}

	out0 := *abi.ConvertType(out[0], new(IPiecePieceInfo)).(*IPiecePieceInfo)

	return out0, err

}

// GetPiece is a free data retrieval call binding the contract method 0x7a87f0c4.
//
// Solidity: function getPiece(uint64 _pi) view returns((uint8,uint8,uint64,uint64,uint64,address,address,uint256,uint256) _pinfo)
func (_Piece *PieceSession) GetPiece(_pi uint64) (IPiecePieceInfo, error) {
	return _Piece.Contract.GetPiece(&_Piece.CallOpts, _pi)
}

// GetPiece is a free data retrieval call binding the contract method 0x7a87f0c4.
//
// Solidity: function getPiece(uint64 _pi) view returns((uint8,uint8,uint64,uint64,uint64,address,address,uint256,uint256) _pinfo)
func (_Piece *PieceCallerSession) GetPiece(_pi uint64) (IPiecePieceInfo, error) {
	return _Piece.Contract.GetPiece(&_Piece.CallOpts, _pi)
}

// GetRIndex is a free data retrieval call binding the contract method 0xba8b2618.
//
// Solidity: function getRIndex(bytes _rn) view returns(uint64)
func (_Piece *PieceCaller) GetRIndex(opts *bind.CallOpts, _rn []byte) (uint64, error) {
	var out []interface{}
	err := _Piece.contract.Call(opts, &out, "getRIndex", _rn)

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// GetRIndex is a free data retrieval call binding the contract method 0xba8b2618.
//
// Solidity: function getRIndex(bytes _rn) view returns(uint64)
func (_Piece *PieceSession) GetRIndex(_rn []byte) (uint64, error) {
	return _Piece.Contract.GetRIndex(&_Piece.CallOpts, _rn)
}

// GetRIndex is a free data retrieval call binding the contract method 0xba8b2618.
//
// Solidity: function getRIndex(bytes _rn) view returns(uint64)
func (_Piece *PieceCallerSession) GetRIndex(_rn []byte) (uint64, error) {
	return _Piece.Contract.GetRIndex(&_Piece.CallOpts, _rn)
}

// GetRS is a free data retrieval call binding the contract method 0x6aadfac7.
//
// Solidity: function getRS(uint64 _pi) view returns(uint8, uint8, uint64)
func (_Piece *PieceCaller) GetRS(opts *bind.CallOpts, _pi uint64) (uint8, uint8, uint64, error) {
	var out []interface{}
	err := _Piece.contract.Call(opts, &out, "getRS", _pi)

	if err != nil {
		return *new(uint8), *new(uint8), *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)
	out1 := *abi.ConvertType(out[1], new(uint8)).(*uint8)
	out2 := *abi.ConvertType(out[2], new(uint64)).(*uint64)

	return out0, out1, out2, err

}

// GetRS is a free data retrieval call binding the contract method 0x6aadfac7.
//
// Solidity: function getRS(uint64 _pi) view returns(uint8, uint8, uint64)
func (_Piece *PieceSession) GetRS(_pi uint64) (uint8, uint8, uint64, error) {
	return _Piece.Contract.GetRS(&_Piece.CallOpts, _pi)
}

// GetRS is a free data retrieval call binding the contract method 0x6aadfac7.
//
// Solidity: function getRS(uint64 _pi) view returns(uint8, uint8, uint64)
func (_Piece *PieceCallerSession) GetRS(_pi uint64) (uint8, uint8, uint64, error) {
	return _Piece.Contract.GetRS(&_Piece.CallOpts, _pi)
}

// GetReplica is a free data retrieval call binding the contract method 0x9108544c.
//
// Solidity: function getReplica(uint64 _r) view returns((uint64,address,bytes32) _ri)
func (_Piece *PieceCaller) GetReplica(opts *bind.CallOpts, _r uint64) (IPieceReplicaInfo, error) {
	var out []interface{}
	err := _Piece.contract.Call(opts, &out, "getReplica", _r)

	if err != nil {
		return *new(IPieceReplicaInfo), err
	}

	out0 := *abi.ConvertType(out[0], new(IPieceReplicaInfo)).(*IPieceReplicaInfo)

	return out0, err

}

// GetReplica is a free data retrieval call binding the contract method 0x9108544c.
//
// Solidity: function getReplica(uint64 _r) view returns((uint64,address,bytes32) _ri)
func (_Piece *PieceSession) GetReplica(_r uint64) (IPieceReplicaInfo, error) {
	return _Piece.Contract.GetReplica(&_Piece.CallOpts, _r)
}

// GetReplica is a free data retrieval call binding the contract method 0x9108544c.
//
// Solidity: function getReplica(uint64 _r) view returns((uint64,address,bytes32) _ri)
func (_Piece *PieceCallerSession) GetReplica(_r uint64) (IPieceReplicaInfo, error) {
	return _Piece.Contract.GetReplica(&_Piece.CallOpts, _r)
}

// GetRevenue is a free data retrieval call binding the contract method 0x85b67095.
//
// Solidity: function getRevenue(address _a) view returns(uint256)
func (_Piece *PieceCaller) GetRevenue(opts *bind.CallOpts, _a common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Piece.contract.Call(opts, &out, "getRevenue", _a)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetRevenue is a free data retrieval call binding the contract method 0x85b67095.
//
// Solidity: function getRevenue(address _a) view returns(uint256)
func (_Piece *PieceSession) GetRevenue(_a common.Address) (*big.Int, error) {
	return _Piece.Contract.GetRevenue(&_Piece.CallOpts, _a)
}

// GetRevenue is a free data retrieval call binding the contract method 0x85b67095.
//
// Solidity: function getRevenue(address _a) view returns(uint256)
func (_Piece *PieceCallerSession) GetRevenue(_a common.Address) (*big.Int, error) {
	return _Piece.Contract.GetRevenue(&_Piece.CallOpts, _a)
}

// GetSRAt is a free data retrieval call binding the contract method 0xbcbfd56b.
//
// Solidity: function getSRAt(address _a, uint64 _ri) view returns(uint64)
func (_Piece *PieceCaller) GetSRAt(opts *bind.CallOpts, _a common.Address, _ri uint64) (uint64, error) {
	var out []interface{}
	err := _Piece.contract.Call(opts, &out, "getSRAt", _a, _ri)

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// GetSRAt is a free data retrieval call binding the contract method 0xbcbfd56b.
//
// Solidity: function getSRAt(address _a, uint64 _ri) view returns(uint64)
func (_Piece *PieceSession) GetSRAt(_a common.Address, _ri uint64) (uint64, error) {
	return _Piece.Contract.GetSRAt(&_Piece.CallOpts, _a, _ri)
}

// GetSRAt is a free data retrieval call binding the contract method 0xbcbfd56b.
//
// Solidity: function getSRAt(address _a, uint64 _ri) view returns(uint64)
func (_Piece *PieceCallerSession) GetSRAt(_a common.Address, _ri uint64) (uint64, error) {
	return _Piece.Contract.GetSRAt(&_Piece.CallOpts, _a, _ri)
}

// GetSStat is a free data retrieval call binding the contract method 0xc3aa3804.
//
// Solidity: function getSStat(address _a, uint64 _e) view returns((uint64,uint256) _si)
func (_Piece *PieceCaller) GetSStat(opts *bind.CallOpts, _a common.Address, _e uint64) (IPieceStoreStat, error) {
	var out []interface{}
	err := _Piece.contract.Call(opts, &out, "getSStat", _a, _e)

	if err != nil {
		return *new(IPieceStoreStat), err
	}

	out0 := *abi.ConvertType(out[0], new(IPieceStoreStat)).(*IPieceStoreStat)

	return out0, err

}

// GetSStat is a free data retrieval call binding the contract method 0xc3aa3804.
//
// Solidity: function getSStat(address _a, uint64 _e) view returns((uint64,uint256) _si)
func (_Piece *PieceSession) GetSStat(_a common.Address, _e uint64) (IPieceStoreStat, error) {
	return _Piece.Contract.GetSStat(&_Piece.CallOpts, _a, _e)
}

// GetSStat is a free data retrieval call binding the contract method 0xc3aa3804.
//
// Solidity: function getSStat(address _a, uint64 _e) view returns((uint64,uint256) _si)
func (_Piece *PieceCallerSession) GetSStat(_a common.Address, _e uint64) (IPieceStoreStat, error) {
	return _Piece.Contract.GetSStat(&_Piece.CallOpts, _a, _e)
}

// GetStore is a free data retrieval call binding the contract method 0x4b4ffccd.
//
// Solidity: function getStore(address _a) view returns((uint64,uint64,uint256,uint256) _si)
func (_Piece *PieceCaller) GetStore(opts *bind.CallOpts, _a common.Address) (IPieceStoreInfo, error) {
	var out []interface{}
	err := _Piece.contract.Call(opts, &out, "getStore", _a)

	if err != nil {
		return *new(IPieceStoreInfo), err
	}

	out0 := *abi.ConvertType(out[0], new(IPieceStoreInfo)).(*IPieceStoreInfo)

	return out0, err

}

// GetStore is a free data retrieval call binding the contract method 0x4b4ffccd.
//
// Solidity: function getStore(address _a) view returns((uint64,uint64,uint256,uint256) _si)
func (_Piece *PieceSession) GetStore(_a common.Address) (IPieceStoreInfo, error) {
	return _Piece.Contract.GetStore(&_Piece.CallOpts, _a)
}

// GetStore is a free data retrieval call binding the contract method 0x4b4ffccd.
//
// Solidity: function getStore(address _a) view returns((uint64,uint64,uint256,uint256) _si)
func (_Piece *PieceCallerSession) GetStore(_a common.Address) (IPieceStoreInfo, error) {
	return _Piece.Contract.GetStore(&_Piece.CallOpts, _a)
}

// GetStoreCount is a free data retrieval call binding the contract method 0x50979d7e.
//
// Solidity: function getStoreCount(address _a, uint64 _e) view returns(uint64)
func (_Piece *PieceCaller) GetStoreCount(opts *bind.CallOpts, _a common.Address, _e uint64) (uint64, error) {
	var out []interface{}
	err := _Piece.contract.Call(opts, &out, "getStoreCount", _a, _e)

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// GetStoreCount is a free data retrieval call binding the contract method 0x50979d7e.
//
// Solidity: function getStoreCount(address _a, uint64 _e) view returns(uint64)
func (_Piece *PieceSession) GetStoreCount(_a common.Address, _e uint64) (uint64, error) {
	return _Piece.Contract.GetStoreCount(&_Piece.CallOpts, _a, _e)
}

// GetStoreCount is a free data retrieval call binding the contract method 0x50979d7e.
//
// Solidity: function getStoreCount(address _a, uint64 _e) view returns(uint64)
func (_Piece *PieceCallerSession) GetStoreCount(_a common.Address, _e uint64) (uint64, error) {
	return _Piece.Contract.GetStoreCount(&_Piece.CallOpts, _a, _e)
}

// GetStoreSalary is a free data retrieval call binding the contract method 0x18198ad7.
//
// Solidity: function getStoreSalary(address _a, uint64 _e) view returns(uint256)
func (_Piece *PieceCaller) GetStoreSalary(opts *bind.CallOpts, _a common.Address, _e uint64) (*big.Int, error) {
	var out []interface{}
	err := _Piece.contract.Call(opts, &out, "getStoreSalary", _a, _e)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetStoreSalary is a free data retrieval call binding the contract method 0x18198ad7.
//
// Solidity: function getStoreSalary(address _a, uint64 _e) view returns(uint256)
func (_Piece *PieceSession) GetStoreSalary(_a common.Address, _e uint64) (*big.Int, error) {
	return _Piece.Contract.GetStoreSalary(&_Piece.CallOpts, _a, _e)
}

// GetStoreSalary is a free data retrieval call binding the contract method 0x18198ad7.
//
// Solidity: function getStoreSalary(address _a, uint64 _e) view returns(uint256)
func (_Piece *PieceCallerSession) GetStoreSalary(_a common.Address, _e uint64) (*big.Int, error) {
	return _Piece.Contract.GetStoreSalary(&_Piece.CallOpts, _a, _e)
}

// MaxSize is a free data retrieval call binding the contract method 0x2565b159.
//
// Solidity: function maxSize() view returns(uint64)
func (_Piece *PieceCaller) MaxSize(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _Piece.contract.Call(opts, &out, "maxSize")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// MaxSize is a free data retrieval call binding the contract method 0x2565b159.
//
// Solidity: function maxSize() view returns(uint64)
func (_Piece *PieceSession) MaxSize() (uint64, error) {
	return _Piece.Contract.MaxSize(&_Piece.CallOpts)
}

// MaxSize is a free data retrieval call binding the contract method 0x2565b159.
//
// Solidity: function maxSize() view returns(uint64)
func (_Piece *PieceCallerSession) MaxSize() (uint64, error) {
	return _Piece.Contract.MaxSize(&_Piece.CallOpts)
}

// MaxStore is a free data retrieval call binding the contract method 0xfca5be7a.
//
// Solidity: function maxStore() view returns(uint64)
func (_Piece *PieceCaller) MaxStore(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _Piece.contract.Call(opts, &out, "maxStore")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// MaxStore is a free data retrieval call binding the contract method 0xfca5be7a.
//
// Solidity: function maxStore() view returns(uint64)
func (_Piece *PieceSession) MaxStore() (uint64, error) {
	return _Piece.Contract.MaxStore(&_Piece.CallOpts)
}

// MaxStore is a free data retrieval call binding the contract method 0xfca5be7a.
//
// Solidity: function maxStore() view returns(uint64)
func (_Piece *PieceCallerSession) MaxStore() (uint64, error) {
	return _Piece.Contract.MaxStore(&_Piece.CallOpts)
}

// MinPrice is a free data retrieval call binding the contract method 0xe45be8eb.
//
// Solidity: function minPrice() view returns(uint256)
func (_Piece *PieceCaller) MinPrice(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Piece.contract.Call(opts, &out, "minPrice")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MinPrice is a free data retrieval call binding the contract method 0xe45be8eb.
//
// Solidity: function minPrice() view returns(uint256)
func (_Piece *PieceSession) MinPrice() (*big.Int, error) {
	return _Piece.Contract.MinPrice(&_Piece.CallOpts)
}

// MinPrice is a free data retrieval call binding the contract method 0xe45be8eb.
//
// Solidity: function minPrice() view returns(uint256)
func (_Piece *PieceCallerSession) MinPrice() (*big.Int, error) {
	return _Piece.Contract.MinPrice(&_Piece.CallOpts)
}

// MinStore is a free data retrieval call binding the contract method 0x88795125.
//
// Solidity: function minStore() view returns(uint64)
func (_Piece *PieceCaller) MinStore(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _Piece.contract.Call(opts, &out, "minStore")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// MinStore is a free data retrieval call binding the contract method 0x88795125.
//
// Solidity: function minStore() view returns(uint64)
func (_Piece *PieceSession) MinStore() (uint64, error) {
	return _Piece.Contract.MinStore(&_Piece.CallOpts)
}

// MinStore is a free data retrieval call binding the contract method 0x88795125.
//
// Solidity: function minStore() view returns(uint64)
func (_Piece *PieceCallerSession) MinStore() (uint64, error) {
	return _Piece.Contract.MinStore(&_Piece.CallOpts)
}

// Node is a free data retrieval call binding the contract method 0xd70754ec.
//
// Solidity: function node() view returns(address)
func (_Piece *PieceCaller) Node(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Piece.contract.Call(opts, &out, "node")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Node is a free data retrieval call binding the contract method 0xd70754ec.
//
// Solidity: function node() view returns(address)
func (_Piece *PieceSession) Node() (common.Address, error) {
	return _Piece.Contract.Node(&_Piece.CallOpts)
}

// Node is a free data retrieval call binding the contract method 0xd70754ec.
//
// Solidity: function node() view returns(address)
func (_Piece *PieceCallerSession) Node() (common.Address, error) {
	return _Piece.Contract.Node(&_Piece.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Piece *PieceCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Piece.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Piece *PieceSession) Owner() (common.Address, error) {
	return _Piece.Contract.Owner(&_Piece.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Piece *PieceCallerSession) Owner() (common.Address, error) {
	return _Piece.Contract.Owner(&_Piece.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_Piece *PieceCaller) ProxiableUUID(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Piece.contract.Call(opts, &out, "proxiableUUID")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_Piece *PieceSession) ProxiableUUID() ([32]byte, error) {
	return _Piece.Contract.ProxiableUUID(&_Piece.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_Piece *PieceCallerSession) ProxiableUUID() ([32]byte, error) {
	return _Piece.Contract.ProxiableUUID(&_Piece.CallOpts)
}

// StreamPrice is a free data retrieval call binding the contract method 0x75011aba.
//
// Solidity: function streamPrice() view returns(uint256)
func (_Piece *PieceCaller) StreamPrice(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Piece.contract.Call(opts, &out, "streamPrice")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// StreamPrice is a free data retrieval call binding the contract method 0x75011aba.
//
// Solidity: function streamPrice() view returns(uint256)
func (_Piece *PieceSession) StreamPrice() (*big.Int, error) {
	return _Piece.Contract.StreamPrice(&_Piece.CallOpts)
}

// StreamPrice is a free data retrieval call binding the contract method 0x75011aba.
//
// Solidity: function streamPrice() view returns(uint256)
func (_Piece *PieceCallerSession) StreamPrice() (*big.Int, error) {
	return _Piece.Contract.StreamPrice(&_Piece.CallOpts)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_Piece *PieceCaller) Token(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Piece.contract.Call(opts, &out, "token")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_Piece *PieceSession) Token() (common.Address, error) {
	return _Piece.Contract.Token(&_Piece.CallOpts)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_Piece *PieceCallerSession) Token() (common.Address, error) {
	return _Piece.Contract.Token(&_Piece.CallOpts)
}

// AddPiece is a paid mutator transaction binding the contract method 0x8ef0e895.
//
// Solidity: function addPiece(bytes _pn, uint256 _price, uint64 _size, uint64 _expire, uint8 rsn, uint8 rsk, address _s) returns()
func (_Piece *PieceTransactor) AddPiece(opts *bind.TransactOpts, _pn []byte, _price *big.Int, _size uint64, _expire uint64, rsn uint8, rsk uint8, _s common.Address) (*types.Transaction, error) {
	return _Piece.contract.Transact(opts, "addPiece", _pn, _price, _size, _expire, rsn, rsk, _s)
}

// AddPiece is a paid mutator transaction binding the contract method 0x8ef0e895.
//
// Solidity: function addPiece(bytes _pn, uint256 _price, uint64 _size, uint64 _expire, uint8 rsn, uint8 rsk, address _s) returns()
func (_Piece *PieceSession) AddPiece(_pn []byte, _price *big.Int, _size uint64, _expire uint64, rsn uint8, rsk uint8, _s common.Address) (*types.Transaction, error) {
	return _Piece.Contract.AddPiece(&_Piece.TransactOpts, _pn, _price, _size, _expire, rsn, rsk, _s)
}

// AddPiece is a paid mutator transaction binding the contract method 0x8ef0e895.
//
// Solidity: function addPiece(bytes _pn, uint256 _price, uint64 _size, uint64 _expire, uint8 rsn, uint8 rsk, address _s) returns()
func (_Piece *PieceTransactorSession) AddPiece(_pn []byte, _price *big.Int, _size uint64, _expire uint64, rsn uint8, rsk uint8, _s common.Address) (*types.Transaction, error) {
	return _Piece.Contract.AddPiece(&_Piece.TransactOpts, _pn, _price, _size, _expire, rsn, rsk, _s)
}

// AddReplica is a paid mutator transaction binding the contract method 0x194cbaf5.
//
// Solidity: function addReplica(bytes _rn, uint64 _pi, uint8 _pri, bytes _proof) returns()
func (_Piece *PieceTransactor) AddReplica(opts *bind.TransactOpts, _rn []byte, _pi uint64, _pri uint8, _proof []byte) (*types.Transaction, error) {
	return _Piece.contract.Transact(opts, "addReplica", _rn, _pi, _pri, _proof)
}

// AddReplica is a paid mutator transaction binding the contract method 0x194cbaf5.
//
// Solidity: function addReplica(bytes _rn, uint64 _pi, uint8 _pri, bytes _proof) returns()
func (_Piece *PieceSession) AddReplica(_rn []byte, _pi uint64, _pri uint8, _proof []byte) (*types.Transaction, error) {
	return _Piece.Contract.AddReplica(&_Piece.TransactOpts, _rn, _pi, _pri, _proof)
}

// AddReplica is a paid mutator transaction binding the contract method 0x194cbaf5.
//
// Solidity: function addReplica(bytes _rn, uint64 _pi, uint8 _pri, bytes _proof) returns()
func (_Piece *PieceTransactorSession) AddReplica(_rn []byte, _pi uint64, _pri uint8, _proof []byte) (*types.Transaction, error) {
	return _Piece.Contract.AddReplica(&_Piece.TransactOpts, _rn, _pi, _pri, _proof)
}

// CheckStore is a paid mutator transaction binding the contract method 0x62b98032.
//
// Solidity: function checkStore(address _a) returns()
func (_Piece *PieceTransactor) CheckStore(opts *bind.TransactOpts, _a common.Address) (*types.Transaction, error) {
	return _Piece.contract.Transact(opts, "checkStore", _a)
}

// CheckStore is a paid mutator transaction binding the contract method 0x62b98032.
//
// Solidity: function checkStore(address _a) returns()
func (_Piece *PieceSession) CheckStore(_a common.Address) (*types.Transaction, error) {
	return _Piece.Contract.CheckStore(&_Piece.TransactOpts, _a)
}

// CheckStore is a paid mutator transaction binding the contract method 0x62b98032.
//
// Solidity: function checkStore(address _a) returns()
func (_Piece *PieceTransactorSession) CheckStore(_a common.Address) (*types.Transaction, error) {
	return _Piece.Contract.CheckStore(&_Piece.TransactOpts, _a)
}

// Initialize is a paid mutator transaction binding the contract method 0x7cf1ee91.
//
// Solidity: function initialize(address _token, address _epoch, address _node, (uint64,uint64,uint64,uint64,uint256,uint256) params, address initialOwner) returns()
func (_Piece *PieceTransactor) Initialize(opts *bind.TransactOpts, _token common.Address, _epoch common.Address, _node common.Address, params IPieceInitParams, initialOwner common.Address) (*types.Transaction, error) {
	return _Piece.contract.Transact(opts, "initialize", _token, _epoch, _node, params, initialOwner)
}

// Initialize is a paid mutator transaction binding the contract method 0x7cf1ee91.
//
// Solidity: function initialize(address _token, address _epoch, address _node, (uint64,uint64,uint64,uint64,uint256,uint256) params, address initialOwner) returns()
func (_Piece *PieceSession) Initialize(_token common.Address, _epoch common.Address, _node common.Address, params IPieceInitParams, initialOwner common.Address) (*types.Transaction, error) {
	return _Piece.Contract.Initialize(&_Piece.TransactOpts, _token, _epoch, _node, params, initialOwner)
}

// Initialize is a paid mutator transaction binding the contract method 0x7cf1ee91.
//
// Solidity: function initialize(address _token, address _epoch, address _node, (uint64,uint64,uint64,uint64,uint256,uint256) params, address initialOwner) returns()
func (_Piece *PieceTransactorSession) Initialize(_token common.Address, _epoch common.Address, _node common.Address, params IPieceInitParams, initialOwner common.Address) (*types.Transaction, error) {
	return _Piece.Contract.Initialize(&_Piece.TransactOpts, _token, _epoch, _node, params, initialOwner)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Piece *PieceTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Piece.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Piece *PieceSession) RenounceOwnership() (*types.Transaction, error) {
	return _Piece.Contract.RenounceOwnership(&_Piece.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Piece *PieceTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Piece.Contract.RenounceOwnership(&_Piece.TransactOpts)
}

// Retake is a paid mutator transaction binding the contract method 0x123f8459.
//
// Solidity: function retake(uint64 _pi) returns()
func (_Piece *PieceTransactor) Retake(opts *bind.TransactOpts, _pi uint64) (*types.Transaction, error) {
	return _Piece.contract.Transact(opts, "retake", _pi)
}

// Retake is a paid mutator transaction binding the contract method 0x123f8459.
//
// Solidity: function retake(uint64 _pi) returns()
func (_Piece *PieceSession) Retake(_pi uint64) (*types.Transaction, error) {
	return _Piece.Contract.Retake(&_Piece.TransactOpts, _pi)
}

// Retake is a paid mutator transaction binding the contract method 0x123f8459.
//
// Solidity: function retake(uint64 _pi) returns()
func (_Piece *PieceTransactorSession) Retake(_pi uint64) (*types.Transaction, error) {
	return _Piece.Contract.Retake(&_Piece.TransactOpts, _pi)
}

// Settle is a paid mutator transaction binding the contract method 0x8df82800.
//
// Solidity: function settle(uint256 _m) returns()
func (_Piece *PieceTransactor) Settle(opts *bind.TransactOpts, _m *big.Int) (*types.Transaction, error) {
	return _Piece.contract.Transact(opts, "settle", _m)
}

// Settle is a paid mutator transaction binding the contract method 0x8df82800.
//
// Solidity: function settle(uint256 _m) returns()
func (_Piece *PieceSession) Settle(_m *big.Int) (*types.Transaction, error) {
	return _Piece.Contract.Settle(&_Piece.TransactOpts, _m)
}

// Settle is a paid mutator transaction binding the contract method 0x8df82800.
//
// Solidity: function settle(uint256 _m) returns()
func (_Piece *PieceTransactorSession) Settle(_m *big.Int) (*types.Transaction, error) {
	return _Piece.Contract.Settle(&_Piece.TransactOpts, _m)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Piece *PieceTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Piece.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Piece *PieceSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Piece.Contract.TransferOwnership(&_Piece.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Piece *PieceTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Piece.Contract.TransferOwnership(&_Piece.TransactOpts, newOwner)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_Piece *PieceTransactor) UpgradeToAndCall(opts *bind.TransactOpts, newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _Piece.contract.Transact(opts, "upgradeToAndCall", newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_Piece *PieceSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _Piece.Contract.UpgradeToAndCall(&_Piece.TransactOpts, newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_Piece *PieceTransactorSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _Piece.Contract.UpgradeToAndCall(&_Piece.TransactOpts, newImplementation, data)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 _m) returns()
func (_Piece *PieceTransactor) Withdraw(opts *bind.TransactOpts, _m *big.Int) (*types.Transaction, error) {
	return _Piece.contract.Transact(opts, "withdraw", _m)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 _m) returns()
func (_Piece *PieceSession) Withdraw(_m *big.Int) (*types.Transaction, error) {
	return _Piece.Contract.Withdraw(&_Piece.TransactOpts, _m)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 _m) returns()
func (_Piece *PieceTransactorSession) Withdraw(_m *big.Int) (*types.Transaction, error) {
	return _Piece.Contract.Withdraw(&_Piece.TransactOpts, _m)
}

// PieceAddPieceIterator is returned from FilterAddPiece and is used to iterate over the raw logs and unpacked data for AddPiece events raised by the Piece contract.
type PieceAddPieceIterator struct {
	Event *PieceAddPiece // Event containing the contract specifics and raw log

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
func (it *PieceAddPieceIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PieceAddPiece)
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
		it.Event = new(PieceAddPiece)
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
func (it *PieceAddPieceIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PieceAddPieceIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PieceAddPiece represents a AddPiece event raised by the Piece contract.
type PieceAddPiece struct {
	A   common.Address
	Pi  uint64
	E   uint64
	Raw types.Log // Blockchain specific contextual infos
}

// FilterAddPiece is a free log retrieval operation binding the contract event 0x09ff0fb5f488978fdffbf3223e63e914f2a9b0c228844b8e2caedad8a85d220c.
//
// Solidity: event AddPiece(address indexed _a, uint64 _pi, uint64 _e)
func (_Piece *PieceFilterer) FilterAddPiece(opts *bind.FilterOpts, _a []common.Address) (*PieceAddPieceIterator, error) {

	var _aRule []interface{}
	for _, _aItem := range _a {
		_aRule = append(_aRule, _aItem)
	}

	logs, sub, err := _Piece.contract.FilterLogs(opts, "AddPiece", _aRule)
	if err != nil {
		return nil, err
	}
	return &PieceAddPieceIterator{contract: _Piece.contract, event: "AddPiece", logs: logs, sub: sub}, nil
}

// WatchAddPiece is a free log subscription operation binding the contract event 0x09ff0fb5f488978fdffbf3223e63e914f2a9b0c228844b8e2caedad8a85d220c.
//
// Solidity: event AddPiece(address indexed _a, uint64 _pi, uint64 _e)
func (_Piece *PieceFilterer) WatchAddPiece(opts *bind.WatchOpts, sink chan<- *PieceAddPiece, _a []common.Address) (event.Subscription, error) {

	var _aRule []interface{}
	for _, _aItem := range _a {
		_aRule = append(_aRule, _aItem)
	}

	logs, sub, err := _Piece.contract.WatchLogs(opts, "AddPiece", _aRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PieceAddPiece)
				if err := _Piece.contract.UnpackLog(event, "AddPiece", log); err != nil {
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
func (_Piece *PieceFilterer) ParseAddPiece(log types.Log) (*PieceAddPiece, error) {
	event := new(PieceAddPiece)
	if err := _Piece.contract.UnpackLog(event, "AddPiece", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PieceAddReplicaIterator is returned from FilterAddReplica and is used to iterate over the raw logs and unpacked data for AddReplica events raised by the Piece contract.
type PieceAddReplicaIterator struct {
	Event *PieceAddReplica // Event containing the contract specifics and raw log

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
func (it *PieceAddReplicaIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PieceAddReplica)
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
		it.Event = new(PieceAddReplica)
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
func (it *PieceAddReplicaIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PieceAddReplicaIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PieceAddReplica represents a AddReplica event raised by the Piece contract.
type PieceAddReplica struct {
	A   common.Address
	Ri  uint64
	Sri uint64
	Raw types.Log // Blockchain specific contextual infos
}

// FilterAddReplica is a free log retrieval operation binding the contract event 0x4b079e585085a3c94b2865036e598391fad4359764b36da39e21897be56349ef.
//
// Solidity: event AddReplica(address indexed _a, uint64 _ri, uint64 _sri)
func (_Piece *PieceFilterer) FilterAddReplica(opts *bind.FilterOpts, _a []common.Address) (*PieceAddReplicaIterator, error) {

	var _aRule []interface{}
	for _, _aItem := range _a {
		_aRule = append(_aRule, _aItem)
	}

	logs, sub, err := _Piece.contract.FilterLogs(opts, "AddReplica", _aRule)
	if err != nil {
		return nil, err
	}
	return &PieceAddReplicaIterator{contract: _Piece.contract, event: "AddReplica", logs: logs, sub: sub}, nil
}

// WatchAddReplica is a free log subscription operation binding the contract event 0x4b079e585085a3c94b2865036e598391fad4359764b36da39e21897be56349ef.
//
// Solidity: event AddReplica(address indexed _a, uint64 _ri, uint64 _sri)
func (_Piece *PieceFilterer) WatchAddReplica(opts *bind.WatchOpts, sink chan<- *PieceAddReplica, _a []common.Address) (event.Subscription, error) {

	var _aRule []interface{}
	for _, _aItem := range _a {
		_aRule = append(_aRule, _aItem)
	}

	logs, sub, err := _Piece.contract.WatchLogs(opts, "AddReplica", _aRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PieceAddReplica)
				if err := _Piece.contract.UnpackLog(event, "AddReplica", log); err != nil {
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
func (_Piece *PieceFilterer) ParseAddReplica(log types.Log) (*PieceAddReplica, error) {
	event := new(PieceAddReplica)
	if err := _Piece.contract.UnpackLog(event, "AddReplica", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PieceInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the Piece contract.
type PieceInitializedIterator struct {
	Event *PieceInitialized // Event containing the contract specifics and raw log

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
func (it *PieceInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PieceInitialized)
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
		it.Event = new(PieceInitialized)
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
func (it *PieceInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PieceInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PieceInitialized represents a Initialized event raised by the Piece contract.
type PieceInitialized struct {
	Version uint64
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_Piece *PieceFilterer) FilterInitialized(opts *bind.FilterOpts) (*PieceInitializedIterator, error) {

	logs, sub, err := _Piece.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &PieceInitializedIterator{contract: _Piece.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_Piece *PieceFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *PieceInitialized) (event.Subscription, error) {

	logs, sub, err := _Piece.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PieceInitialized)
				if err := _Piece.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_Piece *PieceFilterer) ParseInitialized(log types.Log) (*PieceInitialized, error) {
	event := new(PieceInitialized)
	if err := _Piece.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PieceOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Piece contract.
type PieceOwnershipTransferredIterator struct {
	Event *PieceOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *PieceOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PieceOwnershipTransferred)
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
		it.Event = new(PieceOwnershipTransferred)
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
func (it *PieceOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PieceOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PieceOwnershipTransferred represents a OwnershipTransferred event raised by the Piece contract.
type PieceOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Piece *PieceFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*PieceOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Piece.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &PieceOwnershipTransferredIterator{contract: _Piece.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Piece *PieceFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *PieceOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Piece.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PieceOwnershipTransferred)
				if err := _Piece.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_Piece *PieceFilterer) ParseOwnershipTransferred(log types.Log) (*PieceOwnershipTransferred, error) {
	event := new(PieceOwnershipTransferred)
	if err := _Piece.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PieceRetakeIterator is returned from FilterRetake and is used to iterate over the raw logs and unpacked data for Retake events raised by the Piece contract.
type PieceRetakeIterator struct {
	Event *PieceRetake // Event containing the contract specifics and raw log

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
func (it *PieceRetakeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PieceRetake)
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
		it.Event = new(PieceRetake)
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
func (it *PieceRetakeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PieceRetakeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PieceRetake represents a Retake event raised by the Piece contract.
type PieceRetake struct {
	A   common.Address
	Pi  uint64
	M   *big.Int
	Raw types.Log // Blockchain specific contextual infos
}

// FilterRetake is a free log retrieval operation binding the contract event 0x44647ea87405778ba90a8ca28d2b772b1350bcd5950e5111bbc8c506a5c9632a.
//
// Solidity: event Retake(address indexed _a, uint64 _pi, uint256 _m)
func (_Piece *PieceFilterer) FilterRetake(opts *bind.FilterOpts, _a []common.Address) (*PieceRetakeIterator, error) {

	var _aRule []interface{}
	for _, _aItem := range _a {
		_aRule = append(_aRule, _aItem)
	}

	logs, sub, err := _Piece.contract.FilterLogs(opts, "Retake", _aRule)
	if err != nil {
		return nil, err
	}
	return &PieceRetakeIterator{contract: _Piece.contract, event: "Retake", logs: logs, sub: sub}, nil
}

// WatchRetake is a free log subscription operation binding the contract event 0x44647ea87405778ba90a8ca28d2b772b1350bcd5950e5111bbc8c506a5c9632a.
//
// Solidity: event Retake(address indexed _a, uint64 _pi, uint256 _m)
func (_Piece *PieceFilterer) WatchRetake(opts *bind.WatchOpts, sink chan<- *PieceRetake, _a []common.Address) (event.Subscription, error) {

	var _aRule []interface{}
	for _, _aItem := range _a {
		_aRule = append(_aRule, _aItem)
	}

	logs, sub, err := _Piece.contract.WatchLogs(opts, "Retake", _aRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PieceRetake)
				if err := _Piece.contract.UnpackLog(event, "Retake", log); err != nil {
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

// ParseRetake is a log parse operation binding the contract event 0x44647ea87405778ba90a8ca28d2b772b1350bcd5950e5111bbc8c506a5c9632a.
//
// Solidity: event Retake(address indexed _a, uint64 _pi, uint256 _m)
func (_Piece *PieceFilterer) ParseRetake(log types.Log) (*PieceRetake, error) {
	event := new(PieceRetake)
	if err := _Piece.contract.UnpackLog(event, "Retake", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PieceSettleIterator is returned from FilterSettle and is used to iterate over the raw logs and unpacked data for Settle events raised by the Piece contract.
type PieceSettleIterator struct {
	Event *PieceSettle // Event containing the contract specifics and raw log

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
func (it *PieceSettleIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PieceSettle)
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
		it.Event = new(PieceSettle)
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
func (it *PieceSettleIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PieceSettleIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PieceSettle represents a Settle event raised by the Piece contract.
type PieceSettle struct {
	A   common.Address
	M   *big.Int
	Raw types.Log // Blockchain specific contextual infos
}

// FilterSettle is a free log retrieval operation binding the contract event 0xa3788eedc10ef3ec6982384227c562c6666cf2b6af4f6a583b6d5d0f2ba0d204.
//
// Solidity: event Settle(address indexed _a, uint256 _m)
func (_Piece *PieceFilterer) FilterSettle(opts *bind.FilterOpts, _a []common.Address) (*PieceSettleIterator, error) {

	var _aRule []interface{}
	for _, _aItem := range _a {
		_aRule = append(_aRule, _aItem)
	}

	logs, sub, err := _Piece.contract.FilterLogs(opts, "Settle", _aRule)
	if err != nil {
		return nil, err
	}
	return &PieceSettleIterator{contract: _Piece.contract, event: "Settle", logs: logs, sub: sub}, nil
}

// WatchSettle is a free log subscription operation binding the contract event 0xa3788eedc10ef3ec6982384227c562c6666cf2b6af4f6a583b6d5d0f2ba0d204.
//
// Solidity: event Settle(address indexed _a, uint256 _m)
func (_Piece *PieceFilterer) WatchSettle(opts *bind.WatchOpts, sink chan<- *PieceSettle, _a []common.Address) (event.Subscription, error) {

	var _aRule []interface{}
	for _, _aItem := range _a {
		_aRule = append(_aRule, _aItem)
	}

	logs, sub, err := _Piece.contract.WatchLogs(opts, "Settle", _aRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PieceSettle)
				if err := _Piece.contract.UnpackLog(event, "Settle", log); err != nil {
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
func (_Piece *PieceFilterer) ParseSettle(log types.Log) (*PieceSettle, error) {
	event := new(PieceSettle)
	if err := _Piece.contract.UnpackLog(event, "Settle", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PieceUpgradedIterator is returned from FilterUpgraded and is used to iterate over the raw logs and unpacked data for Upgraded events raised by the Piece contract.
type PieceUpgradedIterator struct {
	Event *PieceUpgraded // Event containing the contract specifics and raw log

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
func (it *PieceUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PieceUpgraded)
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
		it.Event = new(PieceUpgraded)
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
func (it *PieceUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PieceUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PieceUpgraded represents a Upgraded event raised by the Piece contract.
type PieceUpgraded struct {
	Implementation common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterUpgraded is a free log retrieval operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_Piece *PieceFilterer) FilterUpgraded(opts *bind.FilterOpts, implementation []common.Address) (*PieceUpgradedIterator, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _Piece.contract.FilterLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return &PieceUpgradedIterator{contract: _Piece.contract, event: "Upgraded", logs: logs, sub: sub}, nil
}

// WatchUpgraded is a free log subscription operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_Piece *PieceFilterer) WatchUpgraded(opts *bind.WatchOpts, sink chan<- *PieceUpgraded, implementation []common.Address) (event.Subscription, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _Piece.contract.WatchLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PieceUpgraded)
				if err := _Piece.contract.UnpackLog(event, "Upgraded", log); err != nil {
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
func (_Piece *PieceFilterer) ParseUpgraded(log types.Log) (*PieceUpgraded, error) {
	event := new(PieceUpgraded)
	if err := _Piece.contract.UnpackLog(event, "Upgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PieceWithdrawIterator is returned from FilterWithdraw and is used to iterate over the raw logs and unpacked data for Withdraw events raised by the Piece contract.
type PieceWithdrawIterator struct {
	Event *PieceWithdraw // Event containing the contract specifics and raw log

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
func (it *PieceWithdrawIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PieceWithdraw)
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
		it.Event = new(PieceWithdraw)
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
func (it *PieceWithdrawIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PieceWithdrawIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PieceWithdraw represents a Withdraw event raised by the Piece contract.
type PieceWithdraw struct {
	A   common.Address
	M   *big.Int
	Raw types.Log // Blockchain specific contextual infos
}

// FilterWithdraw is a free log retrieval operation binding the contract event 0x884edad9ce6fa2440d8a54cc123490eb96d2768479d49ff9c7366125a9424364.
//
// Solidity: event Withdraw(address indexed _a, uint256 _m)
func (_Piece *PieceFilterer) FilterWithdraw(opts *bind.FilterOpts, _a []common.Address) (*PieceWithdrawIterator, error) {

	var _aRule []interface{}
	for _, _aItem := range _a {
		_aRule = append(_aRule, _aItem)
	}

	logs, sub, err := _Piece.contract.FilterLogs(opts, "Withdraw", _aRule)
	if err != nil {
		return nil, err
	}
	return &PieceWithdrawIterator{contract: _Piece.contract, event: "Withdraw", logs: logs, sub: sub}, nil
}

// WatchWithdraw is a free log subscription operation binding the contract event 0x884edad9ce6fa2440d8a54cc123490eb96d2768479d49ff9c7366125a9424364.
//
// Solidity: event Withdraw(address indexed _a, uint256 _m)
func (_Piece *PieceFilterer) WatchWithdraw(opts *bind.WatchOpts, sink chan<- *PieceWithdraw, _a []common.Address) (event.Subscription, error) {

	var _aRule []interface{}
	for _, _aItem := range _a {
		_aRule = append(_aRule, _aItem)
	}

	logs, sub, err := _Piece.contract.WatchLogs(opts, "Withdraw", _aRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PieceWithdraw)
				if err := _Piece.contract.UnpackLog(event, "Withdraw", log); err != nil {
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

// ParseWithdraw is a log parse operation binding the contract event 0x884edad9ce6fa2440d8a54cc123490eb96d2768479d49ff9c7366125a9424364.
//
// Solidity: event Withdraw(address indexed _a, uint256 _m)
func (_Piece *PieceFilterer) ParseWithdraw(log types.Log) (*PieceWithdraw, error) {
	event := new(PieceWithdraw)
	if err := _Piece.contract.UnpackLog(event, "Withdraw", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
