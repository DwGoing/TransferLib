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
	}
	fmt.Printf("USDT Balance ===> %f\n", balance)
}

// func TestTransferBNB(t *testing.T) {
// 	pk, err := account.GetPrivateKeyFromHex("512d8eacf074dea19aa6f8995418076b4519166dec2590d4a7a66e1cf7f320bb")
// 	if err != nil {
// 		t.Error(err)
// 	}
// 	hash, err := client.Transfer(pk, "0x337610d27c682E347C9cD60BD4b3b107C9d34dDd", "BNB", 0.005, nil)
// 	if err != nil {
// 		t.Error(err)
// 	}
// 	fmt.Printf("Tx Hash ===> %s\n", hash)
// }

// func TestTransferUSDT(t *testing.T) {
// 	pk, err := account.GetPrivateKeyFromHex("512d8eacf074dea19aa6f8995418076b4519166dec2590d4a7a66e1cf7f320bb")
// 	if err != nil {
// 		t.Error(err)
// 	}
// 	hash, err := client.Transfer(pk, "0x337610d27c682E347C9cD60BD4b3b107C9d34dDd", "USDT", 0.5, nil)
// 	if err != nil {
// 		t.Error(err)
// 	}
// 	fmt.Printf("Tx Hash ===> %s\n", hash)
// }

// func TestGetBNBTransaction(t *testing.T) {
// 	tx, err := client.GetTransaction("0x3ab5c6df0c9c06c59c5b558c89e08ce372a903874fa6808c8eab31027df443fd")
// 	if err != nil {
// 		t.Error(err)
// 	}
// 	fmt.Printf("Tx ===> %+v\n", tx)
// }

// func TestGetUSDTTransaction(t *testing.T) {
// 	tx, err := client.GetTransaction("0xc942b78e00c25cef39e732c313648e6dec4a50dda11155296ab2f4cc5dc69ef8")
// 	if err != nil {
// 		t.Error(err)
// 	}
// 	fmt.Printf("Tx ===> %+v\n", tx)
// }
