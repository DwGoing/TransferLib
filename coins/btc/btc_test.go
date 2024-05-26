package btc

import (
	"testing"

	"github.com/DwGoing/transfer_lib/common"
	"github.com/btcsuite/btcd/chaincfg"
)

func TestGetAddressFromPrivateKey(t *testing.T) {
	var tests = []struct {
		privateKey  []byte
		addressType common.AddressType
		want        any
	}{
		{
			[]byte{125, 78, 193, 44, 181, 31, 253, 5, 246, 70, 243, 143, 249, 200, 157, 83, 129, 173, 111, 13, 140, 216, 22, 250, 205, 112, 69, 58, 113, 23, 162, 201},
			common.AddressType_BTC_LEGACY,
			"1L9RCzZKvMV3W3CjoCDbBSYBKgoXJmYJxE",
		},
		{
			[]byte{172, 125, 61, 232, 102, 116, 77, 25, 115, 179, 239, 147, 12, 202, 26, 19, 121, 248, 99, 71, 113, 231, 212, 246, 24, 87, 102, 174, 218, 120, 37, 111},
			common.AddressType_BTC_NATIVE_SEGWIT,
			"bc1qr4jwheasg8dnen4jjdvvypttmax4lt4duvq0l7",
		},
		{
			[]byte{88, 142, 11, 215, 238, 10, 45, 141, 42, 93, 174, 176, 91, 206, 41, 221, 45, 186, 119, 168, 82, 181, 233, 18, 230, 120, 143, 13, 110, 143, 237, 4},
			common.AddressType_BTC_NESTED_SEGWIT,
			"3Fsv8QJLy1rXcrJ6RJckUZwpizdcwJVf54",
		},
		{
			[]byte{228, 150, 33, 246, 27, 166, 126, 246, 29, 210, 107, 222, 42, 15, 139, 165, 198, 198, 35, 10, 146, 252, 37, 124, 230, 167, 229, 102, 23, 181, 175, 213},
			common.AddressType_BTC_TAPROOT,
			"bc1pytpzkvlgef75jnh8gz8nfa2des7vd0sg9r2taljjtfckaqkfmqfq0hxa2a",
		},
	}

	for _, test := range tests {
		address, err := GetAddressFromPrivateKey(test.privateKey, test.addressType, &chaincfg.MainNetParams)
		if err != nil {
			if err.Error() != test.want {
				t.Errorf("error not match ===> %s", err.Error())
			}
			continue
		} else {
			if address != test.want {
				t.Errorf("address not match ===> %s", address)
			}
		}
	}
}

func TestGetAddressFromPublicKey(t *testing.T) {
	var tests = []struct {
		publicKey   []byte
		addressType common.AddressType
		want        any
	}{
		{
			[]byte{3, 68, 104, 114, 122, 52, 150, 151, 78, 125, 169, 105, 141, 154, 152, 63, 120, 169, 12, 63, 219, 27, 240, 143, 68, 80, 43, 18, 178, 53, 180, 42, 201},
			common.AddressType_BTC_LEGACY,
			"1L9RCzZKvMV3W3CjoCDbBSYBKgoXJmYJxE",
		},
		{
			[]byte{2, 102, 113, 113, 201, 252, 251, 187, 106, 205, 88, 29, 72, 118, 202, 179, 133, 111, 231, 255, 33, 148, 19, 185, 242, 24, 223, 235, 75, 30, 223, 185, 3},
			common.AddressType_BTC_NATIVE_SEGWIT,
			"bc1qr4jwheasg8dnen4jjdvvypttmax4lt4duvq0l7",
		},
		{
			[]byte{2, 20, 203, 216, 198, 83, 230, 78, 66, 86, 246, 88, 52, 98, 8, 106, 75, 248, 114, 24, 57, 121, 186, 198, 111, 167, 188, 235, 23, 128, 230, 84, 166},
			common.AddressType_BTC_NESTED_SEGWIT,
			"3Fsv8QJLy1rXcrJ6RJckUZwpizdcwJVf54",
		},
		{
			[]byte{2, 46, 222, 89, 130, 115, 13, 182, 163, 71, 190, 12, 206, 171, 138, 2, 241, 80, 239, 91, 130, 31, 64, 201, 233, 71, 209, 240, 174, 127, 88, 50, 246},
			common.AddressType_BTC_TAPROOT,
			"bc1pytpzkvlgef75jnh8gz8nfa2des7vd0sg9r2taljjtfckaqkfmqfq0hxa2a",
		},
	}

	for _, test := range tests {
		address, err := GetAddressFromPublicKey(test.publicKey, test.addressType, &chaincfg.MainNetParams)
		if err != nil {
			if err.Error() != test.want {
				t.Errorf("error not match ===> %s", err.Error())
			}
		} else {
			if address != test.want {
				t.Errorf("address not match ===> %s", address)
			}
		}
	}
}
