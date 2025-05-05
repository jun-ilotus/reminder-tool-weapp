package lottery

import (
	"context"
	"looklook/app/lottery/cmd/api/internal/svc"
	"looklook/app/lottery/cmd/api/internal/types"
	"looklook/app/lottery/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddLotteryParticipationLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAddLotteryParticipationLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddLotteryParticipationLogic {
	return &AddLotteryParticipationLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AddLotteryParticipationLogic) AddLotteryParticipation(req *types.AddLotteryParticipationReq) (resp *types.AddLotteryParticipationResp, err error) {
	//userId := ctxdata.GetUidFromCtx(l.ctx)
	userId := req.UserId

	_, err = l.svcCtx.LotteryRpc.AddLotteryParticipation(l.ctx, &pb.AddLotteryParticipationReq{
		UserId:    userId,
		LotteryId: req.LotteryId,
	})
	if err != nil {
		return nil, err
	}

	return &types.AddLotteryParticipationResp{}, nil
}
