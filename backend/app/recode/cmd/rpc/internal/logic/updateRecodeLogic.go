package logic

import (
	"context"
	"github.com/jinzhu/copier"
	"github.com/pkg/errors"
	"looklook/app/recode/cmd/rpc/internal/svc"
	"looklook/app/recode/cmd/rpc/pb"
	"looklook/app/recode/model"
	"looklook/common/xerr"
	"time"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateRecodeLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateRecodeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateRecodeLogic {
	return &UpdateRecodeLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateRecodeLogic) UpdateRecode(in *pb.UpdateRecodeReq) (*pb.UpdateRecodeResp, error) {
	one, err := l.svcCtx.RecodeModel.FindOne(l.ctx, in.Id)
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "查询记录id失败 err: %v", err)
	}
	if one.UserId != in.UserId { // todo共享也可修改
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "该用户不是事项创建者")
	}
	recode := &model.Recode{}
	_ = copier.Copy(recode, in)
	recode.RecodeTime = time.Unix(in.RecodeTime, 0)
	err = l.svcCtx.RecodeModel.Update(l.ctx, recode)
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "更新失败 err: %v", err)
	}
	return &pb.UpdateRecodeResp{}, nil
}
