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

// func (c *DefaultCache) Get(key string) interface{} {
// 	return c.Cache.Get(key)
// }
// func (c *DefaultCache) Set(key string, value interface{}, duration time.Duration) {
// 	c.Cache.Set(key, value, duration)
// }
// func (c *DefaultCache) Contains(key string) bool {
// 	return c.Cache.Contains(key)
// }
// func (c *DefaultCache) Remove(key string) error {
// 	c.Cache.Remove(key)
// 	return nil
// }
