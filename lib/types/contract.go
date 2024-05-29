package types

import (
	"context"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
)

type IContract interface {
	BalanceOf(ctx context.Context, addr common.Address) *big.Int
	Recharge(ctx context.Context, addr common.Address, val *big.Int) error
}
