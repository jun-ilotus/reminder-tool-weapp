package logic

import (
	"context"

	"looklook/app/lottery/cmd/rpc/internal/svc"
	"looklook/app/lottery/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type DelLotteryLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDelLotteryLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DelLotteryLogic {
	return &DelLotteryLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DelLotteryLogic) DelLottery(in *pb.DelLotteryReq) (*pb.DelLotteryResp, error) {
	// todo: add your logic here and delete this line

	return &pb.DelLotteryResp{}, nil
}
