package server

import (
	"encoding/json"
	"errors"

	"gitee.com/wallesoft/ewa/kernel/encryptor"
	ehttp "gitee.com/wallesoft/ewa/kernel/http"
	"github.com/gogf/gf/encoding/gjson"
	"github.com/gogf/gf/encoding/gxml"
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
	mux       *ServeMux
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
	// message, _ := s.GetMessage()
	// g.Dump(message)
}

//return response
func (s *ServerGuard) handleRequst() {
	originMsg, err := s.GetMessage()
	if err != nil {
		//
	}

	var mtype string
	if originMsg.Contains("MsgType") {
		mtype = originMsg.GetString("MsgType")
	} else if originMsg.Contains("msg_type") {
		mtype = originMsg.GetString("msg_type")
	} else {
		mtype = "text"
	}

	//处理相关信息类型，生成对应map，返回相关response
}

func (s *ServerGuard) dispatch(mtype string, message *Message) {
	handlerGroup := s.mux.GetMuxEntryGroup(mtype)
	if len(handlerGroup) > 0 {
		for _, entry := range handlerGroup {
			if ok := entry.h.ServeMesage(message); !ok {

			}
		}
	}
	// LOOP:
}

//ParseMessage parse message from raw input.
func (s *ServerGuard) parseMessage() (msg *Message, err error) {
	content := s.Request.GetBody()
	mtype := checkDataType(content)
	switch mtype {
	case "xml":
		msg, err = s.parseXMLMessage(content)
		return
	case "json":
		msg, err = s.parseJSONMessage(content)
		return
	default:
		return nil, errors.New("invalid message content: unsupported message type")
	}
	// var err error
	// var msg *gjson.Json
	// content := s.Request.GetBody()
	// mtype := checkDataType(content)
	// switch mtype {
	// case "xml":
	// 	m, xerr := gxml.DecodeWithoutRoot(content)
	// 	if xerr != nil {
	// 		err = xerr
	// 	}
	// 	msg = gjson.New(m)
	// case "json":
	// 	msg = gjson.New(content)
	// default:
	// 	//msg = nil
	// 	err = errors.New("invalid message content: unsupported message type")
	// }
	// // if mtype == "xml" {
	// // 	//with out root 'xml'
	// // 	m, err := gxml.DecodeWithoutRoot(content)
	// // 	if err != nil {
	// // 		return nil, err
	// // 	}
	// // 	msg := gjson.New(m)
	// // }
	// // if mtype == "json" {
	// // 	msg := gjson.New(content)
	// // }
	// if err != nil {
	// 	return nil, err
	// }

	// if msg != nil {
	// 	if s.IsSafeMode() && msg.Contains("Encrypt") {
	// 		//decrypt
	// 		//msg, err := s.DecryptMessage(msg)
	// 		decrypted, decrypterr := s.decryptMessage(msg)
	// 		if decrypterr != nil {
	// 			return nil, decrypterr
	// 		}

	// 		if j, err := gjson.DecodeToJson(decrypted); err == nil {
	// 			return &Message{
	// 				Json: j,
	// 			}, nil
	// 		} else {
	// 			return nil, err
	// 		}

	// 		// }

	// 		// if err != nil {
	// 		// 	return nil, err
	// 		// }
	// 		// j, err := gjson.DecodeToJson(msg)
	// 		// if err != nil {
	// 		// 	return nil, err
	// 		// }
	// 		// return j, nil
	// 	}
	// 	return &Message{Json: msg}, nil
	// }
	//return nil, errors.New("invaild message content")
}
func (s *ServerGuard) parseXMLMessage(content []byte) (message *Message, err error) {
	undecrypted, err := gxml.DecodeWithoutRoot(content)
	if err != nil {
		return nil, err
	}
	if s.IsSafeMode() {
		if val, ok := undecrypted["Encrypt"]; ok {
			decrypted, err := s.decryptMessage(gconv.Bytes(val))
			if err != nil {
				return nil, err
			}
			//out root
			m, err := gxml.DecodeWithoutRoot(decrypted)
			if err != nil {
				return nil, err
			}
			message = &Message{
				Json: gjson.New(m),
			}
			return message, nil
		}
		return nil, errors.New("invalid parse message type of xml: get encrypt content error")
	}
	message = &Message{
		Json: gjson.New(undecrypted),
	}
	return message, nil
}
func (s *ServerGuard) parseJSONMessage(content []byte) (message *Message, err error) {
	j, err := gjson.LoadContent(content)
	if err != nil {
		return nil, err
	}
	if s.IsSafeMode() && j.Contains("Encrypt") {
		decrypted, err := s.decryptMessage(j.GetBytes("Encrypt"))
		if err != nil {
			return nil, err
		}
		message = &Message{
			Json: gjson.New(decrypted),
		}
		return message, nil
	}
	return &Message{
		Json: j,
	}, nil
}

//GetMessage
func (s *ServerGuard) GetMessage() (message *Message, err error) {
	message, err = s.parseMessage()
	//is nil
	if message.IsNil() {
		return nil, errors.New("No message received.")
	}
	if err != nil {
		return nil, err
	}
	return
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
func (s *ServerGuard) decryptMessage(message []byte) ([]byte, error) {
	//token := s.config.Token//gconv.String(s.Config.Get("token"))
	a := []string{s.config.Token, s.Request.GetString("Timestamp"), s.Request.GetString("Nonce"), gconv.String(message)}

	if s.Request.GetString("msg_signature") != encryptor.Signature(a) {
		return nil, encryptor.NewError(encryptor.ERROR_INVALID_SIGNATURE, "Invalid Signature.")
	}
	content, err := s.Encryptor.Decrypt(message)
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
