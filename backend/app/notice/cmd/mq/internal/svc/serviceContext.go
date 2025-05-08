package svc

import (
	"github.com/ArtisanCloud/PowerWeChat/v3/src/miniProgram"
	"looklook/app/notice/cmd/mq/internal/config"
	"looklook/app/notice/cmd/rpc/notice"
	"looklook/app/usercenter/cmd/rpc/usercenter"

	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config config.Config

	NoticeRpc     notice.Notice
	UsercenterRpc usercenter.Usercenter
	WxMiniProgram *miniProgram.MiniProgram // 使用
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,

		NoticeRpc:     notice.NewNotice(zrpc.MustNewClient(c.NoticeRpcConf)),
		UsercenterRpc: usercenter.NewUsercenter(zrpc.MustNewClient(c.UsercenterRpcConf)),
		WxMiniProgram: MustNewMiniProgram(c), // 使用
	}
}
