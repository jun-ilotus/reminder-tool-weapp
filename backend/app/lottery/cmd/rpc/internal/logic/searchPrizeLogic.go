package logic

import (
	"context"

	"looklook/app/lottery/cmd/rpc/internal/svc"
	"looklook/app/lottery/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type SearchPrizeLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSearchPrizeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SearchPrizeLogic {
	return &SearchPrizeLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SearchPrizeLogic) SearchPrize(in *pb.SearchPrizeReq) (*pb.SearchPrizeResp, error) {
	// todo: add your logic here and delete this line

	return &pb.SearchPrizeResp{}, nil
}
