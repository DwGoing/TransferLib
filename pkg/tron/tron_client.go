package tron

import (
	"abao/pkg/common"
	"abao/pkg/hd_wallet"
	"crypto/ecdsa"
	"crypto/sha256"
	"errors"
	"fmt"
	"math"
	"math/big"
	"math/rand"
	"time"

	"github.com/ethereum/go-ethereum/crypto"
	"github.com/fbsobreira/gotron-sdk/pkg/client"
	goTornSdkCommon "github.com/fbsobreira/gotron-sdk/pkg/common"
	"github.com/fbsobreira/gotron-sdk/pkg/proto/api"
	"github.com/fbsobreira/gotron-sdk/pkg/proto/core"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/proto"
)

type TronClient struct {
	common.ChainClient
	apiKeys    []string
	currencies map[string]TronCurrency
}

// @title	创建Tron客户端
// @param	nodes		map[string]int			节点列表
// @param	apiKeys		[]string				ApiKey列表
// @param	currencies	map[string]TronCurrency	币种列表
// @return	_			*TronClient				Tron客户端
func NewTronClient(nodes map[string]int, apiKeys []string, currencies map[string]TronCurrency) *TronClient {
	return &TronClient{
		ChainClient: *common.NewChainClient(common.Chain_TRON, nodes),
		apiKeys:     apiKeys,
		currencies:  currencies,
	}
}

// @title	获取当前高度
// @param	Self		*TronClient
// @return	_			uint64			当前高度
// @return	_			error			异常信息
func (Self *TronClient) GetCurrentHeight() (uint64, error) {
	client, err := Self.getTronRpcClient()
	if err != nil {
		return 0, err
	}
	tx, err := client.GetNowBlock()
	if err != nil {
		return 0, err
	}
	return uint64(tx.BlockHeader.RawData.Number), nil
}

// @title	查询余额
// @param	Self		*TronClient
// @param	address		string		地址
// @param	currency	string		币种
// @param	args		any			参数
// @return	_			float64		余额
// @return	_			error		异常信息
func (Self *TronClient) GetBalance(address string, currency string, args any) (float64, error) {
	currencyInfo, ok := Self.currencies[currency]
	if !ok {
		return 0, errors.New("unsupported currency")
	}
	_, ok = args.(TronGetBalanceParameter)
	if !ok {
		return 0, nil
	}
	client, err := Self.getTronRpcClient()
	if err != nil {
		return 0, err
	}
	var balance float64
	if currencyInfo.Contract == "" {
		account, err := client.GetAccount(address)
		if err != nil {
			return 0, err
		}
		balance, _ = new(big.Float).Quo(new(big.Float).SetInt64(account.Balance), big.NewFloat(math.Pow10(currencyInfo.Decimals))).Float64()
	} else {
		balanceBigInt, err := client.TRC20ContractBalance(address, currencyInfo.Contract)
		if err != nil {
			return 0, err
		}
		balance, _ = new(big.Float).Quo(new(big.Float).SetInt(balanceBigInt), big.NewFloat(math.Pow10(currencyInfo.Decimals))).Float64()
	}
	return balance, nil
}

// @title	转账
// @param	Self		*TronClient
// @param	privateKey	*ecdsa.PrivateKey	私钥
// @param	to			string				交易
// @param	currency	string				币种
// @param	value		float64				金额
// @param	args		any					参数
// @return	_			string				交易哈希
// @return	_			error				异常信息
func (Self *TronClient) Transfer(privateKey string, to string, currency string, value float64, args any) (string, error) {
	account, err := hd_wallet.NewAccountFromPrivateKeyHex(common.AddressType_TRON, privateKey)
	if err != nil {
		return "", err
	}
	from, err := account.GetAddress()
	if err != nil {
		return "", err
	}
	currencyInfo, ok := Self.currencies[currency]
	if !ok {
		return "", errors.New("unsupported currency")
	}
	_, ok = args.(TronTransferParameter)
	if !ok {
		return "", nil
	}
	client, err := Self.getTronRpcClient()
	if err != nil {
		return "", err
	}
	var tx *api.TransactionExtention
	if currencyInfo.Contract == "" {
		valueInt64, _ := new(big.Float).Mul(big.NewFloat(value), big.NewFloat(math.Pow10(currencyInfo.Decimals))).Int64()
		tx, err = client.Transfer(from, to, valueInt64)
		if err != nil {
			return "", err
		}
	} else {
		valueBigInt, _ := new(big.Float).Mul(big.NewFloat(value), big.NewFloat(math.Pow10(currencyInfo.Decimals))).Int(new(big.Int))
		tx, err = client.TRC20Send(from, to, currencyInfo.Contract, valueBigInt, 300000000)
		if err != nil {
			return "", err
		}
	}
	txInfo, err := Self.sendTronTransaction(client, account.GetPrivateKey().ToECDSA(), tx.Transaction)
	if err != nil {
		return "", err
	}
	return goTornSdkCommon.Bytes2Hex(txInfo.GetId()), nil
}

