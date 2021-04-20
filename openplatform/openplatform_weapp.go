package openplatform

import "github.com/gogf/gf/encoding/gjson"

//获取代码草稿列表 @see https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/Mini_Programs/code_template/gettemplatedraftlist.html
func (op *OpenPlatform) GetTemplateDraftList() *gjson.Json {
	client := op.getClientWithToken()
	return client.RequestJson("GET", "wxa/gettemplatedraftlist")
}
