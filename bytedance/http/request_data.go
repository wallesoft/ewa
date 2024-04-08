package http

import "github.com/gogf/gf/v2/encoding/gjson"

//ResponseData 针对抖音接口返回json格式数据设置，err_no err_tips
type ResponseData struct {
	*gjson.Json
}

//HaveError Ok 当返回errc_no=0， err_tips="success" 时为false err_no != 0 时为true
func (r *ResponseData) HaveError() bool {
	if r.Contains("err_no") {
		return r.Get("err_no").Int() != 0
	}
	return false
}
