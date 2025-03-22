package logic

import (
	"context"
	"github.com/pkg/errors"
	"looklook/common/xerr"

	"looklook/app/usercenter/cmd/rpc/internal/svc"
	"looklook/app/usercenter/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type ModifyUserInfoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewModifyUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ModifyUserInfoLogic {
	return &ModifyUserInfoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ModifyUserInfoLogic) ModifyUserInfo(in *pb.ModifyUserInfoReq) (*pb.ModifyUserInfoResp, error) {
	user, err := l.svcCtx.UserModel.FindUserById(l.ctx, in.UserId)
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "ModifyUserInfo find user db err , id:%d , err:%v", in.UserId, err)
	}
	user.Nickname = in.Nickname
	user.Avatar = in.Avatar
	err = l.svcCtx.UserModel.Update(l.ctx, user)
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "ModifyUserInfo find user db err , id:%d , err:%v", in.UserId, err)
	}

	return &pb.ModifyUserInfoResp{}, nil
}
