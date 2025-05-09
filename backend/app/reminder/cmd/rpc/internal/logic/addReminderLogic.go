package logic

import (
	"context"
	"github.com/hibiken/asynq"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"k8s.io/apimachinery/pkg/util/json"
	"looklook/app/mqueue/cmd/job/jobtype"
	"looklook/app/reminder/model"
	"looklook/app/usercenter/cmd/rpc/usercenter"
	"looklook/common/xerr"
	"time"

	"looklook/app/reminder/cmd/rpc/internal/svc"
	"looklook/app/reminder/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddReminderLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddReminderLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddReminderLogic {
	return &AddReminderLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// -----------------------reminder-----------------------
func (l *AddReminderLogic) AddReminder(in *pb.AddReminderReq) (*pb.AddReminderResp, error) {
	var id, memberId int64
	reminder := &model.Reminder{}
	reminder.UserId = in.UserId
	reminder.ReminderTime = time.Unix(in.ReminderTime, 0)
	reminder.Content = in.Content
	reminder.Member = in.Member
	userInfo, err := l.svcCtx.UsercenterRpc.GetUserInfo(l.ctx, &usercenter.GetUserInfoReq{
		Id: in.UserId,
	})
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "usercenter rpc Exception user_id : %+v , err: %v", in.UserId, err)
	}

	// 添加进数据库
	err = l.svcCtx.ReminderModel.Trans(l.ctx, func(context context.Context, session sqlx.Session) error {
		insert, err := l.svcCtx.ReminderModel.TransInsert(l.ctx, session, reminder)
		if err != nil {
			return errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "Reminder Database Exception reminder : %+v , err: %v", reminder, err)
		}
		id, _ = insert.LastInsertId()
		// todo  if member == 1 2 给亲友加   一次性订阅暂时无法完成

		if reminder.Member == 1 || reminder.Member == 2 {
			reminder.UserId = userInfo.User.IntimateId
			insert, err = l.svcCtx.ReminderModel.TransInsert(l.ctx, session, reminder)
			if err != nil {
				return errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "Reminder Database Exception reminder : %+v , err: %v", reminder, err)
			}
			memberId, _ = insert.LastInsertId()
		}

		return nil
	})
	if err != nil {
		return nil, err
	}

	// 把需要发送的信息封装
	p := jobtype.WxMiniProgramNotifyUserPayload{
		ReminderId: id,
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

	if memberId != 0 { // 把亲友消息也装入队列
		// 把需要发送的信息封装
		p := jobtype.WxMiniProgramNotifyUserPayload{
			ReminderId: memberId,
		}
		payload, err := json.Marshal(p)
		if err != nil {
			return nil, errors.Wrapf(xerr.NewErrCode(xerr.SERVER_COMMON_ERROR), "NoticeLotteryStart payload json marshal err:%+v, payload:%+v", err, p)
		}
		_, err = l.svcCtx.AsynqClient.Enqueue(asynq.NewTask(jobtype.MsgWxMiniProgramNotifyUser, payload), asynq.ProcessIn(subTime))
		if err != nil {
			return nil, errors.Wrapf(xerr.NewErrCode(xerr.SERVER_COMMON_ERROR), "NoticeLotteryStart payload json marshal err:%+v, payload:%+v", err, p)
		}
	}

	return &pb.AddReminderResp{
		Id: id,
	}, nil
}
