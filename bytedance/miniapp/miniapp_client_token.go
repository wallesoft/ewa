package miniapp

import (
	"context"
	"io"
	"net/http"

	"gitee.com/wallesoft/ewa/internal/utils"
	ehttp "gitee.com/wallesoft/ewa/kernel/http"
	"github.com/gogf/gf/v2/net/gclient"
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
	Client *ehttp.Client
}

const (
	// 接口需要携带的token key
	clientTokenKey = "access-token"
	// client token 错误码
	clientTokenErrorCode = 28001003
)

// GetToken
func (ct *ClientToken) GetToken(ctx context.Context, refresh ...bool) string {

	return ""
}

// SetToken 前置注入请求token
func (ct *ClientToken) SetToken(ctx context.Context) gclient.HandlerFunc {
	return func(c *gclient.Client, r *http.Request) (*gclient.Response, error) {
		// 在此处注入token
		c.SetHeader(clientTokenKey, ct.GetToken(ctx))
		resp, err := c.Next(r)
		return resp, err
	}
}

// VerifyResponse 后置校验响应状态码, 如果token过期，刷新token
func (ct *ClientToken) VerifyResponse(ctx context.Context) gclient.HandlerFunc {
	return func(c *gclient.Client, r *http.Request) (resp *gclient.Response, err error) {
		reqBodyContent, _ := io.ReadAll(r.Body)
		r.Body = utils.NewReadCloser(reqBodyContent, false)
		resp, err = c.Next(r)
		// 此处校验，刷新逻辑
		bodyContent := resp.ReadAll()
		if ct.isExpired(bodyContent) {
			token := ct.GetToken(ctx, true)
			// 重新请求
			r.Header.Set(clientTokenKey, token)
			r.Body = utils.NewReadCloser(bodyContent, false)
			retryResp, err := c.Do(r)
			resp.Response = retryResp
			return resp, err
		}
		resp.SetBodyContent(bodyContent)
		return resp, err
	}
}

// 检查返回错误
func (ct *ClientToken) isExpired(content []byte) bool {
	// if gjson.Valid(content) {
	// 	res := gjson.New(content)
	// 	if have := res.Cont
	// }
	return false
}

var defaultClientToken = &ClientToken{}

func (app *MiniApp) getClientToken() *ClientToken {
	return defaultClientToken
}

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
