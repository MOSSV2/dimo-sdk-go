package contract

import (
	"context"
	"crypto/ecdsa"
	"encoding/binary"
	"encoding/hex"
	"fmt"
	"math/big"
	"time"

	"github.com/MOSSV2/dimo-sdk-go/lib/log"
	"github.com/MOSSV2/dimo-sdk-go/lib/types"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
)

var logger = log.Logger("contract")

func Set(sk *ecdsa.PrivateKey, _typ string, ca common.Address) error {
	ctx, cancle := context.WithTimeout(context.TODO(), 1*time.Minute)
	defer cancle()
	bi, err := NewBank(ctx)
	if err != nil {
		return err
	}

	au, err := makeAuth(big.NewInt(int64(DevChainID)), sk)
	if err != nil {
		return err
	}

	tx, err := bi.Set(au, _typ, ca)
	if err != nil {
		return err
	}

	err = CheckTx(DevChain, tx.Hash())
	if err != nil {
		return err
	}
	fmt.Println(_typ, " set success to: ", ca.String())
	return nil
}

func GetEpoch(sk *ecdsa.PrivateKey) (uint64, error) {
	ctx, cancle := context.WithTimeout(context.TODO(), 1*time.Minute)
	defer cancle()
	ei, err := NewEpoch(ctx)
	if err != nil {
		return 0, err
	}

	au, err := makeAuth(big.NewInt(int64(DevChainID)), sk)
	if err != nil {
		return 0, err
	}

	tx, err := ei.Check(au)
	if err != nil {
		return 0, err
	}
	err = CheckTx(DevChain, tx.Hash())
	if err != nil {
		return 0, err
	}

	return ei.Current(&bind.CallOpts{From: au.From})
}

func RegisterNode(sk *ecdsa.PrivateKey, _typ uint8) error {
	ctx, cancle := context.WithTimeout(context.TODO(), 3*time.Minute)
	defer cancle()
	ni, err := NewNode(ctx)
	if err != nil {
		return err
	}

	au, err := makeAuth(big.NewInt(int64(DevChainID)), sk)
	if err != nil {
		return err
	}

	_, err = ni.Check(&bind.CallOpts{From: au.From}, au.From, _typ)
	if err == nil {
		logger.Debug("already pledge")
		return nil
	}

	ti, err := NewToken(ctx)
	if err != nil {
		return err
	}

	pval, err := ni.GetMinPledge(&bind.CallOpts{From: au.From}, _typ)
	if err != nil {
		return err
	}
	pinfo, err := ni.GetPledge(&bind.CallOpts{From: au.From}, au.From, _typ)
	if err != nil {
		return err
	}
	pval.Sub(pval, pinfo.Value)

	tx, err := ti.IncreaseAllowance(au, BankAddr, pval)
	if err != nil {
		return err
	}
	err = CheckTx(DevChain, tx.Hash())
	if err != nil {
		return err
	}
	tx, err = ni.Pledge(au, _typ, pval)
	if err != nil {
		return err
	}
	err = CheckTx(DevChain, tx.Hash())
	if err != nil {
		return err
	}
	_, err = ni.Check(&bind.CallOpts{From: au.From}, au.From, _typ)
	return err
}

func AddPiece(sk *ecdsa.PrivateKey, pc types.PieceCore) error {
	logger.Debug("add piece: ", pc)
	ctx, cancle := context.WithTimeout(context.TODO(), 1*time.Minute)
	defer cancle()
	fi, err := NewPiece(ctx)
	if err != nil {
		return err
	}

	au, err := makeAuth(big.NewInt(int64(DevChainID)), sk)
	if err != nil {
		return err
	}

	ce, err := GetEpoch(sk)
	if err != nil {
		return err
	}

	if pc.Expire == 0 {
		pc.Start = ce
		pc.Expire = ce + uint64(DefaultStoreEpoch) + 1
	}
	if pc.Price == nil {
		pc.Price = new(big.Int).SetInt64(int64(DefaultReplicaPrice))
	}

	_size := 1 + (pc.Size-1)/(31*int64(pc.Policy.K))

	_size = (1 + (_size-1)/(32*1024)) * int64(pc.Policy.N)

	val := big.NewInt(int64(pc.Expire-pc.Start) * _size)
	val.Mul(val, pc.Price)

	ti, err := NewToken(ctx)
	if err != nil {
		return err
	}

	tx, err := ti.IncreaseAllowance(au, BankAddr, val)
	if err != nil {
		return err
	}
	err = CheckTx(DevChain, tx.Hash())
	if err != nil {
		return err
	}

	pb, err := G1StringInSolidity(pc.Name)
	if err != nil {
		return err
	}

	fmt.Println("add piece: ", pc)
	tx, err = fi.AddPiece(au, pb, pc.Price, uint64(pc.Size), pc.Expire, pc.Policy.N, pc.Policy.K)
	if err != nil {
		return err
	}
	err = CheckTx(DevChain, tx.Hash())
	if err != nil {
		return err
	}

	return nil
}

