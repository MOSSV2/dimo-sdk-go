package sage

import (
	"bytes"
	"fmt"

	bls "github.com/MOSSV2/dimo-sdk-go/lib/bls12377"
	"github.com/MOSSV2/dimo-sdk-go/lib/lighthash"
	"github.com/MOSSV2/dimo-sdk-go/lib/merkle"
	"github.com/MOSSV2/dimo-sdk-go/lib/types"
)

func BuildProof(b, prev, salt []byte, pindex int) (types.SageMerkleProof, error) {
	shard := 1 + (len(b)-1)/PadSize
	subshard := 1 << 10
	smp := types.SageMerkleProof{
		Leaf: uint64(pindex),
	}
	h := lighthash.New()
	droot, dsets, _, err := merkle.BuildReaderProof(bytes.NewReader(b), h, h.BlockSize(), uint64(pindex))
	if err != nil {
		return smp, err
	}
	smp.PieceRoot = droot
	smp.PiecePath = dsets

	err = EncodeForward(b, prev, salt)
	if err != nil {
		return smp, err
	}

	if subshard < shard {
		hroot, aux, err := merkle.ReaderRootAux(bytes.NewReader(b), h, PadSize, subshard)
		if err != nil {
			return smp, err
		}
		smp.AuxRoot = hroot

		subaux := int64(pindex) / int64(subshard)
		end := (subaux + 1) * int64(subshard)
		if end > int64(shard) {
			end = int64(shard)
		}
		hroot1, hsets1, err := merkle.BuildDataAuxProof(b[subaux*int64(subshard)*PadSize:end*PadSize], aux, h, PadSize, subshard, uint64(pindex))
		if err != nil {
			return smp, fmt.Errorf("create proof: %s", err)
		}
		if !bytes.Equal(hroot1, hroot) {
			return smp, fmt.Errorf("unequal aux root1")
		}
		smp.AuxPath = hsets1

		subaux = int64(pindex-1) / int64(subshard)
		end = (subaux + 1) * int64(subshard)
		if end > int64(shard) {
			end = int64(shard)
		}
		hroot2, hsets2, err := merkle.BuildDataAuxProof(b[subaux*int64(subshard)*PadSize:end*PadSize], aux, h, PadSize, subshard, uint64(pindex-1))
		if err != nil {
			return smp, fmt.Errorf("create proof: %s", err)
		}
		if !bytes.Equal(hroot2, hroot) {
			return smp, fmt.Errorf("unequal aux root2")
		}
		smp.AuxPrevPath = hsets2
	} else {
		hroot1, hsets1, _, err := merkle.BuildReaderProof(bytes.NewReader(b), h, h.BlockSize(), uint64(pindex))
		if err != nil {
			return smp, err
		}
		smp.AuxPath = hsets1
		smp.AuxRoot = hroot1

		hroot2, hsets2, _, err := merkle.BuildReaderProof(bytes.NewReader(b), h, h.BlockSize(), uint64(pindex-1))
		if err != nil {
			return smp, err
		}
		smp.AuxPrevPath = hsets2
		if !bytes.Equal(hroot1, hroot2) {
			return smp, fmt.Errorf("unequal aux root")
		}
	}

	err = EncodeBackward(b, prev, salt)
	if err != nil {
		return smp, err
	}
	if subshard < shard {
		hroot, aux, err := merkle.ReaderRootAux(bytes.NewReader(b), h, PadSize, subshard)
		if err != nil {
			return smp, err
		}
		smp.ReplicaRoot = hroot

		subaux := int64(pindex) / int64(subshard)
		end := (subaux + 1) * int64(subshard)
		if end > int64(shard) {
			end = int64(shard)
		}
		hroot1, hsets1, err := merkle.BuildDataAuxProof(b[subaux*int64(subshard)*PadSize:end*PadSize], aux, h, PadSize, subshard, uint64(pindex))
		if err != nil {
			return smp, fmt.Errorf("create proof: %s", err)
		}
		if !bytes.Equal(hroot1, hroot) {
			return smp, fmt.Errorf("unequal replica root1")
		}
		smp.ReplicaPath = hsets1

		subaux = int64(pindex+1) / int64(subshard)
		end = (subaux + 1) * int64(subshard)
		if end > int64(shard) {
			end = int64(shard)
		}
		hroot2, hsets2, err := merkle.BuildDataAuxProof(b[subaux*int64(subshard)*PadSize:end*PadSize], aux, h, PadSize, subshard, uint64(pindex+1))
		if err != nil {
			return smp, fmt.Errorf("create proof: %s", err)
		}
		if !bytes.Equal(hroot2, hroot) {
			return smp, fmt.Errorf("unequal replica root2")
		}
		smp.ReplicaNextPath = hsets2

		subaux = int64(pindex-1) / int64(subshard)
		end = (subaux + 1) * int64(subshard)
		if end > int64(shard) {
			end = int64(shard)
		}
		hroot3, hsets3, err := merkle.BuildDataAuxProof(b[subaux*int64(subshard)*PadSize:end*PadSize], aux, h, PadSize, subshard, uint64(pindex-1))
		if err != nil {
			return smp, fmt.Errorf("create proof: %s", err)
		}
		if !bytes.Equal(hroot3, hroot) {
			return smp, fmt.Errorf("unequal replica root3")
		}
		smp.ReplicaPrevPath = hsets3
	} else {
		hroot1, hsets1, _, err := merkle.BuildReaderProof(bytes.NewReader(b), h, h.BlockSize(), uint64(pindex))
		if err != nil {
			return smp, err
		}
		smp.ReplicaRoot = hroot1
		smp.ReplicaPath = hsets1

		hroot2, hsets2, _, err := merkle.BuildReaderProof(bytes.NewReader(b), h, h.BlockSize(), uint64(pindex+1))
		if err != nil {
			return smp, err
		}
		smp.ReplicaNextPath = hsets2
		if !bytes.Equal(hroot1, hroot2) {
			return smp, fmt.Errorf("unequal replica root next")
		}

		hroot3, hsets3, _, err := merkle.BuildReaderProof(bytes.NewReader(b), h, h.BlockSize(), uint64(pindex-1))
		if err != nil {
			return smp, err
		}
		smp.ReplicaPrevPath = hsets3
		if !bytes.Equal(hroot1, hroot3) {
			return smp, fmt.Errorf("unequal replica root prev")
		}
	}

	return smp, nil
}

