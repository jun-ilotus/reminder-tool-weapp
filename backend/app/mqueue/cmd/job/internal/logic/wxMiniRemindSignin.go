package logic

import (
	"context"
	"github.com/hibiken/asynq"
	"looklook/app/mqueue/cmd/job/internal/svc"
	"looklook/app/signin/cmd/rpc/signin"
)

// 微信返回的错误码
//const WxErrCodeUserRefuseReceiveMsg = 43101

//var ErrNotifyUserFail = xerr.NewErrMsg("notify user fail")

// WxMiniRemindSigninHandler mini program notify user
type WxMiniRemindSigninHandler struct {
	svcCtx *svc.ServiceContext
}

func NewWxMiniRemindSigninHandler(svcCtx *svc.ServiceContext) *WxMiniRemindSigninHandler {
	return &WxMiniRemindSigninHandler{
		svcCtx: svcCtx,
	}
}

func (l *WxMiniRemindSigninHandler) ProcessTask(ctx context.Context, _ *asynq.Task) error {
	var (
		err error
	)
	_, err = l.svcCtx.SigninRpc.SendRemind(ctx, &signin.SendRemindReq{})
	if err != nil {
		return err
	}

	return nil
}
