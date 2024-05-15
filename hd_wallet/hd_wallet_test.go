package hd_wallet

import (
	"encoding/hex"
	"fmt"
	"testing"

	"github.com/DwGoing/transfer_lib/coins/bsc"
	"github.com/DwGoing/transfer_lib/coins/eth"
	"github.com/DwGoing/transfer_lib/coins/tron"
	"github.com/DwGoing/transfer_lib/common"
	"github.com/btcsuite/btcd/btcutil/base58"
	"github.com/btcsuite/btcd/chaincfg"
)

func TestNewHDWalletFromMnemonic(t *testing.T) {
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
		wallet, err := NewHDWalletFromMnemonic(test.mnemonic, test.password)
		if err != nil {
			if err.Error() != test.want {
				t.Errorf("error not match ===> %s", err.Error())
			}
			continue
		} else {
			privateKeyHex := hex.EncodeToString(wallet.Seed())
			if privateKeyHex != test.want {
				t.Errorf("private key hex not match ===> %s", privateKeyHex)
			}
		}
	}
}

func TestGetPrivateKey(t *testing.T) {
	wallet, err := NewHDWalletFromMnemonic("web firm spy fence blouse skill yard salute drink island thing poem", "")
	if err != nil {
		t.Errorf("recover wallet failed: %s", err.Error())
		return
	}

	var tests = []struct {
		addressType common.AddressType
		want        any
	}{
		// {
		// 	common.AddressType_BTC_LEGACY,
		// 	"L1RHvPZeciCrGZM4vuJayB2sbsesPApvq1nnifg72LkiKqvFTwyk",
		// },
		{
			common.AddressType_ETH,
			"90ba72b1281a54b5f14dfbe4dcdb3fba383ddf1f35dc61d3b2c1fb7579ed536c",
		},
		{
			common.AddressType_BSC,
			"90ba72b1281a54b5f14dfbe4dcdb3fba383ddf1f35dc61d3b2c1fb7579ed536c",
		},
		{
			common.AddressType_TRON,
			"553daccb873dfda98c79949d4922563d4e6c496829ec8afa1cadc44987ec4c27",
		},
		// {
		// 	common.AddressType_SOL,
		// 	0,
		// 	"2jxdp6XXn1tP8qxCiqcthy5w99VGaE5zdYKhgWk8yJePCVXZgz8ux2MK31oaGmsdNQuxb9reueCTHMLXvYC4t5MD",
		// },
	}

	for _, test := range tests {
		path, err := test.addressType.Path()
		if err != nil {
			if err.Error() != test.want {
				t.Errorf("error not match ===> %s", err.Error())
			}
			continue
		}
		path = fmt.Sprintf("%s%d", path, 0)
		privateKey, err := wallet.GetPrivateKeyByPath(path)
		if err != nil {
			if err.Error() != test.want {
				t.Errorf("error not match ===> %s", err.Error())
			}
			continue
		}
		switch test.addressType {
		case common.AddressType_BTC_LEGACY, common.AddressType_BTC_NESTED_SEGWIT, common.AddressType_BTC_NATIVE_SEGWIT, common.AddressType_BTC_TAPROOT, common.AddressType_SOL:
			privateKeyBase58 := base58.Encode(privateKey)
			if privateKeyBase58 != test.want {
				t.Errorf("private key hex not match ===> %s", privateKeyBase58)
			}
		case common.AddressType_ETH, common.AddressType_BSC:
			privateKeyHex := hex.EncodeToString(privateKey)
			if privateKeyHex != test.want {
				t.Errorf("private key hex not match ===> %s", privateKeyHex)
			}
		}
	}
}

