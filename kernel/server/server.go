package server

import (
	"net/http"
)

type Server struct {
	*Config
	Writer  http.ResponseWriter
	Request *http.Request
}
