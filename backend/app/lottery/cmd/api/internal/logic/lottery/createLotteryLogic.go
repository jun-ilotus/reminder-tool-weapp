package lottery

import (
	"context"

	"looklook/app/lottery/cmd/api/internal/svc"
	"looklook/app/lottery/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateLotteryLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreateLotteryLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateLotteryLogic {
	return &CreateLotteryLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateLotteryLogic) CreateLottery(req *types.CreateLotteryReq) (resp *types.CreateLotteryResp, err error) {
	// todo: add your logic here and delete this line

	return
}
