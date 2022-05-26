package officialaccount

import (
	"context"
	"fmt"
	"time"

	"github.com/gogf/gf/v2/container/gvar"
	"github.com/gogf/gf/v2/os/gcache"
)

type JsapiTicket interface {
	GetTicket() string
}

type DefaultJsapiTicket struct {
	oa    *OfficialAccount
	cache *gcache.Cache
}

func (oa *OfficialAccount) JsapiTicket() *DefaultJsapiTicket {
	return &DefaultJsapiTicket{
		oa:    oa,
		cache: oa.config.Cache}
}
func (v *DefaultJsapiTicket) GetTicket(ctx context.Context) string {
	ticket, err := v.cache.Get(ctx, v.getKey())
	if err != nil {
		panic(err.Error())
	}
	if ticket == nil {
		//get & cache
		client := v.oa.GetClientWithToken()
		val := client.RequestJson(ctx, "GET", "cgi-bin/ticket/getticket", "type=jsapi")
		if val.Get("errcode").Int() == 0 && val.Contains("ticket") {
			if err := v.cache.Set(ctx, v.getKey(), val.Get("ticket").String(), time.Second*7200); err != nil {
				panic(err.Error())
			}
			return val.Get("ticket").String()
		} else {
			v.oa.Logger.Stdout(v.oa.Logger.LogStdout).Print(ctx, fmt.Sprintf("[Err] ticket get from api Error: %s", val.MustToJsonString()))
		}
	}
	return gvar.New(ticket).String()
}

func (v *DefaultJsapiTicket) getKey() string {
	return "ewa.officialaccount.jsapi_ticket." + v.oa.config.AppID
}
