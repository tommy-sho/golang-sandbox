package main

import "fmt"

func part(ch chan<- []interface{}, lhs []interface{}, original []interface{}) {
	if len(original) == 0 {
		l := make([]interface{}, len(lhs), cap(lhs))
		copy(l, lhs)
		ch <- l
		return
	}

	for i, v := range original {
		l := make([]interface{}, len(original), cap(original))
		copy(l, original)

		l = append(l[:i], l[i+1:]...)
		lhs := append(lhs, v)
		part(ch, lhs, l)
		fmt.Println("-------------")
	}
}

func perm(original []interface{}) <-chan []interface{} {
	ch := make(chan []interface{})
	go func() {
		defer close(ch)
		buf := make([]interface{}, 0, len(original))
		part(ch, buf, original)
	}()
	return ch
}

func main() {
	for p := range perm([]interface{}{"ピ", "カ", "チュウ"}) {
		fmt.Printf("%s%s%s\n", p...)
	}

	ss := []string{"1", "2", "3", "4", "5"}
	for i, v := range ss {
		c := make([]string, len(ss), cap(ss))
		copy(c, ss)
		c = append(c[:i], c[i+1:]...)
		fmt.Println("i: ", i, ", v: ", v, " :", c)
	}
}
