package openplatform

import (
	"gitee.com/wallesoft/ewa/kernel/http"
	"gitee.com/wallesoft/ewa/miniprogram"
	"github.com/gogf/gf/frame/g"
)

//代码上传 @see https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/Mini_Programs/code/commit.html
//and ext_json @see https://developers.weixin.qq.com/miniprogram/dev/devtools/ext.html#%E5%B0%8F%E7%A8%8B%E5%BA%8F%E6%A8%A1%E6%9D%BF%E5%BC%80%E5%8F%91
func (mp *MiniProgram) Commit(templateId string, extJson string, version string, desc string) *http.ResponseData {
	client := mp.GetClientWithToken()
	return &http.ResponseData{
		Json: client.RequestJson("POST", "wxa/commit", g.Map{
			"template_id":  templateId,
			"ext_json":     extJson,
			"user_version": version,
			"user_desc":    desc,
		}),
	}
}

//获取已上传的代码的页面列表 @see https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/Mini_Programs/code/get_page.html
func (mp *MiniProgram) GetPage() *http.ResponseData {
	client := mp.GetClientWithToken()
	return &http.ResponseData{
		Json: client.RequestJson("GET", "wxa/get_page"),
	}
}

//获取体验二维码 @see https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/Mini_Programs/code/get_qrcode.html
func (mp *MiniProgram) GetQrcode(path string) *miniprogram.AppCode {
	client := mp.GetClientWithToken()
	return &miniprogram.AppCode{
		Mp:  mp.MiniProgram,
		Raw: client.RequestRaw("GET", "wxa/get_qrcode", g.Map{"path": path}),
	}
}
