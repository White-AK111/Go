package semaphores

import (
	"context"
	"fmt"
	"github.com/stretchr/testify/assert"
	"log"
	"testing"
	"time"
)

// Test calc do squaring into map by goroutines 10 times
func TestGoroutinePool_Inc(t *testing.T) {
	var exp int64 = 1000
	pGo := NewGoroutinePool(1000)
	ctx, cancelFunc := context.WithTimeout(context.Background(), time.Second)
	defer cancelFunc()
	res, _ := pGo.Squaring(ctx)

	for i := 0; i < 10; i++ {
		assert.Equal(t, exp, res, "Incorrect result expected 1000 got %d", res)
	}
}

// Example for do squaring into map function by 10000 goroutines
func ExampleGoroutinePool_Inc() {
	sem := NewGoroutinePool(1000)
	ctx, cancelFunc := context.WithTimeout(context.Background(), time.Second)
	defer cancelFunc()
	res, err := sem.Squaring(ctx)
	if err != nil {
		log.Fatalf("Error: %v \n", err)
	}
	fmt.Printf("%d \n", res)
	// Output:
	// 1000
}
