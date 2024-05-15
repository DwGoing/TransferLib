package btc

import (
	"math/rand"
	"time"

	"github.com/DwGoing/transfer_lib/common"
	"github.com/ahmetb/go-linq"
	"github.com/btcsuite/btcd/btcutil"
	"github.com/btcsuite/btcd/chaincfg"
	"github.com/btcsuite/btcd/rpcclient"
)

type Client struct {
	nodes      []Node
	currencies map[string]Currency
}

// Function NewClient 创建客户端
func NewClient(nodes []Node, currencies map[string]Currency) *Client {
	standardNodes := []any{}
	linq.From(nodes).ToSlice(&standardNodes)
	return &Client{
		nodes:      nodes,
		currencies: currencies,
	}
}

// Function newRpcClient 创建RPC客户端
func (Self *Client) newRpcClient() (*rpcclient.Client, error) {
	sum := 0
	for _, item := range Self.nodes {
		sum += item.Weight
	}
	i := rand.New(rand.NewSource(time.Now().UnixNano())).Intn(sum)
	var node Node
	for _, item := range Self.nodes {
		if item.Weight >= i {
			node = item
			break
		}
		i = i - node.Weight
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
	}, nil)
	if err != nil {
		return nil, err
	}
	return client, nil
}

// Function GetCurrentHeight 获取当前高度
func (Self *Client) GetCurrentHeight() (uint64, error) {
	client, err := Self.newRpcClient()
	if err != nil {
		return 0, err
	}
	defer client.Disconnect()
	height, err := client.GetBlockCount()
	if err != nil {
		return 0, err
	}
	return uint64(height), nil
}

// Function GetBalance 查询余额
func (Self *Client) GetBalance(address string, currency string, args any) (float64, error) {
	currencyInfo, ok := Self.currencies[currency]
	if !ok {
		return 0, common.ErrUnsupportedCurrency
	}
	_ = currencyInfo
	client, err := Self.newRpcClient()
	if err != nil {
		return 0, err
	}
	defer client.Disconnect()
	addr, err := btcutil.DecodeAddress(address, &chaincfg.MainNetParams)
	if err != nil {
		return 0, err
	}
	account, err := client.GetAccount(addr)
	if err != nil {
		return 0, err
	}
	a, err := client.GetBalance(account)
	_ = a
	// var balanceBigInt *big.Int
	// if currencyInfo.Contract == "" {
	// 	balanceBigInt, err = client.BalanceAt(context.Background(), goEthereumCommon.HexToAddress(address), nil)
	// 	if err != nil {
	// 		return 0, err
	// 	}
	// } else {
	// 	erc20, err := NewBep20(goEthereumCommon.HexToAddress(currencyInfo.Contract), client)
	// 	if err != nil {
	// 		return 0, err
	// 	}
	// 	balanceBigInt, err = erc20.BalanceOf(nil, goEthereumCommon.HexToAddress(address))
	// 	if err != nil {
	// 		return 0, err
	// 	}
	// }
	// balance, _ := new(big.Float).Quo(new(big.Float).SetInt(balanceBigInt), big.NewFloat(math.Pow10(currencyInfo.Decimals))).Float64()
	// return balance, nil
	return 0, nil
}
