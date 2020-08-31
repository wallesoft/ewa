package main

import (
	"gitee.com/wallesoft/ewa/kernel/server"
	"gitee.com/wallesoft/ewa/openplatform"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/util/gconv"
)

func main() {
	r, _ := server.NewRequest(map[string]interface{}{
		"EncryptType":  "aes",
		"MsgSignature": "9096dc4f1ee13911603d3e436f5b0dfd22b088dd",
		"Nonce":        "199454771",
		"RawBody":      gconv.Bytes("<xml><AppId><![CDATA[wx3832e3725afda621]]></AppId><Encrypt><![CDATA[lT3mY6ueqntYmQV4bl/sj+Zjp/jy1v7eI4NWVqmHKsLTZbe8we6BwfedsWixX5fV5IMljzoNCoK7xt1S4Bld6RCmAipTFDgqudmAkCz8Jjw7S0JVm3zXn/hQgjVnKtE1PId52kFSOyoaRYjQ+bL6mGsPvnDBQnTkX8tl8BiSY9PCRQSurr2P++dX4hazreE7UCzV6wbFJKIpi5F36jtvyzWcbkRS0s/Fix9/qu0IPEg6aW6E91E/OAGE7v4nMa5nU9Fvh/KJF24TKThNoyJvqP8UFhBnmtakGmMM2ZUItXX+pfoLX7pk3SC4yG8KQu5HPUSpbr3oZri0v2gRxf7sW8zqdB5lj4rQxUUcx6GL5xd9Q4kb/RSkN49yzLQhYTW6WzLBLI0sVkVjUfxi9kP/q2cDP9+YuEeuH4Rdn/eiXiEiaWH54EcoocDjF058/cdO5cP4LXzGlNyM15B3fZXPxA==]]></Encrypt></xml>"),
		"Timestamp":    1598631541,
		"Signature":    "04db1597261c5718148ed26d04f96a4aa0ec6b0c",
	})

	s, _ := openplatform.New(map[string]interface{}{
		"Appid":  "wx3832e3725afda621",
		"Secret": "4587be61689594ca7bd877895374031f",
		"Token":  "kB90oxaqQK7Aaj7qXEQVXN2Q21fbWXKO",
		"AesKey": "D8o4EDD2FVfZxf35x23vOF2nxXEv4J33f5jmjfOOK3x",
	})

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
