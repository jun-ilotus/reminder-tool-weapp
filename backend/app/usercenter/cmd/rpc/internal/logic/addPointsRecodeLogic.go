package logic

import (
	"context"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"looklook/app/usercenter/cmd/rpc/internal/svc"
	"looklook/app/usercenter/cmd/rpc/pb"
	"looklook/app/usercenter/model"
	"looklook/common/xerr"
)

type AddPointsRecodeLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddPointsRecodeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddPointsRecodeLogic {
	return &AddPointsRecodeLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// AddPointsRecode 对接dtm
func (l *AddPointsRecodeLogic) AddPointsRecode(in *pb.AddPointsRecodeReq) (*pb.AddPointsRecodeResp, error) {

	id := int64(0)
	//添加事务处理
	if err := l.svcCtx.UserModel.Trans(l.ctx, func(context context.Context, session sqlx.Session) error {
		pointsRecode := &model.PointsRecode{
			UserId:  in.UserId,
			Points:  in.Points,
			Content: in.Content,
		}
		// 插入积分记录表
		insert, err := l.svcCtx.PointsRecodeModel.TransInsert(l.ctx, session, pointsRecode)
		if err != nil {
			return errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "pointsRecode Database Exception pointsRecode : %+v , err: %v", pointsRecode, err)
		}
		id, _ = insert.LastInsertId()

		// 修改用户表中积分值
		err = l.svcCtx.UserModel.TransUpdatePoints(l.ctx, session, pointsRecode.Points, pointsRecode.UserId)
		if err != nil {
			return errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "user Database Exception pointsRecode : %+v , err: %v", pointsRecode.UserId, err)
		}

		return nil
	}); err != nil {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "pointsRecode Database Exception pointsRecode : err: %v", err)
	}

	return &pb.AddPointsRecodeResp{Id: id}, nil
}

//func (l *AddPointsRecodeLogic) AddPointsRecode(in *pb.AddPointsRecodeReq) (*pb.AddPointsRecodeResp, error) {
//	id := int64(0)
//	// 添加事务处理
//	err := l.svcCtx.UserModel.Trans(l.ctx, func(context context.Context, session sqlx.Session) error {
//		pointsRecode := &model.PointsRecode{
//			UserId:  in.UserId,
//			Points:  in.Points,
//			Content: in.Content,
//		}
//		// 插入积分记录表
//		insert, err := l.svcCtx.PointsRecodeModel.TransInsert(l.ctx, session, pointsRecode)
//		if err != nil {
//			return errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "pointsRecode Database Exception pointsRecode : %+v , err: %v", pointsRecode, err)
//		}
//		id, _ = insert.LastInsertId()
//
//		// 修改用户表中积分值
//		err = l.svcCtx.UserModel.TransUpdatePoints(l.ctx, session, pointsRecode.Points, pointsRecode.UserId)
//		if err != nil {
//			return errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "user Database Exception pointsRecode : %+v , err: %v", pointsRecode.UserId, err)
//		}
//
//		return nil
//	})
//	if err != nil {
//		return nil, err
//	}
//
//	return &pb.AddPointsRecodeResp{Id: id}, nil
//}
