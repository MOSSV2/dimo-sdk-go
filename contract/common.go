package contract

import (
	"context"
	"crypto/ecdsa"
	"encoding/hex"
	"fmt"
	"log"
	"math/big"
	"strings"
	"time"

	"github.com/MOSSV2/dimo-sdk-go/contract/go/token"
	dlog "github.com/MOSSV2/dimo-sdk-go/lib/log"
	"github.com/MOSSV2/dimo-sdk-go/lib/utils"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"
)

const (
	KZGVKRoot      = "3b8201b322c65a735690cf82c850b8624c29ec05400e06ba92a9aad12c37c1605812abbc9a1a11f500b3ab28b7751b52"
	RSn6k4VKRoot   = "20124409308349291856393827233330886154936240065532557381303490426078026213022"
	RSn14k7VKRoot  = "12965432607723983577277931168857749896340024793655516799335584916344574509244"
	KZGPlonkVKRoot = "13578975972813045124644274037649423183976186349085914163035805076522753635190"
	MulPlonkVKRoot = "2861925165761505728354689500849584476756532312668674953086797720607020018383"
	AddPlonkVKRoot = "20489533579041455173069175848985790272136441897595303638506439358532630167529"
)

var (
	DevChain   = "http://54.254.72.127:8501"
	DevChainID = 222
	//DevChain     = "http://unibase-sepolia-2052362516.ap-southeast-1.elb.amazonaws.com"
	//DevChainID   = 42069
	DevBlockTime = 2 // seconds/block
	//http://unibasechain-scan-405529765.ap-southeast-1.elb.amazonaws.com/

	// Epoch = 5760 blocks; 7.5 epoch/day; should be less
	DefaultGasLimit = 8_000_000
	DefaultGasPrice = 10

	DefaultReplicaPrice = 1e11 // 1TB*100 epoch cost 10
	DefaultStoreEpoch   = 101  // slight larger than minEpoch
	DelaySubmit         = 7

	DefaultSpacePrice = 1e10
	DefaultSpaceEpoch = 201

	DefaultPenalty = 1e18

	L1Bridge = common.HexToAddress("0x6C0192A83005b0a7c9Daf0b8631b9A01D779967e")

	Base = common.HexToAddress("0x61Ea24745A3F7Bcbb67eD95B674fEcfbb331ABd0")

	BankAddr  = common.HexToAddress("0xDA976D1B21103f847ABCd7f644E84d45203A5C5F")
	TokenAddr = common.HexToAddress("0x6c579D5eF7846E2c6cE255Adc2E0BEF1411fEB5c")
)

var logger = dlog.Logger("contract")

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
	auth.GasLimit = uint64(DefaultGasLimit)
	client, err := ethclient.Dial(DevChain)
	if err != nil {
		return nil, err
	}
	header, err := client.HeaderByNumber(context.TODO(), nil)
	if err != nil {
		return nil, err
	}
	logger.Debugf("height: %d, basefee: %d", header.Number, header.BaseFee)
	auth.GasPrice = header.BaseFee
	return auth, nil
}

func GetTransactionReceipt(endPoint string, hash common.Hash) (*types.Receipt, error) {
	client, err := ethclient.Dial(endPoint)
	if err != nil {
		return nil, err
	}
	defer client.Close()
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	return client.TransactionReceipt(ctx, hash)
}

func GetTransactionRetry(h common.Hash) (*types.Transaction, error) {
	retry := 0
	for retry < 10 {
		tx, err := GetTransaction(h)
		if err == nil {
			return tx, nil
		}
		retry++
		time.Sleep(time.Duration(retry) * time.Second)
	}
	return nil, fmt.Errorf("fail to get tx")
}

func GetTransaction(h common.Hash) (*types.Transaction, error) {
	client, err := ethclient.Dial(DevChain)
	if err != nil {
		return nil, err
	}
	defer client.Close()
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	res, _, err := client.TransactionByHash(ctx, h)
	return res, err
}

func CheckTx(endPoint string, txHash common.Hash) error {
	logger.Debug("check tx: ", txHash.String())
	var receipt *types.Receipt
	var err error

	t := 0
	for i := 0; i < 10; i++ {
		t = 2*t + 1
		time.Sleep(time.Duration(t) * time.Second)
		receipt, err = GetTransactionReceipt(endPoint, txHash)
		if err == nil {
			break
		}
	}

	if receipt == nil {
		return fmt.Errorf("%s not packaged", txHash)
	}

	if receipt.Status == 0 { // 0 means fail
		err = AnalyzeTransactionFailure(txHash)
		log.Println("tx revert: ", err)
		if receipt.GasUsed != receipt.CumulativeGasUsed {
			return fmt.Errorf("%s transaction exceed gas limit", txHash)
		}
		return fmt.Errorf("%s transaction mined but execution failed, check your input", txHash)
	}
	logger.Debugf("%s cost gas: %d", txHash.String(), receipt.GasUsed)
	return nil
}

