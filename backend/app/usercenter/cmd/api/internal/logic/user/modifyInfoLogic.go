package user

import (
	"context"
	"looklook/app/usercenter/cmd/rpc/usercenter"
	"looklook/common/ctxdata"

	"looklook/app/usercenter/cmd/api/internal/svc"
	"looklook/app/usercenter/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ModifyInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewModifyInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ModifyInfoLogic {
	return &ModifyInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ModifyInfoLogic) ModifyInfo(req *types.ModifyUserInfoReq) (resp *types.ModifyUserInfoResp, err error) {
	userId := ctxdata.GetUidFromCtx(l.ctx)
	_, err = l.svcCtx.UsercenterRpc.ModifyUserInfo(l.ctx, &usercenter.ModifyUserInfoReq{
		UserId:   userId,
		Nickname: req.Nickname,
		Avatar:   req.Avatar,
	})
	if err != nil {
		return nil, err
	}

	return
}
