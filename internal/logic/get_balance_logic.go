package logic

import (
	"context"
	"errors"
	"sync"

	"abao/abao"
	"abao/internal/svc"
	"abao/pkg/bsc"
	"abao/pkg/common"
	"abao/pkg/eth"
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
	var balancesLock sync.Mutex
	var client common.IChainClient
	var args any
	switch addressType {
	case common.AddressType_TRON:
		client = tron.NewTronClient(l.svcCtx.Config.Tron.Nodes, l.svcCtx.Config.Tron.ApiKeys, l.svcCtx.Config.Tron.Currencies)
		args = tron.NewTronGetBalanceParameter()
	case common.AddressType_ETH:
		client = eth.NewEthClient(l.svcCtx.Config.Eth.Nodes, l.svcCtx.Config.Eth.Currencies)
		args = eth.NewEthGetBalanceParameter()
	case common.AddressType_BSC:
		client = bsc.NewBscClient(l.svcCtx.Config.Bsc.Nodes, l.svcCtx.Config.Bsc.Currencies)
		args = bsc.NewBscGetBalanceParameter()
	default:
		return nil, errors.New("unsupported address type")
	}
	for _, currency := range in.Currencies {
		waitGroup.Add(1)
		go func(c string) {
			defer waitGroup.Done()
			balance, err := client.GetBalance(in.Address, c, args)
			if err != nil {
				return
			}
			balancesLock.Lock()
			balances[c] = balance
			balancesLock.Unlock()
		}(currency)
	}
	waitGroup.Wait()

	return &abao.GetBalanceResponse{
		Balances: balances,
	}, nil
}
