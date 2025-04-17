package recode

import (
	"context"
	"github.com/jinzhu/copier"
	"looklook/app/signin/cmd/rpc/signin"
	"looklook/common/ctxdata"
	"time"

	"looklook/app/signin/cmd/api/internal/svc"
	"looklook/app/signin/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ListRecodeLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewListRecodeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListRecodeLogic {
	return &ListRecodeLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ListRecodeLogic) ListRecode(req *types.ListRecodeReq) (resp *types.ListRecodeResp, err error) {
	userId := ctxdata.GetUidFromCtx(l.ctx)

	res, err := l.svcCtx.SigninRpc.SearchRecode(l.ctx, &signin.SearchRecodeReq{
		UserId: userId,
	})
	if err != nil {
		return nil, err
	}

	var RecodeList []types.Recode
	var isToday int64 = 0
	if len(res.Recode) > 0 {
		for _, item := range res.Recode {
			var t types.Recode
			_ = copier.Copy(&t, item)
			RecodeList = append(RecodeList, t)
		}
		nowTime := time.Now()
		nowTimeStamp := time.Date(nowTime.Year(), nowTime.Month(), nowTime.Day(), -8, 0, 0, 0, time.UTC).Unix()
		if RecodeList[0].SignDate == nowTimeStamp {
			isToday = 1
		}
	}

	return &types.ListRecodeResp{
		List:    RecodeList,
		Days:    res.Days,
		IsToday: isToday,
	}, nil
}
