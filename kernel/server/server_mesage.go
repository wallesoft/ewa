package server

import "github.com/gogf/gf/encoding/gjson"

//Message message is alias of gjson.Json
//openplatform 参考 https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/api/component_verify_ticket.html
// officialaccount 参考 https://developers.weixin.qq.com/doc/offiaccount/Message_Management/Receiving_standard_messages.html
//gjson.Json 参考 https://www.goframe.org/encoding/gjson/index
type Message struct {
	*gjson.Json
}
