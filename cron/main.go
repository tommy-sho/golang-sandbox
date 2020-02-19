package main

import (
	"fmt"
	"time"

	"github.com/jasonlvhit/gocron"
)

func print() {
	fmt.Println(": ", time.Now())
}
func main() {
	fmt.Println(time.Now())
	gocron.
		Every(10).Seconds().
		From(gocron.NextTick()).Do(print)
	gocron.Start()
	defer gocron.Clear()

	time.Sleep(time.Second * 30)
}
