package cache

import (
	"time"

	"github.com/gogf/gf/os/gcache"
)

type Cache interface {
	Get(key string) interface{}
	Set(key string, val interface{}, duration time.Duration)
	Contains(key string) bool
	Remove(key string) error
}

//DefaultCache is alias of gcache.Cache
type DefaultCache struct {
	*gcache.Cache
}

func NewMemCache() *DefaultCache {
	gc := gcache.New()
	return &DefaultCache{
		Cache: gc,
	}
}
func (c *DefaultCache) Get(key string) interface{} {
	return c.Cache.Get(key)
}
func (c *DefaultCache) Set(key string, value interface{}, duration time.Duration) {
	c.Cache.Set(key, value, duration)
}
func (c *DefaultCache) Contains(key string) bool {
	return c.Cache.Contains(key)
}
func (c *DefaultCache) Remove(key string) error {
	c.Cache.Remove(key)
	return nil
}
