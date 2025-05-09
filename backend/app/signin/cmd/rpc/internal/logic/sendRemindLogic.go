package logic

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/kernel/power"
	"github.com/pkg/errors"
	"looklook/common/kqueue"
	"looklook/common/wxnotice"
	"looklook/common/xerr"
	"time"

	"looklook/app/signin/cmd/rpc/internal/svc"
	"looklook/app/signin/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type SendRemindLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSendRemindLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SendRemindLogic {
	return &SendRemindLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SendRemindLogic) SendRemind(in *pb.SendRemindReq) (*pb.SendRemindResp, error) {
	// 构建微信消息
	now := time.Now()
	msg := wxnotice.MessageReminder{
		ReminderContent: wxnotice.Item{Value: "签到提醒"},
		ReminderTime:    wxnotice.Item{Value: now.Format("2006-01-02 15:04")},
	}
	data, err := power.StructToHashMap(&msg)

	// 拼接小程序页面地址
	//pageAddr := fmt.Sprintf("pages/index/index?lotterId=%d&userId=%d", rpcLotteryInfo.Lottery.Id, in.UserId)
	pageAddr := fmt.Sprintf("pages/signin/index/index")

	reminds, err := SearchRemindFunction(l)
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "WxMiniRemindSigninHandler SearchRemind err:%v", err)
	}
	for _, remind := range reminds.Reminds {

		m := kqueue.WXMiniNoticeSendMessage{
			ToUser: remind.AuthKey,
			Page:   pageAddr,
			UserId: remind.UserId,
			Data:   data,
		}
		body, err := json.Marshal(m)
		err = l.svcCtx.KqueueNoticeSendClient.Push(l.ctx, string(body))
		if err != nil {
			return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "KqueueNoticeSendClient.Push err: %v", err)
		}

	}

	return &pb.SendRemindResp{}, nil
}
