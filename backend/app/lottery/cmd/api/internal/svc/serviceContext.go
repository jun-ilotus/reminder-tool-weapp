package svc

import (
	"github.com/zeromicro/go-zero/zrpc"
	"looklook/app/lottery/cmd/api/internal/config"
	"looklook/app/lottery/cmd/rpc/lottery"
	"looklook/app/usercenter/cmd/rpc/usercenter"
)

type ServiceContext struct {
	Config        config.Config
	UsercenterRpc usercenter.Usercenter
	LotteryRpc    lottery.LotteryZrpcClient
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:        c,
		UsercenterRpc: usercenter.NewUsercenter(zrpc.MustNewClient(c.UsercenterRpcConf)),
		LotteryRpc:    lottery.NewLotteryZrpcClient(zrpc.MustNewClient(c.LotteryRpcConf)),
	}
}
