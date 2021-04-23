package server

import (
	"gitee.com/wallesoft/ewa/kernel/message"
)

const (
	EVENT_AUTHORIZED                 = 16384
	EVENT_UNAUTHORIZED               = 32768
	EVENT_UPDATE_AUTHORIZED          = 65536
	EVENT_COMPONENT_VERIFY_TICKET    = 131072
	EVENT_THIRD_FAST_REGISTER        = 262144
	EVENT_THIRD_FAST_REGISTERBETAAPP = 524288
	EVENT_THIRD_FASTVERIFYBETAAPP    = 1048576
)

var messageType = map[string]message.MessageType{
	"authorized":                       EVENT_AUTHORIZED,
	"unauthorized":                     EVENT_UNAUTHORIZED,
	"updateauthorized":                 EVENT_UPDATE_AUTHORIZED,
	"component_verify_ticket":          EVENT_COMPONENT_VERIFY_TICKET,
	"notify_third_fasteregister":       EVENT_THIRD_FAST_REGISTER,
	"notify_third_fastregisterbetaapp": EVENT_THIRD_FAST_REGISTERBETAAPP,
	"notify_third_fastverifybetaapp":   EVENT_THIRD_FASTVERIFYBETAAPP,
}
