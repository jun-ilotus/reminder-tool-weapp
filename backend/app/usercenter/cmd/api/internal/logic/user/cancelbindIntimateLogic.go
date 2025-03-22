package user

import (
	"context"
	"looklook/app/usercenter/cmd/rpc/usercenter"
	"looklook/common/ctxdata"

	"looklook/app/usercenter/cmd/api/internal/svc"
	"looklook/app/usercenter/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CancelbindIntimateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCancelbindIntimateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CancelbindIntimateLogic {
	return &CancelbindIntimateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CancelbindIntimateLogic) CancelbindIntimate(req *types.CancelBindIntimateReq) (resp *types.CancelBindIntimateResp, err error) {
	userId := ctxdata.GetUidFromCtx(l.ctx)
	userInfo, err := l.svcCtx.UsercenterRpc.GetUserInfo(l.ctx, &usercenter.GetUserInfoReq{
		Id: userId,
	})
	if err != nil {
		return nil, err
	}
	_, err = l.svcCtx.UsercenterRpc.CancelBindIntimate(l.ctx, &usercenter.CancelBindIntimateReq{
		UserId:         userId,
		IntimateUserId: userInfo.User.IntimateId,
	})
	if err != nil {
		return nil, err
	}

	return &types.CancelBindIntimateResp{}, nil
}
