package logic

import (
	"context"
	"github.com/pkg/errors"
	"looklook/app/signin/model"
	"looklook/common/xerr"

	"looklook/app/signin/cmd/rpc/internal/svc"
	"looklook/app/signin/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type ChangeSignRemindLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewChangeSignRemindLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ChangeSignRemindLogic {
	return &ChangeSignRemindLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// -----------------------remind-----------------------
func (l *ChangeSignRemindLogic) ChangeSignRemind(in *pb.ChangeRemindReq) (*pb.ChangeRemindResp, error) {
	remind, err := l.svcCtx.RemindModel.FindOneByUserId(l.ctx, in.UserId)
	if err != nil && err != model.ErrNotFound {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "ChangeSignRemind Database Exception FindOneByUserId : %+v , err: %v", in.UserId, err)
	}
	var id int64
	if remind == nil {
		remind := &model.Remind{
			UserId: in.UserId,
			Status: in.Status,
		}
		insert, err := l.svcCtx.RemindModel.Insert(l.ctx, remind)
		if err != nil {
			return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "ChangeSignRemind Database Exception Insert : %+v , err: %v", remind, err)
		}
		id, _ = insert.LastInsertId()
	} else {
		if remind.Status == in.Status {
			return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "已修改")
		}
		remind.Status = in.Status
		err := l.svcCtx.RemindModel.Update(l.ctx, remind)
		if err != nil {
			return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "ChangeSignRemind Database Exception Update : %+v , err: %v", remind, err)
		}
		id = remind.Id
	}

	return &pb.ChangeRemindResp{Id: id, Status: in.Status}, nil
}
