package miniprogram

import (
	"context"
	"fmt"

	"gitee.com/wallesoft/ewa/kernel/http"
	"github.com/gogf/gf/v2/frame/g"
)

//CheckText @see https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/sec-check/security.msgSecCheck.html
func (mp *MiniProgram) CheckText(ctx context.Context, content string) *http.ResponseData {
	return &http.ResponseData{
		Json: mp.GetClientWithToken().RequestJson(ctx, "POST", "wxa/msg_sec_check", g.Map{"content": content}),
	}
}

//CheckImage @see https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/sec-check/security.imgSecCheck.html
func (mp *MiniProgram) CheckImage(ctx context.Context, path string) *http.ResponseData {
	return &http.ResponseData{
		Json: mp.GetClientWithToken().RequestPost(ctx, "wxa/img_sec_check", fmt.Sprintf("media=@file:%s", path)),
	}
}

//CheckMediaAsync 异步检查图片音频 !!!异步通知 @see https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/sec-check/security.mediaCheckAsync.html
func (mp *MiniProgram) CheckMediaAsync(ctx context.Context, url string, mediaType int) *http.ResponseData {
	return &http.ResponseData{
		Json: mp.GetClientWithToken().RequestJson(ctx, "POST", "wxa/media_check_async", g.Map{"media_url": url, "media_type": mediaType}),
	}
}
