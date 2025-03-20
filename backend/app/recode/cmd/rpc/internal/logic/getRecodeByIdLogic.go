package logic

import (
	"context"
	"github.com/jinzhu/copier"
	"github.com/pkg/errors"
	"looklook/common/xerr"

	"looklook/app/recode/cmd/rpc/internal/svc"
	"looklook/app/recode/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetRecodeByIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetRecodeByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetRecodeByIdLogic {
	return &GetRecodeByIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetRecodeByIdLogic) GetRecodeById(in *pb.GetRecodeByIdReq) (*pb.GetRecodeByIdResp, error) {
	recode, err := l.svcCtx.RecodeModel.FindOne(l.ctx, in.Id)
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "items_id:%d,err:%v", in.Id, err)
	}
	pbItems := &pb.Recode{}
	_ = copier.Copy(pbItems, recode)
	pbItems.RecodeTime = recode.RecodeTime.Unix()
	return &pb.GetRecodeByIdResp{Recode: pbItems}, nil
}
