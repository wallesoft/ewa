package base

import (
	"github.com/gogf/gf/container/gvar"
	"github.com/gogf/gf/os/gcache"
)

//AccessToken
type AccessToken struct {
	Cache    *gcache.Cache
	Appid    string
	Secret   string
	Refresh  bool
	CacheKey string
}

//GetToken
func (at *AccessToken) GetToken() string {
	//cache refresh
	have, err := at.Cache.Contains(at.CacheKey)
	if err != nil {
		panic(err.Error())
	}
	if have && !at.Refresh {
		if token, err := at.Cache.Get(at.CacheKey); err != nil {
			panic(err.Error())
		} else {
			return gvar.New(token).String()
		}
	}
	//request

}

func (at *AccessToken) Refresh() *AccessToken {
	at.Refresh = true
}

func (at *AccessToken) SetToken(token string, expire ...int) *AccessToken {

}
