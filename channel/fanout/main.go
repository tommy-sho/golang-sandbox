package main

import "fmt"

func Merge(out chan int, a, b <-chan int) {
	for a != nil && b != nil {
		select {
		case v, ok := <-a:
			out <- v
			if !ok {
				a = nil
			}
		case v, ok := <-b:
			out <- v
			if !ok {
				b = nil
			}
		}
	}
	close(out)
}

func main() {
	a := make(chan int)
	b := make(chan int)
	out := make(chan int)
	defer func() {
		close(a)
		close(b)
	}()
	go Merge(out, a, b)
	//go func() {
	//	var ok bool
	//	var v int
	//	for !ok {
	//		select {
	//		case v, ok = <-out:
	//			fmt.Println(v)
	//		}
	//	}
	//}()
	go func() {
		for {
			v, ok := <-out
			fmt.Println("receive: ", v)
			if !ok {
				break
			}
		}
	}()
	for i := 0; i < 10; i++ {
		a <- i
		b <- i
	}
}
