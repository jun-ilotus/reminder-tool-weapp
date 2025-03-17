package lottery

import (
	"context"
	"github.com/jinzhu/copier"
	"github.com/pkg/errors"
	"looklook/app/lottery/cmd/api/internal/svc"
	"looklook/app/lottery/cmd/api/internal/types"
	"looklook/app/lottery/cmd/rpc/lottery"
	"looklook/app/lottery/cmd/rpc/pb"
	"looklook/common/ctxdata"
	"looklook/common/xerr"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateLotteryLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreateLotteryLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateLotteryLogic {
	return &CreateLotteryLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateLotteryLogic) CreateLottery(req *types.CreateLotteryReq) (resp *types.CreateLotteryResp, err error) {
	userId := ctxdata.GetUidFromCtx(l.ctx)
	//循环赋值 TODO 寻找更好的方案
	var pbPrizes []*pb.Prize
	for _, reqPrize := range req.Prizes {
		pbPrize := new(pb.Prize)
		err := copier.Copy(&pbPrize, reqPrize)
		if err != nil {
			return nil, err
		}
		pbPrizes = append(pbPrizes, pbPrize)
	}
	addLottery, err := l.svcCtx.LotteryRpc.AddLottery(l.ctx, &lottery.AddLotteryReq{
		UserId:        userId,
		Name:          req.Name,
		Thumb:         req.Thumb,
		PublishType:   req.PublishType,
		PublishTime:   req.PublishTime,
		JoinNumber:    req.JoinNumber,
		Introduce:     req.Introduce,
		AwardDeadline: req.AwardDeadline,
		Prizes:        pbPrizes,
	})
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrMsg("create lottery fail"), "create lottery rpc CreateLottery fail req: %+v , err : %v ", req, err)
	}
	return &types.CreateLotteryResp{
		Id: int(addLottery.Id),
	}, nil
}
