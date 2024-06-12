package types

import (
	"context"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/fxamacker/cbor/v2"
)

type IContract interface {
	BalanceOf(ctx context.Context, addr common.Address) *big.Int
	Recharge(ctx context.Context, addr common.Address, val *big.Int) error
}

type EpochInfo struct {
	Epoch  uint64
	Height *big.Int
	Seed   [32]byte
}

func (ei *EpochInfo) Serialize() ([]byte, error) {
	return cbor.Marshal(ei)
}

func (ei *EpochInfo) Deserialize(b []byte) error {
	return cbor.Unmarshal(b, ei)
}

type RSChallengeInfo struct {
	Store   common.Address
	Replica uint64
}

func (ei *RSChallengeInfo) Serialize() ([]byte, error) {
	return cbor.Marshal(ei)
}

func (ei *RSChallengeInfo) Deserialize(b []byte) error {
	return cbor.Unmarshal(b, ei)
}

type ReplicaInChain struct {
	Name     string
	Serial   uint64
	Piece    uint64
	Index    uint8
	StoredOn common.Address
	Witness  ReplicaWitness
}

type EpochProof struct {
	Epoch uint64
	Store common.Address
	Sum   []byte
	H     []byte
	Value []byte
}

func (ei *EpochProof) Serialize() ([]byte, error) {
	return cbor.Marshal(ei)
}

func (ei *EpochProof) Deserialize(b []byte) error {
	return cbor.Unmarshal(b, ei)
}
