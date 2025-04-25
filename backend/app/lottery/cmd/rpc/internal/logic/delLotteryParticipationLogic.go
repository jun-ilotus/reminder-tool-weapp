package logic

import (
	"context"

	"looklook/app/lottery/cmd/rpc/internal/svc"
	"looklook/app/lottery/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type DelLotteryParticipationLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDelLotteryParticipationLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DelLotteryParticipationLogic {
	return &DelLotteryParticipationLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DelLotteryParticipationLogic) DelLotteryParticipation(in *pb.DelLotteryParticipationReq) (*pb.DelLotteryParticipationResp, error) {
	// todo: add your logic here and delete this line

	return &pb.DelLotteryParticipationResp{}, nil
}
