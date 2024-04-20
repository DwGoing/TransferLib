package common

type Transaction struct {
	Chain     Chain
	Hash      string
	Height    int64
	TimeStamp int64
	Currency  string
	From      string
	To        string
	Value     float64
	Confirms  int64
	Result    bool
}
