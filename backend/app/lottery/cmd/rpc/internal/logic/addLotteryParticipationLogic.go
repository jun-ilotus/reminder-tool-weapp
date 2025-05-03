package logic

import (
	"context"
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

// -----------------------鍙備笌鎶藉-----------------------
func (l *AddLotteryParticipationLogic) AddLotteryParticipation(in *pb.AddLotteryParticipationReq) (*pb.AddLotteryParticipationResp, error) {
	lotteryParticipation := &model.LotteryParticipation{
		LotteryId: in.LotteryId,
		UserId:    in.UserId,
	}

	lottery, err := l.svcCtx.LotteryModel.GetLotteryById(l.ctx, in.LotteryId)
	if err != nil {
		return nil, err
	}
	if lottery.IsAnnounced == 1 {
		return nil, errors.Wrap(xerr.NewErrMsg("该抽奖已开奖"), "该抽奖已开奖")
	}

	// todo 完善一下 防止超卖
	now := time.Now()
	switch lottery.AnnounceType {
	case constants.AnnounceTypeTimeLottery: // 按时间开奖
		if lottery.AnnounceTime.Before(now) {
			return nil, errors.Wrap(xerr.NewErrMsg("抽奖时间已过"), "抽奖时间已过")
		}
	case constants.AnnounceTypePeopleLottery: // 按人数开奖
		if lottery.AnnounceTime.Before(now) {
			return nil, errors.Wrap(xerr.NewErrMsg("抽奖时间已过"), "抽奖时间已过")
		}
		count, err := l.svcCtx.LotteryParticipationModel.GetLotteryParticipationCount(l.ctx, lottery.Id, 0)
		if err != nil {
			return nil, errors.Wrap(xerr.NewErrMsg("参与失败"), "数据出错")
		}
		if count >= lottery.JoinNumber {
			return nil, errors.Wrap(xerr.NewErrMsg("已经抽完了"), "超出抽奖人数限制")
		}
	case 3: // 即抽即中
		if lottery.AnnounceTime.After(now) {
			return nil, errors.Wrap(xerr.NewErrMsg("抽奖暂未开始"), "抽奖暂未开始")
		} else if lottery.AwardDeadline.Before(now) {
			return nil, errors.Wrap(xerr.NewErrMsg("抽奖已结束"), "抽奖时间已过")
		}
		count, err := l.svcCtx.LotteryParticipationModel.GetLotteryParticipationCount(l.ctx, lottery.Id, 0)
		if err != nil {
			return nil, errors.Wrap(xerr.NewErrMsg("参与失败"), "数据出错")
		}
		if count >= lottery.JoinNumber {
			return nil, errors.Wrap(xerr.NewErrMsg("已经抽完了"), "超出抽奖人数限制")
		}
	}

	id, err := l.svcCtx.LotteryParticipationModel.GetLotteryParticipationByUserIdLotteryId(l.ctx, in.UserId, lottery.Id)
	if err != nil && err != model.ErrNotFound {
		return nil, errors.Wrap(xerr.NewErrMsg("参与失败"), "数据出错")
	}
	if id != nil {
		return nil, errors.Wrap(xerr.NewErrMsg("已参与"), "已参与")
	}

	_, err = l.svcCtx.LotteryParticipationModel.Insert(l.ctx, lotteryParticipation)
	if err != nil {
		return nil, errors.Wrap(xerr.NewErrMsg("参与失败"), "数据出错")
	}

	return &pb.AddLotteryParticipationResp{}, nil
}
