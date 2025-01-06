package sdk

import (
	"bytes"
	"encoding/hex"
	"fmt"
	"os"
	"strconv"
	"testing"

	"github.com/MOSSV2/dimo-sdk-go/build"
	"github.com/MOSSV2/dimo-sdk-go/lib/key"
	"github.com/MOSSV2/dimo-sdk-go/lib/types"
	"github.com/MOSSV2/dimo-sdk-go/lib/utils"

	"github.com/ethereum/go-ethereum/crypto"
	"github.com/mitchellh/go-homedir"
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/mem"
)

func TestListReplica(t *testing.T) {
	res, err := ListReplicaByEdge(build.ServerURL, "0xb8d08529f80b7aa61f069530be93112373a48792", 1, 16)
	if err != nil {
		t.Fatal(err)
	}
	for i := 0; i < len(res.Replicas); i++ {
		t.Log(res.Replicas[i].Ordinal, res.Replicas[i].StoredOn)
	}

	pr, err := GetPieceOfEdge(build.ServerURL, res.Replicas[0].Piece)
	if err != nil {
		t.Fatal(err)
	}
	t.Fatal(pr)
}

func TestModel(t *testing.T) {
	skbyte, err := os.ReadFile("/tmp/sk")
	if err != nil {
		t.Fatal(err)
	}
	sks := string(skbyte[:64])
	sk, err := crypto.HexToECDSA(sks)
	if err != nil {
		t.Fatal(err)
	}

	base := ServerURL
	au, err := key.BuildAuth(sk, []byte("test"))
	if err != nil {
		t.Fatal(err)
	}
	ml, err := ListReplica(base, au, "latest")
	if err != nil {
		t.Fatal(err)
	}
	t.Fatal(ml)
}

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

func TestDecode(t *testing.T) {
	pa := "~/dimo-go/bin/stream-edge"
	pb := "~/as"

	pa, _ = homedir.Expand(pa)
	pb, _ = homedir.Expand(pb)

	fa, err := os.Open(pa)
	if err != nil {
		t.Fatal(err)
	}
	fb, err := os.Open(pb)
	if err != nil {
		t.Fatal(err)
	}

	bufa := make([]byte, 8)
	bufb := make([]byte, 8)
	fa.ReadAt(bufa, 9901120*0)
	fb.ReadAt(bufb, 9901120*0)

	if !bytes.Equal(bufa, bufb) {
		t.Fatal("not equal", hex.EncodeToString(bufa), hex.EncodeToString(bufb))
	}
}

func TestCpu(t *testing.T) {
	ci, err := cpu.Info()
	if err != nil {
		t.Fatal(err)
	}

	vms, err := mem.VirtualMemory()
	if err != nil {
		t.Fatal(err)
	}

	cpui := ci[0].ModelName + ", " + strconv.Itoa(len(ci)) + " Cores"
	fmt.Println(cpui)
	fmt.Println(utils.FormatBytes(int64(vms.Total)))
	hi := utils.GetHardwareInfo()
	fmt.Println(hi)
}
