package auth

import "context"

//AccessToken interface
type AccessToken interface {
	GetToken(ctx context.Context, refresh ...bool) string
	GetTokenKey() string
	GetRequestTokenKey() string
}
