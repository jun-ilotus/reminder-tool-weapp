package jobtype

import (
	"looklook/app/order/model"
)

// DeferCloseHomestayOrderPayload defer close homestay order
type DeferCloseHomestayOrderPayload struct {
	Sn string
}

// PaySuccessNotifyUserPayload pay success notify user
type PaySuccessNotifyUserPayload struct {
	Order *model.HomestayOrder
}

// WxMiniProgramNotifyUserPayload mini program notify user
type WxMiniProgramNotifyUserPayload struct {
	MsgType      int
	OpenId       string
	PageAddr     string
	Data         string
	ReminderId   int
	ReminderTime int64
}
