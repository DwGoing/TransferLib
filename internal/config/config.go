package config

import (
	"abao/pkg/tron"

	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	zrpc.RpcServerConf
	Tron TronConf
}

type TronConf struct {
	Nodes      map[string]int               `json:","`
	ApiKeys    []string                     `json:","`
	Currencies map[string]tron.TronCurrency `json:","`
}
