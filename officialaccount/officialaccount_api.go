package officialaccount

import "context"

//GetToken
func (oa *OfficialAccount) GetToken(ctx context.Context) string {
	return oa.accessToken.GetToken(ctx)
}
