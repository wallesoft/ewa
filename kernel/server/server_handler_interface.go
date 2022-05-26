package server

import "context"

//Handler is interface to handle message
type Handler interface {
	Handle(ctx context.Context, message *Message) interface{}
}
