package openplatform

import (
	"context"

	"gitee.com/wallesoft/ewa/kernel/http"
	"github.com/gogf/gf/v2/frame/g"
)

//获取代码草稿列表 @see https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/Mini_Programs/code_template/gettemplatedraftlist.html
func (op *OpenPlatform) GetTemplateDraftList(ctx context.Context) *http.ResponseData {
	client := op.getClientWithToken()
	return &http.ResponseData{
		Json: client.RequestJson(ctx, "GET", "wxa/gettemplatedraftlist"),
	}
}

//将草稿添加到模板库 @see https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/Mini_Programs/code_template/addtotemplate.html
func (op *OpenPlatform) AddToTemplate(ctx context.Context, draftId string) *http.ResponseData {
	client := op.getClientWithToken()
	return &http.ResponseData{
		Json: client.RequestJson(ctx, "POST", "wxa/addtotemplate", g.Map{"draft_id": draftId}),
	}
}

//获取模板库列表 @see https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/Mini_Programs/code_template/gettemplatelist.html
func (op *OpenPlatform) GetTemplateList(ctx context.Context) *http.ResponseData {
	client := op.getClientWithToken()
	return &http.ResponseData{
		Json: client.RequestJson(ctx, "GET", "wxa/gettemplatelist"),
	}
}

//从模板库中删除对应模板 @see https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/Mini_Programs/code_template/deletetemplate.html
func (op *OpenPlatform) DelTemplate(ctx context.Context, templateId string) *http.ResponseData {
	client := op.getClientWithToken()
	return &http.ResponseData{
		Json: client.RequestJson(ctx, "POST", "wxa/deletetemplate", g.Map{"template_id": templateId}),
	}
}
