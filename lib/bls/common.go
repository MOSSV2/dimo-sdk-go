package bls

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"hash"

	"github.com/MOSSV2/dimo-sdk-go/lib/types"

	bls "github.com/consensys/gnark-crypto/ecc/bls12-377"
	"github.com/consensys/gnark-crypto/ecc/bls12-377/fr"
	"github.com/consensys/gnark-crypto/ecc/bw6-761/fr/mimc"
)

const (
	UnPadSize = 31
	PadSize   = 32
	PadByte   = 0x01 // incase zero
	LogShard  = 20
	MinShard  = 2
	MaxShard  = 1 << LogShard
	MaxSize   = MaxShard * UnPadSize // per piece
	SlothDec  = 17
	SlothEnc  = "8c920a4fd422fb02d82062c67a1bc3cdee6ce960787878804bcf0f0f0f0f0f1"
)

type Fr = fr.Element
type G1 = bls.G1Affine
type G2 = bls.G2Affine
type GT = bls.GT
type G1Jac = bls.G1Jac
type G2Jac = bls.G2Jac

const (
	FrSize    = fr.Bytes
	G1Size    = bls.SizeOfG1AffineCompressed
	G2Size    = bls.SizeOfG2AffineCompressed
	GTSize    = bls.SizeOfGT
	ProofSize = FrSize + G1Size
)

func NewFieldHash() hash.Hash {
	return mimc.NewMiMC()
}

var _ types.ICommitment = (*Commitment)(nil)

type Commitment struct {
	Value G1
}

func NewCommitment(b []byte) (types.ICommitment, error) {
	var val G1
	if b != nil {
		_, err := val.SetBytes(b)
		if err != nil {
			return nil, err
		}
	}

	return &Commitment{val}, nil
}

func (c *Commitment) Add(ic types.ICommitment) error {
	com, ok := ic.(*Commitment)
	if !ok {
		return fmt.Errorf("wrong commitment")
	}
	c.Value.Add(&c.Value, &com.Value)
	return nil
}

func (c *Commitment) Serialize() []byte {
	cbyte := c.Value.Bytes()
	return cbyte[:]
}

func (c *Commitment) Raw() []byte {
	return c.Value.Marshal()
}

var _ types.IChallenge = (*Challenge)(nil)

type Challenge struct {
	typ     int
	Random  []byte
	Sum     G1
	Digests []G1
}

func NewChallenge(r []byte, in ...int) types.IChallenge {
	if len(r) != 32 {
		panic("invalid random length, should be 32")
	}
	buf := make([]byte, 32+8*len(in))
	copy(buf, r)
	for i := range in {
		binary.BigEndian.PutUint64(buf[32+i*8:32+(i+1)*8], uint64(in[i]))
	}

	chal := &Challenge{
		typ:     1,
		Random:  buf,
		Digests: make([]G1, 0, 1),
	}

	if len(buf) == 48 {
		chal.typ = 2
	}

	return chal
}

func NewPointChallenge(val int) types.IChallenge {
	buf := make([]byte, 8)
	binary.BigEndian.PutUint64(buf, uint64(val))
	chal := &Challenge{
		typ:     3,
		Random:  buf,
		Digests: make([]G1, 0, 1),
	}
	return chal
}

func (ch *Challenge) Type() int {
	return ch.typ
}

func (ch *Challenge) Add(ic types.ICommitment) error {
	com, ok := ic.(*Commitment)
	if !ok {
		return fmt.Errorf("wrong commitment")
	}

	ch.Sum.Add(&ch.Sum, &com.Value)
	ch.Digests = append(ch.Digests, com.Value)
	return nil
}

func (ch *Challenge) Commitment() types.ICommitment {
	return &Commitment{
		ch.Sum,
	}
}

type EncodeWitness struct {
	Root          G1
	Commits       []G1 // n
	MoveCommits   []G1 // k
	LimitCommits  []G1 // k
	Hs            []G1 // k
	ClaimedValues []Fr // k
}

func NewEncodeWitness(n, k int) *EncodeWitness {
	return &EncodeWitness{
		Commits:       make([]G1, n),
		MoveCommits:   make([]G1, k),
		LimitCommits:  make([]G1, k),
		Hs:            make([]G1, k),
		ClaimedValues: make([]Fr, k),
	}
}

func (ew *EncodeWitness) Serialize() []byte {
	var w bytes.Buffer
	enc := bls.NewEncoder(&w, bls.RawEncoding())
	toEncode := []interface{}{
		&ew.Root,
		ew.Commits,
		ew.MoveCommits,
		ew.LimitCommits,
		ew.Hs,
		ew.ClaimedValues,
	}
	for _, v := range toEncode {
		if err := enc.Encode(v); err != nil {
			panic(err)
		}
	}

	return w.Bytes()
}

func (ew *EncodeWitness) Deserialize(buf []byte) error {
	dec := bls.NewDecoder(bytes.NewReader(buf), bls.NoSubgroupChecks())
	toDecode := []interface{}{
		&ew.Root,
		ew.Commits,
		ew.MoveCommits,
		ew.LimitCommits,
		ew.Hs,
		ew.ClaimedValues,
	}
	for _, v := range toDecode {
		if err := dec.Decode(v); err != nil {
			return err
		}
	}
	return nil
}
