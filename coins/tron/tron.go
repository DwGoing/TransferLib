package tron

import (
	"github.com/DwGoing/transfer_lib/crypto"
	goTornSdkCrypto "github.com/ethereum/go-ethereum/crypto"
	goTornSdkCommon "github.com/fbsobreira/gotron-sdk/pkg/common"
)

// Function GetAddressFromPrivateKey 从私钥获取地址
func GetAddressFromPrivateKey(privateKey []byte) string {
	ethAddress := goTornSdkCrypto.PubkeyToAddress(crypto.ToEcdsa(privateKey).PublicKey)
	tronAddress := make([]byte, 0)
	tronAddress = append(tronAddress, byte(0x41))
	tronAddress = append(tronAddress, ethAddress.Bytes()...)
	return goTornSdkCommon.EncodeCheck(tronAddress)
}
