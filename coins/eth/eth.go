package eth

import (
	"github.com/decred/dcrd/dcrec/secp256k1/v4"
	"github.com/ethereum/go-ethereum/crypto"
)

// Function GetAddressFromPrivateKey 从私钥获取地址
func GetAddressFromPrivateKey(privateKey []byte) string {
	secp256k1PrivateKey := secp256k1.PrivKeyFromBytes(privateKey)
	return crypto.PubkeyToAddress(secp256k1PrivateKey.ToECDSA().PublicKey).Hex()
}
