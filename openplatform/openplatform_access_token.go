package openplatform

import (
	"context"

	"gitee.com/wallesoft/ewa/kernel/auth"
	"gitee.com/wallesoft/ewa/kernel/base"
	"github.com/gogf/gf/v2/crypto/gmd5"
)

type Credentials struct {
	op *OpenPlatform
}

func (c *Credentials) Get(ctx context.Context) map[string]string {
	return map[string]string{
		"component_appid":         c.op.config.AppID,
		"component_appsecret":     c.op.config.AppSecret,
		"component_verify_ticket": c.op.GetVerifyTicket(ctx),
	}
}

var defaultAccessToken = &base.AccessToken{}

func (op *OpenPlatform) getDefaultAccessToken(ctx context.Context) auth.AccessToken {
	defaultAccessToken.Cache = op.config.Cache
	defaultAccessToken.TokenKey = "component_access_token"
	defaultAccessToken.EndPoint = "cgi-bin/component/api_component_token"
	defaultAccessToken.RequestPostMethod = true
	defaultAccessToken.Credentials = &Credentials{op: op}
	defaultAccessToken.CacheKey = "ewa.access_token." + gmd5.MustEncrypt(defaultAccessToken.Credentials.Get(ctx))
	defaultAccessToken.Client = op.getClient()

	return defaultAccessToken
}
