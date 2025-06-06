package svc

import (
	"github.com/zeromicro/go-queue/kq"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/zrpc"
	"looklook/app/signin/cmd/rpc/internal/config"
	"looklook/app/signin/model"
	"looklook/app/usercenter/cmd/rpc/usercenter"
)

type ServiceContext struct {
	Config                 config.Config
	UsercenterRpc          usercenter.Usercenter
	RecodeModel            model.RecodeModel
	TaskModel              model.TaskModel
	TaskFinishModel        model.TaskFinishModel
	RemindModel            model.RemindModel
	KqueueNoticeSendClient *kq.Pusher
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:                 c,
		RecodeModel:            model.NewRecodeModel(sqlx.NewMysql(c.DB.DataSource), c.Cache),
		TaskModel:              model.NewTaskModel(sqlx.NewMysql(c.DB.DataSource), c.Cache),
		TaskFinishModel:        model.NewTaskFinishModel(sqlx.NewMysql(c.DB.DataSource), c.Cache),
		RemindModel:            model.NewRemindModel(sqlx.NewMysql(c.DB.DataSource), c.Cache),
		UsercenterRpc:          usercenter.NewUsercenter(zrpc.MustNewClient(c.UsercenterRpcConf)),
		KqueueNoticeSendClient: kq.NewPusher(c.NoticeSendConf.Brokers, c.NoticeSendConf.Topic),
	}
}
