package contract

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"log"
	"math/big"
	"os"
	"strconv"
	"time"

	"github.com/MOSSV2/dimo-sdk-go/build"
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
	KZGVKRoot = "3b8201b322c65a735690cf82c850b8624c29ec05400e06ba92a9aad12c37c1605812abbc9a1a11f500b3ab28b7751b52"

	RSn6k4VKRoot   = "6182580396057136035349698244013672118159269515303422592734042707265223434376"
	RSn14k7VKRoot  = "10373754662900994857464626600297287716762049534190727150581258556168809741757"
	RSn32k16VKRoot = "17745558306831082913906433185923529664704368402036214777413791674807342977788"
	RSn64k32VKRoot = "16234120671149928180196435146616388668330002147244144092564388600958703384116"

	INKZGVKRoot = "9235750480968672023981154929687398144904220020071947383737661744780443258323"
	INMulVKRoot = "20271222645096249312949571772787165775593093202630226028249238982739427154007"
	INAddVKRoot = "19807320097661814426496674775485679654967276569243947871856251314675094662138"
)

var (
	//DevChain   = "http://54.254.72.127:8501"
	//DevChainID = 222

	//http://unibasechain-scan-405529765.ap-southeast-1.elb.amazonaws.com/
	//L1Bridge   = common.HexToAddress("0xc072613dAaab3E9BcC8dDd23aE7c368DC5751984")
	//DevChain   = "https://chain.unibase.io"
	//DevChainID = 43134

	//DevChain   = "https://ethereum-holesky.publicnode.com"
	//DevChainID = 17000
	//BankAddr   = common.HexToAddress("0x6c579D5eF7846E2c6cE255Adc2E0BEF1411fEB5c")
	//TokenAddr  = common.HexToAddress("0x421BfaFCfa9370c64F65100246D02913Bc9079F4")
	//SyncHeight   = 2_391_000

	//https://sepolia-optimism.etherscan.io/
	//DevChain   = "https://11155420.rpc.thirdweb.com"

	DefaultGasLimit = 600_000
	DefaultGasPrice = 1_000_000_000

	DefaultStreamPrice  = 1e12
	DefaultReplicaPrice = 1e11 // 1TB*100 epoch cost 10
	DefaultStoreEpoch   = 301  // slight larger than minEpoch
	DelaySubmit         = 7

	DefaultSpacePrice = 1e10
	DefaultSpaceEpoch = 201

	DefaultPenalty = 1e18

	Base = common.HexToAddress("0x61Ea24745A3F7Bcbb67eD95B674fEcfbb331ABd0")
)

// op-sepolia
var (
	OPSepolia           = build.OPSepolia
	OPSepoliaExplorer   = "https://sepolia-optimism.etherscan.io/"
	OPSepoliaChainRPC   = "https://optimism-sepolia-rpc.publicnode.com"
	OPSepoliaChainID    = 11155420
	OPSepoliaBankAddr   = common.HexToAddress("0xd1B90aFa21e749f99b2d20d57B31aD96108E4CB1")
	OPSepoliaTokenAddr  = common.HexToAddress("0x96A711D4C093e4BCa8E12653014a9A15530399A1")
	OPSepoliaSyncHeight = 21_696_900
)

// opbnb-testnet
var (
	//https://opbnb-testnet-rpc.bnbchain.org
	OPBNBTestnet           = build.OPBNBTestnet
	OPBNBTestnetExplorer   = "https://opbnb-testnet.bscscan.com/"
	OPBNBTestnetChainRPC   = "https://opbnb-testnet-rpc.publicnode.com"
	OPBNBTestnetChainID    = 5611
	OPBNBTestnetBankAddr   = common.HexToAddress("0x7560B3a48952A05B989C7e2956e12a7f4b534cF5")
	OPBNBTestnetTokenAddr  = common.HexToAddress("0x1600D17EBBB5135837FCA958c1804A249716F393")
	OPBNBTestnetSyncHeight = 48_317_500
)

