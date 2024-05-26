package eth

import (
	"github.com/DwGoing/transfer_lib/crypto"
	goEthereumCrypto "github.com/ethereum/go-ethereum/crypto"
)

// Function GetAddressFromPrivateKey 从私钥获取地址
func GetAddressFromPrivateKey(privateKey []byte) string {
	return goEthereumCrypto.PubkeyToAddress(crypto.ToEcdsa(privateKey).PublicKey).Hex()
}
