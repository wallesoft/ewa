package miniapp

import (
	"context"

	"gitee.com/wallesoft/ewa/kernel/http"
	"github.com/gogf/gf/v2/os/gcache"
)

// 凭证
type Credentials struct {
	// clientKey    string // 应用唯一标识，对应小程序id
	// clientSecret string //应用唯一标识对应的密钥，对应小程序的app secret，可以在开发者后台获取
	miniapp *MiniApp
}

// 实现TokenCredentail接口
func (c *Credentials) Get(ctx context.Context) map[string]string {
	return map[string]string{
		"grant_type":    "client_credential",
		"client_key":    c.miniapp.Config.AppID,
		"client_secret": c.miniapp.Config.Secret,
	}
}

type ClientToken struct {
	Cache  *gcache.Cache
	Client *http.Client
}

func (ct *ClientToken) GetToken(ctx context.Context, refresh ...bool) string {

	return ""
}

var defaultClientToken = &ClientToken{}

// func (app *MiniApp) getDefaultAccessToken(ctx context.Context) *ClientToken {
// 	defaultClientToken.Cache = app.Config.Cache
// 	defaultClientToken.TokenKey = "data.access_token"
// 	defaultClientToken.EndPoint = "oauth/client_token"
// 	defaultClientToken.RequestPostMethod = true // post 请求
// 	defaultClientToken.Credentials = &Credentials{miniapp: app}
// 	defaultClientToken.CacheKey = "ewa.bytedance_miniapp." + gmd5.MustEncrypt(defaultClientToken.Credentials.Get(ctx))
// 	client := app.GetClient()
// 	if app.Config.Sandbox {
// 		client.BaseUri = "https://open-sandbox.douyin.com/"
// 	} else {
// 		client.BaseUri = "https://open.douyin.com/"
// 	}
// 	defaultClientToken.Client = client
// }
