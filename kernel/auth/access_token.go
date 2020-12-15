package auth

//AccessToken interface
type AccessToken interface {
	GetToken() string
	GetTokenKey() string
}
