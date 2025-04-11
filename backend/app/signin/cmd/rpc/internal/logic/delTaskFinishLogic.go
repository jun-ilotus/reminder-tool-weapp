package logic

import (
	"context"

	"looklook/app/signin/cmd/rpc/internal/svc"
	"looklook/app/signin/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type DelTaskFinishLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDelTaskFinishLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DelTaskFinishLogic {
	return &DelTaskFinishLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DelTaskFinishLogic) DelTaskFinish(in *pb.DelTaskFinishReq) (*pb.DelTaskFinishResp, error) {
	// todo: add your logic here and delete this line

	return &pb.DelTaskFinishResp{}, nil
}
