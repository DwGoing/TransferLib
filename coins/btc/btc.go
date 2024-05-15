package btc

import (
	"github.com/DwGoing/transfer_lib/common"
	"github.com/btcsuite/btcd/btcec/v2"
	"github.com/btcsuite/btcd/btcutil"
	"github.com/btcsuite/btcd/chaincfg"
	"github.com/btcsuite/btcd/txscript"
	"github.com/decred/dcrd/dcrec/secp256k1/v4"
)

// Function GetAddressFromPrivateKey 从私钥获取地址
func GetAddressFromPrivateKey(privateKey []byte, addressType common.AddressType, network *chaincfg.Params) (string, error) {
	secp256k1PrivateKey := secp256k1.PrivKeyFromBytes(privateKey)
	x := secp256k1PrivateKey.PubKey().SerializeCompressed()
	_ = x
	return GetAddressFromPublicKey(secp256k1PrivateKey.PubKey().SerializeCompressed(), addressType, network)
}

// Function GetAddressFromPublicKey 从公钥获取地址
func GetAddressFromPublicKey(publicKey []byte, addressType common.AddressType, network *chaincfg.Params) (string, error) {
	if network == nil {
		network = &chaincfg.MainNetParams
	}
	switch addressType {
	case common.AddressType_BTC_LEGACY:
		p2pkh, err := btcutil.NewAddressPubKey(publicKey, network)
		if err != nil {
			return "", err
		}
		return p2pkh.EncodeAddress(), nil
	case common.AddressType_BTC_NATIVE_SEGWIT:
		p2wpkh, err := btcutil.NewAddressWitnessPubKeyHash(btcutil.Hash160(publicKey), network)
		if err != nil {
			return "", err
		}
		return p2wpkh.EncodeAddress(), nil
	case common.AddressType_BTC_NESTED_SEGWIT:
		p2wpkh, err := btcutil.NewAddressWitnessPubKeyHash(btcutil.Hash160(publicKey), network)
		if err != nil {
			return "", err
		}
		redeemScript, err := txscript.PayToAddrScript(p2wpkh)
		if err != nil {
			return "", err
		}
		p2sh, err := btcutil.NewAddressScriptHash(redeemScript, network)
		if err != nil {
			return "", err
		}
		return p2sh.EncodeAddress(), nil
	case common.AddressType_BTC_TAPROOT:
		internalKey, err := btcec.ParsePubKey(publicKey)
		if err != nil {
			return "", err
		}
		p2tr, err := btcutil.NewAddressTaproot(txscript.ComputeTaprootKeyNoScript(internalKey).SerializeCompressed()[1:], network)
		if err != nil {
			return "", err
		}
		return p2tr.EncodeAddress(), nil
	default:
		return "", common.ErrUnsupportedAddressType
	}
}
