package svc

import (
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"looklook/app/reminder/cmd/rpc/internal/config"
	"looklook/app/reminder/model"
)

type ServiceContext struct {
	Config        config.Config
	ReminderModel model.ReminderModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:        c,
		ReminderModel: model.NewReminderModel(sqlx.NewMysql(c.DB.DataSource), c.Cache),
	}
}
