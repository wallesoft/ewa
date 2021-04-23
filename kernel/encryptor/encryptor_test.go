package encryptor_test

// 以下提供测j加解密方法， 因为需要相关敏感配置信息，所以注释掉，如需测试
// 根据自己实际数据进行测试y已下y已经测试通过
// 根据配置更改下++++++++++++++++++++++为你自己配置进行测试
import (
	"testing"

	"gitee.com/wallesoft/ewa/kernel/encryptor"
	"gitee.com/wallesoft/ewa/kernel/server"
	"github.com/gogf/gf/encoding/gbase64"
	"github.com/gogf/gf/encoding/gjson"
	"github.com/gogf/gf/frame/g"
)

// 部分数据敏感，已通过测试，敏感数据去除
var e = &encryptor.Encryptor{
	AppID:     "wx3832e3725afda621",
	Token:     "+++++++++++++++++++++++++",
	BlockSize: 32,
}

func Test_Decrypt(t *testing.T) {
	key, err := gbase64.DecodeToString("+++++++++++++++++++" + "=")
	if err != nil {
		panic(err)
	}
	e.AesKey = key
	g.Dump(key)
	// content := []byte(`lT3mY6ueqntYmQV4bl/sj+Zjp/jy1v7eI4NWVqmHKsLTZbe8we6BwfedsWixX5fV5IMljzoNCoK7xt1S4Bld6RCmAipTFDgqudmAkCz8Jjw7S0JVm3zXn/hQgjVnKtE1PId52kFSOyoaRYjQ+bL6mGsPvnDBQnTkX8tl8BiSY9PCRQSurr2P++dX4hazreE7UCzV6wbFJKIpi5F36jtvyzWcbkRS0s/Fix9/qu0IPEg6aW6E91E/OAGE7v4nMa5nU9Fvh/KJF24TKThNoyJvqP8UFhBnmtakGmMM2ZUItXX+pfoLX7pk3SC4yG8KQu5HPUSpbr3oZri0v2gRxf7sW8zqdB5lj4rQxUUcx6GL5xd9Q4kb/RSkN49yzLQhYTW6WzLBLI0sVkVjUfxi9kP/q2cDP9+YuEeuH4Rdn/eiXiEiaWH54EcoocDjF058/cdO5cP4LXzGlNyM15B3fZXPxA==`)
	content := []byte(`9PlbipOZSBv9WQUf6HV15WPVl461mJPG/l3JzNu959LJzuF/av8LYmzLReZbM2D6+JKhZdePjruNaSPwecYCSaX5e/0O4bxCzp8VZdcB82gID36fZ0OF36qSCR4TWWszmn+aLlnrkgrfo81eZWDLGnAUCNvelWGi38VCzKRF0Vj1LeAfpKDXUI5M8O5m/OC7tYCRqcUqR3wYAGev3vmyijhfcW2vZzzWD7ve8prhUG0b1ZgA3uOAcNQ9tx2R7qc/zKeuiwpNzHkLJyeKyObIRWCAPY0OSi1WFPrhpegZcdhv81NTTSFRQBYVNF7xGaIXYVU7zZ0Cwrb0bsM4eyrgJP3iP3ZQmKQogOxOWCdRR04odZqKsy49sSirMu7C7cTIS3OHRhA3bgwB3nRlKjrxeEbjAurHdLamW4AA/98m8sHs5TnqZ+eoO5eQv7Knym4R1pEIJDk0YNVr2OXnb3ggtrOZgp9uZoxOJVxHlV4yx51qiNWLdxqFEBrZXXEgkYBZC00x9pqTHVCrHOKLV112sUNU8IN+64bq3EI9VYryNffiDerX9Or+YHoXFTI1iP6tLJtr8zYvqaj3sfAkOhAwyuvpkMeyqfswEJTGA0R1nrTds9cjJitvJhZeH8OFJ9QjgyEHnyqWBAaR+zwf8mQsGNEJMco6OLyJZwTxYVvdMChLhZMQYIDUu2SPk2W51N7AzxVsjGCiHYcQmShfAqO9eHrgVts2B54XWnXFgGtc+qyrtrPzhBW4cPoIc3YNM6S3x+3v43c+BZCb4AQ6WJnbfV3/V3XOFBnerOXfQKLnlfSf3UnPZc06Bljy74Ba3e0xVyk4RsiG2YyXC9Oow08IZg==`)
	// decoding, err := gbase64.Decode(content)
	// g.Dump(len(decoding))
	// // g.Dump("aeskey:", e.AesKey)
	// decrypted, err := gaes.Decrypt(decoding, gconv.Bytes(e.AesKey), gconv.Bytes(gstr.SubStr(e.AesKey, 0, 16)))
	// g.Dump(err.Error())
	// g.Dump(len(decrypted))
	//
	decrypted, err := e.Decrypt(content)
	g.Dump(decrypted)
	if err != nil {
		panic(err)
	}
	//g.Dump(descrypted)
	if j, err := gjson.LoadContent(decrypted); err == nil {
		g.Dump(j.Get("xml.AppId"))
		msg := &server.Message{
			Json: j,
		}
		g.Dump(msg.Get("xml.AppId"))
	} else {
		panic(err)
	}

}

func Test_Encrypt(t *testing.T) {
	raw := []byte(`xxxxxx`)
	key, err := gbase64.DecodeToString("xxxxx" + "=")
	if err != nil {
		panic(err)
	}
	e.AesKey = key
	encrypted, err := e.Encrypt(raw, "199454771", 1598631541)
	if err != nil {
		panic(err)
	}
	g.Dump(encrypted)
}
