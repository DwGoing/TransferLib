package hd_wallet

import (
	"github.com/DwGoing/transfer_lib/coins/bsc"
	"github.com/DwGoing/transfer_lib/coins/btc"
	"github.com/DwGoing/transfer_lib/coins/eth"
	"github.com/DwGoing/transfer_lib/coins/sol"
	"github.com/DwGoing/transfer_lib/coins/tron"
	"github.com/DwGoing/transfer_lib/common"
	"github.com/DwGoing/transfer_lib/crypto"
	"github.com/DwGoing/transfer_lib/types"
	"github.com/btcsuite/btcd/chaincfg"
)

type HDWallet struct {
	seed    []byte
	clients map[common.ChainType]any
}

// Function NewHDWalletFromSeed 通过种子创建HD钱包
func NewHDWalletFromSeed(seed []byte) *HDWallet {
	return &HDWallet{
		seed:    seed,
		clients: map[common.ChainType]any{},
	}
}

// Function NewWalletFromMnemonic 通过种子创建钱包
func NewHDWalletFromMnemonic(mnemonic string, password string) (*HDWallet, error) {
	seed, err := crypto.NewSeedFromMnemonic(mnemonic, password)
	if err != nil {
		return nil, err
	}
	return NewHDWalletFromSeed(seed), nil
}

// Function Seed 种子
func (Self *HDWallet) Seed() []byte {
	return Self.seed
}

// Function GetPrivateKey 获取根私钥
func (Self *HDWallet) GetPrivateKey() ([]byte, error) {
	return crypto.NewPrivateKeyFromSeed(Self.seed)
}

// Function GetPrivateKeyByPath 获取子私钥
func (Self *HDWallet) GetPrivateKeyByPath(path string) ([]byte, error) {
	return crypto.NewPrivateKeyFromSeedByPath(Self.seed, path)
}

// Function GetAddress 获取地址
func (Sef *HDWallet) GetAddress(path string, addressType common.AddressType, args any) (string, error) {
	privateKey, err := Sef.GetPrivateKeyByPath(path)
	if err != nil {
		return "", err
	}
	switch addressType {
	case common.AddressType_BTC_LEGACY, common.AddressType_BTC_NESTED_SEGWIT, common.AddressType_BTC_NATIVE_SEGWIT, common.AddressType_BTC_TAPROOT:
		newtwork, ok := args.(chaincfg.Params)
		if !ok {
			return "", common.ErrInvalidParameter
		}
		return btc.GetAddressFromPrivateKey(privateKey, addressType, &newtwork)
	case common.AddressType_ETH:
		return eth.GetAddressFromPrivateKey(privateKey), nil
	case common.AddressType_BSC:
		return bsc.GetAddressFromPrivateKey(privateKey), nil
	case common.AddressType_TRON:
		return tron.GetAddressFromPrivateKey(privateKey), nil
	case common.AddressType_SOL:
		return sol.GetAddressFromPrivateKey(privateKey)
	default:
		return "", common.ErrUnsupportedAddressType
	}
}

// Function NewClient 创建客户端
func (Self *HDWallet) NewClient(chainType common.ChainType, args any) error {
	var client any
	switch chainType {
	case common.ChainType_ETH:
		if nodes, ok := args.([]eth.Node); ok {
			client = eth.NewClient(nodes)
		} else {
			return common.ErrInvalidParameter
		}
	case common.ChainType_BSC:
		if nodes, ok := args.([]bsc.Node); ok {
			client = bsc.NewClient(nodes)
		} else {
			return common.ErrInvalidParameter
		}
	case common.ChainType_TRON:
		if nodes, ok := args.([]tron.Node); ok {
			client = tron.NewClient(nodes)
		} else {
			return common.ErrInvalidParameter
		}
	default:
		return common.ErrUnsupportedChainType
	}
	Self.clients[chainType] = client
	return nil
}

