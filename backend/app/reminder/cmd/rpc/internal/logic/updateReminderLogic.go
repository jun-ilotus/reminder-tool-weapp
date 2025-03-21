package logic

import (
	"context"
	"fmt"
	"github.com/hibiken/asynq"
	"github.com/jinzhu/copier"
	"github.com/pkg/errors"
	"k8s.io/apimachinery/pkg/util/json"
	"looklook/app/mqueue/cmd/job/jobtype"
	"looklook/app/reminder/model"
	"looklook/app/usercenter/cmd/rpc/usercenter"
	usercenterModel "looklook/app/usercenter/model"
	"looklook/common/wxnotice"
	"looklook/common/xerr"
	"time"

	"looklook/app/reminder/cmd/rpc/internal/svc"
	"looklook/app/reminder/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateReminderLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateReminderLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateReminderLogic {
	return &UpdateReminderLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateReminderLogic) UpdateReminder(in *pb.UpdateReminderReq) (*pb.UpdateReminderResp, error) {
	one, err := l.svcCtx.ReminderModel.FindOne(l.ctx, in.Id)
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "查询提醒id失败 err: %v", err)
	}
	if one.UserId != in.UserId {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "该用户不是提醒创建者")
	}
	reminder := &model.Reminder{}
	_ = copier.Copy(reminder, in)
	reminder.ReminderTime = time.Unix(in.ReminderTime, 0)

	// 获取用户openID
	userAuth, err := l.svcCtx.UsercenterRpc.GetUserAuthByUserId(l.ctx, &usercenter.GetUserAuthByUserIdReq{
		UserId:   in.UserId,
		AuthType: usercenterModel.UserAuthTypeSmallWX,
	})
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "usercenter rpc Exception user_id : %+v , err: %v", in.UserId, err)
	}

	// 构建微信消息
	msg := wxnotice.MessageReminder{
		ReminderContent: wxnotice.Item{Value: reminder.Content},
		ReminderTime:    wxnotice.Item{Value: reminder.ReminderTime.Format("2006-01-02 15:04")},
	}
	marshal, err := json.Marshal(msg)

	// 拼接小程序页面地址
	//pageAddr := fmt.Sprintf("pages/index/index?lotterId=%d&userId=%d", rpcLotteryInfo.Lottery.Id, in.UserId)
	pageAddr := fmt.Sprintf("pages/index/index")

	// 把需要发送的信息封装
	p := jobtype.WxMiniProgramNotifyUserPayload{
		MsgType:      msg.Type(),
		OpenId:       userAuth.UserAuth.AuthKey, //必选，接收者（用户）的 openid
		PageAddr:     pageAddr,                  //可选，点击模板卡片后的跳转页面，仅限本小程序内的页面。支持带参数,（示例index?foo=bar）。该字段不填则模板无跳转。
		Data:         string(marshal),
		ReminderId:   int(reminder.Id),
		ReminderTime: reminder.ReminderTime.Unix(),
	}
	payload, err := json.Marshal(p)
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.SERVER_COMMON_ERROR), "NoticeLotteryStart payload json marshal err:%+v, payload:%+v", err, p)
	}

	// 计算发送时间
	reminderTime := time.Unix(in.ReminderTime, 0)
	now := time.Now()
	subTime := reminderTime.Sub(now)

	// 装入队列
	_, err = l.svcCtx.AsynqClient.Enqueue(asynq.NewTask(jobtype.MsgWxMiniProgramNotifyUser, payload), asynq.ProcessIn(subTime))
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.SERVER_COMMON_ERROR), "NoticeLotteryStart payload json marshal err:%+v, payload:%+v", err, p)
	}

	err = l.svcCtx.ReminderModel.Update(l.ctx, reminder)
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "更新失败 err: %v", err)
	}
	return &pb.UpdateReminderResp{}, nil
}
