package common

import "errors"

var (
	ErrInvalidIndex           = errors.New("the index is invalid")
	ErrInvalidMnemonic        = errors.New("the mnemonic is invalid")
	ErrUnsupportedAddressType = errors.New("the address type is unsupported")
	ErrUnsupportedCurrency    = errors.New("the currency is unsupported")
	ErrInvalidTransaction     = errors.New("the transaction is invalid")
	ErrSendTransactionFailed  = errors.New("send transaction failed")
	ErrTransactionNotFound    = errors.New("the transaction is not found")
	ErrUnsupportedChain       = errors.New("the chain is unsupported")
)
