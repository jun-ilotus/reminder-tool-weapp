package logic

import (
	"context"

	"looklook/app/signin/cmd/rpc/internal/svc"
	"looklook/app/signin/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetTaskFinishByIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetTaskFinishByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetTaskFinishByIdLogic {
	return &GetTaskFinishByIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetTaskFinishByIdLogic) GetTaskFinishById(in *pb.GetTaskFinishByIdReq) (*pb.GetTaskFinishByIdResp, error) {
	// todo: add your logic here and delete this line

	return &pb.GetTaskFinishByIdResp{}, nil
}
