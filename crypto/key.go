package crypto

import (
	"strconv"
	"strings"

	"github.com/DwGoing/transfer_lib/common"
	"github.com/tyler-smith/go-bip32"
	"github.com/tyler-smith/go-bip39"
)

// Function NewSeedFromMnemonic 从助记词获取种子
func NewSeedFromMnemonic(mnemonic string, password string) ([]byte, error) {
	if !bip39.IsMnemonicValid(mnemonic) {
		return nil, common.ErrInvalidMnemonic
	}
	seed, err := bip39.NewSeedWithErrorChecking(mnemonic, password)
	if err != nil {
		return nil, err
	}
	return seed, nil
}

// Function NewPrivateKeyFromSeed 从种子中生成私钥
func NewPrivateKeyFromSeed(seed []byte) ([]byte, error) {
	masterKey, err := bip32.NewMasterKey(seed)
	if err != nil {
		return nil, err
	}
	return masterKey.Key, nil
}

// Function NewPrivateKeyFromSeedByPath 从种子中生成子私钥
func NewPrivateKeyFromSeedByPath(seed []byte, path string) ([]byte, error) {
	masterKey, err := bip32.NewMasterKey(seed)
	if err != nil {
		return nil, err
	}
	arr := strings.Split(path, "/")
	currentKey := masterKey
	for _, part := range arr {
		if part == "m" {
			continue
		}
		var harden = false
		if strings.HasSuffix(part, "'") {
			harden = true
			part = strings.TrimSuffix(part, "'")
		}
		id, err := strconv.ParseUint(part, 10, 31)
		if err != nil {
			return nil, err
		}
		var uid = uint32(id)
		if harden {
			uid |= bip32.FirstHardenedChild
		}
		newKey, err := currentKey.NewChildKey(uid)
		if err != nil {
			return nil, err
		}
		currentKey = newKey
	}
	return currentKey.Key, nil
}
