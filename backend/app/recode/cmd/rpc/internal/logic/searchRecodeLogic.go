package logic

import (
	"context"
	"github.com/jinzhu/copier"
	"github.com/pkg/errors"
	"looklook/app/recode/cmd/rpc/internal/svc"
	"looklook/app/recode/cmd/rpc/pb"
	"looklook/app/recode/model"
	"looklook/app/usercenter/cmd/rpc/usercenter"
	"looklook/common/xerr"
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
	userInfo, err := l.svcCtx.UsercenterRpc.GetUserInfo(l.ctx, &usercenter.GetUserInfoReq{Id: in.UserId})
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "查询记录失败 err: %v", err)
	}

	list, err := l.svcCtx.RecodeModel.RecodeList(l.ctx, in.UserId, in.ItemsId)
	if err != nil && !errors.Is(err, model.ErrNotFound) {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "查询记录失败 err: %v", err)
	}

	if userInfo.User.IntimateId != 0 {
		list1, err := l.svcCtx.RecodeModel.RecodeList(l.ctx, userInfo.User.IntimateId, in.ItemsId)
		if err != nil {
			return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "查询记录失败 err: %v", err)
		}
		list = append(list, list1...)
	}

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
