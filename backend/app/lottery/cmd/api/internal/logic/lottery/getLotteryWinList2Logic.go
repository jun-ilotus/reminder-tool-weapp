package lottery

import (
	"context"
	"github.com/jinzhu/copier"
	"looklook/app/lottery/cmd/rpc/lottery"
	"looklook/app/usercenter/cmd/rpc/usercenter"

	"looklook/app/lottery/cmd/api/internal/svc"
	"looklook/app/lottery/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetLotteryWinList2Logic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetLotteryWinList2Logic(ctx context.Context, svcCtx *svc.ServiceContext) *GetLotteryWinList2Logic {
	return &GetLotteryWinList2Logic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetLotteryWinList2Logic) GetLotteryWinList2(req *types.GetLotteryWinList2Req) (resp *types.GetLotteryWinList2Resp, err error) {
	var respList []*types.WonList2

	// 查询出所有奖品
	prizes, err := l.svcCtx.LotteryRpc.SearchPrize(l.ctx, &lottery.SearchPrizeReq{
		LotteryId: req.LotteryId,
	})
	if err != nil {
		return nil, err
	}
	mpPrize := map[int64]*types.WonList2{}
	for _, prize := range prizes.Prize {
		var Tprize types.Prizes
		_ = copier.Copy(&Tprize, &prize)
		mpPrize[prize.Id] = &types.WonList2{ // 该奖品的用户列表  放在一个map中
			Prize:       &Tprize,
			WinnerCount: 0,
			Users:       make([]*types.UserInfo, 0),
		}
		respList = append(respList, mpPrize[prize.Id])
	}

	// 查询所有中奖者
	wins, err := l.svcCtx.LotteryRpc.SearchLotteryParticipationWin(l.ctx, &lottery.SearchLotteryParticipationWinReq{
		LotteryId: req.LotteryId,
	})
	if err != nil {
		return nil, err
	}

	for _, win := range wins.LotteryParticipation {
		// 查询该中奖者个人信息
		info, err := l.svcCtx.UsercenterRpc.GetUserInfo(l.ctx, &usercenter.GetUserInfoReq{Id: win.UserId})
		if err != nil {
			return nil, err
		}
		userInfo := &types.UserInfo{
			Nickname: info.User.Nickname,
			Avatar:   info.User.Avatar,
		}
		// 将该用户信息加在 他获得的奖品用户列表
		mpPrize[win.PrizeId].Users = append(mpPrize[win.PrizeId].Users, userInfo)
		mpPrize[win.PrizeId].WinnerCount++
	}

	return &types.GetLotteryWinList2Resp{List: respList}, nil
}
