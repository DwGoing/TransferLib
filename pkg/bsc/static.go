package bsc

import (
	"abao/pkg/common"

	"github.com/btcsuite/btcd/btcec"
)

func GetPrivateKeyFromSeed(seed []byte, index int64) (*btcec.PrivateKey, error) {
	return common.GetPrivateKeyFromSeed(seed, common.ExtendedKeyVersion_xprv, "m/44'/60'/0'/0/", index)
}

func NewAccountFromPrivateKey(privateKeyHex *btcec.PrivateKey) (*common.Account, error) {
	return common.NewAccountFromPrivateKey(common.AddressType_BSC, privateKeyHex)
}

func NewAccountFromPrivateKeyHex(privateKeyHex string) (*common.Account, error) {
	return common.NewAccountFromPrivateKeyHex(common.AddressType_BSC, privateKeyHex)
}
