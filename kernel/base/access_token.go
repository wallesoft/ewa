package base

import (
	"context"
	"time"

	"github.com/gogf/gf/v2/container/gvar"
	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/os/gcache"
)

//AccessToken
type AccessToken struct {
	Cache *gcache.Cache
	// Appid       string
	// Secret      string
	TokenKey          string //请求结果中
	RequestTokenKey   string //请求时的 获取token key 与请求时不一致时设置，例如第三方平台带小程序公众号请求
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
func (at *AccessToken) GetToken(ctx context.Context, refresh ...bool) string {
	if len(refresh) > 0 && refresh[0] {
		at = at.Refresh()
	}
	//cache refresh
	if token, err := at.Cache.Get(ctx, at.CacheKey); err != nil {
		panic(err.Error())
	} else {
		if token != nil && !at.isRefresh {
			return gvar.New(token).String()
		}
	}
	return at.requestToken(ctx)

}

//GetTokenKey
func (at *AccessToken) GetTokenKey() string {
	return at.TokenKey
}

//GetRequestTokenKey
func (at *AccessToken) GetRequestTokenKey() string {
	if at.RequestTokenKey == "" {
		return at.TokenKey
	}
	return at.RequestTokenKey
}

//Refresh
func (at *AccessToken) Refresh() *AccessToken {
	at.isRefresh = true
	return at
}

//SetToken
func (at *AccessToken) SetToken(ctx context.Context, token string, lifetime time.Duration) *AccessToken {
	if err := at.Cache.Set(ctx, at.CacheKey, token, lifetime); err != nil {
		panic(err.Error())
	}
	if _, err := at.Cache.Contains(ctx, at.CacheKey); err != nil {
		panic("Failed to cache access token.")
	}
	return at
}

func (at *AccessToken) requestToken(ctx context.Context) string {
	var v *gjson.Json
	if at.RequestPostMethod {
		v = at.Client.RequestJson(ctx, "POST", at.EndPoint, at.Credentials.Get(ctx))

	} else {
		v = at.Client.RequestJson(ctx, "GET", at.EndPoint, at.Credentials.Get(ctx))
	}

	if have := v.Contains(at.TokenKey); have {
		at.SetToken(ctx, v.Get(at.TokenKey).String(), v.Get("expires_in", 7200).Duration()*time.Second)
		return v.Get(at.TokenKey).String()
	} else {
		return ""
	}

}
