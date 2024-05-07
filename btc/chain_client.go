package btc

import (
	"math/rand"
	"time"

	"github.com/DwGoing/transfer_lib/chain"
	"github.com/DwGoing/transfer_lib/common"
	"github.com/ahmetb/go-linq"
	"github.com/btcsuite/btcd/rpcclient"
)

type ChainClient struct {
	chainClient chain.ChainClient
	currencies  map[string]Currency
}

/*
@title	创建链客户端
@param	nodes		[]Node				节点列表
@param	currencies	map[string]Currency	币种列表
@return	_			*ChainClient		链客户端
*/
func NewChainClient(nodes []Node, currencies map[string]Currency) *ChainClient {
	standardNodes := []any{}
	linq.From(nodes).ToSlice(&standardNodes)
	return &ChainClient{
		chainClient: *chain.NewChainClient(common.Chain_BTC, standardNodes),
		currencies:  currencies,
	}
}

/*
@title	获取Rpc客户端
@param 	Self 	*ChainClient
@return _ 		*rpcclient.Client	Rpc客户端
@return _ 		error 				异常信息
*/
func (Self *ChainClient) getRpcClient() (*rpcclient.Client, error) {
	sum := 0
	for _, item := range Self.chainClient.Nodes() {
		sum += item.(Node).Weight
	}
	i := rand.New(rand.NewSource(time.Now().UnixNano())).Intn(sum)
	var node Node
	for _, item := range Self.chainClient.Nodes() {
		n := item.(Node)
		if n.Weight >= i {
			node = n
			break
		}
		i = i - n.Weight
	}
	user := node.User
	if user == "" {
		user = "temp"
	}
	password := node.Password
	if password == "" {
		password = "temp"
	}
	client, err := rpcclient.New(&rpcclient.ConnConfig{
		Host:         node.Host,
		User:         user,
		Pass:         password,
		HTTPPostMode: true,
		DisableTLS:   true,
	}, nil)
	if err != nil {
		return nil, err
	}
	return client, nil
}

/*
@title 	链类型
@param 	Self	*ChainClient
@return _ 		common.Chain	链类型
*/
func (Self *ChainClient) Chain() common.Chain {
	return Self.chainClient.Chain()
}

/*
@title	获取当前高度
@param	Self	*ChainClient
@return	_		uint64			当前高度
@return	_		error			异常信息
*/
func (Self *ChainClient) GetCurrentHeight() (uint64, error) {
	client, err := Self.getRpcClient()
	if err != nil {
		return 0, err
	}
	height, err := client.GetBlockCount()
	if err != nil {
		return 0, err
	}
	return uint64(height), nil
}

/*
@title	查询余额
@param	Self		*ChainClient
@param	address		string			地址
@param	currency	string			币种
@param	args		any				参数
@return	_			float64			余额
@return	_			error			异常信息
*/
// func (Self *ChainClient) GetBalance(address string, currency string, args any) (float64, error) {
// 	currencyInfo, ok := Self.currencies[currency]
// 	if !ok {
// 		return 0, common.ErrUnsupportedCurrency
// 	}
// 	_, ok = args.(GetBalanceParameter)
// 	if !ok {
// 		return 0, nil
// 	}
// 	client, err := Self.getRpcClient()
// 	if err != nil {
// 		return 0, err
// 	}
// 	// var balanceBigInt *big.Int
// 	// if currencyInfo.Contract == "" {
// 	// 	balanceBigInt, err = client.BalanceAt(context.Background(), goEthereumCommon.HexToAddress(address), nil)
// 	// 	if err != nil {
// 	// 		return 0, err
// 	// 	}
// 	// } else {
// 	// 	erc20, err := NewBep20(goEthereumCommon.HexToAddress(currencyInfo.Contract), client)
// 	// 	if err != nil {
// 	// 		return 0, err
// 	// 	}
// 	// 	balanceBigInt, err = erc20.BalanceOf(nil, goEthereumCommon.HexToAddress(address))
// 	// 	if err != nil {
// 	// 		return 0, err
// 	// 	}
// 	// }
// 	// balance, _ := new(big.Float).Quo(new(big.Float).SetInt(balanceBigInt), big.NewFloat(math.Pow10(currencyInfo.Decimals))).Float64()
// 	// return balance, nil
// }
