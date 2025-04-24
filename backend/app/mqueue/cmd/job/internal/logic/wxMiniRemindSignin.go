package logic

import (
	"context"
	"fmt"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/basicService/subscribeMessage/request"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/kernel/power"
	"github.com/hibiken/asynq"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
	"looklook/app/mqueue/cmd/job/internal/svc"
	"looklook/app/signin/cmd/rpc/pb"
	"looklook/common/wxnotice"
	"time"
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
		err        error
		templateId string
	)
	// 构建微信消息
	now := time.Now()
	msg := wxnotice.MessageReminder{
		ReminderContent: wxnotice.Item{Value: "签到提醒"},
		ReminderTime:    wxnotice.Item{Value: now.Format("2006-01-02 15:04")},
	}
	data, err := power.StructToHashMap(&msg)

	// 拼接小程序页面地址
	//pageAddr := fmt.Sprintf("pages/index/index?lotterId=%d&userId=%d", rpcLotteryInfo.Lottery.Id, in.UserId)
	pageAddr := fmt.Sprintf("pages/signin/index/index")

	reminds, err := l.svcCtx.SigninRpc.SearchRemind(ctx, &pb.SearchRemindReq{})
	if err != nil {
		return errors.Wrapf(ErrNotifyUserFail, "WxMiniRemindSigninHandler SearchRemind err:%v", err)
	}
	templateId = msg.TemplateId()
	for _, remind := range reminds.Reminds {
		reqData := &request.RequestSubscribeMessageSend{
			ToUser:           remind.AuthKey,
			TemplateID:       templateId,
			Page:             pageAddr,
			MiniProgramState: "formal",
			Lang:             "zh_CN",
			Data:             data,
		}
		// 发送消息
		resp, err := l.svcCtx.WxMiniProgram.SubscribeMessage.Send(ctx, reqData)

		if err != nil {
			// 可重试的错误
			return errors.Wrapf(ErrNotifyUserFail, "WxMiniRemindSigninHandler send message err:%v, reqData:%+v", err, reqData)
		}
		if resp.ErrCode != 0 {
			switch resp.ErrCode {
			// 不可重试的错误码
			case WxErrCodeUserRefuseReceiveMsg:
				logx.Infow("WxMiniRemindSigninHandler user refuse receive msg",
					logx.Field("openid", remind.AuthKey),
				)
				return nil
			default:
				// 可重试的错误码，后续进行细分
				return errors.Wrapf(ErrNotifyUserFail, "WxMiniRemindSigninHandler send message fail, errCode:%v, errMsg: %v, reqData:%+v", resp.ErrCode, resp.ErrMsg, reqData)
			}
		}
	}

	return nil
}
