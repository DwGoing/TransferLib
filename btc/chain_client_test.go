package btc

import (
	"fmt"
	"testing"

	"github.com/DwGoing/transfer_lib/chain"
)

func TestGetCurrentHeight(t *testing.T) {
	client := NewChainClient(
		[]Node{{Node: chain.Node{Host: "bitcoin-testnet.drpc.org", Weight: 100}}},
		map[string]Currency{
			"BTC": {
				Contract: "",
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
