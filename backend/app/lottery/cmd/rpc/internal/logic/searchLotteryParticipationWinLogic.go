package logic

import (
	"context"
	"github.com/jinzhu/copier"
	"github.com/pkg/errors"
	"looklook/common/xerr"

	"looklook/app/lottery/cmd/rpc/internal/svc"
	"looklook/app/lottery/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type SearchLotteryParticipationWinLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSearchLotteryParticipationWinLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SearchLotteryParticipationWinLogic {
	return &SearchLotteryParticipationWinLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SearchLotteryParticipationWinLogic) SearchLotteryParticipationWin(in *pb.SearchLotteryParticipationWinReq) (*pb.SearchLotteryParticipationWinResp, error) {
	count, err := l.svcCtx.LotteryParticipationModel.GetLotteryParticipationCount(l.ctx, in.LotteryId, 1)
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "GetLotteryParticipationCount Database Exception lottery : %+v , err: %v", in.LotteryId, err)
	}
	var resp []*pb.LotteryParticipation
	if count > 0 {
		winList, err := l.svcCtx.LotteryParticipationModel.LotteryParticipationWinList(l.ctx, in.LotteryId)
		if err != nil {
			return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "LotteryParticipationWinList Database Exception lottery : %+v , err: %v", in.LotteryId, err)
		}
		for _, win := range winList {
			var pbWin pb.LotteryParticipation
			_ = copier.Copy(&pbWin, &win)
			resp = append(resp, &pbWin)
		}
	}

	return &pb.SearchLotteryParticipationWinResp{LotteryParticipation: resp, Count: count}, nil
}
