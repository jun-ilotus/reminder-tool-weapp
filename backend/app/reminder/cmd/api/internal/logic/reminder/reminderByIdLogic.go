package reminder

import (
	"context"
	"github.com/jinzhu/copier"
	"looklook/app/reminder/cmd/rpc/reminder"

	"looklook/app/reminder/cmd/api/internal/svc"
	"looklook/app/reminder/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ReminderByIdLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewReminderByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ReminderByIdLogic {
	return &ReminderByIdLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ReminderByIdLogic) ReminderById(req *types.ReminderByIdReq) (resp *types.ReminderByIdResp, err error) {
	Reminder, err := l.svcCtx.ReminderRpc.GetReminderById(l.ctx, &reminder.GetReminderByIdReq{Id: int64(req.Id)})
	if err != nil {
		return nil, err
	}
	var t types.Reminder
	_ = copier.Copy(&t, &Reminder.Reminder)

	return &types.ReminderByIdResp{Reminder: t}, nil
}
