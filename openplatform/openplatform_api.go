package openplatform

import (
	"net/url"

	"github.com/gogf/gf/container/gvar"
	"github.com/gogf/gf/encoding/gjson"
)

//GetPreAuthorizationUrl 获取授权页网址
func (op *OpenPlatform) GetPreAuthorizationUrl(callback string, optional ...map[string]interface{}) string {

	val := &url.Values{}
	authCode := op.GetPreAuthCode()
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

	return "https://mp.weixin.qq.com/cgi-bin/componentloginpage?" + val.Encode()
}

//GetVerifyTicket
func (op *OpenPlatform) GetVerifyTicket() string {
	return op.verifyTicket.GetTicket()
}

//GetAccessToken
func (op *OpenPlatform) GetAccessToken() string {
	return op.accessToken.GetToken()
}

func (op *OpenPlatform) GetPreAuthCode() string {
	client := op.getClientWithToken()
	result := client.PostJson("cgi-bin/component/api_create_preauthcode", map[string]string{
		"component_appid": op.config.AppID,
	})
	v := gjson.New(result)
	if have := v.Contains("errcode"); have {
		panic(v.MustToJsonString())
	}
	if have := v.Contains("pre_auth_code"); have {
		return v.GetString("pre_auth_code")
	}
	panic("Request pre_auth_code fail:" + v.MustToJsonString())

}
