package logic

import (
	"context"

	"looklook/app/signin/cmd/rpc/internal/svc"
	"looklook/app/signin/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type SearchTaskFinishLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSearchTaskFinishLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SearchTaskFinishLogic {
	return &SearchTaskFinishLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SearchTaskFinishLogic) SearchTaskFinish(in *pb.SearchTaskFinishReq) (*pb.SearchTaskFinishResp, error) {
	// todo: add your logic here and delete this line

	return &pb.SearchTaskFinishResp{}, nil
}
