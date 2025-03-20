package items

import (
	"context"
	"github.com/jinzhu/copier"
	"looklook/app/recode/cmd/rpc/recode"
	"looklook/common/ctxdata"

	"looklook/app/recode/cmd/api/internal/svc"
	"looklook/app/recode/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ItemsListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewItemsListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ItemsListLogic {
	return &ItemsListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ItemsListLogic) ItemsList(req *types.ItemsListReq) (resp *types.ItemsListResp, err error) {
	userId := ctxdata.GetUidFromCtx(l.ctx)
	searchItems, err := l.svcCtx.RecodeRpc.SearchItems(l.ctx, &recode.SearchItemsReq{
		UserId: userId,
	})
	if err != nil {
		return nil, err
	}

	var ItemsList []types.Items
	if len(searchItems.Items) > 0 {
		for _, item := range searchItems.Items {
			var t types.Items
			_ = copier.Copy(&t, &item)
			ItemsList = append(ItemsList, t)
		}
	}

	return &types.ItemsListResp{List: ItemsList}, nil
}
