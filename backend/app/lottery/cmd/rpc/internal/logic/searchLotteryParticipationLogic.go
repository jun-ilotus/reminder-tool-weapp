package logic

import (
	"context"
	"github.com/jinzhu/copier"
	"github.com/pkg/errors"
	"looklook/app/lottery/model"
	"looklook/common/xerr"

	"looklook/app/lottery/cmd/rpc/internal/svc"
	"looklook/app/lottery/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type SearchLotteryParticipationLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSearchLotteryParticipationLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SearchLotteryParticipationLogic {
	return &SearchLotteryParticipationLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SearchLotteryParticipationLogic) SearchLotteryParticipation(in *pb.SearchLotteryParticipationReq) (*pb.SearchLotteryParticipationResp, error) {

	if in.LastId == 0 {
		lastId, err := l.svcCtx.LotteryParticipationModel.GetLotteryParticipationLastId(l.ctx, in.LotteryId)
		if err != nil && err != model.ErrNotFound {
			return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "GetLotteryParticipationLastId Database Exception lottery : %+v , err: %v", in.LotteryId, err)
		}
		in.LastId = lastId + 1
	}

	list, err := l.svcCtx.LotteryParticipationModel.LotteryParticipationList(l.ctx, in.LotteryId, in.PageSize, in.LastId)
	if err != nil && err != model.ErrNotFound {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "LotteryParticipationList Database Exception lottery : %+v , err: %v", in.LotteryId, err)
	}
	var pbList []*pb.LotteryParticipation
	for _, lotteryParticipation := range list {
		pbLotteryParticipation := &pb.LotteryParticipation{}
		_ = copier.Copy(pbLotteryParticipation, lotteryParticipation)
		pbList = append(pbList, pbLotteryParticipation)
	}

	count, err := l.svcCtx.LotteryParticipationModel.GetLotteryParticipationCount(l.ctx, in.LotteryId, in.IsWon)
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "GetLotteryParticipationCount Database Exception lottery : %+v , err: %v", in.LotteryId, err)
	}

	return &pb.SearchLotteryParticipationResp{LotteryParticipation: pbList, Count: count}, nil
}
