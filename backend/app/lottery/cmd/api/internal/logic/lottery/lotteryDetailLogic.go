package lottery

import (
	"context"
	"github.com/jinzhu/copier"
	"looklook/app/lottery/cmd/rpc/lottery"
	"looklook/common/ctxdata"

	"looklook/app/lottery/cmd/api/internal/svc"
	"looklook/app/lottery/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type LotteryDetailLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLotteryDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LotteryDetailLogic {
	return &LotteryDetailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LotteryDetailLogic) LotteryDetail(req *types.LotteryDetailReq) (resp *types.LotteryDetailResp, err error) {
	userId := ctxdata.GetUidFromCtx(l.ctx)
	RespLottery, err := l.svcCtx.LotteryRpc.GetLotteryById(l.ctx, &lottery.GetLotteryByIdReq{Id: req.Id})
	if err != nil {
		return nil, err
	}
	var ResLottery types.Lottery
	_ = copier.Copy(&ResLottery, RespLottery.Lottery)
	prizes, err := l.svcCtx.LotteryRpc.SearchPrize(l.ctx, &lottery.SearchPrizeReq{LotteryId: req.Id})
	if err != nil {
		return nil, err
	}
	var res []*types.Prize
	for _, v := range prizes.Prize {
		resPrize := &types.Prize{}
		_ = copier.Copy(resPrize, v)
		res = append(res, resPrize)
	}
	one, err := l.svcCtx.LotteryRpc.GetLotteryParticipationById(l.ctx,
		&lottery.GetLotteryParticipationByIdReq{UserId: userId, LotteryId: req.Id})
	if err != nil {
		return nil, err
	}
	isParticipated := 0
	if one.LotteryParticipation != nil {
		isParticipated = 1
	}

	return &types.LotteryDetailResp{
		Lottery:        ResLottery,
		Prizes:         res,
		IsParticipated: int64(isParticipated),
	}, nil
}
