package logic

import (
	"context"

	"abao/abao"
	"abao/internal/svc"
	"abao/pkg/common"
	"abao/pkg/hd_wallet"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetAddressLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetAddressLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetAddressLogic {
	return &GetAddressLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetAddressLogic) GetAddress(in *abao.GetAddressRequest) (*abao.GetAddressResponse, error) {
	hdWallet, err := hd_wallet.FromMnemonic(in.Mnemonic, in.Password)
	if err != nil {
		return nil, err
	}
	addresses := map[string]string{}
	for _, addressTypeStr := range in.AddressTypes {
		addressType, err := common.ParseAddressType(addressTypeStr)
		if err != nil {
			continue
		}
		account, err := hdWallet.GetAccount(addressType, in.Index)
		if err != nil {
			continue
		}
		address, err := account.GetAddress()
		if err != nil {
			continue
		}
		addresses[addressTypeStr] = address
	}

	return &abao.GetAddressResponse{
		Addresses: addresses,
	}, nil
}
