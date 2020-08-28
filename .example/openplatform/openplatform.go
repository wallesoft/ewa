package main

import (
	"gitee.com/wallesoft/ewa/openplatform"
	"github.com/gogf/gf/frame/g"
)

func main() {
	s, err := openplatform.New(map[string]interface{}{
		"Appid":  "appid",
		"Secret": "secret",
		"Token":  "thisistoken",
	})
	g.Dump(err)
	a := s.Server()
	a.Logger.SetDebug(true)
	//a.Logger.SetStdoutPrint(false)
	a.Logger.SetPath("/tmp/log/")
	//a.Logger.Debug(s.Server().Config.Get("token"), "teshi adfasdfasdfads")
	//a.Logger.Debugf("request: v%", s.Server().Config)
	// a.Logger.Debug(map[string]interface{}{
	// 	"request": s.Server().Config,
	// 	"content": "sasdfasdf",
	// })
	g.Dump(a)
	a.Serve()

}
