package estimator

import (
	"math/big"

	"github.com/ethereum-optimism/optimism/op-e2e/bindings"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/core/types"
)

type Price struct {
	L1BaseFee         *big.Int
	L1BaseFeeScalar   uint32
	BlobBaseFee       *big.Int
	BlobBaseFeeScalar uint32
}

func GetPrice(gpo *bindings.GasPriceOracle, opt *bind.CallOpts) (*Price, error) {
	p := &Price{}
	var err error

	p.L1BaseFee, err = gpo.L1BaseFee(opt)
	if err != nil {
		return p, err
	}

	p.L1BaseFeeScalar, err = gpo.BaseFeeScalar(opt)
	if err != nil {
		return p, err
	}

	p.BlobBaseFee, err = gpo.BlobBaseFee(opt)
	if err != nil {
		return p, err
	}

	p.BlobBaseFeeScalar, err = gpo.BlobBaseFeeScalar(opt)
	if err != nil {
		return p, err
	}
	return p, nil
}

func FlzCompress(data []byte) uint32 {
	return types.FlzCompressLen(data) + 68
}

func GetSize(flzSize uint32) *big.Int {
	estimatedSize := big.NewInt(int64(flzSize))
	estimatedSize.Mul(estimatedSize, types.L1CostFastlzCoef)
	estimatedSize.Add(estimatedSize, types.L1CostIntercept)

	if estimatedSize.Cmp(types.MinTransactionSizeScaled) < 0 {
		estimatedSize.Set(types.MinTransactionSizeScaled)
	}
	//fmt.Println(flzSize, estimatedSize)
	return estimatedSize
}

func L1Fee(p *Price, size *big.Int) *big.Int {
	res := new(big.Int).Set(p.L1BaseFee)
	res.Mul(res, big.NewInt(int64(p.L1BaseFeeScalar)))
	res.Mul(res, big.NewInt(16))
	tmp := new(big.Int).Mul(p.BlobBaseFee, big.NewInt(int64(p.BlobBaseFeeScalar)))
	res.Add(res, tmp)
	res.Mul(res, size)

	res.Div(res, big.NewInt(Decimals))
	res.Div(res, big.NewInt(Decimals))
	//fmt.Println("l1fee: ", utils.FormatEth(res))
	return res
}