// @title	查询交易
// @param	Self		*TronClient
// @param	txHash		string			交易Hash
// @return	_			*Transaction	交易信息
// @return	_			error			异常信息
func (Self *TronClient) GetTransaction(txHash string) (*common.Transaction, error) {
	transaction := common.Transaction{
		Chain: common.Chain_TRON,
		Hash:  txHash,
	}
	client, err := Self.getTronRpcClient()
	if err != nil {
		return nil, err
	}
	tx, err := client.GetTransactionInfoByID(txHash)
	if err != nil {
		return nil, err
	}
	transaction.Height = uint64(tx.BlockNumber)
	transaction.TimeStamp = uint64(tx.GetBlockTimeStamp())
	coreTx, err := client.GetTransactionByID(txHash)
	if err != nil {
		return nil, err
	}
	contracts := coreTx.RawData.GetContract()
	if len(contracts) < 1 {
		return nil, errors.New("not transfer transaction")
	}
	if contracts[0].Type != core.Transaction_Contract_TransferContract &&
		contracts[0].Type != core.Transaction_Contract_TriggerSmartContract {
		return nil, errors.New("not transfer transaction")
	}
	if tx.ContractAddress == nil {
		transaction.Result = tx.Result == core.TransactionInfo_SUCESS
		transaction.Currency = "TRX"
		var contract core.TransferContract
		err = contracts[0].GetParameter().UnmarshalTo(&contract)
		if err != nil {
			return nil, err
		}
		transaction.From = goTornSdkCommon.EncodeCheck(contract.GetOwnerAddress())
		transaction.To = goTornSdkCommon.EncodeCheck(contract.GetToAddress())
		transaction.Value, _ = new(big.Float).Quo(new(big.Float).SetInt64(contract.GetAmount()), big.NewFloat(1e6)).Float64()
	} else {
		receiptResult := tx.GetReceipt().GetResult()
		transaction.Result = receiptResult == core.Transaction_Result_SUCCESS
		contractAddress := goTornSdkCommon.EncodeCheck(tx.ContractAddress)
		currency := ""
		for c, ca := range Self.currencies {
			if ca.Contract == contractAddress {
				currency = c
				break
			}
		}
		if currency == "" {
			return nil, errors.New("unsupported currency")
		}
		transaction.Currency = currency
		var contract core.TriggerSmartContract
		err = contracts[0].GetParameter().UnmarshalTo(&contract)
		if err != nil {
			return nil, err
		}
		logs := tx.GetLog()
		if len(logs) < 1 {
			return nil, errors.New("not transfer transaction")
		}
		log := logs[0]
		topics := log.GetTopics()
		if len(topics) < 3 {
			return nil, errors.New("not transfer transaction")
		}
		// 签名校验
		if goTornSdkCommon.BytesToHexString(topics[0]) != goTornSdkCommon.BytesToHexString(goTornSdkCommon.Keccak256([]byte("Transfer(address,address,uint256)"))) {
			return nil, errors.New("not transfer transaction")
		}
		transaction.From = goTornSdkCommon.EncodeCheck(contract.GetOwnerAddress())
		transaction.To = goTornSdkCommon.EncodeCheck(append([]byte{0x41}, topics[2][12:]...))
		transaction.Value, _ = new(big.Float).Quo(new(big.Float).SetInt(new(big.Int).SetBytes(log.Data)), big.NewFloat(math.Pow10(Self.currencies[currency].Decimals))).Float64()
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

// @title	获取Tron客户端
// @param	Self	*TronClient
// @return	_		*client.GrpcClient	客户端
// @return	_		error				异常信息
func (Self *TronClient) getTronRpcClient() (*client.GrpcClient, error) {
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
	client := client.NewGrpcClient(node)
	apiKey := Self.apiKeys[rand.Int()%len(Self.apiKeys)]
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

// @title	发送Tron交易
// @param	Self		*TronClient
// @param	client		*client.GrpcClient		客户端
// @param	privateKey	*ecdsa.PrivateKey		私钥
// @param	tx			*core.Transaction		交易
// @return	_			*core.TransactionInfo	交易信息
// @return	_			error					异常信息
func (Self *TronClient) sendTronTransaction(client *client.GrpcClient, privateKey *ecdsa.PrivateKey, tx *core.Transaction) (*core.TransactionInfo, error) {
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
		return nil, fmt.Errorf("bad transaction: %v", string(result.GetMessage()))
	}
	var transaction *core.TransactionInfo
	start := 0
	for {
		if start++; start > 10 {
			return nil, errors.New("transaction info not found")
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
