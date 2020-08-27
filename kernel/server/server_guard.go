package server

import (
	"errors"
	"fmt"

	"github.com/gogf/gf/encoding/gjson"
	"github.com/gogf/gf/os/glog"
)

type ServerGuard struct {
	//App			 *Openplatform
	Request        *Request
	Config         Config
	AlwaysValidate bool
	// Response *Response
	Logger *glog.Logger
}
type Config interface {
	Get(pattern string) interface{}
}

// type Config struct {
// 	Appid  string `c:"app_id"`
// 	Secret string `c:"secret"`
// 	Token  string `c:"token"`
// 	AesKey string `c:"aes_key"`
// }

func (s *ServerGuard) Serve() {
	//s.Logger.Debug
}

//ParseMessage parse message from raw input.
func (s *ServerGuard) ParseMessage() (*gjson.Json, error) {
	j, err := gjson.DecodeToJson(s.Request.RawBody)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("Invalid message content: %s", err.Error()))
	}
	return j, nil
}

func (s *ServerGuard) signature() string {
	// a := []string{s.Config.Token, s.Request.Timestamp, s.Request.Nonce}
	// // sort
	// sort.Strings(a)
	// return gsha1.Encrypt(strings.Join(a, ""))
	return ""
}

//Validate validate request source
func (s *ServerGuard) Validate() *ServerGuard {
	if !s.AlwaysValidate && !s.IsSafeMode() {
		return s
	}
	if s.Request.Signature != s.signature() {
		// response
	}
	return s
}

//ForceValidate set to force validation the request
func (s *ServerGuard) ForceValidate() *ServerGuard {
	s.AlwaysValidate = true
	return s
}

//IsSafeMode check the request message is the safe mode.
func (s *ServerGuard) IsSafeMode() bool {
	return s.Request.Signature != "" && s.Request.EncryptType != "aes"
}
