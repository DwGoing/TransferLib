package eth

type EthGetBalanceParameter struct {
}

// @title	创建查询余额参数
// @return	_			EthGetBalanceParameter	参数
func NewEthGetBalanceParameter() EthGetBalanceParameter {
	return EthGetBalanceParameter{}
}

type EthTransferParameter struct {
}

// @title	创建转账参数
// @return	_			EthTransferParameter	参数
func NewEthTransferParameter() EthTransferParameter {
	return EthTransferParameter{}
}
