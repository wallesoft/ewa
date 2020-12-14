package base

import (
	"github.com/gogf/gf/net/ghttp"
)

type Client struct {
	*ghttp.Client
	BaseUri string
}
