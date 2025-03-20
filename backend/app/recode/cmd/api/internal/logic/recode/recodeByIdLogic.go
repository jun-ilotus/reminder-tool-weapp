package recode

import (
	"context"
	"github.com/jinzhu/copier"
	"looklook/app/recode/cmd/rpc/recode"

	"looklook/app/recode/cmd/api/internal/svc"
	"looklook/app/recode/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type RecodeByIdLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRecodeByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RecodeByIdLogic {
	return &RecodeByIdLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RecodeByIdLogic) RecodeById(req *types.RecodeByIdReq) (resp *types.RecodeByIdResp, err error) {
	Recode, err := l.svcCtx.RecodeRpc.GetRecodeById(l.ctx, &recode.GetRecodeByIdReq{Id: int64(req.Id)})
	if err != nil {
		return nil, err
	}
	var t types.Recode
	_ = copier.Copy(&t, &Recode.Recode)

	return &types.RecodeByIdResp{Recode: t}, nil
}
