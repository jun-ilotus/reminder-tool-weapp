package svc

import (
	"github.com/zeromicro/go-zero/zrpc"
	"looklook/app/lottery/cmd/api/internal/config"
	"looklook/app/lottery/cmd/rpc/lottery"
)

type ServiceContext struct {
	Config     config.Config
	LotteryRpc lottery.LotteryZrpcClient
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:     c,
		LotteryRpc: lottery.NewLotteryZrpcClient(zrpc.MustNewClient(c.LotteryRpcConf)),
	}
}
