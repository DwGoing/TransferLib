package sol

import (
	"crypto/ed25519"
	"encoding/hex"

	"github.com/btcsuite/btcutil/base58"
)

// Function GetAddressFromPrivateKey 从私钥获取地址
func GetAddressFromPrivateKey(privateKey []byte) (string, error) {
	ed25519PrivateKey := ed25519.NewKeyFromSeed(privateKey)
	x := hex.EncodeToString(ed25519PrivateKey)
	_ = x
	return base58.Encode(ed25519PrivateKey.Public().(ed25519.PublicKey)), nil
}
