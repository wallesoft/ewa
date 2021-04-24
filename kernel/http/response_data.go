package http

import "github.com/gogf/gf/encoding/gjson"

//ResponseData 针对微信接口返回json格式数据设置，errcode errmsg
type ResponseData struct {
	*gjson.Json
}

//Ok 当返回errcode=0， errmsg="ok" 时为true errcode != 0 时为false
func (r *ResponseData) HaveError() bool {
	return r.GetInt("errcode") != 0
}
