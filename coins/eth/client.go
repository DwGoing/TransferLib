package eth

import (
	"context"
	"math"
	"math/big"
	"math/rand"
	"time"

	"github.com/DwGoing/transfer_lib/common"
	"github.com/DwGoing/transfer_lib/types"
	"github.com/ahmetb/go-linq"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	goEthereumCommon "github.com/ethereum/go-ethereum/common"
	goEthereumTypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

type Client struct {
	nodes    []Node
	decimals map[string]int
}

// 创建客户端
func NewClient(nodes []Node) *Client {
	standardNodes := []any{}
	linq.From(nodes).ToSlice(&standardNodes)
	return &Client{
		nodes:    nodes,
		decimals: map[string]int{},
	}
}

// Function newRpcClient 创建RPC客户端
func (Self *Client) newRpcClient() (*ethclient.Client, error) {
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
		i = i - item.Weight
	}
	client, err := ethclient.Dial(node.Host)
	if err != nil {
		return nil, err
	}
	return client, nil
}

// Function getDecimals 获取小数位数
func (Self *Client) getDecimals(token string) (int, error) {
	if decimals, ok := Self.decimals[token]; ok {
		return decimals, nil
	} else {
		var value int
		if token == "" {
			value = 18
		} else {
			client, err := Self.newRpcClient()
			if err != nil {
				return 0, err
			}
			defer client.Close()
			erc20, err := NewErc20(goEthereumCommon.HexToAddress(token), client)
			if err != nil {
				return 0, err
			}
			decimals, err := erc20.Decimals(nil)
			if err != nil {
				return 0, err
			}
			value = int(decimals)
		}
		Self.decimals[token] = value
		return value, nil
	}
}

// Function GetCurrentHeight 获取当前高度
func (Self *Client) GetCurrentHeight() (uint64, error) {
	client, err := Self.newRpcClient()
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

// Function GetBalance 查询余额
func (Self *Client) GetBalance(address string, token string) (float64, error) {
	client, err := Self.newRpcClient()
	if err != nil {
		return 0, err
	}
	defer client.Close()
	var balanceBigInt *big.Int
	if token == "" {
		balanceBigInt, err = client.BalanceAt(context.Background(), goEthereumCommon.HexToAddress(address), nil)
		if err != nil {
			return 0, err
		}
	} else {
		erc20, err := NewErc20(goEthereumCommon.HexToAddress(token), client)
		if err != nil {
			return 0, err
		}
		balanceBigInt, err = erc20.BalanceOf(nil, goEthereumCommon.HexToAddress(address))
		if err != nil {
			return 0, err
		}
	}
	decimals, err := Self.getDecimals(token)
	if err != nil {
		return 0, err
	}
	balance, _ := new(big.Float).Quo(new(big.Float).SetInt(balanceBigInt), big.NewFloat(math.Pow10(decimals))).Float64()
	return balance, nil
}

// Function Transfer 转账
func (Self *Client) Transfer(privateKey []byte, to string, token string, value float64) (string, error) {
	ecdsaPrivateKey := crypto.ToECDSAUnsafe(privateKey)
	client, err := Self.newRpcClient()
	if err != nil {
		return "", err
	}
	defer client.Close()
	var signedTx *goEthereumTypes.Transaction
	chainId, err := client.ChainID(context.Background())
	if err != nil {
		return "", err
	}
	nonce, err := client.PendingNonceAt(context.Background(), crypto.PubkeyToAddress(ecdsaPrivateKey.PublicKey))
	if err != nil {
		return "", err
	}
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		return "", err
	}
	decimals, err := Self.getDecimals(token)
	if err != nil {
		return "", err
	}
	valueBigInt, _ := new(big.Float).Mul(big.NewFloat(value), big.NewFloat(math.Pow10(decimals))).Int(new(big.Int))
	if token == "" {
		tx := goEthereumTypes.NewTransaction(nonce, goEthereumCommon.HexToAddress(to), valueBigInt, 21000, gasPrice, nil)
		signedTx, err = goEthereumTypes.SignTx(tx, goEthereumTypes.LatestSignerForChainID(chainId), ecdsaPrivateKey)
		if err != nil {
			return "", err
		}
	} else {
		erc20, err := NewErc20(goEthereumCommon.HexToAddress(token), client)
		if err != nil {
			return "", err
		}
		transactOpts, err := bind.NewKeyedTransactorWithChainID(ecdsaPrivateKey, chainId)
		if err != nil {
			return "", err
		}
		transactOpts.NoSend = true
		transactOpts.Nonce = big.NewInt(int64(nonce))
		transactOpts.GasLimit = uint64(300000)
		transactOpts.GasPrice = gasPrice
		signedTx, err = erc20.Transfer(transactOpts, goEthereumCommon.HexToAddress(to), valueBigInt)
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

// Function GetTransaction 查询交易
func (Self *Client) GetTransaction(txHash string) (*types.Transaction, error) {
	transaction := types.Transaction{
		ChainType: common.ChainType_ETH,
		Hash:      txHash,
	}
	client, err := Self.newRpcClient()
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
	var token string
	var from string
	var to string
	var valueBigInt *big.Int
	if receipt.Logs == nil || len(receipt.Logs) == 0 {
		fromAddress, err := goEthereumTypes.Sender(goEthereumTypes.NewEIP155Signer(tx.ChainId()), tx)
		if err != nil {
			return nil, err
		}
		from = fromAddress.Hex()
		to = tx.To().Hex()
		token = ""
		valueBigInt = tx.Value()
	} else {
		token = tx.To().Hex()
		erc20, err := NewErc20(goEthereumCommon.HexToAddress(token), client)
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
	transaction.Token = token
	transaction.From = from
	transaction.To = to
	decimals, err := Self.getDecimals(token)
	if err != nil {
		return nil, err
	}
	transaction.Value, _ = new(big.Float).Quo(new(big.Float).SetInt(valueBigInt), big.NewFloat(math.Pow10(decimals))).Float64()
	if !isPending && transaction.Result {
		height, err := Self.GetCurrentHeight()
		if err != nil {
			return nil, err
		}
		transaction.Confirms = height - transaction.Height
	}
	return &transaction, nil
}
