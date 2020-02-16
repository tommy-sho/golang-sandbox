package main

import (
	"fmt"
	"time"
)

type Job struct {
	Message string
}

var JobQueue chan Job

type Worker struct {
	WorkerPool chan chan Job
	JobChannel chan Job
	quit       chan bool
	number     int
}

func NewWorker(workerPool chan chan Job, number int) Worker {
	return Worker{
		WorkerPool: workerPool,
		JobChannel: make(chan Job),
		quit:       make(chan bool),
		number:     number,
	}
}

func (w Worker) Start() {
	go func() {
		for {
			// register the current worker into the worker queue.
			w.WorkerPool <- w.JobChannel

			select {
			case job := <-w.JobChannel:
				// we have received a work request.
				fmt.Println("worker: ", w.number, job.Message)
				time.Sleep(time.Millisecond)

			case <-w.quit:
				// we have received a signal to stop
				return
			}
		}
	}()
}

func (w Worker) Stop() {
	go func() {
		w.quit <- true
	}()
}
