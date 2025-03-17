package lottery

import (
	"context"
	"github.com/jinzhu/copier"
	"github.com/pkg/errors"
	"looklook/app/lottery/cmd/rpc/lottery"
	"looklook/common/xerr"

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

func (l *LotteryListLogic) LotteryList(req *types.LotteryListReq) (*types.LotteryListResp, error) {
	resp, err := l.svcCtx.LotteryRpc.SearchLottery(l.ctx, &lottery.SearchLotteryReq{
		Page:       req.Page,
		Limit:      req.PageSize,
		IsSelected: req.IsSelected,
	})
	if err != nil {
		//todo 要使用这种写法管理错误，否则Kibana无法收集到错误日志的详情
		return nil, errors.Wrapf(xerr.NewErrMsg("Failed to get SearchLottery"), "Failed to get SearchLottery err : %v ,req:%+v", err, req)
	}

	var LotteryList []types.Lottery
	if len(resp.Lottery) > 0 {
		for _, item := range resp.Lottery {
			var t types.Lottery
			_ = copier.Copy(&t, item)
			LotteryList = append(LotteryList, t)
		}
	}

	return &types.LotteryListResp{List: LotteryList}, nil
}
