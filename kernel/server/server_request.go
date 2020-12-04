package server

import (
	"net/http"

	ehttp "gitee.com/wallesoft/ewa/kernel/http"
	"github.com/gogf/gf/util/gconv"
)

// import (
// 	"github.com/gogf/gf/encoding/gjson"
// 	"github.com/gogf/gf/util/gconv"
// )

// // //Request
// // type Request interface {
// // 	Get(key string, def ...interface{}) interface{}
// // 	GetRaw() []byte
// // 	GetUrl() string
// // }

// Request abstract request.
type Request struct {
	Signature    string
	Timestamp    string
	Nonce        string
	EncryptType  string
	MsgSignature string
	RawBody      []byte
	URL          string
}

// SetRequest
func (s *ServerGuard) SetRequest(r *http.Request) {
	eRequest := &ehttp.Request{Request: r}
	request := &Request{}
	if err := gconv.Struct(eRequest.GetQuery(), request); err != nil {
		// return nil,err
		panic(err)
	}
	request.RawBody = eRequest.GetBody()
	request.URL = eRequest.GetURL()
	s.Request = request
	// return r,nil
}

// // //Get
// // func (r *DefaultRequest) Get(key string, def ...interface{}) interface{} {
// // 	j := gjson.New(r)
// // 	return j.Get(key, def)
// // }

// // //GetRaw get raw body from request
// // func (r *DefaultRequest) GetRaw() []byte {
// // 	return r.RawBody
// // }

// // //GetUrl
// // func (r *DefaultRequest) GetUrl() string {
// // 	return r.URL
// // }

// // //New return new DefaultRequest
// // func NewRequest(c map[string]interface{}) (*DefaultRequest, error) {
// // 	r := &DefaultRequest{}
// // 	if err := gconv.Struct(c, r); err != nil {
// // 		return nil, err
// // 	}
// // 	return r, nil
// // }

// // // // func (r *Request) Config(m map[string]interface{}) r *Reqeust
// // // type Request interface{}{

// // // }
