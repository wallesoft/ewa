package miniprogram

import (
	"gitee.com/wallesoft/ewa/kernel/auth"
	"gitee.com/wallesoft/ewa/kernel/base"
	"github.com/gogf/gf/crypto/gmd5"
)

type Credentials struct {
	miniprogram *MiniProgram
}

func (c *Credentials) Get() map[string]string {
	return map[string]string{
		"appid":      c.miniprogram.Config.AppID,
		"secret":     c.miniprogram.Config.Secret,
		"grant_type": "client_credential",
	}
}

var defaultAccessToken = &base.AccessToken{}

func (mp *MiniProgram) getDefaultAccessToken() auth.AccessToken {
	defaultAccessToken.Cache = mp.Config.Cache
	defaultAccessToken.TokenKey = "access_token"
	defaultAccessToken.EndPoint = "cgi-bin/token"
	defaultAccessToken.RequestPostMethod = false // GET 请求
	defaultAccessToken.Credentials = &Credentials{miniprogram: mp}
	defaultAccessToken.CacheKey = "ewa.weapp_access_token." + gmd5.MustEncrypt(defaultAccessToken.Credentials.Get())
	defaultAccessToken.Client = mp.GetClient()
	return defaultAccessToken
}
