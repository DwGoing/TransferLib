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
