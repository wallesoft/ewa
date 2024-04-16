package miniapp

import (
	// "gitee.com/wallesoft/ewa/bytedance/http"

	"net/http"

	ehttp "gitee.com/wallesoft/ewa/kernel/http"
	"gitee.com/wallesoft/ewa/kernel/log"
	"github.com/gogf/gf/v2/net/gclient"
	"github.com/gogf/gf/v2/os/gcache"
)

type Config struct {
	AppID   string // appid
	Secret  string // secret
	Cache   *gcache.Cache
	Logger  *log.Logger
	Sandbox bool
}

// 默认接口配置地址
func (app *MiniApp) getBaseUri() string {
	return "https://developer.toutiao.com/api/apps/v2/"
}

// 沙盒接口地址
func (app *MiniApp) getSandboxBaseUri() string {
	return "https://open-sandbox.douyin.com/api/apps/v2/"
}

// client  default without token
func (app *MiniApp) GetClient(WithToken ...bool) *ehttp.Client {
	baseUri := app.getBaseUri()
	if app.Config.Sandbox {
		baseUri = app.getSandboxBaseUri()
	}

	client := &ehttp.Client{
		Client:  gclient.New(),
		BaseUri: baseUri, //app.getBaseUri(),
		Logger:  app.Logger,
	}
	if len(WithToken) > 0 && WithToken[0] {
		client.BeforeRequest = handleBeforeRequest
		client.AfterReponse = handleAfterResponse
	}
	return client
}

// client before request  set token etc...
func handleBeforeRequest(c *gclient.Client, r *http.Request) (resp *gclient.Response, err error) {
	// r.Header.Add("")
	resp, err = c.Next(r)
	return
}

// client after resposne
func handleAfterResponse(c *gclient.Client, r *http.Request) (resp *gclient.Response, err error) {
	return
}

// client with token
// func(app *MiniAPP) GetClientWithToken() *base.Client {
// 	return &base.Client{
// 		Client:  gclient.New(),
// 		BaseUri: app.getBaseUri(),
// 		Logger:  app.Logger,
// 		Token:
// 	}
// }
