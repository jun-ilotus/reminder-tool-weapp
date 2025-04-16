package logic

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/dtm-labs/client/dtmgrpc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"looklook/app/signin/cmd/rpc/internal/svc"
	"looklook/app/signin/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddTaskFinishRollbackLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddTaskFinishRollbackLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddTaskFinishRollbackLogic {
	return &AddTaskFinishRollbackLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *AddTaskFinishRollbackLogic) AddTaskFinishRollback(in *pb.AddTaskFinishReq) (*pb.AddTaskFinishResp, error) {
	barrier, err := dtmgrpc.BarrierFromGrpc(l.ctx)
	db, err := sqlx.NewMysql(l.svcCtx.Config.DB.DataSource).RawDB()
	if err != nil {
		return nil, status.Error(codes.Internal, fmt.Sprintf("AddPointsRecodeRollback Database Exception pointsRecode dtm: err: %v", err))
	}

	if err = barrier.CallWithDB(db, func(tx *sql.Tx) error {
		for _, task := range in.Task {
			_ = l.svcCtx.TaskFinishModel.DeleteByUserIdTaskId(l.ctx, in.UserId, task.Id)
		}
		return nil
	}); err != nil {
		return nil, status.Error(codes.Internal, fmt.Sprintf("AddPointsRecodeRollback Database Exception pointsRecode dtm: err: %v", err))
	}

	return &pb.AddTaskFinishResp{}, nil
}
