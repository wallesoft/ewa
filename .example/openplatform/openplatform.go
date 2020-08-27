package main

import (
	"gitee.com/wallesoft/ewa/openplatform"
	"github.com/gogf/gf/frame/g"
)

func main() {
	s, err := openplatform.New(map[string]interface{}{
		"Appid":  "appid",
		"Secret": "secret",
	})
	g.Dump(err)
	a := s.Server()
	a.Logger.SetDebug(true)
	a.Logger.SetStdoutPrint(false)
	a.Logger.SetPath("/tmp/log/")
	a.Logger.Debug(s.Server().Config, "teshi adfasdfasdfads")
	g.Dump()

}
