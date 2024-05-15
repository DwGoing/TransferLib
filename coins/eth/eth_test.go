package eth

import (
	"testing"
)

func TestGetAddressFromPrivateKey(t *testing.T) {
	var tests = []struct {
		privateKey []byte
		want       any
	}{
		{
			[]byte{144, 186, 114, 177, 40, 26, 84, 181, 241, 77, 251, 228, 220, 219, 63, 186, 56, 61, 223, 31, 53, 220, 97, 211, 178, 193, 251, 117, 121, 237, 83, 108},
			"0x69DB56F110a80101c2307e10563a3Dd45653bC8f",
		},
	}

	for _, test := range tests {
		address := GetAddressFromPrivateKey(test.privateKey)
		if address != test.want {
			t.Errorf("address not match ===> %s", address)
		}
	}
}
