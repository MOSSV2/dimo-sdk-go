package contract

import (
	"crypto/ecdsa"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
)

type ContractManage struct {
	ChainID    *big.Int
	EndPoint   string
	BankAddr   common.Address
	TokenAddr  common.Address
	SyncHeight int
	local      common.Address
	sk         *ecdsa.PrivateKey
}

func NewContractManage(hexSk string, local common.Address, chainType string) (*ContractManage, error) {
	cm := &ContractManage{
		local: local,
	}

	switch chainType {
	case "op-sepolia":
		cm.EndPoint = OPSepoliaChain
		cm.ChainID = big.NewInt(int64(OPSepoliaChainID))
		cm.BankAddr = OPSepoliaBankAddr
		cm.TokenAddr = OPSepoliaTokenAddr
		cm.SyncHeight = OPSepoliaSyncHeight
	default:
		return nil, fmt.Errorf("supportted chain type: op-sepolia")
	}

	if hexSk != "" {
		sk, err := crypto.HexToECDSA(hexSk)
		if err != nil {
			return nil, err
		}
		cm.sk = sk
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
