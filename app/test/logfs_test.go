package test

import (
	"encoding/hex"
	"encoding/json"
	"testing"

	"github.com/MOSSV2/dimo-sdk-go/lib/utils"
	"github.com/MOSSV2/dimo-sdk-go/sdk"
	"golang.org/x/exp/rand"
)

const url = "http://localhost:8080"

var jsonaddr = "/testjson"
var dataaddr = "/testdata"

func init() {
	er, err := sdk.Info(url)
	if err != nil {
		jsonaddr = url + jsonaddr
		dataaddr = url + dataaddr
	}
	jsonaddr = er.Name.String() + jsonaddr
	dataaddr = er.Name.String() + dataaddr
}

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

	for i := 0; i < 0; i++ {
		length := rand.Int31n(16) + 16
		nkey := utils.RandomBytes(int(length))
		length = rand.Int31n(1024 * 1024)
		nval := utils.RandomBytes(int(length))

		err := sdk.UploadHub(url, jsonaddr, hex.EncodeToString(nkey), hex.EncodeToString(nval))
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
