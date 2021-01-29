package server

import (
	"gitee.com/wallesoft/ewa/openplatform/auth"
)

//初始化默认事件处理-verifyticket
func (s *Server) initHandler() {
	s.Push(auth.GetDefaultVerifyTicket(), EVENT_COMPONENT_VERIFY_TICKET)

}
