package logic

import (
	"context"
	"looklook/app/lottery/model"

	"looklook/app/lottery/cmd/rpc/internal/svc"
	"looklook/app/lottery/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddPrizeLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddPrizeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddPrizeLogic {
	return &AddPrizeLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// -----------------------濂栧搧琛?----------------------
func (l *AddPrizeLogic) AddPrize(in *pb.AddPrizeReq) (*pb.AddPrizeResp, error) {
	prize := &model.Prize{
		LotteryId: in.LotteryId,
		Type:      in.Type,
		Name:      in.Name,
		Level:     in.Level,
		Thumb:     in.Thumb,
		Count:     in.Count,
		GrantType: in.GrantType,
	}

	if in.Id != 0 {
		prize.Id = in.Id
		err := l.svcCtx.PrizeModel.Update(l.ctx, prize)
		if err != nil {
			return nil, err
		}
	} else {
		_, err := l.svcCtx.PrizeModel.Insert(l.ctx, prize)
		if err != nil {
			return nil, err
		}
	}

	return &pb.AddPrizeResp{}, nil
}