var (
	BNBTestnet           = build.BNBTestnet
	BNBTestnetExplorer   = "https://testnet.bscscan.com/"
	BNBTestnetChainRPC   = "https://bsc-prebsc-dataseed.bnbchain.org"
	BNBTestnetChainID    = 97
	BNBTestnetBankAddr   = common.HexToAddress("0x5903805A3a50Fab318c8650bABC71F58900EE34e")
	BNBTestnetTokenAddr  = common.HexToAddress("0x3259E10E857139a5C58Fa5Dc6C7fF525AaE661F8")
	BNBTestnetSyncHeight = 48_382_314
)

var logger = dlog.Logger("contract")

func init() {
	gasPrice := os.Getenv("GAS_PRICE")
	if gasPrice != "" {
		gasPriceInt, err := strconv.Atoi(gasPrice)
		if err != nil {
			logger.Warn("invalid gas price: ", gasPrice)
		} else {
			DefaultGasPrice = gasPriceInt
		}
	}

	gasLimit := os.Getenv("GAS_LIMIT")
	if gasLimit != "" {
		gasLimitInt, err := strconv.Atoi(gasLimit)
		if err != nil {
			logger.Warn("invalid gas limit: ", gasLimit)
		} else {
			DefaultGasLimit = gasLimitInt
		}
	}

	logger.Infof("gas price: %d", DefaultGasPrice)
	logger.Infof("gas limit: %d", DefaultGasLimit)
}

func MakeAuth(ep string, chainID int64, hexSk string) (*bind.TransactOpts, error) {
	sk, err := crypto.HexToECDSA(hexSk)
	if err != nil {
		return nil, err
	}

	return makeAuth(ep, big.NewInt(chainID), sk)
}

func CheckTx(ep string, txHash common.Hash) error {
	return checkTx(ep, txHash)
}

func makeAuth(ep string, chainID *big.Int, sk *ecdsa.PrivateKey) (*bind.TransactOpts, error) {
	auth := &bind.TransactOpts{}
	auth, err := bind.NewKeyedTransactorWithChainID(sk, chainID)
	if err != nil {
		return nil, fmt.Errorf("NewKeyedTransaction failed %s", err)
	}

	auth.Value = big.NewInt(0)
	auth.GasLimit = uint64(DefaultGasLimit)
	client, err := ethclient.Dial(ep)
	if err != nil {
		return nil, err
	}
	defer client.Close()

	header, err := client.HeaderByNumber(context.TODO(), nil)
	if err != nil {
		return nil, err
	}
	logger.Debugf("height: %d, basefee: %d, blob: %d", header.Number, header.BaseFee, header.BlobGasUsed)

	if header.BaseFee.BitLen() == 0 {
		//auth.GasTipCap = big.NewInt(int64(DefaultGasPrice))
		//auth.GasFeeCap = big.NewInt(int64(DefaultGasPrice))
		auth.GasPrice = big.NewInt(int64(DefaultGasPrice))
	} else {
		auth.GasPrice = header.BaseFee.Mul(header.BaseFee, big.NewInt(6)).Div(header.BaseFee, big.NewInt(5))
	}

	return auth, nil
}

func getTransactionReceipt(endPoint string, hash common.Hash) (*types.Receipt, error) {
	client, err := ethclient.Dial(endPoint)
	if err != nil {
		return nil, err
	}
	defer client.Close()
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	return client.TransactionReceipt(ctx, hash)
}

func getTransactionRetry(endpoint string, h common.Hash) (*types.Transaction, error) {
	retry := 0
	for retry < 10 {
		tx, err := getTransaction(endpoint, h)
		if err == nil {
			return tx, nil
		}
		retry++
		time.Sleep(time.Duration(retry) * time.Second)
	}
	return nil, fmt.Errorf("fail to get tx")
}

func getTransaction(endpoint string, h common.Hash) (*types.Transaction, error) {
	client, err := ethclient.Dial(endpoint)
	if err != nil {
		return nil, err
	}
	defer client.Close()
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	res, _, err := client.TransactionByHash(ctx, h)
	return res, err
}