func TestGetAddress(t *testing.T) {
	wallet, err := NewHDWalletFromMnemonic("web firm spy fence blouse skill yard salute drink island thing poem", "")
	if err != nil {
		t.Errorf("recover wallet failed: %s", err.Error())
		return
	}

	var tests = []struct {
		addressType common.AddressType
		args        any
		want        any
	}{
		{
			common.AddressType_BTC_LEGACY,
			chaincfg.MainNetParams,
			"1L9RCzZKvMV3W3CjoCDbBSYBKgoXJmYJxE",
		},
		{
			common.AddressType_BTC_NATIVE_SEGWIT,
			chaincfg.MainNetParams,
			"bc1qr4jwheasg8dnen4jjdvvypttmax4lt4duvq0l7",
		},
		{
			common.AddressType_BTC_NESTED_SEGWIT,
			chaincfg.MainNetParams,
			"3Fsv8QJLy1rXcrJ6RJckUZwpizdcwJVf54",
		},
		{
			common.AddressType_BTC_TAPROOT,
			chaincfg.MainNetParams,
			"bc1pytpzkvlgef75jnh8gz8nfa2des7vd0sg9r2taljjtfckaqkfmqfq0hxa2a",
		},
		{
			common.AddressType_ETH,
			nil,
			"0x69DB56F110a80101c2307e10563a3Dd45653bC8f",
		},
		{
			common.AddressType_BSC,
			nil,
			"0x69DB56F110a80101c2307e10563a3Dd45653bC8f",
		},
		{
			common.AddressType_TRON,
			nil,
			"TAAsVD38kq6MehYz4BTSJKP42jn4dQjco6",
		},
		{
			common.AddressType_SOL,
			nil,
			"AbYiJFYr7rPUnYCvRZ9DwPMxog3gkWGtEo8JDFgic6kK",
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
		path = fmt.Sprintf("%s%d", path, 0)
		address, err := wallet.GetAddress(path, test.addressType, test.args)
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

func TestGetCurrentHeight(t *testing.T) {
	wallet, err := NewHDWalletFromMnemonic("web firm spy fence blouse skill yard salute drink island thing poem", "")
	if err != nil {
		t.Errorf("recover wallet failed: %s", err.Error())
		return
	}

	var tests = []struct {
		chainType common.ChainType
		args      any
		want      any
	}{
		{
			common.ChainType_ETH,
			[]eth.Node{{Host: "https://1rpc.io/holesky", Weight: 100}},
			nil,
		},
		{
			common.ChainType_BSC,
			[]bsc.Node{{Host: "https://data-seed-prebsc-2-s1.binance.org:8545", Weight: 100}},
			nil,
		},
		{
			common.ChainType_TRON,
			[]tron.Node{{Host: "grpc.nile.trongrid.io:50051", Weight: 100, ApiKeys: []string{"d9b77ec9-39e0-4765-98d8-2c59188344a0"}}},
			nil,
		},
	}

	for _, test := range tests {
		err := wallet.NewClient(test.chainType, test.args)
		if err != nil {
			if err.Error() != test.want {
				t.Errorf("error not match ===> %s", err.Error())
			}
			continue
		}
		height, err := wallet.GetCurrentHeight(test.chainType)
		if err != nil {
			if err.Error() != test.want {
				t.Errorf("error not match ===> %s", err.Error())
			}
			continue
		}
		name, _ := test.chainType.Name()
		fmt.Printf("%s height ===> %d\n", name, height)
	}
}

func TestGetBalance(t *testing.T) {
	wallet, err := NewHDWalletFromMnemonic("web firm spy fence blouse skill yard salute drink island thing poem", "")
	if err != nil {
		t.Errorf("recover wallet failed: %s", err.Error())
		return
	}

	var tests = []struct {
		chainType common.ChainType
		args      any
		address   string
		token     string
		want      any
	}{
		{
			common.ChainType_ETH,
			[]eth.Node{{Host: "https://1rpc.io/holesky", Weight: 100}},
			"0xbb03D2098FAa5867FA3381c9b1CB95F45477916E",
			"",
			nil,
		},
		{
			common.ChainType_ETH,
			[]eth.Node{{Host: "https://1rpc.io/holesky", Weight: 100}},
			"0xbb03D2098FAa5867FA3381c9b1CB95F45477916E",
			"4555Ed1F6D9cb6CC1D52BB88C7525b17a06da0Dd",
			nil,
		},
		{
			common.ChainType_BSC,
			[]bsc.Node{{Host: "https://data-seed-prebsc-2-s1.binance.org:8545", Weight: 100}},
			"0xbb03d2098faa5867fa3381c9b1cb95f45477916e",
			"",
			nil,
		},
		{
			common.ChainType_BSC,
			[]bsc.Node{{Host: "https://data-seed-prebsc-2-s1.binance.org:8545", Weight: 100}},
			"0xbb03d2098faa5867fa3381c9b1cb95f45477916e",
			"337610d27c682E347C9cD60BD4b3b107C9d34dDd",
			nil,
		},
		{
			common.ChainType_TRON,
			[]tron.Node{{Host: "grpc.nile.trongrid.io:50051", Weight: 100, ApiKeys: []string{"d9b77ec9-39e0-4765-98d8-2c59188344a0"}}},
			"TYq73v8nCqi85g5CJxSNYDW5QKvaffAuPh",
			"",
			nil,
		},
		{
			common.ChainType_TRON,
			[]tron.Node{{Host: "grpc.nile.trongrid.io:50051", Weight: 100, ApiKeys: []string{"d9b77ec9-39e0-4765-98d8-2c59188344a0"}}},
			"TYq73v8nCqi85g5CJxSNYDW5QKvaffAuPh",
			"TXLAQ63Xg1NAzckPwKHvzw7CSEmLMEqcdj",
			nil,
		},
	}

	for _, test := range tests {
		err := wallet.NewClient(test.chainType, test.args)
		if err != nil {
			if err.Error() != test.want {
				t.Errorf("error not match ===> %s", err.Error())
			}
			continue
		}
		balance, err := wallet.GetBalance(test.chainType, test.address, test.token)
		if err != nil {
			if err.Error() != test.want {
				t.Errorf("error not match ===> %s", err.Error())
			}
			continue
		}
		fmt.Printf("%s %s balance ===> %f\n", test.address, test.token, balance)
	}
}

func TestTransfer(t *testing.T) {
	wallet, err := NewHDWalletFromMnemonic("web firm spy fence blouse skill yard salute drink island thing poem", "")
	if err != nil {
		t.Errorf("recover wallet failed: %s", err.Error())
		return
	}

	var tests = []struct {
		chainType  common.ChainType
		args       any
		privateKey []byte
		to         string
		token      string
		value      float64
		want       any
	}{
		{
			common.ChainType_ETH,
			[]eth.Node{{Host: "https://1rpc.io/holesky", Weight: 100}},
			[]byte{81, 45, 142, 172, 240, 116, 222, 161, 154, 166, 248, 153, 84, 24, 7, 107, 69, 25, 22, 109, 236, 37, 144, 212, 167, 166, 110, 28, 247, 243, 32, 187},
			"0xbb03D2098FAa5867FA3381c9b1CB95F45477916E",
			"",
			0.5,
			nil,
		},
	}

	for _, test := range tests {
		err := wallet.NewClient(test.chainType, test.args)
		if err != nil {
			if err.Error() != test.want {
				t.Errorf("error not match ===> %s", err.Error())
			}
			continue
		}
		hash, err := wallet.Transfer(test.chainType, test.privateKey, test.to, test.token, test.value)
		if err != nil {
			if err.Error() != test.want {
				t.Errorf("error not match ===> %s", err.Error())
			}
			continue
		}
		fmt.Printf("hash ===> %s\n", hash)
	}
}

func TestGetTransaction(t *testing.T) {
	wallet, err := NewHDWalletFromMnemonic("web firm spy fence blouse skill yard salute drink island thing poem", "")
	if err != nil {
		t.Errorf("recover wallet failed: %s", err.Error())
		return
	}

	var tests = []struct {
		chainType common.ChainType
		args      any
		txHash    string
		want      any
	}{
		{
			common.ChainType_ETH,
			[]eth.Node{{Host: "https://1rpc.io/holesky", Weight: 100}},
			"0xb2c4b35a1c9796ee6e5aebad44efbcad60d2102edb2fa554d3393b7afd125166",
			nil,
		},
		{
			common.ChainType_ETH,
			[]eth.Node{{Host: "https://1rpc.io/holesky", Weight: 100}},
			"0xbede7b0b4ed6ebc4ff603f87ecd7c074e7bbdc17a3c580775181881578d5fe0a",
			nil,
		},
	}

	for _, test := range tests {
		err := wallet.NewClient(test.chainType, test.args)
		if err != nil {
			if err.Error() != test.want {
				t.Errorf("error not match ===> %s", err.Error())
			}
			continue
		}
		tx, err := wallet.GetTransaction(test.chainType, test.txHash)
		if err != nil {
			if err.Error() != test.want {
				t.Errorf("error not match ===> %s", err.Error())
			}
			continue
		}
		fmt.Printf("Tx ===> %+v\n", tx)
	}
}
