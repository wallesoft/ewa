package bytedance

import (
	"gitee.com/wallesoft/ewa/kernel/base"
	"gitee.com/wallesoft/ewa/kernel/log"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/os/gcache"
)

type AppConfig struct {
	AppID  string // appid
	Secret string // secret
	Cache  *gcache.Cache
	Logger *log.Logger
}

func (app *MiniApp) getBaseUri() string {
	return "https://developer.toutiao.com/api/"
}
func (app *MiniApp) GetClient() *base.Client {
	return &base.Client{
		Client:  ghttp.NewClient(),
		BaseUri: app.getBaseUri(),
		Logger:  app.Logger,
	}
}

func (app *MiniApp) GetClientWithToken() *base.Client {
	return &base.Client{
		Client:  ghttp.NewClient(),
		BaseUri: app.getBaseUri(),
		Logger:  app.Logger,
		Token:   app.AccessToken,
	}
}
