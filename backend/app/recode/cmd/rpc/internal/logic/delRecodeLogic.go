package logic

import (
	"context"
	"github.com/pkg/errors"
	"looklook/common/xerr"

	"looklook/app/recode/cmd/rpc/internal/svc"
	"looklook/app/recode/cmd/rpc/pb"

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
	one, err := l.svcCtx.RecodeModel.FindOne(l.ctx, in.Id)
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "查询记录id失败 err: %v", err)
	}
	if one.UserId != in.UserId { // todo intimateId
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "该用户不是记录创建者")
	}
	err = l.svcCtx.RecodeModel.Delete(l.ctx, in.Id)
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "items_id:%d,err:%v", in.Id, err)
	}

	return &pb.DelRecodeResp{}, nil
}
