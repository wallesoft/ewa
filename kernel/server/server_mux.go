package server

import (
	"fmt"

	"gitee.com/wallesoft/ewa/kernel/message"
	"github.com/gogf/gf/container/garray"
)

//muxEntry
type muxEntry struct {
	Handler   Handler
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
	var me muxEntry
	me.Handler = handler
	me.Condition = pattern
	// add to slice head
	s.MuxEntry.InsertBefore(0, me)
}

//GetHandlers
func (s *ServerGuard) GetHandlers() *garray.Array {
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
