package common

import (
	"errors"
	"fmt"

	"github.com/btcsuite/btcd/btcec"
	"github.com/btcsuite/btcd/chaincfg"
	"github.com/btcsuite/btcutil/hdkeychain"
	"github.com/ethereum/go-ethereum/accounts"
	"github.com/tyler-smith/go-bip39"
)

type ExtendedKeyVersion [2][4]byte

var (
	ExtendedKeyVersion_xprv ExtendedKeyVersion = [2][4]byte{{0x04, 0x88, 0xad, 0xe4}, {0x04, 0x88, 0xb2, 0x1e}}
	ExtendedKeyVersion_yprv ExtendedKeyVersion = [2][4]byte{{0x04, 0x9d, 0x78, 0x78}, {0x04, 0x88, 0x7c, 0xb2}}
	ExtendedKeyVersion_Yprv ExtendedKeyVersion = [2][4]byte{{0x02, 0x95, 0xb0, 0x05}, {0x04, 0x88, 0xb4, 0x3f}}
	ExtendedKeyVersion_zprv ExtendedKeyVersion = [2][4]byte{{0x04, 0xb2, 0x43, 0x0c}, {0x04, 0x88, 0x47, 0x46}}
	ExtendedKeyVersion_Zprv ExtendedKeyVersion = [2][4]byte{{0x02, 0xaa, 0x7a, 0x99}, {0x04, 0x88, 0x7e, 0xd3}}
	ExtendedKeyVersion_tprv ExtendedKeyVersion = [2][4]byte{{0x04, 0x35, 0x83, 0x94}, {0x04, 0x88, 0x87, 0xcf}}
	ExtendedKeyVersion_uprv ExtendedKeyVersion = [2][4]byte{{0x04, 0x4a, 0x4e, 0x28}, {0x04, 0x88, 0x52, 0x62}}
	ExtendedKeyVersion_Uprv ExtendedKeyVersion = [2][4]byte{{0x02, 0x42, 0x85, 0xb5}, {0x04, 0x88, 0x89, 0xef}}
	ExtendedKeyVersion_vprv ExtendedKeyVersion = [2][4]byte{{0x04, 0x5f, 0x18, 0xbc}, {0x04, 0x88, 0x1c, 0xf6}}
	ExtendedKeyVersion_Vprv ExtendedKeyVersion = [2][4]byte{{0x02, 0x57, 0x50, 0x48}, {0x04, 0x88, 0x54, 0x83}}
)

/*
@title 	从助记词获取种子
@param 	mnemonic	string 	助记词
@param 	password 	string 	密码
@return _ 			[]byte	种子
@return _ 			error 	异常信息
*/
func GetSeedFromMnemonic(mnemonic string, password string) ([]byte, error) {
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
	return seed, nil
}

/*
@title 	派生子私钥
@param 	version 	[4]byte					私钥ID
@param 	path 		string					派生路径
@return _			*hdkeychain.ExtendedKey 子私钥
@return _ 			error 					异常信息
*/
func DerivePrivateKey(seed []byte, version [4]byte, path string) (*hdkeychain.ExtendedKey, error) {
	masterKey, err := hdkeychain.NewMaster(
		seed,
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
@title 	通过种子获取私钥
@param 	seed 	[]byte				种子
@param 	version	ExtendedKeyVersion	扩展私钥版本
@param 	path	string				派生路径
@param 	index	int64				账户索引
@return _		*btcec.PrivateKey 	私钥
@return _ 		error 				异常信息
*/
func GetPrivateKeyFromSeed(seed []byte, version ExtendedKeyVersion, path string, index int64) (*btcec.PrivateKey, error) {
	path = fmt.Sprintf("%s%d", path, index)
	masterKey, err := DerivePrivateKey(seed, version[0], path)
	if err != nil {
		return nil, err
	}
	return masterKey.ECPrivKey()
}

/*
@title 	通过助记词获取私钥
@param 	seed 	[]byte				种子
@param 	version	ExtendedKeyVersion	扩展私钥版本
@param 	path	string				派生路径
@param 	index	int64				账户索引
@return _		*btcec.PrivateKey 	私钥
@return _ 		error 				异常信息
*/
func GetPrivateKeyFromMnemonic(mnemonic string, password string, version ExtendedKeyVersion, path string, index int64) (*btcec.PrivateKey, error) {
	seed, err := GetSeedFromMnemonic(mnemonic, password)
	if err != nil {
		return nil, err
	}
	return GetPrivateKeyFromSeed(seed, version, path, index)
}
