package account

import (
	"transfer_lib/pkg/common"

	"github.com/decred/dcrd/dcrec/secp256k1/v4"
)

type IAccount interface {
	Chain() common.Chain
	GetPrivateKey() (*secp256k1.PrivateKey, error)
	GetAddress() (string, error)
}

type Account struct {
	seed  []byte
	index int64
}

/*
@title	创建账户
@param	seed	[]byte		种子
@param 	index	int64		账户索引
@return	_		*Account	账户
@return	_		error		异常信息
*/
func NewAccountFromSeed(seed []byte, index int64) *Account {
	return &Account{
		seed:  seed,
		index: index,
	}
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
	seed, err := GetSeedFromMnemonic(mnemonic, password)
	if err != nil {
		return nil, err
	}
	return NewAccountFromSeed(seed, index), nil
}

/*
@title 	种子
@param 	Self	*Account
@return _ 		[]byte		种子
*/
func (Self *Account) Seed() []byte {
	return Self.seed
}

/*
@title 	账户索引
@param 	Self	*Account
@return _ 		int64		账户索引
*/
func (Self *Account) Index() int64 {
	return Self.index
}
