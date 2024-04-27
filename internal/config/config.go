package config

import (
	"github.com/DwGoing/transfer_lib/pkg/bsc"
	"github.com/DwGoing/transfer_lib/pkg/eth"
	"github.com/DwGoing/transfer_lib/pkg/tron"

	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	zrpc.RpcServerConf
	Tron TronConf
	Eth  EthConf
	Bsc  BscConf
}

type TronConf struct {
	Nodes      map[string]int           `json:","`
	ApiKeys    []string                 `json:","`
	Currencies map[string]tron.Currency `json:","`
}

type EthConf struct {
	Nodes      map[string]int          `json:","`
	Currencies map[string]eth.Currency `json:","`
}

type BscConf struct {
	Nodes      map[string]int          `json:","`
	Currencies map[string]bsc.Currency `json:","`
}
