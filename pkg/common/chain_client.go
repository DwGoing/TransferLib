package common

type IChainClient interface {
	GetBalance(address string, addressType AddressType, currency Currency) (float64, error)
}

type ChainClient struct {
	chain Chain
	nodes map[string]int
}

func newChainClient(chain Chain, nodes map[string]int) *ChainClient {
	return &ChainClient{
		chain: chain,
		nodes: nodes,
	}
}
