package contract

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"math/big"
	"math/rand"
	"os"
	"testing"
	"time"

	"github.com/MOSSV2/dimo-sdk-go/lib/bls"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

func TestTransfer(t *testing.T) {
	admin := common.HexToAddress("0xC653B3b33702F3F80336274734f14c2C31885b02")
	has := BalanceOf(DevChain, admin)
	fmt.Println("admin has: ", has)

	sk, addr := makeAccount()
	valt := big.NewInt(1e18)
	valt.Mul(valt, big.NewInt(10000))

	val := big.NewInt(2e17)
	addr = common.HexToAddress("0xcf2bf532adbed038b849416b2346633c57bcc3fe")
	err := transfer(addr, val, valt)
	if err != nil {
		t.Fatal(err)
	}
	t.Fatal(sk)
}

func TestFakeReplica(t *testing.T) {
	valt := big.NewInt(1e18)
	valt.Mul(valt, big.NewInt(10))

	pkey, paddr := makeAccount()

	val := big.NewInt(1e15)
	err := transfer(paddr, val, valt)
	if err != nil {
		t.Fatal(err)
	}

	err = RegisterNode(pkey, 1, nil)
	if err != nil {
		t.Fatal(err)
	}

	ctx, cancle := context.WithTimeout(context.TODO(), 3*time.Minute)
	defer cancle()
	fi, err := NewPiece(ctx)
	if err != nil {
		t.Fatal(err)
	}

	au, err := makeAuth(big.NewInt(int64(DevChainID)), pkey)
	if err != nil {
		t.Fatal(err)
	}

	rb := []byte("test-" + time.Now().String())
	_pi := uint64(76)
	_pri := uint8(1)
	pf := make([]byte, 32)
	fmt.Println("submitreplica0: ", BalanceOf(DevChain, au.From))
	tx, err := fi.AddReplica(au, rb, _pi, _pri, pf)
	if err != nil {
		t.Fatal(err)
	}
	err = CheckTx(DevChain, tx.Hash())
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println("submitreplica1: ", BalanceOf(DevChain, au.From))
}

func TestReplicaStat(t *testing.T) {
	rsp, err := NewRSProof(context.TODO())
	if err != nil {
		t.Fatal(err)
	}
	rspi, err := rsp.GetProof(&bind.CallOpts{From: Base}, 76, 1)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(rspi)
	client, err := ethclient.Dial(DevChain)
	if err != nil {
		t.Fatal(err)
	}
	defer client.Close()
	latest, err := client.BlockNumber(context.TODO())
	if err != nil {
		t.Fatal(err)
	}

	fmt.Println(latest)
}

func TestReward(t *testing.T) {
	sk, addr := makeAccount()
	valt := big.NewInt(1e18)
	val := big.NewInt(1e14)
	err := transfer(addr, val, valt)
	if err != nil {
		t.Fatal(err)
	}
	ri, err := NewReward(context.TODO())
	if err != nil {
		t.Fatal(err)
	}
	au, err := makeAuth(big.NewInt(int64(DevChainID)), sk)
	if err != nil {
		t.Fatal(err)
	}
	for i := uint64(64); i < 3000; i += 64 {
		tx, err := ri.Mint(au, i)
		if err != nil {
			t.Fatal(i, err)
		}

		err = CheckTx(DevChain, tx.Hash())
		if err != nil {
			t.Fatal(i, err)
		}
	}
}

func TestTotalReward(t *testing.T) {
	ri, err := NewReward(context.TODO())
	if err != nil {
		t.Fatal(err)
	}
	ts, err := ri.TotalShare(&bind.CallOpts{From: Base})
	if err != nil {
		t.Fatal(err)
	}
	t.Log(ts)
	cur, err := ri.Current(&bind.CallOpts{From: Base})
	if err != nil {
		t.Fatal(err)
	}
	t.Log(cur)
	val, err := ri.Rest(&bind.CallOpts{From: Base})
	if err != nil {
		t.Fatal(err)
	}
	t.Log(val)
	rsi, err := ri.GetEReward(&bind.CallOpts{From: Base}, cur-1)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(rsi)
	t.Fatal()
}

func TestNodeCheck(t *testing.T) {
	addr := common.HexToAddress("0x5f375a47491a77f8e5bd2d9f179894188142c315")
	err := CheckNode(addr, 1)
	if err != nil {
		t.Fatal(err)
	}
	psk, paddr := makeAccount()
	err = transfer(paddr, big.NewInt(1e14), big.NewInt(2e18))
	if err != nil {
		t.Fatal(err)
	}

	err = RegisterNode(psk, 1, nil)
	if err != nil {
		t.Fatal(err)
	}
	err = CheckNode(paddr, 1)
	if err != nil {
		t.Fatal(err)
	}
	t.Fatal("")
}

