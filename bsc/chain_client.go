package bsc

import (
	"context"
	"math"
	"math/big"
	"math/rand"
	"time"

	"github.com/DwGoing/transfer_lib/chain"
	"github.com/DwGoing/transfer_lib/common"

	"github.com/decred/dcrd/dcrec/secp256k1/v4"

	"github.com/ahmetb/go-linq"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	goEthereumCommon "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
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
		chainClient: *chain.NewChainClient(common.Chain_BSC, standardNodes),
		currencies:  currencies,
	}
}

/*
@title	获取Rpc客户端
@param 	Self 	*ChainClient
@return _ 		*ethclient.Client 	Rpc客户端
@return _ 		error 				异常信息
*/
func (Self *ChainClient) getRpcClient() (*ethclient.Client, error) {
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
	client, err := ethclient.Dial(node.Host)
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
	defer client.Close()
	height, err := client.BlockNumber(context.Background())
	if err != nil {
		return 0, err
	}
	return height, nil
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
func (Self *ChainClient) GetBalance(address string, currency string, args any) (float64, error) {
	currencyInfo, ok := Self.currencies[currency]
	if !ok {
		return 0, common.ErrUnsupportedCurrency
	}
	client, err := Self.getRpcClient()
	if err != nil {
		return 0, err
	}
	defer client.Close()
	var balanceBigInt *big.Int
	if currencyInfo.Contract == "" {
		balanceBigInt, err = client.BalanceAt(context.Background(), goEthereumCommon.HexToAddress(address), nil)
		if err != nil {
			return 0, err
		}
	} else {
		erc20, err := NewBep20(goEthereumCommon.HexToAddress(currencyInfo.Contract), client)
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

/*
@title	转账
@param	Self		*ChainClient
@param	privateKey	*secp256k1.PrivateKey	私钥
@param	to			string					接收方
@param	currency	string					币种
@param	value		float64					金额
@param	args		any						参数
@return	_			string					交易哈希
@return	_			error					异常信息
*/
func (Self *ChainClient) Transfer(privateKey *secp256k1.PrivateKey, to string, currency string, value float64, args any) (string, error) {
	currencyInfo, ok := Self.currencies[currency]
	if !ok {
		return "", common.ErrUnsupportedCurrency
	}
	client, err := Self.getRpcClient()
	if err != nil {
		return "", err
	}
	defer client.Close()
	var signedTx *types.Transaction
	chainId, err := client.ChainID(context.Background())
	if err != nil {
		return "", err
	}
	nonce, err := client.PendingNonceAt(context.Background(), crypto.PubkeyToAddress(privateKey.ToECDSA().PublicKey))
	if err != nil {
		return "", err
	}
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		return "", err
	}
	valueBigInt, _ := new(big.Float).Mul(big.NewFloat(value), big.NewFloat(math.Pow10(currencyInfo.Decimals))).Int(new(big.Int))
	if currencyInfo.Contract == "" {
		tx := types.NewTransaction(nonce, goEthereumCommon.HexToAddress(to), valueBigInt, 21000, gasPrice, nil)
		signedTx, err = types.SignTx(tx, types.LatestSignerForChainID(chainId), privateKey.ToECDSA())
		if err != nil {
			return "", err
		}
	} else {
		ierc20, err := NewBep20(goEthereumCommon.HexToAddress(currencyInfo.Contract), client)
		if err != nil {
			return "", err
		}
		transactOpts, err := bind.NewKeyedTransactorWithChainID(privateKey.ToECDSA(), chainId)
		if err != nil {
			return "", err
		}
		transactOpts.NoSend = true
		transactOpts.Nonce = big.NewInt(int64(nonce))
		transactOpts.GasLimit = uint64(300000)
		transactOpts.GasPrice = gasPrice
		signedTx, err = ierc20.Transfer(transactOpts, goEthereumCommon.HexToAddress(to), valueBigInt)
		if err != nil {
			return "", err
		}
	}
	err = client.SendTransaction(context.Background(), signedTx)
	if err != nil {
		return "", err
	}
	return signedTx.Hash().Hex(), nil
}

/*
@title	查询交易
@param	Self	*ChainClient
@param	txHash	string			交易Hash
@return	_		*Transaction	交易信息
@return	_		error			异常信息
*/
func (Self *ChainClient) GetTransaction(txHash string) (*common.Transaction, error) {
	transaction := common.Transaction{
		Chain: common.Chain_ETH,
		Hash:  txHash,
	}
	client, err := Self.getRpcClient()
	if err != nil {
		return nil, err
	}
	defer client.Close()
	receipt, err := client.TransactionReceipt(context.Background(), goEthereumCommon.HexToHash(txHash))
	if err != nil {
		return nil, err
	}
	transaction.Result = receipt.Status == 1
	transaction.Height = receipt.BlockNumber.Uint64()
	tx, isPending, err := client.TransactionByHash(context.Background(), goEthereumCommon.HexToHash(txHash))
	if err != nil {
		return nil, err
	}
	block, err := client.BlockByNumber(context.Background(), receipt.BlockNumber)
	if err != nil {
		return nil, err
	}
	transaction.TimeStamp = block.Time()
	var currency string
	var currencyInfo Currency
	var from string
	var to string
	var valueBigInt *big.Int
	if receipt.Logs == nil || len(receipt.Logs) == 0 {
		fromAddress, err := types.Sender(types.NewEIP155Signer(tx.ChainId()), tx)
		if err != nil {
			return nil, err
		}
		from = fromAddress.Hex()
		to = tx.To().Hex()
		currency = "BNB"
		valueBigInt = tx.Value()
	} else {
		matchCurrency := linq.From(Self.currencies).FirstWithT(func(item linq.KeyValue) bool {
			currency := item.Value.(Currency)
			toAddress := *tx.To()
			return goEthereumCommon.HexToAddress(currency.Contract) == toAddress
		})
		if matchCurrency == nil {
			return nil, common.ErrUnsupportedCurrency
		}
		currency = matchCurrency.(linq.KeyValue).Key.(string)
		currencyInfo = matchCurrency.(linq.KeyValue).Value.(Currency)
		erc20, err := NewBep20(goEthereumCommon.HexToAddress(currencyInfo.Contract), client)
		if err != nil {
			return nil, err
		}
		for _, log := range receipt.Logs {
			filterer, err := erc20.ParseTransfer(*log)
			if err != nil {
				continue
			}
			from = filterer.From.Hex()
			to = filterer.To.Hex()
			valueBigInt = filterer.Value
			break
		}
	}
	transaction.Currency = currency
	transaction.From = from
	transaction.To = to
	transaction.Value, _ = new(big.Float).Quo(new(big.Float).SetInt(valueBigInt), big.NewFloat(math.Pow10(Self.currencies[currency].Decimals))).Float64()
	if !isPending && transaction.Result {
		height, err := Self.GetCurrentHeight()
		if err != nil {
			return nil, err
		}
		transaction.Confirms = height - transaction.Height
	}
	return &transaction, nil
}
