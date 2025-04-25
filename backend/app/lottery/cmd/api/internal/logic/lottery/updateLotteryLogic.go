package lottery

import (
	"context"

	"looklook/app/lottery/cmd/api/internal/svc"
	"looklook/app/lottery/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateLotteryLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateLotteryLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateLotteryLogic {
	return &UpdateLotteryLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateLotteryLogic) UpdateLottery(req *types.UpdateLotteryReq) (resp *types.UpdateLotteryResp, err error) {
	// todo: add your logic here and delete this line

	return
}
