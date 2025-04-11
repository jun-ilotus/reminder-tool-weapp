package logic

import (
	"context"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"looklook/app/signin/model"
	"looklook/common/xerr"
	"time"

	"looklook/app/signin/cmd/rpc/internal/svc"
	"looklook/app/signin/cmd/rpc/pb"

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
	id := int64(0)
	err := l.svcCtx.RecodeModel.Trans(l.ctx, func(ctx context.Context, session sqlx.Session) error {
		recode := &model.Recode{
			UserId:   in.UserId,
			SignDate: time.Unix(in.SignDate, 0),
		}
		insert, err := l.svcCtx.RecodeModel.TransInsert(l.ctx, session, recode)
		if err != nil {
			return errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "signin recode Database Exception pointsRecode : %+v , err: %v", recode, err)
		}
		id, _ = insert.LastInsertId()

		// todo 添加任务查询 看是否完成

		return nil
	})
	if err != nil {
		return nil, err
	}
	return &pb.AddRecodeResp{Id: id}, nil
}
