package recode

import (
	"context"
	"looklook/app/recode/cmd/rpc/recode"
	"looklook/common/ctxdata"

	"looklook/app/recode/cmd/api/internal/svc"
	"looklook/app/recode/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ModifyRecodeLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewModifyRecodeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ModifyRecodeLogic {
	return &ModifyRecodeLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ModifyRecodeLogic) ModifyRecode(req *types.ModifyRecodeReq) (resp *types.ModifyRecodeResp, err error) {
	userId := ctxdata.GetUidFromCtx(l.ctx)

	Recoder, err := l.svcCtx.RecodeRpc.GetRecodeById(l.ctx, &recode.GetRecodeByIdReq{
		Id: req.Id,
	})
	if err != nil {
		return nil, err
	}

	_, err = l.svcCtx.RecodeRpc.UpdateRecode(l.ctx, &recode.UpdateRecodeReq{
		UserId:     userId,
		Id:         req.Id,
		Content:    req.Content,
		ItemsId:    Recoder.Recode.ItemsId,
		RecodeTime: Recoder.Recode.RecodeTime,
	})
	if err != nil {
		return nil, err
	}

	return &types.ModifyRecodeResp{Id: int(req.Id)}, nil
}
