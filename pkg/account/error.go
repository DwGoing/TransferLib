package account

import "errors"

var (
	ErrInvalidMnemonic        = errors.New("the mnemonic is invalid")
	ErrUnsupportedAddressType = errors.New("the address type is unsupported")
)
