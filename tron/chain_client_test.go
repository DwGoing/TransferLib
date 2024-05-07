package tron

import (
	"fmt"
	"testing"

	"github.com/DwGoing/transfer_lib/account"
	"github.com/DwGoing/transfer_lib/chain"
)

var testClient = NewChainClient(
	[]Node{
		{
			Node:    chain.Node{Host: "grpc.nile.trongrid.io:50051", Weight: 100},
			ApiKeys: []string{"d9b77ec9-39e0-4765-98d8-2c59188344a0"},
		},
	},
	map[string]Currency{
		"TRX": {
			Contract: "",
			Decimals: 6,
		},
		"USDT": {
			Contract: "TXLAQ63Xg1NAzckPwKHvzw7CSEmLMEqcdj",
			Decimals: 6,
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
	balance, err := testClient.GetBalance("TYq73v8nCqi85g5CJxSNYDW5QKvaffAuPh", "TRX", nil)
	if err != nil {
		t.Error(err)
	}
	fmt.Printf("TRX Balance ===> %f\n", balance)
	balance, err = testClient.GetBalance("TYq73v8nCqi85g5CJxSNYDW5QKvaffAuPh", "USDT", nil)
	if err != nil {
		t.Error(err)
	}
	fmt.Printf("USDT Balance ===> %f\n", balance)
}

func TestTransferTRX(t *testing.T) {
	pk, err := account.GetPrivateKeyFromHex("5eb672f4866e58f256c95914c5f96a4a8ca1f117465c544023e9cb0bb2387a65")
	if err != nil {
		t.Error(err)
	}
	hash, err := testClient.Transfer(pk, "TMxdPV49nV1LTrxKLXHMzft5oxmftWgfXD", "TRX", 0.5, nil)
	if err != nil {
		t.Error(err)
	}
	fmt.Printf("Tx Hash ===> %s\n", hash)
}

func TestTransferUSDT(t *testing.T) {
	pk, err := account.GetPrivateKeyFromHex("5eb672f4866e58f256c95914c5f96a4a8ca1f117465c544023e9cb0bb2387a65")
	if err != nil {
		t.Error(err)
	}
	hash, err := testClient.Transfer(pk, "TMxdPV49nV1LTrxKLXHMzft5oxmftWgfXD", "USDT", 5, nil)
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
