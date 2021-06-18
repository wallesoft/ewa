package server

import (
	"fmt"

	"gitee.com/wallesoft/ewa/kernel/message"
	"github.com/gogf/gf/container/garray"
	"github.com/gogf/gf/container/gmap"
)

//muxEntry
type muxEntry struct {
	Handlers  *garray.Array
	Condition message.MessageType
}

//HandlerFunc
type HandlerFunc func(*Message) interface{}

//Handle by func
func (h HandlerFunc) Handle(m *Message) interface{} {
	return h(m)
}

//PushFunc
func (s *ServerGuard) PushFunc(handler HandlerFunc, pattern message.MessageType) {
	s.Push(handler, pattern)
}

//Push
func (s *ServerGuard) Push(handler Handler, pattern message.MessageType) {
	//var me muxEntry
	// var h *garray.Array
	h := garray.New()
	var err error
	// var me *gmap.IntAnyMap
	if s.MuxEntry.Contains(pattern) {
		handlers := s.MuxEntry.Get(pattern)
		h = handlers.(*garray.Array)
	}

	if h.Len() == 0 {
		h.Append(handler)
	} else {
		err = h.InsertBefore(0, handler)
	}
	s.MuxEntry.Set(pattern, h)

	if err != nil {
		panic(err)
	}
}

//GetHandlers
func (s *ServerGuard) GetHandlers() *gmap.IntAnyMap {
	return s.MuxEntry
}

//TypeToEvent
func (s *ServerGuard) TypeToEvent(t string) message.MessageType {
	if s.MessageGroup.Contains(t) {
		return s.MessageGroup.Get(t)
	}
	panic(fmt.Sprintf("Invalid message type: %s", t))
}

//Register message type
func (s *ServerGuard) RegisterMessageType(message map[string]message.MessageType) {
	s.MessageGroup.Sets(message)
}
