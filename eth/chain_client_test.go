package eth

import (
	"fmt"
	"testing"
)

func TestGetCurrentHeight(t *testing.T) {
	client := NewChainClient(
		map[string]int{"https://ethereum-holesky-rpc.publicnode.com": 100},
		map[string]Currency{
			"ETH": {
				Contract: "",
				Decimals: 18,
			},
			"USDT": {
				Contract: "4555Ed1F6D9cb6CC1D52BB88C7525b17a06da0Dd",
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
