package server

import (
	"net/http"

	ehttp "gitee.com/wallesoft/ewa/kernel/http"
)

// SetRequest
func (s *ServerGuard) setRequest(r *http.Request) {
	s.Request = &ehttp.Request{Request: r}
}
