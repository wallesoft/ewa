package encryptor_test

// 以下提供测j加解密方法， 因为需要相关敏感配置信息，所以注释掉，如需测试
// 根据自己实际数据进行测试y已下y已经测试通过
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
	Token:     "kB90oxaqQK7Aaj7qXEQVXN2Q21fbWXKO",
	BlockSize: 32,
}

func Test_Decrypt(t *testing.T) {
	key, err := gbase64.DecodeToString("D8o4EDD2FVfZxf35x23vOF2nxXEv4J33f5jmjfOOK3x" + "=")
	if err != nil {
		panic(err)
	}
	e.AesKey = key
	//content := []byte(`lT3mY6ueqntYmQV4bl/sj+Zjp/jy1v7eI4NWVqmHKsLTZbe8we6BwfedsWixX5fV5IMljzoNCoK7xt1S4Bld6RCmAipTFDgqudmAkCz8Jjw7S0JVm3zXn/hQgjVnKtE1PId52kFSOyoaRYjQ+bL6mGsPvnDBQnTkX8tl8BiSY9PCRQSurr2P++dX4hazreE7UCzV6wbFJKIpi5F36jtvyzWcbkRS0s/Fix9/qu0IPEg6aW6E91E/OAGE7v4nMa5nU9Fvh/KJF24TKThNoyJvqP8UFhBnmtakGmMM2ZUItXX+pfoLX7pk3SC4yG8KQu5HPUSpbr3oZri0v2gRxf7sW8zqdB5lj4rQxUUcx6GL5xd9Q4kb/RSkN49yzLQhYTW6WzLBLI0sVkVjUfxi9kP/q2cDP9+YuEeuH4Rdn/eiXiEiaWH54EcoocDjF058/cdO5cP4LXzGlNyM15B3fZXPxA==`)
	content := []byte(`MDHdIZkEbATQQ6lAD+Wv6Ail9CBHgQY5d0DsQLq2zze5NfZPc7WoFXOU53n1hglCMv3KZGjFEmuoOh1IxBvEpY6ydvDJqR7UJCkEAEdZcTV812RD6DDBJ7HsR7DWu1Ii+l2DbR4qBWEXCcnCPEr7BXEHvnHiYYLM9rxAQpeNfias6Hku+weDs0CipW2EJqUtY7i/BNTkpIEsHguzGQUKZJZ5XNrE4Ji3rKlJ+Ndj/Ic5v1QzGNsCJUa784kZpNrIUh5+3c85buyzpSA6qJMbfPqFiZAfppDdysbOWNUAPU0s8kS8fuEng7Uf269kLy/8ZiAqy7INny9NbgTvHUNLP4R8/gnoKV5Dx8dwOSvFHAIf8qU9lJU6PD4nVsmwFrxZTd0lToXQH/OoRWikKiepNirPI1yIX0YRniN7uC5Ixl/pZgp4035JvkoSAci95ZFx3WfgbElhqNY8DO33Kga47BA+W+w9Ta0zqrB1KkFvqBQ=`)
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
