package kq

import (
	"context"
	"encoding/json"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/basicService/subscribeMessage/request"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
	"looklook/app/notice/cmd/mq/internal/svc"
	"looklook/app/notice/cmd/rpc/notice"
	"looklook/common/kqueue"
	"looklook/common/wxnotice"
	"looklook/common/xerr"
)

// 微信返回的错误码
const WxErrCodeUserRefuseReceiveMsg = 43101

var ErrNotifyUserFail = xerr.NewErrMsg("notify user fail")

/*
*
Listening to the payment flow status change notification message queue
*/
type NoticeSendMq struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewNoticeSendMq(ctx context.Context, svcCtx *svc.ServiceContext) *NoticeSendMq {
	return &NoticeSendMq{
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// Consume todo 老版本go-queue Consume(_, val string) error
// Consume 从队列中消费消息
func (l *NoticeSendMq) Consume(ctx context.Context, _, value string) error {
	var message kqueue.WXMiniNoticeSendMessage
	if err := json.Unmarshal([]byte(value), &message); err != nil {
		logx.WithContext(l.ctx).Error("NewNoticeSendMq->Consume Unmarshal err : %v , val : %s", err, value)
		return err
	}

	insert, err := l.svcCtx.NoticeRpc.AddNoticeLog(l.ctx, &notice.AddNoticeLogReq{
		UserId:     message.UserId,
		UserOpenid: message.ToUser,
		Status:     0,
		Error:      "",
	})
	if err != nil {
		return err
	}

	code, err := l.execService(message)
	if err != nil {
		logx.WithContext(l.ctx).Error("NewNoticeSendMq->execService  err : %v , val : %s , message:%+v", err, value, message)
		_, err = l.svcCtx.NoticeRpc.UpdateNoticeLog(l.ctx, &notice.UpdateNoticeLogReq{
			Id:         insert.Id,
			UserId:     message.UserId,
			UserOpenid: message.ToUser,
			Status:     int64(code),
			Error:      err.Error(),
		})
		return err
	}
	_, err = l.svcCtx.NoticeRpc.UpdateNoticeLog(l.ctx, &notice.UpdateNoticeLogReq{
		Id:         insert.Id,
		UserId:     message.UserId,
		UserOpenid: message.ToUser,
		Status:     int64(code),
	})

	if err != nil {
		return err
	}

	return nil
}

func (l *NoticeSendMq) execService(message kqueue.WXMiniNoticeSendMessage) (code int, err error) {
	reqData := &request.RequestSubscribeMessageSend{
		ToUser:           message.ToUser,
		TemplateID:       wxnotice.ReminderTemplateID,
		Page:             message.Page,
		MiniProgramState: "formal",
		Lang:             "zh_CN",
		Data:             message.Data,
	}
	resp, err := l.svcCtx.WxMiniProgram.SubscribeMessage.Send(l.ctx, reqData)
	if err != nil {
		// 可重试的错误
		return 2, errors.Wrapf(ErrNotifyUserFail, "WxMiniRemindSigninHandler send message err:%v, reqData:%+v", err, reqData)
	}
	if resp.ErrCode != 0 {
		switch resp.ErrCode {
		// 不可重试的错误码
		case WxErrCodeUserRefuseReceiveMsg:
			logx.Infow("WxMiniRemindSigninHandler user refuse receive msg",
				logx.Field("openid", message.ToUser),
			)
			return resp.ErrCode, nil
		default:
			// 可重试的错误码，后续进行细分
			return resp.ErrCode, errors.Wrapf(ErrNotifyUserFail, "WxMiniRemindSigninHandler send message fail, errCode:%v, errMsg: %v, reqData:%+v", resp.ErrCode, resp.ErrMsg, reqData)
		}
	}
	return 1, nil
}
