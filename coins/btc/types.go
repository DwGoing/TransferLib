package btc

type Node struct {
	Host     string
	Weight   int
	User     string
	Password string
}

type Currency struct {
	Contract string `json:"," default:""`
	Decimals int    `json:","`
}
