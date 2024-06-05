package contract

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"log"
	"math/big"
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
	sk, addr := makeAccount()
	val := big.NewInt(1e18)
	val.Mul(val, big.NewInt(100000))
	addr = common.HexToAddress("0xcf2bf532adbed038b849416b2346633c57bcc3fe")
	err := transfer(addr, val)
	if err != nil {
		t.Fatal(err)
	}
	t.Fatal(sk)
}

func TestNodeCheck(t *testing.T) {
	addr := common.HexToAddress("0x5f375a47491a77f8e5bd2d9f179894188142c315")
	err := CheckNode(addr, 1)
	if err != nil {
		t.Fatal(err)
	}
	psk, paddr := makeAccount()
	err = transfer(paddr, big.NewInt(2e18))
	if err != nil {
		t.Fatal(err)
	}

	err = RegisterNode(psk, 1)
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

func transfer(to common.Address, val *big.Int) error {
	skbyte, err := os.ReadFile("/tmp/sk")
	if err != nil {
		return err
	}
	sk := string(skbyte[:64])
	privateKey, err := crypto.HexToECDSA(sk)
	if err != nil {
		return err
	}
	err = Transfer(DevChain, privateKey, to, big.NewInt(1e15))
	if err != nil {
		return err
	}
	return TransferToken(DevChain, privateKey, TokenAddr, to, val)
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
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()
	latest, err := client.BlockNumber(ctx)
	if err != nil {
		t.Fatal(err)
	}

	log.Println(latest)

	for i := latest; i > latest-5; i-- {
		bk, err := client.BlockByNumber(ctx, big.NewInt(int64(i)))
		if err != nil {
			t.Fatal(err)
		}
		t.Log(bk)
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
