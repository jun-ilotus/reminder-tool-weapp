package items

import (
	"context"
	"github.com/jinzhu/copier"
	"looklook/app/recode/cmd/rpc/recode"

	"looklook/app/recode/cmd/api/internal/svc"
	"looklook/app/recode/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ItemsByIdLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewItemsByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ItemsByIdLogic {
	return &ItemsByIdLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ItemsByIdLogic) ItemsById(req *types.ItemsByIdReq) (resp *types.ItemsByIdResp, err error) {
	Recode, err := l.svcCtx.RecodeRpc.GetItemsById(l.ctx, &recode.GetItemsByIdReq{Id: int64(req.Id)})
	if err != nil {
		return nil, err
	}
	var t types.Items
	_ = copier.Copy(&t, &Recode.Items)

	return &types.ItemsByIdResp{Recode: t}, nil
}
