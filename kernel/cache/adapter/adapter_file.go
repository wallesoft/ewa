package adapter

import (
	"errors"
	"time"

	"github.com/gogf/gf/container/gvar"
	"github.com/gogf/gf/os/gcache"
	"github.com/gogf/gf/os/gfile"
	"github.com/gogf/gf/os/gtime"
)

const (
	DEFAULT_EXT = ".tmp"
)

type FileCache struct {
	path string //file dir path
	ext  string
}

//new adapter creates and returns a new cache object
func New(path string, ext ...string) gcache.Adapter {
	if !gfile.IsDir(path) {
		if err := gfile.Mkdir(path); err != nil {
			panic(err.Error())
		}
	}
	extension := DEFAULT_EXT
	if len(ext) > 0 {
		extension = "." + ext[0]
	}
	return &FileCache{
		path: path,
		ext:  extension,
	}
}

//SetPath sets where the cache file create
// func (c *FileCache) SetPath(path string) {
// 	c.path = path
// }
// func (c *FileCache) SetExt(ext string) {
// 	c.ext = ext
// }
// @User gcache.Cache.SetAdapter(New("/path/to",".ext"))

// Set sets cache with <key>-<value> pair, which is expired after <duration>.
//
// It does not expire if <duration> == 0.
// It deletes the <key> if <duration> < 0.
func (c *FileCache) Set(key interface{}, value interface{}, duration time.Duration) error {
	var err error
	f := c.path + gvar.New(key).String() + c.ext
	if value == nil || duration < 0 {

		if gfile.IsFile(f) {
			return gfile.Remove(f)
		}
	} else {
		if duration == 0 {
			var t int64 = 0
			err = gfile.PutBytes(f, append(gvar.New(t).Bytes(), gvar.New(value).Bytes()...))
		} else {
			expired := gtime.Timestamp() + gvar.New(duration.Seconds()).Int64()
			err = gfile.PutBytes(f, append(gvar.New(expired).Bytes(), gvar.New(value).Bytes()...))
		}
	}
	return err
}

// Sets batch sets cache with key-value pairs by <data>, which is expired after <duration>.
//
// It does not expire if <duration> == 0.
// It deletes the keys of <data> if <duration> < 0 or given <value> is nil.
func (c *FileCache) Sets(data map[interface{}]interface{}, duration time.Duration) error {
	if len(data) == 0 {
		return nil
	}
	var err error
	for k, v := range data {
		if err = c.Set(k, v, duration); err != nil {
			return err
		}
	}
	return nil
}

// SetIfNotExist sets cache with <key>-<value> pair which is expired after <duration>
// if <key> does not exist in the cache. It returns true the <key> dose not exist in the
// cache and it sets <value> successfully to the cache, or else it returns false.
//
// The parameter <value> can be type of <func() interface{}>, but it dose nothing if its
// result is nil.
//
// It does not expire if <duration> == 0.
// It deletes the <key> if <duration> < 0 or given <value> is nil.
func (c *FileCache) SetIfNotExist(key interface{}, value interface{}, duration time.Duration) (bool, error) {
	var err error
	if f, ok := value.(func() (interface{}, error)); ok {
		value, err = f()
		if value == nil {
			return false, err
		}
	}

	f := c.path + gvar.New(key).String() + c.ext
	// DEL
	if duration < 0 || value == nil {
		err = gfile.Remove(f)
		if err != nil {
			return false, err
		}
		if ok := gfile.IsFile(f); !ok {
			return true, err
		}
	}

	if exist := gfile.IsFile(f); exist {
		return false, err
	} else {
		err = c.Set(key, value, duration)
		if err != nil {
			return false, err
		}
		return true, err
	}
}

// Get retrieves and returns the associated value of given <key>.
// It returns nil if it does not exist or its value is nil.
func (c *FileCache) Get(key interface{}) (interface{}, error) {
	var err error
	f := c.path + gvar.New(key).String() + c.ext
	//now := gtime.Timestamp()
	if exist := gfile.IsFile(f); exist {
		v := gfile.GetBytes(f)
		if v == nil {
			return nil, err
		}
		//EXPIRED
		if len(v) < 8 {
			return nil, err
		}
		t := v[0:8]
		if gvar.New(t).Int64()-gtime.Timestamp() < 0 && gvar.New(t).Int64() != 0 {
			err = gfile.Remove(f)
			return nil, err
		}
		return gvar.New(v[8:]).Interface(), err
	} else {
		return nil, err
	}
}

// GetOrSet retrieves and returns the value of <key>, or sets <key>-<value> pair and
// returns <value> if <key> does not exist in the cache. The key-value pair expires
// after <duration>.
//
// It does not expire if <duration> == 0.
// It deletes the <key> if <duration> < 0 or given <value> is nil, but it does nothing
// if <value> is a function and the function result is nil.
func (c *FileCache) GetOrSet(key interface{}, value interface{}, duration time.Duration) (interface{}, error) {
	v, err := c.Get(key)
	if err != nil {
		return nil, err
	}
	if v == nil {
		return value, c.Set(key, value, duration)
	} else {
		return v, nil
	}

}

