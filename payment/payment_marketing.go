package payment

import "github.com/gogf/gf/encoding/gjson"

//营销
type Marketing struct {
	payment *Payment
}

//红包
func (m *Marketing) Redpack() *Redpack {
	r := &Redpack{
		config:  gjson.New(nil),
		payment: m.payment,
	}

	r.config.Set("mch_id", m.payment.config.MchID)
	r.config.Set("wxappid", m.payment.config.AppID) // 默认appid，小程序可在配置处修改

	return r
}
