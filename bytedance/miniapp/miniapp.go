package miniapp

import (
	"gitee.com/wallesoft/ewa/kernel/log"
	"github.com/gogf/gf/v2/os/gcache"
)

type MiniAPP struct {
	Config Config // 配置
	// AccessToken auth.
	Logger *log.Logger
	Cache  *gcache.Cache
}
