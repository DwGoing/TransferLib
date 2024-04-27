package btc

import (
	"github.com/DwGoing/transfer_lib/pkg/account"
	"github.com/DwGoing/transfer_lib/pkg/common"

	"github.com/btcsuite/btcd/btcec/v2/schnorr"
	"github.com/btcsuite/btcd/btcutil"
	"github.com/btcsuite/btcd/chaincfg"
	"github.com/btcsuite/btcd/txscript"
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
		chain:   common.Chain_BTC,
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
@param 	Self		*Account
@param 	addressType	common.AddressType		地址类型
@return _ 			*secp256k1.PrivateKey 	私钥
@return _ 			error 					异常信息
*/
func (Self *Account) GetPrivateKey(addressType common.AddressType) (*secp256k1.PrivateKey, error) {
	var path string
	switch addressType {
	case common.AddressType_BTC_LEGACY:
		path = "m/44'/0'/0'/0/"
	case common.AddressType_BTC_NESTED_SEGWIT:
		path = "m/49'/0'/0'/0/"
	case common.AddressType_BTC_NATIVE_SEGWIT:
		path = "m/84'/0'/0'/0/"
	case common.AddressType_BTC_TAPROOT:
		path = "m/86'/0'/0'/0/"
	default:
		return nil, common.ErrUnsupportedAddressType
	}
	return account.GetPrivateKeyFromSeed(Self.account.Seed(), &chaincfg.MainNetParams, path, Self.account.Index())
}

/*
@title 	获取钱包地址
@param 	Self		*Account
@param 	addressType	common.AddressType	地址类型
@return _ 			string				地址
@return _ 			error 				异常信息
*/
func (Self *Account) GetAddress(addressType common.AddressType) (string, error) {
	privateKey, err := Self.GetPrivateKey(addressType)
	if err != nil {
		return "", err
	}
	var address string
	switch addressType {
	case common.AddressType_BTC_LEGACY:
		pubKeyHash := btcutil.Hash160(privateKey.PubKey().SerializeCompressed())
		addressPubKeyHash, err := btcutil.NewAddressPubKeyHash(pubKeyHash, &chaincfg.MainNetParams)
		if err != nil {
			return "", err
		}
		address = addressPubKeyHash.EncodeAddress()
	case common.AddressType_BTC_NESTED_SEGWIT:
		pubKeyHash := btcutil.Hash160(privateKey.PubKey().SerializeCompressed())
		addressWitnessPubKeyHash, err := btcutil.NewAddressWitnessPubKeyHash(pubKeyHash, &chaincfg.MainNetParams)
		if err != nil {
			return "", err
		}
		script, err := txscript.PayToAddrScript(addressWitnessPubKeyHash)
		if err != nil {
			return "", err
		}
		addressScriptHash, err := btcutil.NewAddressScriptHash(script, &chaincfg.MainNetParams)
		if err != nil {
			return "", err
		}
		address = addressScriptHash.EncodeAddress()
	case common.AddressType_BTC_NATIVE_SEGWIT:
		pubKeyHash := btcutil.Hash160(privateKey.PubKey().SerializeCompressed())
		addressWitnessPubKeyHash, err := btcutil.NewAddressWitnessPubKeyHash(pubKeyHash, &chaincfg.MainNetParams)
		if err != nil {
			return "", err
		}
		address = addressWitnessPubKeyHash.EncodeAddress()
	case common.AddressType_BTC_TAPROOT:
		taprootKeyNoScript := txscript.ComputeTaprootKeyNoScript(privateKey.PubKey())
		serializedPubKey := schnorr.SerializePubKey(taprootKeyNoScript)
		addressTaproot, err := btcutil.NewAddressTaproot(serializedPubKey, &chaincfg.MainNetParams)
		if err != nil {
			return "", err
		}
		address = addressTaproot.EncodeAddress()
	default:
		return "", common.ErrUnsupportedAddressType
	}
	return address, nil
}
