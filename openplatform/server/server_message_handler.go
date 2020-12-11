package server

import (
	"time"

	gserver "gitee.com/wallesoft/ewa/kernel/server"
	"github.com/gogf/gf/container/gvar"
	"github.com/gogf/gf/os/gcache"
)

type VerifyTicket struct {
	// responseText string
	server *Server
	cache  *gcache.Cache
}

// var verifyticket = s.VerifyTicket
func (s *Server) initHandler() {
	s.Push(s.VerifyTicket(), EVENT_COMPONENT_VERIFY_TICKET)

}

// Handle handle the message
func (t *VerifyTicket) Handle(m *gserver.Message) interface{} {
	var verifyTicket string
	if have := m.Contains("ComponentVerifyTicket"); have {
		verifyTicket = m.GetString("ComponentVerifyTicket")
		if err := t.cache.Set(t.getKey(), verifyTicket, time.Second*3600); err != nil {
			panic(err.Error())
		}
	}
	return true
}

//verify ticket return
func (s *Server) VerifyTicket() *VerifyTicket {

	vt := &VerifyTicket{
		server: s,
		cache:  s.Cache,
	}
	return vt
}
func (t *VerifyTicket) getKey() string {
	return "easywechat.open_platform.verify_ticket." + t.server.Config.AppID
}

//GetTicket
func (t *VerifyTicket) GetTicket() string {
	ticket, err := t.cache.Get(t.getKey())
	if err != nil {
		panic(err.Error())
	}
	return gvar.New(ticket).String()
}
