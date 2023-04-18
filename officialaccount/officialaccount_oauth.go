package officialaccount

import (
	"context"

	"gitee.com/wallesoft/oauth2"
	"gitee.com/wallesoft/oauth2/client/wechat"
)

// OAuth
// @see https://gitee.com/wallesoft/go/oauth2
func (oa *OfficialAccount) OAuth(ctx context.Context, config *wechat.Config) *oauth2.Oauth {
	if config.ClientID == "" {
		config.ClientID = oa.config.AppID
	}
	if config.ClientSecret == "" {
		config.ClientSecret = oa.config.Secret
	}

	return oauth2.New(ctx, config)
}
