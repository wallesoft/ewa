package openplatform

import (
	"context"

	"gitee.com/wallesoft/ewa/kernel/auth"
	"gitee.com/wallesoft/ewa/kernel/base"
	"github.com/gogf/gf/v2/crypto/gmd5"
)

type MiniProgramCredentials struct {
	op *OpenPlatform
	mp *MiniProgram
}

func (c *MiniProgramCredentials) Get(ctx context.Context) map[string]string {
	return map[string]string{
		"component_appid":          c.op.config.AppID,
		"authorizer_appid":         c.mp.Config.AppID,
		"authorizer_refresh_token": c.mp.RefreshToken,
	}
}

var defaultWeappAccessToken = &base.AccessToken{}

func (op *OpenPlatform) getWeappAccessToken(ctx context.Context, mp *MiniProgram) auth.AccessToken {
	defaultWeappAccessToken.Cache = op.config.Cache
	defaultWeappAccessToken.TokenKey = "authorizer_access_token"
	defaultWeappAccessToken.RequestTokenKey = "access_token"
	defaultWeappAccessToken.EndPoint = "cgi-bin/component/api_authorizer_token"
	defaultWeappAccessToken.RequestPostMethod = true
	defaultWeappAccessToken.Credentials = &MiniProgramCredentials{mp: mp, op: op}
	defaultWeappAccessToken.CacheKey = "ewa.weapp_access_token." + gmd5.MustEncrypt(defaultWeappAccessToken.Credentials.Get(ctx))
	defaultWeappAccessToken.Client = op.getClientWithToken()
	return defaultWeappAccessToken
}
