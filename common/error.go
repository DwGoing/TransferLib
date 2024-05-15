package common

import "errors"

var (
	ErrInvalidParameter       = errors.New("parameter is invalid")
	ErrInvalidIndex           = errors.New("index is invalid")
	ErrInvalidMnemonic        = errors.New("mnemonic is invalid")
	ErrUnsupportedAddressType = errors.New("address type is unsupported")
	ErrUnsupportedCurrency    = errors.New("currency is unsupported")
	ErrInvalidTransaction     = errors.New("transaction is invalid")
	ErrSendTransactionFailed  = errors.New("send transaction failed")
	ErrTransactionNotFound    = errors.New("transaction is not found")
	ErrUnsupportedChainType   = errors.New("chain type is unsupported")
	ErrUninitializedClient    = errors.New("client is uninitialized")
)
