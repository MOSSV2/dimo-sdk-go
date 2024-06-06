package contract

import (
	"fmt"
	"math/big"

	"github.com/MOSSV2/dimo-sdk-go/lib/types"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	etypes "github.com/ethereum/go-ethereum/core/types"
)

func HandleSetEpoch(log etypes.Log, cabi abi.ABI) (types.EpochInfo, error) {
	ei := types.EpochInfo{}

	evInfo, ok := cabi.Events["SetEpoch"]
	if !ok {
		return ei, fmt.Errorf("no event 'SetEpoch' in ABI")
	}

	ld, err := cabi.Unpack(evInfo.Name, log.Data)
	if err != nil {
		return ei, err
	}
	if len(ld) != 1 {
		return ei, fmt.Errorf("invalid log data length")
	}
	ei.Epoch = ld[0].(uint64)
	bh, seed, err := GetEpochInfo(ei.Epoch)
	if err != nil {
		return ei, err
	}
	ei.Height = bh
	ei.Seed = seed
	return ei, nil
}

func HandleAddPiece(log etypes.Log, cabi abi.ABI) (types.PieceCore, error) {
	pc := types.PieceCore{}

	evInfo, ok := cabi.Events["AddPiece"]
	if !ok {
		return pc, fmt.Errorf("no event 'AddPiece' in ABI")
	}

	if len(log.Topics) != 2 {
		return pc, fmt.Errorf("invalid log topic length")
	}
	pc.Owner = common.HexToAddress(log.Topics[1].Hex())

	ld, err := cabi.Unpack(evInfo.Name, log.Data)
	if err != nil {
		return pc, err
	}
	if len(ld) != 2 {
		return pc, fmt.Errorf("invalid log data length")
	}
	pc.Serial = ld[0].(uint64)
	pc.Start = ld[1].(uint64)

	tx, err := GetTransaction(log.TxHash)
	if err != nil {
		return pc, err
	}

	method, ok := cabi.Methods["addPiece"]
	if !ok {
		return pc, fmt.Errorf("no method 'addPiece' in ABI")
	}

	inputData := tx.Data()
	inputs, err := method.Inputs.UnpackValues(inputData[4:])
	if err != nil {
		return pc, err
	}

	if len(inputs) != 6 {
		return pc, fmt.Errorf("invalid input length")
	}
	g1, err := SolidityToG1(inputs[0].([]byte))
	if err != nil {
		return pc, err
	}
	pc.Name = G1ToString(g1)
	pc.Price = inputs[1].(*big.Int)
	pc.Size = int64(inputs[2].(uint64))
	pc.Expire = inputs[3].(uint64)
	pc.Policy.N = inputs[4].(uint8)
	pc.Policy.K = inputs[5].(uint8)

	return pc, nil
}

func HandleAddReplica(log etypes.Log, cabi abi.ABI) (types.ReplicaInChain, error) {
	rc := types.ReplicaInChain{
		Witness: types.ReplicaWitness{},
	}

	evInfo, ok := cabi.Events["AddReplica"]
	if !ok {
		return rc, fmt.Errorf("no event 'AddReplica' in ABI")
	}

	if len(log.Topics) != 2 {
		return rc, fmt.Errorf("invalid log topic length")
	}
	rc.StoredOn = common.HexToAddress(log.Topics[1].Hex())

	ld, err := cabi.Unpack(evInfo.Name, log.Data)
	if err != nil {
		fmt.Println(err)
	}
	if len(ld) != 2 {
		return rc, fmt.Errorf("invalid log data length")
	}
	rc.Serial = ld[0].(uint64)
	rc.Witness.Index = ld[1].(uint64)

	tx, err := GetTransaction(log.TxHash)
	if err != nil {
		return rc, err
	}

	method, ok := cabi.Methods["addReplica"]
	if !ok {
		return rc, fmt.Errorf("no method 'addReplica' in ABI")
	}

	inputData := tx.Data()
	inputs, err := method.Inputs.UnpackValues(inputData[4:])
	if err != nil {
		return rc, err
	}

	if len(inputs) != 4 {
		return rc, fmt.Errorf("invalid input length")
	}

	g1, err := SolidityToG1(inputs[0].([]byte))
	if err != nil {
		return rc, err
	}
	rc.Name = G1ToString(g1)
	rc.Piece = inputs[1].(uint64)
	rc.Index = inputs[2].(uint8)
	rc.Witness.Proof = inputs[3].([]byte)

	return rc, nil
}
