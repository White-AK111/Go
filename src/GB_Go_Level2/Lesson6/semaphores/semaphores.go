// Package semaphores
// Task#1
// Написать программу, которая использует мьютекс для безопасного доступа к данным из нескольких потоков.
// Выполните трассировку программы
package semaphores

import (
	"context"
	"fmt"
	"sync"
	"sync/atomic"
)

// GoroutinePool struct for pool of goroutines
type GoroutinePool struct {
	pool    chan struct{}
	resMap  map[int64]int64
	wg      sync.WaitGroup
	mu      sync.RWMutex
	resCh   chan int64
	counter int64
}

// NewGoroutinePool function initializes new struct GoroutinePool with N goroutines in worker pool
func NewGoroutinePool(N int) *GoroutinePool {
	return &GoroutinePool{
		pool:    make(chan struct{}, N),
		resMap:  map[int64]int64{},
		wg:      sync.WaitGroup{},
		resCh:   make(chan int64, N),
		counter: 0,
	}
}

// Squaring method do squaring into map
func (pGo *GoroutinePool) Squaring(ctx context.Context) (res int64, err error) {
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
				go calcSquaring(pGo)
			}
		default:
			pGo.wg.Wait()
		}
	}
}

// PrintSquaring method print all values in result map
func (pGo *GoroutinePool) PrintSquaring() {
	pGo.mu.RLock()
	defer pGo.mu.RUnlock()
	for i, v := range pGo.resMap {
		fmt.Println(i, v)
	}
}

// calcSquaring function increment counter by atomic function and mutex, set result squaring current counter in map and add to result channel
func calcSquaring(pGo *GoroutinePool) {
	pGo.mu.Lock()
	defer pGo.mu.Unlock()
	defer pGo.wg.Done()

	res := pGo.counter * pGo.counter
	pGo.resMap[pGo.counter] = res
	pGo.resCh <- res
	atomic.AddInt64(&pGo.counter, 1)
}
