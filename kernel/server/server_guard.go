package server

import (
	"errors"
	"fmt"

	"gitee.com/wallesoft/ewa/kernel/encryptor"
	"github.com/gogf/gf/encoding/gjson"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/os/glog"
	"github.com/gogf/gf/util/gconv"
)

type ServerGuard struct {
	//App			*Openplatform
	Request        Request
	Config         Config
	AlwaysValidate bool
	// Response *Response
	Logger    *glog.Logger
	Encryptor *encryptor.Encryptor
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
func New() *ServerGuard {

}

func (s *ServerGuard) Serve() {
	//s.Logger.Debug
	//s.Logger.Debug(map[string]interface{}{"Request received": s.Request})
	s.Logger.Debug(map[string]interface{}{
		"Request Received": map[string]string{
			"uri":     s.Request.GetUrl(),
			"content": gconv.String(s.Request.GetRaw()),
		},
	})
	s.Validate().resolve()
}
func (s *ServerGuard) resolve() {
	message, err := s.GetMessage()
	g.Dump(message)
	g.Dump(err)
}

//ParseMessage parse message from raw input.
func (s *ServerGuard) ParseMessage() (*gjson.Json, error) {
	j, err := gjson.DecodeToJson(s.Request.GetRaw())
	if err != nil {
		return nil, errors.New(fmt.Sprintf("Invalid message content: %s", err.Error()))
	}
	return j, nil
}

//GetMessage
func (s *ServerGuard) GetMessage() (*gjson.Json, error) {
	message, err := s.ParseMessage()
	if err != nil {
		return nil, err
	}
	if s.IsSafeMode() && message.Contains("Encrypt") {
		//decrypt
		msg, err := s.DecryptMessage(message)
		if err != nil {
			return nil, err
		}
		j, err := gjson.DecodeToJson(msg)
		if err != nil {
			return nil, err
		}
		return j, nil
	}
	return message.GetJson("Encrypt"), nil
}
func (s *ServerGuard) signature() string {
	token := gconv.String(s.Config.Get("token"))
	a := []string{token, gconv.String(s.Request.Get("timestamp")), gconv.String(s.Request.Get("nonce"))}
	// sort
	return encryptor.Signature(a)
	// sort.Strings(a)
	// return gsha1.Encrypt(strings.Join(a, ""))
}

//Validate validate request source
func (s *ServerGuard) Validate() *ServerGuard {
	if !s.AlwaysValidate && !s.IsSafeMode() {
		return s
	}
	if gconv.String(s.Request.Get("signature")) != s.signature() {
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
	return gconv.String(s.Request.Get("signature")) != "" && s.Request.EncryptType != "aes"
}

//DecryptMessage decrypt message
func (s *ServerGuard) DecryptMessage(message *gjson.Json) ([]byte, error) {
	token := gconv.String(s.Config.Get("token"))
	a := []string{token, gconv.String(s.Request.Get("Timestamp")), gconv.String(s.Request.Get("Nonce")), message.GetString("Encrypt")}

	if message.GetString("msg_signature") != encryptor.Signature(a) {
		return nil, encryptor.NewError(encryptor.ERROR_INVALID_SIGNATURE, "Invalid Signature.")
	}
	content, err := s.Encryptor.Decrypt(message.GetBytes("Encrypt"))
	if err != nil {
		return nil, err
	}
	return content, nil
}
