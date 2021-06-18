package openplatform

import (
	"net/url"

	"gitee.com/wallesoft/ewa/kernel/http"
)

//通过法人微信快速创建小程序 config参数查看官方文档
func (op *OpenPlatform) FastRegisterWeapp(config map[string]interface{}) *http.ResponseData {
	client := op.getClientWithToken()
	urlVal := url.Values{}
	urlVal.Add("action", "create")
	client.UrlValues = urlVal
	return &http.ResponseData{
		Json: client.RequestJson("POST", "cgi-bin/component/fastregisterweapp", config),
	}
}

//体验小程序创建
func (op *OpenPlatform) FastRegisterBetaWeapp(config map[string]interface{}) *http.ResponseData {

	client := op.getClientWithToken()
	return &http.ResponseData{
		Json: client.RequestJson("POST", "wxa/component/fastregisterbetaweapp", config),
	}
}
