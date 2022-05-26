package openplatform

import (
	"context"

	"gitee.com/wallesoft/ewa/kernel/http"
	"github.com/gogf/gf/v2/frame/g"
)

//小程序登录
func (mp *MiniProgram) Session(ctx context.Context, code string) *http.ResponseData {
	data := g.Map{
		"appid":           mp.Config.AppID,
		"js_code":         code,
		"grant_type":      "authorization_code",
		"component_appid": mp.Component.config.AppID,
	}
	return &http.ResponseData{
		Json: mp.Component.getClientWithToken().RequestJson(ctx, "GET", "sns/component/jscode2session", data),
	}
}
