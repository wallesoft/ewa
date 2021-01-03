package payment

import (
	"github.com/gogf/gf/encoding/gjson"
)

type Order struct {
	config *gjson.Json
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

//Set
func (o *Order) Set(pattern string, value interface{}) {
	o.config.Set(pattern, value)
}

//Jsapi 下单
func (o *Order) Jsapi() string {
	return o.config.MustToJsonString()
}

//H5下单
func (o *Order) H5() {

}

//Query 订单查询
func (o *Order) Query() {

}
