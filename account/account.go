package account

import (
	"github.com/DwGoing/transfer_lib/common"

	"github.com/btcsuite/btcd/btcutil/hdkeychain"
	"github.com/btcsuite/btcd/chaincfg"
	"github.com/decred/dcrd/dcrec/secp256k1/v4"
)

type IAccount interface {
	Chain() common.Chain
	Seed() []byte
	Index() int64
	GetPrivateKey() (*secp256k1.PrivateKey, error)
	GetAddress() (string, error)
}

type Account struct {
	chain common.Chain
	seed  []byte
	index int64
}

/*
@title	创建账户
@param	chain	common.Chain	链类型
@param	seed	[]byte			种子
@param 	index	int64			账户索引
@return	_		*Account		账户
@return	_		error			异常信息
*/
func NewAccountFromSeed(chain common.Chain, seed []byte, index int64) (*Account, error) {
	_, err := hdkeychain.NewMaster(seed, &chaincfg.Params{})
	if err != nil {
		return nil, err
	}
	if index < 0 {
		return nil, common.ErrInvalidIndex
	}
	return &Account{
		chain: chain,
		seed:  seed,
		index: index,
	}, nil
}

/*
@title	创建账户
@param	chain		common.Chain	链类型
@param 	mnemonic	string 			助记词
@param 	password 	string 			密码
@param 	index		int64			账户索引
@return	_			*Account		账户
@return	_			error			异常信息
*/
func NewAccountFromMnemonic(chain common.Chain, mnemonic string, password string, index int64) (*Account, error) {
	seed, err := GetSeedFromMnemonic(mnemonic, password)
	if err != nil {
		return nil, err
	}
	return NewAccountFromSeed(chain, seed, index)
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