func TestCheck(t *testing.T) {
	ctx := context.TODO()
	bi, err := NewBank(ctx)
	if err != nil {
		t.Fatal(err)
	}
	taddr, err := bi.Get(&bind.CallOpts{From: Base}, "token")
	if err != nil {
		t.Fatal(err)
	}
	if taddr != TokenAddr {
		t.Fatal("wrong token")
	}

	si, err := NewSpace(ctx)
	if err != nil {
		t.Fatal(err)
	}
	baddr, err := si.Bank(&bind.CallOpts{From: Base})
	if err != nil {
		t.Fatal(err)
	}
	if baddr != BankAddr {
		t.Fatal("wrong bank in space")
	}

	gi, err := NewGPU(ctx)
	if err != nil {
		t.Fatal(err)
	}
	bgaddr, err := gi.Bank(&bind.CallOpts{From: Base})
	if err != nil {
		t.Fatal(err)
	}
	if bgaddr != BankAddr {
		t.Fatal("wrong bank in gpu")
	}

	mi, err := NewModel(ctx)
	if err != nil {
		t.Fatal(err)
	}
	bmaddr, err := mi.Bank(&bind.CallOpts{From: Base})
	if err != nil {
		t.Fatal(err)
	}
	if bmaddr != BankAddr {
		t.Fatal("wrong bank in model")
	}

	ci, err := NewControl(ctx)
	if err != nil {
		t.Fatal(err)
	}
	bcaddr, err := ci.Bank(&bind.CallOpts{From: Base})
	if err != nil {
		t.Fatal(err)
	}
	if bcaddr != BankAddr {
		t.Fatal("wrong bank in control")
	}
}

func transfer(to common.Address, val, valt *big.Int) error {
	skbyte, err := os.ReadFile("/tmp/sk")
	if err != nil {
		return err
	}
	sk := string(skbyte[:64])
	privateKey, err := crypto.HexToECDSA(sk)
	if err != nil {
		return err
	}
	err = Transfer(DevChain, privateKey, to, val)
	if err != nil {
		return err
	}
	return TransferToken(DevChain, privateKey, TokenAddr, to, valt)
}

func makeAccount() (*ecdsa.PrivateKey, common.Address) {
	privk, err := crypto.GenerateKey()
	if err != nil {
		panic("generate")
	}
	publicKey := privk.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		panic("public")
	}

	addr := crypto.PubkeyToAddress(*publicKeyECDSA)
	fmt.Println("generate: ", addr)
	//skbyte := crypto.FromECDSA(privk)
	//common.Bytes2Hex(skbyte)

	return privk, addr
}

func TestBlock(t *testing.T) {
	client, err := ethclient.Dial(DevChain)
	if err != nil {
		t.Fatal(err)
	}
	defer client.Close()
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	latest, err := client.BlockNumber(ctx)
	if err != nil {
		t.Fatal(err)
	}

	fmt.Println(latest)

	bk, err := client.HeaderByNumber(ctx, big.NewInt(int64(latest-1000)))
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(bk.BaseFee, bk.BlobGasUsed, bk.Hash())

	latest, err = GetEpoch()
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(latest)
	eb, _, err := GetEpochInfo(latest)
	if err != nil {
		t.Fatal(err)
	}

	fmt.Println(eb.Uint64())

	pi, err := NewPiece(context.TODO())
	if err != nil {
		t.Fatal(err)
	}

	pcu, err := pi.Current(&bind.CallOpts{From: Base})
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(pcu)
}

func TestChoose(t *testing.T) {
	_, addr := makeAccount()

	seed := [32]byte{100}
	count := 10
	pcnt := 0
	for i := 0; i < 20; i++ {
		ci, err := choose(addr, seed, uint64(count), uint64(pcnt), uint64(i))
		if err != nil {
			t.Fatal(err)
		}
		nci := Choose(addr, seed, uint64(count), uint64(pcnt), uint64(i))
		if ci != nci {
			t.Fatal("unequal at:", i, ci, nci)
		}
	}
}

func TestOrder(t *testing.T) {
	for i := 0; i < 100; i++ {
		count := rand.Uint64() / 2
		pcnt := rand.Int63n(int64(count))
		ci, err := getOrder(uint64(count), uint64(pcnt))
		if err != nil {
			t.Fatal(err)
		}
		nci, _ := GetOrder(uint64(count), uint64(pcnt))
		if ci != nci {
			t.Fatal("unequal at:", i, ci, nci)
		}
	}
}

func TestMarshal(t *testing.T) {
	var g bls.G1
	g.ScalarMultiplicationBase(big.NewInt(10))

	gsbyte := G1InSolidity(g)
	gr, err := SolidityToG1(gsbyte)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(g.String())
	fmt.Println(gr.String())
	if !g.Equal(&gr) {
		t.Fatal("unequal")
	}
}

func TestReceipt(t *testing.T) {
	tx := common.HexToHash("0x4df93d6807c44fa617bacd1c15fa8fd57ce64f03c9ea41b7a046ea793d6a6afd")
	receipt, err := GetTransactionReceipt(DevChain, tx)
	if err != nil {
		t.Fatal(err)
	}

	fmt.Println("tx receipt: ", receipt)

	fmt.Println("tx status: ", receipt.Status)
	if receipt.Status == 1 {
		fmt.Println("tx success")
		return
	}

	for _, log := range receipt.Logs {
		fmt.Println(log)
	}

	err = AnalyzeTransactionFailure(tx)
	fmt.Println("tx fail: ", err)
}
