package lottery

import (
	"context"
	"looklook/app/lottery/cmd/rpc/pb"
	"looklook/common/ctxdata"

	"looklook/app/lottery/cmd/api/internal/svc"
	"looklook/app/lottery/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateLotteryLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreateLotteryLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateLotteryLogic {
	return &CreateLotteryLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateLotteryLogic) CreateLottery(req *types.CreateLotteryReq) (resp *types.CreateLotteryResp, err error) {
	userId := ctxdata.GetUidFromCtx(l.ctx)

	lottery, err := l.svcCtx.LotteryRpc.AddLottery(l.ctx, &pb.AddLotteryReq{
		Id:            req.Id,
		UserId:        userId,
		Name:          req.Name,
		Thumb:         req.Thumb,
		PublishTime:   req.PublishTime,
		JoinNumber:    req.JoinNumber,
		Introduce:     req.Introduce,
		AwardDeadline: req.AwardDeadline,
		AnnounceType:  req.AnnounceType,
		AnnounceTime:  req.AnnounceTime,
		IsClocked:     req.IsClocked,
		ClockTaskId:   req.ClockTaskId,
	})
	if err != nil {
		return nil, err
	}
	lotteryId := lottery.Id
	for _, prize := range req.Prizes {
		pbPrize := &pb.AddPrizeReq{
			Id:        prize.Id,
			LotteryId: lotteryId,
			Type:      prize.Type,
			Name:      prize.Name,
			Level:     prize.Level,
			Thumb:     prize.Thumb,
			Count:     prize.Count,
			GrantType: prize.GrantType,
		}
		_, err := l.svcCtx.LotteryRpc.AddPrize(l.ctx, pbPrize)
		if err != nil {
			return nil, err
		}
	}

	return &types.CreateLotteryResp{Id: lottery.Id}, nil
}
