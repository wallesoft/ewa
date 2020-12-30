package payment

import (
	"gitee.com/wallesoft/ewa/kernel/cache"
	"gitee.com/wallesoft/ewa/kernel/log"
	"github.com/gogf/gf/0.20201214150022-3517295e9694/text/gstr"
)

type Payment struct {
	config Config
	Logger *log.Logger
	Cache  *gcache.Cache
}

func New(c Config) *Payment {
	return &Payment{
		config: Config,
		Logger: log.New(),
		Cache:  cache.New("ewa.wechat.payment")
	}
}

//Order
func (p *Payment) Order() *Order {
	return &Order{
		payment: p,
	}
}

//Sanbox
func (p *Payment) Sanbox() *Sanbox {

}
func (p *Payment) getKey(endpoint string) key string {
	if endpoint == "sandboxnew/pay/getsignkey" {
		return p.config.Key
	}
	// var key string
	if p.isSanbox {
		key = p.Sanbox().GetKey()
	} else {
		key = p.config.Key
	}
	if key == "" {
		p.Logger.Errorf("config key should not be emtpy")
	}
	if gstr.LenRune(key) != 32 {
		p.Logger.Errorf("%s should be 32 chars length.",key)
	}

	return
}

func (p *Payment) isSandbox() bool {
	return p.config.Sandbox
}
