package items

import (
	"context"
	"looklook/app/recode/cmd/rpc/recode"
	"looklook/common/ctxdata"

	"looklook/app/recode/cmd/api/internal/svc"
	"looklook/app/recode/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ModifyItemsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewModifyItemsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ModifyItemsLogic {
	return &ModifyItemsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ModifyItemsLogic) ModifyItems(req *types.ModifyItemsReq) (resp *types.ModifyItemsResp, err error) {
	userId := ctxdata.GetUidFromCtx(l.ctx)

	_, err = l.svcCtx.RecodeRpc.UpdateItems(l.ctx, &recode.UpdateItemsReq{
		UserId:  userId,
		Id:      req.Id,
		Content: req.Content,
		Member:  req.Member,
	})
	if err != nil {
		return nil, err
	}

	return &types.ModifyItemsResp{Id: int(req.Id)}, nil
}
