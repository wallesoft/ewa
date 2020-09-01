package openplatform

import (
	"gitee.com/wallesoft/ewa/kernel/cache"
	"github.com/gogf/gf/os/glog"
)

type Config struct {
	AppID          string `json:"app_id"`           //app_id
	AppSecret      string `json:"app_secret"`       //app_secret
	Token          string `json:"token"`            //token
	EncodingAESKey string `json:"encoding_aes_key"` //encoding aes key
	Cache          cache.Cache
	Logger         *glog.Logger
}

// func GetConfigWithMap(m map[string]interface{}) error {

// }
