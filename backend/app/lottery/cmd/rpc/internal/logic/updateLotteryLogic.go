package logic

import (
	"context"

	"looklook/app/lottery/cmd/rpc/internal/svc"
	"looklook/app/lottery/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateLotteryLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateLotteryLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateLotteryLogic {
	return &UpdateLotteryLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateLotteryLogic) UpdateLottery(in *pb.UpdateLotteryReq) (*pb.UpdateLotteryResp, error) {
	// todo: add your logic here and delete this line

	return &pb.UpdateLotteryResp{}, nil
}
