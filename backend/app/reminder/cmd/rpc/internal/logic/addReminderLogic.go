package logic

import (
	"context"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"looklook/app/reminder/model"
	"looklook/common/xerr"
	"time"

	"looklook/app/reminder/cmd/rpc/internal/svc"
	"looklook/app/reminder/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddReminderLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddReminderLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddReminderLogic {
	return &AddReminderLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// -----------------------reminder-----------------------
func (l *AddReminderLogic) AddReminder(in *pb.AddReminderReq) (*pb.AddReminderResp, error) {
	var id int64
	err := l.svcCtx.ReminderModel.Trans(l.ctx, func(context context.Context, session sqlx.Session) error {
		logx.Error("in:", in)

		reminder := &model.Reminder{}
		reminder.UserId = in.UserId
		reminder.ReminderTime = time.Unix(in.ReminderTime, 0)
		reminder.Content = in.Content
		reminder.Member = in.Member

		insert, err := l.svcCtx.ReminderModel.TransInsert(l.ctx, session, reminder)

		// todo  if member == 1 2 给亲友加
		if err != nil {
			return errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "Reminder Database Exception reminder : %+v , err: %v", reminder, err)
		}
		id, _ = insert.LastInsertId()
		return nil
	})
	if err != nil {
		return nil, err
	}

	return &pb.AddReminderResp{
		Id: id,
	}, nil
}
