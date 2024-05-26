package bsc

import (
	"fmt"
	"testing"
)

var client = NewClient(
	[]Node{{Host: "https://data-seed-prebsc-2-s1.binance.org:8545", Weight: 100}},
)

func TestGetCurrentHeight(t *testing.T) {
	height, err := client.GetCurrentHeight()
	if err != nil {
		t.Error(err)
		return
	}
	fmt.Printf("Height ===> %d\n", height)
}

func TestGetBalance(t *testing.T) {
	balance, err := client.GetBalance("0xbb03d2098faa5867fa3381c9b1cb95f45477916e", "")
	if err != nil {
		t.Error(err)
	}
	fmt.Printf("BNB Balance ===> %f\n", balance)
	balance, err = client.GetBalance("0xbb03d2098faa5867fa3381c9b1cb95f45477916e", "337610d27c682E347C9cD60BD4b3b107C9d34dDd")
	if err != nil {
		t.Error(err)
		return
	}
	fmt.Printf("USDT Balance ===> %f\n", balance)
}

func TestTransferBNB(t *testing.T) {
	hash, err := client.Transfer(
		[]byte{81, 45, 142, 172, 240, 116, 222, 161, 154, 166, 248, 153, 84, 24, 7, 107, 69, 25, 22, 109, 236, 37, 144, 212, 167, 166, 110, 28, 247, 243, 32, 187},
		"0x337610d27c682E347C9cD60BD4b3b107C9d34dDd", "", 0.005)
	if err != nil {
		t.Error(err)
		return
	}
	fmt.Printf("Tx Hash ===> %s\n", hash)
}

func TestTransferUSDT(t *testing.T) {
	hash, err := client.Transfer(
		[]byte{81, 45, 142, 172, 240, 116, 222, 161, 154, 166, 248, 153, 84, 24, 7, 107, 69, 25, 22, 109, 236, 37, 144, 212, 167, 166, 110, 28, 247, 243, 32, 187},
		"0x337610d27c682E347C9cD60BD4b3b107C9d34dDd", "337610d27c682E347C9cD60BD4b3b107C9d34dDd", 0.5)
	if err != nil {
		t.Error(err)
		return
	}
	fmt.Printf("Tx Hash ===> %s\n", hash)
}

func TestGetBNBTransaction(t *testing.T) {
	tx, err := client.GetTransaction("0x3ab5c6df0c9c06c59c5b558c89e08ce372a903874fa6808c8eab31027df443fd")
	if err != nil {
		t.Error(err)
		return
	}
	fmt.Printf("Tx ===> %+v\n", tx)
}

func TestGetUSDTTransaction(t *testing.T) {
	tx, err := client.GetTransaction("0xc942b78e00c25cef39e732c313648e6dec4a50dda11155296ab2f4cc5dc69ef8")
	if err != nil {
		t.Error(err)
	}
	fmt.Printf("Tx ===> %+v\n", tx)
}
