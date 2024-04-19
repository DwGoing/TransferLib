package tron

type TronCurrency struct {
	Contract string `json:"," default:""`
	Decimals int    `json:","`
}
