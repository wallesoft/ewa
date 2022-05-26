package officialaccount

import (
	"context"

	"gitee.com/wallesoft/ewa/kernel/auth"
	"gitee.com/wallesoft/ewa/kernel/base"
	"github.com/gogf/gf/v2/crypto/gmd5"
)

type Credentials struct {
	oa *OfficialAccount
}

func (c *Credentials) Get(ctx context.Context) map[string]string {
	return map[string]string{
		"appid":      c.oa.config.AppID,
		"secret":     c.oa.config.Secret,
		"grant_type": "client_credential",
	}
}

var defaultAccessToken = &base.AccessToken{}

func (oa *OfficialAccount) getDefaultAccessToken() auth.AccessToken {
	defaultAccessToken.Cache = oa.config.Cache
	defaultAccessToken.TokenKey = "access_token"
	defaultAccessToken.EndPoint = "cgi-bin/token"
	defaultAccessToken.RequestPostMethod = false // GET 请求
	defaultAccessToken.Credentials = &Credentials{oa: oa}
	defaultAccessToken.CacheKey = "ewa.access_token." + gmd5.MustEncrypt(defaultAccessToken.Credentials.Get(context.TODO()))
	defaultAccessToken.Client = oa.GetClient()
	return defaultAccessToken
}
