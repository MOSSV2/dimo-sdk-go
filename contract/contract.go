package contract

import (
	"crypto/ecdsa"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

type ContractManage struct {
	Type       string
	ChainID    *big.Int
	RPC        string
	BankAddr   common.Address
	TokenAddr  common.Address
	SyncHeight int
	sk         *ecdsa.PrivateKey
}

func NewContractManage(sk *ecdsa.PrivateKey, chainType string) (*ContractManage, error) {
	cm := &ContractManage{
		Type: chainType,
		sk:   sk,
	}

	switch chainType {
	case OPSepolia:
		cm.RPC = OPSepoliaChainRPC
		cm.ChainID = big.NewInt(int64(OPSepoliaChainID))
		cm.BankAddr = OPSepoliaBankAddr
		cm.TokenAddr = OPSepoliaTokenAddr
		cm.SyncHeight = OPSepoliaSyncHeight
	case OPBNBTestnet:
		cm.RPC = OPBNBTestnetChainRPC
		cm.ChainID = big.NewInt(int64(OPBNBTestnetChainID))
		cm.BankAddr = OPBNBTestnetBankAddr
		cm.TokenAddr = OPBNBTestnetTokenAddr
		cm.SyncHeight = OPBNBTestnetSyncHeight
	case BNBTestnet:
		cm.RPC = BNBTestnetChainRPC
		cm.ChainID = big.NewInt(int64(BNBTestnetChainID))
		cm.BankAddr = BNBTestnetBankAddr
		cm.TokenAddr = BNBTestnetTokenAddr
		cm.SyncHeight = BNBTestnetSyncHeight
	default:
		return nil, fmt.Errorf("unsupportted chain type: %s, use 'bnb-testnet', 'op-sepolia' or 'opbnb-testnet'", chainType)
	}
	return cm, nil
}

func (c *ContractManage) MakeAuth() (*bind.TransactOpts, error) {
	return makeAuth(c.RPC, c.ChainID, c.sk)
}

func (c *ContractManage) GetTransactionReceipt(hash common.Hash) (*types.Receipt, error) {
	return getTransactionReceipt(c.RPC, hash)
}

func (c *ContractManage) CheckTx(txHash common.Hash) error {
	return checkTx(c.RPC, txHash)
}

func (c *ContractManage) Transfer(toAddr common.Address, value *big.Int) error {
	return transfer(c.RPC, c.sk, toAddr, value)
}

func (c *ContractManage) TransferToken(toAddr common.Address, value *big.Int) error {
	return transferToken(c.RPC, c.ChainID, c.sk, c.TokenAddr, toAddr, value)
}

func (c *ContractManage) BalanceOf(toAddr common.Address) *big.Int {
	return balanceOf(c.RPC, toAddr)
}

func (c *ContractManage) BalanceOfToken(toAddr common.Address) *big.Int {
	return balanceOfToken(c.RPC, c.TokenAddr, toAddr)
}
