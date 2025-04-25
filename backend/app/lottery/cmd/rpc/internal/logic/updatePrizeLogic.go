package logic

import (
	"context"

	"looklook/app/lottery/cmd/rpc/internal/svc"
	"looklook/app/lottery/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdatePrizeLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdatePrizeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdatePrizeLogic {
	return &UpdatePrizeLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdatePrizeLogic) UpdatePrize(in *pb.UpdatePrizeReq) (*pb.UpdatePrizeResp, error) {
	// todo: add your logic here and delete this line

	return &pb.UpdatePrizeResp{}, nil
}
