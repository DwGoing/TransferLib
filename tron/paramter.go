package tron

import "github.com/DwGoing/transfer_lib/chain"

type Node struct {
	chain.Node
	ApiKeys []string
}

type GetBalanceParameter struct {
}

type TransferParameter struct {
}
