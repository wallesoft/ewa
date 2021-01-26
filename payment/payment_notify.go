package payment

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"gitee.com/wallesoft/ewa/kernel/encryptor"
	ehttp "gitee.com/wallesoft/ewa/kernel/http"
	"gitee.com/wallesoft/ewa/kernel/log"
	"github.com/gogf/gf/encoding/gjson"
	"github.com/gogf/gf/encoding/gxml"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/text/gregex"
	"github.com/gogf/gf/util/gconv"
	"github.com/gogf/gf/util/gutil"
)

//支付通知
type Notify struct {
	Request  *ehttp.Request
	Response   *ehttp.Response
	bodyData []byte
	message *NotifyMessage // 解密后的数据
}
type NotifyMessage struct{
	*gjson.Json
}
//获取支付通知
func (p *Payment) Notify(r *http.Request, w http.ResponseWriter) *Notify{
	notify := &Notify{
		Request: &ehttp.Request{
			Request: r,
		},
		Resposne: &ehttp.Request{Request: r},
	}
	notify.bodyData = notify.Request.GetBody()
	//https://pay.weixin.qq.com/wiki/doc/apiv3/apis/chapter3_2_5.shtml
	//解密

	return notify
}

//支付通知处理
func (n *Notify) HandlePaid(f func(message *NotifyMessage)bool{}){

}

