package sage

import (
	"bytes"
	"testing"
	"time"

	bls "github.com/MOSSV2/dimo-sdk-go/lib/bls12377"
	"github.com/MOSSV2/dimo-sdk-go/lib/utils"
)

func TestCode(t *testing.T) {
	salt := utils.RandomBytes(32)
	prev := utils.RandomBytes(32)
	shard := 1024 * 1024
	data := make([]byte, 32*shard)
	var fr_r bls.Fr
	for i := 0; i < 1024*1024; i++ {
		fr_r.SetRandom()
		copy(data[i*32:(i+1)*32], fr_r.Marshal())
	}
	nt := time.Now()
	enc := make([]byte, len(data))
	copy(enc, data)
	err := EncodeForward(enc, prev, salt)
	if err != nil {
		t.Fatal(err)
	}
	err = EncodeBackward(enc, prev, salt)
	if err != nil {
		t.Fatal(err)
	}
	t.Log("enc cost: ", time.Since(nt))
	nt = time.Now()
	err = DecodeBackward(enc, prev, salt)
	if err != nil {
		t.Fatal(err)
	}
	err = DecodeForward(enc, prev, salt)
	if err != nil {
		t.Fatal(err)
	}
	t.Log("dec cost: ", time.Since(nt))
	if !bytes.Equal(data, enc) {
		t.Fatal("unequal")
	}
	t.Fatal()
}

func TestMerkle(t *testing.T) {
	salt := utils.RandomBytes(32)
	prev := utils.RandomBytes(32)
	shard := 1024 * 1024
	data := make([]byte, 32*shard)
	var fr_r bls.Fr
	for i := 0; i < 1024*1024; i++ {
		fr_r.SetRandom()
		copy(data[i*32:(i+1)*32], fr_r.Marshal())
	}
	nt := time.Now()
	smp, err := BuildProof(data, prev, salt, 16)
	if err != nil {
		t.Fatal(err)
	}
	t.Log("gen proof cost: ", time.Since(nt))
	nt = time.Now()
	err = VerifyProof(smp, salt)
	if err != nil {
		t.Fatal(err)
	}
	t.Log("verify cost: ", time.Since(nt))
	t.Fatal()
}
