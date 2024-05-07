package bsc

import (
	"fmt"
	"testing"

	"github.com/DwGoing/transfer_lib/account"
	"github.com/DwGoing/transfer_lib/chain"
)

var client = NewChainClient(
	[]Node{{Node: chain.Node{Host: "https://data-seed-prebsc-2-s1.binance.org:8545", Weight: 100}}},
	map[string]Currency{
		"BNB": {
			Contract: "",
			Decimals: 18,
		},
		"USDT": {
			Contract: "337610d27c682E347C9cD60BD4b3b107C9d34dDd",
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
	balance, err := client.GetBalance("0xbb03d2098faa5867fa3381c9b1cb95f45477916e", "USDT", nil)
	if err != nil {
		t.Error(err)
	}
	fmt.Printf("Balance ===> %f\n", balance)
}

func TestTransfer(t *testing.T) {
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

func TestGetTransaction(t *testing.T) {
	tx, err := client.GetTransaction("0xc942b78e00c25cef39e732c313648e6dec4a50dda11155296ab2f4cc5dc69ef8")
	if err != nil {
		t.Error(err)
	}
	fmt.Printf("Tx ===> %+v\n", tx)
}
