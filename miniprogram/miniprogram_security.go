package miniprogram

import (
	"gitee.com/wallesoft/ewa/kernel/http"
	"github.com/gogf/gf/frame/g"
)

//CheckText @see https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/sec-check/security.msgSecCheck.html
func (mp *MiniProgram) CheckText(content string) *http.ResponseData {
	return &http.ResponseData{
		Json: mp.GetClientWithToken().RequestJson("POST", "wxa/msg_sec_check", g.Map{"content": content}),
	}
}
