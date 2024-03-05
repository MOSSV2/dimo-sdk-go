package bls12377

import (
	"fmt"
	"math/big"
)

func SlothEncode(in, v []byte) error {
	cnt := 1 + (len(in)-1)/PadSize
	if len(in) != PadSize*cnt {
		return fmt.Errorf("invalid in length, should align to %d", PadSize)
	}

	var tmp, fr_v Fr
	fr_v.SetBytes(v)
	big_inv, _ := new(big.Int).SetString(SlothEnc, 16)
	for i := 0; i < cnt; i++ {
		tmp.SetBytes(in[PadSize*i : PadSize*(i+1)])
		tmp.Add(&tmp, &fr_v)
		tmp.Exp(tmp, big_inv)
		copy(in[PadSize*i:PadSize*(i+1)], tmp.Marshal())
	}
	return nil
}

func SlothDecode(in, v []byte) error {
	cnt := 1 + (len(in)-1)/PadSize
	if len(in) != PadSize*cnt {
		return fmt.Errorf("invalid in length, should align to %d", PadSize)
	}
	var tmp, fr_v Fr
	fr_v.SetBytes(v)
	big_e := big.NewInt(SlothDec)
	for i := 0; i < cnt; i++ {
		tmp.SetBytes(in[PadSize*i : PadSize*(i+1)])
		tmp.Exp(tmp, big_e)
		tmp.Sub(&tmp, &fr_v)
		copy(in[PadSize*i:PadSize*(i+1)], tmp.Marshal())
	}
	return nil
}
