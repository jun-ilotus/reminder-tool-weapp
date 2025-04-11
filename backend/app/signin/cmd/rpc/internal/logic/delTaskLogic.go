package logic

import (
	"context"

	"looklook/app/signin/cmd/rpc/internal/svc"
	"looklook/app/signin/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type DelTaskLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDelTaskLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DelTaskLogic {
	return &DelTaskLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DelTaskLogic) DelTask(in *pb.DelTaskReq) (*pb.DelTaskResp, error) {
	// todo: add your logic here and delete this line

	return &pb.DelTaskResp{}, nil
}
