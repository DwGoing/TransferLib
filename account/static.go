package account

import (
	"encoding/hex"
	"fmt"

	"github.com/DwGoing/transfer_lib/common"

	"github.com/btcsuite/btcd/btcec/v2"
	"github.com/btcsuite/btcd/btcutil/hdkeychain"
	"github.com/btcsuite/btcd/chaincfg"
	"github.com/decred/dcrd/dcrec/secp256k1/v4"
	"github.com/ethereum/go-ethereum/accounts"
	goTornSdkCommon "github.com/fbsobreira/gotron-sdk/pkg/common"
	"github.com/tyler-smith/go-bip39"
)

/*
@title 	从助记词获取种子
@param 	mnemonic	string 	助记词
@param 	password 	string 	密码
@return _ 			[]byte	种子
@return _ 			error 	异常信息
*/
func GetSeedFromMnemonic(mnemonic string, password string) ([]byte, error) {
	if mnemonic == "" || !bip39.IsMnemonicValid(mnemonic) {
		return nil, common.ErrInvalidMnemonic
	}
	seed, err := bip39.NewSeedWithErrorChecking(mnemonic, password)
	if err != nil {
		return nil, err
	}
	return seed, nil
}

/*
@title 	派生子私钥
@param 	seed 		[]byte					种子
@param 	chainParams *chaincfg.Params		链参数
@param 	path 		string					派生路径
@return _			*hdkeychain.ExtendedKey 子私钥
@return _ 			error 					异常信息
*/
func DerivePrivateKey(seed []byte, chainParams *chaincfg.Params, path string) (*hdkeychain.ExtendedKey, error) {
	masterKey, err := hdkeychain.NewMaster(seed, chainParams)
	if err != nil {
		return nil, err
	}
	parsedPath, err := accounts.ParseDerivationPath(path)
	if err != nil {
		return nil, err
	}
	childKey := masterKey
	for _, n := range parsedPath {
		childKey, err = childKey.Derive(n)
		if err != nil {
			return nil, err
		}
	}
	return childKey, nil
}

/*
@title 	通过种子获取私钥
@param 	seed 		[]byte					种子
@param 	chainParams *chaincfg.Params		链参数
@param 	path		string					派生路径
@param 	index		int64					账户索引
@return _			*secp256k1.PrivateKey 	私钥
@return _ 			error 					异常信息
*/
func GetPrivateKeyFromSeed(seed []byte, chainParams *chaincfg.Params, path string, index int64) (*secp256k1.PrivateKey, error) {
	path = fmt.Sprintf("%s%d", path, index)
	masterKey, err := DerivePrivateKey(seed, chainParams, path)
	if err != nil {
		return nil, err
	}
	return masterKey.ECPrivKey()
}

/*
@title 	通过助记词获取私钥
@param 	seed 		[]byte					种子
@param 	chainParams *chaincfg.Params		链参数
@param 	path		string					派生路径
@param 	index		int64					账户索引
@return _			*secp256k1.PrivateKey 	私钥
@return _ 			error 					异常信息
*/
func GetPrivateKeyFromMnemonic(mnemonic string, password string, chainParams *chaincfg.Params, path string, index int64) (*secp256k1.PrivateKey, error) {
	seed, err := GetSeedFromMnemonic(mnemonic, password)
	if err != nil {
		return nil, err
	}
	return GetPrivateKeyFromSeed(seed, chainParams, path, index)
}

/*
@title 	十六进制格式私钥
@param 	privateKey	*secp256k1.PrivateKey	私钥
@return _ 			string 					十六进制私钥
@return _ 			error 					异常信息
*/
func PrivateKeyToHex(privateKey *secp256k1.PrivateKey) (string, error) {
	return goTornSdkCommon.Bytes2Hex(privateKey.Serialize()), nil
}

/*
@title 	十六进制格式私钥
@param 	privateKey	string					十六进制私钥
@return _ 			*secp256k1.PrivateKey 	私钥
@return _ 			error 					异常信息
*/
func GetPrivateKeyFromHex(privateKeyHex string) (*secp256k1.PrivateKey, error) {
	bytes, err := hex.DecodeString(privateKeyHex)
	if err != nil {
		return nil, err
	}
	privateKey, _ := btcec.PrivKeyFromBytes(bytes)
	return privateKey, nil
}
