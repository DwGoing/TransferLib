package eth

import (
	"abao/pkg/common"
	"context"
	"errors"
	"math"
	"math/big"
	"math/rand"
	"time"

	goEthereumCommon "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

type EthClient struct {
	common.ChainClient
	currencies map[string]EthCurrency
}

// @title	创建Eth客户端
// @param	nodes		map[string]int			节点列表
// @param	currencies	map[string]TronCurrency	币种列表
// @return	_			*EthClient				Eth客户端
func NewEthClient(nodes map[string]int, currencies map[string]EthCurrency) *EthClient {
	return &EthClient{
		ChainClient: *common.NewChainClient(common.Chain_TRON, nodes),
		currencies:  currencies,
	}
}

// @title	获取当前高度
// @param	Self		*EthClient
// @return	_			int64			当前高度
// @return	_			error			异常信息
func (Self *EthClient) GetCurrentHeight() (int64, error) {
	client, err := Self.GetEthClient()
	if err != nil {
		return 0, err
	}
	height, err := client.BlockNumber(context.Background())
	if err != nil {
		return 0, err
	}
	return int64(height), nil
}

// @title	查询余额
// @param	Self		*EthClient
// @param	address		string		地址
// @param	currency	string		币种
// @param	args		any			参数
// @return	_			float64		余额
// @return	_			error		异常信息
func (Self *EthClient) GetBalance(address string, currency string, args any) (float64, error) {
	client, err := Self.GetEthClient()
	if err != nil {
		return 0, err
	}
	currencyInfo, ok := Self.currencies[currency]
	if !ok {
		return 0, errors.New("unsupported currency")
	}
	var balanceBigInt *big.Int
	if currencyInfo.Contract == "" {
		balanceBigInt, err = client.BalanceAt(context.Background(), goEthereumCommon.HexToAddress(address), nil)
		if err != nil {
			return 0, err
		}
	} else {
		erc20, err := NewErc20(goEthereumCommon.HexToAddress(currencyInfo.Contract), client)
		if err != nil {
			return 0, err
		}
		balanceBigInt, err = erc20.BalanceOf(nil, goEthereumCommon.HexToAddress(address))
		if err != nil {
			return 0, err
		}
	}
	balance, _ := new(big.Float).Quo(new(big.Float).SetInt(balanceBigInt), big.NewFloat(math.Pow10(currencyInfo.Decimals))).Float64()
	return balance, nil
}

// @title	转账
// @param	Self		*EthClient
// @param	privateKey	*ecdsa.PrivateKey	私钥
// @param	to			string				交易
// @param	currency	string				币种
// @param	value		float64				金额
// @param	args		any					参数
// @return	_			string				交易哈希
// @return	_			error				异常信息
func (Self *EthClient) Transfer(privateKey string, to string, currency string, value float64, args any) (string, error) {
	return "", nil
}

// @title	查询交易
// @param	Self		*EthClient
// @param	txHash		string			交易Hash
// @return	_			*Transaction	交易信息
// @return	_			error			异常信息
func (Self *EthClient) GetTransaction(txHash string) (*common.Transaction, error) {
	return nil, nil
}

// @title	获取Eth客户端
// @param 	Self 	*EthClient
// @return 	_ 		*ethclient.Client 	Eth客户端
// @return 	_ 		error 				异常信息
func (Self *EthClient) GetEthClient() (*ethclient.Client, error) {
	sum := 0
	for _, v := range Self.GetNodes() {
		sum += v
	}
	i := rand.New(rand.NewSource(time.Now().UnixNano())).Intn(sum)
	var node string
	for k, v := range Self.GetNodes() {
		if v >= i {
			node = k
			break
		}
		i = i - v
	}
	client, err := ethclient.Dial(node)
	if err != nil {
		return nil, err
	}
	return client, nil
}
