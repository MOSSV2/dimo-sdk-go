package test

import (
	"encoding/hex"
	"testing"

	"github.com/MOSSV2/dimo-sdk-go/lib/types"
	"github.com/MOSSV2/dimo-sdk-go/lib/utils"
	"github.com/MOSSV2/dimo-sdk-go/sdk"
	"golang.org/x/exp/rand"
)

const url = "http://127.0.0.1:8086"

func TestUploadJson(t *testing.T) {

	for i := 0; i < 100; i++ {
		length := rand.Int31n(16) + 16
		nkey := utils.RandomBytes(int(length))
		length = rand.Int31n(1024 * 1024)
		nval := utils.RandomBytes(int(length))

		mm := types.MemeStruct{
			Owner:   "commonjson",
			ID:      hex.EncodeToString(nkey),
			Message: hex.EncodeToString(nval),
		}

		err := sdk.UploadHub(url, mm)
		if err != nil {
			t.Fatal(err)
		}
	}
}

func TestUploadData(t *testing.T) {

	for i := 0; i < 100; i++ {
		length := rand.Int31n(16) + 16
		nkey := utils.RandomBytes(int(length))
		length = rand.Int31n(1024 * 1024)
		nval := utils.RandomBytes(int(length))

		err := sdk.UploadHubData(url, "common", hex.EncodeToString(nkey), nval)
		if err != nil {
			t.Fatal(err)
		}
	}
}
