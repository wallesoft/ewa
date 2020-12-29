package payment

import (
	"gitee.com/wallesoft/ewa/kernel/log"
	"gitee.com/wallesoft/ewa/kernel/cache"
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
func (p *Payment) getKey(endpoint string) string {
	if endpoint == "sandboxnew/pay/getsignkey" {
		return p.config.Key
	}
	var key string
	if p.isSanbox {
		key = p.Sanbox().GetKey()
	}
}

func (p *Payment) isSandbox() bool {
	return p.config.Sandbox
}
