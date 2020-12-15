package base

import (
	"net/url"

	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/os/glog"
)

type Client struct {
	*ghttp.Client
	BaseUri    string
	QueryParam *url.Values
	Logger     *glog.Logger
	// RequestPost bool
}

//PostJson request by post method and return gjson.Json
func (c *Client) PostJson(endpoint string, data ...interface{}) string {
	response, err := c.ContentJson().Post(c.getUri(endpoint), data)
	if err != nil {
		panic(err.Error())
	}
	c.debug(response)
	defer response.Close()
	return response.ReadAllString()
}

// //GetJson request by get method and return gjson.Json
func (c *Client) GetJson(endpoint string, data ...interface{}) string {
	response, err := c.Get(c.getUri(endpoint), data)
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
			"Content:":  response.Raw(),
		},
	})
}

//getUri
func (c *Client) getUri(endpoint string) string {
	if c.BaseUri != "" {
		return c.BaseUri + endpoint
	}
	if c.QueryParam != nil {
		return c.BaseUri + endpoint + "?" + c.QueryParam.Encode()
	}
	return endpoint
}
