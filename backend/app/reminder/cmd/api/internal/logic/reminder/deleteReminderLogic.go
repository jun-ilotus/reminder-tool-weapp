package reminder

import (
	"context"
	"looklook/app/reminder/cmd/rpc/reminder"
	"looklook/common/ctxdata"

	"looklook/app/reminder/cmd/api/internal/svc"
	"looklook/app/reminder/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteReminderLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteReminderLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteReminderLogic {
	return &DeleteReminderLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteReminderLogic) DeleteReminder(req *types.DeleteReminderReq) (resp *types.DeleteReminderResp, err error) {
	userId := ctxdata.GetUidFromCtx(l.ctx)

	_, err = l.svcCtx.ReminderRpc.DelReminder(l.ctx, &reminder.DelReminderReq{
		UserId: userId,
		Id:     int64(req.Id),
	})
	if err != nil {
		return nil, err
	}

	return &types.DeleteReminderResp{Id: req.Id}, nil
}
