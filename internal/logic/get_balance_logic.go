package logic

import (
	"context"
	"errors"
	"sync"

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
	waitGroup := sync.WaitGroup{}
	balances := map[string]float64{}
	switch addressType {
	case common.AddressType_TRON:
		client := tron.NewTronClient(l.svcCtx.Config.Tron.Nodes, l.svcCtx.Config.Tron.ApiKeys, l.svcCtx.Config.Tron.Currencies)
		for _, currency := range in.Currencies {
			waitGroup.Add(1)
			go func(c string) {
				defer waitGroup.Done()
				balance, err := client.GetBalance(in.Address, tron.NewTronGetBalanceParameter(c))
				if err != nil {
					return
				}
				balances[c] = balance
			}(currency)
		}
	default:
		return nil, errors.New("unsupported address type")
	}
	waitGroup.Wait()

	return &abao.GetBalanceResponse{
		Balances: balances,
	}, nil
}
