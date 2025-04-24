package logic

import (
	"fmt"
	"github.com/hibiken/asynq"
	"github.com/zeromicro/go-zero/core/logx"
	"looklook/app/mqueue/cmd/job/jobtype"
)

// scheduler job ------> go-zero-looklook/app/mqueue/cmd/job/internal/logic/remindSignin.go.
func (l *MqueueScheduler) remindSignin() {

	task := asynq.NewTask(jobtype.ScheduleRemindSignin, nil)
	// 这里只有 5 个参数
	entryID, err := l.svcCtx.Scheduler.Register("0 22 * * *", task)
	if err != nil {
		logx.WithContext(l.ctx).Errorf("!!!MqueueSchedulerErr!!! ====> 【remindSigninScheduler】 registered  err:%+v , task:%+v", err, task)
	}
	fmt.Printf("【remindSigninScheduler】 registered an  entry: %q \n", entryID)
}
