package config

import (
	"abao/pkg/eth"
	"abao/pkg/tron"

	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	zrpc.RpcServerConf
	Tron TronConf
	Eth  EthConf
}

type TronConf struct {
	Nodes      map[string]int               `json:","`
	ApiKeys    []string                     `json:","`
	Currencies map[string]tron.TronCurrency `json:","`
}

type EthConf struct {
	Nodes      map[string]int             `json:","`
	Currencies map[string]eth.EthCurrency `json:","`
}
