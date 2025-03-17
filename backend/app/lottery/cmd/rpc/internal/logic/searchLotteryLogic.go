package logic

import (
	"context"
	"github.com/jinzhu/copier"
	"github.com/pkg/errors"
	"looklook/app/order/model"
	"looklook/common/xerr"

	"looklook/app/lottery/cmd/rpc/internal/svc"
	"looklook/app/lottery/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type SearchLotteryLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSearchLotteryLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SearchLotteryLogic {
	return &SearchLotteryLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SearchLotteryLogic) SearchLottery(in *pb.SearchLotteryReq) (*pb.SearchLotteryResp, error) {
	//list, err := l.svcCtx.LotteryModel.FindPageListByIdDESC(l.ctx, whereBuilder, in.LastId, in.PageSize)
	list, err := l.svcCtx.LotteryModel.List(l.ctx, in.Page, in.Limit)
	if err != nil && err != model.ErrNotFound {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "Failed to get user's homestay order err : %v , in :%+v", err, in)
	}

	var resp []*pb.Lottery
	if len(list) > 0 {
		for _, lottery := range list {
			var pbLottery pb.Lottery
			_ = copier.Copy(&pbLottery, lottery)
			resp = append(resp, &pbLottery)
		}
	}

	return &pb.SearchLotteryResp{
		Lottery: resp,
	}, nil
}
