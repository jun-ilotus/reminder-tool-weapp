package config

import (
	"github.com/zeromicro/go-queue/kq"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	service.ServiceConf

	Redis      redis.RedisConf
	WxMiniConf WxMiniConf
	WxMsgConf  WxMsgConf

	// kq : pub sub
	NoticeSendConf kq.KqConf

	// rpc
	NoticeRpcConf     zrpc.RpcClientConf
	UsercenterRpcConf zrpc.RpcClientConf
}
