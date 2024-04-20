package eth

type EthCurrency struct {
	Contract string `json:"," default:""`
	Decimals int    `json:","`
}
