package selfstudy

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

// Test recover panic in goroutine
func TestDontPanic(t *testing.T) {
	assert.NotNil(t, DontPanic(), "error on panic can't be nil")
}
