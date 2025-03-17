package svc

import (
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"looklook/app/lottery/cmd/rpc/internal/config"
	"looklook/app/lottery/model"
)

type ServiceContext struct {
	Config       config.Config
	RedisClient  *redis.Redis
	LotteryModel model.LotteryModel
	PrizeModel   model.PrizeModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:       c,
		LotteryModel: model.NewLotteryModel(sqlx.NewMysql(c.DB.DataSource), c.Cache),
		PrizeModel:   model.NewPrizeModel(sqlx.NewMysql(c.DB.DataSource), c.Cache),
	}
}
