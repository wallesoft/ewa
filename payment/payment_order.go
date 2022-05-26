package payment

import (
	"context"
	"errors"
	"fmt"
	"net/url"

	"github.com/gogf/gf/v2/container/gvar"
	"github.com/gogf/gf/v2/encoding/gjson"
)

type Order struct {
	config  *gjson.Json
	payment *Payment
}

//订单
type OrderConfig struct {
	AppID string `json:"appid"`
	MchID string `json:"mchid"`
	// Description string    `json:"description"`
	// OutTradeNo  string    `json:"out_trade_no"`
	// TimeExpire  string    `json:"time_expire"`
	// Attach      string    `json:"attach"`
	// NotifyUrl   string    `json:"notify_url"`
	// GoodsTag    string    `json:"goods_tag"`
	// Amout       Amount    `json:"amount"`
	// Payer       Payer     `json:"payer"`
	// Detail      Detail    `json:"detail"`
	// SceneInfo   SceneInfo `json:"scene_info"`
}

// //订单金额
// type Amount struct {
// 	Total    int    `json:"total"`
// 	Currency string `json:"currency"`
// }

// //支付者
// type Payer struct {
// 	OpenID string `json:"openid"`
// }

// //优惠功能
// type Detail struct {
// 	CostPrice   int         `json:"cost_price"`
// 	InvoiceID   string      `json:"invoice_id"`
// 	GoodsDetail GoodsDetail `json:"goods_detail"`
// }

// //单品列表
// type GoodsDetail struct {
// 	MerchantGoodsID  string `json:"merchant_goods_id"`
// 	WechatpayGoodsID string `json:"wechatpay_goods_id"`
// 	GoodsName        string `json:"goods_name"`
// 	Quantity         int    `json:"quantity"`
// 	UnitPrice        int    `json:"unit_price"`
// }

// //场景信息
// type SceneInfo struct {
// 	PayerClientIP string    `json:"payer_client_ip"`
// 	DeviceID      string    `json:"device_id"`
// 	StoreInfo     StoreInfo `json:"store_info"`
// }

// //门店信息
// type StoreInfo struct {
// 	ID       string `json:"id"`
// 	Name     string `json:"name"`
// 	AreaCode string `json:"area_code"`
// 	Address  string `json:"address"`
// }

//QueryOrder @see https://pay.weixin.qq.com/wiki/doc/apiv3/apis/chapter3_1_2.shtml 返回参数
type QueryOrder struct {
	Code int
	*gjson.Json
}

//Set
func (o *Order) Set(pattern string, value interface{}) {
	o.config.Set(pattern, value)
}

//Jsapi 下单
func (o *Order) Jsapi(ctx context.Context) (string, error) {
	response := o.payment.getClient().RequestJson(ctx, "POST", "/v3/pay/transactions/jsapi", o.config.MustToJsonString())
	if response.StatusCode == 200 {
		return gjson.New(response.Body).Get("prepay_id").String(), nil
	}
	return gvar.New(response.Body).String(), errors.New(fmt.Sprintf("[Error] %s", response.Status))
}

//H5下单
func (o *Order) H5() {

}

//Query 订单查询
func (o *Order) Query(ctx context.Context) *QueryOrder {
	client := o.payment.getClient()
	client.UrlValues = url.Values{}
	client.UrlValues.Add("mchid", o.payment.config.MchID)
	var response *Response
	if o.config.Contains("transaction_id") {
		//根据微信支付订单号查询
		response = client.RequestJson(ctx, "GET", "/v3/pay/transactions/id/"+o.config.Get("transaction_id").String())
		if response.StatusCode == 200 {
			return &QueryOrder{
				Code: response.StatusCode,
				Json: gjson.New(response.Body),
			}
		}

	}
	if o.config.Contains("out_trade_no") {
		//根据商户订单号查询
		response = client.RequestJson(ctx, "GET", "/v3/pay/transactions/out-trade-no/"+o.config.Get("out_trade_no").String())
		if response.StatusCode == 200 {
			return &QueryOrder{
				Code: 200,
				Json: gjson.New(response.Body),
			}
		}

	}
	return &QueryOrder{
		Code: response.StatusCode,
		Json: gjson.New(response.Body),
	}
}
