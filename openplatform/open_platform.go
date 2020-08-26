package openplatform

import (
	kernel "gitee.com/wallesoft/ewa/kernel/server"
	"gitee.com/wallesoft/ewa/openplatform/server"
	"github.com/gogf/gf/container/gmap"
	"github.com/gogf/gf/os/glog"
)

type OpenPlatform struct {
	config Config
	logger *glog.Logger
}
type Config *gmap.StrAnyMap

// type Config struct {
// 	AppId	string	// third open platform APPID
// 	Secret 	string
// 	Token	string
// 	AesKey	string
// }
//这地方需要改动 参考 glog setconfigfrommap
//Config 结构需要改动
func New(config map[string]interface{}) *OpenPlatform {
	if config == nil || lne(config) == 0 {
		return 
	}
	c := gmap.NewStrAnyMapFrom(config)
	return &OpenPlatform{
		config: c,
		logger: glog.New(),
	}
}

func (op *OpenPlatform) Server() *server.Server {
	return &server.Server{
		&kernel.ServerGuard{
			Logger:op.logger,
			Config: op.config,
		}
	}
}
