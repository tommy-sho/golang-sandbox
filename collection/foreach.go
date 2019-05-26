package main

import (
	"fmt"
)

func (s Seq) ForEach(f func(string) error) error {
	for i := range s {
		if err := f(s[i]); err != nil {
			return err
		}
	}
	return nil
}

func main() {
	f := func(s string) error {
		_, e := fmt.Println(s)
		if e != nil {
			return e
		}
		return nil
	}
	seq := Seq{"1", "2", "3"}
	seq.ForEach(f)
}
