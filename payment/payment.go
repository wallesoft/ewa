package payment

import (
	"gitee.com/wallesoft/ewa/kernel/cache"
	"gitee.com/wallesoft/ewa/kernel/log"
	"github.com/gogf/gf/encoding/gjson"
	"github.com/gogf/gf/os/gcache"
)

type Payment struct {
	config Config
	Logger *log.Logger
	Cache  *gcache.Cache
}

//New
func New(config Config) *Payment {
	return &Payment{
		config: config,
		Logger: log.New(),
		Cache:  cache.New("ewa.wechat.payment"),
	}
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
		config: oj,
	}
}
