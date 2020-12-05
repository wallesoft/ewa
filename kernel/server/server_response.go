package server

import (
	"net/http"

	ehttp "gitee.com/wallesoft/ewa/kernel/http"
)

func (s *ServerGuard) setResponse(w http.ResponseWriter) {
	s.Response = ehttp.GetResponse(w)
}
