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
	defaultAccessToken.SetToken("44_Xtdx5WA2Py9nCM2BNlA9QUvBdlvIwR3vSddYy_-xgiDHwH4pZ9P9mBo0WcMGmwtIG1b6Drjhv5NmHqmVeGHQ_cV85CFdBgJR24v9m-HVNTbQcuIcrm4pkWmEgpd5YLebKCXLAI0pGC8fhWmgMIWdABAZVJ", 7200)
	return defaultAccessToken
}
