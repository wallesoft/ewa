package openplatform

import (
	"context"

	"gitee.com/wallesoft/ewa/kernel/http"
	"github.com/gogf/gf/v2/frame/g"
)

//获取基本信息
func (mp *MiniProgram) GetBasicInfo(ctx context.Context) *http.ResponseData {
	return &http.ResponseData{
		Json: mp.GetClientWithToken().RequestJson(ctx, "GET", "cgi-bin/account/getaccountbasicinfo"),
	}
}

//设置域名等 @see https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/Mini_Programs/Server_Address_Configuration.html
func (mp *MiniProgram) ModifyDomain(ctx context.Context, action string, config ...g.Map) *http.ResponseData {
	client := mp.GetClientWithToken()
	if action == "get" {
		return &http.ResponseData{
			Json: client.RequestJson(ctx, "POST", "wxa/modify_domain", g.Map{"action": action}),
		}
	}

	var param g.Map
	if len(config) > 0 {
		param = config[0]
	}

	param["action"] = action
	return &http.ResponseData{
		Json: client.RequestJson(ctx, "POST", "wxa/modify_domain", param),
	}
}

//设置业务域名
func (mp *MiniProgram) SetWebviewDomain(ctx context.Context, action string, domain ...g.Slice) *http.ResponseData {
	param := g.Map{}
	if action != "" {
		param["action"] = action
	}
	if action != "get" && len(domain) > 0 {
		param["webviewdomain"] = domain[0]
	}
	return &http.ResponseData{
		Json: mp.GetClientWithToken().RequestJson(ctx, "POST", "wxa/setwebviewdomain", param),
	}
}

//设置名称 注意事项 @see https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/Mini_Programs/setnickname.html
func (mp *MiniProgram) SetNickName(ctx context.Context, nickname string, config ...g.Map) *http.ResponseData {
	var data g.Map
	if len(config) > 0 {
		data = config[0]
	}
	data["nick_name"] = nickname
	return &http.ResponseData{
		Json: mp.GetClientWithToken().RequestJson(ctx, "POST", "wxa/setnickname", data),
	}
}

//微信认证名称检测 @see https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/Mini_Programs/wxverify_checknickname.html
//注：该接口只允许通过api创建的小程序使用。
func (mp *MiniProgram) VerifyNickname(ctx context.Context, nickname string) *http.ResponseData {
	return &http.ResponseData{
		Json: mp.GetClientWithToken().RequestJson(ctx, "POST", "cgi-bin/wxverify/checkwxverifynickname", g.Map{"nick_name": nickname}),
	}
}

//查询改名审核状态
func (mp *MiniProgram) QueryNickName(ctx context.Context, audit string) *http.ResponseData {
	return &http.ResponseData{
		Json: mp.GetClientWithToken().RequestJson(ctx, "POST", "wxa/api_wxa_querynickname", g.Map{"audit_id": audit}),
	}
}

//修改头像 @see https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/Mini_Programs/modifyheadimage.html
//注意事项，及config需要配置参数查看官方说明文档
func (mp *MiniProgram) UpdateAvatar(ctx context.Context, config g.Map) *http.ResponseData {
	return &http.ResponseData{
		Json: mp.GetClientWithToken().RequestJson(ctx, "POST", "cgi-bin/account/modifyheadimage", config),
	}
}

//修改简介
func (mp *MiniProgram) SetSignature(ctx context.Context, signature string) *http.ResponseData {
	return &http.ResponseData{
		Json: mp.GetClientWithToken().RequestJson(ctx, "POST", "cgi-bin/account/modifyheadimage", g.Map{"signature": signature}),
	}
}

//查询隐私设置 通过本接口可以查询小程序当前的隐私设置，即是否可被搜索。
func (mp *MiniProgram) GetSearchStatus(ctx context.Context) *http.ResponseData {
	return &http.ResponseData{
		Json: mp.GetClientWithToken().RequestJson(ctx, "GET", "wxa/getwxasearchstatus"),
	}
}

//修改隐私设置 通过本接口修改小程序隐私设置，即修改是否可被搜索
// @param status int  1 表示不可搜索，0 表示可搜索
func (mp *MiniProgram) ChangeSearchStatus(ctx context.Context, status int) *http.ResponseData {
	return &http.ResponseData{
		Json: mp.GetClientWithToken().RequestJson(ctx, "POST", "wxa/changewxasearchstatus", g.Map{"status": status}),
	}
}
