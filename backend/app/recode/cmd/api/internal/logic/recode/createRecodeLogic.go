package recode

import (
	"context"
	"looklook/app/recode/cmd/rpc/recode"
	"looklook/common/ctxdata"

	"looklook/app/recode/cmd/api/internal/svc"
	"looklook/app/recode/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateRecodeLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreateRecodeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateRecodeLogic {
	return &CreateRecodeLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateRecodeLogic) CreateRecode(req *types.CreateRecodeReq) (resp *types.CreateRecodeResp, err error) {
	userId := ctxdata.GetUidFromCtx(l.ctx)
	addRecode, err := l.svcCtx.RecodeRpc.AddRecode(l.ctx, &recode.AddRecodeReq{
		UserId:     userId,
		Content:    req.Content,
		ItemsId:    req.ItemsId,
		RecodeTime: req.RecodeTime,
	})
	if err != nil {
		return nil, err
	}
	return &types.CreateRecodeResp{
		Id: int(addRecode.Id),
	}, nil
}
