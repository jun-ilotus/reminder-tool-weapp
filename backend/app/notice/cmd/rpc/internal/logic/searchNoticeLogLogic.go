package logic

import (
	"context"

	"looklook/app/notice/cmd/rpc/internal/svc"
	"looklook/app/notice/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type SearchNoticeLogLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSearchNoticeLogLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SearchNoticeLogLogic {
	return &SearchNoticeLogLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SearchNoticeLogLogic) SearchNoticeLog(in *pb.SearchNoticeLogReq) (*pb.SearchNoticeLogResp, error) {
	// todo: add your logic here and delete this line

	return &pb.SearchNoticeLogResp{}, nil
}
