package logic

import (
	"context"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"looklook/app/signin/model"
	"looklook/common/xerr"
	"time"

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
	points := int64(0)
	err := l.svcCtx.TaskFinishModel.Trans(l.ctx, func(context context.Context, session sqlx.Session) error {
		for _, task := range in.Task {
			taskFinish := &model.TaskFinish{
				TaskId:     task.Id,
				UserId:     in.UserId,
				FinishDate: time.Unix(in.FinishDate, 0),
				Points:     task.Points,
			}
			points += taskFinish.Points
			_, err := l.svcCtx.TaskFinishModel.TransInsert(l.ctx, session, taskFinish)
			if err != nil {
				return errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "AddTaskFinish recode Database Exception AddTaskFinish : %+v , err: %v", taskFinish, err)
			}
		}
		return nil
	})
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &pb.AddTaskFinishResp{Points: points}, nil
}
