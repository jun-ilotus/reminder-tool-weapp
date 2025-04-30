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

type GetLotteryByIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetLotteryByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetLotteryByIdLogic {
	return &GetLotteryByIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetLotteryByIdLogic) GetLotteryById(in *pb.GetLotteryByIdReq) (*pb.GetLotteryByIdResp, error) {
	lottery, err := l.svcCtx.LotteryModel.GetLotteryById(l.ctx, in.Id)
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "GetLotteryById Database Exception lottery err: %v", err)
	}
	pbLottery := &pb.Lottery{}
	_ = copier.Copy(pbLottery, lottery)
	pbLottery.AnnounceTime = lottery.AnnounceTime.Unix()
	pbLottery.AwardDeadline = lottery.AwardDeadline.Unix()

	return &pb.GetLotteryByIdResp{Lottery: pbLottery}, nil
}
