package miniprogram

import "github.com/gogf/gf/frame/g"

func (mp *MiniProgram) GetUnlimit(scene string, config ...map[string]interface{}) {
	client := mp.getClientWithToken()
	g.Dump(client)
	g.Dump("bbbbbbbb")
	g.Dump(client.RequestJson("POST", "wxa/getwxacodeunlimit", g.Map{
		"scene": scene,
	}).MustToJsonString())
}
