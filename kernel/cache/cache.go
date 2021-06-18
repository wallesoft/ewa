package cache

import (
	"gitee.com/wallesoft/ewa/kernel/cache/adapter"
	"github.com/gogf/gf/os/gcache"
	"github.com/gogf/gf/os/gfile"
)

// type Cache interface {
// 	Get(key string) interface{}
// 	Set(key string, val interface{}, duration time.Duration)
// 	Contains(key string) bool
// 	Remove(key string) error
// }

//DefaultCache is alias of gcache.Cache
// type DefaultCache struct {
// 	*gcache.Cache
// }

var defaultCache = New("ewawechat")

func New(dir ...string) *gcache.Cache {
	tmp := gfile.TempDir() + "/"
	if len(dir) > 0 {
		tmp = gfile.TempDir() + "/" + dir[0] + "/"
	}
	cache := gcache.New()
	adapter := adapter.New(tmp)
	cache.SetAdapter(adapter)
	return cache
}

func Get() *gcache.Cache {
	return defaultCache
}
