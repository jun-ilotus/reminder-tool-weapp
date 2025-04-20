package remind

import (
	"context"
	"looklook/app/signin/cmd/rpc/pb"
	"looklook/common/ctxdata"

	"looklook/app/signin/cmd/api/internal/svc"
	"looklook/app/signin/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetRemindStatusLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetRemindStatusLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetRemindStatusLogic {
	return &GetRemindStatusLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetRemindStatusLogic) GetRemindStatus(req *types.GetRemindStatusReq) (resp *types.GetRemindStatusResp, err error) {
	userId := ctxdata.GetUidFromCtx(l.ctx)
	status, err := l.svcCtx.SigninRpc.GetRemindStatus(l.ctx, &pb.GetRemindStatusReq{
		UserId: userId,
	})
	if err != nil {
		return nil, err
	}

	return &types.GetRemindStatusResp{Status: status.Status}, nil
}
