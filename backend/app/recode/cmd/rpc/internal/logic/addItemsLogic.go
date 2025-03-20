package logic

import (
	"context"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"looklook/app/recode/cmd/rpc/internal/svc"
	"looklook/app/recode/cmd/rpc/pb"
	"looklook/app/recode/model"
	"looklook/common/xerr"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddItemsLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddItemsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddItemsLogic {
	return &AddItemsLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// -----------------------recode-----------------------
func (l *AddItemsLogic) AddItems(in *pb.AddItemsReq) (*pb.AddItemsResp, error) {
	var id int64
	err := l.svcCtx.ItemsModel.Trans(l.ctx, func(context context.Context, session sqlx.Session) error {
		//  logx.Error("in:", in)

		items := &model.Items{}
		items.UserId = in.UserId
		items.Content = in.Content
		items.Member = in.Member

		insert, err := l.svcCtx.ItemsModel.TransInsert(l.ctx, session, items)

		// todo  if member == 1 2 给亲友加
		if err != nil {
			return errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "Reminder Database Exception recode : %+v , err: %v", items, err)
		}
		id, _ = insert.LastInsertId()
		return nil
	})
	if err != nil {
		return nil, err
	}

	return &pb.AddItemsResp{
		Id: id,
	}, nil
}
