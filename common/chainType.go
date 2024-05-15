package common

type ChainType int8

const (
	ChainType_BTC  ChainType = 1
	ChainType_ETH  ChainType = 2
	ChainType_TRON ChainType = 3
	ChainType_BSC  ChainType = 4
	ChainType_SOL  ChainType = 5
)

func ParseChain(chain string) (ChainType, error) {
	switch chain {
	case "BTC":
		return ChainType_BTC, nil
	case "ETH":
		return ChainType_ETH, nil
	case "TRON":
		return ChainType_TRON, nil
	case "BSC":
		return ChainType_BSC, nil
	case "SOL":
		return ChainType_SOL, nil
	default:
		return 0, ErrUnsupportedChainType
	}
}

func (Self ChainType) Name() (string, error) {
	switch Self {
	case ChainType_BTC:
		return "BTC", nil
	case ChainType_ETH:
		return "ETH", nil
	case ChainType_TRON:
		return "TRON", nil
	case ChainType_BSC:
		return "BSC", nil
	case ChainType_SOL:
		return "SOL", nil
	default:
		return "", ErrUnsupportedChainType
	}
}
