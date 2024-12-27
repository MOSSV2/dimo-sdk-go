package test

import (
	"encoding/hex"
	"encoding/json"
	"testing"

	"github.com/MOSSV2/dimo-sdk-go/lib/utils"
	"github.com/MOSSV2/dimo-sdk-go/sdk"
	"github.com/google/uuid"
	"golang.org/x/exp/rand"
)

const url = "http://54.151.130.2:8080"

var jsonaddr = uuid.New().String()
var dataaddr = uuid.New().String()

type JsonStruct struct {
	Name string
	Age  int
}

func TestJson(t *testing.T) {
	js := JsonStruct{
		Name: "test",
		Age:  10,
	}
	jsb, err := json.Marshal(js)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(string(jsb))
	jsbi, err := json.MarshalIndent(js, "", "\t")
	if err != nil {
		t.Fatal(err)
	}
	t.Log(string(jsbi))
	t.Fatal()
}

func TestList(t *testing.T) {
	err := sdk.ListAccountHub(url)
	if err != nil {
		t.Fatal(err)
	}
	err = sdk.ListNeedleHub(url, jsonaddr)
	if err != nil {
		t.Fatal(err)
	}
}

func TestUploadJson(t *testing.T) {
	js := JsonStruct{
		Name: "test",
		Age:  10,
	}
	jsb, err := json.Marshal(js)
	if err != nil {
		t.Fatal(err)
	}
	length := rand.Int31n(16) + 16
	nkey := utils.RandomBytes(int(length))
	err = sdk.UploadHub(url, jsonaddr, hex.EncodeToString(nkey), string(jsb))
	if err != nil {
		t.Fatal(err)
	}

	for i := 0; i < 10; i++ {
		nkey := uuid.New().String()
		length = rand.Int31n(1024 * 1024)
		nval := utils.RandomBytes(int(length))

		err := sdk.UploadHub(url, jsonaddr, nkey, hex.EncodeToString(nval))
		if err != nil {
			t.Fatal(err)
		}
	}

	err = sdk.ListAccountHub(url)
	if err != nil {
		t.Fatal(err)
	}
	err = sdk.ListNeedleHub(url, jsonaddr)
	if err != nil {
		t.Fatal(err)
	}
}

func TestUploadData(t *testing.T) {

	for i := 0; i < 10; i++ {
		length := rand.Int31n(16) + 16
		nkey := utils.RandomBytes(int(length))
		length = rand.Int31n(50 * 1024 * 1024)
		nval := utils.RandomBytes(int(length))

		err := sdk.UploadHubData(url, dataaddr, hex.EncodeToString(nkey), nval)
		if err != nil {
			t.Fatal(err)
		}
	}
}
