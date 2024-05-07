package eth

import (
	"fmt"
	"testing"

	"github.com/DwGoing/transfer_lib/account"
	"github.com/DwGoing/transfer_lib/chain"
)

var client = NewChainClient(
	[]Node{{Node: chain.Node{Host: "https://1rpc.io/holesky", Weight: 100}}},
	map[string]Currency{
		"ETH": {
			Contract: "",
			Decimals: 18,
		},
		"USDT": {
			Contract: "4555Ed1F6D9cb6CC1D52BB88C7525b17a06da0Dd",
			Decimals: 18,
		},
	},
)

func TestGetCurrentHeight(t *testing.T) {
	height, err := client.GetCurrentHeight()
	if err != nil {
		t.Error(err)
	}
	fmt.Printf("Height ===> %d\n", height)
}

func TestGetBalance(t *testing.T) {
	balance, err := client.GetBalance("0xbb03D2098FAa5867FA3381c9b1CB95F45477916E", "ETH", nil)
	if err != nil {
		t.Error(err)
	}
	fmt.Printf("ETH Balance ===> %f\n", balance)
	balance, err = client.GetBalance("0xbb03D2098FAa5867FA3381c9b1CB95F45477916E", "USDT", nil)
	if err != nil {
		t.Error(err)
	}
	fmt.Printf("USDT Balance ===> %f\n", balance)
}

func TestTransferETH(t *testing.T) {
	pk, err := account.GetPrivateKeyFromHex("512d8eacf074dea19aa6f8995418076b4519166dec2590d4a7a66e1cf7f320bb")
	if err != nil {
		t.Error(err)
	}
	hash, err := client.Transfer(pk, "0x337610d27c682E347C9cD60BD4b3b107C9d34dDd", "ETH", 0.005, nil)
	if err != nil {
		t.Error(err)
	}
	fmt.Printf("Tx Hash ===> %s\n", hash)
}

func TestTransferUSDT(t *testing.T) {
	pk, err := account.GetPrivateKeyFromHex("512d8eacf074dea19aa6f8995418076b4519166dec2590d4a7a66e1cf7f320bb")
	if err != nil {
		t.Error(err)
	}
	hash, err := client.Transfer(pk, "0x337610d27c682E347C9cD60BD4b3b107C9d34dDd", "USDT", 0.5, nil)
	if err != nil {
		t.Error(err)
	}
	fmt.Printf("Tx Hash ===> %s\n", hash)
}

func TestGetETHTransaction(t *testing.T) {
	tx, err := client.GetTransaction("0xb2c4b35a1c9796ee6e5aebad44efbcad60d2102edb2fa554d3393b7afd125166")
	if err != nil {
		t.Error(err)
	}
	fmt.Printf("Tx ===> %+v\n", tx)
}

func TestGetUSDTTransaction(t *testing.T) {
	tx, err := client.GetTransaction("0xbede7b0b4ed6ebc4ff603f87ecd7c074e7bbdc17a3c580775181881578d5fe0a")
	if err != nil {
		t.Error(err)
	}
	fmt.Printf("Tx ===> %+v\n", tx)
}
