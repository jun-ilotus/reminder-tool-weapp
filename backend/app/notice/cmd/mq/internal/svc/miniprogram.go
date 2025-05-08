package svc

import (
	"github.com/ArtisanCloud/PowerWeChat/v3/src/miniProgram"
	"github.com/zeromicro/go-zero/core/logx"
	"looklook/app/notice/cmd/mq/internal/config"
)

func MustNewMiniProgram(c config.Config) *miniProgram.MiniProgram {
	app, err := miniProgram.NewMiniProgram(&miniProgram.UserConfig{
		AppID:  c.WxMiniConf.AppId,
		Secret: c.WxMiniConf.Secret,
		Token:  c.WxMsgConf.EventToken,
		AESKey: c.WxMsgConf.EncodingAESKey,
	})
	logx.Must(err)
	return app
}
