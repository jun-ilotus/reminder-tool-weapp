package reminder

import (
	"context"
	"github.com/jinzhu/copier"
	"looklook/app/reminder/cmd/rpc/reminder"
	"looklook/common/ctxdata"

	"looklook/app/reminder/cmd/api/internal/svc"
	"looklook/app/reminder/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ReminderListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewReminderListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ReminderListLogic {
	return &ReminderListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ReminderListLogic) ReminderList(req *types.ReminderListReq) (resp *types.ReminderListResp, err error) {
	userId := ctxdata.GetUidFromCtx(l.ctx)
	searchReminder, err := l.svcCtx.ReminderRpc.SearchReminder(l.ctx, &reminder.SearchReminderReq{
		Status: req.Status,
		UserId: userId,
	})
	if err != nil {
		return nil, err
	}

	var ReminderList []types.Reminder
	if len(searchReminder.Reminder) > 0 {
		for _, item := range searchReminder.Reminder {
			var t types.Reminder
			_ = copier.Copy(&t, &item)
			ReminderList = append(ReminderList, t)
		}
	}

	return &types.ReminderListResp{List: ReminderList}, nil
}
