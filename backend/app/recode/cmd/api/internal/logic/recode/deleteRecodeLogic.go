package recode

import (
	"context"
	"looklook/app/recode/cmd/rpc/recode"
	"looklook/common/ctxdata"

	"looklook/app/recode/cmd/api/internal/svc"
	"looklook/app/recode/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteRecodeLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteRecodeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteRecodeLogic {
	return &DeleteRecodeLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteRecodeLogic) DeleteRecode(req *types.DeleteRecodeReq) (resp *types.DeleteRecodeResp, err error) {
	userId := ctxdata.GetUidFromCtx(l.ctx)

	_, err = l.svcCtx.RecodeRpc.DelRecode(l.ctx, &recode.DelRecodeReq{
		UserId: userId,
		Id:     int64(req.Id),
	})
	if err != nil {
		return nil, err
	}

	return &types.DeleteRecodeResp{Id: req.Id}, nil
}
