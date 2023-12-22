package auth

import (
	"context"
	"time"

	"gitee.com/wallesoft/ewa/kernel/server"
	"github.com/gogf/gf/v2/container/gvar"
	"github.com/gogf/gf/v2/os/gcache"
)

type VerifyTicket interface {
	GetTicket(ctx context.Context) string
}

type DefaultVerifyTicket struct {
	// VerifyTicket string
	appid string
	cache *gcache.Cache
}

var defaultVerifyTicket = &DefaultVerifyTicket{}

func GetVerifyTicket(appid string, cache *gcache.Cache) VerifyTicket {
	defaultVerifyTicket.appid = appid
	defaultVerifyTicket.cache = cache
	return defaultVerifyTicket
}

func GetDefaultVerifyTicket() *DefaultVerifyTicket {
	return defaultVerifyTicket
}

// Handle
func (v *DefaultVerifyTicket) Handle(ctx context.Context, m *server.Message) interface{} {
	var verifyTicket string
	if have := m.Contains("ComponentVerifyTicket"); have {
		verifyTicket = m.Get("ComponentVerifyTicket").String()
		if err := v.cache.Set(ctx, v.getKey(), verifyTicket, time.Second*3600); err != nil {
			panic(err.Error())
		}
	}
	return true
}

// GetTicket
func (v *DefaultVerifyTicket) GetTicket(ctx context.Context) string {
	ticket, err := v.cache.Get(ctx, v.getKey())
	if err != nil {
		panic(err.Error())
	}
	return gvar.New(ticket).String()
}

// getKey函数用于获取DefaultVerifyTicket对象的key值
func (v *DefaultVerifyTicket) getKey() string {
	return "ewawechat.open_platform.verify_ticket." + v.appid
}
