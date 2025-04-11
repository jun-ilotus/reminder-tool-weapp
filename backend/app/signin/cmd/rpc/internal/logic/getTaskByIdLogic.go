package logic

import (
	"context"

	"looklook/app/signin/cmd/rpc/internal/svc"
	"looklook/app/signin/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetTaskByIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetTaskByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetTaskByIdLogic {
	return &GetTaskByIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetTaskByIdLogic) GetTaskById(in *pb.GetTaskByIdReq) (*pb.GetTaskByIdResp, error) {
	// todo: add your logic here and delete this line

	return &pb.GetTaskByIdResp{}, nil
}
