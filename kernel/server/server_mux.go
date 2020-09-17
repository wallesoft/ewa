package server

import (
	"sync"
)

//ServeMux server MUX
type ServeMux struct {
	mu sync.RWMutex
	m  map[int]MuxEntryGroup
}

//MuxEntryGroup is an alice
type MuxEntryGroup []muxEntry

//muxEntry
type muxEntry struct {
	h       Handler
	pattern string
}

var defaultServeMux ServeMux

//Default is the default Server MUX used by serve.
var DefaultServerMux = &defaultServeMux

func (mux *ServeMux) GetMuxEntryGroup(pattern string) MuxEntryGroup {

	// if group, ok := mux.m[pattern]; ok {
	// 	return group
	// }
	// return nil
}
