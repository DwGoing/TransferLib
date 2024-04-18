package hd_wallet

import (
	"abao/pkg/common"
	"errors"

	"github.com/btcsuite/btcd/btcec"
	"github.com/btcsuite/btcutil"
	"github.com/btcsuite/btcutil/base58"
	"github.com/cosmos/btcutil/bech32"
	"github.com/ethereum/go-ethereum/crypto"
	tronCommon "github.com/fbsobreira/gotron-sdk/pkg/common"
)

type Account struct {
	Index       int64
	AddressType common.AddressType
	PrivateKey  *btcec.PrivateKey
}

/*
@title 	获取钱包地址
@param 	Self   	*Account 	Account实例
@return _ 		string 		钱包地址
@return _ 		error 		异常信息
*/
func (Self *Account) GetAddress() (string, error) {
	address := ""
	switch Self.AddressType {
	case common.AddressType_BTC_LEGACY:
		bytes := btcutil.Hash160(Self.PrivateKey.PubKey().SerializeCompressed())
		address = base58.CheckEncode(bytes, 0x00)
	case common.AddressType_BTC_SEGWIT:
		bytes := btcutil.Hash160(Self.PrivateKey.PubKey().SerializeCompressed())
		bytes = append([]byte{0x00, 0x14}, bytes...)
		bytes = btcutil.Hash160(bytes)
		address = base58.CheckEncode(bytes, 0x05)
	case common.AddressType_BTC_NATIVE_SEGWIT:
		bytes := btcutil.Hash160(Self.PrivateKey.PubKey().SerializeCompressed())
		converted, err := bech32.ConvertBits(bytes, 8, 5, true)
		if err != nil {
			break
		}
		combined := make([]byte, len(converted)+1)
		combined[0] = 0x00
		copy(combined[1:], converted)
		address, err = bech32.Encode("bc", combined)
		if err != nil {
			break
		}
	case common.AddressType_ETH:
		address = crypto.PubkeyToAddress(Self.PrivateKey.ToECDSA().PublicKey).Hex()
	case common.AddressType_TRON:
		ethAddress := crypto.PubkeyToAddress(Self.PrivateKey.ToECDSA().PublicKey)
		tronAddress := make([]byte, 0)
		tronAddress = append(tronAddress, byte(0x41))
		tronAddress = append(tronAddress, ethAddress.Bytes()...)
		address = tronCommon.EncodeCheck(tronAddress)
	case common.AddressType_BSC:
		address = crypto.PubkeyToAddress(Self.PrivateKey.ToECDSA().PublicKey).Hex()
	default:
		return "", errors.New("unsupported address type")
	}
	return address, nil
}

func (Self *Account) GetBalance() (int64, error) {
	switch Self.AddressType {
	case common.AddressType_BTC_LEGACY:
	case common.AddressType_BTC_SEGWIT:
	case common.AddressType_BTC_NATIVE_SEGWIT:
	case common.AddressType_ETH:
	case common.AddressType_TRON:
	case common.AddressType_BSC:
	default:
		return 0, errors.New("unsupported address type")
	}
	return 0, nil
}
