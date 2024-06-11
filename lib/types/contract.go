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

type ChallenegRSInfo struct {
	Store   common.Address
	Replica uint64
	Piece   uint64
}

func (ei *ChallenegRSInfo) Serialize() ([]byte, error) {
	return cbor.Marshal(ei)
}

func (ei *ChallenegRSInfo) Deserialize(b []byte) error {
	return cbor.Unmarshal(b, ei)
}
