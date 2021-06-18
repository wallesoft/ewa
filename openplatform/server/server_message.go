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
	// EVENT_WEAPP_AUDIT_SUCCESS        = 8
	// EVENT_WEAPP_AUDIT_FAIL           = 16
	// EVENT_WEAPP_AUDIT_DELAY          = 32
	// EVENT_WXA_NICKNAME_AUDIT         = 64
)

var messageType = map[string]message.MessageType{
	"authorized":                       EVENT_AUTHORIZED,
	"unauthorized":                     EVENT_UNAUTHORIZED,
	"updateauthorized":                 EVENT_UPDATE_AUTHORIZED,
	"component_verify_ticket":          EVENT_COMPONENT_VERIFY_TICKET,
	"notify_third_fasteregister":       EVENT_THIRD_FAST_REGISTER,
	"notify_third_fastregisterbetaapp": EVENT_THIRD_FAST_REGISTERBETAAPP,
	"notify_third_fastverifybetaapp":   EVENT_THIRD_FASTVERIFYBETAAPP,
	// "weapp_audit_success":              EVENT_WEAPP_AUDIT_SUCCESS,
	// "weapp_audit_fail":                 EVENT_WEAPP_AUDIT_FAIL,
	// "weapp_audit_delay":                EVENT_WEAPP_AUDIT_DELAY,
	// "wxa_nickname_audit":               EVENT_WXA_NICKNAME_AUDIT,
}
