package openplatform

import (
	"context"

	"gitee.com/wallesoft/ewa/kernel/http"
	"github.com/gogf/gf/v2/frame/g"
)

//绑定微信用户为体验者
func (mp *MiniProgram) BindTester(wechatId string) *http.ResponseData {
	return &http.ResponseData{
		Json: mp.GetClientWithToken().RequestJson(context.TODO(), "POST", "wxa/bind_tester", g.Map{"wechatid": wechatId}),
	}
}

//解除绑定体验者 userstr 和 wechatid 填写其中一个即可: g.Map{"userstr":"xxxx"} 或 g.Map{"wechatid":"xxxx"}
func (mp *MiniProgram) UnbindTester(tester g.Map) *http.ResponseData {
	return &http.ResponseData{
		Json: mp.GetClientWithToken().RequestJson(context.TODO(), "POST", "wxa/unbind_tester", tester),
	}
}

//获取体验者列表 通过本接口可以获取小程序所有已绑定的体验者列表
func (mp *MiniProgram) GetTesters() *http.ResponseData {
	return &http.ResponseData{
		Json: mp.GetClientWithToken().RequestJson(context.TODO(), "POST", "wxa/memberauth", g.Map{"action": "get_experiencer"}),
	}
}
