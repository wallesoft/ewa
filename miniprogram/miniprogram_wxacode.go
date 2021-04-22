package miniprogram

import "github.com/gogf/gf/frame/g"

func (mp *MiniProgram) GetUnlimit(scene string, config ...map[string]interface{}) {
	client := mp.getClientWithToken()
	client.RequestJson("POST", "wxa/getwxacodeunlimit", g.Map{
		"scene": scene,
	}).MustToJsonString()
}
