package svc

import (
	"github.com/ArtisanCloud/PowerWeChat/v3/src/miniProgram"
	"github.com/hibiken/asynq"
	"github.com/silenceper/wechat/v2/miniprogram"
	"github.com/zeromicro/go-zero/zrpc"
	"looklook/app/mqueue/cmd/job/internal/config"
	"looklook/app/reminder/cmd/rpc/reminder"
	"looklook/app/usercenter/cmd/rpc/usercenter"
)

type ServiceContext struct {
	Config        config.Config
	AsynqServer   *asynq.Server
	MiniProgram   *miniprogram.MiniProgram
	WxMiniProgram *miniProgram.MiniProgram // 使用

	ReminderRpc   reminder.ReminderZrpcClient
	UsercenterRpc usercenter.Usercenter
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:        c,
		AsynqServer:   newAsynqServer(c),
		MiniProgram:   newMiniprogramClient(c),
		WxMiniProgram: MustNewMiniProgram(c), // 使用
		ReminderRpc:   reminder.NewReminderZrpcClient(zrpc.MustNewClient(c.ReminderRpcConf)),
		UsercenterRpc: usercenter.NewUsercenter(zrpc.MustNewClient(c.UsercenterRpcConf)),
	}
}
