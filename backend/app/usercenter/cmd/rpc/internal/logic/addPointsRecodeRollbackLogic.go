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
	"looklook/app/usercenter/model"
	"looklook/common/xerr"

	"looklook/app/usercenter/cmd/rpc/internal/svc"
	"looklook/app/usercenter/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddPointsRecodeRollbackLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddPointsRecodeRollbackLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddPointsRecodeRollbackLogic {
	return &AddPointsRecodeRollbackLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *AddPointsRecodeRollbackLogic) AddPointsRecodeRollback(in *pb.AddPointsRecodeReq) (*pb.AddPointsRecodeResp, error) {

	pointRecode, err := l.svcCtx.PointsRecodeModel.FindLastOneByUserIdTaskId(l.ctx, in.UserId, in.TaskId)
	if err != nil && err != model.ErrNotFound {
		return nil, status.Error(codes.Internal, fmt.Sprintf("AddPointsRecodeRollback Database Exception pointsRecode dtm: err: %v", err))
	}

	if pointRecode != nil {
		barrier, err := dtmgrpc.BarrierFromGrpc(l.ctx)
		db, err := sqlx.NewMysql(l.svcCtx.Config.DB.DataSource).RawDB()
		if err != nil {
			return nil, status.Error(codes.Internal, fmt.Sprintf("AddPointsRecodeRollback Database Exception pointsRecode dtm: err: %v", err))
		}

		if err = barrier.CallWithDB(db, func(tx *sql.Tx) error {
			if err := l.svcCtx.UserModel.Trans(l.ctx, func(context context.Context, session sqlx.Session) error {

				err := l.svcCtx.PointsRecodeModel.Delete(l.ctx, pointRecode.Id)
				if err != nil {
					return errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "pointsRecode Database Exception Delete : %+v , err: %v", pointRecode, err)
				}

				// 修改用户表中积分值
				err = l.svcCtx.UserModel.TransUpdatePoints(l.ctx, session, in.Points, in.UserId)
				if err != nil {
					return errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "user Database Exception pointsRecode : %+v , err: %v", in.UserId, err)
				}

				return nil
			}); err != nil {
				return errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "pointsRecode Database Exception pointsRecode : err: %v", err)
			}
			return nil
		}); err != nil {
			return nil, status.Error(codes.Internal, err.Error())
		}
	}

	return &pb.AddPointsRecodeResp{}, nil
}
