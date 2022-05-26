package main

import (
	"gitee.com/wallesoft/ewa/openplatform"
	"github.com/gogf/gf/v2/frame/g"
)

func main() {
	// r, err := server.NewRequest(map[string]interface{}{
	// 	"EncryptType":  "aes",
	// 	"MsgSignature": "msgsing",
	// 	"Nonce":        "nonce",
	// 	"RawBody":      gconv.Bytes("aaaa"),
	// 	"Timestamp":    "teim",
	// })
	// g.Dump(err)
	// g.Dump(r)
	s := openplatform.New(openplatform.Config{
		AppID:          "appid",
		AppSecret:      "secret",
		Token:          "thisistoken",
		EncodingAESKey: "aeskey",
	})
	// g.Dump(err)
	g.Dump(s)
	// g.Dump(s)
	// g.Dump(err)
	// a := s.Server(r)
	// a.Logger.SetDebug(true)
	// //a.Logger.SetStdoutPrint(false)
	// a.Logger.SetPath("/tmp/log/")
	// //a.Logger.Debug(s.Server().Config.Get("token"), "teshi adfasdfasdfads")
	// //a.Logger.Debugf("request: v%", s.Server().Config)
	// // a.Logger.Debug(map[string]interface{}{
	// // 	"request": s.Server().Config,
	// // 	"content": "sasdfasdf",
	// // })
	// g.Dump(a.IsSafeMode())
	// a.Serve()

}
