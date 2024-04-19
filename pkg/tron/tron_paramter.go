package tron

type TronGetBalanceParameter struct {
	currency string
}

// @title	创建查询余额参数
// @param	currency	string					地址
// @return	_			TronGetBalanceParameter	参数
func NewTronGetBalanceParameter(currency string) TronGetBalanceParameter {
	return TronGetBalanceParameter{
		currency: currency,
	}
}

type TronTransferParameter struct {
	currency string
}

// @title	创建转账参数
// @param	currency	string					地址
// @return	_			TronTransferParameter	参数
func NewTronTransferParameter(currency string) TronTransferParameter {
	return TronTransferParameter{
		currency: currency,
	}
}
