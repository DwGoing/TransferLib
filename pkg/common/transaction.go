package common

type Transaction struct {
	Chain     Chain
	Hash      string
	Height    int64
	TimeStamp int64
	Contract  *string
	From      string
	To        string
	Amount    float64
	Result    bool
}
