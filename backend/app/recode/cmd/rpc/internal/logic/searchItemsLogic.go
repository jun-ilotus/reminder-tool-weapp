package logic

import (
	"context"
	"errors"
	"github.com/jinzhu/copier"
	"looklook/app/recode/model"

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
	list, err := l.svcCtx.ItemsModel.ItemsList(l.ctx, in.UserId)
	if err != nil && !errors.Is(err, model.ErrNotFound) {
		return nil, err
	}
	// todo 查询出共享的事项
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
