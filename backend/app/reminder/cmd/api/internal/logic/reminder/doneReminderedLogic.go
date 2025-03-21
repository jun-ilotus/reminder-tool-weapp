package reminder

import (
	"context"
	"github.com/pkg/errors"
	"looklook/app/reminder/cmd/rpc/reminder"
	"looklook/common/ctxdata"
	"looklook/common/xerr"

	"looklook/app/reminder/cmd/api/internal/svc"
	"looklook/app/reminder/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DoneReminderedLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDoneReminderedLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DoneReminderedLogic {
	return &DoneReminderedLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DoneReminderedLogic) DoneRemindered(req *types.DoneReminderedReq) (resp *types.DoneReminderedResp, err error) {
	userId := ctxdata.GetUidFromCtx(l.ctx)
	reminderById, err := l.svcCtx.ReminderRpc.GetReminderById(l.ctx, &reminder.GetReminderByIdReq{
		Id: req.Id,
	})
	if err != nil {
		return nil, err
	}
	if reminderById.Reminder.UserId != userId {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "该用户不是提醒创建者")
	}

	_, err = l.svcCtx.ReminderRpc.DoneRemindered(l.ctx, &reminder.DoneReminderedReq{
		Id:     req.Id,
		Status: req.Status,
	})
	if err != nil {
		return nil, err
	}

	return &types.DoneReminderedResp{}, nil
}
