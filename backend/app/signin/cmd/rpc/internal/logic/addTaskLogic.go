package logic

import (
	"context"

	"looklook/app/signin/cmd/rpc/internal/svc"
	"looklook/app/signin/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddTaskLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddTaskLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddTaskLogic {
	return &AddTaskLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// -----------------------task-----------------------
func (l *AddTaskLogic) AddTask(in *pb.AddTaskReq) (*pb.AddTaskResp, error) {
	// todo: add your logic here and delete this line

	return &pb.AddTaskResp{}, nil
}
