package miniapp

import (
	"context"
	"fmt"

	"gitee.com/wallesoft/ewa/kernel/cache"
	"gitee.com/wallesoft/ewa/kernel/log"
	"github.com/gogf/gf/v2/os/gcache"
)

type MiniAPP struct {
	Config Config // 配置
	// AccessToken auth.
	Logger *log.Logger
	Cache  *gcache.Cache
}

func New(ctx context.Context, config Config) *MiniAPP {
	app := NewWithOutToken(config)
	app.AccessToken = app.getDefaultAccessToken(ctx)
	return app
}

func NewWithOutToken(config Config) *MiniApp {
	if config.Cache == nil {
		config.Cache = cache.New("ewa.toutiao.miniapp")
	}
	if config.Logger == nil {
		config.Logger = log.New()
		if config.Logger.LogPath != "" {
			if err := config.Logger.SetPath(config.Logger.LogPath);err != nil {
				panic(fmt.Sprintf("[toutiao-miniapp] set log path '%s' err: %v", config.Logger.LogPath, err))
			}
		}
		config.Logger.LogStdout = false
	}

	var app = &MiniAPP{
		Config: config,
		Logger: config.Logger,
		Cache:config.Cache
	}
	return app

}
