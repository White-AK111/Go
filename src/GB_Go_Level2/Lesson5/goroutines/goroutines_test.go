package goroutines

import (
	"context"
	"fmt"
	"github.com/stretchr/testify/assert"
	"log"
	"testing"
	"time"
)

// Test calc to 1000 sheep by goroutines 10 times
func TestGoroutinePool_CalcSheep(t *testing.T) {

	var exp int64 = 1000
	pGo := NewGoroutinePool(1000)
	ctx, cancelFunc := context.WithTimeout(context.Background(), time.Second)
	defer cancelFunc()
	res, _ := pGo.CalcSheep(ctx)

	for i := 0; i < 10; i++ {
		assert.Equal(t, exp, res, "Incorrect result expected 1000 got %d", res)
	}
}

// Example for jump even sheep function by 10 goroutines
func ExampleGoroutinePool_CalcSheep() {
	pGo := NewGoroutinePool(10)
	ctx, cancelFunc := context.WithTimeout(context.Background(), time.Second)
	defer cancelFunc()
	res, err := pGo.CalcSheep(ctx)
	if err != nil {
		log.Fatalf("Error: %v \n", err)
	}
	fmt.Printf("Total sheep: %d \n", res)
	// Output:
	// Sheep with number 2 jumped over the fence
	// Sheep with number 4 jumped over the fence
	// Sheep with number 6 jumped over the fence
	// Sheep with number 8 jumped over the fence
	// Sheep with number 10 jumped over the fence
	// Total sheep: 10
}
