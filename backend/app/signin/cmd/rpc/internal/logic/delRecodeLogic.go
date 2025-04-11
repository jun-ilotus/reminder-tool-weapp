package logic

import (
	"context"

	"looklook/app/signin/cmd/rpc/internal/svc"
	"looklook/app/signin/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type DelRecodeLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDelRecodeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DelRecodeLogic {
	return &DelRecodeLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DelRecodeLogic) DelRecode(in *pb.DelRecodeReq) (*pb.DelRecodeResp, error) {
	// todo: add your logic here and delete this line

	return &pb.DelRecodeResp{}, nil
}
