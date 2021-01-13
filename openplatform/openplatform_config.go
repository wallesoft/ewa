package openplatform

import (
	baseauth "gitee.com/wallesoft/ewa/kernel/auth"
	"gitee.com/wallesoft/ewa/kernel/base"
	"gitee.com/wallesoft/ewa/kernel/log"
	"gitee.com/wallesoft/ewa/openplatform/auth"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/os/gcache"
)

type Config struct {
	AppID          string `json:"app_id"`           //app_id
	AppSecret      string `json:"app_secret"`       //app_secret
	Token          string `json:"token"`            //token
	EncodingAESKey string `json:"encoding_aes_key"` //encoding aes key
	Cache          *gcache.Cache
	Logger         *log.Logger
}

//  --------------------Del----------------------
// // SetLogger
// func (s *OpenPlatform) SetLogger(logger *glog.Logger) {
// 	op.config.Logger.Logger = logger
// }

//logger -------------
func (op *OpenPlatform) ConfigLoggerWithMap(m map[string]interface{}) {
	op.config.Logger.SetConfigWithMap(m)
}

// func (op *OpenPlatform) ConfigLogger(config glog.Config) {
// 	op.config.Logger.SetConfig(config)
// }

// SetCache
func (op *OpenPlatform) SetCache(c *gcache.Cache) {
	op.config.Cache = c
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

func (op *OpenPlatform) getClient() *base.Client {
	return &base.Client{
		Client:  ghttp.NewClient(),
		BaseUri: op.getBaseUri(),
		Logger:  op.config.Logger,
	}
}

func (op *OpenPlatform) getClientWithToken() *base.Client {
	return &base.Client{
		Client:  ghttp.NewClient(),
		BaseUri: op.getBaseUri(),
		Logger:  op.config.Logger,
		Token:   op.getDefaultAccessToken(),
	}
}

// SetLogStdout sets whether output the logging content to stdout.
func (op *OpenPlatform) SetLogStdout(enabled bool) {
	op.config.Logger.LogStdout = enabled
}

// SetAccessLogEnabled enables/disables the access log.
func (op *OpenPlatform) SetAccessLogEnabled(enabled bool) {
	op.config.Logger.AccessLogEnabled = enabled
}

// SetErrorLogEnabled enables/disables the error log.
func (op *OpenPlatform) SetErrorLogEnabled(enabled bool) {
	op.config.Logger.ErrorLogEnabled = enabled
}

// SetErrorStack enables/disables the error stack feature.
func (op *OpenPlatform) SetErrorStack(enabled bool) {
	op.config.Logger.ErrorStack = enabled
}

// GetLogPath returns the log path.
func (op *OpenPlatform) GetLogPath() string {
	return op.config.Logger.LogPath
}

// IsAccessLogEnabled checks whether the access log enabled.
func (op *OpenPlatform) IsAccessLogEnabled() bool {
	return op.config.Logger.AccessLogEnabled
}

// IsErrorLogEnabled checks whether the error log enabled.
func (op *OpenPlatform) IsErrorLogEnabled() bool {
	return op.config.Logger.ErrorLogEnabled
}
