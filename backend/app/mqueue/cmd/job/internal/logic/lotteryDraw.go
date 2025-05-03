package logic

import (
	"context"
	"github.com/hibiken/asynq"
	"github.com/pkg/errors"
	"looklook/app/lottery/cmd/rpc/lottery"
	"looklook/app/mqueue/cmd/job/internal/svc"
	"looklook/common/constants"
	"looklook/common/xerr"
)

// var drawTypeList = []int64{constants.AnnounceTypeTimeLottery, constants.AnnounceTypePeopleLottery}
var drawTypeList = []int64{constants.AnnounceTypeTimeLottery, constants.AnnounceTypePeopleLottery}

var ErrLotteryDrawFail = xerr.NewErrMsg("lottery draw fail")

type LotteryDrawHandler struct {
	svcCtx *svc.ServiceContext
}

func NewLotteryDrawHandler(svcCtx *svc.ServiceContext) *LotteryDrawHandler {
	return &LotteryDrawHandler{
		svcCtx: svcCtx,
	}
}

func (l *LotteryDrawHandler) ProcessTask(ctx context.Context, _ *asynq.Task) error {
	// 遍历开奖类型列表，进行开奖
	for _, drawType := range drawTypeList {
		_, err := l.svcCtx.LotteryRpc.AnnounceLottery(ctx, &lottery.AnnounceLotteryReq{AnnounceType: drawType})
		if err != nil {
			return errors.Wrapf(ErrLotteryDrawFail, "LotteryDrawHandler announce lottery err:%v", err)
		}
	}

	return nil
}
