package tron

import "testing"

func TestGetAddressFromPrivateKey(t *testing.T) {
	var tests = []struct {
		privateKey []byte
		want       any
	}{
		{
			[]byte{94, 182, 114, 244, 134, 110, 88, 242, 86, 201, 89, 20, 197, 249, 106, 74, 140, 161, 241, 23, 70, 92, 84, 64, 35, 233, 203, 11, 178, 56, 122, 101},
			"TYq73v8nCqi85g5CJxSNYDW5QKvaffAuPh",
		},
	}

	for _, test := range tests {
		address := GetAddressFromPrivateKey(test.privateKey)
		if address != test.want {
			t.Errorf("address not match ===> %s", address)
		}
	}
}
