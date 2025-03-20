package items

import (
	"context"
	"looklook/app/recode/cmd/rpc/recode"
	"looklook/common/ctxdata"

	"looklook/app/recode/cmd/api/internal/svc"
	"looklook/app/recode/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteItemsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteItemsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteItemsLogic {
	return &DeleteItemsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteItemsLogic) DeleteItems(req *types.DeleteItemsReq) (resp *types.DeleteItemsResp, err error) {
	userId := ctxdata.GetUidFromCtx(l.ctx)

	_, err = l.svcCtx.RecodeRpc.DelItems(l.ctx, &recode.DelItemsReq{
		UserId: userId,
		Id:     int64(req.Id),
	})
	if err != nil {
		return nil, err
	}

	return &types.DeleteItemsResp{Id: req.Id}, nil
}
