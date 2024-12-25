package param

type paramInfo struct {
	hash string
	size int64
}

var paramMap map[string]paramInfo

func init() {
	paramMap = make(map[string]paramInfo)
	// kzg
	paramMap["kzgsrs-bls12_377-kzg"] = paramInfo{
		size: 3221274340,
		hash: "ecabcd204af80676d39594bb2a1fa7b649f7115a13ac143f5b378f5b3d6d49e2",
	}
	paramMap["kzgsrs-bls12_377-kzg-vk"] = paramInfo{
		size: 48624,
		hash: "3f042498c2fe3cb2bc15816fccd64868c638a3d0a90a81eb05d515b727ddfc37",
	}

	// bn254
	paramMap["kzgsrs-bn254-canonical"] = paramInfo{
		size: 2147517956,
		hash: "4c73afc8e6c59579dd18bf68bf0312ebc975f8a7067b852f6fb616d4acc1d65f",
	}
	paramMap["kzgsrs-bn254-lagrange-33554432"] = paramInfo{
		size: 2147517764,
		hash: "066dd595dd4f89a00e0d67a49b71843b7c5feed7c13a50eee4ad581b9119c5ba",
	}

	// bw6_761
	paramMap["kzgsrs-bw6_761-canonical"] = paramInfo{
		size: 201472900,
		hash: "4157b57e17e08d758efd7141de4340bfbeabf0f2f305e4ee3e24ae605fc8a161",
	}
	paramMap["kzgsrs-bw6_761-lagrange-128"] = paramInfo{
		size: 170308,
		hash: "1d3ba9e4882d3a3ff3e95c4afdc42ab8ef31ae866bc1fbdb3ef8ac94605a3ca8",
	}
	paramMap["kzgsrs-bw6_761-lagrange-256"] = paramInfo{
		size: 194884,
		hash: "0a0e4523857f441350f681cc52ce3ef689ddb93d1e391fe37fe6406108c715c7",
	}
	paramMap["kzgsrs-bw6_761-lagrange-512"] = paramInfo{
		size: 244036,
		hash: "f88197e258b512a2f1dab6bd7dc47b24b6f64cc457da16473f5aafefa8efbfbb",
	}
	paramMap["kzgsrs-bw6_761-lagrange-1024"] = paramInfo{
		size: 342340,
		hash: "32ae9ce229810bcd60f61f47d245f97688eaddeb63f317c5d0e04958788fa9ef",
	}
	paramMap["kzgsrs-bw6_761-lagrange-2048"] = paramInfo{
		size: 538948,
		hash: "246ec7b7592c10f87501c7f222f238f9290efd1d2eb6b93dfa4526f8c0afa2c7",
	}
	paramMap["kzgsrs-bw6_761-lagrange-4096"] = paramInfo{
		size: 932164,
		hash: "51572f8de38d9ab4a32281eb2adc3dabe98df905d3113046261b492c81a5f5a6",
	}
	paramMap["kzgsrs-bw6_761-lagrange-8192"] = paramInfo{
		size: 1718596,
		hash: "3475a7e697e0f45fe127ebb893243b4f1a16766c7055b90021eb9e4e41212338",
	}
	paramMap["kzgsrs-bw6_761-lagrange-16384"] = paramInfo{
		size: 3291460,
		hash: "568b48bfe74af7ef61133c4a1f068de6c0284375038728e7da44ab0cf05b17bf",
	}
	paramMap["kzgsrs-bw6_761-lagrange-32768"] = paramInfo{
		size: 6437188,
		hash: "597aba5a437d7e44bdea72e7cb661f02d5ecdb836bdfd1a6aac0e155db98e0fa",
	}
	paramMap["kzgsrs-bw6_761-lagrange-65536"] = paramInfo{
		size: 12728644,
		hash: "845ca2dfff7364ac77412ff24bcd09b81b7678d604e2bfc2359dfc1c0231b641",
	}
	paramMap["kzgsrs-bw6_761-lagrange-131072"] = paramInfo{
		size: 25311556,
		hash: "eec6e9b551f4597c991630bd53159dd22707a86af13efb6f3fd6853186995025",
	}
	paramMap["kzgsrs-bw6_761-lagrange-262144"] = paramInfo{
		size: 50477380,
		hash: "1a35a5dead576f845004e68a9ad8f3279d721196733862a30d129612e38ae1a7",
	}
	paramMap["kzgsrs-bw6_761-lagrange-524288"] = paramInfo{
		size: 100809028,
		hash: "cbc2b37fa568a0dbdab45a1d1aa356ce7891294bc9362dab492e2930db284064",
	}
	paramMap["kzgsrs-bw6_761-lagrange-524288"] = paramInfo{
		size: 201472324,
		hash: "7f7510fe1822c30b1412efa39297d31374c9a2bdd2634cea025f46f6f04258a8",
	}
}
