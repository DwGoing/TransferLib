package logic

import (
	"context"

	"transfer_lib/internal/svc"
	"transfer_lib/pkg/account"
	"transfer_lib/pkg/bsc"
	"transfer_lib/pkg/chain"
	"transfer_lib/pkg/common"
	"transfer_lib/pkg/eth"
	"transfer_lib/pkg/tron"
	"transfer_lib/transfer_lib"

	"github.com/zeromicro/go-zero/core/logx"
)

type TransferLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewTransferLogic(ctx context.Context, svcCtx *svc.ServiceContext) *TransferLogic {
	return &TransferLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *TransferLogic) Transfer(in *transfer_lib.TransferRequest) (*transfer_lib.TransferResponse, error) {
	chainType, err := common.ParseChain(in.Chain)
	if err != nil {
		return nil, err
	}
	var txHash string
	var client chain.IChainClient
	var args any
	switch chainType {
	case common.Chain_ETH:
		client = eth.NewChainClient(l.svcCtx.Config.Eth.Nodes, l.svcCtx.Config.Eth.Currencies)
		args = eth.NewEthTransferParameter()
	case common.Chain_TRON:
		client = tron.NewChainClient(l.svcCtx.Config.Tron.Nodes, l.svcCtx.Config.Tron.ApiKeys, l.svcCtx.Config.Tron.Currencies)
		args = tron.NewTronTransferParameter()
	case common.Chain_BSC:
		client = bsc.NewChainClient(l.svcCtx.Config.Bsc.Nodes, l.svcCtx.Config.Bsc.Currencies)
		args = bsc.NewBscTransferParameter()
	default:
		return nil, common.ErrUnsupportedChain
	}
	pk, err := account.GetPrivateKeyFromHex(in.PrivateKey)
	txHash, err = client.Transfer(pk.ToECDSA(), in.To, in.Currency, in.Value, args)
	if err != nil {
		return nil, err
	}

	return &transfer_lib.TransferResponse{
		TranscationHash: txHash,
	}, nil
}
