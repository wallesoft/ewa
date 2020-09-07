package server

import (
	"sync"
)

//ServerMux server MUX
type ServerMux struct {
	mu sync.RWMutex
	m  map[string]MuxEntryGroup
}

//MuxEcntryGroup is an alice
type MuxEntryGroup []muxEntry

//muxEntry
type muxEntry struct {
	h       Handler
	pattern string
}

var defaultServerMux ServerMux

//Default is the default Server MUX used by serve.
var DefaultServerMux = &defaultServerMux

func (mux *ServeMux) GetMuxEntryGroup(pattern string) MuxEntryGroup {
	if group, ok := mux.m[pattern]; ok {
		return group
	}
	return nil
}
