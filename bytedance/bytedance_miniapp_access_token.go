package bytedance

import (
	"gitee.com/wallesoft/ewa/kernel/auth"
	"gitee.com/wallesoft/ewa/kernel/base"
	"github.com/gogf/gf/crypto/gmd5"
)

type Credentials struct {
	miniapp *MiniApp
}

func (c *Credentials) Get() map[string]string {
	return map[string]string{
		"appid":      c.miniapp.Config.AppID,
		"secret":     c.miniapp.Config.Secret,
		"grant_type": "client_credential",
	}
}

var defaultAccessToken = &base.AccessToken{}

func (ma *MiniApp) getDefaultAccessToken() auth.AccessToken {
	defaultAccessToken.Cache = ma.Config.Cache
	defaultAccessToken.TokenKey = "access_token"
	defaultAccessToken.EndPoint = "apps/token"
	defaultAccessToken.RequestPostMethod = false // GET 请求
	defaultAccessToken.Credentials = &Credentials{miniapp: ma}
	defaultAccessToken.CacheKey = "ewa.bytedance_miniapp_access_token." + gmd5.MustEncrypt(defaultAccessToken.Credentials.Get())
	defaultAccessToken.Client = ma.GetClient()
	defaultAccessToken.RefreshTokenCode = 40002 // 字节小程序token错误为40002
	return defaultAccessToken
}
