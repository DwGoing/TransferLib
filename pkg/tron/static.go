package tron

import (
	"github.com/decred/dcrd/dcrec/secp256k1/v4"
	"github.com/ethereum/go-ethereum/crypto"
	goTornSdkCommon "github.com/fbsobreira/gotron-sdk/pkg/common"
)

/*
@title 	获取钱包地址
@param 	privateKey	*secp256k1.PrivateKey	私钥
@return _ 			string					地址
*/
func GetAddressFromPrivateKey(privateKey *secp256k1.PrivateKey) string {
	ethAddress := crypto.PubkeyToAddress(privateKey.ToECDSA().PublicKey)
	tronAddress := make([]byte, 0)
	tronAddress = append(tronAddress, byte(0x41))
	tronAddress = append(tronAddress, ethAddress.Bytes()...)
	return goTornSdkCommon.EncodeCheck(tronAddress)
}