func AnalyzeTransactionFailure(txHash common.Hash) error {
	client, err := ethclient.Dial(DevChain)
	if err != nil {
		return err
	}
	defer client.Close()
	tx, isPending, err := client.TransactionByHash(context.Background(), txHash)
	if err != nil {
		return fmt.Errorf("failed to get transaction by hash: %v", err)
	}

	if isPending {
		return fmt.Errorf("transaction is still pending")
	}

	// 获取交易回执
	receipt, err := client.TransactionReceipt(context.Background(), txHash)
	if err != nil {
		return fmt.Errorf("failed to get transaction receipt: %v", err)
	}

	// 获取失败的合约调用信息
	callMsg := ethereum.CallMsg{
		From:     getFrom(tx),
		To:       tx.To(),
		Gas:      tx.Gas(),
		GasPrice: tx.GasPrice(),
		Value:    tx.Value(),
		Data:     tx.Data(),
	}

	_, err = client.CallContract(context.Background(), callMsg, receipt.BlockNumber)
	if err != nil {
		errData := err.Error()
		if strings.HasPrefix(errData, "execution reverted") {
			errorMessage := parseRevertReason(errData)
			return fmt.Errorf("revert reason: %s", errorMessage)
		}
		return fmt.Errorf("contract call error: %v", err)
	} else {
		return fmt.Errorf("no revert reason found, check gas limit or other issues")
	}
}

func getFrom(tx *types.Transaction) common.Address {
	getSigner := func(trans *types.Transaction) types.Signer {
		v, _, _ := trans.RawSignatureValues()
		var isProtectedV bool
		for loop := true; loop; loop = false {
			if v.BitLen() <= 8 {
				vv := v.Uint64()
				isProtectedV = vv != 27 && vv != 28
				break
			}
			isProtectedV = true
		}
		if v.Sign() != 0 && isProtectedV {
			var chainId *big.Int
			for loop := true; loop; loop = false {
				if v.BitLen() <= 64 {
					vv := v.Uint64()
					if vv == 27 || vv == 28 {
						chainId = new(big.Int)
						break
					}
					chainId = new(big.Int).SetUint64((vv - 35) / 2)
					break
				}
				nv := new(big.Int).Sub(v, big.NewInt(35))
				chainId = nv.Div(nv, big.NewInt(2))
			}
			return types.NewEIP155Signer(chainId)
		} else {
			return types.HomesteadSigner{}
		}
	}
	signer := getSigner(tx)
	from, err := types.Sender(signer, tx)
	if err != nil {
		return common.Address{}
	}
	return from
}

func parseRevertReason(errData string) string {
	if len(errData) > 138 {
		reason := errData[138:]
		if reasonBytes, err := hex.DecodeString(reason); err == nil {
			return string(reasonBytes)
		}
	}
	return "unknown revert reason"
}

func Transfer(ep string, sk *ecdsa.PrivateKey, toAddr common.Address, value *big.Int) error {
	ctx, cancle := context.WithTimeout(context.TODO(), 1*time.Minute)
	defer cancle()
	client, err := ethclient.DialContext(ctx, ep)
	if err != nil {
		return fmt.Errorf("dail %s fail %s", ep, err)
	}
	defer client.Close()

	fromAddr := utils.ECDSAToAddr(sk)
	logger.Debugf("%s from has: %d", fromAddr, BalanceOf(ep, fromAddr))
	logger.Debugf("%s to has: %d", toAddr, BalanceOf(ep, toAddr))

	nonce, err := client.PendingNonceAt(ctx, fromAddr)
	if err != nil {
		return err
	}

	header, err := client.HeaderByNumber(ctx, nil)
	if err != nil {
		return err
	}

	gasLimit := uint64(23000)
	gasPrice := header.BaseFee

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
	logger.Debugf("%s to has: %d", toAddr, BalanceOf(ep, toAddr))
	return nil
}

func BalanceOf(ep string, addr common.Address) *big.Int {
	client, err := rpc.Dial(ep)
	if err != nil {
		return big.NewInt(0)
	}
	defer client.Close()

	ctx, cancle := context.WithTimeout(context.TODO(), 10*time.Second)
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
	ctx, cancle := context.WithTimeout(context.TODO(), 1*time.Minute)
	defer cancle()
	client, err := ethclient.DialContext(ctx, ep)
	if err != nil {
		return err
	}
	ti, err := token.NewToken(tokenAddr, client)
	if err != nil {
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
	logger.Debugf("%s from has token: %d", au.From, hasval)

	hasval, err = ti.BalanceOf(&bind.CallOpts{From: Base}, toaddr)
	if err != nil {
		return err
	}
	logger.Debugf("%s to has token: %d", toaddr, hasval)

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
	logger.Debugf("%s to has token: %d", toaddr, hasval)
	return nil
}

func BalanceOfToken(addr common.Address) *big.Int {
	ctx, cancle := context.WithTimeout(context.TODO(), 10*time.Second)
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
	logger.Debugf("%s to has token: %d", addr, hasval)
	return hasval
}
