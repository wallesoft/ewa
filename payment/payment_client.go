package payment

import (
	"github.com/gogf/gf/container/gmap"
	"github.com/gogf/gf/encoding/gjson"
	"github.com/gogf/gf/util/guid"
)

type Client struct {
	payment *Payment
}

func (c *Client) post(endpoint string, parmas map[string]string) *gjson.Json {
	param := gmap.NewStrStrMap(params)
	base := gmap.NewStrStrMap(map[string]string{
		"mch_id":     c.payment.config.MchID,
		"nonce_str":  guid.S(),
		"sub_mch_id": c.payment.config.SubMchID,
		"sub_appid":  c.payment.config.SubAppID,
	})
	requestParam = param.Merge(base)
	requestParam.FliterEmpty()
	secretKey := c.payment.getKey(endpoint)

}

func (c *Client) get(endpoint string, params map[string]string) *gjson.Json {

}

func (c *Client) request(endpoint string, params map[string]string, method string) {

}
