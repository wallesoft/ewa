package miniprogram

import (
	"context"

	"gitee.com/wallesoft/ewa/kernel/http"
	"github.com/gogf/gf/v2/frame/g"
)

//登录凭证校验
func (mp *MiniProgram) Session(ctx context.Context, code string) *http.ResponseData {
	data := g.Map{
		"appid":      mp.Config.AppID,
		"secret":     mp.Config.Secret,
		"grant_type": "authorization_code",
		"js_code":    code,
	}
	return &http.ResponseData{
		Json: mp.GetClient().RequestJson(ctx, "GET", "sns/jscode2session", data),
	}
}
