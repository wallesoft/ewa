package base

import (
	"context"
	"errors"
	"net/url"

	"gitee.com/wallesoft/ewa/kernel/auth"
	"gitee.com/wallesoft/ewa/kernel/log"
	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/util/gconv"

	"github.com/gogf/gf/v2/net/gclient"
)

type Client struct {
	Client    *gclient.Client
	BaseUri   string
	UrlValues url.Values
	Logger    *log.Logger
	Token     auth.AccessToken
}

//PostJson request by post method and return gjson.Json
// **** Deprecated use RequestJson instead
func (c *Client) PostJson(ctx context.Context, endpoint string, data ...interface{}) *gjson.Json {
	// var val interface{}
	// if len(data) > 0 {
	// 	val = data[0]
	// }
	response, err := c.Client.ContentJson().Post(ctx, c.getUri(ctx, endpoint), data...)
	var debugRaw string = response.Raw()
	if err != nil {
		c.handleErrorLog(ctx, err, debugRaw)
	}
	// c.handleAccessLog(response)
	defer response.Close()

	result := gjson.New(response.ReadAllString())

	if have := result.Contains("errcode"); have {
		//40001 refresh token
		if result.Get("errcode").Int() == 40001 {

			c.Token.GetToken(ctx, true)

			resp, err := c.Client.ContentJson().Post(ctx, c.getUri(ctx, endpoint), data...)
			var respRaw string = resp.Raw()
			if err != nil {
				c.handleErrorLog(ctx, err, respRaw)
			}
			res := gjson.New(resp.ReadAllString())
			defer resp.Close()
			if res.Contains("errcode") {
				c.handleErrorLog(ctx, errors.New("Refresh Token Result:"), respRaw)
			} else {
				c.handleAccessLog(ctx, respRaw)
				return res
			}

		}

		c.handleErrorLog(ctx, errors.New("get json with err code."), debugRaw)
		return result
	}
	c.handleAccessLog(ctx, debugRaw)
	return result
}

// //GetJson request by get method and return gjson.Json
// ***Deprecated use RequestJson instead
func (c *Client) GetJson(ctx context.Context, endpoint string, data ...interface{}) *gjson.Json {
	// var val interface{}
	// if len(data) > 0 {
	// 	val = data[0]
	// }
	response, err := c.Client.Get(ctx, c.getUri(ctx, endpoint), data...)
	var debugRaw string = response.Raw()
	if err != nil {
		c.handleErrorLog(ctx, err, debugRaw)
	}

	result := gjson.New(response.ReadAllString())

	if have := result.Contains("errcode"); have {
		//40001 refresh token
		if result.Get("errcode").Int() == 40001 {

			c.Token.GetToken(ctx, true)

			resp, err := c.Client.ContentJson().Get(ctx, c.getUri(ctx, endpoint), data...)
			var respRaw string = resp.Raw()
			if err != nil {
				c.handleErrorLog(ctx, err, respRaw)
			}
			res := gjson.New(resp.ReadAllString())
			defer resp.Close()
			if res.Contains("errcode") {
				c.handleErrorLog(ctx, errors.New("Refresh Token Result:"), respRaw)
			} else {
				c.handleAccessLog(ctx, respRaw)
				return res
			}

		}

		c.handleErrorLog(ctx, errors.New("get json with err code."), debugRaw)
		return result
	}
	c.handleAccessLog(ctx, debugRaw)
	return result
}

//request return json
func (c *Client) RequestJson(ctx context.Context, method string, endpoint string, data ...interface{}) *gjson.Json {

	if method == "POST" {
		c.Client = c.Client.ContentJson()
	}
	raw := c.RequestRaw(ctx, method, endpoint, data...)
	c.handleAccessLog(ctx, gconv.String(raw))
	return gjson.New(raw)
}

//request Post retrun Json !!!上传文件图片 内容安全等接口用 配合 @file:
func (c *Client) RequestPost(ctx context.Context, endpoint string, data ...interface{}) *gjson.Json {
	raw := c.RequestRaw(ctx, "POST", endpoint, data...)
	c.handleAccessLog(ctx, gconv.String(raw))
	return gjson.New(raw)
}

//Request return Raw
func (c *Client) RequestRaw(ctx context.Context, method string, endpoint string, data ...interface{}) []byte {
	var response *gclient.Response
	var err error
	response, err = c.Client.DoRequest(ctx, method, c.getUri(ctx, endpoint), data...)
	if err != nil {
		c.handleErrorLog(ctx, err, response.Raw())
	}
	debugRaw := response.Raw()
	resRaw := response.ReadAll()
	if gjson.Valid(resRaw) {
		result := gjson.New(resRaw)
		if have := result.Contains("errcode"); have {
			//40001 refresh token try once
			if result.Get("errcode").Int() == 40001 {

				c.Token.GetToken(ctx, true)

				resp, err := c.Client.ContentJson().Post(ctx, c.getUri(ctx, endpoint), data...)
				var respRaw string = resp.Raw()
				if err != nil {
					c.handleErrorLog(ctx, err, respRaw)
				}
				res := resp.ReadAll()
				defer resp.Close()
				if gjson.Valid(res) {
					resJson := gjson.New(res)
					if resJson.Contains("errcode") {
						c.handleErrorLog(ctx, errors.New("Refresh Token Result:"), respRaw)
					}
				}
			}
			if result.Get("errcode").Int() != 0 {
				c.handleErrorLog(ctx, errors.New("get json with err code."), debugRaw)
				return resRaw
			}

		}
	}
	return resRaw
}

func (c *Client) handleAccessLog(ctx context.Context, raw string) {
	if !c.Logger.AccessLogEnabled {
		return
	}
	c.Logger.File(c.Logger.AccessLogPattern).Stdout(c.Logger.LogStdout).Printf(ctx, "\n=============Response Raw============\n\n %s \n ", raw)
}

func (c *Client) handleErrorLog(ctx context.Context, err error, raw string) {
	if !c.Logger.ErrorLogEnabled {
		return
	}
	content := "\n\n [Error]:"
	if c.Logger.ErrorStack {
		if stack := gerror.Stack(err); stack != "" {
			content += "\nStack:\n" + stack
		} else {
			content += err.Error()
		}
	} else {
		content += err.Error()
	}
	content += "\n =============Reponse Raw [err] ==============\n" + raw
	c.Logger.File(c.Logger.ErrorLogPattern).Stdout(c.Logger.LogStdout).Print(ctx, content)
}

//getUri 将api与接口地址url进行拼接
func (c *Client) getUri(ctx context.Context, endpoint string) string {

	var param = url.Values{}
	var url string
	//uri params
	if c.UrlValues != nil {
		param = c.UrlValues
	}
	//token
	if c.Token != nil {
		param.Add(c.Token.GetRequestTokenKey(), c.Token.GetToken(ctx))
	}
	//base uri
	if c.BaseUri != "" {
		url = c.BaseUri + endpoint
	} else {
		url = endpoint
	}

	if param != nil && len(param) > 0 {
		url = url + "?" + param.Encode()
	}
	return url

}
