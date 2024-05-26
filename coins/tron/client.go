package tron

import (
	"crypto/ecdsa"
	"crypto/sha256"
	"errors"
	"math"
	"math/big"
	"math/rand"
	"time"

	"github.com/DwGoing/transfer_lib/common"
	"github.com/DwGoing/transfer_lib/types"
	"github.com/ahmetb/go-linq"

	"github.com/ethereum/go-ethereum/crypto"
	"github.com/fbsobreira/gotron-sdk/pkg/client"
	goTornSdkCommon "github.com/fbsobreira/gotron-sdk/pkg/common"
	"github.com/fbsobreira/gotron-sdk/pkg/proto/api"
	"github.com/fbsobreira/gotron-sdk/pkg/proto/core"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/proto"
)

type Client struct {
	nodes    []Node
	decimals map[string]int
}

// Function NewClient 创建客户端
func NewClient(nodes []Node) *Client {
	standardNodes := []any{}
	linq.From(nodes).ToSlice(&standardNodes)
	return &Client{
		nodes:    nodes,
		decimals: map[string]int{},
	}
}

// Function newRpcClient 获取Rpc客户端
func (Self *Client) newRpcClient() (*client.GrpcClient, error) {
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
	client := client.NewGrpcClient(node.Host)
	apiKey := node.ApiKeys[rand.Int()%len(node.ApiKeys)]
	err := client.SetAPIKey(apiKey)
	if err != nil {
		return nil, err
	}
	err = client.Start(grpc.WithTransportCredentials(insecure.NewCredentials()))
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
			value = 6
		} else {
			client, err := Self.newRpcClient()
			if err != nil {
				return 0, err
			}
			defer client.Conn.Close()
			decimalsBigInt, err := client.TRC20GetDecimals(token)
			if err != nil {
				return 0, err
			}
			value = int(decimalsBigInt.Int64())
		}
		Self.decimals[token] = value
		return value, nil
	}
}

// Function sendTransaction 发送交易
func (Self *Client) sendTransaction(client *client.GrpcClient, privateKey *ecdsa.PrivateKey, tx *core.Transaction) (*core.TransactionInfo, error) {
	rawData, err := proto.Marshal(tx.GetRawData())
	if err != nil {
		return nil, err
	}
	h256h := sha256.New()
	h256h.Write(rawData)
	hash := h256h.Sum(nil)

	signature, err := crypto.Sign(hash, privateKey)
	if err != nil {
		return nil, err
	}
	tx.Signature = append(tx.Signature, signature)
	result, err := client.Broadcast(tx)
	if err != nil {
		return nil, err
	}
	if result.Code != 0 {
		return nil, errors.Join(common.ErrSendTransactionFailed, errors.New(string(result.GetMessage())))
	}
	var transaction *core.TransactionInfo
	start := 0
	for {
		if start++; start > 10 {
			return nil, common.ErrTransactionNotFound
		}
		transaction, err = client.GetTransactionInfoByID(goTornSdkCommon.BytesToHexString(hash))
		if err != nil {
			time.Sleep(time.Second)
			continue
		}
		if transaction.Result != 0 {
			return nil, errors.New(string(transaction.ResMessage))
		}
		break
	}
	return transaction, err
}

// Function GetCurrentHeight 获取当前高度
func (Self *Client) GetCurrentHeight() (uint64, error) {
	client, err := Self.newRpcClient()
	if err != nil {
		return 0, err
	}
	defer client.Conn.Close()
	tx, err := client.GetNowBlock()
	if err != nil {
		return 0, err
	}
	return uint64(tx.BlockHeader.RawData.Number), nil
}

