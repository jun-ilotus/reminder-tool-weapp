package svc

import (
	"github.com/zeromicro/go-zero/zrpc"
	"looklook/app/recode/cmd/api/internal/config"
	"looklook/app/recode/cmd/rpc/recode"
)

type ServiceContext struct {
	Config    config.Config
	RecodeRpc recode.RecodeZrpcClient
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:    c,
		RecodeRpc: recode.NewRecodeZrpcClient(zrpc.MustNewClient(c.RecodeRpcConf)),
	}
}
