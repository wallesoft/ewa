package officialaccount

import (
	"fmt"
	"time"

	"github.com/gogf/gf/container/gvar"
	"github.com/gogf/gf/os/gcache"
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
func (v *DefaultJsapiTicket) GetTicket() string {
	ticket, err := v.cache.Get(v.getKey())
	if err != nil {
		panic(err.Error())
	}
	if ticket == nil {
		//get & cache
		client := v.oa.getClientWithToken()
		val := client.RequestJson("GET", "cgi-bin/ticket/getticket", "type=jsapi")
		if val.GetInt("errcode") == 0 && val.Contains("ticket") {
			if err := v.cache.Set(v.getKey(), val.GetString("ticket"), time.Second*7200); err != nil {
				panic(err.Error())
			}
			return val.GetString("ticket")
		} else {
			v.oa.Logger.Stdout(v.oa.Logger.LogStdout).Print(fmt.Sprintf("[Err] ticket get from api Error: %s", val.MustToJsonString()))
		}
	}
	return gvar.New(ticket).String()
}

func (v *DefaultJsapiTicket) getKey() string {
	return "ewa.officialaccount.jsapi_ticket." + v.oa.config.AppID
}
