package astparser

import (
	"github.com/stretchr/testify/assert"
	"log"
	"testing"
)

// TestCountGoroutines test for counter go func by AST
func TestCountGoroutines(t *testing.T) {
	exp := 7

	n, err := CountGoroutines("astparser.go", "ExecGoroutines")
	if err != nil {
		log.Fatalf("Error: %s", err)
	}
	assert.Equal(t, n, exp, "Expected %d, got %d", exp, n)
}
