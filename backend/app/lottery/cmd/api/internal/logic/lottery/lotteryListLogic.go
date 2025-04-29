package lottery

import (
	"context"
	"github.com/jinzhu/copier"
	"looklook/app/lottery/cmd/rpc/lottery"

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
	res, err := l.svcCtx.LotteryRpc.SearchLottery(l.ctx, &lottery.SearchLotteryReq{
		LastId:     req.LastId,
		Limit:      req.PageSize,
		IsSelected: req.IsSelected,
	})
	if err != nil {
		return nil, err
	}

	var LotteryList []types.Lottery
	if len(res.Lottery) > 0 {
		for _, item := range res.Lottery {
			var t types.Lottery
			_ = copier.Copy(&t, item)
			LotteryList = append(LotteryList, t)
		}
	}

	return &types.LotteryListResp{List: LotteryList}, nil
}
