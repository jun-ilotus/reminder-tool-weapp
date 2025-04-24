package logic

import (
	"context"
	"github.com/jinzhu/copier"
	"github.com/pkg/errors"
	usercenterModel "looklook/app/usercenter/model"
	"looklook/common/xerr"

	"looklook/app/signin/cmd/rpc/internal/svc"
	"looklook/app/signin/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
	"looklook/app/usercenter/cmd/rpc/usercenter"
)

type SearchRemindLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSearchRemindLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SearchRemindLogic {
	return &SearchRemindLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SearchRemindLogic) SearchRemind(in *pb.SearchRemindReq) (*pb.SearchRemindResp, error) {
	list, err := l.svcCtx.RemindModel.FindList(l.ctx)
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "FindList 查询记录失败 err: %v", err)
	}

	var resp []*pb.Remind
	for _, v := range list {
		var t pb.Remind
		_ = copier.Copy(&t, v)

		// todo 根据每个 id 去查 openid 求更优雅的方法
		userAuth, err := l.svcCtx.UsercenterRpc.GetUserAuthByUserId(l.ctx, &usercenter.GetUserAuthByUserIdReq{
			UserId:   v.UserId,
			AuthType: usercenterModel.UserAuthTypeSmallWX,
		})
		if err != nil {
			return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "FindList 查询记录失败 err: %v", err)
		}
		t.AuthKey = userAuth.UserAuth.AuthKey

		resp = append(resp, &t)
	}

	return &pb.SearchRemindResp{Reminds: resp}, nil
}
