package items

import (
	"context"
	"looklook/app/recode/cmd/rpc/recode"
	"looklook/common/ctxdata"

	"looklook/app/recode/cmd/api/internal/svc"
	"looklook/app/recode/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateItemsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreateItemsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateItemsLogic {
	return &CreateItemsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateItemsLogic) CreateItems(req *types.CreateItemsReq) (resp *types.CreateItemsResp, err error) {
	userId := ctxdata.GetUidFromCtx(l.ctx)
	addRecode, err := l.svcCtx.RecodeRpc.AddItems(l.ctx, &recode.AddItemsReq{
		UserId:  userId,
		Content: req.Content,
		Member:  req.Member,
	})
	if err != nil {
		return nil, err
	}
	return &types.CreateItemsResp{
		Id: int(addRecode.Id),
	}, nil
}
