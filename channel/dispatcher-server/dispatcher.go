package main

type Dispatcher struct {
	WorkerPool chan chan Job
	MaxWorkers int
	stop       chan bool
}

func NewDispatcher(maxWorkers int) *Dispatcher {
	JobQueue = make(chan Job, 10)
	pool := make(chan chan Job, maxWorkers)
	stop := make(chan bool)
	return &Dispatcher{WorkerPool: pool, MaxWorkers: maxWorkers, stop: stop}
}

func (d *Dispatcher) Run() {
	for i := 0; i < d.MaxWorkers; i++ {
		worker := NewWorker(d.WorkerPool, i)
		worker.Start()
	}
	go d.dispatch()
}

func (d *Dispatcher) dispatch() {
	for {
		select {
		case job := <-JobQueue:
			go func(job Job) {
				jobChannel := <-d.WorkerPool
				jobChannel <- job
			}(job)
		}
	}
}
