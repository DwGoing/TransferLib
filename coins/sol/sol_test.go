package sol

import (
	"testing"

	"github.com/btcsuite/btcd/btcutil/base58"
	"github.com/gagliardetto/solana-go"
)

func TestGetAddressFromPrivateKey(t *testing.T) {
	// s, _ := crypto.NewSeedFromMnemonic("web firm spy fence blouse skill yard salute drink island thing poem", "")
	// pk, _ := crypto.NewPrivateKeyFromSeedByPath(s, "m/44'/501'/0'/0'/0")
	w, _ := solana.WalletFromPrivateKeyBase58("2jxdp6XXn1tP8qxCiqcthy5w99VGaE5zdYKhgWk8yJePCVXZgz8ux2MK31oaGmsdNQuxb9reueCTHMLXvYC4t5MD")
	add := base58.Encode(w.PrivateKey.PublicKey().Bytes())
	_ = add

	var tests = []struct {
		privateKey []byte
		want       any
	}{
		{
			[]byte{14, 253, 232, 28, 18, 241, 157, 15, 133, 123, 204, 88, 173, 37, 185, 205, 209, 25, 230, 88, 186, 27, 245, 221, 85, 92, 163, 126, 114, 37, 74, 193},
			// []byte{87, 16, 58, 217, 145, 66, 55, 60, 174, 181, 93, 236, 168, 156, 98, 140, 30, 59, 139, 211, 12, 90, 89, 55, 62, 55, 26, 129, 145, 105, 136, 192, 142, 146, 201, 104, 52, 81, 201, 64, 2, 185, 90, 226, 89, 42, 57, 117, 247, 101, 211, 72, 170, 60, 251, 250, 244, 143, 184, 7, 217, 196, 78, 108},
			"",
		},
	}

	for _, test := range tests {
		address, err := GetAddressFromPrivateKey(test.privateKey)
		if err != nil {
			if err.Error() != test.want {
				t.Errorf("error not math ===> %s", err.Error())
			}
			continue
		}
		if address != test.want {
			t.Errorf("address not math ===> %s", address)
		}
	}
}
