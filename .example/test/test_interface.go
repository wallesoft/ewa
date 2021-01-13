package main

import (
	"fmt"
)

type A struct {
	Name string
}

type B struct {
	// *A
	Title string
}

type I interface {
	T(i I)
}

func (a *A) Test(name string) {
	fmt.Sprintln("A Test:%s", name)
}
func (a *B) Test(name string) {
	fmt.Sprintln("B Test:%s", name)
}
func (a *A) T(i I) {
	i.T(i)
}
func (b *B) T(i I) {
	i.T(i)
}

func main() {
	a := &A{
		Name: "A",
	}
	b := &B{
		// A:     a,
		Title: "B",
	}
	a.T(b)
}
