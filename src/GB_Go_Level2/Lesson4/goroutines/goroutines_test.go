package goroutines

import (
	"context"
	"fmt"
	"github.com/stretchr/testify/assert"
	"log"
	"testing"
	"time"
)

// Test increment to 1000 by goroutines 10 times
func TestAddWithGoroutines(t *testing.T) {

	calc := NewGoroutineAdd(1000)
	ctx, cancelFunc := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancelFunc()
	res, _ := calc.AddWithGoroutines(ctx)

	for i := 0; i < 10; i++ {
		assert.Equal(t, 1000, res, "Incorrect result expected 1000 got %d", res)
	}
}

// Example increment to 1000 by goroutines
func ExampleAddWithGoroutines() {
	calc := NewGoroutineAdd(1000)
	ctx, cancelFunc := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancelFunc()
	res, err := calc.AddWithGoroutines(ctx)
	if err != nil {
		log.Fatalf("Error: %v \n", err)
	}
	fmt.Print(res)
	// Output:
	// Context done
	// 1000
}
