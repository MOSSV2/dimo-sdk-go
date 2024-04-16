package contract

import (
	"context"

	"github.com/MOSSV2/dimo-sdk-go/contract/go/bank"
	"github.com/MOSSV2/dimo-sdk-go/contract/go/bls"
	"github.com/MOSSV2/dimo-sdk-go/contract/go/control"
	"github.com/MOSSV2/dimo-sdk-go/contract/go/epoch"
	"github.com/MOSSV2/dimo-sdk-go/contract/go/file"
	"github.com/MOSSV2/dimo-sdk-go/contract/go/gpu"
	"github.com/MOSSV2/dimo-sdk-go/contract/go/model"
	"github.com/MOSSV2/dimo-sdk-go/contract/go/node"
	"github.com/MOSSV2/dimo-sdk-go/contract/go/proof"
	"github.com/MOSSV2/dimo-sdk-go/contract/go/round"
	"github.com/MOSSV2/dimo-sdk-go/contract/go/space"
	"github.com/MOSSV2/dimo-sdk-go/contract/go/token"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/ethclient"
)

func NewBank(ctx context.Context) (*bank.Bank, error) {
	client, err := ethclient.DialContext(ctx, DevChain)
	if err != nil {
		return nil, err
	}

	return bank.NewBank(BankAddr, client)
}

func NewToken(ctx context.Context) (*token.Token, error) {
	client, err := ethclient.DialContext(ctx, DevChain)
	if err != nil {
		return nil, err
	}

	bi, err := bank.NewBank(BankAddr, client)
	if err != nil {
		return nil, err
	}
	taddr, err := bi.Get(&bind.CallOpts{From: Base}, "token")
	if err != nil {
		return nil, err
	}
	return token.NewToken(taddr, client)
}

func NewEpoch(ctx context.Context) (*epoch.Epoch, error) {
	client, err := ethclient.DialContext(ctx, DevChain)
	if err != nil {
		return nil, err
	}

	bi, err := bank.NewBank(BankAddr, client)
	if err != nil {
		return nil, err
	}
	eaddr, err := bi.Get(&bind.CallOpts{From: Base}, "epoch")
	if err != nil {
		return nil, err
	}
	return epoch.NewEpoch(eaddr, client)
}

func NewControl(ctx context.Context) (*control.Control, error) {
	client, err := ethclient.DialContext(ctx, DevChain)
	if err != nil {
		return nil, err
	}

	bi, err := bank.NewBank(BankAddr, client)
	if err != nil {
		return nil, err
	}
	caddr, err := bi.Get(&bind.CallOpts{From: Base}, "control")
	if err != nil {
		return nil, err
	}
	return control.NewControl(caddr, client)
}

func NewNode(ctx context.Context) (*node.Node, error) {
	client, err := ethclient.DialContext(ctx, DevChain)
	if err != nil {
		return nil, err
	}

	bi, err := bank.NewBank(BankAddr, client)
	if err != nil {
		return nil, err
	}
	naddr, err := bi.Get(&bind.CallOpts{From: Base}, "node")
	if err != nil {
		return nil, err
	}
	return node.NewNode(naddr, client)
}

func NewFile(ctx context.Context) (*file.File, error) {
	client, err := ethclient.DialContext(ctx, DevChain)
	if err != nil {
		return nil, err
	}

	bi, err := bank.NewBank(BankAddr, client)
	if err != nil {
		return nil, err
	}
	faddr, err := bi.Get(&bind.CallOpts{From: Base}, "file")
	if err != nil {
		return nil, err
	}
	return file.NewFile(faddr, client)
}

func NewRound(ctx context.Context) (*round.Round, error) {
	client, err := ethclient.DialContext(ctx, DevChain)
	if err != nil {
		return nil, err
	}

	bi, err := bank.NewBank(BankAddr, client)
	if err != nil {
		return nil, err
	}
	raddr, err := bi.Get(&bind.CallOpts{From: Base}, "round")
	if err != nil {
		return nil, err
	}
	return round.NewRound(raddr, client)
}

func NewProof(ctx context.Context) (*proof.Proof, error) {
	client, err := ethclient.DialContext(ctx, DevChain)
	if err != nil {
		return nil, err
	}

	bi, err := bank.NewBank(BankAddr, client)
	if err != nil {
		return nil, err
	}
	paddr, err := bi.Get(&bind.CallOpts{From: Base}, "proof")
	if err != nil {
		return nil, err
	}
	return proof.NewProof(paddr, client)
}

func NewGPU(ctx context.Context) (*gpu.Gpu, error) {
	client, err := ethclient.DialContext(ctx, DevChain)
	if err != nil {
		return nil, err
	}

	bi, err := bank.NewBank(BankAddr, client)
	if err != nil {
		return nil, err
	}
	gaddr, err := bi.Get(&bind.CallOpts{From: Base}, "gpu")
	if err != nil {
		return nil, err
	}
	return gpu.NewGpu(gaddr, client)
}

func NewModel(ctx context.Context) (*model.Model, error) {
	client, err := ethclient.DialContext(ctx, DevChain)
	if err != nil {
		return nil, err
	}

	bi, err := bank.NewBank(BankAddr, client)
	if err != nil {
		return nil, err
	}
	maddr, err := bi.Get(&bind.CallOpts{From: Base}, "model")
	if err != nil {
		return nil, err
	}
	return model.NewModel(maddr, client)
}

func NewSpace(ctx context.Context) (*space.Space, error) {
	client, err := ethclient.DialContext(ctx, DevChain)
	if err != nil {
		return nil, err
	}

	bi, err := bank.NewBank(BankAddr, client)
	if err != nil {
		return nil, err
	}
	aaddr, err := bi.Get(&bind.CallOpts{From: Base}, "space")
	if err != nil {
		return nil, err
	}
	return space.NewSpace(aaddr, client)
}

func NewBLS(ctx context.Context) (*bls.BLS, error) {
	client, err := ethclient.DialContext(ctx, DevChain)
	if err != nil {
		return nil, err
	}

	bi, err := bank.NewBank(BankAddr, client)
	if err != nil {
		return nil, err
	}
	aaddr, err := bi.Get(&bind.CallOpts{From: Base}, "bls")
	if err != nil {
		return nil, err
	}
	return bls.NewBLS(aaddr, client)
}
