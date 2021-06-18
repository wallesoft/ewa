package openplatform

import (
	"gitee.com/wallesoft/ewa/kernel/auth"
	"gitee.com/wallesoft/ewa/kernel/base"
	"github.com/gogf/gf/crypto/gmd5"
)

type Credentials struct {
	op *OpenPlatform
}

func (c *Credentials) Get() map[string]string {
	return map[string]string{
		"component_appid":         c.op.config.AppID,
		"component_appsecret":     c.op.config.AppSecret,
		"component_verify_ticket": c.op.GetVerifyTicket(),
	}
}

var defaultAccessToken = &base.AccessToken{}

func (op *OpenPlatform) getDefaultAccessToken() auth.AccessToken {
	defaultAccessToken.Cache = op.config.Cache
	defaultAccessToken.TokenKey = "component_access_token"
	defaultAccessToken.EndPoint = "cgi-bin/component/api_component_token"
	defaultAccessToken.RequestPostMethod = true
	defaultAccessToken.Credentials = &Credentials{op: op}
	defaultAccessToken.CacheKey = "ewa.access_token." + gmd5.MustEncrypt(defaultAccessToken.Credentials.Get())
	defaultAccessToken.Client = op.getClient()
	defaultAccessToken.SetToken("45_KylNwcHHI070YRWlqLdYzpEVtvp0aDiZlPF-85IkZoWdvYkIE9jxThs5XHKKkdOB_vsMZtBbMfEmBvSzywe4FfsKyT6Jrjbg7xuHNkf1Cos3wuFmXvAMVShw9q2Tl7Ozj3FZ_iFFVHyFENUpKNIdAAAADY", 7000)
	return defaultAccessToken
}
