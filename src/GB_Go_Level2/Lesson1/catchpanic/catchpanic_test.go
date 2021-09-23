package catchpanic

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

// Test get and recover panic
func TestGetAndRecoverPanic(t *testing.T) {
	var err error
	assert.NotEqual(t, err, GetAndRecoverPanic(), "error on panic can't be nil")
}
