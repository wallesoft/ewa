package openplatform

import (
	"gitee.com/wallesoft/ewa/kernel/auth"
	"gitee.com/wallesoft/ewa/kernel/base"
	"github.com/gogf/gf/crypto/gmd5"
)

type MiniProgramCredentials struct {
	mp *MiniProgram
}

func (c *MiniProgramCredentials) Get() map[string]string {
	return map[string]string{
		"component_appid":          c.mp.OpenPlatform.config.AppID,
		"authorizer_appid":         c.mp.Appid,
		"authorizer_refresh_token": c.mp.RefreshToken,
	}
}

var defaultWeappAccessToken = &base.AccessToken{}

func (mp *MiniProgram) getDefaultAccessToken() auth.AccessToken {
	defaultWeappAccessToken.Cache = mp.OpenPlatform.config.Cache
	defaultWeappAccessToken.TokenKey = "authorizer_access_token"
	defaultWeappAccessToken.EndPoint = "cgi-bin/component/api_authorizer_token"
	defaultWeappAccessToken.RequestPostMethod = true
	defaultWeappAccessToken.Credentials = &MiniProgramCredentials{mp: mp}
	defaultWeappAccessToken.CacheKey = "ewa.weapp_access_token." + gmd5.MustEncrypt(defaultWeappAccessToken.Credentials.Get())
	defaultWeappAccessToken.Client = mp.OpenPlatform.getClientWithToken()
	return defaultWeappAccessToken
}
