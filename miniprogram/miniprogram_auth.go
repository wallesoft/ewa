package miniprogram

import (
	"gitee.com/wallesoft/ewa/kernel/http"
	"github.com/gogf/gf/frame/g"
)

//登录凭证校验
func (mp *MiniProgram) Session(code string) *http.ResponseData {
	data := g.Map{
		"appid":      mp.Config.AppID,
		"secret":     mp.Config.Secret,
		"grant_type": "authorization_code",
		"js_code":    code,
	}
	return &http.ResponseData{
		Json: mp.GetClient().RequestJson("GET", "sns/jscode2session", data),
	}
}
