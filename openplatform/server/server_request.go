package server

import (
	"gitee.com/wallesoft/ewa/kernel/server"
)

//Request

//CreateRequest  create request
func NewRequest() (request *server.Request) {
	request = &server.Request{}
	return
}
