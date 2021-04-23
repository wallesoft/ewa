package miniprogram

import (
	"gitee.com/wallesoft/ewa/kernel/base"
	"gitee.com/wallesoft/ewa/kernel/log"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/os/gcache"
)

type Config struct {
	AppID  string //appid
	Secret string //secret
	Cache  *gcache.Cache
	Logger *log.Logger
}

func (mp *MiniProgram) getBaseUri() string {
	return "https://api.weixin.qq.com/"
}

func (mp *MiniProgram) getClient() *base.Client {
	return &base.Client{
		Client:  ghttp.NewClient(),
		BaseUri: mp.getBaseUri(),
		Logger:  mp.Config.Logger,
	}
}

func (mp *MiniProgram) getClientWithToken() *base.Client {
	return &base.Client{
		Client:  ghttp.NewClient(),
		BaseUri: mp.getBaseUri(),
		Logger:  mp.Config.Logger,
		Token:   mp.AccessToken,
	}
}
