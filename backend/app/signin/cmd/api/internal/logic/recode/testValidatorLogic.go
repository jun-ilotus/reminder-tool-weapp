package recode

import (
	"context"

	"looklook/app/signin/cmd/api/internal/svc"
	"looklook/app/signin/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type TestValidatorLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewTestValidatorLogic(ctx context.Context, svcCtx *svc.ServiceContext) *TestValidatorLogic {
	return &TestValidatorLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *TestValidatorLogic) TestValidator(req *types.TestReq) (resp *types.TestResp, err error) {
	// todo: add your logic here and delete this line

	return
}
