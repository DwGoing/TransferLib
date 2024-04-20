package logic

import (
	"context"
	"errors"

	"abao/abao"
	"abao/internal/svc"
	"abao/pkg/common"
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
	var transcation *common.Transaction
	switch chain {
	case common.Chain_TRON:
		client := tron.NewTronClient(l.svcCtx.Config.Tron.Nodes, l.svcCtx.Config.Tron.ApiKeys, l.svcCtx.Config.Tron.Currencies)
		tx, err := client.GetTransaction(in.TranscationHash)
		if err != nil {
			return nil, err
		}
		tx.Chain = common.Chain_TRON
		transcation = tx
	default:
		return nil, errors.New("unsupported chain")
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
