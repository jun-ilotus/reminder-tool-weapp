// Code generated by goctl. DO NOT EDIT.
package types

type BindIntimateReq struct {
	IntimateAccessToken string `json:"intimateAccessToken"`
}

type BindIntimateResp struct {
	IntimateUserId int64 `json:"intimateUserId"`
}

type CancelBindIntimateReq struct {
}

type CancelBindIntimateResp struct {
}

type LoginReq struct {
	Mobile   string `json:"mobile"`
	Password string `json:"password"`
}

type LoginResp struct {
	AccessToken  string `json:"accessToken"`
	AccessExpire int64  `json:"accessExpire"`
	RefreshAfter int64  `json:"refreshAfter"`
}

type RegisterReq struct {
	Mobile   string `json:"mobile"`
	Password string `json:"password"`
}

type RegisterResp struct {
	AccessToken  string `json:"accessToken"`
	AccessExpire int64  `json:"accessExpire"`
	RefreshAfter int64  `json:"refreshAfter"`
}

type User struct {
	Id               int64  `json:"id"`
	Mobile           string `json:"mobile"`
	Nickname         string `json:"nickname"`
	Sex              int64  `json:"sex"`
	Avatar           string `json:"avatar"`
	Info             string `json:"info"`
	IntimateId       int64  `json:"intimateId"`
	IntimateNickname string `json:"intimateNickname"`
	IntimateAvatar   string `json:"intimateAvatar"`
}

type UserInfoReq struct {
}

type UserInfoResp struct {
	UserInfo User `json:"userInfo"`
}

type WXMiniAuthReq struct {
	Code          string `json:"code"`
	IV            string `json:"iv,optional"`
	EncryptedData string `json:"encryptedData,optional"`
}

type WXMiniAuthResp struct {
	AccessToken  string `json:"accessToken"`
	AccessExpire int64  `json:"accessExpire"`
	RefreshAfter int64  `json:"refreshAfter"`
}

type ModifyUserInfoReq struct {
	Avatar   string `json:"avatar"`
	Nickname string `json:"nickname"`
}

type ModifyUserInfoResp struct {
}
