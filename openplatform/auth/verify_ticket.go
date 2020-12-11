package auth

import (
	"time"

	"gitee.com/wallesoft/ewa/kernel/server"
	"github.com/gogf/gf/os/gcache"
)

type VerifyTicket interface {
	GetTicket() string
}

type DefaultVerifyTicket struct {
	// VerifyTicket string
	appid string
	cache *gcache.Cache
}

var defaultVerifyTicket = &DefaultVerifyTicket{}

func GetVerifyTicket(appid string, cache *gcache.Cache) *VerifyTicket {
	defaultVerifyTicket.appid = appid
	defaultVerifyTicket.cache = cache
	return defaultVerifyTicket
}

//Handle
func (v *defaultVerifyTicket) Handle(m *server.Message) interface{} {
	var verifyTicket string
	if have := m.Contains("ComponentVerifyTicket"); have {
		verifyTicket = m.GetString("ComponentVerifyTicket")
		if err := v.cache.Set(t.getKey(), verifyTicket, time.Second*3600); err != nil {
			panic(err.Error())
		}
	}
	return true
}

//GetTicket
func (v *defaultVerifyTicket) GetTicket() string {
	if ok := v.cache.Contains(v.getKey()); ok {
		return v.cache.GetString(v.getKey())
	}
}

func (v *defaultVerifyTicket) getKey() string {
	return "easywechat.open_platform.verify_ticket." + v.appid

}
