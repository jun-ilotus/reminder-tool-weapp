package jobtype

// DeferCloseHomestayOrderPayload defer close homestay order
type DeferCloseHomestayOrderPayload struct {
	Sn string
}

// PaySuccessNotifyUserPayload pay success notify user
type PaySuccessNotifyUserPayload struct {
}

// WxMiniProgramNotifyUserPayload mini program notify user
type WxMiniProgramNotifyUserPayload struct {
	ReminderId int64
}
