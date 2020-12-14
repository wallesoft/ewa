package openplatform

import (
	baseauth "gitee.com/wallesoft/ewa/kernel/auth"
	"gitee.com/wallesoft/ewa/openplatform/auth"
	"github.com/gogf/gf/os/gcache"
	"github.com/gogf/gf/os/glog"
)

type Config struct {
	AppID          string `json:"app_id"`           //app_id
	AppSecret      string `json:"app_secret"`       //app_secret
	Token          string `json:"token"`            //token
	EncodingAESKey string `json:"encoding_aes_key"` //encoding aes key
	Cache          *gcache.Cache
	Logger         *glog.Logger
}

// SetLogger
func (s *OpenPlatform) SetLogger(logger *glog.Logger) {
	s.config.Logger = logger
}

// SetCache
func (s *OpenPlatform) SetCache(c *gcache.Cache) {
	s.config.Cache = c
}

//SetVrifyTicket
//需要自定义解决ticket的存储及获取问题时需要设置满足相关接口的对象
func (op *OpenPlatform) SetVerifyTicket(ticket auth.VerifyTicket) {
	op.verifyTicket = ticket
}

//SetAccessToken
//设置的需要满足接口
func (op *OpenPlatform) SetAccessToken(token baseauth.AccessToken) {
	op.accessToken = token
}

//getBaseUri return openplatform baseuri
func (op *OpenPlatform) getBaseUri() string {
	return "https://api.weixin.qq.com/"
}

// func (op *OpenPlatform)
