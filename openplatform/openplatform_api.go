package openplatform

import (
	"net/url"

	"github.com/gogf/gf/container/gvar"
)

//GetPreAuthorizationUrl 获取授权页网址
func (op *OpenPlatform) GetPreAuthorizationUrl(callback string, optional ...map[string]interface{}) string {
	// get_pre_auth_code
	-----------------------------------
	// build query
	val := &url.Values{}
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

	return "https://mp.weixin.qq.com/cgi-bin/componentloginpage?" + val.Encode()
}