func VerifyProof(smp types.SageMerkleProof, salt []byte) error {
	h := lighthash.New()
	ok := merkle.VerifyProof(h, smp.PieceRoot, smp.PiecePath, smp.Leaf)
	if !ok {
		return fmt.Errorf("wrong piece proof at: %d", smp.Leaf)
	}
	ok = merkle.VerifyProof(h, smp.AuxRoot, smp.AuxPath, smp.Leaf)
	if !ok {
		return fmt.Errorf("wrong aux proof at: %d", smp.Leaf)
	}
	ok = merkle.VerifyProof(h, smp.AuxRoot, smp.AuxPrevPath, smp.Leaf-1)
	if !ok {
		return fmt.Errorf("wrong aux prev proof at: %d", smp.Leaf-1)
	}
	ok = merkle.VerifyProof(h, smp.ReplicaRoot, smp.ReplicaPath, smp.Leaf)
	if !ok {
		return fmt.Errorf("wrong replica proof at: %d", smp.Leaf)
	}
	ok = merkle.VerifyProof(h, smp.ReplicaRoot, smp.ReplicaPrevPath, smp.Leaf-1)
	if !ok {
		return fmt.Errorf("wrong replica prev proof at: %d", smp.Leaf-1)
	}
	ok = merkle.VerifyProof(h, smp.ReplicaRoot, smp.ReplicaNextPath, smp.Leaf+1)
	if !ok {
		return fmt.Errorf("wrong replica next proof at: %d", smp.Leaf+1)
	}

	var cur, tmp bls.Fr
	cur.SetBytes(smp.PiecePath[0])
	h.Reset()
	h.Write(smp.AuxPrevPath[0])
	h.Write(salt)
	tmp.SetBytes(h.Sum(nil))
	cur.Add(&cur, &tmp)
	if !bytes.Equal(cur.Marshal(), smp.AuxPath[0]) {
		return fmt.Errorf("wrong encode at forward")
	}

	cur.SetBytes(smp.AuxPath[0])
	h.Reset()
	h.Write(smp.ReplicaNextPath[0])
	h.Write(salt)
	tmp.SetBytes(h.Sum(nil))
	cur.Add(&cur, &tmp)
	if !bytes.Equal(cur.Marshal(), smp.ReplicaPath[0]) {
		return fmt.Errorf("wrong encode at backward next")
	}

	cur.SetBytes(smp.AuxPrevPath[0])
	h.Reset()
	h.Write(smp.ReplicaPath[0])
	h.Write(salt)
	tmp.SetBytes(h.Sum(nil))
	cur.Add(&cur, &tmp)
	if !bytes.Equal(cur.Marshal(), smp.ReplicaPrevPath[0]) {
		return fmt.Errorf("wrong encode at backward prev")
	}

	return nil
}
