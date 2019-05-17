package main

import (
	"fmt"
	"sync"
)

/*
Related article https://blog.kaneshin.co/entry/2016/08/18/190435
*/

// Dispatcher
type Dispatcher struct {
	// pool is channel pool for workers.
	pool chan *worker
	// queue is queue to store some request or data.
	queue chan interface{}
	//workers
	workers []*worker
	wg      sync.WaitGroup
	quit    chan struct{}
}

func (d *Dispatcher) Start() {
	for _, w := range d.workers {
		w.start()
	}

	go func() {
		for {
			select {
			// when data is queued
			case v := <-d.queue:
				(<-d.pool).dest <- v // send data to destination of worker pool.
			case <-d.quit:
				return
			}
		}
	}()
}

func (d *Dispatcher) Send(v interface{}) {
	d.wg.Add(1)
	d.queue <- v
}

func (d *Dispatcher) Sends(v []interface{}) {
	d.wg.Add(len(v))
	for t := range v {
		d.queue <- t
	}
}

func (d *Dispatcher) Wait() {
	d.wg.Wait()
}

type worker struct {
	dispatcher *Dispatcher
	//dest is receiver for dispatcher.
	dest chan interface{} // if set queue on worker, change channel size.
	quit chan struct{}
}

func (w *worker) start() {
	go func() {
		for {
			w.dispatcher.pool <- w // send themselves to dispatcher pool.

			select {
			case v := <-w.dest:
				//do something
				fmt.Println(v)

				w.dispatcher.wg.Done() // Count down wait-group of dispatcher.

			case <-w.quit:
				return
			}
		}
	}()
}

func NewDispatcher(maxWorkers, maxQueues int) *Dispatcher {
	d := &Dispatcher{
		pool:  make(chan *worker, maxWorkers),
		queue: make(chan interface{}, maxQueues),
		quit:  make(chan struct{}),
	}

	// worker の初期化
	d.workers = make([]*worker, cap(d.pool))
	for i := 0; i < cap(d.pool); i++ {
		w := worker{
			dispatcher: d,
			dest:       make(chan interface{}),
			quit:       make(chan struct{}),
		}
		d.workers[i] = &w
	}
	return d
}

func main() {
	d := NewDispatcher(3, 10000)

	d.Start()
	for i := 0; i < 100; i++ {
		meg := fmt.Sprintf("This is %v message", i)
		d.Send(meg)
	}
	d.Wait()
}
