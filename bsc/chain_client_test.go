package bsc

import (
	"fmt"
	"testing"

	"github.com/DwGoing/transfer_lib/chain"
)

func TestGetCurrentHeight(t *testing.T) {
	client := NewChainClient(
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
	height, err := client.GetCurrentHeight()
	if err != nil {
		t.Error(err)
	}
	fmt.Printf("Height ===> %d\n", height)
}
