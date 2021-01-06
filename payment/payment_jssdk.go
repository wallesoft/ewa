package payment

import (
	"fmt"

	"github.com/gogf/gf/container/gvar"
	"github.com/gogf/gf/os/gtime"
	"github.com/gogf/gf/util/grand"
)

type Jssdk struct {
	AppID     string
	Timestamp string
	NonceStr  string
	Package   string
	SignType  string
	PaySign   string
}

//Jssdk
func (p *Payment) Jssdk(prepayId string) *Jssdk {
	jssdk := &Jssdk{
		AppID:     p.config.AppID,
		Timestamp: gtime.TimestampStr(),
		NonceStr:  grand.S(32),
		Package:   "prepay_id=" + prepayId,
		SignType:  "RSA",
	}
	jssdk.setPaySign(p)
	return jssdk
}
func (j *Jssdk) setPaySign(payment *Payment) {
	var err error
	signStr := fmt.Sprintf("%s\n%s\n%s\n%s\n", j.AppID, j.Timestamp, j.NonceStr, j.Package)
	j.PaySign, err = payment.rsaEncrypt(gvar.New(signStr).Bytes())
	if err != nil {
		panic(fmt.Sprintf("[Erro] payment Jssdk set paysign: %s", err.Error()))
	}

}
