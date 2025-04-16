package logic

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/dtm-labs/client/dtmgrpc"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"looklook/app/signin/model"
	"looklook/common/xerr"
	"time"

	"looklook/app/signin/cmd/rpc/internal/svc"
	"looklook/app/signin/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddRecodeRollbackLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddRecodeRollbackLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddRecodeRollbackLogic {
	return &AddRecodeRollbackLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *AddRecodeRollbackLogic) AddRecodeRollback(in *pb.AddRecodeReq) (*pb.AddRecodeResp, error) {
	recode, err := l.svcCtx.RecodeModel.FindLastOneByUserIdSignDate(l.ctx, in.UserId, time.Unix(in.SignDate, 0))
	if err != nil && err != model.ErrNotFound {
		return nil, status.Error(codes.Internal, fmt.Sprintf("AddPointsRecodeRollback Database Exception pointsRecode dtm: err: %v", err))
	}
	if recode != nil {
		barrier, err := dtmgrpc.BarrierFromGrpc(l.ctx)
		db, err := sqlx.NewMysql(l.svcCtx.Config.DB.DataSource).RawDB()
		if err != nil {
			return nil, status.Error(codes.Internal, fmt.Sprintf("AddPointsRecodeRollback Database Exception pointsRecode dtm: err: %v", err))
		}

		if err = barrier.CallWithDB(db, func(tx *sql.Tx) error {
			err := l.svcCtx.RecodeModel.Delete(l.ctx, recode.Id)
			if err != nil {
				return errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "RecodeModel Database Exception Delete : %+v , err: %v", recode, err)
			}
			return nil
		}); err != nil {
			return nil, status.Error(codes.Internal, fmt.Sprintf("AddPointsRecodeRollback Database Exception pointsRecode dtm: err: %v", err))
		}
	}

	return &pb.AddRecodeResp{}, nil
}
