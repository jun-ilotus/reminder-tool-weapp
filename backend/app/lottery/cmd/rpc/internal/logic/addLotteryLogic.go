package logic

import (
	"context"

	"looklook/app/lottery/cmd/rpc/internal/svc"
	"looklook/app/lottery/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddLotteryLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddLotteryLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddLotteryLogic {
	return &AddLotteryLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// -----------------------鎶藉琛?----------------------
func (l *AddLotteryLogic) AddLottery(in *pb.AddLotteryReq) (*pb.AddLotteryResp, error) {
	// todo: add your logic here and delete this line

	return &pb.AddLotteryResp{}, nil
}
