package task

import (
	"context"
	"github.com/jinzhu/copier"
	"looklook/app/signin/cmd/rpc/signin"
	"looklook/common/ctxdata"

	"looklook/app/signin/cmd/api/internal/svc"
	"looklook/app/signin/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ListTaskLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewListTaskLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListTaskLogic {
	return &ListTaskLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ListTaskLogic) ListTask(req *types.ListTaskReq) (resp *types.ListTaskResp, err error) {
	userId := ctxdata.GetUidFromCtx(l.ctx)

	res, err := l.svcCtx.SigninRpc.SearchTask(l.ctx, &signin.SearchTaskReq{
		UserId: userId,
	})
	if err != nil {
		return nil, err
	}

	var RecodeList []types.Task
	if len(res.Task) > 0 {
		for _, item := range res.Task {
			var t types.Task
			_ = copier.Copy(&t, item)
			RecodeList = append(RecodeList, t)
		}
	}
	return &types.ListTaskResp{List: RecodeList}, nil
}
