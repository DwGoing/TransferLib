package logic

import (
	"context"
	"errors"

	"abao/abao"
	"abao/internal/svc"
	"abao/pkg/common"
	"abao/pkg/eth"
	"abao/pkg/tron"

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

func (l *GetTranscationLogic) GetTranscation(in *abao.GetTranscationRequest) (*abao.GetTranscationResponse, error) {
	chain, err := common.ParseChain(in.Chain)
	if err != nil {
		return nil, err
	}
	var client common.IChainClient
	switch chain {
	case common.Chain_TRON:
		client = tron.NewTronClient(l.svcCtx.Config.Tron.Nodes, l.svcCtx.Config.Tron.ApiKeys, l.svcCtx.Config.Tron.Currencies)
	case common.Chain_ETH:
		client = eth.NewEthClient(l.svcCtx.Config.Eth.Nodes, l.svcCtx.Config.Eth.Currencies)
	default:
		return nil, errors.New("unsupported chain")
	}
	transcation, err := client.GetTransaction(in.TranscationHash)
	if err != nil {
		return nil, err
	}

	return &abao.GetTranscationResponse{
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
