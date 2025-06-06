package config

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/zrpc"
)

type KqConfig struct {
	Brokers []string
	Topic   string
}

type Config struct {
	zrpc.RpcServerConf
	DB struct {
		DataSource string
	}
	Cache             cache.CacheConf
	UsercenterRpcConf zrpc.RpcClientConf
	NoticeSendConf    KqConfig
}
