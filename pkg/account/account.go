package account

import (
	"transfer_lib/pkg/common"

	"github.com/btcsuite/btcd/btcec/v2/schnorr"
	"github.com/btcsuite/btcd/btcutil"
	"github.com/btcsuite/btcd/chaincfg"
	"github.com/btcsuite/btcd/txscript"
	"github.com/decred/dcrd/dcrec/secp256k1/v4"
	"github.com/ethereum/go-ethereum/crypto"
	goTornSdkCommon "github.com/fbsobreira/gotron-sdk/pkg/common"
)

type IAccount interface {
	Chain() common.Chain
	GetAddress() (string, error)
	GetPrivateKey() (*secp256k1.PrivateKey, error)
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
func NewAccountFromSeed(seed []byte, index int64) (*Account, error) {
	return &Account{
		seed:  seed,
		index: index,
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
	seed, err := GetSeedFromMnemonic(mnemonic, password)
	if err != nil {
		return nil, err
	}
	return NewAccountFromSeed(seed, index)
}

/*
@title 	获取私钥
@param 	Self   		*Account
@param	addressType common.AddressType 		地址类型
@return _ 			*secp256k1.PrivateKey 	私钥
@return _ 			error 					异常信息
*/
func (Self *Account) GetPrivateKey(addressType common.AddressType) (*secp256k1.PrivateKey, error) {
	var privateKey *secp256k1.PrivateKey
	var err error
	switch addressType {
	case common.AddressType_BTC_LEGACY:
		privateKey, err = GetPrivateKeyFromSeed(Self.seed, &chaincfg.MainNetParams, "m/44'/0'/0'/0/", Self.index)
	case common.AddressType_BTC_NESTED_SEGWIT:
		privateKey, err = GetPrivateKeyFromSeed(Self.seed, &chaincfg.MainNetParams, "m/49'/0'/0'/0/", Self.index)
	case common.AddressType_BTC_NATIVE_SEGWIT:
		privateKey, err = GetPrivateKeyFromSeed(Self.seed, &chaincfg.MainNetParams, "m/84'/0'/0'/0/", Self.index)
	case common.AddressType_BTC_TAPROOT:
		privateKey, err = GetPrivateKeyFromSeed(Self.seed, &chaincfg.MainNetParams, "m/86'/0'/0'/0/", Self.index)
	case common.AddressType_ETH:
		privateKey, err = GetPrivateKeyFromSeed(Self.seed, &chaincfg.MainNetParams, "m/44'/60'/0'/0/", Self.index)
	case common.AddressType_TRON:
		privateKey, err = GetPrivateKeyFromSeed(Self.seed, &chaincfg.MainNetParams, "m/44'/195'/0'/0/", Self.index)
	case common.AddressType_BSC:
		privateKey, err = GetPrivateKeyFromSeed(Self.seed, &chaincfg.MainNetParams, "m/44'/60'/0'/0/", Self.index)
	default:
		return nil, ErrUnsupportedAddressType
	}
	if err != nil {
		return nil, err
	}
	return privateKey, nil
}

/*
@title 	获取十六进制私钥
@param 	Self   		*Account
@param	addressType common.AddressType 	地址类型
@return _ 			string 				私钥
@return _ 			error 				异常信息
*/
func (Self *Account) GetPrivateKeyHex(addressType common.AddressType) (string, error) {
	privateKey, err := Self.GetPrivateKey(addressType)
	if err != nil {
		return "", err
	}
	return goTornSdkCommon.Bytes2Hex(privateKey.Serialize()), nil
}

/*
@title 	获取钱包地址
@param 	Self   		*Account
@return _ 			string 				地址
@return _ 			error 				异常信息
*/
func (Self *Account) GetAddress(addressType common.AddressType) (string, error) {
	privateKey, err := Self.GetPrivateKey(addressType)
	if err != nil {
		return "", err
	}
	address := ""
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
	case common.AddressType_ETH:
		address = crypto.PubkeyToAddress(privateKey.ToECDSA().PublicKey).Hex()
	case common.AddressType_TRON:
		ethAddress := crypto.PubkeyToAddress(privateKey.ToECDSA().PublicKey)
		tronAddress := make([]byte, 0)
		tronAddress = append(tronAddress, byte(0x41))
		tronAddress = append(tronAddress, ethAddress.Bytes()...)
		address = goTornSdkCommon.EncodeCheck(tronAddress)
	case common.AddressType_BSC:
		address = crypto.PubkeyToAddress(privateKey.ToECDSA().PublicKey).Hex()
	default:
		return "", ErrUnsupportedAddressType
	}
	return address, nil
}
