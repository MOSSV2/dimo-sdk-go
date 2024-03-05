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

	SubmitModelService(ctx context.Context, addr common.Address, msg []byte) (ModelServiceMeta, error)
	SubmitModel(ctx context.Context, addr common.Address, msg []byte) (ModelMeta, error)
	SubmitPrompt(ctx context.Context, addr common.Address, msg []byte) (PromptMeta, error)
	SubmitNFT(ctx context.Context, addr common.Address, msg []byte) (NFTMeta, error)

	BuyModel(ctx context.Context, addr common.Address, modelName string) error
	BuyNFT(ctx context.Context, addr common.Address, nftName string) error
}