func checkTx(endPoint string, txHash common.Hash) error {
	logger.Debug("check tx: ", txHash.String())
	var receipt *types.Receipt
	var err error

	t := 0
	for i := 0; i < 10; i++ {
		t = 2*t + 1
		time.Sleep(time.Duration(t) * time.Second)
		receipt, err = getTransactionReceipt(endPoint, txHash)
		if err == nil {
			break
		}
	}

	if receipt == nil {
		return fmt.Errorf("%s not packaged", txHash)
	}

	if receipt.Status == types.ReceiptStatusFailed { // 0 means fail
		for _, elog := range receipt.Logs {
			log.Printf("Log: %v\n", elog) // 打印日志信息
		}
		err = analyzeTransactionFailure(endPoint, txHash)
		if err != nil {
			logger.Warn("tx revert: ", err)
			return err
		}

		if receipt.GasUsed != receipt.CumulativeGasUsed {
			return fmt.Errorf("%s transaction exceed gas limit", txHash)
		}
		return fmt.Errorf("%s transaction mined but execution failed, check your input", txHash)
	}
	logger.Debugf("%s cost gas: %d, price: %d, blob gas: %d, price: %d", txHash.String(), receipt.GasUsed, receipt.EffectiveGasPrice, receipt.BlobGasUsed, receipt.BlobGasPrice)
	return nil
}

func analyzeTransactionFailure(endPoint string, txHash common.Hash) error {
	client, err := ethclient.Dial(endPoint)
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
	return err
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

func transfer(ep string, sk *ecdsa.PrivateKey, toAddr common.Address, value *big.Int) error {
	ctx, cancle := context.WithTimeout(context.TODO(), 1*time.Minute)
	defer cancle()
	client, err := ethclient.DialContext(ctx, ep)
	if err != nil {
		return fmt.Errorf("dail %s fail %s", ep, err)
	}
	defer client.Close()

	fromAddr := utils.ECDSAToAddr(sk)
	logger.Debugf("%s from has: %d", fromAddr, balanceOf(ep, fromAddr))
	logger.Debugf("%s to has: %d", toAddr, balanceOf(ep, toAddr))

	nonce, err := client.PendingNonceAt(ctx, fromAddr)
	if err != nil {
		return err
	}

	header, err := client.HeaderByNumber(ctx, nil)
	if err != nil {
		return err
	}

	gasLimit := uint64(23000)
	gasPrice := header.BaseFee.Add(header.BaseFee, big.NewInt(int64(DefaultGasPrice)))

	tx := types.NewTransaction(nonce, toAddr, value, gasLimit, gasPrice, nil)

	chainID, err := client.NetworkID(ctx)
	if err != nil {
		return err
	}
	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), sk)
	if err != nil {
		return err
	}

	err = client.SendTransaction(ctx, signedTx)
	if err != nil {
		return err
	}

	err = checkTx(ep, signedTx.Hash())
	if err != nil {
		return err
	}
	logger.Debugf("%s to has: %d", toAddr, balanceOf(ep, toAddr))
	return nil
}

func balanceOf(ep string, addr common.Address) *big.Int {
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

func transferToken(ep string, chainID *big.Int, sk *ecdsa.PrivateKey, tokenAddr, toaddr common.Address, val *big.Int) error {
	ctx, cancle := context.WithTimeout(context.TODO(), 1*time.Minute)
	defer cancle()
	client, err := ethclient.DialContext(ctx, ep)
	if err != nil {
		return err
	}
	defer client.Close()
	ti, err := token.NewToken(tokenAddr, client)
	if err != nil {
		return err
	}
	au, err := makeAuth(ep, chainID, sk)
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
	err = checkTx(ep, tx.Hash())
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

func balanceOfToken(ep string, tokenaddr, addr common.Address) *big.Int {
	ctx, cancle := context.WithTimeout(context.TODO(), 10*time.Second)
	defer cancle()

	client, err := ethclient.DialContext(ctx, ep)
	if err != nil {
		return big.NewInt(0)
	}
	defer client.Close()

	ti, err := token.NewToken(tokenaddr, client)
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
