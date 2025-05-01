package lottery

import (
	"context"
	"looklook/app/lottery/cmd/rpc/lottery"
	"looklook/app/usercenter/cmd/rpc/usercenter"

	"looklook/app/lottery/cmd/api/internal/svc"
	"looklook/app/lottery/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SearchParticipationLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSearchParticipationLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SearchParticipationLogic {
	return &SearchParticipationLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SearchParticipationLogic) SearchParticipation(req *types.SearchLotteryParticipationReq) (resp *types.SearchLotteryParticipationResp, err error) {
	participation, err := l.svcCtx.LotteryRpc.SearchLotteryParticipation(l.ctx, &lottery.SearchLotteryParticipationReq{
		LotteryId: req.LotteryId,
		LastId:    req.LastId,
		PageSize:  req.PageSize,
	})
	if err != nil {
		return nil, err
	}

	var list []*types.UserInfo
	for _, p := range participation.LotteryParticipation {
		info, err := l.svcCtx.UsercenterRpc.GetUserInfo(l.ctx, &usercenter.GetUserInfoReq{Id: p.UserId})
		if err != nil {
			return nil, err
		}
		userInfo := &types.UserInfo{
			Nickname: info.User.Nickname,
			Avatar:   info.User.Avatar,
		}
		list = append(list, userInfo)
	}
	resp = &types.SearchLotteryParticipationResp{
		List:  list,
		Count: participation.Count,
	}
	return resp, nil
}
