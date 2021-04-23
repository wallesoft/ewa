package miniprogram

import (
	"fmt"

	"gitee.com/wallesoft/ewa/kernel/auth"
	"gitee.com/wallesoft/ewa/kernel/cache"
	"gitee.com/wallesoft/ewa/kernel/log"
)

type MiniProgram struct {
	Config       Config
	AccessToken  auth.AccessToken
	RefreshToken string // 第三方平台用
}

//New
func New(config Config) *MiniProgram {
	app := NewWithOutToken(config)
	app.AccessToken = app.getDefaultAccessToken()
	return app
}

func NewWithOutToken(config Config) *MiniProgram {
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
	var app = &MiniProgram{
		Config: config,
	}
	return app
}
