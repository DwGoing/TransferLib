package btc

import (
	"testing"

	"github.com/DwGoing/transfer_lib/common"
	goEthereumCommon "github.com/ethereum/go-ethereum/common"
)

func TestNewAccountFromSeed(t *testing.T) {
	var tests = []struct {
		seed  []byte
		index int64
		want  any
	}{
		{[]byte{}, 0, "seed length must be between 128 and 512 bits"},
		{[]byte{0x0}, 0, "seed length must be between 128 and 512 bits"},
		{
			[]byte{63, 246, 203, 185, 240, 225, 36, 163, 229, 205, 213, 143, 158, 228, 228, 216, 124, 210, 170, 182, 12, 145, 228, 90, 229, 62, 188, 127, 142, 179, 80, 3, 161, 96, 210, 204, 94, 236, 113, 11, 143, 196, 229, 50, 116, 130, 247, 147, 239, 165, 149, 40, 30, 97, 61, 178, 57, 198, 38, 43, 53, 193, 147, 98},
			0,
			"3ff6cbb9f0e124a3e5cdd58f9ee4e4d87cd2aab60c91e45ae53ebc7f8eb35003a160d2cc5eec710b8fc4e5327482f793efa595281e613db239c6262b35c19362",
		},
		{
			[]byte{63, 246, 203, 185, 240, 225, 36, 163, 229, 205, 213, 143, 158, 228, 228, 216, 124, 210, 170, 182, 12, 145, 228, 90, 229, 62, 188, 127, 142, 179, 80, 3, 161, 96, 210, 204, 94, 236, 113, 11, 143, 196, 229, 50, 116, 130, 247, 147, 239, 165, 149, 40, 30, 97, 61, 178, 57, 198, 38, 43, 53, 193, 147, 98},
			-1,
			"index is invalid",
		},
	}

	for _, test := range tests {
		account, err := NewAccountFromSeed(test.seed, test.index)
		if err != nil {
			if err.Error() != test.want {
				t.Error()
			}
		} else {
			if account.account.Index() != test.index {
				t.Error("index error")
			}
			if goEthereumCommon.Bytes2Hex(account.account.Seed()) != test.want {
				t.Error()
			}
		}
	}
}

func TestNewAccountFromMnemonicGetAddress(t *testing.T) {
	var tests = []struct {
		mnemonic    string
		password    string
		index       int64
		addressType common.AddressType
		want        any
	}{
		{"unfair enter female during swift radar sell since cross suit buddy slam", "", 0, common.AddressType_BTC_LEGACY, "12hvu9ojYzbrZhSAMjLnCtM14cfuon5hxG"},
		{"unfair enter female during swift radar sell since cross suit buddy slam", "", 0, common.AddressType_BTC_NESTED_SEGWIT, "3KbHiAZpJ1svckinbxd5Sc4QmDv4Zromsw"},
		{"unfair enter female during swift radar sell since cross suit buddy slam", "", 0, common.AddressType_BTC_NATIVE_SEGWIT, "bc1qnkdqwzemsravyhxgyncmsxrucquhd334lzn70e"},
		{"unfair enter female during swift radar sell since cross suit buddy slam", "", 0, common.AddressType_BTC_TAPROOT, "bc1plqqf8e664mysftkr5m7rdnctl6ss6nc57qqm2fle49alr8rq3vgq2d6lpq"},
	}

	for _, test := range tests {
		account, err := NewAccountFromMnemonic(test.mnemonic, test.password, test.index)
		if err != nil {
			if err.Error() != test.want {
				t.Error()
			}
		} else {
			address, err := account.GetAddress(test.addressType)
			if err != nil {
				if err.Error() != test.want {
					t.Error()
				}
			}
			if address != test.want {
				t.Errorf("address not match, %s %s", address, test.want)
			}
		}
	}
}
