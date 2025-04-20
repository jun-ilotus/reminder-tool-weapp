package logic

import (
	"context"
	"github.com/pkg/errors"
	"looklook/app/signin/model"
	"looklook/common/xerr"

	"looklook/app/signin/cmd/rpc/internal/svc"
	"looklook/app/signin/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetRemindStatusLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetRemindStatusLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetRemindStatusLogic {
	return &GetRemindStatusLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetRemindStatusLogic) GetRemindStatus(in *pb.GetRemindStatusReq) (*pb.GetRemindStatusResp, error) {
	remind, err := l.svcCtx.RemindModel.FindOneByUserId(l.ctx, in.UserId)
	if err != nil && err != model.ErrNotFound {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "ChangeSignRemind Database Exception FindOneByUserId : %+v , err: %v", in.UserId, err)
	}
	var status int64
	if remind != nil {
		status = remind.Status
	} else {
		status = 0
	}

	return &pb.GetRemindStatusResp{Status: status}, nil
}
