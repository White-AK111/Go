package catchpanic

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

// Test get and recover panic
func TestGetAndRecoverPanic(t *testing.T) {
	assert.NotNil(t, GetAndRecoverPanic(), "error on panic can't be nil")
}