func AddReplica(sk *ecdsa.PrivateKey, rc types.ReplicaCore, pf []byte) error {
	rb, err := G1StringInSolidity(rc.Name)
	if err != nil {
		return err
	}
	ctx, cancle := context.WithTimeout(context.TODO(), 5*time.Minute)
	defer cancle()
	fi, err := NewPiece(ctx)
	if err != nil {
		return err
	}

	au, err := makeAuth(big.NewInt(int64(DevChainID)), sk)
	if err != nil {
		return err
	}

	pbyte, err := G1StringInSolidity(rc.Piece)
	if err != nil {
		return err
	}
	_pi, err := fi.GetPIndex(&bind.CallOpts{From: au.From}, pbyte)
	if err != nil {
		return err
	}

	tx, err := fi.AddReplica(au, rb, _pi, rc.Index, pf)
	if err != nil {
		return err
	}
	err = CheckTx(DevChain, tx.Hash())
	if err != nil {
		return err
	}

	return nil
}

func SubmitProof(sk *ecdsa.PrivateKey, _ep uint64, _com []byte, pf []byte) error {
	ctx, cancle := context.WithTimeout(context.TODO(), 1*time.Minute)
	defer cancle()
	pi, err := NewEProof(ctx)
	if err != nil {
		return err
	}

	au, err := makeAuth(big.NewInt(int64(DevChainID)), sk)
	if err != nil {
		return err
	}

	tx, err := pi.Submit(au, _ep, _com, pf)
	if err != nil {
		return err
	}

	err = CheckTx(DevChain, tx.Hash())
	if err != nil {
		return err
	}

	return nil
}

func Settle(sk *ecdsa.PrivateKey, _epoch uint64) error {
	ctx, cancle := context.WithTimeout(context.TODO(), 1*time.Minute)
	defer cancle()
	fi, err := NewPiece(ctx)
	if err != nil {
		return err
	}
	au, err := makeAuth(big.NewInt(int64(DevChainID)), sk)
	if err != nil {
		return err
	}

	tx, err := fi.Settle(au, _epoch)
	if err != nil {
		return err
	}

	err = CheckTx(DevChain, tx.Hash())
	if err != nil {
		return err
	}

	return nil
}

func AddModel(sk *ecdsa.PrivateKey, mc types.ModelMeta) error {
	ctx, cancle := context.WithTimeout(context.TODO(), 1*time.Minute)
	defer cancle()
	mi, err := NewModel(ctx)
	if err != nil {
		return err
	}
	au, err := makeAuth(big.NewInt(int64(DevChainID)), sk)
	if err != nil {
		return err
	}

	rt, err := hex.DecodeString(mc.Hash)
	if err != nil {
		return err
	}

	var cnt [32]byte
	binary.BigEndian.PutUint64(cnt[24:32], uint64(mc.Count))

	tx, err := mi.Add(au, mc.Name, rt, cnt)
	if err != nil {
		return err
	}

	err = CheckTx(DevChain, tx.Hash())
	if err != nil {
		return err
	}

	return nil
}

func AddGPU(sk *ecdsa.PrivateKey, _gn string) error {
	ctx, cancle := context.WithTimeout(context.TODO(), 1*time.Minute)
	defer cancle()
	gi, err := NewGPU(ctx)
	if err != nil {
		return err
	}
	au, err := makeAuth(big.NewInt(int64(DevChainID)), sk)
	if err != nil {
		return err
	}

	tx, err := gi.Add(au, _gn)
	if err != nil {
		return err
	}

	err = CheckTx(DevChain, tx.Hash())
	if err != nil {
		return err
	}

	return nil
}

