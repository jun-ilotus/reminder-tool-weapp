package logic

import (
	"context"
	"github.com/jinzhu/copier"
	"github.com/pkg/errors"
	"looklook/common/xerr"

	"looklook/app/usercenter/cmd/rpc/internal/svc"
	"looklook/app/usercenter/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type PointsRecodeListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewPointsRecodeListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PointsRecodeListLogic {
	return &PointsRecodeListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *PointsRecodeListLogic) PointsRecodeList(in *pb.PointsRecodeListReq) (*pb.PointRecodeListResp, error) {

	list, err := l.svcCtx.PointsRecodeModel.FindUserList(l.ctx, in.UserId)
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "user Database Exception PointsRecodeList : %+v , err: %v", in.UserId, err)
	}

	var resp []*pb.PointsRecode
	for _, v := range list {
		var pbPointsRecode pb.PointsRecode
		_ = copier.Copy(&pbPointsRecode, &v)
		resp = append(resp, &pbPointsRecode)
	}

	return &pb.PointRecodeListResp{PointsRecode: resp}, nil
}
