package openplatform

import "sync"

type Handler interface {
	ServeMessage(*Request)
}

type HandlerFunc func(*Request)

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
	Request Request
}

//Handle registers the hanlder for the given pattern
func (mux *ServeMux) Handle(pattern string, handler Handler) (h Handler, patter string) {
	////
}
func (mux *ServeMux) HandleFunc(patter string, handler func(*Request)) {
	/////
}

func Handle(pattern string, handler Handler) {
	DefaultServeMux.Handle(pattern, handler)
}
func HandleFunc(pattern string, handler func(*Request)) {
	DefaultServeMux.HandleFunc(pattern, handler)
}

//Serve
func Serve(request *Request, handler Handler) error {

}

func (mux *ServeMux) ServeMessage(*Request) {

}
