package miniapp

import (
	"context"

	"gitee.com/wallesoft/ewa/bytedance/http"
	"github.com/gogf/gf/v2/frame/g"
)

// 登录凭证校验
func (app *MiniApp) Session(ctx context.Context, code string, anonymousCode ...string) *http.ResponseData {
	data := g.Map{
		"appid":  app.Config.AppID,
		"secret": app.Config.Secret,
	}
	if code != "" {
		data["code"] = code
	}
	if len(anonymousCode) > 0 {
		data["anonymous_code"] = anonymousCode[0]
	}
	return &http.ResponseData{
		Json: app.GetClient(ctx).RequestJson(ctx, "POST", "jscode2session", data),
	}
}
