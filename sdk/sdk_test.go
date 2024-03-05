package sdk

import (
	"encoding/hex"
	"testing"

	"github.com/MOSSV2/dimo-sdk-go/lib/types"
	"github.com/mitchellh/go-homedir"
)

func TestList(t *testing.T) {
	dir, err := homedir.Expand("~/relay")
	if err != nil {
		t.Fatal(err)
	}

	mrm, err := ListLocalDir(dir)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(mrm.Name)
	for k, v := range mrm.Files {
		if len(v) == 64 {
			_, err := hex.DecodeString(v)
			if err == nil {
				t.Log(k, v)
			}
		}
	}

	mrm.Name = "relayd1"

	rootdir, err := homedir.Expand("~/test/test1")
	if err != nil {
		t.Fatal(err)
	}
	err = DownloadModel("", rootdir, types.Auth{}, mrm, nil)
	if err != nil {
		t.Fatal(err)
	}

	t.Fatal()
}
