package server

import (
	"encoding/json"
	"errors"
	"net/http"

	"gitee.com/wallesoft/ewa/kernel/encryptor"
	ehttp "gitee.com/wallesoft/ewa/kernel/http"
	"github.com/gogf/gf/encoding/gjson"
	"github.com/gogf/gf/encoding/gxml"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/os/glog"
	"github.com/gogf/gf/text/gregex"
	"github.com/gogf/gf/util/gconv"
	"github.com/gogf/gf/util/gutil"
)

type ServerGuard struct {
	config         Config
	Request        *ehttp.Request
	AlwaysValidate bool
	Response       *ehttp.Response
	Logger         *glog.Logger
	Encryptor      *encryptor.Encryptor
	muxGroup       string
	queryParam     *queryParam
	bodyData       *bodyData
}
type queryParam struct {
	Signature    string
	Timestamp    string
	Nonce        string
	EncryptType  string
	MsgSignature string
	// RawBody      []byte
	// URL string
}
type bodyData struct {
	RawBody []byte
}

// New
func New(config Config, request *http.Request, writer http.ResponseWriter) *ServerGuard {
	g := &ServerGuard{
		config: config,
	}
	g.setRequest(request)
	g.setResponse(writer)
	return g
}

// SetLogger
func (s *ServerGuard) SetLogger(logger *glog.Logger) {
	s.Logger = logger
}

// Serve
func (s *ServerGuard) Serve() {
	gutil.TryCatch(func() {
		s.parseRequest()
		// log
		s.Logger.Debug(map[string]interface{}{
			"Request Received": map[string]string{
				"uri":     s.Request.GetURL(),
				"content": gconv.String(s.bodyData.RawBody),
			},
		})
		s.Validate().resolve()
	}, func(exception interface{}) {
		switch exception {
		case ehttp.EXCEPTION_EXIT:
			return
		default:
			//LOG
			s.Logger.File("server_error_{Y-m-d}.log").Error(exception)
		}
	})

	//输出缓冲区
	s.Response.Output()
}
func (s *ServerGuard) resolve() {
	//handle Request
	s.handleRequest()

}
func (s *ServerGuard) parseRequest() {
	q := &queryParam{}

	if err := gconv.Struct(s.Request.GetQuery(), q); err != nil {
		//response
	}
	s.queryParam = q
	b := &bodyData{
		RawBody: s.Request.GetBody(),
	}
	s.bodyData = b
}

//return response
func (s *ServerGuard) handleRequest() {
	originMsg, err := s.GetMessage()
	g.Dump("err.....", err)
	g.Dump("msg", originMsg)

	if err != nil {
		panic(err.Error())
	}
	var mtype string
	if originMsg.Contains("MsgType") {
		mtype = originMsg.GetString("MsgType")
	} else if originMsg.Contains("msg_type") {
		mtype = originMsg.GetString("msg_type")
	} else {
		mtype = "text"
	}
	g.Dump(mtype)
	//处理相关信息类型，生成对应map，返回相关response
}

func (s *ServerGuard) dispatch(mtype string, message *Message) {
	// 1 mtype => message.MessageType

	// 2 Get Mux by group name
	// 3 range Mux
	// 4 diff

	// handlerGroup := s.mux.GetMuxEntryGroup(mtype)
	// if len(handlerGroup) > 0 {
	// 	for _, entry := range handlerGroup {
	// 		if ok := entry.h.ServeMesage(message); !ok {

	// 		}
	// 	}
	// }
	// // LOOP:
}

//ParseMessage parse message from raw input.
func (s *ServerGuard) parseMessage() (msg *Message, err error) {
	content := s.bodyData.RawBody
	g.Dump("content is :", content)
	mtype := checkDataType(content)
	g.Dump("type is :", mtype)
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
}
func (s *ServerGuard) parseXMLMessage(content []byte) (message *Message, err error) {
	undecrypted, err := gxml.DecodeWithoutRoot(content)
	g.Dump(undecrypted)
	if err != nil {
		return nil, err
	}
	if s.IsSafeMode() {
		if val, ok := undecrypted["Encrypt"]; ok {
			g.Dump("val", val)
			decrypted, err := s.decryptMessage(gconv.Bytes(val))
			g.Dump(decrypted)
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
	g.Dump("get", message, err)
	//is nil
	if message.IsNil() {
		s.Response.WriteStatusExit(http.StatusNoContent, "No message received")
		// panic(EXCEPTION_EXIT)
	}
	if err != nil {
		return nil, err
	}
	return
}
func (s *ServerGuard) signature() string {
	a := []string{s.config.Token, s.queryParam.Timestamp, s.queryParam.Nonce}
	return encryptor.Signature(a)
}

//Validate validate request source
func (s *ServerGuard) Validate() *ServerGuard {
	if !s.AlwaysValidate && !s.IsSafeMode() {
		return s
	}
	if s.queryParam.Signature != s.signature() {
		s.Response.WriteStatusExit(400, "Invalid request signature")
		// panic(EXCEPTION_EXIT)
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
	return s.queryParam.Signature != "" && s.queryParam.EncryptType == "aes"
}

//DecryptMessage decrypt message
func (s *ServerGuard) decryptMessage(message []byte) ([]byte, error) {
	//token := s.config.Token//gconv.String(s.Config.Get("token"))
	a := []string{s.config.Token, s.queryParam.Timestamp, s.queryParam.Nonce, gconv.String(message)}

	if s.queryParam.MsgSignature != encryptor.Signature(a) {
		return nil, encryptor.NewError(encryptor.ERROR_INVALID_SIGNATURE, "Invalid Signature.")
	}
	content, err := s.Encryptor.Decrypt(message)
	g.Dump("de", content, err)
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
