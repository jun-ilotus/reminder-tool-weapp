package logic

import (
	"context"
	"github.com/pkg/errors"
	"looklook/app/lottery/model"
	"looklook/common/xerr"
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
	if lottery.AnnounceType == 1 && lottery.AnnounceTime.Before(now) {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.PARAM_ERROR_AnnounceTime), "开奖时间在发布时间之前")
	} else if lottery.AnnounceType == 2 && lottery.AnnounceTime.Before(now) {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.PARAM_ERROR_AnnounceTime), "开奖时间在发布时间之前")
	} else if lottery.AnnounceType == 3 && (lottery.AnnounceTime.Before(now) || lottery.AwardDeadline.Before(lottery.AnnounceTime)) {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.PARAM_ERROR_AnnounceTime), "开奖时间在发布时间之前")
	}

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
		}
	} else {
		insert, err := l.svcCtx.LotteryModel.Insert(l.ctx, lottery)
		if err != nil {
			return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "Insert Database Exception lottery : %+v , err: %v", lottery, err)
		}
		Id, _ = insert.LastInsertId()
	}

	return &pb.AddLotteryResp{Id: Id}, nil
}
