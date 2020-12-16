package base

import (
	"errors"
	"net/url"

	"gitee.com/wallesoft/ewa/kernel/auth"
	"gitee.com/wallesoft/ewa/kernel/log"
	"github.com/gogf/gf/encoding/gjson"
	"github.com/gogf/gf/errors/gerror"

	// "github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
)

type Client struct {
	*ghttp.Client
	BaseUri   string
	UrlValues *url.Values
	Logger    *log.Logger
	Token     auth.AccessToken
}

//PostJson request by post method and return gjson.Json
func (c *Client) PostJson(endpoint string, data ...interface{}) *gjson.Json {
	var val interface{}
	if len(data) > 0 {
		val = data[0]
	}
	response, err := c.ContentJson().Post(c.getUri(endpoint), val)
	var debugRaw string = response.Raw()
	if err != nil {
		c.handleErrorLog(err, debugRaw)
	}
	// c.handleAccessLog(response)
	defer response.Close()

	result := gjson.New(response.ReadAllString())

	if have := result.Contains("errcode"); have {
		//40001 refresh token
		if result.GetInt("errcode") == 40001 {

			c.Token.GetToken(true)

			resp, err := c.ContentJson().Post(c.getUri(endpoint), val)
			var respRaw string = resp.Raw()
			if err != nil {
				c.handleErrorLog(err, respRaw)
			}
			res := gjson.New(resp.ReadAllString())
			defer resp.Close()
			if res.Contains("errcode") {
				c.handleErrorLog(errors.New("Refresh Token Result:"), respRaw)
			} else {
				c.handleAccessLog(respRaw)
				return res
			}

		}

		c.handleErrorLog(errors.New("get json with err code."), debugRaw)
		return result
	}
	c.handleAccessLog(debugRaw)
	return result
}

// //GetJson request by get method and return gjson.Json
func (c *Client) GetJson(endpoint string, data ...interface{}) *gjson.Json {
	var val interface{}
	if len(data) > 0 {
		val = data[0]
	}
	response, err := c.Get(c.getUri(endpoint), val)
	var responseRaw string = response.Raw()
	if err != nil {
		c.handleErrorLog(err, responseRaw)
	}
	c.handleAccessLog(responseRaw)
	defer response.Close()
	return gjson.New(response.ReadAllString())
}

func (c *Client) handleAccessLog(raw string) {
	if !c.Logger.AccessLogEnabled {
		return
	}
	c.Logger.File(c.Logger.AccessLogPattern).
		Stdout(c.Logger.LogStdout).
		Printf("\n=============Response Raw============\n\n %s \n ", raw)
}

func (c *Client) handleErrorLog(err error, raw string) {
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
	c.Logger.
		File(c.Logger.ErrorLogPattern).
		Stdout(c.Logger.LogStdout).
		Print(content)
}

//getUri
func (c *Client) getUri(endpoint string) string {

	var param = &url.Values{}
	var url string
	//uri params
	if c.UrlValues != nil {
		param = c.UrlValues
	}
	//token
	if c.Token != nil {
		param.Add(c.Token.GetTokenKey(), c.Token.GetToken())
	}
	//base uri
	if c.BaseUri != "" {
		url = c.BaseUri + endpoint
	} else {
		url = endpoint
	}
	if param != nil {
		url = url + "?" + param.Encode()
	}
	return url

}