// GetOrSetFunc retrieves and returns the value of <key>, or sets <key> with result of
// function <f> and returns its result if <key> does not exist in the cache. The key-value
// pair expires after <duration>.
//
// It does not expire if <duration> == 0.
// It deletes the <key> if <duration> < 0 or given <value> is nil, but it does nothing
// if <value> is a function and the function result is nil.
func (c *FileCache) GetOrSetFunc(key interface{}, f func() (interface{}, error), duration time.Duration) (interface{}, error) {
	v, err := c.Get(key)
	if err != nil {
		return nil, err
	}
	if v == nil {
		value, err := f()
		if err != nil {
			return nil, err
		}
		if value == nil {
			return nil, nil
		}
		return value, c.Set(key, value, duration)
	} else {
		return v, nil
	}
}

// GetOrSetFuncLock retrieves and returns the value of <key>, or sets <key> with result of
// function <f> and returns its result if <key> does not exist in the cache. The key-value
// pair expires after <duration>.
//
// It does not expire if <duration> == 0.
// It does nothing if function <f> returns nil.
//
// Note that the function <f> should be executed within writing mutex lock for concurrent
// safety purpose.
func (c *FileCache) GetOrSetFuncLock(key interface{}, f func() (interface{}, error), duration time.Duration) (interface{}, error) {
	return c.GetOrSetFunc(key, f, duration)
}

// Contains returns true if <key> exists in the cache, or else returns false.
func (c *FileCache) Contains(key interface{}) (bool, error) {
	// f := c.path + gvar.New(key).String() + c.ext
	// return gfile.IsFile(f), nil
	val, err := c.Get(key)
	if val == nil || err != nil {
		return false, err
	}
	return true, nil
}

// GetExpire retrieves and returns the expiration of <key> in the cache.
//
// It returns 0 if the <key> does not expire.
// It returns -1 if the <key> does not exist in the cache.
func (c *FileCache) GetExpire(key interface{}) (time.Duration, error) {
	f := c.path + gvar.New(key).String() + c.ext
	if v := gfile.IsFile(f); !v {
		return -1, nil
	}
	val := gfile.GetBytes(f)
	if len(val) < 8 {
		return 0, errors.New("cache bytes length letter than 8")
	}
	return gvar.New(val[0:8]).Duration() * time.Second, nil
}

// Remove deletes one or more keys from cache, and returns its value.
// If multiple keys are given, it returns the value of the last deleted item.
func (c *FileCache) Remove(keys ...interface{}) (value interface{}, err error) {
	if len(keys) == 0 {
		return nil, nil
	}
	if val, err := c.Get(keys[len(keys)-1]); err != nil {
		return nil, err
	} else {
		value = val
	}
	// DEL ALL
	for _, key := range keys {
		f := c.path + gvar.New(key).String() + c.ext
		err = gfile.Remove(f)
		if err != nil {
			goto LOOP
		}
	}
LOOP:
	return value, err
}

// Update updates the value of <key> without changing its expiration and returns the old value.
// The returned value <exist> is false if the <key> does not exist in the cache.
//
// It deletes the <key> if given <value> is nil.
// It does nothing if <key> does not exist in the cache.
func (c *FileCache) Update(key interface{}, value interface{}) (oldValue interface{}, exist bool, err error) {
	f := c.path + gvar.New(key).String() + c.ext
	if have := gfile.IsFile(f); !have {
		// it does not exist.
		return
	}
	val := gfile.GetBytes(f)
	if len(val) < 8 || gvar.New(val[0:8]).Int64()-gtime.Timestamp() < 0 {
		return nil, false, gfile.Remove(f)
	}
	oldValue, err = c.Get(key)
	if err != nil {
		return nil, false, err
	}
	//UPDATE.
	err = c.Set(key, value, gvar.New(val[0:8]).Duration())

	return oldValue, true, err
}

// UpdateExpire updates the expiration of <key> and returns the old expiration duration value.
//
// It returns -1 and does nothing if the <key> does not exist in the cache.
// It deletes the <key> if <duration> < 0.
func (c *FileCache) UpdateExpire(key interface{}, duration time.Duration) (oldDuration time.Duration, err error) {
	f := c.path + gvar.New(key).String() + c.ext
	if have := gfile.IsFile(f); !have {
		// it does not exist.
		return -1, nil
	}
	val := gfile.GetBytes(f)
	if duration < 0 || gvar.New(val[0:8]).Int64()-gtime.Timestamp() < 0 {
		err = gfile.Remove(f)
		return
	}
	oldDuration = gvar.New(val[0:8]).Duration()
	// UPDATE
	err = c.Set(key, val[8:], duration)
	return
}

// Size returns the number of items in the cache.
func (c *FileCache) Size() (size int, err error) {
	list, err := gfile.ScanDirFile(c.path, "")
	return len(list), err
}

// Data returns a copy of all key-value pairs in the cache as map type.
// Note that this function may leads lots of memory usage, you can implement this function
// if necessary.
func (c *FileCache) Data() (map[interface{}]interface{}, error) {
	//容易引起读写瓶颈，不去实现
	return nil, nil
}

// Keys returns all keys in the cache as slice.
//容易引起读写瓶颈，慎用
func (c *FileCache) Keys() ([]interface{}, error) {
	keys, err := gfile.ScanDirFile(c.path, "")
	return gvar.New(keys).Slice(), err
}

// Values returns all values in the cache as slice.
func (c *FileCache) Values() ([]interface{}, error) {
	//容易引起读写瓶颈，不去实现
	return nil, nil
}

// Clear clears all data of the cache.
// Note that this function is sensitive and should be carefully used.
//小心使用，将删除所有
func (c *FileCache) Clear() error {
	return gfile.Remove(c.path)
}

// Close closes the cache if necessary.
func (c *FileCache) Close() error {
	//it does nothing
	return nil
}
