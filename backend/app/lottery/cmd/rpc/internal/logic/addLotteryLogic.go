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

type AddLotteryLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddLotteryLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddLotteryLogic {
	return &AddLotteryLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

var luaScript = `
local key = KEYS[1]  -- 获取第一个键（库存的 Redis 键）
local num = tonumber(ARGV[1]) -- 获取第一个参数（需要扣减的数量），并转换为数字

-- 获取当前库存值（从 Redis 中获取的值是字符串类型）
local stock = redis.call('get', key)
if stock == false then -- 如果键不存在，返回 -1
    return -1
end

-- 将库存值从字符串转换为数字
stock = tonumber(stock)

-- 检查库存是否足够
if stock >= num then
    -- 如果库存足够，扣减库存
    return redis.call('decrby', key, num)
else
    -- 如果库存不足，将库存减少为 0
    return redis.call('decrby', key, stock)
end
`

// -----------------------鎶藉琛?----------------------
func (l *AddLotteryLogic) AddLottery(in *pb.AddLotteryReq) (*pb.AddLotteryResp, error) {
	Id := int64(0)
	lottery := &model.Lottery{
		UserId:        in.UserId,
		Name:          in.Name,
		Thumb:         in.Thumb,
		PublishTime:   time.Unix(in.PublishTime, 0),
		JoinNumber:    in.JoinNumber,
		Introduce:     in.Introduce,
		AwardDeadline: time.Unix(in.AwardDeadline, 0),
		AnnounceType:  in.AnnounceType,
		AnnounceTime:  time.Unix(in.AnnounceTime, 0),
		IsClocked:     in.IsClocked,
		ClockTaskId:   in.ClockTaskId,
	}
	now := time.Now()
	if lottery.AnnounceType == constants.AnnounceTypeTimeLottery && lottery.AnnounceTime.Before(now) {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.PARAM_ERROR_AnnounceTime), "开奖时间在现在之前")
	} else if lottery.AnnounceType == constants.AnnounceTypePeopleLottery && lottery.AnnounceTime.Before(now) && in.Id == 0 {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.PARAM_ERROR_AnnounceTime), "开始时间在现在时间之前")
	} else if lottery.AnnounceType == constants.AnnounceTypePeopleLottery && lottery.AwardDeadline.Before(lottery.AnnounceTime) {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.PARAM_ERROR_AnnounceTime), "截止时间在开始时间之前")
	}
	// 计算截止时间到现在距离的秒数
	passSeconds := lottery.AwardDeadline.Unix() - now.Unix()

	// todo 绑定打卡任务

	if in.Id != 0 {
		lotteryUpdate, err := l.svcCtx.LotteryModel.FindOne(l.ctx, in.Id)
		if err != nil {
			return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "FindOne Database Exception lottery : %+v , err: %v", lotteryUpdate, err)
		}
		if lotteryUpdate != nil {
			if lotteryUpdate.UserId != lottery.UserId {
				return nil, errors.Wrapf(xerr.NewErrCode(xerr.PARAM_ERROR_UserId), "AddLottery Database Exception lottery : %+v , err: %v", lotteryUpdate, err)
			}
			lottery.Id = lotteryUpdate.Id
			err := l.svcCtx.LotteryModel.Update(l.ctx, lottery)
			if err != nil {
				return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "Update Database Exception lottery : %+v , err: %v", lottery, err)
			}
			Id = lottery.Id

			// 把前缀与id结合成 redisKey
			lotteryJoinNumberKey := fmt.Sprintf("%s%v", model.CacheLotteryJoinNumberIdPrefix, lottery.Id)
			num := lotteryUpdate.JoinNumber - lottery.JoinNumber //更新前 - 更新后   计算需要扣减的库存
			if num != 0 {
				result, err := l.svcCtx.RedisClient.EvalCtx(l.ctx, luaScript, []string{lotteryJoinNumberKey}, num)
				if err != nil {
					return nil, errors.Wrapf(xerr.NewErrMsg("redis库存更新错误"), "redis库存更新错误 err: %+v", err)
				}

				// 库存数据不存在  rdb.Eval 的返回值 result 是一个 interface{} 类型，具体类型取决于 Lua 脚本的返回值。
				if result.(int64) == -1 {
					// 将库存数据保存到 redis 中
					_, err = l.svcCtx.RedisClient.SetnxExCtx(l.ctx, lotteryJoinNumberKey, strconv.FormatInt(lottery.JoinNumber, 10), int(passSeconds))
					if err != nil {
						return nil, errors.Wrapf(xerr.NewErrMsg("redis存放库存错误"), "redis存放库存错误 err: %+v", err)
					}
				}
			}
		}
	} else {
		insert, err := l.svcCtx.LotteryModel.Insert(l.ctx, lottery)
		if err != nil {
			return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "Insert Database Exception lottery : %+v , err: %v", lottery, err)
		}
		Id, _ = insert.LastInsertId()

		// 把前缀与id结合成 redisKey
		lotteryJoinNumberKey := fmt.Sprintf("%s%v", model.CacheLotteryJoinNumberIdPrefix, Id)

		// 将库存数据保存到 redis 中
		_, err = l.svcCtx.RedisClient.SetnxExCtx(l.ctx, lotteryJoinNumberKey, strconv.FormatInt(lottery.JoinNumber, 10), int(passSeconds))
		if err != nil {
			return nil, errors.Wrapf(xerr.NewErrMsg("redis存放库存错误"), "redis存放库存错误 err: %+v", err)
		}
	}

	return &pb.AddLotteryResp{Id: Id}, nil
}
