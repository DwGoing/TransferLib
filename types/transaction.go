package types

import "github.com/DwGoing/transfer_lib/common"

type Transaction struct {
	ChainType common.ChainType
	Hash      string
	Height    uint64
	TimeStamp uint64
	Currency  string
	From      string
	To        string
	Value     float64
	Confirms  uint64
	Result    bool
}
