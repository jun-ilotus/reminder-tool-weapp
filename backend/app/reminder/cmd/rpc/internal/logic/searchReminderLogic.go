package logic

import (
	"context"
	"errors"
	"github.com/jinzhu/copier"
	"looklook/app/reminder/model"

	"looklook/app/reminder/cmd/rpc/internal/svc"
	"looklook/app/reminder/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type SearchReminderLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSearchReminderLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SearchReminderLogic {
	return &SearchReminderLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SearchReminderLogic) SearchReminder(in *pb.SearchReminderReq) (*pb.SearchReminderResp, error) {
	list, err := l.svcCtx.ReminderModel.ReminderList(l.ctx, in.Status, in.UserId)
	if err != nil && !errors.Is(err, model.ErrNotFound) {
		return nil, err
	}

	var resp []*pb.Reminder
	if len(list) > 0 {
		for _, reminder := range list {
			var pbReminder pb.Reminder
			_ = copier.Copy(&pbReminder, &reminder)
			pbReminder.ReminderTime = reminder.ReminderTime.Unix()
			pbReminder.CreateTime = reminder.CreateTime.Unix()
			pbReminder.UpdateTime = reminder.UpdateTime.Unix()
			resp = append(resp, &pbReminder)
		}
	}

	return &pb.SearchReminderResp{
		Reminder: resp,
	}, nil
}
