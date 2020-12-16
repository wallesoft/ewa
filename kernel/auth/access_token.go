package auth

//AccessToken interface
type AccessToken interface {
	GetToken(refresh ...bool) string
	GetTokenKey() string
}
