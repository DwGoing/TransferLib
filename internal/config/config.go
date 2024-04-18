package config

import "github.com/zeromicro/go-zero/zrpc"

type Config struct {
	zrpc.RpcServerConf
	Tron TronConf
}

type TronConf struct {
	Nodes   map[string]int `json:","`
	ApiKeys []string       `json:","`
}
