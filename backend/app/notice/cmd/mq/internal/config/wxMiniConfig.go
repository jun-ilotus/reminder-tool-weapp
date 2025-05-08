package config

// wechat mini config
type WxMiniConf struct {
	AppId  string `json:"AppId"`  //wechat mini appId
	Secret string `json:"Secret"` //wechat mini secret
}

// 微信小程序消息配置
type WxMsgConf struct {
	EventToken     string `json:"EventToken"`     //微信消息回调验证Token
	EncodingAESKey string `json:"EncodingAESKey"` //微信消息加密密钥
}
