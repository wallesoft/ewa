package payment

import (
	"context"
	"fmt"

	"gitee.com/wallesoft/ewa/kernel/cache"
	"gitee.com/wallesoft/ewa/kernel/log"
	"github.com/gogf/gf/v2/os/gcache"
	"github.com/gogf/gf/v2/util/gutil"
)

type Payment struct {
	config Config
	Logger *log.Logger
	Cache  *gcache.Cache
}

// New
func New(ctx context.Context, config Config, compatible ...bool) *Payment {
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

	gutil.TryCatch(ctx, func(ctx context.Context) {
		payment.config = payment.setConfig(config, compatible...)
	}, func(ctx context.Context, err error) {
		payment.Logger.File(payment.Logger.ErrorLogPattern).Error(ctx, fmt.Sprintf("[Erro] %s", err.Error()))
	})

	return payment
}

// Order
func (p *Payment) Order() *Order {
	// c := &OrderConfig{
	// 	AppID: p.config.AppID,
	// 	MchID: p.config.MchID,
	// }
	// oj := gjson.New(c)
	// if len(config) > 0 {
	// 	for pattern, val := range config[0] {
	// 		oj.Set(pattern, val)
	// 	}
	// }
	return &Order{
		// config:  oj,
		payment: p,
	}
}

// Markting
func (p *Payment) Marketing() *Marketing {
	return &Marketing{
		payment: p,
	}
}
