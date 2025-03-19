package logic

import (
	"context"
	"github.com/jinzhu/copier"
	"github.com/pkg/errors"
	"looklook/app/reminder/model"
	"looklook/common/xerr"
	"time"

	"looklook/app/reminder/cmd/rpc/internal/svc"
	"looklook/app/reminder/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateReminderLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateReminderLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateReminderLogic {
	return &UpdateReminderLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateReminderLogic) UpdateReminder(in *pb.UpdateReminderReq) (*pb.UpdateReminderResp, error) {
	one, err := l.svcCtx.ReminderModel.FindOne(l.ctx, in.Id)
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "查询提醒id失败 err: %v", err)
	}
	if one.UserId != in.UserId {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "该用户不是提醒创建者")
	}
	reminder := &model.Reminder{}
	_ = copier.Copy(reminder, in)
	reminder.ReminderTime = time.Unix(in.ReminderTime, 0)
	err = l.svcCtx.ReminderModel.Update(l.ctx, reminder)
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "更新失败 err: %v", err)
	}
	return &pb.UpdateReminderResp{}, nil
}
