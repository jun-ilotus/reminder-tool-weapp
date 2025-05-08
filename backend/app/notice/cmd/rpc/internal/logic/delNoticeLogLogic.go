package logic

import (
	"context"

	"looklook/app/notice/cmd/rpc/internal/svc"
	"looklook/app/notice/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type DelNoticeLogLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDelNoticeLogLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DelNoticeLogLogic {
	return &DelNoticeLogLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DelNoticeLogLogic) DelNoticeLog(in *pb.DelNoticeLogReq) (*pb.DelNoticeLogResp, error) {
	// todo: add your logic here and delete this line

	return &pb.DelNoticeLogResp{}, nil
}
