package logic

import (
	"context"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"looklook/app/recode/model"
	"looklook/common/xerr"
	"time"

	"looklook/app/recode/cmd/rpc/internal/svc"
	"looklook/app/recode/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddRecodeLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddRecodeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddRecodeLogic {
	return &AddRecodeLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// -----------------------recode-----------------------
func (l *AddRecodeLogic) AddRecode(in *pb.AddRecodeReq) (*pb.AddRecodeResp, error) {
	var id int64
	err := l.svcCtx.RecodeModel.Trans(l.ctx, func(context context.Context, session sqlx.Session) error {
		//  logx.Error("in:", in)

		recode := &model.Recode{}
		recode.UserId = in.UserId
		recode.Content = in.Content
		recode.RecodeTime = time.Unix(in.RecodeTime, 0)
		recode.ItemsId = in.ItemsId

		insert, err := l.svcCtx.RecodeModel.TransInsert(l.ctx, session, recode)

		if err != nil {
			return errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "Reminder Database Exception recode : %+v , err: %v", recode, err)
		}
		id, _ = insert.LastInsertId()
		return nil
	})
	if err != nil {
		return nil, err
	}

	return &pb.AddRecodeResp{
		Id: id,
	}, nil
}
