// goroutines - Task#1
// С помощью пула воркеров написать программу, которая запускает 1000 горутин, каждая из которых увеличивает число на 1.
// Дождаться завершения всех горутин и убедиться, что при каждом запуске программы итоговое число равно 1000.
package goroutines

import (
	"context"
	"fmt"
	"sync"
)

// GoroutineInc struct for increment value by goroutines
type GoroutineInc struct {
	pool  chan struct{}
	wg    sync.WaitGroup
	resCh chan int
}

// NewGoroutineAdd function initializes new struct GoroutineInc with N goroutines in worker pool
func NewGoroutineAdd(N int) *GoroutineInc {
	return &GoroutineInc{
		pool:  make(chan struct{}, N),
		wg:    sync.WaitGroup{},
		resCh: make(chan int, N),
	}
}

// AddWithGoroutines method increment values by goroutines
func (calc *GoroutineInc) AddWithGoroutines(ctx context.Context) (res int, err error) {
	defer calc.wg.Wait()
	calc.resCh <- 0

	for {
		select {
		case <-ctx.Done():
			{
				fmt.Println("Context done")
				return <-calc.resCh, nil
			}
		case calc.pool <- struct{}{}:
			{
				calc.wg.Add(1)
				go func() {
					currVal := <-calc.resCh
					calc.resCh <- currVal + 1
					calc.wg.Done()
				}()
			}
		default:
			calc.wg.Wait()
		}
	}
}
