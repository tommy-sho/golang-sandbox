package main

import "fmt"

type Animal interface {
	Bow(string)
}

type A struct{}

func (a *A) Bow(s string) {
	fmt.Println("Hello,", s)
}

type B struct {
	*A
}

func main() {
	a := &A{}
	b := B{a}
	b.Bow("unun")
	var _ Animal = b
}
