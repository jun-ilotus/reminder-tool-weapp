package logic

import (
	"context"
	"github.com/jinzhu/copier"
	"github.com/pkg/errors"
	"looklook/common/xerr"

	"looklook/app/lottery/cmd/rpc/internal/svc"
	"looklook/app/lottery/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type SearchPrizeLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSearchPrizeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SearchPrizeLogic {
	return &SearchPrizeLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SearchPrizeLogic) SearchPrize(in *pb.SearchPrizeReq) (*pb.SearchPrizeResp, error) {
	prizes, err := l.svcCtx.PrizeModel.PrizeListByLotteryId(l.ctx, in.LotteryId)
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "SearchPrize Database Exception lottery err: %v", err)
	}
	var resp []*pb.Prize
	for _, prize := range prizes {
		pbPrize := &pb.Prize{}
		_ = copier.Copy(pbPrize, prize)
		resp = append(resp, pbPrize)
	}

	return &pb.SearchPrizeResp{Prize: resp}, nil
}
