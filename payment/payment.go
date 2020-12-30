package payment

import (
	"gitee.com/wallesoft/ewa/kernel/log"
	"github.com/gogf/gf/os/gcache"
)

type Payment struct {
	config *Config
	Logger *log.Logger
	Cache  *gcache.Cache
}

// func New() *Payment {

// }
