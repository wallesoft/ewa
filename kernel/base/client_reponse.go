package base

import (
	"net/http"

	"github.com/gogf/gf/v2/container/gvar"
)

type Response struct {
	Status     string
	StatusCode int
	Header     http.Header
	Body       []byte
}

// func (r *Response) Raw() string {
// 	return r.ReadAllString()
// }

//ReadAll
func (r *Response) ReadAll() []byte {
	return r.Body
}

//ReadAllString
func (r *Response) ReadAllString() string {
	return gvar.New(r.Body).String()
}
