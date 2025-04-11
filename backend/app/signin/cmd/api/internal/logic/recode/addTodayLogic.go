package recode

import (
	"context"

	"looklook/app/signin/cmd/api/internal/svc"
	"looklook/app/signin/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddTodayLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAddTodayLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddTodayLogic {
	return &AddTodayLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AddTodayLogic) AddToday(req *types.AddRecodeTodayReq) (resp *types.AddRecodeTodayResp, err error) {

	return
}
