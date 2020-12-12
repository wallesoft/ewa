package openplatform

import (
	"net/http"

	"gitee.com/wallesoft/ewa/kernel/base"
	"gitee.com/wallesoft/ewa/kernel/cache"
	"gitee.com/wallesoft/ewa/kernel/encryptor"
	guard "gitee.com/wallesoft/ewa/kernel/server"
	"gitee.com/wallesoft/ewa/openplatform/auth"
	"gitee.com/wallesoft/ewa/openplatform/server"
	"github.com/gogf/gf/os/glog"
)

//OpenPlatform
type OpenPlatform struct {
	config       Config
	accessToken  base.AccessToken
	verifyTicket auth.VerifyTicket
}

//New new OpenPlatform
//@see glog https://goframe.org/os/glog/index
func New(config Config) *OpenPlatform {
	if config.Cache == nil {
		config.Cache = cache.New("ewawechat")
	}
	if config.Logger == nil {
		config.Logger = glog.New()
		// default set close debug / close stdout print
		config.Logger.SetDebug(false)
		config.Logger.SetStdoutPrint(false)
	}

	return &OpenPlatform{
		config:       config,
		verifyTicket: auth.GetVerifyTicket(config.AppID, config.Cache),
	}
}

//Server
func (op *OpenPlatform) Server(request *http.Request, writer http.ResponseWriter) *server.Server {
	gs := guard.New(guard.Config{
		AppID:          op.config.AppID,
		AppSecret:      op.config.AppSecret,
		Token:          op.config.Token,
		EncodingAESKey: op.config.EncodingAESKey,
	}, request, writer)

	gs.Logger = op.config.Logger
	// gs.SetCache(op.config.Cache)

	server := &server.Server{
		ServerGuard: gs,
		// VerifyTicket: op.verifyTicket,
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
