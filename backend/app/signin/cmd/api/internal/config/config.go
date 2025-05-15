package config

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	rest.RestConf
	JwtAuth struct {
		AccessSecret string
	}
	Cache             cache.CacheConf
	SigninRpcConf     zrpc.RpcClientConf
	UsercenterRpcConf zrpc.RpcClientConf
	DtmServerConf     struct {
		DtmServer string
	}
}
