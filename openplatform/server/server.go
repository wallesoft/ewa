package server

import (
	"sync"

	"github.com/gogf/gf/container/gtype"
	"github.com/gogf/gf/os/glog"
)

type Handler interface {
	ServeMessage(Message) bool
}

type HandlerFunc func(Message) bool

func (f HandlerFunc) ServeMessage(m Message) bool {
	return f(m)
}

//ServeMux
type ServeMux struct {
	mu sync.RWMutex
	m  map[string]MuxEntryGroup
}

//MuxEntryGroup is an alice
type MuxEntryGroup []muxEntry

//NewServeMux returns a new ServeMux
func NewServeMux() *ServeMux { return new(ServeMux) }

//DefaultServeMux is the default ServerMux used by Serve.
var DefaultServeMux = &defaultServeMux
var defaultServeMux ServeMux

//muxEntry
type muxEntry struct {
	h       Handler
	pattern string
}

//Server
type Server struct {
	Request *Request
	Message Message
	//Config
	debug   *gtype.Bool
	logger  *glog.Logger
	Handler Handler
}

var defaultHandler = map[string]Handler{
	//"authorized":
	//"updateauthorized":
	//"unauthorized":
	"component_verify_ticket": &TicketHandler{},
}

//Handle registers the hanlder for the given pattern
func (mux *ServeMux) Handle(pattern string, handler Handler) (h Handler, patter string) {
	////
}

func (mux *ServeMux) HandleFunc(patter string, handler func(*Message)) {
	/////
}

// func (mux *ServeMux) Serve(message *Message, handler Handler) error {
// 	// config
// 	// descrypt
// 	// get type then find handler
// 	// handler
// 	// reponse return
// }

func (mux *ServeMux) ServeMessage(m Message) {
	mType := m.Type()
	g := mux.getMuxEntryGroup(mType)
	if len(g) > 0 {
		for _, entry := range g {
			if ok := entry.h.ServeMessage(m); !ok {
				goto LOOP
			}
		}
	} else {
		//default handler
		if entry, ok := defaultHandler[mType]; ok {
			entry.ServeMessage(m)
		}
	}
LOOP:
	//resopnose return

}
func (mux *ServeMux) getMuxEntryGroup(pattern string) MuxEntryGroup {
	if group, ok := mux.m[pattern]; ok {
		return group
	}
	return nil
}

func (s *Server) Serve() error {
	// type := s.Message.Type()
	if s.Handler == nil {
		DefaultServeMux.ServeMessage(s.Message)
	}
	s.Handler.ServeMessage(s.Message)
	return nil
}

func Handle(pattern string, handler Handler) {
	DefaultServeMux.Handle(pattern, handler)
}
func HandleFunc(pattern string, handler func(*Message)) {
	DefaultServeMux.HandleFunc(pattern, handler)
}

//Serve
func Serve(message Message, handler Handler) error {
	server := &Server{Message: message, Handler: handler}
	return server.Serve()
	//return DefaultServeMux.Serve(request, handler)
}
