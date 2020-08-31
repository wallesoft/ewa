package main

import (
	"gitee.com/wallesoft/ewa/kernel/server"
	"gitee.com/wallesoft/ewa/openplatform"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/util/gconv"
)

func main() {
	r, err := server.NewRequest(map[string]interface{}{
		"EncryptType":  "aes",
		"MsgSignature": "msgsing",
		"Nonce":        "nonce",
		"RawBody":      gconv.Bytes("aaaa"),
		"Timestamp":    "teim",
	})
	g.Dump(err)
	g.Dump(r)
	s, err := openplatform.New(map[string]interface{}{
		"Appid":  "appid",
		"Secret": "secret",
		"Token":  "thisistoken",
		"AesKey": "aeskey",
	})
	g.Dump(s)
	g.Dump(err)
	a := s.Server(r)
	a.Logger.SetDebug(true)
	//a.Logger.SetStdoutPrint(false)
	a.Logger.SetPath("/tmp/log/")
	//a.Logger.Debug(s.Server().Config.Get("token"), "teshi adfasdfasdfads")
	//a.Logger.Debugf("request: v%", s.Server().Config)
	// a.Logger.Debug(map[string]interface{}{
	// 	"request": s.Server().Config,
	// 	"content": "sasdfasdf",
	// })
	g.Dump(a.IsSafeMode())
	a.Serve()

}
