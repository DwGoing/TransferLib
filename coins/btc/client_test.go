package btc

import (
	"fmt"
	"testing"
)

var client = NewClient(
	[]Node{{Host: "go.getblock.io/c01f1669b9554b4da3c4881fd1c1d2ca", Weight: 100}},
	map[string]Currency{
		"BTC": {
			Contract: "",
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
	balance, err := client.GetBalance("bc1pew4cpmy60z9ffafu2qe7a0xqjr2cxhclrlh6wldvt4ljgdpzr96s58d29e", "BTC", nil)
	if err != nil {
		t.Error(err)
	}
	fmt.Printf("BTC Balance ===> %f\n", balance)
	// balance, err = client.GetBalance("0xbb03d2098faa5867fa3381c9b1cb95f45477916e", "USDT", nil)
	// if err != nil {
	// 	t.Error(err)
	// }
	// fmt.Printf("USDT Balance ===> %f\n", balance)
}
