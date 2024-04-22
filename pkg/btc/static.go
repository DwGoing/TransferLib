package btc

import (
	"abao/pkg/common"
	"errors"

	"github.com/btcsuite/btcd/btcec"
)

func GetPrivateKeyFromSeed(seed []byte, index int64, addressType common.AddressType) (*btcec.PrivateKey, error) {
	var path string
	switch addressType {
	case common.AddressType_BTC_LEGACY:
		path = "m/44'/0'/0'/0/"
	case common.AddressType_BTC_SEGWIT:
		path = "m/49'/0'/0'/0/"
	case common.AddressType_BTC_NATIVE_SEGWIT:
		path = "m/84'/0'/0'/0/"
	default:
		return nil, errors.New("unsupported address type")
	}
	return common.GetPrivateKeyFromSeed(seed, common.ExtendedKeyVersion_xprv, path, index)
}

func NewAccountFromPrivateKey(privateKeyHex *btcec.PrivateKey) (*common.Account, error) {
	return common.NewAccountFromPrivateKey(common.AddressType_BSC, privateKeyHex)
}

func NewAccountFromPrivateKeyHex(privateKeyHex string) (*common.Account, error) {
	return common.NewAccountFromPrivateKeyHex(common.AddressType_BSC, privateKeyHex)
}
