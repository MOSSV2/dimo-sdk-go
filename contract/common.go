package contract

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"math/big"
	"time"

	"github.com/MOSSV2/dimo-sdk-go/contract/go/token"
	"github.com/MOSSV2/dimo-sdk-go/lib/utils"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"
)

var (
	DevChain   = "http://54.254.72.127:8501"
	DevChainID = 222

	DefaultGasPrice = 10
	ReplicaPrice    = 1e14

	Base = common.HexToAddress("0x61Ea24745A3F7Bcbb67eD95B674fEcfbb331ABd0")

	BankAddr  = common.HexToAddress("0x924465a7d33E0ed365E27907A629168BeBCEb99f")
	TokenAddr = common.HexToAddress("0x4D5c61Cb0329Ee12D52df13C28648cAf69e1aCb2")
	EpochAddr = common.HexToAddress("0x7A1aCF05Eb84e3854Af4A31a6ca471d4E01807f2")
	NodeAddr  = common.HexToAddress("0x07b59a9F8045Dc7E8F213b443ee94BcB8B0f2841")
	BLSAddr   = common.HexToAddress("0xa264f9Da18E7290FDe59Fe6342bE21d9447E3c41")
	RoundAddr = common.HexToAddress("0xbA8dc3116DAAce8D49C45d9ba58A84f80091775D")
	ProofAddr = common.HexToAddress("0x99D9eBfEC77d5dd7fD8AcA81b8338EF9853CeEFA")
	FileAddr  = common.HexToAddress("0x9aeF84fC6e581471284c5b10f063AB2DEd05d2F6")
)

func MakeAuth(chainID *big.Int, hexSk string) (*bind.TransactOpts, error) {
	sk, err := crypto.HexToECDSA(hexSk)
	if err != nil {
		return nil, err
	}

	return makeAuth(chainID, sk)
}

func makeAuth(chainID *big.Int, sk *ecdsa.PrivateKey) (*bind.TransactOpts, error) {
	auth := &bind.TransactOpts{}
	auth, err := bind.NewKeyedTransactorWithChainID(sk, chainID)
	if err != nil {
		return nil, fmt.Errorf("NewKeyedTransaction failed %s", err)
	}

	auth.Value = big.NewInt(0)
	auth.GasPrice = big.NewInt(int64(DefaultGasPrice))
	return auth, nil
}

func GetTransactionReceipt(endPoint string, hash common.Hash) (*types.Receipt, error) {
	client, err := ethclient.Dial(endPoint)
	if err != nil {
		return nil, err
	}
	defer client.Close()
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()
	return client.TransactionReceipt(ctx, hash)
}

func CheckTx(endPoint string, txHash common.Hash) error {
	fmt.Println("check tx: ", txHash.String())
	var receipt *types.Receipt
	var err error

	t := 3
	for i := 0; i < 10; i++ {
		if i != 0 {
			t = 3 * i
		}
		time.Sleep(time.Duration(t) * time.Second)
		receipt, err = GetTransactionReceipt(endPoint, txHash)
		if err == nil {
			break
		}
	}

	if receipt == nil {
		return fmt.Errorf("%s not packaged", txHash)
	}
	fmt.Println("gas cost: ", receipt.GasUsed)
	if receipt.Status == 0 { // 0 means fail
		if receipt.GasUsed != receipt.CumulativeGasUsed {
			return fmt.Errorf("%s transaction exceed gas limit", txHash)
		}
		return fmt.Errorf("%s transaction mined but execution failed, check your input", txHash)
	}
	return nil
}

func Transfer(ep string, sk *ecdsa.PrivateKey, toAddr common.Address, value *big.Int) error {
	ctx, cancle := context.WithTimeout(context.TODO(), 10*time.Second)
	defer cancle()
	client, err := ethclient.DialContext(ctx, ep)
	if err != nil {
		return fmt.Errorf("dail %s fail %s", ep, err)
	}
	defer client.Close()

	fromAddr := utils.ECDSAToAddr(sk)
	fmt.Println("from has: ", BalanceOf(ep, fromAddr))
	fmt.Println("to has: ", BalanceOf(ep, toAddr))

	nonce, err := client.PendingNonceAt(ctx, fromAddr)
	if err != nil {
		return err
	}

	gasPrice, err := client.SuggestGasPrice(ctx)
	if err != nil {
		return err
	}

	gasLimit := uint64(23000)

	tx := types.NewTransaction(nonce, toAddr, value, gasLimit, gasPrice, nil)

	chainID, err := client.NetworkID(ctx)
	if err != nil {
		chainID = big.NewInt(int64(DevChainID))
	}
	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), sk)
	if err != nil {
		return err
	}

	err = client.SendTransaction(ctx, signedTx)
	if err != nil {
		return err
	}

	err = CheckTx(ep, signedTx.Hash())
	if err != nil {
		return err
	}
	fmt.Println("to has: ", BalanceOf(ep, toAddr))
	return nil
}

func BalanceOf(ep string, addr common.Address) *big.Int {
	client, err := rpc.Dial(ep)
	if err != nil {
		return big.NewInt(0)
	}
	defer client.Close()

	ctx, cancle := context.WithTimeout(context.TODO(), 1*time.Second)
	defer cancle()

	var result string
	err = client.CallContext(ctx, &result, "eth_getBalance", addr.String(), "latest")
	if err != nil {
		return big.NewInt(0)
	}

	val, _ := new(big.Int).SetString(result[2:], 16)
	return val
}

func TransferToken(ep string, sk *ecdsa.PrivateKey, tokenAddr, toaddr common.Address, val *big.Int) error {
	ctx, cancle := context.WithTimeout(context.TODO(), 30*time.Second)
	defer cancle()
	client, err := ethclient.DialContext(ctx, ep)
	if err != nil {
		return err
	}
	ti, err := token.NewToken(tokenAddr, client)
	if err != nil {
		fmt.Println(err)
		return err
	}
	au, err := makeAuth(big.NewInt(int64(DevChainID)), sk)
	if err != nil {
		return err
	}
	hasval, err := ti.BalanceOf(&bind.CallOpts{From: Base}, au.From)
	if err != nil {
		return err
	}
	fmt.Println("from has token: ", hasval)

	hasval, err = ti.BalanceOf(&bind.CallOpts{From: Base}, toaddr)
	if err != nil {
		return err
	}
	fmt.Println("to has token: ", hasval)

	tx, err := ti.Transfer(au, toaddr, val)
	if err != nil {
		return err
	}
	err = CheckTx(ep, tx.Hash())
	if err != nil {
		return err
	}

	hasval, err = ti.BalanceOf(&bind.CallOpts{From: Base}, toaddr)
	if err != nil {
		return err
	}
	fmt.Println("to has token: ", hasval)
	return nil
}

func BalanceOfToken(addr common.Address) *big.Int {
	ctx, cancle := context.WithTimeout(context.TODO(), 1*time.Second)
	defer cancle()

	client, err := ethclient.DialContext(ctx, DevChain)
	if err != nil {
		return big.NewInt(0)
	}
	ti, err := token.NewToken(TokenAddr, client)
	if err != nil {
		return big.NewInt(0)
	}

	hasval, err := ti.BalanceOf(&bind.CallOpts{From: addr}, addr)
	if err != nil {
		return big.NewInt(0)
	}
	fmt.Println(addr, " has token: ", hasval)
	return hasval
}
