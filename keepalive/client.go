package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"sync"
	"time"

	"github.com/tcnksm/go-httpstat"
)

func main() {

	tr := &http.Transport{
		DialContext: (&net.Dialer{
			Timeout:   30 * time.Second,
			KeepAlive: 30 * time.Second,
		}).DialContext,
		MaxIdleConnsPerHost: 1024,
		TLSHandshakeTimeout: 0 * time.Second,
		DisableKeepAlives:   false,
	}
	client := &http.Client{Transport: tr}
	sendRequest(client, 0)

	fmt.Println()
	wg := &sync.WaitGroup{}
	for i := 1; i < 10; i++ {
		wg.Add(1)
		time.Sleep(time.Millisecond * 10)
		go func(client *http.Client, i int) {
			sendRequest(client, i)
			wg.Done()
		}(client, i)
	}

	wg.Wait()
}

func sendRequest(c *http.Client, i int) {
	req, err := http.NewRequest("GET", "http://google.com", nil)
	if err != nil {
		log.Fatal(err)
	}
	// Create go-httpstat powered context and pass it to http.Request
	var result httpstat.Result
	result.End(time.Now())

	ctx := httpstat.WithHTTPStat(req.Context(), &result)
	req = req.WithContext(ctx)

	res, err := c.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	if _, err := io.Copy(ioutil.Discard, res.Body); err != nil {
		log.Fatal(err)
	}
	res.Body.Close()
	result.End(time.Now())

	// Show results
	fmt.Printf("--- %d 回目のリクエスト ---\n", i)
	fmt.Printf("%+v\n", result)
}
