package logic

import (
	"context"
	"github.com/pkg/errors"
	"looklook/common/xerr"

	"looklook/app/recode/cmd/rpc/internal/svc"
	"looklook/app/recode/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type DelItemsLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDelItemsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DelItemsLogic {
	return &DelItemsLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DelItemsLogic) DelItems(in *pb.DelItemsReq) (*pb.DelItemsResp, error) {
	one, err := l.svcCtx.ItemsModel.FindOne(l.ctx, in.Id)
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "查询事项id失败 err: %v", err)
	}
	if one.UserId != in.UserId {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "该用户不是事项创建者")
	}
	err = l.svcCtx.ItemsModel.Delete(l.ctx, in.Id)
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "items_id:%d,err:%v", in.Id, err)
	}

	return &pb.DelItemsResp{}, nil
}
