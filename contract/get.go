package contract

import (
	"context"
	"fmt"
	"math/big"
	"time"

	"github.com/MOSSV2/dimo-sdk-go/contract/go/gpu"
	"github.com/MOSSV2/dimo-sdk-go/contract/go/model"
	"github.com/MOSSV2/dimo-sdk-go/contract/go/piece"
	"github.com/MOSSV2/dimo-sdk-go/contract/go/space"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
)

func GetInstAddr(ctx context.Context, typ string) (common.Address, error) {
	bi, err := NewBank(ctx)
	if err != nil {
		return common.Address{}, err
	}

	return bi.Get(&bind.CallOpts{From: Base}, typ)
}

func GetOrder(pcnt, count uint64) (uint8, error) {
	ctx, cancle := context.WithTimeout(context.TODO(), 3*time.Second)
	defer cancle()

	ri, err := NewEVerify(ctx)
	if err != nil {
		return 0, err
	}
	return ri.GetOrder(&bind.CallOpts{From: Base}, pcnt, count)
}

func Choose(addr common.Address, seed [32]byte, pcnt, count uint64, i uint64) (uint64, error) {
	ctx, cancle := context.WithTimeout(context.TODO(), 3*time.Second)
	defer cancle()
	ri, err := NewEVerify(ctx)
	if err != nil {
		return 0, err
	}
	return ri.Choose(&bind.CallOpts{From: Base}, addr, seed, pcnt, count, i)
}

func GetEpochInfo(_epoch uint64) (*big.Int, [32]byte, error) {
	ctx, cancle := context.WithTimeout(context.TODO(), 5*time.Second)
	defer cancle()

	ei, err := NewEpoch(ctx)
	if err != nil {
		return nil, [32]byte{}, err
	}

	return ei.GetEpoch(&bind.CallOpts{From: Base}, _epoch)
}

func CheckNode(addr common.Address, _typ uint8) error {
	ctx, cancle := context.WithTimeout(context.TODO(), 5*time.Second)
	defer cancle()

	ni, err := NewNode(ctx)
	if err != nil {
		return err
	}

	_, err = ni.Check(&bind.CallOpts{From: addr}, addr, _typ)
	return err
}

func GetPieceSerial(_pn string) (uint64, error) {
	pnb, err := G1StringInSolidity(_pn)
	if err != nil {
		return 0, err
	}
	ctx, cancle := context.WithTimeout(context.TODO(), 5*time.Second)
	defer cancle()
	gi, err := NewPiece(ctx)
	if err != nil {
		return 0, err
	}
	return gi.GetPIndex(&bind.CallOpts{From: Base}, pnb)
}

func GetReplicaSerial(_pn string) (uint64, error) {
	pnb, err := G1StringInSolidity(_pn)
	if err != nil {
		return 0, err
	}
	ctx, cancle := context.WithTimeout(context.TODO(), 5*time.Second)
	defer cancle()
	gi, err := NewPiece(ctx)
	if err != nil {
		return 0, err
	}
	return gi.GetRIndex(&bind.CallOpts{From: Base}, pnb)
}

func GetPiece(_pi uint64) (piece.IPiecePieceInfo, error) {
	ctx, cancle := context.WithTimeout(context.TODO(), 5*time.Second)
	defer cancle()
	fi, err := NewPiece(ctx)
	if err != nil {
		return piece.IPiecePieceInfo{}, err
	}
	pb, err := fi.GetPiece(&bind.CallOpts{From: Base}, _pi)
	if err != nil {
		return piece.IPiecePieceInfo{}, err
	}
	return pb, nil
}

func GetReplica(_ri uint64) (piece.IPieceReplicaInfo, error) {
	ctx, cancle := context.WithTimeout(context.TODO(), 5*time.Second)
	defer cancle()
	fi, err := NewPiece(ctx)
	if err != nil {
		return piece.IPieceReplicaInfo{}, err
	}
	pb, err := fi.GetReplica(&bind.CallOpts{From: Base}, _ri)
	if err != nil {
		return piece.IPieceReplicaInfo{}, err
	}

	return pb, nil
}

func GetStore(addr common.Address) (piece.IPieceStoreInfo, error) {
	ctx, cancle := context.WithTimeout(context.TODO(), 5*time.Second)
	defer cancle()
	fi, err := NewPiece(ctx)
	if err != nil {
		return piece.IPieceStoreInfo{}, err
	}
	fss, err := fi.GetStore(&bind.CallOpts{From: addr}, addr)
	if err != nil {
		return piece.IPieceStoreInfo{}, err
	}
	return fss, nil
}

