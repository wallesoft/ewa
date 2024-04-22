package http

import (
	"context"
	"net/url"

	"gitee.com/wallesoft/ewa/kernel/log"
	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/net/gclient"
	"github.com/gogf/gf/v2/text/gstr"
)

type Client struct {
	Client *gclient.Client
	// BeforeRequest gclient.HandlerFunc
	// AfterReponse  gclient.HandlerFunc
	BaseUri   string // 请求接口通用前缀
	Logger    *log.Logger
	UrlValues url.Values
}

// request
func (c *Client) RequestJson(ctx context.Context, method string, endpoint string, data ...interface{}) *gjson.Json {
	if gstr.ToLower(method) == "post" {
		c.Client = c.Client.ContentJson()
	}
	raw := c.RequestRaw(ctx, method, endpoint, data...)
	return gjson.New(raw)
}

// request post
// 需要上传文件等用，配合 @file
func (c *Client) RequestPost(ctx context.Context, endpoint string, data ...interface{}) *gjson.Json {
	raw := c.RequestRaw(ctx, "POST", endpoint, data...)
	return gjson.New(raw)
}

// request raw
// retrun []byte
func (c *Client) RequestRaw(ctx context.Context, method string, endpoint string, data ...interface{}) []byte {
	// if c.BeforeRequest != nil {
	// 	c.Client = c.Client.Use(c.BeforeRequest)
	// }
	// if c.AfterReponse != nil {
	// 	c.Client = c.Client.Use(c.AfterReponse)
	// }
	response, err := c.Client.DoRequest(ctx, method, c.parseUrl(endpoint), data...)

	if err != nil {
		// log
		c.handleErrorLog(ctx, err, response.Raw())
		return response.ReadAll()
	}

	defer response.Close()
	// handle log
	if c.Logger != nil && c.Logger.AccessLogEnabled {
		c.handleAccessLog(ctx, response.Raw())
	}
	resRaw := response.ReadAll()
	return resRaw
}

// 请求前拼接完整url
func (c *Client) parseUrl(endpoint string) string {
	var (
		param = url.Values{}
		url   string
	)
	if c.UrlValues != nil {
		param = c.UrlValues
	}
	if c.BaseUri != "" {
		url = c.BaseUri + endpoint
	} else {
		url = endpoint
	}
	if len(param) > 0 {
		url = url + "?" + param.Encode()
	}
	return url
}

// handlelog
func (c *Client) handleAccessLog(ctx context.Context, raw string) {
	if !c.Logger.AccessLogEnabled {
		return
	}
	c.Logger.File(c.Logger.AccessLogPattern).Stdout(c.Logger.LogStdout).Printf(ctx, "\n================ Request Success Debug ===============\n\n%s\n", raw)
}

// hanle error
func (c *Client) handleErrorLog(ctx context.Context, err error, raw string) {
	if c.Logger == nil {
		return
	}
	if !c.Logger.ErrorLogEnabled {
		return
	}
	content := "\n\n [Client Error]:"
	if c.Logger.ErrorStack {
		if stack := gerror.Stack(err); stack != "" {
			content += "\nStack:\n" + stack
		}
	}
	content += err.Error()
	content += "\n =============== Response Raw [ERR] =============\n" + raw
	c.Logger.File(c.Logger.ErrorLogPattern).Stdout(c.Logger.LogStdout).Print(ctx, content)
}
