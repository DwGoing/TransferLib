package account

import (
	"testing"

	"github.com/btcsuite/btcd/chaincfg"
	"github.com/ethereum/go-ethereum/common"
)

func TestGetSeedFromMnemonic(t *testing.T) {
	var tests = []struct {
		mnemonic string
		password string
		want     any
	}{
		{"", "", "mnemonic is invalid"},
		{"a b c d e f g h i j k l", "", "mnemonic is invalid"},
		{
			"web firm spy fence blouse skill yard salute drink island thing poem",
			"0",
			"55ccd9ac3255467b5e3a25a91ff6f3f791c0fb08a8a21e601788ffec076a731b6960f7ced5662393aef5e91f5e38f234fc57a0f422eef24bd91e70f4cb976ebb",
		},
		{
			"web firm spy fence blouse skill yard salute drink island thing poem",
			"",
			"4eca5371e7471d5d969f73bdd3b2b25a95d7a740281ffa16d3877e37518a4c3ebb03cbcbba634a3532063a35c80bc0d5983f4f21fcfff1ddfdc3ec635e3db733",
		},
	}

	for _, test := range tests {
		seed, err := GetSeedFromMnemonic(test.mnemonic, test.password)
		if err != nil {
			if err.Error() != test.want {
				t.Error(err)
			}
		} else {
			hex := common.Bytes2Hex(seed)
			if hex != test.want {
				t.Errorf("not match want: %s", hex)
			}
		}
	}
}

func TestDerivePrivateKey(t *testing.T) {
	var tests = []struct {
		seed   []byte
		params *chaincfg.Params
		path   string
		want   any
	}{
		{[]byte{}, &chaincfg.MainNetParams, "", "seed length must be between 128 and 512 bits"},
		{
			[]byte{63, 246, 203, 185, 240, 225, 36, 163, 229, 205, 213, 143, 158, 228, 228, 216, 124, 210, 170, 182, 12, 145, 228, 90, 229, 62, 188, 127, 142, 179, 80, 3, 161, 96, 210, 204, 94, 236, 113, 11, 143, 196, 229, 50, 116, 130, 247, 147, 239, 165, 149, 40, 30, 97, 61, 178, 57, 198, 38, 43, 53, 193, 147, 98},
			&chaincfg.MainNetParams,
			"",
			"ambiguous path: use 'm/' prefix for absolute paths, or no leading '/' for relative ones",
		},
		{
			[]byte{63, 246, 203, 185, 240, 225, 36, 163, 229, 205, 213, 143, 158, 228, 228, 216, 124, 210, 170, 182, 12, 145, 228, 90, 229, 62, 188, 127, 142, 179, 80, 3, 161, 96, 210, 204, 94, 236, 113, 11, 143, 196, 229, 50, 116, 130, 247, 147, 239, 165, 149, 40, 30, 97, 61, 178, 57, 198, 38, 43, 53, 193, 147, 98},
			&chaincfg.MainNetParams,
			"m/44'/60'/0'/0/",
			"invalid component: ",
		},
	}

	for _, test := range tests {
		key, err := DerivePrivateKey(test.seed, test.params, test.path)
		if err != nil {
			if err.Error() != test.want {
				t.Error(err)
			}
		} else {
			publicKey, err := key.ECPubKey()
			if err != nil {
				if err.Error() != test.want {
					t.Error(err)
				}
			}
			if common.Bytes2Hex(publicKey.SerializeCompressed()) != test.want {
				t.Error()
			}
		}
	}
}
