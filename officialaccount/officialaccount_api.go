package officialaccount

//GetToken
func (oa *OfficialAccount) GetToken() string {
	return oa.accessToken.GetToken()
}
