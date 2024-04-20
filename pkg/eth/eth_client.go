package eth

import "abao/pkg/common"

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
	return 0, nil
}

// @title	查询余额
// @param	Self	*EthClient
// @param	address	string		地址
// @param	args	any			参数
// @return	_		float64		余额
// @return	_		error		异常信息
func (Self *EthClient) GetBalance(address string, args any) (float64, error) {
	return 0, nil
}

// @title	转账
// @param	Self		*EthClient
// @param	privateKey	*ecdsa.PrivateKey	私钥
// @param	to			string				交易
// @param	value		float64				金额
// @param	args		any					参数
// @return	_			string				交易哈希
// @return	_			error				异常信息
func (Self *EthClient) Transfer(privateKey string, to string, value float64, args any) (string, error) {
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
