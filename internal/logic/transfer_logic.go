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
	addressType, err := common.ParseAddressType(in.AddressType)
	if err != nil {
		return nil, err
	}
	var txHash string
	var client common.IChainClient
	var args any
	switch addressType {
	case common.AddressType_BTC_LEGACY, common.AddressType_BTC_SEGWIT, common.AddressType_BTC_NATIVE_SEGWIT:
	case common.AddressType_ETH:
		client = eth.NewEthClient(l.svcCtx.Config.Eth.Nodes, l.svcCtx.Config.Eth.Currencies)
		args = eth.NewEthTransferParameter()
	case common.AddressType_TRON:
		client = tron.NewTronClient(l.svcCtx.Config.Tron.Nodes, l.svcCtx.Config.Tron.ApiKeys, l.svcCtx.Config.Tron.Currencies)
		args = tron.NewTronTransferParameter()
	case common.AddressType_BSC:
		client = bsc.NewBscClient(l.svcCtx.Config.Bsc.Nodes, l.svcCtx.Config.Bsc.Currencies)
		args = bsc.NewBscTransferParameter()
	default:
		return nil, errors.New("unsupported address type")
	}
	txHash, err = client.Transfer(in.PrivateKey, in.To, in.Currency, in.Value, args)
	if err != nil {
		return nil, err
	}

	return &transfer_lib.TransferResponse{
		TranscationHash: txHash,
	}, nil
}
