package logic

import (
	"context"
	"github.com/jinzhu/copier"
	"github.com/pkg/errors"
	"looklook/app/recode/model"
	"looklook/app/usercenter/cmd/rpc/usercenter"
	"looklook/common/xerr"

	"looklook/app/recode/cmd/rpc/internal/svc"
	"looklook/app/recode/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type SearchItemsLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSearchItemsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SearchItemsLogic {
	return &SearchItemsLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SearchItemsLogic) SearchItems(in *pb.SearchItemsReq) (*pb.SearchItemsResp, error) {
	userInfo, err := l.svcCtx.UsercenterRpc.GetUserInfo(l.ctx, &usercenter.GetUserInfoReq{Id: in.UserId})
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "查询事项失败 err: %v", err)
	}

	list, err := l.svcCtx.ItemsModel.ItemsList(l.ctx, in.UserId)
	if err != nil && !errors.Is(err, model.ErrNotFound) {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "查询事项失败 err: %v", err)
	}

	if userInfo.User.IntimateId != 0 {
		list1, err := l.svcCtx.ItemsModel.ItemsIntimateList(l.ctx, userInfo.User.IntimateId)
		if err != nil {
			return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "查询事项失败 err: %v", err)
		}
		list = append(list, list1...)
	}

	var resp []*pb.Items
	if len(list) > 0 {
		for _, reminder := range list {
			var pbReminder pb.Items
			_ = copier.Copy(&pbReminder, &reminder)
			resp = append(resp, &pbReminder)
		}
	}

	return &pb.SearchItemsResp{
		Items: resp,
	}, nil
}
