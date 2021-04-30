package server

type Guard interface {
	Resolve(message *Message) bool
	ShouldReturnRawResponse() bool
}
