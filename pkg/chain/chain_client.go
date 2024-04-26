package chain

import (
	"crypto/ecdsa"
	"transfer_lib/pkg/common"
)

type IChainClient interface {
	Chain() common.Chain
	GetCurrentHeight() (uint64, error)
	GetBalance(address string, currency string, args any) (float64, error)
	Transfer(privateKey *ecdsa.PrivateKey, to string, currency string, value float64, args any) (string, error)
	GetTransaction(txHash string) (*common.Transaction, error)
}

type ChainClient struct {
	chain common.Chain
	nodes map[string]int
}

/*
@title	创建链客户端
@param	chain	common.Chain	链类型
@param 	nodes 	map[string]int	节点
@return	_		*ChainClient	链客户端
*/
func NewChainClient(chain common.Chain, nodes map[string]int) *ChainClient {
	return &ChainClient{
		chain: chain,
		nodes: nodes,
	}
}

/*
@title 	链类型
@param 	Self	*ChainClient
@return _ 		common.Chain	链类型
*/
func (Self *ChainClient) Chain() common.Chain {
	return Self.chain
}

/*
@title 	节点
@param 	Self	*ChainClient
@return _ 		map[string]int	节点
*/
func (Self *ChainClient) Nodes() map[string]int {
	return Self.nodes
}
