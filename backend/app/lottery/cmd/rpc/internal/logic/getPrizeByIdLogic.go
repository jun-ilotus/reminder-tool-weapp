package logic

import (
	"context"
	"github.com/jinzhu/copier"
	"github.com/pkg/errors"
	"looklook/common/xerr"

	"looklook/app/lottery/cmd/rpc/internal/svc"
	"looklook/app/lottery/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetPrizeByIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetPrizeByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetPrizeByIdLogic {
	return &GetPrizeByIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetPrizeByIdLogic) GetPrizeById(in *pb.GetPrizeByIdReq) (*pb.GetPrizeByIdResp, error) {
	one, err := l.svcCtx.PrizeModel.FindOne(l.ctx, in.Id)
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "GetPrizeById Database Exception lottery err: %v", err)
	}
	prize := pb.Prize{}
	_ = copier.Copy(&prize, one)

	return &pb.GetPrizeByIdResp{Prize: &prize}, nil
}
