package remind

import (
	"context"
	"looklook/app/signin/cmd/rpc/pb"
	"looklook/common/ctxdata"

	"looklook/app/signin/cmd/api/internal/svc"
	"looklook/app/signin/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ChangeSignRemindLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewChangeSignRemindLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ChangeSignRemindLogic {
	return &ChangeSignRemindLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ChangeSignRemindLogic) ChangeSignRemind(req *types.ChangeSignRemindReq) (resp *types.ChangeSignRemindResp, err error) {
	userId := ctxdata.GetUidFromCtx(l.ctx)
	remind, err := l.svcCtx.SigninRpc.ChangeSignRemind(l.ctx, &pb.ChangeRemindReq{
		UserId: userId,
		Status: req.Status,
	})
	if err != nil {
		return nil, err
	}

	return &types.ChangeSignRemindResp{
		Id:     remind.Id,
		Status: remind.Status,
	}, nil
}
