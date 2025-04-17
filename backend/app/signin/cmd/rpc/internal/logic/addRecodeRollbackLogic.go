package logic

import (
	"context"
	"fmt"
	"github.com/pkg/errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"looklook/app/signin/cmd/rpc/internal/svc"
	"looklook/app/signin/cmd/rpc/pb"
	"looklook/app/signin/model"
	"looklook/common/xerr"

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
	recode, err := l.svcCtx.RecodeModel.FindLastOneByUserIdSignDate(l.ctx, in.UserId)
	if err != nil && err != model.ErrNotFound {
		return nil, status.Error(codes.Internal, fmt.Sprintf("AddPointsRecodeRollback Database Exception pointsRecode dtm: err: %v", err))
	}
	if recode != nil {
		err := l.svcCtx.RecodeModel.Delete(l.ctx, recode.Id)
		if err != nil {
			return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "RecodeModel Database Exception Delete : %+v , err: %v", recode, err)
		}
	}

	return &pb.AddRecodeResp{}, nil
}
