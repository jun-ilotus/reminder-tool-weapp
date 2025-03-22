package logic

import (
	"context"
	"github.com/pkg/errors"

	"looklook/app/usercenter/cmd/rpc/internal/svc"
	"looklook/app/usercenter/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type CancelBindIntimateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCancelBindIntimateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CancelBindIntimateLogic {
	return &CancelBindIntimateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CancelBindIntimateLogic) CancelBindIntimate(in *pb.CancelBindIntimateReq) (*pb.CancelBindIntimateResp, error) {
	err := l.svcCtx.UserModel.ModifyIntimateId(l.ctx, in.UserId, 0)
	if err != nil {
		return nil, errors.Wrapf(ErrGenerateTokenError, "BindIntimate err userId:%d , err:%v", in.UserId, err)
	}
	err = l.svcCtx.UserModel.ModifyIntimateId(l.ctx, in.IntimateUserId, 0)
	if err != nil {
		return nil, errors.Wrapf(ErrGenerateTokenError, "BindIntimate err userId:%d , err:%v", in.UserId, err)
	}
	return &pb.CancelBindIntimateResp{}, nil
}
