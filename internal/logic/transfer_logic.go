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

func (l *TransferLogic) Transfer(in *abao.TransferRequest) (*abao.TransferResponse, error) {
	addressType, err := common.ParseAddressType(in.AddressType)
	if err != nil {
		return nil, err
	}
	var txHash string
	var client common.IChainClient
	var args any
	switch addressType {
	case common.AddressType_TRON:
		client = tron.NewTronClient(l.svcCtx.Config.Tron.Nodes, l.svcCtx.Config.Tron.ApiKeys, l.svcCtx.Config.Tron.Currencies)
		args = tron.NewTronTransferParameter()
	default:
		return nil, errors.New("unsupported address type")
	}
	txHash, err = client.Transfer(in.PrivateKey, in.To, in.Currency, in.Value, args)
	if err != nil {
		return nil, err
	}

	return &abao.TransferResponse{
		TranscationHash: txHash,
	}, nil
}
