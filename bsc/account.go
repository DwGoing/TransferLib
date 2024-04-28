/*
 * @Author: dwgoing jianghy1209@gmail.com
 * @Date: 2024-04-27 23:17:14
 * @LastEditors: dwgoing jianghy1209@gmail.com
 * @LastEditTime: 2024-04-28 20:05:32
 * @FilePath: /transfer_lib/bsc/account.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package bsc

import (
	"account"
	"common"

	"github.com/btcsuite/btcd/chaincfg"
	"github.com/decred/dcrd/dcrec/secp256k1/v4"
)

type Account struct {
	account *account.Account
	chain   common.Chain
}

/*
@title	创建账户
@param	seed	[]byte		种子
@param 	index	int64		账户索引
@return	_		*Account	账户
@return	_		error		异常信息
*/
func NewAccountFromSeed(seed []byte, index int64) (*Account, error) {
	account, err := account.NewAccountFromSeed(seed, index)
	if err != nil {
		return nil, err
	}
	return &Account{
		account: account,
		chain:   common.Chain_BSC,
	}, nil
}

/*
@title	创建账户
@param 	mnemonic	string 		助记词
@param 	password 	string 		密码
@param 	index		int64		账户索引
@return	_			*Account	账户
@return	_			error		异常信息
*/
func NewAccountFromMnemonic(mnemonic string, password string, index int64) (*Account, error) {
	seed, err := account.GetSeedFromMnemonic(mnemonic, password)
	if err != nil {
		return nil, err
	}
	return NewAccountFromSeed(seed, index)
}

/*
@title 	链类型
@param 	Self	*Account
@return _ 		common.Chain	链类型
*/
func (Self *Account) Chain() common.Chain {
	return Self.chain
}

/*
@title 	获取私钥
@param 	Self	*Account
@return _ 		*secp256k1.PrivateKey 	私钥
@return _ 		error 					异常信息
*/
func (Self *Account) GetPrivateKey() (*secp256k1.PrivateKey, error) {
	return account.GetPrivateKeyFromSeed(Self.account.Seed(), &chaincfg.MainNetParams, "m/44'/60'/0'/0/", Self.account.Index())
}

/*
@title 	获取钱包地址
@param 	Self	*Account
@return _ 		string		地址
@return _ 		error 		异常信息
*/
func (Self *Account) GetAddress() (string, error) {
	privateKey, err := Self.GetPrivateKey()
	if err != nil {
		return "", err
	}
	return GetAddressFromPrivateKey(privateKey), nil
}
