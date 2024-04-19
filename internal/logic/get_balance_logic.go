package logic

import (
	"context"

	"abao/abao"
	"abao/internal/svc"
	"abao/pkg/common"
	"abao/pkg/tron"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetBalanceLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetBalanceLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetBalanceLogic {
	return &GetBalanceLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetBalanceLogic) GetBalance(in *abao.GetBalanceRequest) (*abao.GetBalanceResponse, error) {
	addressType, err := common.ParseAddressType(in.AddressType)
	if err != nil {
		return nil, err
	}
	balances := map[string]float64{}
	var client common.IChainClient
	for _, currencyStr := range in.Currencies {
		var balance float64
		switch addressType {
		case common.AddressType_TRON:
			client = tron.NewTronClient(l.svcCtx.Config.Tron.Nodes, l.svcCtx.Config.Tron.ApiKeys, l.svcCtx.Config.Tron.Currencies)
			balance, err = client.GetBalance(in.Address, tron.NewTronGetBalanceParameter(currencyStr))
			if err != nil {
				continue
			}
		default:
			continue
		}
		balances[currencyStr] = balance
	}

	return &abao.GetBalanceResponse{
		Balances: balances,
	}, nil
}
