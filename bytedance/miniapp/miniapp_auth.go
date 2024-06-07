package miniapp

import (
	"context"

	"gitee.com/wallesoft/ewa/bytedance/http"
	"github.com/gogf/gf/v2/frame/g"
)

// Validates login credentials and fetches session information.
// The Session function is used to exchange the WeChat Mini Program session key,
// which can be obtained through the user-authorized code or anonymous code.
// Parameters:
//
//	ctx: Context object for sharing information during the request.
//	code: The code returned after user authorization, used to get session info.
//	anonymousCode: Optional parameter for anonymous login code.
//
// Returns:
//
//	*http.ResponseData: A response data object containing the result of the request.
//
// Validates login credentials and fetches session information.
// The Session function is used to exchange the WeChat Mini Program session key,
// which can be obtained through the user-authorized code or anonymous code.
// Parameters:
//
//	ctx: Context object for sharing information during the request.
//	code: The code returned after user authorization, used to get session info.
//	anonymousCode: Optional parameter for anonymous login code.
//
// Returns:
//
//	*http.ResponseData: A response data object containing the result of the request.
func (app *MiniApp) Session(ctx context.Context, code string, anonymousCode ...string) *http.ResponseData {
	// Initializes request data, including the application ID and secret.
	data := g.Map{
		"appid":  app.Config.AppID,
		"secret": app.Config.Secret,
	}
	// If a user-authorized code is provided, it is added to the request data.
	if code != "" {
		data["code"] = code
	}
	// If an anonymous code is provided and this is the first anonymous code,
	// it is added to the request data.
	if len(anonymousCode) > 0 {
		data["anonymous_code"] = anonymousCode[0]
	}
	// Initiates a POST request to exchange session information and returns the response data.
	return &http.ResponseData{
		Json: app.GetClient(ctx).RequestJson(ctx, "POST", "jscode2session", data),
	}
}
