package logic

import (
	"context"

	"looklook/app/lottery/cmd/rpc/internal/svc"
	"looklook/app/lottery/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateLotteryParticipationLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateLotteryParticipationLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateLotteryParticipationLogic {
	return &UpdateLotteryParticipationLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateLotteryParticipationLogic) UpdateLotteryParticipation(in *pb.UpdateLotteryParticipationReq) (*pb.UpdateLotteryParticipationResp, error) {
	// todo: add your logic here and delete this line

	return &pb.UpdateLotteryParticipationResp{}, nil
}
