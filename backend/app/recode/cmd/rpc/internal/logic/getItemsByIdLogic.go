package logic

import (
	"context"
	"github.com/jinzhu/copier"
	"github.com/pkg/errors"
	"looklook/common/xerr"

	"looklook/app/recode/cmd/rpc/internal/svc"
	"looklook/app/recode/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetItemsByIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetItemsByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetItemsByIdLogic {
	return &GetItemsByIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetItemsByIdLogic) GetItemsById(in *pb.GetItemsByIdReq) (*pb.GetItemsByIdResp, error) {
	items, err := l.svcCtx.ItemsModel.FindOne(l.ctx, in.Id)
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "items_id:%d,err:%v", in.Id, err)
	}
	pbItems := &pb.Items{}
	_ = copier.Copy(pbItems, items)
	return &pb.GetItemsByIdResp{Items: pbItems}, nil
}
