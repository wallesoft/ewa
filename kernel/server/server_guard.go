package kernel

import (
	"errors"
	"fmt"
	"sort"
	"strings"

	"github.com/gogf/gf/container/gmap"
	"github.com/gogf/gf/crypto/gsha1"
	"github.com/gogf/gf/encoding/gjson"
	"github.com/gogf/gf/os/glog"
)

type ServerGuard struct {
	//App			 *Openplatform
	Request        *Request
	Config         *gmap.StrAnyMap
	AlwaysValidate bool
	// Response *Response
	Logger *glog.Logger
}

func (s *ServerGuard) Serve() {

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
	token := s.Config.GetVar("token").String()
	a := []string{token, s.Request.Timestamp, s.Request.Nonce}
	// sort
	sort.Strings(a)
	return gsha1.Encrypt(strings.Join(a, ""))
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
