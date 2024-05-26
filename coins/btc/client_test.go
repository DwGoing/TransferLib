package btc

import (
	"fmt"
	"testing"
)

var client = NewClient(
	[]Node{{Host: "nd-987-100-938.p2pify.com/51648a7756d7757e529d48e47431ef7c", User: "hardcore-wozniak", Password: "bony-buggy-anime-kettle-cola-lent", Weight: 100}},
)

func TestGetCurrentHeight(t *testing.T) {
	height, err := client.GetCurrentHeight()
	if err != nil {
		t.Error(err)
	}
	fmt.Printf("Height ===> %d\n", height)
}

func TestGetBalance(t *testing.T) {
	balance, err := client.GetBalance("1L9RCzZKvMV3W3CjoCDbBSYBKgoXJmYJxE", "")
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
