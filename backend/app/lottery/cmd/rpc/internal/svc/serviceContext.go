package svc

import (
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/zrpc"
	"looklook/app/lottery/cmd/rpc/internal/config"
	"looklook/app/lottery/model"

	"looklook/app/usercenter/cmd/rpc/usercenter"
)

type ServiceContext struct {
	Config                    config.Config
	RedisClient               *redis.Redis
	UsercenterRpc             usercenter.Usercenter
	LotteryModel              model.LotteryModel
	PrizeModel                model.PrizeModel
	LotteryParticipationModel model.LotteryParticipationModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
		RedisClient: redis.New(c.Redis.Host, func(r *redis.Redis) {
			r.Type = c.Redis.Type
			r.Pass = c.Redis.Pass
		}),
		UsercenterRpc:             usercenter.NewUsercenter(zrpc.MustNewClient(c.UsercenterRpcConf)),
		LotteryModel:              model.NewLotteryModel(sqlx.NewMysql(c.DB.DataSource), c.Cache),
		PrizeModel:                model.NewPrizeModel(sqlx.NewMysql(c.DB.DataSource), c.Cache),
		LotteryParticipationModel: model.NewLotteryParticipationModel(sqlx.NewMysql(c.DB.DataSource), c.Cache),
	}
}
