package miniapp

import (
	"fmt"
	// "log"

	"gitee.com/wallesoft/ewa/kernel/auth"
	"gitee.com/wallesoft/ewa/kernel/cache"
	"gitee.com/wallesoft/ewa/kernel/log"
	"github.com/gogf/gf/os/gcache"
)

//字节跳动小程序

type MiniApp struct {
	Config      Config
	AccessToken auth.AccessToken
	Logger      *log.Logger
	Cache       *gcache.Cache
}

//New
func New(config Config) *MiniApp {
	app := NewWithOutToken(config)
	app.AccessToken = app.getDefaultAccessToken()
	return app
}

func NewWithOutToken(config Config) *MiniApp {
	if config.Cache == nil {
		config.Cache = cache.New("ewa.wechat.miniprogram")
	}
	if config.Logger == nil {
		config.Logger = log.New()
		if config.Logger.LogPath != "" {
			if err := config.Logger.SetPath(config.Logger.LogPath); err != nil {
				panic(fmt.Sprintf("[miniprogram] set log path '%s' error: %v", config.Logger.LogPath, err))
			}
		}

		// default set close debug / close stdout print
		config.Logger.LogStdout = false
	}
	var app = &MiniApp{
		Config: config,
		Logger: config.Logger,
		Cache:  config.Cache,
	}
	return app
}
