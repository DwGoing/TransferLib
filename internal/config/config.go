package config

import (
	"abao/pkg/bsc"
	"abao/pkg/eth"
	"abao/pkg/tron"

	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	zrpc.RpcServerConf
	Tron TronConf
	Eth  EthConf
	Bsc  BscConf
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

type BscConf struct {
	Nodes      map[string]int             `json:","`
	Currencies map[string]bsc.BscCurrency `json:","`
}
