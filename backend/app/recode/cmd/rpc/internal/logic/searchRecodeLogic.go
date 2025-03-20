package logic

import (
	"context"
	"errors"
	"github.com/jinzhu/copier"
	"looklook/app/recode/cmd/rpc/internal/svc"
	"looklook/app/recode/cmd/rpc/pb"
	"looklook/app/recode/model"
	"time"

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
	list, err := l.svcCtx.RecodeModel.RecodeList(l.ctx, in.UserId, in.ItemsId)
	if err != nil && !errors.Is(err, model.ErrNotFound) {
		return nil, err
	}
	// todo 查询出共享的记录
	recodeDayCount := 0
	var recodeDaySpaced float32
	var resp []*pb.Recode
	if len(list) > 0 {
		for _, recode := range list {
			var pbRecode pb.Recode
			_ = copier.Copy(&pbRecode, &recode)
			pbRecode.RecodeTime = recode.RecodeTime.Unix()
			resp = append(resp, &pbRecode)
		}
		recodeDayCount = len(list)
	} else {
		recodeDayCount = 0
	}

	if recodeDayCount > 1 {
		recodeDaySpaced = float32(list[0].RecodeTime.Sub(list[len(list)-1].RecodeTime)/(24*time.Hour)) / float32(recodeDayCount-1)
	} else {
		recodeDaySpaced = 0
	}

	return &pb.SearchRecodeResp{
		Recode:          resp,
		RecodeDayCount:  int64(recodeDayCount),
		RecodeDaySpaced: recodeDaySpaced,
	}, nil
}
