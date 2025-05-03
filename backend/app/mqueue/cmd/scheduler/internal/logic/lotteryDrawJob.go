package logic

import (
	"fmt"
	"github.com/hibiken/asynq"
	"github.com/zeromicro/go-zero/core/logx"
	"looklook/app/mqueue/cmd/job/jobtype"
)

// scheduler job ------> go-zero-looklook/app/mqueue/cmd/job/internal/logic/settleRecord.go.
func (l *MqueueScheduler) LotteryDrawScheduler() {

	task := asynq.NewTask(jobtype.ScheduleLotteryDraw, nil)
	// every one minute exec
	entryID, err := l.svcCtx.Scheduler.Register("*/1 * * * *", task)
	if err != nil {
		logx.WithContext(l.ctx).Errorf("!!!MqueueSchedulerErr!!! ====> 【LotteryDrawScheduler】 registered  err:%+v , task:%+v", err, task)
	}
	fmt.Printf("【LotteryDrawScheduler】 registered an entry: %q \n", entryID)
}
