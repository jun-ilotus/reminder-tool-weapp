package reminder

import (
	"context"
	"looklook/app/reminder/cmd/rpc/reminder"
	"looklook/common/ctxdata"

	"looklook/app/reminder/cmd/api/internal/svc"
	"looklook/app/reminder/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ModifyReminderLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewModifyReminderLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ModifyReminderLogic {
	return &ModifyReminderLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ModifyReminderLogic) ModifyReminder(req *types.ModifyReminderReq) (resp *types.ModifyReminderResp, err error) {
	userId := ctxdata.GetUidFromCtx(l.ctx)

	_, err = l.svcCtx.ReminderRpc.UpdateReminder(l.ctx, &reminder.UpdateReminderReq{
		UserId:       userId,
		Id:           req.Id,
		ReminderTime: req.ReminderTime,
		Content:      req.Content,
		Member:       req.Member,
		Status:       req.Status,
	})
	if err != nil {
		return nil, err
	}

	return &types.ModifyReminderResp{Id: int(req.Id)}, nil
}
