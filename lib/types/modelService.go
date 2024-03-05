package types

import (
	"context"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/fxamacker/cbor/v2"
)

type ModelServiceStat struct {
	Stat uint8
	Last time.Time // heartbeat
}

type ModelServiceMeta struct {
	ModelServiceStat
	Name     string // hash(ModelName|Owner|Runner)
	Model    string
	Root     string
	Price    *big.Int
	Owner    common.Address
	Runner   common.Address
	Creation time.Time
}

func (msm *ModelServiceMeta) Serialize() ([]byte, error) {
	return cbor.Marshal(msm)
}

func (msm *ModelServiceMeta) Deserialize(b []byte) error {
	return cbor.Unmarshal(b, msm)
}

type IModelService interface {
	SubmitModelService(ctx context.Context, addr common.Address, ms ModelServiceMeta) (ModelServiceMeta, error)
	ConfirmModelService(ctx context.Context, addr common.Address, ms, root string, pf []byte) error
	UpdateModelService(ctx context.Context, addr common.Address, name string) error
	BuyModelService(ctx context.Context, addr common.Address, name string) error
	GetModelService(ctx context.Context, addr common.Address, name string) (ModelServiceResult, error)
	ListModelService(ctx context.Context, addr common.Address, opt Options) ([]ModelServiceResult, error)
}
