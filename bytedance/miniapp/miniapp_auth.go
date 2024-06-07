package miniapp

import (
	"context"

	"gitee.com/wallesoft/ewa/bytedance/http"
	"github.com/gogf/gf/v2/frame/g"
)

// 登录凭证校验
// Session 根据用户授权代码或匿名代码获取会话信息。
// 该方法用于交换微信小程序的用户会话密钥，可以通过用户授权的code或匿名code来获取。
// 参数:
//
//	ctx: 上下文对象，用于传递请求期间的共享信息。
//	code: 用户授权后返回的code，用于获取用户会话信息。
//	anonymousCode: 可选参数，用于匿名登录的code。
//
// 返回值:
//
//	*http.ResponseData: 包含请求结果的响应数据对象。
func (app *MiniApp) Session(ctx context.Context, code string, anonymousCode ...string) *http.ResponseData {
	// 初始化请求数据，包含应用的ID和密钥。
	data := g.Map{
		"appid":  app.Config.AppID,
		"secret": app.Config.Secret,
	}
	// 如果提供了用户授权的code，则将其添加到请求数据中。
	if code != "" {
		data["code"] = code
	}
	// 如果提供了匿名code，并且这是第一个匿名code，则将其添加到请求数据中。
	if len(anonymousCode) > 0 {
		data["anonymous_code"] = anonymousCode[0]
	}
	// 发起POST请求，交换会话信息，并返回响应数据。
	return &http.ResponseData{
		Json: app.GetClient(ctx).RequestJson(ctx, "POST", "jscode2session", data),
	}
}
