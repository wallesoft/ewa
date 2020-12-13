package base

import (
	"time"

	"github.com/gogf/gf/container/gvar"
	"github.com/gogf/gf/os/gcache"
)

//AccessToken
type AccessToken struct {
	Cache *gcache.Cache
	// Appid       string
	// Secret      string
	isRefresh   bool
	CacheKey    string
	Credentials map[string]string
}
type Token struct {
}

//GetToken
func (at *AccessToken) GetToken() string {
	//cache refresh
	have, err := at.Cache.Contains(at.CacheKey)
	if err != nil {
		panic(err.Error())
	}
	if have && !at.isRefresh {
		if token, err := at.Cache.Get(at.CacheKey); err != nil {
			panic(err.Error())
		} else {
			return gvar.New(token).String()
		}
	}
	//request

}

//Refresh
func (at *AccessToken) Refresh() *AccessToken {
	at.isRefresh = true
	return at
}

//SetToken
func (at *AccessToken) SetToken(token string, lifetime time.Duration) *AccessToken {
	if err := at.Cache.Set(at.CacheKey, token, lifetime); err != nil {
		panic(err.Error())
	}
	if have, err := at.Cache.Contains(at.CacheKey); err != nil {
		panic("Failed to cache access token.")
	}
	return at
}

func (at *AccessToken) requestToken() {
	// g.clinet request - content type json
	// parse to gjson
	// gjson contains()
	// err ???
}
