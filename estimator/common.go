package estimator

import "github.com/ethereum/go-ethereum/common"

const (
	// op
	// https://optimism.drpc.org
	// https://optimism-rpc.publicnode.com
	//sep op
	// https://11155420.rpc.thirdweb.com
	// https://optimism-sepolia-rpc.publicnode.com
	L2Endpoint     = "https://optimism.drpc.org"
	L2TestEndpoint = "https://11155420.rpc.thirdweb.com"
	Decimals       = 1_000_000
)

var Base = common.HexToAddress("0x61Ea24745A3F7Bcbb67eD95B674fEcfbb331ABd0")
