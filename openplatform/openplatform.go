package openplatform

import (
	"fmt"
	"net/http"

	baseauth "gitee.com/wallesoft/ewa/kernel/auth"
	"gitee.com/wallesoft/ewa/kernel/cache"
	"gitee.com/wallesoft/ewa/kernel/encryptor"
	"gitee.com/wallesoft/ewa/kernel/log"
	guard "gitee.com/wallesoft/ewa/kernel/server"
	"gitee.com/wallesoft/ewa/miniprogram"
	"gitee.com/wallesoft/ewa/openplatform/auth"
	"gitee.com/wallesoft/ewa/openplatform/server"
)

//OpenPlatform
type OpenPlatform struct {
	config       Config
	accessToken  baseauth.AccessToken
	verifyTicket auth.VerifyTicket
}

//New new OpenPlatform
//@see glog https://goframe.org/os/glog/index
func New(config Config) *OpenPlatform {
	if config.Cache == nil {
		config.Cache = cache.New("ewa.wechat.openplatform")
	}
	if config.Logger == nil {
		config.Logger = log.New()
		if config.Logger.LogPath != "" {
			if err := config.Logger.SetPath(config.Logger.LogPath); err != nil {
				panic(fmt.Sprintf("[openplatform] set log path '%s' error: %v", config.Logger.LogPath, err))
			}
		}
		// default set close debug / close stdout print
		// config.Logger.SetDebug(false)
		// config.Logger.SetStdoutPrint(false)
	}

	var op = &OpenPlatform{
		config:       config,
		verifyTicket: auth.GetVerifyTicket(config.AppID, config.Cache),
	}
	op.accessToken = op.getDefaultAccessToken()
	return op
}

//Server
func (op *OpenPlatform) Server(request *http.Request, writer http.ResponseWriter) *server.Server {
	gs := guard.New(guard.Config{
		AppID:          op.config.AppID,
		AppSecret:      op.config.AppSecret,
		Token:          op.config.Token,
		EncodingAESKey: op.config.EncodingAESKey,
	}, request, writer)

	gs.Logger = log.New()
	gs.Logger.SetStdoutPrint(false)
	server := &server.Server{
		ServerGuard: gs,
	}

	server.SetMux()

	server.Encryptor = encryptor.New(encryptor.Config{
		AppID:          op.config.AppID,
		Token:          op.config.Token,
		EncodingAESKey: op.config.EncodingAESKey,
		BlockSize:      32,
	})
	server.Guard = server
	return server
}

//MiniProgram
func (op *OpenPlatform) MiniProgram(appid string, refreshToken string) *miniprogram.MiniProgram {

	app := miniprogram.New(miniprogram.Config{
		AppID: appid,
	})
	app.RefreshToken = refreshToken
	app.AccessToken = op.getWeappAccessToken(app)
	return app
}
