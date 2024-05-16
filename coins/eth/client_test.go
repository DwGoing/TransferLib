package eth

import (
	"fmt"
	"testing"
)

var client = NewClient(
	[]Node{{Host: "https://1rpc.io/holesky", Weight: 100}},
)

func TestGetCurrentHeight(t *testing.T) {
	height, err := client.GetCurrentHeight()
	if err != nil {
		t.Error(err)
	}
	fmt.Printf("Height ===> %d\n", height)
}

func TestGetBalance(t *testing.T) {
	balance, err := client.GetBalance("0xbb03D2098FAa5867FA3381c9b1CB95F45477916E", "")
	if err != nil {
		t.Error(err)
	}
	fmt.Printf("ETH Balance ===> %f\n", balance)
	balance, err = client.GetBalance("0xbb03D2098FAa5867FA3381c9b1CB95F45477916E", "4555Ed1F6D9cb6CC1D52BB88C7525b17a06da0Dd")
	if err != nil {
		t.Error(err)
	}
	fmt.Printf("USDT Balance ===> %f\n", balance)
}

func TestTransferETH(t *testing.T) {
	hash, err := client.Transfer(
		[]byte{81, 45, 142, 172, 240, 116, 222, 161, 154, 166, 248, 153, 84, 24, 7, 107, 69, 25, 22, 109, 236, 37, 144, 212, 167, 166, 110, 28, 247, 243, 32, 187},
		"0x337610d27c682E347C9cD60BD4b3b107C9d34dDd", "", 0.005)
	if err != nil {
		t.Error(err)
	}
	fmt.Printf("Tx Hash ===> %s\n", hash)
}

func TestTransferUSDT(t *testing.T) {
	hash, err := client.Transfer(
		[]byte{81, 45, 142, 172, 240, 116, 222, 161, 154, 166, 248, 153, 84, 24, 7, 107, 69, 25, 22, 109, 236, 37, 144, 212, 167, 166, 110, 28, 247, 243, 32, 187},
		"0x337610d27c682E347C9cD60BD4b3b107C9d34dDd", "4555Ed1F6D9cb6CC1D52BB88C7525b17a06da0Dd", 0.5)
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
