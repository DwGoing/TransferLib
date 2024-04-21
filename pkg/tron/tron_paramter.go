package tron

type TronGetBalanceParameter struct {
}

// @title	创建查询余额参数
// @return	_			TronGetBalanceParameter	参数
func NewTronGetBalanceParameter() TronGetBalanceParameter {
	return TronGetBalanceParameter{}
}

type TronTransferParameter struct {
}

// @title	创建转账参数
// @return	_			TronTransferParameter	参数
func NewTronTransferParameter() TronTransferParameter {
	return TronTransferParameter{}
}
