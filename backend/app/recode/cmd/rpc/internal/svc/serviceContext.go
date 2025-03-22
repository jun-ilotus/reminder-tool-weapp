package svc

import (
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/zrpc"
	"looklook/app/recode/cmd/rpc/internal/config"
	"looklook/app/recode/model"
	"looklook/app/usercenter/cmd/rpc/usercenter"
)

type ServiceContext struct {
	Config        config.Config
	ItemsModel    model.ItemsModel
	RecodeModel   model.RecodeModel
	UsercenterRpc usercenter.Usercenter
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:        c,
		ItemsModel:    model.NewItemsModel(sqlx.NewMysql(c.DB.DataSource), c.Cache),
		RecodeModel:   model.NewRecodeModel(sqlx.NewMysql(c.DB.DataSource), c.Cache),
		UsercenterRpc: usercenter.NewUsercenter(zrpc.MustNewClient(c.UsercenterRpcConf)),
	}
}
