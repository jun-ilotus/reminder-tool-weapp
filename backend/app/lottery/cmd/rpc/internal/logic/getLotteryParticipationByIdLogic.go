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

type GetLotteryParticipationByIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetLotteryParticipationByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetLotteryParticipationByIdLogic {
	return &GetLotteryParticipationByIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetLotteryParticipationByIdLogic) GetLotteryParticipationById(in *pb.GetLotteryParticipationByIdReq) (*pb.GetLotteryParticipationByIdResp, error) {
	one, err := l.svcCtx.LotteryParticipationModel.GetLotteryParticipationByUserIdLotteryId(l.ctx, in.UserId, in.LotteryId)
	if err != nil && err != model.ErrNotFound {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "GetLotteryParticipationByUserIdLotteryId Database Exception lottery err: %v", err)
	} else if err == model.ErrNotFound {
		return &pb.GetLotteryParticipationByIdResp{}, nil
	}
	resp := &pb.LotteryParticipation{}
	_ = copier.Copy(resp, one)

	return &pb.GetLotteryParticipationByIdResp{LotteryParticipation: resp}, nil
}
