package openplatform

import (
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

// func GetConfigWithMap(m map[string]interface{}) error {

// }
