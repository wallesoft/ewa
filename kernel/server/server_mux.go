package server

import (
	"sync",
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
type MessageGroup  map[string]map[string]message.MessageType
//muxEntry
type muxEntry struct {
	h       Handler
	pattern message.MessageType
}

type serverMux struct {
	sync.RWMutex
	mux     Mux
	mesage  MessageGroup
	// messageType map[string]MessageGroup
}

func init() {
	serverMux.m = make(Mux)
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

func (s *ServerGuard) Push(handler Handler, pattern message.MessageType){
	if _,ok := serverMux.m[s.muxGroup]; ok {
		var me  muxEntry
		me.h = handler
		me.pattern = pattern
		serverMux.m[s.muxGroup] = append(serverMux.m[s.muxGroup],me)
	}
}
func (s *ServerGruad) setGroup(group string) {
	s.muxGroup = group
}
func (s *ServerGuard) InitMux(group string,messageGroup map[string]message.MessageType){
	s.setGroup(group)
	defer serverMux.Unlock()
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
