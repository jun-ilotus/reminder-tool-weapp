package logic

import (
	"context"

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
	// todo: add your logic here and delete this line

	return &pb.SearchLotteryParticipationResp{}, nil
}
