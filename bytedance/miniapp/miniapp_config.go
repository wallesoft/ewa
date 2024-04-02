package miniapp

import (
	"gitee.com/wallesoft/ewa/kernel/base"
	"gitee.com/wallesoft/ewa/kernel/log"
	"github.com/gogf/gf/v2/net/gclient"
	"github.com/gogf/gf/v2/os/gcache"
)

type Config struct {
	AppID  string // appid
	Secret string // secret
	Cache  *gcache.Cache
	Logger *log.Logger
}

// 默认接口配置地址
func (app *MiniAPP) getBaseUri() string {
	return "https://developer.toutiao.com/api/apps/v2/"
}

// 沙盒接口地址
func (app *MiniAPP) getSandboxBaseUri() string {
	return "https://open-sandbox.douyin.com/api/apps/v2/"
}

// client  default without token
func (app *MiniAPP) GetClient() *base.Client {
	return &base.Client{
		Client:  gclient.New(),
		BaseUri: app.getBaseUri(),
		Logger:  app.Logger,
	}
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
