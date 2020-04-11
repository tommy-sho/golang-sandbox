package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"sync"
)

var statusMap = make(map[string]int, 10)
var mu sync.Mutex
var wg = &sync.WaitGroup{}

func main() {

	for i := 0; i < 10000; i++ {
		wg.Add(1) // wgをインクリメント
		go sendRequest(i)
	}
	wg.Wait()
	for status, count := range statusMap {
		fmt.Printf("status: %s - %v\n", status, count)
	}
}

func sendRequest(i int) {
	mu.Lock()
	defer mu.Unlock()

	word := fmt.Sprintf("%v", i)
	res, err := http.Get(fmt.Sprint("http://localhost:8080/greeting?name=hoge") + word)
	if err != nil {
		panic(err)
	}

	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}
	fmt.Printf(string(data))
	statusMap[res.Status]++
	wg.Done()
}
