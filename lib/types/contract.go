package types

import (
	"context"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
)

type IContract interface {
	BalanceOf(ctx context.Context, addr common.Address) *big.Int
	Recharge(ctx context.Context, addr common.Address, val *big.Int) error
	Pay(ctx context.Context, addr common.Address, val *big.Int) error

	Submit(ctx context.Context, addr common.Address, msg []byte) (SpaceMeta, error)
	SubmitModel(ctx context.Context, addr common.Address, msg []byte) (ModelMeta, error)
	SubmitGPU(ctx context.Context, addr common.Address, msg []byte) (GPUMeta, error)

	BuyModel(ctx context.Context, addr common.Address, modelName string) error
}
