package logic

import (
	"context"

	"transfer_lib/internal/svc"
	"transfer_lib/pkg/bsc"
	"transfer_lib/pkg/btc"
	"transfer_lib/pkg/common"
	"transfer_lib/pkg/eth"
	"transfer_lib/pkg/tron"
	"transfer_lib/transfer_lib"

	"github.com/btcsuite/btcd/btcec"
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
		seed, err := common.GetSeedFromMnemonic(in.Mnemonic, in.Password)
		if err != nil {
			return nil, err
		}
		var privateKey *btcec.PrivateKey
		switch addressType {
		case common.AddressType_BTC_LEGACY, common.AddressType_BTC_SEGWIT, common.AddressType_BTC_NATIVE_SEGWIT:
			privateKey, err = btc.GetPrivateKeyFromSeed(seed, in.Index, addressType)
			if err != nil {
				return nil, err
			}
		case common.AddressType_ETH:
			privateKey, err = eth.GetPrivateKeyFromSeed(seed, in.Index)
			if err != nil {
				return nil, err
			}
		case common.AddressType_TRON:
			privateKey, err = tron.GetPrivateKeyFromSeed(seed, in.Index)
			if err != nil {
				return nil, err
			}
		case common.AddressType_BSC:
			privateKey, err = bsc.GetPrivateKeyFromSeed(seed, in.Index)
			if err != nil {
				return nil, err
			}
		default:
			continue
		}
		account, err := common.NewAccountFromPrivateKey(addressType, privateKey)
		if err != nil {
			continue
		}
		address, err := account.GetAddress()
		if err != nil {
			continue
		}
		accounts[addressTypeStr] = &transfer_lib.Account{
			PrivateKey: account.GetPrivateKeyHex(),
			Address:    address,
		}
	}

	return &transfer_lib.GetAccountResponse{
		Accounts: accounts,
	}, nil
}
