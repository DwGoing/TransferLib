package hd_wallet

import (
	"abao/pkg/common"
	"errors"
	"fmt"

	"github.com/btcsuite/btcd/chaincfg"
	"github.com/btcsuite/btcutil/hdkeychain"
	"github.com/ethereum/go-ethereum/accounts"
	"github.com/tyler-smith/go-bip39"
)

type HDWallet struct {
	seed []byte
}

/*
@title 	从种子导入钱包
@param 	seed   			[]byte 			种子
@return _ 				*HDWallet 		钱包实例
@return _ 				error 			异常信息
*/
func NewHDWalletFromSeed(seed []byte) (*HDWallet, error) {
	if len(seed) < 16 || len(seed) > 64 {
		return nil, errors.New("seed invaild")
	}
	return &HDWallet{
		seed: seed,
	}, nil
}

/*
@title 	从助记词导入钱包
@param 	mnemonic   		string 			助记词
@param 	password 		string 			密码
@return _ 				*HDWallet 		钱包实例
@return _ 				error 			异常信息
*/
func NewHDWalletFromMnemonic(mnemonic string, password string) (*HDWallet, error) {
	if mnemonic == "" {
		return nil, errors.New("FromMnemonic Error: mnemonic empty")
	}
	if !bip39.IsMnemonicValid(mnemonic) {
		return nil, errors.New("FromMnemonic Error: mnemonic invaild")
	}
	seed, err := bip39.NewSeedWithErrorChecking(mnemonic, password)
	if err != nil {
		return nil, err
	}
	return NewHDWalletFromSeed(seed)
}

/*
@title 	派生子私钥
@param 	Self   		*HDWallet
@param 	version 	[4]byte					私钥ID
@param 	path 		string					派生路径
@return _			*hdkeychain.ExtendedKey 子私钥
@return _ 			error 					异常信息
*/
func (Self *HDWallet) DerivePrivateKey(version [4]byte, path string) (*hdkeychain.ExtendedKey, error) {
	masterKey, err := hdkeychain.NewMaster(
		Self.seed,
		&chaincfg.Params{HDPrivateKeyID: version},
	)
	if err != nil {
		return nil, err
	}
	parsedPath, err := accounts.ParseDerivationPath(path)
	if err != nil {
		return nil, err
	}
	accountKey := masterKey
	for _, n := range parsedPath {
		accountKey, err = accountKey.Child(n)
		if err != nil {
			return nil, err
		}
	}
	return accountKey, nil
}

/*
@title 	获取子钱包
@param 	Self   		*HDWallet
@param 	addressType AddressType 地址类型
@param 	index 		Currency 	钱包索引
@return _			*Account	Account实例
@return _ 			error 		异常信息
*/
func (Self *HDWallet) GetAccount(addressType common.AddressType, index int64) (*Account, error) {
	var (
		version Version
		path    string
	)
	switch addressType {
	case common.AddressType_BTC_LEGACY:
		version = Version_xprv
		path = "m/44'/0'/0'/0/"
	case common.AddressType_BTC_SEGWIT:
		version = Version_xprv
		path = "m/49'/0'/0'/0/"
	case common.AddressType_BTC_NATIVE_SEGWIT:
		version = Version_xprv
		path = "m/84'/0'/0'/0/"
	case common.AddressType_ETH:
		version = Version_xprv
		path = "m/44'/60'/0'/0/"
	case common.AddressType_TRON:
		version = Version_xprv
		path = "m/44'/195'/0'/0/"
	case common.AddressType_BSC:
		version = Version_xprv
		path = "m/44'/60'/0'/0/"
	default:
		return nil, errors.New("unsupported address type")
	}
	path = fmt.Sprintf("%s%d", path, index)
	privateKey, err := Self.DerivePrivateKey(version[0], path)
	if err != nil {
		return nil, err
	}
	btcecPrivateKey, err := privateKey.ECPrivKey()
	if err != nil {
		return nil, err
	}
	return NewAccountFromPrivateKey(btcecPrivateKey)
}
