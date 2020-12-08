package server

import (
	"net/http"

	guard "gitee.com/wallesoft/ewa/kernel/server"
	// "gitee.com/wallesoft/ewa/kernel/"
	"github.com/gogf/gf/container/gtype"
	"github.com/gogf/gf/os/glog"
)

// type Handler interface {
// 	ServeMessage(Message) bool
// }

// type HandlerFunc func(Message) bool

// func (f HandlerFunc) ServeMessage(m Message) bool {
// 	return f(m)
// }

// //ServeMux
// type ServeMux struct {
// 	mu sync.RWMutex
// 	m  map[string]MuxEntryGroup
// }

// //MuxEntryGroup is an alice
// type MuxEntryGroup []muxEntry

// //NewServeMux returns a new ServeMux
// func NewServeMux() *ServeMux { return new(ServeMux) }

// //DefaultServeMux is the default ServerMux used by Serve.
// var DefaultServeMux = &defaultServeMux
// var defaultServeMux ServeMux

// //muxEntry
// type muxEntry struct {
// 	h       Handler
// 	pattern string
// }

//Server
type Server struct {
	//App   *OpenPlatform
	*guard.ServerGuard
	//Request *Request
	// Message Message
	//Config
	debug  *gtype.Bool
	logger *glog.Logger
	// Handler Handler
	// MuxGroup string
}

const (
	MUX_GROUP = "openplatform" // default config group name
)

// var defaultHandler = map[string]Handler{
// 	//"authorized":
// 	//"updateauthorized":
// 	//"unauthorized":
// 	"component_verify_ticket": &TicketHandler{},
// }

// //Handle registers the hanlder for the given pattern
// func (mux *ServeMux) Handle(pattern string, handler Handler) (h Handler, patter string) {
// 	////
// 	return nil, ""
// }

// func (mux *ServeMux) HandleFunc(patter string, handler func(*Message)) {
// 	/////
// }

// func (mux *ServeMux) Serve(message *Message, handler Handler) error {
// 	// config
// 	// descrypt
// 	// get type then find handler
// 	// handler
// 	// reponse return
// }
// func (s *Server) Push(handler server.Handler) {

// }
func (s *Server) SetMux() {
	// s.MuxGroup = MUX_GROUP
	//init handler
	s.InitMux(MUX_GROUP, messageType)
	s.initHandler()
}

//Resolve
func (s *Server) Resolve() {
	if msg, err := s.GetMessage(); err == nil {
		var t string
		if msg.Contains("InfoType") {
			t = msg.GetString("InfoType")
		} else {
			s.Response.WriteStatusExit(http.StatusBadRequest, "Invalid message info type")
		}
		s.Dispatch(t, msg)
	} else {
		panic(err.Error())
	}
}

// -------------------------------
// func (mux *ServeMux) ServeMessage(m Message) {
// 	mType := m.Type()
// 	g := mux.getMuxEntryGroup(mType)
// 	if len(g) > 0 {
// 		for _, entry := range g {
// 			if ok := entry.h.ServeMessage(m); !ok {
// 				goto LOOP
// 			}
// 		}
// 	} else {
// 		//default handler
// 		if entry, ok := defaultHandler[mType]; ok {
// 			entry.ServeMessage(m)
// 		}
// 	}
// LOOP:
// 	//resopnose return

// }
// // ---------------------------
// func (mux *ServeMux) getMuxEntryGroup(pattern string) MuxEntryGroup {
// 	if group, ok := mux.m[pattern]; ok {
// 		return group
// 	}
// 	return nil
// }

// -------------------------------------------
// func (s *Server) Serve() Response {
// 	// type := s.Message.Type()
// 	s.ServerGuard.Serve()
// 	if s.Handler == nil {
// 		DefaultServeMux.ServeMessage(s.Message)
// 	}
// 	s.Handler.ServeMessage(s.Message)
// 	response := Response{}
// 	return response
// }
// -------------------------------------
// func Handle(pattern string, handler Handler) {
// 	DefaultServeMux.Handle(pattern, handler)
// }
// func HandleFunc(pattern string, handler func(*Message)) {
// 	DefaultServeMux.HandleFunc(pattern, handler)
// }

//Serve
// func Serve(message Message, handler Handler) Response {
// 	server := &Server{Message: message, Handler: handler}
// 	return server.Serve()
// 	//return DefaultServeMux.Serve(request, handler)
// }
