package svc

import (
	"github.com/zeromicro/go-zero/zrpc"
	"looklook/app/signin/cmd/api/internal/config"
	"looklook/app/signin/cmd/rpc/signin"
)

type ServiceContext struct {
	Config    config.Config
	SigninRpc signin.Signin
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:    c,
		SigninRpc: signin.NewSignin(zrpc.MustNewClient(c.SigninRpcConf)),
	}
}
