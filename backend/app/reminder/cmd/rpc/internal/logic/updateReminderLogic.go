package logic

import (
	"context"
	"github.com/hibiken/asynq"
	"github.com/jinzhu/copier"
	"github.com/pkg/errors"
	"k8s.io/apimachinery/pkg/util/json"
	"looklook/app/mqueue/cmd/job/jobtype"
	"looklook/app/reminder/model"
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

	// 提醒时间大于现在的时间才加入提醒
	if reminder.ReminderTime.Unix() > time.Now().Unix() {
		reminder.Status = 0
	}

	if reminder.Status == 0 {
		// 把需要发送的信息封装
		p := jobtype.WxMiniProgramNotifyUserPayload{
			ReminderId: reminder.Id,
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
	}

	err = l.svcCtx.ReminderModel.Update(l.ctx, reminder)
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "更新失败 err: %v", err)
	}
	return &pb.UpdateReminderResp{}, nil
}
