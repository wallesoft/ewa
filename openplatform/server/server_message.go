package server

import (
	"gitee.com/wallesoft/ewa/kernel/message"
)

const (
	EVENT_AUTHORIZED                  = 16384
	EVENT_UNAUTHORIZED                = 32768
	EVENT_UPDATE_AUTHORIZED           = 65536
	EVENT_COMPONENT_VERIFY_TICKET     = 131072
	EVENT_THIRD_FAST_REGISTERED       = 262144
	EVENT_THIRD_FAST_REGISTERDBETAAPP = 524288
)

var messageType = map[string]message.MessageType{
	"authorized":                       EVENT_AUTHORIZED,
	"unauthorized":                     EVENT_UNAUTHORIZED,
	"updateauthorized":                 EVENT_UPDATE_AUTHORIZED,
	"component_verify_ticket":          EVENT_COMPONENT_VERIFY_TICKET,
	"notify_third_fasteregister":       EVENT_THIRD_FAST_REGISTERED,
	"notify_third_fastregisterbetaapp": EVENT_THIRD_FAST_REGISTERDBETAAPP,
}
