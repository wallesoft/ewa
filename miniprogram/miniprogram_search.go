package miniprogram

import (
	"context"
	"fmt"

	"gitee.com/wallesoft/ewa/kernel/http"
	"github.com/gogf/gf/v2/frame/g"
)

//搜索

//小程序的站内搜商品图片搜索
func (mp *MiniProgram) ImageSearch(ctx context.Context, path string) *http.ResponseData {
	return &http.ResponseData{
		Json: mp.GetClientWithToken().RequestPost(ctx, "wxa/imagesearch", fmt.Sprintf("media=@file:%s", path)),
	}
}

//小程序内部搜索API提供针对页面的查询能力，小程序开发者输入搜索词后，将返回自身小程序和搜索词相关的页面。因此，利用该接口，开发者可以查看指定内容的页面被微信平台的收录情况；同时，该接口也可供开发者在小程序内应用，给小程序用户提供搜索能力。
func (mp *MiniProgram) SiteSearch(ctx context.Context, keyword string, nextPage ...string) *http.ResponseData {
	page := ""
	if len(nextPage) > 0 {
		page = nextPage[0]
	}
	return &http.ResponseData{
		Json: mp.GetClientWithToken().RequestJson(ctx, "POST", "wxa/sitesearch", g.Map{"keyword": keyword, "next_page_info": page}),
	}
}

//小程序开发者可以通过本接口提交小程序页面url及参数信息(不要推送webview页面)，让微信可以更及时的收录到小程序的页面信息，开发者提交的页面信息将可能被用于小程序搜索结果展示。
func (mp *MiniProgram) SubmitPages(ctx context.Context, pages g.Array) *http.ResponseData {
	return &http.ResponseData{
		Json: mp.GetClientWithToken().RequestJson(ctx, "POST", "wxa/search/wxaapi_submitpages", g.Map{"pages": pages}),
	}
}
