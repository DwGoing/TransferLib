package crypto

import (
	"encoding/hex"
	"fmt"
	"testing"

	"github.com/DwGoing/transfer_lib/common"
	"github.com/btcsuite/btcd/btcutil/base58"
)

func TestNewSeedFromMnemonic(t *testing.T) {
	var tests = []struct {
		mnemonic string
		password string
		want     any
	}{
		{
			"web firm spy fence blouse skill yard salute drink island thing poem",
			"",
			"4eca5371e7471d5d969f73bdd3b2b25a95d7a740281ffa16d3877e37518a4c3ebb03cbcbba634a3532063a35c80bc0d5983f4f21fcfff1ddfdc3ec635e3db733",
		},
	}

	for _, test := range tests {
		seed, err := NewSeedFromMnemonic(test.mnemonic, test.password)
		if err != nil {
			if err.Error() != test.want {
				t.Errorf("error not match ===> %s", err.Error())
			}
			continue
		}
		if hex.EncodeToString(seed) != test.want {
			t.Errorf("seed not match ===> %s", hex.EncodeToString(seed))
		}
	}
}

func TestNewPrivateKeyFromSeedByPath(t *testing.T) {
	var tests = []struct {
		seed        []byte
		addressType common.AddressType
		want        any
	}{
		{
			[]byte{78, 202, 83, 113, 231, 71, 29, 93, 150, 159, 115, 189, 211, 178, 178, 90, 149, 215, 167, 64, 40, 31, 250, 22, 211, 135, 126, 55, 81, 138, 76, 62, 187, 3, 203, 203, 186, 99, 74, 53, 50, 6, 58, 53, 200, 11, 192, 213, 152, 63, 79, 33, 252, 255, 241, 221, 253, 195, 236, 99, 94, 61, 183, 51},
			common.AddressType_BTC_LEGACY,
			"L1RHvPZeciCrGZM4vuJayB2sbsesPApvq1nnifg72LkiKqvFTwyk",
		},
		{
			[]byte{78, 202, 83, 113, 231, 71, 29, 93, 150, 159, 115, 189, 211, 178, 178, 90, 149, 215, 167, 64, 40, 31, 250, 22, 211, 135, 126, 55, 81, 138, 76, 62, 187, 3, 203, 203, 186, 99, 74, 53, 50, 6, 58, 53, 200, 11, 192, 213, 152, 63, 79, 33, 252, 255, 241, 221, 253, 195, 236, 99, 94, 61, 183, 51},
			common.AddressType_ETH,
			"90ba72b1281a54b5f14dfbe4dcdb3fba383ddf1f35dc61d3b2c1fb7579ed536c",
		},
	}

	for _, test := range tests {
		path, err := test.addressType.Path()
		if err != nil {
			if err.Error() != test.want {
				t.Errorf("error not match ===> %s", err.Error())
			}
			continue
		}
		privateKeyBytes, err := NewPrivateKeyFromSeedByPath(test.seed, fmt.Sprintf("%s%d", path, 0))
		if err != nil {
			if err.Error() != test.want {
				t.Errorf("error not match ===> %s", err.Error())
			}
			continue
		}
		switch test.addressType {
		case common.AddressType_BTC_LEGACY, common.AddressType_BTC_NESTED_SEGWIT, common.AddressType_BTC_NATIVE_SEGWIT, common.AddressType_BTC_TAPROOT:
			privateKeyBase58 := base58.Encode(privateKeyBytes)
			if privateKeyBase58 != test.want {
				t.Errorf("private key hex not match ===> %s", privateKeyBase58)
			}
		case common.AddressType_ETH, common.AddressType_BSC:
			privateKeyHex := hex.EncodeToString(privateKeyBytes)
			if privateKeyHex != test.want {
				t.Errorf("private key hex not match ===> %s", privateKeyHex)
			}
		}
	}
}
