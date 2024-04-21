package bsc

type BscCurrency struct {
	Contract string `json:"," default:""`
	Decimals int    `json:","`
}
