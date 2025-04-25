package lottery

import (
	"context"

	"looklook/app/lottery/cmd/api/internal/svc"
	"looklook/app/lottery/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type LotteryListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLotteryListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LotteryListLogic {
	return &LotteryListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LotteryListLogic) LotteryList(req *types.LotteryListReq) (resp *types.LotteryListResp, err error) {
	// todo: add your logic here and delete this line

	return
}
