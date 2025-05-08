// KqMessage
package kqueue

import "github.com/ArtisanCloud/PowerWeChat/v3/src/kernel/power"

// 第三方支付回调更改支付状态通知
type ThirdPaymentUpdatePayStatusNotifyMessage struct {
	PayStatus int64  `json:"payStatus"`
	OrderSn   string `json:"orderSn"`
}

// 微信小程序发送通知
type WXMiniNoticeSendMessage struct {
	ToUser string         `json:"touser"`
	Page   string         `json:"page,omitempty"`
	Data   *power.HashMap `json:"data"`
	UserId int64          `json:"userId"`
}
