package openplatform

import (
	"context"
	"net/url"

	"gitee.com/wallesoft/ewa/kernel/http"
	"github.com/gogf/gf/v2/container/gvar"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gutil"
)

// StartPushTicket 启动ticket推送服务
func (op *OpenPlatform) StartPushTicket(ctx context.Context) *http.ResponseData {
	return &http.ResponseData{
		Json: op.getClient().RequestJson(ctx, "POST", "cgi-bin/component/api_start_push_ticket", g.Map{"component_appid": op.config.AppID, "component_secret": op.config.AppSecret}),
	}
}

// GetPreAuthCode 获取预授权码
func (op *OpenPlatform) GetPreAuthCode(ctx context.Context) (string, error) {
	var code string
	var err error
	gutil.TryCatch(ctx, func(ctx context.Context) {
		client := op.getClientWithToken()
		v := client.RequestJson(ctx, "POST", "cgi-bin/component/api_create_preauthcode", map[string]string{
			"component_appid": op.config.AppID,
		})

		if have := v.Contains("errcode"); have {
			panic(v.MustToJsonString())
		}
		if have := v.Contains("pre_auth_code"); have {
			code = v.Get("pre_auth_code").String()
		} else {
			panic("Request pre_auth_code fail:" + v.MustToJsonString())
		}
	}, func(ctx context.Context, e error) {
		err = e
		op.Logger.File(op.Logger.ErrorLogPattern).Stdout(op.Logger.LogStdout).Error(ctx, err.Error())
	})

	return code, err
}

// GetPreAuthorizationUrl 获取授权页网址
func (op *OpenPlatform) GetPreAuthorizationUrl(ctx context.Context, callback string, optional ...map[string]interface{}) (string, error) {

	val := url.Values{}
	authCode, err := op.GetPreAuthCode(ctx)
	if err != nil {
		return "", err
	}
	val.Add("pre_auth_code", authCode)
	if len(optional) > 0 {
		options := optional[0]
		if v, ok := options["auth_type"]; ok {
			val.Add("auth_type", gvar.New(v).String())
		}
		if v, ok := options["biz_appid"]; ok {
			val.Add("biz_appid", gvar.New(v).String())
		}
	}

	val.Add("component_appid", op.config.AppID)
	val.Add("redirect_uri", callback)

	return "https://mp.weixin.qq.com/cgi-bin/componentloginpage?" + val.Encode(), nil
}

// GetMobilePreAuthorizationUrl
func (op *OpenPlatform) GetMobilePreAuthorizationUrl(ctx context.Context, callback string, optional ...map[string]interface{}) (string, error) {
	val := url.Values{}
	authCode, err := op.GetPreAuthCode(ctx)
	if err != nil {
		return "", err
	}
	val.Add("pre_auth_code", authCode)
	if len(optional) > 0 {
		options := optional[0]
		if v, ok := options["auth_type"]; ok {
			val.Add("auth_type", gvar.New(v).String())
		}
		if v, ok := options["biz_appid"]; ok {
			val.Add("biz_appid", gvar.New(v).String())
		}
	}

	val.Add("component_appid", op.config.AppID)
	val.Add("redirect_uri", callback)
	val.Add("action", "bindcomponent")
	val.Add("no_scan", "1")

	return "https://mp.weixin.qq.com/safe/bindcomponent?" + val.Encode() + "#wechat_redirect", nil
}

// HandleAuthorize
func (op *OpenPlatform) HandleAuthorize(ctx context.Context, code string) *http.ResponseData {

	client := op.getClientWithToken()
	return &http.ResponseData{
		Json: client.RequestJson(ctx, "POST", "cgi-bin/component/api_query_auth", map[string]string{
			"component_appid":    op.config.AppID,
			"authorization_code": code,
		}),
	}

}

// GetAuthorizer get authorizer info type as gjson.Json
func (op *OpenPlatform) GetAuthorizer(ctx context.Context, appid string) *http.ResponseData {
	client := op.getClientWithToken()
	return &http.ResponseData{
		Json: client.RequestJson(ctx, "POST", "cgi-bin/component/api_get_authorizer_info", map[string]string{
			"component_appid":  op.config.AppID,
			"authorizer_appid": appid,
		}),
	}

}

// GetAuthorizers get authorizer list
func (op *OpenPlatform) GetAuthorizers(ctx context.Context, offset int, count int) *http.ResponseData {
	if count > 500 {
		count = 500
	}
	client := op.getClientWithToken()
	return &http.ResponseData{
		Json: client.RequestJson(ctx, "POST", "cgi-bin/component/api_get_authorizer_list", map[string]interface{}{
			"component_appid": op.config.AppID,
			"offset":          offset,
			"count":           count,
		}),
	}
}

// GetAuthorizerOption get authorizer option info
func (op *OpenPlatform) GetAuthorizerOption(ctx context.Context, appid string, name string) *http.ResponseData {
	client := op.getClientWithToken()
	return &http.ResponseData{
		Json: client.RequestJson(ctx, "POST", "cgi-bin/component/api_get_authorizer_option", map[string]string{
			"component_appid":  op.config.AppID,
			"authorizer_appid": appid,
			"option_name":      name,
		}),
	}
}

// SetAuthorizerOption set authorizer option
func (op *OpenPlatform) SetAuthorizerOption(ctx context.Context, appid string, name string, value string) *http.ResponseData {
	client := op.getClientWithToken()
	return &http.ResponseData{
		Json: client.RequestJson(ctx, "POST", "cgi-bin/component/api_set_authorizer_option", map[string]string{
			"component_appid":  op.config.AppID,
			"authorizer_appid": appid,
			"option_name":      name,
			"option_value":     value,
		}),
	}
}

// GetVerifyTicket
func (op *OpenPlatform) GetVerifyTicket(ctx context.Context) string {
	return op.verifyTicket.GetTicket(ctx)
}

// GetAccessToken
func (op *OpenPlatform) GetAccessToken(ctx context.Context) string {
	return op.accessToken.GetToken(ctx)
}
