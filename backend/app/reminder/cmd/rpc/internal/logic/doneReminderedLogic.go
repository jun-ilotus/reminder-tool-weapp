package logic

import (
	"context"
	"github.com/pkg/errors"
	"looklook/common/xerr"

	"looklook/app/reminder/cmd/rpc/internal/svc"
	"looklook/app/reminder/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type DoneReminderedLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDoneReminderedLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DoneReminderedLogic {
	return &DoneReminderedLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DoneReminderedLogic) DoneRemindered(in *pb.DoneReminderedReq) (*pb.DoneReminderedResp, error) {

	err := l.svcCtx.ReminderModel.DoneRemindered(l.ctx, in.Id, in.Status)
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "更新失败 err: %v", err)
	}

	return &pb.DoneReminderedResp{}, nil
}
