package reminder

import (
	"context"
	"looklook/app/reminder/cmd/rpc/reminder"
	"looklook/common/ctxdata"

	"looklook/app/reminder/cmd/api/internal/svc"
	"looklook/app/reminder/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateReminderLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreateReminderLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateReminderLogic {
	return &CreateReminderLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateReminderLogic) CreateReminder(req *types.CreateReminderReq) (resp *types.CreateReminderResp, err error) {
	userId := ctxdata.GetUidFromCtx(l.ctx)
	addReminder, err := l.svcCtx.ReminderRpc.AddReminder(l.ctx, &reminder.AddReminderReq{
		UserId:       userId,
		ReminderTime: req.ReminderTime,
		Content:      req.Content,
		Member:       req.Member,
	})
	if err != nil {
		return nil, err
	}
	return &types.CreateReminderResp{
		Id: int(addReminder.Id),
	}, nil
}
