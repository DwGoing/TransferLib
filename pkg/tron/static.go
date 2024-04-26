package tron

import (
	"crypto/ecdsa"

	"github.com/ethereum/go-ethereum/crypto"
	goTornSdkCommon "github.com/fbsobreira/gotron-sdk/pkg/common"
)

/*
@title 	获取钱包地址
@param 	privateKey	*ecdsa.PrivateKey	私钥
@return _ 			string					地址
*/
func GetAddressFromPrivateKey(privateKey *ecdsa.PrivateKey) string {
	ethAddress := crypto.PubkeyToAddress(privateKey.PublicKey)
	tronAddress := make([]byte, 0)
	tronAddress = append(tronAddress, byte(0x41))
	tronAddress = append(tronAddress, ethAddress.Bytes()...)
	return goTornSdkCommon.EncodeCheck(tronAddress)
}
