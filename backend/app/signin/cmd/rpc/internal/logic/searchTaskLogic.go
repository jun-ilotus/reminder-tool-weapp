package logic

import (
	"context"

	"looklook/app/signin/cmd/rpc/internal/svc"
	"looklook/app/signin/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type SearchTaskLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSearchTaskLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SearchTaskLogic {
	return &SearchTaskLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SearchTaskLogic) SearchTask(in *pb.SearchTaskReq) (*pb.SearchTaskResp, error) {
	// todo: add your logic here and delete this line

	return &pb.SearchTaskResp{}, nil
}
