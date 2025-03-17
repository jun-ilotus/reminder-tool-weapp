package logic

import (
	"context"

	"looklook/app/lottery/cmd/rpc/internal/svc"
	"looklook/app/lottery/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetLotteryByIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetLotteryByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetLotteryByIdLogic {
	return &GetLotteryByIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetLotteryByIdLogic) GetLotteryById(in *pb.GetLotteryByIdReq) (*pb.GetLotteryByIdResp, error) {
	// todo: add your logic here and delete this line

	return &pb.GetLotteryByIdResp{}, nil
}
