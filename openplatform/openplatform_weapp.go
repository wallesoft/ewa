package openplatform

import (
	"gitee.com/wallesoft/ewa/kernel/http"
	"github.com/gogf/gf/frame/g"
)

//获取代码草稿列表 @see https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/Mini_Programs/code_template/gettemplatedraftlist.html
func (op *OpenPlatform) GetTemplateDraftList() *http.ResponseData {
	client := op.getClientWithToken()
	return &http.ResponseData{
		Json: client.RequestJson("GET", "wxa/gettemplatedraftlist"),
	}
}

//将草稿添加到模板库 @see https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/Mini_Programs/code_template/addtotemplate.html
func (op *OpenPlatform) AddToTemplate(draftId string) *http.ResponseData {
	client := op.getClientWithToken()
	return &http.ResponseData{
		Json: client.RequestJson("POST", "wxa/addtotemplate", g.Map{"draft_id": draftId}),
	}
}

//获取模板库列表 @see https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/Mini_Programs/code_template/gettemplatelist.html
func (op *OpenPlatform) GetTemplateList() *http.ResponseData {
	client := op.getClientWithToken()
	return &http.ResponseData{
		Json: client.RequestJson("GET", "wxa/gettemplatelist"),
	}
}

//从模板库中删除对应模板 @see https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/Mini_Programs/code_template/deletetemplate.html
func (op *OpenPlatform) DelTemplate(templateId string) *http.ResponseData {
	client := op.getClientWithToken()
	return &http.ResponseData{
		Json: client.RequestJson("POST", "wxa/deletetemplate", g.Map{"template_id": templateId}),
	}
}
