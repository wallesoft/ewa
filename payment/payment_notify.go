package payment

import (
	"context"
	"fmt"
	"net/http"

	ehttp "gitee.com/wallesoft/ewa/kernel/http"
	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/util/gconv"
)

//支付通知
type Notify struct {
	Request  *ehttp.Request
	Response *ehttp.Response
	bodyData []byte
	Message  *NotifyMessage // 消息数据包含原始数据
	payment  *Payment
}
type NotifyMessage struct {
	Raw         *gjson.Json //原始数据，未解密 获取消息id，通知时间等用原始数据获取
	*gjson.Json             //解密后的数据 这里面不包含原始数据，只包含解密后的resource
}

type NotifyRes struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

const (
	MessageEventType = "TRANSACTION.SUCCESS" //支付通知状态
	MessageResCode   = "SUCCESS"             //成功返回到code值
	PaySuccess       = "SUCCESS"             //支付成功
	PayRefund        = "REFUND"              //转入退款
	PayNotPay        = "NOTPAY"              //未支付
	PayClosed        = "CLOSED"              //已关闭
	Paying           = "USERPAYING"          //正在支付中(付款码)
	PayRevoked       = "REVOKED"             //已撤销(付款码)
	PayError         = "PAYERROR"            //支付失败
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
	notify.Message = &NotifyMessage{
		Raw: gjson.New(notify.bodyData),
	}

	return notify
}

//支付通知处理
func (n *Notify) HandlePaid(ctx context.Context, f func(message *NotifyMessage) (bool, error)) {

	//验签
	err := n.payment.VerifySignature(n.Request.Header, n.bodyData)
	if err != nil {
		n.output(ctx, http.StatusInternalServerError, "UNVALID_SIGNATURE", err.Error())
	}
	data := n.Message.Raw //gjson.New(n.bodyData)
	if data.Get("event_type").String() == MessageEventType {
		decrypted, err := n.payment.GCMDecryte(data.Get("resource.associated_data").String(), data.Get("resource.ciphertext").String(), data.Get("resource.nonce").String())
		if err != nil {
			n.output(ctx, http.StatusInternalServerError, "AES_256_GCM_UNDECRYPTED", err.Error())
		}
		n.Message.Json = gjson.New(decrypted)
		if ok, err := f(n.Message); ok {
			n.output(ctx, http.StatusOK, MessageResCode, "")
		} else {
			n.output(ctx, http.StatusInternalServerError, "NOTIFY_MSG_HANDLE_ERROR", err.Error())
		}
	}

}

//output response output
func (n *Notify) output(ctx context.Context, status int, code string, message string) {
	if status != http.StatusOK {
		n.payment.Logger.File(n.payment.Logger.ErrorLogPattern).Print(ctx, fmt.Sprintf("[Erro] %s {%s} \n %s \n", code, message, gconv.String(n.bodyData)))
	}
	res := &NotifyRes{
		Code:    code,
		Message: message,
	}
	n.Response.WriteStatus(status, gjson.New(res).MustToJsonString())
	n.Response.Output()
}
