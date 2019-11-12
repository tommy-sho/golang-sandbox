package main

import (
	"strconv"
	"time"
)

func main() {
	dispatcher := NewDispatcher(10)
	dispatcher.Run()
	go dispatcher.dispatch()

	time.Sleep(time.Second)
	for i := 1; i < 100; i++ {
		t := strconv.Itoa(i)
		job := Job{Message: "hello" + t}
		JobQueue <- job
	}
	time.Sleep(time.Second * 10)
}
