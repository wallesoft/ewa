package server

import (
	"github.com/gogf/gf/encoding/gjson"
)

//Message message is alias of gjson.Json
//openplatform 参考 https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/api/component_verify_ticket.html
//officialaccount 参考 https://developers.weixin.qq.com/doc/offiaccount/Message_Management/Receiving_standard_messages.html
//gjson.Json 参考 https://www.goframe.org/encoding/gjson/index
type Message struct {
	*gjson.Json
}

// var messageType = map[string]message.MessageType{
// 	"text":            message.TEXT,
// 	"image":           message.IMAGE,
// 	"voice":           message.VOICE,
// 	"video":           message.VIDEO,
// 	"shortvideo":      message.SHORT_VIDEO,
// 	"location":        message.LOCATION,
// 	"link":            message.LINK,
// 	"device_event":    message.DEVICE_EVENT,
// 	"device_text":     message.DEVICE_TEXT,
// 	"event":           message.EVENT,
// 	"file":            message.FILE,
// 	"miniprogrampage": message.MINIPROGRAM_PAGE,
// }
