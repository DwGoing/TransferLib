package sol

import (
	"bytes"
	"crypto/ed25519"

	"github.com/btcsuite/btcutil/base58"
)

// Function GetAddressFromPrivateKey 从私钥获取地址
func GetAddressFromPrivateKey(privateKey []byte) (string, error) {
	hh := base58.Encode(privateKey)
	_ = hh
	privateKey = base58.Decode("2jxdp6XXn1tP8qxCiqcthy5w99VGaE5zdYKhgWk8yJePCVXZgz8ux2MK31oaGmsdNQuxb9reueCTHMLXvYC4t5MD")
	ed25519PublicKey, _, err := ed25519.GenerateKey(bytes.NewReader(privateKey))
	if err != nil {
		return "", err
	}
	return base58.Encode(ed25519PublicKey), nil
}
