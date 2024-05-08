package account

import (
	"testing"

	"github.com/DwGoing/transfer_lib/common"
	goEthereumCommon "github.com/ethereum/go-ethereum/common"
)

func TestNewAccountFromMnemonic(t *testing.T) {
	var tests = []struct {
		chain    common.Chain
		mnemonic string
		password string
		index    int64
		want     any
	}{
		{common.Chain_BTC, "", "", 0, "mnemonic is invalid"},
		{common.Chain_BTC, "", "", -1, "mnemonic is invalid"},
		{common.Chain_BTC, "a b c d e f g h i j k l", "", 0, "mnemonic is invalid"},
		{common.Chain_BTC, "web firm spy fence blouse skill yard salute drink island thing poem", "", 0, "4eca5371e7471d5d969f73bdd3b2b25a95d7a740281ffa16d3877e37518a4c3ebb03cbcbba634a3532063a35c80bc0d5983f4f21fcfff1ddfdc3ec635e3db733"},
		{common.Chain_BTC, "web firm spy fence blouse skill yard salute drink island thing poem", "", -1, "index is invalid"},
	}

	for _, test := range tests {
		account, err := NewAccountFromMnemonic(test.chain, test.mnemonic, test.password, test.index)
		if err != nil {
			if err.Error() != test.want {
				t.Error(err)
			}
		} else {
			if goEthereumCommon.Bytes2Hex(account.Seed()) != test.want {
				t.Error("seed inconsistent")
			}
		}
	}
}

func TestNewAccountFromSeed(t *testing.T) {
	var tests = []struct {
		chain common.Chain
		seed  []byte
		index int64
		want  any
	}{
		{common.Chain_BTC, []byte{}, 0, "seed length must be between 128 and 512 bits"},
		{common.Chain_BTC, []byte{0x0}, 0, "seed length must be between 128 and 512 bits"},
		{
			common.Chain_BTC,
			[]byte{63, 246, 203, 185, 240, 225, 36, 163, 229, 205, 213, 143, 158, 228, 228, 216, 124, 210, 170, 182, 12, 145, 228, 90, 229, 62, 188, 127, 142, 179, 80, 3, 161, 96, 210, 204, 94, 236, 113, 11, 143, 196, 229, 50, 116, 130, 247, 147, 239, 165, 149, 40, 30, 97, 61, 178, 57, 198, 38, 43, 53, 193, 147, 98},
			0,
			"3ff6cbb9f0e124a3e5cdd58f9ee4e4d87cd2aab60c91e45ae53ebc7f8eb35003a160d2cc5eec710b8fc4e5327482f793efa595281e613db239c6262b35c19362",
		},
		{
			common.Chain_BTC,
			[]byte{63, 246, 203, 185, 240, 225, 36, 163, 229, 205, 213, 143, 158, 228, 228, 216, 124, 210, 170, 182, 12, 145, 228, 90, 229, 62, 188, 127, 142, 179, 80, 3, 161, 96, 210, 204, 94, 236, 113, 11, 143, 196, 229, 50, 116, 130, 247, 147, 239, 165, 149, 40, 30, 97, 61, 178, 57, 198, 38, 43, 53, 193, 147, 98},
			-1,
			"index is invalid",
		},
	}

	for _, test := range tests {
		account, err := NewAccountFromSeed(test.chain, test.seed, test.index)
		if err != nil {
			if err.Error() != test.want {
				t.Error()
			}
		} else {
			if goEthereumCommon.Bytes2Hex(account.Seed()) != test.want {
				t.Error()
			}
		}
	}
}
