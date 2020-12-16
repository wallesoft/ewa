package base

import (
	"time"

	"github.com/gogf/gf/container/gvar"
	"github.com/gogf/gf/encoding/gjson"
	"github.com/gogf/gf/os/gcache"
)

//AccessToken
type AccessToken struct {
	Cache *gcache.Cache
	// Appid       string
	// Secret      string
	TokenKey          string
	isRefresh         bool
	CacheKey          string
	Credentials       TokenCredentail
	EndPoint          string
	RequestPostMethod bool
	Client            *Client
}
type Token struct {
}

//GetToken
func (at *AccessToken) GetToken(refresh ...bool) string {
	if len(refresh) > 0 && refresh[0] {
		at = at.Refresh()
	}
	//cache refresh
	if token, err := at.Cache.Get(at.CacheKey); err != nil {
		panic(err.Error())
	} else {
		if token != nil && !at.isRefresh {
			return gvar.New(token).String()
		}
	}
	return at.requestToken()

}
func (at *AccessToken) GetTokenKey() string {
	return at.TokenKey
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
	if _, err := at.Cache.Contains(at.CacheKey); err != nil {
		panic("Failed to cache access token.")
	}
	return at
}

func (at *AccessToken) requestToken() string {
	var v *gjson.Json
	if at.RequestPostMethod {
		v = at.Client.PostJson(at.EndPoint, at.Credentials.Get())

	} else {
		v = at.Client.GetJson(at.EndPoint, at.Credentials.Get())
	}

	if have := v.Contains(at.TokenKey); have {
		at.SetToken(v.GetString(at.TokenKey), v.GetDuration("expires_in", 7200)*time.Second)
		return v.GetString(at.TokenKey)
	} else {
		return ""
	}

}
