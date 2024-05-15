package tron

import (
	"github.com/decred/dcrd/dcrec/secp256k1/v4"
	"github.com/ethereum/go-ethereum/crypto"
	goTornSdkCommon "github.com/fbsobreira/gotron-sdk/pkg/common"
)

// Function GetAddressFromPrivateKey 从私钥获取地址
func GetAddressFromPrivateKey(privateKey []byte) string {
	secp256k1PrivateKey := secp256k1.PrivKeyFromBytes(privateKey)
	ethAddress := crypto.PubkeyToAddress(secp256k1PrivateKey.ToECDSA().PublicKey)
	tronAddress := make([]byte, 0)
	tronAddress = append(tronAddress, byte(0x41))
	tronAddress = append(tronAddress, ethAddress.Bytes()...)
	return goTornSdkCommon.EncodeCheck(tronAddress)
}
