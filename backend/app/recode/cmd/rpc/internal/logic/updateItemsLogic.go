package logic

import (
	"context"
	"github.com/jinzhu/copier"
	"github.com/pkg/errors"
	"looklook/app/recode/model"
	"looklook/common/xerr"

	"looklook/app/recode/cmd/rpc/internal/svc"
	"looklook/app/recode/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateItemsLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateItemsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateItemsLogic {
	return &UpdateItemsLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateItemsLogic) UpdateItems(in *pb.UpdateItemsReq) (*pb.UpdateItemsResp, error) {
	one, err := l.svcCtx.ItemsModel.FindOne(l.ctx, in.Id)
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "查询事项id失败 err: %v", err)
	}
	if one.UserId != in.UserId {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "该用户不是事项创建者")
	}
	item := &model.Items{}
	_ = copier.Copy(item, in)
	err = l.svcCtx.ItemsModel.Update(l.ctx, item)
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "更新失败 err: %v", err)
	}
	return &pb.UpdateItemsResp{}, nil
}
