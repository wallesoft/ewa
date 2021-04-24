package openplatform

import (
	"gitee.com/wallesoft/ewa/kernel/http"
	"github.com/gogf/gf/frame/g"
)

//代码上传
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
