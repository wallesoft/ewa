package openplatform

import (
	"errors"

	serverguard "gitee.com/wallesoft/ewa/kernel/server"
	"gitee.com/wallesoft/ewa/openplatform/server"
	"github.com/gogf/gf/os/glog"
	"github.com/gogf/gf/util/gconv"
	"github.com/gogf/gf/util/gutil"
)

type OpenPlatform struct {
	config *Config
	logger *glog.Logger
}

// type Config *gmap.StrAnyMap

type Config struct {
	*serverguard.Config
	// AppId  string
	// Secret string
	// Token  string
	// AesKey string
}

//这地方需要改动 参考 glog setconfigfrommap
//Config 结构需要改动
func New(config map[string]interface{}) (*OpenPlatform, error) {
	if config == nil || len(config) == 0 {
		return nil, errors.New("Openplatform configuration cannot be empty")
	}
	config = gutil.MapCopy(config)

	log := glog.New()
	_, logVal := gutil.MapPossibleItemByKey(config, "Logger")
	if logVal != nil {
		if err := log.SetConfigWithMap(gconv.Map(logVal)); err != nil {
			return nil, err
		}
	}
	// debug
	_, debugVal := gutil.MapPossibleItemByKey(config,
		"Debug")
	if debugVal != nil {
		log.SetDebug(gconv.Bool(debugVal))
	}
	c := serverguard.Config{}
	if err := gconv.Struct(config, &c); err != nil {
		return nil, err
	}

	return &OpenPlatform{
		config: &Config{
			&c,
		},
		logger: log,
	}, nil
}

func (op *OpenPlatform) Server() *server.Server {
	return &server.Server{
		&serverguard.ServerGuard{
			Logger: op.logger,
			Config: op.config,
		},
	}
}
