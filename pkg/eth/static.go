package eth

import (
	"github.com/decred/dcrd/dcrec/secp256k1/v4"
	"github.com/ethereum/go-ethereum/crypto"
)

/*
@title 	获取钱包地址
@param 	privateKey	*secp256k1.PrivateKey	私钥
@return _ 			string					地址
*/
func GetAddressFromPrivateKey(privateKey *secp256k1.PrivateKey) string {
	return crypto.PubkeyToAddress(privateKey.ToECDSA().PublicKey).Hex()
}
