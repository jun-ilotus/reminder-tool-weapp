package logic

import (
	"context"
	"looklook/app/signin/cmd/rpc/internal/svc"
	"looklook/app/signin/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddTaskFinishRollbackLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddTaskFinishRollbackLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddTaskFinishRollbackLogic {
	return &AddTaskFinishRollbackLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *AddTaskFinishRollbackLogic) AddTaskFinishRollback(in *pb.AddTaskFinishReq) (*pb.AddTaskFinishResp, error) {
	for _, task := range in.Task {
		_ = l.svcCtx.TaskFinishModel.DeleteByUserIdTaskId(l.ctx, in.UserId, task.Id)
	}
	return &pb.AddTaskFinishResp{}, nil
}
