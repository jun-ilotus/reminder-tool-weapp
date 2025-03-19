package svc

import (
	"github.com/zeromicro/go-zero/zrpc"
	"looklook/app/reminder/cmd/api/internal/config"
	"looklook/app/reminder/cmd/rpc/reminder"
)

type ServiceContext struct {
	Config      config.Config
	ReminderRpc reminder.ReminderZrpcClient
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:      c,
		ReminderRpc: reminder.NewReminderZrpcClient(zrpc.MustNewClient(c.ReminderRpcConf)),
	}
}
