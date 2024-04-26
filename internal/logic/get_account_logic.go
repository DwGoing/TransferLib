package logic

import (
	"context"

	"transfer_lib/internal/svc"
	"transfer_lib/pkg/account"
	"transfer_lib/pkg/bsc"
	"transfer_lib/pkg/btc"
	"transfer_lib/pkg/common"
	"transfer_lib/pkg/eth"
	"transfer_lib/pkg/tron"
	"transfer_lib/transfer_lib"

	"github.com/decred/dcrd/dcrec/secp256k1/v4"
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
		var privateKey *secp256k1.PrivateKey
		var address string
		switch addressType {
		case common.AddressType_BTC_LEGACY, common.AddressType_BTC_NESTED_SEGWIT, common.AddressType_BTC_NATIVE_SEGWIT, common.AddressType_BTC_TAPROOT:
			account, err := btc.NewAccountFromMnemonic(in.Mnemonic, in.Password, in.Index)
			if err != nil {
				return nil, err
			}
			privateKey, err = account.GetPrivateKey(addressType)
			if err != nil {
				return nil, err
			}
			address, err = account.GetAddress(addressType)
			if err != nil {
				return nil, err
			}
		case common.AddressType_ETH:
			account, err := eth.NewAccountFromMnemonic(in.Mnemonic, in.Password, in.Index)
			if err != nil {
				return nil, err
			}
			privateKey, err = account.GetPrivateKey()
			if err != nil {
				return nil, err
			}
			address, err = account.GetAddress()
			if err != nil {
				return nil, err
			}
		case common.AddressType_TRON:
			account, err := tron.NewAccountFromMnemonic(in.Mnemonic, in.Password, in.Index)
			if err != nil {
				return nil, err
			}
			privateKey, err = account.GetPrivateKey()
			if err != nil {
				return nil, err
			}
			address, err = account.GetAddress()
			if err != nil {
				return nil, err
			}
		case common.AddressType_BSC:
			account, err := bsc.NewAccountFromMnemonic(in.Mnemonic, in.Password, in.Index)
			if err != nil {
				return nil, err
			}
			privateKey, err = account.GetPrivateKey()
			if err != nil {
				return nil, err
			}
			address, err = account.GetAddress()
			if err != nil {
				return nil, err
			}
		default:
			return nil, account.ErrUnsupportedAddressType
		}
		privateKeyHex, err := account.ToHex(privateKey)
		if err != nil {
			return nil, err
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
