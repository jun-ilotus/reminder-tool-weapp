package logic

import (
	"context"
	"fmt"
	"github.com/pkg/errors"
	"looklook/app/lottery/model"
	"looklook/common/constants"
	"looklook/common/xerr"
	"time"

	"looklook/app/lottery/cmd/rpc/internal/svc"
	"looklook/app/lottery/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddLotteryParticipationLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddLotteryParticipationLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddLotteryParticipationLogic {
	return &AddLotteryParticipationLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

var addLuaScript = `
local key = KEYS[1]  -- 获取第一个键（库存的 Redis 键）

-- 获取当前库存值
local stock = redis.call('get', key)

-- 将库存值从字符串转换为数字
stock = tonumber(stock)

-- 检查库存是否足够
if stock >= 1 then
    -- 如果库存足够，扣减库存，返回剩余库存
    return redis.call('decrby', key, 1)
else
    -- 如果库存不足，返回-1
    return -1
end
`

// -----------------------鍙備笌鎶藉-----------------------
func (l *AddLotteryParticipationLogic) AddLotteryParticipation(in *pb.AddLotteryParticipationReq) (*pb.AddLotteryParticipationResp, error) {
	lotteryParticipation := &model.LotteryParticipation{
		LotteryId: in.LotteryId,
		UserId:    in.UserId,
	}

	lottery, err := l.svcCtx.LotteryModel.FindOne(l.ctx, in.LotteryId)
	if err != nil {
		return nil, err
	}
	if lottery.IsAnnounced == 1 {
		return nil, errors.Wrap(xerr.NewErrMsg("该抽奖已开奖"), "该抽奖已开奖")
	}
	id, err := l.svcCtx.LotteryParticipationModel.FindOneByLotteryIdUserId(l.ctx, lottery.Id, in.UserId)
	if err != nil && err != model.ErrNotFound {
		return nil, errors.Wrap(xerr.NewErrMsg("参与失败"), "数据出错")
	}
	if id != nil {
		return nil, errors.Wrap(xerr.NewErrMsg("已参与"), "已参与")
	}

	now := time.Now()
	switch lottery.AnnounceType {
	case constants.AnnounceTypeTimeLottery: // 按时间开奖
		if lottery.AnnounceTime.Before(now) {
			return nil, errors.Wrap(xerr.NewErrMsg("抽奖时间已过"), "抽奖时间已过")
		}
	case constants.AnnounceTypePeopleLottery: // 按人数开奖
		if lottery.AnnounceTime.After(now) {
			return nil, errors.Wrap(xerr.NewErrMsg("抽奖还未开始"), "抽奖还未开始")
		}

		//  防止超卖
		// 把前缀与id结合成 redisKey
		lotteryJoinNumberKey := fmt.Sprintf("%s%v", model.CacheLotteryJoinNumberIdPrefix, lottery.Id)
		result, err := l.svcCtx.RedisClient.EvalCtx(l.ctx, addLuaScript, []string{lotteryJoinNumberKey})
		if err != nil {
			return nil, errors.Wrapf(xerr.NewErrMsg("redis库存更新错误"), "redis库存更新错误 err: %+v", err)
		}
		if result.(int64) == -1 {
			return nil, errors.Wrap(xerr.NewErrMsg("已经抽完了"), "超出抽奖人数限制")
		}
	}

	_, err = l.svcCtx.LotteryParticipationModel.Insert(l.ctx, lotteryParticipation)
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrMsg("参与失败"), "数据出错 err: %+v", err)
	}

	return &pb.AddLotteryParticipationResp{}, nil
}
