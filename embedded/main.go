package main

import "fmt"

type Parent struct {
}

func (p Parent) Name() string {
	return "I'm parent"
}

type Child struct {
	Parent
}

func (c Child) Name() string {
	return "I'm child"
}

func main() {
	s := Child{}
	fmt.Println(s.Name())
}
