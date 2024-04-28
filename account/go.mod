module github.com/DwGoing/transfer_lib/account

go 1.22.2

require common v0.0.0

replace common v0.0.0 => ../common

require (
	github.com/btcsuite/btcd v0.24.0
	github.com/btcsuite/btcd/btcec/v2 v2.3.3
	github.com/btcsuite/btcd/btcutil v1.1.5
	github.com/decred/dcrd/dcrec/secp256k1/v4 v4.3.0
	github.com/ethereum/go-ethereum v1.14.0
	github.com/fbsobreira/gotron-sdk v0.0.0-20230907131216-1e824406fe8c
	github.com/tyler-smith/go-bip39 v1.1.0
)

require (
	github.com/bits-and-blooms/bitset v1.13.0 // indirect
	github.com/btcsuite/btcd/chaincfg/chainhash v1.1.0 // indirect
	github.com/consensys/bavard v0.1.13 // indirect
	github.com/consensys/gnark-crypto v0.12.1 // indirect
	github.com/crate-crypto/go-kzg-4844 v1.0.0 // indirect
	github.com/ethereum/c-kzg-4844 v1.0.1 // indirect
	github.com/holiman/uint256 v1.2.4 // indirect
	github.com/mmcloughlin/addchain v0.4.0 // indirect
	github.com/shengdoushi/base58 v1.0.0 // indirect
	github.com/supranational/blst v0.3.11 // indirect
	golang.org/x/crypto v0.22.0 // indirect
	golang.org/x/sync v0.7.0 // indirect
	golang.org/x/sys v0.19.0 // indirect
	rsc.io/tmplfunc v0.0.3 // indirect
)
