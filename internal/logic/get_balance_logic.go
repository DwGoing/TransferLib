package logic

import (
	"context"
	"sync"

	"github.com/DwGoing/transfer_lib/internal/svc"
	"github.com/DwGoing/transfer_lib/pkg/bsc"
	"github.com/DwGoing/transfer_lib/pkg/chain"
	"github.com/DwGoing/transfer_lib/pkg/common"
	"github.com/DwGoing/transfer_lib/pkg/eth"
	"github.com/DwGoing/transfer_lib/pkg/tron"
	"github.com/DwGoing/transfer_lib/transfer_lib"
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

func (l *GetBalanceLogic) GetBalance(in *transfer_lib.GetBalanceRequest) (*transfer_lib.GetBalanceResponse, error) {
	addressType, err := common.ParseAddressType(in.AddressType)
	if err != nil {
		return nil, err
	}
	waitGroup := sync.WaitGroup{}
	balances := map[string]float64{}
	var balancesLock sync.Mutex
	var client chain.IChainClient
	var args any
	switch addressType {
	case common.AddressType_ETH:
		client = eth.NewChainClient(l.svcCtx.Config.Eth.Nodes, l.svcCtx.Config.Eth.Currencies)
		args = eth.GetBalanceParameter{}
	case common.AddressType_TRON:
		client = tron.NewChainClient(l.svcCtx.Config.Tron.Nodes, l.svcCtx.Config.Tron.ApiKeys, l.svcCtx.Config.Tron.Currencies)
		args = tron.GetBalanceParameter{}
	case common.AddressType_BSC:
		client = bsc.NewChainClient(l.svcCtx.Config.Bsc.Nodes, l.svcCtx.Config.Bsc.Currencies)
		args = bsc.GetBalanceParameter{}
	default:
		return nil, common.ErrUnsupportedAddressType
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

	return &transfer_lib.GetBalanceResponse{
		Balances: balances,
	}, nil
}
