package server

type Guard interface {
	Resolve() bool
	ShouldReturnRawResponse() bool
}
