package logic

import (
	"context"
	"github.com/jinzhu/copier"
	"github.com/pkg/errors"
	"looklook/app/signin/model"

	"looklook/app/signin/cmd/rpc/internal/svc"
	"looklook/app/signin/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type SearchTaskLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSearchTaskLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SearchTaskLogic {
	return &SearchTaskLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SearchTaskLogic) SearchTask(in *pb.SearchTaskReq) (*pb.SearchTaskResp, error) {
	list, err := l.svcCtx.TaskModel.FindUserTaskList(l.ctx, in.UserId)
	if err != nil && !errors.Is(err, model.ErrNotFound) {
		return nil, err
	}

	var resp []*pb.Task
	if len(list) > 0 {
		for _, v := range list {
			var pbTask pb.Task
			_ = copier.Copy(&pbTask, v)
			resp = append(resp, &pbTask)
		}
	}

	return &pb.SearchTaskResp{Task: resp}, nil
}
