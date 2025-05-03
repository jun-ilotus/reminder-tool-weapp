package logic

import (
	"context"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
	"looklook/app/reminder/cmd/rpc/internal/svc"
	"looklook/app/reminder/cmd/rpc/pb"
	"looklook/common/xerr"
)

type DoneReminderedLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDoneReminderedLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DoneReminderedLogic {
	return &DoneReminderedLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DoneReminderedLogic) DoneRemindered(in *pb.DoneReminderedReq) (*pb.DoneReminderedResp, error) {

	err := l.svcCtx.ReminderModel.DoneRemindered(l.ctx, in.Id, in.Status)
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "更新失败 err: %v", err)
	}

	return &pb.DoneReminderedResp{}, nil
}

//func (l *DoneReminderedLogic) DoneRemindered(in *pb.DoneReminderedReq) (*pb.DoneReminderedResp, error) {
//
//	reminders, err := l.svcCtx.ReminderModel.ReminderListByLessThanCurrentTime(l.ctx, time.Now())
//	if err != nil && err != model.ErrNotFound {
//		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "更新失败 err: %v", err)
//	}
//	if err == model.ErrNotFound {
//		return &pb.DoneReminderedResp{}, nil
//	}
//
//	for _, reminder := range reminders {
//
//		// 获取用户openID
//		userAuth, err := l.svcCtx.UsercenterRpc.GetUserAuthByUserId(l.ctx, &usercenter.GetUserAuthByUserIdReq{
//			UserId:   reminder.UserId,
//			AuthType: usercenterModel.UserAuthTypeSmallWX,
//		})
//		if err != nil {
//			return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "usercenter rpc Exception user_id : %+v , err: %v", reminder.UserId, err)
//		}
//
//		// 构建微信消息
//		msg := wxnotice.MessageReminder{
//			ReminderContent: wxnotice.Item{Value: reminder.Content},
//			ReminderTime:    wxnotice.Item{Value: reminder.ReminderTime.Format("2006-01-02 15:04")},
//		}
//		data, err := power.StructToHashMap(&msg)
//		// 拼接小程序页面地址
//		//pageAddr := fmt.Sprintf("pages/index/index?lotterId=%d&userId=%d", rpcLotteryInfo.Lottery.Id, in.UserId)
//		pageAddr := fmt.Sprintf("pages/reminderIndex/index")
//
//		templateId := msg.TemplateId()
//		reqData := &request.RequestSubscribeMessageSend{
//			ToUser:           userAuth.UserAuth.AuthKey,
//			TemplateID:       templateId,
//			Page:             pageAddr,
//			MiniProgramState: "formal",
//			Lang:             "zh_CN",
//			Data:             data,
//		}
//		// 发送消息
//		resp, err := l.svcCtx.WxMiniProgram.SubscribeMessage.Send(l.ctx, reqData)
//
//
//		if err != nil {
//			// 可重试的错误
//			return errors.Wrapf(ErrNotifyUserFail, "WxMiniRemindSigninHandler send message err:%v, reqData:%+v", err, reqData)
//		}
//		if resp.ErrCode != 0 {
//			switch resp.ErrCode {
//			// 不可重试的错误码
//			case WxErrCodeUserRefuseReceiveMsg:
//				logx.Infow("WxMiniRemindSigninHandler user refuse receive msg",
//					logx.Field("openid", remind.AuthKey),
//				)
//				return nil
//			default:
//				// 可重试的错误码，后续进行细分
//				return errors.Wrapf(ErrNotifyUserFail, "WxMiniRemindSigninHandler send message fail, errCode:%v, errMsg: %v, reqData:%+v", resp.ErrCode, resp.ErrMsg, reqData)
//			}
//		}
//
//
//		err := l.svcCtx.ReminderModel.DoneRemindered(l.ctx, reminder.Id, 1)
//		if err != nil {
//			return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "更新失败 err: %v", err)
//		}
//	}
//
//	return &pb.DoneReminderedResp{}, nil
//}
