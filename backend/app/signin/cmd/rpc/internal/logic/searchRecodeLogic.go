package logic

import (
	"context"
	"github.com/jinzhu/copier"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"looklook/app/signin/cmd/rpc/internal/svc"
	"looklook/app/signin/cmd/rpc/pb"
	"looklook/app/signin/model"
	"looklook/common/xerr"

	"github.com/zeromicro/go-zero/core/logx"
)

type SearchRecodeLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSearchRecodeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SearchRecodeLogic {
	return &SearchRecodeLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SearchRecodeLogic) SearchRecode(in *pb.SearchRecodeReq) (*pb.SearchRecodeResp, error) {
	list, err := l.svcCtx.RecodeModel.RecodeList(l.ctx, in.UserId)
	if err != nil && !errors.Is(err, model.ErrNotFound) {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "SearchRecode查询记录失败 err: %v", err)
	}
	var days *int

	// 获取用户当前的连续签到天数
	err = l.svcCtx.RecodeModel.Trans(l.ctx, func(ctx context.Context, session sqlx.Session) error {
		days, err = l.svcCtx.RecodeModel.FindCountUserNowConRecodeDays(l.ctx, session, in.UserId)
		if err != nil {
			return errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "SearchRecode recode Database Exception : err: %v", err)
		}
		return nil
	})
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "SearchRecode recode Database Exception : err: %v", err)
	}

	var resp []*pb.Recode
	if len(list) > 0 {
		for _, recode := range list {
			var pbRecode pb.Recode
			_ = copier.Copy(&pbRecode, recode)
			pbRecode.SignDate = recode.SignDate.Unix()
			resp = append(resp, &pbRecode)
		}
	}

	return &pb.SearchRecodeResp{
		Recode: resp,
		Days:   int64(*days),
	}, nil
}