// Function GetBalance 查询余额
func (Self *Client) GetBalance(address string, token string) (float64, error) {
	client, err := Self.newRpcClient()
	if err != nil {
		return 0, err
	}
	defer client.Conn.Close()
	var balanceBigInt *big.Int
	if token == "" {
		account, err := client.GetAccount(address)
		if err != nil {
			return 0, err
		}
		balanceBigInt = big.NewInt(account.Balance)
	} else {
		balanceBigInt, err = client.TRC20ContractBalance(address, token)
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
	from := GetAddressFromPrivateKey(privateKey)
	client, err := Self.newRpcClient()
	if err != nil {
		return "", err
	}
	defer client.Conn.Close()
	var tx *api.TransactionExtention
	decimals, err := Self.getDecimals(token)
	if err != nil {
		return "", err
	}
	if token == "" {
		valueInt64, _ := new(big.Float).Mul(big.NewFloat(value), big.NewFloat(math.Pow10(decimals))).Int64()
		tx, err = client.Transfer(from, to, valueInt64)
		if err != nil {
			return "", err
		}
	} else {
		valueBigInt, _ := new(big.Float).Mul(big.NewFloat(value), big.NewFloat(math.Pow10(decimals))).Int(new(big.Int))
		tx, err = client.TRC20Send(from, to, token, valueBigInt, 300000000)
		if err != nil {
			return "", err
		}
	}
	txInfo, err := Self.sendTransaction(client, ecdsaPrivateKey, tx.Transaction)
	if err != nil {
		return "", err
	}
	return goTornSdkCommon.Bytes2Hex(txInfo.GetId()), nil
}

// Function GetTransaction 查询交易
func (Self *Client) GetTransaction(txHash string) (*types.Transaction, error) {
	transaction := types.Transaction{
		ChainType: common.ChainType_TRON,
		Hash:      txHash,
	}
	client, err := Self.newRpcClient()
	if err != nil {
		return nil, err
	}
	tx, err := client.GetTransactionInfoByID(txHash)
	if err != nil {
		return nil, err
	}
	defer client.Conn.Close()
	transaction.Height = uint64(tx.BlockNumber)
	transaction.TimeStamp = uint64(tx.GetBlockTimeStamp())
	coreTx, err := client.GetTransactionByID(txHash)
	if err != nil {
		return nil, err
	}
	contracts := coreTx.RawData.GetContract()
	if len(contracts) < 1 {
		return nil, common.ErrInvalidTransaction
	}
	if contracts[0].Type != core.Transaction_Contract_TransferContract &&
		contracts[0].Type != core.Transaction_Contract_TriggerSmartContract {
		return nil, common.ErrInvalidTransaction
	}
	if tx.ContractAddress == nil {
		transaction.Result = tx.Result == core.TransactionInfo_SUCESS
		var contract core.TransferContract
		err = contracts[0].GetParameter().UnmarshalTo(&contract)
		if err != nil {
			return nil, err
		}
		transaction.From = goTornSdkCommon.EncodeCheck(contract.GetOwnerAddress())
		transaction.To = goTornSdkCommon.EncodeCheck(contract.GetToAddress())
		decimals, err := Self.getDecimals("")
		if err != nil {
			return nil, err
		}
		transaction.Value, _ = new(big.Float).Quo(new(big.Float).SetInt64(contract.GetAmount()), big.NewFloat(math.Pow10(decimals))).Float64()
	} else {
		receiptResult := tx.GetReceipt().GetResult()
		transaction.Result = receiptResult == core.Transaction_Result_SUCCESS
		contractAddress := goTornSdkCommon.EncodeCheck(tx.ContractAddress)
		transaction.Token = contractAddress
		decimals, err := Self.getDecimals(contractAddress)
		if err != nil {
			return nil, err
		}
		var contract core.TriggerSmartContract
		err = contracts[0].GetParameter().UnmarshalTo(&contract)
		if err != nil {
			return nil, err
		}
		logs := tx.GetLog()
		if len(logs) < 1 {
			return nil, common.ErrInvalidTransaction
		}
		log := logs[0]
		topics := log.GetTopics()
		if len(topics) < 3 {
			return nil, common.ErrInvalidTransaction
		}
		// 签名校验
		if goTornSdkCommon.BytesToHexString(topics[0]) != goTornSdkCommon.BytesToHexString(goTornSdkCommon.Keccak256([]byte("Transfer(address,address,uint256)"))) {
			return nil, common.ErrInvalidTransaction
		}
		transaction.From = goTornSdkCommon.EncodeCheck(contract.GetOwnerAddress())
		transaction.To = goTornSdkCommon.EncodeCheck(append([]byte{0x41}, topics[2][12:]...))
		transaction.Value, _ = new(big.Float).Quo(new(big.Float).SetInt(new(big.Int).SetBytes(log.Data)), big.NewFloat(math.Pow10(decimals))).Float64()
	}
	if transaction.Result {
		height, err := Self.GetCurrentHeight()
		if err != nil {
			return nil, err
		}
		transaction.Confirms = height - transaction.Height
	}
	return &transaction, nil
}
