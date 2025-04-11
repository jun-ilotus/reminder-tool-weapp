package logic

import (
	"context"

	"looklook/app/signin/cmd/rpc/internal/svc"
	"looklook/app/signin/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateTaskFinishLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateTaskFinishLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateTaskFinishLogic {
	return &UpdateTaskFinishLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateTaskFinishLogic) UpdateTaskFinish(in *pb.UpdateTaskFinishReq) (*pb.UpdateTaskFinishResp, error) {
	// todo: add your logic here and delete this line

	return &pb.UpdateTaskFinishResp{}, nil
}
