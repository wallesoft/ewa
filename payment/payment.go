package payment

import (
	"fmt"

	"gitee.com/wallesoft/ewa/kernel/cache"
	"gitee.com/wallesoft/ewa/kernel/log"
	"github.com/gogf/gf/encoding/gjson"
	"github.com/gogf/gf/os/gcache"
	"github.com/gogf/gf/util/gutil"
)

type Payment struct {
	config Config
	Logger *log.Logger
	Cache  *gcache.Cache
}

//New
func New(config Config, compatible ...bool) *Payment {
	if config.Logger == nil {
		config.Logger = log.New()
		if config.LogPath != "" {
			if err := config.Logger.SetPath(config.Logger.LogPath); err != nil {
				panic(fmt.Sprintf("[openplatform] set log path '%s' error: %v", config.LogPath, err))
			}
		} else {
			config.Logger.SetStdoutPrint(false)
			config.Logger.SetPath("/tmp/log")
			config.Logger.ErrorLogPattern = "ewa.payment-error-{Y-m-d}.log"
		}
	}

	payment := &Payment{
		Logger: config.Logger,
		Cache:  cache.New("ewa.wechat.payment"),
	}

	gutil.TryCatch(func() {
		payment.config = payment.setConfig(config, compatible...)
	}, func(err error) {
		payment.Logger.File(payment.Logger.ErrorLogPattern).Print(fmt.Sprintf("[Erro] %s", err.Error()))
	})

	return payment
}

//Order
func (p *Payment) Order(config ...map[string]interface{}) *Order {
	c := &OrderConfig{
		AppID: p.config.AppID,
		MchID: p.config.MchID,
	}
	oj := gjson.New(c)
	if len(config) > 0 {
		for pattern, val := range config[0] {
			oj.Set(pattern, val)
		}
	}
	return &Order{
		config:  oj,
		payment: p,
	}
}

//Markting
func (p *Payment) Marketing() *Marketing {
	return &Marketing{
		payment: p,
	}
}
