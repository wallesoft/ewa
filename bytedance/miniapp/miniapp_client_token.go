package miniapp

import "context"

// 凭证
type Credentials struct {
	clientKey    string // 应用唯一标识，对应小程序id
	clientSecret string //应用唯一标识对应的密钥，对应小程序的app secret，可以在开发者后台获取
}

// 实现TokenCredentail接口
func (c *Credentials) Get(ctx context.Context) map[string]string {
	return map[string]string{
		"grant_type":    "client_credential",
		"client_key":    c.clientKey,
		"client_secret": c.clientSecret,
	}
}
