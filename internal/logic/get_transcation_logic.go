package logic

import (
	"context"

	"github.com/DwGoing/transfer_lib/internal/svc"
	"github.com/DwGoing/transfer_lib/pkg/bsc"
	"github.com/DwGoing/transfer_lib/pkg/chain"
	"github.com/DwGoing/transfer_lib/pkg/common"
	"github.com/DwGoing/transfer_lib/pkg/eth"
	"github.com/DwGoing/transfer_lib/pkg/tron"
	"github.com/DwGoing/transfer_lib/transfer_lib"

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
	chainType, err := common.ParseChain(in.Chain)
	if err != nil {
		return nil, err
	}
	var client chain.IChainClient
	switch chainType {
	case common.Chain_ETH:
		client = eth.NewChainClient(l.svcCtx.Config.Eth.Nodes, l.svcCtx.Config.Eth.Currencies)
	case common.Chain_TRON:
		client = tron.NewChainClient(l.svcCtx.Config.Tron.Nodes, l.svcCtx.Config.Tron.ApiKeys, l.svcCtx.Config.Tron.Currencies)
	case common.Chain_BSC:
		client = bsc.NewChainClient(l.svcCtx.Config.Bsc.Nodes, l.svcCtx.Config.Bsc.Currencies)
	default:
		return nil, common.ErrUnsupportedChain
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