// Function GetCurrentHeight 获取当前高度
func (Self *HDWallet) GetCurrentHeight(chainType common.ChainType) (uint64, error) {
	if client, ok := Self.clients[chainType]; ok {
		switch chainType {
		case common.ChainType_ETH:
			if ethClient, ok := client.(*eth.Client); ok {
				return ethClient.GetCurrentHeight()
			} else {
				return 0, common.ErrUninitializedClient
			}
		case common.ChainType_BSC:
			if bscClient, ok := client.(*bsc.Client); ok {
				return bscClient.GetCurrentHeight()
			} else {
				return 0, common.ErrUninitializedClient
			}
		case common.ChainType_TRON:
			if tronClient, ok := client.(*tron.Client); ok {
				return tronClient.GetCurrentHeight()
			} else {
				return 0, common.ErrUninitializedClient
			}
		default:
			return 0, common.ErrUnsupportedChainType
		}

	} else {
		return 0, common.ErrUninitializedClient
	}
}

// Function GetBalance 查询余额
func (Self *HDWallet) GetBalance(chainType common.ChainType, address string, token string) (float64, error) {
	if client, ok := Self.clients[chainType]; ok {
		switch chainType {
		case common.ChainType_ETH:
			if ethClient, ok := client.(*eth.Client); ok {
				return ethClient.GetBalance(address, token)
			} else {
				return 0, common.ErrUninitializedClient
			}
		case common.ChainType_BSC:
			if bscClient, ok := client.(*bsc.Client); ok {
				return bscClient.GetBalance(address, token)
			} else {
				return 0, common.ErrUninitializedClient
			}
		case common.ChainType_TRON:
			if tronClient, ok := client.(*tron.Client); ok {
				return tronClient.GetBalance(address, token)
			} else {
				return 0, common.ErrUninitializedClient
			}
		default:
			return 0, common.ErrUnsupportedChainType
		}

	} else {
		return 0, common.ErrUninitializedClient
	}
}

// Function Transfer 转账
func (Self *HDWallet) Transfer(chainType common.ChainType, privateKey []byte, to string, token string, value float64) (string, error) {
	if client, ok := Self.clients[chainType]; ok {
		var hash string
		var err error
		switch chainType {
		case common.ChainType_ETH:
			if ethClient, ok := client.(*eth.Client); ok {
				hash, err = ethClient.Transfer(privateKey, to, token, value)
				if err != nil {
					return "", err
				}
			} else {
				return "", common.ErrUninitializedClient
			}
		case common.ChainType_BSC:
			if bscClient, ok := client.(*bsc.Client); ok {
				hash, err = bscClient.Transfer(privateKey, to, token, value)
				if err != nil {
					return "", err
				}
			} else {
				return "", common.ErrUninitializedClient
			}
		default:
			return "", common.ErrUnsupportedChainType
		}
		return hash, nil
	} else {
		return "", common.ErrUninitializedClient
	}
}

// Function GetTransaction 查询交易
func (Self *HDWallet) GetTransaction(chainType common.ChainType, txHash string) (*types.Transaction, error) {
	if client, ok := Self.clients[chainType]; ok {
		var tx *types.Transaction
		var err error
		switch chainType {
		case common.ChainType_ETH:
			if ethClient, ok := client.(*eth.Client); ok {
				tx, err = ethClient.GetTransaction(txHash)
				if err != nil {
					return nil, err
				}
			} else {
				return nil, common.ErrUninitializedClient
			}
		case common.ChainType_BSC:
			if bscClient, ok := client.(*bsc.Client); ok {
				tx, err = bscClient.GetTransaction(txHash)
				if err != nil {
					return nil, err
				}
			} else {
				return nil, common.ErrUninitializedClient
			}
		default:
			return nil, common.ErrUnsupportedChainType
		}
		return tx, nil
	} else {
		return nil, common.ErrUninitializedClient
	}
}
