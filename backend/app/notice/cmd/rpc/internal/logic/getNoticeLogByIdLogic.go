package logic

import (
	"context"

	"looklook/app/notice/cmd/rpc/internal/svc"
	"looklook/app/notice/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetNoticeLogByIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetNoticeLogByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetNoticeLogByIdLogic {
	return &GetNoticeLogByIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetNoticeLogByIdLogic) GetNoticeLogById(in *pb.GetNoticeLogByIdReq) (*pb.GetNoticeLogByIdResp, error) {
	// todo: add your logic here and delete this line

	return &pb.GetNoticeLogByIdResp{}, nil
}
