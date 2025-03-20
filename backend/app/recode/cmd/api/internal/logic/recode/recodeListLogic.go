package recode

import (
	"context"
	"github.com/jinzhu/copier"
	"looklook/app/recode/cmd/rpc/recode"
	"looklook/common/ctxdata"

	"looklook/app/recode/cmd/api/internal/svc"
	"looklook/app/recode/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type RecodeListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRecodeListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RecodeListLogic {
	return &RecodeListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RecodeListLogic) RecodeList(req *types.RecodeListReq) (resp *types.RecodeListResp, err error) {
	userId := ctxdata.GetUidFromCtx(l.ctx)
	searchItems, err := l.svcCtx.RecodeRpc.SearchRecode(l.ctx, &recode.SearchRecodeReq{
		UserId:  userId,
		ItemsId: req.ItemsId,
	})
	if err != nil {
		return nil, err
	}

	var RecodeList []types.Recode
	if len(searchItems.Recode) > 0 {
		for _, item := range searchItems.Recode {
			var t types.Recode
			_ = copier.Copy(&t, &item)
			RecodeList = append(RecodeList, t)
		}
	}

	return &types.RecodeListResp{
		List:            RecodeList,
		RecodeDayCount:  searchItems.RecodeDayCount,
		RecodeDaySpaced: searchItems.RecodeDaySpaced,
	}, nil
}
