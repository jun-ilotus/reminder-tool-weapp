package logic

import (
	"context"

	"looklook/app/lottery/cmd/rpc/internal/svc"
	"looklook/app/lottery/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type DelPrizeLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDelPrizeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DelPrizeLogic {
	return &DelPrizeLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DelPrizeLogic) DelPrize(in *pb.DelPrizeReq) (*pb.DelPrizeResp, error) {
	// todo: add your logic here and delete this line

	return &pb.DelPrizeResp{}, nil
}
