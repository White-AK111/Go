// Package race
// Task#3
// Смоделировать ситуацию “гонки”, и проверить программу на наличии “гонки”
package race

import (
	"context"
	"sync"
)

// GoroutinePool struct for pool of goroutines
type GoroutinePool struct {
	pool    chan struct{}
	wg      sync.WaitGroup
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

// IncRace method increment a counter
func (pGo *GoroutinePool) IncRace(ctx context.Context) (res int64, err error) {
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
				go incCounterRace(pGo)
			}
		default:
			pGo.wg.Wait()
		}
	}
}

// incCounter function increment counter by atomic function without mutex, add to result channel, get race conditions
func incCounterRace(pGo *GoroutinePool) {
	defer pGo.wg.Done()

	// Get race condition this
	pGo.resCh <- pGo.counter
	// And this
	pGo.counter++
}
