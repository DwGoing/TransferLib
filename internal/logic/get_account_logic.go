package logic

import (
	"context"

	"abao/abao"
	"abao/internal/svc"
	"abao/pkg/common"
	"abao/pkg/hd_wallet"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetAccountLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetAccountLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetAccountLogic {
	return &GetAccountLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetAccountLogic) GetAccount(in *abao.GetAccountRequest) (*abao.GetAccountResponse, error) {
	hdWallet, err := hd_wallet.NewHDWalletFromMnemonic(in.Mnemonic, in.Password)
	if err != nil {
		return nil, err
	}
	accounts := map[string]*abao.Account{}
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
		accounts[addressTypeStr] = &abao.Account{
			PrivateKey: account.GetPrivateKeyHex(),
			Address:    address,
		}
	}

	return &abao.GetAccountResponse{
		Accounts: accounts,
	}, nil
}
