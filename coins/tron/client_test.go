package tron

import (
	"fmt"
	"testing"
)

var testClient = NewClient(
	[]Node{
		{
			Host:    "grpc.nile.trongrid.io:50051",
			Weight:  100,
			ApiKeys: []string{"d9b77ec9-39e0-4765-98d8-2c59188344a0"},
		},
	},
)

func TestGetCurrentHeight(t *testing.T) {
	height, err := testClient.GetCurrentHeight()
	if err != nil {
		t.Error(err)
	}
	fmt.Printf("Height ===> %d\n", height)
}

func TestGetBalance(t *testing.T) {
	balance, err := testClient.GetBalance("TYq73v8nCqi85g5CJxSNYDW5QKvaffAuPh", "")
	if err != nil {
		t.Error(err)
	}
	fmt.Printf("TRX Balance ===> %f\n", balance)
	balance, err = testClient.GetBalance("TYq73v8nCqi85g5CJxSNYDW5QKvaffAuPh", "TXLAQ63Xg1NAzckPwKHvzw7CSEmLMEqcdj")
	if err != nil {
		t.Error(err)
	}
	fmt.Printf("USDT Balance ===> %f\n", balance)
}

func TestTransferTRX(t *testing.T) {
	hash, err := testClient.Transfer(
		[]byte{94, 182, 114, 244, 134, 110, 88, 242, 86, 201, 89, 20, 197, 249, 106, 74, 140, 161, 241, 23, 70, 92, 84, 64, 35, 233, 203, 11, 178, 56, 122, 101},
		"TMxdPV49nV1LTrxKLXHMzft5oxmftWgfXD", "", 0.5,
	)
	if err != nil {
		t.Error(err)
	}
	fmt.Printf("Tx Hash ===> %s\n", hash)
}

func TestTransferUSDT(t *testing.T) {
	hash, err := testClient.Transfer(
		[]byte{94, 182, 114, 244, 134, 110, 88, 242, 86, 201, 89, 20, 197, 249, 106, 74, 140, 161, 241, 23, 70, 92, 84, 64, 35, 233, 203, 11, 178, 56, 122, 101},
		"TMxdPV49nV1LTrxKLXHMzft5oxmftWgfXD", "USDT", 5,
	)
	if err != nil {
		t.Error(err)
	}
	fmt.Printf("Tx Hash ===> %s\n", hash)
}

func TestGetTRXTransaction(t *testing.T) {
	tx, err := testClient.GetTransaction("9b94340b77fa9645b80a9c5ab769a7815a2995fde7a37c9721e1754bc5bf5c23")
	if err != nil {
		t.Error(err)
	}
	fmt.Printf("Tx ===> %+v\n", tx)
}

func TestGetUSDTTransaction(t *testing.T) {
	tx, err := testClient.GetTransaction("b0655f29a7a55d29418690dbea7d018df90ee6bc919a40f21b7e95625fdae71e")
	if err != nil {
		t.Error(err)
	}
	fmt.Printf("Tx ===> %+v\n", tx)
}
