package common

type Chain int8

const (
	Chain_BTC  Chain = 1
	Chain_ETH  Chain = 2
	Chain_TRON Chain = 3
	Chain_BSC  Chain = 4
	Chain_SOL  Chain = 5
)

func (chian Chain) ToString() (string, error) {
	switch chian {
	case Chain_BTC:
		return "BTC", nil
	case Chain_ETH:
		return "ETH", nil
	case Chain_TRON:
		return "TRON", nil
	case Chain_BSC:
		return "BSC", nil
	case Chain_SOL:
		return "SOL", nil
	default:
		return "", ErrUnsupportedChain
	}
}

func ParseChain(chain string) (Chain, error) {
	switch chain {
	case "BTC":
		return Chain_BTC, nil
	case "ETH":
		return Chain_ETH, nil
	case "TRON":
		return Chain_TRON, nil
	case "BSC":
		return Chain_BSC, nil
	case "SOL":
		return Chain_SOL, nil
	default:
		return 0, ErrUnsupportedChain
	}
}
