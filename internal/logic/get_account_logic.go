package logic

import (
	"context"

	"transfer_lib/internal/svc"
	"transfer_lib/pkg/account"
	"transfer_lib/pkg/common"
	"transfer_lib/transfer_lib"

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

func (l *GetAccountLogic) GetAccount(in *transfer_lib.GetAccountRequest) (*transfer_lib.GetAccountResponse, error) {
	accounts := map[string]*transfer_lib.Account{}
	for _, addressTypeStr := range in.AddressTypes {
		addressType, err := common.ParseAddressType(addressTypeStr)
		if err != nil {
			continue
		}
		account, err := account.NewAccountFromMnemonic(in.Mnemonic, in.Password, in.Index)
		if err != nil {
			continue
		}
		privateKeyHex, err := account.GetPrivateKeyHex(addressType)
		if err != nil {
			continue
		}
		address, err := account.GetAddress(addressType)
		if err != nil {
			continue
		}
		accounts[addressTypeStr] = &transfer_lib.Account{
			PrivateKey: privateKeyHex,
			Address:    address,
		}
	}

	return &transfer_lib.GetAccountResponse{
		Accounts: accounts,
	}, nil
}
