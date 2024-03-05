package contract

import (
	"context"
	"crypto/ecdsa"
	"encoding/hex"
	"fmt"
	"math/big"
	"time"

	"github.com/MOSSV2/dimo-sdk-go/contract/go/epoch"
	"github.com/MOSSV2/dimo-sdk-go/contract/go/file"
	"github.com/MOSSV2/dimo-sdk-go/contract/go/node"
	"github.com/MOSSV2/dimo-sdk-go/contract/go/proof"
	"github.com/MOSSV2/dimo-sdk-go/contract/go/token"
	"github.com/MOSSV2/dimo-sdk-go/lib/log"
	"github.com/MOSSV2/dimo-sdk-go/lib/types"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

var logger = log.Logger("contract")

func GetEpoch(sk *ecdsa.PrivateKey) (uint64, error) {
	ctx, cancle := context.WithTimeout(context.TODO(), 1*time.Minute)
	defer cancle()
	client, err := ethclient.DialContext(ctx, DevChain)
	if err != nil {
		return 0, err
	}

	au, err := makeAuth(big.NewInt(int64(DevChainID)), sk)
	if err != nil {
		return 0, err
	}

	ei, err := epoch.NewEpoch(EpochAddr, client)
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

	return ei.CurEpoch(&bind.CallOpts{From: au.From})
}

func RegisterNode(sk *ecdsa.PrivateKey, _typ uint8) error {
	ctx, cancle := context.WithTimeout(context.TODO(), 3*time.Minute)
	defer cancle()
	client, err := ethclient.DialContext(ctx, DevChain)
	if err != nil {
		return err
	}

	au, err := makeAuth(big.NewInt(int64(DevChainID)), sk)
	if err != nil {
		return err
	}

	ni, err := node.NewNode(NodeAddr, client)
	if err != nil {
		return err
	}

	err = ni.Check(&bind.CallOpts{From: au.From}, au.From, _typ)
	if err == nil {
		logger.Debug("already pledge")
		return nil
	}

	ti, err := token.NewToken(TokenAddr, client)
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
	return ni.Check(&bind.CallOpts{From: au.From}, au.From, _typ)
}

func AddFile(sk *ecdsa.PrivateKey, _fn string, fc types.FileCore) error {
	ctx, cancle := context.WithTimeout(context.TODO(), 1*time.Minute)
	defer cancle()
	client, err := ethclient.DialContext(ctx, DevChain)
	if err != nil {
		return err
	}

	au, err := makeAuth(big.NewInt(int64(DevChainID)), sk)
	if err != nil {
		return err
	}

	fi, err := file.NewFile(FileAddr, client)
	if err != nil {
		return err
	}

	ce, err := GetEpoch(sk)
	if err != nil {
		return err
	}

	root, err := hex.DecodeString(fc.Hash)
	if err != nil {
		return err
	}

	_f := file.IFileFileBase{
		Owner:  au.From,
		Price:  big.NewInt(int64(ReplicaPrice)),
		Size:   uint64(fc.Size),
		Expire: ce + 100,
		Root:   root,
	}

	logger.Debug("add file: ", _fn)
	fb, err := fi.GetFile(&bind.CallOpts{From: au.From}, _fn)
	if err == nil && fb.Expire > 0 {
		logger.Debug("%s already has: %s", fb.Owner, _fn)
		return fmt.Errorf("%s already has: %s", fb.Owner, _fn)
	}

	tx, err := fi.AddFile(au, _fn, _f)
	if err != nil {
		return err
	}
	err = CheckTx(DevChain, tx.Hash())
	if err != nil {
		return err
	}

	return nil
}

func AddPiece(sk *ecdsa.PrivateKey, _fn string, pieces []string) error {
	ctx, cancle := context.WithTimeout(context.TODO(), 1*time.Minute)
	defer cancle()
	client, err := ethclient.DialContext(ctx, DevChain)
	if err != nil {
		return err
	}

	au, err := makeAuth(big.NewInt(int64(DevChainID)), sk)
	if err != nil {
		return err
	}

	fi, err := file.NewFile(FileAddr, client)
	if err != nil {
		return err
	}

	piecebytes := make([][]byte, 0, len(pieces))
	for i := 0; i < len(pieces); i++ {
		pb, err := StringToSolByte(pieces[i])
		if err != nil {
			return err
		}
		piecebytes = append(piecebytes, pb)
	}

	for i := 0; i < len(pieces); i += 128 {
		end := i + 128
		if end > len(pieces) {
			end = len(pieces)
		}
		logger.Debug("add piece to: ", _fn)
		tx, err := fi.AddPiece(au, _fn, piecebytes[i:end])
		if err != nil {
			return err
		}
		err = CheckTx(DevChain, tx.Hash())
		if err != nil {
			return err
		}
	}

	return nil
}

func SetFileProxy(sk *ecdsa.PrivateKey, _fn string, _proxy common.Address) error {
	ctx, cancle := context.WithTimeout(context.TODO(), 1*time.Minute)
	defer cancle()
	client, err := ethclient.DialContext(ctx, DevChain)
	if err != nil {
		return err
	}

	au, err := makeAuth(big.NewInt(int64(DevChainID)), sk)
	if err != nil {
		return err
	}

	fi, err := file.NewFile(FileAddr, client)
	if err != nil {
		return err
	}

	logger.Debug(_fn, " set proxy to: ", _proxy)
	tx, err := fi.SetProxy(au, _fn, _proxy)
	if err != nil {
		return err
	}
	err = CheckTx(DevChain, tx.Hash())
	if err != nil {
		return err
	}

	logger.Debug(" allowrance to: ", _fn)
	fb, err := GetFile(_fn)
	if err != nil {
		return err
	}

	val := new(big.Int).Mul(fb.Price, big.NewInt(int64(fb.Count*3)))
	val.Mul(val, big.NewInt(int64(fb.Expire-fb.Start)))

	ti, err := token.NewToken(TokenAddr, client)
	if err != nil {
		return err
	}

	tx, err = ti.IncreaseAllowance(au, BankAddr, val)
	if err != nil {
		return err
	}
	err = CheckTx(DevChain, tx.Hash())
	if err != nil {
		return err
	}

	return nil
}

func AddReplica(sk *ecdsa.PrivateKey, _fn string, pIndex uint64, replica string, _sign []byte) error {
	rb, err := StringToSolByte(replica)
	if err != nil {
		return err
	}
	ctx, cancle := context.WithTimeout(context.TODO(), 3*time.Minute)
	defer cancle()
	client, err := ethclient.DialContext(ctx, DevChain)
	if err != nil {
		return err
	}

	au, err := makeAuth(big.NewInt(int64(DevChainID)), sk)
	if err != nil {
		return err
	}

	fi, err := file.NewFile(FileAddr, client)
	if err != nil {
		return err
	}

	fb, err := fi.GetFile(&bind.CallOpts{From: au.From}, _fn)
	if err != nil {
		return err
	}

	if fb.Size == 0 {
		return fmt.Errorf("no file: %s", _fn)
	}

	logger.Debug("add replica: ", replica)
	if fb.Owner == au.From {
		ce, err := GetEpoch(sk)
		if err != nil {
			return err
		}

		val := big.NewInt(int64(fb.Expire - ce))
		val.Mul(val, fb.Price)

		ti, err := token.NewToken(TokenAddr, client)
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
	}

	tx, err := fi.AddReplica(au, _fn, pIndex, rb, _sign)
	if err != nil {
		return err
	}
	err = CheckTx(DevChain, tx.Hash())
	if err != nil {
		return err
	}

	return nil
}

func SubmitProof(sk *ecdsa.PrivateKey, _kp proof.IProofKZGProof) error {
	ctx, cancle := context.WithTimeout(context.TODO(), 1*time.Minute)
	defer cancle()
	client, err := ethclient.DialContext(ctx, DevChain)
	if err != nil {
		return err
	}

	au, err := makeAuth(big.NewInt(int64(DevChainID)), sk)
	if err != nil {
		return err
	}

	pi, err := proof.NewProof(ProofAddr, client)
	if err != nil {
		return err
	}

	tx, err := pi.Submit(au, _kp)
	if err != nil {
		return err
	}

	err = CheckTx(DevChain, tx.Hash())
	if err != nil {
		return err
	}

	return nil
}

func WithdrawFromFile(sk *ecdsa.PrivateKey, val *big.Int) error {
	ctx, cancle := context.WithTimeout(context.TODO(), 1*time.Minute)
	defer cancle()
	client, err := ethclient.DialContext(ctx, DevChain)
	if err != nil {
		return err
	}

	au, err := makeAuth(big.NewInt(int64(DevChainID)), sk)
	if err != nil {
		return err
	}

	fi, err := file.NewFile(FileAddr, client)
	if err != nil {
		return err
	}

	fsb, err := fi.GetStore(&bind.CallOpts{From: au.From}, au.From)
	if err != nil {
		return err
	}

	tx, err := fi.Withdraw(au, fsb.Profit.Revenue)
	if err != nil {
		return err
	}

	err = CheckTx(DevChain, tx.Hash())
	if err != nil {
		return err
	}

	return nil
}
