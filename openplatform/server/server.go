package server

import (
	"net/http"

	guard "gitee.com/wallesoft/ewa/kernel/server"

	"github.com/gogf/gf/v2/container/gtype"
	"github.com/gogf/gf/v2/os/glog"
)

//Server
type Server struct {
	*guard.ServerGuard
	debug  *gtype.Bool  //@deprecated
	logger *glog.Logger //@deprecated
}

const (
	MUX_GROUP = "openplatform" // default config group name
)

func (s *Server) SetMux() {
	//init handler
	s.RegisterMessageType(messageType)
	s.initHandler()
}

//Resolve
func (s *Server) Resolve(msg *guard.Message) bool {

	if msg != nil {
		var t string
		if msg.Contains("InfoType") {
			t = msg.Get("InfoType").String()
		} else if msg.Contains("MsgType") {
			t = msg.Get("MsgType").String()
		} else {
			s.Response.WriteStatusExit(http.StatusBadRequest, "Invalid message type")
		}
		s.Dispatch(t, msg)
		s.Response.Write(guard.SUCCESS_EMPTY_RESPONSE)
	}
	return true
}

//Should return raw response
func (s *Server) ShouldReturnRawResponse() bool {
	return true
}
