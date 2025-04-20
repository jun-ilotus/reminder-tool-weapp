package logic

import (
	"context"
	"github.com/jinzhu/copier"
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

type AddRecodeLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddRecodeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddRecodeLogic {
	return &AddRecodeLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// -----------------------recode-----------------------
func (l *AddRecodeLogic) AddRecode(in *pb.AddRecodeReq) (*pb.AddRecodeResp, error) {
	id := int64(0)
	taskList := make([]*pb.Task, 0)
	recode := &model.Recode{
		UserId:   in.UserId,
		SignDate: time.Unix(in.SignDate, 0),
	}
	// todo 用户两次签到过快 前一次插入事务还未完成 后一次查询不到签过到了  导致两次签到
	err := l.svcCtx.RecodeModel.Trans(l.ctx, func(ctx context.Context, session sqlx.Session) error {
		date, err := l.svcCtx.RecodeModel.FindLastOneByUserIdSignDate(l.ctx, recode.UserId)
		if err != nil && err != model.ErrNotFound {
			return errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "AddRecode recode Database Exception AddRecode : %+v , err: %v", recode, err)
		}
		if date != nil {
			return errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "今天签过到了")
		}

		insert, err := l.svcCtx.RecodeModel.TransInsert(l.ctx, session, recode)
		if err != nil {
			return errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "AddRecode recode Database Exception AddRecode : %+v , err: %v", recode, err)
		}
		id, _ = insert.LastInsertId()

		// 获取用户当前的连续签到天数   必须在同一个事务中才能查询到新插入的数据
		days, err := l.svcCtx.RecodeModel.FindCountUserNowConRecodeDays(l.ctx, session, in.UserId)
		if err != nil {
			return errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "AddRecode recode Database Exception AddRecode : %+v , err: %v", recode, err)
		}

		// 获取完成的任务
		tasks, err := l.svcCtx.TaskModel.FindUserConTaskDone(l.ctx, in.UserId, int64(*days))
		if err != nil && err != model.ErrNotFound {
			return errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "AddRecode recode Database Exception AddRecode : %+v , err: %v", recode, err)
		}
		if tasks != nil && len(tasks) > 0 {
			for _, task := range tasks {
				var pbTask pb.Task
				_ = copier.Copy(&pbTask, &task)
				taskList = append(taskList, &pbTask)
			}
		}

		return nil
	})
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	return &pb.AddRecodeResp{Id: id, Task: taskList}, nil
}
