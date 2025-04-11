package logic

import (
	"context"

	"looklook/app/signin/cmd/rpc/internal/svc"
	"looklook/app/signin/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type SearchRecodeLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSearchRecodeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SearchRecodeLogic {
	return &SearchRecodeLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SearchRecodeLogic) SearchRecode(in *pb.SearchRecodeReq) (*pb.SearchRecodeResp, error) {
	// todo: add your logic here and delete this line

	return &pb.SearchRecodeResp{}, nil
}
