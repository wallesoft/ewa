package payment

import (
	"net/http"

	"github.com/gogf/gf/v2/container/gvar"
	"github.com/gogf/gf/v2/encoding/gjson"
)

type Response struct {
	Status     string
	StatusCode int
	Header     http.Header
	Body       []byte
}

type ResponseResult struct {
	*gjson.Json
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
