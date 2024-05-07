package chain

import (
	"github.com/DwGoing/transfer_lib/common"

	"github.com/decred/dcrd/dcrec/secp256k1/v4"
)

type IChainClient interface {
	Chain() common.Chain
	GetCurrentHeight() (uint64, error)
	GetBalance(address string, currency string, args any) (float64, error)
	Transfer(privateKey *secp256k1.PrivateKey, to string, currency string, value float64, args any) (string, error)
	GetTransaction(txHash string) (*common.Transaction, error)
}

type ChainClient struct {
	chain common.Chain
	nodes []any
}

/*
@title	创建链客户端
@param	chain	common.Chain	链类型
@param 	nodes 	[]any			节点
@return	_		*ChainClient	链客户端
*/
func NewChainClient(chain common.Chain, nodes []any) *ChainClient {
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
@return _ 		[]any			节点
*/
func (Self *ChainClient) Nodes() []any {
	return Self.nodes
}
