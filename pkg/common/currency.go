package common

import "errors"

type Currency int8

const (
	Currency_BTC  Currency = 1
	Currency_ETH  Currency = 2
	Currency_TRX  Currency = 3
	Currency_BNB  Currency = 4
	Currency_USDT Currency = 5
)

func (currency Currency) ToString() (string, error) {
	switch currency {
	case Currency_BTC:
		return "BTC", nil
	case Currency_ETH:
		return "ETH", nil
	case Currency_TRX:
		return "TRX", nil
	case Currency_BNB:
		return "BNB", nil
	case Currency_USDT:
		return "USDT", nil
	default:
		return "", errors.New("unsupported currency")
	}
}

func ParseCurrency(currency string) (Currency, error) {
	switch currency {
	case "BTC":
		return Currency_BTC, nil
	case "ETH":
		return Currency_ETH, nil
	case "TRX":
		return Currency_TRX, nil
	case "BNB":
		return Currency_BNB, nil
	case "USDT":
		return Currency_USDT, nil
	default:
		return 0, errors.New("unsupported currency")
	}
}
