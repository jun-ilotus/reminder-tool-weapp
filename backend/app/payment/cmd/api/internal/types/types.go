// Code generated by goctl. DO NOT EDIT.
package types

type ThirdPaymentWxPayReq struct {
	OrderSn     string `json:"orderSn"`
	ServiceType string `json:"serviceType"`
}

type ThirdPaymentWxPayResp struct {
	Appid     string `json:"appid"`
	NonceStr  string `json:"nonceStr"`
	PaySign   string `json:"paySign"`
	Package   string `json:"package"`
	Timestamp string `json:"timestamp"`
	SignType  string `json:"signType"`
}

type ThirdPaymentWxPayCallbackReq struct {
}

type ThirdPaymentWxPayCallbackResp struct {
	ReturnCode string `json:"return_code"`
}
