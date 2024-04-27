package logic

import (
	"context"

	"github.com/DwGoing/transfer_lib/internal/svc"
	"github.com/DwGoing/transfer_lib/pkg/account"
	"github.com/DwGoing/transfer_lib/pkg/bsc"
	"github.com/DwGoing/transfer_lib/pkg/chain"
	"github.com/DwGoing/transfer_lib/pkg/common"
	"github.com/DwGoing/transfer_lib/pkg/eth"
	"github.com/DwGoing/transfer_lib/pkg/tron"
	"github.com/DwGoing/transfer_lib/transfer_lib"

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
		args = eth.TransferParameter{}
	case common.Chain_TRON:
		client = tron.NewChainClient(l.svcCtx.Config.Tron.Nodes, l.svcCtx.Config.Tron.ApiKeys, l.svcCtx.Config.Tron.Currencies)
		args = tron.TransferParameter{}
	case common.Chain_BSC:
		client = bsc.NewChainClient(l.svcCtx.Config.Bsc.Nodes, l.svcCtx.Config.Bsc.Currencies)
		args = bsc.TransferParameter{}
	default:
		return nil, common.ErrUnsupportedChain
	}
	pk, err := account.GetPrivateKeyFromHex(in.PrivateKey)
	if err != nil {
		return nil, err
	}
	txHash, err = client.Transfer(pk, in.To, in.Currency, in.Value, args)
	if err != nil {
		return nil, err
	}

	return &transfer_lib.TransferResponse{
		TranscationHash: txHash,
	}, nil
}
