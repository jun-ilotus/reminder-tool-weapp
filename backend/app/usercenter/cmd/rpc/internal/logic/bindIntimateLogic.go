package logic

import (
	"context"
	"github.com/pkg/errors"

	"looklook/app/usercenter/cmd/rpc/internal/svc"
	"looklook/app/usercenter/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type BindIntimateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewBindIntimateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *BindIntimateLogic {
	return &BindIntimateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *BindIntimateLogic) BindIntimate(in *pb.BindIntimateReq) (*pb.BindIntimateResp, error) {

	if in.UserId == in.IntimateUserId {
		return nil, errors.Wrapf(ErrGenerateTokenError, "BindIntimate err userId:%d , IntimateUserId: %v, 不能设置自己", in.UserId, in.IntimateUserId)
	}

	err := l.svcCtx.UserModel.ModifyIntimateId(l.ctx, in.UserId, in.IntimateUserId)
	if err != nil {
		return nil, errors.Wrapf(ErrGenerateTokenError, "BindIntimate err userId:%d , err:%v", in.UserId, err)
	}
	err = l.svcCtx.UserModel.ModifyIntimateId(l.ctx, in.IntimateUserId, in.UserId)
	if err != nil {
		return nil, errors.Wrapf(ErrGenerateTokenError, "BindIntimate err userId:%d , err:%v", in.UserId, err)
	}

	return &pb.BindIntimateResp{}, nil
}
