package build

import "os"

const (
	ServerURL    = "http://54.251.11.180:8080"
	OPSepolia    = "op-sepolia"
	OPBNBTestnet = "opbnb-testnet"
)

func CheckChain() string {
	ct := os.Getenv("CHAIN_TYPE")
	if ct == "" {
		panic("please set env 'CHAIN_TYPE' to 'op-sepolia' or 'opbnb-testnet'")
	}
	switch ct {
	case OPSepolia, OPBNBTestnet:
		return ct
	default:
		panic("please set env 'CHAIN_TYPE' to 'op-sepolia' or 'opbnb-testnet'")
	}
}
