package svc

import (
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/zrpc"
	"looklook/app/lottery/cmd/rpc/internal/config"
	"looklook/app/lottery/model"

	"looklook/app/usercenter/cmd/rpc/usercenter"
)

type ServiceContext struct {
	Config                    config.Config
	UsercenterRpc             usercenter.Usercenter
	LotteryModel              model.LotteryModel
	PrizeModel                model.PrizeModel
	LotteryParticipationModel model.LotteryParticipationModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:                    c,
		UsercenterRpc:             usercenter.NewUsercenter(zrpc.MustNewClient(c.UsercenterRpcConf)),
		LotteryModel:              model.NewLotteryModel(sqlx.NewMysql(c.DB.DataSource), c.Cache),
		PrizeModel:                model.NewPrizeModel(sqlx.NewMysql(c.DB.DataSource), c.Cache),
		LotteryParticipationModel: model.NewLotteryParticipationModel(sqlx.NewMysql(c.DB.DataSource), c.Cache),
	}
}
