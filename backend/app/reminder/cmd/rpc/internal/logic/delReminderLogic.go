package logic

import (
	"context"
	"github.com/pkg/errors"
	"looklook/common/xerr"

	"looklook/app/reminder/cmd/rpc/internal/svc"
	"looklook/app/reminder/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type DelReminderLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDelReminderLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DelReminderLogic {
	return &DelReminderLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DelReminderLogic) DelReminder(in *pb.DelReminderReq) (*pb.DelReminderResp, error) {
	one, err := l.svcCtx.ReminderModel.FindOne(l.ctx, in.Id)
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "查询提醒id失败 err: %v", err)
	}
	if one.UserId != in.UserId {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "该用户不是提醒创建者")
	}
	err = l.svcCtx.ReminderModel.Delete(l.ctx, in.Id)
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "reminder_id:%d,err:%v", in.Id, err)
	}

	return &pb.DelReminderResp{}, nil
}
