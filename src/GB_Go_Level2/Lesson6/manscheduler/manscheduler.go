// Package manscheduler
// Task#2
// Написать многопоточную программу, в которой будет использоваться явный вызов планировщика. Выполните трассировку программы
package manscheduler

import (
	"context"
	"runtime"
	"sync"
	"sync/atomic"
)

// GoroutinePool struct for pool of goroutines
type GoroutinePool struct {
	pool    chan struct{}
	wg      sync.WaitGroup
	mu      sync.RWMutex
	resCh   chan int64
	counter int64
}

// NewGoroutinePool function initializes new struct GoroutinePool with N goroutines in worker pool
func NewGoroutinePool(N int) *GoroutinePool {
	return &GoroutinePool{
		pool:    make(chan struct{}, N),
		wg:      sync.WaitGroup{},
		resCh:   make(chan int64, N),
		counter: 0,
	}
}

// Inc method increment a counter
func (pGo *GoroutinePool) Inc(ctx context.Context) (res int64, err error) {
	defer pGo.wg.Wait()

	for {
		select {
		case <-ctx.Done():
			{
				return pGo.counter, nil
			}
		case pGo.pool <- struct{}{}:
			{
				pGo.wg.Add(1)
				go incCounter(pGo)
			}
		default:
			pGo.wg.Wait()
		}
	}
}

// incCounter function increment counter by atomic function and mutex, add to result channel and execute manual scheduler for each 10th counter value
func incCounter(pGo *GoroutinePool) {
	pGo.mu.Lock()
	defer pGo.mu.Unlock()
	defer pGo.wg.Done()

	pGo.resCh <- pGo.counter
	atomic.AddInt64(&pGo.counter, 1)

	// execute manual scheduler for each 10th counter value
	if pGo.counter%10 == 0 {
		runtime.Gosched()
	}
}
