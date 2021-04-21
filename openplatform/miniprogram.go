package openplatform

import baseauth "gitee.com/wallesoft/ewa/kernel/auth"

//小程序
type MiniProgram struct {
	OpenPlatform *OpenPlatform
	accessToken  baseauth.AccessToken
	Appid        string
	RefreshToken string
}
