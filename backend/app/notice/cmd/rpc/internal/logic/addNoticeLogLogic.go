package logic

import (
	"context"
	"github.com/pkg/errors"
	"looklook/app/notice/model"
	"looklook/common/xerr"

	"looklook/app/notice/cmd/rpc/internal/svc"
	"looklook/app/notice/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddNoticeLogLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddNoticeLogLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddNoticeLogLogic {
	return &AddNoticeLogLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// -----------------------noticeLog-----------------------
func (l *AddNoticeLogLogic) AddNoticeLog(in *pb.AddNoticeLogReq) (*pb.AddNoticeLogResp, error) {
	noticeLog := &model.NoticeLog{
		UserId:        in.UserId,
		UserOpenid:    in.UserOpenid,
		MsgTemplateId: in.MsgTemplateId,
		Status:        in.Status,
		Error:         in.Error,
	}
	_, err := l.svcCtx.NoticeModel.Insert(l.ctx, noticeLog)
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "AddNoticeLog Database Exception noticelog : %+v , err: %v", noticeLog, err)
	}

	return &pb.AddNoticeLogResp{}, nil
}
