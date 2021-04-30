package message

//MessageType 注意int为2的n次方
type MessageType = int

// //Message type
const (
	TEXT        = 2
	IMAGE       = 4
	VOICE       = 8
	VIDEO       = 16
	SHORT_VIDEO = 32
	LOCATION    = 64
	LINK        = 128
	// DEVICE_EVENT       = 256
	// DEVICE_TEXT        = 512
	FILE = 1024
	// TEXT_CARD          = 2048 //预留
	// TRANSFER           = 4096
	EVENT            = 8192
	MINIPROGRAM_PAGE = 2097152
	// MINIPROGRAM_NOTICE = 4194304 //企业微信小程序通知消息
)

var DefaultMessage = map[string]MessageType{
	"text":            TEXT,
	"image":           IMAGE,
	"voice":           VOICE,
	"video":           VIDEO,
	"file":            FILE,
	"shortvideo":      SHORT_VIDEO,
	"location":        LOCATION,
	"link":            LINK,
	"event":           EVENT,
	"miniprogrampage": MINIPROGRAM_PAGE,
}
