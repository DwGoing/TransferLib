module github.com/DwGoing/transfer_lib/tron

go 1.22.2

require (
	account v0.0.0
	chain v0.0.0
	common v0.0.0
)

replace (
	account v0.0.0 => ../account
	chain v0.0.0 => ../chain
	common v0.0.0 => ../common
)

require (
	github.com/btcsuite/btcd v0.24.0
	github.com/decred/dcrd/dcrec/secp256k1/v4 v4.3.0
	github.com/ethereum/go-ethereum v1.14.0
	github.com/fbsobreira/gotron-sdk v0.0.0-20230907131216-1e824406fe8c
	google.golang.org/grpc v1.37.0
	google.golang.org/protobuf v1.33.0
)

require (
	github.com/bits-and-blooms/bitset v1.13.0 // indirect
	github.com/btcsuite/btcd/btcec/v2 v2.3.3 // indirect
	github.com/btcsuite/btcd/btcutil v1.1.5 // indirect
	github.com/btcsuite/btcd/chaincfg/chainhash v1.1.0 // indirect
	github.com/consensys/bavard v0.1.13 // indirect
	github.com/consensys/gnark-crypto v0.12.1 // indirect
	github.com/crate-crypto/go-kzg-4844 v1.0.0 // indirect
	github.com/deckarep/golang-set v1.8.0 // indirect
	github.com/ethereum/c-kzg-4844 v1.0.1 // indirect
	github.com/golang/protobuf v1.5.4 // indirect
	github.com/google/uuid v1.3.0 // indirect
	github.com/holiman/uint256 v1.2.4 // indirect
	github.com/mitchellh/go-homedir v1.1.0 // indirect
	github.com/mmcloughlin/addchain v0.4.0 // indirect
	github.com/pborman/uuid v1.2.1 // indirect
	github.com/pkg/errors v0.9.1 // indirect
	github.com/rjeczalik/notify v0.9.3 // indirect
	github.com/shengdoushi/base58 v1.0.0 // indirect
	github.com/supranational/blst v0.3.11 // indirect
	github.com/tyler-smith/go-bip39 v1.1.0 // indirect
	go.uber.org/atomic v1.6.0 // indirect
	go.uber.org/multierr v1.5.0 // indirect
	go.uber.org/zap v1.15.0 // indirect
	golang.org/x/crypto v0.22.0 // indirect
	golang.org/x/net v0.24.0 // indirect
	golang.org/x/sync v0.7.0 // indirect
	golang.org/x/sys v0.19.0 // indirect
	golang.org/x/text v0.14.0 // indirect
	google.golang.org/genproto v0.0.0-20200825200019-8632dd797987 // indirect
	rsc.io/tmplfunc v0.0.3 // indirect
)
