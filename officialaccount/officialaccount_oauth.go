package officialaccount

import (
	"gitee.com/wallesoft/go/oauth2"
)

//OAuth
//@see https://gitee.com/wallesoft/go/oauth2
func (oa *OfficialAccount) OAuth() *oauth2.WechatConfig {
	return oauth2.Wechat(oauth2.Config{
		ClientID:     oa.config.AppID,
		ClientSecret: oa.config.Secret,
	})
}