func AddSpace(sk *ecdsa.PrivateKey, msm types.SpaceMeta) error {
	ctx, cancle := context.WithTimeout(context.TODO(), 3*time.Minute)
	defer cancle()
	si, err := NewSpace(ctx)
	if err != nil {
		return err
	}

	gi, err := NewGPU(ctx)
	if err != nil {
		return err
	}

	mi, err := NewModel(ctx)
	if err != nil {
		return err
	}

	au, err := makeAuth(big.NewInt(int64(DevChainID)), sk)
	if err != nil {
		return err
	}

	_mi, err := mi.GetIndex(&bind.CallOpts{From: au.From}, msm.Model)
	if err != nil {
		return err
	}

	_gi, err := gi.GetIndex(&bind.CallOpts{From: au.From}, msm.GPU)
	if err != nil {
		return err
	}

	ce, err := GetEpoch(sk)
	if err != nil {
		return err
	}

	val := big.NewInt(int64(DefaultSpacePrice))
	val.Mul(val, big.NewInt(int64(DefaultSpaceEpoch)))

	ti, err := NewToken(ctx)
	if err != nil {
		return err
	}

	tx, err := ti.IncreaseAllowance(au, BankAddr, val)
	if err != nil {
		return err
	}
	err = CheckTx(DevChain, tx.Hash())
	if err != nil {
		return err
	}

	tx, err = si.Add(au, msm.Name, _mi, _gi, big.NewInt(int64(DefaultSpacePrice)), ce+uint64(DefaultSpaceEpoch))
	if err != nil {
		return err
	}

	err = CheckTx(DevChain, tx.Hash())
	if err != nil {
		return err
	}

	return nil
}

func ActivateSpace(sk *ecdsa.PrivateKey, sn, root string, pfbyte []byte) error {
	ctx, cancle := context.WithTimeout(context.TODO(), 3*time.Minute)
	defer cancle()
	si, err := NewSpace(ctx)
	if err != nil {
		return err
	}

	au, err := makeAuth(big.NewInt(int64(DevChainID)), sk)
	if err != nil {
		return err
	}

	_ai, err := si.GetIndex(&bind.CallOpts{From: au.From}, sn)
	if err != nil {
		return err
	}

	_rt, err := hex.DecodeString(root)
	if err != nil {
		return err
	}

	tx, err := si.Activate(au, _ai, _rt, pfbyte)
	if err != nil {
		return err
	}

	err = CheckTx(DevChain, tx.Hash())
	if err != nil {
		return err
	}

	return nil
}

func ShutdownSpace(sk *ecdsa.PrivateKey, _ai uint64) error {
	ctx, cancle := context.WithTimeout(context.TODO(), 3*time.Minute)
	defer cancle()
	si, err := NewSpace(ctx)
	if err != nil {
		return err
	}

	au, err := makeAuth(big.NewInt(int64(DevChainID)), sk)
	if err != nil {
		return err
	}

	tx, err := si.Shutdown(au, _ai)
	if err != nil {
		return err
	}

	err = CheckTx(DevChain, tx.Hash())
	if err != nil {
		return err
	}

	return nil
}

func GPUCheck(sk *ecdsa.PrivateKey) error {
	ctx, cancle := context.WithTimeout(context.TODO(), 3*time.Minute)
	defer cancle()
	gi, err := NewGPU(ctx)
	if err != nil {
		return err
	}

	au, err := makeAuth(big.NewInt(int64(DevChainID)), sk)
	if err != nil {
		return err
	}

	tx, err := gi.Check(au)
	if err != nil {
		return err
	}

	return CheckTx(DevChain, tx.Hash())
}

func GPUMint(sk *ecdsa.PrivateKey, _gi uint64, _salt []byte) error {
	ctx, cancle := context.WithTimeout(context.TODO(), 3*time.Minute)
	defer cancle()
	gi, err := NewGPU(ctx)
	if err != nil {
		return err
	}

	au, err := makeAuth(big.NewInt(int64(DevChainID)), sk)
	if err != nil {
		return err
	}

	tx, err := gi.Mint(au, _gi, _salt)
	if err != nil {
		return err
	}

	return CheckTx(DevChain, tx.Hash())
}

func Mint(sk *ecdsa.PrivateKey) error {
	ctx, cancle := context.WithTimeout(context.TODO(), 1*time.Minute)
	defer cancle()
	gi, err := NewControl(ctx)
	if err != nil {
		return err
	}

	baddr, err := gi.Bank(&bind.CallOpts{From: Base})
	if err != nil {
		return err
	}

	fmt.Println(baddr)

	au, err := makeAuth(big.NewInt(int64(DevChainID)), sk)
	if err != nil {
		return err
	}

	tx, err := gi.Mint(au, au.From, big.NewInt(100))
	if err != nil {
		return err
	}

	err = CheckTx(DevChain, tx.Hash())
	if err != nil {
		return err
	}

	return nil
}
