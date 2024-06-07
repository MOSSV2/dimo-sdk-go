package bls

import (
	"encoding/hex"
	"fmt"
	"math/big"

	"github.com/consensys/gnark-crypto/ecc/bls12-377/kzg"
)

type EpochProof struct {
	Sum          G1
	ClaimedValue Fr
	H            G1
}

type AccProof struct {
	pk       *PublicKey
	sum      G1
	elements []Fr
}

func NewAccProof(pk *PublicKey) *AccProof {
	ap := &AccProof{
		pk:       pk,
		elements: make([]Fr, 0, len(pk.Pk.G1)),
	}
	return ap
}

func (ap *AccProof) Reset(rnd []byte) {
	ap.elements = make([]Fr, 0, len(ap.pk.Pk.G1))
	ap.sum.X.SetZero()
	ap.sum.Y.SetZero()
}

func (ap *AccProof) Add(d []byte, slen int, com string, coeByte []byte) error {
	if len(d) == 0 {
		return fmt.Errorf("zero data size")
	}

	comByte, err := hex.DecodeString(com)
	if err != nil {
		return err
	}

	var g1 G1
	g1.SetBytes(comByte)

	var coeFr Fr
	coeFr.SetBytes(coeByte)
	coeBig := new(big.Int)
	coeFr.BigInt(coeBig)
	g1.ScalarMultiplication(&g1, coeBig)
	ap.sum.Add(&ap.sum, &g1)

	shards := Split(slen, d)
	if len(shards) > MaxShard {
		return fmt.Errorf("invalid data size: too large")
	}

	for i := 0; i < len(shards); i++ {
		shards[i].Mul(&shards[i], &coeFr)
	}

	if len(shards) > len(ap.elements) {
		padding := make([]Fr, len(shards)-len(ap.elements))
		ap.elements = append(ap.elements, padding...)
	}

	for i := 0; i < len(ap.elements); i++ {
		ap.elements[i].Add(&ap.elements[i], &shards[i])
	}

	return nil
}

func (ap *AccProof) GenProof(rnd []byte) (EpochProof, error) {
	if len(ap.elements) < MinShard {
		var fr Fr
		ap.elements = append(ap.elements, fr)
	}

	var fr_r Fr
	fr_r.SetBytes(rnd)

	srs := kzg.ProvingKey{
		G1: ap.pk.Pk.G1,
	}

	ep := EpochProof{}
	op, err := kzg.Open(ap.elements, fr_r, srs)
	if err != nil {
		return ep, err
	}

	err = kzg.Verify(&ap.sum, &op, fr_r, ap.pk.Vk)
	if err != nil {
		return ep, err
	}

	ep.Sum.Set(&ap.sum)
	ep.H.Set(&op.H)
	ep.ClaimedValue.Set(&op.ClaimedValue)

	return ep, nil
}
