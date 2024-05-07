package tron

import (
	"fmt"
	"testing"

	"github.com/DwGoing/transfer_lib/chain"
)

func TestGetCurrentHeight(t *testing.T) {
	client := NewChainClient(
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
	height, err := client.GetCurrentHeight()
	if err != nil {
		t.Error(err)
	}
	fmt.Printf("Height ===> %d\n", height)
}
