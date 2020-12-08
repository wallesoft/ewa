package server

import (
	gserver "gitee.com/wallesoft/ewa/kernel/server"
	"github.com/gogf/gf/frame/g"
)

type TicketHandler struct {
	responseText string
}

func (s *Server) initHandler() {
	s.Push(&TicketHandler{}, EVENT_COMPONENT_VERIFY_TICKET)

}
func (t *TicketHandler) Handle(m *gserver.Message) interface{} {
	g.Dump("defautl ticket handler")
	//缓存
	// return true
	return "ok"
}
