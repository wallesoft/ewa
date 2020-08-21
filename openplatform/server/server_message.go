package server

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

func GetMessageFromRequest(r Request) (m Message) {
	//验证来源
	//解密
	//解析类型
	//赋值到相应message
	return
}
