package server

import (
	"gitee.com/wallesoft/ewa/kernel/message"
)

const (
	EVENT_AUTHORIZED              = 16384
	EVENT_UNAUTHORIZED            = 32768
	EVENT_UPDATE_AUTHORIZED       = 65536
	EVENT_COMPONENT_VERIFY_TICKET = 131072
	EVENT_THIRD_FAST_REGISTERED   = 262144
)

var messageType = map[string]message.MessageType{
	"authorized":                 EVENT_AUTHORIZED,
	"unauthorized":               EVENT_UNAUTHORIZED,
	"updateauthorized":           EVENT_UPDATE_AUTHORIZED,
	"component_verify_ticket":    EVENT_COMPONENT_VERIFY_TICKET,
	"notify_third_fasteregister": EVENT_THIRD_FAST_REGISTERED,
}

type Message interface {
	GetAppId() string
	Type() string
	GetXmlString() string
	GetCreateTime() string //*
}
type BaseMessage struct {
	AppId      string
	InfoType   string
	CreateTime string
	XmlString  string
}

type TiketMessage struct {
	BaseMessage
	ComponentVerifyTicket string
}

type AutorizedMessage struct {
	BaseMessage
	AuthorizerAppid              string
	AuthorizationCode            string
	AuthorizationCodeExpiredTime string
	PreAuthCode                  string
}

type UnAuthorizedMessage struct {
	BaseMessage
	AuthorizerAppid string
}

type UpdateAuthorizedMessage struct {
	BaseMessage
	AuthorizerAppid              string
	AuthorizztionCode            string
	AuthoritationCodeExpiredTime string
	PreAuthCode                  string
}

// func GetMessageFromRequest(r Request) (m Message) {
// 	//验证来源
// 	//解密
// 	//解析类型
// 	//赋值到相应message
// 	return
// }
