package contract

import (
	"bytes"
	"context"
	"crypto/ecdsa"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"math/big"
	"os"
	"testing"

	"github.com/MOSSV2/dimo-sdk-go/contract/go/bls"
	bls12377 "github.com/MOSSV2/dimo-sdk-go/lib/bls12377"
	"github.com/MOSSV2/dimo-sdk-go/lib/lighthash"
	"github.com/MOSSV2/dimo-sdk-go/lib/merkle"
	"github.com/MOSSV2/dimo-sdk-go/lib/types"

	"github.com/consensys/gnark-crypto/ecc/bls12-377/kzg"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

func TestTransfer(t *testing.T) {
	sk, addr := makeAccount()
	val := big.NewInt(1e18)
	val.Mul(val, big.NewInt(10000))
	addr = common.HexToAddress("0xcf2bf532adbed038b849416b2346633c57bcc3fe")
	err := transfer(addr, val)
	if err != nil {
		t.Fatal(err)
	}
	t.Fatal(sk)
}

func TestCheck(t *testing.T) {
	addr := common.HexToAddress("0x7552c59444AF49FD16233423075508EfA14c3A5E")
	err := CheckNode(addr, 3)
	if err != nil {
		t.Fatal(err)
	}
	t.Fatal("")
}

func TestGet(t *testing.T) {
	num := uint64(0)
	for i := uint64(0); i < 12; i++ {
		_order, err := GetOrder(num)
		if err != nil {
			t.Fatal(err)
		}
		fmt.Println(num, _order)
		num = num*8 + i + 1
	}
}

func TestGetReplica(t *testing.T) {
	for i := uint64(0); i < 32; i++ {
		fmt.Println(i)
		ri, err := GetReplica(i)
		if err != nil {
			fmt.Println(err)
		}

		if ri.Expire > 0 {
			sn, err := SolByteToString(ri.Name)
			if err == nil {
				fmt.Println(sn)
			}
			fmt.Println(len(ri.Name))
		}
	}
	t.Fatal("")
}

func TestBLS(t *testing.T) {
	client, err := ethclient.DialContext(context.TODO(), DevChain)
	if err != nil {
		t.Fatal(err)
	}

	_, paddr := makeAccount()

	bi, err := bls.NewBLS(BLSAddr, client)
	if err != nil {
		t.Fatal(err)
	}

	pk, err := kzg.NewSRS(10, big.NewInt(12345678))
	if err != nil {
		t.Fatal(err)
	}
	_com := make([][]byte, 2)

	for i := 0; i < 2; i++ {
		xl := G1ToSolByte(pk.Pk.G1[i])
		ng1, err := SolByteToG1(xl)
		if err != nil {
			t.Fatal(err)
		}
		if !ng1.Equal(&pk.Pk.G1[i]) {
			t.Fatal("fail")
		}
		_com[i] = xl
		fmt.Printf("0x%s\n", common.Bytes2Hex(xl))
	}
	pk.Pk.G1[0].Add(&pk.Pk.G1[0], &pk.Pk.G1[1])
	for i := 0; i < 1; i++ {
		xl := G1ToSolByte(pk.Pk.G1[i])
		fmt.Printf("0x%s\n", common.Bytes2Hex(xl))
	}
	res, err := bi.Add(&bind.CallOpts{From: paddr}, _com)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("0x%s\n", common.Bytes2Hex(res))
}

func TestFile(t *testing.T) {
	psk, paddr := makeAccount()
	err := transfer(paddr, big.NewInt(1e18))
	if err != nil {
		t.Fatal(err)
	}

	err = RegisterNode(psk, 1)
	if err != nil {
		t.Fatal(err)
	}

	sk, addr := makeAccount()
	err = transfer(addr, big.NewInt(1e18))
	if err != nil {
		t.Fatal(err)
	}

	_fn, _fc := generateFile("test1", 1)

	fmt.Println("add file: ", _fn)
	err = AddFile(sk, _fn, _fc)
	if err != nil {
		t.Fatal(err)
	}

	fb, err := GetFile(_fn)
	if err != nil {
		t.Fatal(err)
	}

	if fb.Size == 0 {
		t.Fatal("no such file")
	}

	fmt.Println(fb)
	fmt.Println("add piece to: ", _fn)
	err = AddPiece(sk, _fn, _fc.Pieces)
	if err != nil {
		t.Fatal(err)
	}

	fb, err = GetFile(_fn)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(fb)

	proxysk, proxyaddr := makeAccount()
	err = transfer(proxyaddr, big.NewInt(1e18))
	if err != nil {
		t.Fatal(err)
	}

	err = SetFileProxy(sk, _fn, proxyaddr)
	if err != nil {
		t.Fatal(err)
	}

	for i := 0; i < int(fb.Count); i++ {
		pi, err := GetPiece(_fn, uint64(i))
		if err != nil {
			t.Fatal(err)
		}

		replica := make([]byte, len(pi.Name))
		copy(replica, pi.Name)

		hash := crypto.Keccak256(FileAddr.Bytes(), pi.Name, pi.Name)
		sig, err := crypto.Sign(hash, psk)
		if err != nil {
			t.Fatal(err)
		}
		fmt.Println("add replica: ", common.Bytes2Hex(replica))
		err = AddReplica(proxysk, _fn, uint64(i), common.Bytes2Hex(replica), sig)
		if err != nil {
			t.Fatal(err)
		}

		npi, err := GetPiece(_fn, uint64(i))
		if err != nil {
			t.Fatal(err)
		}

		rindex := npi.Replica[len(pi.Replica)]
		ri, err := GetReplica(uint64(rindex))
		if err != nil {
			t.Fatal(err)
		}

		if !bytes.Equal(ri.Name, pi.Name) {
			t.Fatal("unequal")
		}

		if ri.StoredOn != paddr {
			t.Fatal("unequal location")
		}
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
	err = Transfer(DevChain, privateKey, to, big.NewInt(1e17))
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

func generateFile(val string, count int) (_fn string, fc types.FileCore) {
	res := types.FileCore{
		Pieces: make([]string, 0, 1),
	}

	ph := sha256.New()
	for i := 0; i < count; i++ {
		data := []byte(val)
		data = append(data, byte(i))

		ph.Write(data)

		b := bls12377.Pad(data)
		root, err := merkle.ReaderRoot(bytes.NewReader(b), lighthash.New(), bls12377.PadSize)
		if err != nil {
			panic(err)
		}
		res.Pieces = append(res.Pieces, hex.EncodeToString(root))
		res.Size += 101
	}

	name := ph.Sum(nil)
	fc.Hash = hex.EncodeToString(name)
	return hex.EncodeToString(name), res
}
