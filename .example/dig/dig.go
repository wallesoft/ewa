package main

import (
	"go.uber.org/dig"
)

// var Container dig.Container = dig.New()
var Container *dig.Container = dig.New()

// type A struct{
// 	Access Access
// 	Token Token
// }

type Access struct {
	C *dig.Container
	V string
}
type Token struct {
	V string
}

func InitToken() *Token {
	return &Token{
		V: "token",
	}
}

func InitC() *dig.Container {
	return Container
}

func InitAccess(c *dig.Container) *Access {
	return &Access{
		C: c,
		V: "access",
	}
}

func (a *Access) Test() {

	// g.Dump(a.C.Invoke((*Token).P))
}

func (t *Token) P() string {
	return t.V
}

func main() {
	Container.Provide(InitC)
	Container.Provide(InitAccess)
	Container.Invoke((*Access).Test)
}
