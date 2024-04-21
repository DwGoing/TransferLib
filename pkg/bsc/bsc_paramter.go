package bsc

type BscGetBalanceParameter struct {
}

// @title	创建查询余额参数
// @return	_			BscGetBalanceParameter	参数
func NewBscGetBalanceParameter() BscGetBalanceParameter {
	return BscGetBalanceParameter{}
}

type BscTransferParameter struct {
}

// @title	创建转账参数
// @return	_			BscGetBalanceParameter	参数
func NewBscTransferParameter() BscTransferParameter {
	return BscTransferParameter{}
}
