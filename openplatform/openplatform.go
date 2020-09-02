package openplatform

import (
	"net/http"

	"gitee.com/wallesoft/ewa/kernel/cache"
	"gitee.com/wallesoft/ewa/openplatform/server"
	"github.com/gogf/gf/encoding/gjson"
	"github.com/gogf/gf/os/glog"
)

//OpenPlatform
type OpenPlatform struct {
	// config *Config
	// logger *glog.Logger
	*Context
}

//Context struct
type Context struct {
	*Config
}

// type Config *gmap.StrAnyMap

// type Config struct {
// 	Appid  string `c:"app_id"`
// 	Secret string `c:"secret"`
// 	Token  string `c:"token"`
// 	AesKey string `c:"aes_key"`
// }

//Get get value from config
func (c *Config) Get(pattern string) interface{} {
	j := gjson.New(c)
	return j.Get(pattern)
}

//New new OpenPlatform
//@see glog https://goframe.org/os/glog/index
func New(config *Config) *OpenPlatform {
	if config.Cache == nil {
		config.Cache = cache.NewMemCache()
	}
	if config.Logger == nil {
		config.Logger = glog.New()
	}
	ctx := &Context{
		Config: config,
	}
	return &OpenPlatform{ctx}
}

func (op *OpenPlatform) Server(request *http.Request, writer http.ResponseWriter) *server.Server {
	s := server.New(op.Context)
	return s.GetServer(request, writer)
}

//New get openplatform from config
// func New(config map[string]interface{}) (*OpenPlatform, error) {
// 	if config == nil || len(config) == 0 {
// 		return nil, errors.New("Openplatform configuration cannot be empty")
// 	}
// 	config = gutil.MapCopy(config)

// 	log := glog.New()

// 	_, logVal := gutil.MapPossibleItemByKey(config, "Logger")
// 	if logVal != nil {
// 		if err := log.SetConfigWithMap(gconv.Map(logVal)); err != nil {
// 			return nil, err
// 		}
// 	} else {
// 		log.SetDebug(false)
// 	}
// 	// debug
// 	_, debugVal := gutil.MapPossibleItemByKey(config,
// 		"Debug")
// 	if debugVal != nil {
// 		log.SetDebug(gconv.Bool(debugVal))
// 	}
// 	var c *Config
// 	if err := gconv.Struct(config, &c); err != nil {
// 		return nil, err
// 	}
// 	return &OpenPlatform{
// 		config: c,
// 		logger: log,
// 	}, nil
// }

// func (op *OpenPlatform) Server(r guard.Request) *server.Server {
// 	sg := guard.New(r, op.config, op.logger)
// 	return &server.Server{
// 		ServerGuard: sg,
// 	}
// 	///return guard.New()
// 	// s := server.Server{
// 	// 	ServerGuard: guard.ServerGuard{
// 	// 		Logger: op.logger,
// 	// 		Config: op.config,
// 	// 	}}
// 	// return &s

// }
