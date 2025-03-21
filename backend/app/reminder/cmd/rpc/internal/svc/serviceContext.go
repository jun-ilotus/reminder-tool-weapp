package svc

import (
	"github.com/hibiken/asynq"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/zrpc"
	"looklook/app/reminder/cmd/rpc/internal/config"
	"looklook/app/reminder/model"
	"looklook/app/usercenter/cmd/rpc/usercenter"
)

type ServiceContext struct {
	Config        config.Config
	ReminderModel model.ReminderModel
	UsercenterRpc usercenter.Usercenter
	AsynqClient   *asynq.Client
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:        c,
		ReminderModel: model.NewReminderModel(sqlx.NewMysql(c.DB.DataSource), c.Cache),
		UsercenterRpc: usercenter.NewUsercenter(zrpc.MustNewClient(c.UsercenterRpcConf)),
		AsynqClient:   newAsynqClient(c),
	}
}
