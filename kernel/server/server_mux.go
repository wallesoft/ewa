package server

import (
	"fmt"

	"gitee.com/wallesoft/ewa/kernel/message"
)

//ServeMux server MUX
// type ServeMux struct {
// 	sync.RWMutex
// 	m  map[string]MuxEntryGroup
// 	group string
// }
type Mux map[string]MuxEntryGroup

//MuxEntryGroup is an alice
type MuxEntryGroup []muxEntry
type MessageGroup map[string]map[string]message.MessageType

//muxEntry
type muxEntry struct {
	Handler   Handler
	Condition message.MessageType
}

var serverMux struct {
	// sync.RWMutex
	mux     Mux
	message MessageGroup
	// messageType map[string]MessageGroup
}

func init() {
	serverMux.mux = make(Mux)
	serverMux.message = make(MessageGroup)
}

// func setMux(mux Mux) {
// 	defer serverMux.Unlock()
// 	serverMux.m = mux
// }

// func setMuxGroup(group string, mgroup MuxEntryGroup) {
// 	defer serverMux.Unlock()
// 	serverMux.m[group] = mgroup
// }

// // AddMuxEntry
// func AddMuxEntry(group string, entry muxEntry) {
// 	defer serverMux.Unlock()
// 	serverMux.m[group] = append(serverMux.m[group], entry)
// }

func (s *ServerGuard) Push(handler Handler, pattern message.MessageType) {
	// if _, ok := serverMux.mux[s.muxGroup]; ok {
	var me muxEntry
	me.Handler = handler
	me.Condition = pattern
	serverMux.mux[s.muxGroup] = append([]muxEntry{me}, serverMux.mux[s.muxGroup]...)
	// } else {
	// serveMux.mux[s.muxGroup] =
	// }
}

func (s *ServerGuard) setGroup(group string) {
	s.muxGroup = group
}

//GetHandlers
func (s *ServerGuard) GetHandlers() MuxEntryGroup {
	if group, ok := serverMux.mux[s.muxGroup]; ok {
		return group
	} else {
		panic("No mux group")
	}
}

//TypeToEvent
func (s *ServerGuard) TypeToEvent(t string) message.MessageType {
	if mgroup, ok := serverMux.message[s.muxGroup]; ok {
		if event, ok := mgroup[t]; ok {
			return event
		}
	}
	panic(fmt.Sprintf("Invalid message type: %s", t))
}
func (s *ServerGuard) InitMux(group string, messageGroup map[string]message.MessageType) {
	s.setGroup(group)
	// defer serverMux.Unlock()
	serverMux.message[group] = messageGroup
}

// // Push
// func Push(handler Handler, mtype message.Messagetype){

// }
// ------------------------------------------------
// var defaultServeMux ServeMux

// //Default is the default Server MUX used by serve.
// var DefaultServerMux = &defaultServeMux

// func (mux *ServeMux) Notify(event string, message *Message) {

// }
// ---------------------------------------------------
// func (mux *ServeMux) GetMuxEntryGroup(pattern string) MuxEntryGroup {

// 	// if group, ok := mux.m[pattern]; ok {
// 	// 	return group
// 	// }
// 	// return nil
// }
