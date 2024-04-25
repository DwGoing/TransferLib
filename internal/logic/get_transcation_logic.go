package logic

import (
	"context"
	"errors"

	"transfer_lib/internal/svc"
	"transfer_lib/pkg/bsc"
	"transfer_lib/pkg/common"
	"transfer_lib/pkg/eth"
	"transfer_lib/pkg/tron"
	"transfer_lib/transfer_lib"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetTranscationLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetTranscationLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetTranscationLogic {
	return &GetTranscationLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetTranscationLogic) GetTranscation(in *transfer_lib.GetTranscationRequest) (*transfer_lib.GetTranscationResponse, error) {
	chain, err := common.ParseChain(in.Chain)
	if err != nil {
		return nil, err
	}
	var client common.IChainClient
	switch chain {
	case common.Chain_BTC:
	case common.Chain_ETH:
		client = eth.NewEthClient(l.svcCtx.Config.Eth.Nodes, l.svcCtx.Config.Eth.Currencies)
	case common.Chain_TRON:
		client = tron.NewTronClient(l.svcCtx.Config.Tron.Nodes, l.svcCtx.Config.Tron.ApiKeys, l.svcCtx.Config.Tron.Currencies)
	case common.Chain_BSC:
		client = bsc.NewBscClient(l.svcCtx.Config.Bsc.Nodes, l.svcCtx.Config.Bsc.Currencies)
	default:
		return nil, errors.New("unsupported chain")
	}
	transcation, err := client.GetTransaction(in.TranscationHash)
	if err != nil {
		return nil, err
	}

	return &transfer_lib.GetTranscationResponse{
		Result:    transcation.Result,
		Height:    transcation.Height,
		Timestamp: transcation.TimeStamp,
		Currency:  transcation.Currency,
		From:      transcation.From,
		To:        transcation.To,
		Value:     transcation.Value,
		Confirms:  transcation.Confirms,
	}, nil
}
