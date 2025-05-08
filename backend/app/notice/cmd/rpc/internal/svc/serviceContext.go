package svc

import (
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"looklook/app/notice/cmd/rpc/internal/config"
	"looklook/app/notice/model"
)

type ServiceContext struct {
	Config      config.Config
	NoticeModel model.NoticeLogModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:      c,
		NoticeModel: model.NewNoticeLogModel(sqlx.NewMysql(c.DB.DataSource), c.Cache),
	}
}
