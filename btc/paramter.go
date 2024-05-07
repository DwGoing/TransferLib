package btc

import "github.com/DwGoing/transfer_lib/chain"

type Node struct {
	chain.Node
	User     string
	Password string
}

type GetBalanceParameter struct {
}

type TransferParameter struct {
}
