package contract

import (
	"context"
	"fmt"
	"time"

	"github.com/MOSSV2/dimo-sdk-go/contract/go/epoch"
	"github.com/MOSSV2/dimo-sdk-go/contract/go/file"
	"github.com/MOSSV2/dimo-sdk-go/contract/go/node"
	"github.com/MOSSV2/dimo-sdk-go/contract/go/proof"
	"github.com/MOSSV2/dimo-sdk-go/contract/go/round"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

func GetOrder(count uint64) (uint64, error) {
	ctx, cancle := context.WithTimeout(context.TODO(), 3*time.Second)
	defer cancle()
	client, err := ethclient.DialContext(ctx, DevChain)
	if err != nil {
		return 0, err
	}

	ri, err := round.NewRound(RoundAddr, client)
	if err != nil {
		return 0, err
	}
	return ri.GetOrder(&bind.CallOpts{From: Base}, count)
}

func Choose(addr common.Address, seed [32]byte, count uint64, i uint64) (uint64, error) {
	ctx, cancle := context.WithTimeout(context.TODO(), 3*time.Second)
	defer cancle()
	client, err := ethclient.DialContext(ctx, DevChain)
	if err != nil {
		return 0, err
	}

	ri, err := round.NewRound(RoundAddr, client)
	if err != nil {
		return 0, err
	}
	return ri.Choose(&bind.CallOpts{From: Base}, addr, seed, count, i)
}

func GetEpochSeed(_epoch uint64) ([32]byte, error) {
	ctx, cancle := context.WithTimeout(context.TODO(), 5*time.Second)
	defer cancle()
	client, err := ethclient.DialContext(ctx, DevChain)
	if err != nil {
		return [32]byte{}, err
	}

	ei, err := epoch.NewEpoch(EpochAddr, client)
	if err != nil {
		return [32]byte{}, err
	}

	return ei.GetEpochSeed(&bind.CallOpts{From: Base}, _epoch)
}

func CheckNode(addr common.Address, _typ uint8) error {
	ctx, cancle := context.WithTimeout(context.TODO(), 5*time.Second)
	defer cancle()
	client, err := ethclient.DialContext(ctx, DevChain)
	if err != nil {
		return err
	}

	ni, err := node.NewNode(NodeAddr, client)
	if err != nil {
		return err
	}

	return ni.Check(&bind.CallOpts{From: addr}, addr, _typ)
}

func GetFile(_fn string) (file.IFileFileBase, error) {
	ctx, cancle := context.WithTimeout(context.TODO(), 5*time.Second)
	defer cancle()
	client, err := ethclient.DialContext(ctx, DevChain)
	if err != nil {
		return file.IFileFileBase{}, err
	}
	fi, err := file.NewFile(FileAddr, client)
	if err != nil {
		return file.IFileFileBase{}, err
	}
	fb, err := fi.GetFile(&bind.CallOpts{From: Base}, _fn)
	if err != nil {
		return file.IFileFileBase{}, err
	}
	if fb.Size == 0 {
		return file.IFileFileBase{}, fmt.Errorf("no such file")
	}
	return fb, nil
}

func GetPiece(_fn string, _pi uint64) (file.IFilePieceInfo, error) {
	ctx, cancle := context.WithTimeout(context.TODO(), 5*time.Second)
	defer cancle()
	client, err := ethclient.DialContext(ctx, DevChain)
	if err != nil {
		return file.IFilePieceInfo{}, err
	}
	fi, err := file.NewFile(FileAddr, client)
	if err != nil {
		return file.IFilePieceInfo{}, err
	}
	pb, err := fi.GetPiece(&bind.CallOpts{From: Base}, _fn, _pi)
	if err != nil {
		return file.IFilePieceInfo{}, err
	}
	return pb, nil
}

func GetReplica(_ri uint64) (file.IFileReplicaInfo, error) {
	ctx, cancle := context.WithTimeout(context.TODO(), 5*time.Second)
	defer cancle()
	client, err := ethclient.DialContext(ctx, DevChain)
	if err != nil {
		return file.IFileReplicaInfo{}, err
	}
	fi, err := file.NewFile(FileAddr, client)
	if err != nil {
		return file.IFileReplicaInfo{}, err
	}
	pb, err := fi.GetReplica(&bind.CallOpts{From: Base}, _ri)
	if err != nil {
		return file.IFileReplicaInfo{}, err
	}

	return pb, nil
}

func GetStore(addr common.Address) (file.IFileStoreBase, error) {
	ctx, cancle := context.WithTimeout(context.TODO(), 5*time.Second)
	defer cancle()
	client, err := ethclient.DialContext(ctx, DevChain)
	if err != nil {
		return file.IFileStoreBase{}, err
	}
	fi, err := file.NewFile(FileAddr, client)
	if err != nil {
		return file.IFileStoreBase{}, err
	}
	fss, err := fi.GetStore(&bind.CallOpts{From: addr}, addr)
	if err != nil {
		return file.IFileStoreBase{}, err
	}
	return fss, nil
}

func GetStoreStat(addr common.Address, _epoch uint64) (file.IFileStoreStat, error) {
	ctx, cancle := context.WithTimeout(context.TODO(), 5*time.Second)
	defer cancle()
	client, err := ethclient.DialContext(ctx, DevChain)
	if err != nil {
		return file.IFileStoreStat{}, err
	}
	fi, err := file.NewFile(FileAddr, client)
	if err != nil {
		return file.IFileStoreStat{}, err
	}
	fss, err := fi.GetStoreStat(&bind.CallOpts{From: addr}, addr, _epoch)
	if err != nil {
		return file.IFileStoreStat{}, err
	}
	if fss.Count == 0 {
		return file.IFileStoreStat{}, fmt.Errorf("no storage")
	}
	return fss, nil
}

func GetStoreReplica(_a common.Address, _epoch, _ri uint64) (string, error) {
	ctx, cancle := context.WithTimeout(context.TODO(), 5*time.Second)
	defer cancle()
	client, err := ethclient.DialContext(ctx, DevChain)
	if err != nil {
		return "", err
	}
	fi, err := file.NewFile(FileAddr, client)
	if err != nil {
		return "", err
	}
	pb, err := fi.GetReplicaAt(&bind.CallOpts{From: Base}, _a, _epoch, _ri)
	if err != nil {
		return "", err
	}

	return SolByteToString(pb)
}

func GetEpochProof(addr common.Address, _epoch uint64) (proof.IProofEProof, error) {
	ctx, cancle := context.WithTimeout(context.TODO(), 1*time.Minute)
	defer cancle()
	client, err := ethclient.DialContext(ctx, DevChain)
	if err != nil {
		return proof.IProofEProof{}, err
	}
	pi, err := proof.NewProof(ProofAddr, client)
	if err != nil {
		return proof.IProofEProof{}, err
	}
	ep, err := pi.GetEProof(&bind.CallOpts{From: addr}, addr, _epoch)
	if err != nil {
		return proof.IProofEProof{}, err
	}
	return ep, nil
}
