// Package goroutines
//
// Task#1
// Напишите программу, которая запускает n потоков и дожидается завершения их всех
//
// Task#2
// Реализуйте функцию для разблокировки мьютекса с помощью defer
package goroutines

import (
	"context"
	"fmt"
	"sync"
	"sync/atomic"
)

// GoroutinePool struct for pool of goroutines
type GoroutinePool struct {
	pool    chan struct{}
	wg      sync.WaitGroup
	mu      sync.Mutex
	resCh   chan int64
	counter int64
}

// NewGoroutinePool function initializes new struct GoroutinePool with N goroutines in worker pool
func NewGoroutinePool(N int) *GoroutinePool {
	return &GoroutinePool{
		pool:    make(chan struct{}, N),
		wg:      sync.WaitGroup{},
		resCh:   make(chan int64, N/2),
		counter: 0,
	}
}

// CalcSheep method counting even sheep by goroutines, jump only even sheep
func (pGo *GoroutinePool) CalcSheep(ctx context.Context) (res int64, err error) {
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
				go sheepJump(pGo)
			}
		default:
			pGo.wg.Wait()
		}
	}
}

// sheepJump increment sheep by atomic function and mutex, print and add to result channel even sheep
func sheepJump(pGo *GoroutinePool) {
	pGo.mu.Lock()
	defer pGo.mu.Unlock()
	defer pGo.wg.Done()

	// imitation of several outputs of their function
	if (pGo.counter+1)%2 == 0 {
		atomic.AddInt64(&pGo.counter, 1)
		fmt.Printf("Sheep with number %d jumped over the fence\n", pGo.counter)
		pGo.resCh <- pGo.counter
		return
	} else {
		atomic.AddInt64(&pGo.counter, 1)
		return
	}
}
