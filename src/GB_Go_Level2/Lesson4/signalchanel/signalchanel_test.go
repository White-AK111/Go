package signalchanel

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

// Test send SIGNTERM signal
func TestTreatmentSignal(t *testing.T) {
	assert.NotNil(t, TreatmentSignal(), "Error can't be a nil")
}

// Example call ExampleTreatmentSignal function
func ExampleTreatmentSignal() {
	fmt.Println(TreatmentSignal())
	// Output: context canceled
}
