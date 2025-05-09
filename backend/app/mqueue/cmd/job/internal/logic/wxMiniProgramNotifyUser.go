package logic

import (
	"context"
	"encoding/json"
	"github.com/hibiken/asynq"
	"github.com/zeromicro/go-zero/core/logx"
	"looklook/app/mqueue/cmd/job/internal/svc"
	"looklook/app/mqueue/cmd/job/jobtype"
	"looklook/app/reminder/cmd/rpc/reminder"
	"looklook/common/xerr"
)

// 微信返回的错误码
const WxErrCodeUserRefuseReceiveMsg = 43101

var ErrNotifyUserFail = xerr.NewErrMsg("notify user fail")

// WxMiniProgramNotifyUserHandler mini program notify user
type WxMiniProgramNotifyUserHandler struct {
	svcCtx *svc.ServiceContext
}

func NewWxMiniProgramNotifyUserHandler(svcCtx *svc.ServiceContext) *WxMiniProgramNotifyUserHandler {
	return &WxMiniProgramNotifyUserHandler{
		svcCtx: svcCtx,
	}
}

func (l *WxMiniProgramNotifyUserHandler) ProcessTask(ctx context.Context, t *asynq.Task) error {
	var (
		err error
		p   jobtype.WxMiniProgramNotifyUserPayload
	)
	if err = json.Unmarshal(t.Payload(), &p); err != nil {
		// 不可重试的错误，记录日志
		logx.Error("WxMiniProgramNotifyUserHandler ProcessTask payload json unmarshal err",
			logx.Field("err", err))
		return nil
	}

	_, err = l.svcCtx.ReminderRpc.DoneRemindered(ctx, &reminder.DoneReminderedReq{
		Id: p.ReminderId,
	})

	// 判断用户是否允许接收订阅消息
	//preference, err := l.svcCtx.NoticeRpc.GetNoticeSubscribePreference(ctx, &pb.GetNoticeSubscribePreferenceReq{
	//	Openid:     p.OpenId,
	//	TemplateId: templateId,
	//})
	//if err != nil {
	//	// 不可重试的错误
	//	logx.Error("WxMiniProgramNotifyUserHandler ProcessTask NoticeRpc.GetNoticeSubscribePreference err",
	//		logx.Field("err", err),
	//		logx.Field("openid", p.OpenId),
	//		logx.Field("templateId", templateId),
	//	)
	//	return nil
	//}

	//if preference.AcceptCount == 0 {
	//	// 不可重试的错误
	//	logx.WithContext(ctx).Errorw("WxMiniProgramNotifyUser user reject to receive msg",
	//		logx.Field("openid", p.OpenId),
	//		logx.Field("templateId", templateId),
	//	)
	//	return nil
	//}

	return nil
}
