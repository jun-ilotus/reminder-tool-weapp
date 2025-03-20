package svc

import (
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"looklook/app/recode/cmd/rpc/internal/config"
	"looklook/app/recode/model"
)

type ServiceContext struct {
	Config      config.Config
	ItemsModel  model.ItemsModel
	RecodeModel model.RecodeModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:      c,
		ItemsModel:  model.NewItemsModel(sqlx.NewMysql(c.DB.DataSource), c.Cache),
		RecodeModel: model.NewRecodeModel(sqlx.NewMysql(c.DB.DataSource), c.Cache),
	}
}
