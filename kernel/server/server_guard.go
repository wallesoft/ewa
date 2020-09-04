package server

import (
	"encoding/json"
	"errors"

	"gitee.com/wallesoft/ewa/kernel/encryptor"
	ehttp "gitee.com/wallesoft/ewa/kernel/http"
	"github.com/gogf/gf/encoding/gjson"
	"github.com/gogf/gf/encoding/gxml"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/os/glog"
	"github.com/gogf/gf/text/gregex"
	"github.com/gogf/gf/util/gconv"
)

type ServerGuard struct {
	config Config
	//App			*Openplatform
	Request *ehttp.Request
	// Config         Config
	AlwaysValidate bool
	// Response *Response
	Logger    *glog.Logger
	Encryptor *encryptor.Encryptor
}

// type Config interface {
// 	Get(pattern string) interface{}
// }

// type Config struct {
// 	Appid  string `c:"app_id"`
// 	Secret string `c:"secret"`
// 	Token  string `c:"token"`
// 	AesKey string `c:"aes_key"`
// }
// func New(r Request, c Config, l *glog.Logger) *ServerGuard {
// 	g.Dump(r)
// 	encrypt, err := encryptor.New(map[string]interface{}{
// 		"AppId":     c.Get("app_id"),
// 		"Token":     c.Get("token"),
// 		"AesKey":    gconv.String(c.Get("aes_key")) + "=",
// 		"BlockSize": 32,
// 	})
// 	if err != nil {
// 		panic(err)
// 	}

// 	return &ServerGuard{
// 		Request:        r,
// 		Encryptor:      encrypt,
// 		Config:         c,
// 		Logger:         l,
// 		AlwaysValidate: false,
// 	}
// }
func New(config Config) *ServerGuard {
	return &ServerGuard{
		config: config,
	}
}
func (s *ServerGuard) SetLogger(logger *glog.Logger) {
	s.Logger = logger
}
func (s *ServerGuard) Serve() {
	//s.Logger.Debug
	//s.Logger.Debug(map[string]interface{}{"Request received": s.Request})
	s.Logger.Debug(map[string]interface{}{
		"Request Received": map[string]string{
			"uri":     s.Request.GetURL(),
			"content": gconv.String(s.Request.GetBody()),
		},
	})
	s.Validate().resolve()
}
func (s *ServerGuard) resolve() {
	message, _ := s.GetMessage()
	g.Dump(message)
}

//ParseMessage parse message from raw input.
func (s *ServerGuard) ParseMessage() (*gjson.Json, error) {
	//j, err := gjson.DecodeToJson(s.Request.GetRaw())
	content := s.Request.GetBody()
	mtype := checkDataType(content)
	if mtype == "xml" {
		//with out root 'xml'
		m, err := gxml.DecodeWithoutRoot(content)
		if err != nil {
			return nil, err
		}
		return gjson.New(m), nil
		// j, err := gjson.New(m)
		// if err != nil {
		// 	return nil, errors.New(fmt.Sprintf("Invalid message content: %s", err.Error()))
		// }
		// return j, nil
	}
	if mtype == "json" {
		return gjson.New(content), nil
		// j, err := gjson.New(content)
		// if err != nil {
		// 	return nil, errors.New(fmt.Sprintf("Invalid message content: %s", err.Error()))
		// }
		// return j, nil
	}

	return nil, errors.New("Invalid message content: unknow message type.")
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
	//token := gconv.String(s.Config.Get("token"))
	a := []string{s.config.Token, s.Request.GetString("timestamp"), s.Request.GetString("nonce")}
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
	if s.Request.GetString("signature") != s.signature() {
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
	return s.Request.GetString("Signature") != "" && s.Request.GetString("EncryptType") == "aes"
}

//DecryptMessage decrypt message
func (s *ServerGuard) DecryptMessage(message *gjson.Json) ([]byte, error) {
	//token := s.config.Token//gconv.String(s.Config.Get("token"))
	a := []string{s.config.Token, s.Request.GetString("Timestamp"), s.Request.GetString("Nonce"), message.GetString("Encrypt")}

	if message.GetString("msg_signature") != encryptor.Signature(a) {
		return nil, encryptor.NewError(encryptor.ERROR_INVALID_SIGNATURE, "Invalid Signature.")
	}
	content, err := s.Encryptor.Decrypt(message.GetBytes("Encrypt"))
	if err != nil {
		return nil, err
	}
	return content, nil
}
func checkDataType(content []byte) string {
	if json.Valid(content) {
		return "json"
	} else if gregex.IsMatch(`^<.+>[\S\s]+<.+>$`, content) {
		return "xml"
	}
	return ""
}
