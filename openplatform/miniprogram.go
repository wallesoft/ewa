package openplatform

import "gitee.com/wallesoft/ewa/miniprogram"

//MiniProgram
type MiniProgram struct {
	*miniprogram.MiniProgram
	RefreshToken string // 刷新token
}
