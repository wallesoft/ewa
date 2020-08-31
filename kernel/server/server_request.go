package server

import (
	"github.com/gogf/gf/encoding/gjson"
	"github.com/gogf/gf/util/gconv"
)

//Request
type Request interface {
	Get(key string, def ...interface{}) interface{}
	GetRaw() []byte
	GetUrl() string
}

//Request abstract request.
type DefaultRequest struct {
	Signature    string
	Timestamp    string
	Nonce        string
	EncryptType  string
	MsgSignature string
	RawBody      []byte
	URL          string
}

//Get
func (r *DefaultRequest) Get(key string, def ...interface{}) interface{} {
	j := gjson.New(r)
	return j.Get(key, def)
}

//GetRaw get raw body from request
func (r *DefaultRequest) GetRaw() []byte {
	return r.RawBody
}

//GetUrl
func (r *DefaultRequest) GetUrl() string {
	return r.URL
}

//New return new DefaultRequest
func NewRequest(c map[string]interface{}) (*DefaultRequest, error) {
	r := &DefaultRequest{}
	if err := gconv.Struct(c, r); err != nil {
		return nil, err
	}
	return r, nil
}

// // func (r *Request) Config(m map[string]interface{}) r *Reqeust
// type Request interface{}{

// }
