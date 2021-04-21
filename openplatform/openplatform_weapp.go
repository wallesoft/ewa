package openplatform

import (
	"github.com/gogf/gf/encoding/gjson"
	"github.com/gogf/gf/frame/g"
)

//获取代码草稿列表 @see https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/Mini_Programs/code_template/gettemplatedraftlist.html
func (op *OpenPlatform) GetTemplateDraftList() *gjson.Json {
	client := op.getClientWithToken()
	return client.RequestJson("GET", "wxa/gettemplatedraftlist")
}

//将草稿添加到模板库 @see https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/Mini_Programs/code_template/addtotemplate.html
func (op *OpenPlatform) AddToTemplate(draftId string) *gjson.Json {
	client := op.getClientWithToken()
	return client.RequestJson("POST", "wxa/addtotemplate", g.Map{"draft_id": draftId})
}

//获取模板库列表 @see https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/Mini_Programs/code_template/gettemplatelist.html
func (op *OpenPlatform) GetTemplateList() *gjson.Json {
	client := op.getClientWithToken()
	return client.RequestJson("GET", "wxa/gettemplatelist")
}

//从模板库中删除对应模板 @see https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/Mini_Programs/code_template/deletetemplate.html
func (op *OpenPlatform) DelTemplate(templateId string) *gjson.Json {
	client := op.getClientWithToken()
	return client.RequestJson("POST", "wxa/deletetemplate", g.Map{"template_id": templateId})
}
