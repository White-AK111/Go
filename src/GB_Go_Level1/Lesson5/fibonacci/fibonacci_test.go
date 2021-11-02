package fibonacci

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// Benchmark for Fibonacci recurcive without map
func BenchmarkFibonacciRecurciveNoMap(b *testing.B) {
	n := 10
	nf := 0

	for i := 0; i < b.N; i++ {
		nf = FibonacciRecursive(n - 1)
	}

	fmt.Println(nf)
}

// Benchmark for Fibonacci recurcive with map
func BenchmarkFibonacciRecurciveWithMap(b *testing.B) {
	n := 10
	nf := 0
	fibonacciMap := make(map[int]int, 10)

	for i := 0; i < b.N; i++ {
		nf = FibonacciRecursiveWithMap(n-1, fibonacciMap)
	}

	fmt.Println(nf)
}

// Test Fibonacci recursive function without map
func TestFibonacciRecursiveNoMap(t *testing.T) {
	n := 10
	nf := 34

	assert.Equal(t, nf, FibonacciRecursive(n-1), "they should be equal")
}

// Test Fibonacci recursive function with map
func TestFibonacciRecursiveWithMap(t *testing.T) {
	n := 10
	nf := 34
	fibonacciMap := make(map[int]int, 10)

	assert.Equal(t, nf, FibonacciRecursiveWithMap(n-1, fibonacciMap), "they should be equal")
}
