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

type GetLotteryListByUserIdLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetLotteryListByUserIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetLotteryListByUserIdLogic {
	return &GetLotteryListByUserIdLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetLotteryListByUserIdLogic) GetLotteryListByUserId(req *types.GetLotteryListByUserIdReq) (resp *types.GetLotteryListByUserIdResp, err error) {
	userId := ctxdata.GetUidFromCtx(l.ctx)
	var respList []*types.LotteryPrizes
	if req.Type == 1 { // 全部参与的抽奖
		userLottery, err := l.svcCtx.LotteryRpc.SearchUserLotteryParticipation(l.ctx, &lottery.SearchUserLotteryParticipationReq{
			UserId:      userId,
			IsWon:       0,
			IsAnnounced: req.IsAnnounced,
		})
		if err != nil {
			return nil, err
		}
		for _, v := range userLottery.UserLottery {
			res := types.LotteryPrizes{}
			_ = copier.Copy(&res, v)
			respList = append(respList, &res)
		}
	} else if req.Type == 2 { // 用户创建的
		searchLottery, err := l.svcCtx.LotteryRpc.SearchLottery(l.ctx, &lottery.SearchLotteryReq{
			UserId:      userId,
			IsAnnounced: req.IsAnnounced,
		})
		if err != nil {
			return nil, err
		}
		for _, v := range searchLottery.Lottery {
			res := types.LotteryPrizes{
				Name:          v.Name,
				Thumb:         v.Thumb,
				PublishTime:   v.PublishTime,
				JoinNumber:    v.JoinNumber,
				AwardDeadline: v.AwardDeadline,
				AnnounceTime:  v.AnnounceTime,
				AnnounceType:  v.AnnounceType,
				IsAnnounced:   v.IsAnnounced,
				LotteryId:     v.Id,
				UserId:        v.UserId,
			}
			respList = append(respList, &res)
		}
	} else if req.Type == 3 { // 中奖的
		userLottery, err := l.svcCtx.LotteryRpc.SearchUserLotteryParticipation(l.ctx, &lottery.SearchUserLotteryParticipationReq{
			UserId:      userId,
			IsWon:       1,
			IsAnnounced: req.IsAnnounced,
		})
		if err != nil {
			return nil, err
		}
		for _, v := range userLottery.UserLottery {
			res := types.LotteryPrizes{}
			_ = copier.Copy(&res, v)

			// 获取所中奖品信息
			prizeById, err := l.svcCtx.LotteryRpc.GetPrizeById(l.ctx, &lottery.GetPrizeByIdReq{Id: v.PrizeId})
			if err != nil {
				return nil, err
			}
			res.PrizeName = prizeById.Prize.Name
			res.Level = prizeById.Prize.Level
			respList = append(respList, &res)
		}
	}

	return &types.GetLotteryListByUserIdResp{List: respList}, nil
}
