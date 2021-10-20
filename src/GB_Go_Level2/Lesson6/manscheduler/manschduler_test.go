package manscheduler

import (
	"context"
	"fmt"
	"github.com/stretchr/testify/assert"
	"log"
	"testing"
	"time"
)

// Test increment counter to 1000 by goroutines 10 times
func TestGoroutinePool_Inc(t *testing.T) {
	var exp int64 = 1000
	pGo := NewGoroutinePool(1000)
	ctx, cancelFunc := context.WithTimeout(context.Background(), time.Second)
	defer cancelFunc()
	res, _ := pGo.Inc(ctx)

	for i := 0; i < 10; i++ {
		assert.Equal(t, exp, res, "Incorrect result expected 1000 got %d", res)
	}
}

// Example for increment counter function by 10000 goroutines with manual execution scheduler
func ExampleGoroutinePool_Inc() {
	man := NewGoroutinePool(1000)
	ctxMan, cancelFuncMan := context.WithTimeout(context.Background(), time.Second)
	defer cancelFuncMan()
	res, err := man.Inc(ctxMan)
	if err != nil {
		log.Fatalf("Error: %v \n", err)
	}
	fmt.Printf("%d \n", res)
	// Output:
	// 1000
}
