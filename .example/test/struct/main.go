package main

import "fmt"

type A struct {
	name string
}

func (a *A) GetName() string {
	return a.name
}
func (a *A) GetFrom() string {
	return "i am from A"
}

type B struct {
	name string
	age  int
}

type C struct {
	*A
	address string
}

func (b *B) GetAge() int {
	return b.age
}

func (c *C) GetAdress() string {
	return c.address
}

func (c *C) GetName() string {
	return "i am C"
}

func main() {
	a := &A{name: "i am A"}
	c := &C{A: a}
	fmt.Println(fmt.Sprintf("%s", c.A.GetName()))
	fmt.Println(c.GetFrom())
}
