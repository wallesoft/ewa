package base

import "context"

type TokenCredentail interface {
	Get(ctx context.Context) map[string]string
}
