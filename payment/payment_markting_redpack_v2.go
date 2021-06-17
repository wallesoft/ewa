package payment

import (
	"strings"

	"github.com/gogf/gf/encoding/gjson"
	"github.com/gogf/gf/util/grand"
)

// 营销红包 v2
// @see https://pay.weixin.qq.com/wiki/doc/api/tools/cash_coupon.php?chapter=13_4&index=3
type Redpack struct {
	config  *gjson.Json
	payment *Payment
}

// 配置 传入需要参数 校验参数
func (r *Redpack) New(config map[string]interface{}) *Redpack {
	for pattern, val := range config {
		r.Set(pattern, val)
	}

	//随机字符串
	if !r.config.Contains("nonce_str") {
		r.Set("nonce_str", strings.ToUpper(grand.S(32)))
	}

	//signature
	r.Set("sign", r.payment.V2MD5(r.config.ToMap()))

	return r
}

//设置参数
func (r *Redpack) Set(pattern string, value interface{}) {
	r.config.Set(pattern, value)
}

// 红包发送
func (r *Redpack) Send() {

}
