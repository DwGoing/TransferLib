package common

type Transaction struct {
	Chain     Chain
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
