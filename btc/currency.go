package btc

type Currency struct {
	Contract string `json:"," default:""`
	Decimals int    `json:","`
}
