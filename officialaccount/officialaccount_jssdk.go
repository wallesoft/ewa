package officialaccount

import (
	"sort"
	"strings"

	"github.com/gogf/gf/crypto/gsha1"
	"github.com/gogf/gf/encoding/gjson"
	"github.com/gogf/gf/os/gtime"
	"github.com/gogf/gf/text/gstr"
	"github.com/gogf/gf/util/grand"
)

type Jssdk struct {
	Debug     bool     `json:"debug"`
	AppID     string   `json:"appId"`
	Timestamp string   `json:"Timestamp"`
	NonceStr  string   `json:"nonceStr"`
	Signature string   `json:"signature"`
	JsApiList []string `json:"jsApiList"`
}

func (oa *OfficialAccount) Jssdk(url string, list []string) *Jssdk {
	ticket := oa.JsapiTicket().GetTicket()
	timestamp := gtime.TimestampStr()
	nonceStr := grand.S(16)

	j := &Jssdk{
		Debug:     false,
		AppID:     oa.config.AppID,
		Timestamp: timestamp,
		NonceStr:  nonceStr,
		// Signature: ,
		JsApiList: list,
	}
	j.setSignature(url, ticket)
	return j
}

func (j *Jssdk) SetDebug(debug ...bool) {
	if len(debug) > 0 {
		j.Debug = debug[0]
	}
	// return j
}

//JsonString
func (j *Jssdk) JsonString() string {
	return gjson.New(j).MustToJsonString()
}

func (j *Jssdk) setSignature(url, ticket string) {
	if gstr.Contains(url, "#") {
		url = gstr.Split(url, "#")[0]
	}
	str := []string{"nonceStr=" + j.NonceStr, "jsapi_ticket" + ticket, "timestamp" + j.Timestamp, "url" + url}
	//sort
	sort.Strings(str)
	j.Signature = gsha1.Encrypt(strings.Join(str, "&"))
}
