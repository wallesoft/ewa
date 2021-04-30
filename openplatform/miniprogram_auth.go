package openplatform

import (
	"gitee.com/wallesoft/ewa/kernel/http"
	"github.com/gogf/gf/frame/g"
)

//小程序登录
func (mp *MiniProgram) Session(code string) *http.ResponseData {
	data := g.Map{
		"appid":           mp.Config.AppID,
		"js_code":         code,
		"grant_type":      "authorization_code",
		"component_appid": mp.Component.config.AppID,
	}
	return &http.ResponseData{
		Json: mp.Component.getClientWithToken().RequestJson("GET", "sns/component/jscode2session", data),
	}
}
