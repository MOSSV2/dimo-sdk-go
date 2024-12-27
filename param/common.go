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
		hash: "5860ef6cb95d0827a9962fcea3994f044632051c5282dc55f6b28d3b57b743e7",
	}
	paramMap["kzgsrs-bw6_761-lagrange-128"] = paramInfo{
		size: 170308,
		hash: "decd2c7251d3c76d11fbc4d68cda5b8862aef4b380c64d7a84704f24f4e5441c",
	}
	paramMap["kzgsrs-bw6_761-lagrange-256"] = paramInfo{
		size: 194884,
		hash: "1b9543d4f592a9a0e27aec898b8d1c7bba27ee1c3d88662dfcc34c939ed40bf1",
	}
	paramMap["kzgsrs-bw6_761-lagrange-512"] = paramInfo{
		size: 244036,
		hash: "4d4f86a85a12dc15b7e405d37386945190a9be9fd6474c415d889cbba99e83a2",
	}
	paramMap["kzgsrs-bw6_761-lagrange-1024"] = paramInfo{
		size: 342340,
		hash: "f77ae21934d748ed2a8e3915e95ab13f2ab45955b16537632bf35d636b8d41da",
	}
	paramMap["kzgsrs-bw6_761-lagrange-2048"] = paramInfo{
		size: 538948,
		hash: "3b4d21006130451940f9ab8385aca5ed44a4d0bf7dd680eac3cbb548c61657ee",
	}
	paramMap["kzgsrs-bw6_761-lagrange-4096"] = paramInfo{
		size: 932164,
		hash: "a884f61fed2d7d151e863ef037c1fb5a8a8e57d143d5786f15bcef852e6ad80d",
	}
	paramMap["kzgsrs-bw6_761-lagrange-8192"] = paramInfo{
		size: 1718596,
		hash: "32717e7d9bae15f02d87c35975d78f26175f8f408097c7434b08e05bdf0a3f5d",
	}
	paramMap["kzgsrs-bw6_761-lagrange-16384"] = paramInfo{
		size: 3291460,
		hash: "f1800b87fea099a203140d2ba0bcb5a95a6b2cb7ddfc66e9c8852cdae6d8e986",
	}
	paramMap["kzgsrs-bw6_761-lagrange-32768"] = paramInfo{
		size: 6437188,
		hash: "dd98e9a241aaffd7ae1bb65dbaaa5dcfe3323f8ea55c4964595b073b10119742",
	}
	paramMap["kzgsrs-bw6_761-lagrange-65536"] = paramInfo{
		size: 12728644,
		hash: "d97e2c6cf834be8844b0e778f2a28e4e3fda8066823c390258de733e84c5f4bd",
	}
	paramMap["kzgsrs-bw6_761-lagrange-131072"] = paramInfo{
		size: 25311556,
		hash: "52624f500d30b6ca4661d859567cd2e23fb01fc54fd4f9d722c72d92a252cccd",
	}
	paramMap["kzgsrs-bw6_761-lagrange-262144"] = paramInfo{
		size: 50477380,
		hash: "e35ae547ab04b5838cabd28d2627ce381cb4c8e3005a600ef393a41a855cca12",
	}
	paramMap["kzgsrs-bw6_761-lagrange-524288"] = paramInfo{
		size: 100809028,
		hash: "fdc984425e9d0059870b8e43ee15e2b78c80e5cffa155941e3842b884bc98c42",
	}
	paramMap["kzgsrs-bw6_761-lagrange-1048576"] = paramInfo{
		size: 201472324,
		hash: "1f1cae1341e561289c141b01cce9f5cb82a6303267360e16faf9ca944910b4c0",
	}
}
