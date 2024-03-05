package sage

import (
	"fmt"

	bls "github.com/MOSSV2/dimo-sdk-go/lib/bls12377"
	"github.com/MOSSV2/dimo-sdk-go/lib/lighthash"
)

const (
	PadSize  = bls.PadSize
	MaxCount = 32
)

func EncodeForward(in []byte, prev, salt []byte) error {
	cnt := 1 + (len(in)-1)/PadSize
	if len(in) != PadSize*cnt {
		return fmt.Errorf("invalid in length, should align to %d", PadSize)
	}
	h := lighthash.New()
	var cur, tmp bls.Fr
	for i := 0; i < cnt; i++ {
		cur.SetBytes(in[PadSize*i : PadSize*(i+1)])
		h.Reset()
		h.Write(prev)
		h.Write(salt)
		tmp.SetBytes(h.Sum(nil))
		cur.Add(&cur, &tmp)

		prev = cur.Marshal()
		copy(in[PadSize*i:PadSize*(i+1)], prev)
	}

	return nil
}

func EncodeBackward(in []byte, prev, salt []byte) error {
	cnt := 1 + (len(in)-1)/PadSize
	if len(in) != PadSize*cnt {
		return fmt.Errorf("invalid in length, should align to %d", PadSize)
	}
	h := lighthash.New()
	var cur, tmp bls.Fr
	for i := cnt - 1; i >= 0; i-- {
		cur.SetBytes(in[PadSize*i : PadSize*(i+1)])
		h.Reset()
		h.Write(prev)
		h.Write(salt)
		tmp.SetBytes(h.Sum(nil))
		cur.Add(&cur, &tmp)

		prev = cur.Marshal()
		copy(in[PadSize*i:PadSize*(i+1)], prev)
	}

	return nil
}

func DecodeForward(in []byte, prev, salt []byte) error {
	cnt := 1 + (len(in)-1)/PadSize
	if len(in) != PadSize*cnt {
		return fmt.Errorf("invalid in length, should align to %d", PadSize)
	}
	h := lighthash.New()
	var cur, tmp bls.Fr
	for i := 0; i < cnt; i++ {
		cur.SetBytes(in[PadSize*i : PadSize*(i+1)])
		h.Reset()
		h.Write(prev)
		h.Write(salt)
		tmp.SetBytes(h.Sum(nil))
		prev = cur.Marshal()

		cur.Sub(&cur, &tmp)
		copy(in[PadSize*i:PadSize*(i+1)], cur.Marshal())
	}

	return nil
}

func DecodeBackward(in []byte, prev, salt []byte) error {
	cnt := 1 + (len(in)-1)/PadSize
	if len(in) != PadSize*cnt {
		return fmt.Errorf("invalid in length, should align to %d", PadSize)
	}
	h := lighthash.New()
	var cur, tmp bls.Fr
	for i := cnt - 1; i >= 0; i-- {
		cur.SetBytes(in[PadSize*i : PadSize*(i+1)])
		h.Reset()
		h.Write(prev)
		h.Write(salt)
		tmp.SetBytes(h.Sum(nil))
		prev = cur.Marshal()

		cur.Sub(&cur, &tmp)
		copy(in[PadSize*i:PadSize*(i+1)], cur.Marshal())
	}

	return nil
}
