package pointsRecode

import (
	"context"
	"github.com/jinzhu/copier"
	"looklook/app/usercenter/cmd/rpc/pb"
	"looklook/common/ctxdata"

	"looklook/app/usercenter/cmd/api/internal/svc"
	"looklook/app/usercenter/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListLogic {
	return &ListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ListLogic) List(req *types.PointsRecodeListReq) (resp *types.PointsRecodeListResp, err error) {
	userId := ctxdata.GetUidFromCtx(l.ctx)
	list, err := l.svcCtx.UsercenterRpc.PointsRecodeList(l.ctx, &pb.PointsRecodeListReq{UserId: userId})
	if err != nil {
		return nil, err
	}
	var respList []types.PointsRecode
	for _, v := range list.PointsRecode {
		var pointsRecode types.PointsRecode
		_ = copier.Copy(&pointsRecode, v)
		respList = append(respList, pointsRecode)
	}

	return &types.PointsRecodeListResp{List: respList}, nil
}
