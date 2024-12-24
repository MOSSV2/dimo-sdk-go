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
	ChainID    *big.Int
	EndPoint   string
	BankAddr   common.Address
	TokenAddr  common.Address
	SyncHeight int
	sk         *ecdsa.PrivateKey
}

func NewContractManage(sk *ecdsa.PrivateKey, chainType string) (*ContractManage, error) {
	cm := &ContractManage{
		sk: sk,
	}

	switch chainType {
	case "op-sepolia":
		cm.EndPoint = OPSepoliaChain
		cm.ChainID = big.NewInt(int64(OPSepoliaChainID))
		cm.BankAddr = OPSepoliaBankAddr
		cm.TokenAddr = OPSepoliaTokenAddr
		cm.SyncHeight = OPSepoliaSyncHeight
	default:
		return nil, fmt.Errorf("unsupportted chain type: %s, use 'op-sepolia'", chainType)
	}
	return cm, nil
}

func (c *ContractManage) MakeAuth() (*bind.TransactOpts, error) {
	return makeAuth(c.EndPoint, c.ChainID, c.sk)
}

func (c *ContractManage) GetTransactionReceipt(hash common.Hash) (*types.Receipt, error) {
	return getTransactionReceipt(c.EndPoint, hash)
}

func (c *ContractManage) CheckTx(txHash common.Hash) error {
	return checkTx(c.EndPoint, txHash)
}

func (c *ContractManage) Transfer(toAddr common.Address, value *big.Int) error {
	return transfer(c.EndPoint, c.sk, toAddr, value)
}

func (c *ContractManage) TransferToken(toAddr common.Address, value *big.Int) error {
	return transferToken(c.EndPoint, c.ChainID, c.sk, c.TokenAddr, toAddr, value)
}

func (c *ContractManage) BalanceOf(toAddr common.Address) *big.Int {
	return balanceOf(c.EndPoint, toAddr)
}

func (c *ContractManage) BalanceOfToken(toAddr common.Address) *big.Int {
	return balanceOfToken(c.EndPoint, c.TokenAddr, toAddr)
}
