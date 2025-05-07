package logic

import (
	"context"
	"fmt"
	"github.com/pkg/errors"
	"looklook/app/lottery/model"
	"looklook/common/constants"
	"looklook/common/xerr"
	"strconv"
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
		exists, err := l.svcCtx.RedisClient.ExistsCtx(l.ctx, lotteryJoinNumberKey)
		if err != nil {
			return nil, errors.Wrapf(xerr.NewErrMsg("redis库存更新错误"), "redis库存更新错误 err: %+v", err)
		}

		// redis 中没有库存的键值信息
		if !exists {
			// 通过 set 加锁操作 获得到锁的进行数据库查询 将库存保存到redis
			lotteryJoinNumberSetKey := fmt.Sprintf("%sset:%v", model.CacheLotteryJoinNumberIdPrefix, lottery.Id)

			// 若给定的 key 已经存在，则 SETNX 不做任何动作。 且 5 秒后解除锁 防止死锁
			exCtx, err := l.svcCtx.RedisClient.SetnxExCtx(l.ctx, lotteryJoinNumberSetKey, "1", 5)
			if err != nil {
				return nil, errors.Wrapf(xerr.NewErrMsg("redis设置锁错误"), "redis设置锁错误 err: %+v", err)
			}
			if !exCtx { // 没获得到锁  直接返回参与失败   同时也可以防止过快的请求
				return nil, errors.Wrapf(xerr.NewErrMsg("参与失败"), "没获得到锁")
			}

			// 获得到了锁 去数据库中查询库存放入 redis
			count, err := l.svcCtx.LotteryParticipationModel.GetLotteryParticipationCount(l.ctx, lottery.Id, 0)
			if err != nil {
				_, _ = l.svcCtx.RedisClient.DelCtx(l.ctx, lotteryJoinNumberSetKey) // 释放锁防止死锁
				return nil, errors.Wrapf(xerr.NewErrMsg("参与失败"), "数据出错 err: %+v", err)
			}
			count = lottery.JoinNumber - count // 当前剩余库存

			passSeconds := lottery.AwardDeadline.Unix() - now.Unix() // 截止时间减去现在时间
			// 设置库存
			_, err = l.svcCtx.RedisClient.SetnxExCtx(l.ctx, lotteryJoinNumberKey, strconv.FormatInt(count, 10), int(passSeconds))
			_, _ = l.svcCtx.RedisClient.DelCtx(l.ctx, lotteryJoinNumberSetKey) // 释放锁防止死锁
		}

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
