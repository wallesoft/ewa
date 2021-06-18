package officialaccount

import (
	"gitee.com/wallesoft/ewa/kernel/auth"
	"gitee.com/wallesoft/ewa/kernel/base"
	"gitee.com/wallesoft/ewa/kernel/log"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/os/gcache"
)

type Config struct {
	AppID  string        // appid
	Secret string        // secret
	Cache  *gcache.Cache //cache
	Logger *log.Logger   //logger
}

//logger -------------
func (oa *OfficialAccount) ConfigLoggerWithMap(m map[string]interface{}) {
	oa.Logger.SetConfigWithMap(m)
}

//SetAccessToken 需要传入满足接口
func (oa *OfficialAccount) SetAccessToken(token auth.AccessToken) {
	oa.accessToken = token
}

// SetCache
func (oa *OfficialAccount) SetCache(c *gcache.Cache) {
	oa.config.Cache = c
}

func (oa *OfficialAccount) getBaseUri() string {
	return "https://api.weixin.qq.com/"
}

func (oa *OfficialAccount) GetClient() *base.Client {
	return &base.Client{
		Client:  ghttp.NewClient(),
		BaseUri: oa.getBaseUri(),
		Logger:  oa.Logger,
	}
}

func (oa *OfficialAccount) GetClientWithToken() *base.Client {
	return &base.Client{
		Client:  ghttp.NewClient(),
		BaseUri: oa.getBaseUri(),
		Logger:  oa.Logger,
		Token:   oa.getDefaultAccessToken(), // >>>>???????? 能否用oa.accessToken
	}
}
