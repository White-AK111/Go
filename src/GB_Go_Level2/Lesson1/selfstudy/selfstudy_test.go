package selfstudy

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

// Test recover panic in goroutine
func TestDontPanic(t *testing.T) {
	var err error
	assert.NotEqual(t, err, DontPanic(), "error on panic can't be nil")
}
