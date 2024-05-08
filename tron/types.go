package tron

import "github.com/DwGoing/transfer_lib/chain"

type Node struct {
	chain.Node
	ApiKeys []string
}

type Currency struct {
	Contract string `json:"," default:""`
	Decimals int    `json:","`
}
