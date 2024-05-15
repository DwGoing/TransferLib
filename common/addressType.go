package common

type AddressType int8

const (
	AddressType_BTC_LEGACY        AddressType = 1
	AddressType_BTC_NESTED_SEGWIT AddressType = 2
	AddressType_BTC_NATIVE_SEGWIT AddressType = 3
	AddressType_BTC_TAPROOT       AddressType = 4
	AddressType_ETH               AddressType = 5
	AddressType_TRON              AddressType = 6
	AddressType_BSC               AddressType = 7
	AddressType_SOL               AddressType = 8
)

func ParseAddressType(AddressType string) (AddressType, error) {
	switch AddressType {
	case "BTC Legacy":
		return AddressType_BTC_LEGACY, nil
	case "BTC Nested SegWit":
		return AddressType_BTC_NESTED_SEGWIT, nil
	case "BTC Native SegWit":
		return AddressType_BTC_NATIVE_SEGWIT, nil
	case "BTC Taproot":
		return AddressType_BTC_TAPROOT, nil
	case "ETH":
		return AddressType_ETH, nil
	case "TRON":
		return AddressType_TRON, nil
	case "BSC":
		return AddressType_BSC, nil
	case "SOL":
		return AddressType_SOL, nil
	default:
		return 0, ErrUnsupportedAddressType
	}
}

func (Self AddressType) Name() (string, error) {
	switch Self {
	case AddressType_BTC_LEGACY:
		return "BTC Legacy", nil
	case AddressType_BTC_NESTED_SEGWIT:
		return "BTC Nested SegWit", nil
	case AddressType_BTC_NATIVE_SEGWIT:
		return "BTC Native SegWit", nil
	case AddressType_BTC_TAPROOT:
		return "BTC Taproot", nil
	case AddressType_ETH:
		return "ETH", nil
	case AddressType_TRON:
		return "TRON", nil
	case AddressType_BSC:
		return "BSC", nil
	case AddressType_SOL:
		return "SOL", nil
	default:
		return "", ErrUnsupportedAddressType
	}
}

func (Self AddressType) Path() (string, error) {
	switch Self {
	case AddressType_BTC_LEGACY:
		return "m/44'/0'/0'/0/", nil
	case AddressType_BTC_NESTED_SEGWIT:
		return "m/49'/0'/0'/0/", nil
	case AddressType_BTC_NATIVE_SEGWIT:
		return "m/84'/0'/0'/0/", nil
	case AddressType_BTC_TAPROOT:
		return "m/86'/0'/0'/0/", nil
	case AddressType_ETH:
		return "m/44'/60'/0'/0/", nil
	case AddressType_TRON:
		return "m/44'/195'/0'/0/", nil
	case AddressType_BSC:
		return "m/44'/60'/0'/0/", nil
	case AddressType_SOL:
		return "m/44'/501'/0'/0/", nil
	default:
		return "", ErrUnsupportedAddressType
	}
}
