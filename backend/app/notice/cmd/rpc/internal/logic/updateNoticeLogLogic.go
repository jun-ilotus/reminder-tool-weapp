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

type UpdateNoticeLogLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateNoticeLogLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateNoticeLogLogic {
	return &UpdateNoticeLogLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateNoticeLogLogic) UpdateNoticeLog(in *pb.UpdateNoticeLogReq) (*pb.UpdateNoticeLogResp, error) {
	noticeLog := &model.NoticeLog{
		Id:            in.Id,
		UserId:        in.UserId,
		UserOpenid:    in.UserOpenid,
		MsgTemplateId: in.MsgTemplateId,
		Status:        in.Status,
		Error:         in.Error,
	}
	err := l.svcCtx.NoticeModel.Update(l.ctx, noticeLog)
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "UpdateNoticeLog Database Exception noticelog : %+v , err: %v", noticeLog, err)
	}

	return &pb.UpdateNoticeLogResp{}, nil
}
