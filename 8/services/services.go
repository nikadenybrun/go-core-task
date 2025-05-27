package services

import (
	"sync"
)

type CustomWaitGroup struct {
	counter int
	cond    *sync.Cond
}

func NewWaitGroup() *CustomWaitGroup {
	return &CustomWaitGroup{cond: sync.NewCond(&sync.Mutex{})}
}

func (wg *CustomWaitGroup) Add(n int) {
	wg.cond.L.Lock()
	wg.counter += n
	wg.cond.L.Unlock()
}

func (wg *CustomWaitGroup) Done() {
	wg.cond.L.Lock()
	wg.counter--
	if wg.counter < 0 {
		panic("negative counter")
	}
	wg.cond.Signal()
	wg.cond.L.Unlock()
}

func (wg *CustomWaitGroup) Wait() {
	wg.cond.L.Lock()
	for wg.counter > 0 {
		wg.cond.Wait()
	}
	wg.cond.L.Unlock()
}
