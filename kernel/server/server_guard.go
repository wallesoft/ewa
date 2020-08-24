package kernel

import (
	"github.com/gogf/gf/container/gmap"
)

type ServerGuard struct {
	Request        *Request
	Config         *gmap.StrAnyMap
	AlwaysValidate bool
	// Response *Response
}

func (s *ServerGuard) ParseMessage() {

}

func (s *ServerGuard) signature() {
	token := s.Config.GetVar("token").String()
}

//Validate validate request source
func (s *ServerGuard) Validate() *ServerGuard {
	if !s.AlwaysValidate && !s.IsSafeMode() {
		return s
	}

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
