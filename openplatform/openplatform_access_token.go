package openplatform

import (
	"gitee.com/wallesoft/ewa/kernel/auth"
	"gitee.com/wallesoft/ewa/kernel/base"
	"github.com/gogf/gf/crypto/gmd5"
)

var defaultAccessToken = &base.AccessToken{}

func (op *OpenPlatform) getDefaultAccessToken() auth.AccessToken {
	defaultAccessToken.Cache = op.config.Cache
	defaultAccessToken.TokenKey = "component_access_token"
	defaultAccessToken.EndPoint = "cgi-bin/component/api_component_token"
	defaultAccessToken.RequestPostMethod = true
	defaultAccessToken.Credentials = map[string]string{
		"component_appid":         op.config.AppID,
		"component_appsecret":     op.config.AppSecret,
		"component_verify_ticket": op.GetVerifyTicket(),
	}
	defaultAccessToken.CacheKey = "ewa.access_token." + gmd5.MustEncrypt(defaultAccessToken.Credentials)
	defaultAccessToken.Client = op.getClient()
	return defaultAccessToken
}
