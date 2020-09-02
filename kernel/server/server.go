package server

import (
	"net/http"
)

type Server struct {
	*Context
	Writer  http.ResponseWriter
	Request *http.Request
}
