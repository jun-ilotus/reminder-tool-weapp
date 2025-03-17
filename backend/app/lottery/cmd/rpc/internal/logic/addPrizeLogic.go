package logic

import (
	"context"

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
	// todo: add your logic here and delete this line

	return &pb.AddPrizeResp{}, nil
}
