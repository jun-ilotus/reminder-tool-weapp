package logic

import (
	"context"

	"looklook/app/signin/cmd/rpc/internal/svc"
	"looklook/app/signin/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddTaskFinishLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddTaskFinishLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddTaskFinishLogic {
	return &AddTaskFinishLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// -----------------------taskFinish-----------------------
func (l *AddTaskFinishLogic) AddTaskFinish(in *pb.AddTaskFinishReq) (*pb.AddTaskFinishResp, error) {
	// todo: add your logic here and delete this line

	return &pb.AddTaskFinishResp{}, nil
}
