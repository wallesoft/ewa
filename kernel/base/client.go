package base

import (
	"net/url"

	"gitee.com/wallesoft/ewa/kernel/auth"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/os/glog"
)

type Client struct {
	*ghttp.Client
	BaseUri    string
	QueryParam *url.Values
	Logger     *glog.Logger
	Token      auth.AccessToken
}

//PostJson request by post method and return gjson.Json
func (c *Client) PostJson(endpoint string, data ...interface{}) string {
	var val interface{}
	if len(data) > 0 {
		val = data[0]
	}
	response, err := c.ContentJson().Post(c.getUri(endpoint), val)
	if err != nil {
		panic(err.Error())
	}
	c.debug(response)
	defer response.Close()

	return response.ReadAllString()

}

// //GetJson request by get method and return gjson.Json
func (c *Client) GetJson(endpoint string, data ...interface{}) string {
	var val interface{}
	if len(data) > 0 {
		val = data[0]
	}
	response, err := c.Get(c.getUri(endpoint), val)
	if err != nil {
		panic(err.Error())
	}
	c.debug(response)
	defer response.Close()
	return response.ReadAllString()
}

func (c *Client) debug(response *ghttp.ClientResponse) {
	c.Logger.Debug(g.Map{
		"Client Request:": g.Map{
			"Request:":  response.RawRequest(),
			"Response:": response.RawResponse(),
			// "Content:":  response.Raw(),
		},
	})
}

//getUri
func (c *Client) getUri(endpoint string) string {
	param := &url.Values{}
	if c.BaseUri != "" && c.Token != nil {
		param.Add(c.Token.GetTokenKey(), c.Token.GetToken())
		return c.BaseUri + endpoint + "?" + param.Encode()
	} else if c.BaseUri != "" {
		return c.BaseUri + endpoint
	}

	return endpoint
}
