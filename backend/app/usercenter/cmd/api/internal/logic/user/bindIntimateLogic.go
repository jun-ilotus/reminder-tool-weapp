package user

import (
	"context"
	"looklook/app/usercenter/cmd/rpc/usercenter"
	"looklook/common/ctxdata"

	"looklook/app/usercenter/cmd/api/internal/svc"
	"looklook/app/usercenter/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type BindIntimateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewBindIntimateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *BindIntimateLogic {
	return &BindIntimateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *BindIntimateLogic) BindIntimate(req *types.BindIntimateReq) (resp *types.BindIntimateResp, err error) {

	userId := ctxdata.GetUidFromCtx(l.ctx)
	intimateUserId := ctxdata.ParseJWT(req.IntimateAccessToken, l.svcCtx.Config.JwtAuth.AccessSecret)

	_, err = l.svcCtx.UsercenterRpc.BindIntimate(l.ctx, &usercenter.BindIntimateReq{
		UserId:         userId,
		IntimateUserId: intimateUserId,
	})
	if err != nil {
		return nil, err
	}

	return &types.BindIntimateResp{
		IntimateUserId: intimateUserId,
	}, nil
}