func GetStoreStat(addr common.Address, _epoch uint64) (piece.IPieceStoreStat, error) {
	ctx, cancle := context.WithTimeout(context.TODO(), 5*time.Second)
	defer cancle()
	fi, err := NewPiece(ctx)
	if err != nil {
		return piece.IPieceStoreStat{}, err
	}
	fss, err := fi.GetSStat(&bind.CallOpts{From: addr}, addr, _epoch)
	if err != nil {
		return piece.IPieceStoreStat{}, err
	}
	return fss, nil
}

func GetStoreReplica(_a common.Address, _ri uint64) (uint64, error) {
	ctx, cancle := context.WithTimeout(context.TODO(), 5*time.Second)
	defer cancle()
	fi, err := NewPiece(ctx)
	if err != nil {
		return 0, err
	}
	return fi.GetSRAt(&bind.CallOpts{From: Base}, _a, _ri)
}

func GetRevenue(addr common.Address, typ string) (*big.Int, error) {
	res := big.NewInt(0)
	ctx, cancle := context.WithTimeout(context.TODO(), 1*time.Minute)
	defer cancle()

	switch typ {
	case "gpu":
		gi, err := NewGPU(ctx)
		if err != nil {
			return res, err
		}
		return gi.BalanceOf(&bind.CallOpts{From: addr}, addr)
	case "space":
		si, err := NewSpace(ctx)
		if err != nil {
			return res, err
		}
		return si.BalanceOf(&bind.CallOpts{From: addr}, addr)
	default:
		return res, fmt.Errorf("unsupported")
	}
}

func GetGPUIndex(_gn string) (uint64, error) {
	ctx, cancle := context.WithTimeout(context.TODO(), 5*time.Second)
	defer cancle()
	gi, err := NewGPU(ctx)
	if err != nil {
		return 0, err
	}
	return gi.GetIndex(&bind.CallOpts{From: Base}, _gn)
}

func GetGPUInfo(_gi uint64) (gpu.IGPUInfo, error) {
	ctx, cancle := context.WithTimeout(context.TODO(), 5*time.Second)
	defer cancle()
	gi, err := NewGPU(ctx)
	if err != nil {
		return gpu.IGPUInfo{}, err
	}
	return gi.GetGPU(&bind.CallOpts{From: Base}, _gi)
}

func GetGPUOwner(_gi uint64) (common.Address, error) {
	ctx, cancle := context.WithTimeout(context.TODO(), 5*time.Second)
	defer cancle()
	gi, err := NewGPU(ctx)
	if err != nil {
		return common.Address{}, err
	}
	return gi.GetOwner(&bind.CallOpts{From: Base}, _gi)
}

func GetGPUSetting() (*big.Int, [32]byte, error) {
	ctx, cancle := context.WithTimeout(context.TODO(), 5*time.Second)
	defer cancle()
	gi, err := NewGPU(ctx)
	if err != nil {
		return nil, [32]byte{}, err
	}

	dif, err := gi.Difficulty(&bind.CallOpts{From: Base})
	if err != nil {
		return nil, [32]byte{}, err
	}

	seed, err := gi.Seed(&bind.CallOpts{From: Base})
	if err != nil {
		return nil, [32]byte{}, err
	}
	return dif, seed, nil
}

func GetModelIndex(_mn string) (uint64, error) {
	ctx, cancle := context.WithTimeout(context.TODO(), 5*time.Second)
	defer cancle()
	mi, err := NewModel(ctx)
	if err != nil {
		return 0, err
	}
	return mi.GetIndex(&bind.CallOpts{From: Base}, _mn)
}

func GetModelInfo(_mi uint64) (model.IModelInfo, error) {
	ctx, cancle := context.WithTimeout(context.TODO(), 5*time.Second)
	defer cancle()
	mi, err := NewModel(ctx)
	if err != nil {
		return model.IModelInfo{}, err
	}
	return mi.GetModel(&bind.CallOpts{From: Base}, _mi)
}

func GetSpaceIndex(_mn string) (uint64, error) {
	ctx, cancle := context.WithTimeout(context.TODO(), 5*time.Second)
	defer cancle()
	si, err := NewSpace(ctx)
	if err != nil {
		return 0, err
	}
	return si.GetIndex(&bind.CallOpts{From: Base}, _mn)
}

func GetSpaceInfo(_mi uint64) (space.ISpaceInfo, error) {
	ctx, cancle := context.WithTimeout(context.TODO(), 5*time.Second)
	defer cancle()
	si, err := NewSpace(ctx)
	if err != nil {
		return space.ISpaceInfo{}, err
	}
	return si.GetSpace(&bind.CallOpts{From: Base}, _mi)
}
