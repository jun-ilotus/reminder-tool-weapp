package logic

import (
	"context"
	"github.com/jinzhu/copier"
	"github.com/pkg/errors"
	"looklook/common/xerr"

	"looklook/app/reminder/cmd/rpc/internal/svc"
	"looklook/app/reminder/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetReminderByIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetReminderByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetReminderByIdLogic {
	return &GetReminderByIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetReminderByIdLogic) GetReminderById(in *pb.GetReminderByIdReq) (*pb.GetReminderByIdResp, error) {
	reminder, err := l.svcCtx.ReminderModel.FindOne(l.ctx, in.Id)
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "reminder_id:%d,err:%v", in.Id, err)
	}
	pbReminder := &pb.Reminder{}
	_ = copier.Copy(pbReminder, reminder)
	pbReminder.ReminderTime = reminder.ReminderTime.Unix()
	return &pb.GetReminderByIdResp{Reminder: pbReminder}, nil
}
