package payment

import (
	"fmt"
	"net/http"

	ehttp "gitee.com/wallesoft/ewa/kernel/http"
	"github.com/gogf/gf/encoding/gjson"
	"github.com/gogf/gf/util/gconv"
)

//支付通知
type Notify struct {
	Request  *ehttp.Request
	Response *ehttp.Response
	bodyData []byte
	message  *NotifyMessage // 解密后的数据
	payment  *Payment
}
type NotifyMessage struct {
	*gjson.Json
}

type NotifyRes struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

const (
	MessageEventType = "TRANSACTION.SUCCESS"
	MessageResCode   = "SUCCESS" //成功返回到code值
)

//获取支付通知
func (p *Payment) Notify(r *http.Request, w http.ResponseWriter) *Notify {
	notify := &Notify{
		payment: p,
		Request: &ehttp.Request{
			Request: r,
		},
		Response: ehttp.GetResponse(w),
	}
	notify.bodyData = notify.Request.GetBody()
	//https://pay.weixin.qq.com/wiki/doc/apiv3/apis/chapter3_2_5.shtml
	notify.message = &NotifyMessage{
		Json: gjson.New(notify.bodyData),
	}

	return notify
}

//支付通知处理
func (n *Notify) HandlePaid(f func(message *NotifyMessage) (bool, error)) {

	//验签
	err := n.payment.VerifySignature(n.Request.Header, n.bodyData)
	if err != nil {
		n.output(http.StatusInternalServerError, "UNVALID_SIGNATURE", err.Error())
	}
	data := gjson.New(n.bodyData)
	if data.GetString("event_type") == MessageEventType {
		decrypted, err := n.payment.GCMDecryte(data.GetString("resource.associated_data"), data.GetString("resource.ciphertext"), data.GetString("resource.nonce"))
		if err != nil {
			n.output(http.StatusInternalServerError, "AES_256_GCM_UNDECRYPTED", err.Error())
		}
		message := &NotifyMessage{
			Json: gjson.New(decrypted),
		}
		if ok, err := f(message); ok {
			n.output(http.StatusOK, MessageResCode, "")
		} else {
			n.output(http.StatusInternalServerError, "NOTIFY_MSG_HANDLE_ERROR", err.Error())
		}
	}

}

//output response output
func (n *Notify) output(status int, code string, message string) {
	if status != http.StatusOK {
		n.payment.Logger.File(n.payment.Logger.ErrorLogPattern).Print(fmt.Sprintf("[Erro] %s {%s} \n %s \n", code, message, gconv.String(n.bodyData)))
	}
	res := &NotifyRes{
		Code:    code,
		Message: message,
	}
	n.Response.WriteStatus(status, gjson.New(res).MustToJsonString())
	n.Response.Output()
}
