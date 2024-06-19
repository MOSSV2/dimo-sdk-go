package bls

import (
	"fmt"
	"math/big"

	"github.com/consensys/gnark-crypto/ecc/bls12-377/kzg"
	"github.com/consensys/gnark-crypto/hash"
)

const (
	InHashID = hash.MIMC_BW6_761
)

type MoveBatchProof struct {
	H            G1
	ClaimedValue []Fr
}

type MoveAccProof struct {
	pk       *PublicKey
	sum      G1
	rnd      Fr
	elements []Fr
	shardLen int
}

func NewMoveAccProof(pk *PublicKey, shardLen int, rnd []byte) *MoveAccProof {
	ap := &MoveAccProof{
		pk:       pk,
		elements: make([]Fr, len(pk.Pk.G1)),
		shardLen: shardLen,
	}
	ap.rnd.SetBytes(rnd)
	return ap
}

func (ap *MoveAccProof) Add(d []byte, slen int, coms []G1, offset int) (Fr, error) {
	var cy Fr
	if len(d) == 0 {
		return cy, fmt.Errorf("zero data size")
	}

	shards := Split(slen, d)
	if len(shards) > ap.shardLen {
		return cy, fmt.Errorf("invalid data size: too large")
	}

	cy = Eval(shards, ap.rnd)

	h := InHashID.New()
	buf := make([]byte, 48-len(cy.Marshal()))
	buf = append(buf, cy.Marshal()...)
	h.Write(buf)
	var fr Fr
	fr.SetBytes(h.Sum(nil))

	rndbig := new(big.Int)
	cy.BigInt(rndbig)

	coms[2].ScalarMultiplication(&coms[2], rndbig)
	coms[1].Add(&coms[1], &coms[2])
	coms[1].ScalarMultiplication(&coms[1], rndbig)
	coms[0].Add(&coms[0], &coms[1])
	ap.sum.Add(&ap.sum, &coms[0])

	for i := 0; i < len(shards); i++ {
		ap.elements[i].Add(&ap.elements[i], &shards[i])
	}

	for i := 0; i < len(shards); i++ {
		shards[i].Mul(&shards[i], &fr)
	}

	for i := 0; i < len(shards); i++ {
		ap.elements[ap.shardLen*offset+i].Add(&ap.elements[ap.shardLen*offset+i], &shards[i])
	}

	for i := 0; i < len(shards); i++ {
		shards[i].Mul(&shards[i], &fr)
	}

	for i := 0; i < len(shards); i++ {
		ap.elements[MaxShard-ap.shardLen+i].Add(&ap.elements[MaxShard-ap.shardLen+i], &shards[i])
	}

	return cy, nil
}

func (ap *MoveAccProof) GenProof() (G1, error) {
	var res G1
	srs := kzg.ProvingKey{
		G1: ap.pk.Pk.G1,
	}

	op, err := kzg.Open(ap.elements, ap.rnd, srs)
	if err != nil {
		return res, err
	}

	err = kzg.Verify(&ap.sum, &op, ap.rnd, ap.pk.Vk)
	if err != nil {
		return res, err
	}

	return op.H, nil
}
