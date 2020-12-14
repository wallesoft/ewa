package base

import (
	"time"

	"github.com/gogf/gf/container/gvar"
	"github.com/gogf/gf/encoding/gjson"
	"github.com/gogf/gf/frame/g"
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
	Credentials       map[string]string
	EndPoint          string
	RequestPostMethod bool
	Client            *Client
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
	g.Dump("herhrehr")
	return at.requestToken()
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
	client := at.Client.ContentJson()
	g.Dump(client)
	var result string
	if at.RequestPostMethod {
		result = client.PostContent(at.Client.BaseUri+at.EndPoint, at.Credentials)
	} else {
		result = client.GetContent(at.Client.BaseUri+at.EndPoint, at.Credentials)
	}
	g.Dump(at.Credentials)
	v := gjson.New(result)
	// g.clinet request - content type json
	g.Dump("request token :", v.MustToJsonString())
	if have := v.Contains("errcode"); have {
		// err
		panic(v.MustToJsonString())
	}
	if have := v.Contains("component_access_token"); have {
		at.SetToken(v.GetString("component_access_token"), v.GetDuration("expires_in", 7200*time.Second))
		return v.GetString("component_access_token")
	} else {
		panic("Request access_token fail:" + v.MustToJsonString())
	}
	// parse to gjson

	// gjson contains()
	// err ???
}
