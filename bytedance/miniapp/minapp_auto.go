package miniapp

import (
	"gitee.com/wallesoft/ewa/kernel/http"
	"github.com/gogf/gf/frame/g"
)

//获取登录凭证
func (app *MiniApp) Session(code string, anonymousCode ...string) *http.ResponseData {
	data := g.Map{
		"appid":  app.config.AppID,
		"secret": app.config.Secret,
	}
	if code != "" {
		data["code"] = code
	}
	if len(anonymousCode) > 0 {
		data["anonymous_code"] = anonymousCode[0]
	}
	return &http.ResponseData{
		Json: app.GetClient().RequestJson("GET", "apps/jscode2session", data),
	}
}
