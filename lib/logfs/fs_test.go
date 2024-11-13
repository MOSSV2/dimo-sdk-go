package simplefs

import (
	"bytes"
	"math/rand"
	"path/filepath"
	"testing"

	"github.com/MOSSV2/dimo-sdk-go/lib/kv"
	"github.com/MOSSV2/dimo-sdk-go/lib/utils"
	"github.com/mitchellh/go-homedir"
)

func TestFs(t *testing.T) {
	basedir, _ := homedir.Expand("~/test/logfs")
	mdir := filepath.Join(basedir, "kv")
	ds, err := kv.NewBadgerStore(mdir, &kv.DefaultOptions)
	if err != nil {
		t.Fatal(err)
	}

	rdir := filepath.Join(basedir, "data")
	fs, err := New(ds, rdir)
	if err != nil {
		t.Fatal(err)
	}

	key := []byte("testa")
	val := []byte("abcdefg")

	err = fs.Put(key, val)
	if err != nil {
		t.Fatal(err)
	}

	for i := 0; i < 100; i++ {
		length := rand.Int31n(256)
		nkey := utils.RandomBytes(int(length))
		length = rand.Int31n(256 * 1024)
		nval := utils.RandomBytes(int(length))
		err = fs.Put(nkey, nval)
		if err != nil {
			t.Fatal(err)
		}
	}

	nval, err := fs.Get(key)
	if err != nil {
		t.Fatal(err)
	}

	if !bytes.Equal(val, nval) {
		t.Fatal("unequal val")
	}

}
