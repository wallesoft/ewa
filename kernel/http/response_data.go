package http

import "github.com/gogf/gf/v2/encoding/gjson"

//ResponseData 针对微信接口返回json格式数据设置，errcode errmsg
type ResponseData struct {
	*gjson.Json
}

//Ok 当返回errcode=0， errmsg="ok" 时为true errcode != 0 时为false
func (r *ResponseData) HaveError() bool {
	if r.Contains("errcode") {
		return r.Get("errcode").Int() != 0
	}
	return false
}
