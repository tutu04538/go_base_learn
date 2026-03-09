package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

type task struct {
	f func() error
}

type pool struct {
	activeWorker int64
	capacity     int64
	jobCh        chan *task
	sync.Mutex
}

func (p *pool) GetActiveWorker() int64 {
	return atomic.LoadInt64(&p.activeWorker)
}

func (p *pool) decActiveWorker() {
	atomic.AddInt64(&p.activeWorker, -1)
}

func (p *pool) addActiveWorker() {
	atomic.AddInt64(&p.activeWorker, 1)
}

func (p *pool) run() {
	p.addActiveWorker()
	go func() {
		defer p.decActiveWorker()
		for t := range p.jobCh {
			t.f()
		}
	}()
}

func (p *pool) addNewTask(t *task) {
	p.Lock()
	defer p.Unlock()

	if p.GetActiveWorker() < p.capacity {
		p.run()
	}
	p.jobCh <- t
}

func main() {
	p := pool{
		capacity: 3,
		jobCh:    make(chan *task, 10),
	}

	for range 20 {
		p.addNewTask(&task{func() error {
			fmt.Printf("a new task\n")
			return nil
		}})
	}

	time.Sleep(time.Second * 10)
}
