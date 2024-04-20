package eth

type EthGetBalanceParameter struct {
	currency string
}

// @title	创建查询余额参数
// @param	currency	string					地址
// @return	_			EthGetBalanceParameter	参数
func NewEthGetBalanceParameter(currency string) EthGetBalanceParameter {
	return EthGetBalanceParameter{
		currency: currency,
	}
}

type EthTransferParameter struct {
	currency string
}

// @title	创建转账参数
// @param	currency	string					地址
// @return	_			EthTransferParameter	参数
func NewEthTransferParameter(currency string) EthTransferParameter {
	return EthTransferParameter{
		currency: currency,
	}
}
