package server

//Handler is interface to handle message
type Handler interface {
	Handle(message *Message) interface{}
}
