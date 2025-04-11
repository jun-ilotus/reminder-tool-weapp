package logic

import (
	"context"

	"looklook/app/signin/cmd/rpc/internal/svc"
	"looklook/app/signin/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateRecodeLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateRecodeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateRecodeLogic {
	return &UpdateRecodeLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateRecodeLogic) UpdateRecode(in *pb.UpdateRecodeReq) (*pb.UpdateRecodeResp, error) {
	// todo: add your logic here and delete this line

	return &pb.UpdateRecodeResp{}, nil
}
