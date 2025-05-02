package logic

import (
	"context"
	"github.com/jinzhu/copier"
	"github.com/pkg/errors"
	"looklook/app/lottery/model"
	"looklook/common/xerr"

	"looklook/app/lottery/cmd/rpc/internal/svc"
	"looklook/app/lottery/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type SearchUserLotteryParticipationLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSearchUserLotteryParticipationLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SearchUserLotteryParticipationLogic {
	return &SearchUserLotteryParticipationLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SearchUserLotteryParticipationLogic) SearchUserLotteryParticipation(in *pb.SearchUserLotteryParticipationReq) (*pb.SearchUserLotteryParticipationResp, error) {

	list, err := l.svcCtx.LotteryParticipationModel.LotteryParticipationUserList(l.ctx, in.UserId, in.IsWon, in.IsAnnounced)
	if err != nil && err != model.ErrNotFound {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "LotteryParticipationUserList Database Exception lottery err: %v", err)
	}

	var pbList []*pb.UserLottery
	for _, v := range list {
		var pbLotteryParticipation pb.UserLottery
		_ = copier.Copy(&pbLotteryParticipation, v)
		pbLotteryParticipation.UpdateTime = v.UpdateTime.Unix()
		pbLotteryParticipation.AnnounceTime = v.AnnounceTime.Unix()
		pbLotteryParticipation.AwardDeadline = v.AwardDeadline.Unix()
		pbLotteryParticipation.PublishTime = v.PublishTime.Unix()
		pbList = append(pbList, &pbLotteryParticipation)
	}

	return &pb.SearchUserLotteryParticipationResp{UserLottery: pbList}, nil
}
