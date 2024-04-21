package common

type IChainClient interface {
	GetCurrentHeight() (int64, error)
	GetBalance(address string, currency string, args any) (float64, error)
	Transfer(privateKey string, to string, currency string, value float64, args any) (string, error)
	GetTransaction(txHash string) (*Transaction, error)
}

type ChainClient struct {
	chain Chain
	nodes map[string]int
}

func NewChainClient(chain Chain, nodes map[string]int) *ChainClient {
	return &ChainClient{
		chain: chain,
		nodes: nodes,
	}
}

func (Self *ChainClient) GetNodes() map[string]int {
	return Self.nodes
}
