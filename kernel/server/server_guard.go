package kernel

import (
	"sort"
	"strings"

	"github.com/gogf/gf/container/gmap"
	"github.com/gogf/gf/crypto/gsha1"
	"github.com/gogf/gf/encoding/gjson"
	"github.com/gogf/gf/encoding/gxml"
)

type ServerGuard struct {
	Request        *Request
	Config         *gmap.StrAnyMap
	AlwaysValidate bool
	// Response *Response
}

//ParseMessage parse message from raw input.
func (s *ServerGuard) ParseMessage() {
	content := s.Request.RawBody
#########333afasdfaf
	// gjson LoadXml
	// gjson json
	if m, err := gxml.Decode(content); err != nil {
		// try decode json
		n, err := gjson.Decode(content)
		if err != nil {
			return nil, err
		}
		return n, nil
	} else {
		return m, nil
	}
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
